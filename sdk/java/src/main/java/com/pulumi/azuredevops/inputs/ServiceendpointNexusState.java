// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ServiceendpointNexusState extends com.pulumi.resources.ResourceArgs {

    public static final ServiceendpointNexusState Empty = new ServiceendpointNexusState();

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
     * The Service Endpoint password to authenticate at the Nexus IQ Instance.
     * 
     */
    @Import(name="password")
    private @Nullable Output<String> password;

    /**
     * @return The Service Endpoint password to authenticate at the Nexus IQ Instance.
     * 
     */
    public Optional<Output<String>> password() {
        return Optional.ofNullable(this.password);
    }

    /**
     * The ID of the project. Changing this forces a new Service Connection Nexus to be created.
     * 
     */
    @Import(name="projectId")
    private @Nullable Output<String> projectId;

    /**
     * @return The ID of the project. Changing this forces a new Service Connection Nexus to be created.
     * 
     */
    public Optional<Output<String>> projectId() {
        return Optional.ofNullable(this.projectId);
    }

    /**
     * The name of the service endpoint. Changing this forces a new Service Connection Nexus to be created.
     * 
     */
    @Import(name="serviceEndpointName")
    private @Nullable Output<String> serviceEndpointName;

    /**
     * @return The name of the service endpoint. Changing this forces a new Service Connection Nexus to be created.
     * 
     */
    public Optional<Output<String>> serviceEndpointName() {
        return Optional.ofNullable(this.serviceEndpointName);
    }

    /**
     * The Service Endpoint url.
     * 
     */
    @Import(name="url")
    private @Nullable Output<String> url;

    /**
     * @return The Service Endpoint url.
     * 
     */
    public Optional<Output<String>> url() {
        return Optional.ofNullable(this.url);
    }

    /**
     * The Service Endpoint username to authenticate at the Nexus IQ Instance.
     * 
     */
    @Import(name="username")
    private @Nullable Output<String> username;

    /**
     * @return The Service Endpoint username to authenticate at the Nexus IQ Instance.
     * 
     */
    public Optional<Output<String>> username() {
        return Optional.ofNullable(this.username);
    }

    private ServiceendpointNexusState() {}

    private ServiceendpointNexusState(ServiceendpointNexusState $) {
        this.authorization = $.authorization;
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
    public static Builder builder(ServiceendpointNexusState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServiceendpointNexusState $;

        public Builder() {
            $ = new ServiceendpointNexusState();
        }

        public Builder(ServiceendpointNexusState defaults) {
            $ = new ServiceendpointNexusState(Objects.requireNonNull(defaults));
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
         * @param password The Service Endpoint password to authenticate at the Nexus IQ Instance.
         * 
         * @return builder
         * 
         */
        public Builder password(@Nullable Output<String> password) {
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
        public Builder projectId(@Nullable Output<String> projectId) {
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
        public Builder serviceEndpointName(@Nullable Output<String> serviceEndpointName) {
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
        public Builder url(@Nullable Output<String> url) {
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
        public Builder username(@Nullable Output<String> username) {
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

        public ServiceendpointNexusState build() {
            return $;
        }
    }

}