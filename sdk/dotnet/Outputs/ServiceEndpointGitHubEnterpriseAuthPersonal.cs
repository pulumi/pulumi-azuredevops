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
    public sealed class ServiceEndpointGitHubEnterpriseAuthPersonal
    {
        /// <summary>
        /// The Personal Access Token for Github.
        /// </summary>
        public readonly string PersonalAccessToken;
        public readonly string? PersonalAccessTokenHash;

        [OutputConstructor]
        private ServiceEndpointGitHubEnterpriseAuthPersonal(
            string personalAccessToken,

            string? personalAccessTokenHash)
        {
            PersonalAccessToken = personalAccessToken;
            PersonalAccessTokenHash = personalAccessTokenHash;
        }
    }
}
