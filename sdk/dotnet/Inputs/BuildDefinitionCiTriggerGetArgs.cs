// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureDevOps.Inputs
{

    public sealed class BuildDefinitionCiTriggerGetArgs : Pulumi.ResourceArgs
    {
        [Input("override")]
        public Input<Inputs.BuildDefinitionCiTriggerOverrideGetArgs>? Override { get; set; }

        [Input("useYaml")]
        public Input<bool>? UseYaml { get; set; }

        public BuildDefinitionCiTriggerGetArgs()
        {
        }
    }
}
