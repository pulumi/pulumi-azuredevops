// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops;

import com.pulumi.azuredevops.ServiceEndpointKubernetesArgs;
import com.pulumi.azuredevops.Utilities;
import com.pulumi.azuredevops.inputs.ServiceEndpointKubernetesState;
import com.pulumi.azuredevops.outputs.ServiceEndpointKubernetesAzureSubscription;
import com.pulumi.azuredevops.outputs.ServiceEndpointKubernetesKubeconfig;
import com.pulumi.azuredevops.outputs.ServiceEndpointKubernetesServiceAccount;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a Kubernetes service endpoint within Azure DevOps.
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
 * import com.pulumi.azuredevops.ServiceEndpointKubernetes;
 * import com.pulumi.azuredevops.ServiceEndpointKubernetesArgs;
 * import com.pulumi.azuredevops.inputs.ServiceEndpointKubernetesAzureSubscriptionArgs;
 * import com.pulumi.azuredevops.inputs.ServiceEndpointKubernetesKubeconfigArgs;
 * import com.pulumi.azuredevops.inputs.ServiceEndpointKubernetesServiceAccountArgs;
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
 *             .visibility("private")
 *             .versionControl("Git")
 *             .workItemTemplate("Agile")
 *             .description("Managed by Terraform")
 *             .build());
 * 
 *         var example_azure = new ServiceEndpointKubernetes("example-azure", ServiceEndpointKubernetesArgs.builder()        
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
 *         var example_kubeconfig = new ServiceEndpointKubernetes("example-kubeconfig", ServiceEndpointKubernetesArgs.builder()        
 *             .projectId(example.id())
 *             .serviceEndpointName("Example Kubernetes")
 *             .apiserverUrl("https://sample-kubernetes-cluster.hcp.westeurope.azmk8s.io")
 *             .authorizationType("Kubeconfig")
 *             .kubeconfig(ServiceEndpointKubernetesKubeconfigArgs.builder()
 *                 .kubeConfig("""
 *                               apiVersion: v1
 *                               clusters:
 *                               - cluster:
 *                                   certificate-authority: fake-ca-file
 *                                   server: https://1.2.3.4
 *                                 name: development
 *                               contexts:
 *                               - context:
 *                                   cluster: development
 *                                   namespace: frontend
 *                                   user: developer
 *                                 name: dev-frontend
 *                               current-context: dev-frontend
 *                               kind: Config
 *                               preferences: {}
 *                               users:
 *                               - name: developer
 *                                 user:
 *                                   client-certificate: fake-cert-file
 *                                   client-key: fake-key-file
 *                 """)
 *                 .acceptUntrustedCerts(true)
 *                 .clusterContext("dev-frontend")
 *                 .build())
 *             .build());
 * 
 *         var example_service_account = new ServiceEndpointKubernetes("example-service-account", ServiceEndpointKubernetesArgs.builder()        
 *             .projectId(example.id())
 *             .serviceEndpointName("Example Kubernetes")
 *             .apiserverUrl("https://sample-kubernetes-cluster.hcp.westeurope.azmk8s.io")
 *             .authorizationType("ServiceAccount")
 *             .serviceAccount(ServiceEndpointKubernetesServiceAccountArgs.builder()
 *                 .token("000000000000000000000000")
 *                 .caCert("0000000000000000000000000000000")
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
 * - [Azure DevOps Service REST API 7.0 - Endpoints](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-7.0)
 * 
 * ## Import
 * 
 * Azure DevOps Service Endpoint Kubernetes can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**
 * 
 * ```sh
 * $ pulumi import azuredevops:index/serviceEndpointKubernetes:ServiceEndpointKubernetes example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
 * ```
 * 
 */
@ResourceType(type="azuredevops:index/serviceEndpointKubernetes:ServiceEndpointKubernetes")
public class ServiceEndpointKubernetes extends com.pulumi.resources.CustomResource {
    /**
     * The hostname (in form of URI) of the Kubernetes API.
     * 
     */
    @Export(name="apiserverUrl", refs={String.class}, tree="[0]")
    private Output<String> apiserverUrl;

    /**
     * @return The hostname (in form of URI) of the Kubernetes API.
     * 
     */
    public Output<String> apiserverUrl() {
        return this.apiserverUrl;
    }
    @Export(name="authorization", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> authorization;

    public Output<Map<String,String>> authorization() {
        return this.authorization;
    }
    /**
     * The authentication method used to authenticate on the Kubernetes cluster. The value should be one of AzureSubscription, Kubeconfig, ServiceAccount.
     * 
     */
    @Export(name="authorizationType", refs={String.class}, tree="[0]")
    private Output<String> authorizationType;

    /**
     * @return The authentication method used to authenticate on the Kubernetes cluster. The value should be one of AzureSubscription, Kubeconfig, ServiceAccount.
     * 
     */
    public Output<String> authorizationType() {
        return this.authorizationType;
    }
    /**
     * A `azure_subscription` block defined blow.
     * 
     */
    @Export(name="azureSubscriptions", refs={List.class,ServiceEndpointKubernetesAzureSubscription.class}, tree="[0,1]")
    private Output</* @Nullable */ List<ServiceEndpointKubernetesAzureSubscription>> azureSubscriptions;

    /**
     * @return A `azure_subscription` block defined blow.
     * 
     */
    public Output<Optional<List<ServiceEndpointKubernetesAzureSubscription>>> azureSubscriptions() {
        return Codegen.optional(this.azureSubscriptions);
    }
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * A `kubeconfig` block defined blow.
     * 
     */
    @Export(name="kubeconfig", refs={ServiceEndpointKubernetesKubeconfig.class}, tree="[0]")
    private Output</* @Nullable */ ServiceEndpointKubernetesKubeconfig> kubeconfig;

    /**
     * @return A `kubeconfig` block defined blow.
     * 
     */
    public Output<Optional<ServiceEndpointKubernetesKubeconfig>> kubeconfig() {
        return Codegen.optional(this.kubeconfig);
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
     * A `service_account` block defined blow.
     * 
     */
    @Export(name="serviceAccount", refs={ServiceEndpointKubernetesServiceAccount.class}, tree="[0]")
    private Output</* @Nullable */ ServiceEndpointKubernetesServiceAccount> serviceAccount;

    /**
     * @return A `service_account` block defined blow.
     * 
     */
    public Output<Optional<ServiceEndpointKubernetesServiceAccount>> serviceAccount() {
        return Codegen.optional(this.serviceAccount);
    }
    /**
     * The Service Endpoint name.
     * 
     */
    @Export(name="serviceEndpointName", refs={String.class}, tree="[0]")
    private Output<String> serviceEndpointName;

    /**
     * @return The Service Endpoint name.
     * 
     */
    public Output<String> serviceEndpointName() {
        return this.serviceEndpointName;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ServiceEndpointKubernetes(String name) {
        this(name, ServiceEndpointKubernetesArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ServiceEndpointKubernetes(String name, ServiceEndpointKubernetesArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ServiceEndpointKubernetes(String name, ServiceEndpointKubernetesArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/serviceEndpointKubernetes:ServiceEndpointKubernetes", name, args == null ? ServiceEndpointKubernetesArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ServiceEndpointKubernetes(String name, Output<String> id, @Nullable ServiceEndpointKubernetesState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/serviceEndpointKubernetes:ServiceEndpointKubernetes", name, state, makeResourceOptions(options, id));
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
    public static ServiceEndpointKubernetes get(String name, Output<String> id, @Nullable ServiceEndpointKubernetesState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ServiceEndpointKubernetes(name, id, state, options);
    }
}
