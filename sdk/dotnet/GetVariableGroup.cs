// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureDevOps
{
    public static class GetVariableGroup
    {
        /// <summary>
        /// Use this data source to access information about existing Variable Groups within Azure DevOps.
        /// 
        /// &gt; **Note:** Secret values are masked by service and cannot be obtained through API. [Set secret variables](https://docs.microsoft.com/en-us/azure/devops/pipelines/process/variables?view=azure-devops&amp;tabs=yaml%2Cbatch#secret-variables)
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using AzureDevOps = Pulumi.AzureDevOps;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var exampleProject = AzureDevOps.GetProject.Invoke(new()
        ///     {
        ///         Name = "Example Project",
        ///     });
        /// 
        ///     var exampleVariableGroup = AzureDevOps.GetVariableGroup.Invoke(new()
        ///     {
        ///         ProjectId = exampleProject.Apply(getProjectResult =&gt; getProjectResult.Id),
        ///         Name = "Example Variable Group",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["id"] = exampleVariableGroup.Apply(getVariableGroupResult =&gt; getVariableGroupResult.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// ## Relevant Links
        /// 
        /// - [Azure DevOps Service REST API 6.0 - Variable Groups](https://docs.microsoft.com/en-us/rest/api/azure/devops/distributedtask/variablegroups?view=azure-devops-rest-6.0)
        /// </summary>
        public static Task<GetVariableGroupResult> InvokeAsync(GetVariableGroupArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetVariableGroupResult>("azuredevops:index/getVariableGroup:getVariableGroup", args ?? new GetVariableGroupArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to access information about existing Variable Groups within Azure DevOps.
        /// 
        /// &gt; **Note:** Secret values are masked by service and cannot be obtained through API. [Set secret variables](https://docs.microsoft.com/en-us/azure/devops/pipelines/process/variables?view=azure-devops&amp;tabs=yaml%2Cbatch#secret-variables)
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using AzureDevOps = Pulumi.AzureDevOps;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var exampleProject = AzureDevOps.GetProject.Invoke(new()
        ///     {
        ///         Name = "Example Project",
        ///     });
        /// 
        ///     var exampleVariableGroup = AzureDevOps.GetVariableGroup.Invoke(new()
        ///     {
        ///         ProjectId = exampleProject.Apply(getProjectResult =&gt; getProjectResult.Id),
        ///         Name = "Example Variable Group",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["id"] = exampleVariableGroup.Apply(getVariableGroupResult =&gt; getVariableGroupResult.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// ## Relevant Links
        /// 
        /// - [Azure DevOps Service REST API 6.0 - Variable Groups](https://docs.microsoft.com/en-us/rest/api/azure/devops/distributedtask/variablegroups?view=azure-devops-rest-6.0)
        /// </summary>
        public static Output<GetVariableGroupResult> Invoke(GetVariableGroupInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetVariableGroupResult>("azuredevops:index/getVariableGroup:getVariableGroup", args ?? new GetVariableGroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVariableGroupArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Variable Group to retrieve.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The project ID.
        /// </summary>
        [Input("projectId", required: true)]
        public string ProjectId { get; set; } = null!;

        public GetVariableGroupArgs()
        {
        }
        public static new GetVariableGroupArgs Empty => new GetVariableGroupArgs();
    }

    public sealed class GetVariableGroupInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Variable Group to retrieve.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The project ID.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        public GetVariableGroupInvokeArgs()
        {
        }
        public static new GetVariableGroupInvokeArgs Empty => new GetVariableGroupInvokeArgs();
    }


    [OutputType]
    public sealed class GetVariableGroupResult
    {
        /// <summary>
        /// Boolean that indicate if this Variable Group is shared by all pipelines of this project.
        /// </summary>
        public readonly bool AllowAccess;
        /// <summary>
        /// The description of the Variable Group.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of `key_vault` blocks as documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetVariableGroupKeyVaultResult> KeyVaults;
        /// <summary>
        /// The name of the Azure key vault to link secrets from as variables.
        /// </summary>
        public readonly string Name;
        public readonly string ProjectId;
        /// <summary>
        /// One or more `variable` blocks as documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetVariableGroupVariableResult> Variables;

        [OutputConstructor]
        private GetVariableGroupResult(
            bool allowAccess,

            string description,

            string id,

            ImmutableArray<Outputs.GetVariableGroupKeyVaultResult> keyVaults,

            string name,

            string projectId,

            ImmutableArray<Outputs.GetVariableGroupVariableResult> variables)
        {
            AllowAccess = allowAccess;
            Description = description;
            Id = id;
            KeyVaults = keyVaults;
            Name = name;
            ProjectId = projectId;
            Variables = variables;
        }
    }
}
