// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ServiceEndpointAzureDevOpsArgs extends com.pulumi.resources.ResourceArgs {

    public static final ServiceEndpointAzureDevOpsArgs Empty = new ServiceEndpointAzureDevOpsArgs();

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
     * The organization URL.
     * 
     */
    @Import(name="orgUrl", required=true)
    private Output<String> orgUrl;

    /**
     * @return The organization URL.
     * 
     */
    public Output<String> orgUrl() {
        return this.orgUrl;
    }

    /**
     * The Azure DevOps personal access token.
     * 
     */
    @Import(name="personalAccessToken", required=true)
    private Output<String> personalAccessToken;

    /**
     * @return The Azure DevOps personal access token.
     * 
     */
    public Output<String> personalAccessToken() {
        return this.personalAccessToken;
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
     * The URL of the release API.
     * 
     */
    @Import(name="releaseApiUrl", required=true)
    private Output<String> releaseApiUrl;

    /**
     * @return The URL of the release API.
     * 
     */
    public Output<String> releaseApiUrl() {
        return this.releaseApiUrl;
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

    private ServiceEndpointAzureDevOpsArgs() {}

    private ServiceEndpointAzureDevOpsArgs(ServiceEndpointAzureDevOpsArgs $) {
        this.authorization = $.authorization;
        this.description = $.description;
        this.orgUrl = $.orgUrl;
        this.personalAccessToken = $.personalAccessToken;
        this.projectId = $.projectId;
        this.releaseApiUrl = $.releaseApiUrl;
        this.serviceEndpointName = $.serviceEndpointName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServiceEndpointAzureDevOpsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServiceEndpointAzureDevOpsArgs $;

        public Builder() {
            $ = new ServiceEndpointAzureDevOpsArgs();
        }

        public Builder(ServiceEndpointAzureDevOpsArgs defaults) {
            $ = new ServiceEndpointAzureDevOpsArgs(Objects.requireNonNull(defaults));
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
         * @param orgUrl The organization URL.
         * 
         * @return builder
         * 
         */
        public Builder orgUrl(Output<String> orgUrl) {
            $.orgUrl = orgUrl;
            return this;
        }

        /**
         * @param orgUrl The organization URL.
         * 
         * @return builder
         * 
         */
        public Builder orgUrl(String orgUrl) {
            return orgUrl(Output.of(orgUrl));
        }

        /**
         * @param personalAccessToken The Azure DevOps personal access token.
         * 
         * @return builder
         * 
         */
        public Builder personalAccessToken(Output<String> personalAccessToken) {
            $.personalAccessToken = personalAccessToken;
            return this;
        }

        /**
         * @param personalAccessToken The Azure DevOps personal access token.
         * 
         * @return builder
         * 
         */
        public Builder personalAccessToken(String personalAccessToken) {
            return personalAccessToken(Output.of(personalAccessToken));
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
         * @param releaseApiUrl The URL of the release API.
         * 
         * @return builder
         * 
         */
        public Builder releaseApiUrl(Output<String> releaseApiUrl) {
            $.releaseApiUrl = releaseApiUrl;
            return this;
        }

        /**
         * @param releaseApiUrl The URL of the release API.
         * 
         * @return builder
         * 
         */
        public Builder releaseApiUrl(String releaseApiUrl) {
            return releaseApiUrl(Output.of(releaseApiUrl));
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

        public ServiceEndpointAzureDevOpsArgs build() {
            $.orgUrl = Objects.requireNonNull($.orgUrl, "expected parameter 'orgUrl' to be non-null");
            $.personalAccessToken = Objects.requireNonNull($.personalAccessToken, "expected parameter 'personalAccessToken' to be non-null");
            $.projectId = Objects.requireNonNull($.projectId, "expected parameter 'projectId' to be non-null");
            $.releaseApiUrl = Objects.requireNonNull($.releaseApiUrl, "expected parameter 'releaseApiUrl' to be non-null");
            $.serviceEndpointName = Objects.requireNonNull($.serviceEndpointName, "expected parameter 'serviceEndpointName' to be non-null");
            return $;
        }
    }

}