// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureDevOps.Agent.Outputs
{

    [OutputType]
    public sealed class GetPoolsAgentPoolResult
    {
        public readonly bool AutoProvision;
        public readonly int Id;
        public readonly string Name;
        public readonly string PoolType;

        [OutputConstructor]
        private GetPoolsAgentPoolResult(
            bool autoProvision,

            int id,

            string name,

            string poolType)
        {
            AutoProvision = autoProvision;
            Id = id;
            Name = name;
            PoolType = poolType;
        }
    }
}
