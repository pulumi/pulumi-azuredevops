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
    public sealed class BuildDefinitionJobDependency
    {
        /// <summary>
        /// The job reference name that depends on. Reference to `jobs.ref_name`
        /// </summary>
        public readonly string Scope;

        [OutputConstructor]
        private BuildDefinitionJobDependency(string scope)
        {
            Scope = scope;
        }
    }
}
