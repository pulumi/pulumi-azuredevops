// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class ServiceEndpointGitHubAuthPersonalArgs extends com.pulumi.resources.ResourceArgs {

    public static final ServiceEndpointGitHubAuthPersonalArgs Empty = new ServiceEndpointGitHubAuthPersonalArgs();

    /**
     * The Personal Access Token for GitHub.
     * 
     */
    @Import(name="personalAccessToken", required=true)
    private Output<String> personalAccessToken;

    /**
     * @return The Personal Access Token for GitHub.
     * 
     */
    public Output<String> personalAccessToken() {
        return this.personalAccessToken;
    }

    private ServiceEndpointGitHubAuthPersonalArgs() {}

    private ServiceEndpointGitHubAuthPersonalArgs(ServiceEndpointGitHubAuthPersonalArgs $) {
        this.personalAccessToken = $.personalAccessToken;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServiceEndpointGitHubAuthPersonalArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServiceEndpointGitHubAuthPersonalArgs $;

        public Builder() {
            $ = new ServiceEndpointGitHubAuthPersonalArgs();
        }

        public Builder(ServiceEndpointGitHubAuthPersonalArgs defaults) {
            $ = new ServiceEndpointGitHubAuthPersonalArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param personalAccessToken The Personal Access Token for GitHub.
         * 
         * @return builder
         * 
         */
        public Builder personalAccessToken(Output<String> personalAccessToken) {
            $.personalAccessToken = personalAccessToken;
            return this;
        }

        /**
         * @param personalAccessToken The Personal Access Token for GitHub.
         * 
         * @return builder
         * 
         */
        public Builder personalAccessToken(String personalAccessToken) {
            return personalAccessToken(Output.of(personalAccessToken));
        }

        public ServiceEndpointGitHubAuthPersonalArgs build() {
            $.personalAccessToken = Objects.requireNonNull($.personalAccessToken, "expected parameter 'personalAccessToken' to be non-null");
            return $;
        }
    }

}
