// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops;

import com.pulumi.azuredevops.BranchPolicyAutoReviewersArgs;
import com.pulumi.azuredevops.Utilities;
import com.pulumi.azuredevops.inputs.BranchPolicyAutoReviewersState;
import com.pulumi.azuredevops.outputs.BranchPolicyAutoReviewersSettings;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages required reviewer policy branch policy within Azure DevOps.
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
 * import com.pulumi.azuredevops.User;
 * import com.pulumi.azuredevops.UserArgs;
 * import com.pulumi.azuredevops.BranchPolicyAutoReviewers;
 * import com.pulumi.azuredevops.BranchPolicyAutoReviewersArgs;
 * import com.pulumi.azuredevops.inputs.BranchPolicyAutoReviewersSettingsArgs;
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
 *         var exampleUser = new User(&#34;exampleUser&#34;, UserArgs.builder()        
 *             .principalName(&#34;mail@email.com&#34;)
 *             .accountLicenseType(&#34;basic&#34;)
 *             .build());
 * 
 *         var exampleBranchPolicyAutoReviewers = new BranchPolicyAutoReviewers(&#34;exampleBranchPolicyAutoReviewers&#34;, BranchPolicyAutoReviewersArgs.builder()        
 *             .projectId(exampleProject.id())
 *             .enabled(true)
 *             .blocking(true)
 *             .settings(BranchPolicyAutoReviewersSettingsArgs.builder()
 *                 .autoReviewerIds(exampleUser.id())
 *                 .submitterCanVote(false)
 *                 .message(&#34;Auto reviewer&#34;)
 *                 .pathFilters(&#34;*{@literal /}src/*.ts&#34;)
 *                 .scopes(BranchPolicyAutoReviewersSettingsScopeArgs.builder()
 *                     .repositoryId(exampleGit.id())
 *                     .repositoryRef(exampleGit.defaultBranch())
 *                     .matchType(&#34;Exact&#34;)
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ## Relevant Links
 * 
 * - [Azure DevOps Service REST API 6.0 - Policy Configurations](https://docs.microsoft.com/en-us/rest/api/azure/devops/policy/configurations/create?view=azure-devops-rest-6.0)
 * 
 * ## Import
 * 
 * Azure DevOps Branch Policies can be imported using the project ID and policy configuration ID
 * 
 * ```sh
 *  $ pulumi import azuredevops:index/branchPolicyAutoReviewers:BranchPolicyAutoReviewers example 00000000-0000-0000-0000-000000000000/0
 * ```
 * 
 */
@ResourceType(type="azuredevops:index/branchPolicyAutoReviewers:BranchPolicyAutoReviewers")
public class BranchPolicyAutoReviewers extends com.pulumi.resources.CustomResource {
    /**
     * A flag indicating if the policy should be blocking. This relates to the Azure DevOps terms &#34;optional&#34; and &#34;required&#34; reviewers. Defaults to `true`.
     * 
     */
    @Export(name="blocking", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> blocking;

    /**
     * @return A flag indicating if the policy should be blocking. This relates to the Azure DevOps terms &#34;optional&#34; and &#34;required&#34; reviewers. Defaults to `true`.
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
    @Export(name="settings", type=BranchPolicyAutoReviewersSettings.class, parameters={})
    private Output<BranchPolicyAutoReviewersSettings> settings;

    /**
     * @return Configuration for the policy. This block must be defined exactly once.
     * 
     */
    public Output<BranchPolicyAutoReviewersSettings> settings() {
        return this.settings;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public BranchPolicyAutoReviewers(String name) {
        this(name, BranchPolicyAutoReviewersArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public BranchPolicyAutoReviewers(String name, BranchPolicyAutoReviewersArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public BranchPolicyAutoReviewers(String name, BranchPolicyAutoReviewersArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/branchPolicyAutoReviewers:BranchPolicyAutoReviewers", name, args == null ? BranchPolicyAutoReviewersArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private BranchPolicyAutoReviewers(String name, Output<String> id, @Nullable BranchPolicyAutoReviewersState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/branchPolicyAutoReviewers:BranchPolicyAutoReviewers", name, state, makeResourceOptions(options, id));
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
    public static BranchPolicyAutoReviewers get(String name, Output<String> id, @Nullable BranchPolicyAutoReviewersState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new BranchPolicyAutoReviewers(name, id, state, options);
    }
}