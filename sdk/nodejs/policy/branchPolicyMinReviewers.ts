// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages a minimum reviewer branch policy within Azure DevOps.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const project = new azuredevops.Project("project", {});
 * const git = new azuredevops.Git("git", {
 *     projectId: project.id,
 *     initialization: {
 *         initType: "Clean",
 *     },
 * });
 * const branchPolicyMinReviewers = new azuredevops.BranchPolicyMinReviewers("branchPolicyMinReviewers", {
 *     projectId: project.id,
 *     enabled: true,
 *     blocking: true,
 *     settings: {
 *         reviewerCount: 2,
 *         submitterCanVote: false,
 *         scopes: [
 *             {
 *                 repositoryId: git.id,
 *                 repositoryRef: git.defaultBranch,
 *                 matchType: "Exact",
 *             },
 *             {
 *                 repositoryId: git.id,
 *                 repositoryRef: "refs/heads/releases",
 *                 matchType: "Prefix",
 *             },
 *         ],
 *     },
 * });
 * ```
 * ## Relevant Links
 *
 * - [Azure DevOps Service REST API 5.1 - Policy Configurations](https://docs.microsoft.com/en-us/rest/api/azure/devops/policy/configurations/create?view=azure-devops-rest-5.1)
 *
 * @deprecated azuredevops.policy.BranchPolicyMinReviewers has been deprecated in favor of azuredevops.BranchPolicyMinReviewers
 */
export class BranchPolicyMinReviewers extends pulumi.CustomResource {
    /**
     * Get an existing BranchPolicyMinReviewers resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BranchPolicyMinReviewersState, opts?: pulumi.CustomResourceOptions): BranchPolicyMinReviewers {
        pulumi.log.warn("BranchPolicyMinReviewers is deprecated: azuredevops.policy.BranchPolicyMinReviewers has been deprecated in favor of azuredevops.BranchPolicyMinReviewers")
        return new BranchPolicyMinReviewers(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuredevops:Policy/branchPolicyMinReviewers:BranchPolicyMinReviewers';

    /**
     * Returns true if the given object is an instance of BranchPolicyMinReviewers.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BranchPolicyMinReviewers {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BranchPolicyMinReviewers.__pulumiType;
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
    public readonly settings!: pulumi.Output<outputs.Policy.BranchPolicyMinReviewersSettings>;

    /**
     * Create a BranchPolicyMinReviewers resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated azuredevops.policy.BranchPolicyMinReviewers has been deprecated in favor of azuredevops.BranchPolicyMinReviewers */
    constructor(name: string, args: BranchPolicyMinReviewersArgs, opts?: pulumi.CustomResourceOptions)
    /** @deprecated azuredevops.policy.BranchPolicyMinReviewers has been deprecated in favor of azuredevops.BranchPolicyMinReviewers */
    constructor(name: string, argsOrState?: BranchPolicyMinReviewersArgs | BranchPolicyMinReviewersState, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("BranchPolicyMinReviewers is deprecated: azuredevops.policy.BranchPolicyMinReviewers has been deprecated in favor of azuredevops.BranchPolicyMinReviewers")
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as BranchPolicyMinReviewersState | undefined;
            inputs["blocking"] = state ? state.blocking : undefined;
            inputs["enabled"] = state ? state.enabled : undefined;
            inputs["projectId"] = state ? state.projectId : undefined;
            inputs["settings"] = state ? state.settings : undefined;
        } else {
            const args = argsOrState as BranchPolicyMinReviewersArgs | undefined;
            if (!args || args.projectId === undefined) {
                throw new Error("Missing required property 'projectId'");
            }
            if (!args || args.settings === undefined) {
                throw new Error("Missing required property 'settings'");
            }
            inputs["blocking"] = args ? args.blocking : undefined;
            inputs["enabled"] = args ? args.enabled : undefined;
            inputs["projectId"] = args ? args.projectId : undefined;
            inputs["settings"] = args ? args.settings : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(BranchPolicyMinReviewers.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BranchPolicyMinReviewers resources.
 */
export interface BranchPolicyMinReviewersState {
    /**
     * A flag indicating if the policy should be blocking. Defaults to `true`.
     */
    readonly blocking?: pulumi.Input<boolean>;
    /**
     * A flag indicating if the policy should be enabled. Defaults to `true`.
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * The ID of the project in which the policy will be created.
     */
    readonly projectId?: pulumi.Input<string>;
    /**
     * Configuration for the policy. This block must be defined exactly once.
     */
    readonly settings?: pulumi.Input<inputs.Policy.BranchPolicyMinReviewersSettings>;
}

/**
 * The set of arguments for constructing a BranchPolicyMinReviewers resource.
 */
export interface BranchPolicyMinReviewersArgs {
    /**
     * A flag indicating if the policy should be blocking. Defaults to `true`.
     */
    readonly blocking?: pulumi.Input<boolean>;
    /**
     * A flag indicating if the policy should be enabled. Defaults to `true`.
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * The ID of the project in which the policy will be created.
     */
    readonly projectId: pulumi.Input<string>;
    /**
     * Configuration for the policy. This block must be defined exactly once.
     */
    readonly settings: pulumi.Input<inputs.Policy.BranchPolicyMinReviewersSettings>;
}
