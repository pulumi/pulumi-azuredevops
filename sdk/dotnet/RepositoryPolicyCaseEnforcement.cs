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
    /// Manages a case enforcement repository policy within Azure DevOps project.
    /// 
    /// &gt; If both project and project policy are enabled, the project policy has high priority.
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
    ///         var project = new AzureDevOps.Project("project", new AzureDevOps.ProjectArgs
    ///         {
    ///             Description = "Managed by Terraform",
    ///             Visibility = "private",
    ///             VersionControl = "Git",
    ///             WorkItemTemplate = "Agile",
    ///         });
    ///         var git = new AzureDevOps.Git("git", new AzureDevOps.GitArgs
    ///         {
    ///             ProjectId = project.Id,
    ///             Initialization = new AzureDevOps.Inputs.GitInitializationArgs
    ///             {
    ///                 InitType = "Clean",
    ///             },
    ///         });
    ///         var repositoryPolicyCaseEnforcement = new AzureDevOps.RepositoryPolicyCaseEnforcement("repositoryPolicyCaseEnforcement", new AzureDevOps.RepositoryPolicyCaseEnforcementArgs
    ///         {
    ///             ProjectId = project.Id,
    ///             Enabled = true,
    ///             Blocking = true,
    ///             EnforceConsistentCase = true,
    ///             RepositoryIds = 
    ///             {
    ///                 git.Id,
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// # Set project level repository policy
    /// ```csharp
    /// using Pulumi;
    /// using AzureDevOps = Pulumi.AzureDevOps;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var repositoryPolicyCaseEnforcement = new AzureDevOps.RepositoryPolicyCaseEnforcement("repositoryPolicyCaseEnforcement", new AzureDevOps.RepositoryPolicyCaseEnforcementArgs
    ///         {
    ///             ProjectId = azuredevops_project.P.Id,
    ///             Enabled = true,
    ///             Blocking = true,
    ///             EnforceConsistentCase = true,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ## Relevant Links
    /// 
    /// - [Azure DevOps Service REST API 5.1 - Policy Configurations](https://docs.microsoft.com/en-us/rest/api/azure/devops/policy/configurations/create?view=azure-devops-rest-5.1)
    /// 
    /// ## Import
    /// 
    /// Azure DevOps repository policies can be imported using the projectID/policyID or projectName/policyID
    /// 
    /// ```sh
    ///  $ pulumi import azuredevops:index/repositoryPolicyCaseEnforcement:RepositoryPolicyCaseEnforcement p 00000000-0000-0000-0000-000000000000/0
    /// ```
    /// </summary>
    [AzureDevOpsResourceType("azuredevops:index/repositoryPolicyCaseEnforcement:RepositoryPolicyCaseEnforcement")]
    public partial class RepositoryPolicyCaseEnforcement : Pulumi.CustomResource
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
        /// Avoid case-sensitivity conflicts by blocking pushes that change name casing on files, folders, branches, and tags.
        /// </summary>
        [Output("enforceConsistentCase")]
        public Output<bool> EnforceConsistentCase { get; private set; } = null!;

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
        /// Create a RepositoryPolicyCaseEnforcement resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RepositoryPolicyCaseEnforcement(string name, RepositoryPolicyCaseEnforcementArgs args, CustomResourceOptions? options = null)
            : base("azuredevops:index/repositoryPolicyCaseEnforcement:RepositoryPolicyCaseEnforcement", name, args ?? new RepositoryPolicyCaseEnforcementArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RepositoryPolicyCaseEnforcement(string name, Input<string> id, RepositoryPolicyCaseEnforcementState? state = null, CustomResourceOptions? options = null)
            : base("azuredevops:index/repositoryPolicyCaseEnforcement:RepositoryPolicyCaseEnforcement", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RepositoryPolicyCaseEnforcement resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RepositoryPolicyCaseEnforcement Get(string name, Input<string> id, RepositoryPolicyCaseEnforcementState? state = null, CustomResourceOptions? options = null)
        {
            return new RepositoryPolicyCaseEnforcement(name, id, state, options);
        }
    }

    public sealed class RepositoryPolicyCaseEnforcementArgs : Pulumi.ResourceArgs
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
        /// Avoid case-sensitivity conflicts by blocking pushes that change name casing on files, folders, branches, and tags.
        /// </summary>
        [Input("enforceConsistentCase", required: true)]
        public Input<bool> EnforceConsistentCase { get; set; } = null!;

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

        public RepositoryPolicyCaseEnforcementArgs()
        {
        }
    }

    public sealed class RepositoryPolicyCaseEnforcementState : Pulumi.ResourceArgs
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
        /// Avoid case-sensitivity conflicts by blocking pushes that change name casing on files, folders, branches, and tags.
        /// </summary>
        [Input("enforceConsistentCase")]
        public Input<bool>? EnforceConsistentCase { get; set; }

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

        public RepositoryPolicyCaseEnforcementState()
        {
        }
    }
}