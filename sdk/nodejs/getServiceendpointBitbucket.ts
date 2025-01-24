// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Use this data source to access information about an existing Bitbucket service Endpoint.
 *
 * ## Example Usage
 *
 * ### By Service Endpoint ID
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const example = azuredevops.getProject({
 *     name: "Example Project",
 * });
 * const exampleGetServiceendpointBitbucket = example.then(example => azuredevops.getServiceendpointBitbucket({
 *     projectId: example.id,
 *     serviceEndpointId: "00000000-0000-0000-0000-000000000000",
 * }));
 * export const serviceEndpointName = exampleGetServiceendpointBitbucket.then(exampleGetServiceendpointBitbucket => exampleGetServiceendpointBitbucket.serviceEndpointName);
 * ```
 *
 * ### By Service Endpoint Name
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const example = azuredevops.getProject({
 *     name: "Example Project",
 * });
 * const exampleGetServiceendpointBitbucket = example.then(example => azuredevops.getServiceendpointBitbucket({
 *     projectId: example.id,
 *     serviceEndpointName: "Example",
 * }));
 * export const serviceEndpointId = exampleGetServiceendpointBitbucket.then(exampleGetServiceendpointBitbucket => exampleGetServiceendpointBitbucket.id);
 * ```
 *
 * ## PAT Permissions Required
 *
 * - **vso.serviceendpoint**: Grants the ability to read service endpoints.
 */
export function getServiceendpointBitbucket(args: GetServiceendpointBitbucketArgs, opts?: pulumi.InvokeOptions): Promise<GetServiceendpointBitbucketResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azuredevops:index/getServiceendpointBitbucket:getServiceendpointBitbucket", {
        "projectId": args.projectId,
        "serviceEndpointId": args.serviceEndpointId,
        "serviceEndpointName": args.serviceEndpointName,
    }, opts);
}

/**
 * A collection of arguments for invoking getServiceendpointBitbucket.
 */
export interface GetServiceendpointBitbucketArgs {
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
     * > **NOTE:** 1. One of either `serviceEndpointId` or `serviceEndpointName` must be specified.
     * <br>2. When supplying `serviceEndpointName`, take care to ensure that this is a unique name.
     */
    serviceEndpointName?: string;
}

/**
 * A collection of values returned by getServiceendpointBitbucket.
 */
export interface GetServiceendpointBitbucketResult {
    /**
     * The Authorization scheme.
     */
    readonly authorization: {[key: string]: string};
    /**
     * The description of the Service Endpoint.
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
 * Use this data source to access information about an existing Bitbucket service Endpoint.
 *
 * ## Example Usage
 *
 * ### By Service Endpoint ID
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const example = azuredevops.getProject({
 *     name: "Example Project",
 * });
 * const exampleGetServiceendpointBitbucket = example.then(example => azuredevops.getServiceendpointBitbucket({
 *     projectId: example.id,
 *     serviceEndpointId: "00000000-0000-0000-0000-000000000000",
 * }));
 * export const serviceEndpointName = exampleGetServiceendpointBitbucket.then(exampleGetServiceendpointBitbucket => exampleGetServiceendpointBitbucket.serviceEndpointName);
 * ```
 *
 * ### By Service Endpoint Name
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const example = azuredevops.getProject({
 *     name: "Example Project",
 * });
 * const exampleGetServiceendpointBitbucket = example.then(example => azuredevops.getServiceendpointBitbucket({
 *     projectId: example.id,
 *     serviceEndpointName: "Example",
 * }));
 * export const serviceEndpointId = exampleGetServiceendpointBitbucket.then(exampleGetServiceendpointBitbucket => exampleGetServiceendpointBitbucket.id);
 * ```
 *
 * ## PAT Permissions Required
 *
 * - **vso.serviceendpoint**: Grants the ability to read service endpoints.
 */
export function getServiceendpointBitbucketOutput(args: GetServiceendpointBitbucketOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetServiceendpointBitbucketResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azuredevops:index/getServiceendpointBitbucket:getServiceendpointBitbucket", {
        "projectId": args.projectId,
        "serviceEndpointId": args.serviceEndpointId,
        "serviceEndpointName": args.serviceEndpointName,
    }, opts);
}

/**
 * A collection of arguments for invoking getServiceendpointBitbucket.
 */
export interface GetServiceendpointBitbucketOutputArgs {
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
     * > **NOTE:** 1. One of either `serviceEndpointId` or `serviceEndpointName` must be specified.
     * <br>2. When supplying `serviceEndpointName`, take care to ensure that this is a unique name.
     */
    serviceEndpointName?: pulumi.Input<string>;
}
