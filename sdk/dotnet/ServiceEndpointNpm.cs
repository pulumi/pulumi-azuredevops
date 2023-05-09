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
    /// Manages a npm service endpoint within Azure DevOps.
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
    ///     var exampleServiceEndpointNpm = new AzureDevOps.ServiceEndpointNpm("exampleServiceEndpointNpm", new()
    ///     {
    ///         ProjectId = exampleProject.Id,
    ///         ServiceEndpointName = "Example npm",
    ///         Url = "https://registry.npmjs.org",
    ///         AccessToken = "00000000-0000-0000-0000-000000000000",
    ///         Description = "Managed by Terraform",
    ///     });
    /// 
    /// });
    /// ```
    /// ## Relevant Links
    /// 
    /// - [Azure DevOps Service REST API 6.0 - Endpoints](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-6.0)
    /// - [Azure DevOps Service Connections](https://docs.microsoft.com/en-us/azure/devops/pipelines/library/service-endpoints?view=azure-devops&amp;tabs=yaml)
    /// - [npm User Token](https://docs.npmjs.com/about-access-tokens)
    /// 
    /// ## Import
    /// 
    /// Azure DevOps Service Endpoint npm can be imported using the **projectID/serviceEndpointID**, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azuredevops:index/serviceEndpointNpm:ServiceEndpointNpm example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
    /// ```
    /// </summary>
    [AzureDevOpsResourceType("azuredevops:index/serviceEndpointNpm:ServiceEndpointNpm")]
    public partial class ServiceEndpointNpm : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The access token for npm registry.
        /// </summary>
        [Output("accessToken")]
        public Output<string> AccessToken { get; private set; } = null!;

        [Output("authorization")]
        public Output<ImmutableDictionary<string, string>> Authorization { get; private set; } = null!;

        /// <summary>
        /// The Service Endpoint description.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

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
        /// URL of the npm registry to connect with.
        /// </summary>
        [Output("url")]
        public Output<string> Url { get; private set; } = null!;


        /// <summary>
        /// Create a ServiceEndpointNpm resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServiceEndpointNpm(string name, ServiceEndpointNpmArgs args, CustomResourceOptions? options = null)
            : base("azuredevops:index/serviceEndpointNpm:ServiceEndpointNpm", name, args ?? new ServiceEndpointNpmArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServiceEndpointNpm(string name, Input<string> id, ServiceEndpointNpmState? state = null, CustomResourceOptions? options = null)
            : base("azuredevops:index/serviceEndpointNpm:ServiceEndpointNpm", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "accessToken",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ServiceEndpointNpm resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServiceEndpointNpm Get(string name, Input<string> id, ServiceEndpointNpmState? state = null, CustomResourceOptions? options = null)
        {
            return new ServiceEndpointNpm(name, id, state, options);
        }
    }

    public sealed class ServiceEndpointNpmArgs : global::Pulumi.ResourceArgs
    {
        [Input("accessToken", required: true)]
        private Input<string>? _accessToken;

        /// <summary>
        /// The access token for npm registry.
        /// </summary>
        public Input<string>? AccessToken
        {
            get => _accessToken;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _accessToken = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("authorization")]
        private InputMap<string>? _authorization;
        public InputMap<string> Authorization
        {
            get => _authorization ?? (_authorization = new InputMap<string>());
            set => _authorization = value;
        }

        /// <summary>
        /// The Service Endpoint description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

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
        /// URL of the npm registry to connect with.
        /// </summary>
        [Input("url", required: true)]
        public Input<string> Url { get; set; } = null!;

        public ServiceEndpointNpmArgs()
        {
        }
        public static new ServiceEndpointNpmArgs Empty => new ServiceEndpointNpmArgs();
    }

    public sealed class ServiceEndpointNpmState : global::Pulumi.ResourceArgs
    {
        [Input("accessToken")]
        private Input<string>? _accessToken;

        /// <summary>
        /// The access token for npm registry.
        /// </summary>
        public Input<string>? AccessToken
        {
            get => _accessToken;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _accessToken = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("authorization")]
        private InputMap<string>? _authorization;
        public InputMap<string> Authorization
        {
            get => _authorization ?? (_authorization = new InputMap<string>());
            set => _authorization = value;
        }

        /// <summary>
        /// The Service Endpoint description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

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
        /// URL of the npm registry to connect with.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        public ServiceEndpointNpmState()
        {
        }
        public static new ServiceEndpointNpmState Empty => new ServiceEndpointNpmState();
    }
}
