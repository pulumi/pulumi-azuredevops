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
    /// Manages an agent pool within Azure DevOps.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using AzureDevOps = Pulumi.AzureDevOps;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var pool = new AzureDevOps.Pool("pool", new AzureDevOps.PoolArgs
    ///         {
    ///             AutoProvision = false,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ## Relevant Links
    /// 
    /// - [Azure DevOps Service REST API 5.1 - Agent Pools](https://docs.microsoft.com/en-us/rest/api/azure/devops/distributedtask/pools?view=azure-devops-rest-5.1)
    /// </summary>
    public partial class Pool : Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies whether or not a queue should be automatically provisioned for each project collection. Defaults to `false`.
        /// </summary>
        [Output("autoProvision")]
        public Output<bool?> AutoProvision { get; private set; } = null!;

        /// <summary>
        /// The name of the agent pool.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specifies whether the agent pool type is Automation or Deployment. Defaults to `automation`.
        /// </summary>
        [Output("poolType")]
        public Output<string?> PoolType { get; private set; } = null!;


        /// <summary>
        /// Create a Pool resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Pool(string name, PoolArgs? args = null, CustomResourceOptions? options = null)
            : base("azuredevops:index/pool:Pool", name, args ?? new PoolArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Pool(string name, Input<string> id, PoolState? state = null, CustomResourceOptions? options = null)
            : base("azuredevops:index/pool:Pool", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azuredevops:Agent/pool:Pool"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Pool resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Pool Get(string name, Input<string> id, PoolState? state = null, CustomResourceOptions? options = null)
        {
            return new Pool(name, id, state, options);
        }
    }

    public sealed class PoolArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies whether or not a queue should be automatically provisioned for each project collection. Defaults to `false`.
        /// </summary>
        [Input("autoProvision")]
        public Input<bool>? AutoProvision { get; set; }

        /// <summary>
        /// The name of the agent pool.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies whether the agent pool type is Automation or Deployment. Defaults to `automation`.
        /// </summary>
        [Input("poolType")]
        public Input<string>? PoolType { get; set; }

        public PoolArgs()
        {
        }
    }

    public sealed class PoolState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies whether or not a queue should be automatically provisioned for each project collection. Defaults to `false`.
        /// </summary>
        [Input("autoProvision")]
        public Input<bool>? AutoProvision { get; set; }

        /// <summary>
        /// The name of the agent pool.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies whether the agent pool type is Automation or Deployment. Defaults to `automation`.
        /// </summary>
        [Input("poolType")]
        public Input<string>? PoolType { get; set; }

        public PoolState()
        {
        }
    }
}
