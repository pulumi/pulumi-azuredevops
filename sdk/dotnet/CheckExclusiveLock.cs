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
    /// Manages a Exclusive Lock Check.
    /// 
    /// Adding an exclusive lock will only allow a single stage to utilize this resource at a time. If multiple stages are waiting on the lock, only the latest will run. All others will be canceled.
    /// 
    /// ## Example Usage
    /// 
    /// ### Add Exclusive Lock to an environment
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
    ///     var exampleCheckExclusiveLock = new AzureDevOps.CheckExclusiveLock("exampleCheckExclusiveLock", new()
    ///     {
    ///         ProjectId = exampleProject.Id,
    ///         TargetResourceId = exampleServiceEndpointGeneric.Id,
    ///         TargetResourceType = "endpoint",
    ///         Timeout = 43200,
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
    ///     var exampleCheckExclusiveLock = new AzureDevOps.CheckExclusiveLock("exampleCheckExclusiveLock", new()
    ///     {
    ///         ProjectId = exampleProject.Id,
    ///         TargetResourceId = exampleEnvironment.Id,
    ///         TargetResourceType = "environment",
    ///         Timeout = 43200,
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
    [AzureDevOpsResourceType("azuredevops:index/checkExclusiveLock:CheckExclusiveLock")]
    public partial class CheckExclusiveLock : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The project ID. Changing this forces a new Exclusive Lock Check to be created.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// The ID of the resource being protected by the check. Changing this forces a new Exclusive Lock to be created.
        /// </summary>
        [Output("targetResourceId")]
        public Output<string> TargetResourceId { get; private set; } = null!;

        /// <summary>
        /// The type of resource being protected by the check. Valid values: `endpoint`, `environment`, `queue`, `repository`, `securefile`, `variablegroup`. Changing this forces a new Exclusive Lock to be created.
        /// </summary>
        [Output("targetResourceType")]
        public Output<string> TargetResourceType { get; private set; } = null!;

        /// <summary>
        /// The timeout in minutes for the exclusive lock. Defaults to `43200`.
        /// </summary>
        [Output("timeout")]
        public Output<int?> Timeout { get; private set; } = null!;


        /// <summary>
        /// Create a CheckExclusiveLock resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CheckExclusiveLock(string name, CheckExclusiveLockArgs args, CustomResourceOptions? options = null)
            : base("azuredevops:index/checkExclusiveLock:CheckExclusiveLock", name, args ?? new CheckExclusiveLockArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CheckExclusiveLock(string name, Input<string> id, CheckExclusiveLockState? state = null, CustomResourceOptions? options = null)
            : base("azuredevops:index/checkExclusiveLock:CheckExclusiveLock", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing CheckExclusiveLock resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CheckExclusiveLock Get(string name, Input<string> id, CheckExclusiveLockState? state = null, CustomResourceOptions? options = null)
        {
            return new CheckExclusiveLock(name, id, state, options);
        }
    }

    public sealed class CheckExclusiveLockArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The project ID. Changing this forces a new Exclusive Lock Check to be created.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        /// <summary>
        /// The ID of the resource being protected by the check. Changing this forces a new Exclusive Lock to be created.
        /// </summary>
        [Input("targetResourceId", required: true)]
        public Input<string> TargetResourceId { get; set; } = null!;

        /// <summary>
        /// The type of resource being protected by the check. Valid values: `endpoint`, `environment`, `queue`, `repository`, `securefile`, `variablegroup`. Changing this forces a new Exclusive Lock to be created.
        /// </summary>
        [Input("targetResourceType", required: true)]
        public Input<string> TargetResourceType { get; set; } = null!;

        /// <summary>
        /// The timeout in minutes for the exclusive lock. Defaults to `43200`.
        /// </summary>
        [Input("timeout")]
        public Input<int>? Timeout { get; set; }

        public CheckExclusiveLockArgs()
        {
        }
        public static new CheckExclusiveLockArgs Empty => new CheckExclusiveLockArgs();
    }

    public sealed class CheckExclusiveLockState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The project ID. Changing this forces a new Exclusive Lock Check to be created.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The ID of the resource being protected by the check. Changing this forces a new Exclusive Lock to be created.
        /// </summary>
        [Input("targetResourceId")]
        public Input<string>? TargetResourceId { get; set; }

        /// <summary>
        /// The type of resource being protected by the check. Valid values: `endpoint`, `environment`, `queue`, `repository`, `securefile`, `variablegroup`. Changing this forces a new Exclusive Lock to be created.
        /// </summary>
        [Input("targetResourceType")]
        public Input<string>? TargetResourceType { get; set; }

        /// <summary>
        /// The timeout in minutes for the exclusive lock. Defaults to `43200`.
        /// </summary>
        [Input("timeout")]
        public Input<int>? Timeout { get; set; }

        public CheckExclusiveLockState()
        {
        }
        public static new CheckExclusiveLockState Empty => new CheckExclusiveLockState();
    }
}
