// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manages permissions for a AzureDevOps project
 *
 * > **Note** Permissions can be assigned to group principals and not to single user principals.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const example = new azuredevops.Project("example", {
 *     visibility: "private",
 *     versionControl: "Git",
 *     workItemTemplate: "Agile",
 *     description: "Managed by Terraform",
 * });
 * const example-readers = azuredevops.getGroupOutput({
 *     projectId: example.id,
 *     name: "Readers",
 * });
 * const example_permission = new azuredevops.ProjectPermissions("example-permission", {
 *     projectId: example.id,
 *     principal: example_readers.apply(example_readers => example_readers.id),
 *     permissions: {
 *         DELETE: "Deny",
 *         EDIT_BUILD_STATUS: "NotSet",
 *         WORK_ITEM_MOVE: "Allow",
 *         DELETE_TEST_RESULTS: "Deny",
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
export class ProjectPermissions extends pulumi.CustomResource {
    /**
     * Get an existing ProjectPermissions resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProjectPermissionsState, opts?: pulumi.CustomResourceOptions): ProjectPermissions {
        return new ProjectPermissions(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuredevops:index/projectPermissions:ProjectPermissions';

    /**
     * Returns true if the given object is an instance of ProjectPermissions.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ProjectPermissions {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ProjectPermissions.__pulumiType;
    }

    /**
     * the permissions to assign. The following permissions are available
     *
     * | Permission                   | Description                                  |
     * |------------------------------|----------------------------------------------|
     * | GENERIC_READ                 | View project-level information               |
     * | GENERIC_WRITE                | Edit project-level information               |
     * | DELETE                       | Delete team project                          |
     * | PUBLISH_TEST_RESULTS         | Create test runs                             |
     * | ADMINISTER_BUILD             | Administer a build                           |
     * | START_BUILD                  | Start a build                                |
     * | EDIT_BUILD_STATUS            | Edit build quality                           |
     * | UPDATE_BUILD                 | Write to build operational store             |
     * | DELETE_TEST_RESULTS          | Delete test runs                             |
     * | VIEW_TEST_RESULTS            | View test runs                               |
     * | MANAGE_TEST_ENVIRONMENTS     | Manage test environments                     |
     * | MANAGE_TEST_CONFIGURATIONS   | Manage test configurations                   |
     * | WORK_ITEM_DELETE             | Delete and restore work items                |
     * | WORK_ITEM_MOVE               | Move work items out of this project          |
     * | WORK_ITEM_PERMANENTLY_DELETE | Permanently delete work items                |
     * | RENAME                       | Rename team project                          |
     * | MANAGE_PROPERTIES            | Manage project properties                    |
     * | MANAGE_SYSTEM_PROPERTIES     | Manage system project properties             |
     * | BYPASS_PROPERTY_CACHE        | Bypass project property cache                |
     * | BYPASS_RULES                 | Bypass rules on work item updates            |
     * | SUPPRESS_NOTIFICATIONS       | Suppress notifications for work item updates |
     * | UPDATE_VISIBILITY            | Update project visibility                    |
     * | CHANGE_PROCESS               | Change process of team project.              |
     * | AGILETOOLS_BACKLOG           | Agile backlog management.                    |
     * | AGILETOOLS_PLANS             | Agile plans.                                 |
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
     * Create a ProjectPermissions resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProjectPermissionsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProjectPermissionsArgs | ProjectPermissionsState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ProjectPermissionsState | undefined;
            resourceInputs["permissions"] = state ? state.permissions : undefined;
            resourceInputs["principal"] = state ? state.principal : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["replace"] = state ? state.replace : undefined;
        } else {
            const args = argsOrState as ProjectPermissionsArgs | undefined;
            if ((!args || args.permissions === undefined) && !opts.urn) {
                throw new Error("Missing required property 'permissions'");
            }
            if ((!args || args.principal === undefined) && !opts.urn) {
                throw new Error("Missing required property 'principal'");
            }
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            resourceInputs["permissions"] = args ? args.permissions : undefined;
            resourceInputs["principal"] = args ? args.principal : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["replace"] = args ? args.replace : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ProjectPermissions.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ProjectPermissions resources.
 */
export interface ProjectPermissionsState {
    /**
     * the permissions to assign. The following permissions are available
     *
     * | Permission                   | Description                                  |
     * |------------------------------|----------------------------------------------|
     * | GENERIC_READ                 | View project-level information               |
     * | GENERIC_WRITE                | Edit project-level information               |
     * | DELETE                       | Delete team project                          |
     * | PUBLISH_TEST_RESULTS         | Create test runs                             |
     * | ADMINISTER_BUILD             | Administer a build                           |
     * | START_BUILD                  | Start a build                                |
     * | EDIT_BUILD_STATUS            | Edit build quality                           |
     * | UPDATE_BUILD                 | Write to build operational store             |
     * | DELETE_TEST_RESULTS          | Delete test runs                             |
     * | VIEW_TEST_RESULTS            | View test runs                               |
     * | MANAGE_TEST_ENVIRONMENTS     | Manage test environments                     |
     * | MANAGE_TEST_CONFIGURATIONS   | Manage test configurations                   |
     * | WORK_ITEM_DELETE             | Delete and restore work items                |
     * | WORK_ITEM_MOVE               | Move work items out of this project          |
     * | WORK_ITEM_PERMANENTLY_DELETE | Permanently delete work items                |
     * | RENAME                       | Rename team project                          |
     * | MANAGE_PROPERTIES            | Manage project properties                    |
     * | MANAGE_SYSTEM_PROPERTIES     | Manage system project properties             |
     * | BYPASS_PROPERTY_CACHE        | Bypass project property cache                |
     * | BYPASS_RULES                 | Bypass rules on work item updates            |
     * | SUPPRESS_NOTIFICATIONS       | Suppress notifications for work item updates |
     * | UPDATE_VISIBILITY            | Update project visibility                    |
     * | CHANGE_PROCESS               | Change process of team project.              |
     * | AGILETOOLS_BACKLOG           | Agile backlog management.                    |
     * | AGILETOOLS_PLANS             | Agile plans.                                 |
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
}

/**
 * The set of arguments for constructing a ProjectPermissions resource.
 */
export interface ProjectPermissionsArgs {
    /**
     * the permissions to assign. The following permissions are available
     *
     * | Permission                   | Description                                  |
     * |------------------------------|----------------------------------------------|
     * | GENERIC_READ                 | View project-level information               |
     * | GENERIC_WRITE                | Edit project-level information               |
     * | DELETE                       | Delete team project                          |
     * | PUBLISH_TEST_RESULTS         | Create test runs                             |
     * | ADMINISTER_BUILD             | Administer a build                           |
     * | START_BUILD                  | Start a build                                |
     * | EDIT_BUILD_STATUS            | Edit build quality                           |
     * | UPDATE_BUILD                 | Write to build operational store             |
     * | DELETE_TEST_RESULTS          | Delete test runs                             |
     * | VIEW_TEST_RESULTS            | View test runs                               |
     * | MANAGE_TEST_ENVIRONMENTS     | Manage test environments                     |
     * | MANAGE_TEST_CONFIGURATIONS   | Manage test configurations                   |
     * | WORK_ITEM_DELETE             | Delete and restore work items                |
     * | WORK_ITEM_MOVE               | Move work items out of this project          |
     * | WORK_ITEM_PERMANENTLY_DELETE | Permanently delete work items                |
     * | RENAME                       | Rename team project                          |
     * | MANAGE_PROPERTIES            | Manage project properties                    |
     * | MANAGE_SYSTEM_PROPERTIES     | Manage system project properties             |
     * | BYPASS_PROPERTY_CACHE        | Bypass project property cache                |
     * | BYPASS_RULES                 | Bypass rules on work item updates            |
     * | SUPPRESS_NOTIFICATIONS       | Suppress notifications for work item updates |
     * | UPDATE_VISIBILITY            | Update project visibility                    |
     * | CHANGE_PROCESS               | Change process of team project.              |
     * | AGILETOOLS_BACKLOG           | Agile backlog management.                    |
     * | AGILETOOLS_PLANS             | Agile plans.                                 |
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
}
