// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Manages a build validation branch policy within Azure DevOps.
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
 * const buildDefinition = new azuredevops.BuildDefinition("buildDefinition", {
 *     projectId: project.id,
 *     repository: {
 *         repoType: "TfsGit",
 *         repoId: git.id,
 *         ymlPath: "azure-pipelines.yml",
 *     },
 * });
 * const branchPolicyBuildValidation = new azuredevops.BranchPolicyBuildValidation("branchPolicyBuildValidation", {
 *     projectId: project.id,
 *     enabled: true,
 *     blocking: true,
 *     settings: {
 *         displayName: "Don't break the build!",
 *         buildDefinitionId: buildDefinition.id,
 *         validDuration: 720,
 *         filenamePatterns: [
 *             "/WebApp/*",
 *             "!/WebApp/Tests/*",
 *             "*.cs",
 *         ],
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
 * ## Import
 *
 * Azure DevOps Branch Policies can be imported using the project ID and policy configuration ID
 *
 * ```sh
 *  $ pulumi import azuredevops:index/branchPolicyBuildValidation:BranchPolicyBuildValidation p 00000000-0000-0000-0000-000000000000/0
 * ```
 */
export class BranchPolicyBuildValidation extends pulumi.CustomResource {
    /**
     * Get an existing BranchPolicyBuildValidation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BranchPolicyBuildValidationState, opts?: pulumi.CustomResourceOptions): BranchPolicyBuildValidation {
        return new BranchPolicyBuildValidation(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuredevops:index/branchPolicyBuildValidation:BranchPolicyBuildValidation';

    /**
     * Returns true if the given object is an instance of BranchPolicyBuildValidation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BranchPolicyBuildValidation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BranchPolicyBuildValidation.__pulumiType;
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
    public readonly settings!: pulumi.Output<outputs.BranchPolicyBuildValidationSettings>;

    /**
     * Create a BranchPolicyBuildValidation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BranchPolicyBuildValidationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BranchPolicyBuildValidationArgs | BranchPolicyBuildValidationState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BranchPolicyBuildValidationState | undefined;
            inputs["blocking"] = state ? state.blocking : undefined;
            inputs["enabled"] = state ? state.enabled : undefined;
            inputs["projectId"] = state ? state.projectId : undefined;
            inputs["settings"] = state ? state.settings : undefined;
        } else {
            const args = argsOrState as BranchPolicyBuildValidationArgs | undefined;
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
        const aliasOpts = { aliases: [{ type: "azuredevops:Policy/branchPolicyBuildValidation:BranchPolicyBuildValidation" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(BranchPolicyBuildValidation.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BranchPolicyBuildValidation resources.
 */
export interface BranchPolicyBuildValidationState {
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
    settings?: pulumi.Input<inputs.BranchPolicyBuildValidationSettings>;
}

/**
 * The set of arguments for constructing a BranchPolicyBuildValidation resource.
 */
export interface BranchPolicyBuildValidationArgs {
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
    settings: pulumi.Input<inputs.BranchPolicyBuildValidationSettings>;
}
