// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manages an Incoming WebHook service endpoint within Azure DevOps, which can be used as a resource in YAML pipelines to subscribe to a webhook event.
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
 * const exampleServiceendpointIncomingwebhook = new azuredevops.ServiceendpointIncomingwebhook("example", {
 *     projectId: example.id,
 *     webhookName: "example_webhook",
 *     secret: "secret",
 *     httpHeader: "X-Hub-Signature",
 *     serviceEndpointName: "Example IncomingWebhook",
 *     description: "Managed by Pulumi",
 * });
 * ```
 *
 * ## Import
 *
 * Azure DevOps Incoming WebHook Service Endpoint can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**
 *
 * ```sh
 * $ pulumi import azuredevops:index/serviceendpointIncomingwebhook:ServiceendpointIncomingwebhook example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
 * ```
 */
export class ServiceendpointIncomingwebhook extends pulumi.CustomResource {
    /**
     * Get an existing ServiceendpointIncomingwebhook resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServiceendpointIncomingwebhookState, opts?: pulumi.CustomResourceOptions): ServiceendpointIncomingwebhook {
        return new ServiceendpointIncomingwebhook(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuredevops:index/serviceendpointIncomingwebhook:ServiceendpointIncomingwebhook';

    /**
     * Returns true if the given object is an instance of ServiceendpointIncomingwebhook.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServiceendpointIncomingwebhook {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServiceendpointIncomingwebhook.__pulumiType;
    }

    public readonly authorization!: pulumi.Output<{[key: string]: string}>;
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Http header name on which checksum will be sent.
     */
    public readonly httpHeader!: pulumi.Output<string | undefined>;
    /**
     * The ID of the project. Changing this forces a new Service Connection Incoming WebHook to be created.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * Secret for the WebHook. WebHook service will use this secret to calculate the payload checksum.
     */
    public readonly secret!: pulumi.Output<string | undefined>;
    /**
     * The name of the service endpoint. Changing this forces a new Service Connection Incoming WebHook to be created.
     */
    public readonly serviceEndpointName!: pulumi.Output<string>;
    /**
     * The name of the WebHook.
     */
    public readonly webhookName!: pulumi.Output<string>;

    /**
     * Create a ServiceendpointIncomingwebhook resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServiceendpointIncomingwebhookArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServiceendpointIncomingwebhookArgs | ServiceendpointIncomingwebhookState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ServiceendpointIncomingwebhookState | undefined;
            resourceInputs["authorization"] = state ? state.authorization : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["httpHeader"] = state ? state.httpHeader : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["secret"] = state ? state.secret : undefined;
            resourceInputs["serviceEndpointName"] = state ? state.serviceEndpointName : undefined;
            resourceInputs["webhookName"] = state ? state.webhookName : undefined;
        } else {
            const args = argsOrState as ServiceendpointIncomingwebhookArgs | undefined;
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            if ((!args || args.serviceEndpointName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceEndpointName'");
            }
            if ((!args || args.webhookName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'webhookName'");
            }
            resourceInputs["authorization"] = args ? args.authorization : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["httpHeader"] = args ? args.httpHeader : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["secret"] = args?.secret ? pulumi.secret(args.secret) : undefined;
            resourceInputs["serviceEndpointName"] = args ? args.serviceEndpointName : undefined;
            resourceInputs["webhookName"] = args ? args.webhookName : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["secret"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(ServiceendpointIncomingwebhook.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ServiceendpointIncomingwebhook resources.
 */
export interface ServiceendpointIncomingwebhookState {
    authorization?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    description?: pulumi.Input<string>;
    /**
     * Http header name on which checksum will be sent.
     */
    httpHeader?: pulumi.Input<string>;
    /**
     * The ID of the project. Changing this forces a new Service Connection Incoming WebHook to be created.
     */
    projectId?: pulumi.Input<string>;
    /**
     * Secret for the WebHook. WebHook service will use this secret to calculate the payload checksum.
     */
    secret?: pulumi.Input<string>;
    /**
     * The name of the service endpoint. Changing this forces a new Service Connection Incoming WebHook to be created.
     */
    serviceEndpointName?: pulumi.Input<string>;
    /**
     * The name of the WebHook.
     */
    webhookName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ServiceendpointIncomingwebhook resource.
 */
export interface ServiceendpointIncomingwebhookArgs {
    authorization?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    description?: pulumi.Input<string>;
    /**
     * Http header name on which checksum will be sent.
     */
    httpHeader?: pulumi.Input<string>;
    /**
     * The ID of the project. Changing this forces a new Service Connection Incoming WebHook to be created.
     */
    projectId: pulumi.Input<string>;
    /**
     * Secret for the WebHook. WebHook service will use this secret to calculate the payload checksum.
     */
    secret?: pulumi.Input<string>;
    /**
     * The name of the service endpoint. Changing this forces a new Service Connection Incoming WebHook to be created.
     */
    serviceEndpointName: pulumi.Input<string>;
    /**
     * The name of the WebHook.
     */
    webhookName: pulumi.Input<string>;
}
