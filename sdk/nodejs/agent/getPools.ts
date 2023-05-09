// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about existing Agent Pools within Azure DevOps.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const example = azuredevops.getPools({});
 * export const agentPoolName = example.then(example => example.agentPools.map(__item => __item.name));
 * export const autoProvision = example.then(example => example.agentPools.map(__item => __item.autoProvision));
 * export const autoUpdate = example.then(example => example.agentPools.map(__item => __item.autoUpdate));
 * export const poolType = example.then(example => example.agentPools.map(__item => __item.poolType));
 * ```
 * ## Relevant Links
 *
 * - [Azure DevOps Service REST API 6.0 - Agent Pools - Get](https://docs.microsoft.com/en-us/rest/api/azure/devops/distributedtask/pools/get?view=azure-devops-rest-6.0)
 */
/** @deprecated azuredevops.agent.getPools has been deprecated in favor of azuredevops.getPools */
export function getPools(opts?: pulumi.InvokeOptions): Promise<GetPoolsResult> {
    pulumi.log.warn("getPools is deprecated: azuredevops.agent.getPools has been deprecated in favor of azuredevops.getPools")

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azuredevops:Agent/getPools:getPools", {
    }, opts);
}

/**
 * A collection of values returned by getPools.
 */
export interface GetPoolsResult {
    /**
     * A list of existing agent pools in your Azure DevOps Organization with the following details about every agent pool:
     */
    readonly agentPools: outputs.Agent.GetPoolsAgentPool[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
