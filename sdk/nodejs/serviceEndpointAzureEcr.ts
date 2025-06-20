// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Manages a Azure Container Registry service endpoint within Azure DevOps.
 *
 * ## Example Usage
 *
 * ### Authorize with Service Principal
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
 * // azure container registry service connection
 * const exampleServiceEndpointAzureEcr = new azuredevops.ServiceEndpointAzureEcr("example", {
 *     projectId: example.id,
 *     serviceEndpointName: "Example AzureCR",
 *     resourceGroup: "example-rg",
 *     azurecrSpnTenantid: "00000000-0000-0000-0000-000000000000",
 *     azurecrName: "ExampleAcr",
 *     azurecrSubscriptionId: "00000000-0000-0000-0000-000000000000",
 *     azurecrSubscriptionName: "subscription name",
 * });
 * ```
 *
 * ### Authorize with WorkloadIdentityFederation
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const example = new azuredevops.Project("example", {
 *     name: "Example Project",
 *     visibility: "private",
 *     versionControl: "Git",
 *     workItemTemplate: "Agile",
 *     description: "Managed by Pulumi",
 * });
 * const identity = new azure.core.ResourceGroup("identity", {
 *     name: "identity",
 *     location: "UK South",
 * });
 * const exampleUserAssignedIdentity = new azure.authorization.UserAssignedIdentity("example", {
 *     location: identity.location,
 *     name: "example-identity",
 *     resourceGroupName: identity.name,
 * });
 * // azure container registry service connection
 * const exampleServiceEndpointAzureEcr = new azuredevops.ServiceEndpointAzureEcr("example", {
 *     projectId: example.id,
 *     resourceGroup: "Example AzureCR ResourceGroup",
 *     serviceEndpointName: "Example AzureCR",
 *     serviceEndpointAuthenticationScheme: "WorkloadIdentityFederation",
 *     azurecrSpnTenantid: "00000000-0000-0000-0000-000000000000",
 *     azurecrName: "ExampleAcr",
 *     azurecrSubscriptionId: "00000000-0000-0000-0000-000000000000",
 *     azurecrSubscriptionName: "subscription name",
 *     credentials: {
 *         serviceprincipalid: exampleUserAssignedIdentity.clientId,
 *     },
 * });
 * const exampleFederatedIdentityCredential = new azure.armmsi.FederatedIdentityCredential("example", {
 *     name: "example-federated-credential",
 *     resourceGroupName: identity.name,
 *     parentId: exampleUserAssignedIdentity.id,
 *     audience: "api://AzureADTokenExchange",
 *     issuer: exampleServiceEndpointAzureEcr.workloadIdentityFederationIssuer,
 *     subject: exampleServiceEndpointAzureEcr.workloadIdentityFederationSubject,
 * });
 * ```
 *
 * ## Relevant Links
 *
 * - [Azure DevOps Service REST API 7.0 - Service Endpoints](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-7.0)
 * - [Azure Container Registry REST API](https://docs.microsoft.com/en-us/rest/api/containerregistry/)
 *
 * ## Import
 *
 * Azure DevOps Azure Container Registry Service Endpoint can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**
 *
 * ```sh
 * $ pulumi import azuredevops:index/serviceEndpointAzureEcr:ServiceEndpointAzureEcr example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
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

    public /*out*/ readonly appObjectId!: pulumi.Output<string>;
    public /*out*/ readonly authorization!: pulumi.Output<{[key: string]: string}>;
    public /*out*/ readonly azSpnRoleAssignmentId!: pulumi.Output<string>;
    public /*out*/ readonly azSpnRolePermissions!: pulumi.Output<string>;
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
     * A `credentials` block as defined below.
     */
    public readonly credentials!: pulumi.Output<outputs.ServiceEndpointAzureEcrCredentials | undefined>;
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The ID of the project.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * The resource group to which the container registry belongs.
     */
    public readonly resourceGroup!: pulumi.Output<string | undefined>;
    /**
     * Specifies the type of azurerm endpoint, either `WorkloadIdentityFederation`, `ManagedServiceIdentity` or `ServicePrincipal`. Defaults to `ServicePrincipal` for backwards compatibility. `ManagedServiceIdentity` has not yet been implemented for this resource.
     */
    public readonly serviceEndpointAuthenticationScheme!: pulumi.Output<string | undefined>;
    /**
     * The name you will use to refer to this service connection in task inputs.
     */
    public readonly serviceEndpointName!: pulumi.Output<string>;
    /**
     * The Application(Client) ID of the Service Principal.
     */
    public /*out*/ readonly servicePrincipalId!: pulumi.Output<string>;
    public /*out*/ readonly spnObjectId!: pulumi.Output<string>;
    /**
     * The issuer if `serviceEndpointAuthenticationScheme` is set to `WorkloadIdentityFederation`. This looks like `https://vstoken.dev.azure.com/00000000-0000-0000-0000-000000000000`, where the GUID is the Organization ID of your Azure DevOps Organisation.
     */
    public /*out*/ readonly workloadIdentityFederationIssuer!: pulumi.Output<string>;
    /**
     * The subject if `serviceEndpointAuthenticationScheme` is set to `WorkloadIdentityFederation`. This looks like `sc://<organisation>/<project>/<service-connection-name>`.
     */
    public /*out*/ readonly workloadIdentityFederationSubject!: pulumi.Output<string>;

    /**
     * Create a ServiceEndpointAzureEcr resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServiceEndpointAzureEcrArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServiceEndpointAzureEcrArgs | ServiceEndpointAzureEcrState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ServiceEndpointAzureEcrState | undefined;
            resourceInputs["appObjectId"] = state ? state.appObjectId : undefined;
            resourceInputs["authorization"] = state ? state.authorization : undefined;
            resourceInputs["azSpnRoleAssignmentId"] = state ? state.azSpnRoleAssignmentId : undefined;
            resourceInputs["azSpnRolePermissions"] = state ? state.azSpnRolePermissions : undefined;
            resourceInputs["azurecrName"] = state ? state.azurecrName : undefined;
            resourceInputs["azurecrSpnTenantid"] = state ? state.azurecrSpnTenantid : undefined;
            resourceInputs["azurecrSubscriptionId"] = state ? state.azurecrSubscriptionId : undefined;
            resourceInputs["azurecrSubscriptionName"] = state ? state.azurecrSubscriptionName : undefined;
            resourceInputs["credentials"] = state ? state.credentials : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["resourceGroup"] = state ? state.resourceGroup : undefined;
            resourceInputs["serviceEndpointAuthenticationScheme"] = state ? state.serviceEndpointAuthenticationScheme : undefined;
            resourceInputs["serviceEndpointName"] = state ? state.serviceEndpointName : undefined;
            resourceInputs["servicePrincipalId"] = state ? state.servicePrincipalId : undefined;
            resourceInputs["spnObjectId"] = state ? state.spnObjectId : undefined;
            resourceInputs["workloadIdentityFederationIssuer"] = state ? state.workloadIdentityFederationIssuer : undefined;
            resourceInputs["workloadIdentityFederationSubject"] = state ? state.workloadIdentityFederationSubject : undefined;
        } else {
            const args = argsOrState as ServiceEndpointAzureEcrArgs | undefined;
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            if ((!args || args.serviceEndpointName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceEndpointName'");
            }
            resourceInputs["azurecrName"] = args ? args.azurecrName : undefined;
            resourceInputs["azurecrSpnTenantid"] = args ? args.azurecrSpnTenantid : undefined;
            resourceInputs["azurecrSubscriptionId"] = args ? args.azurecrSubscriptionId : undefined;
            resourceInputs["azurecrSubscriptionName"] = args ? args.azurecrSubscriptionName : undefined;
            resourceInputs["credentials"] = args ? args.credentials : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["resourceGroup"] = args ? args.resourceGroup : undefined;
            resourceInputs["serviceEndpointAuthenticationScheme"] = args ? args.serviceEndpointAuthenticationScheme : undefined;
            resourceInputs["serviceEndpointName"] = args ? args.serviceEndpointName : undefined;
            resourceInputs["appObjectId"] = undefined /*out*/;
            resourceInputs["authorization"] = undefined /*out*/;
            resourceInputs["azSpnRoleAssignmentId"] = undefined /*out*/;
            resourceInputs["azSpnRolePermissions"] = undefined /*out*/;
            resourceInputs["servicePrincipalId"] = undefined /*out*/;
            resourceInputs["spnObjectId"] = undefined /*out*/;
            resourceInputs["workloadIdentityFederationIssuer"] = undefined /*out*/;
            resourceInputs["workloadIdentityFederationSubject"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ServiceEndpointAzureEcr.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ServiceEndpointAzureEcr resources.
 */
export interface ServiceEndpointAzureEcrState {
    appObjectId?: pulumi.Input<string>;
    authorization?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    azSpnRoleAssignmentId?: pulumi.Input<string>;
    azSpnRolePermissions?: pulumi.Input<string>;
    /**
     * The Azure container registry name.
     */
    azurecrName?: pulumi.Input<string>;
    /**
     * The tenant id of the service principal.
     */
    azurecrSpnTenantid?: pulumi.Input<string>;
    /**
     * The subscription id of the Azure targets.
     */
    azurecrSubscriptionId?: pulumi.Input<string>;
    /**
     * The subscription name of the Azure targets.
     */
    azurecrSubscriptionName?: pulumi.Input<string>;
    /**
     * A `credentials` block as defined below.
     */
    credentials?: pulumi.Input<inputs.ServiceEndpointAzureEcrCredentials>;
    description?: pulumi.Input<string>;
    /**
     * The ID of the project.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The resource group to which the container registry belongs.
     */
    resourceGroup?: pulumi.Input<string>;
    /**
     * Specifies the type of azurerm endpoint, either `WorkloadIdentityFederation`, `ManagedServiceIdentity` or `ServicePrincipal`. Defaults to `ServicePrincipal` for backwards compatibility. `ManagedServiceIdentity` has not yet been implemented for this resource.
     */
    serviceEndpointAuthenticationScheme?: pulumi.Input<string>;
    /**
     * The name you will use to refer to this service connection in task inputs.
     */
    serviceEndpointName?: pulumi.Input<string>;
    /**
     * The Application(Client) ID of the Service Principal.
     */
    servicePrincipalId?: pulumi.Input<string>;
    spnObjectId?: pulumi.Input<string>;
    /**
     * The issuer if `serviceEndpointAuthenticationScheme` is set to `WorkloadIdentityFederation`. This looks like `https://vstoken.dev.azure.com/00000000-0000-0000-0000-000000000000`, where the GUID is the Organization ID of your Azure DevOps Organisation.
     */
    workloadIdentityFederationIssuer?: pulumi.Input<string>;
    /**
     * The subject if `serviceEndpointAuthenticationScheme` is set to `WorkloadIdentityFederation`. This looks like `sc://<organisation>/<project>/<service-connection-name>`.
     */
    workloadIdentityFederationSubject?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ServiceEndpointAzureEcr resource.
 */
export interface ServiceEndpointAzureEcrArgs {
    /**
     * The Azure container registry name.
     */
    azurecrName?: pulumi.Input<string>;
    /**
     * The tenant id of the service principal.
     */
    azurecrSpnTenantid?: pulumi.Input<string>;
    /**
     * The subscription id of the Azure targets.
     */
    azurecrSubscriptionId?: pulumi.Input<string>;
    /**
     * The subscription name of the Azure targets.
     */
    azurecrSubscriptionName?: pulumi.Input<string>;
    /**
     * A `credentials` block as defined below.
     */
    credentials?: pulumi.Input<inputs.ServiceEndpointAzureEcrCredentials>;
    description?: pulumi.Input<string>;
    /**
     * The ID of the project.
     */
    projectId: pulumi.Input<string>;
    /**
     * The resource group to which the container registry belongs.
     */
    resourceGroup?: pulumi.Input<string>;
    /**
     * Specifies the type of azurerm endpoint, either `WorkloadIdentityFederation`, `ManagedServiceIdentity` or `ServicePrincipal`. Defaults to `ServicePrincipal` for backwards compatibility. `ManagedServiceIdentity` has not yet been implemented for this resource.
     */
    serviceEndpointAuthenticationScheme?: pulumi.Input<string>;
    /**
     * The name you will use to refer to this service connection in task inputs.
     */
    serviceEndpointName: pulumi.Input<string>;
}
