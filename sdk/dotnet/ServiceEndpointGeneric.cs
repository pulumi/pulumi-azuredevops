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
    /// Manages a generic service endpoint within Azure DevOps, which can be used to authenticate to any external server using
    /// basic authentication via a username and password.
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
    ///         Visibility = "private",
    ///         VersionControl = "Git",
    ///         WorkItemTemplate = "Agile",
    ///         Description = "Managed by Terraform",
    ///     });
    /// 
    ///     var exampleServiceEndpointGeneric = new AzureDevOps.ServiceEndpointGeneric("exampleServiceEndpointGeneric", new()
    ///     {
    ///         ProjectId = exampleProject.Id,
    ///         ServerUrl = "https://some-server.example.com",
    ///         Username = "username",
    ///         Password = "password",
    ///         ServiceEndpointName = "Example Generic",
    ///         Description = "Managed by Terraform",
    ///     });
    /// 
    /// });
    /// ```
    /// ## Relevant Links
    /// 
    /// - [Azure DevOps Service REST API 7.0 - Agent Pools](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-7.0)
    /// 
    /// ## Import
    /// 
    /// Azure DevOps Service Endpoint Generic can be imported using **projectID/serviceEndpointID** or
    /// 
    ///  **projectName/serviceEndpointID**
    /// 
    /// ```sh
    /// $ pulumi import azuredevops:index/serviceEndpointGeneric:ServiceEndpointGeneric example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
    /// ```
    /// </summary>
    [AzureDevOpsResourceType("azuredevops:index/serviceEndpointGeneric:ServiceEndpointGeneric")]
    public partial class ServiceEndpointGeneric : global::Pulumi.CustomResource
    {
        [Output("authorization")]
        public Output<ImmutableDictionary<string, string>> Authorization { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The password or token key used to authenticate to the server url using basic authentication.
        /// </summary>
        [Output("password")]
        public Output<string?> Password { get; private set; } = null!;

        /// <summary>
        /// The ID of the project.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// The URL of the server associated with the service endpoint.
        /// </summary>
        [Output("serverUrl")]
        public Output<string> ServerUrl { get; private set; } = null!;

        /// <summary>
        /// The service endpoint name.
        /// </summary>
        [Output("serviceEndpointName")]
        public Output<string> ServiceEndpointName { get; private set; } = null!;

        /// <summary>
        /// The username used to authenticate to the server url using basic authentication.
        /// </summary>
        [Output("username")]
        public Output<string?> Username { get; private set; } = null!;


        /// <summary>
        /// Create a ServiceEndpointGeneric resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServiceEndpointGeneric(string name, ServiceEndpointGenericArgs args, CustomResourceOptions? options = null)
            : base("azuredevops:index/serviceEndpointGeneric:ServiceEndpointGeneric", name, args ?? new ServiceEndpointGenericArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServiceEndpointGeneric(string name, Input<string> id, ServiceEndpointGenericState? state = null, CustomResourceOptions? options = null)
            : base("azuredevops:index/serviceEndpointGeneric:ServiceEndpointGeneric", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "password",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ServiceEndpointGeneric resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServiceEndpointGeneric Get(string name, Input<string> id, ServiceEndpointGenericState? state = null, CustomResourceOptions? options = null)
        {
            return new ServiceEndpointGeneric(name, id, state, options);
        }
    }

    public sealed class ServiceEndpointGenericArgs : global::Pulumi.ResourceArgs
    {
        [Input("authorization")]
        private InputMap<string>? _authorization;
        public InputMap<string> Authorization
        {
            get => _authorization ?? (_authorization = new InputMap<string>());
            set => _authorization = value;
        }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// The password or token key used to authenticate to the server url using basic authentication.
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The ID of the project.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        /// <summary>
        /// The URL of the server associated with the service endpoint.
        /// </summary>
        [Input("serverUrl", required: true)]
        public Input<string> ServerUrl { get; set; } = null!;

        /// <summary>
        /// The service endpoint name.
        /// </summary>
        [Input("serviceEndpointName", required: true)]
        public Input<string> ServiceEndpointName { get; set; } = null!;

        /// <summary>
        /// The username used to authenticate to the server url using basic authentication.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public ServiceEndpointGenericArgs()
        {
        }
        public static new ServiceEndpointGenericArgs Empty => new ServiceEndpointGenericArgs();
    }

    public sealed class ServiceEndpointGenericState : global::Pulumi.ResourceArgs
    {
        [Input("authorization")]
        private InputMap<string>? _authorization;
        public InputMap<string> Authorization
        {
            get => _authorization ?? (_authorization = new InputMap<string>());
            set => _authorization = value;
        }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// The password or token key used to authenticate to the server url using basic authentication.
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The ID of the project.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The URL of the server associated with the service endpoint.
        /// </summary>
        [Input("serverUrl")]
        public Input<string>? ServerUrl { get; set; }

        /// <summary>
        /// The service endpoint name.
        /// </summary>
        [Input("serviceEndpointName")]
        public Input<string>? ServiceEndpointName { get; set; }

        /// <summary>
        /// The username used to authenticate to the server url using basic authentication.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public ServiceEndpointGenericState()
        {
        }
        public static new ServiceEndpointGenericState Empty => new ServiceEndpointGenericState();
    }
}
