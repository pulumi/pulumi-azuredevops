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
    /// Manages a team within a project in a Azure DevOps organization.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
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
    ///     var example_project_readers = AzureDevOps.GetGroup.Invoke(new()
    ///     {
    ///         ProjectId = exampleProject.Id,
    ///         Name = "Readers",
    ///     });
    /// 
    ///     var exampleTeam = new AzureDevOps.Team("exampleTeam", new()
    ///     {
    ///         ProjectId = exampleProject.Id,
    ///         Administrators = new[]
    ///         {
    ///             example_project_contributors.Apply(example_project_contributors =&gt; example_project_contributors.Apply(getGroupResult =&gt; getGroupResult.Descriptor)),
    ///         },
    ///         Members = new[]
    ///         {
    ///             example_project_readers.Apply(example_project_readers =&gt; example_project_readers.Apply(getGroupResult =&gt; getGroupResult.Descriptor)),
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ## Relevant Links
    /// 
    /// - [Azure DevOps Service REST API 6.0 - Teams - Create](https://docs.microsoft.com/en-us/rest/api/azure/devops/core/teams/create?view=azure-devops-rest-6.0)
    /// 
    /// ## PAT Permissions Required
    /// 
    /// - **vso.project_manage**:	Grants the ability to create, read, update, and delete projects and teams.
    /// 
    /// ## Import
    /// 
    /// Azure DevOps teams can be imported using the complete resource id `&lt;project_id&gt;/&lt;team_id&gt;` e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azuredevops:index/team:Team example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
    /// ```
    /// </summary>
    [AzureDevOpsResourceType("azuredevops:index/team:Team")]
    public partial class Team : global::Pulumi.CustomResource
    {
        /// <summary>
        /// List of subject descriptors to define administrators of the team.
        /// </summary>
        [Output("administrators")]
        public Output<ImmutableArray<string>> Administrators { get; private set; } = null!;

        /// <summary>
        /// The description of the Team.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The descriptor of the Team.
        /// </summary>
        [Output("descriptor")]
        public Output<string> Descriptor { get; private set; } = null!;

        /// <summary>
        /// List of subject descriptors to define members of the team.
        /// </summary>
        [Output("members")]
        public Output<ImmutableArray<string>> Members { get; private set; } = null!;

        /// <summary>
        /// The name of the Team.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The Project ID.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;


        /// <summary>
        /// Create a Team resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Team(string name, TeamArgs args, CustomResourceOptions? options = null)
            : base("azuredevops:index/team:Team", name, args ?? new TeamArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Team(string name, Input<string> id, TeamState? state = null, CustomResourceOptions? options = null)
            : base("azuredevops:index/team:Team", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Team resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Team Get(string name, Input<string> id, TeamState? state = null, CustomResourceOptions? options = null)
        {
            return new Team(name, id, state, options);
        }
    }

    public sealed class TeamArgs : global::Pulumi.ResourceArgs
    {
        [Input("administrators")]
        private InputList<string>? _administrators;

        /// <summary>
        /// List of subject descriptors to define administrators of the team.
        /// </summary>
        public InputList<string> Administrators
        {
            get => _administrators ?? (_administrators = new InputList<string>());
            set => _administrators = value;
        }

        /// <summary>
        /// The description of the Team.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

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
        /// The name of the Team.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The Project ID.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        public TeamArgs()
        {
        }
        public static new TeamArgs Empty => new TeamArgs();
    }

    public sealed class TeamState : global::Pulumi.ResourceArgs
    {
        [Input("administrators")]
        private InputList<string>? _administrators;

        /// <summary>
        /// List of subject descriptors to define administrators of the team.
        /// </summary>
        public InputList<string> Administrators
        {
            get => _administrators ?? (_administrators = new InputList<string>());
            set => _administrators = value;
        }

        /// <summary>
        /// The description of the Team.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The descriptor of the Team.
        /// </summary>
        [Input("descriptor")]
        public Input<string>? Descriptor { get; set; }

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
        /// The name of the Team.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The Project ID.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        public TeamState()
        {
        }
        public static new TeamState Empty => new TeamState();
    }
}
