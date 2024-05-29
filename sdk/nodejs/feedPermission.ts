// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manages creation of the Feed Permission within Azure DevOps organization.
 *
 * ## Example Usage
 *
 * ### Create Feed Permission
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const example = new azuredevops.Project("example", {name: "Example Project"});
 * const exampleGroup = new azuredevops.Group("example", {
 *     scope: example.id,
 *     displayName: "Example group",
 *     description: "Example description",
 * });
 * const exampleFeed = new azuredevops.Feed("example", {name: "releases"});
 * const permission = new azuredevops.FeedPermission("permission", {
 *     feedId: exampleFeed.id,
 *     role: "reader",
 *     identityDescriptor: exampleGroup.descriptor,
 * });
 * ```
 *
 * ## Relevant Links
 *
 * - [Azure DevOps Service REST API 7.0 - Feed Management](https://learn.microsoft.com/en-us/rest/api/azure/devops/artifacts/feed-management?view=azure-devops-rest-7.0)
 */
export class FeedPermission extends pulumi.CustomResource {
    /**
     * Get an existing FeedPermission resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FeedPermissionState, opts?: pulumi.CustomResourceOptions): FeedPermission {
        return new FeedPermission(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuredevops:index/feedPermission:FeedPermission';

    /**
     * Returns true if the given object is an instance of FeedPermission.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FeedPermission {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FeedPermission.__pulumiType;
    }

    /**
     * The display name of the assignment
     */
    public readonly displayName!: pulumi.Output<string | undefined>;
    /**
     * The ID of the Feed.
     */
    public readonly feedId!: pulumi.Output<string>;
    /**
     * The Descriptor of identity you want to assign a role.
     */
    public readonly identityDescriptor!: pulumi.Output<string>;
    /**
     * The ID of the identity.
     */
    public /*out*/ readonly identityId!: pulumi.Output<string>;
    /**
     * The ID of the Project Feed is created in. If not specified, feed will be created at the organization level.
     */
    public readonly projectId!: pulumi.Output<string | undefined>;
    /**
     * The role to be assigned, possible values : `reader`, `contributor`, `collaborator`, `administrator`
     */
    public readonly role!: pulumi.Output<string>;

    /**
     * Create a FeedPermission resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FeedPermissionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FeedPermissionArgs | FeedPermissionState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FeedPermissionState | undefined;
            resourceInputs["displayName"] = state ? state.displayName : undefined;
            resourceInputs["feedId"] = state ? state.feedId : undefined;
            resourceInputs["identityDescriptor"] = state ? state.identityDescriptor : undefined;
            resourceInputs["identityId"] = state ? state.identityId : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["role"] = state ? state.role : undefined;
        } else {
            const args = argsOrState as FeedPermissionArgs | undefined;
            if ((!args || args.feedId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'feedId'");
            }
            if ((!args || args.identityDescriptor === undefined) && !opts.urn) {
                throw new Error("Missing required property 'identityDescriptor'");
            }
            if ((!args || args.role === undefined) && !opts.urn) {
                throw new Error("Missing required property 'role'");
            }
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["feedId"] = args ? args.feedId : undefined;
            resourceInputs["identityDescriptor"] = args ? args.identityDescriptor : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["role"] = args ? args.role : undefined;
            resourceInputs["identityId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FeedPermission.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FeedPermission resources.
 */
export interface FeedPermissionState {
    /**
     * The display name of the assignment
     */
    displayName?: pulumi.Input<string>;
    /**
     * The ID of the Feed.
     */
    feedId?: pulumi.Input<string>;
    /**
     * The Descriptor of identity you want to assign a role.
     */
    identityDescriptor?: pulumi.Input<string>;
    /**
     * The ID of the identity.
     */
    identityId?: pulumi.Input<string>;
    /**
     * The ID of the Project Feed is created in. If not specified, feed will be created at the organization level.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The role to be assigned, possible values : `reader`, `contributor`, `collaborator`, `administrator`
     */
    role?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FeedPermission resource.
 */
export interface FeedPermissionArgs {
    /**
     * The display name of the assignment
     */
    displayName?: pulumi.Input<string>;
    /**
     * The ID of the Feed.
     */
    feedId: pulumi.Input<string>;
    /**
     * The Descriptor of identity you want to assign a role.
     */
    identityDescriptor: pulumi.Input<string>;
    /**
     * The ID of the Project Feed is created in. If not specified, feed will be created at the organization level.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The role to be assigned, possible values : `reader`, `contributor`, `collaborator`, `administrator`
     */
    role: pulumi.Input<string>;
}