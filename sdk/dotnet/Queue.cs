// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureDevOps
{
    /// <summary>
    /// Manages an agent queue within Azure DevOps. In the UI, this is equivalent to adding an
    /// Organization defined pool to a project.
    /// 
    /// The created queue is not authorized for use by all pipelines in the project. However,
    /// the `azuredevops.ResourceAuthorization` resource can be used to grant authorization.
    /// 
    /// ## Example Usage
    /// ### Creating a Queue from an organization-level pool
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AzureDevOps = Pulumi.AzureDevOps;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var exampleProject = new AzureDevOps.Project("exampleProject");
    /// 
    ///     var examplePool = AzureDevOps.GetPool.Invoke(new()
    ///     {
    ///         Name = "example-pool",
    ///     });
    /// 
    ///     var exampleQueue = new AzureDevOps.Queue("exampleQueue", new()
    ///     {
    ///         ProjectId = exampleProject.Id,
    ///         AgentPoolId = examplePool.Apply(getPoolResult =&gt; getPoolResult.Id),
    ///     });
    /// 
    ///     // Grant access to queue to all pipelines in the project
    ///     var exampleResourceAuthorization = new AzureDevOps.ResourceAuthorization("exampleResourceAuthorization", new()
    ///     {
    ///         ProjectId = exampleProject.Id,
    ///         ResourceId = exampleQueue.Id,
    ///         Type = "queue",
    ///         Authorized = true,
    ///     });
    /// 
    /// });
    /// ```
    /// ### Creating a Queue at the project level (Organization-level permissions not required)
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AzureDevOps = Pulumi.AzureDevOps;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var exampleProject = AzureDevOps.GetProject.Invoke(new()
    ///     {
    ///         Name = "Example Project",
    ///     });
    /// 
    ///     var exampleQueue = new AzureDevOps.Queue("exampleQueue", new()
    ///     {
    ///         ProjectId = exampleProject.Apply(getProjectResult =&gt; getProjectResult.Id),
    ///     });
    /// 
    /// });
    /// ```
    /// ## Relevant Links
    /// 
    /// - [Azure DevOps Service REST API 7.0 - Agent Queues](https://docs.microsoft.com/en-us/rest/api/azure/devops/distributedtask/queues?view=azure-devops-rest-7.0)
    /// 
    /// ## Import
    /// 
    /// Azure DevOps Agent Pools can be imported using the project ID and agent queue ID, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azuredevops:index/queue:Queue example 00000000-0000-0000-0000-000000000000/0
    /// ```
    /// </summary>
    [AzureDevOpsResourceType("azuredevops:index/queue:Queue")]
    public partial class Queue : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the organization agent pool. Conflicts with `name`.
        /// 
        /// &gt; **NOTE:**
        /// One of `name` or `agent_pool_id` must be specified, but not both.
        /// When `agent_pool_id` is specified, the agent queue name will be derived from the agent pool name.
        /// </summary>
        [Output("agentPoolId")]
        public Output<int> AgentPoolId { get; private set; } = null!;

        /// <summary>
        /// The name of the agent queue. Defaults to the ID of the agent pool. Conflicts with `agent_pool_id`.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which to create the resource.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;


        /// <summary>
        /// Create a Queue resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Queue(string name, QueueArgs args, CustomResourceOptions? options = null)
            : base("azuredevops:index/queue:Queue", name, args ?? new QueueArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Queue(string name, Input<string> id, QueueState? state = null, CustomResourceOptions? options = null)
            : base("azuredevops:index/queue:Queue", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azuredevops:Agent/queue:Queue" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Queue resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Queue Get(string name, Input<string> id, QueueState? state = null, CustomResourceOptions? options = null)
        {
            return new Queue(name, id, state, options);
        }
    }

    public sealed class QueueArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the organization agent pool. Conflicts with `name`.
        /// 
        /// &gt; **NOTE:**
        /// One of `name` or `agent_pool_id` must be specified, but not both.
        /// When `agent_pool_id` is specified, the agent queue name will be derived from the agent pool name.
        /// </summary>
        [Input("agentPoolId")]
        public Input<int>? AgentPoolId { get; set; }

        /// <summary>
        /// The name of the agent queue. Defaults to the ID of the agent pool. Conflicts with `agent_pool_id`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which to create the resource.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        public QueueArgs()
        {
        }
        public static new QueueArgs Empty => new QueueArgs();
    }

    public sealed class QueueState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the organization agent pool. Conflicts with `name`.
        /// 
        /// &gt; **NOTE:**
        /// One of `name` or `agent_pool_id` must be specified, but not both.
        /// When `agent_pool_id` is specified, the agent queue name will be derived from the agent pool name.
        /// </summary>
        [Input("agentPoolId")]
        public Input<int>? AgentPoolId { get; set; }

        /// <summary>
        /// The name of the agent queue. Defaults to the ID of the agent pool. Conflicts with `agent_pool_id`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which to create the resource.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        public QueueState()
        {
        }
        public static new QueueState Empty => new QueueState();
    }
}
