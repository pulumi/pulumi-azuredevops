// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureDevOps
{
    public static class GetBuildDefinition
    {
        /// <summary>
        /// Use this data source to access information about an existing Build Definition.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using AzureDevOps = Pulumi.AzureDevOps;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var exampleProject = Output.Create(AzureDevOps.GetProject.InvokeAsync(new AzureDevOps.GetProjectArgs
        ///         {
        ///             Name = "Example Project",
        ///         }));
        ///         var exampleBuildDefinition = exampleProject.Apply(exampleProject =&gt; Output.Create(AzureDevOps.GetBuildDefinition.InvokeAsync(new AzureDevOps.GetBuildDefinitionArgs
        ///         {
        ///             ProjectId = exampleProject.Id,
        ///             Name = "existing",
        ///         })));
        ///         this.Id = exampleBuildDefinition.Apply(exampleBuildDefinition =&gt; exampleBuildDefinition.Id);
        ///     }
        /// 
        ///     [Output("id")]
        ///     public Output&lt;string&gt; Id { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetBuildDefinitionResult> InvokeAsync(GetBuildDefinitionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetBuildDefinitionResult>("azuredevops:index/getBuildDefinition:getBuildDefinition", args ?? new GetBuildDefinitionArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to access information about an existing Build Definition.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using AzureDevOps = Pulumi.AzureDevOps;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var exampleProject = Output.Create(AzureDevOps.GetProject.InvokeAsync(new AzureDevOps.GetProjectArgs
        ///         {
        ///             Name = "Example Project",
        ///         }));
        ///         var exampleBuildDefinition = exampleProject.Apply(exampleProject =&gt; Output.Create(AzureDevOps.GetBuildDefinition.InvokeAsync(new AzureDevOps.GetBuildDefinitionArgs
        ///         {
        ///             ProjectId = exampleProject.Id,
        ///             Name = "existing",
        ///         })));
        ///         this.Id = exampleBuildDefinition.Apply(exampleBuildDefinition =&gt; exampleBuildDefinition.Id);
        ///     }
        /// 
        ///     [Output("id")]
        ///     public Output&lt;string&gt; Id { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetBuildDefinitionResult> Invoke(GetBuildDefinitionInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetBuildDefinitionResult>("azuredevops:index/getBuildDefinition:getBuildDefinition", args ?? new GetBuildDefinitionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetBuildDefinitionArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of this Build Definition.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The path of the build definition. Default to `\`.
        /// </summary>
        [Input("path")]
        public string? Path { get; set; }

        /// <summary>
        /// The ID of the project.
        /// </summary>
        [Input("projectId", required: true)]
        public string ProjectId { get; set; } = null!;

        public GetBuildDefinitionArgs()
        {
        }
    }

    public sealed class GetBuildDefinitionInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of this Build Definition.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The path of the build definition. Default to `\`.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        /// <summary>
        /// The ID of the project.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        public GetBuildDefinitionInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetBuildDefinitionResult
    {
        /// <summary>
        /// The agent pool that should execute the build.
        /// </summary>
        public readonly string AgentPoolName;
        /// <summary>
        /// A `ci_trigger` block as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetBuildDefinitionCiTriggerResult> CiTriggers;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the variable.
        /// </summary>
        public readonly string Name;
        public readonly string? Path;
        public readonly string ProjectId;
        /// <summary>
        /// A `pull_request_trigger` block as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetBuildDefinitionPullRequestTriggerResult> PullRequestTriggers;
        /// <summary>
        /// A `repository` block as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetBuildDefinitionRepositoryResult> Repositories;
        /// <summary>
        /// The revision of the build definition.
        /// </summary>
        public readonly int Revision;
        /// <summary>
        /// A `schedules` block as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetBuildDefinitionScheduleResult> Schedules;
        /// <summary>
        /// A list of variable group IDs.
        /// </summary>
        public readonly ImmutableArray<int> VariableGroups;
        /// <summary>
        /// A `variable` block as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetBuildDefinitionVariableResult> Variables;

        [OutputConstructor]
        private GetBuildDefinitionResult(
            string agentPoolName,

            ImmutableArray<Outputs.GetBuildDefinitionCiTriggerResult> ciTriggers,

            string id,

            string name,

            string? path,

            string projectId,

            ImmutableArray<Outputs.GetBuildDefinitionPullRequestTriggerResult> pullRequestTriggers,

            ImmutableArray<Outputs.GetBuildDefinitionRepositoryResult> repositories,

            int revision,

            ImmutableArray<Outputs.GetBuildDefinitionScheduleResult> schedules,

            ImmutableArray<int> variableGroups,

            ImmutableArray<Outputs.GetBuildDefinitionVariableResult> variables)
        {
            AgentPoolName = agentPoolName;
            CiTriggers = ciTriggers;
            Id = id;
            Name = name;
            Path = path;
            ProjectId = projectId;
            PullRequestTriggers = pullRequestTriggers;
            Repositories = repositories;
            Revision = revision;
            Schedules = schedules;
            VariableGroups = variableGroups;
            Variables = variables;
        }
    }
}