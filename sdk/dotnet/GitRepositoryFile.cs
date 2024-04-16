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
    /// Manage files within an Azure DevOps Git repository.
    /// 
    /// ## Example Usage
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
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
    ///     });
    /// 
    ///     var exampleGit = new AzureDevOps.Git("example", new()
    ///     {
    ///         ProjectId = example.Id,
    ///         Name = "Example Git Repository",
    ///         Initialization = new AzureDevOps.Inputs.GitInitializationArgs
    ///         {
    ///             InitType = "Clean",
    ///         },
    ///     });
    /// 
    ///     var exampleGitRepositoryFile = new AzureDevOps.GitRepositoryFile("example", new()
    ///     {
    ///         RepositoryId = exampleGit.Id,
    ///         File = ".gitignore",
    ///         Content = "**/*.tfstate",
    ///         Branch = "refs/heads/master",
    ///         CommitMessage = "First commit",
    ///         OverwriteOnCreate = false,
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Relevant Links
    /// 
    /// - [Azure DevOps Service REST API 7.0 - Git API](https://docs.microsoft.com/en-us/rest/api/azure/devops/git/?view=azure-devops-rest-7.0)
    /// 
    /// ## Import
    /// 
    /// Repository files can be imported using a combination of the `repository ID` and `file`, e.g.
    /// 
    /// ```sh
    /// $ pulumi import azuredevops:index/gitRepositoryFile:GitRepositoryFile example 00000000-0000-0000-0000-000000000000/.gitignore
    /// ```
    /// 
    /// To import a file from a branch other than `master`, append `:` and the branch name, e.g.
    /// 
    /// ```sh
    /// $ pulumi import azuredevops:index/gitRepositoryFile:GitRepositoryFile example 00000000-0000-0000-0000-000000000000/.gitignore:refs/heads/master
    /// ```
    /// </summary>
    [AzureDevOpsResourceType("azuredevops:index/gitRepositoryFile:GitRepositoryFile")]
    public partial class GitRepositoryFile : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Git branch (defaults to `refs/heads/master`). The branch must already exist, it will not be created if it
        /// does not already exist.
        /// </summary>
        [Output("branch")]
        public Output<string?> Branch { get; private set; } = null!;

        /// <summary>
        /// Commit message when adding or updating the managed file.
        /// </summary>
        [Output("commitMessage")]
        public Output<string> CommitMessage { get; private set; } = null!;

        /// <summary>
        /// The file content.
        /// </summary>
        [Output("content")]
        public Output<string> Content { get; private set; } = null!;

        /// <summary>
        /// The path of the file to manage.
        /// </summary>
        [Output("file")]
        public Output<string> File { get; private set; } = null!;

        /// <summary>
        /// Enable overwriting existing files (defaults to `false`).
        /// </summary>
        [Output("overwriteOnCreate")]
        public Output<bool?> OverwriteOnCreate { get; private set; } = null!;

        /// <summary>
        /// The ID of the Git repository.
        /// </summary>
        [Output("repositoryId")]
        public Output<string> RepositoryId { get; private set; } = null!;


        /// <summary>
        /// Create a GitRepositoryFile resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GitRepositoryFile(string name, GitRepositoryFileArgs args, CustomResourceOptions? options = null)
            : base("azuredevops:index/gitRepositoryFile:GitRepositoryFile", name, args ?? new GitRepositoryFileArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GitRepositoryFile(string name, Input<string> id, GitRepositoryFileState? state = null, CustomResourceOptions? options = null)
            : base("azuredevops:index/gitRepositoryFile:GitRepositoryFile", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing GitRepositoryFile resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GitRepositoryFile Get(string name, Input<string> id, GitRepositoryFileState? state = null, CustomResourceOptions? options = null)
        {
            return new GitRepositoryFile(name, id, state, options);
        }
    }

    public sealed class GitRepositoryFileArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Git branch (defaults to `refs/heads/master`). The branch must already exist, it will not be created if it
        /// does not already exist.
        /// </summary>
        [Input("branch")]
        public Input<string>? Branch { get; set; }

        /// <summary>
        /// Commit message when adding or updating the managed file.
        /// </summary>
        [Input("commitMessage")]
        public Input<string>? CommitMessage { get; set; }

        /// <summary>
        /// The file content.
        /// </summary>
        [Input("content", required: true)]
        public Input<string> Content { get; set; } = null!;

        /// <summary>
        /// The path of the file to manage.
        /// </summary>
        [Input("file", required: true)]
        public Input<string> File { get; set; } = null!;

        /// <summary>
        /// Enable overwriting existing files (defaults to `false`).
        /// </summary>
        [Input("overwriteOnCreate")]
        public Input<bool>? OverwriteOnCreate { get; set; }

        /// <summary>
        /// The ID of the Git repository.
        /// </summary>
        [Input("repositoryId", required: true)]
        public Input<string> RepositoryId { get; set; } = null!;

        public GitRepositoryFileArgs()
        {
        }
        public static new GitRepositoryFileArgs Empty => new GitRepositoryFileArgs();
    }

    public sealed class GitRepositoryFileState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Git branch (defaults to `refs/heads/master`). The branch must already exist, it will not be created if it
        /// does not already exist.
        /// </summary>
        [Input("branch")]
        public Input<string>? Branch { get; set; }

        /// <summary>
        /// Commit message when adding or updating the managed file.
        /// </summary>
        [Input("commitMessage")]
        public Input<string>? CommitMessage { get; set; }

        /// <summary>
        /// The file content.
        /// </summary>
        [Input("content")]
        public Input<string>? Content { get; set; }

        /// <summary>
        /// The path of the file to manage.
        /// </summary>
        [Input("file")]
        public Input<string>? File { get; set; }

        /// <summary>
        /// Enable overwriting existing files (defaults to `false`).
        /// </summary>
        [Input("overwriteOnCreate")]
        public Input<bool>? OverwriteOnCreate { get; set; }

        /// <summary>
        /// The ID of the Git repository.
        /// </summary>
        [Input("repositoryId")]
        public Input<string>? RepositoryId { get; set; }

        public GitRepositoryFileState()
        {
        }
        public static new GitRepositoryFileState Empty => new GitRepositoryFileState();
    }
}
