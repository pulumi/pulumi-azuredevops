// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureDevOps
{
    /// <summary>
    /// Manage pipeline access permissions to resources.
    /// 
    /// &gt; **Note** This resource is a replacement for `azuredevops.ResourceAuthorization`.  Pipeline authorizations managed by `azuredevops.ResourceAuthorization` can also be managed by this resource.
    /// 
    /// &gt; **Note** If both "All Pipeline Authorization" and "Custom Pipeline Authorization" are configured, "All Pipeline Authorization" has higher priority.
    /// 
    /// ## Example Usage
    /// 
    /// ### Authorization for all pipelines
    /// 
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
    ///         Description = "Managed by Pulumi",
    ///     });
    /// 
    ///     var examplePool = new AzureDevOps.Pool("example", new()
    ///     {
    ///         Name = "Example Pool",
    ///         AutoProvision = false,
    ///         AutoUpdate = false,
    ///     });
    /// 
    ///     var exampleQueue = new AzureDevOps.Queue("example", new()
    ///     {
    ///         ProjectId = example.Id,
    ///         AgentPoolId = examplePool.Id,
    ///     });
    /// 
    ///     var examplePipelineAuthorization = new AzureDevOps.PipelineAuthorization("example", new()
    ///     {
    ///         ProjectId = example.Id,
    ///         ResourceId = exampleQueue.Id,
    ///         Type = "queue",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ### Authorization for specific pipeline
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AzureDevOps = Pulumi.AzureDevOps;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var exampleProject = new AzureDevOps.Project("example", new()
    ///     {
    ///         Name = "Example Project",
    ///         Visibility = "private",
    ///         VersionControl = "Git",
    ///         WorkItemTemplate = "Agile",
    ///         Description = "Managed by Pulumi",
    ///     });
    /// 
    ///     var examplePool = new AzureDevOps.Pool("example", new()
    ///     {
    ///         Name = "Example Pool",
    ///         AutoProvision = false,
    ///         AutoUpdate = false,
    ///     });
    /// 
    ///     var exampleQueue = new AzureDevOps.Queue("example", new()
    ///     {
    ///         ProjectId = exampleProject.Id,
    ///         AgentPoolId = examplePool.Id,
    ///     });
    /// 
    ///     var example = AzureDevOps.GetGitRepository.Invoke(new()
    ///     {
    ///         ProjectId = exampleProject.Id,
    ///         Name = "Example Project",
    ///     });
    /// 
    ///     var exampleBuildDefinition = new AzureDevOps.BuildDefinition("example", new()
    ///     {
    ///         ProjectId = exampleProject.Id,
    ///         Name = "Example Pipeline",
    ///         Repository = new AzureDevOps.Inputs.BuildDefinitionRepositoryArgs
    ///         {
    ///             RepoType = "TfsGit",
    ///             RepoId = example.Apply(getGitRepositoryResult =&gt; getGitRepositoryResult.Id),
    ///             YmlPath = "azure-pipelines.yml",
    ///         },
    ///     });
    /// 
    ///     var examplePipelineAuthorization = new AzureDevOps.PipelineAuthorization("example", new()
    ///     {
    ///         ProjectId = exampleProject.Id,
    ///         ResourceId = exampleQueue.Id,
    ///         Type = "queue",
    ///         PipelineId = exampleBuildDefinition.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Relevant Links
    /// 
    /// - [Azure DevOps Service REST API 7.1 - Pipeline Permissions](https://learn.microsoft.com/en-us/rest/api/azure/devops/approvalsandchecks/pipeline-permissions?view=azure-devops-rest-7.1)
    /// </summary>
    [AzureDevOpsResourceType("azuredevops:index/pipelineAuthorization:PipelineAuthorization")]
    public partial class PipelineAuthorization : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the pipeline. If not configured, all pipelines will be authorized. Changing this forces a new resource to be created.
        /// </summary>
        [Output("pipelineId")]
        public Output<int?> PipelineId { get; private set; } = null!;

        /// <summary>
        /// The ID of the project where the pipeline exists. Defaults to `project_id` if not specified. Changing this forces a new resource to be created
        /// </summary>
        [Output("pipelineProjectId")]
        public Output<string?> PipelineProjectId { get; private set; } = null!;

        /// <summary>
        /// The  ID of the project. Changing this forces a new resource to be created
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// The ID of the resource to authorize. Changing this forces a new resource to be created
        /// </summary>
        [Output("resourceId")]
        public Output<string> ResourceId { get; private set; } = null!;

        /// <summary>
        /// The type of the resource to authorize. Possible values are: `endpoint`, `queue`, `variablegroup`, `environment`, `repository`. Changing this forces a new resource to be created
        /// 
        /// &gt; **Note** `repository` is for AzureDevOps repository. To authorize repository other than
        /// Azure DevOps like GitHub you need to use service connection(`endpoint`)  to connect and authorize.
        /// Typical process for connecting to GitHub:
        /// **Pipeline  &lt;----&gt; Service Connection(`endpoint`) &lt;----&gt; GitHub Repository**
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a PipelineAuthorization resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PipelineAuthorization(string name, PipelineAuthorizationArgs args, CustomResourceOptions? options = null)
            : base("azuredevops:index/pipelineAuthorization:PipelineAuthorization", name, args ?? new PipelineAuthorizationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PipelineAuthorization(string name, Input<string> id, PipelineAuthorizationState? state = null, CustomResourceOptions? options = null)
            : base("azuredevops:index/pipelineAuthorization:PipelineAuthorization", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing PipelineAuthorization resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PipelineAuthorization Get(string name, Input<string> id, PipelineAuthorizationState? state = null, CustomResourceOptions? options = null)
        {
            return new PipelineAuthorization(name, id, state, options);
        }
    }

    public sealed class PipelineAuthorizationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the pipeline. If not configured, all pipelines will be authorized. Changing this forces a new resource to be created.
        /// </summary>
        [Input("pipelineId")]
        public Input<int>? PipelineId { get; set; }

        /// <summary>
        /// The ID of the project where the pipeline exists. Defaults to `project_id` if not specified. Changing this forces a new resource to be created
        /// </summary>
        [Input("pipelineProjectId")]
        public Input<string>? PipelineProjectId { get; set; }

        /// <summary>
        /// The  ID of the project. Changing this forces a new resource to be created
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        /// <summary>
        /// The ID of the resource to authorize. Changing this forces a new resource to be created
        /// </summary>
        [Input("resourceId", required: true)]
        public Input<string> ResourceId { get; set; } = null!;

        /// <summary>
        /// The type of the resource to authorize. Possible values are: `endpoint`, `queue`, `variablegroup`, `environment`, `repository`. Changing this forces a new resource to be created
        /// 
        /// &gt; **Note** `repository` is for AzureDevOps repository. To authorize repository other than
        /// Azure DevOps like GitHub you need to use service connection(`endpoint`)  to connect and authorize.
        /// Typical process for connecting to GitHub:
        /// **Pipeline  &lt;----&gt; Service Connection(`endpoint`) &lt;----&gt; GitHub Repository**
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public PipelineAuthorizationArgs()
        {
        }
        public static new PipelineAuthorizationArgs Empty => new PipelineAuthorizationArgs();
    }

    public sealed class PipelineAuthorizationState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the pipeline. If not configured, all pipelines will be authorized. Changing this forces a new resource to be created.
        /// </summary>
        [Input("pipelineId")]
        public Input<int>? PipelineId { get; set; }

        /// <summary>
        /// The ID of the project where the pipeline exists. Defaults to `project_id` if not specified. Changing this forces a new resource to be created
        /// </summary>
        [Input("pipelineProjectId")]
        public Input<string>? PipelineProjectId { get; set; }

        /// <summary>
        /// The  ID of the project. Changing this forces a new resource to be created
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The ID of the resource to authorize. Changing this forces a new resource to be created
        /// </summary>
        [Input("resourceId")]
        public Input<string>? ResourceId { get; set; }

        /// <summary>
        /// The type of the resource to authorize. Possible values are: `endpoint`, `queue`, `variablegroup`, `environment`, `repository`. Changing this forces a new resource to be created
        /// 
        /// &gt; **Note** `repository` is for AzureDevOps repository. To authorize repository other than
        /// Azure DevOps like GitHub you need to use service connection(`endpoint`)  to connect and authorize.
        /// Typical process for connecting to GitHub:
        /// **Pipeline  &lt;----&gt; Service Connection(`endpoint`) &lt;----&gt; GitHub Repository**
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public PipelineAuthorizationState()
        {
        }
        public static new PipelineAuthorizationState Empty => new PipelineAuthorizationState();
    }
}
