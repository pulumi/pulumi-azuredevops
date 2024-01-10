// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureDevOps
{
    public static class GetEnvironment
    {
        /// <summary>
        /// Use this data source to access information about an Environment.
        /// 
        /// ## Relevant Links
        /// 
        /// * [Azure DevOps Service REST API 7.0 - Environments](https://docs.microsoft.com/en-us/rest/api/azure/devops/distributedtask/environments?view=azure-devops-rest-7.0)
        /// </summary>
        public static Task<GetEnvironmentResult> InvokeAsync(GetEnvironmentArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetEnvironmentResult>("azuredevops:index/getEnvironment:getEnvironment", args ?? new GetEnvironmentArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to access information about an Environment.
        /// 
        /// ## Relevant Links
        /// 
        /// * [Azure DevOps Service REST API 7.0 - Environments](https://docs.microsoft.com/en-us/rest/api/azure/devops/distributedtask/environments?view=azure-devops-rest-7.0)
        /// </summary>
        public static Output<GetEnvironmentResult> Invoke(GetEnvironmentInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetEnvironmentResult>("azuredevops:index/getEnvironment:getEnvironment", args ?? new GetEnvironmentInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetEnvironmentArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the Environment.
        /// </summary>
        [Input("environmentId")]
        public int? EnvironmentId { get; set; }

        /// <summary>
        /// Name of the Environment.
        /// 
        /// &gt; **NOTE:** One of either `environment_id` or `name` must be specified.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// The ID of the project.
        /// </summary>
        [Input("projectId", required: true)]
        public string ProjectId { get; set; } = null!;

        public GetEnvironmentArgs()
        {
        }
        public static new GetEnvironmentArgs Empty => new GetEnvironmentArgs();
    }

    public sealed class GetEnvironmentInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the Environment.
        /// </summary>
        [Input("environmentId")]
        public Input<int>? EnvironmentId { get; set; }

        /// <summary>
        /// Name of the Environment.
        /// 
        /// &gt; **NOTE:** One of either `environment_id` or `name` must be specified.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        public GetEnvironmentInvokeArgs()
        {
        }
        public static new GetEnvironmentInvokeArgs Empty => new GetEnvironmentInvokeArgs();
    }


    [OutputType]
    public sealed class GetEnvironmentResult
    {
        /// <summary>
        /// A description for the Environment.
        /// </summary>
        public readonly string Description;
        public readonly int? EnvironmentId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the Environment.
        /// </summary>
        public readonly string Name;
        public readonly string ProjectId;

        [OutputConstructor]
        private GetEnvironmentResult(
            string description,

            int? environmentId,

            string id,

            string name,

            string projectId)
        {
            Description = description;
            EnvironmentId = environmentId;
            Id = id;
            Name = name;
            ProjectId = projectId;
        }
    }
}
