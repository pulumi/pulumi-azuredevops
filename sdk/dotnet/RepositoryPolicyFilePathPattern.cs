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
    /// Manage a file path pattern repository policy within Azure DevOps project.
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
    ///         var repositoryPolicyFilePathPattern = new AzureDevOps.RepositoryPolicyFilePathPattern("repositoryPolicyFilePathPattern", new AzureDevOps.RepositoryPolicyFilePathPatternArgs
    ///         {
    ///             ProjectId = project.Id,
    ///             Enabled = true,
    ///             Blocking = true,
    ///             Settings = new AzureDevOps.Inputs.RepositoryPolicyFilePathPatternSettingsArgs
    ///             {
    ///                 FilepathPatterns = 
    ///                 {
    ///                     "*.go",
    ///                     "/home/test/*.ts",
    ///                 },
    ///                 Scopes = 
    ///                 {
    ///                     new AzureDevOps.Inputs.RepositoryPolicyFilePathPatternSettingsScopeArgs
    ///                     {
    ///                         RepositoryId = git.Id,
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
    /// - [Azure DevOps Service REST API 5.1 - Policy Configurations](https://docs.microsoft.com/en-us/rest/api/azure/devops/policy/configurations/create?view=azure-devops-rest-5.1)
    /// 
    /// ## Import
    /// 
    /// Azure DevOps repository policies can be imported using the projectID/policyID or projectName/policyID
    /// 
    /// ```sh
    ///  $ pulumi import azuredevops:index/repositoryPolicyFilePathPattern:RepositoryPolicyFilePathPattern p 00000000-0000-0000-0000-000000000000/0
    /// ```
    /// </summary>
    [AzureDevOpsResourceType("azuredevops:index/repositoryPolicyFilePathPattern:RepositoryPolicyFilePathPattern")]
    public partial class RepositoryPolicyFilePathPattern : Pulumi.CustomResource
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
        public Output<Outputs.RepositoryPolicyFilePathPatternSettings> Settings { get; private set; } = null!;


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

    public sealed class RepositoryPolicyFilePathPatternArgs : Pulumi.ResourceArgs
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
        public Input<Inputs.RepositoryPolicyFilePathPatternSettingsArgs> Settings { get; set; } = null!;

        public RepositoryPolicyFilePathPatternArgs()
        {
        }
    }

    public sealed class RepositoryPolicyFilePathPatternState : Pulumi.ResourceArgs
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
        public Input<Inputs.RepositoryPolicyFilePathPatternSettingsGetArgs>? Settings { get; set; }

        public RepositoryPolicyFilePathPatternState()
        {
        }
    }
}
