// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tailscale.Tailnet
{
    public static class GetKey
    {
        public static Task<GetKeyResult> InvokeAsync(GetKeyArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetKeyResult>("tailscale:tailnet:getKey", args ?? new GetKeyArgs(), options.WithDefaults());

        public static Output<GetKeyResult> Invoke(GetKeyInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetKeyResult>("tailscale:tailnet:getKey", args ?? new GetKeyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetKeyArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        [Input("tailnet", required: true)]
        public string Tailnet { get; set; } = null!;

        public GetKeyArgs()
        {
        }
        public static new GetKeyArgs Empty => new GetKeyArgs();
    }

    public sealed class GetKeyInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        [Input("tailnet", required: true)]
        public Input<string> Tailnet { get; set; } = null!;

        public GetKeyInvokeArgs()
        {
        }
        public static new GetKeyInvokeArgs Empty => new GetKeyInvokeArgs();
    }


    [OutputType]
    public sealed class GetKeyResult
    {
        public readonly Outputs.AuthKeyRead Items;

        [OutputConstructor]
        private GetKeyResult(Outputs.AuthKeyRead items)
        {
            Items = items;
        }
    }
}
