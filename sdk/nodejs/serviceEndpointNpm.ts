// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manages a NPM service endpoint within Azure DevOps.
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
 * const exampleServiceEndpointNpm = new azuredevops.ServiceEndpointNpm("example", {
 *     projectId: example.id,
 *     serviceEndpointName: "Example npm",
 *     url: "https://registry.npmjs.org",
 *     accessToken: "00000000-0000-0000-0000-000000000000",
 *     description: "Managed by Pulumi",
 * });
 * ```
 *
 * ## Relevant Links
 *
 * - [Azure DevOps Service REST API 7.0 - Endpoints](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-7.0)
 * - [Azure DevOps Service Connections](https://docs.microsoft.com/en-us/azure/devops/pipelines/library/service-endpoints?view=azure-devops&tabs=yaml)
 * - [npm User Token](https://docs.npmjs.com/about-access-tokens)
 *
 * ## Import
 *
 * Azure DevOps NPM Service Endpoint can be imported using the **projectID/serviceEndpointID**, e.g.
 *
 * ```sh
 * $ pulumi import azuredevops:index/serviceEndpointNpm:ServiceEndpointNpm example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
 * ```
 */
export class ServiceEndpointNpm extends pulumi.CustomResource {
    /**
     * Get an existing ServiceEndpointNpm resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServiceEndpointNpmState, opts?: pulumi.CustomResourceOptions): ServiceEndpointNpm {
        return new ServiceEndpointNpm(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuredevops:index/serviceEndpointNpm:ServiceEndpointNpm';

    /**
     * Returns true if the given object is an instance of ServiceEndpointNpm.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServiceEndpointNpm {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServiceEndpointNpm.__pulumiType;
    }

    /**
     * The access token for npm registry.
     */
    public readonly accessToken!: pulumi.Output<string>;
    public /*out*/ readonly authorization!: pulumi.Output<{[key: string]: string}>;
    /**
     * The Service Endpoint description.
     */
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
     * URL of the npm registry to connect with.
     */
    public readonly url!: pulumi.Output<string>;

    /**
     * Create a ServiceEndpointNpm resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServiceEndpointNpmArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServiceEndpointNpmArgs | ServiceEndpointNpmState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ServiceEndpointNpmState | undefined;
            resourceInputs["accessToken"] = state ? state.accessToken : undefined;
            resourceInputs["authorization"] = state ? state.authorization : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["serviceEndpointName"] = state ? state.serviceEndpointName : undefined;
            resourceInputs["url"] = state ? state.url : undefined;
        } else {
            const args = argsOrState as ServiceEndpointNpmArgs | undefined;
            if ((!args || args.accessToken === undefined) && !opts.urn) {
                throw new Error("Missing required property 'accessToken'");
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
            resourceInputs["accessToken"] = args?.accessToken ? pulumi.secret(args.accessToken) : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["serviceEndpointName"] = args ? args.serviceEndpointName : undefined;
            resourceInputs["url"] = args ? args.url : undefined;
            resourceInputs["authorization"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["accessToken"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(ServiceEndpointNpm.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ServiceEndpointNpm resources.
 */
export interface ServiceEndpointNpmState {
    /**
     * The access token for npm registry.
     */
    accessToken?: pulumi.Input<string>;
    authorization?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The Service Endpoint description.
     */
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
     * URL of the npm registry to connect with.
     */
    url?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ServiceEndpointNpm resource.
 */
export interface ServiceEndpointNpmArgs {
    /**
     * The access token for npm registry.
     */
    accessToken: pulumi.Input<string>;
    /**
     * The Service Endpoint description.
     */
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
     * URL of the npm registry to connect with.
     */
    url: pulumi.Input<string>;
}
