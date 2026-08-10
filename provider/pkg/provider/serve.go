// Copyright 2022, Cloudy Sky Software LLC.

package provider

import (
	pulumiProvider "github.com/pulumi/pulumi/sdk/v3/go/pulumi/provider"

	"github.com/pulumi/pulumi/sdk/v3/go/common/util/cmdutil"
	rpc "github.com/pulumi/pulumi/sdk/v3/proto/go"
)

// Serve launches the gRPC server for the resource provider.
func Serve(providerName, version string, pulumiSchema, openapiDocBytes, metadataBytes []byte) {
	// Start gRPC service.
	err := pulumiProvider.Main(providerName, func(host *pulumiProvider.HostClient) (rpc.ResourceProviderServer, error) {
		return makeProvider(host, providerName, version, pulumiSchema, openapiDocBytes, metadataBytes)
	})
	if err != nil {
		cmdutil.ExitError(err.Error())
	}
}
