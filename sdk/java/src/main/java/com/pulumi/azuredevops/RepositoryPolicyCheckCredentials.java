// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops;

import com.pulumi.azuredevops.RepositoryPolicyCheckCredentialsArgs;
import com.pulumi.azuredevops.Utilities;
import com.pulumi.azuredevops.inputs.RepositoryPolicyCheckCredentialsState;
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
 * Manage a credentials check repository policy within Azure DevOps project. Block pushes that introduce files, folders, or branch names that include platform reserved names or incompatible characters.
 * 
 * &gt; If both project and project policy are enabled, the project policy has high priority.
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
 * import com.pulumi.azuredevops.RepositoryPolicyCheckCredentials;
 * import com.pulumi.azuredevops.RepositoryPolicyCheckCredentialsArgs;
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
 *         var exampleRepositoryPolicyCheckCredentials = new RepositoryPolicyCheckCredentials(&#34;exampleRepositoryPolicyCheckCredentials&#34;, RepositoryPolicyCheckCredentialsArgs.builder()        
 *             .projectId(exampleProject.id())
 *             .enabled(true)
 *             .blocking(true)
 *             .repositoryIds(exampleGit.id())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * # Set project level repository policy
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.azuredevops.Project;
 * import com.pulumi.azuredevops.ProjectArgs;
 * import com.pulumi.azuredevops.RepositoryPolicyCheckCredentials;
 * import com.pulumi.azuredevops.RepositoryPolicyCheckCredentialsArgs;
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
 *         var exampleRepositoryPolicyCheckCredentials = new RepositoryPolicyCheckCredentials(&#34;exampleRepositoryPolicyCheckCredentials&#34;, RepositoryPolicyCheckCredentialsArgs.builder()        
 *             .projectId(exampleProject.id())
 *             .enabled(true)
 *             .blocking(true)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ## Relevant Links
 * 
 * - [Azure DevOps Service REST API 6.0 - Policy Configurations](https://docs.microsoft.com/en-us/rest/api/azure/devops/policy/configurations?view=azure-devops-rest-6.0)
 * 
 * ## Import
 * 
 * Azure DevOps repository policies can be imported using the projectID/policyID or projectName/policyID
 * 
 * ```sh
 *  $ pulumi import azuredevops:index/repositoryPolicyCheckCredentials:RepositoryPolicyCheckCredentials example 00000000-0000-0000-0000-000000000000/0
 * ```
 * 
 */
@ResourceType(type="azuredevops:index/repositoryPolicyCheckCredentials:RepositoryPolicyCheckCredentials")
public class RepositoryPolicyCheckCredentials extends com.pulumi.resources.CustomResource {
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
    public RepositoryPolicyCheckCredentials(String name) {
        this(name, RepositoryPolicyCheckCredentialsArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public RepositoryPolicyCheckCredentials(String name, RepositoryPolicyCheckCredentialsArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public RepositoryPolicyCheckCredentials(String name, RepositoryPolicyCheckCredentialsArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/repositoryPolicyCheckCredentials:RepositoryPolicyCheckCredentials", name, args == null ? RepositoryPolicyCheckCredentialsArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private RepositoryPolicyCheckCredentials(String name, Output<String> id, @Nullable RepositoryPolicyCheckCredentialsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/repositoryPolicyCheckCredentials:RepositoryPolicyCheckCredentials", name, state, makeResourceOptions(options, id));
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
    public static RepositoryPolicyCheckCredentials get(String name, Output<String> id, @Nullable RepositoryPolicyCheckCredentialsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new RepositoryPolicyCheckCredentials(name, id, state, options);
    }
}