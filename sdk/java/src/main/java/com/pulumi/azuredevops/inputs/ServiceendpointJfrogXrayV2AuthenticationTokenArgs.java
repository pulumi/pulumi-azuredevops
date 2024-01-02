// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class ServiceendpointJfrogXrayV2AuthenticationTokenArgs extends com.pulumi.resources.ResourceArgs {

    public static final ServiceendpointJfrogXrayV2AuthenticationTokenArgs Empty = new ServiceendpointJfrogXrayV2AuthenticationTokenArgs();

    /**
     * Authentication Token generated through Artifactory.
     * 
     */
    @Import(name="token", required=true)
    private Output<String> token;

    /**
     * @return Authentication Token generated through Artifactory.
     * 
     */
    public Output<String> token() {
        return this.token;
    }

    private ServiceendpointJfrogXrayV2AuthenticationTokenArgs() {}

    private ServiceendpointJfrogXrayV2AuthenticationTokenArgs(ServiceendpointJfrogXrayV2AuthenticationTokenArgs $) {
        this.token = $.token;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServiceendpointJfrogXrayV2AuthenticationTokenArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServiceendpointJfrogXrayV2AuthenticationTokenArgs $;

        public Builder() {
            $ = new ServiceendpointJfrogXrayV2AuthenticationTokenArgs();
        }

        public Builder(ServiceendpointJfrogXrayV2AuthenticationTokenArgs defaults) {
            $ = new ServiceendpointJfrogXrayV2AuthenticationTokenArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param token Authentication Token generated through Artifactory.
         * 
         * @return builder
         * 
         */
        public Builder token(Output<String> token) {
            $.token = token;
            return this;
        }

        /**
         * @param token Authentication Token generated through Artifactory.
         * 
         * @return builder
         * 
         */
        public Builder token(String token) {
            return token(Output.of(token));
        }

        public ServiceendpointJfrogXrayV2AuthenticationTokenArgs build() {
            if ($.token == null) {
                throw new MissingRequiredPropertyException("ServiceendpointJfrogXrayV2AuthenticationTokenArgs", "token");
            }
            return $;
        }
    }

}
