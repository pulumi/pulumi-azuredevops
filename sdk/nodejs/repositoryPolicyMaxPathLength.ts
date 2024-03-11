// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manage a max path length repository policy within Azure DevOps project.
 *
 * > If both project and project policy are enabled, the repository policy has high priority.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const exampleProject = new azuredevops.Project("exampleProject", {
 *     visibility: "private",
 *     versionControl: "Git",
 *     workItemTemplate: "Agile",
 *     description: "Managed by Terraform",
 * });
 * const exampleGit = new azuredevops.Git("exampleGit", {
 *     projectId: exampleProject.id,
 *     initialization: {
 *         initType: "Clean",
 *     },
 * });
 * const exampleRepositoryPolicyMaxPathLength = new azuredevops.RepositoryPolicyMaxPathLength("exampleRepositoryPolicyMaxPathLength", {
 *     projectId: exampleProject.id,
 *     enabled: true,
 *     blocking: true,
 *     maxPathLength: 500,
 *     repositoryIds: [exampleGit.id],
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * # Set project level repository policy
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const exampleProject = new azuredevops.Project("exampleProject", {
 *     visibility: "private",
 *     versionControl: "Git",
 *     workItemTemplate: "Agile",
 *     description: "Managed by Terraform",
 * });
 * const exampleRepositoryPolicyMaxPathLength = new azuredevops.RepositoryPolicyMaxPathLength("exampleRepositoryPolicyMaxPathLength", {
 *     projectId: exampleProject.id,
 *     enabled: true,
 *     blocking: true,
 *     maxPathLength: 1000,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Relevant Links
 *
 * - [Azure DevOps Service REST API 7.0 - Policy Configurations](https://docs.microsoft.com/en-us/rest/api/azure/devops/policy/configurations?view=azure-devops-rest-7.0)
 *
 * ## Import
 *
 * Azure DevOps repository policies can be imported using the projectID/policyID or projectName/policyID:
 *
 * ```sh
 * $ pulumi import azuredevops:index/repositoryPolicyMaxPathLength:RepositoryPolicyMaxPathLength example 00000000-0000-0000-0000-000000000000/0
 * ```
 */
export class RepositoryPolicyMaxPathLength extends pulumi.CustomResource {
    /**
     * Get an existing RepositoryPolicyMaxPathLength resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RepositoryPolicyMaxPathLengthState, opts?: pulumi.CustomResourceOptions): RepositoryPolicyMaxPathLength {
        return new RepositoryPolicyMaxPathLength(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuredevops:index/repositoryPolicyMaxPathLength:RepositoryPolicyMaxPathLength';

    /**
     * Returns true if the given object is an instance of RepositoryPolicyMaxPathLength.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RepositoryPolicyMaxPathLength {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RepositoryPolicyMaxPathLength.__pulumiType;
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
     * Block pushes that introduce paths that exceed the specified length.
     */
    public readonly maxPathLength!: pulumi.Output<number>;
    /**
     * The ID of the project in which the policy will be created.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * Control whether the policy is enabled for the repository or the project. If `repositoryIds` not configured, the policy will be set to the project.
     */
    public readonly repositoryIds!: pulumi.Output<string[] | undefined>;

    /**
     * Create a RepositoryPolicyMaxPathLength resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RepositoryPolicyMaxPathLengthArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RepositoryPolicyMaxPathLengthArgs | RepositoryPolicyMaxPathLengthState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RepositoryPolicyMaxPathLengthState | undefined;
            resourceInputs["blocking"] = state ? state.blocking : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["maxPathLength"] = state ? state.maxPathLength : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["repositoryIds"] = state ? state.repositoryIds : undefined;
        } else {
            const args = argsOrState as RepositoryPolicyMaxPathLengthArgs | undefined;
            if ((!args || args.maxPathLength === undefined) && !opts.urn) {
                throw new Error("Missing required property 'maxPathLength'");
            }
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            resourceInputs["blocking"] = args ? args.blocking : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["maxPathLength"] = args ? args.maxPathLength : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["repositoryIds"] = args ? args.repositoryIds : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RepositoryPolicyMaxPathLength.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RepositoryPolicyMaxPathLength resources.
 */
export interface RepositoryPolicyMaxPathLengthState {
    /**
     * A flag indicating if the policy should be blocking. Defaults to `true`.
     */
    blocking?: pulumi.Input<boolean>;
    /**
     * A flag indicating if the policy should be enabled. Defaults to `true`.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * Block pushes that introduce paths that exceed the specified length.
     */
    maxPathLength?: pulumi.Input<number>;
    /**
     * The ID of the project in which the policy will be created.
     */
    projectId?: pulumi.Input<string>;
    /**
     * Control whether the policy is enabled for the repository or the project. If `repositoryIds` not configured, the policy will be set to the project.
     */
    repositoryIds?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a RepositoryPolicyMaxPathLength resource.
 */
export interface RepositoryPolicyMaxPathLengthArgs {
    /**
     * A flag indicating if the policy should be blocking. Defaults to `true`.
     */
    blocking?: pulumi.Input<boolean>;
    /**
     * A flag indicating if the policy should be enabled. Defaults to `true`.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * Block pushes that introduce paths that exceed the specified length.
     */
    maxPathLength: pulumi.Input<number>;
    /**
     * The ID of the project in which the policy will be created.
     */
    projectId: pulumi.Input<string>;
    /**
     * Control whether the policy is enabled for the repository or the project. If `repositoryIds` not configured, the policy will be set to the project.
     */
    repositoryIds?: pulumi.Input<pulumi.Input<string>[]>;
}
