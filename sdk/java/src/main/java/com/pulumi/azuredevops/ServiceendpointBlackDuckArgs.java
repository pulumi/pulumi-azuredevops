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


public final class ServiceendpointBlackDuckArgs extends com.pulumi.resources.ResourceArgs {

    public static final ServiceendpointBlackDuckArgs Empty = new ServiceendpointBlackDuckArgs();

    /**
     * The API token of the Black Duck Detect.
     * 
     */
    @Import(name="apiToken", required=true)
    private Output<String> apiToken;

    /**
     * @return The API token of the Black Duck Detect.
     * 
     */
    public Output<String> apiToken() {
        return this.apiToken;
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
     * The server URL of the Black Duck Detect.
     * 
     */
    @Import(name="serverUrl", required=true)
    private Output<String> serverUrl;

    /**
     * @return The server URL of the Black Duck Detect.
     * 
     */
    public Output<String> serverUrl() {
        return this.serverUrl;
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

    private ServiceendpointBlackDuckArgs() {}

    private ServiceendpointBlackDuckArgs(ServiceendpointBlackDuckArgs $) {
        this.apiToken = $.apiToken;
        this.description = $.description;
        this.projectId = $.projectId;
        this.serverUrl = $.serverUrl;
        this.serviceEndpointName = $.serviceEndpointName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServiceendpointBlackDuckArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServiceendpointBlackDuckArgs $;

        public Builder() {
            $ = new ServiceendpointBlackDuckArgs();
        }

        public Builder(ServiceendpointBlackDuckArgs defaults) {
            $ = new ServiceendpointBlackDuckArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param apiToken The API token of the Black Duck Detect.
         * 
         * @return builder
         * 
         */
        public Builder apiToken(Output<String> apiToken) {
            $.apiToken = apiToken;
            return this;
        }

        /**
         * @param apiToken The API token of the Black Duck Detect.
         * 
         * @return builder
         * 
         */
        public Builder apiToken(String apiToken) {
            return apiToken(Output.of(apiToken));
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
         * @param serverUrl The server URL of the Black Duck Detect.
         * 
         * @return builder
         * 
         */
        public Builder serverUrl(Output<String> serverUrl) {
            $.serverUrl = serverUrl;
            return this;
        }

        /**
         * @param serverUrl The server URL of the Black Duck Detect.
         * 
         * @return builder
         * 
         */
        public Builder serverUrl(String serverUrl) {
            return serverUrl(Output.of(serverUrl));
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

        public ServiceendpointBlackDuckArgs build() {
            if ($.apiToken == null) {
                throw new MissingRequiredPropertyException("ServiceendpointBlackDuckArgs", "apiToken");
            }
            if ($.projectId == null) {
                throw new MissingRequiredPropertyException("ServiceendpointBlackDuckArgs", "projectId");
            }
            if ($.serverUrl == null) {
                throw new MissingRequiredPropertyException("ServiceendpointBlackDuckArgs", "serverUrl");
            }
            if ($.serviceEndpointName == null) {
                throw new MissingRequiredPropertyException("ServiceendpointBlackDuckArgs", "serviceEndpointName");
            }
            return $;
        }
    }

}
