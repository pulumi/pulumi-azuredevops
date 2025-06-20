// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Use this data source to access information about an existing Team in a Project within Azure DevOps.
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
 * const example = azuredevops.getTeamOutput({
 *     projectId: exampleProject.id,
 *     name: "Example Project Team",
 * });
 * ```
 *
 * ## Relevant Links
 *
 * - [Azure DevOps Service REST API 7.0 - Teams - Get](https://docs.microsoft.com/en-us/rest/api/azure/devops/core/teams/get?view=azure-devops-rest-7.0)
 *
 * ## PAT Permissions Required
 *
 * - **vso.project**:	Grants the ability to read projects and teams.
 */
export function getTeam(args: GetTeamArgs, opts?: pulumi.InvokeOptions): Promise<GetTeamResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azuredevops:index/getTeam:getTeam", {
        "name": args.name,
        "projectId": args.projectId,
        "top": args.top,
    }, opts);
}

/**
 * A collection of arguments for invoking getTeam.
 */
export interface GetTeamArgs {
    /**
     * The name of the Team.
     */
    name: string;
    /**
     * The Project ID.
     */
    projectId: string;
    /**
     * The maximum number of teams to return. Defaults to `100`. This property is deprecated and will be removed in the feature
     *
     * @deprecated This property is deprecated and will be removed in the feature
     */
    top?: number;
}

/**
 * A collection of values returned by getTeam.
 */
export interface GetTeamResult {
    /**
     * List of subject descriptors for `administrators` of the team.
     */
    readonly administrators: string[];
    /**
     * The description of the team.
     */
    readonly description: string;
    /**
     * The descriptor of the Team.
     */
    readonly descriptor: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * List of subject descriptors for `members` of the team.
     */
    readonly members: string[];
    readonly name: string;
    readonly projectId: string;
    /**
     * @deprecated This property is deprecated and will be removed in the feature
     */
    readonly top?: number;
}
/**
 * Use this data source to access information about an existing Team in a Project within Azure DevOps.
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
 * const example = azuredevops.getTeamOutput({
 *     projectId: exampleProject.id,
 *     name: "Example Project Team",
 * });
 * ```
 *
 * ## Relevant Links
 *
 * - [Azure DevOps Service REST API 7.0 - Teams - Get](https://docs.microsoft.com/en-us/rest/api/azure/devops/core/teams/get?view=azure-devops-rest-7.0)
 *
 * ## PAT Permissions Required
 *
 * - **vso.project**:	Grants the ability to read projects and teams.
 */
export function getTeamOutput(args: GetTeamOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetTeamResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azuredevops:index/getTeam:getTeam", {
        "name": args.name,
        "projectId": args.projectId,
        "top": args.top,
    }, opts);
}

/**
 * A collection of arguments for invoking getTeam.
 */
export interface GetTeamOutputArgs {
    /**
     * The name of the Team.
     */
    name: pulumi.Input<string>;
    /**
     * The Project ID.
     */
    projectId: pulumi.Input<string>;
    /**
     * The maximum number of teams to return. Defaults to `100`. This property is deprecated and will be removed in the feature
     *
     * @deprecated This property is deprecated and will be removed in the feature
     */
    top?: pulumi.Input<number>;
}
