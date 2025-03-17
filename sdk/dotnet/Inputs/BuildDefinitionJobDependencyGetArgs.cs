// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureDevOps.Inputs
{

    public sealed class BuildDefinitionJobDependencyGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The job reference name that depends on. Reference to `jobs.ref_name`
        /// </summary>
        [Input("scope", required: true)]
        public Input<string> Scope { get; set; } = null!;

        public BuildDefinitionJobDependencyGetArgs()
        {
        }
        public static new BuildDefinitionJobDependencyGetArgs Empty => new BuildDefinitionJobDependencyGetArgs();
    }
}
