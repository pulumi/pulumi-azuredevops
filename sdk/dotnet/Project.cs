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
    /// Manages a project within Azure DevOps.
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
    ///     var example = new AzureDevOps.Project("example", new()
    ///     {
    ///         Description = "Managed by Terraform",
    ///         Features = 
    ///         {
    ///             { "artifacts", "disabled" },
    ///             { "testplans", "disabled" },
    ///         },
    ///         VersionControl = "Git",
    ///         Visibility = "private",
    ///         WorkItemTemplate = "Agile",
    ///     });
    /// 
    /// });
    /// ```
    /// ## Relevant Links
    /// 
    /// - [Azure DevOps Service REST API 7.0 - Projects](https://docs.microsoft.com/en-us/rest/api/azure/devops/core/projects?view=azure-devops-rest-7.0)
    /// 
    /// ## PAT Permissions Required
    /// 
    /// - **Project &amp; Team**: Read, Write, &amp; Manage
    /// 
    /// ## Import
    /// 
    /// Azure DevOps Projects can be imported using the project name or by the project Guid, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azuredevops:index/project:Project example "Example Project"
    /// ```
    /// 
    ///  or
    /// 
    /// ```sh
    ///  $ pulumi import azuredevops:index/project:Project example 00000000-0000-0000-0000-000000000000
    /// ```
    /// </summary>
    [AzureDevOpsResourceType("azuredevops:index/project:Project")]
    public partial class Project : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Description of the Project.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Defines the status (`enabled`, `disabled`) of the project features.
        /// Valid features are `boards`, `repositories`, `pipelines`, `testplans`, `artifacts`
        /// 
        /// &gt; **NOTE:**
        /// &gt; It's possible to define project features both within the `azuredevops.ProjectFeatures` resource and
        /// &gt; via the `features` block by using the `azuredevops.Project` resource.
        /// &gt; However it's not possible to use both methods to manage features, since there'll be conflicts.
        /// </summary>
        [Output("features")]
        public Output<ImmutableDictionary<string, string>?> Features { get; private set; } = null!;

        /// <summary>
        /// The Project Name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The Process Template ID used by the Project.
        /// </summary>
        [Output("processTemplateId")]
        public Output<string> ProcessTemplateId { get; private set; } = null!;

        /// <summary>
        /// Specifies the version control system. Valid values: `Git` or `Tfvc`. Defaults to `Git`.
        /// </summary>
        [Output("versionControl")]
        public Output<string?> VersionControl { get; private set; } = null!;

        /// <summary>
        /// Specifies the visibility of the Project. Valid values: `private` or `public`. Defaults to `private`.
        /// </summary>
        [Output("visibility")]
        public Output<string?> Visibility { get; private set; } = null!;

        /// <summary>
        /// Specifies the work item template. Valid values: `Agile`, `Basic`, `CMMI`, `Scrum` or a custom, pre-existing one. Defaults to `Agile`. An empty string will use the parent organization default.
        /// </summary>
        [Output("workItemTemplate")]
        public Output<string?> WorkItemTemplate { get; private set; } = null!;


        /// <summary>
        /// Create a Project resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Project(string name, ProjectArgs? args = null, CustomResourceOptions? options = null)
            : base("azuredevops:index/project:Project", name, args ?? new ProjectArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Project(string name, Input<string> id, ProjectState? state = null, CustomResourceOptions? options = null)
            : base("azuredevops:index/project:Project", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azuredevops:Core/project:Project"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Project resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Project Get(string name, Input<string> id, ProjectState? state = null, CustomResourceOptions? options = null)
        {
            return new Project(name, id, state, options);
        }
    }

    public sealed class ProjectArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Description of the Project.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("features")]
        private InputMap<string>? _features;

        /// <summary>
        /// Defines the status (`enabled`, `disabled`) of the project features.
        /// Valid features are `boards`, `repositories`, `pipelines`, `testplans`, `artifacts`
        /// 
        /// &gt; **NOTE:**
        /// &gt; It's possible to define project features both within the `azuredevops.ProjectFeatures` resource and
        /// &gt; via the `features` block by using the `azuredevops.Project` resource.
        /// &gt; However it's not possible to use both methods to manage features, since there'll be conflicts.
        /// </summary>
        public InputMap<string> Features
        {
            get => _features ?? (_features = new InputMap<string>());
            set => _features = value;
        }

        /// <summary>
        /// The Project Name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the version control system. Valid values: `Git` or `Tfvc`. Defaults to `Git`.
        /// </summary>
        [Input("versionControl")]
        public Input<string>? VersionControl { get; set; }

        /// <summary>
        /// Specifies the visibility of the Project. Valid values: `private` or `public`. Defaults to `private`.
        /// </summary>
        [Input("visibility")]
        public Input<string>? Visibility { get; set; }

        /// <summary>
        /// Specifies the work item template. Valid values: `Agile`, `Basic`, `CMMI`, `Scrum` or a custom, pre-existing one. Defaults to `Agile`. An empty string will use the parent organization default.
        /// </summary>
        [Input("workItemTemplate")]
        public Input<string>? WorkItemTemplate { get; set; }

        public ProjectArgs()
        {
        }
        public static new ProjectArgs Empty => new ProjectArgs();
    }

    public sealed class ProjectState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Description of the Project.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("features")]
        private InputMap<string>? _features;

        /// <summary>
        /// Defines the status (`enabled`, `disabled`) of the project features.
        /// Valid features are `boards`, `repositories`, `pipelines`, `testplans`, `artifacts`
        /// 
        /// &gt; **NOTE:**
        /// &gt; It's possible to define project features both within the `azuredevops.ProjectFeatures` resource and
        /// &gt; via the `features` block by using the `azuredevops.Project` resource.
        /// &gt; However it's not possible to use both methods to manage features, since there'll be conflicts.
        /// </summary>
        public InputMap<string> Features
        {
            get => _features ?? (_features = new InputMap<string>());
            set => _features = value;
        }

        /// <summary>
        /// The Project Name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The Process Template ID used by the Project.
        /// </summary>
        [Input("processTemplateId")]
        public Input<string>? ProcessTemplateId { get; set; }

        /// <summary>
        /// Specifies the version control system. Valid values: `Git` or `Tfvc`. Defaults to `Git`.
        /// </summary>
        [Input("versionControl")]
        public Input<string>? VersionControl { get; set; }

        /// <summary>
        /// Specifies the visibility of the Project. Valid values: `private` or `public`. Defaults to `private`.
        /// </summary>
        [Input("visibility")]
        public Input<string>? Visibility { get; set; }

        /// <summary>
        /// Specifies the work item template. Valid values: `Agile`, `Basic`, `CMMI`, `Scrum` or a custom, pre-existing one. Defaults to `Agile`. An empty string will use the parent organization default.
        /// </summary>
        [Input("workItemTemplate")]
        public Input<string>? WorkItemTemplate { get; set; }

        public ProjectState()
        {
        }
        public static new ProjectState Empty => new ProjectState();
    }
}
