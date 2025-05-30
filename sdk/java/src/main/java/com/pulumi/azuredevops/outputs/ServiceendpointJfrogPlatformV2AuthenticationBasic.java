// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class ServiceendpointJfrogPlatformV2AuthenticationBasic {
    /**
     * @return The Password of the Artifactory.
     * 
     */
    private String password;
    /**
     * @return The Username of the  Artifactory.
     * 
     */
    private String username;

    private ServiceendpointJfrogPlatformV2AuthenticationBasic() {}
    /**
     * @return The Password of the Artifactory.
     * 
     */
    public String password() {
        return this.password;
    }
    /**
     * @return The Username of the  Artifactory.
     * 
     */
    public String username() {
        return this.username;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ServiceendpointJfrogPlatformV2AuthenticationBasic defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String password;
        private String username;
        public Builder() {}
        public Builder(ServiceendpointJfrogPlatformV2AuthenticationBasic defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.password = defaults.password;
    	      this.username = defaults.username;
        }

        @CustomType.Setter
        public Builder password(String password) {
            if (password == null) {
              throw new MissingRequiredPropertyException("ServiceendpointJfrogPlatformV2AuthenticationBasic", "password");
            }
            this.password = password;
            return this;
        }
        @CustomType.Setter
        public Builder username(String username) {
            if (username == null) {
              throw new MissingRequiredPropertyException("ServiceendpointJfrogPlatformV2AuthenticationBasic", "username");
            }
            this.username = username;
            return this;
        }
        public ServiceendpointJfrogPlatformV2AuthenticationBasic build() {
            final var _resultValue = new ServiceendpointJfrogPlatformV2AuthenticationBasic();
            _resultValue.password = password;
            _resultValue.username = username;
            return _resultValue;
        }
    }
}
