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
    public sealed class ServiceendpointExternaltfsAuthPersonal
    {
        /// <summary>
        /// The Personal Access Token for Azure DevOps Organization.
        /// </summary>
        public readonly string PersonalAccessToken;

        [OutputConstructor]
        private ServiceendpointExternaltfsAuthPersonal(string personalAccessToken)
        {
            PersonalAccessToken = personalAccessToken;
        }
    }
}
