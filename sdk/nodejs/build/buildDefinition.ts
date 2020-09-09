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
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const project = new azuredevops.Project("project", {
 *     projectName: "Sample Project",
 *     visibility: "private",
 *     versionControl: "Git",
 *     workItemTemplate: "Agile",
 * });
 * const repository = new azuredevops.Git("repository", {
 *     projectId: project.id,
 *     initialization: {
 *         initType: "Clean",
 *     },
 * });
 * const vars = new azuredevops.VariableGroup("vars", {
 *     projectId: project.id,
 *     description: "Managed by Terraform",
 *     allowAccess: true,
 *     variables: [{
 *         name: "FOO",
 *         value: "BAR",
 *     }],
 * });
 * const build = new azuredevops.BuildDefinition("build", {
 *     projectId: project.id,
 *     path: "\\ExampleFolder",
 *     ciTrigger: {
 *         useYaml: true,
 *     },
 *     repository: {
 *         repoType: "TfsGit",
 *         repoId: repository.id,
 *         branchName: repository.defaultBranch,
 *         ymlPath: "azure-pipelines.yml",
 *     },
 *     variableGroups: [vars.id],
 *     variables: [
 *         {
 *             name: "PipelineVariable",
 *             value: "Go Microsoft!",
 *         },
 *         {
 *             name: "PipelineSecret",
 *             secretValue: "ZGV2cw",
 *             isSecret: true,
 *         },
 *     ],
 * });
 * ```
 * ## Relevant Links
 *
 * * [Azure DevOps Service REST API 5.1 - Build Definitions](https://docs.microsoft.com/en-us/rest/api/azure/devops/build/definitions?view=azure-devops-rest-5.1)
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
     * The agent pool that should execute the build. Defaults to `Hosted Ubuntu 1604`.
     */
    public readonly agentPoolName!: pulumi.Output<string | undefined>;
    /**
     * Continuous Integration Integration trigger.
     */
    public readonly ciTrigger!: pulumi.Output<outputs.Build.BuildDefinitionCiTrigger | undefined>;
    /**
     * The name of the build definition.
     */
    public readonly name!: pulumi.Output<string>;
    public readonly path!: pulumi.Output<string | undefined>;
    /**
     * The project ID or project name.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * Pull Request Integration Integration trigger.
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
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as BuildDefinitionState | undefined;
            inputs["agentPoolName"] = state ? state.agentPoolName : undefined;
            inputs["ciTrigger"] = state ? state.ciTrigger : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["path"] = state ? state.path : undefined;
            inputs["projectId"] = state ? state.projectId : undefined;
            inputs["pullRequestTrigger"] = state ? state.pullRequestTrigger : undefined;
            inputs["repository"] = state ? state.repository : undefined;
            inputs["revision"] = state ? state.revision : undefined;
            inputs["variableGroups"] = state ? state.variableGroups : undefined;
            inputs["variables"] = state ? state.variables : undefined;
        } else {
            const args = argsOrState as BuildDefinitionArgs | undefined;
            if (!args || args.projectId === undefined) {
                throw new Error("Missing required property 'projectId'");
            }
            if (!args || args.repository === undefined) {
                throw new Error("Missing required property 'repository'");
            }
            inputs["agentPoolName"] = args ? args.agentPoolName : undefined;
            inputs["ciTrigger"] = args ? args.ciTrigger : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["path"] = args ? args.path : undefined;
            inputs["projectId"] = args ? args.projectId : undefined;
            inputs["pullRequestTrigger"] = args ? args.pullRequestTrigger : undefined;
            inputs["repository"] = args ? args.repository : undefined;
            inputs["variableGroups"] = args ? args.variableGroups : undefined;
            inputs["variables"] = args ? args.variables : undefined;
            inputs["revision"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(BuildDefinition.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BuildDefinition resources.
 */
export interface BuildDefinitionState {
    /**
     * The agent pool that should execute the build. Defaults to `Hosted Ubuntu 1604`.
     */
    readonly agentPoolName?: pulumi.Input<string>;
    /**
     * Continuous Integration Integration trigger.
     */
    readonly ciTrigger?: pulumi.Input<inputs.Build.BuildDefinitionCiTrigger>;
    /**
     * The name of the build definition.
     */
    readonly name?: pulumi.Input<string>;
    readonly path?: pulumi.Input<string>;
    /**
     * The project ID or project name.
     */
    readonly projectId?: pulumi.Input<string>;
    /**
     * Pull Request Integration Integration trigger.
     */
    readonly pullRequestTrigger?: pulumi.Input<inputs.Build.BuildDefinitionPullRequestTrigger>;
    /**
     * A `repository` block as documented below.
     */
    readonly repository?: pulumi.Input<inputs.Build.BuildDefinitionRepository>;
    /**
     * The revision of the build definition
     */
    readonly revision?: pulumi.Input<number>;
    /**
     * A list of variable group IDs (integers) to link to the build definition.
     */
    readonly variableGroups?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * A list of `variable` blocks, as documented below.
     */
    readonly variables?: pulumi.Input<pulumi.Input<inputs.Build.BuildDefinitionVariable>[]>;
}

/**
 * The set of arguments for constructing a BuildDefinition resource.
 */
export interface BuildDefinitionArgs {
    /**
     * The agent pool that should execute the build. Defaults to `Hosted Ubuntu 1604`.
     */
    readonly agentPoolName?: pulumi.Input<string>;
    /**
     * Continuous Integration Integration trigger.
     */
    readonly ciTrigger?: pulumi.Input<inputs.Build.BuildDefinitionCiTrigger>;
    /**
     * The name of the build definition.
     */
    readonly name?: pulumi.Input<string>;
    readonly path?: pulumi.Input<string>;
    /**
     * The project ID or project name.
     */
    readonly projectId: pulumi.Input<string>;
    /**
     * Pull Request Integration Integration trigger.
     */
    readonly pullRequestTrigger?: pulumi.Input<inputs.Build.BuildDefinitionPullRequestTrigger>;
    /**
     * A `repository` block as documented below.
     */
    readonly repository: pulumi.Input<inputs.Build.BuildDefinitionRepository>;
    /**
     * A list of variable group IDs (integers) to link to the build definition.
     */
    readonly variableGroups?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * A list of `variable` blocks, as documented below.
     */
    readonly variables?: pulumi.Input<pulumi.Input<inputs.Build.BuildDefinitionVariable>[]>;
}
