// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

export function getRoutesConfig(args: GetRoutesConfigArgs, opts?: pulumi.InvokeOptions): Promise<GetRoutesConfigResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tailscale-native:device:getRoutesConfig", {
        "id": args.id,
    }, opts);
}

export interface GetRoutesConfigArgs {
    id: string;
}

export interface GetRoutesConfigResult {
    readonly items: outputs.device.DeviceRoutes;
}
export function getRoutesConfigOutput(args: GetRoutesConfigOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRoutesConfigResult> {
    return pulumi.output(args).apply((a: any) => getRoutesConfig(a, opts))
}

export interface GetRoutesConfigOutputArgs {
    id: pulumi.Input<string>;
}