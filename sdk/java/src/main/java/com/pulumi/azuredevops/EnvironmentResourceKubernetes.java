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
 * <pre>
 * {@code
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
 *         var example = new Project("example", ProjectArgs.builder()
 *             .name("Example Project")
 *             .workItemTemplate("Agile")
 *             .versionControl("Git")
 *             .visibility("private")
 *             .description("Managed by Terraform")
 *             .build());
 * 
 *         var exampleEnvironment = new Environment("exampleEnvironment", EnvironmentArgs.builder()
 *             .projectId(example.id())
 *             .name("Example Environment")
 *             .build());
 * 
 *         var exampleServiceEndpointKubernetes = new ServiceEndpointKubernetes("exampleServiceEndpointKubernetes", ServiceEndpointKubernetesArgs.builder()
 *             .projectId(example.id())
 *             .serviceEndpointName("Example Kubernetes")
 *             .apiserverUrl("https://sample-kubernetes-cluster.hcp.westeurope.azmk8s.io")
 *             .authorizationType("AzureSubscription")
 *             .azureSubscriptions(ServiceEndpointKubernetesAzureSubscriptionArgs.builder()
 *                 .subscriptionId("00000000-0000-0000-0000-000000000000")
 *                 .subscriptionName("Example")
 *                 .tenantId("00000000-0000-0000-0000-000000000000")
 *                 .resourcegroupId("example-rg")
 *                 .namespace("default")
 *                 .clusterName("example-aks")
 *                 .build())
 *             .build());
 * 
 *         var exampleEnvironmentResourceKubernetes = new EnvironmentResourceKubernetes("exampleEnvironmentResourceKubernetes", EnvironmentResourceKubernetesArgs.builder()
 *             .projectId(example.id())
 *             .environmentId(exampleEnvironment.id())
 *             .serviceEndpointId(exampleServiceEndpointKubernetes.id())
 *             .name("Example")
 *             .namespace("default")
 *             .clusterName("example-aks")
 *             .tags(            
 *                 "tag1",
 *                 "tag2")
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
    public EnvironmentResourceKubernetes(java.lang.String name) {
        this(name, EnvironmentResourceKubernetesArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public EnvironmentResourceKubernetes(java.lang.String name, EnvironmentResourceKubernetesArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public EnvironmentResourceKubernetes(java.lang.String name, EnvironmentResourceKubernetesArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/environmentResourceKubernetes:EnvironmentResourceKubernetes", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private EnvironmentResourceKubernetes(java.lang.String name, Output<java.lang.String> id, @Nullable EnvironmentResourceKubernetesState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/environmentResourceKubernetes:EnvironmentResourceKubernetes", name, state, makeResourceOptions(options, id), false);
    }

    private static EnvironmentResourceKubernetesArgs makeArgs(EnvironmentResourceKubernetesArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? EnvironmentResourceKubernetesArgs.Empty : args;
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
    public static EnvironmentResourceKubernetes get(java.lang.String name, Output<java.lang.String> id, @Nullable EnvironmentResourceKubernetesState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new EnvironmentResourceKubernetes(name, id, state, options);
    }
}
