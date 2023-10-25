// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages a Build Definition within Azure DevOps.
 *
 * ## Example Usage
 * ## Remarks
 *
 * The path attribute can not end in `\` unless the path is the root value of `\`.
 *
 * Valid path values (yaml encoded) include:
 * - `\\`
 * - `\\ExampleFolder`
 * - `\\Nested\\Example Folder`
 *
 * The value of `\\ExampleFolder\\` would be invalid.
 *
 * ## Relevant Links
 *
 * - [Azure DevOps Service REST API 7.0 - Build Definitions](https://docs.microsoft.com/en-us/rest/api/azure/devops/build/definitions?view=azure-devops-rest-7.0)
 *
 * ## Import
 *
 * Azure DevOps Build Definitions can be imported using the project name/definitions Id or by the project Guid/definitions Id, e.g.
 *
 * ```sh
 *  $ pulumi import azuredevops:Build/buildDefinition:BuildDefinition example "Example Project"/10
 * ```
 *
 *  or
 *
 * ```sh
 *  $ pulumi import azuredevops:Build/buildDefinition:BuildDefinition example 00000000-0000-0000-0000-000000000000/0
 * ```
 *
 * @deprecated azuredevops.build.BuildDefinition has been deprecated in favor of azuredevops.BuildDefinition
 */
export class BuildDefinition extends pulumi.CustomResource {
    /**
     * Get an existing BuildDefinition resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BuildDefinitionState, opts?: pulumi.CustomResourceOptions): BuildDefinition {
        pulumi.log.warn("BuildDefinition is deprecated: azuredevops.build.BuildDefinition has been deprecated in favor of azuredevops.BuildDefinition")
        return new BuildDefinition(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuredevops:Build/buildDefinition:BuildDefinition';

    /**
     * Returns true if the given object is an instance of BuildDefinition.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BuildDefinition {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BuildDefinition.__pulumiType;
    }

    /**
     * The agent pool that should execute the build. Defaults to `Azure Pipelines`.
     */
    public readonly agentPoolName!: pulumi.Output<string | undefined>;
    /**
     * Continuous Integration trigger.
     */
    public readonly ciTrigger!: pulumi.Output<outputs.Build.BuildDefinitionCiTrigger | undefined>;
    /**
     * A `features` blocks as documented below.
     */
    public readonly features!: pulumi.Output<outputs.Build.BuildDefinitionFeature[] | undefined>;
    /**
     * The name of the build definition.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The folder path of the build definition.
     */
    public readonly path!: pulumi.Output<string | undefined>;
    /**
     * The project ID or project name.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * Pull Request Integration trigger.
     */
    public readonly pullRequestTrigger!: pulumi.Output<outputs.Build.BuildDefinitionPullRequestTrigger | undefined>;
    /**
     * A `repository` block as documented below.
     */
    public readonly repository!: pulumi.Output<outputs.Build.BuildDefinitionRepository>;
    /**
     * The revision of the build definition
     */
    public /*out*/ readonly revision!: pulumi.Output<number>;
    public readonly schedules!: pulumi.Output<outputs.Build.BuildDefinitionSchedule[] | undefined>;
    /**
     * A list of variable group IDs (integers) to link to the build definition.
     */
    public readonly variableGroups!: pulumi.Output<number[] | undefined>;
    /**
     * A list of `variable` blocks, as documented below.
     */
    public readonly variables!: pulumi.Output<outputs.Build.BuildDefinitionVariable[] | undefined>;

    /**
     * Create a BuildDefinition resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated azuredevops.build.BuildDefinition has been deprecated in favor of azuredevops.BuildDefinition */
    constructor(name: string, args: BuildDefinitionArgs, opts?: pulumi.CustomResourceOptions)
    /** @deprecated azuredevops.build.BuildDefinition has been deprecated in favor of azuredevops.BuildDefinition */
    constructor(name: string, argsOrState?: BuildDefinitionArgs | BuildDefinitionState, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("BuildDefinition is deprecated: azuredevops.build.BuildDefinition has been deprecated in favor of azuredevops.BuildDefinition")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BuildDefinitionState | undefined;
            resourceInputs["agentPoolName"] = state ? state.agentPoolName : undefined;
            resourceInputs["ciTrigger"] = state ? state.ciTrigger : undefined;
            resourceInputs["features"] = state ? state.features : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["path"] = state ? state.path : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["pullRequestTrigger"] = state ? state.pullRequestTrigger : undefined;
            resourceInputs["repository"] = state ? state.repository : undefined;
            resourceInputs["revision"] = state ? state.revision : undefined;
            resourceInputs["schedules"] = state ? state.schedules : undefined;
            resourceInputs["variableGroups"] = state ? state.variableGroups : undefined;
            resourceInputs["variables"] = state ? state.variables : undefined;
        } else {
            const args = argsOrState as BuildDefinitionArgs | undefined;
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            if ((!args || args.repository === undefined) && !opts.urn) {
                throw new Error("Missing required property 'repository'");
            }
            resourceInputs["agentPoolName"] = args ? args.agentPoolName : undefined;
            resourceInputs["ciTrigger"] = args ? args.ciTrigger : undefined;
            resourceInputs["features"] = args ? args.features : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["path"] = args ? args.path : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["pullRequestTrigger"] = args ? args.pullRequestTrigger : undefined;
            resourceInputs["repository"] = args ? args.repository : undefined;
            resourceInputs["schedules"] = args ? args.schedules : undefined;
            resourceInputs["variableGroups"] = args ? args.variableGroups : undefined;
            resourceInputs["variables"] = args ? args.variables : undefined;
            resourceInputs["revision"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(BuildDefinition.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BuildDefinition resources.
 */
export interface BuildDefinitionState {
    /**
     * The agent pool that should execute the build. Defaults to `Azure Pipelines`.
     */
    agentPoolName?: pulumi.Input<string>;
    /**
     * Continuous Integration trigger.
     */
    ciTrigger?: pulumi.Input<inputs.Build.BuildDefinitionCiTrigger>;
    /**
     * A `features` blocks as documented below.
     */
    features?: pulumi.Input<pulumi.Input<inputs.Build.BuildDefinitionFeature>[]>;
    /**
     * The name of the build definition.
     */
    name?: pulumi.Input<string>;
    /**
     * The folder path of the build definition.
     */
    path?: pulumi.Input<string>;
    /**
     * The project ID or project name.
     */
    projectId?: pulumi.Input<string>;
    /**
     * Pull Request Integration trigger.
     */
    pullRequestTrigger?: pulumi.Input<inputs.Build.BuildDefinitionPullRequestTrigger>;
    /**
     * A `repository` block as documented below.
     */
    repository?: pulumi.Input<inputs.Build.BuildDefinitionRepository>;
    /**
     * The revision of the build definition
     */
    revision?: pulumi.Input<number>;
    schedules?: pulumi.Input<pulumi.Input<inputs.Build.BuildDefinitionSchedule>[]>;
    /**
     * A list of variable group IDs (integers) to link to the build definition.
     */
    variableGroups?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * A list of `variable` blocks, as documented below.
     */
    variables?: pulumi.Input<pulumi.Input<inputs.Build.BuildDefinitionVariable>[]>;
}

/**
 * The set of arguments for constructing a BuildDefinition resource.
 */
export interface BuildDefinitionArgs {
    /**
     * The agent pool that should execute the build. Defaults to `Azure Pipelines`.
     */
    agentPoolName?: pulumi.Input<string>;
    /**
     * Continuous Integration trigger.
     */
    ciTrigger?: pulumi.Input<inputs.Build.BuildDefinitionCiTrigger>;
    /**
     * A `features` blocks as documented below.
     */
    features?: pulumi.Input<pulumi.Input<inputs.Build.BuildDefinitionFeature>[]>;
    /**
     * The name of the build definition.
     */
    name?: pulumi.Input<string>;
    /**
     * The folder path of the build definition.
     */
    path?: pulumi.Input<string>;
    /**
     * The project ID or project name.
     */
    projectId: pulumi.Input<string>;
    /**
     * Pull Request Integration trigger.
     */
    pullRequestTrigger?: pulumi.Input<inputs.Build.BuildDefinitionPullRequestTrigger>;
    /**
     * A `repository` block as documented below.
     */
    repository: pulumi.Input<inputs.Build.BuildDefinitionRepository>;
    schedules?: pulumi.Input<pulumi.Input<inputs.Build.BuildDefinitionSchedule>[]>;
    /**
     * A list of variable group IDs (integers) to link to the build definition.
     */
    variableGroups?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * A list of `variable` blocks, as documented below.
     */
    variables?: pulumi.Input<pulumi.Input<inputs.Build.BuildDefinitionVariable>[]>;
}
