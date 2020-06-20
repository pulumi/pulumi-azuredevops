// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages Manual or Automatic AzureRM service endpoint within Azure DevOps.
 *
 * ## Requirements (Manual AzureRM Service Endpoint)
 *
 * Before to create a service end point in Azure DevOps, you need to create a Service Principal in your Azure subscription.
 *
 * For detailed steps to create a service principal with Azure cli see the [documentation](https://docs.microsoft.com/en-us/cli/azure/create-an-azure-service-principal-azure-cli?view=azure-cli-latest)
 *
 * ## Example Usage
 *
 * ### Manual AzureRM Service Endpoint
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const project = new azuredevops.Core.Project("project", {
 *     projectName: "Sample Project",
 *     visibility: "private",
 *     versionControl: "Git",
 *     workItemTemplate: "Agile",
 * });
 * const endpointazure = new azuredevops.ServiceEndpoint.AzureRM("endpointazure", {
 *     projectId: project.id,
 *     serviceEndpointName: "TestServiceRM",
 *     credentials: {
 *         serviceprincipalid: "xxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxx",
 *         serviceprincipalkey: "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
 *     },
 *     azurermSpnTenantid: "xxxxxxx-xxxx-xxx-xxxxx-xxxxxxxx",
 *     azurermSubscriptionId: "xxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxx",
 *     azurermSubscriptionName: "Sample Subscription",
 * });
 * ```
 *
 * ### Automatic AzureRM Service Endpoint
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const project = new azuredevops.Core.Project("project", {
 *     projectName: "Sample Project",
 *     visibility: "private",
 *     versionControl: "Git",
 *     workItemTemplate: "Agile",
 * });
 * const endpointazure = new azuredevops.ServiceEndpoint.AzureRM("endpointazure", {
 *     projectId: project.id,
 *     serviceEndpointName: "TestServiceRM",
 *     azurermSpnTenantid: "xxxxxxx-xxxx-xxx-xxxxx-xxxxxxxx",
 *     azurermSubscriptionId: "xxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxx",
 *     azurermSubscriptionName: "Microsoft Azure DEMO",
 * });
 * ```
 *
 * ## Relevant Links
 *
 * * [Azure DevOps Service REST API 5.1 - Service End points](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-5.1)
 */
export class AzureRM extends pulumi.CustomResource {
    /**
     * Get an existing AzureRM resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AzureRMState, opts?: pulumi.CustomResourceOptions): AzureRM {
        return new AzureRM(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuredevops:ServiceEndpoint/azureRM:AzureRM';

    /**
     * Returns true if the given object is an instance of AzureRM.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AzureRM {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AzureRM.__pulumiType;
    }

    public readonly authorization!: pulumi.Output<{[key: string]: string}>;
    /**
     * The tenant id if the service principal.
     */
    public readonly azurermSpnTenantid!: pulumi.Output<string>;
    /**
     * The subscription Id of the Azure targets.
     */
    public readonly azurermSubscriptionId!: pulumi.Output<string>;
    /**
     * The subscription Name of the targets.
     */
    public readonly azurermSubscriptionName!: pulumi.Output<string>;
    /**
     * A `credentials` block.
     */
    public readonly credentials!: pulumi.Output<outputs.ServiceEndpoint.AzureRMCredentials | undefined>;
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The project ID or project name.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * The resource group used for scope of automatic service endpoint.
     */
    public readonly resourceGroup!: pulumi.Output<string | undefined>;
    /**
     * The Service Endpoint name.
     */
    public readonly serviceEndpointName!: pulumi.Output<string>;

    /**
     * Create a AzureRM resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AzureRMArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AzureRMArgs | AzureRMState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as AzureRMState | undefined;
            inputs["authorization"] = state ? state.authorization : undefined;
            inputs["azurermSpnTenantid"] = state ? state.azurermSpnTenantid : undefined;
            inputs["azurermSubscriptionId"] = state ? state.azurermSubscriptionId : undefined;
            inputs["azurermSubscriptionName"] = state ? state.azurermSubscriptionName : undefined;
            inputs["credentials"] = state ? state.credentials : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["projectId"] = state ? state.projectId : undefined;
            inputs["resourceGroup"] = state ? state.resourceGroup : undefined;
            inputs["serviceEndpointName"] = state ? state.serviceEndpointName : undefined;
        } else {
            const args = argsOrState as AzureRMArgs | undefined;
            if (!args || args.azurermSpnTenantid === undefined) {
                throw new Error("Missing required property 'azurermSpnTenantid'");
            }
            if (!args || args.azurermSubscriptionId === undefined) {
                throw new Error("Missing required property 'azurermSubscriptionId'");
            }
            if (!args || args.azurermSubscriptionName === undefined) {
                throw new Error("Missing required property 'azurermSubscriptionName'");
            }
            if (!args || args.projectId === undefined) {
                throw new Error("Missing required property 'projectId'");
            }
            if (!args || args.serviceEndpointName === undefined) {
                throw new Error("Missing required property 'serviceEndpointName'");
            }
            inputs["authorization"] = args ? args.authorization : undefined;
            inputs["azurermSpnTenantid"] = args ? args.azurermSpnTenantid : undefined;
            inputs["azurermSubscriptionId"] = args ? args.azurermSubscriptionId : undefined;
            inputs["azurermSubscriptionName"] = args ? args.azurermSubscriptionName : undefined;
            inputs["credentials"] = args ? args.credentials : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["projectId"] = args ? args.projectId : undefined;
            inputs["resourceGroup"] = args ? args.resourceGroup : undefined;
            inputs["serviceEndpointName"] = args ? args.serviceEndpointName : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(AzureRM.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AzureRM resources.
 */
export interface AzureRMState {
    readonly authorization?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The tenant id if the service principal.
     */
    readonly azurermSpnTenantid?: pulumi.Input<string>;
    /**
     * The subscription Id of the Azure targets.
     */
    readonly azurermSubscriptionId?: pulumi.Input<string>;
    /**
     * The subscription Name of the targets.
     */
    readonly azurermSubscriptionName?: pulumi.Input<string>;
    /**
     * A `credentials` block.
     */
    readonly credentials?: pulumi.Input<inputs.ServiceEndpoint.AzureRMCredentials>;
    readonly description?: pulumi.Input<string>;
    /**
     * The project ID or project name.
     */
    readonly projectId?: pulumi.Input<string>;
    /**
     * The resource group used for scope of automatic service endpoint.
     */
    readonly resourceGroup?: pulumi.Input<string>;
    /**
     * The Service Endpoint name.
     */
    readonly serviceEndpointName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AzureRM resource.
 */
export interface AzureRMArgs {
    readonly authorization?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The tenant id if the service principal.
     */
    readonly azurermSpnTenantid: pulumi.Input<string>;
    /**
     * The subscription Id of the Azure targets.
     */
    readonly azurermSubscriptionId: pulumi.Input<string>;
    /**
     * The subscription Name of the targets.
     */
    readonly azurermSubscriptionName: pulumi.Input<string>;
    /**
     * A `credentials` block.
     */
    readonly credentials?: pulumi.Input<inputs.ServiceEndpoint.AzureRMCredentials>;
    readonly description?: pulumi.Input<string>;
    /**
     * The project ID or project name.
     */
    readonly projectId: pulumi.Input<string>;
    /**
     * The resource group used for scope of automatic service endpoint.
     */
    readonly resourceGroup?: pulumi.Input<string>;
    /**
     * The Service Endpoint name.
     */
    readonly serviceEndpointName: pulumi.Input<string>;
}
