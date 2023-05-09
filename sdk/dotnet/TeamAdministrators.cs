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
    /// Manages administrators of a team within a project in a Azure DevOps organization.
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
    ///         WorkItemTemplate = "Agile",
    ///         VersionControl = "Git",
    ///         Visibility = "private",
    ///         Description = "Managed by Terraform",
    ///     });
    /// 
    ///     var example_project_contributors = AzureDevOps.GetGroup.Invoke(new()
    ///     {
    ///         ProjectId = exampleProject.Id,
    ///         Name = "Contributors",
    ///     });
    /// 
    ///     var exampleTeam = new AzureDevOps.Team("exampleTeam", new()
    ///     {
    ///         ProjectId = exampleProject.Id,
    ///     });
    /// 
    ///     var example_team_administrators = new AzureDevOps.TeamAdministrators("example-team-administrators", new()
    ///     {
    ///         ProjectId = exampleTeam.ProjectId,
    ///         TeamId = exampleTeam.Id,
    ///         Mode = "overwrite",
    ///         Administrators = new[]
    ///         {
    ///             example_project_contributors.Apply(example_project_contributors =&gt; example_project_contributors.Apply(getGroupResult =&gt; getGroupResult.Descriptor)),
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ## Relevant Links
    /// 
    /// - [Azure DevOps Service REST API 6.0 - Teams - Update](https://docs.microsoft.com/en-us/rest/api/azure/devops/core/teams/update?view=azure-devops-rest-6.0)
    /// 
    /// ## PAT Permissions Required
    /// 
    /// - **vso.project_write**:	Grants the ability to read and update projects and teams.
    /// 
    /// ## Import
    /// 
    /// The resource does not support import.
    /// </summary>
    [AzureDevOpsResourceType("azuredevops:index/teamAdministrators:TeamAdministrators")]
    public partial class TeamAdministrators : global::Pulumi.CustomResource
    {
        /// <summary>
        /// List of subject descriptors to define adminitrators of the team.
        /// </summary>
        [Output("administrators")]
        public Output<ImmutableArray<string>> Administrators { get; private set; } = null!;

        /// <summary>
        /// The mode how the resource manages team administrators.
        /// - `mode == add`: the resource will ensure that all specified administrators will be part of the referenced team
        /// - `mode == overwrite`: the resource will replace all existing administrators with the administrators specified within the `administrators` block
        /// </summary>
        [Output("mode")]
        public Output<string?> Mode { get; private set; } = null!;

        /// <summary>
        /// The Project ID.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// The ID of the Team.
        /// </summary>
        [Output("teamId")]
        public Output<string> TeamId { get; private set; } = null!;


        /// <summary>
        /// Create a TeamAdministrators resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TeamAdministrators(string name, TeamAdministratorsArgs args, CustomResourceOptions? options = null)
            : base("azuredevops:index/teamAdministrators:TeamAdministrators", name, args ?? new TeamAdministratorsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TeamAdministrators(string name, Input<string> id, TeamAdministratorsState? state = null, CustomResourceOptions? options = null)
            : base("azuredevops:index/teamAdministrators:TeamAdministrators", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing TeamAdministrators resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TeamAdministrators Get(string name, Input<string> id, TeamAdministratorsState? state = null, CustomResourceOptions? options = null)
        {
            return new TeamAdministrators(name, id, state, options);
        }
    }

    public sealed class TeamAdministratorsArgs : global::Pulumi.ResourceArgs
    {
        [Input("administrators", required: true)]
        private InputList<string>? _administrators;

        /// <summary>
        /// List of subject descriptors to define adminitrators of the team.
        /// </summary>
        public InputList<string> Administrators
        {
            get => _administrators ?? (_administrators = new InputList<string>());
            set => _administrators = value;
        }

        /// <summary>
        /// The mode how the resource manages team administrators.
        /// - `mode == add`: the resource will ensure that all specified administrators will be part of the referenced team
        /// - `mode == overwrite`: the resource will replace all existing administrators with the administrators specified within the `administrators` block
        /// </summary>
        [Input("mode")]
        public Input<string>? Mode { get; set; }

        /// <summary>
        /// The Project ID.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        /// <summary>
        /// The ID of the Team.
        /// </summary>
        [Input("teamId", required: true)]
        public Input<string> TeamId { get; set; } = null!;

        public TeamAdministratorsArgs()
        {
        }
        public static new TeamAdministratorsArgs Empty => new TeamAdministratorsArgs();
    }

    public sealed class TeamAdministratorsState : global::Pulumi.ResourceArgs
    {
        [Input("administrators")]
        private InputList<string>? _administrators;

        /// <summary>
        /// List of subject descriptors to define adminitrators of the team.
        /// </summary>
        public InputList<string> Administrators
        {
            get => _administrators ?? (_administrators = new InputList<string>());
            set => _administrators = value;
        }

        /// <summary>
        /// The mode how the resource manages team administrators.
        /// - `mode == add`: the resource will ensure that all specified administrators will be part of the referenced team
        /// - `mode == overwrite`: the resource will replace all existing administrators with the administrators specified within the `administrators` block
        /// </summary>
        [Input("mode")]
        public Input<string>? Mode { get; set; }

        /// <summary>
        /// The Project ID.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The ID of the Team.
        /// </summary>
        [Input("teamId")]
        public Input<string>? TeamId { get; set; }

        public TeamAdministratorsState()
        {
        }
        public static new TeamAdministratorsState Empty => new TeamAdministratorsState();
    }
}
