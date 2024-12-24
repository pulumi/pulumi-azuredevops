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
    /// Manages a Snyk Security Scan service endpoint within Azure DevOps. Using this service endpoint requires you to install: [Snyk Security Scan](https://marketplace.visualstudio.com/items?itemName=Snyk.snyk-security-scan)
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
    ///     var exampleServiceendpointSnyk = new AzureDevOps.ServiceendpointSnyk("example", new()
    ///     {
    ///         ProjectId = example.Id,
    ///         ServerUrl = "https://snyk.io/",
    ///         ApiToken = "00000000-0000-0000-0000-000000000000",
    ///         ServiceEndpointName = "Example Snyk",
    ///         Description = "Managed by Pulumi",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Relevant Links
    /// 
    /// - [Azure DevOps Service REST API 7.0 - Endpoints](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-7.0)
    /// 
    /// ## Import
    /// 
    /// Azure DevOps Service Endpoint Snyk can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**
    /// 
    /// ```sh
    /// $ pulumi import azuredevops:index/serviceendpointSnyk:ServiceendpointSnyk example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
    /// ```
    /// </summary>
    [AzureDevOpsResourceType("azuredevops:index/serviceendpointSnyk:ServiceendpointSnyk")]
    public partial class ServiceendpointSnyk : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The API token of the Snyk Security Scan.
        /// </summary>
        [Output("apiToken")]
        public Output<string> ApiToken { get; private set; } = null!;

        [Output("authorization")]
        public Output<ImmutableDictionary<string, string>> Authorization { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The ID of the project.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// The server URL of the Snyk Security Scan.
        /// </summary>
        [Output("serverUrl")]
        public Output<string> ServerUrl { get; private set; } = null!;

        /// <summary>
        /// The Service Endpoint name.
        /// </summary>
        [Output("serviceEndpointName")]
        public Output<string> ServiceEndpointName { get; private set; } = null!;


        /// <summary>
        /// Create a ServiceendpointSnyk resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServiceendpointSnyk(string name, ServiceendpointSnykArgs args, CustomResourceOptions? options = null)
            : base("azuredevops:index/serviceendpointSnyk:ServiceendpointSnyk", name, args ?? new ServiceendpointSnykArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServiceendpointSnyk(string name, Input<string> id, ServiceendpointSnykState? state = null, CustomResourceOptions? options = null)
            : base("azuredevops:index/serviceendpointSnyk:ServiceendpointSnyk", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "apiToken",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ServiceendpointSnyk resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServiceendpointSnyk Get(string name, Input<string> id, ServiceendpointSnykState? state = null, CustomResourceOptions? options = null)
        {
            return new ServiceendpointSnyk(name, id, state, options);
        }
    }

    public sealed class ServiceendpointSnykArgs : global::Pulumi.ResourceArgs
    {
        [Input("apiToken", required: true)]
        private Input<string>? _apiToken;

        /// <summary>
        /// The API token of the Snyk Security Scan.
        /// </summary>
        public Input<string>? ApiToken
        {
            get => _apiToken;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _apiToken = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

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
        /// The ID of the project.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        /// <summary>
        /// The server URL of the Snyk Security Scan.
        /// </summary>
        [Input("serverUrl", required: true)]
        public Input<string> ServerUrl { get; set; } = null!;

        /// <summary>
        /// The Service Endpoint name.
        /// </summary>
        [Input("serviceEndpointName", required: true)]
        public Input<string> ServiceEndpointName { get; set; } = null!;

        public ServiceendpointSnykArgs()
        {
        }
        public static new ServiceendpointSnykArgs Empty => new ServiceendpointSnykArgs();
    }

    public sealed class ServiceendpointSnykState : global::Pulumi.ResourceArgs
    {
        [Input("apiToken")]
        private Input<string>? _apiToken;

        /// <summary>
        /// The API token of the Snyk Security Scan.
        /// </summary>
        public Input<string>? ApiToken
        {
            get => _apiToken;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _apiToken = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

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
        /// The ID of the project.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The server URL of the Snyk Security Scan.
        /// </summary>
        [Input("serverUrl")]
        public Input<string>? ServerUrl { get; set; }

        /// <summary>
        /// The Service Endpoint name.
        /// </summary>
        [Input("serviceEndpointName")]
        public Input<string>? ServiceEndpointName { get; set; }

        public ServiceendpointSnykState()
        {
        }
        public static new ServiceendpointSnykState Empty => new ServiceendpointSnykState();
    }
}
