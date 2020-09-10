// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureDevOps.Inputs
{

    public sealed class BuildDefinitionRepositoryArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The branch name for which builds are triggered. Defaults to `master`.
        /// </summary>
        [Input("branchName")]
        public Input<string>? BranchName { get; set; }

        /// <summary>
        /// The Github Enterprise URL. Used if `repo_type` is `GithubEnterprise`.
        /// </summary>
        [Input("githubEnterpriseUrl")]
        public Input<string>? GithubEnterpriseUrl { get; set; }

        /// <summary>
        /// The id of the repository. For `TfsGit` repos, this is simply the ID of the repository. For `Github` repos, this will take the form of `&lt;GitHub Org&gt;/&lt;Repo Name&gt;`. For `Bitbucket` repos, this will take the form of `&lt;Workspace ID&gt;/&lt;Repo Name&gt;`.
        /// </summary>
        [Input("repoId", required: true)]
        public Input<string> RepoId { get; set; } = null!;

        /// <summary>
        /// The repository type. Valid values: `GitHub` or `TfsGit` or `Bitbucket` or `GitHub Enterprise`. Defaults to `Github`. If `repo_type` is `GitHubEnterprise`, must use existing project and GitHub Enterprise service connection.
        /// </summary>
        [Input("repoType", required: true)]
        public Input<string> RepoType { get; set; } = null!;

        /// <summary>
        /// Report build status. Default is true.
        /// </summary>
        [Input("reportBuildStatus")]
        public Input<bool>? ReportBuildStatus { get; set; }

        /// <summary>
        /// The service connection ID. Used if the `repo_type` is `GitHub` or `GitHubEnterprise`.
        /// </summary>
        [Input("serviceConnectionId")]
        public Input<string>? ServiceConnectionId { get; set; }

        /// <summary>
        /// The path of the Yaml file describing the build definition.
        /// </summary>
        [Input("ymlPath", required: true)]
        public Input<string> YmlPath { get; set; } = null!;

        public BuildDefinitionRepositoryArgs()
        {
        }
    }
}
