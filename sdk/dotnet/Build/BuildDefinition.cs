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
    /// ## Relevant Links
    /// 
    /// * [Azure DevOps Service REST API 5.1 - Build Definitions](https://docs.microsoft.com/en-us/rest/api/azure/devops/build/definitions?view=azure-devops-rest-5.1)
    /// </summary>
    public partial class BuildDefinition : Pulumi.CustomResource
    {
        /// <summary>
        /// The agent pool that should execute the build. Defaults to `Hosted Ubuntu 1604`.
        /// </summary>
        [Output("agentPoolName")]
        public Output<string?> AgentPoolName { get; private set; } = null!;

        /// <summary>
        /// Continuous Integration Integration trigger.
        /// </summary>
        [Output("ciTrigger")]
        public Output<Outputs.BuildDefinitionCiTrigger?> CiTrigger { get; private set; } = null!;

        /// <summary>
        /// The name of the build definition.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

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
        /// The agent pool that should execute the build. Defaults to `Hosted Ubuntu 1604`.
        /// </summary>
        [Input("agentPoolName")]
        public Input<string>? AgentPoolName { get; set; }

        /// <summary>
        /// Continuous Integration Integration trigger.
        /// </summary>
        [Input("ciTrigger")]
        public Input<Inputs.BuildDefinitionCiTriggerArgs>? CiTrigger { get; set; }

        /// <summary>
        /// The name of the build definition.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

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
        /// The agent pool that should execute the build. Defaults to `Hosted Ubuntu 1604`.
        /// </summary>
        [Input("agentPoolName")]
        public Input<string>? AgentPoolName { get; set; }

        /// <summary>
        /// Continuous Integration Integration trigger.
        /// </summary>
        [Input("ciTrigger")]
        public Input<Inputs.BuildDefinitionCiTriggerGetArgs>? CiTrigger { get; set; }

        /// <summary>
        /// The name of the build definition.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

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
