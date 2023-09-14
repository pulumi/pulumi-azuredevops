// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manage files within an Azure DevOps Git repository.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const exampleProject = new azuredevops.Project("exampleProject", {
 *     visibility: "private",
 *     versionControl: "Git",
 *     workItemTemplate: "Agile",
 * });
 * const exampleGit = new azuredevops.Git("exampleGit", {
 *     projectId: exampleProject.id,
 *     initialization: {
 *         initType: "Clean",
 *     },
 * });
 * const exampleGitRepositoryFile = new azuredevops.GitRepositoryFile("exampleGitRepositoryFile", {
 *     repositoryId: exampleGit.id,
 *     file: ".gitignore",
 *     content: "**&#47;*.tfstate",
 *     branch: "refs/heads/master",
 *     commitMessage: "First commit",
 *     overwriteOnCreate: false,
 * });
 * ```
 * ## Relevant Links
 *
 * - [Azure DevOps Service REST API 7.0 - Git API](https://docs.microsoft.com/en-us/rest/api/azure/devops/git/?view=azure-devops-rest-7.0)
 *
 * ## Import
 *
 * Repository files can be imported using a combination of the `repository ID` and `file`, e.g.
 *
 * ```sh
 *  $ pulumi import azuredevops:index/gitRepositoryFile:GitRepositoryFile example 00000000-0000-0000-0000-000000000000/.gitignore
 * ```
 *
 *  To import a file from a branch other than `master`, append `:` and the branch name, e.g.
 *
 * ```sh
 *  $ pulumi import azuredevops:index/gitRepositoryFile:GitRepositoryFile example 00000000-0000-0000-0000-000000000000/.gitignore:refs/heads/master
 * ```
 */
export class GitRepositoryFile extends pulumi.CustomResource {
    /**
     * Get an existing GitRepositoryFile resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GitRepositoryFileState, opts?: pulumi.CustomResourceOptions): GitRepositoryFile {
        return new GitRepositoryFile(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuredevops:index/gitRepositoryFile:GitRepositoryFile';

    /**
     * Returns true if the given object is an instance of GitRepositoryFile.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GitRepositoryFile {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GitRepositoryFile.__pulumiType;
    }

    /**
     * Git branch (defaults to `refs/heads/master`). The branch must already exist, it will not be created if it
     * does not already exist.
     */
    public readonly branch!: pulumi.Output<string | undefined>;
    /**
     * Commit message when adding or updating the managed file.
     */
    public readonly commitMessage!: pulumi.Output<string>;
    /**
     * The file content.
     */
    public readonly content!: pulumi.Output<string>;
    /**
     * The path of the file to manage.
     */
    public readonly file!: pulumi.Output<string>;
    /**
     * Enable overwriting existing files (defaults to `false`).
     */
    public readonly overwriteOnCreate!: pulumi.Output<boolean | undefined>;
    /**
     * The ID of the Git repository.
     */
    public readonly repositoryId!: pulumi.Output<string>;

    /**
     * Create a GitRepositoryFile resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GitRepositoryFileArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GitRepositoryFileArgs | GitRepositoryFileState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GitRepositoryFileState | undefined;
            resourceInputs["branch"] = state ? state.branch : undefined;
            resourceInputs["commitMessage"] = state ? state.commitMessage : undefined;
            resourceInputs["content"] = state ? state.content : undefined;
            resourceInputs["file"] = state ? state.file : undefined;
            resourceInputs["overwriteOnCreate"] = state ? state.overwriteOnCreate : undefined;
            resourceInputs["repositoryId"] = state ? state.repositoryId : undefined;
        } else {
            const args = argsOrState as GitRepositoryFileArgs | undefined;
            if ((!args || args.content === undefined) && !opts.urn) {
                throw new Error("Missing required property 'content'");
            }
            if ((!args || args.file === undefined) && !opts.urn) {
                throw new Error("Missing required property 'file'");
            }
            if ((!args || args.repositoryId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'repositoryId'");
            }
            resourceInputs["branch"] = args ? args.branch : undefined;
            resourceInputs["commitMessage"] = args ? args.commitMessage : undefined;
            resourceInputs["content"] = args ? args.content : undefined;
            resourceInputs["file"] = args ? args.file : undefined;
            resourceInputs["overwriteOnCreate"] = args ? args.overwriteOnCreate : undefined;
            resourceInputs["repositoryId"] = args ? args.repositoryId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(GitRepositoryFile.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GitRepositoryFile resources.
 */
export interface GitRepositoryFileState {
    /**
     * Git branch (defaults to `refs/heads/master`). The branch must already exist, it will not be created if it
     * does not already exist.
     */
    branch?: pulumi.Input<string>;
    /**
     * Commit message when adding or updating the managed file.
     */
    commitMessage?: pulumi.Input<string>;
    /**
     * The file content.
     */
    content?: pulumi.Input<string>;
    /**
     * The path of the file to manage.
     */
    file?: pulumi.Input<string>;
    /**
     * Enable overwriting existing files (defaults to `false`).
     */
    overwriteOnCreate?: pulumi.Input<boolean>;
    /**
     * The ID of the Git repository.
     */
    repositoryId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a GitRepositoryFile resource.
 */
export interface GitRepositoryFileArgs {
    /**
     * Git branch (defaults to `refs/heads/master`). The branch must already exist, it will not be created if it
     * does not already exist.
     */
    branch?: pulumi.Input<string>;
    /**
     * Commit message when adding or updating the managed file.
     */
    commitMessage?: pulumi.Input<string>;
    /**
     * The file content.
     */
    content: pulumi.Input<string>;
    /**
     * The path of the file to manage.
     */
    file: pulumi.Input<string>;
    /**
     * Enable overwriting existing files (defaults to `false`).
     */
    overwriteOnCreate?: pulumi.Input<boolean>;
    /**
     * The ID of the Git repository.
     */
    repositoryId: pulumi.Input<string>;
}
