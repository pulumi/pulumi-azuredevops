// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops;

import com.pulumi.azuredevops.FeedPermissionArgs;
import com.pulumi.azuredevops.Utilities;
import com.pulumi.azuredevops.inputs.FeedPermissionState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages creation of the Feed Permission within Azure DevOps organization.
 * 
 * ## Example Usage
 * 
 * ### Create Feed Permission
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
 * import com.pulumi.azuredevops.Group;
 * import com.pulumi.azuredevops.GroupArgs;
 * import com.pulumi.azuredevops.Feed;
 * import com.pulumi.azuredevops.FeedArgs;
 * import com.pulumi.azuredevops.FeedPermission;
 * import com.pulumi.azuredevops.FeedPermissionArgs;
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
 *         var exampleGroup = new Group("exampleGroup", GroupArgs.builder()
 *             .scope(example.id())
 *             .displayName("Example group")
 *             .description("Example description")
 *             .build());
 * 
 *         var exampleFeed = new Feed("exampleFeed", FeedArgs.builder()
 *             .name("examplefeed")
 *             .build());
 * 
 *         var permission = new FeedPermission("permission", FeedPermissionArgs.builder()
 *             .feedId(exampleFeed.id())
 *             .role("reader")
 *             .identityDescriptor(exampleGroup.descriptor())
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
 * - [Azure DevOps Service REST API 7.0 - Feed Management](https://learn.microsoft.com/en-us/rest/api/azure/devops/artifacts/feed-management?view=azure-devops-rest-7.0)
 * 
 */
@ResourceType(type="azuredevops:index/feedPermission:FeedPermission")
public class FeedPermission extends com.pulumi.resources.CustomResource {
    /**
     * The display name of the assignment
     * 
     */
    @Export(name="displayName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> displayName;

    /**
     * @return The display name of the assignment
     * 
     */
    public Output<Optional<String>> displayName() {
        return Codegen.optional(this.displayName);
    }
    /**
     * The ID of the Feed.
     * 
     */
    @Export(name="feedId", refs={String.class}, tree="[0]")
    private Output<String> feedId;

    /**
     * @return The ID of the Feed.
     * 
     */
    public Output<String> feedId() {
        return this.feedId;
    }
    /**
     * The Descriptor of identity you want to assign a role.
     * 
     */
    @Export(name="identityDescriptor", refs={String.class}, tree="[0]")
    private Output<String> identityDescriptor;

    /**
     * @return The Descriptor of identity you want to assign a role.
     * 
     */
    public Output<String> identityDescriptor() {
        return this.identityDescriptor;
    }
    /**
     * The ID of the identity.
     * 
     */
    @Export(name="identityId", refs={String.class}, tree="[0]")
    private Output<String> identityId;

    /**
     * @return The ID of the identity.
     * 
     */
    public Output<String> identityId() {
        return this.identityId;
    }
    /**
     * The ID of the Project Feed is created in. If not specified, feed will be created at the organization level.
     * 
     */
    @Export(name="projectId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> projectId;

    /**
     * @return The ID of the Project Feed is created in. If not specified, feed will be created at the organization level.
     * 
     */
    public Output<Optional<String>> projectId() {
        return Codegen.optional(this.projectId);
    }
    /**
     * The role to be assigned. Possible values are: `reader`, `contributor`, `collaborator`, `administrator`
     * 
     */
    @Export(name="role", refs={String.class}, tree="[0]")
    private Output<String> role;

    /**
     * @return The role to be assigned. Possible values are: `reader`, `contributor`, `collaborator`, `administrator`
     * 
     */
    public Output<String> role() {
        return this.role;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public FeedPermission(java.lang.String name) {
        this(name, FeedPermissionArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public FeedPermission(java.lang.String name, FeedPermissionArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public FeedPermission(java.lang.String name, FeedPermissionArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/feedPermission:FeedPermission", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private FeedPermission(java.lang.String name, Output<java.lang.String> id, @Nullable FeedPermissionState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/feedPermission:FeedPermission", name, state, makeResourceOptions(options, id), false);
    }

    private static FeedPermissionArgs makeArgs(FeedPermissionArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? FeedPermissionArgs.Empty : args;
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
    public static FeedPermission get(java.lang.String name, Output<java.lang.String> id, @Nullable FeedPermissionState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new FeedPermission(name, id, state, options);
    }
}
