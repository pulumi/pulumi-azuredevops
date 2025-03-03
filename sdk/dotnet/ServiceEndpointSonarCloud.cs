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
    /// Manages a SonarQube Cloud service endpoint within Azure DevOps.
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
    ///     var exampleServiceEndpointSonarCloud = new AzureDevOps.ServiceEndpointSonarCloud("example", new()
    ///     {
    ///         ProjectId = example.Id,
    ///         ServiceEndpointName = "Example SonarCloud",
    ///         Token = "0000000000000000000000000000000000000000",
    ///         Description = "Managed by Pulumi",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Relevant Links
    /// 
    /// - [Azure DevOps Service REST API 7.0 - Endpoints](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-7.0)
    /// - [Azure DevOps Service Connections](https://docs.microsoft.com/en-us/azure/devops/pipelines/library/service-endpoints?view=azure-devops&amp;tabs=yaml)
    /// - [SonarCloud User Token](https://docs.sonarcloud.io/advanced-setup/user-accounts/)
    /// 
    /// ## Import
    /// 
    /// Azure DevOps SonarQube Cloud Service Endpoint can be imported using the **projectID/serviceEndpointID**, e.g.
    /// 
    /// ```sh
    /// $ pulumi import azuredevops:index/serviceEndpointSonarCloud:ServiceEndpointSonarCloud example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
    /// ```
    /// </summary>
    [AzureDevOpsResourceType("azuredevops:index/serviceEndpointSonarCloud:ServiceEndpointSonarCloud")]
    public partial class ServiceEndpointSonarCloud : global::Pulumi.CustomResource
    {
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
        /// The Authentication Token generated through SonarCloud (go to `My Account &gt; Security &gt; Generate Tokens`).
        /// </summary>
        [Output("token")]
        public Output<string> Token { get; private set; } = null!;


        /// <summary>
        /// Create a ServiceEndpointSonarCloud resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServiceEndpointSonarCloud(string name, ServiceEndpointSonarCloudArgs args, CustomResourceOptions? options = null)
            : base("azuredevops:index/serviceEndpointSonarCloud:ServiceEndpointSonarCloud", name, args ?? new ServiceEndpointSonarCloudArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServiceEndpointSonarCloud(string name, Input<string> id, ServiceEndpointSonarCloudState? state = null, CustomResourceOptions? options = null)
            : base("azuredevops:index/serviceEndpointSonarCloud:ServiceEndpointSonarCloud", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "token",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ServiceEndpointSonarCloud resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServiceEndpointSonarCloud Get(string name, Input<string> id, ServiceEndpointSonarCloudState? state = null, CustomResourceOptions? options = null)
        {
            return new ServiceEndpointSonarCloud(name, id, state, options);
        }
    }

    public sealed class ServiceEndpointSonarCloudArgs : global::Pulumi.ResourceArgs
    {
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

        [Input("token", required: true)]
        private Input<string>? _token;

        /// <summary>
        /// The Authentication Token generated through SonarCloud (go to `My Account &gt; Security &gt; Generate Tokens`).
        /// </summary>
        public Input<string>? Token
        {
            get => _token;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _token = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        public ServiceEndpointSonarCloudArgs()
        {
        }
        public static new ServiceEndpointSonarCloudArgs Empty => new ServiceEndpointSonarCloudArgs();
    }

    public sealed class ServiceEndpointSonarCloudState : global::Pulumi.ResourceArgs
    {
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

        [Input("token")]
        private Input<string>? _token;

        /// <summary>
        /// The Authentication Token generated through SonarCloud (go to `My Account &gt; Security &gt; Generate Tokens`).
        /// </summary>
        public Input<string>? Token
        {
            get => _token;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _token = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        public ServiceEndpointSonarCloudState()
        {
        }
        public static new ServiceEndpointSonarCloudState Empty => new ServiceEndpointSonarCloudState();
    }
}
