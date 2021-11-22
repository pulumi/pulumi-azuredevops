// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.AzureDevOps
{
    public static class GetTeams
    {
        public static Task<GetTeamsResult> InvokeAsync(GetTeamsArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetTeamsResult>("azuredevops:index/getTeams:getTeams", args ?? new GetTeamsArgs(), options.WithVersion());

        public static Output<GetTeamsResult> Invoke(GetTeamsInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetTeamsResult>("azuredevops:index/getTeams:getTeams", args ?? new GetTeamsInvokeArgs(), options.WithVersion());
    }


    public sealed class GetTeamsArgs : Pulumi.InvokeArgs
    {
        [Input("projectId")]
        public string? ProjectId { get; set; }

        public GetTeamsArgs()
        {
        }
    }

    public sealed class GetTeamsInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        public GetTeamsInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetTeamsResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? ProjectId;
        public readonly ImmutableArray<Outputs.GetTeamsTeamResult> Teams;

        [OutputConstructor]
        private GetTeamsResult(
            string id,

            string? projectId,

            ImmutableArray<Outputs.GetTeamsTeamResult> teams)
        {
            Id = id;
            ProjectId = projectId;
            Teams = teams;
        }
    }
}
