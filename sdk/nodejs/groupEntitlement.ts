// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * layout: "azuredevops"
 * page_title: "AzureDevops: azuredevops.GroupEntitlement"
 * description: |-
 * Manages a group entitlement within Azure DevOps organization.
 * <!-- yaml: line 5: could not find expected ':' -->
 *
 * # azuredevops.GroupEntitlement
 *
 * Manages a group entitlement within Azure DevOps.
 *
 * ## Example Usage
 *
 * ### With an Azure DevOps local group managed by this resource
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const example = new azuredevops.GroupEntitlement("example", {displayName: "Group Name"});
 * ```
 *
 * ### With group origin ID
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const example = new azuredevops.GroupEntitlement("example", {
 *     origin: "aad",
 *     originId: "00000000-0000-0000-0000-000000000000",
 * });
 * ```
 *
 * ## Relevant Links
 *
 * - [Azure DevOps Service REST API 7.0 - Group Entitlements](https://learn.microsoft.com/en-us/rest/api/azure/devops/memberentitlementmanagement/group-entitlements?view=azure-devops-rest-7.1)
 * - [Programmatic mapping of access levels](https://docs.microsoft.com/en-us/azure/devops/organizations/security/access-levels?view=azure-devops#programmatic-mapping-of-access-levels)
 *
 * ## PAT Permissions Required
 *
 * - **Member Entitlement Management**: Read & Write
 *
 * ## Import
 *
 * The resource allows the import via the ID of a group entitlement, which is a
 *
 * UUID.
 *
 * ```sh
 * $ pulumi import azuredevops:index/groupEntitlement:GroupEntitlement example 00000000-0000-0000-0000-000000000000
 * ```
 */
export class GroupEntitlement extends pulumi.CustomResource {
    /**
     * Get an existing GroupEntitlement resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GroupEntitlementState, opts?: pulumi.CustomResourceOptions): GroupEntitlement {
        return new GroupEntitlement(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuredevops:index/groupEntitlement:GroupEntitlement';

    /**
     * Returns true if the given object is an instance of GroupEntitlement.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GroupEntitlement {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GroupEntitlement.__pulumiType;
    }

    /**
     * Type of Account License. Valid values: `advanced`, `earlyAdopter`, `express`, `none`, `professional`, or `stakeholder`. Defaults to `express`. In addition, the value `basic` is allowed which is an alias for `express` and reflects the name of the `express` license used in the Azure DevOps web interface.
     */
    public readonly accountLicenseType!: pulumi.Output<string | undefined>;
    /**
     * The descriptor is the primary way to reference the graph subject while the system is running. This field will uniquely identify the group graph subject.
     */
    public /*out*/ readonly descriptor!: pulumi.Output<string>;
    /**
     * The display name is the name used in Azure DevOps UI. Cannot be set together with `originId` and `origin`.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * The source of the licensing (e.g. Account. MSDN etc.) Valid values: `account` (Default), `auto`, `msdn`, `none`, `profile`, `trial`
     *
     * > **NOTE:** A existing group in Azure AD can only be referenced by the combination of `originId` and `origin`.
     */
    public readonly licensingSource!: pulumi.Output<string | undefined>;
    /**
     * The type of source provider for the origin identifier.
     */
    public readonly origin!: pulumi.Output<string>;
    /**
     * The unique identifier from the system of origin. Typically, a sid, object id or Guid. e.g. Used for member of other tenant on Azure Active Directory.
     */
    public readonly originId!: pulumi.Output<string>;
    /**
     * The principal name of a graph member on Azure DevOps
     */
    public /*out*/ readonly principalName!: pulumi.Output<string>;

    /**
     * Create a GroupEntitlement resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: GroupEntitlementArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GroupEntitlementArgs | GroupEntitlementState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GroupEntitlementState | undefined;
            resourceInputs["accountLicenseType"] = state ? state.accountLicenseType : undefined;
            resourceInputs["descriptor"] = state ? state.descriptor : undefined;
            resourceInputs["displayName"] = state ? state.displayName : undefined;
            resourceInputs["licensingSource"] = state ? state.licensingSource : undefined;
            resourceInputs["origin"] = state ? state.origin : undefined;
            resourceInputs["originId"] = state ? state.originId : undefined;
            resourceInputs["principalName"] = state ? state.principalName : undefined;
        } else {
            const args = argsOrState as GroupEntitlementArgs | undefined;
            resourceInputs["accountLicenseType"] = args ? args.accountLicenseType : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["licensingSource"] = args ? args.licensingSource : undefined;
            resourceInputs["origin"] = args ? args.origin : undefined;
            resourceInputs["originId"] = args ? args.originId : undefined;
            resourceInputs["descriptor"] = undefined /*out*/;
            resourceInputs["principalName"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(GroupEntitlement.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GroupEntitlement resources.
 */
export interface GroupEntitlementState {
    /**
     * Type of Account License. Valid values: `advanced`, `earlyAdopter`, `express`, `none`, `professional`, or `stakeholder`. Defaults to `express`. In addition, the value `basic` is allowed which is an alias for `express` and reflects the name of the `express` license used in the Azure DevOps web interface.
     */
    accountLicenseType?: pulumi.Input<string>;
    /**
     * The descriptor is the primary way to reference the graph subject while the system is running. This field will uniquely identify the group graph subject.
     */
    descriptor?: pulumi.Input<string>;
    /**
     * The display name is the name used in Azure DevOps UI. Cannot be set together with `originId` and `origin`.
     */
    displayName?: pulumi.Input<string>;
    /**
     * The source of the licensing (e.g. Account. MSDN etc.) Valid values: `account` (Default), `auto`, `msdn`, `none`, `profile`, `trial`
     *
     * > **NOTE:** A existing group in Azure AD can only be referenced by the combination of `originId` and `origin`.
     */
    licensingSource?: pulumi.Input<string>;
    /**
     * The type of source provider for the origin identifier.
     */
    origin?: pulumi.Input<string>;
    /**
     * The unique identifier from the system of origin. Typically, a sid, object id or Guid. e.g. Used for member of other tenant on Azure Active Directory.
     */
    originId?: pulumi.Input<string>;
    /**
     * The principal name of a graph member on Azure DevOps
     */
    principalName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a GroupEntitlement resource.
 */
export interface GroupEntitlementArgs {
    /**
     * Type of Account License. Valid values: `advanced`, `earlyAdopter`, `express`, `none`, `professional`, or `stakeholder`. Defaults to `express`. In addition, the value `basic` is allowed which is an alias for `express` and reflects the name of the `express` license used in the Azure DevOps web interface.
     */
    accountLicenseType?: pulumi.Input<string>;
    /**
     * The display name is the name used in Azure DevOps UI. Cannot be set together with `originId` and `origin`.
     */
    displayName?: pulumi.Input<string>;
    /**
     * The source of the licensing (e.g. Account. MSDN etc.) Valid values: `account` (Default), `auto`, `msdn`, `none`, `profile`, `trial`
     *
     * > **NOTE:** A existing group in Azure AD can only be referenced by the combination of `originId` and `origin`.
     */
    licensingSource?: pulumi.Input<string>;
    /**
     * The type of source provider for the origin identifier.
     */
    origin?: pulumi.Input<string>;
    /**
     * The unique identifier from the system of origin. Typically, a sid, object id or Guid. e.g. Used for member of other tenant on Azure Active Directory.
     */
    originId?: pulumi.Input<string>;
}
