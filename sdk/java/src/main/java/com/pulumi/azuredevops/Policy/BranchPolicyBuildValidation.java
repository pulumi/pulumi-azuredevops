// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.Policy;

import com.pulumi.azuredevops.Policy.BranchPolicyBuildValidationArgs;
import com.pulumi.azuredevops.Policy.inputs.BranchPolicyBuildValidationState;
import com.pulumi.azuredevops.Policy.outputs.BranchPolicyBuildValidationSettings;
import com.pulumi.azuredevops.Utilities;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a build validation branch policy within Azure DevOps.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.azuredevops.Project;
 * import com.pulumi.azuredevops.Git;
 * import com.pulumi.azuredevops.GitArgs;
 * import com.pulumi.azuredevops.inputs.GitInitializationArgs;
 * import com.pulumi.azuredevops.BuildDefinition;
 * import com.pulumi.azuredevops.BuildDefinitionArgs;
 * import com.pulumi.azuredevops.inputs.BuildDefinitionRepositoryArgs;
 * import com.pulumi.azuredevops.BranchPolicyBuildValidation;
 * import com.pulumi.azuredevops.BranchPolicyBuildValidationArgs;
 * import com.pulumi.azuredevops.inputs.BranchPolicyBuildValidationSettingsArgs;
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
 *         var exampleProject = new Project(&#34;exampleProject&#34;);
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
 *             .repository(BuildDefinitionRepositoryArgs.builder()
 *                 .repoType(&#34;TfsGit&#34;)
 *                 .repoId(exampleGit.id())
 *                 .ymlPath(&#34;azure-pipelines.yml&#34;)
 *                 .build())
 *             .build());
 * 
 *         var exampleBranchPolicyBuildValidation = new BranchPolicyBuildValidation(&#34;exampleBranchPolicyBuildValidation&#34;, BranchPolicyBuildValidationArgs.builder()        
 *             .projectId(exampleProject.id())
 *             .enabled(true)
 *             .blocking(true)
 *             .settings(BranchPolicyBuildValidationSettingsArgs.builder()
 *                 .displayName(&#34;Example build validation policy&#34;)
 *                 .buildDefinitionId(exampleBuildDefinition.id())
 *                 .validDuration(720)
 *                 .filenamePatterns(                
 *                     &#34;/WebApp/*&#34;,
 *                     &#34;!/WebApp/Tests/*&#34;,
 *                     &#34;*.cs&#34;)
 *                 .scopes(                
 *                     BranchPolicyBuildValidationSettingsScopeArgs.builder()
 *                         .repositoryId(exampleGit.id())
 *                         .repositoryRef(exampleGit.defaultBranch())
 *                         .matchType(&#34;Exact&#34;)
 *                         .build(),
 *                     BranchPolicyBuildValidationSettingsScopeArgs.builder()
 *                         .repositoryId(exampleGit.id())
 *                         .repositoryRef(&#34;refs/heads/releases&#34;)
 *                         .matchType(&#34;Prefix&#34;)
 *                         .build(),
 *                     BranchPolicyBuildValidationSettingsScopeArgs.builder()
 *                         .matchType(&#34;DefaultBranch&#34;)
 *                         .build())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ## Relevant Links
 * 
 * - [Azure DevOps Service REST API 7.0 - Policy Configurations](https://docs.microsoft.com/en-us/rest/api/azure/devops/policy/configurations/create?view=azure-devops-rest-7.0)
 * 
 * ## Import
 * 
 * Azure DevOps Branch Policies can be imported using the project ID and policy configuration ID
 * 
 * ```sh
 *  $ pulumi import azuredevops:Policy/branchPolicyBuildValidation:BranchPolicyBuildValidation example 00000000-0000-0000-0000-000000000000/0
 * ```
 * 
 * @deprecated
 * azuredevops.policy.BranchPolicyBuildValidation has been deprecated in favor of azuredevops.BranchPolicyBuildValidation
 * 
 */
@Deprecated /* azuredevops.policy.BranchPolicyBuildValidation has been deprecated in favor of azuredevops.BranchPolicyBuildValidation */
@ResourceType(type="azuredevops:Policy/branchPolicyBuildValidation:BranchPolicyBuildValidation")
public class BranchPolicyBuildValidation extends com.pulumi.resources.CustomResource {
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
     * Configuration for the policy. This block must be defined exactly once.
     * 
     */
    @Export(name="settings", type=BranchPolicyBuildValidationSettings.class, parameters={})
    private Output<BranchPolicyBuildValidationSettings> settings;

    /**
     * @return Configuration for the policy. This block must be defined exactly once.
     * 
     */
    public Output<BranchPolicyBuildValidationSettings> settings() {
        return this.settings;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public BranchPolicyBuildValidation(String name) {
        this(name, BranchPolicyBuildValidationArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public BranchPolicyBuildValidation(String name, BranchPolicyBuildValidationArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public BranchPolicyBuildValidation(String name, BranchPolicyBuildValidationArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:Policy/branchPolicyBuildValidation:BranchPolicyBuildValidation", name, args == null ? BranchPolicyBuildValidationArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private BranchPolicyBuildValidation(String name, Output<String> id, @Nullable BranchPolicyBuildValidationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:Policy/branchPolicyBuildValidation:BranchPolicyBuildValidation", name, state, makeResourceOptions(options, id));
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
    public static BranchPolicyBuildValidation get(String name, Output<String> id, @Nullable BranchPolicyBuildValidationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new BranchPolicyBuildValidation(name, id, state, options);
    }
}
