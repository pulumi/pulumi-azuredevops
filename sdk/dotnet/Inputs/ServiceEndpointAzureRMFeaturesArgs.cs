// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureDevOps.Inputs
{

    public sealed class ServiceEndpointAzureRMFeaturesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether or not to validate connection with Azure after create or update operations. Defaults to `false`
        /// </summary>
        [Input("validate")]
        public Input<bool>? Validate { get; set; }

        public ServiceEndpointAzureRMFeaturesArgs()
        {
        }
        public static new ServiceEndpointAzureRMFeaturesArgs Empty => new ServiceEndpointAzureRMFeaturesArgs();
    }
}