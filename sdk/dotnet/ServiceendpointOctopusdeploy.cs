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
    /// Manages an Octopus Deploy service endpoint within Azure DevOps. Using this service endpoint requires you to install [Octopus Deploy](https://marketplace.visualstudio.com/items?itemName=octopusdeploy.octopus-deploy-build-release-tasks).
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
    ///     var exampleServiceendpointOctopusdeploy = new AzureDevOps.ServiceendpointOctopusdeploy("example", new()
    ///     {
    ///         ProjectId = example.Id,
    ///         Url = "https://octopus.com",
    ///         ApiKey = "000000000000000000000000000000000000",
    ///         ServiceEndpointName = "Example Octopus Deploy",
    ///         Description = "Managed by Pulumi",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Relevant Links
    /// 
    /// - [Azure DevOps Service REST API 7.0 - Agent Pools](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-7.0)
    /// 
    /// ## Import
    /// 
    /// Azure DevOps Octopus Deploy Service Endpoint can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**
    /// 
    /// ```sh
    /// $ pulumi import azuredevops:index/serviceendpointOctopusdeploy:ServiceendpointOctopusdeploy example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
    /// ```
    /// </summary>
    [AzureDevOpsResourceType("azuredevops:index/serviceendpointOctopusdeploy:ServiceendpointOctopusdeploy")]
    public partial class ServiceendpointOctopusdeploy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// API key to connect to Octopus Deploy.
        /// </summary>
        [Output("apiKey")]
        public Output<string> ApiKey { get; private set; } = null!;

        [Output("authorization")]
        public Output<ImmutableDictionary<string, string>> Authorization { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Whether to ignore SSL errors when connecting to the Octopus server from the agent. Default to `false`.
        /// </summary>
        [Output("ignoreSslError")]
        public Output<bool?> IgnoreSslError { get; private set; } = null!;

        /// <summary>
        /// The ID of the project.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// The Service Endpoint name.
        /// </summary>
        [Output("serviceEndpointName")]
        public Output<string> ServiceEndpointName { get; private set; } = null!;

        /// <summary>
        /// Octopus Server url.
        /// </summary>
        [Output("url")]
        public Output<string> Url { get; private set; } = null!;


        /// <summary>
        /// Create a ServiceendpointOctopusdeploy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServiceendpointOctopusdeploy(string name, ServiceendpointOctopusdeployArgs args, CustomResourceOptions? options = null)
            : base("azuredevops:index/serviceendpointOctopusdeploy:ServiceendpointOctopusdeploy", name, args ?? new ServiceendpointOctopusdeployArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServiceendpointOctopusdeploy(string name, Input<string> id, ServiceendpointOctopusdeployState? state = null, CustomResourceOptions? options = null)
            : base("azuredevops:index/serviceendpointOctopusdeploy:ServiceendpointOctopusdeploy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ServiceendpointOctopusdeploy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServiceendpointOctopusdeploy Get(string name, Input<string> id, ServiceendpointOctopusdeployState? state = null, CustomResourceOptions? options = null)
        {
            return new ServiceendpointOctopusdeploy(name, id, state, options);
        }
    }

    public sealed class ServiceendpointOctopusdeployArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// API key to connect to Octopus Deploy.
        /// </summary>
        [Input("apiKey", required: true)]
        public Input<string> ApiKey { get; set; } = null!;

        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Whether to ignore SSL errors when connecting to the Octopus server from the agent. Default to `false`.
        /// </summary>
        [Input("ignoreSslError")]
        public Input<bool>? IgnoreSslError { get; set; }

        /// <summary>
        /// The ID of the project.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        /// <summary>
        /// The Service Endpoint name.
        /// </summary>
        [Input("serviceEndpointName", required: true)]
        public Input<string> ServiceEndpointName { get; set; } = null!;

        /// <summary>
        /// Octopus Server url.
        /// </summary>
        [Input("url", required: true)]
        public Input<string> Url { get; set; } = null!;

        public ServiceendpointOctopusdeployArgs()
        {
        }
        public static new ServiceendpointOctopusdeployArgs Empty => new ServiceendpointOctopusdeployArgs();
    }

    public sealed class ServiceendpointOctopusdeployState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// API key to connect to Octopus Deploy.
        /// </summary>
        [Input("apiKey")]
        public Input<string>? ApiKey { get; set; }

        [Input("authorization")]
        private InputMap<string>? _authorization;
        public InputMap<string> Authorization
        {
            get => _authorization ?? (_authorization = new InputMap<string>());
            set => _authorization = value;
        }

        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Whether to ignore SSL errors when connecting to the Octopus server from the agent. Default to `false`.
        /// </summary>
        [Input("ignoreSslError")]
        public Input<bool>? IgnoreSslError { get; set; }

        /// <summary>
        /// The ID of the project.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The Service Endpoint name.
        /// </summary>
        [Input("serviceEndpointName")]
        public Input<string>? ServiceEndpointName { get; set; }

        /// <summary>
        /// Octopus Server url.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        public ServiceendpointOctopusdeployState()
        {
        }
        public static new ServiceendpointOctopusdeployState Empty => new ServiceendpointOctopusdeployState();
    }
}
