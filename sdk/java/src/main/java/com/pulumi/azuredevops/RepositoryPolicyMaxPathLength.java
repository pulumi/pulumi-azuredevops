// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops;

import com.pulumi.azuredevops.RepositoryPolicyMaxPathLengthArgs;
import com.pulumi.azuredevops.Utilities;
import com.pulumi.azuredevops.inputs.RepositoryPolicyMaxPathLengthState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manage a max path length repository policy within Azure DevOps project.
 * 
 * &gt; If both project and project policy are enabled, the repository policy has high priority.
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
 * import com.pulumi.azuredevops.Git;
 * import com.pulumi.azuredevops.GitArgs;
 * import com.pulumi.azuredevops.inputs.GitInitializationArgs;
 * import com.pulumi.azuredevops.RepositoryPolicyMaxPathLength;
 * import com.pulumi.azuredevops.RepositoryPolicyMaxPathLengthArgs;
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
 *         var exampleGit = new Git("exampleGit", GitArgs.builder()
 *             .projectId(example.id())
 *             .name("Sample Repo")
 *             .initialization(GitInitializationArgs.builder()
 *                 .initType("Clean")
 *                 .build())
 *             .build());
 * 
 *         var exampleRepositoryPolicyMaxPathLength = new RepositoryPolicyMaxPathLength("exampleRepositoryPolicyMaxPathLength", RepositoryPolicyMaxPathLengthArgs.builder()
 *             .projectId(example.id())
 *             .enabled(true)
 *             .blocking(true)
 *             .maxPathLength(500)
 *             .repositoryIds(exampleGit.id())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * # Set project level repository policy
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
 * import com.pulumi.azuredevops.RepositoryPolicyMaxPathLength;
 * import com.pulumi.azuredevops.RepositoryPolicyMaxPathLengthArgs;
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
 *         var exampleRepositoryPolicyMaxPathLength = new RepositoryPolicyMaxPathLength("exampleRepositoryPolicyMaxPathLength", RepositoryPolicyMaxPathLengthArgs.builder()
 *             .projectId(example.id())
 *             .enabled(true)
 *             .blocking(true)
 *             .maxPathLength(1000)
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
 * - [Azure DevOps Service REST API 7.0 - Policy Configurations](https://docs.microsoft.com/en-us/rest/api/azure/devops/policy/configurations?view=azure-devops-rest-7.0)
 * 
 * ## Import
 * 
 * Azure DevOps repository policies can be imported using the projectID/policyID or projectName/policyID:
 * 
 * ```sh
 * $ pulumi import azuredevops:index/repositoryPolicyMaxPathLength:RepositoryPolicyMaxPathLength example 00000000-0000-0000-0000-000000000000/0
 * ```
 * 
 */
@ResourceType(type="azuredevops:index/repositoryPolicyMaxPathLength:RepositoryPolicyMaxPathLength")
public class RepositoryPolicyMaxPathLength extends com.pulumi.resources.CustomResource {
    /**
     * A flag indicating if the policy should be blocking. Defaults to `true`.
     * 
     */
    @Export(name="blocking", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> blocking;

    /**
     * @return A flag indicating if the policy should be blocking. Defaults to `true`.
     * 
     */
    public Output<Optional<Boolean>> blocking() {
        return Codegen.optional(this.blocking);
    }
    /**
     * A flag indicating if the policy should be enabled. Defaults to `true`.
     * 
     */
    @Export(name="enabled", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> enabled;

    /**
     * @return A flag indicating if the policy should be enabled. Defaults to `true`.
     * 
     */
    public Output<Optional<Boolean>> enabled() {
        return Codegen.optional(this.enabled);
    }
    /**
     * Block pushes that introduce paths that exceed the specified length.
     * 
     */
    @Export(name="maxPathLength", refs={Integer.class}, tree="[0]")
    private Output<Integer> maxPathLength;

    /**
     * @return Block pushes that introduce paths that exceed the specified length.
     * 
     */
    public Output<Integer> maxPathLength() {
        return this.maxPathLength;
    }
    /**
     * The ID of the project in which the policy will be created.
     * 
     */
    @Export(name="projectId", refs={String.class}, tree="[0]")
    private Output<String> projectId;

    /**
     * @return The ID of the project in which the policy will be created.
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }
    /**
     * Control whether the policy is enabled for the repository or the project. If `repository_ids` not configured, the policy will be set to the project.
     * 
     */
    @Export(name="repositoryIds", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> repositoryIds;

    /**
     * @return Control whether the policy is enabled for the repository or the project. If `repository_ids` not configured, the policy will be set to the project.
     * 
     */
    public Output<Optional<List<String>>> repositoryIds() {
        return Codegen.optional(this.repositoryIds);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public RepositoryPolicyMaxPathLength(java.lang.String name) {
        this(name, RepositoryPolicyMaxPathLengthArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public RepositoryPolicyMaxPathLength(java.lang.String name, RepositoryPolicyMaxPathLengthArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public RepositoryPolicyMaxPathLength(java.lang.String name, RepositoryPolicyMaxPathLengthArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/repositoryPolicyMaxPathLength:RepositoryPolicyMaxPathLength", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private RepositoryPolicyMaxPathLength(java.lang.String name, Output<java.lang.String> id, @Nullable RepositoryPolicyMaxPathLengthState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/repositoryPolicyMaxPathLength:RepositoryPolicyMaxPathLength", name, state, makeResourceOptions(options, id), false);
    }

    private static RepositoryPolicyMaxPathLengthArgs makeArgs(RepositoryPolicyMaxPathLengthArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? RepositoryPolicyMaxPathLengthArgs.Empty : args;
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
    public static RepositoryPolicyMaxPathLength get(java.lang.String name, Output<java.lang.String> id, @Nullable RepositoryPolicyMaxPathLengthState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new RepositoryPolicyMaxPathLength(name, id, state, options);
    }
}
