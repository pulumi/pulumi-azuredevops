// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manages an Azure DevOps service endpoint within Azure DevOps.
 *
 * > **Note** This resource is duplicate with `azuredevops.ServiceEndpointPipeline`,  will be removed in the future, use `azuredevops.ServiceEndpointPipeline` instead.
 *
 * > **Note** Prerequisite: Extension [Configurable Pipeline Runner](https://marketplace.visualstudio.com/items?itemName=CSE-DevOps.RunPipelines) has been installed for the organization.
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
 * const exampleServiceEndpointAzureDevOps = new azuredevops.ServiceEndpointAzureDevOps("exampleServiceEndpointAzureDevOps", {
 *     projectId: exampleProject.id,
 *     serviceEndpointName: "Example Azure DevOps",
 *     orgUrl: "https://dev.azure.com/testorganization",
 *     releaseApiUrl: "https://vsrm.dev.azure.com/testorganization",
 *     personalAccessToken: "0000000000000000000000000000000000000000000000000000",
 *     description: "Managed by Terraform",
 * });
 * ```
 * ## Relevant Links
 *
 * - [Azure DevOps Service REST API 6.0 - Service Endpoints](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-6.0)
 *
 * ## Import
 *
 * Azure DevOps Service Endpoint Azure DevOps can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**
 *
 * ```sh
 *  $ pulumi import azuredevops:index/serviceEndpointAzureDevOps:ServiceEndpointAzureDevOps example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
 * ```
 */
export class ServiceEndpointAzureDevOps extends pulumi.CustomResource {
    /**
     * Get an existing ServiceEndpointAzureDevOps resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServiceEndpointAzureDevOpsState, opts?: pulumi.CustomResourceOptions): ServiceEndpointAzureDevOps {
        return new ServiceEndpointAzureDevOps(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuredevops:index/serviceEndpointAzureDevOps:ServiceEndpointAzureDevOps';

    /**
     * Returns true if the given object is an instance of ServiceEndpointAzureDevOps.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServiceEndpointAzureDevOps {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServiceEndpointAzureDevOps.__pulumiType;
    }

    public readonly authorization!: pulumi.Output<{[key: string]: string}>;
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The organization URL.
     */
    public readonly orgUrl!: pulumi.Output<string>;
    /**
     * The Azure DevOps personal access token.
     */
    public readonly personalAccessToken!: pulumi.Output<string>;
    /**
     * The ID of the project.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * The URL of the release API.
     */
    public readonly releaseApiUrl!: pulumi.Output<string>;
    /**
     * The Service Endpoint name.
     */
    public readonly serviceEndpointName!: pulumi.Output<string>;

    /**
     * Create a ServiceEndpointAzureDevOps resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServiceEndpointAzureDevOpsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServiceEndpointAzureDevOpsArgs | ServiceEndpointAzureDevOpsState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ServiceEndpointAzureDevOpsState | undefined;
            resourceInputs["authorization"] = state ? state.authorization : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["orgUrl"] = state ? state.orgUrl : undefined;
            resourceInputs["personalAccessToken"] = state ? state.personalAccessToken : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["releaseApiUrl"] = state ? state.releaseApiUrl : undefined;
            resourceInputs["serviceEndpointName"] = state ? state.serviceEndpointName : undefined;
        } else {
            const args = argsOrState as ServiceEndpointAzureDevOpsArgs | undefined;
            if ((!args || args.orgUrl === undefined) && !opts.urn) {
                throw new Error("Missing required property 'orgUrl'");
            }
            if ((!args || args.personalAccessToken === undefined) && !opts.urn) {
                throw new Error("Missing required property 'personalAccessToken'");
            }
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            if ((!args || args.releaseApiUrl === undefined) && !opts.urn) {
                throw new Error("Missing required property 'releaseApiUrl'");
            }
            if ((!args || args.serviceEndpointName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceEndpointName'");
            }
            resourceInputs["authorization"] = args ? args.authorization : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["orgUrl"] = args ? args.orgUrl : undefined;
            resourceInputs["personalAccessToken"] = args?.personalAccessToken ? pulumi.secret(args.personalAccessToken) : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["releaseApiUrl"] = args ? args.releaseApiUrl : undefined;
            resourceInputs["serviceEndpointName"] = args ? args.serviceEndpointName : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["personalAccessToken"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(ServiceEndpointAzureDevOps.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ServiceEndpointAzureDevOps resources.
 */
export interface ServiceEndpointAzureDevOpsState {
    authorization?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    description?: pulumi.Input<string>;
    /**
     * The organization URL.
     */
    orgUrl?: pulumi.Input<string>;
    /**
     * The Azure DevOps personal access token.
     */
    personalAccessToken?: pulumi.Input<string>;
    /**
     * The ID of the project.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The URL of the release API.
     */
    releaseApiUrl?: pulumi.Input<string>;
    /**
     * The Service Endpoint name.
     */
    serviceEndpointName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ServiceEndpointAzureDevOps resource.
 */
export interface ServiceEndpointAzureDevOpsArgs {
    authorization?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    description?: pulumi.Input<string>;
    /**
     * The organization URL.
     */
    orgUrl: pulumi.Input<string>;
    /**
     * The Azure DevOps personal access token.
     */
    personalAccessToken: pulumi.Input<string>;
    /**
     * The ID of the project.
     */
    projectId: pulumi.Input<string>;
    /**
     * The URL of the release API.
     */
    releaseApiUrl: pulumi.Input<string>;
    /**
     * The Service Endpoint name.
     */
    serviceEndpointName: pulumi.Input<string>;
}
