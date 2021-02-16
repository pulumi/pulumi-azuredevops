// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manages a Azure Container Registry service endpoint within Azure DevOps.
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
 * // azure container registry service connection
 * const azurecr = new azuredevops.ServiceEndpointAzureEcr("azurecr", {
 *     projectId: project.id,
 *     serviceEndpointName: "Sample AzureCR",
 *     resourceGroup: "sample-rg",
 *     azurecrSpnTenantid: "00000000-0000-0000-0000-000000000000",
 *     azurecrName: "sampleAcr",
 *     azurecrSubscriptionId: "00000000-0000-0000-0000-000000000000",
 *     azurecrSubscriptionName: "sampleSub",
 * });
 * ```
 * ## Relevant Links
 *
 * - [Azure DevOps Service REST API 5.1 - Service Endpoints](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-5.1)
 * - [Azure Container Registry REST API](https://docs.microsoft.com/en-us/rest/api/containerregistry/)
 *
 * ## Import
 *
 * Azure DevOps Service Endpoint Azure Container Registry can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**
 *
 * ```sh
 *  $ pulumi import azuredevops:index/serviceEndpointAzureEcr:ServiceEndpointAzureEcr serviceendpoint 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
 * ```
 */
export class ServiceEndpointAzureEcr extends pulumi.CustomResource {
    /**
     * Get an existing ServiceEndpointAzureEcr resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServiceEndpointAzureEcrState, opts?: pulumi.CustomResourceOptions): ServiceEndpointAzureEcr {
        return new ServiceEndpointAzureEcr(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuredevops:index/serviceEndpointAzureEcr:ServiceEndpointAzureEcr';

    /**
     * Returns true if the given object is an instance of ServiceEndpointAzureEcr.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServiceEndpointAzureEcr {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServiceEndpointAzureEcr.__pulumiType;
    }

    public readonly authorization!: pulumi.Output<{[key: string]: string}>;
    /**
     * The Azure container registry name.
     */
    public readonly azurecrName!: pulumi.Output<string>;
    /**
     * The tenant id of the service principal.
     */
    public readonly azurecrSpnTenantid!: pulumi.Output<string>;
    /**
     * The subscription id of the Azure targets.
     */
    public readonly azurecrSubscriptionId!: pulumi.Output<string>;
    /**
     * The subscription name of the Azure targets.
     */
    public readonly azurecrSubscriptionName!: pulumi.Output<string>;
    /**
     * The name you will use to refer to this service connection in task inputs.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The project ID or project name.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * The resource group to which the container registry belongs.
     */
    public readonly resourceGroup!: pulumi.Output<string>;
    /**
     * The name you will use to refer to this service connection in task inputs.
     */
    public readonly serviceEndpointName!: pulumi.Output<string>;

    /**
     * Create a ServiceEndpointAzureEcr resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServiceEndpointAzureEcrArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServiceEndpointAzureEcrArgs | ServiceEndpointAzureEcrState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ServiceEndpointAzureEcrState | undefined;
            inputs["authorization"] = state ? state.authorization : undefined;
            inputs["azurecrName"] = state ? state.azurecrName : undefined;
            inputs["azurecrSpnTenantid"] = state ? state.azurecrSpnTenantid : undefined;
            inputs["azurecrSubscriptionId"] = state ? state.azurecrSubscriptionId : undefined;
            inputs["azurecrSubscriptionName"] = state ? state.azurecrSubscriptionName : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["projectId"] = state ? state.projectId : undefined;
            inputs["resourceGroup"] = state ? state.resourceGroup : undefined;
            inputs["serviceEndpointName"] = state ? state.serviceEndpointName : undefined;
        } else {
            const args = argsOrState as ServiceEndpointAzureEcrArgs | undefined;
            if ((!args || args.azurecrName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'azurecrName'");
            }
            if ((!args || args.azurecrSpnTenantid === undefined) && !opts.urn) {
                throw new Error("Missing required property 'azurecrSpnTenantid'");
            }
            if ((!args || args.azurecrSubscriptionId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'azurecrSubscriptionId'");
            }
            if ((!args || args.azurecrSubscriptionName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'azurecrSubscriptionName'");
            }
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            if ((!args || args.resourceGroup === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroup'");
            }
            if ((!args || args.serviceEndpointName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceEndpointName'");
            }
            inputs["authorization"] = args ? args.authorization : undefined;
            inputs["azurecrName"] = args ? args.azurecrName : undefined;
            inputs["azurecrSpnTenantid"] = args ? args.azurecrSpnTenantid : undefined;
            inputs["azurecrSubscriptionId"] = args ? args.azurecrSubscriptionId : undefined;
            inputs["azurecrSubscriptionName"] = args ? args.azurecrSubscriptionName : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["projectId"] = args ? args.projectId : undefined;
            inputs["resourceGroup"] = args ? args.resourceGroup : undefined;
            inputs["serviceEndpointName"] = args ? args.serviceEndpointName : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ServiceEndpointAzureEcr.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ServiceEndpointAzureEcr resources.
 */
export interface ServiceEndpointAzureEcrState {
    readonly authorization?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The Azure container registry name.
     */
    readonly azurecrName?: pulumi.Input<string>;
    /**
     * The tenant id of the service principal.
     */
    readonly azurecrSpnTenantid?: pulumi.Input<string>;
    /**
     * The subscription id of the Azure targets.
     */
    readonly azurecrSubscriptionId?: pulumi.Input<string>;
    /**
     * The subscription name of the Azure targets.
     */
    readonly azurecrSubscriptionName?: pulumi.Input<string>;
    /**
     * The name you will use to refer to this service connection in task inputs.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The project ID or project name.
     */
    readonly projectId?: pulumi.Input<string>;
    /**
     * The resource group to which the container registry belongs.
     */
    readonly resourceGroup?: pulumi.Input<string>;
    /**
     * The name you will use to refer to this service connection in task inputs.
     */
    readonly serviceEndpointName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ServiceEndpointAzureEcr resource.
 */
export interface ServiceEndpointAzureEcrArgs {
    readonly authorization?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The Azure container registry name.
     */
    readonly azurecrName: pulumi.Input<string>;
    /**
     * The tenant id of the service principal.
     */
    readonly azurecrSpnTenantid: pulumi.Input<string>;
    /**
     * The subscription id of the Azure targets.
     */
    readonly azurecrSubscriptionId: pulumi.Input<string>;
    /**
     * The subscription name of the Azure targets.
     */
    readonly azurecrSubscriptionName: pulumi.Input<string>;
    /**
     * The name you will use to refer to this service connection in task inputs.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The project ID or project name.
     */
    readonly projectId: pulumi.Input<string>;
    /**
     * The resource group to which the container registry belongs.
     */
    readonly resourceGroup: pulumi.Input<string>;
    /**
     * The name you will use to refer to this service connection in task inputs.
     */
    readonly serviceEndpointName: pulumi.Input<string>;
}
