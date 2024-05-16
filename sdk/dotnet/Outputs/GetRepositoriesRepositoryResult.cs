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
    public sealed class GetRepositoriesRepositoryResult
    {
        /// <summary>
        /// The ref of the default branch.
        /// </summary>
        public readonly string DefaultBranch;
        /// <summary>
        /// Is the repository disabled?
        /// </summary>
        public readonly bool Disabled;
        /// <summary>
        /// Git repository identifier.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Name of the Git repository to retrieve; requires `project_id` to be specified as well
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// ID of project to list Git repositories
        /// </summary>
        public readonly string ProjectId;
        /// <summary>
        /// HTTPS Url to clone the Git repository
        /// </summary>
        public readonly string RemoteUrl;
        /// <summary>
        /// Compressed size (bytes) of the repository.
        /// </summary>
        public readonly int Size;
        /// <summary>
        /// SSH Url to clone the Git repository
        /// </summary>
        public readonly string SshUrl;
        /// <summary>
        /// Details REST API endpoint for the Git Repository.
        /// </summary>
        public readonly string Url;
        /// <summary>
        /// Url of the Git repository web view
        /// </summary>
        public readonly string WebUrl;

        [OutputConstructor]
        private GetRepositoriesRepositoryResult(
            string defaultBranch,

            bool disabled,

            string id,

            string name,

            string projectId,

            string remoteUrl,

            int size,

            string sshUrl,

            string url,

            string webUrl)
        {
            DefaultBranch = defaultBranch;
            Disabled = disabled;
            Id = id;
            Name = name;
            ProjectId = projectId;
            RemoteUrl = remoteUrl;
            Size = size;
            SshUrl = sshUrl;
            Url = url;
            WebUrl = webUrl;
        }
    }
}
