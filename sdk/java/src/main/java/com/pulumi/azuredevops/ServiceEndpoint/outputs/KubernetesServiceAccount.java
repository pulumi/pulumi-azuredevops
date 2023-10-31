// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.ServiceEndpoint.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class KubernetesServiceAccount {
    /**
     * @return The certificate from a Kubernetes secret object.
     * 
     */
    private String caCert;
    /**
     * @return The token from a Kubernetes secret object.
     * 
     */
    private String token;

    private KubernetesServiceAccount() {}
    /**
     * @return The certificate from a Kubernetes secret object.
     * 
     */
    public String caCert() {
        return this.caCert;
    }
    /**
     * @return The token from a Kubernetes secret object.
     * 
     */
    public String token() {
        return this.token;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(KubernetesServiceAccount defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String caCert;
        private String token;
        public Builder() {}
        public Builder(KubernetesServiceAccount defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.caCert = defaults.caCert;
    	      this.token = defaults.token;
        }

        @CustomType.Setter
        public Builder caCert(String caCert) {
            this.caCert = Objects.requireNonNull(caCert);
            return this;
        }
        @CustomType.Setter
        public Builder token(String token) {
            this.token = Objects.requireNonNull(token);
            return this;
        }
        public KubernetesServiceAccount build() {
            final var _resultValue = new KubernetesServiceAccount();
            _resultValue.caCert = caCert;
            _resultValue.token = token;
            return _resultValue;
        }
    }
}
