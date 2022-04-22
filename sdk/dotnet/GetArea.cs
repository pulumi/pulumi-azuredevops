// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
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
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using AzureDevOps = Pulumi.AzureDevOps;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var exampleProject = new AzureDevOps.Project("exampleProject", new AzureDevOps.ProjectArgs
        ///         {
        ///             WorkItemTemplate = "Agile",
        ///             VersionControl = "Git",
        ///             Visibility = "private",
        ///             Description = "Managed by Terraform",
        ///         });
        ///         var exampleArea = exampleProject.Id.Apply(id =&gt; AzureDevOps.GetArea.Invoke(new AzureDevOps.GetAreaInvokeArgs
        ///         {
        ///             ProjectId = id,
        ///             Path = "/",
        ///             FetchChildren = false,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// ## Relevant Links
        /// 
        /// - [Azure DevOps Service REST API 6.0 - Classification Nodes - Get Classification Nodes](https://docs.microsoft.com/en-us/rest/api/azure/devops/wit/classification-nodes/create-or-update?view=azure-devops-rest-6.0)
        /// 
        /// ## PAT Permissions Required
        /// 
        /// - **Project &amp; Team**: vso.work - Grants the ability to read work items, queries, boards, area and iterations paths, and other work item tracking related metadata. Also grants the ability to execute queries, search work items and to receive notifications about work item events via service hooks.
        /// </summary>
        public static Task<GetAreaResult> InvokeAsync(GetAreaArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAreaResult>("azuredevops:index/getArea:getArea", args ?? new GetAreaArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to access information about an existing Area (Component) within Azure DevOps.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using AzureDevOps = Pulumi.AzureDevOps;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var exampleProject = new AzureDevOps.Project("exampleProject", new AzureDevOps.ProjectArgs
        ///         {
        ///             WorkItemTemplate = "Agile",
        ///             VersionControl = "Git",
        ///             Visibility = "private",
        ///             Description = "Managed by Terraform",
        ///         });
        ///         var exampleArea = exampleProject.Id.Apply(id =&gt; AzureDevOps.GetArea.Invoke(new AzureDevOps.GetAreaInvokeArgs
        ///         {
        ///             ProjectId = id,
        ///             Path = "/",
        ///             FetchChildren = false,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// ## Relevant Links
        /// 
        /// - [Azure DevOps Service REST API 6.0 - Classification Nodes - Get Classification Nodes](https://docs.microsoft.com/en-us/rest/api/azure/devops/wit/classification-nodes/create-or-update?view=azure-devops-rest-6.0)
        /// 
        /// ## PAT Permissions Required
        /// 
        /// - **Project &amp; Team**: vso.work - Grants the ability to read work items, queries, boards, area and iterations paths, and other work item tracking related metadata. Also grants the ability to execute queries, search work items and to receive notifications about work item events via service hooks.
        /// </summary>
        public static Output<GetAreaResult> Invoke(GetAreaInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetAreaResult>("azuredevops:index/getArea:getArea", args ?? new GetAreaInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAreaArgs : Pulumi.InvokeArgs
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
    }

    public sealed class GetAreaInvokeArgs : Pulumi.InvokeArgs
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
        /// The project ID of the child Area node
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
