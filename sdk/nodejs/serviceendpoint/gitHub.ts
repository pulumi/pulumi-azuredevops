// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages a GitHub service endpoint within Azure DevOps.
 *
 * ## Relevant Links
 *
 * - [Azure DevOps Service REST API 7.0 - Service Endpoints](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-7.0)
 *
 * ## Import
 *
 * Azure DevOps Service Endpoint GitHub can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**
 *
 * ```sh
 *  $ pulumi import azuredevops:ServiceEndpoint/gitHub:GitHub example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
 * ```
 *
 * @deprecated azuredevops.serviceendpoint.GitHub has been deprecated in favor of azuredevops.ServiceEndpointGitHub
 */
export class GitHub extends pulumi.CustomResource {
    /**
     * Get an existing GitHub resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GitHubState, opts?: pulumi.CustomResourceOptions): GitHub {
        pulumi.log.warn("GitHub is deprecated: azuredevops.serviceendpoint.GitHub has been deprecated in favor of azuredevops.ServiceEndpointGitHub")
        return new GitHub(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuredevops:ServiceEndpoint/gitHub:GitHub';

    /**
     * Returns true if the given object is an instance of GitHub.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GitHub {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GitHub.__pulumiType;
    }

    public readonly authOauth!: pulumi.Output<outputs.ServiceEndpoint.GitHubAuthOauth | undefined>;
    /**
     * An `authPersonal` block as documented below. Allows connecting using a personal access token.
     */
    public readonly authPersonal!: pulumi.Output<outputs.ServiceEndpoint.GitHubAuthPersonal | undefined>;
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
     * Create a GitHub resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated azuredevops.serviceendpoint.GitHub has been deprecated in favor of azuredevops.ServiceEndpointGitHub */
    constructor(name: string, args: GitHubArgs, opts?: pulumi.CustomResourceOptions)
    /** @deprecated azuredevops.serviceendpoint.GitHub has been deprecated in favor of azuredevops.ServiceEndpointGitHub */
    constructor(name: string, argsOrState?: GitHubArgs | GitHubState, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("GitHub is deprecated: azuredevops.serviceendpoint.GitHub has been deprecated in favor of azuredevops.ServiceEndpointGitHub")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GitHubState | undefined;
            resourceInputs["authOauth"] = state ? state.authOauth : undefined;
            resourceInputs["authPersonal"] = state ? state.authPersonal : undefined;
            resourceInputs["authorization"] = state ? state.authorization : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["serviceEndpointName"] = state ? state.serviceEndpointName : undefined;
        } else {
            const args = argsOrState as GitHubArgs | undefined;
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            if ((!args || args.serviceEndpointName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceEndpointName'");
            }
            resourceInputs["authOauth"] = args ? args.authOauth : undefined;
            resourceInputs["authPersonal"] = args ? args.authPersonal : undefined;
            resourceInputs["authorization"] = args ? args.authorization : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["serviceEndpointName"] = args ? args.serviceEndpointName : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(GitHub.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GitHub resources.
 */
export interface GitHubState {
    authOauth?: pulumi.Input<inputs.ServiceEndpoint.GitHubAuthOauth>;
    /**
     * An `authPersonal` block as documented below. Allows connecting using a personal access token.
     */
    authPersonal?: pulumi.Input<inputs.ServiceEndpoint.GitHubAuthPersonal>;
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
}

/**
 * The set of arguments for constructing a GitHub resource.
 */
export interface GitHubArgs {
    authOauth?: pulumi.Input<inputs.ServiceEndpoint.GitHubAuthOauth>;
    /**
     * An `authPersonal` block as documented below. Allows connecting using a personal access token.
     */
    authPersonal?: pulumi.Input<inputs.ServiceEndpoint.GitHubAuthPersonal>;
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
}
