// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Use this data source to access information about an existing Area (Component) within Azure DevOps.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const exampleProject = new azuredevops.Project("example", {
 *     name: "Example Project",
 *     workItemTemplate: "Agile",
 *     versionControl: "Git",
 *     visibility: "private",
 *     description: "Managed by Pulumi",
 * });
 * const example = exampleProject.id.apply(id => azuredevops.getAreaOutput({
 *     projectId: id,
 *     path: "/",
 *     fetchChildren: false,
 * }));
 * ```
 *
 * ## Relevant Links
 *
 * - [Azure DevOps Service REST API 7.0 - Classification Nodes - Get Classification Nodes](https://docs.microsoft.com/en-us/rest/api/azure/devops/wit/classification-nodes/create-or-update?view=azure-devops-rest-7.0)
 *
 * ## PAT Permissions Required
 *
 * - **Project & Team**: vso.work - Grants the ability to read work items, queries, boards, area and iterations paths, and other work item tracking related metadata. Also grants the ability to execute queries, search work items and to receive notifications about work item events via service hooks.
 */
export function getArea(args: GetAreaArgs, opts?: pulumi.InvokeOptions): Promise<GetAreaResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azuredevops:index/getArea:getArea", {
        "fetchChildren": args.fetchChildren,
        "path": args.path,
        "projectId": args.projectId,
    }, opts);
}

/**
 * A collection of arguments for invoking getArea.
 */
export interface GetAreaArgs {
    /**
     * Read children nodes, _Depth_: 1, _Default_: `true`
     */
    fetchChildren?: boolean;
    /**
     * The path to the Area; _Format_: URL relative; if omitted, or value `"/"` is used, the root Area will be returned
     */
    path?: string;
    /**
     * The project ID.
     */
    projectId: string;
}

/**
 * A collection of values returned by getArea.
 */
export interface GetAreaResult {
    /**
     * A list of `children` blocks as defined below, empty if `hasChildren == false`
     */
    readonly childrens: outputs.GetAreaChildren[];
    readonly fetchChildren?: boolean;
    /**
     * Indicator if the child Area node has child nodes
     */
    readonly hasChildren: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The name of the child Area node
     */
    readonly name: string;
    /**
     * The complete path (in relative URL format) of the child Area
     */
    readonly path: string;
    /**
     * The ID of project.
     */
    readonly projectId: string;
}
/**
 * Use this data source to access information about an existing Area (Component) within Azure DevOps.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const exampleProject = new azuredevops.Project("example", {
 *     name: "Example Project",
 *     workItemTemplate: "Agile",
 *     versionControl: "Git",
 *     visibility: "private",
 *     description: "Managed by Pulumi",
 * });
 * const example = exampleProject.id.apply(id => azuredevops.getAreaOutput({
 *     projectId: id,
 *     path: "/",
 *     fetchChildren: false,
 * }));
 * ```
 *
 * ## Relevant Links
 *
 * - [Azure DevOps Service REST API 7.0 - Classification Nodes - Get Classification Nodes](https://docs.microsoft.com/en-us/rest/api/azure/devops/wit/classification-nodes/create-or-update?view=azure-devops-rest-7.0)
 *
 * ## PAT Permissions Required
 *
 * - **Project & Team**: vso.work - Grants the ability to read work items, queries, boards, area and iterations paths, and other work item tracking related metadata. Also grants the ability to execute queries, search work items and to receive notifications about work item events via service hooks.
 */
export function getAreaOutput(args: GetAreaOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetAreaResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azuredevops:index/getArea:getArea", {
        "fetchChildren": args.fetchChildren,
        "path": args.path,
        "projectId": args.projectId,
    }, opts);
}

/**
 * A collection of arguments for invoking getArea.
 */
export interface GetAreaOutputArgs {
    /**
     * Read children nodes, _Depth_: 1, _Default_: `true`
     */
    fetchChildren?: pulumi.Input<boolean>;
    /**
     * The path to the Area; _Format_: URL relative; if omitted, or value `"/"` is used, the root Area will be returned
     */
    path?: pulumi.Input<string>;
    /**
     * The project ID.
     */
    projectId: pulumi.Input<string>;
}
