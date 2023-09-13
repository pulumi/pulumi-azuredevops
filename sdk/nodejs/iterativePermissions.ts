// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manages permissions for an Iteration (Sprint)
 *
 * > **Note** Permissions can be assigned to group principals and not to single user principals.
 *
 * ## Permission levels
 *
 * Permission for Iterations within Azure DevOps can be applied on two different levels.
 * Those levels are reflected by specifying (or omitting) values for the arguments `projectId` and `path`.
 *
 * ## Example Usage
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
 * const example_root_permissions = new azuredevops.IterativePermissions("example-root-permissions", {
 *     projectId: example.id,
 *     principal: example_readers.apply(example_readers => example_readers.id),
 *     permissions: {
 *         CREATE_CHILDREN: "Deny",
 *         GENERIC_READ: "NotSet",
 *         DELETE: "Deny",
 *     },
 * });
 * const example_iteration_permissions = new azuredevops.IterativePermissions("example-iteration-permissions", {
 *     projectId: example.id,
 *     principal: example_readers.apply(example_readers => example_readers.id),
 *     path: "Iteration 1",
 *     permissions: {
 *         CREATE_CHILDREN: "Allow",
 *         GENERIC_READ: "NotSet",
 *         DELETE: "Allow",
 *     },
 * });
 * ```
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
export class IterativePermissions extends pulumi.CustomResource {
    /**
     * Get an existing IterativePermissions resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IterativePermissionsState, opts?: pulumi.CustomResourceOptions): IterativePermissions {
        return new IterativePermissions(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuredevops:index/iterativePermissions:IterativePermissions';

    /**
     * Returns true if the given object is an instance of IterativePermissions.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IterativePermissions {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IterativePermissions.__pulumiType;
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
     * Replace (`true`) or merge (`false`) the permissions. Default: `true`
     *
     * | Permission      | Description                    |
     * |-----------------|--------------------------------|
     * | GENERIC_READ    | View permissions for this node |
     * | GENERIC_WRITE   | Edit this node                 |
     * | CREATE_CHILDREN | Create child nodes             |
     * | DELETE          | Delete this node               |
     */
    public readonly replace!: pulumi.Output<boolean | undefined>;

    /**
     * Create a IterativePermissions resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IterativePermissionsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IterativePermissionsArgs | IterativePermissionsState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as IterativePermissionsState | undefined;
            resourceInputs["path"] = state ? state.path : undefined;
            resourceInputs["permissions"] = state ? state.permissions : undefined;
            resourceInputs["principal"] = state ? state.principal : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["replace"] = state ? state.replace : undefined;
        } else {
            const args = argsOrState as IterativePermissionsArgs | undefined;
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
        super(IterativePermissions.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IterativePermissions resources.
 */
export interface IterativePermissionsState {
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
     * Replace (`true`) or merge (`false`) the permissions. Default: `true`
     *
     * | Permission      | Description                    |
     * |-----------------|--------------------------------|
     * | GENERIC_READ    | View permissions for this node |
     * | GENERIC_WRITE   | Edit this node                 |
     * | CREATE_CHILDREN | Create child nodes             |
     * | DELETE          | Delete this node               |
     */
    replace?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a IterativePermissions resource.
 */
export interface IterativePermissionsArgs {
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
     * Replace (`true`) or merge (`false`) the permissions. Default: `true`
     *
     * | Permission      | Description                    |
     * |-----------------|--------------------------------|
     * | GENERIC_READ    | View permissions for this node |
     * | GENERIC_WRITE   | Edit this node                 |
     * | CREATE_CHILDREN | Create child nodes             |
     * | DELETE          | Delete this node               |
     */
    replace?: pulumi.Input<boolean>;
}
