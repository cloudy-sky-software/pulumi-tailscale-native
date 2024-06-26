// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.TailscaleNative.Tailnet
{
    public static class GetDNSPreference
    {
        public static Task<GetDNSPreferenceResult> InvokeAsync(GetDNSPreferenceArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDNSPreferenceResult>("tailscale-native:tailnet:getDNSPreference", args ?? new GetDNSPreferenceArgs(), options.WithDefaults());

        public static Output<GetDNSPreferenceResult> Invoke(GetDNSPreferenceInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDNSPreferenceResult>("tailscale-native:tailnet:getDNSPreference", args ?? new GetDNSPreferenceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDNSPreferenceArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// For paid plans, your domain is your tailnet. For solo plans, the tailnet is the email you signed up with. So `alice@gmail.com` has the tailnet `alice@gmail.com` since `@gmail.com` is a shared email host. Alternatively, you can specify the value "-" to refer to the default tailnet of the authenticated user making the API call.
        /// </summary>
        [Input("tailnet", required: true)]
        public string Tailnet { get; set; } = null!;

        public GetDNSPreferenceArgs()
        {
        }
        public static new GetDNSPreferenceArgs Empty => new GetDNSPreferenceArgs();
    }

    public sealed class GetDNSPreferenceInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// For paid plans, your domain is your tailnet. For solo plans, the tailnet is the email you signed up with. So `alice@gmail.com` has the tailnet `alice@gmail.com` since `@gmail.com` is a shared email host. Alternatively, you can specify the value "-" to refer to the default tailnet of the authenticated user making the API call.
        /// </summary>
        [Input("tailnet", required: true)]
        public Input<string> Tailnet { get; set; } = null!;

        public GetDNSPreferenceInvokeArgs()
        {
        }
        public static new GetDNSPreferenceInvokeArgs Empty => new GetDNSPreferenceInvokeArgs();
    }


    [OutputType]
    public sealed class GetDNSPreferenceResult
    {
        public readonly Outputs.NameServersPreference Items;

        [OutputConstructor]
        private GetDNSPreferenceResult(Outputs.NameServersPreference items)
        {
            Items = items;
        }
    }
}
