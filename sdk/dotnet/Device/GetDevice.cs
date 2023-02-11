// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tailscale.Device
{
    public static class GetDevice
    {
        public static Task<GetDeviceResult> InvokeAsync(GetDeviceArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDeviceResult>("tailscale:device:getDevice", args ?? new GetDeviceArgs(), options.WithDefaults());

        public static Output<GetDeviceResult> Invoke(GetDeviceInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDeviceResult>("tailscale:device:getDevice", args ?? new GetDeviceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDeviceArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetDeviceArgs()
        {
        }
        public static new GetDeviceArgs Empty => new GetDeviceArgs();
    }

    public sealed class GetDeviceInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetDeviceInvokeArgs()
        {
        }
        public static new GetDeviceInvokeArgs Empty => new GetDeviceInvokeArgs();
    }


    [OutputType]
    public sealed class GetDeviceResult
    {
        public readonly Outputs.Device Items;

        [OutputConstructor]
        private GetDeviceResult(Outputs.Device items)
        {
            Items = items;
        }
    }
}
