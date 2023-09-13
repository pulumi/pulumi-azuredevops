// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops;

import com.pulumi.azuredevops.TeamArgs;
import com.pulumi.azuredevops.Utilities;
import com.pulumi.azuredevops.inputs.TeamState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a team within a project in a Azure DevOps organization.
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
 *         final var example-project-readers = AzuredevopsFunctions.getGroup(GetGroupArgs.builder()
 *             .projectId(exampleProject.id())
 *             .name(&#34;Readers&#34;)
 *             .build());
 * 
 *         var exampleTeam = new Team(&#34;exampleTeam&#34;, TeamArgs.builder()        
 *             .projectId(exampleProject.id())
 *             .administrators(example_project_contributors.applyValue(example_project_contributors -&gt; example_project_contributors.descriptor()))
 *             .members(example_project_readers.applyValue(example_project_readers -&gt; example_project_readers.descriptor()))
 *             .build());
 * 
 *     }
 * }
 * ```
 * ## Relevant Links
 * 
 * - [Azure DevOps Service REST API 7.0 - Teams - Create](https://docs.microsoft.com/en-us/rest/api/azure/devops/core/teams/create?view=azure-devops-rest-7.0)
 * 
 * ## PAT Permissions Required
 * 
 * - **vso.project_manage**:	Grants the ability to create, read, update, and delete projects and teams.
 * 
 * ## Import
 * 
 * Azure DevOps teams can be imported using the complete resource id `&lt;project_id&gt;/&lt;team_id&gt;` e.g.
 * 
 * ```sh
 *  $ pulumi import azuredevops:index/team:Team example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
 * ```
 * 
 */
@ResourceType(type="azuredevops:index/team:Team")
public class Team extends com.pulumi.resources.CustomResource {
    /**
     * List of subject descriptors to define administrators of the team.
     * 
     * &gt; NOTE: It&#39;s possible to define team administrators both within the
     * `azuredevops.Team` resource via the `administrators` block and by using the
     * `azuredevops.TeamAdministrators` resource. However it&#39;s not possible to use
     * both methods to manage team administrators, since there&#39;ll be conflicts.
     * 
     */
    @Export(name="administrators", type=List.class, parameters={String.class})
    private Output<List<String>> administrators;

    /**
     * @return List of subject descriptors to define administrators of the team.
     * 
     * &gt; NOTE: It&#39;s possible to define team administrators both within the
     * `azuredevops.Team` resource via the `administrators` block and by using the
     * `azuredevops.TeamAdministrators` resource. However it&#39;s not possible to use
     * both methods to manage team administrators, since there&#39;ll be conflicts.
     * 
     */
    public Output<List<String>> administrators() {
        return this.administrators;
    }
    /**
     * The description of the Team.
     * 
     */
    @Export(name="description", type=String.class, parameters={})
    private Output</* @Nullable */ String> description;

    /**
     * @return The description of the Team.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The descriptor of the Team.
     * 
     */
    @Export(name="descriptor", type=String.class, parameters={})
    private Output<String> descriptor;

    /**
     * @return The descriptor of the Team.
     * 
     */
    public Output<String> descriptor() {
        return this.descriptor;
    }
    /**
     * List of subject descriptors to define members of the team.
     * 
     * &gt; NOTE: It&#39;s possible to define team members both within the
     * `azuredevops.Team` resource via the `members` block and by using the
     * `azuredevops.TeamMembers` resource. However it&#39;s not possible to use
     * both methods to manage team members, since there&#39;ll be conflicts.
     * 
     */
    @Export(name="members", type=List.class, parameters={String.class})
    private Output<List<String>> members;

    /**
     * @return List of subject descriptors to define members of the team.
     * 
     * &gt; NOTE: It&#39;s possible to define team members both within the
     * `azuredevops.Team` resource via the `members` block and by using the
     * `azuredevops.TeamMembers` resource. However it&#39;s not possible to use
     * both methods to manage team members, since there&#39;ll be conflicts.
     * 
     */
    public Output<List<String>> members() {
        return this.members;
    }
    /**
     * The name of the Team.
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return The name of the Team.
     * 
     */
    public Output<String> name() {
        return this.name;
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
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Team(String name) {
        this(name, TeamArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Team(String name, TeamArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Team(String name, TeamArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/team:Team", name, args == null ? TeamArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Team(String name, Output<String> id, @Nullable TeamState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/team:Team", name, state, makeResourceOptions(options, id));
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
    public static Team get(String name, Output<String> id, @Nullable TeamState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Team(name, id, state, options);
    }
}
