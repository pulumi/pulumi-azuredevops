// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * ## Example Usage
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
 * const exampleVariableGroup = new azuredevops.VariableGroup("exampleVariableGroup", {
 *     projectId: exampleProject.id,
 *     description: "Example Variable Group Description",
 *     allowAccess: true,
 *     variables: [
 *         {
 *             name: "key1",
 *             value: "val1",
 *         },
 *         {
 *             name: "key2",
 *             secretValue: "val2",
 *             isSecret: true,
 *         },
 *     ],
 * });
 * ```
 * ### With AzureRM Key Vault
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
 * const exampleServiceEndpointAzureRM = new azuredevops.ServiceEndpointAzureRM("exampleServiceEndpointAzureRM", {
 *     projectId: exampleProject.id,
 *     serviceEndpointName: "Example AzureRM",
 *     description: "Managed by Terraform",
 *     credentials: {
 *         serviceprincipalid: "00000000-0000-0000-0000-000000000000",
 *         serviceprincipalkey: "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
 *     },
 *     azurermSpnTenantid: "00000000-0000-0000-0000-000000000000",
 *     azurermSubscriptionId: "00000000-0000-0000-0000-000000000000",
 *     azurermSubscriptionName: "Example Subscription Name",
 * });
 * const exampleVariableGroup = new azuredevops.VariableGroup("exampleVariableGroup", {
 *     projectId: exampleProject.id,
 *     description: "Example Variable Group Description",
 *     allowAccess: true,
 *     keyVault: {
 *         name: "example-kv",
 *         serviceEndpointId: exampleServiceEndpointAzureRM.id,
 *     },
 *     variables: [
 *         {
 *             name: "key1",
 *         },
 *         {
 *             name: "key2",
 *         },
 *     ],
 * });
 * ```
 * ## Relevant Links
 *
 * - [Azure DevOps Service REST API 6.0 - Variable Groups](https://docs.microsoft.com/en-us/rest/api/azure/devops/distributedtask/variablegroups?view=azure-devops-rest-6.0)
 * - [Azure DevOps Service REST API 6.0 - Authorized Resources](https://docs.microsoft.com/en-us/rest/api/azure/devops/build/authorizedresources?view=azure-devops-rest-6.0)
 *
 * ## PAT Permissions Required
 *
 * - **Variable Groups**: Read, Create, & Manage
 * - **Build**: Read & execute
 * - **Project and Team**: Read
 * - **Token Administration**: Read & manage
 * - **Tokens**: Read & manage
 * - **Work Items**: Read
 *
 * ## Import
 *
 * **Variable groups containing secret values cannot be imported.** Azure DevOps Variable groups can be imported using the project name/variable group ID or by the project Guid/variable group ID, e.g.
 *
 * ```sh
 *  $ pulumi import azuredevops:Pipeline/variableGroup:VariableGroup example "Example Project/10"
 * ```
 *
 *  or
 *
 * ```sh
 *  $ pulumi import azuredevops:Pipeline/variableGroup:VariableGroup example 00000000-0000-0000-0000-000000000000/0
 * ```
 *
 *  _Note that for secret variables, the import command retrieve blank value in the tfstate._
 *
 * @deprecated azuredevops.pipeline.VariableGroup has been deprecated in favor of azuredevops.VariableGroup
 */
export class VariableGroup extends pulumi.CustomResource {
    /**
     * Get an existing VariableGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VariableGroupState, opts?: pulumi.CustomResourceOptions): VariableGroup {
        pulumi.log.warn("VariableGroup is deprecated: azuredevops.pipeline.VariableGroup has been deprecated in favor of azuredevops.VariableGroup")
        return new VariableGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuredevops:Pipeline/variableGroup:VariableGroup';

    /**
     * Returns true if the given object is an instance of VariableGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VariableGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VariableGroup.__pulumiType;
    }

    /**
     * Boolean that indicate if this variable group is shared by all pipelines of this project.
     */
    public readonly allowAccess!: pulumi.Output<boolean | undefined>;
    /**
     * The description of the Variable Group.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * A list of `keyVault` blocks as documented below.
     */
    public readonly keyVault!: pulumi.Output<outputs.Pipeline.VariableGroupKeyVault | undefined>;
    /**
     * The name of the Variable Group.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the project.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * One or more `variable` blocks as documented below.
     */
    public readonly variables!: pulumi.Output<outputs.Pipeline.VariableGroupVariable[]>;

    /**
     * Create a VariableGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated azuredevops.pipeline.VariableGroup has been deprecated in favor of azuredevops.VariableGroup */
    constructor(name: string, args: VariableGroupArgs, opts?: pulumi.CustomResourceOptions)
    /** @deprecated azuredevops.pipeline.VariableGroup has been deprecated in favor of azuredevops.VariableGroup */
    constructor(name: string, argsOrState?: VariableGroupArgs | VariableGroupState, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("VariableGroup is deprecated: azuredevops.pipeline.VariableGroup has been deprecated in favor of azuredevops.VariableGroup")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VariableGroupState | undefined;
            resourceInputs["allowAccess"] = state ? state.allowAccess : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["keyVault"] = state ? state.keyVault : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["variables"] = state ? state.variables : undefined;
        } else {
            const args = argsOrState as VariableGroupArgs | undefined;
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            if ((!args || args.variables === undefined) && !opts.urn) {
                throw new Error("Missing required property 'variables'");
            }
            resourceInputs["allowAccess"] = args ? args.allowAccess : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["keyVault"] = args ? args.keyVault : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["variables"] = args ? args.variables : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(VariableGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VariableGroup resources.
 */
export interface VariableGroupState {
    /**
     * Boolean that indicate if this variable group is shared by all pipelines of this project.
     */
    allowAccess?: pulumi.Input<boolean>;
    /**
     * The description of the Variable Group.
     */
    description?: pulumi.Input<string>;
    /**
     * A list of `keyVault` blocks as documented below.
     */
    keyVault?: pulumi.Input<inputs.Pipeline.VariableGroupKeyVault>;
    /**
     * The name of the Variable Group.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project.
     */
    projectId?: pulumi.Input<string>;
    /**
     * One or more `variable` blocks as documented below.
     */
    variables?: pulumi.Input<pulumi.Input<inputs.Pipeline.VariableGroupVariable>[]>;
}

/**
 * The set of arguments for constructing a VariableGroup resource.
 */
export interface VariableGroupArgs {
    /**
     * Boolean that indicate if this variable group is shared by all pipelines of this project.
     */
    allowAccess?: pulumi.Input<boolean>;
    /**
     * The description of the Variable Group.
     */
    description?: pulumi.Input<string>;
    /**
     * A list of `keyVault` blocks as documented below.
     */
    keyVault?: pulumi.Input<inputs.Pipeline.VariableGroupKeyVault>;
    /**
     * The name of the Variable Group.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project.
     */
    projectId: pulumi.Input<string>;
    /**
     * One or more `variable` blocks as documented below.
     */
    variables: pulumi.Input<pulumi.Input<inputs.Pipeline.VariableGroupVariable>[]>;
}
