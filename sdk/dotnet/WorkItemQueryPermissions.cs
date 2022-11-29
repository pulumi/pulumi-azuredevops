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
    /// Manages permissions for Work Item Queries.
    /// 
    /// &gt; **Note** Permissions can be assigned to group principals and not to single user principals.
    /// 
    /// ## Permission levels
    /// 
    /// Permission for Work Item Queries within Azure DevOps can be applied on two different levels.
    /// Those levels are reflected by specifying (or omitting) values for the arguments `project_id` and `path`.
    /// 
    /// ### Project level
    /// 
    /// Permissions for all Work Item Queries inside a project (existing or newly created ones) are specified, if only the argument `project_id` has a value.
    /// 
    /// #### Example usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using AzureDevOps = Pulumi.AzureDevOps;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new AzureDevOps.Project("example", new()
    ///     {
    ///         WorkItemTemplate = "Agile",
    ///         VersionControl = "Git",
    ///         Visibility = "private",
    ///         Description = "Managed by Terraform",
    ///     });
    /// 
    ///     var example_readers = AzureDevOps.GetGroup.Invoke(new()
    ///     {
    ///         ProjectId = example.Id,
    ///         Name = "Readers",
    ///     });
    /// 
    ///     var project_wiq_root_permissions = new AzureDevOps.WorkItemQueryPermissions("project-wiq-root-permissions", new()
    ///     {
    ///         ProjectId = example.Id,
    ///         Principal = example_readers.Apply(getGroupResult =&gt; getGroupResult).Apply(example_readers =&gt; example_readers.Apply(getGroupResult =&gt; getGroupResult.Id)),
    ///         Permissions = 
    ///         {
    ///             { "CreateRepository", "Deny" },
    ///             { "DeleteRepository", "Deny" },
    ///             { "RenameRepository", "NotSet" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ### Shared Queries folder level
    /// 
    /// Permissions for a specific folder inside Shared Queries are specified if the arguments `project_id` and `path` are set.
    /// 
    /// &gt; **Note** To set permissions for the Shared Queries folder itself use `/` as path value
    /// 
    /// #### Example usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using AzureDevOps = Pulumi.AzureDevOps;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new AzureDevOps.Project("example", new()
    ///     {
    ///         WorkItemTemplate = "Agile",
    ///         VersionControl = "Git",
    ///         Visibility = "private",
    ///         Description = "Managed by Terraform",
    ///     });
    /// 
    ///     var example_readers = AzureDevOps.GetGroup.Invoke(new()
    ///     {
    ///         ProjectId = example.Id,
    ///         Name = "Readers",
    ///     });
    /// 
    ///     var example_permissions = new AzureDevOps.WorkItemQueryPermissions("example-permissions", new()
    ///     {
    ///         ProjectId = example.Id,
    ///         Path = "/Team",
    ///         Principal = example_readers.Apply(getGroupResult =&gt; getGroupResult).Apply(example_readers =&gt; example_readers.Apply(getGroupResult =&gt; getGroupResult.Id)),
    ///         Permissions = 
    ///         {
    ///             { "Contribute", "Allow" },
    ///             { "Delete", "Deny" },
    ///             { "Read", "NotSet" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
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
    ///     var example = new AzureDevOps.Project("example", new()
    ///     {
    ///         WorkItemTemplate = "Agile",
    ///         VersionControl = "Git",
    ///         Visibility = "private",
    ///         Description = "Managed by Terraform",
    ///     });
    /// 
    ///     var example_readers = AzureDevOps.GetGroup.Invoke(new()
    ///     {
    ///         ProjectId = example.Id,
    ///         Name = "Readers",
    ///     });
    /// 
    ///     var example_contributors = AzureDevOps.GetGroup.Invoke(new()
    ///     {
    ///         ProjectId = example.Id,
    ///         Name = "Contributors",
    ///     });
    /// 
    ///     var example_project_permissions = new AzureDevOps.WorkItemQueryPermissions("example-project-permissions", new()
    ///     {
    ///         ProjectId = example.Id,
    ///         Principal = example_readers.Apply(getGroupResult =&gt; getGroupResult).Apply(example_readers =&gt; example_readers.Apply(getGroupResult =&gt; getGroupResult.Id)),
    ///         Permissions = 
    ///         {
    ///             { "Read", "Allow" },
    ///             { "Delete", "Deny" },
    ///             { "Contribute", "Deny" },
    ///             { "ManagePermissions", "Deny" },
    ///         },
    ///     });
    /// 
    ///     var example_sharedqueries_permissions = new AzureDevOps.WorkItemQueryPermissions("example-sharedqueries-permissions", new()
    ///     {
    ///         ProjectId = example.Id,
    ///         Path = "/",
    ///         Principal = example_contributors.Apply(getGroupResult =&gt; getGroupResult).Apply(example_contributors =&gt; example_contributors.Apply(getGroupResult =&gt; getGroupResult.Id)),
    ///         Permissions = 
    ///         {
    ///             { "Read", "Allow" },
    ///             { "Delete", "Deny" },
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
    [AzureDevOpsResourceType("azuredevops:index/workItemQueryPermissions:WorkItemQueryPermissions")]
    public partial class WorkItemQueryPermissions : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Path to a query or folder beneath `Shared Queries`
        /// </summary>
        [Output("path")]
        public Output<string?> Path { get; private set; } = null!;

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
        /// Create a WorkItemQueryPermissions resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public WorkItemQueryPermissions(string name, WorkItemQueryPermissionsArgs args, CustomResourceOptions? options = null)
            : base("azuredevops:index/workItemQueryPermissions:WorkItemQueryPermissions", name, args ?? new WorkItemQueryPermissionsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private WorkItemQueryPermissions(string name, Input<string> id, WorkItemQueryPermissionsState? state = null, CustomResourceOptions? options = null)
            : base("azuredevops:index/workItemQueryPermissions:WorkItemQueryPermissions", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing WorkItemQueryPermissions resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static WorkItemQueryPermissions Get(string name, Input<string> id, WorkItemQueryPermissionsState? state = null, CustomResourceOptions? options = null)
        {
            return new WorkItemQueryPermissions(name, id, state, options);
        }
    }

    public sealed class WorkItemQueryPermissionsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Path to a query or folder beneath `Shared Queries`
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

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

        public WorkItemQueryPermissionsArgs()
        {
        }
        public static new WorkItemQueryPermissionsArgs Empty => new WorkItemQueryPermissionsArgs();
    }

    public sealed class WorkItemQueryPermissionsState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Path to a query or folder beneath `Shared Queries`
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

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

        public WorkItemQueryPermissionsState()
        {
        }
        public static new WorkItemQueryPermissionsState Empty => new WorkItemQueryPermissionsState();
    }
}
