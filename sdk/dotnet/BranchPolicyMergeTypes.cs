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
    /// Branch policy for merge types allowed on a specified branch.
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
    ///     var exampleProject = new AzureDevOps.Project("exampleProject");
    /// 
    ///     var exampleGit = new AzureDevOps.Git("exampleGit", new()
    ///     {
    ///         ProjectId = exampleProject.Id,
    ///         Initialization = new AzureDevOps.Inputs.GitInitializationArgs
    ///         {
    ///             InitType = "Clean",
    ///         },
    ///     });
    /// 
    ///     var exampleBranchPolicyMergeTypes = new AzureDevOps.BranchPolicyMergeTypes("exampleBranchPolicyMergeTypes", new()
    ///     {
    ///         ProjectId = exampleProject.Id,
    ///         Enabled = true,
    ///         Blocking = true,
    ///         Settings = new AzureDevOps.Inputs.BranchPolicyMergeTypesSettingsArgs
    ///         {
    ///             AllowSquash = true,
    ///             AllowRebaseAndFastForward = true,
    ///             AllowBasicNoFastForward = true,
    ///             AllowRebaseWithMerge = true,
    ///             Scopes = new[]
    ///             {
    ///                 new AzureDevOps.Inputs.BranchPolicyMergeTypesSettingsScopeArgs
    ///                 {
    ///                     RepositoryId = exampleGit.Id,
    ///                     RepositoryRef = exampleGit.DefaultBranch,
    ///                     MatchType = "Exact",
    ///                 },
    ///                 new AzureDevOps.Inputs.BranchPolicyMergeTypesSettingsScopeArgs
    ///                 {
    ///                     RepositoryId = null,
    ///                     RepositoryRef = "refs/heads/releases",
    ///                     MatchType = "Prefix",
    ///                 },
    ///                 new AzureDevOps.Inputs.BranchPolicyMergeTypesSettingsScopeArgs
    ///                 {
    ///                     MatchType = "DefaultBranch",
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ## Relevant Links
    /// 
    /// - [Azure DevOps Service REST API 7.0 - Policy Configurations](https://docs.microsoft.com/en-us/rest/api/azure/devops/policy/configurations/create?view=azure-devops-rest-7.0)
    /// 
    /// ## Import
    /// 
    /// Azure DevOps Branch Policies can be imported using the project ID and policy configuration ID:
    /// 
    /// ```sh
    /// $ pulumi import azuredevops:index/branchPolicyMergeTypes:BranchPolicyMergeTypes example 00000000-0000-0000-0000-000000000000/0
    /// ```
    /// </summary>
    [AzureDevOpsResourceType("azuredevops:index/branchPolicyMergeTypes:BranchPolicyMergeTypes")]
    public partial class BranchPolicyMergeTypes : global::Pulumi.CustomResource
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
        public Output<Outputs.BranchPolicyMergeTypesSettings> Settings { get; private set; } = null!;


        /// <summary>
        /// Create a BranchPolicyMergeTypes resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BranchPolicyMergeTypes(string name, BranchPolicyMergeTypesArgs args, CustomResourceOptions? options = null)
            : base("azuredevops:index/branchPolicyMergeTypes:BranchPolicyMergeTypes", name, args ?? new BranchPolicyMergeTypesArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BranchPolicyMergeTypes(string name, Input<string> id, BranchPolicyMergeTypesState? state = null, CustomResourceOptions? options = null)
            : base("azuredevops:index/branchPolicyMergeTypes:BranchPolicyMergeTypes", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing BranchPolicyMergeTypes resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BranchPolicyMergeTypes Get(string name, Input<string> id, BranchPolicyMergeTypesState? state = null, CustomResourceOptions? options = null)
        {
            return new BranchPolicyMergeTypes(name, id, state, options);
        }
    }

    public sealed class BranchPolicyMergeTypesArgs : global::Pulumi.ResourceArgs
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
        public Input<Inputs.BranchPolicyMergeTypesSettingsArgs> Settings { get; set; } = null!;

        public BranchPolicyMergeTypesArgs()
        {
        }
        public static new BranchPolicyMergeTypesArgs Empty => new BranchPolicyMergeTypesArgs();
    }

    public sealed class BranchPolicyMergeTypesState : global::Pulumi.ResourceArgs
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
        public Input<Inputs.BranchPolicyMergeTypesSettingsGetArgs>? Settings { get; set; }

        public BranchPolicyMergeTypesState()
        {
        }
        public static new BranchPolicyMergeTypesState Empty => new BranchPolicyMergeTypesState();
    }
}
