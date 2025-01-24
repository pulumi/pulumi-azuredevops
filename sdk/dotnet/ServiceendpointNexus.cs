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
    /// Manages a Nexus IQ service endpoint within Azure DevOps, which can be used as a resource in YAML pipelines to connect to a Nexus IQ instance.
    /// Nexus IQ is not supported by default, to manage a nexus service connection resource, it is necessary to install the [Nexus Extension](https://marketplace.visualstudio.com/items?itemName=SonatypeIntegrations.nexus-iq-azure-extension) in Azure DevOps.
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
    ///         Visibility = "private",
    ///         VersionControl = "Git",
    ///         WorkItemTemplate = "Agile",
    ///         Description = "Managed by Pulumi",
    ///     });
    /// 
    ///     var exampleServiceendpointNexus = new AzureDevOps.ServiceendpointNexus("example", new()
    ///     {
    ///         ProjectId = example.Id,
    ///         ServiceEndpointName = "nexus-example",
    ///         Description = "Service Endpoint for 'Nexus IQ' (Managed by Terraform)",
    ///         Url = "https://example.com",
    ///         Username = "username",
    ///         Password = "password",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Azure DevOps Nexus Service Connection can be imported using the `projectId/id` or or `projectName/id`, e.g.
    /// 
    /// ```sh
    /// $ pulumi import azuredevops:index/serviceendpointNexus:ServiceendpointNexus example projectName/00000000-0000-0000-0000-000000000000
    /// ```
    /// </summary>
    [AzureDevOpsResourceType("azuredevops:index/serviceendpointNexus:ServiceendpointNexus")]
    public partial class ServiceendpointNexus : global::Pulumi.CustomResource
    {
        [Output("authorization")]
        public Output<ImmutableDictionary<string, string>> Authorization { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The Service Endpoint password to authenticate at the Nexus IQ Instance.
        /// </summary>
        [Output("password")]
        public Output<string> Password { get; private set; } = null!;

        /// <summary>
        /// The ID of the project. Changing this forces a new Service Connection Nexus to be created.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// The name of the service endpoint. Changing this forces a new Service Connection Nexus to be created.
        /// </summary>
        [Output("serviceEndpointName")]
        public Output<string> ServiceEndpointName { get; private set; } = null!;

        /// <summary>
        /// The Service Endpoint url.
        /// </summary>
        [Output("url")]
        public Output<string> Url { get; private set; } = null!;

        /// <summary>
        /// The Service Endpoint username to authenticate at the Nexus IQ Instance.
        /// </summary>
        [Output("username")]
        public Output<string> Username { get; private set; } = null!;


        /// <summary>
        /// Create a ServiceendpointNexus resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServiceendpointNexus(string name, ServiceendpointNexusArgs args, CustomResourceOptions? options = null)
            : base("azuredevops:index/serviceendpointNexus:ServiceendpointNexus", name, args ?? new ServiceendpointNexusArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServiceendpointNexus(string name, Input<string> id, ServiceendpointNexusState? state = null, CustomResourceOptions? options = null)
            : base("azuredevops:index/serviceendpointNexus:ServiceendpointNexus", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ServiceendpointNexus resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServiceendpointNexus Get(string name, Input<string> id, ServiceendpointNexusState? state = null, CustomResourceOptions? options = null)
        {
            return new ServiceendpointNexus(name, id, state, options);
        }
    }

    public sealed class ServiceendpointNexusArgs : global::Pulumi.ResourceArgs
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

        [Input("password", required: true)]
        private Input<string>? _password;

        /// <summary>
        /// The Service Endpoint password to authenticate at the Nexus IQ Instance.
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
        /// The ID of the project. Changing this forces a new Service Connection Nexus to be created.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        /// <summary>
        /// The name of the service endpoint. Changing this forces a new Service Connection Nexus to be created.
        /// </summary>
        [Input("serviceEndpointName", required: true)]
        public Input<string> ServiceEndpointName { get; set; } = null!;

        /// <summary>
        /// The Service Endpoint url.
        /// </summary>
        [Input("url", required: true)]
        public Input<string> Url { get; set; } = null!;

        /// <summary>
        /// The Service Endpoint username to authenticate at the Nexus IQ Instance.
        /// </summary>
        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public ServiceendpointNexusArgs()
        {
        }
        public static new ServiceendpointNexusArgs Empty => new ServiceendpointNexusArgs();
    }

    public sealed class ServiceendpointNexusState : global::Pulumi.ResourceArgs
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
        /// The Service Endpoint password to authenticate at the Nexus IQ Instance.
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
        /// The ID of the project. Changing this forces a new Service Connection Nexus to be created.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The name of the service endpoint. Changing this forces a new Service Connection Nexus to be created.
        /// </summary>
        [Input("serviceEndpointName")]
        public Input<string>? ServiceEndpointName { get; set; }

        /// <summary>
        /// The Service Endpoint url.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        /// <summary>
        /// The Service Endpoint username to authenticate at the Nexus IQ Instance.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public ServiceendpointNexusState()
        {
        }
        public static new ServiceendpointNexusState Empty => new ServiceendpointNexusState();
    }
}
