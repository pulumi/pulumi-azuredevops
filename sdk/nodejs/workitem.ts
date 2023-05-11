// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manages a Work Item in Azure Devops.
 *
 * ## Example Usage
 * ### Basic usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const exampleProject = new azuredevops.Project("exampleProject", {
 *     workItemTemplate: "Agile",
 *     versionControl: "Git",
 *     visibility: "private",
 *     description: "Managed by Terraform",
 * });
 * const exampleWorkitem = new azuredevops.Workitem("exampleWorkitem", {
 *     projectId: exampleProject.id,
 *     title: "Example Work Item",
 *     type: "Issue",
 *     state: "Active",
 *     tags: ["Tag"],
 * });
 * ```
 * ### With custom fields
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const exampleProject = new azuredevops.Project("exampleProject", {
 *     workItemTemplate: "Agile",
 *     versionControl: "Git",
 *     visibility: "private",
 *     description: "Managed by Terraform",
 * });
 * const exampleWorkitem = new azuredevops.Workitem("exampleWorkitem", {
 *     projectId: exampleProject.id,
 *     title: "Example Work Item",
 *     type: "Issue",
 *     state: "Active",
 *     tags: ["Tag"],
 *     customFields: {
 *         example: "example",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Work Item resource does not support import.
 */
export class Workitem extends pulumi.CustomResource {
    /**
     * Get an existing Workitem resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WorkitemState, opts?: pulumi.CustomResourceOptions): Workitem {
        return new Workitem(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuredevops:index/workitem:Workitem';

    /**
     * Returns true if the given object is an instance of Workitem.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Workitem {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Workitem.__pulumiType;
    }

    /**
     * Specifies a list with Custom Fields for the Work Item.
     */
    public readonly customFields!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The ID of the Project.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * The state of the Work Item. The four main states that are defined for the User Story (`Agile`) are `New`, `Active`, `Resolved`, and `Closed`. See [Workflow states](https://learn.microsoft.com/en-us/azure/devops/boards/work-items/workflow-and-state-categories?view=azure-devops&tabs=agile-process#workflow-states) for more details.
     */
    public readonly state!: pulumi.Output<string>;
    /**
     * Specifies a list of Tags.
     */
    public readonly tags!: pulumi.Output<string[] | undefined>;
    /**
     * The Title of the Work Item.
     */
    public readonly title!: pulumi.Output<string>;
    /**
     * The Type of the Work Item. The work item type varies depending on the process used when creating the project(`Agile`, `Basic`, `Scrum`, `Scrum`). See [Work Item Types](https://learn.microsoft.com/en-us/azure/devops/boards/work-items/about-work-items?view=azure-devops) for more details.
     */
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a Workitem resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WorkitemArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WorkitemArgs | WorkitemState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WorkitemState | undefined;
            resourceInputs["customFields"] = state ? state.customFields : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["state"] = state ? state.state : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["title"] = state ? state.title : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as WorkitemArgs | undefined;
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            if ((!args || args.title === undefined) && !opts.urn) {
                throw new Error("Missing required property 'title'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["customFields"] = args ? args.customFields : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["state"] = args ? args.state : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["title"] = args ? args.title : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Workitem.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Workitem resources.
 */
export interface WorkitemState {
    /**
     * Specifies a list with Custom Fields for the Work Item.
     */
    customFields?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The ID of the Project.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The state of the Work Item. The four main states that are defined for the User Story (`Agile`) are `New`, `Active`, `Resolved`, and `Closed`. See [Workflow states](https://learn.microsoft.com/en-us/azure/devops/boards/work-items/workflow-and-state-categories?view=azure-devops&tabs=agile-process#workflow-states) for more details.
     */
    state?: pulumi.Input<string>;
    /**
     * Specifies a list of Tags.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The Title of the Work Item.
     */
    title?: pulumi.Input<string>;
    /**
     * The Type of the Work Item. The work item type varies depending on the process used when creating the project(`Agile`, `Basic`, `Scrum`, `Scrum`). See [Work Item Types](https://learn.microsoft.com/en-us/azure/devops/boards/work-items/about-work-items?view=azure-devops) for more details.
     */
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Workitem resource.
 */
export interface WorkitemArgs {
    /**
     * Specifies a list with Custom Fields for the Work Item.
     */
    customFields?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The ID of the Project.
     */
    projectId: pulumi.Input<string>;
    /**
     * The state of the Work Item. The four main states that are defined for the User Story (`Agile`) are `New`, `Active`, `Resolved`, and `Closed`. See [Workflow states](https://learn.microsoft.com/en-us/azure/devops/boards/work-items/workflow-and-state-categories?view=azure-devops&tabs=agile-process#workflow-states) for more details.
     */
    state?: pulumi.Input<string>;
    /**
     * Specifies a list of Tags.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The Title of the Work Item.
     */
    title: pulumi.Input<string>;
    /**
     * The Type of the Work Item. The work item type varies depending on the process used when creating the project(`Agile`, `Basic`, `Scrum`, `Scrum`). See [Work Item Types](https://learn.microsoft.com/en-us/azure/devops/boards/work-items/about-work-items?view=azure-devops) for more details.
     */
    type: pulumi.Input<string>;
}