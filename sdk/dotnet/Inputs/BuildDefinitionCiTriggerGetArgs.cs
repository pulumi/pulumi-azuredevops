// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureDevOps.Inputs
{

    public sealed class BuildDefinitionCiTriggerGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Override the azure-pipeline file and use a this configuration for all builds.
        /// </summary>
        [Input("override")]
        public Input<Inputs.BuildDefinitionCiTriggerOverrideGetArgs>? Override { get; set; }

        /// <summary>
        /// Use the azure-pipeline file for the build configuration. Defaults to `false`.
        /// </summary>
        [Input("useYaml")]
        public Input<bool>? UseYaml { get; set; }

        public BuildDefinitionCiTriggerGetArgs()
        {
        }
        public static new BuildDefinitionCiTriggerGetArgs Empty => new BuildDefinitionCiTriggerGetArgs();
    }
}
