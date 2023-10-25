// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Use this data source to access information about existing Variable Groups within Azure DevOps.
 *
 * > **Note:** Secret values are masked by service and cannot be obtained through API. [Set secret variables](https://docs.microsoft.com/en-us/azure/devops/pipelines/process/variables?view=azure-devops&tabs=yaml%2Cbatch#secret-variables)
 *
 * ## Relevant Links
 *
 * - [Azure DevOps Service REST API 7.0 - Variable Groups](https://docs.microsoft.com/en-us/rest/api/azure/devops/distributedtask/variablegroups?view=azure-devops-rest-7.0)
 */
export function getVariableGroup(args: GetVariableGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetVariableGroupResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azuredevops:index/getVariableGroup:getVariableGroup", {
        "name": args.name,
        "projectId": args.projectId,
    }, opts);
}

/**
 * A collection of arguments for invoking getVariableGroup.
 */
export interface GetVariableGroupArgs {
    /**
     * The name of the Variable Group to retrieve.
     */
    name: string;
    /**
     * The project ID.
     */
    projectId: string;
}

/**
 * A collection of values returned by getVariableGroup.
 */
export interface GetVariableGroupResult {
    /**
     * Boolean that indicate if this Variable Group is shared by all pipelines of this project.
     */
    readonly allowAccess: boolean;
    /**
     * The description of the Variable Group.
     */
    readonly description: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of `keyVault` blocks as documented below.
     */
    readonly keyVaults: outputs.GetVariableGroupKeyVault[];
    /**
     * The name of the Azure key vault to link secrets from as variables.
     */
    readonly name: string;
    readonly projectId: string;
    /**
     * One or more `variable` blocks as documented below.
     */
    readonly variables: outputs.GetVariableGroupVariable[];
}
/**
 * Use this data source to access information about existing Variable Groups within Azure DevOps.
 *
 * > **Note:** Secret values are masked by service and cannot be obtained through API. [Set secret variables](https://docs.microsoft.com/en-us/azure/devops/pipelines/process/variables?view=azure-devops&tabs=yaml%2Cbatch#secret-variables)
 *
 * ## Relevant Links
 *
 * - [Azure DevOps Service REST API 7.0 - Variable Groups](https://docs.microsoft.com/en-us/rest/api/azure/devops/distributedtask/variablegroups?view=azure-devops-rest-7.0)
 */
export function getVariableGroupOutput(args: GetVariableGroupOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetVariableGroupResult> {
    return pulumi.output(args).apply((a: any) => getVariableGroup(a, opts))
}

/**
 * A collection of arguments for invoking getVariableGroup.
 */
export interface GetVariableGroupOutputArgs {
    /**
     * The name of the Variable Group to retrieve.
     */
    name: pulumi.Input<string>;
    /**
     * The project ID.
     */
    projectId: pulumi.Input<string>;
}
