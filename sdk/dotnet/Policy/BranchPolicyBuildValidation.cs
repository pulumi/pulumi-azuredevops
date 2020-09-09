// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureDevOps.Policy
{
    /// <summary>
    /// Manages a build validation branch policy within Azure DevOps.
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
    ///             ProjectName = "Sample Project",
    ///         });
    ///         var git = new AzureDevOps.Git("git", new AzureDevOps.GitArgs
    ///         {
    ///             ProjectId = project.Id,
    ///             Initialization = new AzureDevOps.Inputs.GitInitializationArgs
    ///             {
    ///                 InitType = "Clean",
    ///             },
    ///         });
    ///         var buildDefinition = new AzureDevOps.BuildDefinition("buildDefinition", new AzureDevOps.BuildDefinitionArgs
    ///         {
    ///             ProjectId = project.Id,
    ///             Repository = new AzureDevOps.Inputs.BuildDefinitionRepositoryArgs
    ///             {
    ///                 RepoType = "TfsGit",
    ///                 RepoId = git.Id,
    ///                 YmlPath = "azure-pipelines.yml",
    ///             },
    ///         });
    ///         var branchPolicyBuildValidation = new AzureDevOps.BranchPolicyBuildValidation("branchPolicyBuildValidation", new AzureDevOps.BranchPolicyBuildValidationArgs
    ///         {
    ///             ProjectId = project.Id,
    ///             Enabled = true,
    ///             Blocking = true,
    ///             Settings = new AzureDevOps.Inputs.BranchPolicyBuildValidationSettingsArgs
    ///             {
    ///                 DisplayName = "Don't break the build!",
    ///                 BuildDefinitionId = buildDefinition.Id,
    ///                 ValidDuration = 720,
    ///                 Scopes = 
    ///                 {
    ///                     new AzureDevOps.Inputs.BranchPolicyBuildValidationSettingsScopeArgs
    ///                     {
    ///                         RepositoryId = git.Id,
    ///                         RepositoryRef = git.DefaultBranch,
    ///                         MatchType = "Exact",
    ///                     },
    ///                     new AzureDevOps.Inputs.BranchPolicyBuildValidationSettingsScopeArgs
    ///                     {
    ///                         RepositoryId = git.Id,
    ///                         RepositoryRef = "refs/heads/releases",
    ///                         MatchType = "Prefix",
    ///                     },
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ## Relevant Links
    /// 
    /// * [Azure DevOps Service REST API 5.1 - Policy Configurations](https://docs.microsoft.com/en-us/rest/api/azure/devops/policy/configurations/create?view=azure-devops-rest-5.1)
    /// </summary>
    [Obsolete(@"azuredevops.policy.BranchPolicyBuildValidation has been deprecated in favor of azuredevops.BranchPolicyBuildValidation")]
    public partial class BranchPolicyBuildValidation : Pulumi.CustomResource
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
        /// The ID of the project in which the policy will be created.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// Configuration for the policy. This block must be defined exactly once.
        /// </summary>
        [Output("settings")]
        public Output<Outputs.BranchPolicyBuildValidationSettings> Settings { get; private set; } = null!;


        /// <summary>
        /// Create a BranchPolicyBuildValidation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BranchPolicyBuildValidation(string name, BranchPolicyBuildValidationArgs args, CustomResourceOptions? options = null)
            : base("azuredevops:Policy/branchPolicyBuildValidation:BranchPolicyBuildValidation", name, args ?? new BranchPolicyBuildValidationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BranchPolicyBuildValidation(string name, Input<string> id, BranchPolicyBuildValidationState? state = null, CustomResourceOptions? options = null)
            : base("azuredevops:Policy/branchPolicyBuildValidation:BranchPolicyBuildValidation", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing BranchPolicyBuildValidation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BranchPolicyBuildValidation Get(string name, Input<string> id, BranchPolicyBuildValidationState? state = null, CustomResourceOptions? options = null)
        {
            return new BranchPolicyBuildValidation(name, id, state, options);
        }
    }

    public sealed class BranchPolicyBuildValidationArgs : Pulumi.ResourceArgs
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
        /// The ID of the project in which the policy will be created.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        /// <summary>
        /// Configuration for the policy. This block must be defined exactly once.
        /// </summary>
        [Input("settings", required: true)]
        public Input<Inputs.BranchPolicyBuildValidationSettingsArgs> Settings { get; set; } = null!;

        public BranchPolicyBuildValidationArgs()
        {
        }
    }

    public sealed class BranchPolicyBuildValidationState : Pulumi.ResourceArgs
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
        /// The ID of the project in which the policy will be created.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// Configuration for the policy. This block must be defined exactly once.
        /// </summary>
        [Input("settings")]
        public Input<Inputs.BranchPolicyBuildValidationSettingsGetArgs>? Settings { get; set; }

        public BranchPolicyBuildValidationState()
        {
        }
    }
}
