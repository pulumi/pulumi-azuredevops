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
    public sealed class GetProjectsProjectResult
    {
        /// <summary>
        /// Name of the Project, if not specified all projects will be returned.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The ID of the Project.
        /// </summary>
        public readonly string ProjectId;
        /// <summary>
        /// Url to the full version of the object.
        /// </summary>
        public readonly string ProjectUrl;
        /// <summary>
        /// State of the Project, if not specified all projects will be returned. Valid values are `all`, `deleting`, `new`, `wellFormed`, `createPending`, `unchanged`,`deleted`.
        /// </summary>
        public readonly string State;

        [OutputConstructor]
        private GetProjectsProjectResult(
            string name,

            string projectId,

            string projectUrl,

            string state)
        {
            Name = name;
            ProjectId = projectId;
            ProjectUrl = projectUrl;
            State = state;
        }
    }
}
