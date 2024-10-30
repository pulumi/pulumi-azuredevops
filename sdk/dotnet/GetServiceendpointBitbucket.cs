// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureDevOps
{
    public static class GetServiceendpointBitbucket
    {
        /// <summary>
        /// Use this data source to access information about an existing Bitbucket service Endpoint.
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
        ///     var example = AzureDevOps.GetProject.Invoke(new()
        ///     {
        ///         Name = "Example Project",
        ///     });
        /// 
        ///     var exampleGetServiceendpointBitbucket = AzureDevOps.GetServiceendpointBitbucket.Invoke(new()
        ///     {
        ///         ProjectId = example.Apply(getProjectResult =&gt; getProjectResult.Id),
        ///         ServiceEndpointId = "00000000-0000-0000-0000-000000000000",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["serviceEndpointName"] = exampleGetServiceendpointBitbucket.Apply(getServiceendpointBitbucketResult =&gt; getServiceendpointBitbucketResult.ServiceEndpointName),
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
        ///     var example = AzureDevOps.GetProject.Invoke(new()
        ///     {
        ///         Name = "Example Project",
        ///     });
        /// 
        ///     var exampleGetServiceendpointBitbucket = AzureDevOps.GetServiceendpointBitbucket.Invoke(new()
        ///     {
        ///         ProjectId = example.Apply(getProjectResult =&gt; getProjectResult.Id),
        ///         ServiceEndpointName = "Example",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["serviceEndpointId"] = exampleGetServiceendpointBitbucket.Apply(getServiceendpointBitbucketResult =&gt; getServiceendpointBitbucketResult.Id),
        ///     };
        /// });
        /// ```
        /// 
        /// ## PAT Permissions Required
        /// 
        /// - **vso.serviceendpoint**: Grants the ability to read service endpoints.
        /// </summary>
        public static Task<GetServiceendpointBitbucketResult> InvokeAsync(GetServiceendpointBitbucketArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetServiceendpointBitbucketResult>("azuredevops:index/getServiceendpointBitbucket:getServiceendpointBitbucket", args ?? new GetServiceendpointBitbucketArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to access information about an existing Bitbucket service Endpoint.
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
        ///     var example = AzureDevOps.GetProject.Invoke(new()
        ///     {
        ///         Name = "Example Project",
        ///     });
        /// 
        ///     var exampleGetServiceendpointBitbucket = AzureDevOps.GetServiceendpointBitbucket.Invoke(new()
        ///     {
        ///         ProjectId = example.Apply(getProjectResult =&gt; getProjectResult.Id),
        ///         ServiceEndpointId = "00000000-0000-0000-0000-000000000000",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["serviceEndpointName"] = exampleGetServiceendpointBitbucket.Apply(getServiceendpointBitbucketResult =&gt; getServiceendpointBitbucketResult.ServiceEndpointName),
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
        ///     var example = AzureDevOps.GetProject.Invoke(new()
        ///     {
        ///         Name = "Example Project",
        ///     });
        /// 
        ///     var exampleGetServiceendpointBitbucket = AzureDevOps.GetServiceendpointBitbucket.Invoke(new()
        ///     {
        ///         ProjectId = example.Apply(getProjectResult =&gt; getProjectResult.Id),
        ///         ServiceEndpointName = "Example",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["serviceEndpointId"] = exampleGetServiceendpointBitbucket.Apply(getServiceendpointBitbucketResult =&gt; getServiceendpointBitbucketResult.Id),
        ///     };
        /// });
        /// ```
        /// 
        /// ## PAT Permissions Required
        /// 
        /// - **vso.serviceendpoint**: Grants the ability to read service endpoints.
        /// </summary>
        public static Output<GetServiceendpointBitbucketResult> Invoke(GetServiceendpointBitbucketInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetServiceendpointBitbucketResult>("azuredevops:index/getServiceendpointBitbucket:getServiceendpointBitbucket", args ?? new GetServiceendpointBitbucketInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetServiceendpointBitbucketArgs : global::Pulumi.InvokeArgs
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

        public GetServiceendpointBitbucketArgs()
        {
        }
        public static new GetServiceendpointBitbucketArgs Empty => new GetServiceendpointBitbucketArgs();
    }

    public sealed class GetServiceendpointBitbucketInvokeArgs : global::Pulumi.InvokeArgs
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

        public GetServiceendpointBitbucketInvokeArgs()
        {
        }
        public static new GetServiceendpointBitbucketInvokeArgs Empty => new GetServiceendpointBitbucketInvokeArgs();
    }


    [OutputType]
    public sealed class GetServiceendpointBitbucketResult
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
        private GetServiceendpointBitbucketResult(
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
