// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops;

import com.pulumi.azuredevops.TeamMembersArgs;
import com.pulumi.azuredevops.Utilities;
import com.pulumi.azuredevops.inputs.TeamMembersState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages members of a team within a project in a Azure DevOps organization.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.azuredevops.Project;
 * import com.pulumi.azuredevops.ProjectArgs;
 * import com.pulumi.azuredevops.AzuredevopsFunctions;
 * import com.pulumi.azuredevops.inputs.GetGroupArgs;
 * import com.pulumi.azuredevops.Team;
 * import com.pulumi.azuredevops.TeamArgs;
 * import com.pulumi.azuredevops.TeamMembers;
 * import com.pulumi.azuredevops.TeamMembersArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var exampleProject = new Project(&#34;exampleProject&#34;, ProjectArgs.builder()        
 *             .workItemTemplate(&#34;Agile&#34;)
 *             .versionControl(&#34;Git&#34;)
 *             .visibility(&#34;private&#34;)
 *             .description(&#34;Managed by Terraform&#34;)
 *             .build());
 * 
 *         final var example-project-readers = AzuredevopsFunctions.getGroup(GetGroupArgs.builder()
 *             .projectId(exampleProject.id())
 *             .name(&#34;Readers&#34;)
 *             .build());
 * 
 *         var exampleTeam = new Team(&#34;exampleTeam&#34;, TeamArgs.builder()        
 *             .projectId(exampleProject.id())
 *             .build());
 * 
 *         var example_team_members = new TeamMembers(&#34;example-team-members&#34;, TeamMembersArgs.builder()        
 *             .projectId(exampleTeam.projectId())
 *             .teamId(exampleTeam.id())
 *             .mode(&#34;overwrite&#34;)
 *             .members(example_project_readers.applyValue(example_project_readers -&gt; example_project_readers.descriptor()))
 *             .build());
 * 
 *     }
 * }
 * ```
 * ## Relevant Links
 * 
 * - [Azure DevOps Service REST API 6.0 - Teams - Update](https://docs.microsoft.com/en-us/rest/api/azure/devops/core/teams/update?view=azure-devops-rest-6.0)
 * 
 * ## PAT Permissions Required
 * 
 * - **vso.project_write**:	Grants the ability to read and update projects and teams.
 * 
 * ## Import
 * 
 * The resource does not support import.
 * 
 */
@ResourceType(type="azuredevops:index/teamMembers:TeamMembers")
public class TeamMembers extends com.pulumi.resources.CustomResource {
    /**
     * List of subject descriptors to define members of the team.
     * 
     */
    @Export(name="members", type=List.class, parameters={String.class})
    private Output<List<String>> members;

    /**
     * @return List of subject descriptors to define members of the team.
     * 
     */
    public Output<List<String>> members() {
        return this.members;
    }
    /**
     * The mode how the resource manages team members.
     * - `mode == add`: the resource will ensure that all specified members will be part of the referenced team
     * - `mode == overwrite`: the resource will replace all existing members with the members specified within the `members` block
     * 
     */
    @Export(name="mode", type=String.class, parameters={})
    private Output</* @Nullable */ String> mode;

    /**
     * @return The mode how the resource manages team members.
     * - `mode == add`: the resource will ensure that all specified members will be part of the referenced team
     * - `mode == overwrite`: the resource will replace all existing members with the members specified within the `members` block
     * 
     */
    public Output<Optional<String>> mode() {
        return Codegen.optional(this.mode);
    }
    /**
     * The Project ID.
     * 
     */
    @Export(name="projectId", type=String.class, parameters={})
    private Output<String> projectId;

    /**
     * @return The Project ID.
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }
    /**
     * The ID of the Team.
     * 
     */
    @Export(name="teamId", type=String.class, parameters={})
    private Output<String> teamId;

    /**
     * @return The ID of the Team.
     * 
     */
    public Output<String> teamId() {
        return this.teamId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public TeamMembers(String name) {
        this(name, TeamMembersArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public TeamMembers(String name, TeamMembersArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public TeamMembers(String name, TeamMembersArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/teamMembers:TeamMembers", name, args == null ? TeamMembersArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private TeamMembers(String name, Output<String> id, @Nullable TeamMembersState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/teamMembers:TeamMembers", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static TeamMembers get(String name, Output<String> id, @Nullable TeamMembersState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new TeamMembers(name, id, state, options);
    }
}
