// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureDevOps.Repository
{
    [Obsolete(@"azuredevops.repository.getRepositories has been deprecated in favor of azuredevops.getRepositories")]
    public static class GetRepositories
    {
        /// <summary>
        /// Use this data source to access information about **multiple** existing Git Repositories within Azure DevOps.
        /// To read informations about a **single** Git Repository use the data source `azuredevops.Git`
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
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
        ///     var example_all_repos = AzureDevOps.GetRepositories.Invoke(new()
        ///     {
        ///         ProjectId = example.Apply(getProjectResult =&gt; getProjectResult.Id),
        ///         IncludeHidden = true,
        ///     });
        /// 
        ///     var example_single_repo = AzureDevOps.GetRepositories.Invoke(new()
        ///     {
        ///         ProjectId = example.Apply(getProjectResult =&gt; getProjectResult.Id),
        ///         Name = "Example Repository",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// ## Relevant Links
        /// 
        /// - [Azure DevOps Service REST API 6.0 - Git API](https://docs.microsoft.com/en-us/rest/api/azure/devops/git/?view=azure-devops-rest-6.0)
        /// </summary>
        public static Task<GetRepositoriesResult> InvokeAsync(GetRepositoriesArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetRepositoriesResult>("azuredevops:Repository/getRepositories:getRepositories", args ?? new GetRepositoriesArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to access information about **multiple** existing Git Repositories within Azure DevOps.
        /// To read informations about a **single** Git Repository use the data source `azuredevops.Git`
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
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
        ///     var example_all_repos = AzureDevOps.GetRepositories.Invoke(new()
        ///     {
        ///         ProjectId = example.Apply(getProjectResult =&gt; getProjectResult.Id),
        ///         IncludeHidden = true,
        ///     });
        /// 
        ///     var example_single_repo = AzureDevOps.GetRepositories.Invoke(new()
        ///     {
        ///         ProjectId = example.Apply(getProjectResult =&gt; getProjectResult.Id),
        ///         Name = "Example Repository",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// ## Relevant Links
        /// 
        /// - [Azure DevOps Service REST API 6.0 - Git API](https://docs.microsoft.com/en-us/rest/api/azure/devops/git/?view=azure-devops-rest-6.0)
        /// </summary>
        public static Output<GetRepositoriesResult> Invoke(GetRepositoriesInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetRepositoriesResult>("azuredevops:Repository/getRepositories:getRepositories", args ?? new GetRepositoriesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRepositoriesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// DataSource without specifying any arguments will return all Git repositories of an organization.
        /// </summary>
        [Input("includeHidden")]
        public bool? IncludeHidden { get; set; }

        /// <summary>
        /// Name of the Git repository to retrieve; requires `project_id` to be specified as well
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// ID of project to list Git repositories
        /// </summary>
        [Input("projectId")]
        public string? ProjectId { get; set; }

        public GetRepositoriesArgs()
        {
        }
        public static new GetRepositoriesArgs Empty => new GetRepositoriesArgs();
    }

    public sealed class GetRepositoriesInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// DataSource without specifying any arguments will return all Git repositories of an organization.
        /// </summary>
        [Input("includeHidden")]
        public Input<bool>? IncludeHidden { get; set; }

        /// <summary>
        /// Name of the Git repository to retrieve; requires `project_id` to be specified as well
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// ID of project to list Git repositories
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        public GetRepositoriesInvokeArgs()
        {
        }
        public static new GetRepositoriesInvokeArgs Empty => new GetRepositoriesInvokeArgs();
    }


    [OutputType]
    public sealed class GetRepositoriesResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly bool? IncludeHidden;
        /// <summary>
        /// Git repository name.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Project identifier to which the Git repository belongs.
        /// </summary>
        public readonly string? ProjectId;
        /// <summary>
        /// A list of existing projects in your Azure DevOps Organization with details about every project which includes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetRepositoriesRepositoryResult> Repositories;

        [OutputConstructor]
        private GetRepositoriesResult(
            string id,

            bool? includeHidden,

            string? name,

            string? projectId,

            ImmutableArray<Outputs.GetRepositoriesRepositoryResult> repositories)
        {
            Id = id;
            IncludeHidden = includeHidden;
            Name = name;
            ProjectId = projectId;
            Repositories = repositories;
        }
    }
}
