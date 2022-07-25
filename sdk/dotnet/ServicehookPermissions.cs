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
    /// Manages permissions for service hooks
    /// 
    /// ## Permission levels
    /// 
    /// Permissions for service hooks within Azure DevOps can be applied on the Organizational level or, if the optional attribute `project_id` is specified, on Project level.
    /// Those levels are reflected by specifying (or omitting) values for the argument `project_id`.
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
    ///         var example = new AzureDevOps.Project("example", new AzureDevOps.ProjectArgs
    ///         {
    ///             WorkItemTemplate = "Agile",
    ///             VersionControl = "Git",
    ///             Visibility = "private",
    ///             Description = "Managed by Terraform",
    ///         });
    ///         var example_readers = AzureDevOps.GetGroup.Invoke(new AzureDevOps.GetGroupInvokeArgs
    ///         {
    ///             ProjectId = example.Id,
    ///             Name = "Readers",
    ///         });
    ///         var example_permissions = new AzureDevOps.ServicehookPermissions("example-permissions", new AzureDevOps.ServicehookPermissionsArgs
    ///         {
    ///             ProjectId = example.Id,
    ///             Principal = example_readers.Apply(example_readers =&gt; example_readers.Id),
    ///             Permissions = 
    ///             {
    ///                 { "ViewSubscriptions", "allow" },
    ///                 { "EditSubscriptions", "allow" },
    ///                 { "DeleteSubscriptions", "allow" },
    ///                 { "PublishEvents", "allow" },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
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
    [AzureDevOpsResourceType("azuredevops:index/servicehookPermissions:ServicehookPermissions")]
    public partial class ServicehookPermissions : Pulumi.CustomResource
    {
        /// <summary>
        /// the permissions to assign. The following permissions are available.
        /// </summary>
        [Output("permissions")]
        public Output<ImmutableDictionary<string, string>> Permissions { get; private set; } = null!;

        /// <summary>
        /// The **group** principal to assign the permissions.
        /// </summary>
        [Output("principal")]
        public Output<string> Principal { get; private set; } = null!;

        /// <summary>
        /// The ID of the project.
        /// </summary>
        [Output("projectId")]
        public Output<string?> ProjectId { get; private set; } = null!;

        /// <summary>
        /// Replace (`true`) or merge (`false`) the permissions. Default: `true`
        /// </summary>
        [Output("replace")]
        public Output<bool?> Replace { get; private set; } = null!;


        /// <summary>
        /// Create a ServicehookPermissions resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServicehookPermissions(string name, ServicehookPermissionsArgs args, CustomResourceOptions? options = null)
            : base("azuredevops:index/servicehookPermissions:ServicehookPermissions", name, args ?? new ServicehookPermissionsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServicehookPermissions(string name, Input<string> id, ServicehookPermissionsState? state = null, CustomResourceOptions? options = null)
            : base("azuredevops:index/servicehookPermissions:ServicehookPermissions", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ServicehookPermissions resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServicehookPermissions Get(string name, Input<string> id, ServicehookPermissionsState? state = null, CustomResourceOptions? options = null)
        {
            return new ServicehookPermissions(name, id, state, options);
        }
    }

    public sealed class ServicehookPermissionsArgs : Pulumi.ResourceArgs
    {
        [Input("permissions", required: true)]
        private InputMap<string>? _permissions;

        /// <summary>
        /// the permissions to assign. The following permissions are available.
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
        /// The ID of the project.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// Replace (`true`) or merge (`false`) the permissions. Default: `true`
        /// </summary>
        [Input("replace")]
        public Input<bool>? Replace { get; set; }

        public ServicehookPermissionsArgs()
        {
        }
    }

    public sealed class ServicehookPermissionsState : Pulumi.ResourceArgs
    {
        [Input("permissions")]
        private InputMap<string>? _permissions;

        /// <summary>
        /// the permissions to assign. The following permissions are available.
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
        /// The ID of the project.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// Replace (`true`) or merge (`false`) the permissions. Default: `true`
        /// </summary>
        [Input("replace")]
        public Input<bool>? Replace { get; set; }

        public ServicehookPermissionsState()
        {
        }
    }
}