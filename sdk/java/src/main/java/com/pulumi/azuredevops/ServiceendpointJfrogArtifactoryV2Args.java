// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops;

import com.pulumi.azuredevops.inputs.ServiceendpointJfrogArtifactoryV2AuthenticationBasicArgs;
import com.pulumi.azuredevops.inputs.ServiceendpointJfrogArtifactoryV2AuthenticationTokenArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ServiceendpointJfrogArtifactoryV2Args extends com.pulumi.resources.ResourceArgs {

    public static final ServiceendpointJfrogArtifactoryV2Args Empty = new ServiceendpointJfrogArtifactoryV2Args();

    /**
     * A `authentication_basic` block as documented below.
     * 
     */
    @Import(name="authenticationBasic")
    private @Nullable Output<ServiceendpointJfrogArtifactoryV2AuthenticationBasicArgs> authenticationBasic;

    /**
     * @return A `authentication_basic` block as documented below.
     * 
     */
    public Optional<Output<ServiceendpointJfrogArtifactoryV2AuthenticationBasicArgs>> authenticationBasic() {
        return Optional.ofNullable(this.authenticationBasic);
    }

    /**
     * A `authentication_token` block as documented below.
     * 
     */
    @Import(name="authenticationToken")
    private @Nullable Output<ServiceendpointJfrogArtifactoryV2AuthenticationTokenArgs> authenticationToken;

    /**
     * @return A `authentication_token` block as documented below.
     * 
     */
    public Optional<Output<ServiceendpointJfrogArtifactoryV2AuthenticationTokenArgs>> authenticationToken() {
        return Optional.ofNullable(this.authenticationToken);
    }

    @Import(name="authorization")
    private @Nullable Output<Map<String,String>> authorization;

    public Optional<Output<Map<String,String>>> authorization() {
        return Optional.ofNullable(this.authorization);
    }

    /**
     * The Service Endpoint description.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return The Service Endpoint description.
     * 
     */
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
     * URL of the Artifactory server to connect with.
     * 
     */
    @Import(name="url", required=true)
    private Output<String> url;

    /**
     * @return URL of the Artifactory server to connect with.
     * 
     */
    public Output<String> url() {
        return this.url;
    }

    private ServiceendpointJfrogArtifactoryV2Args() {}

    private ServiceendpointJfrogArtifactoryV2Args(ServiceendpointJfrogArtifactoryV2Args $) {
        this.authenticationBasic = $.authenticationBasic;
        this.authenticationToken = $.authenticationToken;
        this.authorization = $.authorization;
        this.description = $.description;
        this.projectId = $.projectId;
        this.serviceEndpointName = $.serviceEndpointName;
        this.url = $.url;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServiceendpointJfrogArtifactoryV2Args defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServiceendpointJfrogArtifactoryV2Args $;

        public Builder() {
            $ = new ServiceendpointJfrogArtifactoryV2Args();
        }

        public Builder(ServiceendpointJfrogArtifactoryV2Args defaults) {
            $ = new ServiceendpointJfrogArtifactoryV2Args(Objects.requireNonNull(defaults));
        }

        /**
         * @param authenticationBasic A `authentication_basic` block as documented below.
         * 
         * @return builder
         * 
         */
        public Builder authenticationBasic(@Nullable Output<ServiceendpointJfrogArtifactoryV2AuthenticationBasicArgs> authenticationBasic) {
            $.authenticationBasic = authenticationBasic;
            return this;
        }

        /**
         * @param authenticationBasic A `authentication_basic` block as documented below.
         * 
         * @return builder
         * 
         */
        public Builder authenticationBasic(ServiceendpointJfrogArtifactoryV2AuthenticationBasicArgs authenticationBasic) {
            return authenticationBasic(Output.of(authenticationBasic));
        }

        /**
         * @param authenticationToken A `authentication_token` block as documented below.
         * 
         * @return builder
         * 
         */
        public Builder authenticationToken(@Nullable Output<ServiceendpointJfrogArtifactoryV2AuthenticationTokenArgs> authenticationToken) {
            $.authenticationToken = authenticationToken;
            return this;
        }

        /**
         * @param authenticationToken A `authentication_token` block as documented below.
         * 
         * @return builder
         * 
         */
        public Builder authenticationToken(ServiceendpointJfrogArtifactoryV2AuthenticationTokenArgs authenticationToken) {
            return authenticationToken(Output.of(authenticationToken));
        }

        public Builder authorization(@Nullable Output<Map<String,String>> authorization) {
            $.authorization = authorization;
            return this;
        }

        public Builder authorization(Map<String,String> authorization) {
            return authorization(Output.of(authorization));
        }

        /**
         * @param description The Service Endpoint description.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The Service Endpoint description.
         * 
         * @return builder
         * 
         */
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
         * @param url URL of the Artifactory server to connect with.
         * 
         * @return builder
         * 
         */
        public Builder url(Output<String> url) {
            $.url = url;
            return this;
        }

        /**
         * @param url URL of the Artifactory server to connect with.
         * 
         * @return builder
         * 
         */
        public Builder url(String url) {
            return url(Output.of(url));
        }

        public ServiceendpointJfrogArtifactoryV2Args build() {
            $.projectId = Objects.requireNonNull($.projectId, "expected parameter 'projectId' to be non-null");
            $.serviceEndpointName = Objects.requireNonNull($.serviceEndpointName, "expected parameter 'serviceEndpointName' to be non-null");
            $.url = Objects.requireNonNull($.url, "expected parameter 'url' to be non-null");
            return $;
        }
    }

}