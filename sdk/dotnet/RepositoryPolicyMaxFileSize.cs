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
    /// Manage a max file size repository policy within Azure DevOps project.
    /// 
    /// &gt; If both project and project policy are enabled, the repository policy has high priority.
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
    ///     var exampleRepositoryPolicyMaxFileSize = new AzureDevOps.RepositoryPolicyMaxFileSize("example", new()
    ///     {
    ///         ProjectId = example.Id,
    ///         Enabled = true,
    ///         Blocking = true,
    ///         MaxFileSize = 1,
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
    ///     var exampleRepositoryPolicyMaxFileSize = new AzureDevOps.RepositoryPolicyMaxFileSize("example", new()
    ///     {
    ///         ProjectId = example.Id,
    ///         Enabled = true,
    ///         Blocking = true,
    ///         MaxFileSize = 1,
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
    /// $ pulumi import azuredevops:index/repositoryPolicyMaxFileSize:RepositoryPolicyMaxFileSize example 00000000-0000-0000-0000-000000000000/0
    /// ```
    /// </summary>
    [AzureDevOpsResourceType("azuredevops:index/repositoryPolicyMaxFileSize:RepositoryPolicyMaxFileSize")]
    public partial class RepositoryPolicyMaxFileSize : global::Pulumi.CustomResource
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
        /// Block pushes that contain new or updated files larger than this limit. Possible values are: `1, 2, 5, 10, 50, 100, 200` (MB).
        /// </summary>
        [Output("maxFileSize")]
        public Output<int> MaxFileSize { get; private set; } = null!;

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
        /// Create a RepositoryPolicyMaxFileSize resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RepositoryPolicyMaxFileSize(string name, RepositoryPolicyMaxFileSizeArgs args, CustomResourceOptions? options = null)
            : base("azuredevops:index/repositoryPolicyMaxFileSize:RepositoryPolicyMaxFileSize", name, args ?? new RepositoryPolicyMaxFileSizeArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RepositoryPolicyMaxFileSize(string name, Input<string> id, RepositoryPolicyMaxFileSizeState? state = null, CustomResourceOptions? options = null)
            : base("azuredevops:index/repositoryPolicyMaxFileSize:RepositoryPolicyMaxFileSize", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RepositoryPolicyMaxFileSize resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RepositoryPolicyMaxFileSize Get(string name, Input<string> id, RepositoryPolicyMaxFileSizeState? state = null, CustomResourceOptions? options = null)
        {
            return new RepositoryPolicyMaxFileSize(name, id, state, options);
        }
    }

    public sealed class RepositoryPolicyMaxFileSizeArgs : global::Pulumi.ResourceArgs
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

        /// <summary>
        /// Block pushes that contain new or updated files larger than this limit. Possible values are: `1, 2, 5, 10, 50, 100, 200` (MB).
        /// </summary>
        [Input("maxFileSize", required: true)]
        public Input<int> MaxFileSize { get; set; } = null!;

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

        public RepositoryPolicyMaxFileSizeArgs()
        {
        }
        public static new RepositoryPolicyMaxFileSizeArgs Empty => new RepositoryPolicyMaxFileSizeArgs();
    }

    public sealed class RepositoryPolicyMaxFileSizeState : global::Pulumi.ResourceArgs
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

        /// <summary>
        /// Block pushes that contain new or updated files larger than this limit. Possible values are: `1, 2, 5, 10, 50, 100, 200` (MB).
        /// </summary>
        [Input("maxFileSize")]
        public Input<int>? MaxFileSize { get; set; }

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

        public RepositoryPolicyMaxFileSizeState()
        {
        }
        public static new RepositoryPolicyMaxFileSizeState Empty => new RepositoryPolicyMaxFileSizeState();
    }
}
