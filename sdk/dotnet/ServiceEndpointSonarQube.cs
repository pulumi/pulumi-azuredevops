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
    /// Manages a SonarQube service endpoint within Azure DevOps.
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
    ///     var exampleServiceEndpointSonarQube = new AzureDevOps.ServiceEndpointSonarQube("exampleServiceEndpointSonarQube", new()
    ///     {
    ///         ProjectId = exampleProject.Id,
    ///         ServiceEndpointName = "Example SonarQube",
    ///         Url = "https://sonarqube.my.com",
    ///         Token = "0000000000000000000000000000000000000000",
    ///         Description = "Managed by Terraform",
    ///     });
    /// 
    /// });
    /// ```
    /// ## Relevant Links
    /// 
    /// - [Azure DevOps Service REST API 6.0 - Endpoints](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-6.0)
    /// - [Azure DevOps Service Connections](https://docs.microsoft.com/en-us/azure/devops/pipelines/library/service-endpoints?view=azure-devops&amp;tabs=yaml)
    /// - [SonarQube User Token](https://docs.sonarqube.org/latest/user-guide/user-token/)
    /// 
    /// ## Import
    /// 
    /// Azure DevOps Service Endpoint SonarQube can be imported using the **projectID/serviceEndpointID**, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azuredevops:index/serviceEndpointSonarQube:ServiceEndpointSonarQube example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
    /// ```
    /// </summary>
    [AzureDevOpsResourceType("azuredevops:index/serviceEndpointSonarQube:ServiceEndpointSonarQube")]
    public partial class ServiceEndpointSonarQube : global::Pulumi.CustomResource
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
        /// Authentication Token generated through SonarQube (go to My Account &gt; Security &gt; Generate Tokens).
        /// </summary>
        [Output("token")]
        public Output<string> Token { get; private set; } = null!;

        /// <summary>
        /// URL of the SonarQube server to connect with.
        /// </summary>
        [Output("url")]
        public Output<string> Url { get; private set; } = null!;


        /// <summary>
        /// Create a ServiceEndpointSonarQube resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServiceEndpointSonarQube(string name, ServiceEndpointSonarQubeArgs args, CustomResourceOptions? options = null)
            : base("azuredevops:index/serviceEndpointSonarQube:ServiceEndpointSonarQube", name, args ?? new ServiceEndpointSonarQubeArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServiceEndpointSonarQube(string name, Input<string> id, ServiceEndpointSonarQubeState? state = null, CustomResourceOptions? options = null)
            : base("azuredevops:index/serviceEndpointSonarQube:ServiceEndpointSonarQube", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ServiceEndpointSonarQube resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServiceEndpointSonarQube Get(string name, Input<string> id, ServiceEndpointSonarQubeState? state = null, CustomResourceOptions? options = null)
        {
            return new ServiceEndpointSonarQube(name, id, state, options);
        }
    }

    public sealed class ServiceEndpointSonarQubeArgs : global::Pulumi.ResourceArgs
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
        /// Authentication Token generated through SonarQube (go to My Account &gt; Security &gt; Generate Tokens).
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

        /// <summary>
        /// URL of the SonarQube server to connect with.
        /// </summary>
        [Input("url", required: true)]
        public Input<string> Url { get; set; } = null!;

        public ServiceEndpointSonarQubeArgs()
        {
        }
        public static new ServiceEndpointSonarQubeArgs Empty => new ServiceEndpointSonarQubeArgs();
    }

    public sealed class ServiceEndpointSonarQubeState : global::Pulumi.ResourceArgs
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
        /// Authentication Token generated through SonarQube (go to My Account &gt; Security &gt; Generate Tokens).
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

        /// <summary>
        /// URL of the SonarQube server to connect with.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        public ServiceEndpointSonarQubeState()
        {
        }
        public static new ServiceEndpointSonarQubeState Empty => new ServiceEndpointSonarQubeState();
    }
}
