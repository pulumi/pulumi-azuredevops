// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manages permissions for service hooks
 *
 * ## Permission levels
 *
 * Permissions for service hooks within Azure DevOps can be applied on the Organizational level or, if the optional attribute `projectId` is specified, on Project level.
 * Those levels are reflected by specifying (or omitting) values for the argument `projectId`.
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
 * const example-readers = azuredevops.getGroupOutput({
 *     projectId: example.id,
 *     name: "Readers",
 * });
 * const example_permissions = new azuredevops.ServicehookPermissions("example-permissions", {
 *     projectId: example.id,
 *     principal: example_readers.apply(example_readers => example_readers.id),
 *     permissions: {
 *         ViewSubscriptions: "allow",
 *         EditSubscriptions: "allow",
 *         DeleteSubscriptions: "allow",
 *         PublishEvents: "allow",
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
export class ServicehookPermissions extends pulumi.CustomResource {
    /**
     * Get an existing ServicehookPermissions resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServicehookPermissionsState, opts?: pulumi.CustomResourceOptions): ServicehookPermissions {
        return new ServicehookPermissions(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuredevops:index/servicehookPermissions:ServicehookPermissions';

    /**
     * Returns true if the given object is an instance of ServicehookPermissions.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServicehookPermissions {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServicehookPermissions.__pulumiType;
    }

    /**
     * the permissions to assign. The following permissions are available.
     */
    public readonly permissions!: pulumi.Output<{[key: string]: string}>;
    /**
     * The **group** principal to assign the permissions.
     */
    public readonly principal!: pulumi.Output<string>;
    /**
     * The ID of the project.
     */
    public readonly projectId!: pulumi.Output<string | undefined>;
    /**
     * Replace (`true`) or merge (`false`) the permissions. Default: `true`
     * <table>
     * <thead>
     * <tr>
     * <th>Name</th>
     * <th>Permission Description</th>
     * </tr>
     * </thead>
     * <tbody>
     * <tr>
     * <td>ViewSubscriptions</td>
     * <td>View Subscriptions</td>
     * </tr>
     * <tr>
     * <td>EditSubscriptions</td>
     * <td>Edit Subscription</td>
     * </tr>
     * <tr>
     * <td>DeleteSubscriptions</td>
     * <td>Delete Subscriptions</td>
     * </tr>
     * <tr>
     * <td>PublishEvents</td>
     * <td>Publish Events</td>
     * </tr>
     * </tbody>
     * </table>
     */
    public readonly replace!: pulumi.Output<boolean | undefined>;

    /**
     * Create a ServicehookPermissions resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServicehookPermissionsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServicehookPermissionsArgs | ServicehookPermissionsState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ServicehookPermissionsState | undefined;
            resourceInputs["permissions"] = state ? state.permissions : undefined;
            resourceInputs["principal"] = state ? state.principal : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["replace"] = state ? state.replace : undefined;
        } else {
            const args = argsOrState as ServicehookPermissionsArgs | undefined;
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
        super(ServicehookPermissions.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ServicehookPermissions resources.
 */
export interface ServicehookPermissionsState {
    /**
     * the permissions to assign. The following permissions are available.
     */
    permissions?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The **group** principal to assign the permissions.
     */
    principal?: pulumi.Input<string>;
    /**
     * The ID of the project.
     */
    projectId?: pulumi.Input<string>;
    /**
     * Replace (`true`) or merge (`false`) the permissions. Default: `true`
     * <table>
     * <thead>
     * <tr>
     * <th>Name</th>
     * <th>Permission Description</th>
     * </tr>
     * </thead>
     * <tbody>
     * <tr>
     * <td>ViewSubscriptions</td>
     * <td>View Subscriptions</td>
     * </tr>
     * <tr>
     * <td>EditSubscriptions</td>
     * <td>Edit Subscription</td>
     * </tr>
     * <tr>
     * <td>DeleteSubscriptions</td>
     * <td>Delete Subscriptions</td>
     * </tr>
     * <tr>
     * <td>PublishEvents</td>
     * <td>Publish Events</td>
     * </tr>
     * </tbody>
     * </table>
     */
    replace?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a ServicehookPermissions resource.
 */
export interface ServicehookPermissionsArgs {
    /**
     * the permissions to assign. The following permissions are available.
     */
    permissions: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The **group** principal to assign the permissions.
     */
    principal: pulumi.Input<string>;
    /**
     * The ID of the project.
     */
    projectId?: pulumi.Input<string>;
    /**
     * Replace (`true`) or merge (`false`) the permissions. Default: `true`
     * <table>
     * <thead>
     * <tr>
     * <th>Name</th>
     * <th>Permission Description</th>
     * </tr>
     * </thead>
     * <tbody>
     * <tr>
     * <td>ViewSubscriptions</td>
     * <td>View Subscriptions</td>
     * </tr>
     * <tr>
     * <td>EditSubscriptions</td>
     * <td>Edit Subscription</td>
     * </tr>
     * <tr>
     * <td>DeleteSubscriptions</td>
     * <td>Delete Subscriptions</td>
     * </tr>
     * <tr>
     * <td>PublishEvents</td>
     * <td>Publish Events</td>
     * </tr>
     * </tbody>
     * </table>
     */
    replace?: pulumi.Input<boolean>;
}
