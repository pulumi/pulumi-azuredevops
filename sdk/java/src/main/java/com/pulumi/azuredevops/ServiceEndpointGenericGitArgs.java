// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ServiceEndpointGenericGitArgs extends com.pulumi.resources.ResourceArgs {

    public static final ServiceEndpointGenericGitArgs Empty = new ServiceEndpointGenericGitArgs();

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
     * A value indicating whether or not to attempt accessing this git server from Azure Pipelines.
     * 
     */
    @Import(name="enablePipelinesAccess")
    private @Nullable Output<Boolean> enablePipelinesAccess;

    /**
     * @return A value indicating whether or not to attempt accessing this git server from Azure Pipelines.
     * 
     */
    public Optional<Output<Boolean>> enablePipelinesAccess() {
        return Optional.ofNullable(this.enablePipelinesAccess);
    }

    /**
     * The PAT or password used to authenticate to the git repository.
     * 
     */
    @Import(name="password")
    private @Nullable Output<String> password;

    /**
     * @return The PAT or password used to authenticate to the git repository.
     * 
     */
    public Optional<Output<String>> password() {
        return Optional.ofNullable(this.password);
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
     * The URL of the repository associated with the service endpoint.
     * 
     */
    @Import(name="repositoryUrl", required=true)
    private Output<String> repositoryUrl;

    /**
     * @return The URL of the repository associated with the service endpoint.
     * 
     */
    public Output<String> repositoryUrl() {
        return this.repositoryUrl;
    }

    /**
     * The name of the service endpoint.
     * 
     */
    @Import(name="serviceEndpointName", required=true)
    private Output<String> serviceEndpointName;

    /**
     * @return The name of the service endpoint.
     * 
     */
    public Output<String> serviceEndpointName() {
        return this.serviceEndpointName;
    }

    /**
     * The username used to authenticate to the git repository.
     * 
     */
    @Import(name="username")
    private @Nullable Output<String> username;

    /**
     * @return The username used to authenticate to the git repository.
     * 
     */
    public Optional<Output<String>> username() {
        return Optional.ofNullable(this.username);
    }

    private ServiceEndpointGenericGitArgs() {}

    private ServiceEndpointGenericGitArgs(ServiceEndpointGenericGitArgs $) {
        this.authorization = $.authorization;
        this.description = $.description;
        this.enablePipelinesAccess = $.enablePipelinesAccess;
        this.password = $.password;
        this.projectId = $.projectId;
        this.repositoryUrl = $.repositoryUrl;
        this.serviceEndpointName = $.serviceEndpointName;
        this.username = $.username;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServiceEndpointGenericGitArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServiceEndpointGenericGitArgs $;

        public Builder() {
            $ = new ServiceEndpointGenericGitArgs();
        }

        public Builder(ServiceEndpointGenericGitArgs defaults) {
            $ = new ServiceEndpointGenericGitArgs(Objects.requireNonNull(defaults));
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
         * @param enablePipelinesAccess A value indicating whether or not to attempt accessing this git server from Azure Pipelines.
         * 
         * @return builder
         * 
         */
        public Builder enablePipelinesAccess(@Nullable Output<Boolean> enablePipelinesAccess) {
            $.enablePipelinesAccess = enablePipelinesAccess;
            return this;
        }

        /**
         * @param enablePipelinesAccess A value indicating whether or not to attempt accessing this git server from Azure Pipelines.
         * 
         * @return builder
         * 
         */
        public Builder enablePipelinesAccess(Boolean enablePipelinesAccess) {
            return enablePipelinesAccess(Output.of(enablePipelinesAccess));
        }

        /**
         * @param password The PAT or password used to authenticate to the git repository.
         * 
         * @return builder
         * 
         */
        public Builder password(@Nullable Output<String> password) {
            $.password = password;
            return this;
        }

        /**
         * @param password The PAT or password used to authenticate to the git repository.
         * 
         * @return builder
         * 
         */
        public Builder password(String password) {
            return password(Output.of(password));
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
         * @param repositoryUrl The URL of the repository associated with the service endpoint.
         * 
         * @return builder
         * 
         */
        public Builder repositoryUrl(Output<String> repositoryUrl) {
            $.repositoryUrl = repositoryUrl;
            return this;
        }

        /**
         * @param repositoryUrl The URL of the repository associated with the service endpoint.
         * 
         * @return builder
         * 
         */
        public Builder repositoryUrl(String repositoryUrl) {
            return repositoryUrl(Output.of(repositoryUrl));
        }

        /**
         * @param serviceEndpointName The name of the service endpoint.
         * 
         * @return builder
         * 
         */
        public Builder serviceEndpointName(Output<String> serviceEndpointName) {
            $.serviceEndpointName = serviceEndpointName;
            return this;
        }

        /**
         * @param serviceEndpointName The name of the service endpoint.
         * 
         * @return builder
         * 
         */
        public Builder serviceEndpointName(String serviceEndpointName) {
            return serviceEndpointName(Output.of(serviceEndpointName));
        }

        /**
         * @param username The username used to authenticate to the git repository.
         * 
         * @return builder
         * 
         */
        public Builder username(@Nullable Output<String> username) {
            $.username = username;
            return this;
        }

        /**
         * @param username The username used to authenticate to the git repository.
         * 
         * @return builder
         * 
         */
        public Builder username(String username) {
            return username(Output.of(username));
        }

        public ServiceEndpointGenericGitArgs build() {
            $.projectId = Objects.requireNonNull($.projectId, "expected parameter 'projectId' to be non-null");
            $.repositoryUrl = Objects.requireNonNull($.repositoryUrl, "expected parameter 'repositoryUrl' to be non-null");
            $.serviceEndpointName = Objects.requireNonNull($.serviceEndpointName, "expected parameter 'serviceEndpointName' to be non-null");
            return $;
        }
    }

}