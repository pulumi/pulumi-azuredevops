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
    public sealed class GetAreaChildrenResult
    {
        public readonly bool HasChildren;
        public readonly string Id;
        public readonly string Name;
        public readonly string Path;
        public readonly string ProjectId;

        [OutputConstructor]
        private GetAreaChildrenResult(
            bool hasChildren,

            string id,

            string name,

            string path,

            string projectId)
        {
            HasChildren = hasChildren;
            Id = id;
            Name = name;
            Path = path;
            ProjectId = projectId;
        }
    }
}
