// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops;

import com.pulumi.azuredevops.CheckBranchControlArgs;
import com.pulumi.azuredevops.Utilities;
import com.pulumi.azuredevops.inputs.CheckBranchControlState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a branch control check on a resource within Azure DevOps.
 * 
 * ## Example Usage
 * 
 * ### Protect a service connection
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.azuredevops.Project;
 * import com.pulumi.azuredevops.ProjectArgs;
 * import com.pulumi.azuredevops.ServiceEndpointGeneric;
 * import com.pulumi.azuredevops.ServiceEndpointGenericArgs;
 * import com.pulumi.azuredevops.CheckBranchControl;
 * import com.pulumi.azuredevops.CheckBranchControlArgs;
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
 *         var example = new Project(&#34;example&#34;, ProjectArgs.builder()        
 *             .name(&#34;Example Project&#34;)
 *             .build());
 * 
 *         var exampleServiceEndpointGeneric = new ServiceEndpointGeneric(&#34;exampleServiceEndpointGeneric&#34;, ServiceEndpointGenericArgs.builder()        
 *             .projectId(example.id())
 *             .serverUrl(&#34;https://some-server.example.com&#34;)
 *             .username(&#34;username&#34;)
 *             .password(&#34;password&#34;)
 *             .serviceEndpointName(&#34;Example Generic&#34;)
 *             .description(&#34;Managed by Terraform&#34;)
 *             .build());
 * 
 *         var exampleCheckBranchControl = new CheckBranchControl(&#34;exampleCheckBranchControl&#34;, CheckBranchControlArgs.builder()        
 *             .projectId(example.id())
 *             .displayName(&#34;Managed by Terraform&#34;)
 *             .targetResourceId(exampleServiceEndpointGeneric.id())
 *             .targetResourceType(&#34;endpoint&#34;)
 *             .allowedBranches(&#34;refs/heads/main, refs/heads/features/*&#34;)
 *             .timeout(1440)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ### Protect an environment
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.azuredevops.Project;
 * import com.pulumi.azuredevops.ProjectArgs;
 * import com.pulumi.azuredevops.Environment;
 * import com.pulumi.azuredevops.EnvironmentArgs;
 * import com.pulumi.azuredevops.CheckBranchControl;
 * import com.pulumi.azuredevops.CheckBranchControlArgs;
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
 *         var example = new Project(&#34;example&#34;, ProjectArgs.builder()        
 *             .name(&#34;Example Project&#34;)
 *             .build());
 * 
 *         var exampleEnvironment = new Environment(&#34;exampleEnvironment&#34;, EnvironmentArgs.builder()        
 *             .projectId(example.id())
 *             .name(&#34;Example Environment&#34;)
 *             .build());
 * 
 *         var exampleCheckBranchControl = new CheckBranchControl(&#34;exampleCheckBranchControl&#34;, CheckBranchControlArgs.builder()        
 *             .projectId(example.id())
 *             .displayName(&#34;Managed by Terraform&#34;)
 *             .targetResourceId(exampleEnvironment.id())
 *             .targetResourceType(&#34;environment&#34;)
 *             .allowedBranches(&#34;refs/heads/main, refs/heads/features/*&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ### Protect an agent queue
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.azuredevops.Project;
 * import com.pulumi.azuredevops.ProjectArgs;
 * import com.pulumi.azuredevops.Pool;
 * import com.pulumi.azuredevops.PoolArgs;
 * import com.pulumi.azuredevops.Queue;
 * import com.pulumi.azuredevops.QueueArgs;
 * import com.pulumi.azuredevops.CheckBranchControl;
 * import com.pulumi.azuredevops.CheckBranchControlArgs;
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
 *         var example = new Project(&#34;example&#34;, ProjectArgs.builder()        
 *             .name(&#34;Example Project&#34;)
 *             .build());
 * 
 *         var examplePool = new Pool(&#34;examplePool&#34;, PoolArgs.builder()        
 *             .name(&#34;example-pool&#34;)
 *             .build());
 * 
 *         var exampleQueue = new Queue(&#34;exampleQueue&#34;, QueueArgs.builder()        
 *             .projectId(example.id())
 *             .agentPoolId(examplePool.id())
 *             .build());
 * 
 *         var exampleCheckBranchControl = new CheckBranchControl(&#34;exampleCheckBranchControl&#34;, CheckBranchControlArgs.builder()        
 *             .projectId(example.id())
 *             .displayName(&#34;Managed by Terraform&#34;)
 *             .targetResourceId(exampleQueue.id())
 *             .targetResourceType(&#34;queue&#34;)
 *             .allowedBranches(&#34;refs/heads/main, refs/heads/features/*&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ### Protect a repository
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
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
 * import com.pulumi.azuredevops.CheckBranchControl;
 * import com.pulumi.azuredevops.CheckBranchControlArgs;
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
 *         var example = new Project(&#34;example&#34;, ProjectArgs.builder()        
 *             .name(&#34;Example Project&#34;)
 *             .build());
 * 
 *         var exampleGit = new Git(&#34;exampleGit&#34;, GitArgs.builder()        
 *             .projectId(example.id())
 *             .name(&#34;Example Empty Git Repository&#34;)
 *             .initialization(GitInitializationArgs.builder()
 *                 .initType(&#34;Clean&#34;)
 *                 .build())
 *             .build());
 * 
 *         var exampleCheckBranchControl = new CheckBranchControl(&#34;exampleCheckBranchControl&#34;, CheckBranchControlArgs.builder()        
 *             .projectId(example.id())
 *             .displayName(&#34;Managed by Terraform&#34;)
 *             .targetResourceId(Output.tuple(example.id(), exampleGit.id()).applyValue(values -&gt; {
 *                 var exampleId = values.t1;
 *                 var exampleGitId = values.t2;
 *                 return String.format(&#34;%s.%s&#34;, exampleId,exampleGitId);
 *             }))
 *             .targetResourceType(&#34;repository&#34;)
 *             .allowedBranches(&#34;refs/heads/main, refs/heads/features/*&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ### Protect a variable group
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.azuredevops.Project;
 * import com.pulumi.azuredevops.ProjectArgs;
 * import com.pulumi.azuredevops.VariableGroup;
 * import com.pulumi.azuredevops.VariableGroupArgs;
 * import com.pulumi.azuredevops.inputs.VariableGroupVariableArgs;
 * import com.pulumi.azuredevops.CheckBranchControl;
 * import com.pulumi.azuredevops.CheckBranchControlArgs;
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
 *         var example = new Project(&#34;example&#34;, ProjectArgs.builder()        
 *             .name(&#34;Example Project&#34;)
 *             .build());
 * 
 *         var exampleVariableGroup = new VariableGroup(&#34;exampleVariableGroup&#34;, VariableGroupArgs.builder()        
 *             .projectId(example.id())
 *             .name(&#34;Example Variable Group&#34;)
 *             .description(&#34;Example Variable Group Description&#34;)
 *             .allowAccess(true)
 *             .variables(            
 *                 VariableGroupVariableArgs.builder()
 *                     .name(&#34;key1&#34;)
 *                     .value(&#34;val1&#34;)
 *                     .build(),
 *                 VariableGroupVariableArgs.builder()
 *                     .name(&#34;key2&#34;)
 *                     .secretValue(&#34;val2&#34;)
 *                     .isSecret(true)
 *                     .build())
 *             .build());
 * 
 *         var exampleCheckBranchControl = new CheckBranchControl(&#34;exampleCheckBranchControl&#34;, CheckBranchControlArgs.builder()        
 *             .projectId(example.id())
 *             .displayName(&#34;Managed by Terraform&#34;)
 *             .targetResourceId(exampleVariableGroup.id())
 *             .targetResourceType(&#34;variablegroup&#34;)
 *             .allowedBranches(&#34;refs/heads/main, refs/heads/features/*&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Relevant Links
 * 
 * - [Define approvals and checks](https://learn.microsoft.com/en-us/azure/devops/pipelines/process/approvals?view=azure-devops&amp;tabs=check-pass)
 * 
 * ## Import
 * 
 * Importing this resource is not supported.
 * 
 */
@ResourceType(type="azuredevops:index/checkBranchControl:CheckBranchControl")
public class CheckBranchControl extends com.pulumi.resources.CustomResource {
    /**
     * The branches allowed to use the resource. Specify a comma separated list of allowed branches in `refs/heads/branch_name` format. To allow deployments from all branches, specify `*` . `refs/heads/features/* , refs/heads/releases/*` restricts deployments to all branches under features/ or releases/ . Defaults to `*`.
     * 
     */
    @Export(name="allowedBranches", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> allowedBranches;

    /**
     * @return The branches allowed to use the resource. Specify a comma separated list of allowed branches in `refs/heads/branch_name` format. To allow deployments from all branches, specify `*` . `refs/heads/features/* , refs/heads/releases/*` restricts deployments to all branches under features/ or releases/ . Defaults to `*`.
     * 
     */
    public Output<Optional<String>> allowedBranches() {
        return Codegen.optional(this.allowedBranches);
    }
    /**
     * The name of the branch control check displayed in the web UI.
     * 
     */
    @Export(name="displayName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> displayName;

    /**
     * @return The name of the branch control check displayed in the web UI.
     * 
     */
    public Output<Optional<String>> displayName() {
        return Codegen.optional(this.displayName);
    }
    /**
     * Allow deployment from branches for which protection status could not be obtained. Only relevant when verify_branch_protection is `true`. Defaults to `false`.
     * 
     */
    @Export(name="ignoreUnknownProtectionStatus", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> ignoreUnknownProtectionStatus;

    /**
     * @return Allow deployment from branches for which protection status could not be obtained. Only relevant when verify_branch_protection is `true`. Defaults to `false`.
     * 
     */
    public Output<Optional<Boolean>> ignoreUnknownProtectionStatus() {
        return Codegen.optional(this.ignoreUnknownProtectionStatus);
    }
    /**
     * The project ID.
     * 
     */
    @Export(name="projectId", refs={String.class}, tree="[0]")
    private Output<String> projectId;

    /**
     * @return The project ID.
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }
    /**
     * The ID of the resource being protected by the check.
     * 
     */
    @Export(name="targetResourceId", refs={String.class}, tree="[0]")
    private Output<String> targetResourceId;

    /**
     * @return The ID of the resource being protected by the check.
     * 
     */
    public Output<String> targetResourceId() {
        return this.targetResourceId;
    }
    /**
     * The type of resource being protected by the check. Valid values: `endpoint`, `environment`, `queue`, `repository`, `securefile`, `variablegroup`.
     * 
     */
    @Export(name="targetResourceType", refs={String.class}, tree="[0]")
    private Output<String> targetResourceType;

    /**
     * @return The type of resource being protected by the check. Valid values: `endpoint`, `environment`, `queue`, `repository`, `securefile`, `variablegroup`.
     * 
     */
    public Output<String> targetResourceType() {
        return this.targetResourceType;
    }
    /**
     * The timeout in minutes for the branch control check. Defaults to `1440`.
     * 
     */
    @Export(name="timeout", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> timeout;

    /**
     * @return The timeout in minutes for the branch control check. Defaults to `1440`.
     * 
     */
    public Output<Optional<Integer>> timeout() {
        return Codegen.optional(this.timeout);
    }
    /**
     * Validate the branches being deployed are protected. Defaults to `false`.
     * 
     */
    @Export(name="verifyBranchProtection", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> verifyBranchProtection;

    /**
     * @return Validate the branches being deployed are protected. Defaults to `false`.
     * 
     */
    public Output<Optional<Boolean>> verifyBranchProtection() {
        return Codegen.optional(this.verifyBranchProtection);
    }
    /**
     * The version of the check.
     * 
     */
    @Export(name="version", refs={Integer.class}, tree="[0]")
    private Output<Integer> version;

    /**
     * @return The version of the check.
     * 
     */
    public Output<Integer> version() {
        return this.version;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public CheckBranchControl(String name) {
        this(name, CheckBranchControlArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public CheckBranchControl(String name, CheckBranchControlArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public CheckBranchControl(String name, CheckBranchControlArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/checkBranchControl:CheckBranchControl", name, args == null ? CheckBranchControlArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private CheckBranchControl(String name, Output<String> id, @Nullable CheckBranchControlState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/checkBranchControl:CheckBranchControl", name, state, makeResourceOptions(options, id));
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
    public static CheckBranchControl get(String name, Output<String> id, @Nullable CheckBranchControlState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new CheckBranchControl(name, id, state, options);
    }
}
