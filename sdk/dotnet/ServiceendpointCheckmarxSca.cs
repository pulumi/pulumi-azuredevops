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
    /// Manages a Checkmarx SCA service endpoint within Azure DevOps. Using this service endpoint requires you to install: [Checkmarx SAST](https://marketplace.visualstudio.com/items?itemName=checkmarx.cxsast)
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
    ///     var exampleServiceendpointCheckmarxSca = new AzureDevOps.ServiceendpointCheckmarxSca("example", new()
    ///     {
    ///         ProjectId = example.Id,
    ///         ServiceEndpointName = "Example Checkmarx SCA",
    ///         AccessControlUrl = "https://accesscontrol.com",
    ///         ServerUrl = "https://server.com",
    ///         WebAppUrl = "https://webapp.com",
    ///         Account = "account",
    ///         Username = "username",
    ///         Password = "password",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Relevant Links
    /// 
    /// - [Azure DevOps Service REST API 7.0 - Service Endpoints](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-7.0)
    /// 
    /// ## Import
    /// 
    /// Azure DevOps Service Endpoint Check Marx SCA can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**
    /// 
    /// ```sh
    /// $ pulumi import azuredevops:index/serviceendpointCheckmarxSca:ServiceendpointCheckmarxSca example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
    /// ```
    /// </summary>
    [AzureDevOpsResourceType("azuredevops:index/serviceendpointCheckmarxSca:ServiceendpointCheckmarxSca")]
    public partial class ServiceendpointCheckmarxSca : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Access Control URL of the Checkmarx SCA.
        /// </summary>
        [Output("accessControlUrl")]
        public Output<string> AccessControlUrl { get; private set; } = null!;

        /// <summary>
        /// The account of the Checkmarx SCA.
        /// </summary>
        [Output("account")]
        public Output<string> Account { get; private set; } = null!;

        [Output("authorization")]
        public Output<ImmutableDictionary<string, string>> Authorization { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The password of the Checkmarx SCA.
        /// </summary>
        [Output("password")]
        public Output<string> Password { get; private set; } = null!;

        /// <summary>
        /// The ID of the project.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// The Server URL of the Checkmarx SCA.
        /// </summary>
        [Output("serverUrl")]
        public Output<string> ServerUrl { get; private set; } = null!;

        /// <summary>
        /// The Service Endpoint name.
        /// </summary>
        [Output("serviceEndpointName")]
        public Output<string> ServiceEndpointName { get; private set; } = null!;

        /// <summary>
        /// The full team name of the Checkmarx.
        /// </summary>
        [Output("team")]
        public Output<string?> Team { get; private set; } = null!;

        /// <summary>
        /// The username of the Checkmarx SCA.
        /// </summary>
        [Output("username")]
        public Output<string> Username { get; private set; } = null!;

        /// <summary>
        /// The Web App URL of the Checkmarx SCA.
        /// </summary>
        [Output("webAppUrl")]
        public Output<string> WebAppUrl { get; private set; } = null!;


        /// <summary>
        /// Create a ServiceendpointCheckmarxSca resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServiceendpointCheckmarxSca(string name, ServiceendpointCheckmarxScaArgs args, CustomResourceOptions? options = null)
            : base("azuredevops:index/serviceendpointCheckmarxSca:ServiceendpointCheckmarxSca", name, args ?? new ServiceendpointCheckmarxScaArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServiceendpointCheckmarxSca(string name, Input<string> id, ServiceendpointCheckmarxScaState? state = null, CustomResourceOptions? options = null)
            : base("azuredevops:index/serviceendpointCheckmarxSca:ServiceendpointCheckmarxSca", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ServiceendpointCheckmarxSca resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServiceendpointCheckmarxSca Get(string name, Input<string> id, ServiceendpointCheckmarxScaState? state = null, CustomResourceOptions? options = null)
        {
            return new ServiceendpointCheckmarxSca(name, id, state, options);
        }
    }

    public sealed class ServiceendpointCheckmarxScaArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Access Control URL of the Checkmarx SCA.
        /// </summary>
        [Input("accessControlUrl", required: true)]
        public Input<string> AccessControlUrl { get; set; } = null!;

        /// <summary>
        /// The account of the Checkmarx SCA.
        /// </summary>
        [Input("account", required: true)]
        public Input<string> Account { get; set; } = null!;

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
        /// The password of the Checkmarx SCA.
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
        /// The Server URL of the Checkmarx SCA.
        /// </summary>
        [Input("serverUrl", required: true)]
        public Input<string> ServerUrl { get; set; } = null!;

        /// <summary>
        /// The Service Endpoint name.
        /// </summary>
        [Input("serviceEndpointName", required: true)]
        public Input<string> ServiceEndpointName { get; set; } = null!;

        /// <summary>
        /// The full team name of the Checkmarx.
        /// </summary>
        [Input("team")]
        public Input<string>? Team { get; set; }

        /// <summary>
        /// The username of the Checkmarx SCA.
        /// </summary>
        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        /// <summary>
        /// The Web App URL of the Checkmarx SCA.
        /// </summary>
        [Input("webAppUrl", required: true)]
        public Input<string> WebAppUrl { get; set; } = null!;

        public ServiceendpointCheckmarxScaArgs()
        {
        }
        public static new ServiceendpointCheckmarxScaArgs Empty => new ServiceendpointCheckmarxScaArgs();
    }

    public sealed class ServiceendpointCheckmarxScaState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Access Control URL of the Checkmarx SCA.
        /// </summary>
        [Input("accessControlUrl")]
        public Input<string>? AccessControlUrl { get; set; }

        /// <summary>
        /// The account of the Checkmarx SCA.
        /// </summary>
        [Input("account")]
        public Input<string>? Account { get; set; }

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
        /// The password of the Checkmarx SCA.
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
        /// The Server URL of the Checkmarx SCA.
        /// </summary>
        [Input("serverUrl")]
        public Input<string>? ServerUrl { get; set; }

        /// <summary>
        /// The Service Endpoint name.
        /// </summary>
        [Input("serviceEndpointName")]
        public Input<string>? ServiceEndpointName { get; set; }

        /// <summary>
        /// The full team name of the Checkmarx.
        /// </summary>
        [Input("team")]
        public Input<string>? Team { get; set; }

        /// <summary>
        /// The username of the Checkmarx SCA.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        /// <summary>
        /// The Web App URL of the Checkmarx SCA.
        /// </summary>
        [Input("webAppUrl")]
        public Input<string>? WebAppUrl { get; set; }

        public ServiceendpointCheckmarxScaState()
        {
        }
        public static new ServiceendpointCheckmarxScaState Empty => new ServiceendpointCheckmarxScaState();
    }
}
