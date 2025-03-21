// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops;

import com.pulumi.azuredevops.inputs.ServiceEndpointGitHubAuthOauthArgs;
import com.pulumi.azuredevops.inputs.ServiceEndpointGitHubAuthPersonalArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ServiceEndpointGitHubArgs extends com.pulumi.resources.ResourceArgs {

    public static final ServiceEndpointGitHubArgs Empty = new ServiceEndpointGitHubArgs();

    /**
     * An `auth_oauth` block as documented below. Allows connecting using an Oauth token.
     * 
     */
    @Import(name="authOauth")
    private @Nullable Output<ServiceEndpointGitHubAuthOauthArgs> authOauth;

    /**
     * @return An `auth_oauth` block as documented below. Allows connecting using an Oauth token.
     * 
     */
    public Optional<Output<ServiceEndpointGitHubAuthOauthArgs>> authOauth() {
        return Optional.ofNullable(this.authOauth);
    }

    /**
     * An `auth_personal` block as documented below. Allows connecting using a personal access token.
     * 
     */
    @Import(name="authPersonal")
    private @Nullable Output<ServiceEndpointGitHubAuthPersonalArgs> authPersonal;

    /**
     * @return An `auth_personal` block as documented below. Allows connecting using a personal access token.
     * 
     */
    public Optional<Output<ServiceEndpointGitHubAuthPersonalArgs>> authPersonal() {
        return Optional.ofNullable(this.authPersonal);
    }

    @Import(name="description")
    private @Nullable Output<String> description;

    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The ID of the project.
     * 
     */
    @Import(name="projectId", required=true)
    private Output<String> projectId;

    /**
     * @return The ID of the project.
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }

    /**
     * The Service Endpoint name.
     * 
     */
    @Import(name="serviceEndpointName", required=true)
    private Output<String> serviceEndpointName;

    /**
     * @return The Service Endpoint name.
     * 
     */
    public Output<String> serviceEndpointName() {
        return this.serviceEndpointName;
    }

    private ServiceEndpointGitHubArgs() {}

    private ServiceEndpointGitHubArgs(ServiceEndpointGitHubArgs $) {
        this.authOauth = $.authOauth;
        this.authPersonal = $.authPersonal;
        this.description = $.description;
        this.projectId = $.projectId;
        this.serviceEndpointName = $.serviceEndpointName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServiceEndpointGitHubArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServiceEndpointGitHubArgs $;

        public Builder() {
            $ = new ServiceEndpointGitHubArgs();
        }

        public Builder(ServiceEndpointGitHubArgs defaults) {
            $ = new ServiceEndpointGitHubArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param authOauth An `auth_oauth` block as documented below. Allows connecting using an Oauth token.
         * 
         * @return builder
         * 
         */
        public Builder authOauth(@Nullable Output<ServiceEndpointGitHubAuthOauthArgs> authOauth) {
            $.authOauth = authOauth;
            return this;
        }

        /**
         * @param authOauth An `auth_oauth` block as documented below. Allows connecting using an Oauth token.
         * 
         * @return builder
         * 
         */
        public Builder authOauth(ServiceEndpointGitHubAuthOauthArgs authOauth) {
            return authOauth(Output.of(authOauth));
        }

        /**
         * @param authPersonal An `auth_personal` block as documented below. Allows connecting using a personal access token.
         * 
         * @return builder
         * 
         */
        public Builder authPersonal(@Nullable Output<ServiceEndpointGitHubAuthPersonalArgs> authPersonal) {
            $.authPersonal = authPersonal;
            return this;
        }

        /**
         * @param authPersonal An `auth_personal` block as documented below. Allows connecting using a personal access token.
         * 
         * @return builder
         * 
         */
        public Builder authPersonal(ServiceEndpointGitHubAuthPersonalArgs authPersonal) {
            return authPersonal(Output.of(authPersonal));
        }

        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param projectId The ID of the project.
         * 
         * @return builder
         * 
         */
        public Builder projectId(Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param projectId The ID of the project.
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
        }

        /**
         * @param serviceEndpointName The Service Endpoint name.
         * 
         * @return builder
         * 
         */
        public Builder serviceEndpointName(Output<String> serviceEndpointName) {
            $.serviceEndpointName = serviceEndpointName;
            return this;
        }

        /**
         * @param serviceEndpointName The Service Endpoint name.
         * 
         * @return builder
         * 
         */
        public Builder serviceEndpointName(String serviceEndpointName) {
            return serviceEndpointName(Output.of(serviceEndpointName));
        }

        public ServiceEndpointGitHubArgs build() {
            if ($.projectId == null) {
                throw new MissingRequiredPropertyException("ServiceEndpointGitHubArgs", "projectId");
            }
            if ($.serviceEndpointName == null) {
                throw new MissingRequiredPropertyException("ServiceEndpointGitHubArgs", "serviceEndpointName");
            }
            return $;
        }
    }

}
