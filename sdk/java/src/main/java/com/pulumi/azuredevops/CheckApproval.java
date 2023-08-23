// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops;

import com.pulumi.azuredevops.CheckApprovalArgs;
import com.pulumi.azuredevops.Utilities;
import com.pulumi.azuredevops.inputs.CheckApprovalState;
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
 * Manages a Approval Check.
 * 
 * ## Example Usage
 * ### Protect an environment
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.azuredevops.Project;
 * import com.pulumi.azuredevops.Environment;
 * import com.pulumi.azuredevops.EnvironmentArgs;
 * import com.pulumi.azuredevops.Group;
 * import com.pulumi.azuredevops.GroupArgs;
 * import com.pulumi.azuredevops.CheckApproval;
 * import com.pulumi.azuredevops.CheckApprovalArgs;
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
 *         var exampleEnvironment = new Environment(&#34;exampleEnvironment&#34;, EnvironmentArgs.builder()        
 *             .projectId(exampleProject.id())
 *             .build());
 * 
 *         var exampleGroup = new Group(&#34;exampleGroup&#34;, GroupArgs.builder()        
 *             .displayName(&#34;some-azdo-group&#34;)
 *             .build());
 * 
 *         var exampleCheckApproval = new CheckApproval(&#34;exampleCheckApproval&#34;, CheckApprovalArgs.builder()        
 *             .projectId(exampleProject.id())
 *             .targetResourceId(exampleEnvironment.id())
 *             .targetResourceType(&#34;environment&#34;)
 *             .requesterCanApprove(true)
 *             .approvers(exampleGroup.originId())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Importing this resource is not supported.
 * 
 */
@ResourceType(type="azuredevops:index/checkApproval:CheckApproval")
public class CheckApproval extends com.pulumi.resources.CustomResource {
    /**
     * Specifies a list of approver IDs.
     * 
     */
    @Export(name="approvers", type=List.class, parameters={String.class})
    private Output<List<String>> approvers;

    /**
     * @return Specifies a list of approver IDs.
     * 
     */
    public Output<List<String>> approvers() {
        return this.approvers;
    }
    /**
     * The instructions for the approvers.
     * 
     */
    @Export(name="instructions", type=String.class, parameters={})
    private Output</* @Nullable */ String> instructions;

    /**
     * @return The instructions for the approvers.
     * 
     */
    public Output<Optional<String>> instructions() {
        return Codegen.optional(this.instructions);
    }
    /**
     * The minimum number of approvers. This property is applicable when there is more than 1 approver.
     * 
     */
    @Export(name="minimumRequiredApprovers", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> minimumRequiredApprovers;

    /**
     * @return The minimum number of approvers. This property is applicable when there is more than 1 approver.
     * 
     */
    public Output<Optional<Integer>> minimumRequiredApprovers() {
        return Codegen.optional(this.minimumRequiredApprovers);
    }
    /**
     * The project ID. Changing this forces a new Approval Check to be created.
     * 
     */
    @Export(name="projectId", type=String.class, parameters={})
    private Output<String> projectId;

    /**
     * @return The project ID. Changing this forces a new Approval Check to be created.
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }
    /**
     * Can the requestor approve? Defaults to `false`.
     * 
     */
    @Export(name="requesterCanApprove", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> requesterCanApprove;

    /**
     * @return Can the requestor approve? Defaults to `false`.
     * 
     */
    public Output<Optional<Boolean>> requesterCanApprove() {
        return Codegen.optional(this.requesterCanApprove);
    }
    /**
     * The ID of the resource being protected by the check. Changing this forces a new Approval Check to be created.
     * 
     */
    @Export(name="targetResourceId", type=String.class, parameters={})
    private Output<String> targetResourceId;

    /**
     * @return The ID of the resource being protected by the check. Changing this forces a new Approval Check to be created.
     * 
     */
    public Output<String> targetResourceId() {
        return this.targetResourceId;
    }
    /**
     * The type of resource being protected by the check. Valid values: `endpoint`, `environment`, `queue`, `repository`, `securefile`, `variablegroup`. Changing this forces a new Approval Check to be created.
     * 
     */
    @Export(name="targetResourceType", type=String.class, parameters={})
    private Output<String> targetResourceType;

    /**
     * @return The type of resource being protected by the check. Valid values: `endpoint`, `environment`, `queue`, `repository`, `securefile`, `variablegroup`. Changing this forces a new Approval Check to be created.
     * 
     */
    public Output<String> targetResourceType() {
        return this.targetResourceType;
    }
    /**
     * The timeout in minutes for the approval.  Defaults to `43200`.
     * 
     */
    @Export(name="timeout", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> timeout;

    /**
     * @return The timeout in minutes for the approval.  Defaults to `43200`.
     * 
     */
    public Output<Optional<Integer>> timeout() {
        return Codegen.optional(this.timeout);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public CheckApproval(String name) {
        this(name, CheckApprovalArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public CheckApproval(String name, CheckApprovalArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public CheckApproval(String name, CheckApprovalArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/checkApproval:CheckApproval", name, args == null ? CheckApprovalArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private CheckApproval(String name, Output<String> id, @Nullable CheckApprovalState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/checkApproval:CheckApproval", name, state, makeResourceOptions(options, id));
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
    public static CheckApproval get(String name, Output<String> id, @Nullable CheckApprovalState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new CheckApproval(name, id, state, options);
    }
}