// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops;

import com.pulumi.azuredevops.BranchPolicyWorkItemLinkingArgs;
import com.pulumi.azuredevops.Utilities;
import com.pulumi.azuredevops.inputs.BranchPolicyWorkItemLinkingState;
import com.pulumi.azuredevops.outputs.BranchPolicyWorkItemLinkingSettings;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Require associations between branches and a work item within Azure DevOps.
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
 * import com.pulumi.azuredevops.BranchPolicyWorkItemLinking;
 * import com.pulumi.azuredevops.BranchPolicyWorkItemLinkingArgs;
 * import com.pulumi.azuredevops.inputs.BranchPolicyWorkItemLinkingSettingsArgs;
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
 *         var exampleBranchPolicyWorkItemLinking = new BranchPolicyWorkItemLinking("exampleBranchPolicyWorkItemLinking", BranchPolicyWorkItemLinkingArgs.builder()
 *             .projectId(example.id())
 *             .enabled(true)
 *             .blocking(true)
 *             .settings(BranchPolicyWorkItemLinkingSettingsArgs.builder()
 *                 .scopes(                
 *                     BranchPolicyWorkItemLinkingSettingsScopeArgs.builder()
 *                         .repositoryId(exampleGit.id())
 *                         .repositoryRef(exampleGit.defaultBranch())
 *                         .matchType("Exact")
 *                         .build(),
 *                     BranchPolicyWorkItemLinkingSettingsScopeArgs.builder()
 *                         .repositoryId(exampleGit.id())
 *                         .repositoryRef("refs/heads/releases")
 *                         .matchType("Prefix")
 *                         .build(),
 *                     BranchPolicyWorkItemLinkingSettingsScopeArgs.builder()
 *                         .matchType("DefaultBranch")
 *                         .build())
 *                 .build())
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
 * - [Azure DevOps Service REST API 7.0 - Policy Configurations](https://docs.microsoft.com/en-us/rest/api/azure/devops/policy/configurations/create?view=azure-devops-rest-7.0)
 * 
 * ## Import
 * 
 * Azure DevOps Branch Policies can be imported using the project ID and policy configuration ID:
 * 
 * ```sh
 * $ pulumi import azuredevops:index/branchPolicyWorkItemLinking:BranchPolicyWorkItemLinking example 00000000-0000-0000-0000-000000000000/0
 * ```
 * 
 */
@ResourceType(type="azuredevops:index/branchPolicyWorkItemLinking:BranchPolicyWorkItemLinking")
public class BranchPolicyWorkItemLinking extends com.pulumi.resources.CustomResource {
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
     * A `settings` block as defined below. Configuration for the policy. This block must be defined exactly once.
     * 
     */
    @Export(name="settings", refs={BranchPolicyWorkItemLinkingSettings.class}, tree="[0]")
    private Output<BranchPolicyWorkItemLinkingSettings> settings;

    /**
     * @return A `settings` block as defined below. Configuration for the policy. This block must be defined exactly once.
     * 
     */
    public Output<BranchPolicyWorkItemLinkingSettings> settings() {
        return this.settings;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public BranchPolicyWorkItemLinking(java.lang.String name) {
        this(name, BranchPolicyWorkItemLinkingArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public BranchPolicyWorkItemLinking(java.lang.String name, BranchPolicyWorkItemLinkingArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public BranchPolicyWorkItemLinking(java.lang.String name, BranchPolicyWorkItemLinkingArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/branchPolicyWorkItemLinking:BranchPolicyWorkItemLinking", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private BranchPolicyWorkItemLinking(java.lang.String name, Output<java.lang.String> id, @Nullable BranchPolicyWorkItemLinkingState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/branchPolicyWorkItemLinking:BranchPolicyWorkItemLinking", name, state, makeResourceOptions(options, id), false);
    }

    private static BranchPolicyWorkItemLinkingArgs makeArgs(BranchPolicyWorkItemLinkingArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? BranchPolicyWorkItemLinkingArgs.Empty : args;
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
    public static BranchPolicyWorkItemLinking get(java.lang.String name, Output<java.lang.String> id, @Nullable BranchPolicyWorkItemLinkingState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new BranchPolicyWorkItemLinking(name, id, state, options);
    }
}
