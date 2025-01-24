// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops;

import com.pulumi.azuredevops.inputs.ServiceEndpointServiceFabricAzureActiveDirectoryArgs;
import com.pulumi.azuredevops.inputs.ServiceEndpointServiceFabricCertificateArgs;
import com.pulumi.azuredevops.inputs.ServiceEndpointServiceFabricNoneArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ServiceEndpointServiceFabricArgs extends com.pulumi.resources.ResourceArgs {

    public static final ServiceEndpointServiceFabricArgs Empty = new ServiceEndpointServiceFabricArgs();

    @Import(name="authorization")
    private @Nullable Output<Map<String,String>> authorization;

    public Optional<Output<Map<String,String>>> authorization() {
        return Optional.ofNullable(this.authorization);
    }

    /**
     * An `azure_active_directory` block as documented below.
     * 
     */
    @Import(name="azureActiveDirectory")
    private @Nullable Output<ServiceEndpointServiceFabricAzureActiveDirectoryArgs> azureActiveDirectory;

    /**
     * @return An `azure_active_directory` block as documented below.
     * 
     */
    public Optional<Output<ServiceEndpointServiceFabricAzureActiveDirectoryArgs>> azureActiveDirectory() {
        return Optional.ofNullable(this.azureActiveDirectory);
    }

    /**
     * A `certificate` block as documented below.
     * 
     */
    @Import(name="certificate")
    private @Nullable Output<ServiceEndpointServiceFabricCertificateArgs> certificate;

    /**
     * @return A `certificate` block as documented below.
     * 
     */
    public Optional<Output<ServiceEndpointServiceFabricCertificateArgs>> certificate() {
        return Optional.ofNullable(this.certificate);
    }

    /**
     * Client connection endpoint for the cluster. Prefix the value with `tcp://`;. This value overrides the publish profile.
     * 
     */
    @Import(name="clusterEndpoint", required=true)
    private Output<String> clusterEndpoint;

    /**
     * @return Client connection endpoint for the cluster. Prefix the value with `tcp://`;. This value overrides the publish profile.
     * 
     */
    public Output<String> clusterEndpoint() {
        return this.clusterEndpoint;
    }

    @Import(name="description")
    private @Nullable Output<String> description;

    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * A `none` block as documented below.
     * 
     */
    @Import(name="none")
    private @Nullable Output<ServiceEndpointServiceFabricNoneArgs> none;

    /**
     * @return A `none` block as documented below.
     * 
     */
    public Optional<Output<ServiceEndpointServiceFabricNoneArgs>> none() {
        return Optional.ofNullable(this.none);
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

    private ServiceEndpointServiceFabricArgs() {}

    private ServiceEndpointServiceFabricArgs(ServiceEndpointServiceFabricArgs $) {
        this.authorization = $.authorization;
        this.azureActiveDirectory = $.azureActiveDirectory;
        this.certificate = $.certificate;
        this.clusterEndpoint = $.clusterEndpoint;
        this.description = $.description;
        this.none = $.none;
        this.projectId = $.projectId;
        this.serviceEndpointName = $.serviceEndpointName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServiceEndpointServiceFabricArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServiceEndpointServiceFabricArgs $;

        public Builder() {
            $ = new ServiceEndpointServiceFabricArgs();
        }

        public Builder(ServiceEndpointServiceFabricArgs defaults) {
            $ = new ServiceEndpointServiceFabricArgs(Objects.requireNonNull(defaults));
        }

        public Builder authorization(@Nullable Output<Map<String,String>> authorization) {
            $.authorization = authorization;
            return this;
        }

        public Builder authorization(Map<String,String> authorization) {
            return authorization(Output.of(authorization));
        }

        /**
         * @param azureActiveDirectory An `azure_active_directory` block as documented below.
         * 
         * @return builder
         * 
         */
        public Builder azureActiveDirectory(@Nullable Output<ServiceEndpointServiceFabricAzureActiveDirectoryArgs> azureActiveDirectory) {
            $.azureActiveDirectory = azureActiveDirectory;
            return this;
        }

        /**
         * @param azureActiveDirectory An `azure_active_directory` block as documented below.
         * 
         * @return builder
         * 
         */
        public Builder azureActiveDirectory(ServiceEndpointServiceFabricAzureActiveDirectoryArgs azureActiveDirectory) {
            return azureActiveDirectory(Output.of(azureActiveDirectory));
        }

        /**
         * @param certificate A `certificate` block as documented below.
         * 
         * @return builder
         * 
         */
        public Builder certificate(@Nullable Output<ServiceEndpointServiceFabricCertificateArgs> certificate) {
            $.certificate = certificate;
            return this;
        }

        /**
         * @param certificate A `certificate` block as documented below.
         * 
         * @return builder
         * 
         */
        public Builder certificate(ServiceEndpointServiceFabricCertificateArgs certificate) {
            return certificate(Output.of(certificate));
        }

        /**
         * @param clusterEndpoint Client connection endpoint for the cluster. Prefix the value with `tcp://`;. This value overrides the publish profile.
         * 
         * @return builder
         * 
         */
        public Builder clusterEndpoint(Output<String> clusterEndpoint) {
            $.clusterEndpoint = clusterEndpoint;
            return this;
        }

        /**
         * @param clusterEndpoint Client connection endpoint for the cluster. Prefix the value with `tcp://`;. This value overrides the publish profile.
         * 
         * @return builder
         * 
         */
        public Builder clusterEndpoint(String clusterEndpoint) {
            return clusterEndpoint(Output.of(clusterEndpoint));
        }

        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param none A `none` block as documented below.
         * 
         * @return builder
         * 
         */
        public Builder none(@Nullable Output<ServiceEndpointServiceFabricNoneArgs> none) {
            $.none = none;
            return this;
        }

        /**
         * @param none A `none` block as documented below.
         * 
         * @return builder
         * 
         */
        public Builder none(ServiceEndpointServiceFabricNoneArgs none) {
            return none(Output.of(none));
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

        public ServiceEndpointServiceFabricArgs build() {
            if ($.clusterEndpoint == null) {
                throw new MissingRequiredPropertyException("ServiceEndpointServiceFabricArgs", "clusterEndpoint");
            }
            if ($.projectId == null) {
                throw new MissingRequiredPropertyException("ServiceEndpointServiceFabricArgs", "projectId");
            }
            if ($.serviceEndpointName == null) {
                throw new MissingRequiredPropertyException("ServiceEndpointServiceFabricArgs", "serviceEndpointName");
            }
            return $;
        }
    }

}
