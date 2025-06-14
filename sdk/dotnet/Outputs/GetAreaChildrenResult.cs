// *** WARNING: this file was generated by pulumi-language-dotnet. ***
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
        /// <summary>
        /// Indicator if the child Area node has child nodes
        /// </summary>
        public readonly bool HasChildren;
        /// <summary>
        /// The ID of the child Area node
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the child Area node
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The path to the Area; _Format_: URL relative; if omitted, or value `"/"` is used, the root Area will be returned
        /// </summary>
        public readonly string Path;
        /// <summary>
        /// The project ID.
        /// </summary>
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
