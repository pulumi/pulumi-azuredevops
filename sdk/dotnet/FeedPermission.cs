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
    /// Manages creation of the Feed Permission within Azure DevOps organization.
    /// 
    /// ## Example Usage
    /// 
    /// ### Create Feed Permission
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AzureDevOps = Pulumi.AzureDevOps;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new AzureDevOps.Project("example", new()
    ///     {
    ///         Name = "Example Project",
    ///     });
    /// 
    ///     var exampleGroup = new AzureDevOps.Group("example", new()
    ///     {
    ///         Scope = example.Id,
    ///         DisplayName = "Example group",
    ///         Description = "Example description",
    ///     });
    /// 
    ///     var exampleFeed = new AzureDevOps.Feed("example", new()
    ///     {
    ///         Name = "releases",
    ///     });
    /// 
    ///     var permission = new AzureDevOps.FeedPermission("permission", new()
    ///     {
    ///         FeedId = exampleFeed.Id,
    ///         Role = "reader",
    ///         IdentityDescriptor = exampleGroup.Descriptor,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Relevant Links
    /// 
    /// - [Azure DevOps Service REST API 7.0 - Feed Management](https://learn.microsoft.com/en-us/rest/api/azure/devops/artifacts/feed-management?view=azure-devops-rest-7.0)
    /// </summary>
    [AzureDevOpsResourceType("azuredevops:index/feedPermission:FeedPermission")]
    public partial class FeedPermission : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The display name of the assignment
        /// </summary>
        [Output("displayName")]
        public Output<string?> DisplayName { get; private set; } = null!;

        /// <summary>
        /// The ID of the Feed.
        /// </summary>
        [Output("feedId")]
        public Output<string> FeedId { get; private set; } = null!;

        /// <summary>
        /// The Descriptor of identity you want to assign a role.
        /// </summary>
        [Output("identityDescriptor")]
        public Output<string> IdentityDescriptor { get; private set; } = null!;

        /// <summary>
        /// The ID of the identity.
        /// </summary>
        [Output("identityId")]
        public Output<string> IdentityId { get; private set; } = null!;

        /// <summary>
        /// The ID of the Project Feed is created in. If not specified, feed will be created at the organization level.
        /// </summary>
        [Output("projectId")]
        public Output<string?> ProjectId { get; private set; } = null!;

        /// <summary>
        /// The role to be assigned, possible values : `reader`, `contributor`, `collaborator`, `administrator`
        /// </summary>
        [Output("role")]
        public Output<string> Role { get; private set; } = null!;


        /// <summary>
        /// Create a FeedPermission resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FeedPermission(string name, FeedPermissionArgs args, CustomResourceOptions? options = null)
            : base("azuredevops:index/feedPermission:FeedPermission", name, args ?? new FeedPermissionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FeedPermission(string name, Input<string> id, FeedPermissionState? state = null, CustomResourceOptions? options = null)
            : base("azuredevops:index/feedPermission:FeedPermission", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing FeedPermission resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FeedPermission Get(string name, Input<string> id, FeedPermissionState? state = null, CustomResourceOptions? options = null)
        {
            return new FeedPermission(name, id, state, options);
        }
    }

    public sealed class FeedPermissionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The display name of the assignment
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// The ID of the Feed.
        /// </summary>
        [Input("feedId", required: true)]
        public Input<string> FeedId { get; set; } = null!;

        /// <summary>
        /// The Descriptor of identity you want to assign a role.
        /// </summary>
        [Input("identityDescriptor", required: true)]
        public Input<string> IdentityDescriptor { get; set; } = null!;

        /// <summary>
        /// The ID of the Project Feed is created in. If not specified, feed will be created at the organization level.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The role to be assigned, possible values : `reader`, `contributor`, `collaborator`, `administrator`
        /// </summary>
        [Input("role", required: true)]
        public Input<string> Role { get; set; } = null!;

        public FeedPermissionArgs()
        {
        }
        public static new FeedPermissionArgs Empty => new FeedPermissionArgs();
    }

    public sealed class FeedPermissionState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The display name of the assignment
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// The ID of the Feed.
        /// </summary>
        [Input("feedId")]
        public Input<string>? FeedId { get; set; }

        /// <summary>
        /// The Descriptor of identity you want to assign a role.
        /// </summary>
        [Input("identityDescriptor")]
        public Input<string>? IdentityDescriptor { get; set; }

        /// <summary>
        /// The ID of the identity.
        /// </summary>
        [Input("identityId")]
        public Input<string>? IdentityId { get; set; }

        /// <summary>
        /// The ID of the Project Feed is created in. If not specified, feed will be created at the organization level.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The role to be assigned, possible values : `reader`, `contributor`, `collaborator`, `administrator`
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        public FeedPermissionState()
        {
        }
        public static new FeedPermissionState Empty => new FeedPermissionState();
    }
}
