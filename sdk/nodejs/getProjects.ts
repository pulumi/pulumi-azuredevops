// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Use this data source to access information about existing Projects within Azure DevOps.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const example = azuredevops.getProjects({
 *     name: "Example Project",
 *     state: "wellFormed",
 * });
 * export const projectId = example.then(example => example.projects.map(__item => __item.projectId));
 * export const name = example.then(example => example.projects.map(__item => __item.name));
 * export const projectUrl = example.then(example => example.projects.map(__item => __item.projectUrl));
 * export const state = example.then(example => example.projects.map(__item => __item.state));
 * ```
 * ## Relevant Links
 *
 * - [Azure DevOps Service REST API 7.0 - Projects - Get](https://docs.microsoft.com/en-us/rest/api/azure/devops/core/projects/get?view=azure-devops-rest-7.0)
 */
export function getProjects(args?: GetProjectsArgs, opts?: pulumi.InvokeOptions): Promise<GetProjectsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azuredevops:index/getProjects:getProjects", {
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
     *
     * DataSource without specifying any arguments will return all projects.
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
     * The name of the Project.
     */
    readonly name?: string;
    /**
     * A list of existing projects in your Azure DevOps Organization with details about every project which includes:
     */
    readonly projects: outputs.GetProjectsProject[];
    /**
     * Project state.
     */
    readonly state?: string;
}
/**
 * Use this data source to access information about existing Projects within Azure DevOps.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const example = azuredevops.getProjects({
 *     name: "Example Project",
 *     state: "wellFormed",
 * });
 * export const projectId = example.then(example => example.projects.map(__item => __item.projectId));
 * export const name = example.then(example => example.projects.map(__item => __item.name));
 * export const projectUrl = example.then(example => example.projects.map(__item => __item.projectUrl));
 * export const state = example.then(example => example.projects.map(__item => __item.state));
 * ```
 * ## Relevant Links
 *
 * - [Azure DevOps Service REST API 7.0 - Projects - Get](https://docs.microsoft.com/en-us/rest/api/azure/devops/core/projects/get?view=azure-devops-rest-7.0)
 */
export function getProjectsOutput(args?: GetProjectsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetProjectsResult> {
    return pulumi.output(args).apply((a: any) => getProjects(a, opts))
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
     *
     * DataSource without specifying any arguments will return all projects.
     */
    state?: pulumi.Input<string>;
}
