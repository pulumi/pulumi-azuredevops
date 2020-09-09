// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manages a user entitlement within Azure DevOps.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const user = new azuredevops.User("user", {
 *     principalName: "foo@contoso.com",
 * });
 * ```
 * ## Relevant Links
 *
 * * [Azure DevOps Service REST API 5.1 - User Entitlements - Add](https://docs.microsoft.com/en-us/rest/api/azure/devops/memberentitlementmanagement/user%20entitlements/add?view=azure-devops-rest-5.1)
 *
 * ## PAT Permissions Required
 *
 * - **Member Entitlement Management**: Read & Write
 */
export class User extends pulumi.CustomResource {
    /**
     * Get an existing User resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UserState, opts?: pulumi.CustomResourceOptions): User {
        return new User(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuredevops:index/user:User';

    /**
     * Returns true if the given object is an instance of User.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is User {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === User.__pulumiType;
    }

    /**
     * Type of Account License. Valid values: `advanced`, `earlyAdopter`, `express`, `none`, `professional`, or `stakeholder`. Defaults to `express`. In addition the value `basic` is allowed which is an alias for `express` and reflects the name of the `express` license used in the Azure DevOps web interface.
     */
    public readonly accountLicenseType!: pulumi.Output<string | undefined>;
    /**
     * The descriptor is the primary way to reference the graph subject while the system is running. This field will uniquely identify the user graph subject.
     */
    public /*out*/ readonly descriptor!: pulumi.Output<string>;
    /**
     * The source of the licensing (e.g. Account. MSDN etc.) Valid values: `account` (Default), `auto`, `msdn`, `none`, `profile`, `trail`
     */
    public readonly licensingSource!: pulumi.Output<string | undefined>;
    /**
     * The type of source provider for the origin identifier.
     */
    public readonly origin!: pulumi.Output<string>;
    /**
     * The unique identifier from the system of origin. Typically a sid, object id or Guid. e.g. Used for member of other tenant on Azure Active Directory.
     */
    public readonly originId!: pulumi.Output<string>;
    /**
     * The principal name is the PrincipalName of a graph member from the source provider. Usually, e-mail address.
     */
    public readonly principalName!: pulumi.Output<string>;

    /**
     * Create a User resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: UserArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UserArgs | UserState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as UserState | undefined;
            inputs["accountLicenseType"] = state ? state.accountLicenseType : undefined;
            inputs["descriptor"] = state ? state.descriptor : undefined;
            inputs["licensingSource"] = state ? state.licensingSource : undefined;
            inputs["origin"] = state ? state.origin : undefined;
            inputs["originId"] = state ? state.originId : undefined;
            inputs["principalName"] = state ? state.principalName : undefined;
        } else {
            const args = argsOrState as UserArgs | undefined;
            inputs["accountLicenseType"] = args ? args.accountLicenseType : undefined;
            inputs["licensingSource"] = args ? args.licensingSource : undefined;
            inputs["origin"] = args ? args.origin : undefined;
            inputs["originId"] = args ? args.originId : undefined;
            inputs["principalName"] = args ? args.principalName : undefined;
            inputs["descriptor"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azuredevops:Entitlement/user:User" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(User.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering User resources.
 */
export interface UserState {
    /**
     * Type of Account License. Valid values: `advanced`, `earlyAdopter`, `express`, `none`, `professional`, or `stakeholder`. Defaults to `express`. In addition the value `basic` is allowed which is an alias for `express` and reflects the name of the `express` license used in the Azure DevOps web interface.
     */
    readonly accountLicenseType?: pulumi.Input<string>;
    /**
     * The descriptor is the primary way to reference the graph subject while the system is running. This field will uniquely identify the user graph subject.
     */
    readonly descriptor?: pulumi.Input<string>;
    /**
     * The source of the licensing (e.g. Account. MSDN etc.) Valid values: `account` (Default), `auto`, `msdn`, `none`, `profile`, `trail`
     */
    readonly licensingSource?: pulumi.Input<string>;
    /**
     * The type of source provider for the origin identifier.
     */
    readonly origin?: pulumi.Input<string>;
    /**
     * The unique identifier from the system of origin. Typically a sid, object id or Guid. e.g. Used for member of other tenant on Azure Active Directory.
     */
    readonly originId?: pulumi.Input<string>;
    /**
     * The principal name is the PrincipalName of a graph member from the source provider. Usually, e-mail address.
     */
    readonly principalName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a User resource.
 */
export interface UserArgs {
    /**
     * Type of Account License. Valid values: `advanced`, `earlyAdopter`, `express`, `none`, `professional`, or `stakeholder`. Defaults to `express`. In addition the value `basic` is allowed which is an alias for `express` and reflects the name of the `express` license used in the Azure DevOps web interface.
     */
    readonly accountLicenseType?: pulumi.Input<string>;
    /**
     * The source of the licensing (e.g. Account. MSDN etc.) Valid values: `account` (Default), `auto`, `msdn`, `none`, `profile`, `trail`
     */
    readonly licensingSource?: pulumi.Input<string>;
    /**
     * The type of source provider for the origin identifier.
     */
    readonly origin?: pulumi.Input<string>;
    /**
     * The unique identifier from the system of origin. Typically a sid, object id or Guid. e.g. Used for member of other tenant on Azure Active Directory.
     */
    readonly originId?: pulumi.Input<string>;
    /**
     * The principal name is the PrincipalName of a graph member from the source provider. Usually, e-mail address.
     */
    readonly principalName?: pulumi.Input<string>;
}
