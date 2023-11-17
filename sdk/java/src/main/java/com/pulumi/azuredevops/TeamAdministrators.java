// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops;

import com.pulumi.azuredevops.TeamAdministratorsArgs;
import com.pulumi.azuredevops.Utilities;
import com.pulumi.azuredevops.inputs.TeamAdministratorsState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages administrators of a team within a project in a Azure DevOps organization.
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
 * import com.pulumi.azuredevops.TeamAdministrators;
 * import com.pulumi.azuredevops.TeamAdministratorsArgs;
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
 *         final var example-project-contributors = AzuredevopsFunctions.getGroup(GetGroupArgs.builder()
 *             .projectId(exampleProject.id())
 *             .name(&#34;Contributors&#34;)
 *             .build());
 * 
 *         var exampleTeam = new Team(&#34;exampleTeam&#34;, TeamArgs.builder()        
 *             .projectId(exampleProject.id())
 *             .build());
 * 
 *         var example_team_administrators = new TeamAdministrators(&#34;example-team-administrators&#34;, TeamAdministratorsArgs.builder()        
 *             .projectId(exampleTeam.projectId())
 *             .teamId(exampleTeam.id())
 *             .mode(&#34;overwrite&#34;)
 *             .administrators(example_project_contributors.applyValue(example_project_contributors -&gt; example_project_contributors.descriptor()))
 *             .build());
 * 
 *     }
 * }
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
 * 
 */
@ResourceType(type="azuredevops:index/teamAdministrators:TeamAdministrators")
public class TeamAdministrators extends com.pulumi.resources.CustomResource {
    /**
     * List of subject descriptors to define adminitrators of the team.NOTE: It&#39;s possible to define team administrators both within the
     * `azuredevops.Team` resource via the `administrators` block and by using the
     * `azuredevops.TeamAdministrators` resource. However it&#39;s not possible to use
     * both methods to manage team administrators, since there&#39;ll be conflicts.
     * 
     */
    @Export(name="administrators", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> administrators;

    /**
     * @return List of subject descriptors to define adminitrators of the team.NOTE: It&#39;s possible to define team administrators both within the
     * `azuredevops.Team` resource via the `administrators` block and by using the
     * `azuredevops.TeamAdministrators` resource. However it&#39;s not possible to use
     * both methods to manage team administrators, since there&#39;ll be conflicts.
     * 
     */
    public Output<List<String>> administrators() {
        return this.administrators;
    }
    /**
     * The mode how the resource manages team administrators.
     * 
     */
    @Export(name="mode", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> mode;

    /**
     * @return The mode how the resource manages team administrators.
     * 
     */
    public Output<Optional<String>> mode() {
        return Codegen.optional(this.mode);
    }
    /**
     * The Project ID.
     * 
     */
    @Export(name="projectId", refs={String.class}, tree="[0]")
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
    @Export(name="teamId", refs={String.class}, tree="[0]")
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
    public TeamAdministrators(String name) {
        this(name, TeamAdministratorsArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public TeamAdministrators(String name, TeamAdministratorsArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public TeamAdministrators(String name, TeamAdministratorsArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/teamAdministrators:TeamAdministrators", name, args == null ? TeamAdministratorsArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private TeamAdministrators(String name, Output<String> id, @Nullable TeamAdministratorsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/teamAdministrators:TeamAdministrators", name, state, makeResourceOptions(options, id));
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
    public static TeamAdministrators get(String name, Output<String> id, @Nullable TeamAdministratorsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new TeamAdministrators(name, id, state, options);
    }
}
