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
    public sealed class RepositoryPolicyAuthorEmailPatternSettingsScope
    {
        /// <summary>
        /// The repository ID.
        /// </summary>
        public readonly string RepositoryId;

        [OutputConstructor]
        private RepositoryPolicyAuthorEmailPatternSettingsScope(string repositoryId)
        {
            RepositoryId = repositoryId;
        }
    }
}
