// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class ServiceendpointJfrogArtifactoryV2AuthenticationBasic {
    /**
     * @return Artifactory Password.
     * 
     */
    private String password;
    /**
     * @return Artifactory Username.
     * 
     */
    private String username;

    private ServiceendpointJfrogArtifactoryV2AuthenticationBasic() {}
    /**
     * @return Artifactory Password.
     * 
     */
    public String password() {
        return this.password;
    }
    /**
     * @return Artifactory Username.
     * 
     */
    public String username() {
        return this.username;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ServiceendpointJfrogArtifactoryV2AuthenticationBasic defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String password;
        private String username;
        public Builder() {}
        public Builder(ServiceendpointJfrogArtifactoryV2AuthenticationBasic defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.password = defaults.password;
    	      this.username = defaults.username;
        }

        @CustomType.Setter
        public Builder password(String password) {
            this.password = Objects.requireNonNull(password);
            return this;
        }
        @CustomType.Setter
        public Builder username(String username) {
            this.username = Objects.requireNonNull(username);
            return this;
        }
        public ServiceendpointJfrogArtifactoryV2AuthenticationBasic build() {
            final var _resultValue = new ServiceendpointJfrogArtifactoryV2AuthenticationBasic();
            _resultValue.password = password;
            _resultValue.username = username;
            return _resultValue;
        }
    }
}
