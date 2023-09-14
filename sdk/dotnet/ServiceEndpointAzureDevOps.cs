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
    /// Manages an Azure DevOps service endpoint within Azure DevOps.
    /// 
    /// &gt; **Note** This resource is duplicate with `azuredevops.ServiceEndpointPipeline`,  will be removed in the future, use `azuredevops.ServiceEndpointPipeline` instead.
    /// 
    /// &gt; **Note** Prerequisite: Extension [Configurable Pipeline Runner](https://marketplace.visualstudio.com/items?itemName=CSE-DevOps.RunPipelines) has been installed for the organization.
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
    ///     var exampleServiceEndpointAzureDevOps = new AzureDevOps.ServiceEndpointAzureDevOps("exampleServiceEndpointAzureDevOps", new()
    ///     {
    ///         ProjectId = exampleProject.Id,
    ///         ServiceEndpointName = "Example Azure DevOps",
    ///         OrgUrl = "https://dev.azure.com/testorganization",
    ///         ReleaseApiUrl = "https://vsrm.dev.azure.com/testorganization",
    ///         PersonalAccessToken = "0000000000000000000000000000000000000000000000000000",
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
    /// Azure DevOps Service Endpoint Azure DevOps can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**
    /// 
    /// ```sh
    ///  $ pulumi import azuredevops:index/serviceEndpointAzureDevOps:ServiceEndpointAzureDevOps example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
    /// ```
    /// </summary>
    [AzureDevOpsResourceType("azuredevops:index/serviceEndpointAzureDevOps:ServiceEndpointAzureDevOps")]
    public partial class ServiceEndpointAzureDevOps : global::Pulumi.CustomResource
    {
        [Output("authorization")]
        public Output<ImmutableDictionary<string, string>> Authorization { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The organization URL.
        /// </summary>
        [Output("orgUrl")]
        public Output<string> OrgUrl { get; private set; } = null!;

        /// <summary>
        /// The Azure DevOps personal access token.
        /// </summary>
        [Output("personalAccessToken")]
        public Output<string> PersonalAccessToken { get; private set; } = null!;

        /// <summary>
        /// The ID of the project.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// The URL of the release API.
        /// </summary>
        [Output("releaseApiUrl")]
        public Output<string> ReleaseApiUrl { get; private set; } = null!;

        /// <summary>
        /// The Service Endpoint name.
        /// </summary>
        [Output("serviceEndpointName")]
        public Output<string> ServiceEndpointName { get; private set; } = null!;


        /// <summary>
        /// Create a ServiceEndpointAzureDevOps resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServiceEndpointAzureDevOps(string name, ServiceEndpointAzureDevOpsArgs args, CustomResourceOptions? options = null)
            : base("azuredevops:index/serviceEndpointAzureDevOps:ServiceEndpointAzureDevOps", name, args ?? new ServiceEndpointAzureDevOpsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServiceEndpointAzureDevOps(string name, Input<string> id, ServiceEndpointAzureDevOpsState? state = null, CustomResourceOptions? options = null)
            : base("azuredevops:index/serviceEndpointAzureDevOps:ServiceEndpointAzureDevOps", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "personalAccessToken",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ServiceEndpointAzureDevOps resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServiceEndpointAzureDevOps Get(string name, Input<string> id, ServiceEndpointAzureDevOpsState? state = null, CustomResourceOptions? options = null)
        {
            return new ServiceEndpointAzureDevOps(name, id, state, options);
        }
    }

    public sealed class ServiceEndpointAzureDevOpsArgs : global::Pulumi.ResourceArgs
    {
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
        /// The organization URL.
        /// </summary>
        [Input("orgUrl", required: true)]
        public Input<string> OrgUrl { get; set; } = null!;

        [Input("personalAccessToken", required: true)]
        private Input<string>? _personalAccessToken;

        /// <summary>
        /// The Azure DevOps personal access token.
        /// </summary>
        public Input<string>? PersonalAccessToken
        {
            get => _personalAccessToken;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _personalAccessToken = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The ID of the project.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        /// <summary>
        /// The URL of the release API.
        /// </summary>
        [Input("releaseApiUrl", required: true)]
        public Input<string> ReleaseApiUrl { get; set; } = null!;

        /// <summary>
        /// The Service Endpoint name.
        /// </summary>
        [Input("serviceEndpointName", required: true)]
        public Input<string> ServiceEndpointName { get; set; } = null!;

        public ServiceEndpointAzureDevOpsArgs()
        {
        }
        public static new ServiceEndpointAzureDevOpsArgs Empty => new ServiceEndpointAzureDevOpsArgs();
    }

    public sealed class ServiceEndpointAzureDevOpsState : global::Pulumi.ResourceArgs
    {
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
        /// The organization URL.
        /// </summary>
        [Input("orgUrl")]
        public Input<string>? OrgUrl { get; set; }

        [Input("personalAccessToken")]
        private Input<string>? _personalAccessToken;

        /// <summary>
        /// The Azure DevOps personal access token.
        /// </summary>
        public Input<string>? PersonalAccessToken
        {
            get => _personalAccessToken;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _personalAccessToken = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The ID of the project.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The URL of the release API.
        /// </summary>
        [Input("releaseApiUrl")]
        public Input<string>? ReleaseApiUrl { get; set; }

        /// <summary>
        /// The Service Endpoint name.
        /// </summary>
        [Input("serviceEndpointName")]
        public Input<string>? ServiceEndpointName { get; set; }

        public ServiceEndpointAzureDevOpsState()
        {
        }
        public static new ServiceEndpointAzureDevOpsState Empty => new ServiceEndpointAzureDevOpsState();
    }
}
