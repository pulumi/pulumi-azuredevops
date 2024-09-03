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
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
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
 *         var example = new Project("example", ProjectArgs.builder()
 *             .name("Example Project")
 *             .workItemTemplate("Agile")
 *             .versionControl("Git")
 *             .visibility("private")
 *             .description("Managed by Terraform")
 *             .build());
 * 
 *         final var example-readers = AzuredevopsFunctions.getGroup(GetGroupArgs.builder()
 *             .projectId(example.id())
 *             .name("Readers")
 *             .build());
 * 
 *         var exampleGit = new Git("exampleGit", GitArgs.builder()
 *             .projectId(example.id())
 *             .name("Example Repository")
 *             .initialization(GitInitializationArgs.builder()
 *                 .initType("Clean")
 *                 .build())
 *             .build());
 * 
 *         var exampleBuildDefinition = new BuildDefinition("exampleBuildDefinition", BuildDefinitionArgs.builder()
 *             .projectId(example.id())
 *             .name("Example Build Definition")
 *             .path("\\ExampleFolder")
 *             .ciTrigger(BuildDefinitionCiTriggerArgs.builder()
 *                 .useYaml(true)
 *                 .build())
 *             .repository(BuildDefinitionRepositoryArgs.builder()
 *                 .repoType("TfsGit")
 *                 .repoId(exampleGit.id())
 *                 .branchName(exampleGit.defaultBranch())
 *                 .ymlPath("azure-pipelines.yml")
 *                 .build())
 *             .build());
 * 
 *         var exampleBuildDefinitionPermissions = new BuildDefinitionPermissions("exampleBuildDefinitionPermissions", BuildDefinitionPermissionsArgs.builder()
 *             .projectId(example.id())
 *             .principal(example_readers.applyValue(example_readers -> example_readers.id()))
 *             .buildDefinitionId(exampleBuildDefinition.id())
 *             .permissions(Map.ofEntries(
 *                 Map.entry("ViewBuilds", "Allow"),
 *                 Map.entry("EditBuildQuality", "Deny"),
 *                 Map.entry("DeleteBuilds", "Deny"),
 *                 Map.entry("StopBuilds", "Allow")
 *             ))
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Relevant Links
 * 
 * * [Azure DevOps Service REST API 7.0 - Security](https://docs.microsoft.com/en-us/rest/api/azure/devops/security/?view=azure-devops-rest-7.0)
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
    @Export(name="buildDefinitionId", refs={String.class}, tree="[0]")
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
     * &lt;table&gt;
     * &lt;thead&gt;
     * &lt;tr&gt;
     * &lt;th&gt;Permission&lt;/th&gt;
     * &lt;th&gt;Description&lt;/th&gt;
     * &lt;/tr&gt;
     * &lt;/thead&gt;
     * &lt;tbody&gt;
     * &lt;tr&gt;
     * &lt;td&gt;ViewBuilds&lt;/td&gt;
     * &lt;td&gt;View builds&lt;/td&gt;
     * &lt;/tr&gt;
     * &lt;tr&gt;
     * &lt;td&gt;EditBuildQuality&lt;/td&gt;
     * &lt;td&gt;Edit build quality&lt;/td&gt;
     * &lt;/tr&gt;
     * &lt;tr&gt;
     * &lt;td&gt;RetainIndefinitely&lt;/td&gt;
     * &lt;td&gt;Retain indefinitely&lt;/td&gt;
     * &lt;/tr&gt;
     * &lt;tr&gt;
     * &lt;td&gt;DeleteBuilds&lt;/td&gt;
     * &lt;td&gt;Delete builds&lt;/td&gt;
     * &lt;/tr&gt;
     * &lt;tr&gt;
     * &lt;td&gt;ManageBuildQualities&lt;/td&gt;
     * &lt;td&gt;Manage build qualities&lt;/td&gt;
     * &lt;/tr&gt;
     * &lt;tr&gt;
     * &lt;td&gt;DestroyBuilds&lt;/td&gt;
     * &lt;td&gt;Destroy builds&lt;/td&gt;
     * &lt;/tr&gt;
     * &lt;tr&gt;
     * &lt;td&gt;UpdateBuildInformation&lt;/td&gt;
     * &lt;td&gt;Update build information&lt;/td&gt;
     * &lt;/tr&gt;
     * &lt;tr&gt;
     * &lt;td&gt;QueueBuilds&lt;/td&gt;
     * &lt;td&gt;Queue builds&lt;/td&gt;
     * &lt;/tr&gt;
     * &lt;tr&gt;
     * &lt;td&gt;ManageBuildQueue&lt;/td&gt;
     * &lt;td&gt;Manage build queue&lt;/td&gt;
     * &lt;/tr&gt;
     * &lt;tr&gt;
     * &lt;td&gt;StopBuilds&lt;/td&gt;
     * &lt;td&gt;Stop builds&lt;/td&gt;
     * &lt;/tr&gt;
     * &lt;tr&gt;
     * &lt;td&gt;ViewBuildDefinition&lt;/td&gt;
     * &lt;td&gt;View build pipeline&lt;/td&gt;
     * &lt;/tr&gt;
     * &lt;tr&gt;
     * &lt;td&gt;EditBuildDefinition&lt;/td&gt;
     * &lt;td&gt;Edit build pipeline&lt;/td&gt;
     * &lt;/tr&gt;
     * &lt;tr&gt;
     * &lt;td&gt;DeleteBuildDefinition&lt;/td&gt;
     * &lt;td&gt;Delete build pipeline&lt;/td&gt;
     * &lt;/tr&gt;
     * &lt;tr&gt;
     * &lt;td&gt;OverrideBuildCheckInValidation&lt;/td&gt;
     * &lt;td&gt;Override check-in validation by build&lt;/td&gt;
     * &lt;/tr&gt;
     * &lt;tr&gt;
     * &lt;td&gt;AdministerBuildPermissions&lt;/td&gt;
     * &lt;td&gt;Administer build permissions&lt;/td&gt;
     * &lt;/tr&gt;
     * &lt;/tbody&gt;
     * &lt;/table&gt;
     * 
     */
    @Export(name="permissions", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> permissions;

    /**
     * @return the permissions to assign. The following permissions are available.
     * &lt;table&gt;
     * &lt;thead&gt;
     * &lt;tr&gt;
     * &lt;th&gt;Permission&lt;/th&gt;
     * &lt;th&gt;Description&lt;/th&gt;
     * &lt;/tr&gt;
     * &lt;/thead&gt;
     * &lt;tbody&gt;
     * &lt;tr&gt;
     * &lt;td&gt;ViewBuilds&lt;/td&gt;
     * &lt;td&gt;View builds&lt;/td&gt;
     * &lt;/tr&gt;
     * &lt;tr&gt;
     * &lt;td&gt;EditBuildQuality&lt;/td&gt;
     * &lt;td&gt;Edit build quality&lt;/td&gt;
     * &lt;/tr&gt;
     * &lt;tr&gt;
     * &lt;td&gt;RetainIndefinitely&lt;/td&gt;
     * &lt;td&gt;Retain indefinitely&lt;/td&gt;
     * &lt;/tr&gt;
     * &lt;tr&gt;
     * &lt;td&gt;DeleteBuilds&lt;/td&gt;
     * &lt;td&gt;Delete builds&lt;/td&gt;
     * &lt;/tr&gt;
     * &lt;tr&gt;
     * &lt;td&gt;ManageBuildQualities&lt;/td&gt;
     * &lt;td&gt;Manage build qualities&lt;/td&gt;
     * &lt;/tr&gt;
     * &lt;tr&gt;
     * &lt;td&gt;DestroyBuilds&lt;/td&gt;
     * &lt;td&gt;Destroy builds&lt;/td&gt;
     * &lt;/tr&gt;
     * &lt;tr&gt;
     * &lt;td&gt;UpdateBuildInformation&lt;/td&gt;
     * &lt;td&gt;Update build information&lt;/td&gt;
     * &lt;/tr&gt;
     * &lt;tr&gt;
     * &lt;td&gt;QueueBuilds&lt;/td&gt;
     * &lt;td&gt;Queue builds&lt;/td&gt;
     * &lt;/tr&gt;
     * &lt;tr&gt;
     * &lt;td&gt;ManageBuildQueue&lt;/td&gt;
     * &lt;td&gt;Manage build queue&lt;/td&gt;
     * &lt;/tr&gt;
     * &lt;tr&gt;
     * &lt;td&gt;StopBuilds&lt;/td&gt;
     * &lt;td&gt;Stop builds&lt;/td&gt;
     * &lt;/tr&gt;
     * &lt;tr&gt;
     * &lt;td&gt;ViewBuildDefinition&lt;/td&gt;
     * &lt;td&gt;View build pipeline&lt;/td&gt;
     * &lt;/tr&gt;
     * &lt;tr&gt;
     * &lt;td&gt;EditBuildDefinition&lt;/td&gt;
     * &lt;td&gt;Edit build pipeline&lt;/td&gt;
     * &lt;/tr&gt;
     * &lt;tr&gt;
     * &lt;td&gt;DeleteBuildDefinition&lt;/td&gt;
     * &lt;td&gt;Delete build pipeline&lt;/td&gt;
     * &lt;/tr&gt;
     * &lt;tr&gt;
     * &lt;td&gt;OverrideBuildCheckInValidation&lt;/td&gt;
     * &lt;td&gt;Override check-in validation by build&lt;/td&gt;
     * &lt;/tr&gt;
     * &lt;tr&gt;
     * &lt;td&gt;AdministerBuildPermissions&lt;/td&gt;
     * &lt;td&gt;Administer build permissions&lt;/td&gt;
     * &lt;/tr&gt;
     * &lt;/tbody&gt;
     * &lt;/table&gt;
     * 
     */
    public Output<Map<String,String>> permissions() {
        return this.permissions;
    }
    /**
     * The **group** principal to assign the permissions.
     * 
     */
    @Export(name="principal", refs={String.class}, tree="[0]")
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
    @Export(name="projectId", refs={String.class}, tree="[0]")
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
    @Export(name="replace", refs={Boolean.class}, tree="[0]")
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
    public BuildDefinitionPermissions(java.lang.String name) {
        this(name, BuildDefinitionPermissionsArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public BuildDefinitionPermissions(java.lang.String name, BuildDefinitionPermissionsArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public BuildDefinitionPermissions(java.lang.String name, BuildDefinitionPermissionsArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/buildDefinitionPermissions:BuildDefinitionPermissions", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private BuildDefinitionPermissions(java.lang.String name, Output<java.lang.String> id, @Nullable BuildDefinitionPermissionsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/buildDefinitionPermissions:BuildDefinitionPermissions", name, state, makeResourceOptions(options, id), false);
    }

    private static BuildDefinitionPermissionsArgs makeArgs(BuildDefinitionPermissionsArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? BuildDefinitionPermissionsArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
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
    public static BuildDefinitionPermissions get(java.lang.String name, Output<java.lang.String> id, @Nullable BuildDefinitionPermissionsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new BuildDefinitionPermissions(name, id, state, options);
    }
}
