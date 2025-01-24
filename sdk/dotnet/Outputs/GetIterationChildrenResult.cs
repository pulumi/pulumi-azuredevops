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
    public sealed class GetIterationChildrenResult
    {
        /// <summary>
        /// Indicator if the child Iteration node has child nodes
        /// </summary>
        public readonly bool HasChildren;
        /// <summary>
        /// The ID of the child Iteration node
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the child Iteration node
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The path to the Iteration, _Format_: URL relative; if omitted, or value `"/"` is used, the root Iteration will be returned
        /// </summary>
        public readonly string Path;
        /// <summary>
        /// The project ID.
        /// </summary>
        public readonly string ProjectId;

        [OutputConstructor]
        private GetIterationChildrenResult(
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
