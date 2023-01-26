// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.inputs;

import com.pulumi.azuredevops.inputs.ServiceEndpointKubernetesAzureSubscriptionArgs;
import com.pulumi.azuredevops.inputs.ServiceEndpointKubernetesKubeconfigArgs;
import com.pulumi.azuredevops.inputs.ServiceEndpointKubernetesServiceAccountArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ServiceEndpointKubernetesState extends com.pulumi.resources.ResourceArgs {

    public static final ServiceEndpointKubernetesState Empty = new ServiceEndpointKubernetesState();

    /**
     * The hostname (in form of URI) of the Kubernetes API.
     * 
     */
    @Import(name="apiserverUrl")
    private @Nullable Output<String> apiserverUrl;

    /**
     * @return The hostname (in form of URI) of the Kubernetes API.
     * 
     */
    public Optional<Output<String>> apiserverUrl() {
        return Optional.ofNullable(this.apiserverUrl);
    }

    @Import(name="authorization")
    private @Nullable Output<Map<String,String>> authorization;

    public Optional<Output<Map<String,String>>> authorization() {
        return Optional.ofNullable(this.authorization);
    }

    /**
     * The authentication method used to authenticate on the Kubernetes cluster. The value should be one of AzureSubscription, Kubeconfig, ServiceAccount.
     * 
     */
    @Import(name="authorizationType")
    private @Nullable Output<String> authorizationType;

    /**
     * @return The authentication method used to authenticate on the Kubernetes cluster. The value should be one of AzureSubscription, Kubeconfig, ServiceAccount.
     * 
     */
    public Optional<Output<String>> authorizationType() {
        return Optional.ofNullable(this.authorizationType);
    }

    /**
     * A `azure_subscription` block defined blow.
     * 
     */
    @Import(name="azureSubscriptions")
    private @Nullable Output<List<ServiceEndpointKubernetesAzureSubscriptionArgs>> azureSubscriptions;

    /**
     * @return A `azure_subscription` block defined blow.
     * 
     */
    public Optional<Output<List<ServiceEndpointKubernetesAzureSubscriptionArgs>>> azureSubscriptions() {
        return Optional.ofNullable(this.azureSubscriptions);
    }

    @Import(name="description")
    private @Nullable Output<String> description;

    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * A `kubeconfig` block defined blow.
     * 
     */
    @Import(name="kubeconfigs")
    private @Nullable Output<List<ServiceEndpointKubernetesKubeconfigArgs>> kubeconfigs;

    /**
     * @return A `kubeconfig` block defined blow.
     * 
     */
    public Optional<Output<List<ServiceEndpointKubernetesKubeconfigArgs>>> kubeconfigs() {
        return Optional.ofNullable(this.kubeconfigs);
    }

    /**
     * The ID of the project.
     * 
     */
    @Import(name="projectId")
    private @Nullable Output<String> projectId;

    /**
     * @return The ID of the project.
     * 
     */
    public Optional<Output<String>> projectId() {
        return Optional.ofNullable(this.projectId);
    }

    /**
     * A `service_account` block defined blow.
     * 
     */
    @Import(name="serviceAccount")
    private @Nullable Output<ServiceEndpointKubernetesServiceAccountArgs> serviceAccount;

    /**
     * @return A `service_account` block defined blow.
     * 
     */
    public Optional<Output<ServiceEndpointKubernetesServiceAccountArgs>> serviceAccount() {
        return Optional.ofNullable(this.serviceAccount);
    }

    /**
     * The Service Endpoint name.
     * 
     */
    @Import(name="serviceEndpointName")
    private @Nullable Output<String> serviceEndpointName;

    /**
     * @return The Service Endpoint name.
     * 
     */
    public Optional<Output<String>> serviceEndpointName() {
        return Optional.ofNullable(this.serviceEndpointName);
    }

    private ServiceEndpointKubernetesState() {}

    private ServiceEndpointKubernetesState(ServiceEndpointKubernetesState $) {
        this.apiserverUrl = $.apiserverUrl;
        this.authorization = $.authorization;
        this.authorizationType = $.authorizationType;
        this.azureSubscriptions = $.azureSubscriptions;
        this.description = $.description;
        this.kubeconfigs = $.kubeconfigs;
        this.projectId = $.projectId;
        this.serviceAccount = $.serviceAccount;
        this.serviceEndpointName = $.serviceEndpointName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServiceEndpointKubernetesState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServiceEndpointKubernetesState $;

        public Builder() {
            $ = new ServiceEndpointKubernetesState();
        }

        public Builder(ServiceEndpointKubernetesState defaults) {
            $ = new ServiceEndpointKubernetesState(Objects.requireNonNull(defaults));
        }

        /**
         * @param apiserverUrl The hostname (in form of URI) of the Kubernetes API.
         * 
         * @return builder
         * 
         */
        public Builder apiserverUrl(@Nullable Output<String> apiserverUrl) {
            $.apiserverUrl = apiserverUrl;
            return this;
        }

        /**
         * @param apiserverUrl The hostname (in form of URI) of the Kubernetes API.
         * 
         * @return builder
         * 
         */
        public Builder apiserverUrl(String apiserverUrl) {
            return apiserverUrl(Output.of(apiserverUrl));
        }

        public Builder authorization(@Nullable Output<Map<String,String>> authorization) {
            $.authorization = authorization;
            return this;
        }

        public Builder authorization(Map<String,String> authorization) {
            return authorization(Output.of(authorization));
        }

        /**
         * @param authorizationType The authentication method used to authenticate on the Kubernetes cluster. The value should be one of AzureSubscription, Kubeconfig, ServiceAccount.
         * 
         * @return builder
         * 
         */
        public Builder authorizationType(@Nullable Output<String> authorizationType) {
            $.authorizationType = authorizationType;
            return this;
        }

        /**
         * @param authorizationType The authentication method used to authenticate on the Kubernetes cluster. The value should be one of AzureSubscription, Kubeconfig, ServiceAccount.
         * 
         * @return builder
         * 
         */
        public Builder authorizationType(String authorizationType) {
            return authorizationType(Output.of(authorizationType));
        }

        /**
         * @param azureSubscriptions A `azure_subscription` block defined blow.
         * 
         * @return builder
         * 
         */
        public Builder azureSubscriptions(@Nullable Output<List<ServiceEndpointKubernetesAzureSubscriptionArgs>> azureSubscriptions) {
            $.azureSubscriptions = azureSubscriptions;
            return this;
        }

        /**
         * @param azureSubscriptions A `azure_subscription` block defined blow.
         * 
         * @return builder
         * 
         */
        public Builder azureSubscriptions(List<ServiceEndpointKubernetesAzureSubscriptionArgs> azureSubscriptions) {
            return azureSubscriptions(Output.of(azureSubscriptions));
        }

        /**
         * @param azureSubscriptions A `azure_subscription` block defined blow.
         * 
         * @return builder
         * 
         */
        public Builder azureSubscriptions(ServiceEndpointKubernetesAzureSubscriptionArgs... azureSubscriptions) {
            return azureSubscriptions(List.of(azureSubscriptions));
        }

        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param kubeconfigs A `kubeconfig` block defined blow.
         * 
         * @return builder
         * 
         */
        public Builder kubeconfigs(@Nullable Output<List<ServiceEndpointKubernetesKubeconfigArgs>> kubeconfigs) {
            $.kubeconfigs = kubeconfigs;
            return this;
        }

        /**
         * @param kubeconfigs A `kubeconfig` block defined blow.
         * 
         * @return builder
         * 
         */
        public Builder kubeconfigs(List<ServiceEndpointKubernetesKubeconfigArgs> kubeconfigs) {
            return kubeconfigs(Output.of(kubeconfigs));
        }

        /**
         * @param kubeconfigs A `kubeconfig` block defined blow.
         * 
         * @return builder
         * 
         */
        public Builder kubeconfigs(ServiceEndpointKubernetesKubeconfigArgs... kubeconfigs) {
            return kubeconfigs(List.of(kubeconfigs));
        }

        /**
         * @param projectId The ID of the project.
         * 
         * @return builder
         * 
         */
        public Builder projectId(@Nullable Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param projectId The ID of the project.
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
        }

        /**
         * @param serviceAccount A `service_account` block defined blow.
         * 
         * @return builder
         * 
         */
        public Builder serviceAccount(@Nullable Output<ServiceEndpointKubernetesServiceAccountArgs> serviceAccount) {
            $.serviceAccount = serviceAccount;
            return this;
        }

        /**
         * @param serviceAccount A `service_account` block defined blow.
         * 
         * @return builder
         * 
         */
        public Builder serviceAccount(ServiceEndpointKubernetesServiceAccountArgs serviceAccount) {
            return serviceAccount(Output.of(serviceAccount));
        }

        /**
         * @param serviceEndpointName The Service Endpoint name.
         * 
         * @return builder
         * 
         */
        public Builder serviceEndpointName(@Nullable Output<String> serviceEndpointName) {
            $.serviceEndpointName = serviceEndpointName;
            return this;
        }

        /**
         * @param serviceEndpointName The Service Endpoint name.
         * 
         * @return builder
         * 
         */
        public Builder serviceEndpointName(String serviceEndpointName) {
            return serviceEndpointName(Output.of(serviceEndpointName));
        }

        public ServiceEndpointKubernetesState build() {
            return $;
        }
    }

}
