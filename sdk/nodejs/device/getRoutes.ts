// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

export function getRoutes(args: GetRoutesArgs, opts?: pulumi.InvokeOptions): Promise<GetRoutesResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tailscale:device:getRoutes", {
        "id": args.id,
    }, opts);
}

export interface GetRoutesArgs {
    id: string;
}

export interface GetRoutesResult {
    readonly items: outputs.device.DeviceRoutes;
}
export function getRoutesOutput(args: GetRoutesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRoutesResult> {
    return pulumi.output(args).apply((a: any) => getRoutes(a, opts))
}

export interface GetRoutesOutputArgs {
    id: pulumi.Input<string>;
}
