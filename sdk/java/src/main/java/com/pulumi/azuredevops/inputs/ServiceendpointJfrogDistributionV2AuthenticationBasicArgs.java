// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class ServiceendpointJfrogDistributionV2AuthenticationBasicArgs extends com.pulumi.resources.ResourceArgs {

    public static final ServiceendpointJfrogDistributionV2AuthenticationBasicArgs Empty = new ServiceendpointJfrogDistributionV2AuthenticationBasicArgs();

    /**
     * The Password of the Artifactory.
     * 
     */
    @Import(name="password", required=true)
    private Output<String> password;

    /**
     * @return The Password of the Artifactory.
     * 
     */
    public Output<String> password() {
        return this.password;
    }

    /**
     * The Username of the Artifactory.
     * 
     */
    @Import(name="username", required=true)
    private Output<String> username;

    /**
     * @return The Username of the Artifactory.
     * 
     */
    public Output<String> username() {
        return this.username;
    }

    private ServiceendpointJfrogDistributionV2AuthenticationBasicArgs() {}

    private ServiceendpointJfrogDistributionV2AuthenticationBasicArgs(ServiceendpointJfrogDistributionV2AuthenticationBasicArgs $) {
        this.password = $.password;
        this.username = $.username;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServiceendpointJfrogDistributionV2AuthenticationBasicArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServiceendpointJfrogDistributionV2AuthenticationBasicArgs $;

        public Builder() {
            $ = new ServiceendpointJfrogDistributionV2AuthenticationBasicArgs();
        }

        public Builder(ServiceendpointJfrogDistributionV2AuthenticationBasicArgs defaults) {
            $ = new ServiceendpointJfrogDistributionV2AuthenticationBasicArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param password The Password of the Artifactory.
         * 
         * @return builder
         * 
         */
        public Builder password(Output<String> password) {
            $.password = password;
            return this;
        }

        /**
         * @param password The Password of the Artifactory.
         * 
         * @return builder
         * 
         */
        public Builder password(String password) {
            return password(Output.of(password));
        }

        /**
         * @param username The Username of the Artifactory.
         * 
         * @return builder
         * 
         */
        public Builder username(Output<String> username) {
            $.username = username;
            return this;
        }

        /**
         * @param username The Username of the Artifactory.
         * 
         * @return builder
         * 
         */
        public Builder username(String username) {
            return username(Output.of(username));
        }

        public ServiceendpointJfrogDistributionV2AuthenticationBasicArgs build() {
            if ($.password == null) {
                throw new MissingRequiredPropertyException("ServiceendpointJfrogDistributionV2AuthenticationBasicArgs", "password");
            }
            if ($.username == null) {
                throw new MissingRequiredPropertyException("ServiceendpointJfrogDistributionV2AuthenticationBasicArgs", "username");
            }
            return $;
        }
    }

}
