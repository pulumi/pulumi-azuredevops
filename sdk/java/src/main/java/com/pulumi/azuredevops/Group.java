// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops;

import com.pulumi.azuredevops.GroupArgs;
import com.pulumi.azuredevops.Utilities;
import com.pulumi.azuredevops.inputs.GroupState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a group within Azure DevOps.
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
 * import com.pulumi.azuredevops.Group;
 * import com.pulumi.azuredevops.GroupArgs;
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
 *         final var example-readers = AzuredevopsFunctions.getGroup(GetGroupArgs.builder()
 *             .projectId(example.id())
 *             .name("Readers")
 *             .build());
 * 
 *         final var example-contributors = AzuredevopsFunctions.getGroup(GetGroupArgs.builder()
 *             .projectId(example.id())
 *             .name("Contributors")
 *             .build());
 * 
 *         var exampleGroup = new Group("exampleGroup", GroupArgs.builder()
 *             .scope(example.id())
 *             .displayName("Example group")
 *             .description("Example description")
 *             .members(            
 *                 example_readers.applyValue(example_readers -> example_readers.descriptor()),
 *                 example_contributors.applyValue(example_contributors -> example_contributors.descriptor()))
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
 * - [Azure DevOps Service REST API 7.0 - Groups](https://docs.microsoft.com/en-us/rest/api/azure/devops/graph/groups?view=azure-devops-rest-7.0)
 * 
 * ## PAT Permissions Required
 * 
 * - **Project &amp; Team**: Read, Write, &amp; Manage
 * 
 * ## Import
 * 
 * Azure DevOps groups can be imported using the group identity descriptor, e.g.
 * 
 * ```sh
 * $ pulumi import azuredevops:index/group:Group example aadgp.Uy0xLTktMTU1MTM3NDI0NS0xMjA0NDAwOTY5LTI0MDI5ODY0MTMtMjE3OTQwODYxNi0zLTIxNjc2NjQyNTMtMzI1Nzg0NDI4OS0yMjU4MjcwOTc0LTI2MDYxODY2NDU
 * ```
 * 
 */
@ResourceType(type="azuredevops:index/group:Group")
public class Group extends com.pulumi.resources.CustomResource {
    /**
     * The Description of the Project.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return The Description of the Project.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The identity (subject) descriptor of the Group.
     * 
     */
    @Export(name="descriptor", refs={String.class}, tree="[0]")
    private Output<String> descriptor;

    /**
     * @return The identity (subject) descriptor of the Group.
     * 
     */
    public Output<String> descriptor() {
        return this.descriptor;
    }
    /**
     * The name of a new Azure DevOps group that is not backed by an external provider. The `origin_id` and `mail` arguments cannot be used simultaneously with `display_name`.
     * 
     */
    @Export(name="displayName", refs={String.class}, tree="[0]")
    private Output<String> displayName;

    /**
     * @return The name of a new Azure DevOps group that is not backed by an external provider. The `origin_id` and `mail` arguments cannot be used simultaneously with `display_name`.
     * 
     */
    public Output<String> displayName() {
        return this.displayName;
    }
    /**
     * This represents the name of the container of origin for a graph member.
     * 
     */
    @Export(name="domain", refs={String.class}, tree="[0]")
    private Output<String> domain;

    /**
     * @return This represents the name of the container of origin for a graph member.
     * 
     */
    public Output<String> domain() {
        return this.domain;
    }
    /**
     * The ID of the Group.
     * 
     */
    @Export(name="groupId", refs={String.class}, tree="[0]")
    private Output<String> groupId;

    /**
     * @return The ID of the Group.
     * 
     */
    public Output<String> groupId() {
        return this.groupId;
    }
    /**
     * The mail address as a reference to an existing group from an external AD or AAD backed provider. The `scope`, `origin_id` and `display_name` arguments cannot be used simultaneously with `mail`.
     * 
     */
    @Export(name="mail", refs={String.class}, tree="[0]")
    private Output<String> mail;

    /**
     * @return The mail address as a reference to an existing group from an external AD or AAD backed provider. The `scope`, `origin_id` and `display_name` arguments cannot be used simultaneously with `mail`.
     * 
     */
    public Output<String> mail() {
        return this.mail;
    }
    /**
     * The member of the Group.
     * 
     * &gt; **NOTE:** It&#39;s possible to define group members both within the `azuredevops.Group` resource via the members block and by using the `azuredevops.GroupMembership` resource. However it&#39;s not possible to use both methods to manage group members, since there&#39;ll be conflicts.
     * 
     */
    @Export(name="members", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> members;

    /**
     * @return The member of the Group.
     * 
     * &gt; **NOTE:** It&#39;s possible to define group members both within the `azuredevops.Group` resource via the members block and by using the `azuredevops.GroupMembership` resource. However it&#39;s not possible to use both methods to manage group members, since there&#39;ll be conflicts.
     * 
     */
    public Output<List<String>> members() {
        return this.members;
    }
    /**
     * The type of source provider for the origin identifier (ex:AD, AAD, MSA)
     * 
     */
    @Export(name="origin", refs={String.class}, tree="[0]")
    private Output<String> origin;

    /**
     * @return The type of source provider for the origin identifier (ex:AD, AAD, MSA)
     * 
     */
    public Output<String> origin() {
        return this.origin;
    }
    /**
     * The OriginID as a reference to a group from an external AD or AAD backed provider. The `scope`, `mail` and `display_name` arguments cannot be used simultaneously with `origin_id`.
     * 
     */
    @Export(name="originId", refs={String.class}, tree="[0]")
    private Output<String> originId;

    /**
     * @return The OriginID as a reference to a group from an external AD or AAD backed provider. The `scope`, `mail` and `display_name` arguments cannot be used simultaneously with `origin_id`.
     * 
     */
    public Output<String> originId() {
        return this.originId;
    }
    /**
     * This is the PrincipalName of this graph member from the source provider.
     * 
     */
    @Export(name="principalName", refs={String.class}, tree="[0]")
    private Output<String> principalName;

    /**
     * @return This is the PrincipalName of this graph member from the source provider.
     * 
     */
    public Output<String> principalName() {
        return this.principalName;
    }
    /**
     * The scope of the group. A descriptor referencing the scope (collection, project) in which the group should be created. If omitted, will be created in the scope of the enclosing account or organization.x
     * 
     */
    @Export(name="scope", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> scope;

    /**
     * @return The scope of the group. A descriptor referencing the scope (collection, project) in which the group should be created. If omitted, will be created in the scope of the enclosing account or organization.x
     * 
     */
    public Output<Optional<String>> scope() {
        return Codegen.optional(this.scope);
    }
    /**
     * This field identifies the type of the graph subject (ex: Group, Scope, User).
     * 
     */
    @Export(name="subjectKind", refs={String.class}, tree="[0]")
    private Output<String> subjectKind;

    /**
     * @return This field identifies the type of the graph subject (ex: Group, Scope, User).
     * 
     */
    public Output<String> subjectKind() {
        return this.subjectKind;
    }
    /**
     * This url is the full route to the source resource of this graph subject.
     * 
     */
    @Export(name="url", refs={String.class}, tree="[0]")
    private Output<String> url;

    /**
     * @return This url is the full route to the source resource of this graph subject.
     * 
     */
    public Output<String> url() {
        return this.url;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Group(java.lang.String name) {
        this(name, GroupArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Group(java.lang.String name, @Nullable GroupArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Group(java.lang.String name, @Nullable GroupArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/group:Group", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private Group(java.lang.String name, Output<java.lang.String> id, @Nullable GroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/group:Group", name, state, makeResourceOptions(options, id), false);
    }

    private static GroupArgs makeArgs(@Nullable GroupArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? GroupArgs.Empty : args;
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
    public static Group get(java.lang.String name, Output<java.lang.String> id, @Nullable GroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Group(name, id, state, options);
    }
}
