// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manages Pipeline Settings for Azure DevOps projects
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
 * const exampleProjectPipelineSettings = new azuredevops.ProjectPipelineSettings("exampleProjectPipelineSettings", {
 *     projectId: exampleProject.id,
 *     enforceJobScope: true,
 *     enforceReferencedRepoScopedToken: false,
 *     enforceSettableVar: true,
 *     publishPipelineMetadata: false,
 *     statusBadgesArePrivate: true,
 * });
 * ```
 * ## Relevant Links
 *
 * No official documentation available
 *
 * ## PAT Permissions Required
 *
 * - Full Access
 *
 * ## Import
 *
 * Azure DevOps feature settings can be imported using the project id, e.g.
 *
 * ```sh
 *  $ pulumi import azuredevops:index/projectPipelineSettings:ProjectPipelineSettings example 00000000-0000-0000-0000-000000000000
 * ```
 */
export class ProjectPipelineSettings extends pulumi.CustomResource {
    /**
     * Get an existing ProjectPipelineSettings resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProjectPipelineSettingsState, opts?: pulumi.CustomResourceOptions): ProjectPipelineSettings {
        return new ProjectPipelineSettings(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuredevops:index/projectPipelineSettings:ProjectPipelineSettings';

    /**
     * Returns true if the given object is an instance of ProjectPipelineSettings.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ProjectPipelineSettings {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ProjectPipelineSettings.__pulumiType;
    }

    /**
     * Limit job authorization scope to current project for non-release pipelines.
     */
    public readonly enforceJobScope!: pulumi.Output<boolean>;
    /**
     * Protect access to repositories in YAML pipelines.
     */
    public readonly enforceReferencedRepoScopedToken!: pulumi.Output<boolean>;
    /**
     * Limit variables that can be set at queue time.
     */
    public readonly enforceSettableVar!: pulumi.Output<boolean>;
    /**
     * The `id` of the project for which the project pipeline settings will be managed.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * Publish metadata from pipelines.
     */
    public readonly publishPipelineMetadata!: pulumi.Output<boolean>;
    /**
     * Disable anonymous access to badges.
     */
    public readonly statusBadgesArePrivate!: pulumi.Output<boolean>;

    /**
     * Create a ProjectPipelineSettings resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProjectPipelineSettingsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProjectPipelineSettingsArgs | ProjectPipelineSettingsState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ProjectPipelineSettingsState | undefined;
            resourceInputs["enforceJobScope"] = state ? state.enforceJobScope : undefined;
            resourceInputs["enforceReferencedRepoScopedToken"] = state ? state.enforceReferencedRepoScopedToken : undefined;
            resourceInputs["enforceSettableVar"] = state ? state.enforceSettableVar : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["publishPipelineMetadata"] = state ? state.publishPipelineMetadata : undefined;
            resourceInputs["statusBadgesArePrivate"] = state ? state.statusBadgesArePrivate : undefined;
        } else {
            const args = argsOrState as ProjectPipelineSettingsArgs | undefined;
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            resourceInputs["enforceJobScope"] = args ? args.enforceJobScope : undefined;
            resourceInputs["enforceReferencedRepoScopedToken"] = args ? args.enforceReferencedRepoScopedToken : undefined;
            resourceInputs["enforceSettableVar"] = args ? args.enforceSettableVar : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["publishPipelineMetadata"] = args ? args.publishPipelineMetadata : undefined;
            resourceInputs["statusBadgesArePrivate"] = args ? args.statusBadgesArePrivate : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ProjectPipelineSettings.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ProjectPipelineSettings resources.
 */
export interface ProjectPipelineSettingsState {
    /**
     * Limit job authorization scope to current project for non-release pipelines.
     */
    enforceJobScope?: pulumi.Input<boolean>;
    /**
     * Protect access to repositories in YAML pipelines.
     */
    enforceReferencedRepoScopedToken?: pulumi.Input<boolean>;
    /**
     * Limit variables that can be set at queue time.
     */
    enforceSettableVar?: pulumi.Input<boolean>;
    /**
     * The `id` of the project for which the project pipeline settings will be managed.
     */
    projectId?: pulumi.Input<string>;
    /**
     * Publish metadata from pipelines.
     */
    publishPipelineMetadata?: pulumi.Input<boolean>;
    /**
     * Disable anonymous access to badges.
     */
    statusBadgesArePrivate?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a ProjectPipelineSettings resource.
 */
export interface ProjectPipelineSettingsArgs {
    /**
     * Limit job authorization scope to current project for non-release pipelines.
     */
    enforceJobScope?: pulumi.Input<boolean>;
    /**
     * Protect access to repositories in YAML pipelines.
     */
    enforceReferencedRepoScopedToken?: pulumi.Input<boolean>;
    /**
     * Limit variables that can be set at queue time.
     */
    enforceSettableVar?: pulumi.Input<boolean>;
    /**
     * The `id` of the project for which the project pipeline settings will be managed.
     */
    projectId: pulumi.Input<string>;
    /**
     * Publish metadata from pipelines.
     */
    publishPipelineMetadata?: pulumi.Input<boolean>;
    /**
     * Disable anonymous access to badges.
     */
    statusBadgesArePrivate?: pulumi.Input<boolean>;
}
