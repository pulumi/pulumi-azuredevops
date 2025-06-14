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
    public sealed class GetBuildDefinitionJobTargetResult
    {
        /// <summary>
        /// A list of demands that represents the agent capabilities required by this build. Example: `git`
        /// </summary>
        public readonly ImmutableArray<string> Demands;
        /// <summary>
        /// A `execution_options` blocks as documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetBuildDefinitionJobTargetExecutionOptionResult> ExecutionOptions;
        /// <summary>
        /// The execution type of the Job.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetBuildDefinitionJobTargetResult(
            ImmutableArray<string> demands,

            ImmutableArray<Outputs.GetBuildDefinitionJobTargetExecutionOptionResult> executionOptions,

            string type)
        {
            Demands = demands;
            ExecutionOptions = executionOptions;
            Type = type;
        }
    }
}
