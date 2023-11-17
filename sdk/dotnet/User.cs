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
    /// Manages a user entitlement within Azure DevOps.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AzureDevOps = Pulumi.AzureDevOps;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new AzureDevOps.User("example", new()
    ///     {
    ///         PrincipalName = "foo@contoso.com",
    ///     });
    /// 
    /// });
    /// ```
    /// ## Relevant Links
    /// 
    /// - [Azure DevOps Service REST API 7.0 - User Entitlements - Add](https://docs.microsoft.com/en-us/rest/api/azure/devops/memberentitlementmanagement/user-entitlements/add?view=azure-devops-rest-7.0)
    /// - [Programmatic mapping of access levels](https://docs.microsoft.com/en-us/azure/devops/organizations/security/access-levels?view=azure-devops#programmatic-mapping-of-access-levels)
    /// 
    /// ## PAT Permissions Required
    /// 
    /// - **Member Entitlement Management**: Read &amp; Write
    /// 
    /// ## Import
    /// 
    /// The resources allows the import via the UUID of a user entitlement or by using the principal name of a user owning an entitlement.
    /// </summary>
    [AzureDevOpsResourceType("azuredevops:index/user:User")]
    public partial class User : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Type of Account License. Valid values: `advanced`, `earlyAdopter`, `express`, `none`, `professional`, or `stakeholder`. Defaults to `express`. In addition the value `basic` is allowed which is an alias for `express` and reflects the name of the `express` license used in the Azure DevOps web interface.
        /// </summary>
        [Output("accountLicenseType")]
        public Output<string?> AccountLicenseType { get; private set; } = null!;

        /// <summary>
        /// The descriptor is the primary way to reference the graph subject while the system is running. This field will uniquely identify the user graph subject.
        /// </summary>
        [Output("descriptor")]
        public Output<string> Descriptor { get; private set; } = null!;

        /// <summary>
        /// The source of the licensing (e.g. Account. MSDN etc.) Valid values: `account` (Default), `auto`, `msdn`, `none`, `profile`, `trial`**NOTE:** A user can only be referenced by it's `principal_name` or by the combination of `origin_id` and `origin`.
        /// </summary>
        [Output("licensingSource")]
        public Output<string?> LicensingSource { get; private set; } = null!;

        /// <summary>
        /// The type of source provider for the origin identifier.
        /// </summary>
        [Output("origin")]
        public Output<string> Origin { get; private set; } = null!;

        /// <summary>
        /// The unique identifier from the system of origin. Typically a sid, object id or Guid. e.g. Used for member of other tenant on Azure Active Directory.
        /// </summary>
        [Output("originId")]
        public Output<string> OriginId { get; private set; } = null!;

        /// <summary>
        /// The principal name is the PrincipalName of a graph member from the source provider. Usually, e-mail address.
        /// </summary>
        [Output("principalName")]
        public Output<string> PrincipalName { get; private set; } = null!;


        /// <summary>
        /// Create a User resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public User(string name, UserArgs? args = null, CustomResourceOptions? options = null)
            : base("azuredevops:index/user:User", name, args ?? new UserArgs(), MakeResourceOptions(options, ""))
        {
        }

        private User(string name, Input<string> id, UserState? state = null, CustomResourceOptions? options = null)
            : base("azuredevops:index/user:User", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azuredevops:Entitlement/user:User"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing User resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static User Get(string name, Input<string> id, UserState? state = null, CustomResourceOptions? options = null)
        {
            return new User(name, id, state, options);
        }
    }

    public sealed class UserArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Type of Account License. Valid values: `advanced`, `earlyAdopter`, `express`, `none`, `professional`, or `stakeholder`. Defaults to `express`. In addition the value `basic` is allowed which is an alias for `express` and reflects the name of the `express` license used in the Azure DevOps web interface.
        /// </summary>
        [Input("accountLicenseType")]
        public Input<string>? AccountLicenseType { get; set; }

        /// <summary>
        /// The source of the licensing (e.g. Account. MSDN etc.) Valid values: `account` (Default), `auto`, `msdn`, `none`, `profile`, `trial`**NOTE:** A user can only be referenced by it's `principal_name` or by the combination of `origin_id` and `origin`.
        /// </summary>
        [Input("licensingSource")]
        public Input<string>? LicensingSource { get; set; }

        /// <summary>
        /// The type of source provider for the origin identifier.
        /// </summary>
        [Input("origin")]
        public Input<string>? Origin { get; set; }

        /// <summary>
        /// The unique identifier from the system of origin. Typically a sid, object id or Guid. e.g. Used for member of other tenant on Azure Active Directory.
        /// </summary>
        [Input("originId")]
        public Input<string>? OriginId { get; set; }

        /// <summary>
        /// The principal name is the PrincipalName of a graph member from the source provider. Usually, e-mail address.
        /// </summary>
        [Input("principalName")]
        public Input<string>? PrincipalName { get; set; }

        public UserArgs()
        {
        }
        public static new UserArgs Empty => new UserArgs();
    }

    public sealed class UserState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Type of Account License. Valid values: `advanced`, `earlyAdopter`, `express`, `none`, `professional`, or `stakeholder`. Defaults to `express`. In addition the value `basic` is allowed which is an alias for `express` and reflects the name of the `express` license used in the Azure DevOps web interface.
        /// </summary>
        [Input("accountLicenseType")]
        public Input<string>? AccountLicenseType { get; set; }

        /// <summary>
        /// The descriptor is the primary way to reference the graph subject while the system is running. This field will uniquely identify the user graph subject.
        /// </summary>
        [Input("descriptor")]
        public Input<string>? Descriptor { get; set; }

        /// <summary>
        /// The source of the licensing (e.g. Account. MSDN etc.) Valid values: `account` (Default), `auto`, `msdn`, `none`, `profile`, `trial`**NOTE:** A user can only be referenced by it's `principal_name` or by the combination of `origin_id` and `origin`.
        /// </summary>
        [Input("licensingSource")]
        public Input<string>? LicensingSource { get; set; }

        /// <summary>
        /// The type of source provider for the origin identifier.
        /// </summary>
        [Input("origin")]
        public Input<string>? Origin { get; set; }

        /// <summary>
        /// The unique identifier from the system of origin. Typically a sid, object id or Guid. e.g. Used for member of other tenant on Azure Active Directory.
        /// </summary>
        [Input("originId")]
        public Input<string>? OriginId { get; set; }

        /// <summary>
        /// The principal name is the PrincipalName of a graph member from the source provider. Usually, e-mail address.
        /// </summary>
        [Input("principalName")]
        public Input<string>? PrincipalName { get; set; }

        public UserState()
        {
        }
        public static new UserState Empty => new UserState();
    }
}
