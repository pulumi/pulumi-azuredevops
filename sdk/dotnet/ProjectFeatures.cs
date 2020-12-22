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
    /// using Pulumi;
    /// using AzureDevOps = Pulumi.AzureDevOps;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var tf_project_test_001 = Output.Create(AzureDevOps.GetProject.InvokeAsync(new AzureDevOps.GetProjectArgs
    ///         {
    ///             Name = "Test Project",
    ///         }));
    ///         var my_project_features = new AzureDevOps.ProjectFeatures("my-project-features", new AzureDevOps.ProjectFeaturesArgs
    ///         {
    ///             ProjectId = tf_project_test_001.Apply(tf_project_test_001 =&gt; tf_project_test_001.Id),
    ///             Features = 
    ///             {
    ///                 { "testplans", "disabled" },
    ///                 { "artifacts", "enabled" },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
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
    ///  $ pulumi import azuredevops:index/projectFeatures:ProjectFeatures project_id 2785562e-8f45-4534-a10e-b9ca1666b17e
    /// ```
    /// </summary>
    public partial class ProjectFeatures : Pulumi.CustomResource
    {
        /// <summary>
        /// Defines the status (`enabled`, `disabled`) of the project features.  
        /// Valid features `boards`, `repositories`, `pipelines`, `testplans`, `artifacts`
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
                Aliases =
                {
                    new Pulumi.Alias { Type = "azuredevops:Core/projectFeatures:ProjectFeatures"},
                },
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

    public sealed class ProjectFeaturesArgs : Pulumi.ResourceArgs
    {
        [Input("features", required: true)]
        private InputMap<string>? _features;

        /// <summary>
        /// Defines the status (`enabled`, `disabled`) of the project features.  
        /// Valid features `boards`, `repositories`, `pipelines`, `testplans`, `artifacts`
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
    }

    public sealed class ProjectFeaturesState : Pulumi.ResourceArgs
    {
        [Input("features")]
        private InputMap<string>? _features;

        /// <summary>
        /// Defines the status (`enabled`, `disabled`) of the project features.  
        /// Valid features `boards`, `repositories`, `pipelines`, `testplans`, `artifacts`
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
    }
}
