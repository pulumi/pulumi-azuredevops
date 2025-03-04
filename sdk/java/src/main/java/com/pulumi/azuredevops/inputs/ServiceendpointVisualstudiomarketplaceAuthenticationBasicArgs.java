// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class ServiceendpointVisualstudiomarketplaceAuthenticationBasicArgs extends com.pulumi.resources.ResourceArgs {

    public static final ServiceendpointVisualstudiomarketplaceAuthenticationBasicArgs Empty = new ServiceendpointVisualstudiomarketplaceAuthenticationBasicArgs();

    /**
     * The password of the marketplace.
     * 
     */
    @Import(name="password", required=true)
    private Output<String> password;

    /**
     * @return The password of the marketplace.
     * 
     */
    public Output<String> password() {
        return this.password;
    }

    /**
     * The username of the marketplace.
     * 
     */
    @Import(name="username", required=true)
    private Output<String> username;

    /**
     * @return The username of the marketplace.
     * 
     */
    public Output<String> username() {
        return this.username;
    }

    private ServiceendpointVisualstudiomarketplaceAuthenticationBasicArgs() {}

    private ServiceendpointVisualstudiomarketplaceAuthenticationBasicArgs(ServiceendpointVisualstudiomarketplaceAuthenticationBasicArgs $) {
        this.password = $.password;
        this.username = $.username;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServiceendpointVisualstudiomarketplaceAuthenticationBasicArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServiceendpointVisualstudiomarketplaceAuthenticationBasicArgs $;

        public Builder() {
            $ = new ServiceendpointVisualstudiomarketplaceAuthenticationBasicArgs();
        }

        public Builder(ServiceendpointVisualstudiomarketplaceAuthenticationBasicArgs defaults) {
            $ = new ServiceendpointVisualstudiomarketplaceAuthenticationBasicArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param password The password of the marketplace.
         * 
         * @return builder
         * 
         */
        public Builder password(Output<String> password) {
            $.password = password;
            return this;
        }

        /**
         * @param password The password of the marketplace.
         * 
         * @return builder
         * 
         */
        public Builder password(String password) {
            return password(Output.of(password));
        }

        /**
         * @param username The username of the marketplace.
         * 
         * @return builder
         * 
         */
        public Builder username(Output<String> username) {
            $.username = username;
            return this;
        }

        /**
         * @param username The username of the marketplace.
         * 
         * @return builder
         * 
         */
        public Builder username(String username) {
            return username(Output.of(username));
        }

        public ServiceendpointVisualstudiomarketplaceAuthenticationBasicArgs build() {
            if ($.password == null) {
                throw new MissingRequiredPropertyException("ServiceendpointVisualstudiomarketplaceAuthenticationBasicArgs", "password");
            }
            if ($.username == null) {
                throw new MissingRequiredPropertyException("ServiceendpointVisualstudiomarketplaceAuthenticationBasicArgs", "username");
            }
            return $;
        }
    }

}
