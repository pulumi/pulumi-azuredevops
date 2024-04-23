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
    /// Manages features for Azure DevOps projects
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
    ///         Name = "Example Project",
    ///         Visibility = "private",
    ///         VersionControl = "Git",
    ///         WorkItemTemplate = "Agile",
    ///         Description = "Managed by Terraform",
    ///     });
    /// 
    ///     var example_features = new AzureDevOps.ProjectFeatures("example-features", new()
    ///     {
    ///         ProjectId = example.Id,
    ///         Features = 
    ///         {
    ///             { "testplans", "disabled" },
    ///             { "artifacts", "enabled" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Relevant Links
    /// 
    /// No official documentation available
    /// 
    /// ## PAT Permissions Required
    /// 
    /// - **Project &amp; Team**: Read, Write, &amp; Manage
    /// 
    /// ## Import
    /// 
    /// Azure DevOps feature settings can be imported using the project id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import azuredevops:index/projectFeatures:ProjectFeatures example 00000000-0000-0000-0000-000000000000
    /// ```
    /// </summary>
    [AzureDevOpsResourceType("azuredevops:index/projectFeatures:ProjectFeatures")]
    public partial class ProjectFeatures : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Defines the status (`enabled`, `disabled`) of the project features.  
        /// Valid features `boards`, `repositories`, `pipelines`, `testplans`, `artifacts`
        /// 
        /// &gt; **NOTE:**
        /// &gt; It's possible to define project features both within the `azuredevops.ProjectFeatures` resource and
        /// &gt; via the `features` block by using the `azuredevops.Project` resource.
        /// &gt; However it's not possible to use both methods to manage features, since there'll be conflicts.
        /// </summary>
        [Output("features")]
        public Output<ImmutableDictionary<string, string>> Features { get; private set; } = null!;

        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;


        /// <summary>
        /// Create a ProjectFeatures resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ProjectFeatures(string name, ProjectFeaturesArgs args, CustomResourceOptions? options = null)
            : base("azuredevops:index/projectFeatures:ProjectFeatures", name, args ?? new ProjectFeaturesArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ProjectFeatures(string name, Input<string> id, ProjectFeaturesState? state = null, CustomResourceOptions? options = null)
            : base("azuredevops:index/projectFeatures:ProjectFeatures", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ProjectFeatures resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ProjectFeatures Get(string name, Input<string> id, ProjectFeaturesState? state = null, CustomResourceOptions? options = null)
        {
            return new ProjectFeatures(name, id, state, options);
        }
    }

    public sealed class ProjectFeaturesArgs : global::Pulumi.ResourceArgs
    {
        [Input("features", required: true)]
        private InputMap<string>? _features;

        /// <summary>
        /// Defines the status (`enabled`, `disabled`) of the project features.  
        /// Valid features `boards`, `repositories`, `pipelines`, `testplans`, `artifacts`
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

        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        public ProjectFeaturesArgs()
        {
        }
        public static new ProjectFeaturesArgs Empty => new ProjectFeaturesArgs();
    }

    public sealed class ProjectFeaturesState : global::Pulumi.ResourceArgs
    {
        [Input("features")]
        private InputMap<string>? _features;

        /// <summary>
        /// Defines the status (`enabled`, `disabled`) of the project features.  
        /// Valid features `boards`, `repositories`, `pipelines`, `testplans`, `artifacts`
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

        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        public ProjectFeaturesState()
        {
        }
        public static new ProjectFeaturesState Empty => new ProjectFeaturesState();
    }
}
