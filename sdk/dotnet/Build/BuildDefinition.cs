// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureDevOps.Build
{
    /// <summary>
    /// Manages a Build Definition within Azure DevOps.
    /// 
    /// ## Example Usage
    /// ### Tfs
    /// ```csharp
    /// using Pulumi;
    /// using AzureDevOps = Pulumi.AzureDevOps;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var project = new AzureDevOps.Project("project", new AzureDevOps.ProjectArgs
    ///         {
    ///             Visibility = "private",
    ///             VersionControl = "Git",
    ///             WorkItemTemplate = "Agile",
    ///         });
    ///         var repository = new AzureDevOps.Git("repository", new AzureDevOps.GitArgs
    ///         {
    ///             ProjectId = project.Id,
    ///             Initialization = new AzureDevOps.Inputs.GitInitializationArgs
    ///             {
    ///                 InitType = "Clean",
    ///             },
    ///         });
    ///         var vars = new AzureDevOps.VariableGroup("vars", new AzureDevOps.VariableGroupArgs
    ///         {
    ///             ProjectId = project.Id,
    ///             Description = "Managed by Terraform",
    ///             AllowAccess = true,
    ///             Variables = 
    ///             {
    ///                 new AzureDevOps.Inputs.VariableGroupVariableArgs
    ///                 {
    ///                     Name = "FOO",
    ///                     Value = "BAR",
    ///                 },
    ///             },
    ///         });
    ///         var build = new AzureDevOps.BuildDefinition("build", new AzureDevOps.BuildDefinitionArgs
    ///         {
    ///             ProjectId = project.Id,
    ///             Path = "\\ExampleFolder",
    ///             CiTrigger = new AzureDevOps.Inputs.BuildDefinitionCiTriggerArgs
    ///             {
    ///                 UseYaml = true,
    ///             },
    ///             Schedules = 
    ///             {
    ///                 new AzureDevOps.Inputs.BuildDefinitionScheduleArgs
    ///                 {
    ///                     BranchFilters = 
    ///                     {
    ///                         new AzureDevOps.Inputs.BuildDefinitionScheduleBranchFilterArgs
    ///                         {
    ///                             Includes = 
    ///                             {
    ///                                 "master",
    ///                             },
    ///                             Excludes = 
    ///                             {
    ///                                 "test",
    ///                                 "regression",
    ///                             },
    ///                         },
    ///                     },
    ///                     DaysToBuilds = 
    ///                     {
    ///                         "Wed",
    ///                         "Sun",
    ///                     },
    ///                     ScheduleOnlyWithChanges = true,
    ///                     StartHours = 10,
    ///                     StartMinutes = 59,
    ///                     TimeZone = "(UTC) Coordinated Universal Time",
    ///                 },
    ///             },
    ///             Repository = new AzureDevOps.Inputs.BuildDefinitionRepositoryArgs
    ///             {
    ///                 RepoType = "TfsGit",
    ///                 RepoId = repository.Id,
    ///                 BranchName = repository.DefaultBranch,
    ///                 YmlPath = "azure-pipelines.yml",
    ///             },
    ///             VariableGroups = 
    ///             {
    ///                 vars.Id,
    ///             },
    ///             Variables = 
    ///             {
    ///                 new AzureDevOps.Inputs.BuildDefinitionVariableArgs
    ///                 {
    ///                     Name = "PipelineVariable",
    ///                     Value = "Go Microsoft!",
    ///                 },
    ///                 new AzureDevOps.Inputs.BuildDefinitionVariableArgs
    ///                 {
    ///                     Name = "PipelineSecret",
    ///                     SecretValue = "ZGV2cw",
    ///                     IsSecret = true,
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### GitHub Enterprise
    /// ```csharp
    /// using Pulumi;
    /// using AzureDevOps = Pulumi.AzureDevOps;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var sampleDotnetcoreAppRelease = new AzureDevOps.BuildDefinition("sampleDotnetcoreAppRelease", new AzureDevOps.BuildDefinitionArgs
    ///         {
    ///             ProjectId = azuredevops_project.Project.Id,
    ///             Path = "\\ExampleFolder",
    ///             CiTrigger = new AzureDevOps.Inputs.BuildDefinitionCiTriggerArgs
    ///             {
    ///                 UseYaml = true,
    ///             },
    ///             Repository = new AzureDevOps.Inputs.BuildDefinitionRepositoryArgs
    ///             {
    ///                 RepoType = "GitHubEnterprise",
    ///                 RepoId = "&lt;GitHub Org&gt;/&lt;Repo Name&gt;",
    ///                 GithubEnterpriseUrl = "https://github.company.com",
    ///                 BranchName = "master",
    ///                 YmlPath = "azure-pipelines.yml",
    ///                 ServiceConnectionId = "...",
    ///             },
    ///             Schedules = 
    ///             {
    ///                 new AzureDevOps.Inputs.BuildDefinitionScheduleArgs
    ///                 {
    ///                     BranchFilters = 
    ///                     {
    ///                         new AzureDevOps.Inputs.BuildDefinitionScheduleBranchFilterArgs
    ///                         {
    ///                             Includes = 
    ///                             {
    ///                                 "main",
    ///                             },
    ///                             Excludes = 
    ///                             {
    ///                                 "test",
    ///                                 "regression",
    ///                             },
    ///                         },
    ///                     },
    ///                     DaysToBuilds = 
    ///                     {
    ///                         "Wed",
    ///                         "Sun",
    ///                     },
    ///                     ScheduleOnlyWithChanges = true,
    ///                     StartHours = 10,
    ///                     StartMinutes = 59,
    ///                     TimeZone = "(UTC) Coordinated Universal Time",
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ## Relevant Links
    /// 
    /// - [Azure DevOps Service REST API 5.1 - Build Definitions](https://docs.microsoft.com/en-us/rest/api/azure/devops/build/definitions?view=azure-devops-rest-5.1)
    /// 
    /// ## Import
    /// 
    /// Azure DevOps Build Definitions can be imported using the project name/definitions Id or by the project Guid/definitions Id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azuredevops:Build/buildDefinition:BuildDefinition build "Test Project"/10
    /// ```
    /// 
    ///  or
    /// 
    /// ```sh
    ///  $ pulumi import azuredevops:Build/buildDefinition:BuildDefinition build 00000000-0000-0000-0000-000000000000/0
    /// ```
    /// </summary>
    [Obsolete(@"azuredevops.build.BuildDefinition has been deprecated in favor of azuredevops.BuildDefinition")]
    [AzureDevOpsResourceType("azuredevops:Build/buildDefinition:BuildDefinition")]
    public partial class BuildDefinition : Pulumi.CustomResource
    {
        /// <summary>
        /// The agent pool that should execute the build. Defaults to `Azure Pipelines`.
        /// </summary>
        [Output("agentPoolName")]
        public Output<string?> AgentPoolName { get; private set; } = null!;

        /// <summary>
        /// Continuous Integration trigger.
        /// </summary>
        [Output("ciTrigger")]
        public Output<Outputs.BuildDefinitionCiTrigger?> CiTrigger { get; private set; } = null!;

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
        /// Pull Request Integration Integration trigger.
        /// </summary>
        [Output("pullRequestTrigger")]
        public Output<Outputs.BuildDefinitionPullRequestTrigger?> PullRequestTrigger { get; private set; } = null!;

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
            : base("azuredevops:Build/buildDefinition:BuildDefinition", name, args ?? new BuildDefinitionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BuildDefinition(string name, Input<string> id, BuildDefinitionState? state = null, CustomResourceOptions? options = null)
            : base("azuredevops:Build/buildDefinition:BuildDefinition", name, state, MakeResourceOptions(options, id))
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

    public sealed class BuildDefinitionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The agent pool that should execute the build. Defaults to `Azure Pipelines`.
        /// </summary>
        [Input("agentPoolName")]
        public Input<string>? AgentPoolName { get; set; }

        /// <summary>
        /// Continuous Integration trigger.
        /// </summary>
        [Input("ciTrigger")]
        public Input<Inputs.BuildDefinitionCiTriggerArgs>? CiTrigger { get; set; }

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
        /// Pull Request Integration Integration trigger.
        /// </summary>
        [Input("pullRequestTrigger")]
        public Input<Inputs.BuildDefinitionPullRequestTriggerArgs>? PullRequestTrigger { get; set; }

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
    }

    public sealed class BuildDefinitionState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The agent pool that should execute the build. Defaults to `Azure Pipelines`.
        /// </summary>
        [Input("agentPoolName")]
        public Input<string>? AgentPoolName { get; set; }

        /// <summary>
        /// Continuous Integration trigger.
        /// </summary>
        [Input("ciTrigger")]
        public Input<Inputs.BuildDefinitionCiTriggerGetArgs>? CiTrigger { get; set; }

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
        /// Pull Request Integration Integration trigger.
        /// </summary>
        [Input("pullRequestTrigger")]
        public Input<Inputs.BuildDefinitionPullRequestTriggerGetArgs>? PullRequestTrigger { get; set; }

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
    }
}
