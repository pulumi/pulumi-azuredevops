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
    public sealed class GetBuildDefinitionJobResult
    {
        /// <summary>
        /// Enables scripts and other processes launched by tasks to access the OAuth token through the `System.AccessToken` variable.
        /// </summary>
        public readonly bool AllowScriptsAuthAccessOption;
        /// <summary>
        /// Specifies when this job should run. Can **Custom conditions** to specify more complex conditions. More details: [Pipeline conditions](https://learn.microsoft.com/en-us/azure/devops/pipelines/process/conditions?view=azure-devops)
        /// </summary>
        public readonly string Condition;
        /// <summary>
        /// A `dependencies` blocks as documented below. Define the job dependencies.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetBuildDefinitionJobDependencyResult> Dependencies;
        /// <summary>
        /// The job authorization scope for builds queued against this definition.
        /// </summary>
        public readonly string JobAuthorizationScope;
        /// <summary>
        /// The job cancel timeout (in minutes) for builds cancelled by user for this definition.
        /// </summary>
        public readonly int JobCancelTimeoutInMinutes;
        /// <summary>
        /// The job execution timeout (in minutes) for builds queued against this definition.
        /// </summary>
        public readonly int JobTimeoutInMinutes;
        /// <summary>
        /// The name of this Build Definition.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The reference name of the job, can be used to define the job dependencies.
        /// </summary>
        public readonly string RefName;
        /// <summary>
        /// A `target` blocks as documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetBuildDefinitionJobTargetResult> Targets;

        [OutputConstructor]
        private GetBuildDefinitionJobResult(
            bool allowScriptsAuthAccessOption,

            string condition,

            ImmutableArray<Outputs.GetBuildDefinitionJobDependencyResult> dependencies,

            string jobAuthorizationScope,

            int jobCancelTimeoutInMinutes,

            int jobTimeoutInMinutes,

            string name,

            string refName,

            ImmutableArray<Outputs.GetBuildDefinitionJobTargetResult> targets)
        {
            AllowScriptsAuthAccessOption = allowScriptsAuthAccessOption;
            Condition = condition;
            Dependencies = dependencies;
            JobAuthorizationScope = jobAuthorizationScope;
            JobCancelTimeoutInMinutes = jobCancelTimeoutInMinutes;
            JobTimeoutInMinutes = jobTimeoutInMinutes;
            Name = name;
            RefName = refName;
            Targets = targets;
        }
    }
}
