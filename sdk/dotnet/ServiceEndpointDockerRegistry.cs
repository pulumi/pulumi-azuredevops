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
    /// Manages a Docker Registry service endpoint within Azure DevOps.
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
    ///             Visibility = "private",
    ///             VersionControl = "Git",
    ///             WorkItemTemplate = "Agile",
    ///         });
    ///         // dockerhub registry service connection
    ///         var dockerhubregistry = new AzureDevOps.ServiceEndpointDockerRegistry("dockerhubregistry", new AzureDevOps.ServiceEndpointDockerRegistryArgs
    ///         {
    ///             ProjectId = project.Id,
    ///             ServiceEndpointName = "Sample Docker Hub",
    ///             DockerUsername = "sample",
    ///             DockerEmail = "email@example.com",
    ///             DockerPassword = "12345",
    ///             RegistryType = "DockerHub",
    ///         });
    ///         // other docker registry service connection
    ///         var otherregistry = new AzureDevOps.ServiceEndpointDockerRegistry("otherregistry", new AzureDevOps.ServiceEndpointDockerRegistryArgs
    ///         {
    ///             ProjectId = project.Id,
    ///             ServiceEndpointName = "Sample Docker Registry",
    ///             DockerRegistry = "https://sample.azurecr.io/v1",
    ///             DockerUsername = "sample",
    ///             DockerPassword = "12345",
    ///             RegistryType = "Others",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ## Relevant Links
    /// 
    /// - [Azure DevOps Service REST API 5.1 - Service Endpoints](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-5.1)
    /// - [Docker Registry Service Connection](https://docs.microsoft.com/en-us/azure/devops/pipelines/library/service-endpoints?view=azure-devops&amp;tabs=yaml#sep-docreg)
    /// </summary>
    public partial class ServiceEndpointDockerRegistry : Pulumi.CustomResource
    {
        [Output("authorization")]
        public Output<ImmutableDictionary<string, string>> Authorization { get; private set; } = null!;

        /// <summary>
        /// The name you will use to refer to this service connection in task inputs.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The email for Docker account user.
        /// </summary>
        [Output("dockerEmail")]
        public Output<string?> DockerEmail { get; private set; } = null!;

        /// <summary>
        /// The password for the account user identified above.
        /// </summary>
        [Output("dockerPassword")]
        public Output<string?> DockerPassword { get; private set; } = null!;

        /// <summary>
        /// A bcrypted hash of the attribute 'docker_password'
        /// </summary>
        [Output("dockerPasswordHash")]
        public Output<string> DockerPasswordHash { get; private set; } = null!;

        /// <summary>
        /// The URL of the Docker registry. (Default: "https://index.docker.io/v1/")
        /// </summary>
        [Output("dockerRegistry")]
        public Output<string> DockerRegistryUrl { get; private set; } = null!;

        /// <summary>
        /// The identifier of the Docker account user.
        /// </summary>
        [Output("dockerUsername")]
        public Output<string?> DockerUsername { get; private set; } = null!;

        /// <summary>
        /// The project ID or project name.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// Can be "DockerHub" or "Others" (Default "DockerHub")
        /// </summary>
        [Output("registryType")]
        public Output<string> RegistryType { get; private set; } = null!;

        /// <summary>
        /// The name you will use to refer to this service connection in task inputs.
        /// </summary>
        [Output("serviceEndpointName")]
        public Output<string> ServiceEndpointName { get; private set; } = null!;


        /// <summary>
        /// Create a ServiceEndpointDockerRegistry resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServiceEndpointDockerRegistry(string name, ServiceEndpointDockerRegistryArgs args, CustomResourceOptions? options = null)
            : base("azuredevops:index/serviceEndpointDockerRegistry:ServiceEndpointDockerRegistry", name, args ?? new ServiceEndpointDockerRegistryArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServiceEndpointDockerRegistry(string name, Input<string> id, ServiceEndpointDockerRegistryState? state = null, CustomResourceOptions? options = null)
            : base("azuredevops:index/serviceEndpointDockerRegistry:ServiceEndpointDockerRegistry", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azuredevops:ServiceEndpoint/dockerRegistry:DockerRegistry"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ServiceEndpointDockerRegistry resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServiceEndpointDockerRegistry Get(string name, Input<string> id, ServiceEndpointDockerRegistryState? state = null, CustomResourceOptions? options = null)
        {
            return new ServiceEndpointDockerRegistry(name, id, state, options);
        }
    }

    public sealed class ServiceEndpointDockerRegistryArgs : Pulumi.ResourceArgs
    {
        [Input("authorization")]
        private InputMap<string>? _authorization;
        public InputMap<string> Authorization
        {
            get => _authorization ?? (_authorization = new InputMap<string>());
            set => _authorization = value;
        }

        /// <summary>
        /// The name you will use to refer to this service connection in task inputs.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The email for Docker account user.
        /// </summary>
        [Input("dockerEmail")]
        public Input<string>? DockerEmail { get; set; }

        /// <summary>
        /// The password for the account user identified above.
        /// </summary>
        [Input("dockerPassword")]
        public Input<string>? DockerPassword { get; set; }

        /// <summary>
        /// The URL of the Docker registry. (Default: "https://index.docker.io/v1/")
        /// </summary>
        [Input("dockerRegistry", required: true)]
        public Input<string> DockerRegistryUrl { get; set; } = null!;

        /// <summary>
        /// The identifier of the Docker account user.
        /// </summary>
        [Input("dockerUsername")]
        public Input<string>? DockerUsername { get; set; }

        /// <summary>
        /// The project ID or project name.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        /// <summary>
        /// Can be "DockerHub" or "Others" (Default "DockerHub")
        /// </summary>
        [Input("registryType", required: true)]
        public Input<string> RegistryType { get; set; } = null!;

        /// <summary>
        /// The name you will use to refer to this service connection in task inputs.
        /// </summary>
        [Input("serviceEndpointName", required: true)]
        public Input<string> ServiceEndpointName { get; set; } = null!;

        public ServiceEndpointDockerRegistryArgs()
        {
        }
    }

    public sealed class ServiceEndpointDockerRegistryState : Pulumi.ResourceArgs
    {
        [Input("authorization")]
        private InputMap<string>? _authorization;
        public InputMap<string> Authorization
        {
            get => _authorization ?? (_authorization = new InputMap<string>());
            set => _authorization = value;
        }

        /// <summary>
        /// The name you will use to refer to this service connection in task inputs.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The email for Docker account user.
        /// </summary>
        [Input("dockerEmail")]
        public Input<string>? DockerEmail { get; set; }

        /// <summary>
        /// The password for the account user identified above.
        /// </summary>
        [Input("dockerPassword")]
        public Input<string>? DockerPassword { get; set; }

        /// <summary>
        /// A bcrypted hash of the attribute 'docker_password'
        /// </summary>
        [Input("dockerPasswordHash")]
        public Input<string>? DockerPasswordHash { get; set; }

        /// <summary>
        /// The URL of the Docker registry. (Default: "https://index.docker.io/v1/")
        /// </summary>
        [Input("dockerRegistry")]
        public Input<string>? DockerRegistryUrl { get; set; }

        /// <summary>
        /// The identifier of the Docker account user.
        /// </summary>
        [Input("dockerUsername")]
        public Input<string>? DockerUsername { get; set; }

        /// <summary>
        /// The project ID or project name.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// Can be "DockerHub" or "Others" (Default "DockerHub")
        /// </summary>
        [Input("registryType")]
        public Input<string>? RegistryType { get; set; }

        /// <summary>
        /// The name you will use to refer to this service connection in task inputs.
        /// </summary>
        [Input("serviceEndpointName")]
        public Input<string>? ServiceEndpointName { get; set; }

        public ServiceEndpointDockerRegistryState()
        {
        }
    }
}