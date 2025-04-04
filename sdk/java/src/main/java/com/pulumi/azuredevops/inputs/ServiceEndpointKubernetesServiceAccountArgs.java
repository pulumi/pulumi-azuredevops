// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ServiceEndpointKubernetesServiceAccountArgs extends com.pulumi.resources.ResourceArgs {

    public static final ServiceEndpointKubernetesServiceAccountArgs Empty = new ServiceEndpointKubernetesServiceAccountArgs();

    /**
     * Set this option to allow clients to accept a self-signed certificate. Defaults to `false`.
     * 
     */
    @Import(name="acceptUntrustedCerts")
    private @Nullable Output<Boolean> acceptUntrustedCerts;

    /**
     * @return Set this option to allow clients to accept a self-signed certificate. Defaults to `false`.
     * 
     */
    public Optional<Output<Boolean>> acceptUntrustedCerts() {
        return Optional.ofNullable(this.acceptUntrustedCerts);
    }

    /**
     * The certificate from a Kubernetes secret object.
     * 
     */
    @Import(name="caCert", required=true)
    private Output<String> caCert;

    /**
     * @return The certificate from a Kubernetes secret object.
     * 
     */
    public Output<String> caCert() {
        return this.caCert;
    }

    /**
     * The token from a Kubernetes secret object.
     * 
     */
    @Import(name="token", required=true)
    private Output<String> token;

    /**
     * @return The token from a Kubernetes secret object.
     * 
     */
    public Output<String> token() {
        return this.token;
    }

    private ServiceEndpointKubernetesServiceAccountArgs() {}

    private ServiceEndpointKubernetesServiceAccountArgs(ServiceEndpointKubernetesServiceAccountArgs $) {
        this.acceptUntrustedCerts = $.acceptUntrustedCerts;
        this.caCert = $.caCert;
        this.token = $.token;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServiceEndpointKubernetesServiceAccountArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServiceEndpointKubernetesServiceAccountArgs $;

        public Builder() {
            $ = new ServiceEndpointKubernetesServiceAccountArgs();
        }

        public Builder(ServiceEndpointKubernetesServiceAccountArgs defaults) {
            $ = new ServiceEndpointKubernetesServiceAccountArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param acceptUntrustedCerts Set this option to allow clients to accept a self-signed certificate. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder acceptUntrustedCerts(@Nullable Output<Boolean> acceptUntrustedCerts) {
            $.acceptUntrustedCerts = acceptUntrustedCerts;
            return this;
        }

        /**
         * @param acceptUntrustedCerts Set this option to allow clients to accept a self-signed certificate. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder acceptUntrustedCerts(Boolean acceptUntrustedCerts) {
            return acceptUntrustedCerts(Output.of(acceptUntrustedCerts));
        }

        /**
         * @param caCert The certificate from a Kubernetes secret object.
         * 
         * @return builder
         * 
         */
        public Builder caCert(Output<String> caCert) {
            $.caCert = caCert;
            return this;
        }

        /**
         * @param caCert The certificate from a Kubernetes secret object.
         * 
         * @return builder
         * 
         */
        public Builder caCert(String caCert) {
            return caCert(Output.of(caCert));
        }

        /**
         * @param token The token from a Kubernetes secret object.
         * 
         * @return builder
         * 
         */
        public Builder token(Output<String> token) {
            $.token = token;
            return this;
        }

        /**
         * @param token The token from a Kubernetes secret object.
         * 
         * @return builder
         * 
         */
        public Builder token(String token) {
            return token(Output.of(token));
        }

        public ServiceEndpointKubernetesServiceAccountArgs build() {
            if ($.caCert == null) {
                throw new MissingRequiredPropertyException("ServiceEndpointKubernetesServiceAccountArgs", "caCert");
            }
            if ($.token == null) {
                throw new MissingRequiredPropertyException("ServiceEndpointKubernetesServiceAccountArgs", "token");
            }
            return $;
        }
    }

}
