// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureDevOps
{
    public static class GetIdentityGroups
    {
        /// <summary>
        /// Use this data source to access information about existing Groups within Azure DevOps On-Premise(Azure DevOps Server).
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using AzureDevOps = Pulumi.AzureDevOps;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = AzureDevOps.GetProject.Invoke(new()
        ///     {
        ///         Name = "Example Project",
        ///     });
        /// 
        ///     // load all existing groups inside an organization
        ///     var example_all_groups = AzureDevOps.GetIdentityGroups.Invoke();
        /// 
        ///     // load all existing groups inside a specific project
        ///     var example_project_groups = AzureDevOps.GetIdentityGroups.Invoke(new()
        ///     {
        ///         ProjectId = example.Apply(getProjectResult =&gt; getProjectResult.Id),
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// ## Relevant Links
        /// 
        /// - [Azure DevOps Service REST API 7.0 - Identities](https://docs.microsoft.com/en-us/rest/api/azure/devops/ims/?view=azure-devops-rest-7.2)
        /// </summary>
        public static Task<GetIdentityGroupsResult> InvokeAsync(GetIdentityGroupsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetIdentityGroupsResult>("azuredevops:index/getIdentityGroups:getIdentityGroups", args ?? new GetIdentityGroupsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to access information about existing Groups within Azure DevOps On-Premise(Azure DevOps Server).
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using AzureDevOps = Pulumi.AzureDevOps;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = AzureDevOps.GetProject.Invoke(new()
        ///     {
        ///         Name = "Example Project",
        ///     });
        /// 
        ///     // load all existing groups inside an organization
        ///     var example_all_groups = AzureDevOps.GetIdentityGroups.Invoke();
        /// 
        ///     // load all existing groups inside a specific project
        ///     var example_project_groups = AzureDevOps.GetIdentityGroups.Invoke(new()
        ///     {
        ///         ProjectId = example.Apply(getProjectResult =&gt; getProjectResult.Id),
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// ## Relevant Links
        /// 
        /// - [Azure DevOps Service REST API 7.0 - Identities](https://docs.microsoft.com/en-us/rest/api/azure/devops/ims/?view=azure-devops-rest-7.2)
        /// </summary>
        public static Output<GetIdentityGroupsResult> Invoke(GetIdentityGroupsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetIdentityGroupsResult>("azuredevops:index/getIdentityGroups:getIdentityGroups", args ?? new GetIdentityGroupsInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to access information about existing Groups within Azure DevOps On-Premise(Azure DevOps Server).
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using AzureDevOps = Pulumi.AzureDevOps;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = AzureDevOps.GetProject.Invoke(new()
        ///     {
        ///         Name = "Example Project",
        ///     });
        /// 
        ///     // load all existing groups inside an organization
        ///     var example_all_groups = AzureDevOps.GetIdentityGroups.Invoke();
        /// 
        ///     // load all existing groups inside a specific project
        ///     var example_project_groups = AzureDevOps.GetIdentityGroups.Invoke(new()
        ///     {
        ///         ProjectId = example.Apply(getProjectResult =&gt; getProjectResult.Id),
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// ## Relevant Links
        /// 
        /// - [Azure DevOps Service REST API 7.0 - Identities](https://docs.microsoft.com/en-us/rest/api/azure/devops/ims/?view=azure-devops-rest-7.2)
        /// </summary>
        public static Output<GetIdentityGroupsResult> Invoke(GetIdentityGroupsInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetIdentityGroupsResult>("azuredevops:index/getIdentityGroups:getIdentityGroups", args ?? new GetIdentityGroupsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetIdentityGroupsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Project ID. If no project ID is specified all groups of an organization will be returned
        /// </summary>
        [Input("projectId")]
        public string? ProjectId { get; set; }

        public GetIdentityGroupsArgs()
        {
        }
        public static new GetIdentityGroupsArgs Empty => new GetIdentityGroupsArgs();
    }

    public sealed class GetIdentityGroupsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Project ID. If no project ID is specified all groups of an organization will be returned
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        public GetIdentityGroupsInvokeArgs()
        {
        }
        public static new GetIdentityGroupsInvokeArgs Empty => new GetIdentityGroupsInvokeArgs();
    }


    [OutputType]
    public sealed class GetIdentityGroupsResult
    {
        /// <summary>
        /// A `groups` blocks as documented below. A set of existing groups in your Azure DevOps Organization or project with details about every single group.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetIdentityGroupsGroupResult> Groups;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? ProjectId;

        [OutputConstructor]
        private GetIdentityGroupsResult(
            ImmutableArray<Outputs.GetIdentityGroupsGroupResult> groups,

            string id,

            string? projectId)
        {
            Groups = groups;
            Id = id;
            ProjectId = projectId;
        }
    }
}
