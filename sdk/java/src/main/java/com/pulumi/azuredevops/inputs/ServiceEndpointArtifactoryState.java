// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.inputs;

import com.pulumi.azuredevops.inputs.ServiceEndpointArtifactoryAuthenticationBasicArgs;
import com.pulumi.azuredevops.inputs.ServiceEndpointArtifactoryAuthenticationTokenArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ServiceEndpointArtifactoryState extends com.pulumi.resources.ResourceArgs {

    public static final ServiceEndpointArtifactoryState Empty = new ServiceEndpointArtifactoryState();

    @Import(name="authenticationBasic")
    private @Nullable Output<ServiceEndpointArtifactoryAuthenticationBasicArgs> authenticationBasic;

    public Optional<Output<ServiceEndpointArtifactoryAuthenticationBasicArgs>> authenticationBasic() {
        return Optional.ofNullable(this.authenticationBasic);
    }

    /**
     * A `authentication_basic` block as defined below.
     * 
     */
    @Import(name="authenticationToken")
    private @Nullable Output<ServiceEndpointArtifactoryAuthenticationTokenArgs> authenticationToken;

    /**
     * @return A `authentication_basic` block as defined below.
     * 
     */
    public Optional<Output<ServiceEndpointArtifactoryAuthenticationTokenArgs>> authenticationToken() {
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
     * URL of the Artifactory server to connect with.
     * 
     * _**Note: URL should not end in a slash character.**_
     * 
     */
    @Import(name="url")
    private @Nullable Output<String> url;

    /**
     * @return URL of the Artifactory server to connect with.
     * 
     * _**Note: URL should not end in a slash character.**_
     * 
     */
    public Optional<Output<String>> url() {
        return Optional.ofNullable(this.url);
    }

    private ServiceEndpointArtifactoryState() {}

    private ServiceEndpointArtifactoryState(ServiceEndpointArtifactoryState $) {
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
    public static Builder builder(ServiceEndpointArtifactoryState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServiceEndpointArtifactoryState $;

        public Builder() {
            $ = new ServiceEndpointArtifactoryState();
        }

        public Builder(ServiceEndpointArtifactoryState defaults) {
            $ = new ServiceEndpointArtifactoryState(Objects.requireNonNull(defaults));
        }

        public Builder authenticationBasic(@Nullable Output<ServiceEndpointArtifactoryAuthenticationBasicArgs> authenticationBasic) {
            $.authenticationBasic = authenticationBasic;
            return this;
        }

        public Builder authenticationBasic(ServiceEndpointArtifactoryAuthenticationBasicArgs authenticationBasic) {
            return authenticationBasic(Output.of(authenticationBasic));
        }

        /**
         * @param authenticationToken A `authentication_basic` block as defined below.
         * 
         * @return builder
         * 
         */
        public Builder authenticationToken(@Nullable Output<ServiceEndpointArtifactoryAuthenticationTokenArgs> authenticationToken) {
            $.authenticationToken = authenticationToken;
            return this;
        }

        /**
         * @param authenticationToken A `authentication_basic` block as defined below.
         * 
         * @return builder
         * 
         */
        public Builder authenticationToken(ServiceEndpointArtifactoryAuthenticationTokenArgs authenticationToken) {
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
         * @param url URL of the Artifactory server to connect with.
         * 
         * _**Note: URL should not end in a slash character.**_
         * 
         * @return builder
         * 
         */
        public Builder url(@Nullable Output<String> url) {
            $.url = url;
            return this;
        }

        /**
         * @param url URL of the Artifactory server to connect with.
         * 
         * _**Note: URL should not end in a slash character.**_
         * 
         * @return builder
         * 
         */
        public Builder url(String url) {
            return url(Output.of(url));
        }

        public ServiceEndpointArtifactoryState build() {
            return $;
        }
    }

}
