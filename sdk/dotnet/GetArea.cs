// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureDevOps
{
    public static class GetArea
    {
        /// <summary>
        /// Use this data source to access information about an existing Area (Component) within Azure DevOps.
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
        ///     var exampleProject = new AzureDevOps.Project("example", new()
        ///     {
        ///         Name = "Example Project",
        ///         WorkItemTemplate = "Agile",
        ///         VersionControl = "Git",
        ///         Visibility = "private",
        ///         Description = "Managed by Pulumi",
        ///     });
        /// 
        ///     var example = AzureDevOps.GetArea.Invoke(new()
        ///     {
        ///         ProjectId = exampleProject.Id,
        ///         Path = "/",
        ///         FetchChildren = false,
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// ## Relevant Links
        /// 
        /// - [Azure DevOps Service REST API 7.0 - Classification Nodes - Get Classification Nodes](https://docs.microsoft.com/en-us/rest/api/azure/devops/wit/classification-nodes/create-or-update?view=azure-devops-rest-7.0)
        /// 
        /// ## PAT Permissions Required
        /// 
        /// - **Project &amp; Team**: vso.work - Grants the ability to read work items, queries, boards, area and iterations paths, and other work item tracking related metadata. Also grants the ability to execute queries, search work items and to receive notifications about work item events via service hooks.
        /// </summary>
        public static Task<GetAreaResult> InvokeAsync(GetAreaArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAreaResult>("azuredevops:index/getArea:getArea", args ?? new GetAreaArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to access information about an existing Area (Component) within Azure DevOps.
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
        ///     var exampleProject = new AzureDevOps.Project("example", new()
        ///     {
        ///         Name = "Example Project",
        ///         WorkItemTemplate = "Agile",
        ///         VersionControl = "Git",
        ///         Visibility = "private",
        ///         Description = "Managed by Pulumi",
        ///     });
        /// 
        ///     var example = AzureDevOps.GetArea.Invoke(new()
        ///     {
        ///         ProjectId = exampleProject.Id,
        ///         Path = "/",
        ///         FetchChildren = false,
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// ## Relevant Links
        /// 
        /// - [Azure DevOps Service REST API 7.0 - Classification Nodes - Get Classification Nodes](https://docs.microsoft.com/en-us/rest/api/azure/devops/wit/classification-nodes/create-or-update?view=azure-devops-rest-7.0)
        /// 
        /// ## PAT Permissions Required
        /// 
        /// - **Project &amp; Team**: vso.work - Grants the ability to read work items, queries, boards, area and iterations paths, and other work item tracking related metadata. Also grants the ability to execute queries, search work items and to receive notifications about work item events via service hooks.
        /// </summary>
        public static Output<GetAreaResult> Invoke(GetAreaInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAreaResult>("azuredevops:index/getArea:getArea", args ?? new GetAreaInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to access information about an existing Area (Component) within Azure DevOps.
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
        ///     var exampleProject = new AzureDevOps.Project("example", new()
        ///     {
        ///         Name = "Example Project",
        ///         WorkItemTemplate = "Agile",
        ///         VersionControl = "Git",
        ///         Visibility = "private",
        ///         Description = "Managed by Pulumi",
        ///     });
        /// 
        ///     var example = AzureDevOps.GetArea.Invoke(new()
        ///     {
        ///         ProjectId = exampleProject.Id,
        ///         Path = "/",
        ///         FetchChildren = false,
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// ## Relevant Links
        /// 
        /// - [Azure DevOps Service REST API 7.0 - Classification Nodes - Get Classification Nodes](https://docs.microsoft.com/en-us/rest/api/azure/devops/wit/classification-nodes/create-or-update?view=azure-devops-rest-7.0)
        /// 
        /// ## PAT Permissions Required
        /// 
        /// - **Project &amp; Team**: vso.work - Grants the ability to read work items, queries, boards, area and iterations paths, and other work item tracking related metadata. Also grants the ability to execute queries, search work items and to receive notifications about work item events via service hooks.
        /// </summary>
        public static Output<GetAreaResult> Invoke(GetAreaInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetAreaResult>("azuredevops:index/getArea:getArea", args ?? new GetAreaInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAreaArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Read children nodes, _Depth_: 1, _Default_: `true`
        /// </summary>
        [Input("fetchChildren")]
        public bool? FetchChildren { get; set; }

        /// <summary>
        /// The path to the Area; _Format_: URL relative; if omitted, or value `"/"` is used, the root Area will be returned
        /// </summary>
        [Input("path")]
        public string? Path { get; set; }

        /// <summary>
        /// The project ID.
        /// </summary>
        [Input("projectId", required: true)]
        public string ProjectId { get; set; } = null!;

        public GetAreaArgs()
        {
        }
        public static new GetAreaArgs Empty => new GetAreaArgs();
    }

    public sealed class GetAreaInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Read children nodes, _Depth_: 1, _Default_: `true`
        /// </summary>
        [Input("fetchChildren")]
        public Input<bool>? FetchChildren { get; set; }

        /// <summary>
        /// The path to the Area; _Format_: URL relative; if omitted, or value `"/"` is used, the root Area will be returned
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        /// <summary>
        /// The project ID.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        public GetAreaInvokeArgs()
        {
        }
        public static new GetAreaInvokeArgs Empty => new GetAreaInvokeArgs();
    }


    [OutputType]
    public sealed class GetAreaResult
    {
        /// <summary>
        /// A list of `children` blocks as defined below, empty if `has_children == false`
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAreaChildrenResult> Childrens;
        public readonly bool? FetchChildren;
        /// <summary>
        /// Indicator if the child Area node has child nodes
        /// </summary>
        public readonly bool HasChildren;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the child Area node
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The complete path (in relative URL format) of the child Area
        /// </summary>
        public readonly string Path;
        /// <summary>
        /// The ID of project.
        /// </summary>
        public readonly string ProjectId;

        [OutputConstructor]
        private GetAreaResult(
            ImmutableArray<Outputs.GetAreaChildrenResult> childrens,

            bool? fetchChildren,

            bool hasChildren,

            string id,

            string name,

            string path,

            string projectId)
        {
            Childrens = childrens;
            FetchChildren = fetchChildren;
            HasChildren = hasChildren;
            Id = id;
            Name = name;
            Path = path;
            ProjectId = projectId;
        }
    }
}
