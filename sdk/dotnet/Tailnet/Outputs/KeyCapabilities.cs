// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tailscale.Tailnet.Outputs
{

    [OutputType]
    public sealed class KeyCapabilities
    {
        public readonly Outputs.DeviceKeyCapabilities Devices;

        [OutputConstructor]
        private KeyCapabilities(Outputs.DeviceKeyCapabilities devices)
        {
            Devices = devices;
        }
    }
}
