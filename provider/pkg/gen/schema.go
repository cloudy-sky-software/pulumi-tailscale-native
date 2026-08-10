// Copyright 2022, Cloudy Sky Software LLC.

package gen

import (
	"bytes"
	"encoding/json"

	"github.com/getkin/kin-openapi/openapi3"

	dotnetgen "github.com/pulumi/pulumi-dotnet/pulumi-language-dotnet/v3/codegen"
	gogen "github.com/pulumi/pulumi/pkg/v3/codegen/go"
	nodejsgen "github.com/pulumi/pulumi/pkg/v3/codegen/nodejs"
	pythongen "github.com/pulumi/pulumi/pkg/v3/codegen/python"
	pschema "github.com/pulumi/pulumi/pkg/v3/codegen/schema"

	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"

	openapigen "github.com/cloudy-sky-software/pulschema/pkg"

	"github.com/cloudy-sky-software/pulumi-tailscale-native/provider/pkg/gen/examples"
)

const packageName = "tailscale-native"

// PulumiSchema will generate a Pulumi schema for the given k8s schema.
func PulumiSchema(openapiDoc openapi3.T) (pschema.PackageSpec, openapigen.ProviderMetadata, openapi3.T) {
	pkg := pschema.PackageSpec{
		Name:        packageName,
		Description: "A native Pulumi package for creating and managing Tailscale resources.",
		DisplayName: "Tailscale Native",
		License:     "Apache-2.0",
		Keywords: []string{
			"pulumi",
			packageName,
			"category/cloud",
			"kind/native",
		},
		Homepage:   "https://cloudysky.software",
		Publisher:  "Cloudy Sky Software",
		Repository: "https://github.com/cloudy-sky-software/pulumi-tailscale-native",

		Config: pschema.ConfigSpec{
			Variables: map[string]pschema.PropertySpec{
				"apiKey": {
					Description: "The API key",
					TypeSpec:    pschema.TypeSpec{Type: pschema.StringType.String()},
					Language: map[string]pschema.RawMessage{
						//nolint: goconst
						"csharp": rawMessage(map[string]interface{}{
							"name": "ApiKey",
						}),
					},
					Secret: true,
				},
				"clientId": {
					Description: "The OAuth client ID",
					TypeSpec:    pschema.TypeSpec{Type: pschema.StringType.String()},
					Language: map[string]pschema.RawMessage{
						//nolint: goconst
						"csharp": rawMessage(map[string]interface{}{
							"name": "ClientId",
						}),
					},
					Secret: false,
				},
				"clientSecret": {
					Description: "The OAuth client secret",
					TypeSpec:    pschema.TypeSpec{Type: pschema.StringType.String()},
					Language: map[string]pschema.RawMessage{
						//nolint: goconst
						"csharp": rawMessage(map[string]interface{}{
							"name": "ClientSecret",
						}),
					},
					Secret: true,
				},
			},
		},

		Provider: &pschema.ResourceSpec{
			ObjectTypeSpec: pschema.ObjectTypeSpec{
				Description: "The provider type for the Tailscale package.",
				Type:        "object",
			},
			InputProperties: map[string]pschema.PropertySpec{
				"apiKey": {
					DefaultInfo: &pschema.DefaultSpec{
						Environment: []string{
							"TAILSCALE_NATIVE_APIKEY",
							"TAILSCALE_APIKEY",
						},
					},
					Description: "The Tailscale API key.",
					TypeSpec:    pschema.TypeSpec{Type: pschema.StringType.String()},
					Language: map[string]pschema.RawMessage{
						"csharp": rawMessage(map[string]interface{}{
							"name": "ApiKey",
						}),
					},
					Secret: true,
				},
				"clientId": {
					DefaultInfo: &pschema.DefaultSpec{
						Environment: []string{
							"TAILSCALE_NATIVE_CLIENT_ID",
							"TAILSCALE_CLIENT_ID",
						},
					},
					Description: "The Tailscale OAuth client ID.",
					TypeSpec:    pschema.TypeSpec{Type: "string"},
					Language: map[string]pschema.RawMessage{
						"csharp": rawMessage(map[string]interface{}{
							"name": "ClientId",
						}),
					},
					Secret: false,
				},
				"clientSecret": {
					DefaultInfo: &pschema.DefaultSpec{
						Environment: []string{
							"TAILSCALE_NATIVE_CLIENT_SECRET",
							"TAILSCALE_CLIENT_SECRET",
						},
					},
					Description: "The Tailscale OAuth client secret.",
					TypeSpec:    pschema.TypeSpec{Type: "string"},
					Language: map[string]pschema.RawMessage{
						"csharp": rawMessage(map[string]interface{}{
							"name": "ClientSecret",
						}),
					},
					Secret: true,
				},
			},
		},

		PluginDownloadURL: "github://api.github.com/cloudy-sky-software/pulumi-tailscale-native",
		Types:             map[string]pschema.ComplexTypeSpec{},
		Resources:         map[string]pschema.ResourceSpec{},
		Functions:         map[string]pschema.FunctionSpec{},
		Language:          map[string]pschema.RawMessage{},
	}

	csharpNamespaces := map[string]string{
		"tailscale-native": "TailscaleNative",
		// TODO: Is this needed?
		"": "Provider",
	}

	openAPICtx := &openapigen.OpenAPIContext{
		Doc:                               openapiDoc,
		Pkg:                               &pkg,
		ExcludedPaths:                     []string{},
		OperationIDsHaveTypeSpecNamespace: true,
		TypeSpecNamespaceSeparator:        "_",
		AllowedPluralResources:            []string{"Routes", "Tags", "NameServers", "SearchPaths", "DNSPreferences"},
	}

	providerMetadata, updatedOpenAPIDoc, err := openAPICtx.GatherResourcesFromAPI(csharpNamespaces)
	if err != nil {
		contract.Failf("generating resources from OpenAPI spec: %v", err)
	}

	// Add examples to resources
	for k, v := range examples.ResourceExample {
		if r, ok := pkg.Resources[k]; ok {
			r.Description += "\n\n" + v
			pkg.Resources[k] = r
		}
	}

	//nolint: goconst
	pkg.Language["csharp"] = rawMessage(dotnetgen.CSharpPackageInfo{
		Namespaces: csharpNamespaces,
		PackageReferences: map[string]string{
			"Pulumi": "3.*",
		},
		RootNamespace: "Pulumi",
	})

	pkg.Language["go"] = rawMessage(gogen.GoPackageInfo{
		ImportBasePath: "github.com/cloudy-sky-software/pulumi-tailscale-native/sdk/go/tailscale",
		ModuleToPackage: map[string]string{
			"github.com/cloudy-sky-software/pulumi-tailscale-native/sdk/go/tailscale": "tailscale",
		},
	})
	pkg.Language["nodejs"] = rawMessage(nodejsgen.NodePackageInfo{
		PackageName: "@cloudyskysoftware/pulumi-tailscale-native",
	})
	pkg.Language["python"] = rawMessage(pythongen.PackageInfo{
		PackageName: "pulumi_tailscale_native",
		Requires: map[string]string{
			"pulumi": ">=3.0.0,<4.0.0",
		},
		PyProject: struct {
			Enabled bool `json:"enabled,omitempty"`
		}{Enabled: true},
	})

	metadata := openapigen.ProviderMetadata{
		ResourceCRUDMap:  providerMetadata.ResourceCRUDMap,
		AutoNameMap:      providerMetadata.AutoNameMap,
		SDKToAPINameMap:  providerMetadata.SDKToAPINameMap,
		APIToSDKNameMap:  providerMetadata.APIToSDKNameMap,
		PathParamNameMap: providerMetadata.PathParamNameMap,
	}
	return pkg, metadata, updatedOpenAPIDoc
}

func rawMessage(v interface{}) pschema.RawMessage {
	var out bytes.Buffer
	encoder := json.NewEncoder(&out)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v)
	contract.Assertf(err == nil, "Failed to convert input to raw JSON bytes")
	return out.Bytes()
}
