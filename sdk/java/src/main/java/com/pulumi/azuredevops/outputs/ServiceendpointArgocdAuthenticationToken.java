// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ServiceendpointArgocdAuthenticationToken {
    /**
     * @return Authentication Token generated through ArgoCD.
     * 
     */
    private String token;
    private @Nullable String tokenHash;

    private ServiceendpointArgocdAuthenticationToken() {}
    /**
     * @return Authentication Token generated through ArgoCD.
     * 
     */
    public String token() {
        return this.token;
    }
    public Optional<String> tokenHash() {
        return Optional.ofNullable(this.tokenHash);
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
        private @Nullable String tokenHash;
        public Builder() {}
        public Builder(ServiceendpointArgocdAuthenticationToken defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.token = defaults.token;
    	      this.tokenHash = defaults.tokenHash;
        }

        @CustomType.Setter
        public Builder token(String token) {
            this.token = Objects.requireNonNull(token);
            return this;
        }
        @CustomType.Setter
        public Builder tokenHash(@Nullable String tokenHash) {
            this.tokenHash = tokenHash;
            return this;
        }
        public ServiceendpointArgocdAuthenticationToken build() {
            final var o = new ServiceendpointArgocdAuthenticationToken();
            o.token = token;
            o.tokenHash = tokenHash;
            return o;
        }
    }
}
