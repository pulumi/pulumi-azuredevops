// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class ServiceendpointArgocdAuthenticationToken {
    /**
     * @return The ArgoCD access token.
     * 
     */
    private String token;

    private ServiceendpointArgocdAuthenticationToken() {}
    /**
     * @return The ArgoCD access token.
     * 
     */
    public String token() {
        return this.token;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ServiceendpointArgocdAuthenticationToken defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String token;
        public Builder() {}
        public Builder(ServiceendpointArgocdAuthenticationToken defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.token = defaults.token;
        }

        @CustomType.Setter
        public Builder token(String token) {
            if (token == null) {
              throw new MissingRequiredPropertyException("ServiceendpointArgocdAuthenticationToken", "token");
            }
            this.token = token;
            return this;
        }
        public ServiceendpointArgocdAuthenticationToken build() {
            final var _resultValue = new ServiceendpointArgocdAuthenticationToken();
            _resultValue.token = token;
            return _resultValue;
        }
    }
}
