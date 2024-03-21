// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops;

import com.pulumi.azuredevops.EnvironmentResourceKubernetesArgs;
import com.pulumi.azuredevops.Utilities;
import com.pulumi.azuredevops.inputs.EnvironmentResourceKubernetesState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a Kubernetes Resource for an Environment.
 * 
 * ## Example Usage
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
 * import com.pulumi.azuredevops.ServiceEndpointKubernetes;
 * import com.pulumi.azuredevops.ServiceEndpointKubernetesArgs;
 * import com.pulumi.azuredevops.inputs.ServiceEndpointKubernetesAzureSubscriptionArgs;
 * import com.pulumi.azuredevops.EnvironmentResourceKubernetes;
 * import com.pulumi.azuredevops.EnvironmentResourceKubernetesArgs;
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
 *             .workItemTemplate(&#34;Agile&#34;)
 *             .versionControl(&#34;Git&#34;)
 *             .visibility(&#34;private&#34;)
 *             .description(&#34;Managed by Terraform&#34;)
 *             .build());
 * 
 *         var exampleEnvironment = new Environment(&#34;exampleEnvironment&#34;, EnvironmentArgs.builder()        
 *             .projectId(exampleProject.id())
 *             .build());
 * 
 *         var exampleServiceEndpointKubernetes = new ServiceEndpointKubernetes(&#34;exampleServiceEndpointKubernetes&#34;, ServiceEndpointKubernetesArgs.builder()        
 *             .projectId(exampleProject.id())
 *             .serviceEndpointName(&#34;Example Kubernetes&#34;)
 *             .apiserverUrl(&#34;https://sample-kubernetes-cluster.hcp.westeurope.azmk8s.io&#34;)
 *             .authorizationType(&#34;AzureSubscription&#34;)
 *             .azureSubscriptions(ServiceEndpointKubernetesAzureSubscriptionArgs.builder()
 *                 .subscriptionId(&#34;00000000-0000-0000-0000-000000000000&#34;)
 *                 .subscriptionName(&#34;Example&#34;)
 *                 .tenantId(&#34;00000000-0000-0000-0000-000000000000&#34;)
 *                 .resourcegroupId(&#34;example-rg&#34;)
 *                 .namespace(&#34;default&#34;)
 *                 .clusterName(&#34;example-aks&#34;)
 *                 .build())
 *             .build());
 * 
 *         var exampleEnvironmentResourceKubernetes = new EnvironmentResourceKubernetes(&#34;exampleEnvironmentResourceKubernetes&#34;, EnvironmentResourceKubernetesArgs.builder()        
 *             .projectId(exampleProject.id())
 *             .environmentId(exampleEnvironment.id())
 *             .serviceEndpointId(exampleServiceEndpointKubernetes.id())
 *             .namespace(&#34;default&#34;)
 *             .clusterName(&#34;example-aks&#34;)
 *             .tags(            
 *                 &#34;tag1&#34;,
 *                 &#34;tag2&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Relevant Links
 * 
 * * [Azure DevOps Service REST API 6.0 - Kubernetes](https://docs.microsoft.com/en-us/rest/api/azure/devops/distributedtask/kubernetes?view=azure-devops-rest-6.0)
 * 
 * ## Import
 * 
 * The resource does not support import.
 * 
 */
@ResourceType(type="azuredevops:index/environmentResourceKubernetes:EnvironmentResourceKubernetes")
public class EnvironmentResourceKubernetes extends com.pulumi.resources.CustomResource {
    /**
     * A cluster name for the Kubernetes Resource.
     * 
     */
    @Export(name="clusterName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> clusterName;

    /**
     * @return A cluster name for the Kubernetes Resource.
     * 
     */
    public Output<Optional<String>> clusterName() {
        return Codegen.optional(this.clusterName);
    }
    /**
     * The ID of the environment under which to create the Kubernetes Resource.
     * 
     */
    @Export(name="environmentId", refs={Integer.class}, tree="[0]")
    private Output<Integer> environmentId;

    /**
     * @return The ID of the environment under which to create the Kubernetes Resource.
     * 
     */
    public Output<Integer> environmentId() {
        return this.environmentId;
    }
    /**
     * The name for the Kubernetes Resource.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name for the Kubernetes Resource.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The namespace for the Kubernetes Resource.
     * 
     */
    @Export(name="namespace", refs={String.class}, tree="[0]")
    private Output<String> namespace;

    /**
     * @return The namespace for the Kubernetes Resource.
     * 
     */
    public Output<String> namespace() {
        return this.namespace;
    }
    /**
     * The ID of the project.
     * 
     */
    @Export(name="projectId", refs={String.class}, tree="[0]")
    private Output<String> projectId;

    /**
     * @return The ID of the project.
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }
    /**
     * The ID of the service endpoint to associate with the Kubernetes Resource.
     * 
     */
    @Export(name="serviceEndpointId", refs={String.class}, tree="[0]")
    private Output<String> serviceEndpointId;

    /**
     * @return The ID of the service endpoint to associate with the Kubernetes Resource.
     * 
     */
    public Output<String> serviceEndpointId() {
        return this.serviceEndpointId;
    }
    /**
     * A set of tags for the Kubernetes Resource.
     * 
     */
    @Export(name="tags", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> tags;

    /**
     * @return A set of tags for the Kubernetes Resource.
     * 
     */
    public Output<Optional<List<String>>> tags() {
        return Codegen.optional(this.tags);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public EnvironmentResourceKubernetes(String name) {
        this(name, EnvironmentResourceKubernetesArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public EnvironmentResourceKubernetes(String name, EnvironmentResourceKubernetesArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public EnvironmentResourceKubernetes(String name, EnvironmentResourceKubernetesArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/environmentResourceKubernetes:EnvironmentResourceKubernetes", name, args == null ? EnvironmentResourceKubernetesArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private EnvironmentResourceKubernetes(String name, Output<String> id, @Nullable EnvironmentResourceKubernetesState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/environmentResourceKubernetes:EnvironmentResourceKubernetes", name, state, makeResourceOptions(options, id));
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
    public static EnvironmentResourceKubernetes get(String name, Output<String> id, @Nullable EnvironmentResourceKubernetesState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new EnvironmentResourceKubernetes(name, id, state, options);
    }
}
