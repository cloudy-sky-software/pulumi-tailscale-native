// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class Routes extends pulumi.CustomResource {
    /**
     * Get an existing Routes resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Routes {
        return new Routes(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tailscale-native:device:Routes';

    /**
     * Returns true if the given object is an instance of Routes.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Routes {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Routes.__pulumiType;
    }

    public /*out*/ readonly advertisedRoutes!: pulumi.Output<string[]>;
    public /*out*/ readonly enabledRoutes!: pulumi.Output<string[]>;
    public readonly routes!: pulumi.Output<string[]>;

    /**
     * Create a Routes resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RoutesArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.routes === undefined) && !opts.urn) {
                throw new Error("Missing required property 'routes'");
            }
            resourceInputs["id"] = args ? args.id : undefined;
            resourceInputs["routes"] = args ? args.routes : undefined;
            resourceInputs["advertisedRoutes"] = undefined /*out*/;
            resourceInputs["enabledRoutes"] = undefined /*out*/;
        } else {
            resourceInputs["advertisedRoutes"] = undefined /*out*/;
            resourceInputs["enabledRoutes"] = undefined /*out*/;
            resourceInputs["routes"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Routes.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Routes resource.
 */
export interface RoutesArgs {
    id?: pulumi.Input<string>;
    routes: pulumi.Input<pulumi.Input<string>[]>;
}
