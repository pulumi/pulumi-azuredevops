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


public final class ServiceendpointNexusArgs extends com.pulumi.resources.ResourceArgs {

    public static final ServiceendpointNexusArgs Empty = new ServiceendpointNexusArgs();

    @Import(name="description")
    private @Nullable Output<String> description;

    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The Service Endpoint password to authenticate at the Nexus IQ Instance.
     * 
     */
    @Import(name="password", required=true)
    private Output<String> password;

    /**
     * @return The Service Endpoint password to authenticate at the Nexus IQ Instance.
     * 
     */
    public Output<String> password() {
        return this.password;
    }

    /**
     * The ID of the project. Changing this forces a new Service Connection Nexus to be created.
     * 
     */
    @Import(name="projectId", required=true)
    private Output<String> projectId;

    /**
     * @return The ID of the project. Changing this forces a new Service Connection Nexus to be created.
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }

    /**
     * The name of the service endpoint. Changing this forces a new Service Connection Nexus to be created.
     * 
     */
    @Import(name="serviceEndpointName", required=true)
    private Output<String> serviceEndpointName;

    /**
     * @return The name of the service endpoint. Changing this forces a new Service Connection Nexus to be created.
     * 
     */
    public Output<String> serviceEndpointName() {
        return this.serviceEndpointName;
    }

    /**
     * The Service Endpoint url.
     * 
     */
    @Import(name="url", required=true)
    private Output<String> url;

    /**
     * @return The Service Endpoint url.
     * 
     */
    public Output<String> url() {
        return this.url;
    }

    /**
     * The Service Endpoint username to authenticate at the Nexus IQ Instance.
     * 
     */
    @Import(name="username", required=true)
    private Output<String> username;

    /**
     * @return The Service Endpoint username to authenticate at the Nexus IQ Instance.
     * 
     */
    public Output<String> username() {
        return this.username;
    }

    private ServiceendpointNexusArgs() {}

    private ServiceendpointNexusArgs(ServiceendpointNexusArgs $) {
        this.description = $.description;
        this.password = $.password;
        this.projectId = $.projectId;
        this.serviceEndpointName = $.serviceEndpointName;
        this.url = $.url;
        this.username = $.username;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServiceendpointNexusArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServiceendpointNexusArgs $;

        public Builder() {
            $ = new ServiceendpointNexusArgs();
        }

        public Builder(ServiceendpointNexusArgs defaults) {
            $ = new ServiceendpointNexusArgs(Objects.requireNonNull(defaults));
        }

        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param password The Service Endpoint password to authenticate at the Nexus IQ Instance.
         * 
         * @return builder
         * 
         */
        public Builder password(Output<String> password) {
            $.password = password;
            return this;
        }

        /**
         * @param password The Service Endpoint password to authenticate at the Nexus IQ Instance.
         * 
         * @return builder
         * 
         */
        public Builder password(String password) {
            return password(Output.of(password));
        }

        /**
         * @param projectId The ID of the project. Changing this forces a new Service Connection Nexus to be created.
         * 
         * @return builder
         * 
         */
        public Builder projectId(Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param projectId The ID of the project. Changing this forces a new Service Connection Nexus to be created.
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
        }

        /**
         * @param serviceEndpointName The name of the service endpoint. Changing this forces a new Service Connection Nexus to be created.
         * 
         * @return builder
         * 
         */
        public Builder serviceEndpointName(Output<String> serviceEndpointName) {
            $.serviceEndpointName = serviceEndpointName;
            return this;
        }

        /**
         * @param serviceEndpointName The name of the service endpoint. Changing this forces a new Service Connection Nexus to be created.
         * 
         * @return builder
         * 
         */
        public Builder serviceEndpointName(String serviceEndpointName) {
            return serviceEndpointName(Output.of(serviceEndpointName));
        }

        /**
         * @param url The Service Endpoint url.
         * 
         * @return builder
         * 
         */
        public Builder url(Output<String> url) {
            $.url = url;
            return this;
        }

        /**
         * @param url The Service Endpoint url.
         * 
         * @return builder
         * 
         */
        public Builder url(String url) {
            return url(Output.of(url));
        }

        /**
         * @param username The Service Endpoint username to authenticate at the Nexus IQ Instance.
         * 
         * @return builder
         * 
         */
        public Builder username(Output<String> username) {
            $.username = username;
            return this;
        }

        /**
         * @param username The Service Endpoint username to authenticate at the Nexus IQ Instance.
         * 
         * @return builder
         * 
         */
        public Builder username(String username) {
            return username(Output.of(username));
        }

        public ServiceendpointNexusArgs build() {
            if ($.password == null) {
                throw new MissingRequiredPropertyException("ServiceendpointNexusArgs", "password");
            }
            if ($.projectId == null) {
                throw new MissingRequiredPropertyException("ServiceendpointNexusArgs", "projectId");
            }
            if ($.serviceEndpointName == null) {
                throw new MissingRequiredPropertyException("ServiceendpointNexusArgs", "serviceEndpointName");
            }
            if ($.url == null) {
                throw new MissingRequiredPropertyException("ServiceendpointNexusArgs", "url");
            }
            if ($.username == null) {
                throw new MissingRequiredPropertyException("ServiceendpointNexusArgs", "username");
            }
            return $;
        }
    }

}
