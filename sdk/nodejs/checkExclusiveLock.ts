// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manages a Exclusive Lock Check.
 *
 * Adding an exclusive lock will only allow a single stage to utilize this resource at a time. If multiple stages are waiting on the lock, only the latest will run. All others will be canceled.
 *
 * ## Example Usage
 *
 * ## Import
 *
 * Importing this resource is not supported.
 */
export class CheckExclusiveLock extends pulumi.CustomResource {
    /**
     * Get an existing CheckExclusiveLock resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CheckExclusiveLockState, opts?: pulumi.CustomResourceOptions): CheckExclusiveLock {
        return new CheckExclusiveLock(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuredevops:index/checkExclusiveLock:CheckExclusiveLock';

    /**
     * Returns true if the given object is an instance of CheckExclusiveLock.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CheckExclusiveLock {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CheckExclusiveLock.__pulumiType;
    }

    /**
     * The project ID. Changing this forces a new Exclusive Lock Check to be created.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * The ID of the resource being protected by the check. Changing this forces a new Exclusive Lock to be created.
     */
    public readonly targetResourceId!: pulumi.Output<string>;
    /**
     * The type of resource being protected by the check. Valid values: `endpoint`, `environment`, `queue`, `repository`, `securefile`, `variablegroup`. Changing this forces a new Exclusive Lock to be created.
     */
    public readonly targetResourceType!: pulumi.Output<string>;
    /**
     * The timeout in minutes for the exclusive lock. Defaults to `43200`.
     */
    public readonly timeout!: pulumi.Output<number | undefined>;

    /**
     * Create a CheckExclusiveLock resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CheckExclusiveLockArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CheckExclusiveLockArgs | CheckExclusiveLockState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CheckExclusiveLockState | undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["targetResourceId"] = state ? state.targetResourceId : undefined;
            resourceInputs["targetResourceType"] = state ? state.targetResourceType : undefined;
            resourceInputs["timeout"] = state ? state.timeout : undefined;
        } else {
            const args = argsOrState as CheckExclusiveLockArgs | undefined;
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            if ((!args || args.targetResourceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'targetResourceId'");
            }
            if ((!args || args.targetResourceType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'targetResourceType'");
            }
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["targetResourceId"] = args ? args.targetResourceId : undefined;
            resourceInputs["targetResourceType"] = args ? args.targetResourceType : undefined;
            resourceInputs["timeout"] = args ? args.timeout : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CheckExclusiveLock.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CheckExclusiveLock resources.
 */
export interface CheckExclusiveLockState {
    /**
     * The project ID. Changing this forces a new Exclusive Lock Check to be created.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The ID of the resource being protected by the check. Changing this forces a new Exclusive Lock to be created.
     */
    targetResourceId?: pulumi.Input<string>;
    /**
     * The type of resource being protected by the check. Valid values: `endpoint`, `environment`, `queue`, `repository`, `securefile`, `variablegroup`. Changing this forces a new Exclusive Lock to be created.
     */
    targetResourceType?: pulumi.Input<string>;
    /**
     * The timeout in minutes for the exclusive lock. Defaults to `43200`.
     */
    timeout?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a CheckExclusiveLock resource.
 */
export interface CheckExclusiveLockArgs {
    /**
     * The project ID. Changing this forces a new Exclusive Lock Check to be created.
     */
    projectId: pulumi.Input<string>;
    /**
     * The ID of the resource being protected by the check. Changing this forces a new Exclusive Lock to be created.
     */
    targetResourceId: pulumi.Input<string>;
    /**
     * The type of resource being protected by the check. Valid values: `endpoint`, `environment`, `queue`, `repository`, `securefile`, `variablegroup`. Changing this forces a new Exclusive Lock to be created.
     */
    targetResourceType: pulumi.Input<string>;
    /**
     * The timeout in minutes for the exclusive lock. Defaults to `43200`.
     */
    timeout?: pulumi.Input<number>;
}
