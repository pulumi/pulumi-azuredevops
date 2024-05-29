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
    /// Manages assignment of security roles to various resources within Azure DevOps organization.
    /// </summary>
    [AzureDevOpsResourceType("azuredevops:index/securityroleAssignment:SecurityroleAssignment")]
    public partial class SecurityroleAssignment : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the identity to authorize.
        /// </summary>
        [Output("identityId")]
        public Output<string> IdentityId { get; private set; } = null!;

        /// <summary>
        /// The ID of the resource on which the role is to be assigned.
        /// </summary>
        [Output("resourceId")]
        public Output<string> ResourceId { get; private set; } = null!;

        /// <summary>
        /// Name of the role to assign.
        /// </summary>
        [Output("roleName")]
        public Output<string> RoleName { get; private set; } = null!;

        /// <summary>
        /// The scope in which this assignment should exist.
        /// </summary>
        [Output("scope")]
        public Output<string> Scope { get; private set; } = null!;


        /// <summary>
        /// Create a SecurityroleAssignment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecurityroleAssignment(string name, SecurityroleAssignmentArgs args, CustomResourceOptions? options = null)
            : base("azuredevops:index/securityroleAssignment:SecurityroleAssignment", name, args ?? new SecurityroleAssignmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SecurityroleAssignment(string name, Input<string> id, SecurityroleAssignmentState? state = null, CustomResourceOptions? options = null)
            : base("azuredevops:index/securityroleAssignment:SecurityroleAssignment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SecurityroleAssignment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SecurityroleAssignment Get(string name, Input<string> id, SecurityroleAssignmentState? state = null, CustomResourceOptions? options = null)
        {
            return new SecurityroleAssignment(name, id, state, options);
        }
    }

    public sealed class SecurityroleAssignmentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the identity to authorize.
        /// </summary>
        [Input("identityId", required: true)]
        public Input<string> IdentityId { get; set; } = null!;

        /// <summary>
        /// The ID of the resource on which the role is to be assigned.
        /// </summary>
        [Input("resourceId", required: true)]
        public Input<string> ResourceId { get; set; } = null!;

        /// <summary>
        /// Name of the role to assign.
        /// </summary>
        [Input("roleName", required: true)]
        public Input<string> RoleName { get; set; } = null!;

        /// <summary>
        /// The scope in which this assignment should exist.
        /// </summary>
        [Input("scope", required: true)]
        public Input<string> Scope { get; set; } = null!;

        public SecurityroleAssignmentArgs()
        {
        }
        public static new SecurityroleAssignmentArgs Empty => new SecurityroleAssignmentArgs();
    }

    public sealed class SecurityroleAssignmentState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the identity to authorize.
        /// </summary>
        [Input("identityId")]
        public Input<string>? IdentityId { get; set; }

        /// <summary>
        /// The ID of the resource on which the role is to be assigned.
        /// </summary>
        [Input("resourceId")]
        public Input<string>? ResourceId { get; set; }

        /// <summary>
        /// Name of the role to assign.
        /// </summary>
        [Input("roleName")]
        public Input<string>? RoleName { get; set; }

        /// <summary>
        /// The scope in which this assignment should exist.
        /// </summary>
        [Input("scope")]
        public Input<string>? Scope { get; set; }

        public SecurityroleAssignmentState()
        {
        }
        public static new SecurityroleAssignmentState Empty => new SecurityroleAssignmentState();
    }
}