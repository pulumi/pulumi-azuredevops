// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureDevOps
{
    public static class GetIteration
    {
        /// <summary>
        /// Use this data source to access information about an existing Iteration (Sprint) within Azure DevOps.
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
        ///         var project = new AzureDevOps.Project("project", new AzureDevOps.ProjectArgs
        ///         {
        ///             WorkItemTemplate = "Agile",
        ///             VersionControl = "Git",
        ///             Visibility = "private",
        ///             Description = "Managed by Terraform",
        ///         });
        ///         var root_iteration = project.Id.Apply(id =&gt; AzureDevOps.GetIteration.InvokeAsync(new AzureDevOps.GetIterationArgs
        ///         {
        ///             ProjectId = id,
        ///             Path = "/",
        ///             FetchChildren = true,
        ///         }));
        ///         var child_iteration = project.Id.Apply(id =&gt; AzureDevOps.GetIteration.InvokeAsync(new AzureDevOps.GetIterationArgs
        ///         {
        ///             ProjectId = id,
        ///             Path = "/Iteration 1",
        ///             FetchChildren = true,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// ## Relevant Links
        /// 
        /// - [Azure DevOps Service REST API 5.1 - Classification Nodes - Get Classification Nodes](https://docs.microsoft.com/en-us/rest/api/azure/devops/wit/classification%20nodes/get%20classification%20nodes?view=azure-devops-rest-5.1)
        /// 
        /// ## PAT Permissions Required
        /// 
        /// - **Project &amp; Team**: vso.work - Grants the ability to read work items, queries, boards, area and iterations paths, and other work item tracking related metadata. Also grants the ability to execute queries, search work items and to receive notifications about work item events via service hooks.
        /// </summary>
        public static Task<GetIterationResult> InvokeAsync(GetIterationArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetIterationResult>("azuredevops:index/getIteration:getIteration", args ?? new GetIterationArgs(), options.WithVersion());
    }


    public sealed class GetIterationArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Read children nodes, _Depth_: 1, _Default_: `true`
        /// </summary>
        [Input("fetchChildren")]
        public bool? FetchChildren { get; set; }

        /// <summary>
        /// The path to the Iteration, _Format_: URL relative; if omitted, or value `"/"` is used, the root Iteration will be returned
        /// </summary>
        [Input("path")]
        public string? Path { get; set; }

        /// <summary>
        /// The project ID.
        /// </summary>
        [Input("projectId", required: true)]
        public string ProjectId { get; set; } = null!;

        public GetIterationArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetIterationResult
    {
        /// <summary>
        /// A list of `children` blocks as defined below, empty if `has_children == false`
        /// </summary>
        public readonly ImmutableArray<Outputs.GetIterationChildrenResult> Childrens;
        public readonly bool? FetchChildren;
        /// <summary>
        /// Indicator if the child Iteration node has child nodes
        /// </summary>
        public readonly bool HasChildren;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the child Iteration node
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The complete path (in relative URL format) of the child Iteration
        /// </summary>
        public readonly string Path;
        /// <summary>
        /// The project ID of the child Iteration node
        /// </summary>
        public readonly string ProjectId;

        [OutputConstructor]
        private GetIterationResult(
            ImmutableArray<Outputs.GetIterationChildrenResult> childrens,

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
