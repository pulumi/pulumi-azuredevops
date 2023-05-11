// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops;

import com.pulumi.azuredevops.RepositoryPolicyAuthorEmailPatternArgs;
import com.pulumi.azuredevops.Utilities;
import com.pulumi.azuredevops.inputs.RepositoryPolicyAuthorEmailPatternState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manage author email pattern repository policy within Azure DevOps project.
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
 * import com.pulumi.azuredevops.Git;
 * import com.pulumi.azuredevops.GitArgs;
 * import com.pulumi.azuredevops.inputs.GitInitializationArgs;
 * import com.pulumi.azuredevops.RepositoryPolicyAuthorEmailPattern;
 * import com.pulumi.azuredevops.RepositoryPolicyAuthorEmailPatternArgs;
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
 *             .description(&#34;Managed by Terraform&#34;)
 *             .build());
 * 
 *         var exampleGit = new Git(&#34;exampleGit&#34;, GitArgs.builder()        
 *             .projectId(exampleProject.id())
 *             .initialization(GitInitializationArgs.builder()
 *                 .initType(&#34;Clean&#34;)
 *                 .build())
 *             .build());
 * 
 *         var exampleRepositoryPolicyAuthorEmailPattern = new RepositoryPolicyAuthorEmailPattern(&#34;exampleRepositoryPolicyAuthorEmailPattern&#34;, RepositoryPolicyAuthorEmailPatternArgs.builder()        
 *             .projectId(exampleProject.id())
 *             .enabled(true)
 *             .blocking(true)
 *             .authorEmailPatterns(            
 *                 &#34;user1@test.com&#34;,
 *                 &#34;user2@test.com&#34;)
 *             .repositoryIds(exampleGit.id())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ## Set project level repository policy
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.azuredevops.Project;
 * import com.pulumi.azuredevops.ProjectArgs;
 * import com.pulumi.azuredevops.RepositoryPolicyAuthorEmailPattern;
 * import com.pulumi.azuredevops.RepositoryPolicyAuthorEmailPatternArgs;
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
 *             .description(&#34;Managed by Terraform&#34;)
 *             .build());
 * 
 *         var exampleRepositoryPolicyAuthorEmailPattern = new RepositoryPolicyAuthorEmailPattern(&#34;exampleRepositoryPolicyAuthorEmailPattern&#34;, RepositoryPolicyAuthorEmailPatternArgs.builder()        
 *             .projectId(exampleProject.id())
 *             .enabled(true)
 *             .blocking(true)
 *             .authorEmailPatterns(            
 *                 &#34;user1@test.com&#34;,
 *                 &#34;user2@test.com&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Relevant Links
 * 
 * - [Azure DevOps Service REST API 6.0 - Policy Configurations](https://docs.microsoft.com/en-us/rest/api/azure/devops/policy/configurations?view=azure-devops-rest-6.0)
 * 
 * ## Import
 * 
 * Azure DevOps Branch Policies can be imported using the project ID and policy configuration ID
 * 
 * ```sh
 *  $ pulumi import azuredevops:index/repositoryPolicyAuthorEmailPattern:RepositoryPolicyAuthorEmailPattern example 00000000-0000-0000-0000-000000000000/0
 * ```
 * 
 */
@ResourceType(type="azuredevops:index/repositoryPolicyAuthorEmailPattern:RepositoryPolicyAuthorEmailPattern")
public class RepositoryPolicyAuthorEmailPattern extends com.pulumi.resources.CustomResource {
    /**
     * Block pushes with a commit author email that does not match the patterns. You can specify exact emails or use wildcards.
     * Email patterns prefixed with &#34;!&#34; are excluded. Order is important.
     * 
     */
    @Export(name="authorEmailPatterns", type=List.class, parameters={String.class})
    private Output<List<String>> authorEmailPatterns;

    /**
     * @return Block pushes with a commit author email that does not match the patterns. You can specify exact emails or use wildcards.
     * Email patterns prefixed with &#34;!&#34; are excluded. Order is important.
     * 
     */
    public Output<List<String>> authorEmailPatterns() {
        return this.authorEmailPatterns;
    }
    /**
     * A flag indicating if the policy should be blocking. Defaults to `true`.
     * 
     */
    @Export(name="blocking", type=Boolean.class, parameters={})
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
    @Export(name="enabled", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> enabled;

    /**
     * @return A flag indicating if the policy should be enabled. Defaults to `true`.
     * 
     */
    public Output<Optional<Boolean>> enabled() {
        return Codegen.optional(this.enabled);
    }
    /**
     * The ID of the project in which the policy will be created.
     * 
     */
    @Export(name="projectId", type=String.class, parameters={})
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
    @Export(name="repositoryIds", type=List.class, parameters={String.class})
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
    public RepositoryPolicyAuthorEmailPattern(String name) {
        this(name, RepositoryPolicyAuthorEmailPatternArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public RepositoryPolicyAuthorEmailPattern(String name, RepositoryPolicyAuthorEmailPatternArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public RepositoryPolicyAuthorEmailPattern(String name, RepositoryPolicyAuthorEmailPatternArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/repositoryPolicyAuthorEmailPattern:RepositoryPolicyAuthorEmailPattern", name, args == null ? RepositoryPolicyAuthorEmailPatternArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private RepositoryPolicyAuthorEmailPattern(String name, Output<String> id, @Nullable RepositoryPolicyAuthorEmailPatternState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/repositoryPolicyAuthorEmailPattern:RepositoryPolicyAuthorEmailPattern", name, state, makeResourceOptions(options, id));
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
    public static RepositoryPolicyAuthorEmailPattern get(String name, Output<String> id, @Nullable RepositoryPolicyAuthorEmailPatternState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new RepositoryPolicyAuthorEmailPattern(name, id, state, options);
    }
}