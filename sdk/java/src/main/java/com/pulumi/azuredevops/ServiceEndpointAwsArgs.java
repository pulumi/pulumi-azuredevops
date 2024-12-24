// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ServiceEndpointAwsArgs extends com.pulumi.resources.ResourceArgs {

    public static final ServiceEndpointAwsArgs Empty = new ServiceEndpointAwsArgs();

    /**
     * The AWS access key ID for signing programmatic requests.
     * 
     */
    @Import(name="accessKeyId")
    private @Nullable Output<String> accessKeyId;

    /**
     * @return The AWS access key ID for signing programmatic requests.
     * 
     */
    public Optional<Output<String>> accessKeyId() {
        return Optional.ofNullable(this.accessKeyId);
    }

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
     * A unique identifier that is used by third parties when assuming roles in their customers&#39; accounts, aka cross-account role access.
     * 
     */
    @Import(name="externalId")
    private @Nullable Output<String> externalId;

    /**
     * @return A unique identifier that is used by third parties when assuming roles in their customers&#39; accounts, aka cross-account role access.
     * 
     */
    public Optional<Output<String>> externalId() {
        return Optional.ofNullable(this.externalId);
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
     * Optional identifier for the assumed role session.
     * 
     */
    @Import(name="roleSessionName")
    private @Nullable Output<String> roleSessionName;

    /**
     * @return Optional identifier for the assumed role session.
     * 
     */
    public Optional<Output<String>> roleSessionName() {
        return Optional.ofNullable(this.roleSessionName);
    }

    /**
     * The Amazon Resource Name (ARN) of the role to assume.
     * 
     */
    @Import(name="roleToAssume")
    private @Nullable Output<String> roleToAssume;

    /**
     * @return The Amazon Resource Name (ARN) of the role to assume.
     * 
     */
    public Optional<Output<String>> roleToAssume() {
        return Optional.ofNullable(this.roleToAssume);
    }

    /**
     * The AWS secret access key for signing programmatic requests.
     * 
     */
    @Import(name="secretAccessKey")
    private @Nullable Output<String> secretAccessKey;

    /**
     * @return The AWS secret access key for signing programmatic requests.
     * 
     */
    public Optional<Output<String>> secretAccessKey() {
        return Optional.ofNullable(this.secretAccessKey);
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
     * The AWS session token for signing programmatic requests.
     * 
     */
    @Import(name="sessionToken")
    private @Nullable Output<String> sessionToken;

    /**
     * @return The AWS session token for signing programmatic requests.
     * 
     */
    public Optional<Output<String>> sessionToken() {
        return Optional.ofNullable(this.sessionToken);
    }

    /**
     * Enable this to attempt getting credentials with OIDC token from Azure Devops.
     * 
     */
    @Import(name="useOidc")
    private @Nullable Output<Boolean> useOidc;

    /**
     * @return Enable this to attempt getting credentials with OIDC token from Azure Devops.
     * 
     */
    public Optional<Output<Boolean>> useOidc() {
        return Optional.ofNullable(this.useOidc);
    }

    private ServiceEndpointAwsArgs() {}

    private ServiceEndpointAwsArgs(ServiceEndpointAwsArgs $) {
        this.accessKeyId = $.accessKeyId;
        this.authorization = $.authorization;
        this.description = $.description;
        this.externalId = $.externalId;
        this.projectId = $.projectId;
        this.roleSessionName = $.roleSessionName;
        this.roleToAssume = $.roleToAssume;
        this.secretAccessKey = $.secretAccessKey;
        this.serviceEndpointName = $.serviceEndpointName;
        this.sessionToken = $.sessionToken;
        this.useOidc = $.useOidc;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServiceEndpointAwsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServiceEndpointAwsArgs $;

        public Builder() {
            $ = new ServiceEndpointAwsArgs();
        }

        public Builder(ServiceEndpointAwsArgs defaults) {
            $ = new ServiceEndpointAwsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param accessKeyId The AWS access key ID for signing programmatic requests.
         * 
         * @return builder
         * 
         */
        public Builder accessKeyId(@Nullable Output<String> accessKeyId) {
            $.accessKeyId = accessKeyId;
            return this;
        }

        /**
         * @param accessKeyId The AWS access key ID for signing programmatic requests.
         * 
         * @return builder
         * 
         */
        public Builder accessKeyId(String accessKeyId) {
            return accessKeyId(Output.of(accessKeyId));
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
         * @param externalId A unique identifier that is used by third parties when assuming roles in their customers&#39; accounts, aka cross-account role access.
         * 
         * @return builder
         * 
         */
        public Builder externalId(@Nullable Output<String> externalId) {
            $.externalId = externalId;
            return this;
        }

        /**
         * @param externalId A unique identifier that is used by third parties when assuming roles in their customers&#39; accounts, aka cross-account role access.
         * 
         * @return builder
         * 
         */
        public Builder externalId(String externalId) {
            return externalId(Output.of(externalId));
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
         * @param roleSessionName Optional identifier for the assumed role session.
         * 
         * @return builder
         * 
         */
        public Builder roleSessionName(@Nullable Output<String> roleSessionName) {
            $.roleSessionName = roleSessionName;
            return this;
        }

        /**
         * @param roleSessionName Optional identifier for the assumed role session.
         * 
         * @return builder
         * 
         */
        public Builder roleSessionName(String roleSessionName) {
            return roleSessionName(Output.of(roleSessionName));
        }

        /**
         * @param roleToAssume The Amazon Resource Name (ARN) of the role to assume.
         * 
         * @return builder
         * 
         */
        public Builder roleToAssume(@Nullable Output<String> roleToAssume) {
            $.roleToAssume = roleToAssume;
            return this;
        }

        /**
         * @param roleToAssume The Amazon Resource Name (ARN) of the role to assume.
         * 
         * @return builder
         * 
         */
        public Builder roleToAssume(String roleToAssume) {
            return roleToAssume(Output.of(roleToAssume));
        }

        /**
         * @param secretAccessKey The AWS secret access key for signing programmatic requests.
         * 
         * @return builder
         * 
         */
        public Builder secretAccessKey(@Nullable Output<String> secretAccessKey) {
            $.secretAccessKey = secretAccessKey;
            return this;
        }

        /**
         * @param secretAccessKey The AWS secret access key for signing programmatic requests.
         * 
         * @return builder
         * 
         */
        public Builder secretAccessKey(String secretAccessKey) {
            return secretAccessKey(Output.of(secretAccessKey));
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
         * @param sessionToken The AWS session token for signing programmatic requests.
         * 
         * @return builder
         * 
         */
        public Builder sessionToken(@Nullable Output<String> sessionToken) {
            $.sessionToken = sessionToken;
            return this;
        }

        /**
         * @param sessionToken The AWS session token for signing programmatic requests.
         * 
         * @return builder
         * 
         */
        public Builder sessionToken(String sessionToken) {
            return sessionToken(Output.of(sessionToken));
        }

        /**
         * @param useOidc Enable this to attempt getting credentials with OIDC token from Azure Devops.
         * 
         * @return builder
         * 
         */
        public Builder useOidc(@Nullable Output<Boolean> useOidc) {
            $.useOidc = useOidc;
            return this;
        }

        /**
         * @param useOidc Enable this to attempt getting credentials with OIDC token from Azure Devops.
         * 
         * @return builder
         * 
         */
        public Builder useOidc(Boolean useOidc) {
            return useOidc(Output.of(useOidc));
        }

        public ServiceEndpointAwsArgs build() {
            if ($.projectId == null) {
                throw new MissingRequiredPropertyException("ServiceEndpointAwsArgs", "projectId");
            }
            if ($.serviceEndpointName == null) {
                throw new MissingRequiredPropertyException("ServiceEndpointAwsArgs", "serviceEndpointName");
            }
            return $;
        }
    }

}
