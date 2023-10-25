// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureDevOps.Agent
{
    [Obsolete(@"azuredevops.agent.getPool has been deprecated in favor of azuredevops.getPool")]
    public static class GetPool
    {
        /// <summary>
        /// Use this data source to access information about an existing Agent Pool within Azure DevOps.
        /// 
        /// ## Relevant Links
        /// 
        /// - [Azure DevOps Service REST API 7.0 - Agent Pools - Get](https://docs.microsoft.com/en-us/rest/api/azure/devops/distributedtask/pools/get?view=azure-devops-rest-7.0)
        /// </summary>
        public static Task<GetPoolResult> InvokeAsync(GetPoolArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetPoolResult>("azuredevops:Agent/getPool:getPool", args ?? new GetPoolArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to access information about an existing Agent Pool within Azure DevOps.
        /// 
        /// ## Relevant Links
        /// 
        /// - [Azure DevOps Service REST API 7.0 - Agent Pools - Get](https://docs.microsoft.com/en-us/rest/api/azure/devops/distributedtask/pools/get?view=azure-devops-rest-7.0)
        /// </summary>
        public static Output<GetPoolResult> Invoke(GetPoolInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetPoolResult>("azuredevops:Agent/getPool:getPool", args ?? new GetPoolInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPoolArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the Agent Pool.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetPoolArgs()
        {
        }
        public static new GetPoolArgs Empty => new GetPoolArgs();
    }

    public sealed class GetPoolInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the Agent Pool.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetPoolInvokeArgs()
        {
        }
        public static new GetPoolInvokeArgs Empty => new GetPoolInvokeArgs();
    }


    [OutputType]
    public sealed class GetPoolResult
    {
        public readonly bool AutoProvision;
        public readonly bool AutoUpdate;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string PoolType;

        [OutputConstructor]
        private GetPoolResult(
            bool autoProvision,

            bool autoUpdate,

            string id,

            string name,

            string poolType)
        {
            AutoProvision = autoProvision;
            AutoUpdate = autoUpdate;
            Id = id;
            Name = name;
            PoolType = poolType;
        }
    }
}
