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
    /// Manages a GitHub service endpoint within Azure DevOps.
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
    ///     var exampleServiceEndpointGitHub = new AzureDevOps.ServiceEndpointGitHub("exampleServiceEndpointGitHub", new()
    ///     {
    ///         ProjectId = exampleProject.Id,
    ///         ServiceEndpointName = "Example GitHub Personal Access Token",
    ///         AuthPersonal = new AzureDevOps.Inputs.ServiceEndpointGitHubAuthPersonalArgs
    ///         {
    ///             PersonalAccessToken = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
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
    ///     var exampleServiceEndpointGitHub = new AzureDevOps.ServiceEndpointGitHub("exampleServiceEndpointGitHub", new()
    ///     {
    ///         ProjectId = exampleProject.Id,
    ///         ServiceEndpointName = "Example GitHub",
    ///         AuthOauth = new AzureDevOps.Inputs.ServiceEndpointGitHubAuthOauthArgs
    ///         {
    ///             OauthConfigurationId = "00000000-0000-0000-0000-000000000000",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
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
    ///     var exampleServiceEndpointGitHub = new AzureDevOps.ServiceEndpointGitHub("exampleServiceEndpointGitHub", new()
    ///     {
    ///         ProjectId = exampleProject.Id,
    ///         ServiceEndpointName = "Example GitHub Apps: Azure Pipelines",
    ///         Description = "Managed by Terraform",
    ///     });
    /// 
    /// });
    /// ```
    /// ## Relevant Links
    /// 
    /// - [Azure DevOps Service REST API 7.0 - Service Endpoints](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-7.0)
    /// 
    /// ## Import
    /// 
    /// Azure DevOps Service Endpoint GitHub can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**
    /// 
    /// ```sh
    ///  $ pulumi import azuredevops:index/serviceEndpointGitHub:ServiceEndpointGitHub example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
    /// ```
    /// </summary>
    [AzureDevOpsResourceType("azuredevops:index/serviceEndpointGitHub:ServiceEndpointGitHub")]
    public partial class ServiceEndpointGitHub : global::Pulumi.CustomResource
    {
        [Output("authOauth")]
        public Output<Outputs.ServiceEndpointGitHubAuthOauth?> AuthOauth { get; private set; } = null!;

        /// <summary>
        /// An `auth_personal` block as documented below. Allows connecting using a personal access token.
        /// </summary>
        [Output("authPersonal")]
        public Output<Outputs.ServiceEndpointGitHubAuthPersonal?> AuthPersonal { get; private set; } = null!;

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
        /// The Service Endpoint name.
        /// </summary>
        [Output("serviceEndpointName")]
        public Output<string> ServiceEndpointName { get; private set; } = null!;


        /// <summary>
        /// Create a ServiceEndpointGitHub resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServiceEndpointGitHub(string name, ServiceEndpointGitHubArgs args, CustomResourceOptions? options = null)
            : base("azuredevops:index/serviceEndpointGitHub:ServiceEndpointGitHub", name, args ?? new ServiceEndpointGitHubArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServiceEndpointGitHub(string name, Input<string> id, ServiceEndpointGitHubState? state = null, CustomResourceOptions? options = null)
            : base("azuredevops:index/serviceEndpointGitHub:ServiceEndpointGitHub", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azuredevops:ServiceEndpoint/gitHub:GitHub"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ServiceEndpointGitHub resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServiceEndpointGitHub Get(string name, Input<string> id, ServiceEndpointGitHubState? state = null, CustomResourceOptions? options = null)
        {
            return new ServiceEndpointGitHub(name, id, state, options);
        }
    }

    public sealed class ServiceEndpointGitHubArgs : global::Pulumi.ResourceArgs
    {
        [Input("authOauth")]
        public Input<Inputs.ServiceEndpointGitHubAuthOauthArgs>? AuthOauth { get; set; }

        /// <summary>
        /// An `auth_personal` block as documented below. Allows connecting using a personal access token.
        /// </summary>
        [Input("authPersonal")]
        public Input<Inputs.ServiceEndpointGitHubAuthPersonalArgs>? AuthPersonal { get; set; }

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
        /// The Service Endpoint name.
        /// </summary>
        [Input("serviceEndpointName", required: true)]
        public Input<string> ServiceEndpointName { get; set; } = null!;

        public ServiceEndpointGitHubArgs()
        {
        }
        public static new ServiceEndpointGitHubArgs Empty => new ServiceEndpointGitHubArgs();
    }

    public sealed class ServiceEndpointGitHubState : global::Pulumi.ResourceArgs
    {
        [Input("authOauth")]
        public Input<Inputs.ServiceEndpointGitHubAuthOauthGetArgs>? AuthOauth { get; set; }

        /// <summary>
        /// An `auth_personal` block as documented below. Allows connecting using a personal access token.
        /// </summary>
        [Input("authPersonal")]
        public Input<Inputs.ServiceEndpointGitHubAuthPersonalGetArgs>? AuthPersonal { get; set; }

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
        /// The Service Endpoint name.
        /// </summary>
        [Input("serviceEndpointName")]
        public Input<string>? ServiceEndpointName { get; set; }

        public ServiceEndpointGitHubState()
        {
        }
        public static new ServiceEndpointGitHubState Empty => new ServiceEndpointGitHubState();
    }
}
