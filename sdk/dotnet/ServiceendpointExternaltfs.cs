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
    /// Manages an Azure Repository/Team Foundation Server service endpoint within Azure DevOps.
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
    ///     var exampleServiceendpointExternaltfs = new AzureDevOps.ServiceendpointExternaltfs("example", new()
    ///     {
    ///         ProjectId = example.Id,
    ///         ServiceEndpointName = "Example External TFS Name",
    ///         ConnectionUrl = "https://dev.azure.com/myorganization",
    ///         Description = "Managed by Pulumi",
    ///         AuthPersonal = new AzureDevOps.Inputs.ServiceendpointExternaltfsAuthPersonalArgs
    ///         {
    ///             PersonalAccessToken = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
    ///         },
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
    /// Azure DevOps Azure Repository/Team Foundation Server Service Endpoint can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**
    /// 
    /// ```sh
    /// $ pulumi import azuredevops:index/serviceendpointExternaltfs:ServiceendpointExternaltfs example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
    /// ```
    /// </summary>
    [AzureDevOpsResourceType("azuredevops:index/serviceendpointExternaltfs:ServiceendpointExternaltfs")]
    public partial class ServiceendpointExternaltfs : global::Pulumi.CustomResource
    {
        /// <summary>
        /// An `auth_personal` block as documented below. Allows connecting using a personal access token.
        /// </summary>
        [Output("authPersonal")]
        public Output<Outputs.ServiceendpointExternaltfsAuthPersonal> AuthPersonal { get; private set; } = null!;

        [Output("authorization")]
        public Output<ImmutableDictionary<string, string>> Authorization { get; private set; } = null!;

        /// <summary>
        /// Azure DevOps Organization or TFS Project Collection Url.
        /// </summary>
        [Output("connectionUrl")]
        public Output<string> ConnectionUrl { get; private set; } = null!;

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
        /// Create a ServiceendpointExternaltfs resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServiceendpointExternaltfs(string name, ServiceendpointExternaltfsArgs args, CustomResourceOptions? options = null)
            : base("azuredevops:index/serviceendpointExternaltfs:ServiceendpointExternaltfs", name, args ?? new ServiceendpointExternaltfsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServiceendpointExternaltfs(string name, Input<string> id, ServiceendpointExternaltfsState? state = null, CustomResourceOptions? options = null)
            : base("azuredevops:index/serviceendpointExternaltfs:ServiceendpointExternaltfs", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ServiceendpointExternaltfs resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServiceendpointExternaltfs Get(string name, Input<string> id, ServiceendpointExternaltfsState? state = null, CustomResourceOptions? options = null)
        {
            return new ServiceendpointExternaltfs(name, id, state, options);
        }
    }

    public sealed class ServiceendpointExternaltfsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// An `auth_personal` block as documented below. Allows connecting using a personal access token.
        /// </summary>
        [Input("authPersonal", required: true)]
        public Input<Inputs.ServiceendpointExternaltfsAuthPersonalArgs> AuthPersonal { get; set; } = null!;

        /// <summary>
        /// Azure DevOps Organization or TFS Project Collection Url.
        /// </summary>
        [Input("connectionUrl", required: true)]
        public Input<string> ConnectionUrl { get; set; } = null!;

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

        public ServiceendpointExternaltfsArgs()
        {
        }
        public static new ServiceendpointExternaltfsArgs Empty => new ServiceendpointExternaltfsArgs();
    }

    public sealed class ServiceendpointExternaltfsState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// An `auth_personal` block as documented below. Allows connecting using a personal access token.
        /// </summary>
        [Input("authPersonal")]
        public Input<Inputs.ServiceendpointExternaltfsAuthPersonalGetArgs>? AuthPersonal { get; set; }

        [Input("authorization")]
        private InputMap<string>? _authorization;
        public InputMap<string> Authorization
        {
            get => _authorization ?? (_authorization = new InputMap<string>());
            set => _authorization = value;
        }

        /// <summary>
        /// Azure DevOps Organization or TFS Project Collection Url.
        /// </summary>
        [Input("connectionUrl")]
        public Input<string>? ConnectionUrl { get; set; }

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

        public ServiceendpointExternaltfsState()
        {
        }
        public static new ServiceendpointExternaltfsState Empty => new ServiceendpointExternaltfsState();
    }
}
