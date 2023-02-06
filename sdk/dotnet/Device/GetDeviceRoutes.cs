// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tailscale.Device
{
    public static class GetDeviceRoutes
    {
        public static Task<GetDeviceRoutesResult> InvokeAsync(GetDeviceRoutesArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDeviceRoutesResult>("tailscale:device:getDeviceRoutes", args ?? new GetDeviceRoutesArgs(), options.WithDefaults());

        public static Output<GetDeviceRoutesResult> Invoke(GetDeviceRoutesInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDeviceRoutesResult>("tailscale:device:getDeviceRoutes", args ?? new GetDeviceRoutesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDeviceRoutesArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetDeviceRoutesArgs()
        {
        }
        public static new GetDeviceRoutesArgs Empty => new GetDeviceRoutesArgs();
    }

    public sealed class GetDeviceRoutesInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetDeviceRoutesInvokeArgs()
        {
        }
        public static new GetDeviceRoutesInvokeArgs Empty => new GetDeviceRoutesInvokeArgs();
    }


    [OutputType]
    public sealed class GetDeviceRoutesResult
    {
        public readonly Outputs.DeviceRoutes Items;

        [OutputConstructor]
        private GetDeviceRoutesResult(Outputs.DeviceRoutes items)
        {
            Items = items;
        }
    }
}
