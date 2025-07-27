package provider

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/pulumi/pulumi/pkg/v3/resource/provider"

	"github.com/pulumi/pulumi/sdk/v3/go/common/util/logging"
	pulumirpc "github.com/pulumi/pulumi/sdk/v3/proto/go"

	fwCallback "github.com/cloudy-sky-software/pulumi-provider-framework/callback"
	fwRest "github.com/cloudy-sky-software/pulumi-provider-framework/rest"

	"golang.org/x/oauth2/clientcredentials"
)

type tailscaleProvider struct {
	fwCallback.UnimplementedProviderCallback

	name    string
	version string

	apiKey            *string
	oauthClientID     *string
	oauthClientSecret *string
}

var (
	handler  *fwRest.Provider
	callback fwCallback.ProviderCallback
)

func makeProvider(host *provider.HostClient, name, version string, pulumiSchemaBytes, openapiDocBytes, metadataBytes []byte) (pulumirpc.ResourceProviderServer, error) {
	p := &tailscaleProvider{
		name:    name,
		version: version,
	}

	callback = p
	rp, err := fwRest.MakeProvider(host, name, version, pulumiSchemaBytes, openapiDocBytes, metadataBytes, callback)

	handler = rp.(*fwRest.Provider)

	return rp, err
}

func (p *tailscaleProvider) GetAuthorizationHeader() string {
	if p.apiKey != nil {
		basicAuth := base64.URLEncoding.EncodeToString([]byte(*p.apiKey + ":"))
		return fmt.Sprintf("%s %s", basicAuthSchemePrefix, basicAuth)
	}

	if p.oauthClientID == nil || p.oauthClientSecret == nil {
		logging.Errorf("None of the supported credentials were configured. Calls will fail!")
		return ""
	}

	apiDoc := handler.GetOpenAPIDoc()
	for _, scheme := range apiDoc.Components.SecuritySchemes {
		if strings.ToLower(scheme.Value.Type) != "oauth2" {
			continue
		}

		if scheme.Value.Flows == nil || scheme.Value.Flows.ClientCredentials == nil {
			logging.Errorf("Invalid OAuth2 flow declared in the OpenAPI doc")
			continue
		}

		tokenURL := scheme.Value.Flows.ClientCredentials.TokenURL

		config := clientcredentials.Config{
			ClientID:     *p.oauthClientID,
			ClientSecret: *p.oauthClientSecret,
			TokenURL:     tokenURL,
		}
		token, err := config.Token(context.Background())
		if err != nil {
			logging.Errorf("Failed to fetch token using OAuth2 client credentials: %v", err)
			return ""
		}

		return fmt.Sprintf("%s %s", bearerAuthSchemePrefix, token.AccessToken)
	}

	logging.Errorf("Could not determine authorization header")

	return ""
}

// OnConfigure is called by the provider framework when Pulumi calls Configure on
// the resource provider server.
func (p *tailscaleProvider) OnConfigure(_ context.Context, req *pulumirpc.ConfigureRequest) (*pulumirpc.ConfigureResponse, error) {
	apiKey, ok := req.GetVariables()[fmt.Sprintf("%s:config:apiKey", p.name)]
	if !ok {
		// Check if it's set as an env var.
		envVarNames := handler.GetSchemaSpec().Provider.InputProperties["apiKey"].DefaultInfo.Environment
		for _, n := range envVarNames {
			v := os.Getenv(n)
			if v != "" {
				apiKey = v
			}
		}

		// Return an error if the API key is still empty.
		if apiKey == "" {
			logging.V(3).Info("Did not find an API key in the config")
		}
	}

	if apiKey != "" {
		logging.V(3).Info("Configuring Tailscale API key")
		p.apiKey = &apiKey
	}

	clientID, ok := req.GetVariables()[fmt.Sprintf("%s:config:clientId", p.name)]
	if !ok {
		// Check if it's set as an env var.
		envVarNames := handler.GetSchemaSpec().Provider.InputProperties["clientId"].DefaultInfo.Environment
		for _, n := range envVarNames {
			v := os.Getenv(n)
			if v != "" {
				clientID = v
			}
		}
	}

	if apiKey != "" && clientID != "" {
		return nil, errors.New("only one of apiKey or clientId and secret must be specified")
	}

	clientSecret, ok := req.GetVariables()[fmt.Sprintf("%s:config:clientSecret", p.name)]
	if !ok {
		// Check if it's set as an env var.
		envVarNames := handler.GetSchemaSpec().Provider.InputProperties["clientSecret"].DefaultInfo.Environment
		for _, n := range envVarNames {
			v := os.Getenv(n)
			if v != "" {
				clientSecret = v
			}
		}
	}

	logging.V(3).Info("Configuring Tailscale OAuth client credentials")
	p.oauthClientID = &clientID
	p.oauthClientSecret = &clientSecret

	return &pulumirpc.ConfigureResponse{
		AcceptSecrets: true,
	}, nil
}
