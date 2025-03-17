// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureDevOps.Inputs
{

    public sealed class BuildDefinitionJobTargetExecutionOptionsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to continue the job when an error occurs. Possible values are: `true`, `false`.
        /// </summary>
        [Input("continueOnError")]
        public Input<bool>? ContinueOnError { get; set; }

        /// <summary>
        /// Limit the number of agents to be used. If job type is `AgentlessJob`, the concurrency is not configurable and is fixed to 50.
        /// </summary>
        [Input("maxConcurrency")]
        public Input<int>? MaxConcurrency { get; set; }

        /// <summary>
        /// A list of comma separated configuration variables to use. These are defined on the Variables tab. For example, OperatingSystem, Browser will run the tasks for both variables. Available when `execution_options.type` is `Multi-Configuration`.
        /// </summary>
        [Input("multipliers")]
        public Input<string>? Multipliers { get; set; }

        /// <summary>
        /// The execution type of the Job. Possible values are: `None`, `Multi-Configuration`, `Multi-Agent`.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public BuildDefinitionJobTargetExecutionOptionsArgs()
        {
        }
        public static new BuildDefinitionJobTargetExecutionOptionsArgs Empty => new BuildDefinitionJobTargetExecutionOptionsArgs();
    }
}
