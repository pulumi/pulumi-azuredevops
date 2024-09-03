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
    /// Manages a ArgoCD service endpoint within Azure DevOps. Using this service endpoint requires you to first install [Argo CD Extension](https://marketplace.visualstudio.com/items?itemName=scb-tomasmortensen.vsix-argocd).
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
    ///     });
    /// 
    ///     var exampleServiceendpointArgocd = new AzureDevOps.ServiceendpointArgocd("example", new()
    ///     {
    ///         ProjectId = example.Id,
    ///         ServiceEndpointName = "Example ArgoCD",
    ///         Description = "Managed by Terraform",
    ///         Url = "https://argocd.my.com",
    ///         AuthenticationToken = new AzureDevOps.Inputs.ServiceendpointArgocdAuthenticationTokenArgs
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
    ///     var exampleServiceendpointArgocd = new AzureDevOps.ServiceendpointArgocd("example", new()
    ///     {
    ///         ProjectId = example.Id,
    ///         ServiceEndpointName = "Example ArgoCD",
    ///         Description = "Managed by Terraform",
    ///         Url = "https://argocd.my.com",
    ///         AuthenticationBasic = new AzureDevOps.Inputs.ServiceendpointArgocdAuthenticationBasicArgs
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
    /// - [Azure DevOps Service Connections](https://docs.microsoft.com/en-us/azure/devops/pipelines/library/service-endpoints?view=azure-devops&amp;tabs=yaml)
    /// - [ArgoCD Project/User Token](https://argo-cd.readthedocs.io/en/stable/user-guide/commands/argocd_account_generate-token/)
    /// - [Argo CD Extension](https://marketplace.visualstudio.com/items?itemName=scb-tomasmortensen.vsix-argocd)
    /// 
    /// ## Import
    /// 
    /// Azure DevOps Service Endpoint ArgoCD can be imported using the **projectID/serviceEndpointID**, e.g.
    /// 
    /// ```sh
    /// $ pulumi import azuredevops:index/serviceendpointArgocd:ServiceendpointArgocd example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
    /// ```
    /// </summary>
    [AzureDevOpsResourceType("azuredevops:index/serviceendpointArgocd:ServiceendpointArgocd")]
    public partial class ServiceendpointArgocd : global::Pulumi.CustomResource
    {
        /// <summary>
        /// An `authentication_basic` block for the ArgoCD as documented below.
        /// 
        /// &gt; **NOTE:** `authentication_basic` and `authentication_token` conflict with each other, only one is required.
        /// </summary>
        [Output("authenticationBasic")]
        public Output<Outputs.ServiceendpointArgocdAuthenticationBasic?> AuthenticationBasic { get; private set; } = null!;

        /// <summary>
        /// An `authentication_token` block for the ArgoCD as documented below.
        /// </summary>
        [Output("authenticationToken")]
        public Output<Outputs.ServiceendpointArgocdAuthenticationToken?> AuthenticationToken { get; private set; } = null!;

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
        /// URL of the ArgoCD server to connect with.
        /// </summary>
        [Output("url")]
        public Output<string> Url { get; private set; } = null!;


        /// <summary>
        /// Create a ServiceendpointArgocd resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServiceendpointArgocd(string name, ServiceendpointArgocdArgs args, CustomResourceOptions? options = null)
            : base("azuredevops:index/serviceendpointArgocd:ServiceendpointArgocd", name, args ?? new ServiceendpointArgocdArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServiceendpointArgocd(string name, Input<string> id, ServiceendpointArgocdState? state = null, CustomResourceOptions? options = null)
            : base("azuredevops:index/serviceendpointArgocd:ServiceendpointArgocd", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ServiceendpointArgocd resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServiceendpointArgocd Get(string name, Input<string> id, ServiceendpointArgocdState? state = null, CustomResourceOptions? options = null)
        {
            return new ServiceendpointArgocd(name, id, state, options);
        }
    }

    public sealed class ServiceendpointArgocdArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// An `authentication_basic` block for the ArgoCD as documented below.
        /// 
        /// &gt; **NOTE:** `authentication_basic` and `authentication_token` conflict with each other, only one is required.
        /// </summary>
        [Input("authenticationBasic")]
        public Input<Inputs.ServiceendpointArgocdAuthenticationBasicArgs>? AuthenticationBasic { get; set; }

        /// <summary>
        /// An `authentication_token` block for the ArgoCD as documented below.
        /// </summary>
        [Input("authenticationToken")]
        public Input<Inputs.ServiceendpointArgocdAuthenticationTokenArgs>? AuthenticationToken { get; set; }

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
        /// URL of the ArgoCD server to connect with.
        /// </summary>
        [Input("url", required: true)]
        public Input<string> Url { get; set; } = null!;

        public ServiceendpointArgocdArgs()
        {
        }
        public static new ServiceendpointArgocdArgs Empty => new ServiceendpointArgocdArgs();
    }

    public sealed class ServiceendpointArgocdState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// An `authentication_basic` block for the ArgoCD as documented below.
        /// 
        /// &gt; **NOTE:** `authentication_basic` and `authentication_token` conflict with each other, only one is required.
        /// </summary>
        [Input("authenticationBasic")]
        public Input<Inputs.ServiceendpointArgocdAuthenticationBasicGetArgs>? AuthenticationBasic { get; set; }

        /// <summary>
        /// An `authentication_token` block for the ArgoCD as documented below.
        /// </summary>
        [Input("authenticationToken")]
        public Input<Inputs.ServiceendpointArgocdAuthenticationTokenGetArgs>? AuthenticationToken { get; set; }

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
        /// URL of the ArgoCD server to connect with.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        public ServiceendpointArgocdState()
        {
        }
        public static new ServiceendpointArgocdState Empty => new ServiceendpointArgocdState();
    }
}
