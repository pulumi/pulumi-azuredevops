// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureDevOps
{
    /// <summary>
    /// Manage a file path pattern repository policy within Azure DevOps project.
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
    ///         Description = "Managed by Pulumi",
    ///     });
    /// 
    ///     var exampleGit = new AzureDevOps.Git("example", new()
    ///     {
    ///         ProjectId = example.Id,
    ///         Name = "Example Repository",
    ///         Initialization = new AzureDevOps.Inputs.GitInitializationArgs
    ///         {
    ///             InitType = "Clean",
    ///         },
    ///     });
    /// 
    ///     var exampleRepositoryPolicyFilePathPattern = new AzureDevOps.RepositoryPolicyFilePathPattern("example", new()
    ///     {
    ///         ProjectId = example.Id,
    ///         Enabled = true,
    ///         Blocking = true,
    ///         FilepathPatterns = new[]
    ///         {
    ///             "*.go",
    ///             "/home/test/*.ts",
    ///         },
    ///         RepositoryIds = new[]
    ///         {
    ///             exampleGit.Id,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// # Set project level repository policy
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
    ///         Description = "Managed by Pulumi",
    ///     });
    /// 
    ///     var examplep = new AzureDevOps.RepositoryPolicyFilePathPattern("examplep", new()
    ///     {
    ///         ProjectId = example.Id,
    ///         Enabled = true,
    ///         Blocking = true,
    ///         FilepathPatterns = new[]
    ///         {
    ///             "*.go",
    ///             "/home/test/*.ts",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Relevant Links
    /// 
    /// - [Azure DevOps Service REST API 7.0 - Policy Configurations](https://docs.microsoft.com/en-us/rest/api/azure/devops/policy/configurations?view=azure-devops-rest-7.0)
    /// 
    /// ## Import
    /// 
    /// Azure DevOps repository policies can be imported using the projectID/policyID or projectName/policyID:
    /// 
    /// ```sh
    /// $ pulumi import azuredevops:index/repositoryPolicyFilePathPattern:RepositoryPolicyFilePathPattern example 00000000-0000-0000-0000-000000000000/0
    /// ```
    /// </summary>
    [AzureDevOpsResourceType("azuredevops:index/repositoryPolicyFilePathPattern:RepositoryPolicyFilePathPattern")]
    public partial class RepositoryPolicyFilePathPattern : global::Pulumi.CustomResource
    {
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
        /// Block pushes from introducing file paths that match the following patterns. Exact paths begin with "/". You can specify exact paths and wildcards. You can also specify multiple paths using ";" as a separator. Paths prefixed with "!" are excluded. Order is important.
        /// </summary>
        [Output("filepathPatterns")]
        public Output<ImmutableArray<string>> FilepathPatterns { get; private set; } = null!;

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
        /// Create a RepositoryPolicyFilePathPattern resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RepositoryPolicyFilePathPattern(string name, RepositoryPolicyFilePathPatternArgs args, CustomResourceOptions? options = null)
            : base("azuredevops:index/repositoryPolicyFilePathPattern:RepositoryPolicyFilePathPattern", name, args ?? new RepositoryPolicyFilePathPatternArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RepositoryPolicyFilePathPattern(string name, Input<string> id, RepositoryPolicyFilePathPatternState? state = null, CustomResourceOptions? options = null)
            : base("azuredevops:index/repositoryPolicyFilePathPattern:RepositoryPolicyFilePathPattern", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RepositoryPolicyFilePathPattern resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RepositoryPolicyFilePathPattern Get(string name, Input<string> id, RepositoryPolicyFilePathPatternState? state = null, CustomResourceOptions? options = null)
        {
            return new RepositoryPolicyFilePathPattern(name, id, state, options);
        }
    }

    public sealed class RepositoryPolicyFilePathPatternArgs : global::Pulumi.ResourceArgs
    {
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

        [Input("filepathPatterns", required: true)]
        private InputList<string>? _filepathPatterns;

        /// <summary>
        /// Block pushes from introducing file paths that match the following patterns. Exact paths begin with "/". You can specify exact paths and wildcards. You can also specify multiple paths using ";" as a separator. Paths prefixed with "!" are excluded. Order is important.
        /// </summary>
        public InputList<string> FilepathPatterns
        {
            get => _filepathPatterns ?? (_filepathPatterns = new InputList<string>());
            set => _filepathPatterns = value;
        }

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

        public RepositoryPolicyFilePathPatternArgs()
        {
        }
        public static new RepositoryPolicyFilePathPatternArgs Empty => new RepositoryPolicyFilePathPatternArgs();
    }

    public sealed class RepositoryPolicyFilePathPatternState : global::Pulumi.ResourceArgs
    {
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

        [Input("filepathPatterns")]
        private InputList<string>? _filepathPatterns;

        /// <summary>
        /// Block pushes from introducing file paths that match the following patterns. Exact paths begin with "/". You can specify exact paths and wildcards. You can also specify multiple paths using ";" as a separator. Paths prefixed with "!" are excluded. Order is important.
        /// </summary>
        public InputList<string> FilepathPatterns
        {
            get => _filepathPatterns ?? (_filepathPatterns = new InputList<string>());
            set => _filepathPatterns = value;
        }

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

        public RepositoryPolicyFilePathPatternState()
        {
        }
        public static new RepositoryPolicyFilePathPatternState Empty => new RepositoryPolicyFilePathPatternState();
    }
}
