// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.TailscaleNative.Tailnet
{
    public static class GetNameServersConfig
    {
        public static Task<GetNameServersConfigResult> InvokeAsync(GetNameServersConfigArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetNameServersConfigResult>("tailscale-native:tailnet:getNameServersConfig", args ?? new GetNameServersConfigArgs(), options.WithDefaults());

        public static Output<GetNameServersConfigResult> Invoke(GetNameServersConfigInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetNameServersConfigResult>("tailscale-native:tailnet:getNameServersConfig", args ?? new GetNameServersConfigInvokeArgs(), options.WithDefaults());

        public static Output<GetNameServersConfigResult> Invoke(GetNameServersConfigInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetNameServersConfigResult>("tailscale-native:tailnet:getNameServersConfig", args ?? new GetNameServersConfigInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNameServersConfigArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// For paid plans, your domain is your tailnet. For solo plans, the tailnet is the email you signed up with. So `alice@gmail.com` has the tailnet `alice@gmail.com` since `@gmail.com` is a shared email host. Alternatively, you can specify the value "-" to refer to the default tailnet of the authenticated user making the API call.
        /// </summary>
        [Input("tailnet", required: true)]
        public string Tailnet { get; set; } = null!;

        public GetNameServersConfigArgs()
        {
        }
        public static new GetNameServersConfigArgs Empty => new GetNameServersConfigArgs();
    }

    public sealed class GetNameServersConfigInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// For paid plans, your domain is your tailnet. For solo plans, the tailnet is the email you signed up with. So `alice@gmail.com` has the tailnet `alice@gmail.com` since `@gmail.com` is a shared email host. Alternatively, you can specify the value "-" to refer to the default tailnet of the authenticated user making the API call.
        /// </summary>
        [Input("tailnet", required: true)]
        public Input<string> Tailnet { get; set; } = null!;

        public GetNameServersConfigInvokeArgs()
        {
        }
        public static new GetNameServersConfigInvokeArgs Empty => new GetNameServersConfigInvokeArgs();
    }


    [OutputType]
    public sealed class GetNameServersConfigResult
    {
        public readonly ImmutableArray<string> Dns;
        public readonly bool MagicDNS;

        [OutputConstructor]
        private GetNameServersConfigResult(
            ImmutableArray<string> dns,

            bool magicDNS)
        {
            Dns = dns;
            MagicDNS = magicDNS;
        }
    }
}
