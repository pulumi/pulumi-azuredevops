// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manages members of a team within a project in a Azure DevOps organization.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const exampleProject = new azuredevops.Project("exampleProject", {
 *     workItemTemplate: "Agile",
 *     versionControl: "Git",
 *     visibility: "private",
 *     description: "Managed by Terraform",
 * });
 * const example-project-readers = azuredevops.getGroupOutput({
 *     projectId: exampleProject.id,
 *     name: "Readers",
 * });
 * const exampleTeam = new azuredevops.Team("exampleTeam", {projectId: exampleProject.id});
 * const example_team_members = new azuredevops.TeamMembers("example-team-members", {
 *     projectId: exampleTeam.projectId,
 *     teamId: exampleTeam.id,
 *     mode: "overwrite",
 *     members: [example_project_readers.apply(example_project_readers => example_project_readers.descriptor)],
 * });
 * ```
 * ## Relevant Links
 *
 * - [Azure DevOps Service REST API 7.0 - Teams - Update](https://docs.microsoft.com/en-us/rest/api/azure/devops/core/teams/update?view=azure-devops-rest-7.0)
 *
 * ## PAT Permissions Required
 *
 * - **vso.project_write**:	Grants the ability to read and update projects and teams.
 *
 * ## Import
 *
 * The resource does not support import.
 */
export class TeamMembers extends pulumi.CustomResource {
    /**
     * Get an existing TeamMembers resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TeamMembersState, opts?: pulumi.CustomResourceOptions): TeamMembers {
        return new TeamMembers(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuredevops:index/teamMembers:TeamMembers';

    /**
     * Returns true if the given object is an instance of TeamMembers.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TeamMembers {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TeamMembers.__pulumiType;
    }

    /**
     * List of subject descriptors to define members of the team.NOTE: It's possible to define team members both within the
     * `azuredevops.Team` resource via the `members` block and by using the
     * `azuredevops.TeamMembers` resource. However it's not possible to use
     * both methods to manage team members, since there'll be conflicts.
     */
    public readonly members!: pulumi.Output<string[]>;
    /**
     * The mode how the resource manages team members.
     */
    public readonly mode!: pulumi.Output<string | undefined>;
    /**
     * The Project ID.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * The ID of the Team.
     */
    public readonly teamId!: pulumi.Output<string>;

    /**
     * Create a TeamMembers resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TeamMembersArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TeamMembersArgs | TeamMembersState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TeamMembersState | undefined;
            resourceInputs["members"] = state ? state.members : undefined;
            resourceInputs["mode"] = state ? state.mode : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["teamId"] = state ? state.teamId : undefined;
        } else {
            const args = argsOrState as TeamMembersArgs | undefined;
            if ((!args || args.members === undefined) && !opts.urn) {
                throw new Error("Missing required property 'members'");
            }
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            if ((!args || args.teamId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'teamId'");
            }
            resourceInputs["members"] = args ? args.members : undefined;
            resourceInputs["mode"] = args ? args.mode : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["teamId"] = args ? args.teamId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(TeamMembers.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TeamMembers resources.
 */
export interface TeamMembersState {
    /**
     * List of subject descriptors to define members of the team.NOTE: It's possible to define team members both within the
     * `azuredevops.Team` resource via the `members` block and by using the
     * `azuredevops.TeamMembers` resource. However it's not possible to use
     * both methods to manage team members, since there'll be conflicts.
     */
    members?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The mode how the resource manages team members.
     */
    mode?: pulumi.Input<string>;
    /**
     * The Project ID.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The ID of the Team.
     */
    teamId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a TeamMembers resource.
 */
export interface TeamMembersArgs {
    /**
     * List of subject descriptors to define members of the team.NOTE: It's possible to define team members both within the
     * `azuredevops.Team` resource via the `members` block and by using the
     * `azuredevops.TeamMembers` resource. However it's not possible to use
     * both methods to manage team members, since there'll be conflicts.
     */
    members: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The mode how the resource manages team members.
     */
    mode?: pulumi.Input<string>;
    /**
     * The Project ID.
     */
    projectId: pulumi.Input<string>;
    /**
     * The ID of the Team.
     */
    teamId: pulumi.Input<string>;
}
