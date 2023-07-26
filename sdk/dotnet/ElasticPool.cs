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
    /// Manages Elastic pool within Azure DevOps.
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
    ///     var exampleProject = new AzureDevOps.Project("exampleProject", new()
    ///     {
    ///         Visibility = "private",
    ///         VersionControl = "Git",
    ///         WorkItemTemplate = "Agile",
    ///         Description = "Managed by Terraform",
    ///     });
    /// 
    ///     var exampleServiceEndpointAzureRM = new AzureDevOps.ServiceEndpointAzureRM("exampleServiceEndpointAzureRM", new()
    ///     {
    ///         ProjectId = exampleProject.Id,
    ///         ServiceEndpointName = "Example Azure Connection",
    ///         Description = "Managed by Terraform",
    ///         ServiceEndpointAuthenticationScheme = "ServicePrincipal",
    ///         Credentials = new AzureDevOps.Inputs.ServiceEndpointAzureRMCredentialsArgs
    ///         {
    ///             Serviceprincipalid = "00000000-0000-0000-0000-000000000000",
    ///             Serviceprincipalkey = "00000000-0000-0000-0000-000000000000",
    ///         },
    ///         AzurermSpnTenantid = "00000000-0000-0000-0000-000000000000",
    ///         AzurermSubscriptionId = "00000000-0000-0000-0000-000000000000",
    ///         AzurermSubscriptionName = "Subscription Name",
    ///     });
    /// 
    ///     var exampleElasticPool = new AzureDevOps.ElasticPool("exampleElasticPool", new()
    ///     {
    ///         ServiceEndpointId = exampleServiceEndpointAzureRM.Id,
    ///         ServiceEndpointScope = exampleProject.Id,
    ///         DesiredIdle = 2,
    ///         MaxCapacity = 3,
    ///         AzureResourceId = "/subscriptions/&lt;Subscription Id&gt;/resourceGroups/&lt;Resource Name&gt;/providers/Microsoft.Compute/virtualMachineScaleSets/&lt;VMSS Name&gt;",
    ///     });
    /// 
    /// });
    /// ```
    /// ## Relevant Links
    /// 
    /// - [Azure DevOps Service REST API 7.0 - Elastic Pools](https://learn.microsoft.com/en-us/rest/api/azure/devops/distributedtask/elasticpools/create?view=azure-devops-rest-7.0)
    /// 
    /// ## Import
    /// 
    /// Azure DevOps Agent Pools can be imported using the Elastic pool ID, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azuredevops:index/elasticPool:ElasticPool example 0
    /// ```
    /// </summary>
    [AzureDevOpsResourceType("azuredevops:index/elasticPool:ElasticPool")]
    public partial class ElasticPool : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Set whether agents should be configured to run with interactive UI. Defaults to `false`.
        /// </summary>
        [Output("agentInteractiveUi")]
        public Output<bool?> AgentInteractiveUi { get; private set; } = null!;

        /// <summary>
        /// Specifies whether a queue should be automatically provisioned for each project collection. Defaults to `false`.
        /// </summary>
        [Output("autoProvision")]
        public Output<bool?> AutoProvision { get; private set; } = null!;

        /// <summary>
        /// Specifies whether or not agents within the pool should be automatically updated. Defaults to `true`.
        /// </summary>
        [Output("autoUpdate")]
        public Output<bool?> AutoUpdate { get; private set; } = null!;

        /// <summary>
        /// The ID of the Azure resource.
        /// </summary>
        [Output("azureResourceId")]
        public Output<string> AzureResourceId { get; private set; } = null!;

        /// <summary>
        /// Number of agents to keep on standby.
        /// </summary>
        [Output("desiredIdle")]
        public Output<int> DesiredIdle { get; private set; } = null!;

        /// <summary>
        /// Maximum number of virtual machines in the scale set.
        /// </summary>
        [Output("maxCapacity")]
        public Output<int> MaxCapacity { get; private set; } = null!;

        /// <summary>
        /// The name of the Elastic pool.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Tear down virtual machines after every use. Defaults to `false`.
        /// </summary>
        [Output("recycleAfterEachUse")]
        public Output<bool?> RecycleAfterEachUse { get; private set; } = null!;

        /// <summary>
        /// The ID of Service Endpoint used to connect to Azure.
        /// </summary>
        [Output("serviceEndpointId")]
        public Output<string> ServiceEndpointId { get; private set; } = null!;

        /// <summary>
        /// The Project ID of Service Endpoint belongs to.
        /// </summary>
        [Output("serviceEndpointScope")]
        public Output<string> ServiceEndpointScope { get; private set; } = null!;

        /// <summary>
        /// Delay in minutes before deleting excess idle agents. Defaults to `30`.
        /// </summary>
        [Output("timeToLiveMinutes")]
        public Output<int?> TimeToLiveMinutes { get; private set; } = null!;


        /// <summary>
        /// Create a ElasticPool resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ElasticPool(string name, ElasticPoolArgs args, CustomResourceOptions? options = null)
            : base("azuredevops:index/elasticPool:ElasticPool", name, args ?? new ElasticPoolArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ElasticPool(string name, Input<string> id, ElasticPoolState? state = null, CustomResourceOptions? options = null)
            : base("azuredevops:index/elasticPool:ElasticPool", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ElasticPool resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ElasticPool Get(string name, Input<string> id, ElasticPoolState? state = null, CustomResourceOptions? options = null)
        {
            return new ElasticPool(name, id, state, options);
        }
    }

    public sealed class ElasticPoolArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Set whether agents should be configured to run with interactive UI. Defaults to `false`.
        /// </summary>
        [Input("agentInteractiveUi")]
        public Input<bool>? AgentInteractiveUi { get; set; }

        /// <summary>
        /// Specifies whether a queue should be automatically provisioned for each project collection. Defaults to `false`.
        /// </summary>
        [Input("autoProvision")]
        public Input<bool>? AutoProvision { get; set; }

        /// <summary>
        /// Specifies whether or not agents within the pool should be automatically updated. Defaults to `true`.
        /// </summary>
        [Input("autoUpdate")]
        public Input<bool>? AutoUpdate { get; set; }

        /// <summary>
        /// The ID of the Azure resource.
        /// </summary>
        [Input("azureResourceId", required: true)]
        public Input<string> AzureResourceId { get; set; } = null!;

        /// <summary>
        /// Number of agents to keep on standby.
        /// </summary>
        [Input("desiredIdle", required: true)]
        public Input<int> DesiredIdle { get; set; } = null!;

        /// <summary>
        /// Maximum number of virtual machines in the scale set.
        /// </summary>
        [Input("maxCapacity", required: true)]
        public Input<int> MaxCapacity { get; set; } = null!;

        /// <summary>
        /// The name of the Elastic pool.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Tear down virtual machines after every use. Defaults to `false`.
        /// </summary>
        [Input("recycleAfterEachUse")]
        public Input<bool>? RecycleAfterEachUse { get; set; }

        /// <summary>
        /// The ID of Service Endpoint used to connect to Azure.
        /// </summary>
        [Input("serviceEndpointId", required: true)]
        public Input<string> ServiceEndpointId { get; set; } = null!;

        /// <summary>
        /// The Project ID of Service Endpoint belongs to.
        /// </summary>
        [Input("serviceEndpointScope", required: true)]
        public Input<string> ServiceEndpointScope { get; set; } = null!;

        /// <summary>
        /// Delay in minutes before deleting excess idle agents. Defaults to `30`.
        /// </summary>
        [Input("timeToLiveMinutes")]
        public Input<int>? TimeToLiveMinutes { get; set; }

        public ElasticPoolArgs()
        {
        }
        public static new ElasticPoolArgs Empty => new ElasticPoolArgs();
    }

    public sealed class ElasticPoolState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Set whether agents should be configured to run with interactive UI. Defaults to `false`.
        /// </summary>
        [Input("agentInteractiveUi")]
        public Input<bool>? AgentInteractiveUi { get; set; }

        /// <summary>
        /// Specifies whether a queue should be automatically provisioned for each project collection. Defaults to `false`.
        /// </summary>
        [Input("autoProvision")]
        public Input<bool>? AutoProvision { get; set; }

        /// <summary>
        /// Specifies whether or not agents within the pool should be automatically updated. Defaults to `true`.
        /// </summary>
        [Input("autoUpdate")]
        public Input<bool>? AutoUpdate { get; set; }

        /// <summary>
        /// The ID of the Azure resource.
        /// </summary>
        [Input("azureResourceId")]
        public Input<string>? AzureResourceId { get; set; }

        /// <summary>
        /// Number of agents to keep on standby.
        /// </summary>
        [Input("desiredIdle")]
        public Input<int>? DesiredIdle { get; set; }

        /// <summary>
        /// Maximum number of virtual machines in the scale set.
        /// </summary>
        [Input("maxCapacity")]
        public Input<int>? MaxCapacity { get; set; }

        /// <summary>
        /// The name of the Elastic pool.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Tear down virtual machines after every use. Defaults to `false`.
        /// </summary>
        [Input("recycleAfterEachUse")]
        public Input<bool>? RecycleAfterEachUse { get; set; }

        /// <summary>
        /// The ID of Service Endpoint used to connect to Azure.
        /// </summary>
        [Input("serviceEndpointId")]
        public Input<string>? ServiceEndpointId { get; set; }

        /// <summary>
        /// The Project ID of Service Endpoint belongs to.
        /// </summary>
        [Input("serviceEndpointScope")]
        public Input<string>? ServiceEndpointScope { get; set; }

        /// <summary>
        /// Delay in minutes before deleting excess idle agents. Defaults to `30`.
        /// </summary>
        [Input("timeToLiveMinutes")]
        public Input<int>? TimeToLiveMinutes { get; set; }

        public ElasticPoolState()
        {
        }
        public static new ElasticPoolState Empty => new ElasticPoolState();
    }
}
