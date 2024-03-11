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
    /// Manages a Required Template Check.
    /// 
    /// ## Example Usage
    /// 
    /// ### Protect a service connection
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AzureDevOps = Pulumi.AzureDevOps;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var exampleProject = new AzureDevOps.Project("exampleProject");
    /// 
    ///     var exampleServiceEndpointGeneric = new AzureDevOps.ServiceEndpointGeneric("exampleServiceEndpointGeneric", new()
    ///     {
    ///         ProjectId = exampleProject.Id,
    ///         ServerUrl = "https://some-server.example.com",
    ///         Username = "username",
    ///         Password = "password",
    ///         ServiceEndpointName = "Example Generic",
    ///         Description = "Managed by Terraform",
    ///     });
    /// 
    ///     var exampleCheckRequiredTemplate = new AzureDevOps.CheckRequiredTemplate("exampleCheckRequiredTemplate", new()
    ///     {
    ///         ProjectId = exampleProject.Id,
    ///         TargetResourceId = exampleServiceEndpointGeneric.Id,
    ///         TargetResourceType = "endpoint",
    ///         RequiredTemplates = new[]
    ///         {
    ///             new AzureDevOps.Inputs.CheckRequiredTemplateRequiredTemplateArgs
    ///             {
    ///                 RepositoryType = "azuregit",
    ///                 RepositoryName = "project/repository",
    ///                 RepositoryRef = "refs/heads/main",
    ///                 TemplatePath = "template/path.yml",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ### Protect an environment
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AzureDevOps = Pulumi.AzureDevOps;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var exampleProject = new AzureDevOps.Project("exampleProject");
    /// 
    ///     var exampleEnvironment = new AzureDevOps.Environment("exampleEnvironment", new()
    ///     {
    ///         ProjectId = exampleProject.Id,
    ///     });
    /// 
    ///     var exampleCheckRequiredTemplate = new AzureDevOps.CheckRequiredTemplate("exampleCheckRequiredTemplate", new()
    ///     {
    ///         ProjectId = exampleProject.Id,
    ///         TargetResourceId = exampleEnvironment.Id,
    ///         TargetResourceType = "environment",
    ///         RequiredTemplates = new[]
    ///         {
    ///             new AzureDevOps.Inputs.CheckRequiredTemplateRequiredTemplateArgs
    ///             {
    ///                 RepositoryName = "project/repository",
    ///                 RepositoryRef = "refs/heads/main",
    ///                 TemplatePath = "template/path.yml",
    ///             },
    ///             new AzureDevOps.Inputs.CheckRequiredTemplateRequiredTemplateArgs
    ///             {
    ///                 RepositoryName = "project/repository",
    ///                 RepositoryRef = "refs/heads/main",
    ///                 TemplatePath = "template/alternate-path.yml",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Importing this resource is not supported.
    /// </summary>
    [AzureDevOpsResourceType("azuredevops:index/checkRequiredTemplate:CheckRequiredTemplate")]
    public partial class CheckRequiredTemplate : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The project ID. Changing this forces a new Required Template Check to be created.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// One or more `required_template` blocks documented below.
        /// </summary>
        [Output("requiredTemplates")]
        public Output<ImmutableArray<Outputs.CheckRequiredTemplateRequiredTemplate>> RequiredTemplates { get; private set; } = null!;

        /// <summary>
        /// The ID of the resource being protected by the check. Changing this forces a new Required Template Check to be created.
        /// </summary>
        [Output("targetResourceId")]
        public Output<string> TargetResourceId { get; private set; } = null!;

        /// <summary>
        /// The type of resource being protected by the check. Valid values: `endpoint`, `environment`, `queue`, `repository`, `securefile`, `variablegroup`. Changing this forces a new Required Template Check to be created.
        /// </summary>
        [Output("targetResourceType")]
        public Output<string> TargetResourceType { get; private set; } = null!;


        /// <summary>
        /// Create a CheckRequiredTemplate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CheckRequiredTemplate(string name, CheckRequiredTemplateArgs args, CustomResourceOptions? options = null)
            : base("azuredevops:index/checkRequiredTemplate:CheckRequiredTemplate", name, args ?? new CheckRequiredTemplateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CheckRequiredTemplate(string name, Input<string> id, CheckRequiredTemplateState? state = null, CustomResourceOptions? options = null)
            : base("azuredevops:index/checkRequiredTemplate:CheckRequiredTemplate", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing CheckRequiredTemplate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CheckRequiredTemplate Get(string name, Input<string> id, CheckRequiredTemplateState? state = null, CustomResourceOptions? options = null)
        {
            return new CheckRequiredTemplate(name, id, state, options);
        }
    }

    public sealed class CheckRequiredTemplateArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The project ID. Changing this forces a new Required Template Check to be created.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        [Input("requiredTemplates", required: true)]
        private InputList<Inputs.CheckRequiredTemplateRequiredTemplateArgs>? _requiredTemplates;

        /// <summary>
        /// One or more `required_template` blocks documented below.
        /// </summary>
        public InputList<Inputs.CheckRequiredTemplateRequiredTemplateArgs> RequiredTemplates
        {
            get => _requiredTemplates ?? (_requiredTemplates = new InputList<Inputs.CheckRequiredTemplateRequiredTemplateArgs>());
            set => _requiredTemplates = value;
        }

        /// <summary>
        /// The ID of the resource being protected by the check. Changing this forces a new Required Template Check to be created.
        /// </summary>
        [Input("targetResourceId", required: true)]
        public Input<string> TargetResourceId { get; set; } = null!;

        /// <summary>
        /// The type of resource being protected by the check. Valid values: `endpoint`, `environment`, `queue`, `repository`, `securefile`, `variablegroup`. Changing this forces a new Required Template Check to be created.
        /// </summary>
        [Input("targetResourceType", required: true)]
        public Input<string> TargetResourceType { get; set; } = null!;

        public CheckRequiredTemplateArgs()
        {
        }
        public static new CheckRequiredTemplateArgs Empty => new CheckRequiredTemplateArgs();
    }

    public sealed class CheckRequiredTemplateState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The project ID. Changing this forces a new Required Template Check to be created.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        [Input("requiredTemplates")]
        private InputList<Inputs.CheckRequiredTemplateRequiredTemplateGetArgs>? _requiredTemplates;

        /// <summary>
        /// One or more `required_template` blocks documented below.
        /// </summary>
        public InputList<Inputs.CheckRequiredTemplateRequiredTemplateGetArgs> RequiredTemplates
        {
            get => _requiredTemplates ?? (_requiredTemplates = new InputList<Inputs.CheckRequiredTemplateRequiredTemplateGetArgs>());
            set => _requiredTemplates = value;
        }

        /// <summary>
        /// The ID of the resource being protected by the check. Changing this forces a new Required Template Check to be created.
        /// </summary>
        [Input("targetResourceId")]
        public Input<string>? TargetResourceId { get; set; }

        /// <summary>
        /// The type of resource being protected by the check. Valid values: `endpoint`, `environment`, `queue`, `repository`, `securefile`, `variablegroup`. Changing this forces a new Required Template Check to be created.
        /// </summary>
        [Input("targetResourceType")]
        public Input<string>? TargetResourceType { get; set; }

        public CheckRequiredTemplateState()
        {
        }
        public static new CheckRequiredTemplateState Empty => new CheckRequiredTemplateState();
    }
}
