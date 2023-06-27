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
    /// Manages permissions for a Build Definition
    /// 
    /// &gt; **Note** Permissions can be assigned to group principals and not to single user principals.
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
    ///     var example_readers = AzureDevOps.GetGroup.Invoke(new()
    ///     {
    ///         ProjectId = exampleProject.Id,
    ///         Name = "Readers",
    ///     });
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
    ///     var exampleBuildDefinition = new AzureDevOps.BuildDefinition("exampleBuildDefinition", new()
    ///     {
    ///         ProjectId = exampleProject.Id,
    ///         Path = "\\ExampleFolder",
    ///         CiTrigger = new AzureDevOps.Inputs.BuildDefinitionCiTriggerArgs
    ///         {
    ///             UseYaml = true,
    ///         },
    ///         Repository = new AzureDevOps.Inputs.BuildDefinitionRepositoryArgs
    ///         {
    ///             RepoType = "TfsGit",
    ///             RepoId = exampleGit.Id,
    ///             BranchName = exampleGit.DefaultBranch,
    ///             YmlPath = "azure-pipelines.yml",
    ///         },
    ///     });
    /// 
    ///     var exampleBuildDefinitionPermissions = new AzureDevOps.BuildDefinitionPermissions("exampleBuildDefinitionPermissions", new()
    ///     {
    ///         ProjectId = exampleProject.Id,
    ///         Principal = example_readers.Apply(example_readers =&gt; example_readers.Apply(getGroupResult =&gt; getGroupResult.Id)),
    ///         BuildDefinitionId = exampleBuildDefinition.Id,
    ///         Permissions = 
    ///         {
    ///             { "ViewBuilds", "Allow" },
    ///             { "EditBuildQuality", "Deny" },
    ///             { "DeleteBuilds", "Deny" },
    ///             { "StopBuilds", "Allow" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ## Relevant Links
    /// 
    /// * [Azure DevOps Service REST API 6.0 - Security](https://docs.microsoft.com/en-us/rest/api/azure/devops/security/?view=azure-devops-rest-6.0)
    /// 
    /// ## PAT Permissions Required
    /// 
    /// - **Project &amp; Team**: vso.security_manage - Grants the ability to read, write, and manage security permissions.
    /// 
    /// ## Import
    /// 
    /// The resource does not support import.
    /// </summary>
    [AzureDevOpsResourceType("azuredevops:index/buildDefinitionPermissions:BuildDefinitionPermissions")]
    public partial class BuildDefinitionPermissions : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The id of the build definition to assign the permissions.
        /// </summary>
        [Output("buildDefinitionId")]
        public Output<string> BuildDefinitionId { get; private set; } = null!;

        /// <summary>
        /// the permissions to assign. The following permissions are available.
        /// 
        /// | Permission                     | Description                           |
        /// |--------------------------------|---------------------------------------|
        /// | ViewBuilds                     | View builds                           |
        /// | EditBuildQuality               | Edit build quality                    |
        /// | RetainIndefinitely             | Retain indefinitely                   |
        /// | DeleteBuilds                   | Delete builds                         |
        /// | ManageBuildQualities           | Manage build qualities                |
        /// | DestroyBuilds                  | Destroy builds                        |
        /// | UpdateBuildInformation         | Update build information              |
        /// | QueueBuilds                    | Queue builds                          |
        /// | ManageBuildQueue               | Manage build queue                    |
        /// | StopBuilds                     | Stop builds                           |
        /// | ViewBuildDefinition            | View build pipeline                   |
        /// | EditBuildDefinition            | Edit build pipeline                   |
        /// | DeleteBuildDefinition          | Delete build pipeline                 |
        /// | OverrideBuildCheckInValidation | Override check-in validation by build |
        /// | AdministerBuildPermissions     | Administer build permissions          |
        /// </summary>
        [Output("permissions")]
        public Output<ImmutableDictionary<string, string>> Permissions { get; private set; } = null!;

        /// <summary>
        /// The **group** principal to assign the permissions.
        /// </summary>
        [Output("principal")]
        public Output<string> Principal { get; private set; } = null!;

        /// <summary>
        /// The ID of the project to assign the permissions.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// Replace (`true`) or merge (`false`) the permissions. Default: `true`.
        /// </summary>
        [Output("replace")]
        public Output<bool?> Replace { get; private set; } = null!;


        /// <summary>
        /// Create a BuildDefinitionPermissions resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BuildDefinitionPermissions(string name, BuildDefinitionPermissionsArgs args, CustomResourceOptions? options = null)
            : base("azuredevops:index/buildDefinitionPermissions:BuildDefinitionPermissions", name, args ?? new BuildDefinitionPermissionsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BuildDefinitionPermissions(string name, Input<string> id, BuildDefinitionPermissionsState? state = null, CustomResourceOptions? options = null)
            : base("azuredevops:index/buildDefinitionPermissions:BuildDefinitionPermissions", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing BuildDefinitionPermissions resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BuildDefinitionPermissions Get(string name, Input<string> id, BuildDefinitionPermissionsState? state = null, CustomResourceOptions? options = null)
        {
            return new BuildDefinitionPermissions(name, id, state, options);
        }
    }

    public sealed class BuildDefinitionPermissionsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The id of the build definition to assign the permissions.
        /// </summary>
        [Input("buildDefinitionId", required: true)]
        public Input<string> BuildDefinitionId { get; set; } = null!;

        [Input("permissions", required: true)]
        private InputMap<string>? _permissions;

        /// <summary>
        /// the permissions to assign. The following permissions are available.
        /// 
        /// | Permission                     | Description                           |
        /// |--------------------------------|---------------------------------------|
        /// | ViewBuilds                     | View builds                           |
        /// | EditBuildQuality               | Edit build quality                    |
        /// | RetainIndefinitely             | Retain indefinitely                   |
        /// | DeleteBuilds                   | Delete builds                         |
        /// | ManageBuildQualities           | Manage build qualities                |
        /// | DestroyBuilds                  | Destroy builds                        |
        /// | UpdateBuildInformation         | Update build information              |
        /// | QueueBuilds                    | Queue builds                          |
        /// | ManageBuildQueue               | Manage build queue                    |
        /// | StopBuilds                     | Stop builds                           |
        /// | ViewBuildDefinition            | View build pipeline                   |
        /// | EditBuildDefinition            | Edit build pipeline                   |
        /// | DeleteBuildDefinition          | Delete build pipeline                 |
        /// | OverrideBuildCheckInValidation | Override check-in validation by build |
        /// | AdministerBuildPermissions     | Administer build permissions          |
        /// </summary>
        public InputMap<string> Permissions
        {
            get => _permissions ?? (_permissions = new InputMap<string>());
            set => _permissions = value;
        }

        /// <summary>
        /// The **group** principal to assign the permissions.
        /// </summary>
        [Input("principal", required: true)]
        public Input<string> Principal { get; set; } = null!;

        /// <summary>
        /// The ID of the project to assign the permissions.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        /// <summary>
        /// Replace (`true`) or merge (`false`) the permissions. Default: `true`.
        /// </summary>
        [Input("replace")]
        public Input<bool>? Replace { get; set; }

        public BuildDefinitionPermissionsArgs()
        {
        }
        public static new BuildDefinitionPermissionsArgs Empty => new BuildDefinitionPermissionsArgs();
    }

    public sealed class BuildDefinitionPermissionsState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The id of the build definition to assign the permissions.
        /// </summary>
        [Input("buildDefinitionId")]
        public Input<string>? BuildDefinitionId { get; set; }

        [Input("permissions")]
        private InputMap<string>? _permissions;

        /// <summary>
        /// the permissions to assign. The following permissions are available.
        /// 
        /// | Permission                     | Description                           |
        /// |--------------------------------|---------------------------------------|
        /// | ViewBuilds                     | View builds                           |
        /// | EditBuildQuality               | Edit build quality                    |
        /// | RetainIndefinitely             | Retain indefinitely                   |
        /// | DeleteBuilds                   | Delete builds                         |
        /// | ManageBuildQualities           | Manage build qualities                |
        /// | DestroyBuilds                  | Destroy builds                        |
        /// | UpdateBuildInformation         | Update build information              |
        /// | QueueBuilds                    | Queue builds                          |
        /// | ManageBuildQueue               | Manage build queue                    |
        /// | StopBuilds                     | Stop builds                           |
        /// | ViewBuildDefinition            | View build pipeline                   |
        /// | EditBuildDefinition            | Edit build pipeline                   |
        /// | DeleteBuildDefinition          | Delete build pipeline                 |
        /// | OverrideBuildCheckInValidation | Override check-in validation by build |
        /// | AdministerBuildPermissions     | Administer build permissions          |
        /// </summary>
        public InputMap<string> Permissions
        {
            get => _permissions ?? (_permissions = new InputMap<string>());
            set => _permissions = value;
        }

        /// <summary>
        /// The **group** principal to assign the permissions.
        /// </summary>
        [Input("principal")]
        public Input<string>? Principal { get; set; }

        /// <summary>
        /// The ID of the project to assign the permissions.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// Replace (`true`) or merge (`false`) the permissions. Default: `true`.
        /// </summary>
        [Input("replace")]
        public Input<bool>? Replace { get; set; }

        public BuildDefinitionPermissionsState()
        {
        }
        public static new BuildDefinitionPermissionsState Empty => new BuildDefinitionPermissionsState();
    }
}
