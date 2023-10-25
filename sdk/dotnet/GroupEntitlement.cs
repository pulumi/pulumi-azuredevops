// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureDevOps
{
    /// <summary>
    /// Manages a group entitlement within Azure DevOps.
    /// 
    /// ## Example Usage
    /// ## Relevant Links
    /// 
    /// - [Azure DevOps Service REST API 7.0 - Group Entitlements](https://learn.microsoft.com/en-us/rest/api/azure/devops/memberentitlementmanagement/group-entitlements?view=azure-devops-rest-7.1)
    /// - [Programmatic mapping of access levels](https://docs.microsoft.com/en-us/azure/devops/organizations/security/access-levels?view=azure-devops#programmatic-mapping-of-access-levels)
    /// 
    /// ## PAT Permissions Required
    /// 
    /// - **Member Entitlement Management**: Read &amp; Write
    /// 
    /// ## Import
    /// 
    /// The resource allows the import via the ID of a group entitlement, which is a UUID.
    /// 
    /// ```sh
    ///  $ pulumi import azuredevops:index/groupEntitlement:GroupEntitlement example 00000000-0000-0000-0000-000000000000
    /// ```
    /// </summary>
    [AzureDevOpsResourceType("azuredevops:index/groupEntitlement:GroupEntitlement")]
    public partial class GroupEntitlement : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Type of Account License. Valid values: `advanced`, `earlyAdopter`, `express`, `none`, `professional`, or `stakeholder`. Defaults to `express`. In addition, the value `basic` is allowed which is an alias for `express` and reflects the name of the `express` license used in the Azure DevOps web interface.
        /// </summary>
        [Output("accountLicenseType")]
        public Output<string?> AccountLicenseType { get; private set; } = null!;

        /// <summary>
        /// The descriptor is the primary way to reference the graph subject while the system is running. This field will uniquely identify the group graph subject.
        /// </summary>
        [Output("descriptor")]
        public Output<string> Descriptor { get; private set; } = null!;

        /// <summary>
        /// The display name is the name used in Azure DevOps UI. Cannot be set together with `origin_id` and `origin`.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// The source of the licensing (e.g. Account. MSDN etc.) Valid values: `account` (Default), `auto`, `msdn`, `none`, `profile`, `trial`
        /// 
        /// &gt; **NOTE:** A existing group in Azure AD can only be referenced by the combination of `origin_id` and `origin`.
        /// </summary>
        [Output("licensingSource")]
        public Output<string?> LicensingSource { get; private set; } = null!;

        /// <summary>
        /// The type of source provider for the origin identifier.
        /// </summary>
        [Output("origin")]
        public Output<string> Origin { get; private set; } = null!;

        /// <summary>
        /// The unique identifier from the system of origin. Typically, a sid, object id or Guid. e.g. Used for member of other tenant on Azure Active Directory.
        /// </summary>
        [Output("originId")]
        public Output<string> OriginId { get; private set; } = null!;

        /// <summary>
        /// The principal name of a graph member on Azure DevOps
        /// </summary>
        [Output("principalName")]
        public Output<string> PrincipalName { get; private set; } = null!;


        /// <summary>
        /// Create a GroupEntitlement resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GroupEntitlement(string name, GroupEntitlementArgs? args = null, CustomResourceOptions? options = null)
            : base("azuredevops:index/groupEntitlement:GroupEntitlement", name, args ?? new GroupEntitlementArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GroupEntitlement(string name, Input<string> id, GroupEntitlementState? state = null, CustomResourceOptions? options = null)
            : base("azuredevops:index/groupEntitlement:GroupEntitlement", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing GroupEntitlement resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GroupEntitlement Get(string name, Input<string> id, GroupEntitlementState? state = null, CustomResourceOptions? options = null)
        {
            return new GroupEntitlement(name, id, state, options);
        }
    }

    public sealed class GroupEntitlementArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Type of Account License. Valid values: `advanced`, `earlyAdopter`, `express`, `none`, `professional`, or `stakeholder`. Defaults to `express`. In addition, the value `basic` is allowed which is an alias for `express` and reflects the name of the `express` license used in the Azure DevOps web interface.
        /// </summary>
        [Input("accountLicenseType")]
        public Input<string>? AccountLicenseType { get; set; }

        /// <summary>
        /// The display name is the name used in Azure DevOps UI. Cannot be set together with `origin_id` and `origin`.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// The source of the licensing (e.g. Account. MSDN etc.) Valid values: `account` (Default), `auto`, `msdn`, `none`, `profile`, `trial`
        /// 
        /// &gt; **NOTE:** A existing group in Azure AD can only be referenced by the combination of `origin_id` and `origin`.
        /// </summary>
        [Input("licensingSource")]
        public Input<string>? LicensingSource { get; set; }

        /// <summary>
        /// The type of source provider for the origin identifier.
        /// </summary>
        [Input("origin")]
        public Input<string>? Origin { get; set; }

        /// <summary>
        /// The unique identifier from the system of origin. Typically, a sid, object id or Guid. e.g. Used for member of other tenant on Azure Active Directory.
        /// </summary>
        [Input("originId")]
        public Input<string>? OriginId { get; set; }

        public GroupEntitlementArgs()
        {
        }
        public static new GroupEntitlementArgs Empty => new GroupEntitlementArgs();
    }

    public sealed class GroupEntitlementState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Type of Account License. Valid values: `advanced`, `earlyAdopter`, `express`, `none`, `professional`, or `stakeholder`. Defaults to `express`. In addition, the value `basic` is allowed which is an alias for `express` and reflects the name of the `express` license used in the Azure DevOps web interface.
        /// </summary>
        [Input("accountLicenseType")]
        public Input<string>? AccountLicenseType { get; set; }

        /// <summary>
        /// The descriptor is the primary way to reference the graph subject while the system is running. This field will uniquely identify the group graph subject.
        /// </summary>
        [Input("descriptor")]
        public Input<string>? Descriptor { get; set; }

        /// <summary>
        /// The display name is the name used in Azure DevOps UI. Cannot be set together with `origin_id` and `origin`.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// The source of the licensing (e.g. Account. MSDN etc.) Valid values: `account` (Default), `auto`, `msdn`, `none`, `profile`, `trial`
        /// 
        /// &gt; **NOTE:** A existing group in Azure AD can only be referenced by the combination of `origin_id` and `origin`.
        /// </summary>
        [Input("licensingSource")]
        public Input<string>? LicensingSource { get; set; }

        /// <summary>
        /// The type of source provider for the origin identifier.
        /// </summary>
        [Input("origin")]
        public Input<string>? Origin { get; set; }

        /// <summary>
        /// The unique identifier from the system of origin. Typically, a sid, object id or Guid. e.g. Used for member of other tenant on Azure Active Directory.
        /// </summary>
        [Input("originId")]
        public Input<string>? OriginId { get; set; }

        /// <summary>
        /// The principal name of a graph member on Azure DevOps
        /// </summary>
        [Input("principalName")]
        public Input<string>? PrincipalName { get; set; }

        public GroupEntitlementState()
        {
        }
        public static new GroupEntitlementState Empty => new GroupEntitlementState();
    }
}
