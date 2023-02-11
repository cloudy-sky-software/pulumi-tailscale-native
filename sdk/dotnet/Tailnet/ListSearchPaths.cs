// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tailscale.Tailnet
{
    public static class ListSearchPaths
    {
        public static Task<ListSearchPathsResult> InvokeAsync(ListSearchPathsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListSearchPathsResult>("tailscale:tailnet:listSearchPaths", args ?? new ListSearchPathsArgs(), options.WithDefaults());

        public static Output<ListSearchPathsResult> Invoke(ListSearchPathsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListSearchPathsResult>("tailscale:tailnet:listSearchPaths", args ?? new ListSearchPathsInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListSearchPathsArgs : global::Pulumi.InvokeArgs
    {
        [Input("tailnet", required: true)]
        public string Tailnet { get; set; } = null!;

        public ListSearchPathsArgs()
        {
        }
        public static new ListSearchPathsArgs Empty => new ListSearchPathsArgs();
    }

    public sealed class ListSearchPathsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("tailnet", required: true)]
        public Input<string> Tailnet { get; set; } = null!;

        public ListSearchPathsInvokeArgs()
        {
        }
        public static new ListSearchPathsInvokeArgs Empty => new ListSearchPathsInvokeArgs();
    }


    [OutputType]
    public sealed class ListSearchPathsResult
    {
        public readonly Outputs.DnsSearchPaths Items;

        [OutputConstructor]
        private ListSearchPathsResult(Outputs.DnsSearchPaths items)
        {
            Items = items;
        }
    }
}
