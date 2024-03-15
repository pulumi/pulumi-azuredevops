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
    /// ## Example Usage
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
    ///     var exampleProject = new AzureDevOps.Project("exampleProject", new()
    ///     {
    ///         WorkItemTemplate = "Agile",
    ///         VersionControl = "Git",
    ///         Visibility = "private",
    ///         Description = "Managed by Terraform",
    ///     });
    /// 
    ///     var exampleVariableGroup = new AzureDevOps.VariableGroup("exampleVariableGroup", new()
    ///     {
    ///         ProjectId = exampleProject.Id,
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
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ### With AzureRM Key Vault
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
    ///     var exampleProject = new AzureDevOps.Project("exampleProject", new()
    ///     {
    ///         WorkItemTemplate = "Agile",
    ///         VersionControl = "Git",
    ///         Visibility = "private",
    ///         Description = "Managed by Terraform",
    ///     });
    /// 
    ///     var exampleServiceEndpointAzureRM = new AzureDevOps.ServiceEndpointAzureRM("exampleServiceEndpointAzureRM", new()
    ///     {
    ///         ProjectId = exampleProject.Id,
    ///         ServiceEndpointName = "Example AzureRM",
    ///         Description = "Managed by Terraform",
    ///         Credentials = new AzureDevOps.Inputs.ServiceEndpointAzureRMCredentialsArgs
    ///         {
    ///             Serviceprincipalid = "00000000-0000-0000-0000-000000000000",
    ///             Serviceprincipalkey = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
    ///         },
    ///         AzurermSpnTenantid = "00000000-0000-0000-0000-000000000000",
    ///         AzurermSubscriptionId = "00000000-0000-0000-0000-000000000000",
    ///         AzurermSubscriptionName = "Example Subscription Name",
    ///     });
    /// 
    ///     var exampleVariableGroup = new AzureDevOps.VariableGroup("exampleVariableGroup", new()
    ///     {
    ///         ProjectId = exampleProject.Id,
    ///         Description = "Example Variable Group Description",
    ///         AllowAccess = true,
    ///         KeyVault = new AzureDevOps.Inputs.VariableGroupKeyVaultArgs
    ///         {
    ///             Name = "example-kv",
    ///             ServiceEndpointId = exampleServiceEndpointAzureRM.Id,
    ///         },
    ///         Variables = new[]
    ///         {
    ///             new AzureDevOps.Inputs.VariableGroupVariableArgs
    ///             {
    ///                 Name = "key1",
    ///             },
    ///             new AzureDevOps.Inputs.VariableGroupVariableArgs
    ///             {
    ///                 Name = "key2",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Relevant Links
    /// 
    /// - [Azure DevOps Service REST API 7.0 - Variable Groups](https://docs.microsoft.com/en-us/rest/api/azure/devops/distributedtask/variablegroups?view=azure-devops-rest-7.0)
    /// - [Azure DevOps Service REST API 7.0 - Authorized Resources](https://docs.microsoft.com/en-us/rest/api/azure/devops/build/authorizedresources?view=azure-devops-rest-7.0)
    /// 
    /// ## PAT Permissions Required
    /// 
    /// - **Variable Groups**: Read, Create, &amp; Manage
    /// - **Build**: Read &amp; execute
    /// - **Project and Team**: Read
    /// - **Token Administration**: Read &amp; manage
    /// - **Tokens**: Read &amp; manage
    /// - **Work Items**: Read
    /// 
    /// ## Import
    /// 
    /// **Variable groups containing secret values cannot be imported.**
    /// 
    /// Azure DevOps Variable groups can be imported using the project name/variable group ID or by the project Guid/variable group ID, e.g.
    /// 
    /// ```sh
    /// $ pulumi import azuredevops:index/variableGroup:VariableGroup example "Example Project/10"
    /// ```
    /// 
    /// or
    /// 
    /// ```sh
    /// $ pulumi import azuredevops:index/variableGroup:VariableGroup example 00000000-0000-0000-0000-000000000000/0
    /// ```
    /// 
    /// _Note that for secret variables, the import command retrieve blank value in the tfstate._
    /// </summary>
    [AzureDevOpsResourceType("azuredevops:index/variableGroup:VariableGroup")]
    public partial class VariableGroup : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Boolean that indicate if this variable group is shared by all pipelines of this project.
        /// </summary>
        [Output("allowAccess")]
        public Output<bool?> AllowAccess { get; private set; } = null!;

        /// <summary>
        /// The description of the Variable Group.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// A list of `key_vault` blocks as documented below.
        /// </summary>
        [Output("keyVault")]
        public Output<Outputs.VariableGroupKeyVault?> KeyVault { get; private set; } = null!;

        /// <summary>
        /// The name of the Variable Group.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The ID of the project.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// One or more `variable` blocks as documented below.
        /// </summary>
        [Output("variables")]
        public Output<ImmutableArray<Outputs.VariableGroupVariable>> Variables { get; private set; } = null!;


        /// <summary>
        /// Create a VariableGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VariableGroup(string name, VariableGroupArgs args, CustomResourceOptions? options = null)
            : base("azuredevops:index/variableGroup:VariableGroup", name, args ?? new VariableGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VariableGroup(string name, Input<string> id, VariableGroupState? state = null, CustomResourceOptions? options = null)
            : base("azuredevops:index/variableGroup:VariableGroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing VariableGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VariableGroup Get(string name, Input<string> id, VariableGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new VariableGroup(name, id, state, options);
        }
    }

    public sealed class VariableGroupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Boolean that indicate if this variable group is shared by all pipelines of this project.
        /// </summary>
        [Input("allowAccess")]
        public Input<bool>? AllowAccess { get; set; }

        /// <summary>
        /// The description of the Variable Group.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// A list of `key_vault` blocks as documented below.
        /// </summary>
        [Input("keyVault")]
        public Input<Inputs.VariableGroupKeyVaultArgs>? KeyVault { get; set; }

        /// <summary>
        /// The name of the Variable Group.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        [Input("variables", required: true)]
        private InputList<Inputs.VariableGroupVariableArgs>? _variables;

        /// <summary>
        /// One or more `variable` blocks as documented below.
        /// </summary>
        public InputList<Inputs.VariableGroupVariableArgs> Variables
        {
            get => _variables ?? (_variables = new InputList<Inputs.VariableGroupVariableArgs>());
            set => _variables = value;
        }

        public VariableGroupArgs()
        {
        }
        public static new VariableGroupArgs Empty => new VariableGroupArgs();
    }

    public sealed class VariableGroupState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Boolean that indicate if this variable group is shared by all pipelines of this project.
        /// </summary>
        [Input("allowAccess")]
        public Input<bool>? AllowAccess { get; set; }

        /// <summary>
        /// The description of the Variable Group.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// A list of `key_vault` blocks as documented below.
        /// </summary>
        [Input("keyVault")]
        public Input<Inputs.VariableGroupKeyVaultGetArgs>? KeyVault { get; set; }

        /// <summary>
        /// The name of the Variable Group.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        [Input("variables")]
        private InputList<Inputs.VariableGroupVariableGetArgs>? _variables;

        /// <summary>
        /// One or more `variable` blocks as documented below.
        /// </summary>
        public InputList<Inputs.VariableGroupVariableGetArgs> Variables
        {
            get => _variables ?? (_variables = new InputList<Inputs.VariableGroupVariableGetArgs>());
            set => _variables = value;
        }

        public VariableGroupState()
        {
        }
        public static new VariableGroupState Empty => new VariableGroupState();
    }
}
