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
    /// Manages Dashboard within Azure DevOps project.
    /// 
    /// &gt; **NOTE:** Project level Dashboard allows to be created with the same name. Dashboard held by a team must have a different name.
    /// 
    /// ## Example Usage
    /// 
    /// ### Manage Project dashboard
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
    ///         Description = "Managed by Pulumi",
    ///     });
    /// 
    ///     var exampleDashboard = new AzureDevOps.Dashboard("example", new()
    ///     {
    ///         ProjectId = example.Id,
    ///         Name = "Example dashboard",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ### Manage Team dashboard
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
    ///         Description = "Managed by Pulumi",
    ///     });
    /// 
    ///     var exampleTeam = new AzureDevOps.Team("example", new()
    ///     {
    ///         ProjectId = example.Id,
    ///         Name = "Example team",
    ///     });
    /// 
    ///     var exampleDashboard = new AzureDevOps.Dashboard("example", new()
    ///     {
    ///         ProjectId = example.Id,
    ///         Name = "Example dashboard",
    ///         TeamId = exampleTeam.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Relevant Links
    /// 
    /// - [Azure DevOps dashboards REST API 7.1 - Dashboard ](https://learn.microsoft.com/en-us/rest/api/azure/devops/dashboard/dashboards?view=azure-devops-rest-7.1)
    /// 
    /// ## Import
    /// 
    /// Azure DevOps Dashboard can be imported using the `projectId/dasboardId` or `projectId/teamId/dasboardId`
    /// 
    /// ```sh
    /// $ pulumi import azuredevops:index/dashboard:Dashboard dashboard 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
    /// ```
    /// 
    /// or
    /// 
    /// ```sh
    /// $ pulumi import azuredevops:index/dashboard:Dashboard dashboard 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
    /// ```
    /// </summary>
    [AzureDevOpsResourceType("azuredevops:index/dashboard:Dashboard")]
    public partial class Dashboard : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The description of the dashboard.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The name of the Dashboard.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The owner of the Dashboard, could be the project or a team.
        /// </summary>
        [Output("ownerId")]
        public Output<string> OwnerId { get; private set; } = null!;

        /// <summary>
        /// The ID of the Project. Changing this forces a new resource to be created.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// The interval for client to automatically refresh the dashboard. Expressed in minutes. Possible values are: `0`, `5`.Defaults to `0`.
        /// </summary>
        [Output("refreshInterval")]
        public Output<int?> RefreshInterval { get; private set; } = null!;

        /// <summary>
        /// The ID of the Team.
        /// </summary>
        [Output("teamId")]
        public Output<string?> TeamId { get; private set; } = null!;


        /// <summary>
        /// Create a Dashboard resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Dashboard(string name, DashboardArgs args, CustomResourceOptions? options = null)
            : base("azuredevops:index/dashboard:Dashboard", name, args ?? new DashboardArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Dashboard(string name, Input<string> id, DashboardState? state = null, CustomResourceOptions? options = null)
            : base("azuredevops:index/dashboard:Dashboard", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Dashboard resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Dashboard Get(string name, Input<string> id, DashboardState? state = null, CustomResourceOptions? options = null)
        {
            return new Dashboard(name, id, state, options);
        }
    }

    public sealed class DashboardArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the dashboard.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the Dashboard.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the Project. Changing this forces a new resource to be created.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        /// <summary>
        /// The interval for client to automatically refresh the dashboard. Expressed in minutes. Possible values are: `0`, `5`.Defaults to `0`.
        /// </summary>
        [Input("refreshInterval")]
        public Input<int>? RefreshInterval { get; set; }

        /// <summary>
        /// The ID of the Team.
        /// </summary>
        [Input("teamId")]
        public Input<string>? TeamId { get; set; }

        public DashboardArgs()
        {
        }
        public static new DashboardArgs Empty => new DashboardArgs();
    }

    public sealed class DashboardState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the dashboard.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the Dashboard.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The owner of the Dashboard, could be the project or a team.
        /// </summary>
        [Input("ownerId")]
        public Input<string>? OwnerId { get; set; }

        /// <summary>
        /// The ID of the Project. Changing this forces a new resource to be created.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The interval for client to automatically refresh the dashboard. Expressed in minutes. Possible values are: `0`, `5`.Defaults to `0`.
        /// </summary>
        [Input("refreshInterval")]
        public Input<int>? RefreshInterval { get; set; }

        /// <summary>
        /// The ID of the Team.
        /// </summary>
        [Input("teamId")]
        public Input<string>? TeamId { get; set; }

        public DashboardState()
        {
        }
        public static new DashboardState Empty => new DashboardState();
    }
}
