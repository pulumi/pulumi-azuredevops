// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Manages a GitHub service endpoint within Azure DevOps.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const project = new azuredevops.Project("project", {
 *     visibility: "private",
 *     versionControl: "Git",
 *     workItemTemplate: "Agile",
 * });
 * const serviceendpointGh1 = new azuredevops.ServiceEndpointGitHub("serviceendpointGh1", {
 *     projectId: project.id,
 *     serviceEndpointName: "Sample GithHub Personal Access Token",
 *     authPersonal: {
 *         personalAccessToken: "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
 *     },
 * });
 * ```
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const serviceendpointGh2 = new azuredevops.ServiceEndpointGitHub("serviceendpointGh2", {
 *     projectId: azuredevops_project.project.id,
 *     serviceEndpointName: "Sample GithHub Grant",
 *     authOauth: {
 *         oauthConfigurationId: "00000000-0000-0000-0000-000000000000",
 *     },
 * });
 * ```
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const serviceendpointGh3 = new azuredevops.ServiceEndpointGitHub("serviceendpointGh3", {
 *     projectId: azuredevops_project.project.id,
 *     serviceEndpointName: "Sample GithHub Apps: Azure Pipelines",
 *     description: "",
 * });
 * ```
 * ## Relevant Links
 *
 * - [Azure DevOps Service REST API 5.1 - Service Endpoints](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-5.1)
 *
 * ## Import
 *
 * Azure DevOps Service Endpoint GitHub can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**
 *
 * ```sh
 *  $ pulumi import azuredevops:index/serviceEndpointGitHub:ServiceEndpointGitHub serviceendpoint 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
 * ```
 */
export class ServiceEndpointGitHub extends pulumi.CustomResource {
    /**
     * Get an existing ServiceEndpointGitHub resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServiceEndpointGitHubState, opts?: pulumi.CustomResourceOptions): ServiceEndpointGitHub {
        return new ServiceEndpointGitHub(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuredevops:index/serviceEndpointGitHub:ServiceEndpointGitHub';

    /**
     * Returns true if the given object is an instance of ServiceEndpointGitHub.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServiceEndpointGitHub {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServiceEndpointGitHub.__pulumiType;
    }

    /**
     * An `authOauth` block as documented below. Allows connecting using an Oauth token.
     */
    public readonly authOauth!: pulumi.Output<outputs.ServiceEndpointGitHubAuthOauth | undefined>;
    /**
     * An `authPersonal` block as documented below. Allows connecting using a personal access token.
     */
    public readonly authPersonal!: pulumi.Output<outputs.ServiceEndpointGitHubAuthPersonal | undefined>;
    public readonly authorization!: pulumi.Output<{[key: string]: string}>;
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The project ID or project name.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * The Service Endpoint name.
     */
    public readonly serviceEndpointName!: pulumi.Output<string>;

    /**
     * Create a ServiceEndpointGitHub resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServiceEndpointGitHubArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServiceEndpointGitHubArgs | ServiceEndpointGitHubState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ServiceEndpointGitHubState | undefined;
            inputs["authOauth"] = state ? state.authOauth : undefined;
            inputs["authPersonal"] = state ? state.authPersonal : undefined;
            inputs["authorization"] = state ? state.authorization : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["projectId"] = state ? state.projectId : undefined;
            inputs["serviceEndpointName"] = state ? state.serviceEndpointName : undefined;
        } else {
            const args = argsOrState as ServiceEndpointGitHubArgs | undefined;
            if ((!args || args.projectId === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'projectId'");
            }
            if ((!args || args.serviceEndpointName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'serviceEndpointName'");
            }
            inputs["authOauth"] = args ? args.authOauth : undefined;
            inputs["authPersonal"] = args ? args.authPersonal : undefined;
            inputs["authorization"] = args ? args.authorization : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["projectId"] = args ? args.projectId : undefined;
            inputs["serviceEndpointName"] = args ? args.serviceEndpointName : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azuredevops:ServiceEndpoint/gitHub:GitHub" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(ServiceEndpointGitHub.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ServiceEndpointGitHub resources.
 */
export interface ServiceEndpointGitHubState {
    /**
     * An `authOauth` block as documented below. Allows connecting using an Oauth token.
     */
    readonly authOauth?: pulumi.Input<inputs.ServiceEndpointGitHubAuthOauth>;
    /**
     * An `authPersonal` block as documented below. Allows connecting using a personal access token.
     */
    readonly authPersonal?: pulumi.Input<inputs.ServiceEndpointGitHubAuthPersonal>;
    readonly authorization?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly description?: pulumi.Input<string>;
    /**
     * The project ID or project name.
     */
    readonly projectId?: pulumi.Input<string>;
    /**
     * The Service Endpoint name.
     */
    readonly serviceEndpointName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ServiceEndpointGitHub resource.
 */
export interface ServiceEndpointGitHubArgs {
    /**
     * An `authOauth` block as documented below. Allows connecting using an Oauth token.
     */
    readonly authOauth?: pulumi.Input<inputs.ServiceEndpointGitHubAuthOauth>;
    /**
     * An `authPersonal` block as documented below. Allows connecting using a personal access token.
     */
    readonly authPersonal?: pulumi.Input<inputs.ServiceEndpointGitHubAuthPersonal>;
    readonly authorization?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly description?: pulumi.Input<string>;
    /**
     * The project ID or project name.
     */
    readonly projectId: pulumi.Input<string>;
    /**
     * The Service Endpoint name.
     */
    readonly serviceEndpointName: pulumi.Input<string>;
}
