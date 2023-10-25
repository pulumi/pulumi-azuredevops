// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Use this data source to access information about the Azure DevOps organization configured for the provider.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const example = azuredevops.getClientConfig({});
 * export const orgUrl = example.then(example => example.organizationUrl);
 * ```
 */
export function getClientConfig(opts?: pulumi.InvokeOptions): Promise<GetClientConfigResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azuredevops:index/getClientConfig:getClientConfig", {
    }, opts);
}

/**
 * A collection of values returned by getClientConfig.
 */
export interface GetClientConfigResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly organizationUrl: string;
}
/**
 * Use this data source to access information about the Azure DevOps organization configured for the provider.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const example = azuredevops.getClientConfig({});
 * export const orgUrl = example.then(example => example.organizationUrl);
 * ```
 */
export function getClientConfigOutput(opts?: pulumi.InvokeOptions): pulumi.Output<GetClientConfigResult> {
    return pulumi.output(getClientConfig(opts))
}
