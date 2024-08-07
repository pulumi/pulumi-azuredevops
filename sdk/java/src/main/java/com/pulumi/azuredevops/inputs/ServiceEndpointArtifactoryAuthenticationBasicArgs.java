// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class ServiceEndpointArtifactoryAuthenticationBasicArgs extends com.pulumi.resources.ResourceArgs {

    public static final ServiceEndpointArtifactoryAuthenticationBasicArgs Empty = new ServiceEndpointArtifactoryAuthenticationBasicArgs();

    /**
     * The Artifactory password.
     * 
     */
    @Import(name="password", required=true)
    private Output<String> password;

    /**
     * @return The Artifactory password.
     * 
     */
    public Output<String> password() {
        return this.password;
    }

    /**
     * The Artifactory user name.
     * 
     */
    @Import(name="username", required=true)
    private Output<String> username;

    /**
     * @return The Artifactory user name.
     * 
     */
    public Output<String> username() {
        return this.username;
    }

    private ServiceEndpointArtifactoryAuthenticationBasicArgs() {}

    private ServiceEndpointArtifactoryAuthenticationBasicArgs(ServiceEndpointArtifactoryAuthenticationBasicArgs $) {
        this.password = $.password;
        this.username = $.username;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServiceEndpointArtifactoryAuthenticationBasicArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServiceEndpointArtifactoryAuthenticationBasicArgs $;

        public Builder() {
            $ = new ServiceEndpointArtifactoryAuthenticationBasicArgs();
        }

        public Builder(ServiceEndpointArtifactoryAuthenticationBasicArgs defaults) {
            $ = new ServiceEndpointArtifactoryAuthenticationBasicArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param password The Artifactory password.
         * 
         * @return builder
         * 
         */
        public Builder password(Output<String> password) {
            $.password = password;
            return this;
        }

        /**
         * @param password The Artifactory password.
         * 
         * @return builder
         * 
         */
        public Builder password(String password) {
            return password(Output.of(password));
        }

        /**
         * @param username The Artifactory user name.
         * 
         * @return builder
         * 
         */
        public Builder username(Output<String> username) {
            $.username = username;
            return this;
        }

        /**
         * @param username The Artifactory user name.
         * 
         * @return builder
         * 
         */
        public Builder username(String username) {
            return username(Output.of(username));
        }

        public ServiceEndpointArtifactoryAuthenticationBasicArgs build() {
            if ($.password == null) {
                throw new MissingRequiredPropertyException("ServiceEndpointArtifactoryAuthenticationBasicArgs", "password");
            }
            if ($.username == null) {
                throw new MissingRequiredPropertyException("ServiceEndpointArtifactoryAuthenticationBasicArgs", "username");
            }
            return $;
        }
    }

}
