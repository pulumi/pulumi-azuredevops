// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Manage a file path pattern repository policy within Azure DevOps project.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const project = new azuredevops.Project("project", {
 *     description: "Managed by Terraform",
 *     visibility: "private",
 *     versionControl: "Git",
 *     workItemTemplate: "Agile",
 * });
 * const git = new azuredevops.Git("git", {
 *     projectId: project.id,
 *     initialization: {
 *         initType: "Clean",
 *     },
 * });
 * const repositoryPolicyFilePathPattern = new azuredevops.RepositoryPolicyFilePathPattern("repositoryPolicyFilePathPattern", {
 *     projectId: project.id,
 *     enabled: true,
 *     blocking: true,
 *     settings: {
 *         filepathPatterns: [
 *             "*.go",
 *             "/home/test/*.ts",
 *         ],
 *         scopes: [{
 *             repositoryId: git.id,
 *         }],
 *     },
 * });
 * ```
 * ## Relevant Links
 *
 * - [Azure DevOps Service REST API 5.1 - Policy Configurations](https://docs.microsoft.com/en-us/rest/api/azure/devops/policy/configurations/create?view=azure-devops-rest-5.1)
 *
 * ## Import
 *
 * Azure DevOps repository policies can be imported using the projectID/policyID or projectName/policyID
 *
 * ```sh
 *  $ pulumi import azuredevops:index/repositoryPolicyFilePathPattern:RepositoryPolicyFilePathPattern p 00000000-0000-0000-0000-000000000000/0
 * ```
 */
export class RepositoryPolicyFilePathPattern extends pulumi.CustomResource {
    /**
     * Get an existing RepositoryPolicyFilePathPattern resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RepositoryPolicyFilePathPatternState, opts?: pulumi.CustomResourceOptions): RepositoryPolicyFilePathPattern {
        return new RepositoryPolicyFilePathPattern(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuredevops:index/repositoryPolicyFilePathPattern:RepositoryPolicyFilePathPattern';

    /**
     * Returns true if the given object is an instance of RepositoryPolicyFilePathPattern.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RepositoryPolicyFilePathPattern {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RepositoryPolicyFilePathPattern.__pulumiType;
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
    public readonly settings!: pulumi.Output<outputs.RepositoryPolicyFilePathPatternSettings>;

    /**
     * Create a RepositoryPolicyFilePathPattern resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RepositoryPolicyFilePathPatternArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RepositoryPolicyFilePathPatternArgs | RepositoryPolicyFilePathPatternState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RepositoryPolicyFilePathPatternState | undefined;
            inputs["blocking"] = state ? state.blocking : undefined;
            inputs["enabled"] = state ? state.enabled : undefined;
            inputs["projectId"] = state ? state.projectId : undefined;
            inputs["settings"] = state ? state.settings : undefined;
        } else {
            const args = argsOrState as RepositoryPolicyFilePathPatternArgs | undefined;
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            if ((!args || args.settings === undefined) && !opts.urn) {
                throw new Error("Missing required property 'settings'");
            }
            inputs["blocking"] = args ? args.blocking : undefined;
            inputs["enabled"] = args ? args.enabled : undefined;
            inputs["projectId"] = args ? args.projectId : undefined;
            inputs["settings"] = args ? args.settings : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(RepositoryPolicyFilePathPattern.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RepositoryPolicyFilePathPattern resources.
 */
export interface RepositoryPolicyFilePathPatternState {
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
    settings?: pulumi.Input<inputs.RepositoryPolicyFilePathPatternSettings>;
}

/**
 * The set of arguments for constructing a RepositoryPolicyFilePathPattern resource.
 */
export interface RepositoryPolicyFilePathPatternArgs {
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
    settings: pulumi.Input<inputs.RepositoryPolicyFilePathPatternSettings>;
}
