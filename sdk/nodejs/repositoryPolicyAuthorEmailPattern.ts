// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manage author email pattern repository policy within Azure DevOps project.
 *
 * ## Example Usage
 *
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
 * const exampleRepositoryPolicyAuthorEmailPattern = new azuredevops.RepositoryPolicyAuthorEmailPattern("exampleRepositoryPolicyAuthorEmailPattern", {
 *     projectId: exampleProject.id,
 *     enabled: true,
 *     blocking: true,
 *     authorEmailPatterns: [
 *         "user1@test.com",
 *         "user2@test.com",
 *     ],
 *     repositoryIds: [exampleGit.id],
 * });
 * ```
 * ## Set project level repository policy
 *
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
 * const exampleRepositoryPolicyAuthorEmailPattern = new azuredevops.RepositoryPolicyAuthorEmailPattern("exampleRepositoryPolicyAuthorEmailPattern", {
 *     projectId: exampleProject.id,
 *     enabled: true,
 *     blocking: true,
 *     authorEmailPatterns: [
 *         "user1@test.com",
 *         "user2@test.com",
 *     ],
 * });
 * ```
 *
 * ## Relevant Links
 *
 * - [Azure DevOps Service REST API 6.0 - Policy Configurations](https://docs.microsoft.com/en-us/rest/api/azure/devops/policy/configurations?view=azure-devops-rest-6.0)
 *
 * ## Import
 *
 * Azure DevOps Branch Policies can be imported using the project ID and policy configuration ID
 *
 * ```sh
 *  $ pulumi import azuredevops:index/repositoryPolicyAuthorEmailPattern:RepositoryPolicyAuthorEmailPattern example 00000000-0000-0000-0000-000000000000/0
 * ```
 */
export class RepositoryPolicyAuthorEmailPattern extends pulumi.CustomResource {
    /**
     * Get an existing RepositoryPolicyAuthorEmailPattern resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RepositoryPolicyAuthorEmailPatternState, opts?: pulumi.CustomResourceOptions): RepositoryPolicyAuthorEmailPattern {
        return new RepositoryPolicyAuthorEmailPattern(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuredevops:index/repositoryPolicyAuthorEmailPattern:RepositoryPolicyAuthorEmailPattern';

    /**
     * Returns true if the given object is an instance of RepositoryPolicyAuthorEmailPattern.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RepositoryPolicyAuthorEmailPattern {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RepositoryPolicyAuthorEmailPattern.__pulumiType;
    }

    /**
     * Block pushes with a commit author email that does not match the patterns. You can specify exact emails or use wildcards. 
     * Email patterns prefixed with "!" are excluded. Order is important.
     */
    public readonly authorEmailPatterns!: pulumi.Output<string[]>;
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
     * Control whether the policy is enabled for the repository or the project. If `repositoryIds` not configured, the policy will be set to the project.
     */
    public readonly repositoryIds!: pulumi.Output<string[] | undefined>;

    /**
     * Create a RepositoryPolicyAuthorEmailPattern resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RepositoryPolicyAuthorEmailPatternArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RepositoryPolicyAuthorEmailPatternArgs | RepositoryPolicyAuthorEmailPatternState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RepositoryPolicyAuthorEmailPatternState | undefined;
            resourceInputs["authorEmailPatterns"] = state ? state.authorEmailPatterns : undefined;
            resourceInputs["blocking"] = state ? state.blocking : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["repositoryIds"] = state ? state.repositoryIds : undefined;
        } else {
            const args = argsOrState as RepositoryPolicyAuthorEmailPatternArgs | undefined;
            if ((!args || args.authorEmailPatterns === undefined) && !opts.urn) {
                throw new Error("Missing required property 'authorEmailPatterns'");
            }
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            resourceInputs["authorEmailPatterns"] = args ? args.authorEmailPatterns : undefined;
            resourceInputs["blocking"] = args ? args.blocking : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["repositoryIds"] = args ? args.repositoryIds : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RepositoryPolicyAuthorEmailPattern.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RepositoryPolicyAuthorEmailPattern resources.
 */
export interface RepositoryPolicyAuthorEmailPatternState {
    /**
     * Block pushes with a commit author email that does not match the patterns. You can specify exact emails or use wildcards. 
     * Email patterns prefixed with "!" are excluded. Order is important.
     */
    authorEmailPatterns?: pulumi.Input<pulumi.Input<string>[]>;
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
     * Control whether the policy is enabled for the repository or the project. If `repositoryIds` not configured, the policy will be set to the project.
     */
    repositoryIds?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a RepositoryPolicyAuthorEmailPattern resource.
 */
export interface RepositoryPolicyAuthorEmailPatternArgs {
    /**
     * Block pushes with a commit author email that does not match the patterns. You can specify exact emails or use wildcards. 
     * Email patterns prefixed with "!" are excluded. Order is important.
     */
    authorEmailPatterns: pulumi.Input<pulumi.Input<string>[]>;
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
     * Control whether the policy is enabled for the repository or the project. If `repositoryIds` not configured, the policy will be set to the project.
     */
    repositoryIds?: pulumi.Input<pulumi.Input<string>[]>;
}
