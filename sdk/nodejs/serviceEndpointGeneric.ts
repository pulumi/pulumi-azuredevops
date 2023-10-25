// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manages a generic service endpoint within Azure DevOps, which can be used to authenticate to any external server using
 * basic authentication via a username and password.
 *
 * ## Relevant Links
 *
 * - [Azure DevOps Service REST API 7.0 - Agent Pools](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-7.0)
 *
 * ## Import
 *
 * Azure DevOps Service Endpoint Generic can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**
 *
 * ```sh
 *  $ pulumi import azuredevops:index/serviceEndpointGeneric:ServiceEndpointGeneric example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
 * ```
 */
export class ServiceEndpointGeneric extends pulumi.CustomResource {
    /**
     * Get an existing ServiceEndpointGeneric resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServiceEndpointGenericState, opts?: pulumi.CustomResourceOptions): ServiceEndpointGeneric {
        return new ServiceEndpointGeneric(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuredevops:index/serviceEndpointGeneric:ServiceEndpointGeneric';

    /**
     * Returns true if the given object is an instance of ServiceEndpointGeneric.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServiceEndpointGeneric {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServiceEndpointGeneric.__pulumiType;
    }

    public readonly authorization!: pulumi.Output<{[key: string]: string}>;
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The password or token key used to authenticate to the server url using basic authentication.
     */
    public readonly password!: pulumi.Output<string | undefined>;
    /**
     * The ID of the project.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * The URL of the server associated with the service endpoint.
     */
    public readonly serverUrl!: pulumi.Output<string>;
    /**
     * The service endpoint name.
     */
    public readonly serviceEndpointName!: pulumi.Output<string>;
    /**
     * The username used to authenticate to the server url using basic authentication.
     */
    public readonly username!: pulumi.Output<string | undefined>;

    /**
     * Create a ServiceEndpointGeneric resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServiceEndpointGenericArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServiceEndpointGenericArgs | ServiceEndpointGenericState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ServiceEndpointGenericState | undefined;
            resourceInputs["authorization"] = state ? state.authorization : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["serverUrl"] = state ? state.serverUrl : undefined;
            resourceInputs["serviceEndpointName"] = state ? state.serviceEndpointName : undefined;
            resourceInputs["username"] = state ? state.username : undefined;
        } else {
            const args = argsOrState as ServiceEndpointGenericArgs | undefined;
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            if ((!args || args.serverUrl === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serverUrl'");
            }
            if ((!args || args.serviceEndpointName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceEndpointName'");
            }
            resourceInputs["authorization"] = args ? args.authorization : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["password"] = args?.password ? pulumi.secret(args.password) : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["serverUrl"] = args ? args.serverUrl : undefined;
            resourceInputs["serviceEndpointName"] = args ? args.serviceEndpointName : undefined;
            resourceInputs["username"] = args ? args.username : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["password"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(ServiceEndpointGeneric.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ServiceEndpointGeneric resources.
 */
export interface ServiceEndpointGenericState {
    authorization?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    description?: pulumi.Input<string>;
    /**
     * The password or token key used to authenticate to the server url using basic authentication.
     */
    password?: pulumi.Input<string>;
    /**
     * The ID of the project.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The URL of the server associated with the service endpoint.
     */
    serverUrl?: pulumi.Input<string>;
    /**
     * The service endpoint name.
     */
    serviceEndpointName?: pulumi.Input<string>;
    /**
     * The username used to authenticate to the server url using basic authentication.
     */
    username?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ServiceEndpointGeneric resource.
 */
export interface ServiceEndpointGenericArgs {
    authorization?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    description?: pulumi.Input<string>;
    /**
     * The password or token key used to authenticate to the server url using basic authentication.
     */
    password?: pulumi.Input<string>;
    /**
     * The ID of the project.
     */
    projectId: pulumi.Input<string>;
    /**
     * The URL of the server associated with the service endpoint.
     */
    serverUrl: pulumi.Input<string>;
    /**
     * The service endpoint name.
     */
    serviceEndpointName: pulumi.Input<string>;
    /**
     * The username used to authenticate to the server url using basic authentication.
     */
    username?: pulumi.Input<string>;
}
