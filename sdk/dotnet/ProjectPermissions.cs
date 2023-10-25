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
    /// Manages permissions for a AzureDevOps project
    /// 
    /// &gt; **Note** Permissions can be assigned to group principals and not to single user principals.
    /// 
    /// ## Relevant Links
    /// 
    /// * [Azure DevOps Service REST API 7.0 - Security](https://docs.microsoft.com/en-us/rest/api/azure/devops/security/?view=azure-devops-rest-7.0)
    /// 
    /// ## PAT Permissions Required
    /// 
    /// - **Project &amp; Team**: vso.security_manage - Grants the ability to read, write, and manage security permissions.
    /// 
    /// ## Import
    /// 
    /// The resource does not support import.
    /// </summary>
    [AzureDevOpsResourceType("azuredevops:index/projectPermissions:ProjectPermissions")]
    public partial class ProjectPermissions : global::Pulumi.CustomResource
    {
        /// <summary>
        /// the permissions to assign. The following permissions are available
        /// 
        /// | Permission                   | Description                                  |
        /// |------------------------------|----------------------------------------------|
        /// | GENERIC_READ                 | View project-level information               |
        /// | GENERIC_WRITE                | Edit project-level information               |
        /// | DELETE                       | Delete team project                          |
        /// | PUBLISH_TEST_RESULTS         | Create test runs                             |
        /// | ADMINISTER_BUILD             | Administer a build                           |
        /// | START_BUILD                  | Start a build                                |
        /// | EDIT_BUILD_STATUS            | Edit build quality                           |
        /// | UPDATE_BUILD                 | Write to build operational store             |
        /// | DELETE_TEST_RESULTS          | Delete test runs                             |
        /// | VIEW_TEST_RESULTS            | View test runs                               |
        /// | MANAGE_TEST_ENVIRONMENTS     | Manage test environments                     |
        /// | MANAGE_TEST_CONFIGURATIONS   | Manage test configurations                   |
        /// | WORK_ITEM_DELETE             | Delete and restore work items                |
        /// | WORK_ITEM_MOVE               | Move work items out of this project          |
        /// | WORK_ITEM_PERMANENTLY_DELETE | Permanently delete work items                |
        /// | RENAME                       | Rename team project                          |
        /// | MANAGE_PROPERTIES            | Manage project properties                    |
        /// | MANAGE_SYSTEM_PROPERTIES     | Manage system project properties             |
        /// | BYPASS_PROPERTY_CACHE        | Bypass project property cache                |
        /// | BYPASS_RULES                 | Bypass rules on work item updates            |
        /// | SUPPRESS_NOTIFICATIONS       | Suppress notifications for work item updates |
        /// | UPDATE_VISIBILITY            | Update project visibility                    |
        /// | CHANGE_PROCESS               | Change process of team project.              |
        /// | AGILETOOLS_BACKLOG           | Agile backlog management.                    |
        /// | AGILETOOLS_PLANS             | Agile plans.                                 |
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
        /// Replace (`true`) or merge (`false`) the permissions. Default: `true`
        /// </summary>
        [Output("replace")]
        public Output<bool?> Replace { get; private set; } = null!;


        /// <summary>
        /// Create a ProjectPermissions resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ProjectPermissions(string name, ProjectPermissionsArgs args, CustomResourceOptions? options = null)
            : base("azuredevops:index/projectPermissions:ProjectPermissions", name, args ?? new ProjectPermissionsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ProjectPermissions(string name, Input<string> id, ProjectPermissionsState? state = null, CustomResourceOptions? options = null)
            : base("azuredevops:index/projectPermissions:ProjectPermissions", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ProjectPermissions resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ProjectPermissions Get(string name, Input<string> id, ProjectPermissionsState? state = null, CustomResourceOptions? options = null)
        {
            return new ProjectPermissions(name, id, state, options);
        }
    }

    public sealed class ProjectPermissionsArgs : global::Pulumi.ResourceArgs
    {
        [Input("permissions", required: true)]
        private InputMap<string>? _permissions;

        /// <summary>
        /// the permissions to assign. The following permissions are available
        /// 
        /// | Permission                   | Description                                  |
        /// |------------------------------|----------------------------------------------|
        /// | GENERIC_READ                 | View project-level information               |
        /// | GENERIC_WRITE                | Edit project-level information               |
        /// | DELETE                       | Delete team project                          |
        /// | PUBLISH_TEST_RESULTS         | Create test runs                             |
        /// | ADMINISTER_BUILD             | Administer a build                           |
        /// | START_BUILD                  | Start a build                                |
        /// | EDIT_BUILD_STATUS            | Edit build quality                           |
        /// | UPDATE_BUILD                 | Write to build operational store             |
        /// | DELETE_TEST_RESULTS          | Delete test runs                             |
        /// | VIEW_TEST_RESULTS            | View test runs                               |
        /// | MANAGE_TEST_ENVIRONMENTS     | Manage test environments                     |
        /// | MANAGE_TEST_CONFIGURATIONS   | Manage test configurations                   |
        /// | WORK_ITEM_DELETE             | Delete and restore work items                |
        /// | WORK_ITEM_MOVE               | Move work items out of this project          |
        /// | WORK_ITEM_PERMANENTLY_DELETE | Permanently delete work items                |
        /// | RENAME                       | Rename team project                          |
        /// | MANAGE_PROPERTIES            | Manage project properties                    |
        /// | MANAGE_SYSTEM_PROPERTIES     | Manage system project properties             |
        /// | BYPASS_PROPERTY_CACHE        | Bypass project property cache                |
        /// | BYPASS_RULES                 | Bypass rules on work item updates            |
        /// | SUPPRESS_NOTIFICATIONS       | Suppress notifications for work item updates |
        /// | UPDATE_VISIBILITY            | Update project visibility                    |
        /// | CHANGE_PROCESS               | Change process of team project.              |
        /// | AGILETOOLS_BACKLOG           | Agile backlog management.                    |
        /// | AGILETOOLS_PLANS             | Agile plans.                                 |
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
        /// Replace (`true`) or merge (`false`) the permissions. Default: `true`
        /// </summary>
        [Input("replace")]
        public Input<bool>? Replace { get; set; }

        public ProjectPermissionsArgs()
        {
        }
        public static new ProjectPermissionsArgs Empty => new ProjectPermissionsArgs();
    }

    public sealed class ProjectPermissionsState : global::Pulumi.ResourceArgs
    {
        [Input("permissions")]
        private InputMap<string>? _permissions;

        /// <summary>
        /// the permissions to assign. The following permissions are available
        /// 
        /// | Permission                   | Description                                  |
        /// |------------------------------|----------------------------------------------|
        /// | GENERIC_READ                 | View project-level information               |
        /// | GENERIC_WRITE                | Edit project-level information               |
        /// | DELETE                       | Delete team project                          |
        /// | PUBLISH_TEST_RESULTS         | Create test runs                             |
        /// | ADMINISTER_BUILD             | Administer a build                           |
        /// | START_BUILD                  | Start a build                                |
        /// | EDIT_BUILD_STATUS            | Edit build quality                           |
        /// | UPDATE_BUILD                 | Write to build operational store             |
        /// | DELETE_TEST_RESULTS          | Delete test runs                             |
        /// | VIEW_TEST_RESULTS            | View test runs                               |
        /// | MANAGE_TEST_ENVIRONMENTS     | Manage test environments                     |
        /// | MANAGE_TEST_CONFIGURATIONS   | Manage test configurations                   |
        /// | WORK_ITEM_DELETE             | Delete and restore work items                |
        /// | WORK_ITEM_MOVE               | Move work items out of this project          |
        /// | WORK_ITEM_PERMANENTLY_DELETE | Permanently delete work items                |
        /// | RENAME                       | Rename team project                          |
        /// | MANAGE_PROPERTIES            | Manage project properties                    |
        /// | MANAGE_SYSTEM_PROPERTIES     | Manage system project properties             |
        /// | BYPASS_PROPERTY_CACHE        | Bypass project property cache                |
        /// | BYPASS_RULES                 | Bypass rules on work item updates            |
        /// | SUPPRESS_NOTIFICATIONS       | Suppress notifications for work item updates |
        /// | UPDATE_VISIBILITY            | Update project visibility                    |
        /// | CHANGE_PROCESS               | Change process of team project.              |
        /// | AGILETOOLS_BACKLOG           | Agile backlog management.                    |
        /// | AGILETOOLS_PLANS             | Agile plans.                                 |
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
        /// Replace (`true`) or merge (`false`) the permissions. Default: `true`
        /// </summary>
        [Input("replace")]
        public Input<bool>? Replace { get; set; }

        public ProjectPermissionsState()
        {
        }
        public static new ProjectPermissionsState Empty => new ProjectPermissionsState();
    }
}
