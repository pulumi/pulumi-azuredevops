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
    /// Manages an Artifactory server endpoint within an Azure DevOps organization. Using this service endpoint requires you to first install [JFrog Artifactory Extension](https://marketplace.visualstudio.com/items?itemName=JFrog.jfrog-artifactory-vsts-extension).
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
    ///         Description = "Managed by Terraform",
    ///     });
    /// 
    ///     var exampleServiceEndpointArtifactory = new AzureDevOps.ServiceEndpointArtifactory("example", new()
    ///     {
    ///         ProjectId = example.Id,
    ///         ServiceEndpointName = "Example Artifactory",
    ///         Description = "Managed by Terraform",
    ///         Url = "https://artifactory.my.com",
    ///         AuthenticationToken = new AzureDevOps.Inputs.ServiceEndpointArtifactoryAuthenticationTokenArgs
    ///         {
    ///             Token = "0000000000000000000000000000000000000000",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// Alternatively a username and password may be used.
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
    ///         Description = "Managed by Terraform",
    ///     });
    /// 
    ///     var exampleServiceEndpointArtifactory = new AzureDevOps.ServiceEndpointArtifactory("example", new()
    ///     {
    ///         ProjectId = example.Id,
    ///         ServiceEndpointName = "Example Artifactory",
    ///         Description = "Managed by Terraform",
    ///         Url = "https://artifactory.my.com",
    ///         AuthenticationBasic = new AzureDevOps.Inputs.ServiceEndpointArtifactoryAuthenticationBasicArgs
    ///         {
    ///             Username = "username",
    ///             Password = "password",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Relevant Links
    /// 
    /// * [Azure DevOps Service Connections](https://docs.microsoft.com/en-us/azure/devops/pipelines/library/service-endpoints?view=azure-devops&amp;tabs=yaml)
    /// * [Artifactory User Token](https://docs.artifactory.org/latest/user-guide/user-token/)
    /// 
    /// ## Import
    /// 
    /// Azure DevOps Service Endpoint Artifactory can be imported using the **projectID/serviceEndpointID**, e.g.
    /// 
    /// ```sh
    /// $ pulumi import azuredevops:index/serviceEndpointArtifactory:ServiceEndpointArtifactory example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
    /// ```
    /// </summary>
    [AzureDevOpsResourceType("azuredevops:index/serviceEndpointArtifactory:ServiceEndpointArtifactory")]
    public partial class ServiceEndpointArtifactory : global::Pulumi.CustomResource
    {
        [Output("authenticationBasic")]
        public Output<Outputs.ServiceEndpointArtifactoryAuthenticationBasic?> AuthenticationBasic { get; private set; } = null!;

        [Output("authenticationToken")]
        public Output<Outputs.ServiceEndpointArtifactoryAuthenticationToken?> AuthenticationToken { get; private set; } = null!;

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
        /// URL of the Artifactory server to connect with.
        /// 
        /// *Note: URL should not end in a slash character.*
        /// * either `authentication_token` or `authentication_basic` (one is required)
        /// </summary>
        [Output("url")]
        public Output<string> Url { get; private set; } = null!;


        /// <summary>
        /// Create a ServiceEndpointArtifactory resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServiceEndpointArtifactory(string name, ServiceEndpointArtifactoryArgs args, CustomResourceOptions? options = null)
            : base("azuredevops:index/serviceEndpointArtifactory:ServiceEndpointArtifactory", name, args ?? new ServiceEndpointArtifactoryArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServiceEndpointArtifactory(string name, Input<string> id, ServiceEndpointArtifactoryState? state = null, CustomResourceOptions? options = null)
            : base("azuredevops:index/serviceEndpointArtifactory:ServiceEndpointArtifactory", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ServiceEndpointArtifactory resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServiceEndpointArtifactory Get(string name, Input<string> id, ServiceEndpointArtifactoryState? state = null, CustomResourceOptions? options = null)
        {
            return new ServiceEndpointArtifactory(name, id, state, options);
        }
    }

    public sealed class ServiceEndpointArtifactoryArgs : global::Pulumi.ResourceArgs
    {
        [Input("authenticationBasic")]
        public Input<Inputs.ServiceEndpointArtifactoryAuthenticationBasicArgs>? AuthenticationBasic { get; set; }

        [Input("authenticationToken")]
        public Input<Inputs.ServiceEndpointArtifactoryAuthenticationTokenArgs>? AuthenticationToken { get; set; }

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
        /// URL of the Artifactory server to connect with.
        /// 
        /// *Note: URL should not end in a slash character.*
        /// * either `authentication_token` or `authentication_basic` (one is required)
        /// </summary>
        [Input("url", required: true)]
        public Input<string> Url { get; set; } = null!;

        public ServiceEndpointArtifactoryArgs()
        {
        }
        public static new ServiceEndpointArtifactoryArgs Empty => new ServiceEndpointArtifactoryArgs();
    }

    public sealed class ServiceEndpointArtifactoryState : global::Pulumi.ResourceArgs
    {
        [Input("authenticationBasic")]
        public Input<Inputs.ServiceEndpointArtifactoryAuthenticationBasicGetArgs>? AuthenticationBasic { get; set; }

        [Input("authenticationToken")]
        public Input<Inputs.ServiceEndpointArtifactoryAuthenticationTokenGetArgs>? AuthenticationToken { get; set; }

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
        /// URL of the Artifactory server to connect with.
        /// 
        /// *Note: URL should not end in a slash character.*
        /// * either `authentication_token` or `authentication_basic` (one is required)
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        public ServiceEndpointArtifactoryState()
        {
        }
        public static new ServiceEndpointArtifactoryState Empty => new ServiceEndpointArtifactoryState();
    }
}
