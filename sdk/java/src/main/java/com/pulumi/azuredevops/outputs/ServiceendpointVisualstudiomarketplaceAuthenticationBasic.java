// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class ServiceendpointVisualstudiomarketplaceAuthenticationBasic {
    /**
     * @return The password of the marketplace.
     * 
     */
    private String password;
    /**
     * @return The username of the marketplace.
     * 
     */
    private String username;

    private ServiceendpointVisualstudiomarketplaceAuthenticationBasic() {}
    /**
     * @return The password of the marketplace.
     * 
     */
    public String password() {
        return this.password;
    }
    /**
     * @return The username of the marketplace.
     * 
     */
    public String username() {
        return this.username;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ServiceendpointVisualstudiomarketplaceAuthenticationBasic defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String password;
        private String username;
        public Builder() {}
        public Builder(ServiceendpointVisualstudiomarketplaceAuthenticationBasic defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.password = defaults.password;
    	      this.username = defaults.username;
        }

        @CustomType.Setter
        public Builder password(String password) {
            if (password == null) {
              throw new MissingRequiredPropertyException("ServiceendpointVisualstudiomarketplaceAuthenticationBasic", "password");
            }
            this.password = password;
            return this;
        }
        @CustomType.Setter
        public Builder username(String username) {
            if (username == null) {
              throw new MissingRequiredPropertyException("ServiceendpointVisualstudiomarketplaceAuthenticationBasic", "username");
            }
            this.username = username;
            return this;
        }
        public ServiceendpointVisualstudiomarketplaceAuthenticationBasic build() {
            final var _resultValue = new ServiceendpointVisualstudiomarketplaceAuthenticationBasic();
            _resultValue.password = password;
            _resultValue.username = username;
            return _resultValue;
        }
    }
}