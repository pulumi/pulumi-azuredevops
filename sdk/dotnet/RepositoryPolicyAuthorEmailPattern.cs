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
    /// Manage author email pattern repository policy within Azure DevOps project.
    /// 
    /// ## Example Usage
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
    ///             Visibility = "private",
    ///             VersionControl = "Git",
    ///             WorkItemTemplate = "Agile",
    ///             Description = "Managed by Terraform",
    ///         });
    ///         var exampleGit = new AzureDevOps.Git("exampleGit", new AzureDevOps.GitArgs
    ///         {
    ///             ProjectId = exampleProject.Id,
    ///             Initialization = new AzureDevOps.Inputs.GitInitializationArgs
    ///             {
    ///                 InitType = "Clean",
    ///             },
    ///         });
    ///         var exampleRepositoryPolicyAuthorEmailPattern = new AzureDevOps.RepositoryPolicyAuthorEmailPattern("exampleRepositoryPolicyAuthorEmailPattern", new AzureDevOps.RepositoryPolicyAuthorEmailPatternArgs
    ///         {
    ///             ProjectId = exampleProject.Id,
    ///             Enabled = true,
    ///             Blocking = true,
    ///             AuthorEmailPatterns = 
    ///             {
    ///                 "user1@test.com",
    ///                 "user2@test.com",
    ///             },
    ///             RepositoryIds = 
    ///             {
    ///                 exampleGit.Id,
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ## Set project level repository policy
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
    ///             Visibility = "private",
    ///             VersionControl = "Git",
    ///             WorkItemTemplate = "Agile",
    ///             Description = "Managed by Terraform",
    ///         });
    ///         var exampleRepositoryPolicyAuthorEmailPattern = new AzureDevOps.RepositoryPolicyAuthorEmailPattern("exampleRepositoryPolicyAuthorEmailPattern", new AzureDevOps.RepositoryPolicyAuthorEmailPatternArgs
    ///         {
    ///             ProjectId = exampleProject.Id,
    ///             Enabled = true,
    ///             Blocking = true,
    ///             AuthorEmailPatterns = 
    ///             {
    ///                 "user1@test.com",
    ///                 "user2@test.com",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Relevant Links
    /// 
    /// - [Azure DevOps Service REST API 6.0 - Policy Configurations](https://docs.microsoft.com/en-us/rest/api/azure/devops/policy/configurations?view=azure-devops-rest-6.0)
    /// 
    /// ## Import
    /// 
    /// Azure DevOps Branch Policies can be imported using the project ID and policy configuration ID
    /// 
    /// ```sh
    ///  $ pulumi import azuredevops:index/repositoryPolicyAuthorEmailPattern:RepositoryPolicyAuthorEmailPattern example 00000000-0000-0000-0000-000000000000/0
    /// ```
    /// </summary>
    [AzureDevOpsResourceType("azuredevops:index/repositoryPolicyAuthorEmailPattern:RepositoryPolicyAuthorEmailPattern")]
    public partial class RepositoryPolicyAuthorEmailPattern : Pulumi.CustomResource
    {
        /// <summary>
        /// Block pushes with a commit author email that does not match the patterns. You can specify exact emails or use wildcards. 
        /// Email patterns prefixed with "!" are excluded. Order is important.
        /// </summary>
        [Output("authorEmailPatterns")]
        public Output<ImmutableArray<string>> AuthorEmailPatterns { get; private set; } = null!;

        /// <summary>
        /// A flag indicating if the policy should be blocking. Defaults to `true`.
        /// </summary>
        [Output("blocking")]
        public Output<bool?> Blocking { get; private set; } = null!;

        /// <summary>
        /// A flag indicating if the policy should be enabled. Defaults to `true`.
        /// </summary>
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the policy will be created.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// Control whether the policy is enabled for the repository or the project. If `repository_ids` not configured, the policy will be set to the project.
        /// </summary>
        [Output("repositoryIds")]
        public Output<ImmutableArray<string>> RepositoryIds { get; private set; } = null!;


        /// <summary>
        /// Create a RepositoryPolicyAuthorEmailPattern resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RepositoryPolicyAuthorEmailPattern(string name, RepositoryPolicyAuthorEmailPatternArgs args, CustomResourceOptions? options = null)
            : base("azuredevops:index/repositoryPolicyAuthorEmailPattern:RepositoryPolicyAuthorEmailPattern", name, args ?? new RepositoryPolicyAuthorEmailPatternArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RepositoryPolicyAuthorEmailPattern(string name, Input<string> id, RepositoryPolicyAuthorEmailPatternState? state = null, CustomResourceOptions? options = null)
            : base("azuredevops:index/repositoryPolicyAuthorEmailPattern:RepositoryPolicyAuthorEmailPattern", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RepositoryPolicyAuthorEmailPattern resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RepositoryPolicyAuthorEmailPattern Get(string name, Input<string> id, RepositoryPolicyAuthorEmailPatternState? state = null, CustomResourceOptions? options = null)
        {
            return new RepositoryPolicyAuthorEmailPattern(name, id, state, options);
        }
    }

    public sealed class RepositoryPolicyAuthorEmailPatternArgs : Pulumi.ResourceArgs
    {
        [Input("authorEmailPatterns", required: true)]
        private InputList<string>? _authorEmailPatterns;

        /// <summary>
        /// Block pushes with a commit author email that does not match the patterns. You can specify exact emails or use wildcards. 
        /// Email patterns prefixed with "!" are excluded. Order is important.
        /// </summary>
        public InputList<string> AuthorEmailPatterns
        {
            get => _authorEmailPatterns ?? (_authorEmailPatterns = new InputList<string>());
            set => _authorEmailPatterns = value;
        }

        /// <summary>
        /// A flag indicating if the policy should be blocking. Defaults to `true`.
        /// </summary>
        [Input("blocking")]
        public Input<bool>? Blocking { get; set; }

        /// <summary>
        /// A flag indicating if the policy should be enabled. Defaults to `true`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The ID of the project in which the policy will be created.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        [Input("repositoryIds")]
        private InputList<string>? _repositoryIds;

        /// <summary>
        /// Control whether the policy is enabled for the repository or the project. If `repository_ids` not configured, the policy will be set to the project.
        /// </summary>
        public InputList<string> RepositoryIds
        {
            get => _repositoryIds ?? (_repositoryIds = new InputList<string>());
            set => _repositoryIds = value;
        }

        public RepositoryPolicyAuthorEmailPatternArgs()
        {
        }
    }

    public sealed class RepositoryPolicyAuthorEmailPatternState : Pulumi.ResourceArgs
    {
        [Input("authorEmailPatterns")]
        private InputList<string>? _authorEmailPatterns;

        /// <summary>
        /// Block pushes with a commit author email that does not match the patterns. You can specify exact emails or use wildcards. 
        /// Email patterns prefixed with "!" are excluded. Order is important.
        /// </summary>
        public InputList<string> AuthorEmailPatterns
        {
            get => _authorEmailPatterns ?? (_authorEmailPatterns = new InputList<string>());
            set => _authorEmailPatterns = value;
        }

        /// <summary>
        /// A flag indicating if the policy should be blocking. Defaults to `true`.
        /// </summary>
        [Input("blocking")]
        public Input<bool>? Blocking { get; set; }

        /// <summary>
        /// A flag indicating if the policy should be enabled. Defaults to `true`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The ID of the project in which the policy will be created.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        [Input("repositoryIds")]
        private InputList<string>? _repositoryIds;

        /// <summary>
        /// Control whether the policy is enabled for the repository or the project. If `repository_ids` not configured, the policy will be set to the project.
        /// </summary>
        public InputList<string> RepositoryIds
        {
            get => _repositoryIds ?? (_repositoryIds = new InputList<string>());
            set => _repositoryIds = value;
        }

        public RepositoryPolicyAuthorEmailPatternState()
        {
        }
    }
}
