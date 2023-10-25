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
    /// Manages a Azure Container Registry service endpoint within Azure DevOps.
    /// 
    /// ## Relevant Links
    /// 
    /// - [Azure DevOps Service REST API 7.0 - Service Endpoints](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-7.0)
    /// - [Azure Container Registry REST API](https://docs.microsoft.com/en-us/rest/api/containerregistry/)
    /// 
    /// ## Import
    /// 
    /// Azure DevOps Service Endpoint Azure Container Registry can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**
    /// 
    /// ```sh
    ///  $ pulumi import azuredevops:index/serviceEndpointAzureEcr:ServiceEndpointAzureEcr example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
    /// ```
    /// </summary>
    [AzureDevOpsResourceType("azuredevops:index/serviceEndpointAzureEcr:ServiceEndpointAzureEcr")]
    public partial class ServiceEndpointAzureEcr : global::Pulumi.CustomResource
    {
        [Output("appObjectId")]
        public Output<string> AppObjectId { get; private set; } = null!;

        [Output("authorization")]
        public Output<ImmutableDictionary<string, string>> Authorization { get; private set; } = null!;

        [Output("azSpnRoleAssignmentId")]
        public Output<string> AzSpnRoleAssignmentId { get; private set; } = null!;

        [Output("azSpnRolePermissions")]
        public Output<string> AzSpnRolePermissions { get; private set; } = null!;

        /// <summary>
        /// The Azure container registry name.
        /// </summary>
        [Output("azurecrName")]
        public Output<string> AzurecrName { get; private set; } = null!;

        /// <summary>
        /// The tenant id of the service principal.
        /// </summary>
        [Output("azurecrSpnTenantid")]
        public Output<string> AzurecrSpnTenantid { get; private set; } = null!;

        /// <summary>
        /// The subscription id of the Azure targets.
        /// </summary>
        [Output("azurecrSubscriptionId")]
        public Output<string> AzurecrSubscriptionId { get; private set; } = null!;

        /// <summary>
        /// The subscription name of the Azure targets.
        /// </summary>
        [Output("azurecrSubscriptionName")]
        public Output<string> AzurecrSubscriptionName { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The ID of the project.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// The resource group to which the container registry belongs.
        /// </summary>
        [Output("resourceGroup")]
        public Output<string> ResourceGroup { get; private set; } = null!;

        /// <summary>
        /// The name you will use to refer to this service connection in task inputs.
        /// </summary>
        [Output("serviceEndpointName")]
        public Output<string> ServiceEndpointName { get; private set; } = null!;

        /// <summary>
        /// The service principal ID.
        /// </summary>
        [Output("servicePrincipalId")]
        public Output<string> ServicePrincipalId { get; private set; } = null!;

        [Output("spnObjectId")]
        public Output<string> SpnObjectId { get; private set; } = null!;


        /// <summary>
        /// Create a ServiceEndpointAzureEcr resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServiceEndpointAzureEcr(string name, ServiceEndpointAzureEcrArgs args, CustomResourceOptions? options = null)
            : base("azuredevops:index/serviceEndpointAzureEcr:ServiceEndpointAzureEcr", name, args ?? new ServiceEndpointAzureEcrArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServiceEndpointAzureEcr(string name, Input<string> id, ServiceEndpointAzureEcrState? state = null, CustomResourceOptions? options = null)
            : base("azuredevops:index/serviceEndpointAzureEcr:ServiceEndpointAzureEcr", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ServiceEndpointAzureEcr resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServiceEndpointAzureEcr Get(string name, Input<string> id, ServiceEndpointAzureEcrState? state = null, CustomResourceOptions? options = null)
        {
            return new ServiceEndpointAzureEcr(name, id, state, options);
        }
    }

    public sealed class ServiceEndpointAzureEcrArgs : global::Pulumi.ResourceArgs
    {
        [Input("authorization")]
        private InputMap<string>? _authorization;
        public InputMap<string> Authorization
        {
            get => _authorization ?? (_authorization = new InputMap<string>());
            set => _authorization = value;
        }

        /// <summary>
        /// The Azure container registry name.
        /// </summary>
        [Input("azurecrName", required: true)]
        public Input<string> AzurecrName { get; set; } = null!;

        /// <summary>
        /// The tenant id of the service principal.
        /// </summary>
        [Input("azurecrSpnTenantid", required: true)]
        public Input<string> AzurecrSpnTenantid { get; set; } = null!;

        /// <summary>
        /// The subscription id of the Azure targets.
        /// </summary>
        [Input("azurecrSubscriptionId", required: true)]
        public Input<string> AzurecrSubscriptionId { get; set; } = null!;

        /// <summary>
        /// The subscription name of the Azure targets.
        /// </summary>
        [Input("azurecrSubscriptionName", required: true)]
        public Input<string> AzurecrSubscriptionName { get; set; } = null!;

        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The ID of the project.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        /// <summary>
        /// The resource group to which the container registry belongs.
        /// </summary>
        [Input("resourceGroup", required: true)]
        public Input<string> ResourceGroup { get; set; } = null!;

        /// <summary>
        /// The name you will use to refer to this service connection in task inputs.
        /// </summary>
        [Input("serviceEndpointName", required: true)]
        public Input<string> ServiceEndpointName { get; set; } = null!;

        public ServiceEndpointAzureEcrArgs()
        {
        }
        public static new ServiceEndpointAzureEcrArgs Empty => new ServiceEndpointAzureEcrArgs();
    }

    public sealed class ServiceEndpointAzureEcrState : global::Pulumi.ResourceArgs
    {
        [Input("appObjectId")]
        public Input<string>? AppObjectId { get; set; }

        [Input("authorization")]
        private InputMap<string>? _authorization;
        public InputMap<string> Authorization
        {
            get => _authorization ?? (_authorization = new InputMap<string>());
            set => _authorization = value;
        }

        [Input("azSpnRoleAssignmentId")]
        public Input<string>? AzSpnRoleAssignmentId { get; set; }

        [Input("azSpnRolePermissions")]
        public Input<string>? AzSpnRolePermissions { get; set; }

        /// <summary>
        /// The Azure container registry name.
        /// </summary>
        [Input("azurecrName")]
        public Input<string>? AzurecrName { get; set; }

        /// <summary>
        /// The tenant id of the service principal.
        /// </summary>
        [Input("azurecrSpnTenantid")]
        public Input<string>? AzurecrSpnTenantid { get; set; }

        /// <summary>
        /// The subscription id of the Azure targets.
        /// </summary>
        [Input("azurecrSubscriptionId")]
        public Input<string>? AzurecrSubscriptionId { get; set; }

        /// <summary>
        /// The subscription name of the Azure targets.
        /// </summary>
        [Input("azurecrSubscriptionName")]
        public Input<string>? AzurecrSubscriptionName { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The ID of the project.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The resource group to which the container registry belongs.
        /// </summary>
        [Input("resourceGroup")]
        public Input<string>? ResourceGroup { get; set; }

        /// <summary>
        /// The name you will use to refer to this service connection in task inputs.
        /// </summary>
        [Input("serviceEndpointName")]
        public Input<string>? ServiceEndpointName { get; set; }

        /// <summary>
        /// The service principal ID.
        /// </summary>
        [Input("servicePrincipalId")]
        public Input<string>? ServicePrincipalId { get; set; }

        [Input("spnObjectId")]
        public Input<string>? SpnObjectId { get; set; }

        public ServiceEndpointAzureEcrState()
        {
        }
        public static new ServiceEndpointAzureEcrState Empty => new ServiceEndpointAzureEcrState();
    }
}
