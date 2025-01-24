// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureDevOps
{
    public static class GetServiceendpointNpm
    {
        /// <summary>
        /// Use this data source to access information about an existing NPM Service Endpoint.
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
        ///     var example = AzureDevOps.GetServiceendpointNpm.Invoke(new()
        ///     {
        ///         ProjectId = exampleAzuredevopsProject.Id,
        ///         ServiceEndpointName = "Example npm",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["serviceEndpointId"] = example.Apply(getServiceendpointNpmResult =&gt; getServiceendpointNpmResult.Id),
        ///     };
        /// });
        /// ```
        /// </summary>
        public static Task<GetServiceendpointNpmResult> InvokeAsync(GetServiceendpointNpmArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetServiceendpointNpmResult>("azuredevops:index/getServiceendpointNpm:getServiceendpointNpm", args ?? new GetServiceendpointNpmArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to access information about an existing NPM Service Endpoint.
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
        ///     var example = AzureDevOps.GetServiceendpointNpm.Invoke(new()
        ///     {
        ///         ProjectId = exampleAzuredevopsProject.Id,
        ///         ServiceEndpointName = "Example npm",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["serviceEndpointId"] = example.Apply(getServiceendpointNpmResult =&gt; getServiceendpointNpmResult.Id),
        ///     };
        /// });
        /// ```
        /// </summary>
        public static Output<GetServiceendpointNpmResult> Invoke(GetServiceendpointNpmInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetServiceendpointNpmResult>("azuredevops:index/getServiceendpointNpm:getServiceendpointNpm", args ?? new GetServiceendpointNpmInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to access information about an existing NPM Service Endpoint.
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
        ///     var example = AzureDevOps.GetServiceendpointNpm.Invoke(new()
        ///     {
        ///         ProjectId = exampleAzuredevopsProject.Id,
        ///         ServiceEndpointName = "Example npm",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["serviceEndpointId"] = example.Apply(getServiceendpointNpmResult =&gt; getServiceendpointNpmResult.Id),
        ///     };
        /// });
        /// ```
        /// </summary>
        public static Output<GetServiceendpointNpmResult> Invoke(GetServiceendpointNpmInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetServiceendpointNpmResult>("azuredevops:index/getServiceendpointNpm:getServiceendpointNpm", args ?? new GetServiceendpointNpmInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetServiceendpointNpmArgs : global::Pulumi.InvokeArgs
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

        public GetServiceendpointNpmArgs()
        {
        }
        public static new GetServiceendpointNpmArgs Empty => new GetServiceendpointNpmArgs();
    }

    public sealed class GetServiceendpointNpmInvokeArgs : global::Pulumi.InvokeArgs
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

        public GetServiceendpointNpmInvokeArgs()
        {
        }
        public static new GetServiceendpointNpmInvokeArgs Empty => new GetServiceendpointNpmInvokeArgs();
    }


    [OutputType]
    public sealed class GetServiceendpointNpmResult
    {
        /// <summary>
        /// The Authorization scheme.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Authorization;
        /// <summary>
        /// The description of the Service Endpoint.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string ProjectId;
        public readonly string ServiceEndpointId;
        public readonly string ServiceEndpointName;
        /// <summary>
        /// The URL of the NPM registry to connect with.
        /// </summary>
        public readonly string Url;

        [OutputConstructor]
        private GetServiceendpointNpmResult(
            ImmutableDictionary<string, string> authorization,

            string description,

            string id,

            string projectId,

            string serviceEndpointId,

            string serviceEndpointName,

            string url)
        {
            Authorization = authorization;
            Description = description;
            Id = id;
            ProjectId = projectId;
            ServiceEndpointId = serviceEndpointId;
            ServiceEndpointName = serviceEndpointName;
            Url = url;
        }
    }
}
