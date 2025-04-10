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


public final class ServiceEndpointGenericArgs extends com.pulumi.resources.ResourceArgs {

    public static final ServiceEndpointGenericArgs Empty = new ServiceEndpointGenericArgs();

    @Import(name="description")
    private @Nullable Output<String> description;

    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The password or token key used to authenticate to the server url using basic authentication.
     * 
     */
    @Import(name="password")
    private @Nullable Output<String> password;

    /**
     * @return The password or token key used to authenticate to the server url using basic authentication.
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
     * The URL of the server associated with the service endpoint.
     * 
     */
    @Import(name="serverUrl", required=true)
    private Output<String> serverUrl;

    /**
     * @return The URL of the server associated with the service endpoint.
     * 
     */
    public Output<String> serverUrl() {
        return this.serverUrl;
    }

    /**
     * The service endpoint name.
     * 
     */
    @Import(name="serviceEndpointName", required=true)
    private Output<String> serviceEndpointName;

    /**
     * @return The service endpoint name.
     * 
     */
    public Output<String> serviceEndpointName() {
        return this.serviceEndpointName;
    }

    /**
     * The username used to authenticate to the server url using basic authentication.
     * 
     */
    @Import(name="username")
    private @Nullable Output<String> username;

    /**
     * @return The username used to authenticate to the server url using basic authentication.
     * 
     */
    public Optional<Output<String>> username() {
        return Optional.ofNullable(this.username);
    }

    private ServiceEndpointGenericArgs() {}

    private ServiceEndpointGenericArgs(ServiceEndpointGenericArgs $) {
        this.description = $.description;
        this.password = $.password;
        this.projectId = $.projectId;
        this.serverUrl = $.serverUrl;
        this.serviceEndpointName = $.serviceEndpointName;
        this.username = $.username;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServiceEndpointGenericArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServiceEndpointGenericArgs $;

        public Builder() {
            $ = new ServiceEndpointGenericArgs();
        }

        public Builder(ServiceEndpointGenericArgs defaults) {
            $ = new ServiceEndpointGenericArgs(Objects.requireNonNull(defaults));
        }

        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param password The password or token key used to authenticate to the server url using basic authentication.
         * 
         * @return builder
         * 
         */
        public Builder password(@Nullable Output<String> password) {
            $.password = password;
            return this;
        }

        /**
         * @param password The password or token key used to authenticate to the server url using basic authentication.
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
         * @param serverUrl The URL of the server associated with the service endpoint.
         * 
         * @return builder
         * 
         */
        public Builder serverUrl(Output<String> serverUrl) {
            $.serverUrl = serverUrl;
            return this;
        }

        /**
         * @param serverUrl The URL of the server associated with the service endpoint.
         * 
         * @return builder
         * 
         */
        public Builder serverUrl(String serverUrl) {
            return serverUrl(Output.of(serverUrl));
        }

        /**
         * @param serviceEndpointName The service endpoint name.
         * 
         * @return builder
         * 
         */
        public Builder serviceEndpointName(Output<String> serviceEndpointName) {
            $.serviceEndpointName = serviceEndpointName;
            return this;
        }

        /**
         * @param serviceEndpointName The service endpoint name.
         * 
         * @return builder
         * 
         */
        public Builder serviceEndpointName(String serviceEndpointName) {
            return serviceEndpointName(Output.of(serviceEndpointName));
        }

        /**
         * @param username The username used to authenticate to the server url using basic authentication.
         * 
         * @return builder
         * 
         */
        public Builder username(@Nullable Output<String> username) {
            $.username = username;
            return this;
        }

        /**
         * @param username The username used to authenticate to the server url using basic authentication.
         * 
         * @return builder
         * 
         */
        public Builder username(String username) {
            return username(Output.of(username));
        }

        public ServiceEndpointGenericArgs build() {
            if ($.projectId == null) {
                throw new MissingRequiredPropertyException("ServiceEndpointGenericArgs", "projectId");
            }
            if ($.serverUrl == null) {
                throw new MissingRequiredPropertyException("ServiceEndpointGenericArgs", "serverUrl");
            }
            if ($.serviceEndpointName == null) {
                throw new MissingRequiredPropertyException("ServiceEndpointGenericArgs", "serviceEndpointName");
            }
            return $;
        }
    }

}
