// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops;

import com.pulumi.azuredevops.GroupMembershipArgs;
import com.pulumi.azuredevops.Utilities;
import com.pulumi.azuredevops.inputs.GroupMembershipState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages group membership within Azure DevOps.
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
 * import com.pulumi.azuredevops.User;
 * import com.pulumi.azuredevops.UserArgs;
 * import com.pulumi.azuredevops.AzuredevopsFunctions;
 * import com.pulumi.azuredevops.inputs.GetGroupArgs;
 * import com.pulumi.azuredevops.GroupMembership;
 * import com.pulumi.azuredevops.GroupMembershipArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App }{{@code
 *     public static void main(String[] args) }{{@code
 *         Pulumi.run(App::stack);
 *     }}{@code
 * 
 *     public static void stack(Context ctx) }{{@code
 *         var exampleProject = new Project("exampleProject", ProjectArgs.builder()
 *             .name("Example Project")
 *             .build());
 * 
 *         var exampleUser = new User("exampleUser", UserArgs.builder()
 *             .principalName("foo}{@literal @}{@code contoso.com")
 *             .build());
 * 
 *         final var example = AzuredevopsFunctions.getGroup(GetGroupArgs.builder()
 *             .projectId(exampleProject.id())
 *             .name("Build Administrators")
 *             .build());
 * 
 *         var exampleGroupMembership = new GroupMembership("exampleGroupMembership", GroupMembershipArgs.builder()
 *             .group(example.applyValue(getGroupResult -> getGroupResult).applyValue(example -> example.applyValue(getGroupResult -> getGroupResult.descriptor())))
 *             .members(exampleUser.descriptor())
 *             .build());
 * 
 *     }}{@code
 * }}{@code
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Relevant Links
 * 
 * - [Azure DevOps Service REST API 7.0 - Memberships](https://docs.microsoft.com/en-us/rest/api/azure/devops/graph/memberships?view=azure-devops-rest-7.0)
 * 
 * ## PAT Permissions Required
 * 
 * - **Deployment Groups**: Read &amp; Manage
 * 
 * ## Import
 * 
 * Not supported.
 * 
 */
@ResourceType(type="azuredevops:index/groupMembership:GroupMembership")
public class GroupMembership extends com.pulumi.resources.CustomResource {
    /**
     * The descriptor of the group being managed.
     * 
     */
    @Export(name="group", refs={String.class}, tree="[0]")
    private Output<String> group;

    /**
     * @return The descriptor of the group being managed.
     * 
     */
    public Output<String> group() {
        return this.group;
    }
    /**
     * A list of user or group descriptors that will become members of the group.
     * 
     * &gt; **NOTE** It&#39;s possible to define group members both within the `azuredevops.GroupMembership resource` via the members block and by using the `azuredevops.Group` resource. However it&#39;s not possible to use both methods to manage group members, since there&#39;ll be conflicts.
     * 
     * &gt; **NOTE**  The `members` uses `descriptor` as the identifier not Resource ID or others.
     * 
     */
    @Export(name="members", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> members;

    /**
     * @return A list of user or group descriptors that will become members of the group.
     * 
     * &gt; **NOTE** It&#39;s possible to define group members both within the `azuredevops.GroupMembership resource` via the members block and by using the `azuredevops.Group` resource. However it&#39;s not possible to use both methods to manage group members, since there&#39;ll be conflicts.
     * 
     * &gt; **NOTE**  The `members` uses `descriptor` as the identifier not Resource ID or others.
     * 
     */
    public Output<List<String>> members() {
        return this.members;
    }
    /**
     * The mode how the resource manages group members.
     * - `mode == add`: the resource will ensure that all specified members will be part of the referenced group
     * - `mode == overwrite`: the resource will replace all existing members with the members specified within the `members` block
     * 
     * &gt; **NOTE** To clear all members from a group, specify an empty list of descriptors in the `members` attribute and set the `mode` member to `overwrite`.
     * 
     */
    @Export(name="mode", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> mode;

    /**
     * @return The mode how the resource manages group members.
     * - `mode == add`: the resource will ensure that all specified members will be part of the referenced group
     * - `mode == overwrite`: the resource will replace all existing members with the members specified within the `members` block
     * 
     * &gt; **NOTE** To clear all members from a group, specify an empty list of descriptors in the `members` attribute and set the `mode` member to `overwrite`.
     * 
     */
    public Output<Optional<String>> mode() {
        return Codegen.optional(this.mode);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public GroupMembership(java.lang.String name) {
        this(name, GroupMembershipArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public GroupMembership(java.lang.String name, GroupMembershipArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public GroupMembership(java.lang.String name, GroupMembershipArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/groupMembership:GroupMembership", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private GroupMembership(java.lang.String name, Output<java.lang.String> id, @Nullable GroupMembershipState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/groupMembership:GroupMembership", name, state, makeResourceOptions(options, id), false);
    }

    private static GroupMembershipArgs makeArgs(GroupMembershipArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? GroupMembershipArgs.Empty : args;
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
    public static GroupMembership get(java.lang.String name, Output<java.lang.String> id, @Nullable GroupMembershipState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new GroupMembership(name, id, state, options);
    }
}
