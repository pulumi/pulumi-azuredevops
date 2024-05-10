// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops;

import com.pulumi.azuredevops.ProjectPipelineSettingsArgs;
import com.pulumi.azuredevops.Utilities;
import com.pulumi.azuredevops.inputs.ProjectPipelineSettingsState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Manages Pipeline Settings for Azure DevOps projects
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
 * import com.pulumi.azuredevops.ProjectPipelineSettings;
 * import com.pulumi.azuredevops.ProjectPipelineSettingsArgs;
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
 *             .visibility("private")
 *             .versionControl("Git")
 *             .workItemTemplate("Agile")
 *             .description("Managed by Terraform")
 *             .build());
 * 
 *         var exampleProjectPipelineSettings = new ProjectPipelineSettings("exampleProjectPipelineSettings", ProjectPipelineSettingsArgs.builder()        
 *             .projectId(example.id())
 *             .enforceJobScope(true)
 *             .enforceReferencedRepoScopedToken(false)
 *             .enforceSettableVar(true)
 *             .publishPipelineMetadata(false)
 *             .statusBadgesArePrivate(true)
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
 * No official documentation available
 * 
 * ## PAT Permissions Required
 * 
 * - Full Access
 * 
 * ## Import
 * 
 * Azure DevOps feature settings can be imported using the project id, e.g.
 * 
 * ```sh
 * $ pulumi import azuredevops:index/projectPipelineSettings:ProjectPipelineSettings example 00000000-0000-0000-0000-000000000000
 * ```
 * 
 */
@ResourceType(type="azuredevops:index/projectPipelineSettings:ProjectPipelineSettings")
public class ProjectPipelineSettings extends com.pulumi.resources.CustomResource {
    /**
     * Limit job authorization scope to current project for non-release pipelines.
     * 
     */
    @Export(name="enforceJobScope", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> enforceJobScope;

    /**
     * @return Limit job authorization scope to current project for non-release pipelines.
     * 
     */
    public Output<Boolean> enforceJobScope() {
        return this.enforceJobScope;
    }
    /**
     * Limit job authorization scope to current project for release pipelines.
     * 
     * &gt; **NOTE:**
     * The settings at the organization will override settings specified on the project.
     * For example, if `enforce_job_scope` is true at the organization, the `azuredevops.ProjectPipelineSettings` resource cannot set it to false.
     * In this scenario, the plan will always show that the resource is trying to change `enforce_job_scope` from `true` to `false`.
     * 
     */
    @Export(name="enforceJobScopeForRelease", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> enforceJobScopeForRelease;

    /**
     * @return Limit job authorization scope to current project for release pipelines.
     * 
     * &gt; **NOTE:**
     * The settings at the organization will override settings specified on the project.
     * For example, if `enforce_job_scope` is true at the organization, the `azuredevops.ProjectPipelineSettings` resource cannot set it to false.
     * In this scenario, the plan will always show that the resource is trying to change `enforce_job_scope` from `true` to `false`.
     * 
     */
    public Output<Boolean> enforceJobScopeForRelease() {
        return this.enforceJobScopeForRelease;
    }
    /**
     * Protect access to repositories in YAML pipelines.
     * 
     */
    @Export(name="enforceReferencedRepoScopedToken", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> enforceReferencedRepoScopedToken;

    /**
     * @return Protect access to repositories in YAML pipelines.
     * 
     */
    public Output<Boolean> enforceReferencedRepoScopedToken() {
        return this.enforceReferencedRepoScopedToken;
    }
    /**
     * Limit variables that can be set at queue time.
     * 
     */
    @Export(name="enforceSettableVar", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> enforceSettableVar;

    /**
     * @return Limit variables that can be set at queue time.
     * 
     */
    public Output<Boolean> enforceSettableVar() {
        return this.enforceSettableVar;
    }
    /**
     * The `id` of the project for which the project pipeline settings will be managed.
     * 
     */
    @Export(name="projectId", refs={String.class}, tree="[0]")
    private Output<String> projectId;

    /**
     * @return The `id` of the project for which the project pipeline settings will be managed.
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }
    /**
     * Publish metadata from pipelines.
     * 
     */
    @Export(name="publishPipelineMetadata", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> publishPipelineMetadata;

    /**
     * @return Publish metadata from pipelines.
     * 
     */
    public Output<Boolean> publishPipelineMetadata() {
        return this.publishPipelineMetadata;
    }
    /**
     * Disable anonymous access to badges.
     * 
     */
    @Export(name="statusBadgesArePrivate", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> statusBadgesArePrivate;

    /**
     * @return Disable anonymous access to badges.
     * 
     */
    public Output<Boolean> statusBadgesArePrivate() {
        return this.statusBadgesArePrivate;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ProjectPipelineSettings(String name) {
        this(name, ProjectPipelineSettingsArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ProjectPipelineSettings(String name, ProjectPipelineSettingsArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ProjectPipelineSettings(String name, ProjectPipelineSettingsArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/projectPipelineSettings:ProjectPipelineSettings", name, args == null ? ProjectPipelineSettingsArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ProjectPipelineSettings(String name, Output<String> id, @Nullable ProjectPipelineSettingsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/projectPipelineSettings:ProjectPipelineSettings", name, state, makeResourceOptions(options, id));
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
    public static ProjectPipelineSettings get(String name, Output<String> id, @Nullable ProjectPipelineSettingsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ProjectPipelineSettings(name, id, state, options);
    }
}
