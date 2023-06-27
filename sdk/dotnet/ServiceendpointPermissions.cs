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
    /// Manages permissions for a Service Endpoint
    /// 
    /// &gt; **Note** Permissions can be assigned to group principals and not to single user principals.
    /// 
    /// ## Permission levels
    /// 
    /// Permission for Service Endpoints within Azure DevOps can be applied on two different levels.
    /// Those levels are reflected by specifying (or omitting) values for the arguments `project_id` and `serviceendpoint_id`.
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
    ///     var example_root_permissions = new AzureDevOps.ServiceendpointPermissions("example-root-permissions", new()
    ///     {
    ///         ProjectId = exampleProject.Id,
    ///         Principal = example_readers.Apply(example_readers =&gt; example_readers.Apply(getGroupResult =&gt; getGroupResult.Id)),
    ///         Permissions = 
    ///         {
    ///             { "Use", "allow" },
    ///             { "Administer", "allow" },
    ///             { "Create", "allow" },
    ///             { "ViewAuthorization", "allow" },
    ///             { "ViewEndpoint", "allow" },
    ///         },
    ///     });
    /// 
    ///     var exampleServiceEndpointDockerRegistry = new AzureDevOps.ServiceEndpointDockerRegistry("exampleServiceEndpointDockerRegistry", new()
    ///     {
    ///         ProjectId = exampleProject.Id,
    ///         ServiceEndpointName = "Example Docker Hub",
    ///         DockerUsername = "username",
    ///         DockerEmail = "email@example.com",
    ///         DockerPassword = "password",
    ///         RegistryType = "DockerHub",
    ///     });
    /// 
    ///     var example_permissions = new AzureDevOps.ServiceendpointPermissions("example-permissions", new()
    ///     {
    ///         ProjectId = exampleProject.Id,
    ///         Principal = example_readers.Apply(example_readers =&gt; example_readers.Apply(getGroupResult =&gt; getGroupResult.Id)),
    ///         ServiceendpointId = exampleServiceEndpointDockerRegistry.Id,
    ///         Permissions = 
    ///         {
    ///             { "Use", "allow" },
    ///             { "Administer", "deny" },
    ///             { "Create", "deny" },
    ///             { "ViewAuthorization", "allow" },
    ///             { "ViewEndpoint", "allow" },
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
    [AzureDevOpsResourceType("azuredevops:index/serviceendpointPermissions:ServiceendpointPermissions")]
    public partial class ServiceendpointPermissions : global::Pulumi.CustomResource
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
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// Replace (`true`) or merge (`false`) the permissions. Default: `true`
        /// 
        /// | Permission        | Description                         |
        /// | ----------------- | ----------------------------------- |
        /// | Use               | Use service endpoint                |
        /// | Administer        | Full control over service endpoints |
        /// | Create            | Create service endpoints            |
        /// | ViewAuthorization | View authorizations                 |
        /// | ViewEndpoint      | View service endpoint properties    |
        /// </summary>
        [Output("replace")]
        public Output<bool?> Replace { get; private set; } = null!;

        /// <summary>
        /// The id of the service endpoint to assign the permissions.
        /// </summary>
        [Output("serviceendpointId")]
        public Output<string?> ServiceendpointId { get; private set; } = null!;


        /// <summary>
        /// Create a ServiceendpointPermissions resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServiceendpointPermissions(string name, ServiceendpointPermissionsArgs args, CustomResourceOptions? options = null)
            : base("azuredevops:index/serviceendpointPermissions:ServiceendpointPermissions", name, args ?? new ServiceendpointPermissionsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServiceendpointPermissions(string name, Input<string> id, ServiceendpointPermissionsState? state = null, CustomResourceOptions? options = null)
            : base("azuredevops:index/serviceendpointPermissions:ServiceendpointPermissions", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ServiceendpointPermissions resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServiceendpointPermissions Get(string name, Input<string> id, ServiceendpointPermissionsState? state = null, CustomResourceOptions? options = null)
        {
            return new ServiceendpointPermissions(name, id, state, options);
        }
    }

    public sealed class ServiceendpointPermissionsArgs : global::Pulumi.ResourceArgs
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
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        /// <summary>
        /// Replace (`true`) or merge (`false`) the permissions. Default: `true`
        /// 
        /// | Permission        | Description                         |
        /// | ----------------- | ----------------------------------- |
        /// | Use               | Use service endpoint                |
        /// | Administer        | Full control over service endpoints |
        /// | Create            | Create service endpoints            |
        /// | ViewAuthorization | View authorizations                 |
        /// | ViewEndpoint      | View service endpoint properties    |
        /// </summary>
        [Input("replace")]
        public Input<bool>? Replace { get; set; }

        /// <summary>
        /// The id of the service endpoint to assign the permissions.
        /// </summary>
        [Input("serviceendpointId")]
        public Input<string>? ServiceendpointId { get; set; }

        public ServiceendpointPermissionsArgs()
        {
        }
        public static new ServiceendpointPermissionsArgs Empty => new ServiceendpointPermissionsArgs();
    }

    public sealed class ServiceendpointPermissionsState : global::Pulumi.ResourceArgs
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
        /// 
        /// | Permission        | Description                         |
        /// | ----------------- | ----------------------------------- |
        /// | Use               | Use service endpoint                |
        /// | Administer        | Full control over service endpoints |
        /// | Create            | Create service endpoints            |
        /// | ViewAuthorization | View authorizations                 |
        /// | ViewEndpoint      | View service endpoint properties    |
        /// </summary>
        [Input("replace")]
        public Input<bool>? Replace { get; set; }

        /// <summary>
        /// The id of the service endpoint to assign the permissions.
        /// </summary>
        [Input("serviceendpointId")]
        public Input<string>? ServiceendpointId { get; set; }

        public ServiceendpointPermissionsState()
        {
        }
        public static new ServiceendpointPermissionsState Empty => new ServiceendpointPermissionsState();
    }
}
