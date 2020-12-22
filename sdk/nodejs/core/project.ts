// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a project within Azure DevOps.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const project = new azuredevops.Project("project", {
 *     description: "Test Project Description",
 *     features: {
 *         artifacts: "disabled",
 *         testplans: "disabled",
 *     },
 *     versionControl: "Git",
 *     visibility: "private",
 *     workItemTemplate: "Agile",
 * });
 * ```
 * ## Relevant Links
 *
 * - [Azure DevOps Service REST API 5.1 - Projects](https://docs.microsoft.com/en-us/rest/api/azure/devops/core/projects?view=azure-devops-rest-5.1)
 *
 * ## PAT Permissions Required
 *
 * - **Project & Team**: Read, Write, & Manage
 *
 * ## Import
 *
 * Azure DevOps Projects can be imported using the project name or by the project Guid, e.g.
 *
 * ```sh
 *  $ pulumi import azuredevops:Core/project:Project project "Test Project"
 * ```
 *
 *  or
 *
 * ```sh
 *  $ pulumi import azuredevops:Core/project:Project project 00000000-0000-0000-0000-000000000000
 * ```
 *
 * @deprecated azuredevops.core.Project has been deprecated in favor of azuredevops.Project
 */
export class Project extends pulumi.CustomResource {
    /**
     * Get an existing Project resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProjectState, opts?: pulumi.CustomResourceOptions): Project {
        pulumi.log.warn("Project is deprecated: azuredevops.core.Project has been deprecated in favor of azuredevops.Project")
        return new Project(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuredevops:Core/project:Project';

    /**
     * Returns true if the given object is an instance of Project.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Project {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Project.__pulumiType;
    }

    /**
     * The Description of the Project.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Defines the status (`enabled`, `disabled`) of the project features.  
     * Valid features `boards`, `repositories`, `pipelines`, `testplans`, `artifacts`
     */
    public readonly features!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The Project Name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The Process Template ID used by the Project.
     */
    public /*out*/ readonly processTemplateId!: pulumi.Output<string>;
    /**
     * Specifies the version control system. Valid values: `Git` or `Tfvc`. Defaults to `Git`.
     */
    public readonly versionControl!: pulumi.Output<string | undefined>;
    /**
     * Specifies the visibility of the Project. Valid values: `private` or `public`. Defaults to `private`.
     */
    public readonly visibility!: pulumi.Output<string | undefined>;
    /**
     * Specifies the work item template. Defaults to `Agile`.
     */
    public readonly workItemTemplate!: pulumi.Output<string | undefined>;

    /**
     * Create a Project resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated azuredevops.core.Project has been deprecated in favor of azuredevops.Project */
    constructor(name: string, args?: ProjectArgs, opts?: pulumi.CustomResourceOptions)
    /** @deprecated azuredevops.core.Project has been deprecated in favor of azuredevops.Project */
    constructor(name: string, argsOrState?: ProjectArgs | ProjectState, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("Project is deprecated: azuredevops.core.Project has been deprecated in favor of azuredevops.Project")
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ProjectState | undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["features"] = state ? state.features : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["processTemplateId"] = state ? state.processTemplateId : undefined;
            inputs["versionControl"] = state ? state.versionControl : undefined;
            inputs["visibility"] = state ? state.visibility : undefined;
            inputs["workItemTemplate"] = state ? state.workItemTemplate : undefined;
        } else {
            const args = argsOrState as ProjectArgs | undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["features"] = args ? args.features : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["versionControl"] = args ? args.versionControl : undefined;
            inputs["visibility"] = args ? args.visibility : undefined;
            inputs["workItemTemplate"] = args ? args.workItemTemplate : undefined;
            inputs["processTemplateId"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Project.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Project resources.
 */
export interface ProjectState {
    /**
     * The Description of the Project.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Defines the status (`enabled`, `disabled`) of the project features.  
     * Valid features `boards`, `repositories`, `pipelines`, `testplans`, `artifacts`
     */
    readonly features?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The Project Name.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The Process Template ID used by the Project.
     */
    readonly processTemplateId?: pulumi.Input<string>;
    /**
     * Specifies the version control system. Valid values: `Git` or `Tfvc`. Defaults to `Git`.
     */
    readonly versionControl?: pulumi.Input<string>;
    /**
     * Specifies the visibility of the Project. Valid values: `private` or `public`. Defaults to `private`.
     */
    readonly visibility?: pulumi.Input<string>;
    /**
     * Specifies the work item template. Defaults to `Agile`.
     */
    readonly workItemTemplate?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Project resource.
 */
export interface ProjectArgs {
    /**
     * The Description of the Project.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Defines the status (`enabled`, `disabled`) of the project features.  
     * Valid features `boards`, `repositories`, `pipelines`, `testplans`, `artifacts`
     */
    readonly features?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The Project Name.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Specifies the version control system. Valid values: `Git` or `Tfvc`. Defaults to `Git`.
     */
    readonly versionControl?: pulumi.Input<string>;
    /**
     * Specifies the visibility of the Project. Valid values: `private` or `public`. Defaults to `private`.
     */
    readonly visibility?: pulumi.Input<string>;
    /**
     * Specifies the work item template. Defaults to `Agile`.
     */
    readonly workItemTemplate?: pulumi.Input<string>;
}
