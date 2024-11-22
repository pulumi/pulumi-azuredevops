// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manages a case enforcement repository policy within Azure DevOps project.
 *
 * > If both project and project policy are enabled, the project policy has high priority.
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
 *     description: "Managed by Pulumi",
 * });
 * const exampleGit = new azuredevops.Git("example", {
 *     projectId: example.id,
 *     name: "Example Repository",
 *     initialization: {
 *         initType: "Clean",
 *     },
 * });
 * const exampleRepositoryPolicyCaseEnforcement = new azuredevops.RepositoryPolicyCaseEnforcement("example", {
 *     projectId: example.id,
 *     enabled: true,
 *     blocking: true,
 *     enforceConsistentCase: true,
 *     repositoryIds: [exampleGit.id],
 * });
 * ```
 *
 * # Set project level repository policy
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const example = new azuredevops.Project("example", {
 *     name: "Example Project",
 *     visibility: "private",
 *     versionControl: "Git",
 *     workItemTemplate: "Agile",
 *     description: "Managed by Pulumi",
 * });
 * const exampleRepositoryPolicyCaseEnforcement = new azuredevops.RepositoryPolicyCaseEnforcement("example", {
 *     projectId: example.id,
 *     enabled: true,
 *     blocking: true,
 *     enforceConsistentCase: true,
 * });
 * ```
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
 * $ pulumi import azuredevops:index/repositoryPolicyCaseEnforcement:RepositoryPolicyCaseEnforcement example 00000000-0000-0000-0000-000000000000/0
 * ```
 */
export class RepositoryPolicyCaseEnforcement extends pulumi.CustomResource {
    /**
     * Get an existing RepositoryPolicyCaseEnforcement resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RepositoryPolicyCaseEnforcementState, opts?: pulumi.CustomResourceOptions): RepositoryPolicyCaseEnforcement {
        return new RepositoryPolicyCaseEnforcement(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuredevops:index/repositoryPolicyCaseEnforcement:RepositoryPolicyCaseEnforcement';

    /**
     * Returns true if the given object is an instance of RepositoryPolicyCaseEnforcement.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RepositoryPolicyCaseEnforcement {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RepositoryPolicyCaseEnforcement.__pulumiType;
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
     * Avoid case-sensitivity conflicts by blocking pushes that change name casing on files, folders, branches, and tags.
     */
    public readonly enforceConsistentCase!: pulumi.Output<boolean>;
    /**
     * The ID of the project in which the policy will be created.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * Control whether the policy is enabled for the repository or the project. If `repositoryIds` not configured, the policy will be set to the project.
     */
    public readonly repositoryIds!: pulumi.Output<string[] | undefined>;

    /**
     * Create a RepositoryPolicyCaseEnforcement resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RepositoryPolicyCaseEnforcementArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RepositoryPolicyCaseEnforcementArgs | RepositoryPolicyCaseEnforcementState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RepositoryPolicyCaseEnforcementState | undefined;
            resourceInputs["blocking"] = state ? state.blocking : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["enforceConsistentCase"] = state ? state.enforceConsistentCase : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["repositoryIds"] = state ? state.repositoryIds : undefined;
        } else {
            const args = argsOrState as RepositoryPolicyCaseEnforcementArgs | undefined;
            if ((!args || args.enforceConsistentCase === undefined) && !opts.urn) {
                throw new Error("Missing required property 'enforceConsistentCase'");
            }
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            resourceInputs["blocking"] = args ? args.blocking : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["enforceConsistentCase"] = args ? args.enforceConsistentCase : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["repositoryIds"] = args ? args.repositoryIds : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RepositoryPolicyCaseEnforcement.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RepositoryPolicyCaseEnforcement resources.
 */
export interface RepositoryPolicyCaseEnforcementState {
    /**
     * A flag indicating if the policy should be blocking. Defaults to `true`.
     */
    blocking?: pulumi.Input<boolean>;
    /**
     * A flag indicating if the policy should be enabled. Defaults to `true`.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * Avoid case-sensitivity conflicts by blocking pushes that change name casing on files, folders, branches, and tags.
     */
    enforceConsistentCase?: pulumi.Input<boolean>;
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
 * The set of arguments for constructing a RepositoryPolicyCaseEnforcement resource.
 */
export interface RepositoryPolicyCaseEnforcementArgs {
    /**
     * A flag indicating if the policy should be blocking. Defaults to `true`.
     */
    blocking?: pulumi.Input<boolean>;
    /**
     * A flag indicating if the policy should be enabled. Defaults to `true`.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * Avoid case-sensitivity conflicts by blocking pushes that change name casing on files, folders, branches, and tags.
     */
    enforceConsistentCase: pulumi.Input<boolean>;
    /**
     * The ID of the project in which the policy will be created.
     */
    projectId: pulumi.Input<string>;
    /**
     * Control whether the policy is enabled for the repository or the project. If `repositoryIds` not configured, the policy will be set to the project.
     */
    repositoryIds?: pulumi.Input<pulumi.Input<string>[]>;
}
