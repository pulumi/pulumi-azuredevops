// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Use this data source to access information about a **single** (existing) Git Repository within Azure DevOps.
 * To read information about **multiple** Git Repositories use the data source `azuredevops.getRepositories`
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const example = azuredevops.getProject({
 *     name: "Example Project",
 * });
 * // Load a specific Git repository by name
 * const example-single-repo = example.then(example => azuredevops.getGitRepository({
 *     projectId: example.id,
 *     name: "Example Repository",
 * }));
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Relevant Links
 *
 * - [Azure DevOps Service REST API 7.0 - Git API](https://docs.microsoft.com/en-us/rest/api/azure/devops/git/?view=azure-devops-rest-7.0)
 */
export function getGitRepository(args: GetGitRepositoryArgs, opts?: pulumi.InvokeOptions): Promise<GetGitRepositoryResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azuredevops:index/getGitRepository:getGitRepository", {
        "name": args.name,
        "projectId": args.projectId,
    }, opts);
}

/**
 * A collection of arguments for invoking getGitRepository.
 */
export interface GetGitRepositoryArgs {
    /**
     * Name of the Git repository to retrieve
     */
    name: string;
    /**
     * ID of project to list Git repositories
     */
    projectId: string;
}

/**
 * A collection of values returned by getGitRepository.
 */
export interface GetGitRepositoryResult {
    /**
     * The ref of the default branch.
     */
    readonly defaultBranch: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly isFork: boolean;
    /**
     * Git repository name.
     */
    readonly name: string;
    /**
     * Project identifier to which the Git repository belongs.
     */
    readonly projectId: string;
    /**
     * HTTPS Url to clone the Git repository
     */
    readonly remoteUrl: string;
    /**
     * Compressed size (bytes) of the repository.
     */
    readonly size: number;
    /**
     * SSH Url to clone the Git repository
     */
    readonly sshUrl: string;
    /**
     * Details REST API endpoint for the Git Repository.
     */
    readonly url: string;
    /**
     * Url of the Git repository web view
     */
    readonly webUrl: string;
}
/**
 * Use this data source to access information about a **single** (existing) Git Repository within Azure DevOps.
 * To read information about **multiple** Git Repositories use the data source `azuredevops.getRepositories`
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const example = azuredevops.getProject({
 *     name: "Example Project",
 * });
 * // Load a specific Git repository by name
 * const example-single-repo = example.then(example => azuredevops.getGitRepository({
 *     projectId: example.id,
 *     name: "Example Repository",
 * }));
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Relevant Links
 *
 * - [Azure DevOps Service REST API 7.0 - Git API](https://docs.microsoft.com/en-us/rest/api/azure/devops/git/?view=azure-devops-rest-7.0)
 */
export function getGitRepositoryOutput(args: GetGitRepositoryOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetGitRepositoryResult> {
    return pulumi.output(args).apply((a: any) => getGitRepository(a, opts))
}

/**
 * A collection of arguments for invoking getGitRepository.
 */
export interface GetGitRepositoryOutputArgs {
    /**
     * Name of the Git repository to retrieve
     */
    name: pulumi.Input<string>;
    /**
     * ID of project to list Git repositories
     */
    projectId: pulumi.Input<string>;
}
