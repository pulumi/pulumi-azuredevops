// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GitInitializationArgs extends com.pulumi.resources.ResourceArgs {

    public static final GitInitializationArgs Empty = new GitInitializationArgs();

    /**
     * The type of repository to create. Valid values: `Uninitialized`, `Clean` or `Import`.
     * 
     */
    @Import(name="initType", required=true)
    private Output<String> initType;

    /**
     * @return The type of repository to create. Valid values: `Uninitialized`, `Clean` or `Import`.
     * 
     */
    public Output<String> initType() {
        return this.initType;
    }

    /**
     * The ID of service connection used to authenticate to a private repository for import initialization. Conflicts with `username` and `password`.
     * 
     */
    @Import(name="serviceConnectionId")
    private @Nullable Output<String> serviceConnectionId;

    /**
     * @return The ID of service connection used to authenticate to a private repository for import initialization. Conflicts with `username` and `password`.
     * 
     */
    public Optional<Output<String>> serviceConnectionId() {
        return Optional.ofNullable(this.serviceConnectionId);
    }

    /**
     * Type of the source repository. Used if the `init_type` is `Import`. Valid values: `Git`.
     * 
     */
    @Import(name="sourceType")
    private @Nullable Output<String> sourceType;

    /**
     * @return Type of the source repository. Used if the `init_type` is `Import`. Valid values: `Git`.
     * 
     */
    public Optional<Output<String>> sourceType() {
        return Optional.ofNullable(this.sourceType);
    }

    /**
     * The URL of the source repository. Used if the `init_type` is `Import`.
     * 
     */
    @Import(name="sourceUrl")
    private @Nullable Output<String> sourceUrl;

    /**
     * @return The URL of the source repository. Used if the `init_type` is `Import`.
     * 
     */
    public Optional<Output<String>> sourceUrl() {
        return Optional.ofNullable(this.sourceUrl);
    }

    /**
     * The username used to authenticate to a private repository for import initialization. Conflicts with `service_connection_id`.
     * 
     */
    @Import(name="username")
    private @Nullable Output<String> username;

    /**
     * @return The username used to authenticate to a private repository for import initialization. Conflicts with `service_connection_id`.
     * 
     */
    public Optional<Output<String>> username() {
        return Optional.ofNullable(this.username);
    }

    private GitInitializationArgs() {}

    private GitInitializationArgs(GitInitializationArgs $) {
        this.initType = $.initType;
        this.serviceConnectionId = $.serviceConnectionId;
        this.sourceType = $.sourceType;
        this.sourceUrl = $.sourceUrl;
        this.username = $.username;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GitInitializationArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GitInitializationArgs $;

        public Builder() {
            $ = new GitInitializationArgs();
        }

        public Builder(GitInitializationArgs defaults) {
            $ = new GitInitializationArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param initType The type of repository to create. Valid values: `Uninitialized`, `Clean` or `Import`.
         * 
         * @return builder
         * 
         */
        public Builder initType(Output<String> initType) {
            $.initType = initType;
            return this;
        }

        /**
         * @param initType The type of repository to create. Valid values: `Uninitialized`, `Clean` or `Import`.
         * 
         * @return builder
         * 
         */
        public Builder initType(String initType) {
            return initType(Output.of(initType));
        }

        /**
         * @param serviceConnectionId The ID of service connection used to authenticate to a private repository for import initialization. Conflicts with `username` and `password`.
         * 
         * @return builder
         * 
         */
        public Builder serviceConnectionId(@Nullable Output<String> serviceConnectionId) {
            $.serviceConnectionId = serviceConnectionId;
            return this;
        }

        /**
         * @param serviceConnectionId The ID of service connection used to authenticate to a private repository for import initialization. Conflicts with `username` and `password`.
         * 
         * @return builder
         * 
         */
        public Builder serviceConnectionId(String serviceConnectionId) {
            return serviceConnectionId(Output.of(serviceConnectionId));
        }

        /**
         * @param sourceType Type of the source repository. Used if the `init_type` is `Import`. Valid values: `Git`.
         * 
         * @return builder
         * 
         */
        public Builder sourceType(@Nullable Output<String> sourceType) {
            $.sourceType = sourceType;
            return this;
        }

        /**
         * @param sourceType Type of the source repository. Used if the `init_type` is `Import`. Valid values: `Git`.
         * 
         * @return builder
         * 
         */
        public Builder sourceType(String sourceType) {
            return sourceType(Output.of(sourceType));
        }

        /**
         * @param sourceUrl The URL of the source repository. Used if the `init_type` is `Import`.
         * 
         * @return builder
         * 
         */
        public Builder sourceUrl(@Nullable Output<String> sourceUrl) {
            $.sourceUrl = sourceUrl;
            return this;
        }

        /**
         * @param sourceUrl The URL of the source repository. Used if the `init_type` is `Import`.
         * 
         * @return builder
         * 
         */
        public Builder sourceUrl(String sourceUrl) {
            return sourceUrl(Output.of(sourceUrl));
        }

        /**
         * @param username The username used to authenticate to a private repository for import initialization. Conflicts with `service_connection_id`.
         * 
         * @return builder
         * 
         */
        public Builder username(@Nullable Output<String> username) {
            $.username = username;
            return this;
        }

        /**
         * @param username The username used to authenticate to a private repository for import initialization. Conflicts with `service_connection_id`.
         * 
         * @return builder
         * 
         */
        public Builder username(String username) {
            return username(Output.of(username));
        }

        public GitInitializationArgs build() {
            if ($.initType == null) {
                throw new MissingRequiredPropertyException("GitInitializationArgs", "initType");
            }
            return $;
        }
    }

}
