// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Use this data source to access information about existing Groups within Azure DevOps On-Premise(Azure DevOps Server).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const example = azuredevops.getProject({
 *     name: "Example Project",
 * });
 * // load all existing groups inside an organization
 * const example_all_groups = azuredevops.getIdentityGroups({});
 * // load all existing groups inside a specific project
 * const example_project_groups = example.then(example => azuredevops.getIdentityGroups({
 *     projectId: example.id,
 * }));
 * ```
 *
 * ## Relevant Links
 *
 * - [Azure DevOps Service REST API 7.0 - Identities](https://docs.microsoft.com/en-us/rest/api/azure/devops/ims/?view=azure-devops-rest-7.2)
 */
export function getIdentityGroups(args?: GetIdentityGroupsArgs, opts?: pulumi.InvokeOptions): Promise<GetIdentityGroupsResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azuredevops:index/getIdentityGroups:getIdentityGroups", {
        "projectId": args.projectId,
    }, opts);
}

/**
 * A collection of arguments for invoking getIdentityGroups.
 */
export interface GetIdentityGroupsArgs {
    /**
     * The Project ID. If no project ID is specified all groups of an organization will be returned
     */
    projectId?: string;
}

/**
 * A collection of values returned by getIdentityGroups.
 */
export interface GetIdentityGroupsResult {
    /**
     * A `groups` blocks as documented below. A set of existing groups in your Azure DevOps Organization or project with details about every single group.
     */
    readonly groups: outputs.GetIdentityGroupsGroup[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly projectId?: string;
}
/**
 * Use this data source to access information about existing Groups within Azure DevOps On-Premise(Azure DevOps Server).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const example = azuredevops.getProject({
 *     name: "Example Project",
 * });
 * // load all existing groups inside an organization
 * const example_all_groups = azuredevops.getIdentityGroups({});
 * // load all existing groups inside a specific project
 * const example_project_groups = example.then(example => azuredevops.getIdentityGroups({
 *     projectId: example.id,
 * }));
 * ```
 *
 * ## Relevant Links
 *
 * - [Azure DevOps Service REST API 7.0 - Identities](https://docs.microsoft.com/en-us/rest/api/azure/devops/ims/?view=azure-devops-rest-7.2)
 */
export function getIdentityGroupsOutput(args?: GetIdentityGroupsOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetIdentityGroupsResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azuredevops:index/getIdentityGroups:getIdentityGroups", {
        "projectId": args.projectId,
    }, opts);
}

/**
 * A collection of arguments for invoking getIdentityGroups.
 */
export interface GetIdentityGroupsOutputArgs {
    /**
     * The Project ID. If no project ID is specified all groups of an organization will be returned
     */
    projectId?: pulumi.Input<string>;
}
