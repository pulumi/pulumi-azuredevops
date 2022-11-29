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


public final class ServiceendpointIncomingwebhookArgs extends com.pulumi.resources.ResourceArgs {

    public static final ServiceendpointIncomingwebhookArgs Empty = new ServiceendpointIncomingwebhookArgs();

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
     * Http header name on which checksum will be sent.
     * 
     */
    @Import(name="httpHeader")
    private @Nullable Output<String> httpHeader;

    /**
     * @return Http header name on which checksum will be sent.
     * 
     */
    public Optional<Output<String>> httpHeader() {
        return Optional.ofNullable(this.httpHeader);
    }

    /**
     * The ID of the project. Changing this forces a new Service Connection Incoming WebHook to be created.
     * 
     */
    @Import(name="projectId", required=true)
    private Output<String> projectId;

    /**
     * @return The ID of the project. Changing this forces a new Service Connection Incoming WebHook to be created.
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }

    /**
     * Secret for the WebHook. WebHook service will use this secret to calculate the payload checksum.
     * 
     */
    @Import(name="secret")
    private @Nullable Output<String> secret;

    /**
     * @return Secret for the WebHook. WebHook service will use this secret to calculate the payload checksum.
     * 
     */
    public Optional<Output<String>> secret() {
        return Optional.ofNullable(this.secret);
    }

    /**
     * The name of the service endpoint. Changing this forces a new Service Connection Incoming WebHook to be created.
     * 
     */
    @Import(name="serviceEndpointName", required=true)
    private Output<String> serviceEndpointName;

    /**
     * @return The name of the service endpoint. Changing this forces a new Service Connection Incoming WebHook to be created.
     * 
     */
    public Output<String> serviceEndpointName() {
        return this.serviceEndpointName;
    }

    /**
     * The name of the WebHook.
     * 
     */
    @Import(name="webhookName", required=true)
    private Output<String> webhookName;

    /**
     * @return The name of the WebHook.
     * 
     */
    public Output<String> webhookName() {
        return this.webhookName;
    }

    private ServiceendpointIncomingwebhookArgs() {}

    private ServiceendpointIncomingwebhookArgs(ServiceendpointIncomingwebhookArgs $) {
        this.authorization = $.authorization;
        this.description = $.description;
        this.httpHeader = $.httpHeader;
        this.projectId = $.projectId;
        this.secret = $.secret;
        this.serviceEndpointName = $.serviceEndpointName;
        this.webhookName = $.webhookName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServiceendpointIncomingwebhookArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServiceendpointIncomingwebhookArgs $;

        public Builder() {
            $ = new ServiceendpointIncomingwebhookArgs();
        }

        public Builder(ServiceendpointIncomingwebhookArgs defaults) {
            $ = new ServiceendpointIncomingwebhookArgs(Objects.requireNonNull(defaults));
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
         * @param httpHeader Http header name on which checksum will be sent.
         * 
         * @return builder
         * 
         */
        public Builder httpHeader(@Nullable Output<String> httpHeader) {
            $.httpHeader = httpHeader;
            return this;
        }

        /**
         * @param httpHeader Http header name on which checksum will be sent.
         * 
         * @return builder
         * 
         */
        public Builder httpHeader(String httpHeader) {
            return httpHeader(Output.of(httpHeader));
        }

        /**
         * @param projectId The ID of the project. Changing this forces a new Service Connection Incoming WebHook to be created.
         * 
         * @return builder
         * 
         */
        public Builder projectId(Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param projectId The ID of the project. Changing this forces a new Service Connection Incoming WebHook to be created.
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
        }

        /**
         * @param secret Secret for the WebHook. WebHook service will use this secret to calculate the payload checksum.
         * 
         * @return builder
         * 
         */
        public Builder secret(@Nullable Output<String> secret) {
            $.secret = secret;
            return this;
        }

        /**
         * @param secret Secret for the WebHook. WebHook service will use this secret to calculate the payload checksum.
         * 
         * @return builder
         * 
         */
        public Builder secret(String secret) {
            return secret(Output.of(secret));
        }

        /**
         * @param serviceEndpointName The name of the service endpoint. Changing this forces a new Service Connection Incoming WebHook to be created.
         * 
         * @return builder
         * 
         */
        public Builder serviceEndpointName(Output<String> serviceEndpointName) {
            $.serviceEndpointName = serviceEndpointName;
            return this;
        }

        /**
         * @param serviceEndpointName The name of the service endpoint. Changing this forces a new Service Connection Incoming WebHook to be created.
         * 
         * @return builder
         * 
         */
        public Builder serviceEndpointName(String serviceEndpointName) {
            return serviceEndpointName(Output.of(serviceEndpointName));
        }

        /**
         * @param webhookName The name of the WebHook.
         * 
         * @return builder
         * 
         */
        public Builder webhookName(Output<String> webhookName) {
            $.webhookName = webhookName;
            return this;
        }

        /**
         * @param webhookName The name of the WebHook.
         * 
         * @return builder
         * 
         */
        public Builder webhookName(String webhookName) {
            return webhookName(Output.of(webhookName));
        }

        public ServiceendpointIncomingwebhookArgs build() {
            $.projectId = Objects.requireNonNull($.projectId, "expected parameter 'projectId' to be non-null");
            $.serviceEndpointName = Objects.requireNonNull($.serviceEndpointName, "expected parameter 'serviceEndpointName' to be non-null");
            $.webhookName = Objects.requireNonNull($.webhookName, "expected parameter 'webhookName' to be non-null");
            return $;
        }
    }

}
