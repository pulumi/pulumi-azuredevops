// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureDevOps
{
    public static class GetServiceEndpointGithub
    {
        /// <summary>
        /// Use this data source to access information about an existing GitHub service Endpoint.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
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
        ///     var serviceendpoint = AzureDevOps.GetServiceEndpointGithub.Invoke(new()
        ///     {
        ///         ProjectId = sample.Apply(getProjectResult =&gt; getProjectResult.Id),
        ///         ServiceEndpointId = "00000000-0000-0000-0000-000000000000",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["serviceEndpointName"] = serviceendpoint.Apply(getServiceEndpointGithubResult =&gt; getServiceEndpointGithubResult.ServiceEndpointName),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% example %}}
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
        ///     var serviceendpoint = AzureDevOps.GetServiceEndpointGithub.Invoke(new()
        ///     {
        ///         ProjectId = sample.Apply(getProjectResult =&gt; getProjectResult.Id),
        ///         ServiceEndpointName = "Example-Service-Endpoint",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["serviceEndpointId"] = serviceendpoint.Apply(getServiceEndpointGithubResult =&gt; getServiceEndpointGithubResult.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetServiceEndpointGithubResult> InvokeAsync(GetServiceEndpointGithubArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetServiceEndpointGithubResult>("azuredevops:index/getServiceEndpointGithub:getServiceEndpointGithub", args ?? new GetServiceEndpointGithubArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to access information about an existing GitHub service Endpoint.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
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
        ///     var serviceendpoint = AzureDevOps.GetServiceEndpointGithub.Invoke(new()
        ///     {
        ///         ProjectId = sample.Apply(getProjectResult =&gt; getProjectResult.Id),
        ///         ServiceEndpointId = "00000000-0000-0000-0000-000000000000",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["serviceEndpointName"] = serviceendpoint.Apply(getServiceEndpointGithubResult =&gt; getServiceEndpointGithubResult.ServiceEndpointName),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% example %}}
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
        ///     var serviceendpoint = AzureDevOps.GetServiceEndpointGithub.Invoke(new()
        ///     {
        ///         ProjectId = sample.Apply(getProjectResult =&gt; getProjectResult.Id),
        ///         ServiceEndpointName = "Example-Service-Endpoint",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["serviceEndpointId"] = serviceendpoint.Apply(getServiceEndpointGithubResult =&gt; getServiceEndpointGithubResult.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetServiceEndpointGithubResult> Invoke(GetServiceEndpointGithubInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetServiceEndpointGithubResult>("azuredevops:index/getServiceEndpointGithub:getServiceEndpointGithub", args ?? new GetServiceEndpointGithubInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetServiceEndpointGithubArgs : global::Pulumi.InvokeArgs
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
        /// </summary>
        [Input("serviceEndpointName")]
        public string? ServiceEndpointName { get; set; }

        public GetServiceEndpointGithubArgs()
        {
        }
        public static new GetServiceEndpointGithubArgs Empty => new GetServiceEndpointGithubArgs();
    }

    public sealed class GetServiceEndpointGithubInvokeArgs : global::Pulumi.InvokeArgs
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
        /// </summary>
        [Input("serviceEndpointName")]
        public Input<string>? ServiceEndpointName { get; set; }

        public GetServiceEndpointGithubInvokeArgs()
        {
        }
        public static new GetServiceEndpointGithubInvokeArgs Empty => new GetServiceEndpointGithubInvokeArgs();
    }


    [OutputType]
    public sealed class GetServiceEndpointGithubResult
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
        private GetServiceEndpointGithubResult(
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
