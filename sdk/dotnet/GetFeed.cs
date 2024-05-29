// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureDevOps
{
    public static class GetFeed
    {
        /// <summary>
        /// Use this data source to access information about existing Feed within a given project in Azure DevOps.
        /// 
        /// ## Example Usage
        /// 
        /// ### Basic Example
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using AzureDevOps = Pulumi.AzureDevOps;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = AzureDevOps.GetFeed.Invoke(new()
        ///     {
        ///         Name = "releases",
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// ### Access feed within a project
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using AzureDevOps = Pulumi.AzureDevOps;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = AzureDevOps.GetProject.Invoke(new()
        ///     {
        ///         Name = "Example Project",
        ///     });
        /// 
        ///     var exampleGetFeed = AzureDevOps.GetFeed.Invoke(new()
        ///     {
        ///         Name = "releases",
        ///         ProjectId = example.Apply(getProjectResult =&gt; getProjectResult.Id),
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// 
        /// ## Relevant Links
        /// 
        /// - [Azure DevOps Service REST API 7.0 - Feed - Get](https://learn.microsoft.com/en-us/rest/api/azure/devops/artifacts/feed-management/get-feed?view=azure-devops-rest-7.0)
        /// </summary>
        public static Task<GetFeedResult> InvokeAsync(GetFeedArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetFeedResult>("azuredevops:index/getFeed:getFeed", args ?? new GetFeedArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to access information about existing Feed within a given project in Azure DevOps.
        /// 
        /// ## Example Usage
        /// 
        /// ### Basic Example
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using AzureDevOps = Pulumi.AzureDevOps;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = AzureDevOps.GetFeed.Invoke(new()
        ///     {
        ///         Name = "releases",
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// ### Access feed within a project
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using AzureDevOps = Pulumi.AzureDevOps;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = AzureDevOps.GetProject.Invoke(new()
        ///     {
        ///         Name = "Example Project",
        ///     });
        /// 
        ///     var exampleGetFeed = AzureDevOps.GetFeed.Invoke(new()
        ///     {
        ///         Name = "releases",
        ///         ProjectId = example.Apply(getProjectResult =&gt; getProjectResult.Id),
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// 
        /// ## Relevant Links
        /// 
        /// - [Azure DevOps Service REST API 7.0 - Feed - Get](https://learn.microsoft.com/en-us/rest/api/azure/devops/artifacts/feed-management/get-feed?view=azure-devops-rest-7.0)
        /// </summary>
        public static Output<GetFeedResult> Invoke(GetFeedInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetFeedResult>("azuredevops:index/getFeed:getFeed", args ?? new GetFeedInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFeedArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// ID of the Feed.
        /// 
        /// &gt; **Note** Only one of `name` or `feed_id` can be set at the same time.
        /// </summary>
        [Input("feedId")]
        public string? FeedId { get; set; }

        /// <summary>
        /// Name of the Feed.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// ID of the Project Feed is created in.
        /// </summary>
        [Input("projectId")]
        public string? ProjectId { get; set; }

        public GetFeedArgs()
        {
        }
        public static new GetFeedArgs Empty => new GetFeedArgs();
    }

    public sealed class GetFeedInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// ID of the Feed.
        /// 
        /// &gt; **Note** Only one of `name` or `feed_id` can be set at the same time.
        /// </summary>
        [Input("feedId")]
        public Input<string>? FeedId { get; set; }

        /// <summary>
        /// Name of the Feed.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// ID of the Project Feed is created in.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        public GetFeedInvokeArgs()
        {
        }
        public static new GetFeedInvokeArgs Empty => new GetFeedInvokeArgs();
    }


    [OutputType]
    public sealed class GetFeedResult
    {
        /// <summary>
        /// The ID of the Feed.
        /// </summary>
        public readonly string? FeedId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the Feed.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The ID of the Project.
        /// </summary>
        public readonly string? ProjectId;

        [OutputConstructor]
        private GetFeedResult(
            string? feedId,

            string id,

            string? name,

            string? projectId)
        {
            FeedId = feedId;
            Id = id;
            Name = name;
            ProjectId = projectId;
        }
    }
}