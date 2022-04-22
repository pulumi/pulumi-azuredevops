// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureDevOps.Build.Outputs
{

    [OutputType]
    public sealed class BuildDefinitionScheduleBranchFilter
    {
        public readonly ImmutableArray<string> Excludes;
        public readonly ImmutableArray<string> Includes;

        [OutputConstructor]
        private BuildDefinitionScheduleBranchFilter(
            ImmutableArray<string> excludes,

            ImmutableArray<string> includes)
        {
            Excludes = excludes;
            Includes = includes;
        }
    }
}
