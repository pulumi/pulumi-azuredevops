// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Use this data source to access information about an existing Build Definition.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const exampleProject = azuredevops.getProject({
 *     name: "Example Project",
 * });
 * const exampleBuildDefinition = exampleProject.then(exampleProject => azuredevops.getBuildDefinition({
 *     projectId: exampleProject.id,
 *     name: "existing",
 * }));
 * export const id = exampleBuildDefinition.then(exampleBuildDefinition => exampleBuildDefinition.id);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getBuildDefinition(args: GetBuildDefinitionArgs, opts?: pulumi.InvokeOptions): Promise<GetBuildDefinitionResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azuredevops:index/getBuildDefinition:getBuildDefinition", {
        "name": args.name,
        "path": args.path,
        "projectId": args.projectId,
    }, opts);
}

/**
 * A collection of arguments for invoking getBuildDefinition.
 */
export interface GetBuildDefinitionArgs {
    /**
     * The name of this Build Definition.
     */
    name: string;
    /**
     * The path of the build definition. Default to `\`.
     */
    path?: string;
    /**
     * The ID of the project.
     */
    projectId: string;
}

/**
 * A collection of values returned by getBuildDefinition.
 */
export interface GetBuildDefinitionResult {
    /**
     * The agent pool that should execute the build.
     */
    readonly agentPoolName: string;
    /**
     * A `ciTrigger` block as defined below.
     */
    readonly ciTriggers: outputs.GetBuildDefinitionCiTrigger[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The name of the variable.
     */
    readonly name: string;
    readonly path?: string;
    readonly projectId: string;
    /**
     * A `pullRequestTrigger` block as defined below.
     */
    readonly pullRequestTriggers: outputs.GetBuildDefinitionPullRequestTrigger[];
    /**
     * The queue status of the build definition.
     */
    readonly queueStatus: string;
    /**
     * A `repository` block as defined below.
     */
    readonly repositories: outputs.GetBuildDefinitionRepository[];
    /**
     * The revision of the build definition.
     */
    readonly revision: number;
    /**
     * A `schedules` block as defined below.
     */
    readonly schedules: outputs.GetBuildDefinitionSchedule[];
    /**
     * A list of variable group IDs.
     */
    readonly variableGroups: number[];
    /**
     * A `variable` block as defined below.
     */
    readonly variables: outputs.GetBuildDefinitionVariable[];
}
/**
 * Use this data source to access information about an existing Build Definition.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const exampleProject = azuredevops.getProject({
 *     name: "Example Project",
 * });
 * const exampleBuildDefinition = exampleProject.then(exampleProject => azuredevops.getBuildDefinition({
 *     projectId: exampleProject.id,
 *     name: "existing",
 * }));
 * export const id = exampleBuildDefinition.then(exampleBuildDefinition => exampleBuildDefinition.id);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getBuildDefinitionOutput(args: GetBuildDefinitionOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetBuildDefinitionResult> {
    return pulumi.output(args).apply((a: any) => getBuildDefinition(a, opts))
}

/**
 * A collection of arguments for invoking getBuildDefinition.
 */
export interface GetBuildDefinitionOutputArgs {
    /**
     * The name of this Build Definition.
     */
    name: pulumi.Input<string>;
    /**
     * The path of the build definition. Default to `\`.
     */
    path?: pulumi.Input<string>;
    /**
     * The ID of the project.
     */
    projectId: pulumi.Input<string>;
}
