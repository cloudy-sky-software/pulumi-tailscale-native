// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

export function getDevice(args: GetDeviceArgs, opts?: pulumi.InvokeOptions): Promise<GetDeviceResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tailscale-native:device:getDevice", {
        "id": args.id,
    }, opts);
}

export interface GetDeviceArgs {
    id: string;
}

export interface GetDeviceResult {
    readonly addresses: string[];
    readonly advertisedRoutes: string[];
    readonly authorized: boolean;
    readonly blocksIncomingConnections: boolean;
    readonly clientConnectivity: outputs.device.ClientConnectivity;
    readonly clientVersion: string;
    readonly created: string;
    readonly enabledRoutes: string[];
    readonly expires: string;
    readonly hostname: string;
    readonly id: string;
    readonly isExternal: boolean;
    readonly keyExpiryDisabled: boolean;
    readonly lastSeen: string;
    readonly machineKey: string;
    readonly name: string;
    readonly nodeKey: string;
    readonly os: string;
    readonly updateAvailable: boolean;
    readonly user: string;
}
export function getDeviceOutput(args: GetDeviceOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetDeviceResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("tailscale-native:device:getDevice", {
        "id": args.id,
    }, opts);
}

export interface GetDeviceOutputArgs {
    id: pulumi.Input<string>;
}
