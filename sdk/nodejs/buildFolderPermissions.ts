// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manages permissions for a Build Folder
 *
 * > **Note** Permissions can be assigned to group principals and not to single user principals.
 *
 * ## Example Usage
 *
 * ### Set specific folder permissions
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
 *     description: "Managed by Pulumi",
 * });
 * const example_readers = azuredevops.getGroupOutput({
 *     projectId: example.id,
 *     name: "Readers",
 * });
 * const exampleBuildFolder = new azuredevops.BuildFolder("example", {
 *     projectId: example.id,
 *     path: "\\ExampleFolder",
 *     description: "ExampleFolder description",
 * });
 * const exampleBuildFolderPermissions = new azuredevops.BuildFolderPermissions("example", {
 *     projectId: example.id,
 *     path: "\\ExampleFolder",
 *     principal: example_readers.apply(example_readers => example_readers.id),
 *     permissions: {
 *         ViewBuilds: "Allow",
 *         EditBuildQuality: "Allow",
 *         RetainIndefinitely: "Allow",
 *         DeleteBuilds: "Deny",
 *         ManageBuildQualities: "Deny",
 *         DestroyBuilds: "Deny",
 *         UpdateBuildInformation: "Deny",
 *         QueueBuilds: "Allow",
 *         ManageBuildQueue: "Deny",
 *         StopBuilds: "Allow",
 *         ViewBuildDefinition: "Allow",
 *         EditBuildDefinition: "Deny",
 *         DeleteBuildDefinition: "Deny",
 *         AdministerBuildPermissions: "NotSet",
 *     },
 * });
 * ```
 * ### Set root folder permissions
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const example = new azuredevops.Project("example", {
 *     name: "Example Project",
 *     workItemTemplate: "Agile",
 *     versionControl: "Git",
 *     visibility: "private",
 *     description: "Managed by Pulumi",
 * });
 * const example_readers = azuredevops.getGroupOutput({
 *     projectId: example.id,
 *     name: "Readers",
 * });
 * const exampleBuildFolderPermissions = new azuredevops.BuildFolderPermissions("example", {
 *     projectId: example.id,
 *     path: "\\",
 *     principal: example_readers.apply(example_readers => example_readers.id),
 *     permissions: {
 *         RetainIndefinitely: "Allow",
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
export class BuildFolderPermissions extends pulumi.CustomResource {
    /**
     * Get an existing BuildFolderPermissions resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BuildFolderPermissionsState, opts?: pulumi.CustomResourceOptions): BuildFolderPermissions {
        return new BuildFolderPermissions(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuredevops:index/buildFolderPermissions:BuildFolderPermissions';

    /**
     * Returns true if the given object is an instance of BuildFolderPermissions.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BuildFolderPermissions {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BuildFolderPermissions.__pulumiType;
    }

    /**
     * The folder path to assign the permissions.
     */
    public readonly path!: pulumi.Output<string>;
    /**
     * the permissions to assign. The following permissions are available.
     *
     * | Permission                     | Description                           |
     * |--------------------------------|---------------------------------------|
     * | ViewBuilds                     | View builds                           |
     * | EditBuildQuality               | Edit build quality                    |
     * | RetainIndefinitely             | Retain indefinitely                   |
     * | DeleteBuilds                   | Delete builds                         |
     * | ManageBuildQualities           | Manage build qualities                |
     * | DestroyBuilds                  | Destroy builds                        |
     * | UpdateBuildInformation         | Update build information              |
     * | QueueBuilds                    | Queue builds                          |
     * | ManageBuildQueue               | Manage build queue                    |
     * | StopBuilds                     | Stop builds                           |
     * | ViewBuildDefinition            | View build pipeline                   |
     * | EditBuildDefinition            | Edit build pipeline                   |
     * | DeleteBuildDefinition          | Delete build pipeline                 |
     * | OverrideBuildCheckInValidation | Override check-in validation by build |
     * | AdministerBuildPermissions     | Administer build permissions          |
     * | CreateBuildDefinition          | Create build pipeline                 |
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
     */
    public readonly replace!: pulumi.Output<boolean | undefined>;

    /**
     * Create a BuildFolderPermissions resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BuildFolderPermissionsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BuildFolderPermissionsArgs | BuildFolderPermissionsState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BuildFolderPermissionsState | undefined;
            resourceInputs["path"] = state ? state.path : undefined;
            resourceInputs["permissions"] = state ? state.permissions : undefined;
            resourceInputs["principal"] = state ? state.principal : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["replace"] = state ? state.replace : undefined;
        } else {
            const args = argsOrState as BuildFolderPermissionsArgs | undefined;
            if ((!args || args.path === undefined) && !opts.urn) {
                throw new Error("Missing required property 'path'");
            }
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
        super(BuildFolderPermissions.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BuildFolderPermissions resources.
 */
export interface BuildFolderPermissionsState {
    /**
     * The folder path to assign the permissions.
     */
    path?: pulumi.Input<string>;
    /**
     * the permissions to assign. The following permissions are available.
     *
     * | Permission                     | Description                           |
     * |--------------------------------|---------------------------------------|
     * | ViewBuilds                     | View builds                           |
     * | EditBuildQuality               | Edit build quality                    |
     * | RetainIndefinitely             | Retain indefinitely                   |
     * | DeleteBuilds                   | Delete builds                         |
     * | ManageBuildQualities           | Manage build qualities                |
     * | DestroyBuilds                  | Destroy builds                        |
     * | UpdateBuildInformation         | Update build information              |
     * | QueueBuilds                    | Queue builds                          |
     * | ManageBuildQueue               | Manage build queue                    |
     * | StopBuilds                     | Stop builds                           |
     * | ViewBuildDefinition            | View build pipeline                   |
     * | EditBuildDefinition            | Edit build pipeline                   |
     * | DeleteBuildDefinition          | Delete build pipeline                 |
     * | OverrideBuildCheckInValidation | Override check-in validation by build |
     * | AdministerBuildPermissions     | Administer build permissions          |
     * | CreateBuildDefinition          | Create build pipeline                 |
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
     */
    replace?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a BuildFolderPermissions resource.
 */
export interface BuildFolderPermissionsArgs {
    /**
     * The folder path to assign the permissions.
     */
    path: pulumi.Input<string>;
    /**
     * the permissions to assign. The following permissions are available.
     *
     * | Permission                     | Description                           |
     * |--------------------------------|---------------------------------------|
     * | ViewBuilds                     | View builds                           |
     * | EditBuildQuality               | Edit build quality                    |
     * | RetainIndefinitely             | Retain indefinitely                   |
     * | DeleteBuilds                   | Delete builds                         |
     * | ManageBuildQualities           | Manage build qualities                |
     * | DestroyBuilds                  | Destroy builds                        |
     * | UpdateBuildInformation         | Update build information              |
     * | QueueBuilds                    | Queue builds                          |
     * | ManageBuildQueue               | Manage build queue                    |
     * | StopBuilds                     | Stop builds                           |
     * | ViewBuildDefinition            | View build pipeline                   |
     * | EditBuildDefinition            | Edit build pipeline                   |
     * | DeleteBuildDefinition          | Delete build pipeline                 |
     * | OverrideBuildCheckInValidation | Override check-in validation by build |
     * | AdministerBuildPermissions     | Administer build permissions          |
     * | CreateBuildDefinition          | Create build pipeline                 |
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
     */
    replace?: pulumi.Input<boolean>;
}
