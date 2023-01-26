// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureDevOps.Agent
{
    [Obsolete(@"azuredevops.agent.getPools has been deprecated in favor of azuredevops.getPools")]
    public static class GetPools
    {
        /// <summary>
        /// Use this data source to access information about existing Agent Pools within Azure DevOps.
        /// 
        /// ## Relevant Links
        /// 
        /// - [Azure DevOps Service REST API 6.0 - Agent Pools - Get](https://docs.microsoft.com/en-us/rest/api/azure/devops/distributedtask/pools/get?view=azure-devops-rest-6.0)
        /// </summary>
        public static Task<GetPoolsResult> InvokeAsync(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetPoolsResult>("azuredevops:Agent/getPools:getPools", InvokeArgs.Empty, options.WithDefaults());
    }


    [OutputType]
    public sealed class GetPoolsResult
    {
        /// <summary>
        /// A list of existing agent pools in your Azure DevOps Organization with the following details about every agent pool:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetPoolsAgentPoolResult> AgentPools;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetPoolsResult(
            ImmutableArray<Outputs.GetPoolsAgentPoolResult> agentPools,

            string id)
        {
            AgentPools = agentPools;
            Id = id;
        }
    }
}
