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
    /// Manages a Git Repository Branch.
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
    ///     var exampleGitRepositoryBranch = new AzureDevOps.GitRepositoryBranch("example", new()
    ///     {
    ///         RepositoryId = exampleGit.Id,
    ///         Name = "example-branch-name",
    ///         RefBranch = exampleGit.DefaultBranch,
    ///     });
    /// 
    ///     var exampleFromCommitId = new AzureDevOps.GitRepositoryBranch("example_from_commit_id", new()
    ///     {
    ///         RepositoryId = exampleGit.Id,
    ///         Name = "example-from-commit-id",
    ///         RefCommitId = exampleGitRepositoryBranch.LastCommitId,
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [AzureDevOpsResourceType("azuredevops:index/gitRepositoryBranch:GitRepositoryBranch")]
    public partial class GitRepositoryBranch : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The commit object ID of last commit on the branch.
        /// </summary>
        [Output("lastCommitId")]
        public Output<string> LastCommitId { get; private set; } = null!;

        /// <summary>
        /// The name of the branch in short format not prefixed with `refs/heads/`.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The reference to the source branch to create the branch from, in `&lt;name&gt;` or `refs/heads/&lt;name&gt;` format. Conflict with `ref_tag`, `ref_commit_id`.
        /// </summary>
        [Output("refBranch")]
        public Output<string?> RefBranch { get; private set; } = null!;

        /// <summary>
        /// The commit object ID to create the branch from. Conflict with `ref_branch`, `ref_tag`.
        /// </summary>
        [Output("refCommitId")]
        public Output<string?> RefCommitId { get; private set; } = null!;

        /// <summary>
        /// The reference to the tag to create the branch from, in `&lt;name&gt;` or `refs/tags/&lt;name&gt;` format. Conflict with `ref_branch`, `ref_commit_id`.
        /// </summary>
        [Output("refTag")]
        public Output<string?> RefTag { get; private set; } = null!;

        /// <summary>
        /// The ID of the repository the branch is created against.
        /// </summary>
        [Output("repositoryId")]
        public Output<string> RepositoryId { get; private set; } = null!;


        /// <summary>
        /// Create a GitRepositoryBranch resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GitRepositoryBranch(string name, GitRepositoryBranchArgs args, CustomResourceOptions? options = null)
            : base("azuredevops:index/gitRepositoryBranch:GitRepositoryBranch", name, args ?? new GitRepositoryBranchArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GitRepositoryBranch(string name, Input<string> id, GitRepositoryBranchState? state = null, CustomResourceOptions? options = null)
            : base("azuredevops:index/gitRepositoryBranch:GitRepositoryBranch", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing GitRepositoryBranch resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GitRepositoryBranch Get(string name, Input<string> id, GitRepositoryBranchState? state = null, CustomResourceOptions? options = null)
        {
            return new GitRepositoryBranch(name, id, state, options);
        }
    }

    public sealed class GitRepositoryBranchArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the branch in short format not prefixed with `refs/heads/`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The reference to the source branch to create the branch from, in `&lt;name&gt;` or `refs/heads/&lt;name&gt;` format. Conflict with `ref_tag`, `ref_commit_id`.
        /// </summary>
        [Input("refBranch")]
        public Input<string>? RefBranch { get; set; }

        /// <summary>
        /// The commit object ID to create the branch from. Conflict with `ref_branch`, `ref_tag`.
        /// </summary>
        [Input("refCommitId")]
        public Input<string>? RefCommitId { get; set; }

        /// <summary>
        /// The reference to the tag to create the branch from, in `&lt;name&gt;` or `refs/tags/&lt;name&gt;` format. Conflict with `ref_branch`, `ref_commit_id`.
        /// </summary>
        [Input("refTag")]
        public Input<string>? RefTag { get; set; }

        /// <summary>
        /// The ID of the repository the branch is created against.
        /// </summary>
        [Input("repositoryId", required: true)]
        public Input<string> RepositoryId { get; set; } = null!;

        public GitRepositoryBranchArgs()
        {
        }
        public static new GitRepositoryBranchArgs Empty => new GitRepositoryBranchArgs();
    }

    public sealed class GitRepositoryBranchState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The commit object ID of last commit on the branch.
        /// </summary>
        [Input("lastCommitId")]
        public Input<string>? LastCommitId { get; set; }

        /// <summary>
        /// The name of the branch in short format not prefixed with `refs/heads/`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The reference to the source branch to create the branch from, in `&lt;name&gt;` or `refs/heads/&lt;name&gt;` format. Conflict with `ref_tag`, `ref_commit_id`.
        /// </summary>
        [Input("refBranch")]
        public Input<string>? RefBranch { get; set; }

        /// <summary>
        /// The commit object ID to create the branch from. Conflict with `ref_branch`, `ref_tag`.
        /// </summary>
        [Input("refCommitId")]
        public Input<string>? RefCommitId { get; set; }

        /// <summary>
        /// The reference to the tag to create the branch from, in `&lt;name&gt;` or `refs/tags/&lt;name&gt;` format. Conflict with `ref_branch`, `ref_commit_id`.
        /// </summary>
        [Input("refTag")]
        public Input<string>? RefTag { get; set; }

        /// <summary>
        /// The ID of the repository the branch is created against.
        /// </summary>
        [Input("repositoryId")]
        public Input<string>? RepositoryId { get; set; }

        public GitRepositoryBranchState()
        {
        }
        public static new GitRepositoryBranchState Empty => new GitRepositoryBranchState();
    }
}
