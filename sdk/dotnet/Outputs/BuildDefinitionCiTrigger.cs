// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureDevOps.Outputs
{

    [OutputType]
    public sealed class BuildDefinitionCiTrigger
    {
        /// <summary>
        /// Override the azure-pipeline file and use a this configuration for all builds.
        /// </summary>
        public readonly Outputs.BuildDefinitionCiTriggerOverride? Override;
        /// <summary>
        /// Use the azure-pipeline file for the build configuration. Defaults to `false`.
        /// </summary>
        public readonly bool? UseYaml;

        [OutputConstructor]
        private BuildDefinitionCiTrigger(
            Outputs.BuildDefinitionCiTriggerOverride? @override,

            bool? useYaml)
        {
            Override = @override;
            UseYaml = useYaml;
        }
    }
}
