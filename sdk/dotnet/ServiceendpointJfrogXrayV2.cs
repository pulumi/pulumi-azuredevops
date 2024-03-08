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
    /// Manages an JFrog XRay V2 server endpoint within an Azure DevOps organization.
    /// 
    /// &gt; **Note:** Using this service endpoint requires you to first install [JFrog Extension](https://marketplace.visualstudio.com/items?itemName=JFrog.jfrog-azure-devops-extension).
    /// 
    /// ## Example Usage
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
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
    ///     var exampleServiceendpointJfrogXrayV2 = new AzureDevOps.ServiceendpointJfrogXrayV2("exampleServiceendpointJfrogXrayV2", new()
    ///     {
    ///         ProjectId = exampleProject.Id,
    ///         ServiceEndpointName = "Example Artifactory",
    ///         Description = "Managed by Terraform",
    ///         Url = "https://artifactory.my.com",
    ///         AuthenticationToken = new AzureDevOps.Inputs.ServiceendpointJfrogXrayV2AuthenticationTokenArgs
    ///         {
    ///             Token = "0000000000000000000000000000000000000000",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// Alternatively a username and password may be used.
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
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
    ///     var exampleServiceendpointJfrogXrayV2 = new AzureDevOps.ServiceendpointJfrogXrayV2("exampleServiceendpointJfrogXrayV2", new()
    ///     {
    ///         ProjectId = exampleProject.Id,
    ///         ServiceEndpointName = "Example Artifactory",
    ///         Description = "Managed by Terraform",
    ///         Url = "https://artifactory.my.com",
    ///         AuthenticationBasic = new AzureDevOps.Inputs.ServiceendpointJfrogXrayV2AuthenticationBasicArgs
    ///         {
    ///             Username = "username",
    ///             Password = "password",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Relevant Links
    /// 
    /// * [Azure DevOps Service Connections](https://docs.microsoft.com/en-us/azure/devops/pipelines/library/service-endpoints?view=azure-devops&amp;tabs=yaml)
    /// * [Artifactory User Token](https://docs.artifactory.org/latest/user-guide/user-token/)
    /// 
    /// ## Import
    /// 
    /// Azure DevOps Service Endpoint JFrog XRay V2 can be imported using the **projectID/serviceEndpointID**, e.g.
    /// 
    /// ```sh
    /// $ pulumi import azuredevops:index/serviceendpointJfrogXrayV2:ServiceendpointJfrogXrayV2 example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
    /// ```
    /// </summary>
    [AzureDevOpsResourceType("azuredevops:index/serviceendpointJfrogXrayV2:ServiceendpointJfrogXrayV2")]
    public partial class ServiceendpointJfrogXrayV2 : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A `authentication_basic` block as documented below.
        /// </summary>
        [Output("authenticationBasic")]
        public Output<Outputs.ServiceendpointJfrogXrayV2AuthenticationBasic?> AuthenticationBasic { get; private set; } = null!;

        /// <summary>
        /// A `authentication_token` block as documented below.
        /// </summary>
        [Output("authenticationToken")]
        public Output<Outputs.ServiceendpointJfrogXrayV2AuthenticationToken?> AuthenticationToken { get; private set; } = null!;

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
        /// &gt; **NOTE:** URL should not end in a slash character.
        /// </summary>
        [Output("url")]
        public Output<string> Url { get; private set; } = null!;


        /// <summary>
        /// Create a ServiceendpointJfrogXrayV2 resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServiceendpointJfrogXrayV2(string name, ServiceendpointJfrogXrayV2Args args, CustomResourceOptions? options = null)
            : base("azuredevops:index/serviceendpointJfrogXrayV2:ServiceendpointJfrogXrayV2", name, args ?? new ServiceendpointJfrogXrayV2Args(), MakeResourceOptions(options, ""))
        {
        }

        private ServiceendpointJfrogXrayV2(string name, Input<string> id, ServiceendpointJfrogXrayV2State? state = null, CustomResourceOptions? options = null)
            : base("azuredevops:index/serviceendpointJfrogXrayV2:ServiceendpointJfrogXrayV2", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ServiceendpointJfrogXrayV2 resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServiceendpointJfrogXrayV2 Get(string name, Input<string> id, ServiceendpointJfrogXrayV2State? state = null, CustomResourceOptions? options = null)
        {
            return new ServiceendpointJfrogXrayV2(name, id, state, options);
        }
    }

    public sealed class ServiceendpointJfrogXrayV2Args : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A `authentication_basic` block as documented below.
        /// </summary>
        [Input("authenticationBasic")]
        public Input<Inputs.ServiceendpointJfrogXrayV2AuthenticationBasicArgs>? AuthenticationBasic { get; set; }

        /// <summary>
        /// A `authentication_token` block as documented below.
        /// </summary>
        [Input("authenticationToken")]
        public Input<Inputs.ServiceendpointJfrogXrayV2AuthenticationTokenArgs>? AuthenticationToken { get; set; }

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
        /// &gt; **NOTE:** URL should not end in a slash character.
        /// </summary>
        [Input("url", required: true)]
        public Input<string> Url { get; set; } = null!;

        public ServiceendpointJfrogXrayV2Args()
        {
        }
        public static new ServiceendpointJfrogXrayV2Args Empty => new ServiceendpointJfrogXrayV2Args();
    }

    public sealed class ServiceendpointJfrogXrayV2State : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A `authentication_basic` block as documented below.
        /// </summary>
        [Input("authenticationBasic")]
        public Input<Inputs.ServiceendpointJfrogXrayV2AuthenticationBasicGetArgs>? AuthenticationBasic { get; set; }

        /// <summary>
        /// A `authentication_token` block as documented below.
        /// </summary>
        [Input("authenticationToken")]
        public Input<Inputs.ServiceendpointJfrogXrayV2AuthenticationTokenGetArgs>? AuthenticationToken { get; set; }

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
        /// &gt; **NOTE:** URL should not end in a slash character.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        public ServiceendpointJfrogXrayV2State()
        {
        }
        public static new ServiceendpointJfrogXrayV2State Empty => new ServiceendpointJfrogXrayV2State();
    }
}
