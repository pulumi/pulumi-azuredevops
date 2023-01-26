// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureDevOps
{
    public static class GetGitRepository
    {
        /// <summary>
        /// Use this data source to access information about a **single** (existing) Git Repository within Azure DevOps.
        /// To read information about **multiple** Git Repositories use the data source `azuredevops.getRepositories`
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
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
        ///     var example_single_repo = AzureDevOps.GetGitRepository.Invoke(new()
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
        public static Task<GetGitRepositoryResult> InvokeAsync(GetGitRepositoryArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetGitRepositoryResult>("azuredevops:index/getGitRepository:getGitRepository", args ?? new GetGitRepositoryArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to access information about a **single** (existing) Git Repository within Azure DevOps.
        /// To read information about **multiple** Git Repositories use the data source `azuredevops.getRepositories`
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
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
        ///     var example_single_repo = AzureDevOps.GetGitRepository.Invoke(new()
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
        public static Output<GetGitRepositoryResult> Invoke(GetGitRepositoryInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetGitRepositoryResult>("azuredevops:index/getGitRepository:getGitRepository", args ?? new GetGitRepositoryInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetGitRepositoryArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the Git repository to retrieve
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// ID of project to list Git repositories
        /// </summary>
        [Input("projectId", required: true)]
        public string ProjectId { get; set; } = null!;

        public GetGitRepositoryArgs()
        {
        }
        public static new GetGitRepositoryArgs Empty => new GetGitRepositoryArgs();
    }

    public sealed class GetGitRepositoryInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the Git repository to retrieve
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// ID of project to list Git repositories
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        public GetGitRepositoryInvokeArgs()
        {
        }
        public static new GetGitRepositoryInvokeArgs Empty => new GetGitRepositoryInvokeArgs();
    }


    [OutputType]
    public sealed class GetGitRepositoryResult
    {
        /// <summary>
        /// The ref of the default branch.
        /// </summary>
        public readonly string DefaultBranch;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly bool IsFork;
        /// <summary>
        /// Git repository name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Project identifier to which the Git repository belongs.
        /// </summary>
        public readonly string ProjectId;
        /// <summary>
        /// HTTPS Url to clone the Git repository
        /// </summary>
        public readonly string RemoteUrl;
        /// <summary>
        /// Compressed size (bytes) of the repository.
        /// </summary>
        public readonly int Size;
        /// <summary>
        /// SSH Url to clone the Git repository
        /// </summary>
        public readonly string SshUrl;
        /// <summary>
        /// Details REST API endpoint for the Git Repository.
        /// </summary>
        public readonly string Url;
        /// <summary>
        /// Url of the Git repository web view
        /// </summary>
        public readonly string WebUrl;

        [OutputConstructor]
        private GetGitRepositoryResult(
            string defaultBranch,

            string id,

            bool isFork,

            string name,

            string projectId,

            string remoteUrl,

            int size,

            string sshUrl,

            string url,

            string webUrl)
        {
            DefaultBranch = defaultBranch;
            Id = id;
            IsFork = isFork;
            Name = name;
            ProjectId = projectId;
            RemoteUrl = remoteUrl;
            Size = size;
            SshUrl = sshUrl;
            Url = url;
            WebUrl = webUrl;
        }
    }
}
