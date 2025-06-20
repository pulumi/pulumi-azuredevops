// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureDevOps
{
    /// <summary>
    /// Manages permissions for tagging
    /// 
    /// ## Permission levels
    /// 
    /// Permissions for tagging within Azure DevOps can be applied only on Organizational and Project level.
    /// The project level is reflected by specifying the argument `project_id`, otherwise the permissions are set on the organizational level.
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
    ///     var example_readers = AzureDevOps.GetGroup.Invoke(new()
    ///     {
    ///         ProjectId = example.Id,
    ///         Name = "Readers",
    ///     });
    /// 
    ///     var example_permissions = new AzureDevOps.TaggingPermissions("example-permissions", new()
    ///     {
    ///         ProjectId = example.Id,
    ///         Principal = example_readers.Apply(example_readers =&gt; example_readers.Apply(getGroupResult =&gt; getGroupResult.Id)),
    ///         Permissions = 
    ///         {
    ///             { "Enumerate", "allow" },
    ///             { "Create", "allow" },
    ///             { "Update", "allow" },
    ///             { "Delete", "allow" },
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
    [AzureDevOpsResourceType("azuredevops:index/taggingPermissions:TaggingPermissions")]
    public partial class TaggingPermissions : global::Pulumi.CustomResource
    {
        /// <summary>
        /// the permissions to assign. The following permissions are available.
        /// 
        /// | Name      | Permission Description    |
        /// |-----------|---------------------------|
        /// | Enumerate | Enumerate tag definitions |
        /// | Create    | Create tag definition     |
        /// | Update    | Update tag definition     |
        /// | Delete    | Delete tag definition     |
        /// </summary>
        [Output("permissions")]
        public Output<ImmutableDictionary<string, string>> Permissions { get; private set; } = null!;

        /// <summary>
        /// The **group or user** principal to assign the permissions.
        /// </summary>
        [Output("principal")]
        public Output<string> Principal { get; private set; } = null!;

        /// <summary>
        /// The ID of the project to assign the permissions. If omitted, organization wide permissions for tagging are managed.
        /// </summary>
        [Output("projectId")]
        public Output<string?> ProjectId { get; private set; } = null!;

        /// <summary>
        /// Replace (`true`) or merge (`false`) the permissions. Default: `true`
        /// </summary>
        [Output("replace")]
        public Output<bool?> Replace { get; private set; } = null!;


        /// <summary>
        /// Create a TaggingPermissions resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TaggingPermissions(string name, TaggingPermissionsArgs args, CustomResourceOptions? options = null)
            : base("azuredevops:index/taggingPermissions:TaggingPermissions", name, args ?? new TaggingPermissionsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TaggingPermissions(string name, Input<string> id, TaggingPermissionsState? state = null, CustomResourceOptions? options = null)
            : base("azuredevops:index/taggingPermissions:TaggingPermissions", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing TaggingPermissions resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TaggingPermissions Get(string name, Input<string> id, TaggingPermissionsState? state = null, CustomResourceOptions? options = null)
        {
            return new TaggingPermissions(name, id, state, options);
        }
    }

    public sealed class TaggingPermissionsArgs : global::Pulumi.ResourceArgs
    {
        [Input("permissions", required: true)]
        private InputMap<string>? _permissions;

        /// <summary>
        /// the permissions to assign. The following permissions are available.
        /// 
        /// | Name      | Permission Description    |
        /// |-----------|---------------------------|
        /// | Enumerate | Enumerate tag definitions |
        /// | Create    | Create tag definition     |
        /// | Update    | Update tag definition     |
        /// | Delete    | Delete tag definition     |
        /// </summary>
        public InputMap<string> Permissions
        {
            get => _permissions ?? (_permissions = new InputMap<string>());
            set => _permissions = value;
        }

        /// <summary>
        /// The **group or user** principal to assign the permissions.
        /// </summary>
        [Input("principal", required: true)]
        public Input<string> Principal { get; set; } = null!;

        /// <summary>
        /// The ID of the project to assign the permissions. If omitted, organization wide permissions for tagging are managed.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// Replace (`true`) or merge (`false`) the permissions. Default: `true`
        /// </summary>
        [Input("replace")]
        public Input<bool>? Replace { get; set; }

        public TaggingPermissionsArgs()
        {
        }
        public static new TaggingPermissionsArgs Empty => new TaggingPermissionsArgs();
    }

    public sealed class TaggingPermissionsState : global::Pulumi.ResourceArgs
    {
        [Input("permissions")]
        private InputMap<string>? _permissions;

        /// <summary>
        /// the permissions to assign. The following permissions are available.
        /// 
        /// | Name      | Permission Description    |
        /// |-----------|---------------------------|
        /// | Enumerate | Enumerate tag definitions |
        /// | Create    | Create tag definition     |
        /// | Update    | Update tag definition     |
        /// | Delete    | Delete tag definition     |
        /// </summary>
        public InputMap<string> Permissions
        {
            get => _permissions ?? (_permissions = new InputMap<string>());
            set => _permissions = value;
        }

        /// <summary>
        /// The **group or user** principal to assign the permissions.
        /// </summary>
        [Input("principal")]
        public Input<string>? Principal { get; set; }

        /// <summary>
        /// The ID of the project to assign the permissions. If omitted, organization wide permissions for tagging are managed.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// Replace (`true`) or merge (`false`) the permissions. Default: `true`
        /// </summary>
        [Input("replace")]
        public Input<bool>? Replace { get; set; }

        public TaggingPermissionsState()
        {
        }
        public static new TaggingPermissionsState Empty => new TaggingPermissionsState();
    }
}
