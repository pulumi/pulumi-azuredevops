// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops;

import com.pulumi.azuredevops.inputs.ServiceendpointExternaltfsAuthPersonalArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ServiceendpointExternaltfsArgs extends com.pulumi.resources.ResourceArgs {

    public static final ServiceendpointExternaltfsArgs Empty = new ServiceendpointExternaltfsArgs();

    /**
     * An `auth_personal` block as documented below. Allows connecting using a personal access token.
     * 
     */
    @Import(name="authPersonal", required=true)
    private Output<ServiceendpointExternaltfsAuthPersonalArgs> authPersonal;

    /**
     * @return An `auth_personal` block as documented below. Allows connecting using a personal access token.
     * 
     */
    public Output<ServiceendpointExternaltfsAuthPersonalArgs> authPersonal() {
        return this.authPersonal;
    }

    @Import(name="authorization")
    private @Nullable Output<Map<String,String>> authorization;

    public Optional<Output<Map<String,String>>> authorization() {
        return Optional.ofNullable(this.authorization);
    }

    /**
     * Azure DevOps Organization or TFS Project Collection Url.
     * 
     */
    @Import(name="connectionUrl", required=true)
    private Output<String> connectionUrl;

    /**
     * @return Azure DevOps Organization or TFS Project Collection Url.
     * 
     */
    public Output<String> connectionUrl() {
        return this.connectionUrl;
    }

    @Import(name="description")
    private @Nullable Output<String> description;

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

    private ServiceendpointExternaltfsArgs() {}

    private ServiceendpointExternaltfsArgs(ServiceendpointExternaltfsArgs $) {
        this.authPersonal = $.authPersonal;
        this.authorization = $.authorization;
        this.connectionUrl = $.connectionUrl;
        this.description = $.description;
        this.projectId = $.projectId;
        this.serviceEndpointName = $.serviceEndpointName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServiceendpointExternaltfsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServiceendpointExternaltfsArgs $;

        public Builder() {
            $ = new ServiceendpointExternaltfsArgs();
        }

        public Builder(ServiceendpointExternaltfsArgs defaults) {
            $ = new ServiceendpointExternaltfsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param authPersonal An `auth_personal` block as documented below. Allows connecting using a personal access token.
         * 
         * @return builder
         * 
         */
        public Builder authPersonal(Output<ServiceendpointExternaltfsAuthPersonalArgs> authPersonal) {
            $.authPersonal = authPersonal;
            return this;
        }

        /**
         * @param authPersonal An `auth_personal` block as documented below. Allows connecting using a personal access token.
         * 
         * @return builder
         * 
         */
        public Builder authPersonal(ServiceendpointExternaltfsAuthPersonalArgs authPersonal) {
            return authPersonal(Output.of(authPersonal));
        }

        public Builder authorization(@Nullable Output<Map<String,String>> authorization) {
            $.authorization = authorization;
            return this;
        }

        public Builder authorization(Map<String,String> authorization) {
            return authorization(Output.of(authorization));
        }

        /**
         * @param connectionUrl Azure DevOps Organization or TFS Project Collection Url.
         * 
         * @return builder
         * 
         */
        public Builder connectionUrl(Output<String> connectionUrl) {
            $.connectionUrl = connectionUrl;
            return this;
        }

        /**
         * @param connectionUrl Azure DevOps Organization or TFS Project Collection Url.
         * 
         * @return builder
         * 
         */
        public Builder connectionUrl(String connectionUrl) {
            return connectionUrl(Output.of(connectionUrl));
        }

        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

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

        public ServiceendpointExternaltfsArgs build() {
            if ($.authPersonal == null) {
                throw new MissingRequiredPropertyException("ServiceendpointExternaltfsArgs", "authPersonal");
            }
            if ($.connectionUrl == null) {
                throw new MissingRequiredPropertyException("ServiceendpointExternaltfsArgs", "connectionUrl");
            }
            if ($.projectId == null) {
                throw new MissingRequiredPropertyException("ServiceendpointExternaltfsArgs", "projectId");
            }
            if ($.serviceEndpointName == null) {
                throw new MissingRequiredPropertyException("ServiceendpointExternaltfsArgs", "serviceEndpointName");
            }
            return $;
        }
    }

}
