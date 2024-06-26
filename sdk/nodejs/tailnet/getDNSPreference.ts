// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

export function getDNSPreference(args: GetDNSPreferenceArgs, opts?: pulumi.InvokeOptions): Promise<GetDNSPreferenceResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tailscale-native:tailnet:getDNSPreference", {
        "tailnet": args.tailnet,
    }, opts);
}

export interface GetDNSPreferenceArgs {
    /**
     * For paid plans, your domain is your tailnet. For solo plans, the tailnet is the email you signed up with. So `alice@gmail.com` has the tailnet `alice@gmail.com` since `@gmail.com` is a shared email host. Alternatively, you can specify the value "-" to refer to the default tailnet of the authenticated user making the API call.
     */
    tailnet: string;
}

export interface GetDNSPreferenceResult {
    readonly items: outputs.tailnet.NameServersPreference;
}
export function getDNSPreferenceOutput(args: GetDNSPreferenceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDNSPreferenceResult> {
    return pulumi.output(args).apply((a: any) => getDNSPreference(a, opts))
}

export interface GetDNSPreferenceOutputArgs {
    /**
     * For paid plans, your domain is your tailnet. For solo plans, the tailnet is the email you signed up with. So `alice@gmail.com` has the tailnet `alice@gmail.com` since `@gmail.com` is a shared email host. Alternatively, you can specify the value "-" to refer to the default tailnet of the authenticated user making the API call.
     */
    tailnet: pulumi.Input<string>;
}
