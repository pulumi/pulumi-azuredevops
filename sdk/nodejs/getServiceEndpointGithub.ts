// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Use this data source to access information about an existing GitHub service Endpoint.
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
 * const serviceendpoint = sample.then(sample => azuredevops.getServiceEndpointGithub({
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
 * const serviceendpoint = sample.then(sample => azuredevops.getServiceEndpointGithub({
 *     projectId: sample.id,
 *     serviceEndpointName: "Example-Service-Endpoint",
 * }));
 * export const serviceEndpointId = serviceendpoint.then(serviceendpoint => serviceendpoint.id);
 * ```
 */
export function getServiceEndpointGithub(args: GetServiceEndpointGithubArgs, opts?: pulumi.InvokeOptions): Promise<GetServiceEndpointGithubResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azuredevops:index/getServiceEndpointGithub:getServiceEndpointGithub", {
        "projectId": args.projectId,
        "serviceEndpointId": args.serviceEndpointId,
        "serviceEndpointName": args.serviceEndpointName,
    }, opts);
}

/**
 * A collection of arguments for invoking getServiceEndpointGithub.
 */
export interface GetServiceEndpointGithubArgs {
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
 * A collection of values returned by getServiceEndpointGithub.
 */
export interface GetServiceEndpointGithubResult {
    /**
     * Specifies the Authorization Scheme Map.
     */
    readonly authorization: {[key: string]: string};
    /**
     * Specifies the description of the Service Endpoint.
     */
    readonly description: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly projectId: string;
    readonly serviceEndpointId: string;
    readonly serviceEndpointName: string;
}
/**
 * Use this data source to access information about an existing GitHub service Endpoint.
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
 * const serviceendpoint = sample.then(sample => azuredevops.getServiceEndpointGithub({
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
 * const serviceendpoint = sample.then(sample => azuredevops.getServiceEndpointGithub({
 *     projectId: sample.id,
 *     serviceEndpointName: "Example-Service-Endpoint",
 * }));
 * export const serviceEndpointId = serviceendpoint.then(serviceendpoint => serviceendpoint.id);
 * ```
 */
export function getServiceEndpointGithubOutput(args: GetServiceEndpointGithubOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetServiceEndpointGithubResult> {
    return pulumi.output(args).apply((a: any) => getServiceEndpointGithub(a, opts))
}

/**
 * A collection of arguments for invoking getServiceEndpointGithub.
 */
export interface GetServiceEndpointGithubOutputArgs {
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