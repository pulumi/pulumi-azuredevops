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


public final class ServiceEndpointSonarQubeState extends com.pulumi.resources.ResourceArgs {

    public static final ServiceEndpointSonarQubeState Empty = new ServiceEndpointSonarQubeState();

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
     * Authentication Token generated through SonarQube (go to My Account &gt; Security &gt; Generate Tokens).
     * 
     */
    @Import(name="token")
    private @Nullable Output<String> token;

    /**
     * @return Authentication Token generated through SonarQube (go to My Account &gt; Security &gt; Generate Tokens).
     * 
     */
    public Optional<Output<String>> token() {
        return Optional.ofNullable(this.token);
    }

    /**
     * A bcrypted hash of the attribute &#39;token&#39;
     * 
     */
    @Import(name="tokenHash")
    private @Nullable Output<String> tokenHash;

    /**
     * @return A bcrypted hash of the attribute &#39;token&#39;
     * 
     */
    public Optional<Output<String>> tokenHash() {
        return Optional.ofNullable(this.tokenHash);
    }

    /**
     * URL of the SonarQube server to connect with.
     * 
     */
    @Import(name="url")
    private @Nullable Output<String> url;

    /**
     * @return URL of the SonarQube server to connect with.
     * 
     */
    public Optional<Output<String>> url() {
        return Optional.ofNullable(this.url);
    }

    private ServiceEndpointSonarQubeState() {}

    private ServiceEndpointSonarQubeState(ServiceEndpointSonarQubeState $) {
        this.authorization = $.authorization;
        this.description = $.description;
        this.projectId = $.projectId;
        this.serviceEndpointName = $.serviceEndpointName;
        this.token = $.token;
        this.tokenHash = $.tokenHash;
        this.url = $.url;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServiceEndpointSonarQubeState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServiceEndpointSonarQubeState $;

        public Builder() {
            $ = new ServiceEndpointSonarQubeState();
        }

        public Builder(ServiceEndpointSonarQubeState defaults) {
            $ = new ServiceEndpointSonarQubeState(Objects.requireNonNull(defaults));
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
         * @param token Authentication Token generated through SonarQube (go to My Account &gt; Security &gt; Generate Tokens).
         * 
         * @return builder
         * 
         */
        public Builder token(@Nullable Output<String> token) {
            $.token = token;
            return this;
        }

        /**
         * @param token Authentication Token generated through SonarQube (go to My Account &gt; Security &gt; Generate Tokens).
         * 
         * @return builder
         * 
         */
        public Builder token(String token) {
            return token(Output.of(token));
        }

        /**
         * @param tokenHash A bcrypted hash of the attribute &#39;token&#39;
         * 
         * @return builder
         * 
         */
        public Builder tokenHash(@Nullable Output<String> tokenHash) {
            $.tokenHash = tokenHash;
            return this;
        }

        /**
         * @param tokenHash A bcrypted hash of the attribute &#39;token&#39;
         * 
         * @return builder
         * 
         */
        public Builder tokenHash(String tokenHash) {
            return tokenHash(Output.of(tokenHash));
        }

        /**
         * @param url URL of the SonarQube server to connect with.
         * 
         * @return builder
         * 
         */
        public Builder url(@Nullable Output<String> url) {
            $.url = url;
            return this;
        }

        /**
         * @param url URL of the SonarQube server to connect with.
         * 
         * @return builder
         * 
         */
        public Builder url(String url) {
            return url(Output.of(url));
        }

        public ServiceEndpointSonarQubeState build() {
            return $;
        }
    }

}
