// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops;

import com.pulumi.azuredevops.BuildDefinitionArgs;
import com.pulumi.azuredevops.Utilities;
import com.pulumi.azuredevops.inputs.BuildDefinitionState;
import com.pulumi.azuredevops.outputs.BuildDefinitionCiTrigger;
import com.pulumi.azuredevops.outputs.BuildDefinitionFeature;
import com.pulumi.azuredevops.outputs.BuildDefinitionPullRequestTrigger;
import com.pulumi.azuredevops.outputs.BuildDefinitionRepository;
import com.pulumi.azuredevops.outputs.BuildDefinitionSchedule;
import com.pulumi.azuredevops.outputs.BuildDefinitionVariable;
import com.pulumi.core.Alias;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a Build Definition within Azure DevOps.
 * 
 * ## Example Usage
 * 
 * ### Tfs
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.azuredevops.Project;
 * import com.pulumi.azuredevops.ProjectArgs;
 * import com.pulumi.azuredevops.Git;
 * import com.pulumi.azuredevops.GitArgs;
 * import com.pulumi.azuredevops.inputs.GitInitializationArgs;
 * import com.pulumi.azuredevops.VariableGroup;
 * import com.pulumi.azuredevops.VariableGroupArgs;
 * import com.pulumi.azuredevops.inputs.VariableGroupVariableArgs;
 * import com.pulumi.azuredevops.BuildDefinition;
 * import com.pulumi.azuredevops.BuildDefinitionArgs;
 * import com.pulumi.azuredevops.inputs.BuildDefinitionCiTriggerArgs;
 * import com.pulumi.azuredevops.inputs.BuildDefinitionScheduleArgs;
 * import com.pulumi.azuredevops.inputs.BuildDefinitionRepositoryArgs;
 * import com.pulumi.azuredevops.inputs.BuildDefinitionVariableArgs;
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
 *             .visibility(&#34;private&#34;)
 *             .versionControl(&#34;Git&#34;)
 *             .workItemTemplate(&#34;Agile&#34;)
 *             .build());
 * 
 *         var exampleGit = new Git(&#34;exampleGit&#34;, GitArgs.builder()        
 *             .projectId(exampleProject.id())
 *             .initialization(GitInitializationArgs.builder()
 *                 .initType(&#34;Clean&#34;)
 *                 .build())
 *             .build());
 * 
 *         var exampleVariableGroup = new VariableGroup(&#34;exampleVariableGroup&#34;, VariableGroupArgs.builder()        
 *             .projectId(exampleProject.id())
 *             .description(&#34;Managed by Terraform&#34;)
 *             .allowAccess(true)
 *             .variables(VariableGroupVariableArgs.builder()
 *                 .name(&#34;FOO&#34;)
 *                 .value(&#34;BAR&#34;)
 *                 .build())
 *             .build());
 * 
 *         var exampleBuildDefinition = new BuildDefinition(&#34;exampleBuildDefinition&#34;, BuildDefinitionArgs.builder()        
 *             .projectId(exampleProject.id())
 *             .path(&#34;\\ExampleFolder&#34;)
 *             .ciTrigger(BuildDefinitionCiTriggerArgs.builder()
 *                 .useYaml(false)
 *                 .build())
 *             .schedules(BuildDefinitionScheduleArgs.builder()
 *                 .branchFilters(BuildDefinitionScheduleBranchFilterArgs.builder()
 *                     .includes(&#34;master&#34;)
 *                     .excludes(                    
 *                         &#34;test&#34;,
 *                         &#34;regression&#34;)
 *                     .build())
 *                 .daysToBuilds(                
 *                     &#34;Wed&#34;,
 *                     &#34;Sun&#34;)
 *                 .scheduleOnlyWithChanges(true)
 *                 .startHours(10)
 *                 .startMinutes(59)
 *                 .timeZone(&#34;(UTC) Coordinated Universal Time&#34;)
 *                 .build())
 *             .repository(BuildDefinitionRepositoryArgs.builder()
 *                 .repoType(&#34;TfsGit&#34;)
 *                 .repoId(exampleGit.id())
 *                 .branchName(exampleGit.defaultBranch())
 *                 .ymlPath(&#34;azure-pipelines.yml&#34;)
 *                 .build())
 *             .variableGroups(exampleVariableGroup.id())
 *             .variables(            
 *                 BuildDefinitionVariableArgs.builder()
 *                     .name(&#34;PipelineVariable&#34;)
 *                     .value(&#34;Go Microsoft!&#34;)
 *                     .build(),
 *                 BuildDefinitionVariableArgs.builder()
 *                     .name(&#34;PipelineSecret&#34;)
 *                     .secretValue(&#34;ZGV2cw&#34;)
 *                     .isSecret(true)
 *                     .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ### GitHub Enterprise
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.azuredevops.Project;
 * import com.pulumi.azuredevops.ProjectArgs;
 * import com.pulumi.azuredevops.ServiceEndpointGitHubEnterprise;
 * import com.pulumi.azuredevops.ServiceEndpointGitHubEnterpriseArgs;
 * import com.pulumi.azuredevops.inputs.ServiceEndpointGitHubEnterpriseAuthPersonalArgs;
 * import com.pulumi.azuredevops.BuildDefinition;
 * import com.pulumi.azuredevops.BuildDefinitionArgs;
 * import com.pulumi.azuredevops.inputs.BuildDefinitionCiTriggerArgs;
 * import com.pulumi.azuredevops.inputs.BuildDefinitionRepositoryArgs;
 * import com.pulumi.azuredevops.inputs.BuildDefinitionScheduleArgs;
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
 *             .visibility(&#34;private&#34;)
 *             .versionControl(&#34;Git&#34;)
 *             .workItemTemplate(&#34;Agile&#34;)
 *             .build());
 * 
 *         var exampleServiceEndpointGitHubEnterprise = new ServiceEndpointGitHubEnterprise(&#34;exampleServiceEndpointGitHubEnterprise&#34;, ServiceEndpointGitHubEnterpriseArgs.builder()        
 *             .projectId(exampleProject.id())
 *             .serviceEndpointName(&#34;Example GitHub Enterprise&#34;)
 *             .url(&#34;https://github.contoso.com&#34;)
 *             .description(&#34;Managed by Terraform&#34;)
 *             .authPersonal(ServiceEndpointGitHubEnterpriseAuthPersonalArgs.builder()
 *                 .personalAccessToken(&#34;xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx&#34;)
 *                 .build())
 *             .build());
 * 
 *         var exampleBuildDefinition = new BuildDefinition(&#34;exampleBuildDefinition&#34;, BuildDefinitionArgs.builder()        
 *             .projectId(exampleProject.id())
 *             .path(&#34;\\ExampleFolder&#34;)
 *             .ciTrigger(BuildDefinitionCiTriggerArgs.builder()
 *                 .useYaml(false)
 *                 .build())
 *             .repository(BuildDefinitionRepositoryArgs.builder()
 *                 .repoType(&#34;GitHubEnterprise&#34;)
 *                 .repoId(&#34;&lt;GitHub Org&gt;/&lt;Repo Name&gt;&#34;)
 *                 .githubEnterpriseUrl(&#34;https://github.company.com&#34;)
 *                 .branchName(&#34;master&#34;)
 *                 .ymlPath(&#34;azure-pipelines.yml&#34;)
 *                 .serviceConnectionId(exampleServiceEndpointGitHubEnterprise.id())
 *                 .build())
 *             .schedules(BuildDefinitionScheduleArgs.builder()
 *                 .branchFilters(BuildDefinitionScheduleBranchFilterArgs.builder()
 *                     .includes(&#34;main&#34;)
 *                     .excludes(                    
 *                         &#34;test&#34;,
 *                         &#34;regression&#34;)
 *                     .build())
 *                 .daysToBuilds(                
 *                     &#34;Wed&#34;,
 *                     &#34;Sun&#34;)
 *                 .scheduleOnlyWithChanges(true)
 *                 .startHours(10)
 *                 .startMinutes(59)
 *                 .timeZone(&#34;(UTC) Coordinated Universal Time&#34;)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Remarks
 * 
 * The path attribute can not end in `\` unless the path is the root value of `\`.
 * 
 * Valid path values (yaml encoded) include:
 * - `\\`
 * - `\\ExampleFolder`
 * - `\\Nested\\Example Folder`
 * 
 * The value of `\\ExampleFolder\\` would be invalid.
 * 
 * ## Relevant Links
 * 
 * - [Azure DevOps Service REST API 7.0 - Build Definitions](https://docs.microsoft.com/en-us/rest/api/azure/devops/build/definitions?view=azure-devops-rest-7.0)
 * 
 * ## Import
 * 
 * Azure DevOps Build Definitions can be imported using the project name/definitions Id or by the project Guid/definitions Id, e.g.
 * 
 * ```sh
 * $ pulumi import azuredevops:index/buildDefinition:BuildDefinition example &#34;Example Project&#34;/10
 * ```
 * 
 * or
 * 
 * ```sh
 * $ pulumi import azuredevops:index/buildDefinition:BuildDefinition example 00000000-0000-0000-0000-000000000000/0
 * ```
 * 
 */
@ResourceType(type="azuredevops:index/buildDefinition:BuildDefinition")
public class BuildDefinition extends com.pulumi.resources.CustomResource {
    /**
     * The agent pool that should execute the build. Defaults to `Azure Pipelines`.
     * 
     */
    @Export(name="agentPoolName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> agentPoolName;

    /**
     * @return The agent pool that should execute the build. Defaults to `Azure Pipelines`.
     * 
     */
    public Output<Optional<String>> agentPoolName() {
        return Codegen.optional(this.agentPoolName);
    }
    /**
     * Continuous Integration trigger.
     * 
     */
    @Export(name="ciTrigger", refs={BuildDefinitionCiTrigger.class}, tree="[0]")
    private Output</* @Nullable */ BuildDefinitionCiTrigger> ciTrigger;

    /**
     * @return Continuous Integration trigger.
     * 
     */
    public Output<Optional<BuildDefinitionCiTrigger>> ciTrigger() {
        return Codegen.optional(this.ciTrigger);
    }
    /**
     * A `features` blocks as documented below.
     * 
     */
    @Export(name="features", refs={List.class,BuildDefinitionFeature.class}, tree="[0,1]")
    private Output</* @Nullable */ List<BuildDefinitionFeature>> features;

    /**
     * @return A `features` blocks as documented below.
     * 
     */
    public Output<Optional<List<BuildDefinitionFeature>>> features() {
        return Codegen.optional(this.features);
    }
    /**
     * The name of the build definition.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the build definition.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The folder path of the build definition.
     * 
     */
    @Export(name="path", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> path;

    /**
     * @return The folder path of the build definition.
     * 
     */
    public Output<Optional<String>> path() {
        return Codegen.optional(this.path);
    }
    /**
     * The project ID or project name.
     * 
     */
    @Export(name="projectId", refs={String.class}, tree="[0]")
    private Output<String> projectId;

    /**
     * @return The project ID or project name.
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }
    /**
     * Pull Request Integration trigger.
     * 
     */
    @Export(name="pullRequestTrigger", refs={BuildDefinitionPullRequestTrigger.class}, tree="[0]")
    private Output</* @Nullable */ BuildDefinitionPullRequestTrigger> pullRequestTrigger;

    /**
     * @return Pull Request Integration trigger.
     * 
     */
    public Output<Optional<BuildDefinitionPullRequestTrigger>> pullRequestTrigger() {
        return Codegen.optional(this.pullRequestTrigger);
    }
    /**
     * The queue status of the build definition. Valid values: `enabled` or `paused` or `disabled`. Defaults to `enabled`.
     * 
     */
    @Export(name="queueStatus", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> queueStatus;

    /**
     * @return The queue status of the build definition. Valid values: `enabled` or `paused` or `disabled`. Defaults to `enabled`.
     * 
     */
    public Output<Optional<String>> queueStatus() {
        return Codegen.optional(this.queueStatus);
    }
    /**
     * A `repository` block as documented below.
     * 
     */
    @Export(name="repository", refs={BuildDefinitionRepository.class}, tree="[0]")
    private Output<BuildDefinitionRepository> repository;

    /**
     * @return A `repository` block as documented below.
     * 
     */
    public Output<BuildDefinitionRepository> repository() {
        return this.repository;
    }
    /**
     * The revision of the build definition
     * 
     */
    @Export(name="revision", refs={Integer.class}, tree="[0]")
    private Output<Integer> revision;

    /**
     * @return The revision of the build definition
     * 
     */
    public Output<Integer> revision() {
        return this.revision;
    }
    @Export(name="schedules", refs={List.class,BuildDefinitionSchedule.class}, tree="[0,1]")
    private Output</* @Nullable */ List<BuildDefinitionSchedule>> schedules;

    public Output<Optional<List<BuildDefinitionSchedule>>> schedules() {
        return Codegen.optional(this.schedules);
    }
    /**
     * A list of variable group IDs (integers) to link to the build definition.
     * 
     */
    @Export(name="variableGroups", refs={List.class,Integer.class}, tree="[0,1]")
    private Output</* @Nullable */ List<Integer>> variableGroups;

    /**
     * @return A list of variable group IDs (integers) to link to the build definition.
     * 
     */
    public Output<Optional<List<Integer>>> variableGroups() {
        return Codegen.optional(this.variableGroups);
    }
    /**
     * A list of `variable` blocks, as documented below.
     * 
     */
    @Export(name="variables", refs={List.class,BuildDefinitionVariable.class}, tree="[0,1]")
    private Output</* @Nullable */ List<BuildDefinitionVariable>> variables;

    /**
     * @return A list of `variable` blocks, as documented below.
     * 
     */
    public Output<Optional<List<BuildDefinitionVariable>>> variables() {
        return Codegen.optional(this.variables);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public BuildDefinition(String name) {
        this(name, BuildDefinitionArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public BuildDefinition(String name, BuildDefinitionArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public BuildDefinition(String name, BuildDefinitionArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/buildDefinition:BuildDefinition", name, args == null ? BuildDefinitionArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private BuildDefinition(String name, Output<String> id, @Nullable BuildDefinitionState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/buildDefinition:BuildDefinition", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .aliases(List.of(
                Output.of(Alias.builder().type("azuredevops:Build/buildDefinition:BuildDefinition").build())
            ))
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
    public static BuildDefinition get(String name, Output<String> id, @Nullable BuildDefinitionState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new BuildDefinition(name, id, state, options);
    }
}
