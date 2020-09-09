// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Use this data source to access information about an existing Project within Azure DevOps.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const project = azuredevops.getProject({
 *     projectName: "Sample Project",
 * });
 * export const id = project.then(project => project.id);
 * export const projectName = project.then(project => project.projectName);
 * export const visibility = project.then(project => project.visibility);
 * export const versionControl = project.then(project => project.versionControl);
 * export const workItemTemplate = project.then(project => project.workItemTemplate);
 * export const processTemplateId = project.then(project => project.processTemplateId);
 * ```
 * ## Relevant Links
 *
 * - [Azure DevOps Service REST API 5.1 - Projects - Get](https://docs.microsoft.com/en-us/rest/api/azure/devops/core/projects/get?view=azure-devops-rest-5.1)
 */
export function getProject(args: GetProjectArgs, opts?: pulumi.InvokeOptions): Promise<GetProjectResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azuredevops:index/getProject:getProject", {
        "projectName": args.projectName,
    }, opts);
}

/**
 * A collection of arguments for invoking getProject.
 */
export interface GetProjectArgs {
    /**
     * Name of the Project.
     */
    readonly projectName: string;
}

/**
 * A collection of values returned by getProject.
 */
export interface GetProjectResult {
    readonly description: string;
    readonly features: {[key: string]: any};
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly processTemplateId: string;
    readonly projectName: string;
    readonly versionControl: string;
    readonly visibility: string;
    readonly workItemTemplate: string;
}
