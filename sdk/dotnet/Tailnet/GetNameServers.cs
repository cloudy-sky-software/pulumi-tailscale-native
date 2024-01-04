// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.TailscaleNative.Tailnet
{
    public static class GetNameServers
    {
        public static Task<GetNameServersResult> InvokeAsync(GetNameServersArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetNameServersResult>("tailscale-native:tailnet:getNameServers", args ?? new GetNameServersArgs(), options.WithDefaults());

        public static Output<GetNameServersResult> Invoke(GetNameServersInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetNameServersResult>("tailscale-native:tailnet:getNameServers", args ?? new GetNameServersInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNameServersArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// For paid plans, your domain is your tailnet. For solo plans, the tailnet is the email you signed up with. So `alice@gmail.com` has the tailnet `alice@gmail.com` since `@gmail.com` is a shared email host. Alternatively, you can specify the value "-" to refer to the default tailnet of the authenticated user making the API call.
        /// </summary>
        [Input("tailnet", required: true)]
        public string Tailnet { get; set; } = null!;

        public GetNameServersArgs()
        {
        }
        public static new GetNameServersArgs Empty => new GetNameServersArgs();
    }

    public sealed class GetNameServersInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// For paid plans, your domain is your tailnet. For solo plans, the tailnet is the email you signed up with. So `alice@gmail.com` has the tailnet `alice@gmail.com` since `@gmail.com` is a shared email host. Alternatively, you can specify the value "-" to refer to the default tailnet of the authenticated user making the API call.
        /// </summary>
        [Input("tailnet", required: true)]
        public Input<string> Tailnet { get; set; } = null!;

        public GetNameServersInvokeArgs()
        {
        }
        public static new GetNameServersInvokeArgs Empty => new GetNameServersInvokeArgs();
    }


    [OutputType]
    public sealed class GetNameServersResult
    {
        public readonly Outputs.NameServers Items;

        [OutputConstructor]
        private GetNameServersResult(Outputs.NameServers items)
        {
            Items = items;
        }
    }
}
