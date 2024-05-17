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
    /// Manages creation of the Feed within Azure DevOps organization.
    /// 
    /// ## Example Usage
    /// 
    /// ### Create Feed in the scope of whole Organization
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AzureDevOps = Pulumi.AzureDevOps;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new AzureDevOps.Feed("example", new()
    ///     {
    ///         Name = "releases",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ### Create Feed in the scope of a Project
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
    ///         Visibility = "private",
    ///         VersionControl = "Git",
    ///         WorkItemTemplate = "Agile",
    ///         Description = "Managed by Terraform",
    ///     });
    /// 
    ///     var exampleFeed = new AzureDevOps.Feed("example", new()
    ///     {
    ///         Name = "releases",
    ///         ProjectId = example.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ### Create Feed with Soft Delete
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AzureDevOps = Pulumi.AzureDevOps;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new AzureDevOps.Feed("example", new()
    ///     {
    ///         Name = "releases",
    ///         Features = new[]
    ///         {
    ///             new AzureDevOps.Inputs.FeedFeatureArgs
    ///             {
    ///                 PermanentDelete = false,
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Relevant Links
    /// 
    /// - [Azure DevOps Service REST API 7.0 - Feed Management](https://learn.microsoft.com/en-us/rest/api/azure/devops/artifacts/feed-management?view=azure-devops-rest-7.0)
    /// </summary>
    [AzureDevOpsResourceType("azuredevops:index/feed:Feed")]
    public partial class Feed : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A `features` blocks as documented below.
        /// 
        /// &gt; **Note** *Because of ADO limitations feed name can be **reserved** for up to 15 minutes after permanent delete of the feed*
        /// </summary>
        [Output("features")]
        public Output<ImmutableArray<Outputs.FeedFeature>> Features { get; private set; } = null!;

        /// <summary>
        /// The name of the Feed.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The ID of the Project Feed is created in. If not specified, feed will be created at the organization level.
        /// </summary>
        [Output("projectId")]
        public Output<string?> ProjectId { get; private set; } = null!;


        /// <summary>
        /// Create a Feed resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Feed(string name, FeedArgs? args = null, CustomResourceOptions? options = null)
            : base("azuredevops:index/feed:Feed", name, args ?? new FeedArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Feed(string name, Input<string> id, FeedState? state = null, CustomResourceOptions? options = null)
            : base("azuredevops:index/feed:Feed", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Feed resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Feed Get(string name, Input<string> id, FeedState? state = null, CustomResourceOptions? options = null)
        {
            return new Feed(name, id, state, options);
        }
    }

    public sealed class FeedArgs : global::Pulumi.ResourceArgs
    {
        [Input("features")]
        private InputList<Inputs.FeedFeatureArgs>? _features;

        /// <summary>
        /// A `features` blocks as documented below.
        /// 
        /// &gt; **Note** *Because of ADO limitations feed name can be **reserved** for up to 15 minutes after permanent delete of the feed*
        /// </summary>
        public InputList<Inputs.FeedFeatureArgs> Features
        {
            get => _features ?? (_features = new InputList<Inputs.FeedFeatureArgs>());
            set => _features = value;
        }

        /// <summary>
        /// The name of the Feed.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the Project Feed is created in. If not specified, feed will be created at the organization level.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        public FeedArgs()
        {
        }
        public static new FeedArgs Empty => new FeedArgs();
    }

    public sealed class FeedState : global::Pulumi.ResourceArgs
    {
        [Input("features")]
        private InputList<Inputs.FeedFeatureGetArgs>? _features;

        /// <summary>
        /// A `features` blocks as documented below.
        /// 
        /// &gt; **Note** *Because of ADO limitations feed name can be **reserved** for up to 15 minutes after permanent delete of the feed*
        /// </summary>
        public InputList<Inputs.FeedFeatureGetArgs> Features
        {
            get => _features ?? (_features = new InputList<Inputs.FeedFeatureGetArgs>());
            set => _features = value;
        }

        /// <summary>
        /// The name of the Feed.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the Project Feed is created in. If not specified, feed will be created at the organization level.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        public FeedState()
        {
        }
        public static new FeedState Empty => new FeedState();
    }
}
