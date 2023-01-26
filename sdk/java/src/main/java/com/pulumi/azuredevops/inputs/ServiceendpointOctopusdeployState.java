// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ServiceendpointOctopusdeployState extends com.pulumi.resources.ResourceArgs {

    public static final ServiceendpointOctopusdeployState Empty = new ServiceendpointOctopusdeployState();

    /**
     * API key to connect to Octopus Deploy.
     * 
     */
    @Import(name="apiKey")
    private @Nullable Output<String> apiKey;

    /**
     * @return API key to connect to Octopus Deploy.
     * 
     */
    public Optional<Output<String>> apiKey() {
        return Optional.ofNullable(this.apiKey);
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
     * Whether to ignore SSL errors when connecting to the Octopus server from the agent. Default to `false`.
     * 
     */
    @Import(name="ignoreSslError")
    private @Nullable Output<Boolean> ignoreSslError;

    /**
     * @return Whether to ignore SSL errors when connecting to the Octopus server from the agent. Default to `false`.
     * 
     */
    public Optional<Output<Boolean>> ignoreSslError() {
        return Optional.ofNullable(this.ignoreSslError);
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
     * Octopus Server url.
     * 
     */
    @Import(name="url")
    private @Nullable Output<String> url;

    /**
     * @return Octopus Server url.
     * 
     */
    public Optional<Output<String>> url() {
        return Optional.ofNullable(this.url);
    }

    private ServiceendpointOctopusdeployState() {}

    private ServiceendpointOctopusdeployState(ServiceendpointOctopusdeployState $) {
        this.apiKey = $.apiKey;
        this.authorization = $.authorization;
        this.description = $.description;
        this.ignoreSslError = $.ignoreSslError;
        this.projectId = $.projectId;
        this.serviceEndpointName = $.serviceEndpointName;
        this.url = $.url;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServiceendpointOctopusdeployState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServiceendpointOctopusdeployState $;

        public Builder() {
            $ = new ServiceendpointOctopusdeployState();
        }

        public Builder(ServiceendpointOctopusdeployState defaults) {
            $ = new ServiceendpointOctopusdeployState(Objects.requireNonNull(defaults));
        }

        /**
         * @param apiKey API key to connect to Octopus Deploy.
         * 
         * @return builder
         * 
         */
        public Builder apiKey(@Nullable Output<String> apiKey) {
            $.apiKey = apiKey;
            return this;
        }

        /**
         * @param apiKey API key to connect to Octopus Deploy.
         * 
         * @return builder
         * 
         */
        public Builder apiKey(String apiKey) {
            return apiKey(Output.of(apiKey));
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
         * @param ignoreSslError Whether to ignore SSL errors when connecting to the Octopus server from the agent. Default to `false`.
         * 
         * @return builder
         * 
         */
        public Builder ignoreSslError(@Nullable Output<Boolean> ignoreSslError) {
            $.ignoreSslError = ignoreSslError;
            return this;
        }

        /**
         * @param ignoreSslError Whether to ignore SSL errors when connecting to the Octopus server from the agent. Default to `false`.
         * 
         * @return builder
         * 
         */
        public Builder ignoreSslError(Boolean ignoreSslError) {
            return ignoreSslError(Output.of(ignoreSslError));
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
         * @param url Octopus Server url.
         * 
         * @return builder
         * 
         */
        public Builder url(@Nullable Output<String> url) {
            $.url = url;
            return this;
        }

        /**
         * @param url Octopus Server url.
         * 
         * @return builder
         * 
         */
        public Builder url(String url) {
            return url(Output.of(url));
        }

        public ServiceendpointOctopusdeployState build() {
            return $;
        }
    }

}
