// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

export function getKey(args: GetKeyArgs, opts?: pulumi.InvokeOptions): Promise<GetKeyResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tailscale-native:tailnet:getKey", {
        "id": args.id,
        "tailnet": args.tailnet,
    }, opts);
}

export interface GetKeyArgs {
    id: string;
    tailnet: string;
}

export interface GetKeyResult {
    readonly items: outputs.tailnet.AuthKey;
}
export function getKeyOutput(args: GetKeyOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetKeyResult> {
    return pulumi.output(args).apply((a: any) => getKey(a, opts))
}

export interface GetKeyOutputArgs {
    id: pulumi.Input<string>;
    tailnet: pulumi.Input<string>;
}
