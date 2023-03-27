// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops;

import com.pulumi.azuredevops.BuildDefinitionPermissionsArgs;
import com.pulumi.azuredevops.Utilities;
import com.pulumi.azuredevops.inputs.BuildDefinitionPermissionsState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages permissions for a Build Definition
 * 
 * &gt; **Note** Permissions can be assigned to group principals and not to single user principals.
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
 * import com.pulumi.azuredevops.Git;
 * import com.pulumi.azuredevops.GitArgs;
 * import com.pulumi.azuredevops.inputs.GitInitializationArgs;
 * import com.pulumi.azuredevops.BuildDefinition;
 * import com.pulumi.azuredevops.BuildDefinitionArgs;
 * import com.pulumi.azuredevops.inputs.BuildDefinitionCiTriggerArgs;
 * import com.pulumi.azuredevops.inputs.BuildDefinitionRepositoryArgs;
 * import com.pulumi.azuredevops.BuildDefinitionPermissions;
 * import com.pulumi.azuredevops.BuildDefinitionPermissionsArgs;
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
 *         final var example-readers = AzuredevopsFunctions.getGroup(GetGroupArgs.builder()
 *             .projectId(exampleProject.id())
 *             .name(&#34;Readers&#34;)
 *             .build());
 * 
 *         var exampleGit = new Git(&#34;exampleGit&#34;, GitArgs.builder()        
 *             .projectId(exampleProject.id())
 *             .initialization(GitInitializationArgs.builder()
 *                 .initType(&#34;Clean&#34;)
 *                 .build())
 *             .build());
 * 
 *         var exampleBuildDefinition = new BuildDefinition(&#34;exampleBuildDefinition&#34;, BuildDefinitionArgs.builder()        
 *             .projectId(exampleProject.id())
 *             .path(&#34;\\ExampleFolder&#34;)
 *             .ciTrigger(BuildDefinitionCiTriggerArgs.builder()
 *                 .useYaml(true)
 *                 .build())
 *             .repository(BuildDefinitionRepositoryArgs.builder()
 *                 .repoType(&#34;TfsGit&#34;)
 *                 .repoId(exampleGit.id())
 *                 .branchName(exampleGit.defaultBranch())
 *                 .ymlPath(&#34;azure-pipelines.yml&#34;)
 *                 .build())
 *             .build());
 * 
 *         var exampleBuildDefinitionPermissions = new BuildDefinitionPermissions(&#34;exampleBuildDefinitionPermissions&#34;, BuildDefinitionPermissionsArgs.builder()        
 *             .projectId(exampleProject.id())
 *             .principal(example_readers.applyValue(example_readers -&gt; example_readers.id()))
 *             .buildDefinitionId(exampleBuildDefinition.id())
 *             .permissions(Map.ofEntries(
 *                 Map.entry(&#34;ViewBuilds&#34;, &#34;Allow&#34;),
 *                 Map.entry(&#34;EditBuildQuality&#34;, &#34;Deny&#34;),
 *                 Map.entry(&#34;DeleteBuilds&#34;, &#34;Deny&#34;),
 *                 Map.entry(&#34;StopBuilds&#34;, &#34;Allow&#34;)
 *             ))
 *             .build());
 * 
 *     }
 * }
 * ```
 * ## Relevant Links
 * 
 * * [Azure DevOps Service REST API 6.0 - Security](https://docs.microsoft.com/en-us/rest/api/azure/devops/security/?view=azure-devops-rest-6.0)
 * 
 * ## PAT Permissions Required
 * 
 * - **Project &amp; Team**: vso.security_manage - Grants the ability to read, write, and manage security permissions.
 * 
 * ## Import
 * 
 * The resource does not support import.
 * 
 */
@ResourceType(type="azuredevops:index/buildDefinitionPermissions:BuildDefinitionPermissions")
public class BuildDefinitionPermissions extends com.pulumi.resources.CustomResource {
    /**
     * The id of the build definition to assign the permissions.
     * 
     */
    @Export(name="buildDefinitionId", type=String.class, parameters={})
    private Output<String> buildDefinitionId;

    /**
     * @return The id of the build definition to assign the permissions.
     * 
     */
    public Output<String> buildDefinitionId() {
        return this.buildDefinitionId;
    }
    /**
     * the permissions to assign. The following permissions are available.
     * 
     */
    @Export(name="permissions", type=Map.class, parameters={String.class, String.class})
    private Output<Map<String,String>> permissions;

    /**
     * @return the permissions to assign. The following permissions are available.
     * 
     */
    public Output<Map<String,String>> permissions() {
        return this.permissions;
    }
    /**
     * The **group** principal to assign the permissions.
     * 
     */
    @Export(name="principal", type=String.class, parameters={})
    private Output<String> principal;

    /**
     * @return The **group** principal to assign the permissions.
     * 
     */
    public Output<String> principal() {
        return this.principal;
    }
    /**
     * The ID of the project to assign the permissions.
     * 
     */
    @Export(name="projectId", type=String.class, parameters={})
    private Output<String> projectId;

    /**
     * @return The ID of the project to assign the permissions.
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }
    /**
     * Replace (`true`) or merge (`false`) the permissions. Default: `true`.
     * 
     */
    @Export(name="replace", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> replace;

    /**
     * @return Replace (`true`) or merge (`false`) the permissions. Default: `true`.
     * 
     */
    public Output<Optional<Boolean>> replace() {
        return Codegen.optional(this.replace);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public BuildDefinitionPermissions(String name) {
        this(name, BuildDefinitionPermissionsArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public BuildDefinitionPermissions(String name, BuildDefinitionPermissionsArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public BuildDefinitionPermissions(String name, BuildDefinitionPermissionsArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/buildDefinitionPermissions:BuildDefinitionPermissions", name, args == null ? BuildDefinitionPermissionsArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private BuildDefinitionPermissions(String name, Output<String> id, @Nullable BuildDefinitionPermissionsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/buildDefinitionPermissions:BuildDefinitionPermissions", name, state, makeResourceOptions(options, id));
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
    public static BuildDefinitionPermissions get(String name, Output<String> id, @Nullable BuildDefinitionPermissionsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new BuildDefinitionPermissions(name, id, state, options);
    }
}
