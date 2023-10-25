// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Use this data source to access information about existing Agent Pools within Azure DevOps.
 *
 * ## Relevant Links
 *
 * - [Azure DevOps Service REST API 7.0 - Agent Pools - Get](https://docs.microsoft.com/en-us/rest/api/azure/devops/distributedtask/pools/get?view=azure-devops-rest-7.0)
 */
export function getPools(opts?: pulumi.InvokeOptions): Promise<GetPoolsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azuredevops:index/getPools:getPools", {
    }, opts);
}

/**
 * A collection of values returned by getPools.
 */
export interface GetPoolsResult {
    /**
     * A list of existing agent pools in your Azure DevOps Organization with the following details about every agent pool:
     */
    readonly agentPools: outputs.GetPoolsAgentPool[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
/**
 * Use this data source to access information about existing Agent Pools within Azure DevOps.
 *
 * ## Relevant Links
 *
 * - [Azure DevOps Service REST API 7.0 - Agent Pools - Get](https://docs.microsoft.com/en-us/rest/api/azure/devops/distributedtask/pools/get?view=azure-devops-rest-7.0)
 */
export function getPoolsOutput(opts?: pulumi.InvokeOptions): pulumi.Output<GetPoolsResult> {
    return pulumi.output(getPools(opts))
}
