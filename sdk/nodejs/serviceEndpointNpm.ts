// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manages a npm service endpoint within Azure DevOps.
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
 * const serviceendpoint = new azuredevops.ServiceEndpointNpm("serviceendpoint", {
 *     projectId: project.id,
 *     serviceEndpointName: "Sample npm",
 *     url: "https://registry.npmjs.org",
 *     accessToken: "00000000-0000-0000-0000-000000000000",
 *     description: "Managed by Terraform",
 * });
 * ```
 * ## Relevant Links
 *
 * - [Azure DevOps Service Connections](https://docs.microsoft.com/en-us/azure/devops/pipelines/library/service-endpoints?view=azure-devops&tabs=yaml)
 * - [npm User Token](https://docs.npmjs.com/about-access-tokens)
 *
 * ## Import
 *
 * Azure DevOps Service Endpoint npm can be imported using the **projectID/serviceEndpointID**, e.g.
 *
 * ```sh
 *  $ pulumi import azuredevops:index/serviceEndpointNpm:ServiceEndpointNpm serviceendpoint 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
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
    /**
     * A bcrypted hash of the attribute 'access_token'
     */
    public /*out*/ readonly accessTokenHash!: pulumi.Output<string>;
    public readonly authorization!: pulumi.Output<{[key: string]: string}>;
    /**
     * The Service Endpoint description.
     */
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
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ServiceEndpointNpmState | undefined;
            inputs["accessToken"] = state ? state.accessToken : undefined;
            inputs["accessTokenHash"] = state ? state.accessTokenHash : undefined;
            inputs["authorization"] = state ? state.authorization : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["projectId"] = state ? state.projectId : undefined;
            inputs["serviceEndpointName"] = state ? state.serviceEndpointName : undefined;
            inputs["url"] = state ? state.url : undefined;
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
            inputs["accessToken"] = args ? args.accessToken : undefined;
            inputs["authorization"] = args ? args.authorization : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["projectId"] = args ? args.projectId : undefined;
            inputs["serviceEndpointName"] = args ? args.serviceEndpointName : undefined;
            inputs["url"] = args ? args.url : undefined;
            inputs["accessTokenHash"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ServiceEndpointNpm.__pulumiType, name, inputs, opts);
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
    /**
     * A bcrypted hash of the attribute 'access_token'
     */
    accessTokenHash?: pulumi.Input<string>;
    authorization?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The Service Endpoint description.
     */
    description?: pulumi.Input<string>;
    /**
     * The project ID or project name.
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
    authorization?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The Service Endpoint description.
     */
    description?: pulumi.Input<string>;
    /**
     * The project ID or project name.
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
