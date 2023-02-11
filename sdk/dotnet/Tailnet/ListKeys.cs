// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tailscale.Tailnet
{
    public static class ListKeys
    {
        public static Task<ListKeysResult> InvokeAsync(ListKeysArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListKeysResult>("tailscale:tailnet:listKeys", args ?? new ListKeysArgs(), options.WithDefaults());

        public static Output<ListKeysResult> Invoke(ListKeysInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListKeysResult>("tailscale:tailnet:listKeys", args ?? new ListKeysInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListKeysArgs : global::Pulumi.InvokeArgs
    {
        [Input("tailnet", required: true)]
        public string Tailnet { get; set; } = null!;

        public ListKeysArgs()
        {
        }
        public static new ListKeysArgs Empty => new ListKeysArgs();
    }

    public sealed class ListKeysInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("tailnet", required: true)]
        public Input<string> Tailnet { get; set; } = null!;

        public ListKeysInvokeArgs()
        {
        }
        public static new ListKeysInvokeArgs Empty => new ListKeysInvokeArgs();
    }


    [OutputType]
    public sealed class ListKeysResult
    {
        public readonly ImmutableArray<Outputs.AuthKeyReadItem> Items;

        [OutputConstructor]
        private ListKeysResult(ImmutableArray<Outputs.AuthKeyReadItem> items)
        {
            Items = items;
        }
    }
}
