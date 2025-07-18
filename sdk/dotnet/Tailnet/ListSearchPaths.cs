// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.TailscaleNative.Tailnet
{
    public static class ListSearchPaths
    {
        public static Task<ListSearchPathsResult> InvokeAsync(ListSearchPathsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListSearchPathsResult>("tailscale-native:tailnet:listSearchPaths", args ?? new ListSearchPathsArgs(), options.WithDefaults());

        public static Output<ListSearchPathsResult> Invoke(ListSearchPathsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListSearchPathsResult>("tailscale-native:tailnet:listSearchPaths", args ?? new ListSearchPathsInvokeArgs(), options.WithDefaults());

        public static Output<ListSearchPathsResult> Invoke(ListSearchPathsInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<ListSearchPathsResult>("tailscale-native:tailnet:listSearchPaths", args ?? new ListSearchPathsInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListSearchPathsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// For paid plans, your domain is your tailnet. For solo plans, the tailnet is the email you signed up with. So `alice@gmail.com` has the tailnet `alice@gmail.com` since `@gmail.com` is a shared email host. Alternatively, you can specify the value "-" to refer to the default tailnet of the authenticated user making the API call.
        /// </summary>
        [Input("tailnet", required: true)]
        public string Tailnet { get; set; } = null!;

        public ListSearchPathsArgs()
        {
        }
        public static new ListSearchPathsArgs Empty => new ListSearchPathsArgs();
    }

    public sealed class ListSearchPathsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// For paid plans, your domain is your tailnet. For solo plans, the tailnet is the email you signed up with. So `alice@gmail.com` has the tailnet `alice@gmail.com` since `@gmail.com` is a shared email host. Alternatively, you can specify the value "-" to refer to the default tailnet of the authenticated user making the API call.
        /// </summary>
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
        public readonly ImmutableArray<string> SearchPaths;

        [OutputConstructor]
        private ListSearchPathsResult(ImmutableArray<string> searchPaths)
        {
            SearchPaths = searchPaths;
        }
    }
}
