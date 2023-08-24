// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manages a generic service endpoint within Azure DevOps, which can be used to authenticate to any external git service
 * using basic authentication via a username and password. This is mostly useful for importing private git repositories.
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
 * const exampleServiceEndpointGenericGit = new azuredevops.ServiceEndpointGenericGit("exampleServiceEndpointGenericGit", {
 *     projectId: exampleProject.id,
 *     repositoryUrl: "https://dev.azure.com/org/project/_git/repository",
 *     username: "username",
 *     password: "password",
 *     serviceEndpointName: "Example Generic Git",
 *     description: "Managed by Terraform",
 * });
 * ```
 * ## Relevant Links
 *
 * - [Azure DevOps Service REST API 6.0 - Agent Pools](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-6.0)
 *
 * ## Import
 *
 * Azure DevOps Service Endpoint Generic Git can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**
 *
 * ```sh
 *  $ pulumi import azuredevops:index/serviceEndpointGenericGit:ServiceEndpointGenericGit example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
 * ```
 */
export class ServiceEndpointGenericGit extends pulumi.CustomResource {
    /**
     * Get an existing ServiceEndpointGenericGit resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServiceEndpointGenericGitState, opts?: pulumi.CustomResourceOptions): ServiceEndpointGenericGit {
        return new ServiceEndpointGenericGit(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuredevops:index/serviceEndpointGenericGit:ServiceEndpointGenericGit';

    /**
     * Returns true if the given object is an instance of ServiceEndpointGenericGit.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServiceEndpointGenericGit {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServiceEndpointGenericGit.__pulumiType;
    }

    public readonly authorization!: pulumi.Output<{[key: string]: string}>;
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * A value indicating whether or not to attempt accessing this git server from Azure Pipelines.
     */
    public readonly enablePipelinesAccess!: pulumi.Output<boolean | undefined>;
    /**
     * The PAT or password used to authenticate to the git repository.
     *
     * > **Note** For AzureDevOps Git, PAT should be used as the password.
     */
    public readonly password!: pulumi.Output<string | undefined>;
    /**
     * The ID of the project.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * The URL of the repository associated with the service endpoint.
     */
    public readonly repositoryUrl!: pulumi.Output<string>;
    /**
     * The name of the service endpoint.
     */
    public readonly serviceEndpointName!: pulumi.Output<string>;
    /**
     * The username used to authenticate to the git repository.
     */
    public readonly username!: pulumi.Output<string | undefined>;

    /**
     * Create a ServiceEndpointGenericGit resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServiceEndpointGenericGitArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServiceEndpointGenericGitArgs | ServiceEndpointGenericGitState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ServiceEndpointGenericGitState | undefined;
            resourceInputs["authorization"] = state ? state.authorization : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["enablePipelinesAccess"] = state ? state.enablePipelinesAccess : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["repositoryUrl"] = state ? state.repositoryUrl : undefined;
            resourceInputs["serviceEndpointName"] = state ? state.serviceEndpointName : undefined;
            resourceInputs["username"] = state ? state.username : undefined;
        } else {
            const args = argsOrState as ServiceEndpointGenericGitArgs | undefined;
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            if ((!args || args.repositoryUrl === undefined) && !opts.urn) {
                throw new Error("Missing required property 'repositoryUrl'");
            }
            if ((!args || args.serviceEndpointName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceEndpointName'");
            }
            resourceInputs["authorization"] = args ? args.authorization : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["enablePipelinesAccess"] = args ? args.enablePipelinesAccess : undefined;
            resourceInputs["password"] = args?.password ? pulumi.secret(args.password) : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["repositoryUrl"] = args ? args.repositoryUrl : undefined;
            resourceInputs["serviceEndpointName"] = args ? args.serviceEndpointName : undefined;
            resourceInputs["username"] = args ? args.username : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["password"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(ServiceEndpointGenericGit.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ServiceEndpointGenericGit resources.
 */
export interface ServiceEndpointGenericGitState {
    authorization?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    description?: pulumi.Input<string>;
    /**
     * A value indicating whether or not to attempt accessing this git server from Azure Pipelines.
     */
    enablePipelinesAccess?: pulumi.Input<boolean>;
    /**
     * The PAT or password used to authenticate to the git repository.
     *
     * > **Note** For AzureDevOps Git, PAT should be used as the password.
     */
    password?: pulumi.Input<string>;
    /**
     * The ID of the project.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The URL of the repository associated with the service endpoint.
     */
    repositoryUrl?: pulumi.Input<string>;
    /**
     * The name of the service endpoint.
     */
    serviceEndpointName?: pulumi.Input<string>;
    /**
     * The username used to authenticate to the git repository.
     */
    username?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ServiceEndpointGenericGit resource.
 */
export interface ServiceEndpointGenericGitArgs {
    authorization?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    description?: pulumi.Input<string>;
    /**
     * A value indicating whether or not to attempt accessing this git server from Azure Pipelines.
     */
    enablePipelinesAccess?: pulumi.Input<boolean>;
    /**
     * The PAT or password used to authenticate to the git repository.
     *
     * > **Note** For AzureDevOps Git, PAT should be used as the password.
     */
    password?: pulumi.Input<string>;
    /**
     * The ID of the project.
     */
    projectId: pulumi.Input<string>;
    /**
     * The URL of the repository associated with the service endpoint.
     */
    repositoryUrl: pulumi.Input<string>;
    /**
     * The name of the service endpoint.
     */
    serviceEndpointName: pulumi.Input<string>;
    /**
     * The username used to authenticate to the git repository.
     */
    username?: pulumi.Input<string>;
}
