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
    public sealed class GetBuildDefinitionCiTriggerResult
    {
        /// <summary>
        /// A `override` block as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetBuildDefinitionCiTriggerOverrideResult> Overrides;
        /// <summary>
        /// Use the azure-pipeline file for the build configuration.
        /// </summary>
        public readonly bool UseYaml;

        [OutputConstructor]
        private GetBuildDefinitionCiTriggerResult(
            ImmutableArray<Outputs.GetBuildDefinitionCiTriggerOverrideResult> overrides,

            bool useYaml)
        {
            Overrides = overrides;
            UseYaml = useYaml;
        }
    }
}
