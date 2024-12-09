// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Use this data source to access information about an existing Group within Azure DevOps
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
 * const exampleGetGroup = example.then(example => azuredevops.getGroup({
 *     projectId: example.id,
 *     name: "Example Group",
 * }));
 * export const groupId = exampleGetGroup.then(exampleGetGroup => exampleGetGroup.id);
 * export const groupDescriptor = exampleGetGroup.then(exampleGetGroup => exampleGetGroup.descriptor);
 * const example-collection-group = azuredevops.getGroup({
 *     name: "Project Collection Administrators",
 * });
 * export const collectionGroupId = exampleGetGroup.then(exampleGetGroup => exampleGetGroup.id);
 * export const collectionGroupDescriptor = exampleGetGroup.then(exampleGetGroup => exampleGetGroup.descriptor);
 * ```
 *
 * ## Relevant Links
 *
 * - [Azure DevOps Service REST API 7.0 - Groups - Get](https://docs.microsoft.com/en-us/rest/api/azure/devops/graph/groups/get?view=azure-devops-rest-7.0)
 *
 * ## PAT Permissions Required
 *
 * - **Graph**: Read
 * - **Work Items**: Read
 */
export function getGroup(args: GetGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetGroupResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azuredevops:index/getGroup:getGroup", {
        "name": args.name,
        "projectId": args.projectId,
    }, opts);
}

/**
 * A collection of arguments for invoking getGroup.
 */
export interface GetGroupArgs {
    /**
     * The Group Name.
     */
    name: string;
    /**
     * The Project ID. If no project ID is specified the project collection groups will be searched.
     */
    projectId?: string;
}

/**
 * A collection of values returned by getGroup.
 */
export interface GetGroupResult {
    /**
     * The Descriptor is the primary way to reference the graph subject. This field will uniquely identify the same graph subject across both Accounts and Organizations.
     */
    readonly descriptor: string;
    /**
     * The ID of the group.
     */
    readonly groupId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    /**
     * The type of source provider for the origin identifier (ex:AD, AAD, MSA)
     */
    readonly origin: string;
    /**
     * The unique identifier from the system of origin. Typically a sid, object id or Guid. Linking and unlinking operations can cause this value to change for a user because the user is not backed by a different provider and has a different unique id in the new provider.
     */
    readonly originId: string;
    readonly projectId?: string;
}
/**
 * Use this data source to access information about an existing Group within Azure DevOps
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
 * const exampleGetGroup = example.then(example => azuredevops.getGroup({
 *     projectId: example.id,
 *     name: "Example Group",
 * }));
 * export const groupId = exampleGetGroup.then(exampleGetGroup => exampleGetGroup.id);
 * export const groupDescriptor = exampleGetGroup.then(exampleGetGroup => exampleGetGroup.descriptor);
 * const example-collection-group = azuredevops.getGroup({
 *     name: "Project Collection Administrators",
 * });
 * export const collectionGroupId = exampleGetGroup.then(exampleGetGroup => exampleGetGroup.id);
 * export const collectionGroupDescriptor = exampleGetGroup.then(exampleGetGroup => exampleGetGroup.descriptor);
 * ```
 *
 * ## Relevant Links
 *
 * - [Azure DevOps Service REST API 7.0 - Groups - Get](https://docs.microsoft.com/en-us/rest/api/azure/devops/graph/groups/get?view=azure-devops-rest-7.0)
 *
 * ## PAT Permissions Required
 *
 * - **Graph**: Read
 * - **Work Items**: Read
 */
export function getGroupOutput(args: GetGroupOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetGroupResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azuredevops:index/getGroup:getGroup", {
        "name": args.name,
        "projectId": args.projectId,
    }, opts);
}

/**
 * A collection of arguments for invoking getGroup.
 */
export interface GetGroupOutputArgs {
    /**
     * The Group Name.
     */
    name: pulumi.Input<string>;
    /**
     * The Project ID. If no project ID is specified the project collection groups will be searched.
     */
    projectId?: pulumi.Input<string>;
}
