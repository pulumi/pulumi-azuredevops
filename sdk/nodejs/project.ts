// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manages a project within Azure DevOps.
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
 *     features: {
 *         testplans: "disabled",
 *         artifacts: "disabled",
 *     },
 * });
 * ```
 *
 * ## Relevant Links
 *
 * - [Azure DevOps Service REST API 7.0 - Projects](https://docs.microsoft.com/en-us/rest/api/azure/devops/core/projects?view=azure-devops-rest-7.0)
 *
 * ## PAT Permissions Required
 *
 * - **Project & Team**: Read, Write, & Manage
 * - **Work Items**: Read
 *
 * ## Import
 *
 * Azure DevOps Projects can be imported using the project name or by the project Guid, e.g.
 *
 * ```sh
 * $ pulumi import azuredevops:index/project:Project example "Example Project"
 * ```
 *
 * or
 *
 * ```sh
 * $ pulumi import azuredevops:index/project:Project example 00000000-0000-0000-0000-000000000000
 * ```
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
        return new Project(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuredevops:index/project:Project';

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
     * *
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Defines the status (`enabled`, `disabled`) of the project features. Valid features are `boards`, `repositories`, `pipelines`, `testplans`, `artifacts`
     *
     * | Features     | Possible Values   |
     * |--------------|-------------------|
     * | boards       | enabled, disabled |
     * | repositories | enabled, disabled |
     * | pipelines    | enabled, disabled |
     * | testplans    | enabled, disabled |
     * | artifacts    | enabled, disabled |
     *
     * > **NOTE:** It's possible to define project features both within the `azuredevops.ProjectFeatures` resource and
     * via the `features` block by using the `azuredevops.Project` resource.
     * However it's not possible to use both methods to manage features, since there'll be conflicts.
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
     * Specifies the version control system. Possbile values are: `Git` or `Tfvc`. Defaults to `Git`.
     */
    public readonly versionControl!: pulumi.Output<string | undefined>;
    /**
     * Specifies the visibility of the Project. Possible values are: `private` or `public`. Defaults to `private`.
     */
    public readonly visibility!: pulumi.Output<string | undefined>;
    /**
     * Specifies the work item template. Possible values are: `Agile`, `Basic`, `CMMI`, `Scrum` or a custom, pre-existing one. Defaults to `Agile`. An empty string will use the parent organization default.
     */
    public readonly workItemTemplate!: pulumi.Output<string | undefined>;

    /**
     * Create a Project resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ProjectArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProjectArgs | ProjectState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ProjectState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["features"] = state ? state.features : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["processTemplateId"] = state ? state.processTemplateId : undefined;
            resourceInputs["versionControl"] = state ? state.versionControl : undefined;
            resourceInputs["visibility"] = state ? state.visibility : undefined;
            resourceInputs["workItemTemplate"] = state ? state.workItemTemplate : undefined;
        } else {
            const args = argsOrState as ProjectArgs | undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["features"] = args ? args.features : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["versionControl"] = args ? args.versionControl : undefined;
            resourceInputs["visibility"] = args ? args.visibility : undefined;
            resourceInputs["workItemTemplate"] = args ? args.workItemTemplate : undefined;
            resourceInputs["processTemplateId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Project.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Project resources.
 */
export interface ProjectState {
    /**
     * The Description of the Project.
     * *
     */
    description?: pulumi.Input<string>;
    /**
     * Defines the status (`enabled`, `disabled`) of the project features. Valid features are `boards`, `repositories`, `pipelines`, `testplans`, `artifacts`
     *
     * | Features     | Possible Values   |
     * |--------------|-------------------|
     * | boards       | enabled, disabled |
     * | repositories | enabled, disabled |
     * | pipelines    | enabled, disabled |
     * | testplans    | enabled, disabled |
     * | artifacts    | enabled, disabled |
     *
     * > **NOTE:** It's possible to define project features both within the `azuredevops.ProjectFeatures` resource and
     * via the `features` block by using the `azuredevops.Project` resource.
     * However it's not possible to use both methods to manage features, since there'll be conflicts.
     */
    features?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The Project Name.
     */
    name?: pulumi.Input<string>;
    /**
     * The Process Template ID used by the Project.
     */
    processTemplateId?: pulumi.Input<string>;
    /**
     * Specifies the version control system. Possbile values are: `Git` or `Tfvc`. Defaults to `Git`.
     */
    versionControl?: pulumi.Input<string>;
    /**
     * Specifies the visibility of the Project. Possible values are: `private` or `public`. Defaults to `private`.
     */
    visibility?: pulumi.Input<string>;
    /**
     * Specifies the work item template. Possible values are: `Agile`, `Basic`, `CMMI`, `Scrum` or a custom, pre-existing one. Defaults to `Agile`. An empty string will use the parent organization default.
     */
    workItemTemplate?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Project resource.
 */
export interface ProjectArgs {
    /**
     * The Description of the Project.
     * *
     */
    description?: pulumi.Input<string>;
    /**
     * Defines the status (`enabled`, `disabled`) of the project features. Valid features are `boards`, `repositories`, `pipelines`, `testplans`, `artifacts`
     *
     * | Features     | Possible Values   |
     * |--------------|-------------------|
     * | boards       | enabled, disabled |
     * | repositories | enabled, disabled |
     * | pipelines    | enabled, disabled |
     * | testplans    | enabled, disabled |
     * | artifacts    | enabled, disabled |
     *
     * > **NOTE:** It's possible to define project features both within the `azuredevops.ProjectFeatures` resource and
     * via the `features` block by using the `azuredevops.Project` resource.
     * However it's not possible to use both methods to manage features, since there'll be conflicts.
     */
    features?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The Project Name.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the version control system. Possbile values are: `Git` or `Tfvc`. Defaults to `Git`.
     */
    versionControl?: pulumi.Input<string>;
    /**
     * Specifies the visibility of the Project. Possible values are: `private` or `public`. Defaults to `private`.
     */
    visibility?: pulumi.Input<string>;
    /**
     * Specifies the work item template. Possible values are: `Agile`, `Basic`, `CMMI`, `Scrum` or a custom, pre-existing one. Defaults to `Agile`. An empty string will use the parent organization default.
     */
    workItemTemplate?: pulumi.Input<string>;
}
