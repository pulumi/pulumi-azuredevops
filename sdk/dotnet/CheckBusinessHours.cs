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
    /// Manages a business hours check on a resource within Azure DevOps.
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
    ///     var exampleCheckBusinessHours = new AzureDevOps.CheckBusinessHours("exampleCheckBusinessHours", new()
    ///     {
    ///         ProjectId = exampleProject.Id,
    ///         DisplayName = "Managed by Terraform",
    ///         TargetResourceId = exampleServiceEndpointGeneric.Id,
    ///         TargetResourceType = "endpoint",
    ///         StartTime = "07:00",
    ///         EndTime = "15:30",
    ///         TimeZone = "UTC",
    ///         Monday = true,
    ///         Tuesday = true,
    ///         Timeout = 1440,
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
    ///     var exampleCheckBusinessHours = new AzureDevOps.CheckBusinessHours("exampleCheckBusinessHours", new()
    ///     {
    ///         ProjectId = exampleProject.Id,
    ///         DisplayName = "Managed by Terraform",
    ///         TargetResourceId = exampleEnvironment.Id,
    ///         TargetResourceType = "environment",
    ///         StartTime = "07:00",
    ///         EndTime = "15:30",
    ///         TimeZone = "UTC",
    ///         Monday = true,
    ///         Tuesday = true,
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ### Protect an agent queue
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
    ///     var examplePool = new AzureDevOps.Pool("examplePool");
    /// 
    ///     var exampleQueue = new AzureDevOps.Queue("exampleQueue", new()
    ///     {
    ///         ProjectId = exampleProject.Id,
    ///         AgentPoolId = examplePool.Id,
    ///     });
    /// 
    ///     var exampleCheckBusinessHours = new AzureDevOps.CheckBusinessHours("exampleCheckBusinessHours", new()
    ///     {
    ///         ProjectId = exampleProject.Id,
    ///         DisplayName = "Managed by Terraform",
    ///         TargetResourceId = exampleQueue.Id,
    ///         TargetResourceType = "queue",
    ///         StartTime = "07:00",
    ///         EndTime = "15:30",
    ///         TimeZone = "UTC",
    ///         Monday = true,
    ///         Tuesday = true,
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ### Protect a repository
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
    ///     var exampleGit = new AzureDevOps.Git("exampleGit", new()
    ///     {
    ///         ProjectId = exampleProject.Id,
    ///         Initialization = new AzureDevOps.Inputs.GitInitializationArgs
    ///         {
    ///             InitType = "Clean",
    ///         },
    ///     });
    /// 
    ///     var exampleCheckBusinessHours = new AzureDevOps.CheckBusinessHours("exampleCheckBusinessHours", new()
    ///     {
    ///         ProjectId = exampleProject.Id,
    ///         DisplayName = "Managed by Terraform",
    ///         TargetResourceId = Output.Tuple(exampleProject.Id, exampleGit.Id).Apply(values =&gt;
    ///         {
    ///             var exampleProjectId = values.Item1;
    ///             var exampleGitId = values.Item2;
    ///             return $"{exampleProjectId}.{exampleGitId}";
    ///         }),
    ///         TargetResourceType = "repository",
    ///         StartTime = "07:00",
    ///         EndTime = "15:30",
    ///         TimeZone = "UTC",
    ///         Monday = true,
    ///         Tuesday = true,
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ### Protect a variable group
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
    ///     var exampleCheckBusinessHours = new AzureDevOps.CheckBusinessHours("exampleCheckBusinessHours", new()
    ///     {
    ///         ProjectId = exampleProject.Id,
    ///         DisplayName = "Managed by Terraform",
    ///         TargetResourceId = exampleVariableGroup.Id,
    ///         TargetResourceType = "variablegroup",
    ///         StartTime = "07:00",
    ///         EndTime = "15:30",
    ///         TimeZone = "UTC",
    ///         Monday = true,
    ///         Tuesday = true,
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Relevant Links
    /// 
    /// - [Define approvals and checks](https://learn.microsoft.com/en-us/azure/devops/pipelines/process/approvals?view=azure-devops&amp;tabs=check-pass)
    /// 
    /// ## Supported Time Zones
    /// 
    /// - AUS Central Standard Time
    /// - AUS Eastern Standard Time
    /// - Afghanistan Standard Time
    /// - Alaskan Standard Time
    /// - Aleutian Standard Time
    /// - Altai Standard Time
    /// - Arab Standard Time
    /// - Arabian Standard Time
    /// - Arabic Standard Time
    /// - Argentina Standard Time
    /// - Astrakhan Standard Time
    /// - Atlantic Standard Time
    /// - Aus Central W. Standard Time
    /// - Azerbaijan Standard Time
    /// - Azores Standard Time
    /// - Bahia Standard Time
    /// - Bangladesh Standard Time
    /// - Belarus Standard Time
    /// - Bougainville Standard Time
    /// - Canada Central Standard Time
    /// - Cape Verde Standard Time
    /// - Caucasus Standard Time
    /// - Cen. Australia Standard Time
    /// - Central America Standard Time
    /// - Central Asia Standard Time
    /// - Central Brazilian Standard Time
    /// - Central Europe Standard Time
    /// - Central European Standard Time
    /// - Central Pacific Standard Time
    /// - Central Standard Time (Mexico)
    /// - Central Standard Time
    /// - Chatham Islands Standard Time
    /// - China Standard Time
    /// - Cuba Standard Time
    /// - Dateline Standard Time
    /// - E. Africa Standard Time
    /// - E. Australia Standard Time
    /// - E. Europe Standard Time
    /// - E. South America Standard Time
    /// - Easter Island Standard Time
    /// - Eastern Standard Time (Mexico)
    /// - Eastern Standard Time
    /// - Egypt Standard Time
    /// - Ekaterinburg Standard Time
    /// - FLE Standard Time
    /// - Fiji Standard Time
    /// - GMT Standard Time
    /// - GTB Standard Time
    /// - Georgian Standard Time
    /// - Greenland Standard Time
    /// - Greenwich Standard Time
    /// - Haiti Standard Time
    /// - Hawaiian Standard Time
    /// - India Standard Time
    /// - Iran Standard Time
    /// - Israel Standard Time
    /// - Jordan Standard Time
    /// - Kaliningrad Standard Time
    /// - Kamchatka Standard Time
    /// - Korea Standard Time
    /// - Libya Standard Time
    /// - Line Islands Standard Time
    /// - Lord Howe Standard Time
    /// - Magadan Standard Time
    /// - Magallanes Standard Time
    /// - Marquesas Standard Time
    /// - Mauritius Standard Time
    /// - Mid-Atlantic Standard Time
    /// - Middle East Standard Time
    /// - Montevideo Standard Time
    /// - Morocco Standard Time
    /// - Mountain Standard Time (Mexico)
    /// - Mountain Standard Time
    /// - Myanmar Standard Time
    /// - N. Central Asia Standard Time
    /// - Namibia Standard Time
    /// - Nepal Standard Time
    /// - New Zealand Standard Time
    /// - Newfoundland Standard Time
    /// - Norfolk Standard Time
    /// - North Asia East Standard Time
    /// - North Asia Standard Time
    /// - North Korea Standard Time
    /// - Omsk Standard Time
    /// - Pacific SA Standard Time
    /// - Pacific Standard Time (Mexico)
    /// - Pacific Standard Time
    /// - Pakistan Standard Time
    /// - Paraguay Standard Time
    /// - Qyzylorda Standard Time
    /// - Romance Standard Time
    /// - Russia Time Zone 10
    /// - Russia Time Zone 11
    /// - Russia Time Zone 3
    /// - Russian Standard Time
    /// - SA Eastern Standard Time
    /// - SA Pacific Standard Time
    /// - SA Western Standard Time
    /// - SE Asia Standard Time
    /// - Saint Pierre Standard Time
    /// - Sakhalin Standard Time
    /// - Samoa Standard Time
    /// - Sao Tome Standard Time
    /// - Saratov Standard Time
    /// - Singapore Standard Time
    /// - South Africa Standard Time
    /// - South Sudan Standard Time
    /// - Sri Lanka Standard Time
    /// - Sudan Standard Time
    /// - Syria Standard Time
    /// - Taipei Standard Time
    /// - Tasmania Standard Time
    /// - Tocantins Standard Time
    /// - Tokyo Standard Time
    /// - Tomsk Standard Time
    /// - Tonga Standard Time
    /// - Transbaikal Standard Time
    /// - Turkey Standard Time
    /// - Turks And Caicos Standard Time
    /// - US Eastern Standard Time
    /// - US Mountain Standard Time
    /// - UTC
    /// - UTC+12
    /// - UTC+13
    /// - UTC-02
    /// - UTC-08
    /// - UTC-09
    /// - UTC-11
    /// - Ulaanbaatar Standard Time
    /// - Venezuela Standard Time
    /// - Vladivostok Standard Time
    /// - Volgograd Standard Time
    /// - W. Australia Standard Time
    /// - W. Central Africa Standard Time
    /// - W. Europe Standard Time
    /// - W. Mongolia Standard Time
    /// - West Asia Standard Time
    /// - West Bank Standard Time
    /// - West Pacific Standard Time
    /// - Yakutsk Standard Time
    /// - Yukon Standard Time
    /// 
    /// ## Import
    /// 
    /// Importing this resource is not supported.
    /// </summary>
    [AzureDevOpsResourceType("azuredevops:index/checkBusinessHours:CheckBusinessHours")]
    public partial class CheckBusinessHours : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the business hours check displayed in the web UI.
        /// </summary>
        [Output("displayName")]
        public Output<string?> DisplayName { get; private set; } = null!;

        /// <summary>
        /// The end of the time period that this check will be allowed to pass, specified as 24-hour time with leading zeros.
        /// </summary>
        [Output("endTime")]
        public Output<string> EndTime { get; private set; } = null!;

        /// <summary>
        /// This check will pass on Fridays. Defaults to `false`.
        /// </summary>
        [Output("friday")]
        public Output<bool?> Friday { get; private set; } = null!;

        /// <summary>
        /// This check will pass on Mondays. Defaults to `false`.
        /// </summary>
        [Output("monday")]
        public Output<bool?> Monday { get; private set; } = null!;

        /// <summary>
        /// The project ID.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// This check will pass on Saturdays. Defaults to `false`.
        /// </summary>
        [Output("saturday")]
        public Output<bool?> Saturday { get; private set; } = null!;

        /// <summary>
        /// The beginning of the time period that this check will be allowed to pass, specified as 24-hour time with leading zeros.
        /// </summary>
        [Output("startTime")]
        public Output<string> StartTime { get; private set; } = null!;

        /// <summary>
        /// This check will pass on Sundays. Defaults to `false`.
        /// </summary>
        [Output("sunday")]
        public Output<bool?> Sunday { get; private set; } = null!;

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
        /// This check will pass on Thursdays. Defaults to `false`.
        /// </summary>
        [Output("thursday")]
        public Output<bool?> Thursday { get; private set; } = null!;

        /// <summary>
        /// The time zone this check will be evaluated in. See below for supported values.
        /// </summary>
        [Output("timeZone")]
        public Output<string> TimeZone { get; private set; } = null!;

        /// <summary>
        /// The timeout in minutes for the business hours check. Defaults to `1440`.
        /// </summary>
        [Output("timeout")]
        public Output<int?> Timeout { get; private set; } = null!;

        /// <summary>
        /// This check will pass on Tuesday. Defaults to `false`.
        /// </summary>
        [Output("tuesday")]
        public Output<bool?> Tuesday { get; private set; } = null!;

        /// <summary>
        /// The version of the check.
        /// </summary>
        [Output("version")]
        public Output<int> Version { get; private set; } = null!;

        /// <summary>
        /// This check will pass on Wednesdays. Defaults to `false`.
        /// </summary>
        [Output("wednesday")]
        public Output<bool?> Wednesday { get; private set; } = null!;


        /// <summary>
        /// Create a CheckBusinessHours resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CheckBusinessHours(string name, CheckBusinessHoursArgs args, CustomResourceOptions? options = null)
            : base("azuredevops:index/checkBusinessHours:CheckBusinessHours", name, args ?? new CheckBusinessHoursArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CheckBusinessHours(string name, Input<string> id, CheckBusinessHoursState? state = null, CustomResourceOptions? options = null)
            : base("azuredevops:index/checkBusinessHours:CheckBusinessHours", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing CheckBusinessHours resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CheckBusinessHours Get(string name, Input<string> id, CheckBusinessHoursState? state = null, CustomResourceOptions? options = null)
        {
            return new CheckBusinessHours(name, id, state, options);
        }
    }

    public sealed class CheckBusinessHoursArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the business hours check displayed in the web UI.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// The end of the time period that this check will be allowed to pass, specified as 24-hour time with leading zeros.
        /// </summary>
        [Input("endTime", required: true)]
        public Input<string> EndTime { get; set; } = null!;

        /// <summary>
        /// This check will pass on Fridays. Defaults to `false`.
        /// </summary>
        [Input("friday")]
        public Input<bool>? Friday { get; set; }

        /// <summary>
        /// This check will pass on Mondays. Defaults to `false`.
        /// </summary>
        [Input("monday")]
        public Input<bool>? Monday { get; set; }

        /// <summary>
        /// The project ID.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        /// <summary>
        /// This check will pass on Saturdays. Defaults to `false`.
        /// </summary>
        [Input("saturday")]
        public Input<bool>? Saturday { get; set; }

        /// <summary>
        /// The beginning of the time period that this check will be allowed to pass, specified as 24-hour time with leading zeros.
        /// </summary>
        [Input("startTime", required: true)]
        public Input<string> StartTime { get; set; } = null!;

        /// <summary>
        /// This check will pass on Sundays. Defaults to `false`.
        /// </summary>
        [Input("sunday")]
        public Input<bool>? Sunday { get; set; }

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
        /// This check will pass on Thursdays. Defaults to `false`.
        /// </summary>
        [Input("thursday")]
        public Input<bool>? Thursday { get; set; }

        /// <summary>
        /// The time zone this check will be evaluated in. See below for supported values.
        /// </summary>
        [Input("timeZone", required: true)]
        public Input<string> TimeZone { get; set; } = null!;

        /// <summary>
        /// The timeout in minutes for the business hours check. Defaults to `1440`.
        /// </summary>
        [Input("timeout")]
        public Input<int>? Timeout { get; set; }

        /// <summary>
        /// This check will pass on Tuesday. Defaults to `false`.
        /// </summary>
        [Input("tuesday")]
        public Input<bool>? Tuesday { get; set; }

        /// <summary>
        /// This check will pass on Wednesdays. Defaults to `false`.
        /// </summary>
        [Input("wednesday")]
        public Input<bool>? Wednesday { get; set; }

        public CheckBusinessHoursArgs()
        {
        }
        public static new CheckBusinessHoursArgs Empty => new CheckBusinessHoursArgs();
    }

    public sealed class CheckBusinessHoursState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the business hours check displayed in the web UI.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// The end of the time period that this check will be allowed to pass, specified as 24-hour time with leading zeros.
        /// </summary>
        [Input("endTime")]
        public Input<string>? EndTime { get; set; }

        /// <summary>
        /// This check will pass on Fridays. Defaults to `false`.
        /// </summary>
        [Input("friday")]
        public Input<bool>? Friday { get; set; }

        /// <summary>
        /// This check will pass on Mondays. Defaults to `false`.
        /// </summary>
        [Input("monday")]
        public Input<bool>? Monday { get; set; }

        /// <summary>
        /// The project ID.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// This check will pass on Saturdays. Defaults to `false`.
        /// </summary>
        [Input("saturday")]
        public Input<bool>? Saturday { get; set; }

        /// <summary>
        /// The beginning of the time period that this check will be allowed to pass, specified as 24-hour time with leading zeros.
        /// </summary>
        [Input("startTime")]
        public Input<string>? StartTime { get; set; }

        /// <summary>
        /// This check will pass on Sundays. Defaults to `false`.
        /// </summary>
        [Input("sunday")]
        public Input<bool>? Sunday { get; set; }

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
        /// This check will pass on Thursdays. Defaults to `false`.
        /// </summary>
        [Input("thursday")]
        public Input<bool>? Thursday { get; set; }

        /// <summary>
        /// The time zone this check will be evaluated in. See below for supported values.
        /// </summary>
        [Input("timeZone")]
        public Input<string>? TimeZone { get; set; }

        /// <summary>
        /// The timeout in minutes for the business hours check. Defaults to `1440`.
        /// </summary>
        [Input("timeout")]
        public Input<int>? Timeout { get; set; }

        /// <summary>
        /// This check will pass on Tuesday. Defaults to `false`.
        /// </summary>
        [Input("tuesday")]
        public Input<bool>? Tuesday { get; set; }

        /// <summary>
        /// The version of the check.
        /// </summary>
        [Input("version")]
        public Input<int>? Version { get; set; }

        /// <summary>
        /// This check will pass on Wednesdays. Defaults to `false`.
        /// </summary>
        [Input("wednesday")]
        public Input<bool>? Wednesday { get; set; }

        public CheckBusinessHoursState()
        {
        }
        public static new CheckBusinessHoursState Empty => new CheckBusinessHoursState();
    }
}
