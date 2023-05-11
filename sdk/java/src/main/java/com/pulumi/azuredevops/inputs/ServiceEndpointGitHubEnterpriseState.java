// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.inputs;

import com.pulumi.azuredevops.inputs.ServiceEndpointGitHubEnterpriseAuthPersonalArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ServiceEndpointGitHubEnterpriseState extends com.pulumi.resources.ResourceArgs {

    public static final ServiceEndpointGitHubEnterpriseState Empty = new ServiceEndpointGitHubEnterpriseState();

    /**
     * An `auth_personal` block as documented below. Allows connecting using a personal access token.
     * 
     */
    @Import(name="authPersonal")
    private @Nullable Output<ServiceEndpointGitHubEnterpriseAuthPersonalArgs> authPersonal;

    /**
     * @return An `auth_personal` block as documented below. Allows connecting using a personal access token.
     * 
     */
    public Optional<Output<ServiceEndpointGitHubEnterpriseAuthPersonalArgs>> authPersonal() {
        return Optional.ofNullable(this.authPersonal);
    }

    @Import(name="authorization")
    private @Nullable Output<Map<String,String>> authorization;

    public Optional<Output<Map<String,String>>> authorization() {
        return Optional.ofNullable(this.authorization);
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
    @Import(name="projectId")
    private @Nullable Output<String> projectId;

    /**
     * @return The ID of the project.
     * 
     */
    public Optional<Output<String>> projectId() {
        return Optional.ofNullable(this.projectId);
    }

    /**
     * The Service Endpoint name.
     * 
     */
    @Import(name="serviceEndpointName")
    private @Nullable Output<String> serviceEndpointName;

    /**
     * @return The Service Endpoint name.
     * 
     */
    public Optional<Output<String>> serviceEndpointName() {
        return Optional.ofNullable(this.serviceEndpointName);
    }

    /**
     * GitHub Enterprise Server Url.
     * 
     */
    @Import(name="url")
    private @Nullable Output<String> url;

    /**
     * @return GitHub Enterprise Server Url.
     * 
     */
    public Optional<Output<String>> url() {
        return Optional.ofNullable(this.url);
    }

    private ServiceEndpointGitHubEnterpriseState() {}

    private ServiceEndpointGitHubEnterpriseState(ServiceEndpointGitHubEnterpriseState $) {
        this.authPersonal = $.authPersonal;
        this.authorization = $.authorization;
        this.description = $.description;
        this.projectId = $.projectId;
        this.serviceEndpointName = $.serviceEndpointName;
        this.url = $.url;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServiceEndpointGitHubEnterpriseState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServiceEndpointGitHubEnterpriseState $;

        public Builder() {
            $ = new ServiceEndpointGitHubEnterpriseState();
        }

        public Builder(ServiceEndpointGitHubEnterpriseState defaults) {
            $ = new ServiceEndpointGitHubEnterpriseState(Objects.requireNonNull(defaults));
        }

        /**
         * @param authPersonal An `auth_personal` block as documented below. Allows connecting using a personal access token.
         * 
         * @return builder
         * 
         */
        public Builder authPersonal(@Nullable Output<ServiceEndpointGitHubEnterpriseAuthPersonalArgs> authPersonal) {
            $.authPersonal = authPersonal;
            return this;
        }

        /**
         * @param authPersonal An `auth_personal` block as documented below. Allows connecting using a personal access token.
         * 
         * @return builder
         * 
         */
        public Builder authPersonal(ServiceEndpointGitHubEnterpriseAuthPersonalArgs authPersonal) {
            return authPersonal(Output.of(authPersonal));
        }

        public Builder authorization(@Nullable Output<Map<String,String>> authorization) {
            $.authorization = authorization;
            return this;
        }

        public Builder authorization(Map<String,String> authorization) {
            return authorization(Output.of(authorization));
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
        public Builder projectId(@Nullable Output<String> projectId) {
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
        public Builder serviceEndpointName(@Nullable Output<String> serviceEndpointName) {
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

        /**
         * @param url GitHub Enterprise Server Url.
         * 
         * @return builder
         * 
         */
        public Builder url(@Nullable Output<String> url) {
            $.url = url;
            return this;
        }

        /**
         * @param url GitHub Enterprise Server Url.
         * 
         * @return builder
         * 
         */
        public Builder url(String url) {
            return url(Output.of(url));
        }

        public ServiceEndpointGitHubEnterpriseState build() {
            return $;
        }
    }

}