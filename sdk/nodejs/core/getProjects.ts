// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about existing Projects within Azure DevOps.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const test = azuredevops.getProjects({
 *     name: "contoso",
 *     state: "wellFormed",
 * });
 * export const projectId = [test.then(test => test.projects)].map(__item => __item?.projectId);
 * export const name = [test.then(test => test.projects)].map(__item => __item?.name);
 * export const projectUrl = [test.then(test => test.projects)].map(__item => __item?.projectUrl);
 * export const state = [test.then(test => test.projects)].map(__item => __item?.state);
 * ```
 * ## Relevant Links
 *
 * - [Azure DevOps Service REST API 5.1 - Projects - Get](https://docs.microsoft.com/en-us/rest/api/azure/devops/core/projects/get?view=azure-devops-rest-5.1)
 */
/** @deprecated azuredevops.core.getProjects has been deprecated in favor of azuredevops.getProjects */
export function getProjects(args?: GetProjectsArgs, opts?: pulumi.InvokeOptions): Promise<GetProjectsResult> {
    pulumi.log.warn("getProjects is deprecated: azuredevops.core.getProjects has been deprecated in favor of azuredevops.getProjects")
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("azuredevops:Core/getProjects:getProjects", {
        "name": args.name,
        "state": args.state,
    }, opts);
}

/**
 * A collection of arguments for invoking getProjects.
 */
export interface GetProjectsArgs {
    /**
     * Name of the Project, if not specified all projects will be returned.
     */
    name?: string;
    /**
     * State of the Project, if not specified all projects will be returned. Valid values are `all`, `deleting`, `new`, `wellFormed`, `createPending`, `unchanged`,`deleted`.
     */
    state?: string;
}

/**
 * A collection of values returned by getProjects.
 */
export interface GetProjectsResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Project name.
     */
    readonly name?: string;
    /**
     * A list of existing projects in your Azure DevOps Organization with details about every project which includes:
     */
    readonly projects: outputs.Core.GetProjectsProject[];
    /**
     * Project state.
     */
    readonly state?: string;
}

export function getProjectsOutput(args?: GetProjectsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetProjectsResult> {
    return pulumi.output(args).apply(a => getProjects(a, opts))
}

/**
 * A collection of arguments for invoking getProjects.
 */
export interface GetProjectsOutputArgs {
    /**
     * Name of the Project, if not specified all projects will be returned.
     */
    name?: pulumi.Input<string>;
    /**
     * State of the Project, if not specified all projects will be returned. Valid values are `all`, `deleting`, `new`, `wellFormed`, `createPending`, `unchanged`,`deleted`.
     */
    state?: pulumi.Input<string>;
}
