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


public final class ServiceEndpointSonarQubeArgs extends com.pulumi.resources.ResourceArgs {

    public static final ServiceEndpointSonarQubeArgs Empty = new ServiceEndpointSonarQubeArgs();

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
     * The Authentication Token generated through SonarQube (go to My Account &gt; Security &gt; Generate Tokens).
     * 
     */
    @Import(name="token", required=true)
    private Output<String> token;

    /**
     * @return The Authentication Token generated through SonarQube (go to My Account &gt; Security &gt; Generate Tokens).
     * 
     */
    public Output<String> token() {
        return this.token;
    }

    /**
     * URL of the SonarQube server to connect with.
     * 
     */
    @Import(name="url", required=true)
    private Output<String> url;

    /**
     * @return URL of the SonarQube server to connect with.
     * 
     */
    public Output<String> url() {
        return this.url;
    }

    private ServiceEndpointSonarQubeArgs() {}

    private ServiceEndpointSonarQubeArgs(ServiceEndpointSonarQubeArgs $) {
        this.description = $.description;
        this.projectId = $.projectId;
        this.serviceEndpointName = $.serviceEndpointName;
        this.token = $.token;
        this.url = $.url;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServiceEndpointSonarQubeArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServiceEndpointSonarQubeArgs $;

        public Builder() {
            $ = new ServiceEndpointSonarQubeArgs();
        }

        public Builder(ServiceEndpointSonarQubeArgs defaults) {
            $ = new ServiceEndpointSonarQubeArgs(Objects.requireNonNull(defaults));
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
         * @param token The Authentication Token generated through SonarQube (go to My Account &gt; Security &gt; Generate Tokens).
         * 
         * @return builder
         * 
         */
        public Builder token(Output<String> token) {
            $.token = token;
            return this;
        }

        /**
         * @param token The Authentication Token generated through SonarQube (go to My Account &gt; Security &gt; Generate Tokens).
         * 
         * @return builder
         * 
         */
        public Builder token(String token) {
            return token(Output.of(token));
        }

        /**
         * @param url URL of the SonarQube server to connect with.
         * 
         * @return builder
         * 
         */
        public Builder url(Output<String> url) {
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

        public ServiceEndpointSonarQubeArgs build() {
            if ($.projectId == null) {
                throw new MissingRequiredPropertyException("ServiceEndpointSonarQubeArgs", "projectId");
            }
            if ($.serviceEndpointName == null) {
                throw new MissingRequiredPropertyException("ServiceEndpointSonarQubeArgs", "serviceEndpointName");
            }
            if ($.token == null) {
                throw new MissingRequiredPropertyException("ServiceEndpointSonarQubeArgs", "token");
            }
            if ($.url == null) {
                throw new MissingRequiredPropertyException("ServiceEndpointSonarQubeArgs", "url");
            }
            return $;
        }
    }

}
