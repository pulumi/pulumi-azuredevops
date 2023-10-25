// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing Agent Pool within Azure DevOps.
 *
 * ## Relevant Links
 *
 * - [Azure DevOps Service REST API 7.0 - Agent Pools - Get](https://docs.microsoft.com/en-us/rest/api/azure/devops/distributedtask/pools/get?view=azure-devops-rest-7.0)
 */
/** @deprecated azuredevops.agent.getPool has been deprecated in favor of azuredevops.getPool */
export function getPool(args: GetPoolArgs, opts?: pulumi.InvokeOptions): Promise<GetPoolResult> {
    pulumi.log.warn("getPool is deprecated: azuredevops.agent.getPool has been deprecated in favor of azuredevops.getPool")

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azuredevops:Agent/getPool:getPool", {
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getPool.
 */
export interface GetPoolArgs {
    /**
     * Name of the Agent Pool.
     */
    name: string;
}

/**
 * A collection of values returned by getPool.
 */
export interface GetPoolResult {
    readonly autoProvision: boolean;
    readonly autoUpdate: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    readonly poolType: string;
}
/**
 * Use this data source to access information about an existing Agent Pool within Azure DevOps.
 *
 * ## Relevant Links
 *
 * - [Azure DevOps Service REST API 7.0 - Agent Pools - Get](https://docs.microsoft.com/en-us/rest/api/azure/devops/distributedtask/pools/get?view=azure-devops-rest-7.0)
 */
/** @deprecated azuredevops.agent.getPool has been deprecated in favor of azuredevops.getPool */
export function getPoolOutput(args: GetPoolOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPoolResult> {
    return pulumi.output(args).apply((a: any) => getPool(a, opts))
}

/**
 * A collection of arguments for invoking getPool.
 */
export interface GetPoolOutputArgs {
    /**
     * Name of the Agent Pool.
     */
    name: pulumi.Input<string>;
}
