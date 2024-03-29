// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class ServiceendpointMavenAuthenticationBasic {
    /**
     * @return The password Maven Repository.
     * 
     */
    private String password;
    /**
     * @return The Username of the Maven Repository.
     * 
     */
    private String username;

    private ServiceendpointMavenAuthenticationBasic() {}
    /**
     * @return The password Maven Repository.
     * 
     */
    public String password() {
        return this.password;
    }
    /**
     * @return The Username of the Maven Repository.
     * 
     */
    public String username() {
        return this.username;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ServiceendpointMavenAuthenticationBasic defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String password;
        private String username;
        public Builder() {}
        public Builder(ServiceendpointMavenAuthenticationBasic defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.password = defaults.password;
    	      this.username = defaults.username;
        }

        @CustomType.Setter
        public Builder password(String password) {
            if (password == null) {
              throw new MissingRequiredPropertyException("ServiceendpointMavenAuthenticationBasic", "password");
            }
            this.password = password;
            return this;
        }
        @CustomType.Setter
        public Builder username(String username) {
            if (username == null) {
              throw new MissingRequiredPropertyException("ServiceendpointMavenAuthenticationBasic", "username");
            }
            this.username = username;
            return this;
        }
        public ServiceendpointMavenAuthenticationBasic build() {
            final var _resultValue = new ServiceendpointMavenAuthenticationBasic();
            _resultValue.password = password;
            _resultValue.username = username;
            return _resultValue;
        }
    }
}
