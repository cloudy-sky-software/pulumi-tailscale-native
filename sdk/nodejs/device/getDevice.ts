// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

export function getDevice(args: GetDeviceArgs, opts?: pulumi.InvokeOptions): Promise<GetDeviceResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tailscale:device:getDevice", {
        "id": args.id,
    }, opts);
}

export interface GetDeviceArgs {
    id: string;
}

export interface GetDeviceResult {
    readonly items: outputs.device.Device;
}
export function getDeviceOutput(args: GetDeviceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDeviceResult> {
    return pulumi.output(args).apply((a: any) => getDevice(a, opts))
}

export interface GetDeviceOutputArgs {
    id: pulumi.Input<string>;
}
