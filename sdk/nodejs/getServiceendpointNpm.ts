// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Use this data source to access information about an existing NPM Service Endpoint.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const example = azuredevops.getServiceendpointNpm({
 *     projectId: azuredevops_project.example.id,
 *     serviceEndpointName: "Example npm",
 * });
 * export const serviceEndpointId = example.then(example => example.id);
 * ```
 */
export function getServiceendpointNpm(args: GetServiceendpointNpmArgs, opts?: pulumi.InvokeOptions): Promise<GetServiceendpointNpmResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azuredevops:index/getServiceendpointNpm:getServiceendpointNpm", {
        "projectId": args.projectId,
        "serviceEndpointId": args.serviceEndpointId,
        "serviceEndpointName": args.serviceEndpointName,
    }, opts);
}

/**
 * A collection of arguments for invoking getServiceendpointNpm.
 */
export interface GetServiceendpointNpmArgs {
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
 * A collection of values returned by getServiceendpointNpm.
 */
export interface GetServiceendpointNpmResult {
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
    /**
     * Specifies the URL of the npm registry to connect with.
     */
    readonly url: string;
}
/**
 * Use this data source to access information about an existing NPM Service Endpoint.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const example = azuredevops.getServiceendpointNpm({
 *     projectId: azuredevops_project.example.id,
 *     serviceEndpointName: "Example npm",
 * });
 * export const serviceEndpointId = example.then(example => example.id);
 * ```
 */
export function getServiceendpointNpmOutput(args: GetServiceendpointNpmOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetServiceendpointNpmResult> {
    return pulumi.output(args).apply((a: any) => getServiceendpointNpm(a, opts))
}

/**
 * A collection of arguments for invoking getServiceendpointNpm.
 */
export interface GetServiceendpointNpmOutputArgs {
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
