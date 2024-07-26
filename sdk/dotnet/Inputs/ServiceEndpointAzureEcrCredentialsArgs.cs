// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureDevOps.Inputs
{

    public sealed class ServiceEndpointAzureEcrCredentialsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The service principal application Id
        /// </summary>
        [Input("serviceprincipalid", required: true)]
        public Input<string> Serviceprincipalid { get; set; } = null!;

        public ServiceEndpointAzureEcrCredentialsArgs()
        {
        }
        public static new ServiceEndpointAzureEcrCredentialsArgs Empty => new ServiceEndpointAzureEcrCredentialsArgs();
    }
}
