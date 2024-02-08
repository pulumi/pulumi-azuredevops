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
    /// Manages Pipeline Settings for Azure DevOps projects
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
    ///     var exampleProject = new AzureDevOps.Project("exampleProject", new()
    ///     {
    ///         Visibility = "private",
    ///         VersionControl = "Git",
    ///         WorkItemTemplate = "Agile",
    ///         Description = "Managed by Terraform",
    ///     });
    /// 
    ///     var exampleProjectPipelineSettings = new AzureDevOps.ProjectPipelineSettings("exampleProjectPipelineSettings", new()
    ///     {
    ///         ProjectId = exampleProject.Id,
    ///         EnforceJobScope = true,
    ///         EnforceReferencedRepoScopedToken = false,
    ///         EnforceSettableVar = true,
    ///         PublishPipelineMetadata = false,
    ///         StatusBadgesArePrivate = true,
    ///     });
    /// 
    /// });
    /// ```
    /// ## Relevant Links
    /// 
    /// No official documentation available
    /// 
    /// ## PAT Permissions Required
    /// 
    /// - Full Access
    /// 
    /// ## Import
    /// 
    /// Azure DevOps feature settings can be imported using the project id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import azuredevops:index/projectPipelineSettings:ProjectPipelineSettings example 00000000-0000-0000-0000-000000000000
    /// ```
    /// </summary>
    [AzureDevOpsResourceType("azuredevops:index/projectPipelineSettings:ProjectPipelineSettings")]
    public partial class ProjectPipelineSettings : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Limit job authorization scope to current project for non-release pipelines.
        /// </summary>
        [Output("enforceJobScope")]
        public Output<bool> EnforceJobScope { get; private set; } = null!;

        /// <summary>
        /// Limit job authorization scope to current project for release pipelines.
        /// 
        /// &gt; **NOTE:**
        /// &gt; The settings at the organization will override settings specified on the project.
        /// &gt; For example, if `enforce_job_scope` is true at the organization, the `azuredevops.ProjectPipelineSettings` resource cannot set it to false.
        /// &gt; In this scenario, the plan will always show that the resource is trying to change `enforce_job_scope` from `true` to `false`.
        /// </summary>
        [Output("enforceJobScopeForRelease")]
        public Output<bool> EnforceJobScopeForRelease { get; private set; } = null!;

        /// <summary>
        /// Protect access to repositories in YAML pipelines.
        /// </summary>
        [Output("enforceReferencedRepoScopedToken")]
        public Output<bool> EnforceReferencedRepoScopedToken { get; private set; } = null!;

        /// <summary>
        /// Limit variables that can be set at queue time.
        /// </summary>
        [Output("enforceSettableVar")]
        public Output<bool> EnforceSettableVar { get; private set; } = null!;

        /// <summary>
        /// The `id` of the project for which the project pipeline settings will be managed.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// Publish metadata from pipelines.
        /// </summary>
        [Output("publishPipelineMetadata")]
        public Output<bool> PublishPipelineMetadata { get; private set; } = null!;

        /// <summary>
        /// Disable anonymous access to badges.
        /// </summary>
        [Output("statusBadgesArePrivate")]
        public Output<bool> StatusBadgesArePrivate { get; private set; } = null!;


        /// <summary>
        /// Create a ProjectPipelineSettings resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ProjectPipelineSettings(string name, ProjectPipelineSettingsArgs args, CustomResourceOptions? options = null)
            : base("azuredevops:index/projectPipelineSettings:ProjectPipelineSettings", name, args ?? new ProjectPipelineSettingsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ProjectPipelineSettings(string name, Input<string> id, ProjectPipelineSettingsState? state = null, CustomResourceOptions? options = null)
            : base("azuredevops:index/projectPipelineSettings:ProjectPipelineSettings", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ProjectPipelineSettings resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ProjectPipelineSettings Get(string name, Input<string> id, ProjectPipelineSettingsState? state = null, CustomResourceOptions? options = null)
        {
            return new ProjectPipelineSettings(name, id, state, options);
        }
    }

    public sealed class ProjectPipelineSettingsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Limit job authorization scope to current project for non-release pipelines.
        /// </summary>
        [Input("enforceJobScope")]
        public Input<bool>? EnforceJobScope { get; set; }

        /// <summary>
        /// Limit job authorization scope to current project for release pipelines.
        /// 
        /// &gt; **NOTE:**
        /// &gt; The settings at the organization will override settings specified on the project.
        /// &gt; For example, if `enforce_job_scope` is true at the organization, the `azuredevops.ProjectPipelineSettings` resource cannot set it to false.
        /// &gt; In this scenario, the plan will always show that the resource is trying to change `enforce_job_scope` from `true` to `false`.
        /// </summary>
        [Input("enforceJobScopeForRelease")]
        public Input<bool>? EnforceJobScopeForRelease { get; set; }

        /// <summary>
        /// Protect access to repositories in YAML pipelines.
        /// </summary>
        [Input("enforceReferencedRepoScopedToken")]
        public Input<bool>? EnforceReferencedRepoScopedToken { get; set; }

        /// <summary>
        /// Limit variables that can be set at queue time.
        /// </summary>
        [Input("enforceSettableVar")]
        public Input<bool>? EnforceSettableVar { get; set; }

        /// <summary>
        /// The `id` of the project for which the project pipeline settings will be managed.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        /// <summary>
        /// Publish metadata from pipelines.
        /// </summary>
        [Input("publishPipelineMetadata")]
        public Input<bool>? PublishPipelineMetadata { get; set; }

        /// <summary>
        /// Disable anonymous access to badges.
        /// </summary>
        [Input("statusBadgesArePrivate")]
        public Input<bool>? StatusBadgesArePrivate { get; set; }

        public ProjectPipelineSettingsArgs()
        {
        }
        public static new ProjectPipelineSettingsArgs Empty => new ProjectPipelineSettingsArgs();
    }

    public sealed class ProjectPipelineSettingsState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Limit job authorization scope to current project for non-release pipelines.
        /// </summary>
        [Input("enforceJobScope")]
        public Input<bool>? EnforceJobScope { get; set; }

        /// <summary>
        /// Limit job authorization scope to current project for release pipelines.
        /// 
        /// &gt; **NOTE:**
        /// &gt; The settings at the organization will override settings specified on the project.
        /// &gt; For example, if `enforce_job_scope` is true at the organization, the `azuredevops.ProjectPipelineSettings` resource cannot set it to false.
        /// &gt; In this scenario, the plan will always show that the resource is trying to change `enforce_job_scope` from `true` to `false`.
        /// </summary>
        [Input("enforceJobScopeForRelease")]
        public Input<bool>? EnforceJobScopeForRelease { get; set; }

        /// <summary>
        /// Protect access to repositories in YAML pipelines.
        /// </summary>
        [Input("enforceReferencedRepoScopedToken")]
        public Input<bool>? EnforceReferencedRepoScopedToken { get; set; }

        /// <summary>
        /// Limit variables that can be set at queue time.
        /// </summary>
        [Input("enforceSettableVar")]
        public Input<bool>? EnforceSettableVar { get; set; }

        /// <summary>
        /// The `id` of the project for which the project pipeline settings will be managed.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// Publish metadata from pipelines.
        /// </summary>
        [Input("publishPipelineMetadata")]
        public Input<bool>? PublishPipelineMetadata { get; set; }

        /// <summary>
        /// Disable anonymous access to badges.
        /// </summary>
        [Input("statusBadgesArePrivate")]
        public Input<bool>? StatusBadgesArePrivate { get; set; }

        public ProjectPipelineSettingsState()
        {
        }
        public static new ProjectPipelineSettingsState Empty => new ProjectPipelineSettingsState();
    }
}
