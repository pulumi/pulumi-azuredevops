// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureDevOps.Repository
{
    /// <summary>
    /// Manages a git repository within Azure DevOps.
    /// 
    /// ## Example Usage
    /// ### Create Git repository
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
    ///             Visibility = "private",
    ///             VersionControl = "Git",
    ///             WorkItemTemplate = "Agile",
    ///         });
    ///         var repo = new AzureDevOps.Git("repo", new AzureDevOps.GitArgs
    ///         {
    ///             ProjectId = project.Id,
    ///             Initialization = new AzureDevOps.Inputs.GitInitializationArgs
    ///             {
    ///                 InitType = "Clean",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Create Fork of another Azure DevOps Git repository
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using AzureDevOps = Pulumi.AzureDevOps;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var repo = new AzureDevOps.Git("repo", new AzureDevOps.GitArgs
    ///         {
    ///             ProjectId = azuredevops_project.Project.Id,
    ///             ParentRepositoryId = azuredevops_git_repository.Parent.Id,
    ///             Initialization = new AzureDevOps.Inputs.GitInitializationArgs
    ///             {
    ///                 InitType = "Clean",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Create Import from another Git repository
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using AzureDevOps = Pulumi.AzureDevOps;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var repo = new AzureDevOps.Git("repo", new AzureDevOps.GitArgs
    ///         {
    ///             ProjectId = azuredevops_project.Project.Id,
    ///             Initialization = new AzureDevOps.Inputs.GitInitializationArgs
    ///             {
    ///                 InitType = "Import",
    ///                 SourceType = "Git",
    ///                 SourceUrl = "https://github.com/microsoft/terraform-provider-azuredevops.git",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Import from a Private Repository
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using AzureDevOps = Pulumi.AzureDevOps;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var serviceendpoint = new AzureDevOps.ServiceEndpointGenericGit("serviceendpoint", new AzureDevOps.ServiceEndpointGenericGitArgs
    ///         {
    ///             ProjectId = azuredevops_project.Project.Id,
    ///             RepositoryUrl = "https://dev.azure.com/org/project/_git/repository",
    ///             Username = "username",
    ///             Password = "&lt;password&gt;/&lt;PAT&gt;",
    ///             ServiceEndpointName = "Sample Generic Git",
    ///             Description = "Managed by Terraform",
    ///         });
    ///         var repo = new AzureDevOps.Git("repo", new AzureDevOps.GitArgs
    ///         {
    ///             ProjectId = azuredevops_project.Project.Id,
    ///             Initialization = new AzureDevOps.Inputs.GitInitializationArgs
    ///             {
    ///                 InitType = "Import",
    ///                 SourceType = "Git",
    ///                 SourceUrl = "https://dev.azure.com/example-org/private-repository.git",
    ///                 ServiceConnectionId = serviceendpoint.Id,
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ## Relevant Links
    /// 
    /// - [Azure DevOps Service REST API 5.1 - Git Repositories](https://docs.microsoft.com/en-us/rest/api/azure/devops/git/repositories?view=azure-devops-rest-5.1)
    /// 
    /// ## Import
    /// 
    /// Azure DevOps Repositories can be imported using the repo name or by the repo Guid e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azuredevops:Repository/git:Git repository projectName/repoName
    /// ```
    /// 
    ///  or
    /// 
    /// ```sh
    ///  $ pulumi import azuredevops:Repository/git:Git repository projectName/00000000-0000-0000-0000-000000000000
    /// ```
    /// </summary>
    [Obsolete(@"azuredevops.repository.Git has been deprecated in favor of azuredevops.Git")]
    [AzureDevOpsResourceType("azuredevops:Repository/git:Git")]
    public partial class Git : Pulumi.CustomResource
    {
        /// <summary>
        /// The ref of the default branch. Will be used as the branch name for initialized repositories.
        /// </summary>
        [Output("defaultBranch")]
        public Output<string> DefaultBranch { get; private set; } = null!;

        /// <summary>
        /// An `initialization` block as documented below.
        /// </summary>
        [Output("initialization")]
        public Output<Outputs.GitInitialization> Initialization { get; private set; } = null!;

        /// <summary>
        /// True if the repository was created as a fork.
        /// </summary>
        [Output("isFork")]
        public Output<bool> IsFork { get; private set; } = null!;

        /// <summary>
        /// The name of the git repository.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The ID of a Git project from which a fork is to be created.
        /// </summary>
        [Output("parentRepositoryId")]
        public Output<string?> ParentRepositoryId { get; private set; } = null!;

        /// <summary>
        /// The project ID or project name.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// Git HTTPS URL of the repository
        /// </summary>
        [Output("remoteUrl")]
        public Output<string> RemoteUrl { get; private set; } = null!;

        /// <summary>
        /// Size in bytes.
        /// </summary>
        [Output("size")]
        public Output<int> Size { get; private set; } = null!;

        /// <summary>
        /// Git SSH URL of the repository.
        /// </summary>
        [Output("sshUrl")]
        public Output<string> SshUrl { get; private set; } = null!;

        /// <summary>
        /// REST API URL of the repository.
        /// </summary>
        [Output("url")]
        public Output<string> Url { get; private set; } = null!;

        /// <summary>
        /// Web link to the repository.
        /// </summary>
        [Output("webUrl")]
        public Output<string> WebUrl { get; private set; } = null!;


        /// <summary>
        /// Create a Git resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Git(string name, GitArgs args, CustomResourceOptions? options = null)
            : base("azuredevops:Repository/git:Git", name, args ?? new GitArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Git(string name, Input<string> id, GitState? state = null, CustomResourceOptions? options = null)
            : base("azuredevops:Repository/git:Git", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Git resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Git Get(string name, Input<string> id, GitState? state = null, CustomResourceOptions? options = null)
        {
            return new Git(name, id, state, options);
        }
    }

    public sealed class GitArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ref of the default branch. Will be used as the branch name for initialized repositories.
        /// </summary>
        [Input("defaultBranch")]
        public Input<string>? DefaultBranch { get; set; }

        /// <summary>
        /// An `initialization` block as documented below.
        /// </summary>
        [Input("initialization", required: true)]
        public Input<Inputs.GitInitializationArgs> Initialization { get; set; } = null!;

        /// <summary>
        /// The name of the git repository.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of a Git project from which a fork is to be created.
        /// </summary>
        [Input("parentRepositoryId")]
        public Input<string>? ParentRepositoryId { get; set; }

        /// <summary>
        /// The project ID or project name.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        public GitArgs()
        {
        }
    }

    public sealed class GitState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ref of the default branch. Will be used as the branch name for initialized repositories.
        /// </summary>
        [Input("defaultBranch")]
        public Input<string>? DefaultBranch { get; set; }

        /// <summary>
        /// An `initialization` block as documented below.
        /// </summary>
        [Input("initialization")]
        public Input<Inputs.GitInitializationGetArgs>? Initialization { get; set; }

        /// <summary>
        /// True if the repository was created as a fork.
        /// </summary>
        [Input("isFork")]
        public Input<bool>? IsFork { get; set; }

        /// <summary>
        /// The name of the git repository.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of a Git project from which a fork is to be created.
        /// </summary>
        [Input("parentRepositoryId")]
        public Input<string>? ParentRepositoryId { get; set; }

        /// <summary>
        /// The project ID or project name.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// Git HTTPS URL of the repository
        /// </summary>
        [Input("remoteUrl")]
        public Input<string>? RemoteUrl { get; set; }

        /// <summary>
        /// Size in bytes.
        /// </summary>
        [Input("size")]
        public Input<int>? Size { get; set; }

        /// <summary>
        /// Git SSH URL of the repository.
        /// </summary>
        [Input("sshUrl")]
        public Input<string>? SshUrl { get; set; }

        /// <summary>
        /// REST API URL of the repository.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        /// <summary>
        /// Web link to the repository.
        /// </summary>
        [Input("webUrl")]
        public Input<string>? WebUrl { get; set; }

        public GitState()
        {
        }
    }
}
