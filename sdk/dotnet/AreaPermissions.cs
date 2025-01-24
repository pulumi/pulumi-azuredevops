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
    /// Manages permissions for an Area (Component)
    /// 
    /// &gt; **Note** Permissions can be assigned to group principals and not to single user principals.
    /// 
    /// ## Permission levels
    /// 
    /// Permission for Areas within Azure DevOps can be applied on two different levels.
    /// Those levels are reflected by specifying (or omitting) values for the arguments `project_id` and `path`.
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
    ///     var example = new AzureDevOps.Project("example", new()
    ///     {
    ///         Name = "Example Project",
    ///         WorkItemTemplate = "Agile",
    ///         VersionControl = "Git",
    ///         Visibility = "private",
    ///         Description = "Managed by Pulumi",
    ///     });
    /// 
    ///     var example_project_readers = AzureDevOps.GetGroup.Invoke(new()
    ///     {
    ///         ProjectId = example.Id,
    ///         Name = "Readers",
    ///     });
    /// 
    ///     var example_root_permissions = new AzureDevOps.AreaPermissions("example-root-permissions", new()
    ///     {
    ///         ProjectId = example.Id,
    ///         Principal = example_project_readers.Apply(example_project_readers =&gt; example_project_readers.Apply(getGroupResult =&gt; getGroupResult.Id)),
    ///         Path = "/",
    ///         Permissions = 
    ///         {
    ///             { "CREATE_CHILDREN", "Deny" },
    ///             { "GENERIC_READ", "Allow" },
    ///             { "DELETE", "Deny" },
    ///             { "WORK_ITEM_READ", "Allow" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
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
    [AzureDevOpsResourceType("azuredevops:index/areaPermissions:AreaPermissions")]
    public partial class AreaPermissions : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the branch to assign the permissions.
        /// </summary>
        [Output("path")]
        public Output<string?> Path { get; private set; } = null!;

        /// <summary>
        /// the permissions to assign. The following permissions are available.
        /// 
        /// | Permission             | Description                          |
        /// |------------------------|--------------------------------------|
        /// | GENERIC_READ           | View permissions for this node       |
        /// | GENERIC_WRITE          | Edit this node                       |
        /// | CREATE_CHILDREN        | Create child nodes                   |
        /// | DELETE                 | Delete this node                     |
        /// | WORK_ITEM_READ         | View work items in this node         |
        /// | WORK_ITEM_WRITE        | Edit work items in this node         |
        /// | MANAGE_TEST_PLANS      | Manage test plans                    |
        /// | MANAGE_TEST_SUITES     | Manage test suites                   |
        /// | WORK_ITEM_SAVE_COMMENT | Edit work item comments in this node |
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
        /// Create a AreaPermissions resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AreaPermissions(string name, AreaPermissionsArgs args, CustomResourceOptions? options = null)
            : base("azuredevops:index/areaPermissions:AreaPermissions", name, args ?? new AreaPermissionsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AreaPermissions(string name, Input<string> id, AreaPermissionsState? state = null, CustomResourceOptions? options = null)
            : base("azuredevops:index/areaPermissions:AreaPermissions", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AreaPermissions resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AreaPermissions Get(string name, Input<string> id, AreaPermissionsState? state = null, CustomResourceOptions? options = null)
        {
            return new AreaPermissions(name, id, state, options);
        }
    }

    public sealed class AreaPermissionsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the branch to assign the permissions.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        [Input("permissions", required: true)]
        private InputMap<string>? _permissions;

        /// <summary>
        /// the permissions to assign. The following permissions are available.
        /// 
        /// | Permission             | Description                          |
        /// |------------------------|--------------------------------------|
        /// | GENERIC_READ           | View permissions for this node       |
        /// | GENERIC_WRITE          | Edit this node                       |
        /// | CREATE_CHILDREN        | Create child nodes                   |
        /// | DELETE                 | Delete this node                     |
        /// | WORK_ITEM_READ         | View work items in this node         |
        /// | WORK_ITEM_WRITE        | Edit work items in this node         |
        /// | MANAGE_TEST_PLANS      | Manage test plans                    |
        /// | MANAGE_TEST_SUITES     | Manage test suites                   |
        /// | WORK_ITEM_SAVE_COMMENT | Edit work item comments in this node |
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

        public AreaPermissionsArgs()
        {
        }
        public static new AreaPermissionsArgs Empty => new AreaPermissionsArgs();
    }

    public sealed class AreaPermissionsState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the branch to assign the permissions.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        [Input("permissions")]
        private InputMap<string>? _permissions;

        /// <summary>
        /// the permissions to assign. The following permissions are available.
        /// 
        /// | Permission             | Description                          |
        /// |------------------------|--------------------------------------|
        /// | GENERIC_READ           | View permissions for this node       |
        /// | GENERIC_WRITE          | Edit this node                       |
        /// | CREATE_CHILDREN        | Create child nodes                   |
        /// | DELETE                 | Delete this node                     |
        /// | WORK_ITEM_READ         | View work items in this node         |
        /// | WORK_ITEM_WRITE        | Edit work items in this node         |
        /// | MANAGE_TEST_PLANS      | Manage test plans                    |
        /// | MANAGE_TEST_SUITES     | Manage test suites                   |
        /// | WORK_ITEM_SAVE_COMMENT | Edit work item comments in this node |
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

        public AreaPermissionsState()
        {
        }
        public static new AreaPermissionsState Empty => new AreaPermissionsState();
    }
}
