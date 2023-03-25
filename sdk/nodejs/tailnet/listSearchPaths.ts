// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

export function listSearchPaths(args: ListSearchPathsArgs, opts?: pulumi.InvokeOptions): Promise<ListSearchPathsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tailscale-native:tailnet:listSearchPaths", {
        "tailnet": args.tailnet,
    }, opts);
}

export interface ListSearchPathsArgs {
    tailnet: string;
}

export interface ListSearchPathsResult {
    readonly items: outputs.tailnet.DnsSearchPaths;
}
export function listSearchPathsOutput(args: ListSearchPathsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<ListSearchPathsResult> {
    return pulumi.output(args).apply((a: any) => listSearchPaths(a, opts))
}

export interface ListSearchPathsOutputArgs {
    tailnet: pulumi.Input<string>;
}
