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
 * ### Service Principal Manual AzureRM Service Endpoint (Subscription Scoped)
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const exampleProject = new azuredevops.Project("exampleProject", {
 *     visibility: "private",
 *     versionControl: "Git",
 *     workItemTemplate: "Agile",
 *     description: "Managed by Terraform",
 * });
 * const exampleServiceEndpointAzureRM = new azuredevops.ServiceEndpointAzureRM("exampleServiceEndpointAzureRM", {
 *     projectId: exampleProject.id,
 *     serviceEndpointName: "Example AzureRM",
 *     description: "Managed by Terraform",
 *     serviceEndpointAuthenticationScheme: "ServicePrincipal",
 *     credentials: {
 *         serviceprincipalid: "00000000-0000-0000-0000-000000000000",
 *         serviceprincipalkey: "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
 *     },
 *     azurermSpnTenantid: "00000000-0000-0000-0000-000000000000",
 *     azurermSubscriptionId: "00000000-0000-0000-0000-000000000000",
 *     azurermSubscriptionName: "Example Subscription Name",
 * });
 * ```
 * ### Service Principal Manual AzureRM Service Endpoint (ManagementGroup Scoped)
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const exampleProject = new azuredevops.Project("exampleProject", {
 *     visibility: "private",
 *     versionControl: "Git",
 *     workItemTemplate: "Agile",
 *     description: "Managed by Terraform",
 * });
 * const exampleServiceEndpointAzureRM = new azuredevops.ServiceEndpointAzureRM("exampleServiceEndpointAzureRM", {
 *     projectId: exampleProject.id,
 *     serviceEndpointName: "Example AzureRM",
 *     description: "Managed by Terraform",
 *     serviceEndpointAuthenticationScheme: "ServicePrincipal",
 *     credentials: {
 *         serviceprincipalid: "00000000-0000-0000-0000-000000000000",
 *         serviceprincipalkey: "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
 *     },
 *     azurermSpnTenantid: "00000000-0000-0000-0000-000000000000",
 *     azurermManagementGroupId: "managementGroup",
 *     azurermManagementGroupName: "managementGroup",
 * });
 * ```
 * ### Service Principal Automatic AzureRM Service Endpoint
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const exampleProject = new azuredevops.Project("exampleProject", {
 *     visibility: "private",
 *     versionControl: "Git",
 *     workItemTemplate: "Agile",
 * });
 * const exampleServiceEndpointAzureRM = new azuredevops.ServiceEndpointAzureRM("exampleServiceEndpointAzureRM", {
 *     projectId: exampleProject.id,
 *     serviceEndpointName: "Example AzureRM",
 *     serviceEndpointAuthenticationScheme: "ServicePrincipal",
 *     azurermSpnTenantid: "00000000-0000-0000-0000-000000000000",
 *     azurermSubscriptionId: "00000000-0000-0000-0000-000000000000",
 *     azurermSubscriptionName: "Example Subscription Name",
 * });
 * ```
 * ### Workload Identity Federation Manual AzureRM Service Endpoint (Subscription Scoped)
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const serviceConnectionName = "example-federated-sc";
 * const exampleProject = new azuredevops.Project("exampleProject", {
 *     visibility: "private",
 *     versionControl: "Git",
 *     workItemTemplate: "Agile",
 *     description: "Managed by Terraform",
 * });
 * const identity = new azurerm.index.Azurerm_resource_group("identity", {
 *     name: "identity",
 *     location: "UK South",
 * });
 * const exampleazurerm_user_assigned_identity = new azurerm.index.Azurerm_user_assigned_identity("exampleazurerm_user_assigned_identity", {
 *     location: identity.location,
 *     name: "example-identity",
 *     resourceGroupName: "azurerm_resource_group.identity.name",
 * });
 * const exampleServiceEndpointAzureRM = new azuredevops.ServiceEndpointAzureRM("exampleServiceEndpointAzureRM", {
 *     projectId: exampleProject.id,
 *     serviceEndpointName: serviceConnectionName,
 *     description: "Managed by Terraform",
 *     serviceEndpointAuthenticationScheme: "WorkloadIdentityFederation",
 *     credentials: {
 *         serviceprincipalid: exampleazurerm_user_assigned_identity.clientId,
 *     },
 *     azurermSpnTenantid: "00000000-0000-0000-0000-000000000000",
 *     azurermSubscriptionId: "00000000-0000-0000-0000-000000000000",
 *     azurermSubscriptionName: "Example Subscription Name",
 * });
 * const exampleazurerm_federated_identity_credential = new azurerm.index.Azurerm_federated_identity_credential("exampleazurerm_federated_identity_credential", {
 *     name: "example-federated-credential",
 *     resourceGroupName: identity.name,
 *     parentId: exampleazurerm_user_assigned_identity.id,
 *     audience: ["api://AzureADTokenExchange"],
 *     issuer: exampleServiceEndpointAzureRM.workloadIdentityFederationIssuer,
 *     subject: exampleServiceEndpointAzureRM.workloadIdentityFederationSubject,
 * });
 * ```
 * ### Workload Identity Federation Automatic AzureRM Service Endpoint
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const exampleProject = new azuredevops.Project("exampleProject", {
 *     visibility: "private",
 *     versionControl: "Git",
 *     workItemTemplate: "Agile",
 * });
 * const exampleServiceEndpointAzureRM = new azuredevops.ServiceEndpointAzureRM("exampleServiceEndpointAzureRM", {
 *     projectId: exampleProject.id,
 *     serviceEndpointName: "Example AzureRM",
 *     serviceEndpointAuthenticationScheme: "WorkloadIdentityFederation",
 *     azurermSpnTenantid: "00000000-0000-0000-0000-000000000000",
 *     azurermSubscriptionId: "00000000-0000-0000-0000-000000000000",
 *     azurermSubscriptionName: "Example Subscription Name",
 * });
 * ```
 * ### Managed Identity AzureRM Service Endpoint
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const exampleProject = new azuredevops.Project("exampleProject", {
 *     visibility: "private",
 *     versionControl: "Git",
 *     workItemTemplate: "Agile",
 * });
 * const exampleServiceEndpointAzureRM = new azuredevops.ServiceEndpointAzureRM("exampleServiceEndpointAzureRM", {
 *     projectId: exampleProject.id,
 *     serviceEndpointName: "Example AzureRM",
 *     serviceEndpointAuthenticationScheme: "ManagedServiceIdentity",
 *     azurermSpnTenantid: "00000000-0000-0000-0000-000000000000",
 *     azurermSubscriptionId: "00000000-0000-0000-0000-000000000000",
 *     azurermSubscriptionName: "Example Subscription Name",
 * });
 * ```
 * ## Relevant Links
 *
 * - [Azure DevOps Service REST API 7.0 - Service End points](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-7.0)
 *
 * ## Import
 *
 * Azure DevOps Service Endpoint Azure Resource Manage can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**
 *
 * ```sh
 * $ pulumi import azuredevops:ServiceEndpoint/azureRM:AzureRM example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
 * ```
 *
 * @deprecated azuredevops.serviceendpoint.AzureRM has been deprecated in favor of azuredevops.ServiceEndpointAzureRM
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
        pulumi.log.warn("AzureRM is deprecated: azuredevops.serviceendpoint.AzureRM has been deprecated in favor of azuredevops.ServiceEndpointAzureRM")
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
     * The Management group ID of the Azure targets.
     */
    public readonly azurermManagementGroupId!: pulumi.Output<string | undefined>;
    /**
     * The Management group Name of the targets.
     */
    public readonly azurermManagementGroupName!: pulumi.Output<string | undefined>;
    /**
     * The Tenant ID if the service principal.
     */
    public readonly azurermSpnTenantid!: pulumi.Output<string>;
    /**
     * The Subscription ID of the Azure targets.
     */
    public readonly azurermSubscriptionId!: pulumi.Output<string | undefined>;
    /**
     * The Subscription Name of the targets.
     */
    public readonly azurermSubscriptionName!: pulumi.Output<string | undefined>;
    /**
     * A `credentials` block.
     */
    public readonly credentials!: pulumi.Output<outputs.ServiceEndpoint.AzureRMCredentials | undefined>;
    /**
     * Service connection description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The Cloud Environment to use. Defaults to `AzureCloud`. Possible values are `AzureCloud`, `AzureChinaCloud`. Changing this forces a new resource to be created.
     *
     * > **NOTE:** One of either `Subscription` scoped i.e. `azurermSubscriptionId`, `azurermSubscriptionName` or `ManagementGroup` scoped i.e. `azurermManagementGroupId`, `azurermManagementGroupName` values must be specified.
     */
    public readonly environment!: pulumi.Output<string | undefined>;
    /**
     * A `features` block.
     */
    public readonly features!: pulumi.Output<outputs.ServiceEndpoint.AzureRMFeatures | undefined>;
    /**
     * The ID of the project.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * The resource group used for scope of automatic service endpoint.
     */
    public readonly resourceGroup!: pulumi.Output<string | undefined>;
    /**
     * Specifies the type of azurerm endpoint, either `WorkloadIdentityFederation`, `ManagedServiceIdentity` or `ServicePrincipal`. Defaults to `ServicePrincipal` for backwards compatibility.
     *
     * > **NOTE:** The `WorkloadIdentityFederation` authentication scheme is currently in private preview. Your organisation must be part of the preview and the feature toggle must be turned on to use it. More details can be found [here](https://aka.ms/azdo-rm-workload-identity).
     */
    public readonly serviceEndpointAuthenticationScheme!: pulumi.Output<string | undefined>;
    /**
     * The Service Endpoint Name.
     */
    public readonly serviceEndpointName!: pulumi.Output<string>;
    /**
     * The Application(Client) ID of the Service Principal.
     */
    public /*out*/ readonly servicePrincipalId!: pulumi.Output<string>;
    /**
     * The issuer if `serviceEndpointAuthenticationScheme` is set to `WorkloadIdentityFederation`. This looks like `https://vstoken.dev.azure.com/00000000-0000-0000-0000-000000000000`, where the GUID is the Organization ID of your Azure DevOps Organisation.
     */
    public /*out*/ readonly workloadIdentityFederationIssuer!: pulumi.Output<string>;
    /**
     * The subject if `serviceEndpointAuthenticationScheme` is set to `WorkloadIdentityFederation`. This looks like `sc://<organisation>/<project>/<service-connection-name>`.
     */
    public /*out*/ readonly workloadIdentityFederationSubject!: pulumi.Output<string>;

    /**
     * Create a AzureRM resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated azuredevops.serviceendpoint.AzureRM has been deprecated in favor of azuredevops.ServiceEndpointAzureRM */
    constructor(name: string, args: AzureRMArgs, opts?: pulumi.CustomResourceOptions)
    /** @deprecated azuredevops.serviceendpoint.AzureRM has been deprecated in favor of azuredevops.ServiceEndpointAzureRM */
    constructor(name: string, argsOrState?: AzureRMArgs | AzureRMState, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("AzureRM is deprecated: azuredevops.serviceendpoint.AzureRM has been deprecated in favor of azuredevops.ServiceEndpointAzureRM")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AzureRMState | undefined;
            resourceInputs["authorization"] = state ? state.authorization : undefined;
            resourceInputs["azurermManagementGroupId"] = state ? state.azurermManagementGroupId : undefined;
            resourceInputs["azurermManagementGroupName"] = state ? state.azurermManagementGroupName : undefined;
            resourceInputs["azurermSpnTenantid"] = state ? state.azurermSpnTenantid : undefined;
            resourceInputs["azurermSubscriptionId"] = state ? state.azurermSubscriptionId : undefined;
            resourceInputs["azurermSubscriptionName"] = state ? state.azurermSubscriptionName : undefined;
            resourceInputs["credentials"] = state ? state.credentials : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["environment"] = state ? state.environment : undefined;
            resourceInputs["features"] = state ? state.features : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["resourceGroup"] = state ? state.resourceGroup : undefined;
            resourceInputs["serviceEndpointAuthenticationScheme"] = state ? state.serviceEndpointAuthenticationScheme : undefined;
            resourceInputs["serviceEndpointName"] = state ? state.serviceEndpointName : undefined;
            resourceInputs["servicePrincipalId"] = state ? state.servicePrincipalId : undefined;
            resourceInputs["workloadIdentityFederationIssuer"] = state ? state.workloadIdentityFederationIssuer : undefined;
            resourceInputs["workloadIdentityFederationSubject"] = state ? state.workloadIdentityFederationSubject : undefined;
        } else {
            const args = argsOrState as AzureRMArgs | undefined;
            if ((!args || args.azurermSpnTenantid === undefined) && !opts.urn) {
                throw new Error("Missing required property 'azurermSpnTenantid'");
            }
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            if ((!args || args.serviceEndpointName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceEndpointName'");
            }
            resourceInputs["authorization"] = args ? args.authorization : undefined;
            resourceInputs["azurermManagementGroupId"] = args ? args.azurermManagementGroupId : undefined;
            resourceInputs["azurermManagementGroupName"] = args ? args.azurermManagementGroupName : undefined;
            resourceInputs["azurermSpnTenantid"] = args ? args.azurermSpnTenantid : undefined;
            resourceInputs["azurermSubscriptionId"] = args ? args.azurermSubscriptionId : undefined;
            resourceInputs["azurermSubscriptionName"] = args ? args.azurermSubscriptionName : undefined;
            resourceInputs["credentials"] = args ? args.credentials : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["environment"] = args ? args.environment : undefined;
            resourceInputs["features"] = args ? args.features : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["resourceGroup"] = args ? args.resourceGroup : undefined;
            resourceInputs["serviceEndpointAuthenticationScheme"] = args ? args.serviceEndpointAuthenticationScheme : undefined;
            resourceInputs["serviceEndpointName"] = args ? args.serviceEndpointName : undefined;
            resourceInputs["servicePrincipalId"] = undefined /*out*/;
            resourceInputs["workloadIdentityFederationIssuer"] = undefined /*out*/;
            resourceInputs["workloadIdentityFederationSubject"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AzureRM.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AzureRM resources.
 */
export interface AzureRMState {
    authorization?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The Management group ID of the Azure targets.
     */
    azurermManagementGroupId?: pulumi.Input<string>;
    /**
     * The Management group Name of the targets.
     */
    azurermManagementGroupName?: pulumi.Input<string>;
    /**
     * The Tenant ID if the service principal.
     */
    azurermSpnTenantid?: pulumi.Input<string>;
    /**
     * The Subscription ID of the Azure targets.
     */
    azurermSubscriptionId?: pulumi.Input<string>;
    /**
     * The Subscription Name of the targets.
     */
    azurermSubscriptionName?: pulumi.Input<string>;
    /**
     * A `credentials` block.
     */
    credentials?: pulumi.Input<inputs.ServiceEndpoint.AzureRMCredentials>;
    /**
     * Service connection description.
     */
    description?: pulumi.Input<string>;
    /**
     * The Cloud Environment to use. Defaults to `AzureCloud`. Possible values are `AzureCloud`, `AzureChinaCloud`. Changing this forces a new resource to be created.
     *
     * > **NOTE:** One of either `Subscription` scoped i.e. `azurermSubscriptionId`, `azurermSubscriptionName` or `ManagementGroup` scoped i.e. `azurermManagementGroupId`, `azurermManagementGroupName` values must be specified.
     */
    environment?: pulumi.Input<string>;
    /**
     * A `features` block.
     */
    features?: pulumi.Input<inputs.ServiceEndpoint.AzureRMFeatures>;
    /**
     * The ID of the project.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The resource group used for scope of automatic service endpoint.
     */
    resourceGroup?: pulumi.Input<string>;
    /**
     * Specifies the type of azurerm endpoint, either `WorkloadIdentityFederation`, `ManagedServiceIdentity` or `ServicePrincipal`. Defaults to `ServicePrincipal` for backwards compatibility.
     *
     * > **NOTE:** The `WorkloadIdentityFederation` authentication scheme is currently in private preview. Your organisation must be part of the preview and the feature toggle must be turned on to use it. More details can be found [here](https://aka.ms/azdo-rm-workload-identity).
     */
    serviceEndpointAuthenticationScheme?: pulumi.Input<string>;
    /**
     * The Service Endpoint Name.
     */
    serviceEndpointName?: pulumi.Input<string>;
    /**
     * The Application(Client) ID of the Service Principal.
     */
    servicePrincipalId?: pulumi.Input<string>;
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
 * The set of arguments for constructing a AzureRM resource.
 */
export interface AzureRMArgs {
    authorization?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The Management group ID of the Azure targets.
     */
    azurermManagementGroupId?: pulumi.Input<string>;
    /**
     * The Management group Name of the targets.
     */
    azurermManagementGroupName?: pulumi.Input<string>;
    /**
     * The Tenant ID if the service principal.
     */
    azurermSpnTenantid: pulumi.Input<string>;
    /**
     * The Subscription ID of the Azure targets.
     */
    azurermSubscriptionId?: pulumi.Input<string>;
    /**
     * The Subscription Name of the targets.
     */
    azurermSubscriptionName?: pulumi.Input<string>;
    /**
     * A `credentials` block.
     */
    credentials?: pulumi.Input<inputs.ServiceEndpoint.AzureRMCredentials>;
    /**
     * Service connection description.
     */
    description?: pulumi.Input<string>;
    /**
     * The Cloud Environment to use. Defaults to `AzureCloud`. Possible values are `AzureCloud`, `AzureChinaCloud`. Changing this forces a new resource to be created.
     *
     * > **NOTE:** One of either `Subscription` scoped i.e. `azurermSubscriptionId`, `azurermSubscriptionName` or `ManagementGroup` scoped i.e. `azurermManagementGroupId`, `azurermManagementGroupName` values must be specified.
     */
    environment?: pulumi.Input<string>;
    /**
     * A `features` block.
     */
    features?: pulumi.Input<inputs.ServiceEndpoint.AzureRMFeatures>;
    /**
     * The ID of the project.
     */
    projectId: pulumi.Input<string>;
    /**
     * The resource group used for scope of automatic service endpoint.
     */
    resourceGroup?: pulumi.Input<string>;
    /**
     * Specifies the type of azurerm endpoint, either `WorkloadIdentityFederation`, `ManagedServiceIdentity` or `ServicePrincipal`. Defaults to `ServicePrincipal` for backwards compatibility.
     *
     * > **NOTE:** The `WorkloadIdentityFederation` authentication scheme is currently in private preview. Your organisation must be part of the preview and the feature toggle must be turned on to use it. More details can be found [here](https://aka.ms/azdo-rm-workload-identity).
     */
    serviceEndpointAuthenticationScheme?: pulumi.Input<string>;
    /**
     * The Service Endpoint Name.
     */
    serviceEndpointName: pulumi.Input<string>;
}
