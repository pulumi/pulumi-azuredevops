// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Branch policy for merge types allowed on a specified branch.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const example = new azuredevops.Project("example", {name: "Example Project"});
 * const exampleGit = new azuredevops.Git("example", {
 *     projectId: example.id,
 *     name: "Example Repository",
 *     initialization: {
 *         initType: "Clean",
 *     },
 * });
 * const exampleBranchPolicyMergeTypes = new azuredevops.BranchPolicyMergeTypes("example", {
 *     projectId: example.id,
 *     enabled: true,
 *     blocking: true,
 *     settings: {
 *         allowSquash: true,
 *         allowRebaseAndFastForward: true,
 *         allowBasicNoFastForward: true,
 *         allowRebaseWithMerge: true,
 *         scopes: [
 *             {
 *                 repositoryId: exampleGit.id,
 *                 repositoryRef: exampleGit.defaultBranch,
 *                 matchType: "Exact",
 *             },
 *             {
 *                 repositoryId: undefined,
 *                 repositoryRef: "refs/heads/releases",
 *                 matchType: "Prefix",
 *             },
 *             {
 *                 matchType: "DefaultBranch",
 *             },
 *         ],
 *     },
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Relevant Links
 *
 * - [Azure DevOps Service REST API 7.0 - Policy Configurations](https://docs.microsoft.com/en-us/rest/api/azure/devops/policy/configurations/create?view=azure-devops-rest-7.0)
 *
 * ## Import
 *
 * Azure DevOps Branch Policies can be imported using the project ID and policy configuration ID:
 *
 * ```sh
 * $ pulumi import azuredevops:index/branchPolicyMergeTypes:BranchPolicyMergeTypes example 00000000-0000-0000-0000-000000000000/0
 * ```
 */
export class BranchPolicyMergeTypes extends pulumi.CustomResource {
    /**
     * Get an existing BranchPolicyMergeTypes resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BranchPolicyMergeTypesState, opts?: pulumi.CustomResourceOptions): BranchPolicyMergeTypes {
        return new BranchPolicyMergeTypes(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuredevops:index/branchPolicyMergeTypes:BranchPolicyMergeTypes';

    /**
     * Returns true if the given object is an instance of BranchPolicyMergeTypes.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BranchPolicyMergeTypes {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BranchPolicyMergeTypes.__pulumiType;
    }

    /**
     * A flag indicating if the policy should be blocking. Defaults to `true`.
     */
    public readonly blocking!: pulumi.Output<boolean | undefined>;
    /**
     * A flag indicating if the policy should be enabled. Defaults to `true`.
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * The ID of the project in which the policy will be created.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * Configuration for the policy. This block must be defined exactly once.
     */
    public readonly settings!: pulumi.Output<outputs.BranchPolicyMergeTypesSettings>;

    /**
     * Create a BranchPolicyMergeTypes resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BranchPolicyMergeTypesArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BranchPolicyMergeTypesArgs | BranchPolicyMergeTypesState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BranchPolicyMergeTypesState | undefined;
            resourceInputs["blocking"] = state ? state.blocking : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["settings"] = state ? state.settings : undefined;
        } else {
            const args = argsOrState as BranchPolicyMergeTypesArgs | undefined;
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            if ((!args || args.settings === undefined) && !opts.urn) {
                throw new Error("Missing required property 'settings'");
            }
            resourceInputs["blocking"] = args ? args.blocking : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["settings"] = args ? args.settings : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(BranchPolicyMergeTypes.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BranchPolicyMergeTypes resources.
 */
export interface BranchPolicyMergeTypesState {
    /**
     * A flag indicating if the policy should be blocking. Defaults to `true`.
     */
    blocking?: pulumi.Input<boolean>;
    /**
     * A flag indicating if the policy should be enabled. Defaults to `true`.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * The ID of the project in which the policy will be created.
     */
    projectId?: pulumi.Input<string>;
    /**
     * Configuration for the policy. This block must be defined exactly once.
     */
    settings?: pulumi.Input<inputs.BranchPolicyMergeTypesSettings>;
}

/**
 * The set of arguments for constructing a BranchPolicyMergeTypes resource.
 */
export interface BranchPolicyMergeTypesArgs {
    /**
     * A flag indicating if the policy should be blocking. Defaults to `true`.
     */
    blocking?: pulumi.Input<boolean>;
    /**
     * A flag indicating if the policy should be enabled. Defaults to `true`.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * The ID of the project in which the policy will be created.
     */
    projectId: pulumi.Input<string>;
    /**
     * Configuration for the policy. This block must be defined exactly once.
     */
    settings: pulumi.Input<inputs.BranchPolicyMergeTypesSettings>;
}
