// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manages permissions for Git repositories.
 *
 * > **Note** Permissions can be assigned to group principals and not to single user principals.
 *
 * ## Permission levels
 *
 * Permission for Git Repositories within Azure DevOps can be applied on three different levels.
 * Those levels are reflected by specifying (or omitting) values for the arguments `projectId`, `repositoryId` and `branchName`.
 *
 * ### Project level
 *
 * Permissions for all Git Repositories inside a project (existing or newly created ones) are specified, if only the argument `projectId` has a value.
 *
 * #### Example usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const example = new azuredevops.Project("example", {
 *     workItemTemplate: "Agile",
 *     versionControl: "Git",
 *     visibility: "private",
 *     description: "Managed by Terraform",
 * });
 * const example-readers = azuredevops.getGroupOutput({
 *     projectId: example.id,
 *     name: "Readers",
 * });
 * const example_permissions = new azuredevops.GitPermissions("example-permissions", {
 *     projectId: example.id,
 *     principal: example_readers.apply(example_readers => example_readers.id),
 *     permissions: {
 *         CreateRepository: "Deny",
 *         DeleteRepository: "Deny",
 *         RenameRepository: "NotSet",
 *     },
 * });
 * ```
 *
 * ### Repository level
 *
 * Permissions for a specific Git Repository and all existing or newly created branches are specified if the arguments `projectId` and `repositoryId` are set.
 *
 * #### Example usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const exampleProject = new azuredevops.Project("exampleProject", {
 *     workItemTemplate: "Agile",
 *     versionControl: "Git",
 *     visibility: "private",
 *     description: "Managed by Terraform",
 * });
 * const example-group = azuredevops.getGroup({
 *     name: "Project Collection Administrators",
 * });
 * const exampleGit = new azuredevops.Git("exampleGit", {
 *     projectId: exampleProject.id,
 *     initialization: {
 *         initType: "Clean",
 *     },
 * });
 * const example_permissions = new azuredevops.GitPermissions("example-permissions", {
 *     projectId: exampleGit.projectId,
 *     repositoryId: exampleGit.id,
 *     principal: example_group.then(example_group => example_group.id),
 *     permissions: {
 *         RemoveOthersLocks: "Allow",
 *         ManagePermissions: "Deny",
 *         CreateTag: "Deny",
 *         CreateBranch: "NotSet",
 *     },
 * });
 * ```
 *
 * ### Branch level
 *
 * Permissions for a specific branch inside a Git Repository are specified if all above mentioned the arguments are set.
 *
 * #### Example usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const exampleProject = new azuredevops.Project("exampleProject", {
 *     workItemTemplate: "Agile",
 *     versionControl: "Git",
 *     visibility: "private",
 *     description: "Managed by Terraform",
 * });
 * const exampleGit = new azuredevops.Git("exampleGit", {
 *     projectId: exampleProject.id,
 *     initialization: {
 *         initType: "Clean",
 *     },
 * });
 * const example-group = azuredevops.getGroup({
 *     name: "Project Collection Administrators",
 * });
 * const example_permissions = new azuredevops.GitPermissions("example-permissions", {
 *     projectId: exampleGit.projectId,
 *     repositoryId: exampleGit.id,
 *     branchName: "refs/heads/master",
 *     principal: example_group.then(example_group => example_group.id),
 *     permissions: {
 *         RemoveOthersLocks: "Allow",
 *         ForcePush: "Deny",
 *     },
 * });
 * ```
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
 *     description: "Managed by Terraform",
 * });
 * const example-project-readers = azuredevops.getGroupOutput({
 *     projectId: exampleProject.id,
 *     name: "Readers",
 * });
 * const example-project-contributors = azuredevops.getGroupOutput({
 *     projectId: exampleProject.id,
 *     name: "Contributors",
 * });
 * const example-project-administrators = azuredevops.getGroupOutput({
 *     projectId: exampleProject.id,
 *     name: "Project administrators",
 * });
 * const example_permissions = new azuredevops.GitPermissions("example-permissions", {
 *     projectId: exampleProject.id,
 *     principal: example_project_readers.apply(example_project_readers => example_project_readers.id),
 *     permissions: {
 *         CreateRepository: "Deny",
 *         DeleteRepository: "Deny",
 *         RenameRepository: "NotSet",
 *     },
 * });
 * const exampleGit = new azuredevops.Git("exampleGit", {
 *     projectId: exampleProject.id,
 *     defaultBranch: "refs/heads/master",
 *     initialization: {
 *         initType: "Clean",
 *     },
 * });
 * const example_repo_permissions = new azuredevops.GitPermissions("example-repo-permissions", {
 *     projectId: exampleGit.projectId,
 *     repositoryId: exampleGit.id,
 *     principal: example_project_administrators.apply(example_project_administrators => example_project_administrators.id),
 *     permissions: {
 *         RemoveOthersLocks: "Allow",
 *         ManagePermissions: "Deny",
 *         CreateTag: "Deny",
 *         CreateBranch: "NotSet",
 *     },
 * });
 * const example_branch_permissions = new azuredevops.GitPermissions("example-branch-permissions", {
 *     projectId: exampleGit.projectId,
 *     repositoryId: exampleGit.id,
 *     branchName: "master",
 *     principal: example_project_contributors.apply(example_project_contributors => example_project_contributors.id),
 *     permissions: {
 *         RemoveOthersLocks: "Allow",
 *         ForcePush: "Deny",
 *     },
 * });
 * ```
 * ## Relevant Links
 *
 * * [Azure DevOps Service REST API 6.0 - Security](https://docs.microsoft.com/en-us/rest/api/azure/devops/security/?view=azure-devops-rest-6.0)
 *
 * ## PAT Permissions Required
 *
 * - **Project & Team**: vso.security_manage - Grants the ability to read, write, and manage security permissions.
 *
 * ## Import
 *
 * The resource does not support import.
 */
export class GitPermissions extends pulumi.CustomResource {
    /**
     * Get an existing GitPermissions resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GitPermissionsState, opts?: pulumi.CustomResourceOptions): GitPermissions {
        return new GitPermissions(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuredevops:index/gitPermissions:GitPermissions';

    /**
     * Returns true if the given object is an instance of GitPermissions.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GitPermissions {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GitPermissions.__pulumiType;
    }

    /**
     * The name of the branch to assign the permissions.
     */
    public readonly branchName!: pulumi.Output<string | undefined>;
    /**
     * the permissions to assign. The follwing permissions are available
     */
    public readonly permissions!: pulumi.Output<{[key: string]: string}>;
    /**
     * The **group** principal to assign the permissions.
     */
    public readonly principal!: pulumi.Output<string>;
    /**
     * The ID of the project to assign the permissions.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * Replace (`true`) or merge (`false`) the permissions. Default: `true`
     */
    public readonly replace!: pulumi.Output<boolean | undefined>;
    /**
     * The ID of the GIT repository to assign the permissions
     */
    public readonly repositoryId!: pulumi.Output<string | undefined>;

    /**
     * Create a GitPermissions resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GitPermissionsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GitPermissionsArgs | GitPermissionsState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GitPermissionsState | undefined;
            resourceInputs["branchName"] = state ? state.branchName : undefined;
            resourceInputs["permissions"] = state ? state.permissions : undefined;
            resourceInputs["principal"] = state ? state.principal : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["replace"] = state ? state.replace : undefined;
            resourceInputs["repositoryId"] = state ? state.repositoryId : undefined;
        } else {
            const args = argsOrState as GitPermissionsArgs | undefined;
            if ((!args || args.permissions === undefined) && !opts.urn) {
                throw new Error("Missing required property 'permissions'");
            }
            if ((!args || args.principal === undefined) && !opts.urn) {
                throw new Error("Missing required property 'principal'");
            }
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            resourceInputs["branchName"] = args ? args.branchName : undefined;
            resourceInputs["permissions"] = args ? args.permissions : undefined;
            resourceInputs["principal"] = args ? args.principal : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["replace"] = args ? args.replace : undefined;
            resourceInputs["repositoryId"] = args ? args.repositoryId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(GitPermissions.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GitPermissions resources.
 */
export interface GitPermissionsState {
    /**
     * The name of the branch to assign the permissions.
     */
    branchName?: pulumi.Input<string>;
    /**
     * the permissions to assign. The follwing permissions are available
     */
    permissions?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The **group** principal to assign the permissions.
     */
    principal?: pulumi.Input<string>;
    /**
     * The ID of the project to assign the permissions.
     */
    projectId?: pulumi.Input<string>;
    /**
     * Replace (`true`) or merge (`false`) the permissions. Default: `true`
     */
    replace?: pulumi.Input<boolean>;
    /**
     * The ID of the GIT repository to assign the permissions
     */
    repositoryId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a GitPermissions resource.
 */
export interface GitPermissionsArgs {
    /**
     * The name of the branch to assign the permissions.
     */
    branchName?: pulumi.Input<string>;
    /**
     * the permissions to assign. The follwing permissions are available
     */
    permissions: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The **group** principal to assign the permissions.
     */
    principal: pulumi.Input<string>;
    /**
     * The ID of the project to assign the permissions.
     */
    projectId: pulumi.Input<string>;
    /**
     * Replace (`true`) or merge (`false`) the permissions. Default: `true`
     */
    replace?: pulumi.Input<boolean>;
    /**
     * The ID of the GIT repository to assign the permissions
     */
    repositoryId?: pulumi.Input<string>;
}
