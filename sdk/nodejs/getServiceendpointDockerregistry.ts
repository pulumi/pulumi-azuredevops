// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Use this data source to access information about an existing Docker Registry Service Endpoint.
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
 * const exampleGetServiceendpointDockerregistry = example.then(example => azuredevops.getServiceendpointDockerregistry({
 *     projectId: example.id,
 *     serviceEndpointId: "00000000-0000-0000-0000-000000000000",
 * }));
 * export const serviceEndpointName = exampleGetServiceendpointDockerregistry.then(exampleGetServiceendpointDockerregistry => exampleGetServiceendpointDockerregistry.serviceEndpointName);
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
 * const exampleGetServiceendpointDockerregistry = example.then(example => azuredevops.getServiceendpointDockerregistry({
 *     projectId: example.id,
 *     serviceEndpointName: "Example-Service-Endpoint",
 * }));
 * export const serviceEndpointId = serviceendpoint.id;
 * ```
 */
export function getServiceendpointDockerregistry(args: GetServiceendpointDockerregistryArgs, opts?: pulumi.InvokeOptions): Promise<GetServiceendpointDockerregistryResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azuredevops:index/getServiceendpointDockerregistry:getServiceendpointDockerregistry", {
        "projectId": args.projectId,
        "serviceEndpointId": args.serviceEndpointId,
        "serviceEndpointName": args.serviceEndpointName,
    }, opts);
}

/**
 * A collection of arguments for invoking getServiceendpointDockerregistry.
 */
export interface GetServiceendpointDockerregistryArgs {
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
 * A collection of values returned by getServiceendpointDockerregistry.
 */
export interface GetServiceendpointDockerregistryResult {
    /**
     * The Authorization scheme.
     */
    readonly authorization: {[key: string]: string};
    /**
     * The Service Endpoint description.
     */
    readonly description: string;
    /**
     * The email for Docker account user.
     */
    readonly dockerEmail: string;
    /**
     * The password for the account user identified above.
     */
    readonly dockerPassword: string;
    /**
     * The URL of the Docker registry.
     */
    readonly dockerRegistry: string;
    /**
     * The identifier of the Docker account user.
     */
    readonly dockerUsername: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly projectId: string;
    /**
     * Can be "DockerHub" or "Others" (Default "DockerHub")
     */
    readonly registryType: string;
    readonly serviceEndpointId: string;
    readonly serviceEndpointName: string;
}
/**
 * Use this data source to access information about an existing Docker Registry Service Endpoint.
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
 * const exampleGetServiceendpointDockerregistry = example.then(example => azuredevops.getServiceendpointDockerregistry({
 *     projectId: example.id,
 *     serviceEndpointId: "00000000-0000-0000-0000-000000000000",
 * }));
 * export const serviceEndpointName = exampleGetServiceendpointDockerregistry.then(exampleGetServiceendpointDockerregistry => exampleGetServiceendpointDockerregistry.serviceEndpointName);
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
 * const exampleGetServiceendpointDockerregistry = example.then(example => azuredevops.getServiceendpointDockerregistry({
 *     projectId: example.id,
 *     serviceEndpointName: "Example-Service-Endpoint",
 * }));
 * export const serviceEndpointId = serviceendpoint.id;
 * ```
 */
export function getServiceendpointDockerregistryOutput(args: GetServiceendpointDockerregistryOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetServiceendpointDockerregistryResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azuredevops:index/getServiceendpointDockerregistry:getServiceendpointDockerregistry", {
        "projectId": args.projectId,
        "serviceEndpointId": args.serviceEndpointId,
        "serviceEndpointName": args.serviceEndpointName,
    }, opts);
}

/**
 * A collection of arguments for invoking getServiceendpointDockerregistry.
 */
export interface GetServiceendpointDockerregistryOutputArgs {
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
