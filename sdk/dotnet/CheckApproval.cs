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
    /// Manages a Approval Check.
    /// 
    /// ## Example Usage
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
    ///     var exampleGroup = new AzureDevOps.Group("exampleGroup", new()
    ///     {
    ///         DisplayName = "some-azdo-group",
    ///     });
    /// 
    ///     var exampleCheckApproval = new AzureDevOps.CheckApproval("exampleCheckApproval", new()
    ///     {
    ///         ProjectId = exampleProject.Id,
    ///         TargetResourceId = exampleEnvironment.Id,
    ///         TargetResourceType = "environment",
    ///         RequesterCanApprove = true,
    ///         Approvers = new[]
    ///         {
    ///             exampleGroup.OriginId,
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
    [AzureDevOpsResourceType("azuredevops:index/checkApproval:CheckApproval")]
    public partial class CheckApproval : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies a list of approver IDs.
        /// </summary>
        [Output("approvers")]
        public Output<ImmutableArray<string>> Approvers { get; private set; } = null!;

        /// <summary>
        /// The instructions for the approvers.
        /// </summary>
        [Output("instructions")]
        public Output<string?> Instructions { get; private set; } = null!;

        /// <summary>
        /// The minimum number of approvers. This property is applicable when there is more than 1 approver.
        /// </summary>
        [Output("minimumRequiredApprovers")]
        public Output<int?> MinimumRequiredApprovers { get; private set; } = null!;

        /// <summary>
        /// The project ID. Changing this forces a new Approval Check to be created.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// Can the requestor approve? Defaults to `false`.
        /// </summary>
        [Output("requesterCanApprove")]
        public Output<bool?> RequesterCanApprove { get; private set; } = null!;

        /// <summary>
        /// The ID of the resource being protected by the check. Changing this forces a new Approval Check to be created.
        /// </summary>
        [Output("targetResourceId")]
        public Output<string> TargetResourceId { get; private set; } = null!;

        /// <summary>
        /// The type of resource being protected by the check. Valid values: `endpoint`, `environment`, `queue`, `repository`, `securefile`, `variablegroup`. Changing this forces a new Approval Check to be created.
        /// </summary>
        [Output("targetResourceType")]
        public Output<string> TargetResourceType { get; private set; } = null!;

        /// <summary>
        /// The timeout in minutes for the approval.  Defaults to `43200`.
        /// </summary>
        [Output("timeout")]
        public Output<int?> Timeout { get; private set; } = null!;

        /// <summary>
        /// The version of the check.
        /// </summary>
        [Output("version")]
        public Output<int> Version { get; private set; } = null!;


        /// <summary>
        /// Create a CheckApproval resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CheckApproval(string name, CheckApprovalArgs args, CustomResourceOptions? options = null)
            : base("azuredevops:index/checkApproval:CheckApproval", name, args ?? new CheckApprovalArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CheckApproval(string name, Input<string> id, CheckApprovalState? state = null, CustomResourceOptions? options = null)
            : base("azuredevops:index/checkApproval:CheckApproval", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing CheckApproval resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CheckApproval Get(string name, Input<string> id, CheckApprovalState? state = null, CustomResourceOptions? options = null)
        {
            return new CheckApproval(name, id, state, options);
        }
    }

    public sealed class CheckApprovalArgs : global::Pulumi.ResourceArgs
    {
        [Input("approvers", required: true)]
        private InputList<string>? _approvers;

        /// <summary>
        /// Specifies a list of approver IDs.
        /// </summary>
        public InputList<string> Approvers
        {
            get => _approvers ?? (_approvers = new InputList<string>());
            set => _approvers = value;
        }

        /// <summary>
        /// The instructions for the approvers.
        /// </summary>
        [Input("instructions")]
        public Input<string>? Instructions { get; set; }

        /// <summary>
        /// The minimum number of approvers. This property is applicable when there is more than 1 approver.
        /// </summary>
        [Input("minimumRequiredApprovers")]
        public Input<int>? MinimumRequiredApprovers { get; set; }

        /// <summary>
        /// The project ID. Changing this forces a new Approval Check to be created.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        /// <summary>
        /// Can the requestor approve? Defaults to `false`.
        /// </summary>
        [Input("requesterCanApprove")]
        public Input<bool>? RequesterCanApprove { get; set; }

        /// <summary>
        /// The ID of the resource being protected by the check. Changing this forces a new Approval Check to be created.
        /// </summary>
        [Input("targetResourceId", required: true)]
        public Input<string> TargetResourceId { get; set; } = null!;

        /// <summary>
        /// The type of resource being protected by the check. Valid values: `endpoint`, `environment`, `queue`, `repository`, `securefile`, `variablegroup`. Changing this forces a new Approval Check to be created.
        /// </summary>
        [Input("targetResourceType", required: true)]
        public Input<string> TargetResourceType { get; set; } = null!;

        /// <summary>
        /// The timeout in minutes for the approval.  Defaults to `43200`.
        /// </summary>
        [Input("timeout")]
        public Input<int>? Timeout { get; set; }

        public CheckApprovalArgs()
        {
        }
        public static new CheckApprovalArgs Empty => new CheckApprovalArgs();
    }

    public sealed class CheckApprovalState : global::Pulumi.ResourceArgs
    {
        [Input("approvers")]
        private InputList<string>? _approvers;

        /// <summary>
        /// Specifies a list of approver IDs.
        /// </summary>
        public InputList<string> Approvers
        {
            get => _approvers ?? (_approvers = new InputList<string>());
            set => _approvers = value;
        }

        /// <summary>
        /// The instructions for the approvers.
        /// </summary>
        [Input("instructions")]
        public Input<string>? Instructions { get; set; }

        /// <summary>
        /// The minimum number of approvers. This property is applicable when there is more than 1 approver.
        /// </summary>
        [Input("minimumRequiredApprovers")]
        public Input<int>? MinimumRequiredApprovers { get; set; }

        /// <summary>
        /// The project ID. Changing this forces a new Approval Check to be created.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// Can the requestor approve? Defaults to `false`.
        /// </summary>
        [Input("requesterCanApprove")]
        public Input<bool>? RequesterCanApprove { get; set; }

        /// <summary>
        /// The ID of the resource being protected by the check. Changing this forces a new Approval Check to be created.
        /// </summary>
        [Input("targetResourceId")]
        public Input<string>? TargetResourceId { get; set; }

        /// <summary>
        /// The type of resource being protected by the check. Valid values: `endpoint`, `environment`, `queue`, `repository`, `securefile`, `variablegroup`. Changing this forces a new Approval Check to be created.
        /// </summary>
        [Input("targetResourceType")]
        public Input<string>? TargetResourceType { get; set; }

        /// <summary>
        /// The timeout in minutes for the approval.  Defaults to `43200`.
        /// </summary>
        [Input("timeout")]
        public Input<int>? Timeout { get; set; }

        /// <summary>
        /// The version of the check.
        /// </summary>
        [Input("version")]
        public Input<int>? Version { get; set; }

        public CheckApprovalState()
        {
        }
        public static new CheckApprovalState Empty => new CheckApprovalState();
    }
}
