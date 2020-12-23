// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages group membership within Azure DevOps.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const project = new azuredevops.Project("project", {});
 * const user = new azuredevops.User("user", {principalName: "foo@contoso.com"});
 * const group = project.id.apply(id => azuredevops.getGroup({
 *     projectId: id,
 *     name: "Build Administrators",
 * }));
 * const membership = new azuredevops.GroupMembership("membership", {
 *     group: group.descriptor,
 *     members: [user.descriptor],
 * });
 * ```
 * ## Relevant Links
 *
 * - [Azure DevOps Service REST API 5.1 - Memberships](https://docs.microsoft.com/en-us/rest/api/azure/devops/graph/memberships?view=azure-devops-rest-5.0)
 *
 * ## PAT Permissions Required
 *
 * - **Deployment Groups**: Read & Manage
 *
 * ## Import
 *
 * Not supported.
 *
 * @deprecated azuredevops.identities.GroupMembership has been deprecated in favor of azuredevops.GroupMembership
 */
export class GroupMembership extends pulumi.CustomResource {
    /**
     * Get an existing GroupMembership resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GroupMembershipState, opts?: pulumi.CustomResourceOptions): GroupMembership {
        pulumi.log.warn("GroupMembership is deprecated: azuredevops.identities.GroupMembership has been deprecated in favor of azuredevops.GroupMembership")
        return new GroupMembership(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuredevops:Identities/groupMembership:GroupMembership';

    /**
     * Returns true if the given object is an instance of GroupMembership.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GroupMembership {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GroupMembership.__pulumiType;
    }

    /**
     * The descriptor of the group being managed.
     */
    public readonly group!: pulumi.Output<string>;
    /**
     * A list of user or group descriptors that will become members of the group.
     * > NOTE: It's possible to define group members both within the `azuredevops.GroupMembership resource` via the members block and by using the `azuredevops.Group` resource. However it's not possible to use both methods to manage group members, since there'll be conflicts.
     */
    public readonly members!: pulumi.Output<string[]>;
    /**
     * The mode how the resource manages group members.
     * - `mode == add`: the resource will ensure that all specified members will be part of the referenced group
     * - `mode == overwrite`: the resource will replace all existing members with the members specified within the `members` block
     * > NOTE: To clear all members from a group, specify an empty list of descriptors in the `members` attribute and set the `mode` member to `overwrite`.
     */
    public readonly mode!: pulumi.Output<string | undefined>;

    /**
     * Create a GroupMembership resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated azuredevops.identities.GroupMembership has been deprecated in favor of azuredevops.GroupMembership */
    constructor(name: string, args: GroupMembershipArgs, opts?: pulumi.CustomResourceOptions)
    /** @deprecated azuredevops.identities.GroupMembership has been deprecated in favor of azuredevops.GroupMembership */
    constructor(name: string, argsOrState?: GroupMembershipArgs | GroupMembershipState, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("GroupMembership is deprecated: azuredevops.identities.GroupMembership has been deprecated in favor of azuredevops.GroupMembership")
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as GroupMembershipState | undefined;
            inputs["group"] = state ? state.group : undefined;
            inputs["members"] = state ? state.members : undefined;
            inputs["mode"] = state ? state.mode : undefined;
        } else {
            const args = argsOrState as GroupMembershipArgs | undefined;
            if ((!args || args.group === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'group'");
            }
            if ((!args || args.members === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'members'");
            }
            inputs["group"] = args ? args.group : undefined;
            inputs["members"] = args ? args.members : undefined;
            inputs["mode"] = args ? args.mode : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(GroupMembership.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GroupMembership resources.
 */
export interface GroupMembershipState {
    /**
     * The descriptor of the group being managed.
     */
    readonly group?: pulumi.Input<string>;
    /**
     * A list of user or group descriptors that will become members of the group.
     * > NOTE: It's possible to define group members both within the `azuredevops.GroupMembership resource` via the members block and by using the `azuredevops.Group` resource. However it's not possible to use both methods to manage group members, since there'll be conflicts.
     */
    readonly members?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The mode how the resource manages group members.
     * - `mode == add`: the resource will ensure that all specified members will be part of the referenced group
     * - `mode == overwrite`: the resource will replace all existing members with the members specified within the `members` block
     * > NOTE: To clear all members from a group, specify an empty list of descriptors in the `members` attribute and set the `mode` member to `overwrite`.
     */
    readonly mode?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a GroupMembership resource.
 */
export interface GroupMembershipArgs {
    /**
     * The descriptor of the group being managed.
     */
    readonly group: pulumi.Input<string>;
    /**
     * A list of user or group descriptors that will become members of the group.
     * > NOTE: It's possible to define group members both within the `azuredevops.GroupMembership resource` via the members block and by using the `azuredevops.Group` resource. However it's not possible to use both methods to manage group members, since there'll be conflicts.
     */
    readonly members: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The mode how the resource manages group members.
     * - `mode == add`: the resource will ensure that all specified members will be part of the referenced group
     * - `mode == overwrite`: the resource will replace all existing members with the members specified within the `members` block
     * > NOTE: To clear all members from a group, specify an empty list of descriptors in the `members` attribute and set the `mode` member to `overwrite`.
     */
    readonly mode?: pulumi.Input<string>;
}
