// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class ServiceendpointJfrogArtifactoryV2AuthenticationTokenArgs extends com.pulumi.resources.ResourceArgs {

    public static final ServiceendpointJfrogArtifactoryV2AuthenticationTokenArgs Empty = new ServiceendpointJfrogArtifactoryV2AuthenticationTokenArgs();

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

    private ServiceendpointJfrogArtifactoryV2AuthenticationTokenArgs() {}

    private ServiceendpointJfrogArtifactoryV2AuthenticationTokenArgs(ServiceendpointJfrogArtifactoryV2AuthenticationTokenArgs $) {
        this.token = $.token;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServiceendpointJfrogArtifactoryV2AuthenticationTokenArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServiceendpointJfrogArtifactoryV2AuthenticationTokenArgs $;

        public Builder() {
            $ = new ServiceendpointJfrogArtifactoryV2AuthenticationTokenArgs();
        }

        public Builder(ServiceendpointJfrogArtifactoryV2AuthenticationTokenArgs defaults) {
            $ = new ServiceendpointJfrogArtifactoryV2AuthenticationTokenArgs(Objects.requireNonNull(defaults));
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

        public ServiceendpointJfrogArtifactoryV2AuthenticationTokenArgs build() {
            $.token = Objects.requireNonNull($.token, "expected parameter 'token' to be non-null");
            return $;
        }
    }

}
