// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureDevOps.Outputs
{

    [OutputType]
    public sealed class GetBuildDefinitionJobTargetExecutionOptionResult
    {
        /// <summary>
        /// Whether to continue the job when an error occurs.
        /// </summary>
        public readonly bool ContinueOnError;
        /// <summary>
        /// Limit the number of agents to be used. If job type is `AgentlessJob`, the concurrency is not configurable and is fixed to 50.
        /// </summary>
        public readonly int MaxConcurrency;
        /// <summary>
        /// A list of comma separated configuration variables to use. These are defined on the Variables tab. For example, OperatingSystem, Browser will run the tasks for both variables.
        /// </summary>
        public readonly string Multipliers;
        /// <summary>
        /// The execution type of the Job.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetBuildDefinitionJobTargetExecutionOptionResult(
            bool continueOnError,

            int maxConcurrency,

            string multipliers,

            string type)
        {
            ContinueOnError = continueOnError;
            MaxConcurrency = maxConcurrency;
            Multipliers = multipliers;
            Type = type;
        }
    }
}
