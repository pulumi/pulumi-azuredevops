// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Use this data source to access information about an existing Azure Container Registry Service Endpoint.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const example = azuredevops.getServiceendpointAzurecr({
 *     projectId: exampleAzuredevopsProject.id,
 *     serviceEndpointName: "Example Azure Container Registry",
 * });
 * export const serviceEndpointId = example.then(example => example.id);
 * ```
 */
export function getServiceendpointAzurecr(args: GetServiceendpointAzurecrArgs, opts?: pulumi.InvokeOptions): Promise<GetServiceendpointAzurecrResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azuredevops:index/getServiceendpointAzurecr:getServiceendpointAzurecr", {
        "projectId": args.projectId,
        "serviceEndpointId": args.serviceEndpointId,
        "serviceEndpointName": args.serviceEndpointName,
    }, opts);
}

/**
 * A collection of arguments for invoking getServiceendpointAzurecr.
 */
export interface GetServiceendpointAzurecrArgs {
    /**
     * The ID of the project.
     */
    projectId: string;
    /**
     * the ID of the Service Endpoint.
     */
    serviceEndpointId?: string;
    /**
     * the Name of the Service Endpoint.
     *
     * > **NOTE:** One of either `serviceEndpointId` or `serviceEndpointName` must be specified.
     */
    serviceEndpointName?: string;
}

/**
 * A collection of values returned by getServiceendpointAzurecr.
 */
export interface GetServiceendpointAzurecrResult {
    /**
     * The Object ID of the Service Principal.
     */
    readonly appObjectId: string;
    /**
     * Specifies the Authorization Scheme Map.
     */
    readonly authorization: {[key: string]: string};
    /**
     * The ID of Service Principal Role Assignment.
     */
    readonly azSpnRoleAssignmentId: string;
    /**
     * The Service Principal Role Permissions.
     */
    readonly azSpnRolePermissions: string;
    /**
     * The Azure Container Registry name.
     */
    readonly azurecrName: string;
    /**
     * The Tenant ID of the service principal.
     */
    readonly azurecrSpnTenantid: string;
    /**
     * The Subscription ID of the Azure targets.
     */
    readonly azurecrSubscriptionId: string;
    /**
     * The Subscription Name of the Azure targets.
     */
    readonly azurecrSubscriptionName: string;
    /**
     * The Service Endpoint description.
     */
    readonly description: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly projectId: string;
    /**
     * The Resource Group to which the Container Registry belongs.
     */
    readonly resourceGroup: string;
    readonly serviceEndpointId: string;
    readonly serviceEndpointName: string;
    /**
     * The Application(Client) ID of the Service Principal.
     */
    readonly servicePrincipalId: string;
    /**
     * The ID of the Service Principal.
     */
    readonly spnObjectId: string;
}
/**
 * Use this data source to access information about an existing Azure Container Registry Service Endpoint.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const example = azuredevops.getServiceendpointAzurecr({
 *     projectId: exampleAzuredevopsProject.id,
 *     serviceEndpointName: "Example Azure Container Registry",
 * });
 * export const serviceEndpointId = example.then(example => example.id);
 * ```
 */
export function getServiceendpointAzurecrOutput(args: GetServiceendpointAzurecrOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetServiceendpointAzurecrResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azuredevops:index/getServiceendpointAzurecr:getServiceendpointAzurecr", {
        "projectId": args.projectId,
        "serviceEndpointId": args.serviceEndpointId,
        "serviceEndpointName": args.serviceEndpointName,
    }, opts);
}

/**
 * A collection of arguments for invoking getServiceendpointAzurecr.
 */
export interface GetServiceendpointAzurecrOutputArgs {
    /**
     * The ID of the project.
     */
    projectId: pulumi.Input<string>;
    /**
     * the ID of the Service Endpoint.
     */
    serviceEndpointId?: pulumi.Input<string>;
    /**
     * the Name of the Service Endpoint.
     *
     * > **NOTE:** One of either `serviceEndpointId` or `serviceEndpointName` must be specified.
     */
    serviceEndpointName?: pulumi.Input<string>;
}
