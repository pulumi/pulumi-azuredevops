// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Use this data source to access information about an existing AzureRM service Endpoint.
 *
 * ## Example Usage
 * ### By Service Endpoint ID
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const sample = azuredevops.getProject({
 *     name: "Sample Project",
 * });
 * const serviceendpoint = sample.then(sample => azuredevops.getServiceEndpointAzureRM({
 *     projectId: sample.id,
 *     serviceEndpointId: "00000000-0000-0000-0000-000000000000",
 * }));
 * export const serviceEndpointName = serviceendpoint.then(serviceendpoint => serviceendpoint.serviceEndpointName);
 * ```
 * ### By Service Endpoint Name
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const sample = azuredevops.getProject({
 *     name: "Sample Project",
 * });
 * const serviceendpoint = sample.then(sample => azuredevops.getServiceEndpointAzureRM({
 *     projectId: sample.id,
 *     serviceEndpointName: "Example-Service-Endpoint",
 * }));
 * export const serviceEndpointId = serviceendpoint.then(serviceendpoint => serviceendpoint.id);
 * ```
 */
export function getServiceEndpointAzureRM(args: GetServiceEndpointAzureRMArgs, opts?: pulumi.InvokeOptions): Promise<GetServiceEndpointAzureRMResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("azuredevops:index/getServiceEndpointAzureRM:getServiceEndpointAzureRM", {
        "projectId": args.projectId,
        "serviceEndpointId": args.serviceEndpointId,
        "serviceEndpointName": args.serviceEndpointName,
    }, opts);
}

/**
 * A collection of arguments for invoking getServiceEndpointAzureRM.
 */
export interface GetServiceEndpointAzureRMArgs {
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
     */
    serviceEndpointName?: string;
}

/**
 * A collection of values returned by getServiceEndpointAzureRM.
 */
export interface GetServiceEndpointAzureRMResult {
    /**
     * Specifies the Authorization Scheme Map.
     */
    readonly authorization: {[key: string]: string};
    /**
     * Specified the Management Group ID of the Service Endpoint is target, if available.
     */
    readonly azurermManagementGroupId: string;
    /**
     * Specified the Management Group Name of the Service Endpoint target, if available.
     */
    readonly azurermManagementGroupName: string;
    /**
     * Specifies the Tenant ID of the Azure targets.
     */
    readonly azurermSpnTenantid: string;
    /**
     * Specifies the Subscription ID of the Service Endpoint target, if available.
     */
    readonly azurermSubscriptionId: string;
    /**
     * Specifies the Subscription Name of the Service Endpoint target, if available.
     */
    readonly azurermSubscriptionName: string;
    /**
     * Specifies the description of the Service Endpoint.
     */
    readonly description: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly projectId: string;
    /**
     * Specifies the Resource Group of the Service Endpoint target, if available.
     */
    readonly resourceGroup: string;
    readonly serviceEndpointId: string;
    readonly serviceEndpointName: string;
}

export function getServiceEndpointAzureRMOutput(args: GetServiceEndpointAzureRMOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetServiceEndpointAzureRMResult> {
    return pulumi.output(args).apply(a => getServiceEndpointAzureRM(a, opts))
}

/**
 * A collection of arguments for invoking getServiceEndpointAzureRM.
 */
export interface GetServiceEndpointAzureRMOutputArgs {
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
     */
    serviceEndpointName?: pulumi.Input<string>;
}
