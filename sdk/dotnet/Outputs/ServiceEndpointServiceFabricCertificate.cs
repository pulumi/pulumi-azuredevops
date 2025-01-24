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
    public sealed class ServiceEndpointServiceFabricCertificate
    {
        /// <summary>
        /// Base64 encoding of the cluster's client certificate file.
        /// </summary>
        public readonly string ClientCertificate;
        /// <summary>
        /// Password for the certificate.
        /// </summary>
        public readonly string? ClientCertificatePassword;
        /// <summary>
        /// The common name(s) of the cluster's certificate(s). This is used to verify the identity of the cluster. This value overrides the publish profile. Separate multiple common names with a comma (',')
        /// </summary>
        public readonly string? ServerCertificateCommonName;
        /// <summary>
        /// Verification mode for the cluster. Possible values are: `Thumbprint`, `CommonName`.
        /// </summary>
        public readonly string ServerCertificateLookup;
        /// <summary>
        /// The thumbprint(s) of the cluster's certificate(s). This is used to verify the identity of the cluster. This value overrides the publish profile. Separate multiple thumbprints with a comma (',')
        /// </summary>
        public readonly string? ServerCertificateThumbprint;

        [OutputConstructor]
        private ServiceEndpointServiceFabricCertificate(
            string clientCertificate,

            string? clientCertificatePassword,

            string? serverCertificateCommonName,

            string serverCertificateLookup,

            string? serverCertificateThumbprint)
        {
            ClientCertificate = clientCertificate;
            ClientCertificatePassword = clientCertificatePassword;
            ServerCertificateCommonName = serverCertificateCommonName;
            ServerCertificateLookup = serverCertificateLookup;
            ServerCertificateThumbprint = serverCertificateThumbprint;
        }
    }
}
