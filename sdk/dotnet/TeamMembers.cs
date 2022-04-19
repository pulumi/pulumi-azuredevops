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
    /// Manages members of a team within a project in a Azure DevOps organization.
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
    ///             WorkItemTemplate = "Agile",
    ///             VersionControl = "Git",
    ///             Visibility = "private",
    ///             Description = "Managed by Terraform",
    ///         });
    ///         var example_project_readers = AzureDevOps.GetGroup.Invoke(new AzureDevOps.GetGroupInvokeArgs
    ///         {
    ///             ProjectId = exampleProject.Id,
    ///             Name = "Readers",
    ///         });
    ///         var exampleTeam = new AzureDevOps.Team("exampleTeam", new AzureDevOps.TeamArgs
    ///         {
    ///             ProjectId = exampleProject.Id,
    ///         });
    ///         var example_team_members = new AzureDevOps.TeamMembers("example-team-members", new AzureDevOps.TeamMembersArgs
    ///         {
    ///             ProjectId = exampleTeam.ProjectId,
    ///             TeamId = exampleTeam.Id,
    ///             Mode = "overwrite",
    ///             Members = 
    ///             {
    ///                 example_project_readers.Apply(example_project_readers =&gt; example_project_readers.Descriptor),
    ///             },
    ///         });
    ///     }
    /// 
    /// }
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
    [AzureDevOpsResourceType("azuredevops:index/teamMembers:TeamMembers")]
    public partial class TeamMembers : Pulumi.CustomResource
    {
        /// <summary>
        /// List of subject descriptors to define members of the team.
        /// </summary>
        [Output("members")]
        public Output<ImmutableArray<string>> Members { get; private set; } = null!;

        /// <summary>
        /// The mode how the resource manages team members.
        /// - `mode == add`: the resource will ensure that all specified members will be part of the referenced team
        /// - `mode == overwrite`: the resource will replace all existing members with the members specified within the `members` block
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
        /// Create a TeamMembers resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TeamMembers(string name, TeamMembersArgs args, CustomResourceOptions? options = null)
            : base("azuredevops:index/teamMembers:TeamMembers", name, args ?? new TeamMembersArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TeamMembers(string name, Input<string> id, TeamMembersState? state = null, CustomResourceOptions? options = null)
            : base("azuredevops:index/teamMembers:TeamMembers", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing TeamMembers resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TeamMembers Get(string name, Input<string> id, TeamMembersState? state = null, CustomResourceOptions? options = null)
        {
            return new TeamMembers(name, id, state, options);
        }
    }

    public sealed class TeamMembersArgs : Pulumi.ResourceArgs
    {
        [Input("members", required: true)]
        private InputList<string>? _members;

        /// <summary>
        /// List of subject descriptors to define members of the team.
        /// </summary>
        public InputList<string> Members
        {
            get => _members ?? (_members = new InputList<string>());
            set => _members = value;
        }

        /// <summary>
        /// The mode how the resource manages team members.
        /// - `mode == add`: the resource will ensure that all specified members will be part of the referenced team
        /// - `mode == overwrite`: the resource will replace all existing members with the members specified within the `members` block
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

        public TeamMembersArgs()
        {
        }
    }

    public sealed class TeamMembersState : Pulumi.ResourceArgs
    {
        [Input("members")]
        private InputList<string>? _members;

        /// <summary>
        /// List of subject descriptors to define members of the team.
        /// </summary>
        public InputList<string> Members
        {
            get => _members ?? (_members = new InputList<string>());
            set => _members = value;
        }

        /// <summary>
        /// The mode how the resource manages team members.
        /// - `mode == add`: the resource will ensure that all specified members will be part of the referenced team
        /// - `mode == overwrite`: the resource will replace all existing members with the members specified within the `members` block
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

        public TeamMembersState()
        {
        }
    }
}
