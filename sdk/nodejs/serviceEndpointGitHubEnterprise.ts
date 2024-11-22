// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Manages a GitHub Enterprise Server service endpoint within Azure DevOps.
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
 * });
 * const exampleServiceEndpointGitHubEnterprise = new azuredevops.ServiceEndpointGitHubEnterprise("example", {
 *     projectId: example.id,
 *     serviceEndpointName: "Example GitHub Enterprise",
 *     url: "https://github.contoso.com",
 *     description: "Managed by Pulumi",
 *     authPersonal: {
 *         personalAccessToken: "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
 *     },
 * });
 * ```
 *
 * ## Relevant Links
 *
 * - [Azure DevOps Service REST API 7.0 - Service Endpoints](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-7.0)
 *
 * ## Import
 *
 * Azure DevOps Service Endpoint GitHub Enterprise Server can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**
 *
 * ```sh
 * $ pulumi import azuredevops:index/serviceEndpointGitHubEnterprise:ServiceEndpointGitHubEnterprise example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
 * ```
 */
export class ServiceEndpointGitHubEnterprise extends pulumi.CustomResource {
    /**
     * Get an existing ServiceEndpointGitHubEnterprise resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServiceEndpointGitHubEnterpriseState, opts?: pulumi.CustomResourceOptions): ServiceEndpointGitHubEnterprise {
        return new ServiceEndpointGitHubEnterprise(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuredevops:index/serviceEndpointGitHubEnterprise:ServiceEndpointGitHubEnterprise';

    /**
     * Returns true if the given object is an instance of ServiceEndpointGitHubEnterprise.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServiceEndpointGitHubEnterprise {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServiceEndpointGitHubEnterprise.__pulumiType;
    }

    public readonly authPersonal!: pulumi.Output<outputs.ServiceEndpointGitHubEnterpriseAuthPersonal>;
    public readonly authorization!: pulumi.Output<{[key: string]: string}>;
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The ID of the project.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * The Service Endpoint name.
     */
    public readonly serviceEndpointName!: pulumi.Output<string>;
    /**
     * GitHub Enterprise Server Url.
     */
    public readonly url!: pulumi.Output<string>;

    /**
     * Create a ServiceEndpointGitHubEnterprise resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServiceEndpointGitHubEnterpriseArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServiceEndpointGitHubEnterpriseArgs | ServiceEndpointGitHubEnterpriseState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ServiceEndpointGitHubEnterpriseState | undefined;
            resourceInputs["authPersonal"] = state ? state.authPersonal : undefined;
            resourceInputs["authorization"] = state ? state.authorization : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["serviceEndpointName"] = state ? state.serviceEndpointName : undefined;
            resourceInputs["url"] = state ? state.url : undefined;
        } else {
            const args = argsOrState as ServiceEndpointGitHubEnterpriseArgs | undefined;
            if ((!args || args.authPersonal === undefined) && !opts.urn) {
                throw new Error("Missing required property 'authPersonal'");
            }
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            if ((!args || args.serviceEndpointName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceEndpointName'");
            }
            if ((!args || args.url === undefined) && !opts.urn) {
                throw new Error("Missing required property 'url'");
            }
            resourceInputs["authPersonal"] = args ? args.authPersonal : undefined;
            resourceInputs["authorization"] = args ? args.authorization : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["serviceEndpointName"] = args ? args.serviceEndpointName : undefined;
            resourceInputs["url"] = args ? args.url : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ServiceEndpointGitHubEnterprise.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ServiceEndpointGitHubEnterprise resources.
 */
export interface ServiceEndpointGitHubEnterpriseState {
    authPersonal?: pulumi.Input<inputs.ServiceEndpointGitHubEnterpriseAuthPersonal>;
    authorization?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    description?: pulumi.Input<string>;
    /**
     * The ID of the project.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The Service Endpoint name.
     */
    serviceEndpointName?: pulumi.Input<string>;
    /**
     * GitHub Enterprise Server Url.
     */
    url?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ServiceEndpointGitHubEnterprise resource.
 */
export interface ServiceEndpointGitHubEnterpriseArgs {
    authPersonal: pulumi.Input<inputs.ServiceEndpointGitHubEnterpriseAuthPersonal>;
    authorization?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    description?: pulumi.Input<string>;
    /**
     * The ID of the project.
     */
    projectId: pulumi.Input<string>;
    /**
     * The Service Endpoint name.
     */
    serviceEndpointName: pulumi.Input<string>;
    /**
     * GitHub Enterprise Server Url.
     */
    url: pulumi.Input<string>;
}
