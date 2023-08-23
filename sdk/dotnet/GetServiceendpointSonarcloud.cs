// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureDevOps
{
    public static class GetServiceendpointSonarcloud
    {
        /// <summary>
        /// Use this data source to access information about an existing Sonar Cloud Service Endpoint.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using AzureDevOps = Pulumi.AzureDevOps;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = AzureDevOps.GetServiceendpointSonarcloud.Invoke(new()
        ///     {
        ///         ProjectId = azuredevops_project.Example.Id,
        ///         ServiceEndpointName = "Example Sonar Cloud",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["serviceEndpointId"] = example.Apply(getServiceendpointSonarcloudResult =&gt; getServiceendpointSonarcloudResult.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetServiceendpointSonarcloudResult> InvokeAsync(GetServiceendpointSonarcloudArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetServiceendpointSonarcloudResult>("azuredevops:index/getServiceendpointSonarcloud:getServiceendpointSonarcloud", args ?? new GetServiceendpointSonarcloudArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to access information about an existing Sonar Cloud Service Endpoint.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using AzureDevOps = Pulumi.AzureDevOps;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = AzureDevOps.GetServiceendpointSonarcloud.Invoke(new()
        ///     {
        ///         ProjectId = azuredevops_project.Example.Id,
        ///         ServiceEndpointName = "Example Sonar Cloud",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["serviceEndpointId"] = example.Apply(getServiceendpointSonarcloudResult =&gt; getServiceendpointSonarcloudResult.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetServiceendpointSonarcloudResult> Invoke(GetServiceendpointSonarcloudInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetServiceendpointSonarcloudResult>("azuredevops:index/getServiceendpointSonarcloud:getServiceendpointSonarcloud", args ?? new GetServiceendpointSonarcloudInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetServiceendpointSonarcloudArgs : global::Pulumi.InvokeArgs
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

        public GetServiceendpointSonarcloudArgs()
        {
        }
        public static new GetServiceendpointSonarcloudArgs Empty => new GetServiceendpointSonarcloudArgs();
    }

    public sealed class GetServiceendpointSonarcloudInvokeArgs : global::Pulumi.InvokeArgs
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

        public GetServiceendpointSonarcloudInvokeArgs()
        {
        }
        public static new GetServiceendpointSonarcloudInvokeArgs Empty => new GetServiceendpointSonarcloudInvokeArgs();
    }


    [OutputType]
    public sealed class GetServiceendpointSonarcloudResult
    {
        /// <summary>
        /// Specifies the Authorization Scheme Map.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Authorization;
        /// <summary>
        /// Specifies the description of the Service Endpoint.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string ProjectId;
        public readonly string ServiceEndpointId;
        public readonly string ServiceEndpointName;

        [OutputConstructor]
        private GetServiceendpointSonarcloudResult(
            ImmutableDictionary<string, string> authorization,

            string description,

            string id,

            string projectId,

            string serviceEndpointId,

            string serviceEndpointName)
        {
            Authorization = authorization;
            Description = description;
            Id = id;
            ProjectId = projectId;
            ServiceEndpointId = serviceEndpointId;
            ServiceEndpointName = serviceEndpointName;
        }
    }
}