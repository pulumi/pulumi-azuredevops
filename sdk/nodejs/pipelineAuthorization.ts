// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manage pipeline access permissions to resources.
 *
 * > **Note** This resource is a replacement for `azuredevops.ResourceAuthorization`.  Pipeline authorizations managed by `azuredevops.ResourceAuthorization` can also be managed by this resource.
 *
 * > **Note** If both "All Pipeline Authorization" and "Custom Pipeline Authorization" are configured, "All Pipeline Authorization" has higher priority.
 *
 * ## Example Usage
 *
 * ### Authorization for all pipelines
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
 * });
 * const examplePool = new azuredevops.Pool("example", {
 *     name: "Example Pool",
 *     autoProvision: false,
 *     autoUpdate: false,
 * });
 * const exampleQueue = new azuredevops.Queue("example", {
 *     projectId: example.id,
 *     agentPoolId: examplePool.id,
 * });
 * const examplePipelineAuthorization = new azuredevops.PipelineAuthorization("example", {
 *     projectId: example.id,
 *     resourceId: exampleQueue.id,
 *     type: "queue",
 * });
 * ```
 *
 * ### Authorization for specific pipeline
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const exampleProject = new azuredevops.Project("example", {
 *     name: "Example Project",
 *     visibility: "private",
 *     versionControl: "Git",
 *     workItemTemplate: "Agile",
 *     description: "Managed by Pulumi",
 * });
 * const examplePool = new azuredevops.Pool("example", {
 *     name: "Example Pool",
 *     autoProvision: false,
 *     autoUpdate: false,
 * });
 * const exampleQueue = new azuredevops.Queue("example", {
 *     projectId: exampleProject.id,
 *     agentPoolId: examplePool.id,
 * });
 * const example = azuredevops.getGitRepositoryOutput({
 *     projectId: exampleProject.id,
 *     name: "Example Project",
 * });
 * const exampleBuildDefinition = new azuredevops.BuildDefinition("example", {
 *     projectId: exampleProject.id,
 *     name: "Example Pipeline",
 *     repository: {
 *         repoType: "TfsGit",
 *         repoId: example.apply(example => example.id),
 *         ymlPath: "azure-pipelines.yml",
 *     },
 * });
 * const examplePipelineAuthorization = new azuredevops.PipelineAuthorization("example", {
 *     projectId: exampleProject.id,
 *     resourceId: exampleQueue.id,
 *     type: "queue",
 *     pipelineId: exampleBuildDefinition.id,
 * });
 * ```
 *
 * ## Relevant Links
 *
 * - [Azure DevOps Service REST API 7.1 - Pipeline Permissions](https://learn.microsoft.com/en-us/rest/api/azure/devops/approvalsandchecks/pipeline-permissions?view=azure-devops-rest-7.1)
 */
export class PipelineAuthorization extends pulumi.CustomResource {
    /**
     * Get an existing PipelineAuthorization resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PipelineAuthorizationState, opts?: pulumi.CustomResourceOptions): PipelineAuthorization {
        return new PipelineAuthorization(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuredevops:index/pipelineAuthorization:PipelineAuthorization';

    /**
     * Returns true if the given object is an instance of PipelineAuthorization.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PipelineAuthorization {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PipelineAuthorization.__pulumiType;
    }

    /**
     * The ID of the pipeline. If not configured, all pipelines will be authorized. Changing this forces a new resource to be created.
     */
    public readonly pipelineId!: pulumi.Output<number | undefined>;
    /**
     * The ID of the project where the pipeline exists. Defaults to `projectId` if not specified. Changing this forces a new resource to be created
     */
    public readonly pipelineProjectId!: pulumi.Output<string | undefined>;
    /**
     * The  ID of the project. Changing this forces a new resource to be created
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * The ID of the resource to authorize. Changing this forces a new resource to be created
     */
    public readonly resourceId!: pulumi.Output<string>;
    /**
     * The type of the resource to authorize. Possible values are: `endpoint`, `queue`, `variablegroup`, `environment`, `repository`. Changing this forces a new resource to be created
     *
     * > **Note** `repository` is for AzureDevOps repository. To authorize repository other than
     * Azure DevOps like GitHub you need to use service connection(`endpoint`)  to connect and authorize.
     * Typical process for connecting to GitHub:
     * **Pipeline  <----> Service Connection(`endpoint`) <----> GitHub Repository**
     */
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a PipelineAuthorization resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PipelineAuthorizationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PipelineAuthorizationArgs | PipelineAuthorizationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PipelineAuthorizationState | undefined;
            resourceInputs["pipelineId"] = state ? state.pipelineId : undefined;
            resourceInputs["pipelineProjectId"] = state ? state.pipelineProjectId : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["resourceId"] = state ? state.resourceId : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as PipelineAuthorizationArgs | undefined;
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            if ((!args || args.resourceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceId'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["pipelineId"] = args ? args.pipelineId : undefined;
            resourceInputs["pipelineProjectId"] = args ? args.pipelineProjectId : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["resourceId"] = args ? args.resourceId : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(PipelineAuthorization.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PipelineAuthorization resources.
 */
export interface PipelineAuthorizationState {
    /**
     * The ID of the pipeline. If not configured, all pipelines will be authorized. Changing this forces a new resource to be created.
     */
    pipelineId?: pulumi.Input<number>;
    /**
     * The ID of the project where the pipeline exists. Defaults to `projectId` if not specified. Changing this forces a new resource to be created
     */
    pipelineProjectId?: pulumi.Input<string>;
    /**
     * The  ID of the project. Changing this forces a new resource to be created
     */
    projectId?: pulumi.Input<string>;
    /**
     * The ID of the resource to authorize. Changing this forces a new resource to be created
     */
    resourceId?: pulumi.Input<string>;
    /**
     * The type of the resource to authorize. Possible values are: `endpoint`, `queue`, `variablegroup`, `environment`, `repository`. Changing this forces a new resource to be created
     *
     * > **Note** `repository` is for AzureDevOps repository. To authorize repository other than
     * Azure DevOps like GitHub you need to use service connection(`endpoint`)  to connect and authorize.
     * Typical process for connecting to GitHub:
     * **Pipeline  <----> Service Connection(`endpoint`) <----> GitHub Repository**
     */
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a PipelineAuthorization resource.
 */
export interface PipelineAuthorizationArgs {
    /**
     * The ID of the pipeline. If not configured, all pipelines will be authorized. Changing this forces a new resource to be created.
     */
    pipelineId?: pulumi.Input<number>;
    /**
     * The ID of the project where the pipeline exists. Defaults to `projectId` if not specified. Changing this forces a new resource to be created
     */
    pipelineProjectId?: pulumi.Input<string>;
    /**
     * The  ID of the project. Changing this forces a new resource to be created
     */
    projectId: pulumi.Input<string>;
    /**
     * The ID of the resource to authorize. Changing this forces a new resource to be created
     */
    resourceId: pulumi.Input<string>;
    /**
     * The type of the resource to authorize. Possible values are: `endpoint`, `queue`, `variablegroup`, `environment`, `repository`. Changing this forces a new resource to be created
     *
     * > **Note** `repository` is for AzureDevOps repository. To authorize repository other than
     * Azure DevOps like GitHub you need to use service connection(`endpoint`)  to connect and authorize.
     * Typical process for connecting to GitHub:
     * **Pipeline  <----> Service Connection(`endpoint`) <----> GitHub Repository**
     */
    type: pulumi.Input<string>;
}
