// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops;

import com.pulumi.azuredevops.PipelineAuthorizationArgs;
import com.pulumi.azuredevops.Utilities;
import com.pulumi.azuredevops.inputs.PipelineAuthorizationState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manage pipeline access permissions to resources.
 * 
 * &gt; **Note** This resource is a replacement for `azuredevops.ResourceAuthorization`.  Pipeline authorizations managed by `azuredevops.ResourceAuthorization` can also
 * be managed by this resource
 * 
 * ## Example Usage
 * ### Authorization for all pipelines
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
 * import com.pulumi.azuredevops.PipelineAuthorization;
 * import com.pulumi.azuredevops.PipelineAuthorizationArgs;
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
 *         var examplePool = new Pool(&#34;examplePool&#34;, PoolArgs.builder()        
 *             .autoProvision(false)
 *             .autoUpdate(false)
 *             .build());
 * 
 *         var exampleQueue = new Queue(&#34;exampleQueue&#34;, QueueArgs.builder()        
 *             .projectId(exampleProject.id())
 *             .agentPoolId(examplePool.id())
 *             .build());
 * 
 *         var examplePipelineAuthorization = new PipelineAuthorization(&#34;examplePipelineAuthorization&#34;, PipelineAuthorizationArgs.builder()        
 *             .projectId(exampleProject.id())
 *             .resourceId(exampleQueue.id())
 *             .type(&#34;queue&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Authorization for specific pipeline
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
 * import com.pulumi.azuredevops.AzuredevopsFunctions;
 * import com.pulumi.azuredevops.inputs.GetGitRepositoryArgs;
 * import com.pulumi.azuredevops.BuildDefinition;
 * import com.pulumi.azuredevops.BuildDefinitionArgs;
 * import com.pulumi.azuredevops.inputs.BuildDefinitionRepositoryArgs;
 * import com.pulumi.azuredevops.PipelineAuthorization;
 * import com.pulumi.azuredevops.PipelineAuthorizationArgs;
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
 *         var examplePool = new Pool(&#34;examplePool&#34;, PoolArgs.builder()        
 *             .autoProvision(false)
 *             .autoUpdate(false)
 *             .build());
 * 
 *         var exampleQueue = new Queue(&#34;exampleQueue&#34;, QueueArgs.builder()        
 *             .projectId(exampleProject.id())
 *             .agentPoolId(examplePool.id())
 *             .build());
 * 
 *         final var exampleGitRepository = AzuredevopsFunctions.getGitRepository(GetGitRepositoryArgs.builder()
 *             .projectId(exampleProject.id())
 *             .name(&#34;Example Project&#34;)
 *             .build());
 * 
 *         var exampleBuildDefinition = new BuildDefinition(&#34;exampleBuildDefinition&#34;, BuildDefinitionArgs.builder()        
 *             .projectId(exampleProject.id())
 *             .repository(BuildDefinitionRepositoryArgs.builder()
 *                 .repoType(&#34;TfsGit&#34;)
 *                 .repoId(exampleGitRepository.applyValue(getGitRepositoryResult -&gt; getGitRepositoryResult).applyValue(exampleGitRepository -&gt; exampleGitRepository.applyValue(getGitRepositoryResult -&gt; getGitRepositoryResult.id())))
 *                 .ymlPath(&#34;azure-pipelines.yml&#34;)
 *                 .build())
 *             .build());
 * 
 *         var examplePipelineAuthorization = new PipelineAuthorization(&#34;examplePipelineAuthorization&#34;, PipelineAuthorizationArgs.builder()        
 *             .projectId(exampleProject.id())
 *             .resourceId(exampleQueue.id())
 *             .type(&#34;queue&#34;)
 *             .pipelineId(exampleBuildDefinition.id())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ## Relevant Links
 * 
 * - [Azure DevOps Service REST API 7.1 - Pipeline Permissions](https://learn.microsoft.com/en-us/rest/api/azure/devops/approvalsandchecks/pipeline-permissions?view=azure-devops-rest-7.1)
 * 
 */
@ResourceType(type="azuredevops:index/pipelineAuthorization:PipelineAuthorization")
public class PipelineAuthorization extends com.pulumi.resources.CustomResource {
    /**
     * The ID of the pipeline. Changing this forces a new resource to be created
     * 
     */
    @Export(name="pipelineId", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> pipelineId;

    /**
     * @return The ID of the pipeline. Changing this forces a new resource to be created
     * 
     */
    public Output<Optional<Integer>> pipelineId() {
        return Codegen.optional(this.pipelineId);
    }
    /**
     * The  ID of the project. Changing this forces a new resource to be created
     * 
     */
    @Export(name="projectId", type=String.class, parameters={})
    private Output<String> projectId;

    /**
     * @return The  ID of the project. Changing this forces a new resource to be created
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }
    /**
     * The ID of the resource to authorize. Changing this forces a new resource to be created
     * 
     */
    @Export(name="resourceId", type=String.class, parameters={})
    private Output<String> resourceId;

    /**
     * @return The ID of the resource to authorize. Changing this forces a new resource to be created
     * 
     */
    public Output<String> resourceId() {
        return this.resourceId;
    }
    /**
     * The type of the resource to authorize. Valid values: `endpoint`, `queue`, `variablegroup`, `environment`. Changing this forces a new resource to be created
     * 
     */
    @Export(name="type", type=String.class, parameters={})
    private Output<String> type;

    /**
     * @return The type of the resource to authorize. Valid values: `endpoint`, `queue`, `variablegroup`, `environment`. Changing this forces a new resource to be created
     * 
     */
    public Output<String> type() {
        return this.type;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public PipelineAuthorization(String name) {
        this(name, PipelineAuthorizationArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public PipelineAuthorization(String name, PipelineAuthorizationArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public PipelineAuthorization(String name, PipelineAuthorizationArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/pipelineAuthorization:PipelineAuthorization", name, args == null ? PipelineAuthorizationArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private PipelineAuthorization(String name, Output<String> id, @Nullable PipelineAuthorizationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/pipelineAuthorization:PipelineAuthorization", name, state, makeResourceOptions(options, id));
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
    public static PipelineAuthorization get(String name, Output<String> id, @Nullable PipelineAuthorizationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new PipelineAuthorization(name, id, state, options);
    }
}
