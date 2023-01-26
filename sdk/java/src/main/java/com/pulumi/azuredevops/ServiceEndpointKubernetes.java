// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops;

import com.pulumi.azuredevops.ServiceEndpointKubernetesArgs;
import com.pulumi.azuredevops.Utilities;
import com.pulumi.azuredevops.inputs.ServiceEndpointKubernetesState;
import com.pulumi.azuredevops.outputs.ServiceEndpointKubernetesAzureSubscription;
import com.pulumi.azuredevops.outputs.ServiceEndpointKubernetesKubeconfig;
import com.pulumi.azuredevops.outputs.ServiceEndpointKubernetesServiceAccount;
import com.pulumi.core.Alias;
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
 * ## Relevant Links
 * 
 * - [Azure DevOps Service REST API 6.0 - Endpoints](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-6.0)
 * 
 * ## Import
 * 
 * Azure DevOps Service Endpoint Kubernetes can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**
 * 
 * ```sh
 *  $ pulumi import azuredevops:index/serviceEndpointKubernetes:ServiceEndpointKubernetes example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
 * ```
 * 
 */
@ResourceType(type="azuredevops:index/serviceEndpointKubernetes:ServiceEndpointKubernetes")
public class ServiceEndpointKubernetes extends com.pulumi.resources.CustomResource {
    /**
     * The hostname (in form of URI) of the Kubernetes API.
     * 
     */
    @Export(name="apiserverUrl", type=String.class, parameters={})
    private Output<String> apiserverUrl;

    /**
     * @return The hostname (in form of URI) of the Kubernetes API.
     * 
     */
    public Output<String> apiserverUrl() {
        return this.apiserverUrl;
    }
    @Export(name="authorization", type=Map.class, parameters={String.class, String.class})
    private Output<Map<String,String>> authorization;

    public Output<Map<String,String>> authorization() {
        return this.authorization;
    }
    /**
     * The authentication method used to authenticate on the Kubernetes cluster. The value should be one of AzureSubscription, Kubeconfig, ServiceAccount.
     * 
     */
    @Export(name="authorizationType", type=String.class, parameters={})
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
    @Export(name="azureSubscriptions", type=List.class, parameters={ServiceEndpointKubernetesAzureSubscription.class})
    private Output</* @Nullable */ List<ServiceEndpointKubernetesAzureSubscription>> azureSubscriptions;

    /**
     * @return A `azure_subscription` block defined blow.
     * 
     */
    public Output<Optional<List<ServiceEndpointKubernetesAzureSubscription>>> azureSubscriptions() {
        return Codegen.optional(this.azureSubscriptions);
    }
    @Export(name="description", type=String.class, parameters={})
    private Output</* @Nullable */ String> description;

    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * A `kubeconfig` block defined blow.
     * 
     */
    @Export(name="kubeconfig", type=ServiceEndpointKubernetesKubeconfig.class, parameters={})
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
    @Export(name="projectId", type=String.class, parameters={})
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
    @Export(name="serviceAccount", type=ServiceEndpointKubernetesServiceAccount.class, parameters={})
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
    @Export(name="serviceEndpointName", type=String.class, parameters={})
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
            .aliases(List.of(
                Output.of(Alias.builder().type("azuredevops:ServiceEndpoint/kubernetes:Kubernetes").build())
            ))
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
