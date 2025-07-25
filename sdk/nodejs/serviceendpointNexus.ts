// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manages a Nexus IQ service endpoint within Azure DevOps, which can be used as a resource in YAML pipelines to connect to a Nexus IQ instance.
 * Nexus IQ is not supported by default, to manage a nexus service connection resource, it is necessary to install the [Nexus Extension](https://marketplace.visualstudio.com/items?itemName=SonatypeIntegrations.nexus-iq-azure-extension) in Azure DevOps.
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
 * const exampleServiceendpointNexus = new azuredevops.ServiceendpointNexus("example", {
 *     projectId: example.id,
 *     serviceEndpointName: "nexus-example",
 *     description: "Service Endpoint for 'Nexus IQ' (Managed by Terraform)",
 *     url: "https://example.com",
 *     username: "username",
 *     password: "password",
 * });
 * ```
 *
 * ## Import
 *
 * Azure DevOps Nexus Service Connection can be imported using the `projectId/id` or or `projectName/id`, e.g.
 *
 * ```sh
 * $ pulumi import azuredevops:index/serviceendpointNexus:ServiceendpointNexus example projectName/00000000-0000-0000-0000-000000000000
 * ```
 */
export class ServiceendpointNexus extends pulumi.CustomResource {
    /**
     * Get an existing ServiceendpointNexus resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServiceendpointNexusState, opts?: pulumi.CustomResourceOptions): ServiceendpointNexus {
        return new ServiceendpointNexus(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuredevops:index/serviceendpointNexus:ServiceendpointNexus';

    /**
     * Returns true if the given object is an instance of ServiceendpointNexus.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServiceendpointNexus {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServiceendpointNexus.__pulumiType;
    }

    public /*out*/ readonly authorization!: pulumi.Output<{[key: string]: string}>;
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The Service Endpoint password to authenticate at the Nexus IQ Instance.
     */
    public readonly password!: pulumi.Output<string>;
    /**
     * The ID of the project. Changing this forces a new Service Connection Nexus to be created.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * The name of the service endpoint. Changing this forces a new Service Connection Nexus to be created.
     */
    public readonly serviceEndpointName!: pulumi.Output<string>;
    /**
     * The Service Endpoint url.
     */
    public readonly url!: pulumi.Output<string>;
    /**
     * The Service Endpoint username to authenticate at the Nexus IQ Instance.
     */
    public readonly username!: pulumi.Output<string>;

    /**
     * Create a ServiceendpointNexus resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServiceendpointNexusArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServiceendpointNexusArgs | ServiceendpointNexusState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ServiceendpointNexusState | undefined;
            resourceInputs["authorization"] = state ? state.authorization : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["serviceEndpointName"] = state ? state.serviceEndpointName : undefined;
            resourceInputs["url"] = state ? state.url : undefined;
            resourceInputs["username"] = state ? state.username : undefined;
        } else {
            const args = argsOrState as ServiceendpointNexusArgs | undefined;
            if ((!args || args.password === undefined) && !opts.urn) {
                throw new Error("Missing required property 'password'");
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
            if ((!args || args.username === undefined) && !opts.urn) {
                throw new Error("Missing required property 'username'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["password"] = args?.password ? pulumi.secret(args.password) : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["serviceEndpointName"] = args ? args.serviceEndpointName : undefined;
            resourceInputs["url"] = args ? args.url : undefined;
            resourceInputs["username"] = args ? args.username : undefined;
            resourceInputs["authorization"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["password"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(ServiceendpointNexus.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ServiceendpointNexus resources.
 */
export interface ServiceendpointNexusState {
    authorization?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    description?: pulumi.Input<string>;
    /**
     * The Service Endpoint password to authenticate at the Nexus IQ Instance.
     */
    password?: pulumi.Input<string>;
    /**
     * The ID of the project. Changing this forces a new Service Connection Nexus to be created.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The name of the service endpoint. Changing this forces a new Service Connection Nexus to be created.
     */
    serviceEndpointName?: pulumi.Input<string>;
    /**
     * The Service Endpoint url.
     */
    url?: pulumi.Input<string>;
    /**
     * The Service Endpoint username to authenticate at the Nexus IQ Instance.
     */
    username?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ServiceendpointNexus resource.
 */
export interface ServiceendpointNexusArgs {
    description?: pulumi.Input<string>;
    /**
     * The Service Endpoint password to authenticate at the Nexus IQ Instance.
     */
    password: pulumi.Input<string>;
    /**
     * The ID of the project. Changing this forces a new Service Connection Nexus to be created.
     */
    projectId: pulumi.Input<string>;
    /**
     * The name of the service endpoint. Changing this forces a new Service Connection Nexus to be created.
     */
    serviceEndpointName: pulumi.Input<string>;
    /**
     * The Service Endpoint url.
     */
    url: pulumi.Input<string>;
    /**
     * The Service Endpoint username to authenticate at the Nexus IQ Instance.
     */
    username: pulumi.Input<string>;
}
