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
    public sealed class BuildDefinitionRepository
    {
        /// <summary>
        /// The branch name for which builds are triggered. Defaults to `master`.
        /// </summary>
        public readonly string? BranchName;
        /// <summary>
        /// The Github Enterprise URL. Used if `repo_type` is `GithubEnterprise`.
        /// </summary>
        public readonly string? GithubEnterpriseUrl;
        /// <summary>
        /// The id of the repository. For `TfsGit` repos, this is simply the ID of the repository. For `Github` repos, this will take the form of `&lt;GitHub Org&gt;/&lt;Repo Name&gt;`. For `Bitbucket` repos, this will take the form of `&lt;Workspace ID&gt;/&lt;Repo Name&gt;`.
        /// </summary>
        public readonly string RepoId;
        /// <summary>
        /// The repository type. Possible values are: `GitHub` or `TfsGit` or `Bitbucket` or `GitHub Enterprise`. Defaults to `GitHub`. If `repo_type` is `GitHubEnterprise`, must use existing project and GitHub Enterprise service connection.
        /// </summary>
        public readonly string RepoType;
        /// <summary>
        /// Report build status. Default is true.
        /// </summary>
        public readonly bool? ReportBuildStatus;
        /// <summary>
        /// The service connection ID. Used if the `repo_type` is `GitHub` or `GitHubEnterprise`.
        /// </summary>
        public readonly string? ServiceConnectionId;
        /// <summary>
        /// The path of the Yaml file describing the build definition.
        /// </summary>
        public readonly string YmlPath;

        [OutputConstructor]
        private BuildDefinitionRepository(
            string? branchName,

            string? githubEnterpriseUrl,

            string repoId,

            string repoType,

            bool? reportBuildStatus,

            string? serviceConnectionId,

            string ymlPath)
        {
            BranchName = branchName;
            GithubEnterpriseUrl = githubEnterpriseUrl;
            RepoId = repoId;
            RepoType = repoType;
            ReportBuildStatus = reportBuildStatus;
            ServiceConnectionId = serviceConnectionId;
            YmlPath = ymlPath;
        }
    }
}
