// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureDevOps
{
    public static class GetServiceEndpointAzureRM
    {
        /// <summary>
        /// Use this data source to access information about an existing AzureRM service Endpoint.
        /// 
        /// ## Example Usage
        /// 
        /// ### By Service Endpoint ID
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using AzureDevOps = Pulumi.AzureDevOps;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var sample = AzureDevOps.GetProject.Invoke(new()
        ///     {
        ///         Name = "Sample Project",
        ///     });
        /// 
        ///     var serviceendpoint = AzureDevOps.GetServiceEndpointAzureRM.Invoke(new()
        ///     {
        ///         ProjectId = sample.Apply(getProjectResult =&gt; getProjectResult.Id),
        ///         ServiceEndpointId = "00000000-0000-0000-0000-000000000000",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["serviceEndpointName"] = serviceendpoint.Apply(getServiceEndpointAzureRMResult =&gt; getServiceEndpointAzureRMResult.ServiceEndpointName),
        ///     };
        /// });
        /// ```
        /// 
        /// ### By Service Endpoint Name
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using AzureDevOps = Pulumi.AzureDevOps;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var sample = AzureDevOps.GetProject.Invoke(new()
        ///     {
        ///         Name = "Sample Project",
        ///     });
        /// 
        ///     var serviceendpoint = AzureDevOps.GetServiceEndpointAzureRM.Invoke(new()
        ///     {
        ///         ProjectId = sample.Apply(getProjectResult =&gt; getProjectResult.Id),
        ///         ServiceEndpointName = "Example-Service-Endpoint",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["serviceEndpointId"] = serviceendpoint.Apply(getServiceEndpointAzureRMResult =&gt; getServiceEndpointAzureRMResult.Id),
        ///     };
        /// });
        /// ```
        /// </summary>
        public static Task<GetServiceEndpointAzureRMResult> InvokeAsync(GetServiceEndpointAzureRMArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetServiceEndpointAzureRMResult>("azuredevops:index/getServiceEndpointAzureRM:getServiceEndpointAzureRM", args ?? new GetServiceEndpointAzureRMArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to access information about an existing AzureRM service Endpoint.
        /// 
        /// ## Example Usage
        /// 
        /// ### By Service Endpoint ID
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using AzureDevOps = Pulumi.AzureDevOps;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var sample = AzureDevOps.GetProject.Invoke(new()
        ///     {
        ///         Name = "Sample Project",
        ///     });
        /// 
        ///     var serviceendpoint = AzureDevOps.GetServiceEndpointAzureRM.Invoke(new()
        ///     {
        ///         ProjectId = sample.Apply(getProjectResult =&gt; getProjectResult.Id),
        ///         ServiceEndpointId = "00000000-0000-0000-0000-000000000000",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["serviceEndpointName"] = serviceendpoint.Apply(getServiceEndpointAzureRMResult =&gt; getServiceEndpointAzureRMResult.ServiceEndpointName),
        ///     };
        /// });
        /// ```
        /// 
        /// ### By Service Endpoint Name
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using AzureDevOps = Pulumi.AzureDevOps;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var sample = AzureDevOps.GetProject.Invoke(new()
        ///     {
        ///         Name = "Sample Project",
        ///     });
        /// 
        ///     var serviceendpoint = AzureDevOps.GetServiceEndpointAzureRM.Invoke(new()
        ///     {
        ///         ProjectId = sample.Apply(getProjectResult =&gt; getProjectResult.Id),
        ///         ServiceEndpointName = "Example-Service-Endpoint",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["serviceEndpointId"] = serviceendpoint.Apply(getServiceEndpointAzureRMResult =&gt; getServiceEndpointAzureRMResult.Id),
        ///     };
        /// });
        /// ```
        /// </summary>
        public static Output<GetServiceEndpointAzureRMResult> Invoke(GetServiceEndpointAzureRMInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetServiceEndpointAzureRMResult>("azuredevops:index/getServiceEndpointAzureRM:getServiceEndpointAzureRM", args ?? new GetServiceEndpointAzureRMInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetServiceEndpointAzureRMArgs : global::Pulumi.InvokeArgs
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
        /// &gt; **NOTE:** When supplying `service_endpoint_name`, take care to ensure that this is a unique name.
        /// </summary>
        [Input("serviceEndpointName")]
        public string? ServiceEndpointName { get; set; }

        public GetServiceEndpointAzureRMArgs()
        {
        }
        public static new GetServiceEndpointAzureRMArgs Empty => new GetServiceEndpointAzureRMArgs();
    }

    public sealed class GetServiceEndpointAzureRMInvokeArgs : global::Pulumi.InvokeArgs
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
        /// &gt; **NOTE:** When supplying `service_endpoint_name`, take care to ensure that this is a unique name.
        /// </summary>
        [Input("serviceEndpointName")]
        public Input<string>? ServiceEndpointName { get; set; }

        public GetServiceEndpointAzureRMInvokeArgs()
        {
        }
        public static new GetServiceEndpointAzureRMInvokeArgs Empty => new GetServiceEndpointAzureRMInvokeArgs();
    }


    [OutputType]
    public sealed class GetServiceEndpointAzureRMResult
    {
        /// <summary>
        /// Specifies the Authorization Scheme Map.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Authorization;
        /// <summary>
        /// Specified the Management Group ID of the Service Endpoint is target, if available.
        /// </summary>
        public readonly string AzurermManagementGroupId;
        /// <summary>
        /// Specified the Management Group Name of the Service Endpoint target, if available.
        /// </summary>
        public readonly string AzurermManagementGroupName;
        /// <summary>
        /// Specifies the Tenant ID of the Azure targets.
        /// </summary>
        public readonly string AzurermSpnTenantid;
        /// <summary>
        /// Specifies the Subscription ID of the Service Endpoint target, if available.
        /// </summary>
        public readonly string AzurermSubscriptionId;
        /// <summary>
        /// Specifies the Subscription Name of the Service Endpoint target, if available.
        /// </summary>
        public readonly string AzurermSubscriptionName;
        /// <summary>
        /// Specifies the description of the Service Endpoint.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The Cloud Environment. Possible values are `AzureCloud`, `AzureChinaCloud`, `AzureUSGovernment`, and `AzureGermanCloud`.
        /// </summary>
        public readonly string Environment;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string ProjectId;
        /// <summary>
        /// Specifies the Resource Group of the Service Endpoint target, if available.
        /// </summary>
        public readonly string ResourceGroup;
        /// <summary>
        /// Specifies the authentication scheme of azurerm endpoint, either `WorkloadIdentityFederation`, `ManagedServiceIdentity` or `ServicePrincipal`.
        /// </summary>
        public readonly string ServiceEndpointAuthenticationScheme;
        public readonly string ServiceEndpointId;
        public readonly string ServiceEndpointName;
        /// <summary>
        /// The Application(Client) ID of the Service Principal.
        /// </summary>
        public readonly string ServicePrincipalId;
        /// <summary>
        /// The issuer if `service_endpoint_authentication_scheme` is set to `WorkloadIdentityFederation`. This looks like `https://vstoken.dev.azure.com/f66a4bc2-08ad-4ec0-a25e-e769d6b3b294`, where the GUID is the Organization ID of your Azure DevOps Organisation.
        /// </summary>
        public readonly string WorkloadIdentityFederationIssuer;
        /// <summary>
        /// The subject if `service_endpoint_authentication_scheme` is set to `WorkloadIdentityFederation`. This looks like `sc://my-organisation/my-project/my-service-connection-name`.
        /// </summary>
        public readonly string WorkloadIdentityFederationSubject;

        [OutputConstructor]
        private GetServiceEndpointAzureRMResult(
            ImmutableDictionary<string, string> authorization,

            string azurermManagementGroupId,

            string azurermManagementGroupName,

            string azurermSpnTenantid,

            string azurermSubscriptionId,

            string azurermSubscriptionName,

            string description,

            string environment,

            string id,

            string projectId,

            string resourceGroup,

            string serviceEndpointAuthenticationScheme,

            string serviceEndpointId,

            string serviceEndpointName,

            string servicePrincipalId,

            string workloadIdentityFederationIssuer,

            string workloadIdentityFederationSubject)
        {
            Authorization = authorization;
            AzurermManagementGroupId = azurermManagementGroupId;
            AzurermManagementGroupName = azurermManagementGroupName;
            AzurermSpnTenantid = azurermSpnTenantid;
            AzurermSubscriptionId = azurermSubscriptionId;
            AzurermSubscriptionName = azurermSubscriptionName;
            Description = description;
            Environment = environment;
            Id = id;
            ProjectId = projectId;
            ResourceGroup = resourceGroup;
            ServiceEndpointAuthenticationScheme = serviceEndpointAuthenticationScheme;
            ServiceEndpointId = serviceEndpointId;
            ServiceEndpointName = serviceEndpointName;
            ServicePrincipalId = servicePrincipalId;
            WorkloadIdentityFederationIssuer = workloadIdentityFederationIssuer;
            WorkloadIdentityFederationSubject = workloadIdentityFederationSubject;
        }
    }
}
