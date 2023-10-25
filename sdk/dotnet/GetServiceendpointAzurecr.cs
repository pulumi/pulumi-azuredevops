// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureDevOps
{
    public static class GetServiceendpointAzurecr
    {
        /// <summary>
        /// Use this data source to access information about an existing Azure Container Registry Service Endpoint.
        /// </summary>
        public static Task<GetServiceendpointAzurecrResult> InvokeAsync(GetServiceendpointAzurecrArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetServiceendpointAzurecrResult>("azuredevops:index/getServiceendpointAzurecr:getServiceendpointAzurecr", args ?? new GetServiceendpointAzurecrArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to access information about an existing Azure Container Registry Service Endpoint.
        /// </summary>
        public static Output<GetServiceendpointAzurecrResult> Invoke(GetServiceendpointAzurecrInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetServiceendpointAzurecrResult>("azuredevops:index/getServiceendpointAzurecr:getServiceendpointAzurecr", args ?? new GetServiceendpointAzurecrInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetServiceendpointAzurecrArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the project.
        /// </summary>
        [Input("projectId", required: true)]
        public string ProjectId { get; set; } = null!;

        /// <summary>
        /// the ID of the Service Endpoint.
        /// </summary>
        [Input("serviceEndpointId")]
        public string? ServiceEndpointId { get; set; }

        /// <summary>
        /// the Name of the Service Endpoint.
        /// 
        /// &gt; **NOTE:** One of either `service_endpoint_id` or `service_endpoint_name` must be specified.
        /// </summary>
        [Input("serviceEndpointName")]
        public string? ServiceEndpointName { get; set; }

        public GetServiceendpointAzurecrArgs()
        {
        }
        public static new GetServiceendpointAzurecrArgs Empty => new GetServiceendpointAzurecrArgs();
    }

    public sealed class GetServiceendpointAzurecrInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the project.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        /// <summary>
        /// the ID of the Service Endpoint.
        /// </summary>
        [Input("serviceEndpointId")]
        public Input<string>? ServiceEndpointId { get; set; }

        /// <summary>
        /// the Name of the Service Endpoint.
        /// 
        /// &gt; **NOTE:** One of either `service_endpoint_id` or `service_endpoint_name` must be specified.
        /// </summary>
        [Input("serviceEndpointName")]
        public Input<string>? ServiceEndpointName { get; set; }

        public GetServiceendpointAzurecrInvokeArgs()
        {
        }
        public static new GetServiceendpointAzurecrInvokeArgs Empty => new GetServiceendpointAzurecrInvokeArgs();
    }


    [OutputType]
    public sealed class GetServiceendpointAzurecrResult
    {
        /// <summary>
        /// The Object ID of the Service Principal.
        /// </summary>
        public readonly string AppObjectId;
        /// <summary>
        /// Specifies the Authorization Scheme Map.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Authorization;
        /// <summary>
        /// The ID of Service Principal Role Assignment.
        /// </summary>
        public readonly string AzSpnRoleAssignmentId;
        /// <summary>
        /// The Service Principal Role Permissions.
        /// </summary>
        public readonly string AzSpnRolePermissions;
        /// <summary>
        /// The Azure Container Registry name.
        /// </summary>
        public readonly string AzurecrName;
        /// <summary>
        /// The Tenant ID of the service principal.
        /// </summary>
        public readonly string AzurecrSpnTenantid;
        /// <summary>
        /// The Subscription ID of the Azure targets.
        /// </summary>
        public readonly string AzurecrSubscriptionId;
        /// <summary>
        /// The Subscription Name of the Azure targets.
        /// </summary>
        public readonly string AzurecrSubscriptionName;
        /// <summary>
        /// The Service Endpoint description.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string ProjectId;
        /// <summary>
        /// The Resource Group to which the Container Registry belongs.
        /// </summary>
        public readonly string ResourceGroup;
        public readonly string ServiceEndpointId;
        public readonly string ServiceEndpointName;
        /// <summary>
        /// The Application(Client) ID of the Service Principal.
        /// </summary>
        public readonly string ServicePrincipalId;
        /// <summary>
        /// The ID of the Service Principal.
        /// </summary>
        public readonly string SpnObjectId;

        [OutputConstructor]
        private GetServiceendpointAzurecrResult(
            string appObjectId,

            ImmutableDictionary<string, string> authorization,

            string azSpnRoleAssignmentId,

            string azSpnRolePermissions,

            string azurecrName,

            string azurecrSpnTenantid,

            string azurecrSubscriptionId,

            string azurecrSubscriptionName,

            string description,

            string id,

            string projectId,

            string resourceGroup,

            string serviceEndpointId,

            string serviceEndpointName,

            string servicePrincipalId,

            string spnObjectId)
        {
            AppObjectId = appObjectId;
            Authorization = authorization;
            AzSpnRoleAssignmentId = azSpnRoleAssignmentId;
            AzSpnRolePermissions = azSpnRolePermissions;
            AzurecrName = azurecrName;
            AzurecrSpnTenantid = azurecrSpnTenantid;
            AzurecrSubscriptionId = azurecrSubscriptionId;
            AzurecrSubscriptionName = azurecrSubscriptionName;
            Description = description;
            Id = id;
            ProjectId = projectId;
            ResourceGroup = resourceGroup;
            ServiceEndpointId = serviceEndpointId;
            ServiceEndpointName = serviceEndpointName;
            ServicePrincipalId = servicePrincipalId;
            SpnObjectId = spnObjectId;
        }
    }
}
