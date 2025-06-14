// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manages an Azure Service Bus service endpoint within Azure DevOps.
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
 * const exampleServiceendpointAzureServiceBus = new azuredevops.ServiceendpointAzureServiceBus("example", {
 *     projectId: example.id,
 *     serviceEndpointName: "Example Azure Service Bus",
 *     queueName: "queue",
 *     connectionString: "connection string",
 *     description: "Managed by Pulumi",
 * });
 * ```
 *
 * ## Relevant Links
 *
 * - [Azure DevOps Service REST API 7.0 - Endpoints](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-7.0)
 *
 * ## Import
 *
 * Azure DevOps Azure Service Bus Service Endpoint can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**
 *
 * ```sh
 * $ pulumi import azuredevops:index/serviceendpointAzureServiceBus:ServiceendpointAzureServiceBus example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
 * ```
 */
export class ServiceendpointAzureServiceBus extends pulumi.CustomResource {
    /**
     * Get an existing ServiceendpointAzureServiceBus resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServiceendpointAzureServiceBusState, opts?: pulumi.CustomResourceOptions): ServiceendpointAzureServiceBus {
        return new ServiceendpointAzureServiceBus(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuredevops:index/serviceendpointAzureServiceBus:ServiceendpointAzureServiceBus';

    /**
     * Returns true if the given object is an instance of ServiceendpointAzureServiceBus.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServiceendpointAzureServiceBus {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServiceendpointAzureServiceBus.__pulumiType;
    }

    public /*out*/ readonly authorization!: pulumi.Output<{[key: string]: string}>;
    /**
     * The  Azure Service Bus Connection string.
     */
    public readonly connectionString!: pulumi.Output<string>;
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The ID of the project.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * The Azure Service Bus Queue Name.
     */
    public readonly queueName!: pulumi.Output<string>;
    /**
     * The Service Endpoint name.
     */
    public readonly serviceEndpointName!: pulumi.Output<string>;

    /**
     * Create a ServiceendpointAzureServiceBus resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServiceendpointAzureServiceBusArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServiceendpointAzureServiceBusArgs | ServiceendpointAzureServiceBusState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ServiceendpointAzureServiceBusState | undefined;
            resourceInputs["authorization"] = state ? state.authorization : undefined;
            resourceInputs["connectionString"] = state ? state.connectionString : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["queueName"] = state ? state.queueName : undefined;
            resourceInputs["serviceEndpointName"] = state ? state.serviceEndpointName : undefined;
        } else {
            const args = argsOrState as ServiceendpointAzureServiceBusArgs | undefined;
            if ((!args || args.connectionString === undefined) && !opts.urn) {
                throw new Error("Missing required property 'connectionString'");
            }
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            if ((!args || args.queueName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'queueName'");
            }
            if ((!args || args.serviceEndpointName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceEndpointName'");
            }
            resourceInputs["connectionString"] = args?.connectionString ? pulumi.secret(args.connectionString) : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["queueName"] = args ? args.queueName : undefined;
            resourceInputs["serviceEndpointName"] = args ? args.serviceEndpointName : undefined;
            resourceInputs["authorization"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["connectionString"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(ServiceendpointAzureServiceBus.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ServiceendpointAzureServiceBus resources.
 */
export interface ServiceendpointAzureServiceBusState {
    authorization?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The  Azure Service Bus Connection string.
     */
    connectionString?: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    /**
     * The ID of the project.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The Azure Service Bus Queue Name.
     */
    queueName?: pulumi.Input<string>;
    /**
     * The Service Endpoint name.
     */
    serviceEndpointName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ServiceendpointAzureServiceBus resource.
 */
export interface ServiceendpointAzureServiceBusArgs {
    /**
     * The  Azure Service Bus Connection string.
     */
    connectionString: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    /**
     * The ID of the project.
     */
    projectId: pulumi.Input<string>;
    /**
     * The Azure Service Bus Queue Name.
     */
    queueName: pulumi.Input<string>;
    /**
     * The Service Endpoint name.
     */
    serviceEndpointName: pulumi.Input<string>;
}
