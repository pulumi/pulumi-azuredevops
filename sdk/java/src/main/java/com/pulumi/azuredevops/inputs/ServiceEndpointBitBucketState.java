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


public final class ServiceEndpointBitBucketState extends com.pulumi.resources.ResourceArgs {

    public static final ServiceEndpointBitBucketState Empty = new ServiceEndpointBitBucketState();

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
     * Bitbucket account password.
     * 
     */
    @Import(name="password")
    private @Nullable Output<String> password;

    /**
     * @return Bitbucket account password.
     * 
     */
    public Optional<Output<String>> password() {
        return Optional.ofNullable(this.password);
    }

    /**
     * A bcrypted hash of the attribute &#39;password&#39;
     * 
     */
    @Import(name="passwordHash")
    private @Nullable Output<String> passwordHash;

    /**
     * @return A bcrypted hash of the attribute &#39;password&#39;
     * 
     */
    public Optional<Output<String>> passwordHash() {
        return Optional.ofNullable(this.passwordHash);
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
     * Bitbucket account username.
     * 
     */
    @Import(name="username")
    private @Nullable Output<String> username;

    /**
     * @return Bitbucket account username.
     * 
     */
    public Optional<Output<String>> username() {
        return Optional.ofNullable(this.username);
    }

    private ServiceEndpointBitBucketState() {}

    private ServiceEndpointBitBucketState(ServiceEndpointBitBucketState $) {
        this.authorization = $.authorization;
        this.description = $.description;
        this.password = $.password;
        this.passwordHash = $.passwordHash;
        this.projectId = $.projectId;
        this.serviceEndpointName = $.serviceEndpointName;
        this.username = $.username;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServiceEndpointBitBucketState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServiceEndpointBitBucketState $;

        public Builder() {
            $ = new ServiceEndpointBitBucketState();
        }

        public Builder(ServiceEndpointBitBucketState defaults) {
            $ = new ServiceEndpointBitBucketState(Objects.requireNonNull(defaults));
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
         * @param password Bitbucket account password.
         * 
         * @return builder
         * 
         */
        public Builder password(@Nullable Output<String> password) {
            $.password = password;
            return this;
        }

        /**
         * @param password Bitbucket account password.
         * 
         * @return builder
         * 
         */
        public Builder password(String password) {
            return password(Output.of(password));
        }

        /**
         * @param passwordHash A bcrypted hash of the attribute &#39;password&#39;
         * 
         * @return builder
         * 
         */
        public Builder passwordHash(@Nullable Output<String> passwordHash) {
            $.passwordHash = passwordHash;
            return this;
        }

        /**
         * @param passwordHash A bcrypted hash of the attribute &#39;password&#39;
         * 
         * @return builder
         * 
         */
        public Builder passwordHash(String passwordHash) {
            return passwordHash(Output.of(passwordHash));
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
         * @param username Bitbucket account username.
         * 
         * @return builder
         * 
         */
        public Builder username(@Nullable Output<String> username) {
            $.username = username;
            return this;
        }

        /**
         * @param username Bitbucket account username.
         * 
         * @return builder
         * 
         */
        public Builder username(String username) {
            return username(Output.of(username));
        }

        public ServiceEndpointBitBucketState build() {
            return $;
        }
    }

}
