// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureDevOps.Inputs
{

    public sealed class ServiceEndpointServiceFabricCertificateArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Base64 encoding of the cluster's client certificate file.
        /// </summary>
        [Input("clientCertificate", required: true)]
        public Input<string> ClientCertificate { get; set; } = null!;

        [Input("clientCertificateHash")]
        public Input<string>? ClientCertificateHash { get; set; }

        /// <summary>
        /// Password for the certificate.
        /// </summary>
        [Input("clientCertificatePassword")]
        public Input<string>? ClientCertificatePassword { get; set; }

        [Input("clientCertificatePasswordHash")]
        public Input<string>? ClientCertificatePasswordHash { get; set; }

        /// <summary>
        /// The common name(s) of the cluster's certificate(s). This is used to verify the identity of the cluster. This value overrides the publish profile. Separate multiple common names with a comma (',')
        /// </summary>
        [Input("serverCertificateCommonName")]
        public Input<string>? ServerCertificateCommonName { get; set; }

        /// <summary>
        /// Verification mode for the cluster. Possible values include `Thumbprint` or `CommonName`.
        /// </summary>
        [Input("serverCertificateLookup", required: true)]
        public Input<string> ServerCertificateLookup { get; set; } = null!;

        /// <summary>
        /// The thumbprint(s) of the cluster's certificate(s). This is used to verify the identity of the cluster. This value overrides the publish profile. Separate multiple thumbprints with a comma (',')
        /// </summary>
        [Input("serverCertificateThumbprint")]
        public Input<string>? ServerCertificateThumbprint { get; set; }

        public ServiceEndpointServiceFabricCertificateArgs()
        {
        }
    }
}
