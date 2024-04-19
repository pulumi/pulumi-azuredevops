// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manages permissions for tagging
 *
 * ## Permission levels
 *
 * Permissions for tagging within Azure DevOps can be applied only on Organizational and Project level.
 * The project level is reflected by specifying the argument `projectId`, otherwise the permissions are set on the organizational level.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
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
 * const example-readers = azuredevops.getGroupOutput({
 *     projectId: example.id,
 *     name: "Readers",
 * });
 * const example_permissions = new azuredevops.TaggingPermissions("example-permissions", {
 *     projectId: example.id,
 *     principal: example_readers.apply(example_readers => example_readers.id),
 *     permissions: {
 *         Enumerate: "allow",
 *         Create: "allow",
 *         Update: "allow",
 *         Delete: "allow",
 *     },
 * });
 * ```
 * <!--End PulumiCodeChooser -->
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
export class TaggingPermissions extends pulumi.CustomResource {
    /**
     * Get an existing TaggingPermissions resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TaggingPermissionsState, opts?: pulumi.CustomResourceOptions): TaggingPermissions {
        return new TaggingPermissions(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuredevops:index/taggingPermissions:TaggingPermissions';

    /**
     * Returns true if the given object is an instance of TaggingPermissions.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TaggingPermissions {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TaggingPermissions.__pulumiType;
    }

    /**
     * the permissions to assign. The following permissions are available.
     */
    public readonly permissions!: pulumi.Output<{[key: string]: string}>;
    /**
     * The **group or user** principal to assign the permissions.
     */
    public readonly principal!: pulumi.Output<string>;
    /**
     * The ID of the project to assign the permissions. If omitted, organization wide permissions for tagging are managed.
     */
    public readonly projectId!: pulumi.Output<string | undefined>;
    /**
     * Replace (`true`) or merge (`false`) the permissions. Default: `true`
     *
     * | Name               | Permission Description     |
     * | ------------------ | -------------------------- |
     * | Enumerate          | Enumerate tag definitions  |
     * | Create             | Create tag definition      |
     * | Update             | Update tag definition      |
     * | Delete             | Delete tag definition      |
     */
    public readonly replace!: pulumi.Output<boolean | undefined>;

    /**
     * Create a TaggingPermissions resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TaggingPermissionsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TaggingPermissionsArgs | TaggingPermissionsState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TaggingPermissionsState | undefined;
            resourceInputs["permissions"] = state ? state.permissions : undefined;
            resourceInputs["principal"] = state ? state.principal : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["replace"] = state ? state.replace : undefined;
        } else {
            const args = argsOrState as TaggingPermissionsArgs | undefined;
            if ((!args || args.permissions === undefined) && !opts.urn) {
                throw new Error("Missing required property 'permissions'");
            }
            if ((!args || args.principal === undefined) && !opts.urn) {
                throw new Error("Missing required property 'principal'");
            }
            resourceInputs["permissions"] = args ? args.permissions : undefined;
            resourceInputs["principal"] = args ? args.principal : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["replace"] = args ? args.replace : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(TaggingPermissions.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TaggingPermissions resources.
 */
export interface TaggingPermissionsState {
    /**
     * the permissions to assign. The following permissions are available.
     */
    permissions?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The **group or user** principal to assign the permissions.
     */
    principal?: pulumi.Input<string>;
    /**
     * The ID of the project to assign the permissions. If omitted, organization wide permissions for tagging are managed.
     */
    projectId?: pulumi.Input<string>;
    /**
     * Replace (`true`) or merge (`false`) the permissions. Default: `true`
     *
     * | Name               | Permission Description     |
     * | ------------------ | -------------------------- |
     * | Enumerate          | Enumerate tag definitions  |
     * | Create             | Create tag definition      |
     * | Update             | Update tag definition      |
     * | Delete             | Delete tag definition      |
     */
    replace?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a TaggingPermissions resource.
 */
export interface TaggingPermissionsArgs {
    /**
     * the permissions to assign. The following permissions are available.
     */
    permissions: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The **group or user** principal to assign the permissions.
     */
    principal: pulumi.Input<string>;
    /**
     * The ID of the project to assign the permissions. If omitted, organization wide permissions for tagging are managed.
     */
    projectId?: pulumi.Input<string>;
    /**
     * Replace (`true`) or merge (`false`) the permissions. Default: `true`
     *
     * | Name               | Permission Description     |
     * | ------------------ | -------------------------- |
     * | Enumerate          | Enumerate tag definitions  |
     * | Create             | Create tag definition      |
     * | Update             | Update tag definition      |
     * | Delete             | Delete tag definition      |
     */
    replace?: pulumi.Input<boolean>;
}
