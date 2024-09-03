// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manages permissions for an Area (Component)
 *
 * > **Note** Permissions can be assigned to group principals and not to single user principals.
 *
 * ## Permission levels
 *
 * Permission for Areas within Azure DevOps can be applied on two different levels.
 * Those levels are reflected by specifying (or omitting) values for the arguments `projectId` and `path`.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const example = new azuredevops.Project("example", {
 *     name: "Example Project",
 *     workItemTemplate: "Agile",
 *     versionControl: "Git",
 *     visibility: "private",
 *     description: "Managed by Terraform",
 * });
 * const example-project-readers = azuredevops.getGroupOutput({
 *     projectId: example.id,
 *     name: "Readers",
 * });
 * const example_root_permissions = new azuredevops.AreaPermissions("example-root-permissions", {
 *     projectId: example.id,
 *     principal: example_project_readers.apply(example_project_readers => example_project_readers.id),
 *     path: "/",
 *     permissions: {
 *         CREATE_CHILDREN: "Deny",
 *         GENERIC_READ: "Allow",
 *         DELETE: "Deny",
 *         WORK_ITEM_READ: "Allow",
 *     },
 * });
 * ```
 *
 * ## Relevant Links
 *
 * * [Azure DevOps Service REST API 7.0 - Security](https://docs.microsoft.com/en-us/rest/api/azure/devops/security/?view=azure-devops-rest-7.0)
 *
 * ## PAT Permissions Required
 *
 * - **Project & Team**: vso.security_manage - Grants the ability to read, write, and manage security permissions.
 *
 * ## Import
 *
 * The resource does not support import.
 */
export class AreaPermissions extends pulumi.CustomResource {
    /**
     * Get an existing AreaPermissions resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AreaPermissionsState, opts?: pulumi.CustomResourceOptions): AreaPermissions {
        return new AreaPermissions(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuredevops:index/areaPermissions:AreaPermissions';

    /**
     * Returns true if the given object is an instance of AreaPermissions.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AreaPermissions {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AreaPermissions.__pulumiType;
    }

    /**
     * The name of the branch to assign the permissions.
     */
    public readonly path!: pulumi.Output<string | undefined>;
    /**
     * the permissions to assign. The following permissions are available.
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
     * Replace (`true`) or merge (`false`) the permissions. Default: `true`.
     * <table>
     * <thead>
     * <tr>
     * <th>Permission</th>
     * <th>Description</th>
     * </tr>
     * </thead>
     * <tbody>
     * <tr>
     * <td>GENERIC_READ</td>
     * <td>View permissions for this node</td>
     * </tr>
     * <tr>
     * <td>GENERIC_WRITE</td>
     * <td>Edit this node</td>
     * </tr>
     * <tr>
     * <td>CREATE_CHILDREN</td>
     * <td>Create child nodes</td>
     * </tr>
     * <tr>
     * <td>DELETE</td>
     * <td>Delete this node</td>
     * </tr>
     * <tr>
     * <td>WORK_ITEM_READ</td>
     * <td>View work items in this node</td>
     * </tr>
     * <tr>
     * <td>WORK_ITEM_WRITE</td>
     * <td>Edit work items in this node</td>
     * </tr>
     * <tr>
     * <td>MANAGE_TEST_PLANS</td>
     * <td>Manage test plans</td>
     * </tr>
     * <tr>
     * <td>MANAGE_TEST_SUITES</td>
     * <td>Manage test suites</td>
     * </tr>
     * <tr>
     * <td>WORK_ITEM_SAVE_COMMENT</td>
     * <td>Edit work item comments in this node</td>
     * </tr>
     * </tbody>
     * </table>
     */
    public readonly replace!: pulumi.Output<boolean | undefined>;

    /**
     * Create a AreaPermissions resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AreaPermissionsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AreaPermissionsArgs | AreaPermissionsState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AreaPermissionsState | undefined;
            resourceInputs["path"] = state ? state.path : undefined;
            resourceInputs["permissions"] = state ? state.permissions : undefined;
            resourceInputs["principal"] = state ? state.principal : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["replace"] = state ? state.replace : undefined;
        } else {
            const args = argsOrState as AreaPermissionsArgs | undefined;
            if ((!args || args.permissions === undefined) && !opts.urn) {
                throw new Error("Missing required property 'permissions'");
            }
            if ((!args || args.principal === undefined) && !opts.urn) {
                throw new Error("Missing required property 'principal'");
            }
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            resourceInputs["path"] = args ? args.path : undefined;
            resourceInputs["permissions"] = args ? args.permissions : undefined;
            resourceInputs["principal"] = args ? args.principal : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["replace"] = args ? args.replace : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AreaPermissions.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AreaPermissions resources.
 */
export interface AreaPermissionsState {
    /**
     * The name of the branch to assign the permissions.
     */
    path?: pulumi.Input<string>;
    /**
     * the permissions to assign. The following permissions are available.
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
     * Replace (`true`) or merge (`false`) the permissions. Default: `true`.
     * <table>
     * <thead>
     * <tr>
     * <th>Permission</th>
     * <th>Description</th>
     * </tr>
     * </thead>
     * <tbody>
     * <tr>
     * <td>GENERIC_READ</td>
     * <td>View permissions for this node</td>
     * </tr>
     * <tr>
     * <td>GENERIC_WRITE</td>
     * <td>Edit this node</td>
     * </tr>
     * <tr>
     * <td>CREATE_CHILDREN</td>
     * <td>Create child nodes</td>
     * </tr>
     * <tr>
     * <td>DELETE</td>
     * <td>Delete this node</td>
     * </tr>
     * <tr>
     * <td>WORK_ITEM_READ</td>
     * <td>View work items in this node</td>
     * </tr>
     * <tr>
     * <td>WORK_ITEM_WRITE</td>
     * <td>Edit work items in this node</td>
     * </tr>
     * <tr>
     * <td>MANAGE_TEST_PLANS</td>
     * <td>Manage test plans</td>
     * </tr>
     * <tr>
     * <td>MANAGE_TEST_SUITES</td>
     * <td>Manage test suites</td>
     * </tr>
     * <tr>
     * <td>WORK_ITEM_SAVE_COMMENT</td>
     * <td>Edit work item comments in this node</td>
     * </tr>
     * </tbody>
     * </table>
     */
    replace?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a AreaPermissions resource.
 */
export interface AreaPermissionsArgs {
    /**
     * The name of the branch to assign the permissions.
     */
    path?: pulumi.Input<string>;
    /**
     * the permissions to assign. The following permissions are available.
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
     * Replace (`true`) or merge (`false`) the permissions. Default: `true`.
     * <table>
     * <thead>
     * <tr>
     * <th>Permission</th>
     * <th>Description</th>
     * </tr>
     * </thead>
     * <tbody>
     * <tr>
     * <td>GENERIC_READ</td>
     * <td>View permissions for this node</td>
     * </tr>
     * <tr>
     * <td>GENERIC_WRITE</td>
     * <td>Edit this node</td>
     * </tr>
     * <tr>
     * <td>CREATE_CHILDREN</td>
     * <td>Create child nodes</td>
     * </tr>
     * <tr>
     * <td>DELETE</td>
     * <td>Delete this node</td>
     * </tr>
     * <tr>
     * <td>WORK_ITEM_READ</td>
     * <td>View work items in this node</td>
     * </tr>
     * <tr>
     * <td>WORK_ITEM_WRITE</td>
     * <td>Edit work items in this node</td>
     * </tr>
     * <tr>
     * <td>MANAGE_TEST_PLANS</td>
     * <td>Manage test plans</td>
     * </tr>
     * <tr>
     * <td>MANAGE_TEST_SUITES</td>
     * <td>Manage test suites</td>
     * </tr>
     * <tr>
     * <td>WORK_ITEM_SAVE_COMMENT</td>
     * <td>Edit work item comments in this node</td>
     * </tr>
     * </tbody>
     * </table>
     */
    replace?: pulumi.Input<boolean>;
}
