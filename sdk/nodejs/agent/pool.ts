// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages an agent pool within Azure DevOps.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const pool = new azuredevops.Pool("pool", {
 *     autoProvision: false,
 * });
 * ```
 * ## Relevant Links
 *
 * * [Azure DevOps Service REST API 5.1 - Agent Pools](https://docs.microsoft.com/en-us/rest/api/azure/devops/distributedtask/pools?view=azure-devops-rest-5.1)
 *
 * @deprecated azuredevops.agent.Pool has been deprecated in favor of azuredevops.Pool
 */
export class Pool extends pulumi.CustomResource {
    /**
     * Get an existing Pool resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PoolState, opts?: pulumi.CustomResourceOptions): Pool {
        pulumi.log.warn("Pool is deprecated: azuredevops.agent.Pool has been deprecated in favor of azuredevops.Pool")
        return new Pool(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuredevops:Agent/pool:Pool';

    /**
     * Returns true if the given object is an instance of Pool.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Pool {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Pool.__pulumiType;
    }

    /**
     * Specifies whether or not a queue should be automatically provisioned for each project collection. Defaults to `false`.
     */
    public readonly autoProvision!: pulumi.Output<boolean | undefined>;
    /**
     * The name of the agent pool.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies whether the agent pool type is Automation or Deployment.  Defaults to `automation`.
     */
    public readonly poolType!: pulumi.Output<string | undefined>;

    /**
     * Create a Pool resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated azuredevops.agent.Pool has been deprecated in favor of azuredevops.Pool */
    constructor(name: string, args?: PoolArgs, opts?: pulumi.CustomResourceOptions)
    /** @deprecated azuredevops.agent.Pool has been deprecated in favor of azuredevops.Pool */
    constructor(name: string, argsOrState?: PoolArgs | PoolState, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("Pool is deprecated: azuredevops.agent.Pool has been deprecated in favor of azuredevops.Pool")
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as PoolState | undefined;
            inputs["autoProvision"] = state ? state.autoProvision : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["poolType"] = state ? state.poolType : undefined;
        } else {
            const args = argsOrState as PoolArgs | undefined;
            inputs["autoProvision"] = args ? args.autoProvision : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["poolType"] = args ? args.poolType : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Pool.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Pool resources.
 */
export interface PoolState {
    /**
     * Specifies whether or not a queue should be automatically provisioned for each project collection. Defaults to `false`.
     */
    readonly autoProvision?: pulumi.Input<boolean>;
    /**
     * The name of the agent pool.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Specifies whether the agent pool type is Automation or Deployment.  Defaults to `automation`.
     */
    readonly poolType?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Pool resource.
 */
export interface PoolArgs {
    /**
     * Specifies whether or not a queue should be automatically provisioned for each project collection. Defaults to `false`.
     */
    readonly autoProvision?: pulumi.Input<boolean>;
    /**
     * The name of the agent pool.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Specifies whether the agent pool type is Automation or Deployment.  Defaults to `automation`.
     */
    readonly poolType?: pulumi.Input<string>;
}
