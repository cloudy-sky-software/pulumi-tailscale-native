// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.TailscaleNative.Tailnet
{
    public static class GetKey
    {
        public static Task<GetKeyResult> InvokeAsync(GetKeyArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetKeyResult>("tailscale-native:tailnet:getKey", args ?? new GetKeyArgs(), options.WithDefaults());

        public static Output<GetKeyResult> Invoke(GetKeyInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetKeyResult>("tailscale-native:tailnet:getKey", args ?? new GetKeyInvokeArgs(), options.WithDefaults());

        public static Output<GetKeyResult> Invoke(GetKeyInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetKeyResult>("tailscale-native:tailnet:getKey", args ?? new GetKeyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetKeyArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        /// <summary>
        /// For paid plans, your domain is your tailnet. For solo plans, the tailnet is the email you signed up with. So `alice@gmail.com` has the tailnet `alice@gmail.com` since `@gmail.com` is a shared email host. Alternatively, you can specify the value "-" to refer to the default tailnet of the authenticated user making the API call.
        /// </summary>
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

        /// <summary>
        /// For paid plans, your domain is your tailnet. For solo plans, the tailnet is the email you signed up with. So `alice@gmail.com` has the tailnet `alice@gmail.com` since `@gmail.com` is a shared email host. Alternatively, you can specify the value "-" to refer to the default tailnet of the authenticated user making the API call.
        /// </summary>
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
        public readonly string? Created;
        public readonly string Expires;
        public readonly string Key;

        [OutputConstructor]
        private GetKeyResult(
            string? created,

            string expires,

            string key)
        {
            Created = created;
            Expires = expires;
            Key = key;
        }
    }
}
