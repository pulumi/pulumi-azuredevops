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
    public sealed class ServiceEndpointKubernetesAzureSubscription
    {
        /// <summary>
        /// Azure environment refers to whether the public cloud offering or domestic (government) clouds are being used. Currently, only the public cloud is supported. The value must be AzureCloud. This is also the default-value.
        /// </summary>
        public readonly string? AzureEnvironment;
        /// <summary>
        /// Set this option to allow use cluster admin credentials.
        /// </summary>
        public readonly bool? ClusterAdmin;
        /// <summary>
        /// The name of the Kubernetes cluster.
        /// </summary>
        public readonly string ClusterName;
        /// <summary>
        /// The Kubernetes namespace. Default value is "default".
        /// </summary>
        public readonly string? Namespace;
        /// <summary>
        /// The resource group name, to which the Kubernetes cluster is deployed.
        /// </summary>
        public readonly string ResourcegroupId;
        /// <summary>
        /// The id of the Azure subscription.
        /// </summary>
        public readonly string SubscriptionId;
        /// <summary>
        /// The name of the Azure subscription.
        /// </summary>
        public readonly string SubscriptionName;
        /// <summary>
        /// The id of the tenant used by the subscription.
        /// </summary>
        public readonly string TenantId;

        [OutputConstructor]
        private ServiceEndpointKubernetesAzureSubscription(
            string? azureEnvironment,

            bool? clusterAdmin,

            string clusterName,

            string? @namespace,

            string resourcegroupId,

            string subscriptionId,

            string subscriptionName,

            string tenantId)
        {
            AzureEnvironment = azureEnvironment;
            ClusterAdmin = clusterAdmin;
            ClusterName = clusterName;
            Namespace = @namespace;
            ResourcegroupId = resourcegroupId;
            SubscriptionId = subscriptionId;
            SubscriptionName = subscriptionName;
            TenantId = tenantId;
        }
    }
}
