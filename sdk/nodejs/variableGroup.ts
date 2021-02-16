// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Manages variable groups within Azure DevOps.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const project = new azuredevops.Project("project", {});
 * const variablegroup = new azuredevops.VariableGroup("variablegroup", {
 *     projectId: project.id,
 *     description: "Test Variable Group Description",
 *     allowAccess: true,
 *     variables: [
 *         {
 *             name: "key",
 *             value: "value",
 *         },
 *         {
 *             name: "Account Password",
 *             secretValue: "p@ssword123",
 *             isSecret: true,
 *         },
 *     ],
 * });
 * ```
 * ## Relevant Links
 *
 * - [Azure DevOps Service REST API 5.1 - Variable Groups](https://docs.microsoft.com/en-us/rest/api/azure/devops/distributedtask/variablegroups?view=azure-devops-rest-5.1)
 * - [Azure DevOps Service REST API 5.1 - Authorized Resources](https://docs.microsoft.com/en-us/rest/api/azure/devops/build/authorizedresources?view=azure-devops-rest-5.1)
 *
 * ## PAT Permissions Required
 *
 * - **Variable Groups**: Read, Create, & Manage
 *
 * ## Import
 *
 * **Variable groups containing secret values cannot be imported.** Azure DevOps Variable groups can be imported using the project name/variable group ID or by the project Guid/variable group ID, e.g.
 *
 * ```sh
 *  $ pulumi import azuredevops:index/variableGroup:VariableGroup variablegroup "Test Project/10"
 * ```
 *
 *  or
 *
 * ```sh
 *  $ pulumi import azuredevops:index/variableGroup:VariableGroup variablegroup 00000000-0000-0000-0000-000000000000/0
 * ```
 *
 *  _Note that for secret variables, the import command retrieve blank value in the tfstate._
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
        return new VariableGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuredevops:index/variableGroup:VariableGroup';

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
    public readonly keyVault!: pulumi.Output<outputs.VariableGroupKeyVault | undefined>;
    /**
     * The name of the Variable Group.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The project ID or project name.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * One or more `variable` blocks as documented below.
     */
    public readonly variables!: pulumi.Output<outputs.VariableGroupVariable[]>;

    /**
     * Create a VariableGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VariableGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VariableGroupArgs | VariableGroupState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VariableGroupState | undefined;
            inputs["allowAccess"] = state ? state.allowAccess : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["keyVault"] = state ? state.keyVault : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["projectId"] = state ? state.projectId : undefined;
            inputs["variables"] = state ? state.variables : undefined;
        } else {
            const args = argsOrState as VariableGroupArgs | undefined;
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            if ((!args || args.variables === undefined) && !opts.urn) {
                throw new Error("Missing required property 'variables'");
            }
            inputs["allowAccess"] = args ? args.allowAccess : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["keyVault"] = args ? args.keyVault : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["projectId"] = args ? args.projectId : undefined;
            inputs["variables"] = args ? args.variables : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        const aliasOpts = { aliases: [{ type: "azuredevops:Pipeline/variableGroup:VariableGroup" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(VariableGroup.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VariableGroup resources.
 */
export interface VariableGroupState {
    /**
     * Boolean that indicate if this variable group is shared by all pipelines of this project.
     */
    readonly allowAccess?: pulumi.Input<boolean>;
    /**
     * The description of the Variable Group.
     */
    readonly description?: pulumi.Input<string>;
    readonly keyVault?: pulumi.Input<inputs.VariableGroupKeyVault>;
    /**
     * The name of the Variable Group.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The project ID or project name.
     */
    readonly projectId?: pulumi.Input<string>;
    /**
     * One or more `variable` blocks as documented below.
     */
    readonly variables?: pulumi.Input<pulumi.Input<inputs.VariableGroupVariable>[]>;
}

/**
 * The set of arguments for constructing a VariableGroup resource.
 */
export interface VariableGroupArgs {
    /**
     * Boolean that indicate if this variable group is shared by all pipelines of this project.
     */
    readonly allowAccess?: pulumi.Input<boolean>;
    /**
     * The description of the Variable Group.
     */
    readonly description?: pulumi.Input<string>;
    readonly keyVault?: pulumi.Input<inputs.VariableGroupKeyVault>;
    /**
     * The name of the Variable Group.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The project ID or project name.
     */
    readonly projectId: pulumi.Input<string>;
    /**
     * One or more `variable` blocks as documented below.
     */
    readonly variables: pulumi.Input<pulumi.Input<inputs.VariableGroupVariable>[]>;
}
