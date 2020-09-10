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
    ///             Description = "Test Project Description",
    ///             Visibility = "private",
    ///             VersionControl = "Git",
    ///             WorkItemTemplate = "Agile",
    ///         });
    ///         var project_readers = project.Id.Apply(id =&gt; AzureDevOps.GetGroup.InvokeAsync(new AzureDevOps.GetGroupArgs
    ///         {
    ///             ProjectId = id,
    ///             Name = "Readers",
    ///         }));
    ///         var project_perm = new AzureDevOps.ProjectPermissions("project-perm", new AzureDevOps.ProjectPermissionsArgs
    ///         {
    ///             ProjectId = project.Id,
    ///             Principal = project_readers.Apply(project_readers =&gt; project_readers.Id),
    ///             Permissions = 
    ///             {
    ///                 { "DELETE", "Deny" },
    ///                 { "EDIT_BUILD_STATUS", "NotSet" },
    ///                 { "WORK_ITEM_MOVE", "Allow" },
    ///                 { "DELETE_TEST_RESULTS", "Deny" },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ## Relevant Links
    /// 
    /// * [Azure DevOps Service REST API 5.1 - Security](https://docs.microsoft.com/en-us/rest/api/azure/devops/security/?view=azure-devops-rest-5.1)
    /// 
    /// ## PAT Permissions Required
    /// 
    /// - **Project &amp; Team**: vso.security_manage - Grants the ability to read, write, and manage security permissions.
    /// </summary>
    public partial class ProjectPermissions : Pulumi.CustomResource
    {
        /// <summary>
        /// the permissions to assign. The following permissions are available
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

    public sealed class ProjectPermissionsArgs : Pulumi.ResourceArgs
    {
        [Input("permissions", required: true)]
        private InputMap<string>? _permissions;

        /// <summary>
        /// the permissions to assign. The following permissions are available
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
    }

    public sealed class ProjectPermissionsState : Pulumi.ResourceArgs
    {
        [Input("permissions")]
        private InputMap<string>? _permissions;

        /// <summary>
        /// the permissions to assign. The following permissions are available
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
    }
}
