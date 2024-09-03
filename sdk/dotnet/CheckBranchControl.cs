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
    /// Manages a branch control check on a resource within Azure DevOps.
    /// 
    /// ## Example Usage
    /// 
    /// ### Protect a service connection
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
    ///     });
    /// 
    ///     var exampleServiceEndpointGeneric = new AzureDevOps.ServiceEndpointGeneric("example", new()
    ///     {
    ///         ProjectId = example.Id,
    ///         ServerUrl = "https://some-server.example.com",
    ///         Username = "username",
    ///         Password = "password",
    ///         ServiceEndpointName = "Example Generic",
    ///         Description = "Managed by Terraform",
    ///     });
    /// 
    ///     var exampleCheckBranchControl = new AzureDevOps.CheckBranchControl("example", new()
    ///     {
    ///         ProjectId = example.Id,
    ///         DisplayName = "Managed by Terraform",
    ///         TargetResourceId = exampleServiceEndpointGeneric.Id,
    ///         TargetResourceType = "endpoint",
    ///         AllowedBranches = "refs/heads/main, refs/heads/features/*",
    ///         Timeout = 1440,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ### Protect an environment
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
    ///     });
    /// 
    ///     var exampleEnvironment = new AzureDevOps.Environment("example", new()
    ///     {
    ///         ProjectId = example.Id,
    ///         Name = "Example Environment",
    ///     });
    /// 
    ///     var exampleCheckBranchControl = new AzureDevOps.CheckBranchControl("example", new()
    ///     {
    ///         ProjectId = example.Id,
    ///         DisplayName = "Managed by Terraform",
    ///         TargetResourceId = exampleEnvironment.Id,
    ///         TargetResourceType = "environment",
    ///         AllowedBranches = "refs/heads/main, refs/heads/features/*",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ### Protect an agent queue
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
    ///     });
    /// 
    ///     var examplePool = new AzureDevOps.Pool("example", new()
    ///     {
    ///         Name = "example-pool",
    ///     });
    /// 
    ///     var exampleQueue = new AzureDevOps.Queue("example", new()
    ///     {
    ///         ProjectId = example.Id,
    ///         AgentPoolId = examplePool.Id,
    ///     });
    /// 
    ///     var exampleCheckBranchControl = new AzureDevOps.CheckBranchControl("example", new()
    ///     {
    ///         ProjectId = example.Id,
    ///         DisplayName = "Managed by Terraform",
    ///         TargetResourceId = exampleQueue.Id,
    ///         TargetResourceType = "queue",
    ///         AllowedBranches = "refs/heads/main, refs/heads/features/*",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ### Protect a repository
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
    ///     });
    /// 
    ///     var exampleGit = new AzureDevOps.Git("example", new()
    ///     {
    ///         ProjectId = example.Id,
    ///         Name = "Example Empty Git Repository",
    ///         Initialization = new AzureDevOps.Inputs.GitInitializationArgs
    ///         {
    ///             InitType = "Clean",
    ///         },
    ///     });
    /// 
    ///     var exampleCheckBranchControl = new AzureDevOps.CheckBranchControl("example", new()
    ///     {
    ///         ProjectId = example.Id,
    ///         DisplayName = "Managed by Terraform",
    ///         TargetResourceId = Output.Tuple(example.Id, exampleGit.Id).Apply(values =&gt;
    ///         {
    ///             var exampleId = values.Item1;
    ///             var exampleGitId = values.Item2;
    ///             return $"{exampleId}.{exampleGitId}";
    ///         }),
    ///         TargetResourceType = "repository",
    ///         AllowedBranches = "refs/heads/main, refs/heads/features/*",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ### Protect a variable group
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
    ///     });
    /// 
    ///     var exampleVariableGroup = new AzureDevOps.VariableGroup("example", new()
    ///     {
    ///         ProjectId = example.Id,
    ///         Name = "Example Variable Group",
    ///         Description = "Example Variable Group Description",
    ///         AllowAccess = true,
    ///         Variables = new[]
    ///         {
    ///             new AzureDevOps.Inputs.VariableGroupVariableArgs
    ///             {
    ///                 Name = "key1",
    ///                 Value = "val1",
    ///             },
    ///             new AzureDevOps.Inputs.VariableGroupVariableArgs
    ///             {
    ///                 Name = "key2",
    ///                 SecretValue = "val2",
    ///                 IsSecret = true,
    ///             },
    ///         },
    ///     });
    /// 
    ///     var exampleCheckBranchControl = new AzureDevOps.CheckBranchControl("example", new()
    ///     {
    ///         ProjectId = example.Id,
    ///         DisplayName = "Managed by Terraform",
    ///         TargetResourceId = exampleVariableGroup.Id,
    ///         TargetResourceType = "variablegroup",
    ///         AllowedBranches = "refs/heads/main, refs/heads/features/*",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Relevant Links
    /// 
    /// - [Define approvals and checks](https://learn.microsoft.com/en-us/azure/devops/pipelines/process/approvals?view=azure-devops&amp;tabs=check-pass)
    /// 
    /// ## Import
    /// 
    /// Importing this resource is not supported.
    /// </summary>
    [AzureDevOpsResourceType("azuredevops:index/checkBranchControl:CheckBranchControl")]
    public partial class CheckBranchControl : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The branches allowed to use the resource. Specify a comma separated list of allowed branches in `refs/heads/branch_name` format. To allow deployments from all branches, specify `*` . `refs/heads/features/* , refs/heads/releases/*` restricts deployments to all branches under features/ or releases/ . Defaults to `*`.
        /// </summary>
        [Output("allowedBranches")]
        public Output<string?> AllowedBranches { get; private set; } = null!;

        /// <summary>
        /// The name of the branch control check displayed in the web UI.
        /// </summary>
        [Output("displayName")]
        public Output<string?> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Allow deployment from branches for which protection status could not be obtained. Only relevant when verify_branch_protection is `true`. Defaults to `false`.
        /// </summary>
        [Output("ignoreUnknownProtectionStatus")]
        public Output<bool?> IgnoreUnknownProtectionStatus { get; private set; } = null!;

        /// <summary>
        /// The project ID.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// The ID of the resource being protected by the check.
        /// </summary>
        [Output("targetResourceId")]
        public Output<string> TargetResourceId { get; private set; } = null!;

        /// <summary>
        /// The type of resource being protected by the check. Valid values: `endpoint`, `environment`, `queue`, `repository`, `securefile`, `variablegroup`.
        /// </summary>
        [Output("targetResourceType")]
        public Output<string> TargetResourceType { get; private set; } = null!;

        /// <summary>
        /// The timeout in minutes for the branch control check. Defaults to `1440`.
        /// </summary>
        [Output("timeout")]
        public Output<int?> Timeout { get; private set; } = null!;

        /// <summary>
        /// Validate the branches being deployed are protected. Defaults to `false`.
        /// </summary>
        [Output("verifyBranchProtection")]
        public Output<bool?> VerifyBranchProtection { get; private set; } = null!;

        /// <summary>
        /// The version of the check.
        /// </summary>
        [Output("version")]
        public Output<int> Version { get; private set; } = null!;


        /// <summary>
        /// Create a CheckBranchControl resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CheckBranchControl(string name, CheckBranchControlArgs args, CustomResourceOptions? options = null)
            : base("azuredevops:index/checkBranchControl:CheckBranchControl", name, args ?? new CheckBranchControlArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CheckBranchControl(string name, Input<string> id, CheckBranchControlState? state = null, CustomResourceOptions? options = null)
            : base("azuredevops:index/checkBranchControl:CheckBranchControl", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing CheckBranchControl resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CheckBranchControl Get(string name, Input<string> id, CheckBranchControlState? state = null, CustomResourceOptions? options = null)
        {
            return new CheckBranchControl(name, id, state, options);
        }
    }

    public sealed class CheckBranchControlArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The branches allowed to use the resource. Specify a comma separated list of allowed branches in `refs/heads/branch_name` format. To allow deployments from all branches, specify `*` . `refs/heads/features/* , refs/heads/releases/*` restricts deployments to all branches under features/ or releases/ . Defaults to `*`.
        /// </summary>
        [Input("allowedBranches")]
        public Input<string>? AllowedBranches { get; set; }

        /// <summary>
        /// The name of the branch control check displayed in the web UI.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Allow deployment from branches for which protection status could not be obtained. Only relevant when verify_branch_protection is `true`. Defaults to `false`.
        /// </summary>
        [Input("ignoreUnknownProtectionStatus")]
        public Input<bool>? IgnoreUnknownProtectionStatus { get; set; }

        /// <summary>
        /// The project ID.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        /// <summary>
        /// The ID of the resource being protected by the check.
        /// </summary>
        [Input("targetResourceId", required: true)]
        public Input<string> TargetResourceId { get; set; } = null!;

        /// <summary>
        /// The type of resource being protected by the check. Valid values: `endpoint`, `environment`, `queue`, `repository`, `securefile`, `variablegroup`.
        /// </summary>
        [Input("targetResourceType", required: true)]
        public Input<string> TargetResourceType { get; set; } = null!;

        /// <summary>
        /// The timeout in minutes for the branch control check. Defaults to `1440`.
        /// </summary>
        [Input("timeout")]
        public Input<int>? Timeout { get; set; }

        /// <summary>
        /// Validate the branches being deployed are protected. Defaults to `false`.
        /// </summary>
        [Input("verifyBranchProtection")]
        public Input<bool>? VerifyBranchProtection { get; set; }

        public CheckBranchControlArgs()
        {
        }
        public static new CheckBranchControlArgs Empty => new CheckBranchControlArgs();
    }

    public sealed class CheckBranchControlState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The branches allowed to use the resource. Specify a comma separated list of allowed branches in `refs/heads/branch_name` format. To allow deployments from all branches, specify `*` . `refs/heads/features/* , refs/heads/releases/*` restricts deployments to all branches under features/ or releases/ . Defaults to `*`.
        /// </summary>
        [Input("allowedBranches")]
        public Input<string>? AllowedBranches { get; set; }

        /// <summary>
        /// The name of the branch control check displayed in the web UI.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Allow deployment from branches for which protection status could not be obtained. Only relevant when verify_branch_protection is `true`. Defaults to `false`.
        /// </summary>
        [Input("ignoreUnknownProtectionStatus")]
        public Input<bool>? IgnoreUnknownProtectionStatus { get; set; }

        /// <summary>
        /// The project ID.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The ID of the resource being protected by the check.
        /// </summary>
        [Input("targetResourceId")]
        public Input<string>? TargetResourceId { get; set; }

        /// <summary>
        /// The type of resource being protected by the check. Valid values: `endpoint`, `environment`, `queue`, `repository`, `securefile`, `variablegroup`.
        /// </summary>
        [Input("targetResourceType")]
        public Input<string>? TargetResourceType { get; set; }

        /// <summary>
        /// The timeout in minutes for the branch control check. Defaults to `1440`.
        /// </summary>
        [Input("timeout")]
        public Input<int>? Timeout { get; set; }

        /// <summary>
        /// Validate the branches being deployed are protected. Defaults to `false`.
        /// </summary>
        [Input("verifyBranchProtection")]
        public Input<bool>? VerifyBranchProtection { get; set; }

        /// <summary>
        /// The version of the check.
        /// </summary>
        [Input("version")]
        public Input<int>? Version { get; set; }

        public CheckBranchControlState()
        {
        }
        public static new CheckBranchControlState Empty => new CheckBranchControlState();
    }
}
