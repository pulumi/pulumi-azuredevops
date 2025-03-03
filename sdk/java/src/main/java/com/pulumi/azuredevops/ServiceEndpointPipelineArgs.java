// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops;

import com.pulumi.azuredevops.inputs.ServiceEndpointPipelineAuthPersonalArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ServiceEndpointPipelineArgs extends com.pulumi.resources.ResourceArgs {

    public static final ServiceEndpointPipelineArgs Empty = new ServiceEndpointPipelineArgs();

    /**
     * An `auth_personal` block as documented below. Allows connecting using a personal access token.
     * 
     */
    @Import(name="authPersonal", required=true)
    private Output<ServiceEndpointPipelineAuthPersonalArgs> authPersonal;

    /**
     * @return An `auth_personal` block as documented below. Allows connecting using a personal access token.
     * 
     */
    public Output<ServiceEndpointPipelineAuthPersonalArgs> authPersonal() {
        return this.authPersonal;
    }

    @Import(name="description")
    private @Nullable Output<String> description;

    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The organization name used for `Organization Url` and `Release API Url` fields.
     * 
     */
    @Import(name="organizationName", required=true)
    private Output<String> organizationName;

    /**
     * @return The organization name used for `Organization Url` and `Release API Url` fields.
     * 
     */
    public Output<String> organizationName() {
        return this.organizationName;
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

    private ServiceEndpointPipelineArgs() {}

    private ServiceEndpointPipelineArgs(ServiceEndpointPipelineArgs $) {
        this.authPersonal = $.authPersonal;
        this.description = $.description;
        this.organizationName = $.organizationName;
        this.projectId = $.projectId;
        this.serviceEndpointName = $.serviceEndpointName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServiceEndpointPipelineArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServiceEndpointPipelineArgs $;

        public Builder() {
            $ = new ServiceEndpointPipelineArgs();
        }

        public Builder(ServiceEndpointPipelineArgs defaults) {
            $ = new ServiceEndpointPipelineArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param authPersonal An `auth_personal` block as documented below. Allows connecting using a personal access token.
         * 
         * @return builder
         * 
         */
        public Builder authPersonal(Output<ServiceEndpointPipelineAuthPersonalArgs> authPersonal) {
            $.authPersonal = authPersonal;
            return this;
        }

        /**
         * @param authPersonal An `auth_personal` block as documented below. Allows connecting using a personal access token.
         * 
         * @return builder
         * 
         */
        public Builder authPersonal(ServiceEndpointPipelineAuthPersonalArgs authPersonal) {
            return authPersonal(Output.of(authPersonal));
        }

        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param organizationName The organization name used for `Organization Url` and `Release API Url` fields.
         * 
         * @return builder
         * 
         */
        public Builder organizationName(Output<String> organizationName) {
            $.organizationName = organizationName;
            return this;
        }

        /**
         * @param organizationName The organization name used for `Organization Url` and `Release API Url` fields.
         * 
         * @return builder
         * 
         */
        public Builder organizationName(String organizationName) {
            return organizationName(Output.of(organizationName));
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

        public ServiceEndpointPipelineArgs build() {
            if ($.authPersonal == null) {
                throw new MissingRequiredPropertyException("ServiceEndpointPipelineArgs", "authPersonal");
            }
            if ($.organizationName == null) {
                throw new MissingRequiredPropertyException("ServiceEndpointPipelineArgs", "organizationName");
            }
            if ($.projectId == null) {
                throw new MissingRequiredPropertyException("ServiceEndpointPipelineArgs", "projectId");
            }
            if ($.serviceEndpointName == null) {
                throw new MissingRequiredPropertyException("ServiceEndpointPipelineArgs", "serviceEndpointName");
            }
            return $;
        }
    }

}
