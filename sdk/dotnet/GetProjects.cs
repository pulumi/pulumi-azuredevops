// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureDevOps
{
    public static class GetProjects
    {
        /// <summary>
        /// Use this data source to access information about existing Projects within Azure DevOps.
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
        ///     var example = AzureDevOps.GetProjects.Invoke(new()
        ///     {
        ///         Name = "Example Project",
        ///         State = "wellFormed",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["projectId"] = example.Apply(getProjectsResult =&gt; getProjectsResult.Projects).Select(__item =&gt; __item.ProjectId).ToList(),
        ///         ["name"] = example.Apply(getProjectsResult =&gt; getProjectsResult.Projects).Select(__item =&gt; __item.Name).ToList(),
        ///         ["projectUrl"] = example.Apply(getProjectsResult =&gt; getProjectsResult.Projects).Select(__item =&gt; __item.ProjectUrl).ToList(),
        ///         ["state"] = example.Apply(getProjectsResult =&gt; getProjectsResult.Projects).Select(__item =&gt; __item.State).ToList(),
        ///     };
        /// });
        /// ```
        /// 
        /// ## Relevant Links
        /// 
        /// - [Azure DevOps Service REST API 7.0 - Projects - Get](https://docs.microsoft.com/en-us/rest/api/azure/devops/core/projects/get?view=azure-devops-rest-7.0)
        /// </summary>
        public static Task<GetProjectsResult> InvokeAsync(GetProjectsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetProjectsResult>("azuredevops:index/getProjects:getProjects", args ?? new GetProjectsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to access information about existing Projects within Azure DevOps.
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
        ///     var example = AzureDevOps.GetProjects.Invoke(new()
        ///     {
        ///         Name = "Example Project",
        ///         State = "wellFormed",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["projectId"] = example.Apply(getProjectsResult =&gt; getProjectsResult.Projects).Select(__item =&gt; __item.ProjectId).ToList(),
        ///         ["name"] = example.Apply(getProjectsResult =&gt; getProjectsResult.Projects).Select(__item =&gt; __item.Name).ToList(),
        ///         ["projectUrl"] = example.Apply(getProjectsResult =&gt; getProjectsResult.Projects).Select(__item =&gt; __item.ProjectUrl).ToList(),
        ///         ["state"] = example.Apply(getProjectsResult =&gt; getProjectsResult.Projects).Select(__item =&gt; __item.State).ToList(),
        ///     };
        /// });
        /// ```
        /// 
        /// ## Relevant Links
        /// 
        /// - [Azure DevOps Service REST API 7.0 - Projects - Get](https://docs.microsoft.com/en-us/rest/api/azure/devops/core/projects/get?view=azure-devops-rest-7.0)
        /// </summary>
        public static Output<GetProjectsResult> Invoke(GetProjectsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetProjectsResult>("azuredevops:index/getProjects:getProjects", args ?? new GetProjectsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetProjectsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the Project, if not specified all projects will be returned.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// State of the Project, if not specified all projects will be returned. Valid values are `all`, `deleting`, `new`, `wellFormed`, `createPending`, `unchanged`,`deleted`.
        /// 
        /// DataSource without specifying any arguments will return all projects.
        /// </summary>
        [Input("state")]
        public string? State { get; set; }

        public GetProjectsArgs()
        {
        }
        public static new GetProjectsArgs Empty => new GetProjectsArgs();
    }

    public sealed class GetProjectsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the Project, if not specified all projects will be returned.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// State of the Project, if not specified all projects will be returned. Valid values are `all`, `deleting`, `new`, `wellFormed`, `createPending`, `unchanged`,`deleted`.
        /// 
        /// DataSource without specifying any arguments will return all projects.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        public GetProjectsInvokeArgs()
        {
        }
        public static new GetProjectsInvokeArgs Empty => new GetProjectsInvokeArgs();
    }


    [OutputType]
    public sealed class GetProjectsResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the Project.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// A list of existing projects in your Azure DevOps Organization with details about every project which includes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetProjectsProjectResult> Projects;
        /// <summary>
        /// Project state.
        /// </summary>
        public readonly string? State;

        [OutputConstructor]
        private GetProjectsResult(
            string id,

            string? name,

            ImmutableArray<Outputs.GetProjectsProjectResult> projects,

            string? state)
        {
            Id = id;
            Name = name;
            Projects = projects;
            State = state;
        }
    }
}
