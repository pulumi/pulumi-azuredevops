// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Manages a status check branch policy within Azure DevOps.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const example = new azuredevops.Project("example", {
 *     name: "Example Project",
 *     visibility: "private",
 *     versionControl: "Git",
 *     workItemTemplate: "Agile",
 *     features: {
 *         testplans: "disabled",
 *         artifacts: "disabled",
 *     },
 *     description: "Managed by Pulumi",
 * });
 * const exampleGit = new azuredevops.Git("example", {
 *     projectId: example.id,
 *     name: "Example Repository",
 *     initialization: {
 *         initType: "Clean",
 *     },
 * });
 * const exampleUser = new azuredevops.User("example", {
 *     principalName: "mail@email.com",
 *     accountLicenseType: "basic",
 * });
 * const exampleBranchPolicyStatusCheck = new azuredevops.BranchPolicyStatusCheck("example", {
 *     projectId: example.id,
 *     enabled: true,
 *     blocking: true,
 *     settings: {
 *         name: "Release",
 *         authorId: exampleUser.id,
 *         invalidateOnUpdate: true,
 *         applicability: "conditional",
 *         displayName: "PreCheck",
 *         scopes: [
 *             {
 *                 repositoryId: exampleGit.id,
 *                 repositoryRef: exampleGit.defaultBranch,
 *                 matchType: "Exact",
 *             },
 *             {
 *                 matchType: "DefaultBranch",
 *             },
 *         ],
 *     },
 * });
 * ```
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
 * $ pulumi import azuredevops:index/branchPolicyStatusCheck:BranchPolicyStatusCheck example 00000000-0000-0000-0000-000000000000/0
 * ```
 */
export class BranchPolicyStatusCheck extends pulumi.CustomResource {
    /**
     * Get an existing BranchPolicyStatusCheck resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BranchPolicyStatusCheckState, opts?: pulumi.CustomResourceOptions): BranchPolicyStatusCheck {
        return new BranchPolicyStatusCheck(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuredevops:index/branchPolicyStatusCheck:BranchPolicyStatusCheck';

    /**
     * Returns true if the given object is an instance of BranchPolicyStatusCheck.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BranchPolicyStatusCheck {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BranchPolicyStatusCheck.__pulumiType;
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
     * A `settings` block as defined below. Configuration for the policy. This block must be defined exactly once.
     */
    public readonly settings!: pulumi.Output<outputs.BranchPolicyStatusCheckSettings>;

    /**
     * Create a BranchPolicyStatusCheck resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BranchPolicyStatusCheckArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BranchPolicyStatusCheckArgs | BranchPolicyStatusCheckState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BranchPolicyStatusCheckState | undefined;
            resourceInputs["blocking"] = state ? state.blocking : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["settings"] = state ? state.settings : undefined;
        } else {
            const args = argsOrState as BranchPolicyStatusCheckArgs | undefined;
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
        super(BranchPolicyStatusCheck.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BranchPolicyStatusCheck resources.
 */
export interface BranchPolicyStatusCheckState {
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
     * A `settings` block as defined below. Configuration for the policy. This block must be defined exactly once.
     */
    settings?: pulumi.Input<inputs.BranchPolicyStatusCheckSettings>;
}

/**
 * The set of arguments for constructing a BranchPolicyStatusCheck resource.
 */
export interface BranchPolicyStatusCheckArgs {
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
     * A `settings` block as defined below. Configuration for the policy. This block must be defined exactly once.
     */
    settings: pulumi.Input<inputs.BranchPolicyStatusCheckSettings>;
}
