// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Use this data source to access information about existing Feed within a given project in Azure DevOps.
 *
 * ## Example Usage
 *
 * ### Basic Example
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const example = azuredevops.getFeed({
 *     name: "releases",
 * });
 * ```
 *
 * ### Access feed within a project
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const example = azuredevops.getProject({
 *     name: "Example Project",
 * });
 * const exampleGetFeed = example.then(example => azuredevops.getFeed({
 *     name: "releases",
 *     projectId: example.id,
 * }));
 * ```
 *
 * ## Relevant Links
 *
 * - [Azure DevOps Service REST API 7.0 - Feed - Get](https://learn.microsoft.com/en-us/rest/api/azure/devops/artifacts/feed-management/get-feed?view=azure-devops-rest-7.0)
 */
export function getFeed(args?: GetFeedArgs, opts?: pulumi.InvokeOptions): Promise<GetFeedResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azuredevops:index/getFeed:getFeed", {
        "feedId": args.feedId,
        "name": args.name,
        "projectId": args.projectId,
    }, opts);
}

/**
 * A collection of arguments for invoking getFeed.
 */
export interface GetFeedArgs {
    /**
     * ID of the Feed.
     *
     * > **Note** Only one of `name` or `feedId` can be set at the same time.
     */
    feedId?: string;
    /**
     * Name of the Feed.
     */
    name?: string;
    /**
     * ID of the Project Feed is created in.
     */
    projectId?: string;
}

/**
 * A collection of values returned by getFeed.
 */
export interface GetFeedResult {
    /**
     * The ID of the Feed.
     */
    readonly feedId?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The name of the Feed.
     */
    readonly name?: string;
    /**
     * The ID of the Project.
     */
    readonly projectId?: string;
}
/**
 * Use this data source to access information about existing Feed within a given project in Azure DevOps.
 *
 * ## Example Usage
 *
 * ### Basic Example
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const example = azuredevops.getFeed({
 *     name: "releases",
 * });
 * ```
 *
 * ### Access feed within a project
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const example = azuredevops.getProject({
 *     name: "Example Project",
 * });
 * const exampleGetFeed = example.then(example => azuredevops.getFeed({
 *     name: "releases",
 *     projectId: example.id,
 * }));
 * ```
 *
 * ## Relevant Links
 *
 * - [Azure DevOps Service REST API 7.0 - Feed - Get](https://learn.microsoft.com/en-us/rest/api/azure/devops/artifacts/feed-management/get-feed?view=azure-devops-rest-7.0)
 */
export function getFeedOutput(args?: GetFeedOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFeedResult> {
    return pulumi.output(args).apply((a: any) => getFeed(a, opts))
}

/**
 * A collection of arguments for invoking getFeed.
 */
export interface GetFeedOutputArgs {
    /**
     * ID of the Feed.
     *
     * > **Note** Only one of `name` or `feedId` can be set at the same time.
     */
    feedId?: pulumi.Input<string>;
    /**
     * Name of the Feed.
     */
    name?: pulumi.Input<string>;
    /**
     * ID of the Project Feed is created in.
     */
    projectId?: pulumi.Input<string>;
}