// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureDevOps
{
    public static class GetAgentQueue
    {
        /// <summary>
        /// Use this data source to access information about an existing Agent Queue within Azure DevOps.
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
        ///     var exampleProject = new AzureDevOps.Project("exampleProject", new()
        ///     {
        ///         WorkItemTemplate = "Agile",
        ///         VersionControl = "Git",
        ///         Visibility = "private",
        ///         Description = "Managed by Terraform",
        ///     });
        /// 
        ///     var exampleAgentQueue = AzureDevOps.GetAgentQueue.Invoke(new()
        ///     {
        ///         ProjectId = exampleProject.Id,
        ///         Name = "Example Agent Queue",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["name"] = exampleAgentQueue.Apply(getAgentQueueResult =&gt; getAgentQueueResult.Name),
        ///         ["poolId"] = exampleAgentQueue.Apply(getAgentQueueResult =&gt; getAgentQueueResult.AgentPoolId),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// ## Relevant Links
        /// 
        /// - [Azure DevOps Service REST API 7.0 - Agent Queues - Get](https://docs.microsoft.com/en-us/rest/api/azure/devops/distributedtask/queues/get?view=azure-devops-rest-7.0)
        /// </summary>
        public static Task<GetAgentQueueResult> InvokeAsync(GetAgentQueueArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAgentQueueResult>("azuredevops:index/getAgentQueue:getAgentQueue", args ?? new GetAgentQueueArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to access information about an existing Agent Queue within Azure DevOps.
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
        ///     var exampleProject = new AzureDevOps.Project("exampleProject", new()
        ///     {
        ///         WorkItemTemplate = "Agile",
        ///         VersionControl = "Git",
        ///         Visibility = "private",
        ///         Description = "Managed by Terraform",
        ///     });
        /// 
        ///     var exampleAgentQueue = AzureDevOps.GetAgentQueue.Invoke(new()
        ///     {
        ///         ProjectId = exampleProject.Id,
        ///         Name = "Example Agent Queue",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["name"] = exampleAgentQueue.Apply(getAgentQueueResult =&gt; getAgentQueueResult.Name),
        ///         ["poolId"] = exampleAgentQueue.Apply(getAgentQueueResult =&gt; getAgentQueueResult.AgentPoolId),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// ## Relevant Links
        /// 
        /// - [Azure DevOps Service REST API 7.0 - Agent Queues - Get](https://docs.microsoft.com/en-us/rest/api/azure/devops/distributedtask/queues/get?view=azure-devops-rest-7.0)
        /// </summary>
        public static Output<GetAgentQueueResult> Invoke(GetAgentQueueInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAgentQueueResult>("azuredevops:index/getAgentQueue:getAgentQueue", args ?? new GetAgentQueueInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAgentQueueArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the Agent Queue.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The Project Id.
        /// </summary>
        [Input("projectId", required: true)]
        public string ProjectId { get; set; } = null!;

        public GetAgentQueueArgs()
        {
        }
        public static new GetAgentQueueArgs Empty => new GetAgentQueueArgs();
    }

    public sealed class GetAgentQueueInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the Agent Queue.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The Project Id.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        public GetAgentQueueInvokeArgs()
        {
        }
        public static new GetAgentQueueInvokeArgs Empty => new GetAgentQueueInvokeArgs();
    }


    [OutputType]
    public sealed class GetAgentQueueResult
    {
        /// <summary>
        /// Agent pool identifier to which the agent queue belongs.
        /// </summary>
        public readonly int AgentPoolId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the agent queue.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Project identifier to which the agent queue belongs.
        /// </summary>
        public readonly string ProjectId;

        [OutputConstructor]
        private GetAgentQueueResult(
            int agentPoolId,

            string id,

            string name,

            string projectId)
        {
            AgentPoolId = agentPoolId;
            Id = id;
            Name = name;
            ProjectId = projectId;
        }
    }
}
