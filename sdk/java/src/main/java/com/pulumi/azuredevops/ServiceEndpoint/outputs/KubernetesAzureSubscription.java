// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.ServiceEndpoint.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class KubernetesAzureSubscription {
    /**
     * @return Azure environment refers to whether the public cloud offering or domestic (government) clouds are being used. Currently, only the public cloud is supported. The value must be AzureCloud. This is also the default-value.
     * 
     */
    private @Nullable String azureEnvironment;
    /**
     * @return Set this option to allow use cluster admin credentials.
     * 
     */
    private @Nullable Boolean clusterAdmin;
    /**
     * @return The name of the Kubernetes cluster.
     * 
     */
    private String clusterName;
    /**
     * @return The Kubernetes namespace. Default value is &#34;default&#34;.
     * 
     */
    private @Nullable String namespace;
    /**
     * @return The resource group name, to which the Kubernetes cluster is deployed.
     * 
     */
    private String resourcegroupId;
    /**
     * @return The id of the Azure subscription.
     * 
     */
    private String subscriptionId;
    /**
     * @return The name of the Azure subscription.
     * 
     */
    private String subscriptionName;
    /**
     * @return The id of the tenant used by the subscription.
     * 
     */
    private String tenantId;

    private KubernetesAzureSubscription() {}
    /**
     * @return Azure environment refers to whether the public cloud offering or domestic (government) clouds are being used. Currently, only the public cloud is supported. The value must be AzureCloud. This is also the default-value.
     * 
     */
    public Optional<String> azureEnvironment() {
        return Optional.ofNullable(this.azureEnvironment);
    }
    /**
     * @return Set this option to allow use cluster admin credentials.
     * 
     */
    public Optional<Boolean> clusterAdmin() {
        return Optional.ofNullable(this.clusterAdmin);
    }
    /**
     * @return The name of the Kubernetes cluster.
     * 
     */
    public String clusterName() {
        return this.clusterName;
    }
    /**
     * @return The Kubernetes namespace. Default value is &#34;default&#34;.
     * 
     */
    public Optional<String> namespace() {
        return Optional.ofNullable(this.namespace);
    }
    /**
     * @return The resource group name, to which the Kubernetes cluster is deployed.
     * 
     */
    public String resourcegroupId() {
        return this.resourcegroupId;
    }
    /**
     * @return The id of the Azure subscription.
     * 
     */
    public String subscriptionId() {
        return this.subscriptionId;
    }
    /**
     * @return The name of the Azure subscription.
     * 
     */
    public String subscriptionName() {
        return this.subscriptionName;
    }
    /**
     * @return The id of the tenant used by the subscription.
     * 
     */
    public String tenantId() {
        return this.tenantId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(KubernetesAzureSubscription defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String azureEnvironment;
        private @Nullable Boolean clusterAdmin;
        private String clusterName;
        private @Nullable String namespace;
        private String resourcegroupId;
        private String subscriptionId;
        private String subscriptionName;
        private String tenantId;
        public Builder() {}
        public Builder(KubernetesAzureSubscription defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.azureEnvironment = defaults.azureEnvironment;
    	      this.clusterAdmin = defaults.clusterAdmin;
    	      this.clusterName = defaults.clusterName;
    	      this.namespace = defaults.namespace;
    	      this.resourcegroupId = defaults.resourcegroupId;
    	      this.subscriptionId = defaults.subscriptionId;
    	      this.subscriptionName = defaults.subscriptionName;
    	      this.tenantId = defaults.tenantId;
        }

        @CustomType.Setter
        public Builder azureEnvironment(@Nullable String azureEnvironment) {
            this.azureEnvironment = azureEnvironment;
            return this;
        }
        @CustomType.Setter
        public Builder clusterAdmin(@Nullable Boolean clusterAdmin) {
            this.clusterAdmin = clusterAdmin;
            return this;
        }
        @CustomType.Setter
        public Builder clusterName(String clusterName) {
            this.clusterName = Objects.requireNonNull(clusterName);
            return this;
        }
        @CustomType.Setter
        public Builder namespace(@Nullable String namespace) {
            this.namespace = namespace;
            return this;
        }
        @CustomType.Setter
        public Builder resourcegroupId(String resourcegroupId) {
            this.resourcegroupId = Objects.requireNonNull(resourcegroupId);
            return this;
        }
        @CustomType.Setter
        public Builder subscriptionId(String subscriptionId) {
            this.subscriptionId = Objects.requireNonNull(subscriptionId);
            return this;
        }
        @CustomType.Setter
        public Builder subscriptionName(String subscriptionName) {
            this.subscriptionName = Objects.requireNonNull(subscriptionName);
            return this;
        }
        @CustomType.Setter
        public Builder tenantId(String tenantId) {
            this.tenantId = Objects.requireNonNull(tenantId);
            return this;
        }
        public KubernetesAzureSubscription build() {
            final var o = new KubernetesAzureSubscription();
            o.azureEnvironment = azureEnvironment;
            o.clusterAdmin = clusterAdmin;
            o.clusterName = clusterName;
            o.namespace = namespace;
            o.resourcegroupId = resourcegroupId;
            o.subscriptionId = subscriptionId;
            o.subscriptionName = subscriptionName;
            o.tenantId = tenantId;
            return o;
        }
    }
}