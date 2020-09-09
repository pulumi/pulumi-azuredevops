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
    public sealed class AzureRMCredentials
    {
        /// <summary>
        /// The service principal application Id
        /// </summary>
        public readonly string Serviceprincipalid;
        /// <summary>
        /// The service principal secret.
        /// </summary>
        public readonly string Serviceprincipalkey;
        public readonly string? ServiceprincipalkeyHash;

        [OutputConstructor]
        private AzureRMCredentials(
            string serviceprincipalid,

            string serviceprincipalkey,

            string? serviceprincipalkeyHash)
        {
            Serviceprincipalid = serviceprincipalid;
            Serviceprincipalkey = serviceprincipalkey;
            ServiceprincipalkeyHash = serviceprincipalkeyHash;
        }
    }
}
