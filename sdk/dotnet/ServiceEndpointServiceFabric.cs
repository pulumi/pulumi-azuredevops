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
    /// Manages a Service Fabric service endpoint within Azure DevOps.
    /// 
    /// ## Example Usage
    /// ### Client Certificate Authentication
    /// 
    /// ```csharp
    /// using System;
    /// using System.IO;
    /// using Pulumi;
    /// using AzureDevOps = Pulumi.AzureDevOps;
    /// 
    /// class MyStack : Stack
    /// {
    /// 	private static string ReadFileBase64(string path) {
    /// 		return Convert.ToBase64String(System.Text.UTF8.GetBytes(File.ReadAllText(path)))
    /// 	}
    /// 
    ///     public MyStack()
    ///     {
    ///         var project = new AzureDevOps.Project("project", new AzureDevOps.ProjectArgs
    ///         {
    ///             Visibility = "private",
    ///             VersionControl = "Git",
    ///             WorkItemTemplate = "Agile",
    ///         });
    ///         var test = new AzureDevOps.ServiceEndpointServiceFabric("test", new AzureDevOps.ServiceEndpointServiceFabricArgs
    ///         {
    ///             ProjectId = project.Id,
    ///             ServiceEndpointName = "Sample Service Fabric",
    ///             Description = "Managed by Terraform",
    ///             ClusterEndpoint = "tcp://test",
    ///             Certificate = new AzureDevOps.Inputs.ServiceEndpointServiceFabricCertificateArgs
    ///             {
    ///                 ServerCertificateLookup = "Thumbprint",
    ///                 ServerCertificateThumbprint = "0000000000000000000000000000000000000000",
    ///                 ClientCertificate = ReadFileBase64("certificate.pfx"),
    ///                 ClientCertificatePassword = "password",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Azure Active Directory Authentication
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
    ///             Visibility = "private",
    ///             VersionControl = "Git",
    ///             WorkItemTemplate = "Agile",
    ///         });
    ///         var test = new AzureDevOps.ServiceEndpointServiceFabric("test", new AzureDevOps.ServiceEndpointServiceFabricArgs
    ///         {
    ///             ProjectId = project.Id,
    ///             ServiceEndpointName = "Sample Service Fabric",
    ///             Description = "Managed by Terraform",
    ///             ClusterEndpoint = "tcp://test",
    ///             AzureActiveDirectory = new AzureDevOps.Inputs.ServiceEndpointServiceFabricAzureActiveDirectoryArgs
    ///             {
    ///                 ServerCertificateLookup = "Thumbprint",
    ///                 ServerCertificateThumbprint = "0000000000000000000000000000000000000000",
    ///                 Username = "username",
    ///                 Password = "password",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Windows Authentication
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
    ///             Visibility = "private",
    ///             VersionControl = "Git",
    ///             WorkItemTemplate = "Agile",
    ///         });
    ///         var test = new AzureDevOps.ServiceEndpointServiceFabric("test", new AzureDevOps.ServiceEndpointServiceFabricArgs
    ///         {
    ///             ProjectId = project.Id,
    ///             ServiceEndpointName = "Sample Service Fabric",
    ///             Description = "Managed by Terraform",
    ///             ClusterEndpoint = "tcp://test",
    ///             None = new AzureDevOps.Inputs.ServiceEndpointServiceFabricNoneArgs
    ///             {
    ///                 Unsecured = false,
    ///                 ClusterSpn = "HTTP/www.contoso.com",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ## Relevant Links
    /// 
    /// - [Azure DevOps Service REST API 5.1 - Endpoints](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-5.1)
    /// 
    /// ## Import
    /// 
    /// Azure DevOps Service Endpoint Service Fabric can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**
    /// 
    /// ```sh
    ///  $ pulumi import azuredevops:index/serviceEndpointServiceFabric:ServiceEndpointServiceFabric serviceendpoint 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
    /// ```
    /// </summary>
    [AzureDevOpsResourceType("azuredevops:index/serviceEndpointServiceFabric:ServiceEndpointServiceFabric")]
    public partial class ServiceEndpointServiceFabric : Pulumi.CustomResource
    {
        [Output("authorization")]
        public Output<ImmutableDictionary<string, string>> Authorization { get; private set; } = null!;

        [Output("azureActiveDirectory")]
        public Output<Outputs.ServiceEndpointServiceFabricAzureActiveDirectory?> AzureActiveDirectory { get; private set; } = null!;

        [Output("certificate")]
        public Output<Outputs.ServiceEndpointServiceFabricCertificate?> Certificate { get; private set; } = null!;

        /// <summary>
        /// Client connection endpoint for the cluster. Prefix the value with 'tcp://';. This value overrides the publish profile.
        /// </summary>
        [Output("clusterEndpoint")]
        public Output<string> ClusterEndpoint { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("none")]
        public Output<Outputs.ServiceEndpointServiceFabricNone?> None { get; private set; } = null!;

        /// <summary>
        /// The project ID or project name.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// The Service Endpoint name.
        /// </summary>
        [Output("serviceEndpointName")]
        public Output<string> ServiceEndpointName { get; private set; } = null!;


        /// <summary>
        /// Create a ServiceEndpointServiceFabric resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServiceEndpointServiceFabric(string name, ServiceEndpointServiceFabricArgs args, CustomResourceOptions? options = null)
            : base("azuredevops:index/serviceEndpointServiceFabric:ServiceEndpointServiceFabric", name, args ?? new ServiceEndpointServiceFabricArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServiceEndpointServiceFabric(string name, Input<string> id, ServiceEndpointServiceFabricState? state = null, CustomResourceOptions? options = null)
            : base("azuredevops:index/serviceEndpointServiceFabric:ServiceEndpointServiceFabric", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ServiceEndpointServiceFabric resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServiceEndpointServiceFabric Get(string name, Input<string> id, ServiceEndpointServiceFabricState? state = null, CustomResourceOptions? options = null)
        {
            return new ServiceEndpointServiceFabric(name, id, state, options);
        }
    }

    public sealed class ServiceEndpointServiceFabricArgs : Pulumi.ResourceArgs
    {
        [Input("authorization")]
        private InputMap<string>? _authorization;
        public InputMap<string> Authorization
        {
            get => _authorization ?? (_authorization = new InputMap<string>());
            set => _authorization = value;
        }

        [Input("azureActiveDirectory")]
        public Input<Inputs.ServiceEndpointServiceFabricAzureActiveDirectoryArgs>? AzureActiveDirectory { get; set; }

        [Input("certificate")]
        public Input<Inputs.ServiceEndpointServiceFabricCertificateArgs>? Certificate { get; set; }

        /// <summary>
        /// Client connection endpoint for the cluster. Prefix the value with 'tcp://';. This value overrides the publish profile.
        /// </summary>
        [Input("clusterEndpoint", required: true)]
        public Input<string> ClusterEndpoint { get; set; } = null!;

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("none")]
        public Input<Inputs.ServiceEndpointServiceFabricNoneArgs>? None { get; set; }

        /// <summary>
        /// The project ID or project name.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        /// <summary>
        /// The Service Endpoint name.
        /// </summary>
        [Input("serviceEndpointName", required: true)]
        public Input<string> ServiceEndpointName { get; set; } = null!;

        public ServiceEndpointServiceFabricArgs()
        {
        }
    }

    public sealed class ServiceEndpointServiceFabricState : Pulumi.ResourceArgs
    {
        [Input("authorization")]
        private InputMap<string>? _authorization;
        public InputMap<string> Authorization
        {
            get => _authorization ?? (_authorization = new InputMap<string>());
            set => _authorization = value;
        }

        [Input("azureActiveDirectory")]
        public Input<Inputs.ServiceEndpointServiceFabricAzureActiveDirectoryGetArgs>? AzureActiveDirectory { get; set; }

        [Input("certificate")]
        public Input<Inputs.ServiceEndpointServiceFabricCertificateGetArgs>? Certificate { get; set; }

        /// <summary>
        /// Client connection endpoint for the cluster. Prefix the value with 'tcp://';. This value overrides the publish profile.
        /// </summary>
        [Input("clusterEndpoint")]
        public Input<string>? ClusterEndpoint { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("none")]
        public Input<Inputs.ServiceEndpointServiceFabricNoneGetArgs>? None { get; set; }

        /// <summary>
        /// The project ID or project name.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The Service Endpoint name.
        /// </summary>
        [Input("serviceEndpointName")]
        public Input<string>? ServiceEndpointName { get; set; }

        public ServiceEndpointServiceFabricState()
        {
        }
    }
}