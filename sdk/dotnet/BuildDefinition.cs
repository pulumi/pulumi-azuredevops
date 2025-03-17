// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureDevOps
{
    /// <summary>
    /// Manages a Build Definition within Azure DevOps.
    /// 
    /// ## Example Usage
    /// 
    /// ### Azure DevOps
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AzureDevOps = Pulumi.AzureDevOps;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new AzureDevOps.Project("example", new()
    ///     {
    ///         Name = "Example Project",
    ///         Visibility = "private",
    ///         VersionControl = "Git",
    ///         WorkItemTemplate = "Agile",
    ///     });
    /// 
    ///     var exampleGit = new AzureDevOps.Git("example", new()
    ///     {
    ///         ProjectId = example.Id,
    ///         Name = "Example Repository",
    ///         Initialization = new AzureDevOps.Inputs.GitInitializationArgs
    ///         {
    ///             InitType = "Clean",
    ///         },
    ///     });
    /// 
    ///     var exampleVariableGroup = new AzureDevOps.VariableGroup("example", new()
    ///     {
    ///         ProjectId = example.Id,
    ///         Name = "Example Pipeline Variables",
    ///         Description = "Managed by Pulumi",
    ///         AllowAccess = true,
    ///         Variables = new[]
    ///         {
    ///             new AzureDevOps.Inputs.VariableGroupVariableArgs
    ///             {
    ///                 Name = "FOO",
    ///                 Value = "BAR",
    ///             },
    ///         },
    ///     });
    /// 
    ///     var exampleBuildDefinition = new AzureDevOps.BuildDefinition("example", new()
    ///     {
    ///         ProjectId = example.Id,
    ///         Name = "Example Build Definition",
    ///         Path = "\\ExampleFolder",
    ///         CiTrigger = new AzureDevOps.Inputs.BuildDefinitionCiTriggerArgs
    ///         {
    ///             UseYaml = false,
    ///         },
    ///         Schedules = new[]
    ///         {
    ///             new AzureDevOps.Inputs.BuildDefinitionScheduleArgs
    ///             {
    ///                 BranchFilters = new[]
    ///                 {
    ///                     new AzureDevOps.Inputs.BuildDefinitionScheduleBranchFilterArgs
    ///                     {
    ///                         Includes = new[]
    ///                         {
    ///                             "master",
    ///                         },
    ///                         Excludes = new[]
    ///                         {
    ///                             "test",
    ///                             "regression",
    ///                         },
    ///                     },
    ///                 },
    ///                 DaysToBuilds = new[]
    ///                 {
    ///                     "Wed",
    ///                     "Sun",
    ///                 },
    ///                 ScheduleOnlyWithChanges = true,
    ///                 StartHours = 10,
    ///                 StartMinutes = 59,
    ///                 TimeZone = "(UTC) Coordinated Universal Time",
    ///             },
    ///         },
    ///         Repository = new AzureDevOps.Inputs.BuildDefinitionRepositoryArgs
    ///         {
    ///             RepoType = "TfsGit",
    ///             RepoId = exampleGit.Id,
    ///             BranchName = exampleGit.DefaultBranch,
    ///             YmlPath = "azure-pipelines.yml",
    ///         },
    ///         VariableGroups = new[]
    ///         {
    ///             exampleVariableGroup.Id,
    ///         },
    ///         Variables = new[]
    ///         {
    ///             new AzureDevOps.Inputs.BuildDefinitionVariableArgs
    ///             {
    ///                 Name = "PipelineVariable",
    ///                 Value = "Go Microsoft!",
    ///             },
    ///             new AzureDevOps.Inputs.BuildDefinitionVariableArgs
    ///             {
    ///                 Name = "PipelineSecret",
    ///                 SecretValue = "ZGV2cw",
    ///                 IsSecret = true,
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ### GitHub Enterprise
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AzureDevOps = Pulumi.AzureDevOps;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new AzureDevOps.Project("example", new()
    ///     {
    ///         Name = "Example Project",
    ///         Visibility = "private",
    ///         VersionControl = "Git",
    ///         WorkItemTemplate = "Agile",
    ///     });
    /// 
    ///     var exampleServiceEndpointGitHubEnterprise = new AzureDevOps.ServiceEndpointGitHubEnterprise("example", new()
    ///     {
    ///         ProjectId = example.Id,
    ///         ServiceEndpointName = "Example GitHub Enterprise",
    ///         Url = "https://github.contoso.com",
    ///         Description = "Managed by Pulumi",
    ///         AuthPersonal = new AzureDevOps.Inputs.ServiceEndpointGitHubEnterpriseAuthPersonalArgs
    ///         {
    ///             PersonalAccessToken = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
    ///         },
    ///     });
    /// 
    ///     var exampleBuildDefinition = new AzureDevOps.BuildDefinition("example", new()
    ///     {
    ///         ProjectId = example.Id,
    ///         Name = "Example Build Definition",
    ///         Path = "\\ExampleFolder",
    ///         CiTrigger = new AzureDevOps.Inputs.BuildDefinitionCiTriggerArgs
    ///         {
    ///             UseYaml = false,
    ///         },
    ///         Repository = new AzureDevOps.Inputs.BuildDefinitionRepositoryArgs
    ///         {
    ///             RepoType = "GitHubEnterprise",
    ///             RepoId = "&lt;GitHub Org&gt;/&lt;Repo Name&gt;",
    ///             GithubEnterpriseUrl = "https://github.company.com",
    ///             BranchName = "master",
    ///             YmlPath = "azure-pipelines.yml",
    ///             ServiceConnectionId = exampleServiceEndpointGitHubEnterprise.Id,
    ///         },
    ///         Schedules = new[]
    ///         {
    ///             new AzureDevOps.Inputs.BuildDefinitionScheduleArgs
    ///             {
    ///                 BranchFilters = new[]
    ///                 {
    ///                     new AzureDevOps.Inputs.BuildDefinitionScheduleBranchFilterArgs
    ///                     {
    ///                         Includes = new[]
    ///                         {
    ///                             "main",
    ///                         },
    ///                         Excludes = new[]
    ///                         {
    ///                             "test",
    ///                             "regression",
    ///                         },
    ///                     },
    ///                 },
    ///                 DaysToBuilds = new[]
    ///                 {
    ///                     "Wed",
    ///                     "Sun",
    ///                 },
    ///                 ScheduleOnlyWithChanges = true,
    ///                 StartHours = 10,
    ///                 StartMinutes = 59,
    ///                 TimeZone = "(UTC) Coordinated Universal Time",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ### Build Completion Trigger
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AzureDevOps = Pulumi.AzureDevOps;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new AzureDevOps.BuildDefinition("example", new()
    ///     {
    ///         ProjectId = exampleAzuredevopsProject.Id,
    ///         Name = "Example Build Definition",
    ///         Path = "\\ExampleFolder",
    ///         CiTrigger = new AzureDevOps.Inputs.BuildDefinitionCiTriggerArgs
    ///         {
    ///             UseYaml = false,
    ///         },
    ///         Repository = new AzureDevOps.Inputs.BuildDefinitionRepositoryArgs
    ///         {
    ///             RepoType = "GitHubEnterprise",
    ///             RepoId = "&lt;GitHub Org&gt;/&lt;Repo Name&gt;",
    ///             GithubEnterpriseUrl = "https://github.company.com",
    ///             BranchName = "main",
    ///             YmlPath = "azure-pipelines.yml",
    ///             ServiceConnectionId = exampleAzuredevopsServiceendpointGithubEnterprise.Id,
    ///         },
    ///         BuildCompletionTriggers = new[]
    ///         {
    ///             new AzureDevOps.Inputs.BuildDefinitionBuildCompletionTriggerArgs
    ///             {
    ///                 BuildDefinitionId = 10,
    ///                 BranchFilters = new[]
    ///                 {
    ///                     new AzureDevOps.Inputs.BuildDefinitionBuildCompletionTriggerBranchFilterArgs
    ///                     {
    ///                         Includes = new[]
    ///                         {
    ///                             "main",
    ///                         },
    ///                         Excludes = new[]
    ///                         {
    ///                             "test",
    ///                         },
    ///                     },
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ### Pull Request Trigger
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AzureDevOps = Pulumi.AzureDevOps;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = AzureDevOps.GetServiceEndpointGithub.Invoke(new()
    ///     {
    ///         ProjectId = exampleAzuredevopsProject.Id,
    ///         ServiceEndpointId = "00000000-0000-0000-0000-000000000000",
    ///     });
    /// 
    ///     var exampleBuildDefinition = new AzureDevOps.BuildDefinition("example", new()
    ///     {
    ///         ProjectId = exampleAzuredevopsProject2.Id,
    ///         Name = "Example Build Definition",
    ///         Path = "\\ExampleFolder",
    ///         CiTrigger = new AzureDevOps.Inputs.BuildDefinitionCiTriggerArgs
    ///         {
    ///             UseYaml = false,
    ///         },
    ///         Repository = new AzureDevOps.Inputs.BuildDefinitionRepositoryArgs
    ///         {
    ///             RepoType = "GitHub",
    ///             RepoId = "&lt;GitHub Org&gt;/&lt;Repo Name&gt;",
    ///             BranchName = "main",
    ///             YmlPath = "azure-pipelines.yml",
    ///             ServiceConnectionId = example.Apply(getServiceEndpointGithubResult =&gt; getServiceEndpointGithubResult.Id),
    ///         },
    ///         PullRequestTrigger = new AzureDevOps.Inputs.BuildDefinitionPullRequestTriggerArgs
    ///         {
    ///             Override = new AzureDevOps.Inputs.BuildDefinitionPullRequestTriggerOverrideArgs
    ///             {
    ///                 BranchFilters = new[]
    ///                 {
    ///                     new AzureDevOps.Inputs.BuildDefinitionPullRequestTriggerOverrideBranchFilterArgs
    ///                     {
    ///                         Includes = new[]
    ///                         {
    ///                             "main",
    ///                         },
    ///                     },
    ///                 },
    ///             },
    ///             Forks = new AzureDevOps.Inputs.BuildDefinitionPullRequestTriggerForksArgs
    ///             {
    ///                 Enabled = false,
    ///                 ShareSecrets = false,
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ### Using Other Git and Agent Jobs
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AzureDevOps = Pulumi.AzureDevOps;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new AzureDevOps.ServiceEndpointGenericGit("example", new()
    ///     {
    ///         ProjectId = exampleAzuredevopsProject.Id,
    ///         RepositoryUrl = "https://gitlab.com/example/example.git",
    ///         Password = "token",
    ///         ServiceEndpointName = "Example Generic Git",
    ///     });
    /// 
    ///     var exampleBuildDefinition = new AzureDevOps.BuildDefinition("example", new()
    ///     {
    ///         ProjectId = exampleAzuredevopsProject2.Id,
    ///         Name = "Example Build Definition",
    ///         Path = "\\ExampleFolder",
    ///         CiTrigger = new AzureDevOps.Inputs.BuildDefinitionCiTriggerArgs
    ///         {
    ///             UseYaml = false,
    ///         },
    ///         Repository = new AzureDevOps.Inputs.BuildDefinitionRepositoryArgs
    ///         {
    ///             RepoType = "Git",
    ///             RepoId = example.RepositoryUrl,
    ///             BranchName = "refs/heads/main",
    ///             Url = example.RepositoryUrl,
    ///             ServiceConnectionId = example.Id,
    ///         },
    ///         Jobs = new[]
    ///         {
    ///             new AzureDevOps.Inputs.BuildDefinitionJobArgs
    ///             {
    ///                 Name = "Agent Job1",
    ///                 RefName = "agent_job1",
    ///                 Condition = "succeededOrFailed()",
    ///                 Target = new AzureDevOps.Inputs.BuildDefinitionJobTargetArgs
    ///                 {
    ///                     Type = "AgentJob",
    ///                     ExecutionOptions = new AzureDevOps.Inputs.BuildDefinitionJobTargetExecutionOptionsArgs
    ///                     {
    ///                         Type = "None",
    ///                     },
    ///                 },
    ///             },
    ///             new AzureDevOps.Inputs.BuildDefinitionJobArgs
    ///             {
    ///                 Name = "Agent Job2",
    ///                 RefName = "agent_job2",
    ///                 Condition = "succeededOrFailed()",
    ///                 Dependencies = new[]
    ///                 {
    ///                     new AzureDevOps.Inputs.BuildDefinitionJobDependencyArgs
    ///                     {
    ///                         Scope = "agent_job1",
    ///                     },
    ///                 },
    ///                 Target = new AzureDevOps.Inputs.BuildDefinitionJobTargetArgs
    ///                 {
    ///                     Type = "AgentJob",
    ///                     Demands = new[]
    ///                     {
    ///                         "git",
    ///                     },
    ///                     ExecutionOptions = new AzureDevOps.Inputs.BuildDefinitionJobTargetExecutionOptionsArgs
    ///                     {
    ///                         Type = "Multi-Configuration",
    ///                         ContinueOnError = true,
    ///                         Multipliers = "multipliers",
    ///                         MaxConcurrency = 2,
    ///                     },
    ///                 },
    ///             },
    ///             new AzureDevOps.Inputs.BuildDefinitionJobArgs
    ///             {
    ///                 Name = "Agentless Job1",
    ///                 RefName = "agentless_job1",
    ///                 Condition = "succeeded()",
    ///                 Target = new AzureDevOps.Inputs.BuildDefinitionJobTargetArgs
    ///                 {
    ///                     Type = "AgentlessJob",
    ///                     ExecutionOptions = new AzureDevOps.Inputs.BuildDefinitionJobTargetExecutionOptionsArgs
    ///                     {
    ///                         Type = "None",
    ///                     },
    ///                 },
    ///             },
    ///             new AzureDevOps.Inputs.BuildDefinitionJobArgs
    ///             {
    ///                 Name = "Agentless Job2",
    ///                 RefName = "agentless_job2",
    ///                 Condition = "succeeded()",
    ///                 JobAuthorizationScope = "project",
    ///                 Dependencies = new[]
    ///                 {
    ///                     new AzureDevOps.Inputs.BuildDefinitionJobDependencyArgs
    ///                     {
    ///                         Scope = "agent_job2",
    ///                     },
    ///                     new AzureDevOps.Inputs.BuildDefinitionJobDependencyArgs
    ///                     {
    ///                         Scope = "agentless_job1",
    ///                     },
    ///                 },
    ///                 Target = new AzureDevOps.Inputs.BuildDefinitionJobTargetArgs
    ///                 {
    ///                     Type = "AgentlessJob",
    ///                     ExecutionOptions = new AzureDevOps.Inputs.BuildDefinitionJobTargetExecutionOptionsArgs
    ///                     {
    ///                         Type = "Multi-Configuration",
    ///                         ContinueOnError = true,
    ///                         Multipliers = "multipliers",
    ///                     },
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Remarks
    /// 
    /// The path attribute can not end in `\` unless the path is the root value of `\`.
    /// 
    /// Valid path values (yaml encoded) include:
    /// - `\\`
    /// - `\\ExampleFolder`
    /// - `\\Nested\\Example Folder`
    /// 
    /// The value of `\\ExampleFolder\\` would be invalid.
    /// 
    /// ## Relevant Links
    /// 
    /// - [Azure DevOps Service REST API 7.0 - Build Definitions](https://docs.microsoft.com/en-us/rest/api/azure/devops/build/definitions?view=azure-devops-rest-7.0)
    /// 
    /// ## Import
    /// 
    /// Azure DevOps Build Definitions can be imported using the project name/definitions Id or by the project Guid/definitions Id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import azuredevops:index/buildDefinition:BuildDefinition example "Example Project"/10
    /// ```
    /// 
    /// or
    /// 
    /// ```sh
    /// $ pulumi import azuredevops:index/buildDefinition:BuildDefinition example 00000000-0000-0000-0000-000000000000/0
    /// ```
    /// </summary>
    [AzureDevOpsResourceType("azuredevops:index/buildDefinition:BuildDefinition")]
    public partial class BuildDefinition : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The agent pool that should execute the build. Defaults to `Azure Pipelines`.
        /// </summary>
        [Output("agentPoolName")]
        public Output<string?> AgentPoolName { get; private set; } = null!;

        /// <summary>
        /// The Agent Specification to run the pipelines. Required when `repo_type` is `Git`. Example: `windows-2019`, `windows-latest`, `macos-13` etc.
        /// </summary>
        [Output("agentSpecification")]
        public Output<string?> AgentSpecification { get; private set; } = null!;

        /// <summary>
        /// A `build_completion_trigger` block as documented below.
        /// </summary>
        [Output("buildCompletionTriggers")]
        public Output<ImmutableArray<Outputs.BuildDefinitionBuildCompletionTrigger>> BuildCompletionTriggers { get; private set; } = null!;

        /// <summary>
        /// A `ci_trigger` block as documented below.
        /// </summary>
        [Output("ciTrigger")]
        public Output<Outputs.BuildDefinitionCiTrigger?> CiTrigger { get; private set; } = null!;

        /// <summary>
        /// A `features` blocks as documented below.
        /// </summary>
        [Output("features")]
        public Output<ImmutableArray<Outputs.BuildDefinitionFeature>> Features { get; private set; } = null!;

        /// <summary>
        /// The job authorization scope for builds queued against this definition. Possible values are: `project`, `projectCollection`. Defaults to `projectCollection`.
        /// </summary>
        [Output("jobAuthorizationScope")]
        public Output<string?> JobAuthorizationScope { get; private set; } = null!;

        /// <summary>
        /// A `jobs` blocks as documented below.
        /// 
        /// &gt; **NOTE:** The `jobs` are classic pipelines, you need to enable the classic pipeline feature for your organization to use this feature.
        /// </summary>
        [Output("jobs")]
        public Output<ImmutableArray<Outputs.BuildDefinitionJob>> Jobs { get; private set; } = null!;

        /// <summary>
        /// The name of the build definition.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The folder path of the build definition.
        /// </summary>
        [Output("path")]
        public Output<string?> Path { get; private set; } = null!;

        /// <summary>
        /// The project ID or project name.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// A `pull_request_trigger` block as documented below.
        /// </summary>
        [Output("pullRequestTrigger")]
        public Output<Outputs.BuildDefinitionPullRequestTrigger?> PullRequestTrigger { get; private set; } = null!;

        /// <summary>
        /// The queue status of the build definition. Possible values are: `enabled` or `paused` or `disabled`. Defaults to `enabled`.
        /// </summary>
        [Output("queueStatus")]
        public Output<string?> QueueStatus { get; private set; } = null!;

        /// <summary>
        /// A `repository` block as documented below.
        /// </summary>
        [Output("repository")]
        public Output<Outputs.BuildDefinitionRepository> Repository { get; private set; } = null!;

        /// <summary>
        /// The revision of the build definition
        /// </summary>
        [Output("revision")]
        public Output<int> Revision { get; private set; } = null!;

        [Output("schedules")]
        public Output<ImmutableArray<Outputs.BuildDefinitionSchedule>> Schedules { get; private set; } = null!;

        /// <summary>
        /// A list of variable group IDs (integers) to link to the build definition.
        /// </summary>
        [Output("variableGroups")]
        public Output<ImmutableArray<int>> VariableGroups { get; private set; } = null!;

        /// <summary>
        /// A list of `variable` blocks, as documented below.
        /// </summary>
        [Output("variables")]
        public Output<ImmutableArray<Outputs.BuildDefinitionVariable>> Variables { get; private set; } = null!;


        /// <summary>
        /// Create a BuildDefinition resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BuildDefinition(string name, BuildDefinitionArgs args, CustomResourceOptions? options = null)
            : base("azuredevops:index/buildDefinition:BuildDefinition", name, args ?? new BuildDefinitionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BuildDefinition(string name, Input<string> id, BuildDefinitionState? state = null, CustomResourceOptions? options = null)
            : base("azuredevops:index/buildDefinition:BuildDefinition", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing BuildDefinition resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BuildDefinition Get(string name, Input<string> id, BuildDefinitionState? state = null, CustomResourceOptions? options = null)
        {
            return new BuildDefinition(name, id, state, options);
        }
    }

    public sealed class BuildDefinitionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The agent pool that should execute the build. Defaults to `Azure Pipelines`.
        /// </summary>
        [Input("agentPoolName")]
        public Input<string>? AgentPoolName { get; set; }

        /// <summary>
        /// The Agent Specification to run the pipelines. Required when `repo_type` is `Git`. Example: `windows-2019`, `windows-latest`, `macos-13` etc.
        /// </summary>
        [Input("agentSpecification")]
        public Input<string>? AgentSpecification { get; set; }

        [Input("buildCompletionTriggers")]
        private InputList<Inputs.BuildDefinitionBuildCompletionTriggerArgs>? _buildCompletionTriggers;

        /// <summary>
        /// A `build_completion_trigger` block as documented below.
        /// </summary>
        public InputList<Inputs.BuildDefinitionBuildCompletionTriggerArgs> BuildCompletionTriggers
        {
            get => _buildCompletionTriggers ?? (_buildCompletionTriggers = new InputList<Inputs.BuildDefinitionBuildCompletionTriggerArgs>());
            set => _buildCompletionTriggers = value;
        }

        /// <summary>
        /// A `ci_trigger` block as documented below.
        /// </summary>
        [Input("ciTrigger")]
        public Input<Inputs.BuildDefinitionCiTriggerArgs>? CiTrigger { get; set; }

        [Input("features")]
        private InputList<Inputs.BuildDefinitionFeatureArgs>? _features;

        /// <summary>
        /// A `features` blocks as documented below.
        /// </summary>
        public InputList<Inputs.BuildDefinitionFeatureArgs> Features
        {
            get => _features ?? (_features = new InputList<Inputs.BuildDefinitionFeatureArgs>());
            set => _features = value;
        }

        /// <summary>
        /// The job authorization scope for builds queued against this definition. Possible values are: `project`, `projectCollection`. Defaults to `projectCollection`.
        /// </summary>
        [Input("jobAuthorizationScope")]
        public Input<string>? JobAuthorizationScope { get; set; }

        [Input("jobs")]
        private InputList<Inputs.BuildDefinitionJobArgs>? _jobs;

        /// <summary>
        /// A `jobs` blocks as documented below.
        /// 
        /// &gt; **NOTE:** The `jobs` are classic pipelines, you need to enable the classic pipeline feature for your organization to use this feature.
        /// </summary>
        public InputList<Inputs.BuildDefinitionJobArgs> Jobs
        {
            get => _jobs ?? (_jobs = new InputList<Inputs.BuildDefinitionJobArgs>());
            set => _jobs = value;
        }

        /// <summary>
        /// The name of the build definition.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The folder path of the build definition.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        /// <summary>
        /// The project ID or project name.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        /// <summary>
        /// A `pull_request_trigger` block as documented below.
        /// </summary>
        [Input("pullRequestTrigger")]
        public Input<Inputs.BuildDefinitionPullRequestTriggerArgs>? PullRequestTrigger { get; set; }

        /// <summary>
        /// The queue status of the build definition. Possible values are: `enabled` or `paused` or `disabled`. Defaults to `enabled`.
        /// </summary>
        [Input("queueStatus")]
        public Input<string>? QueueStatus { get; set; }

        /// <summary>
        /// A `repository` block as documented below.
        /// </summary>
        [Input("repository", required: true)]
        public Input<Inputs.BuildDefinitionRepositoryArgs> Repository { get; set; } = null!;

        [Input("schedules")]
        private InputList<Inputs.BuildDefinitionScheduleArgs>? _schedules;
        public InputList<Inputs.BuildDefinitionScheduleArgs> Schedules
        {
            get => _schedules ?? (_schedules = new InputList<Inputs.BuildDefinitionScheduleArgs>());
            set => _schedules = value;
        }

        [Input("variableGroups")]
        private InputList<int>? _variableGroups;

        /// <summary>
        /// A list of variable group IDs (integers) to link to the build definition.
        /// </summary>
        public InputList<int> VariableGroups
        {
            get => _variableGroups ?? (_variableGroups = new InputList<int>());
            set => _variableGroups = value;
        }

        [Input("variables")]
        private InputList<Inputs.BuildDefinitionVariableArgs>? _variables;

        /// <summary>
        /// A list of `variable` blocks, as documented below.
        /// </summary>
        public InputList<Inputs.BuildDefinitionVariableArgs> Variables
        {
            get => _variables ?? (_variables = new InputList<Inputs.BuildDefinitionVariableArgs>());
            set => _variables = value;
        }

        public BuildDefinitionArgs()
        {
        }
        public static new BuildDefinitionArgs Empty => new BuildDefinitionArgs();
    }

    public sealed class BuildDefinitionState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The agent pool that should execute the build. Defaults to `Azure Pipelines`.
        /// </summary>
        [Input("agentPoolName")]
        public Input<string>? AgentPoolName { get; set; }

        /// <summary>
        /// The Agent Specification to run the pipelines. Required when `repo_type` is `Git`. Example: `windows-2019`, `windows-latest`, `macos-13` etc.
        /// </summary>
        [Input("agentSpecification")]
        public Input<string>? AgentSpecification { get; set; }

        [Input("buildCompletionTriggers")]
        private InputList<Inputs.BuildDefinitionBuildCompletionTriggerGetArgs>? _buildCompletionTriggers;

        /// <summary>
        /// A `build_completion_trigger` block as documented below.
        /// </summary>
        public InputList<Inputs.BuildDefinitionBuildCompletionTriggerGetArgs> BuildCompletionTriggers
        {
            get => _buildCompletionTriggers ?? (_buildCompletionTriggers = new InputList<Inputs.BuildDefinitionBuildCompletionTriggerGetArgs>());
            set => _buildCompletionTriggers = value;
        }

        /// <summary>
        /// A `ci_trigger` block as documented below.
        /// </summary>
        [Input("ciTrigger")]
        public Input<Inputs.BuildDefinitionCiTriggerGetArgs>? CiTrigger { get; set; }

        [Input("features")]
        private InputList<Inputs.BuildDefinitionFeatureGetArgs>? _features;

        /// <summary>
        /// A `features` blocks as documented below.
        /// </summary>
        public InputList<Inputs.BuildDefinitionFeatureGetArgs> Features
        {
            get => _features ?? (_features = new InputList<Inputs.BuildDefinitionFeatureGetArgs>());
            set => _features = value;
        }

        /// <summary>
        /// The job authorization scope for builds queued against this definition. Possible values are: `project`, `projectCollection`. Defaults to `projectCollection`.
        /// </summary>
        [Input("jobAuthorizationScope")]
        public Input<string>? JobAuthorizationScope { get; set; }

        [Input("jobs")]
        private InputList<Inputs.BuildDefinitionJobGetArgs>? _jobs;

        /// <summary>
        /// A `jobs` blocks as documented below.
        /// 
        /// &gt; **NOTE:** The `jobs` are classic pipelines, you need to enable the classic pipeline feature for your organization to use this feature.
        /// </summary>
        public InputList<Inputs.BuildDefinitionJobGetArgs> Jobs
        {
            get => _jobs ?? (_jobs = new InputList<Inputs.BuildDefinitionJobGetArgs>());
            set => _jobs = value;
        }

        /// <summary>
        /// The name of the build definition.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The folder path of the build definition.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        /// <summary>
        /// The project ID or project name.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// A `pull_request_trigger` block as documented below.
        /// </summary>
        [Input("pullRequestTrigger")]
        public Input<Inputs.BuildDefinitionPullRequestTriggerGetArgs>? PullRequestTrigger { get; set; }

        /// <summary>
        /// The queue status of the build definition. Possible values are: `enabled` or `paused` or `disabled`. Defaults to `enabled`.
        /// </summary>
        [Input("queueStatus")]
        public Input<string>? QueueStatus { get; set; }

        /// <summary>
        /// A `repository` block as documented below.
        /// </summary>
        [Input("repository")]
        public Input<Inputs.BuildDefinitionRepositoryGetArgs>? Repository { get; set; }

        /// <summary>
        /// The revision of the build definition
        /// </summary>
        [Input("revision")]
        public Input<int>? Revision { get; set; }

        [Input("schedules")]
        private InputList<Inputs.BuildDefinitionScheduleGetArgs>? _schedules;
        public InputList<Inputs.BuildDefinitionScheduleGetArgs> Schedules
        {
            get => _schedules ?? (_schedules = new InputList<Inputs.BuildDefinitionScheduleGetArgs>());
            set => _schedules = value;
        }

        [Input("variableGroups")]
        private InputList<int>? _variableGroups;

        /// <summary>
        /// A list of variable group IDs (integers) to link to the build definition.
        /// </summary>
        public InputList<int> VariableGroups
        {
            get => _variableGroups ?? (_variableGroups = new InputList<int>());
            set => _variableGroups = value;
        }

        [Input("variables")]
        private InputList<Inputs.BuildDefinitionVariableGetArgs>? _variables;

        /// <summary>
        /// A list of `variable` blocks, as documented below.
        /// </summary>
        public InputList<Inputs.BuildDefinitionVariableGetArgs> Variables
        {
            get => _variables ?? (_variables = new InputList<Inputs.BuildDefinitionVariableGetArgs>());
            set => _variables = value;
        }

        public BuildDefinitionState()
        {
        }
        public static new BuildDefinitionState Empty => new BuildDefinitionState();
    }
}
