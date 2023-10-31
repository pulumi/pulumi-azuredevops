// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class ServiceEndpointGitHubAuthOauth {
    private String oauthConfigurationId;

    private ServiceEndpointGitHubAuthOauth() {}
    public String oauthConfigurationId() {
        return this.oauthConfigurationId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ServiceEndpointGitHubAuthOauth defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String oauthConfigurationId;
        public Builder() {}
        public Builder(ServiceEndpointGitHubAuthOauth defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.oauthConfigurationId = defaults.oauthConfigurationId;
        }

        @CustomType.Setter
        public Builder oauthConfigurationId(String oauthConfigurationId) {
            this.oauthConfigurationId = Objects.requireNonNull(oauthConfigurationId);
            return this;
        }
        public ServiceEndpointGitHubAuthOauth build() {
            final var _resultValue = new ServiceEndpointGitHubAuthOauth();
            _resultValue.oauthConfigurationId = oauthConfigurationId;
            return _resultValue;
        }
    }
}
