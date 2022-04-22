// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureDevOps.Inputs
{

    public sealed class BuildDefinitionRepositoryGetArgs : Pulumi.ResourceArgs
    {
        [Input("branchName")]
        public Input<string>? BranchName { get; set; }

        [Input("githubEnterpriseUrl")]
        public Input<string>? GithubEnterpriseUrl { get; set; }

        [Input("repoId", required: true)]
        public Input<string> RepoId { get; set; } = null!;

        [Input("repoType", required: true)]
        public Input<string> RepoType { get; set; } = null!;

        [Input("reportBuildStatus")]
        public Input<bool>? ReportBuildStatus { get; set; }

        [Input("serviceConnectionId")]
        public Input<string>? ServiceConnectionId { get; set; }

        [Input("ymlPath", required: true)]
        public Input<string> YmlPath { get; set; } = null!;

        public BuildDefinitionRepositoryGetArgs()
        {
        }
    }
}
