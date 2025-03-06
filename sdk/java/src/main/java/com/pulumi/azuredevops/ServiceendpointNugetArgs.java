// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ServiceendpointNugetArgs extends com.pulumi.resources.ResourceArgs {

    public static final ServiceendpointNugetArgs Empty = new ServiceendpointNugetArgs();

    /**
     * The API Key used to connect to the endpoint.
     * 
     */
    @Import(name="apiKey")
    private @Nullable Output<String> apiKey;

    /**
     * @return The API Key used to connect to the endpoint.
     * 
     */
    public Optional<Output<String>> apiKey() {
        return Optional.ofNullable(this.apiKey);
    }

    @Import(name="description")
    private @Nullable Output<String> description;

    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The URL for the feed. This will generally end with `index.json`.
     * 
     */
    @Import(name="feedUrl", required=true)
    private Output<String> feedUrl;

    /**
     * @return The URL for the feed. This will generally end with `index.json`.
     * 
     */
    public Output<String> feedUrl() {
        return this.feedUrl;
    }

    /**
     * The account password used to connect to the endpoint
     * 
     * &gt; **Note** Only one of `api_key` or `personal_access_token` or  `username`, `password` can be set at the same time.
     * 
     */
    @Import(name="password")
    private @Nullable Output<String> password;

    /**
     * @return The account password used to connect to the endpoint
     * 
     * &gt; **Note** Only one of `api_key` or `personal_access_token` or  `username`, `password` can be set at the same time.
     * 
     */
    public Optional<Output<String>> password() {
        return Optional.ofNullable(this.password);
    }

    /**
     * The Personal access token used to  connect to the endpoint. Personal access tokens are applicable only for NuGet feeds hosted on other Azure DevOps Services organizations or Azure DevOps Server 2019 (or later).
     * 
     */
    @Import(name="personalAccessToken")
    private @Nullable Output<String> personalAccessToken;

    /**
     * @return The Personal access token used to  connect to the endpoint. Personal access tokens are applicable only for NuGet feeds hosted on other Azure DevOps Services organizations or Azure DevOps Server 2019 (or later).
     * 
     */
    public Optional<Output<String>> personalAccessToken() {
        return Optional.ofNullable(this.personalAccessToken);
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

    /**
     * The account username used to connect to the endpoint.
     * 
     */
    @Import(name="username")
    private @Nullable Output<String> username;

    /**
     * @return The account username used to connect to the endpoint.
     * 
     */
    public Optional<Output<String>> username() {
        return Optional.ofNullable(this.username);
    }

    private ServiceendpointNugetArgs() {}

    private ServiceendpointNugetArgs(ServiceendpointNugetArgs $) {
        this.apiKey = $.apiKey;
        this.description = $.description;
        this.feedUrl = $.feedUrl;
        this.password = $.password;
        this.personalAccessToken = $.personalAccessToken;
        this.projectId = $.projectId;
        this.serviceEndpointName = $.serviceEndpointName;
        this.username = $.username;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServiceendpointNugetArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServiceendpointNugetArgs $;

        public Builder() {
            $ = new ServiceendpointNugetArgs();
        }

        public Builder(ServiceendpointNugetArgs defaults) {
            $ = new ServiceendpointNugetArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param apiKey The API Key used to connect to the endpoint.
         * 
         * @return builder
         * 
         */
        public Builder apiKey(@Nullable Output<String> apiKey) {
            $.apiKey = apiKey;
            return this;
        }

        /**
         * @param apiKey The API Key used to connect to the endpoint.
         * 
         * @return builder
         * 
         */
        public Builder apiKey(String apiKey) {
            return apiKey(Output.of(apiKey));
        }

        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param feedUrl The URL for the feed. This will generally end with `index.json`.
         * 
         * @return builder
         * 
         */
        public Builder feedUrl(Output<String> feedUrl) {
            $.feedUrl = feedUrl;
            return this;
        }

        /**
         * @param feedUrl The URL for the feed. This will generally end with `index.json`.
         * 
         * @return builder
         * 
         */
        public Builder feedUrl(String feedUrl) {
            return feedUrl(Output.of(feedUrl));
        }

        /**
         * @param password The account password used to connect to the endpoint
         * 
         * &gt; **Note** Only one of `api_key` or `personal_access_token` or  `username`, `password` can be set at the same time.
         * 
         * @return builder
         * 
         */
        public Builder password(@Nullable Output<String> password) {
            $.password = password;
            return this;
        }

        /**
         * @param password The account password used to connect to the endpoint
         * 
         * &gt; **Note** Only one of `api_key` or `personal_access_token` or  `username`, `password` can be set at the same time.
         * 
         * @return builder
         * 
         */
        public Builder password(String password) {
            return password(Output.of(password));
        }

        /**
         * @param personalAccessToken The Personal access token used to  connect to the endpoint. Personal access tokens are applicable only for NuGet feeds hosted on other Azure DevOps Services organizations or Azure DevOps Server 2019 (or later).
         * 
         * @return builder
         * 
         */
        public Builder personalAccessToken(@Nullable Output<String> personalAccessToken) {
            $.personalAccessToken = personalAccessToken;
            return this;
        }

        /**
         * @param personalAccessToken The Personal access token used to  connect to the endpoint. Personal access tokens are applicable only for NuGet feeds hosted on other Azure DevOps Services organizations or Azure DevOps Server 2019 (or later).
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

        /**
         * @param username The account username used to connect to the endpoint.
         * 
         * @return builder
         * 
         */
        public Builder username(@Nullable Output<String> username) {
            $.username = username;
            return this;
        }

        /**
         * @param username The account username used to connect to the endpoint.
         * 
         * @return builder
         * 
         */
        public Builder username(String username) {
            return username(Output.of(username));
        }

        public ServiceendpointNugetArgs build() {
            if ($.feedUrl == null) {
                throw new MissingRequiredPropertyException("ServiceendpointNugetArgs", "feedUrl");
            }
            if ($.projectId == null) {
                throw new MissingRequiredPropertyException("ServiceendpointNugetArgs", "projectId");
            }
            if ($.serviceEndpointName == null) {
                throw new MissingRequiredPropertyException("ServiceendpointNugetArgs", "serviceEndpointName");
            }
            return $;
        }
    }

}
