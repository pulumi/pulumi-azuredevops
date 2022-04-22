// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a group within Azure DevOps.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const exampleProject = new azuredevops.Project("exampleProject", {});
 * const example-readers = azuredevops.getGroupOutput({
 *     projectId: exampleProject.id,
 *     name: "Readers",
 * });
 * const example-contributors = azuredevops.getGroupOutput({
 *     projectId: exampleProject.id,
 *     name: "Contributors",
 * });
 * const exampleGroup = new azuredevops.Group("exampleGroup", {
 *     scope: exampleProject.id,
 *     displayName: "Example group",
 *     description: "Example description",
 *     members: [
 *         example_readers.apply(example_readers => example_readers.descriptor),
 *         example_contributors.apply(example_contributors => example_contributors.descriptor),
 *     ],
 * });
 * ```
 * ## Relevant Links
 *
 * - [Azure DevOps Service REST API 6.0 - Groups](https://docs.microsoft.com/en-us/rest/api/azure/devops/graph/groups?view=azure-devops-rest-6.0)
 *
 * ## PAT Permissions Required
 *
 * - **Project & Team**: Read, Write, & Manage
 *
 * ## Import
 *
 * Azure DevOps groups can be imported using the group identity descriptor, e.g.
 *
 * ```sh
 *  $ pulumi import azuredevops:Identities/group:Group example aadgp.Uy0xLTktMTU1MTM3NDI0NS0xMjA0NDAwOTY5LTI0MDI5ODY0MTMtMjE3OTQwODYxNi0zLTIxNjc2NjQyNTMtMzI1Nzg0NDI4OS0yMjU4MjcwOTc0LTI2MDYxODY2NDU
 * ```
 *
 * @deprecated azuredevops.identities.Group has been deprecated in favor of azuredevops.Group
 */
export class Group extends pulumi.CustomResource {
    /**
     * Get an existing Group resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GroupState, opts?: pulumi.CustomResourceOptions): Group {
        pulumi.log.warn("Group is deprecated: azuredevops.identities.Group has been deprecated in favor of azuredevops.Group")
        return new Group(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuredevops:Identities/group:Group';

    /**
     * Returns true if the given object is an instance of Group.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Group {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Group.__pulumiType;
    }

    /**
     * The Description of the Project.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The identity (subject) descriptor of the Group.
     */
    public /*out*/ readonly descriptor!: pulumi.Output<string>;
    /**
     * The name of a new Azure DevOps group that is not backed by an external provider. The `originId` and `mail` arguments cannot be used simultaneously with `displayName`.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * This represents the name of the container of origin for a graph member.
     */
    public /*out*/ readonly domain!: pulumi.Output<string>;
    /**
     * The mail address as a reference to an existing group from an external AD or AAD backed provider. The `scope`, `originId` and `displayName` arguments cannot be used simultaneously with `mail`.
     */
    public readonly mail!: pulumi.Output<string>;
    /**
     * > NOTE: It's possible to define group members both within the `azuredevops.Group` resource via the members block and by using the `azuredevops.GroupMembership` resource. However it's not possible to use both methods to manage group members, since there'll be conflicts.
     */
    public readonly members!: pulumi.Output<string[]>;
    /**
     * The type of source provider for the origin identifier (ex:AD, AAD, MSA)
     */
    public /*out*/ readonly origin!: pulumi.Output<string>;
    /**
     * The OriginID as a reference to a group from an external AD or AAD backed provider. The `scope`, `mail` and `displayName` arguments cannot be used simultaneously with `originId`.
     */
    public readonly originId!: pulumi.Output<string>;
    /**
     * This is the PrincipalName of this graph member from the source provider.
     */
    public /*out*/ readonly principalName!: pulumi.Output<string>;
    /**
     * The scope of the group. A descriptor referencing the scope (collection, project) in which the group should be created. If omitted, will be created in the scope of the enclosing account or organization.x
     */
    public readonly scope!: pulumi.Output<string | undefined>;
    /**
     * This field identifies the type of the graph subject (ex: Group, Scope, User).
     */
    public /*out*/ readonly subjectKind!: pulumi.Output<string>;
    /**
     * This url is the full route to the source resource of this graph subject.
     */
    public /*out*/ readonly url!: pulumi.Output<string>;

    /**
     * Create a Group resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated azuredevops.identities.Group has been deprecated in favor of azuredevops.Group */
    constructor(name: string, args?: GroupArgs, opts?: pulumi.CustomResourceOptions)
    /** @deprecated azuredevops.identities.Group has been deprecated in favor of azuredevops.Group */
    constructor(name: string, argsOrState?: GroupArgs | GroupState, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("Group is deprecated: azuredevops.identities.Group has been deprecated in favor of azuredevops.Group")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GroupState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["descriptor"] = state ? state.descriptor : undefined;
            resourceInputs["displayName"] = state ? state.displayName : undefined;
            resourceInputs["domain"] = state ? state.domain : undefined;
            resourceInputs["mail"] = state ? state.mail : undefined;
            resourceInputs["members"] = state ? state.members : undefined;
            resourceInputs["origin"] = state ? state.origin : undefined;
            resourceInputs["originId"] = state ? state.originId : undefined;
            resourceInputs["principalName"] = state ? state.principalName : undefined;
            resourceInputs["scope"] = state ? state.scope : undefined;
            resourceInputs["subjectKind"] = state ? state.subjectKind : undefined;
            resourceInputs["url"] = state ? state.url : undefined;
        } else {
            const args = argsOrState as GroupArgs | undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["mail"] = args ? args.mail : undefined;
            resourceInputs["members"] = args ? args.members : undefined;
            resourceInputs["originId"] = args ? args.originId : undefined;
            resourceInputs["scope"] = args ? args.scope : undefined;
            resourceInputs["descriptor"] = undefined /*out*/;
            resourceInputs["domain"] = undefined /*out*/;
            resourceInputs["origin"] = undefined /*out*/;
            resourceInputs["principalName"] = undefined /*out*/;
            resourceInputs["subjectKind"] = undefined /*out*/;
            resourceInputs["url"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Group.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Group resources.
 */
export interface GroupState {
    /**
     * The Description of the Project.
     */
    description?: pulumi.Input<string>;
    /**
     * The identity (subject) descriptor of the Group.
     */
    descriptor?: pulumi.Input<string>;
    /**
     * The name of a new Azure DevOps group that is not backed by an external provider. The `originId` and `mail` arguments cannot be used simultaneously with `displayName`.
     */
    displayName?: pulumi.Input<string>;
    /**
     * This represents the name of the container of origin for a graph member.
     */
    domain?: pulumi.Input<string>;
    /**
     * The mail address as a reference to an existing group from an external AD or AAD backed provider. The `scope`, `originId` and `displayName` arguments cannot be used simultaneously with `mail`.
     */
    mail?: pulumi.Input<string>;
    /**
     * > NOTE: It's possible to define group members both within the `azuredevops.Group` resource via the members block and by using the `azuredevops.GroupMembership` resource. However it's not possible to use both methods to manage group members, since there'll be conflicts.
     */
    members?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The type of source provider for the origin identifier (ex:AD, AAD, MSA)
     */
    origin?: pulumi.Input<string>;
    /**
     * The OriginID as a reference to a group from an external AD or AAD backed provider. The `scope`, `mail` and `displayName` arguments cannot be used simultaneously with `originId`.
     */
    originId?: pulumi.Input<string>;
    /**
     * This is the PrincipalName of this graph member from the source provider.
     */
    principalName?: pulumi.Input<string>;
    /**
     * The scope of the group. A descriptor referencing the scope (collection, project) in which the group should be created. If omitted, will be created in the scope of the enclosing account or organization.x
     */
    scope?: pulumi.Input<string>;
    /**
     * This field identifies the type of the graph subject (ex: Group, Scope, User).
     */
    subjectKind?: pulumi.Input<string>;
    /**
     * This url is the full route to the source resource of this graph subject.
     */
    url?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Group resource.
 */
export interface GroupArgs {
    /**
     * The Description of the Project.
     */
    description?: pulumi.Input<string>;
    /**
     * The name of a new Azure DevOps group that is not backed by an external provider. The `originId` and `mail` arguments cannot be used simultaneously with `displayName`.
     */
    displayName?: pulumi.Input<string>;
    /**
     * The mail address as a reference to an existing group from an external AD or AAD backed provider. The `scope`, `originId` and `displayName` arguments cannot be used simultaneously with `mail`.
     */
    mail?: pulumi.Input<string>;
    /**
     * > NOTE: It's possible to define group members both within the `azuredevops.Group` resource via the members block and by using the `azuredevops.GroupMembership` resource. However it's not possible to use both methods to manage group members, since there'll be conflicts.
     */
    members?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The OriginID as a reference to a group from an external AD or AAD backed provider. The `scope`, `mail` and `displayName` arguments cannot be used simultaneously with `originId`.
     */
    originId?: pulumi.Input<string>;
    /**
     * The scope of the group. A descriptor referencing the scope (collection, project) in which the group should be created. If omitted, will be created in the scope of the enclosing account or organization.x
     */
    scope?: pulumi.Input<string>;
}
