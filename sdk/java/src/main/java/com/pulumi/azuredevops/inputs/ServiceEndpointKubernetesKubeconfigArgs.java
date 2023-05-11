// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ServiceEndpointKubernetesKubeconfigArgs extends com.pulumi.resources.ResourceArgs {

    public static final ServiceEndpointKubernetesKubeconfigArgs Empty = new ServiceEndpointKubernetesKubeconfigArgs();

    /**
     * Set this option to allow clients to accept a self-signed certificate.
     * 
     */
    @Import(name="acceptUntrustedCerts")
    private @Nullable Output<Boolean> acceptUntrustedCerts;

    /**
     * @return Set this option to allow clients to accept a self-signed certificate.
     * 
     */
    public Optional<Output<Boolean>> acceptUntrustedCerts() {
        return Optional.ofNullable(this.acceptUntrustedCerts);
    }

    /**
     * Context within the kubeconfig file that is to be used for identifying the cluster. Default value is the current-context set in kubeconfig.
     * 
     */
    @Import(name="clusterContext")
    private @Nullable Output<String> clusterContext;

    /**
     * @return Context within the kubeconfig file that is to be used for identifying the cluster. Default value is the current-context set in kubeconfig.
     * 
     */
    public Optional<Output<String>> clusterContext() {
        return Optional.ofNullable(this.clusterContext);
    }

    /**
     * The content of the kubeconfig in yaml notation to be used to communicate with the API-Server of Kubernetes.
     * 
     */
    @Import(name="kubeConfig", required=true)
    private Output<String> kubeConfig;

    /**
     * @return The content of the kubeconfig in yaml notation to be used to communicate with the API-Server of Kubernetes.
     * 
     */
    public Output<String> kubeConfig() {
        return this.kubeConfig;
    }

    private ServiceEndpointKubernetesKubeconfigArgs() {}

    private ServiceEndpointKubernetesKubeconfigArgs(ServiceEndpointKubernetesKubeconfigArgs $) {
        this.acceptUntrustedCerts = $.acceptUntrustedCerts;
        this.clusterContext = $.clusterContext;
        this.kubeConfig = $.kubeConfig;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServiceEndpointKubernetesKubeconfigArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServiceEndpointKubernetesKubeconfigArgs $;

        public Builder() {
            $ = new ServiceEndpointKubernetesKubeconfigArgs();
        }

        public Builder(ServiceEndpointKubernetesKubeconfigArgs defaults) {
            $ = new ServiceEndpointKubernetesKubeconfigArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param acceptUntrustedCerts Set this option to allow clients to accept a self-signed certificate.
         * 
         * @return builder
         * 
         */
        public Builder acceptUntrustedCerts(@Nullable Output<Boolean> acceptUntrustedCerts) {
            $.acceptUntrustedCerts = acceptUntrustedCerts;
            return this;
        }

        /**
         * @param acceptUntrustedCerts Set this option to allow clients to accept a self-signed certificate.
         * 
         * @return builder
         * 
         */
        public Builder acceptUntrustedCerts(Boolean acceptUntrustedCerts) {
            return acceptUntrustedCerts(Output.of(acceptUntrustedCerts));
        }

        /**
         * @param clusterContext Context within the kubeconfig file that is to be used for identifying the cluster. Default value is the current-context set in kubeconfig.
         * 
         * @return builder
         * 
         */
        public Builder clusterContext(@Nullable Output<String> clusterContext) {
            $.clusterContext = clusterContext;
            return this;
        }

        /**
         * @param clusterContext Context within the kubeconfig file that is to be used for identifying the cluster. Default value is the current-context set in kubeconfig.
         * 
         * @return builder
         * 
         */
        public Builder clusterContext(String clusterContext) {
            return clusterContext(Output.of(clusterContext));
        }

        /**
         * @param kubeConfig The content of the kubeconfig in yaml notation to be used to communicate with the API-Server of Kubernetes.
         * 
         * @return builder
         * 
         */
        public Builder kubeConfig(Output<String> kubeConfig) {
            $.kubeConfig = kubeConfig;
            return this;
        }

        /**
         * @param kubeConfig The content of the kubeconfig in yaml notation to be used to communicate with the API-Server of Kubernetes.
         * 
         * @return builder
         * 
         */
        public Builder kubeConfig(String kubeConfig) {
            return kubeConfig(Output.of(kubeConfig));
        }

        public ServiceEndpointKubernetesKubeconfigArgs build() {
            $.kubeConfig = Objects.requireNonNull($.kubeConfig, "expected parameter 'kubeConfig' to be non-null");
            return $;
        }
    }

}