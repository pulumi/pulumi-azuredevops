// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops;

import com.pulumi.azuredevops.inputs.ServiceEndpointAzureEcrCredentialsArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ServiceEndpointAzureEcrArgs extends com.pulumi.resources.ResourceArgs {

    public static final ServiceEndpointAzureEcrArgs Empty = new ServiceEndpointAzureEcrArgs();

    @Import(name="authorization")
    private @Nullable Output<Map<String,String>> authorization;

    public Optional<Output<Map<String,String>>> authorization() {
        return Optional.ofNullable(this.authorization);
    }

    /**
     * The Azure container registry name.
     * 
     */
    @Import(name="azurecrName", required=true)
    private Output<String> azurecrName;

    /**
     * @return The Azure container registry name.
     * 
     */
    public Output<String> azurecrName() {
        return this.azurecrName;
    }

    /**
     * The tenant id of the service principal.
     * 
     */
    @Import(name="azurecrSpnTenantid", required=true)
    private Output<String> azurecrSpnTenantid;

    /**
     * @return The tenant id of the service principal.
     * 
     */
    public Output<String> azurecrSpnTenantid() {
        return this.azurecrSpnTenantid;
    }

    /**
     * The subscription id of the Azure targets.
     * 
     */
    @Import(name="azurecrSubscriptionId", required=true)
    private Output<String> azurecrSubscriptionId;

    /**
     * @return The subscription id of the Azure targets.
     * 
     */
    public Output<String> azurecrSubscriptionId() {
        return this.azurecrSubscriptionId;
    }

    /**
     * The subscription name of the Azure targets.
     * 
     */
    @Import(name="azurecrSubscriptionName", required=true)
    private Output<String> azurecrSubscriptionName;

    /**
     * @return The subscription name of the Azure targets.
     * 
     */
    public Output<String> azurecrSubscriptionName() {
        return this.azurecrSubscriptionName;
    }

    /**
     * A `credentials` block.
     * 
     */
    @Import(name="credentials")
    private @Nullable Output<ServiceEndpointAzureEcrCredentialsArgs> credentials;

    /**
     * @return A `credentials` block.
     * 
     */
    public Optional<Output<ServiceEndpointAzureEcrCredentialsArgs>> credentials() {
        return Optional.ofNullable(this.credentials);
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
     * The resource group to which the container registry belongs.
     * 
     */
    @Import(name="resourceGroup")
    private @Nullable Output<String> resourceGroup;

    /**
     * @return The resource group to which the container registry belongs.
     * 
     */
    public Optional<Output<String>> resourceGroup() {
        return Optional.ofNullable(this.resourceGroup);
    }

    /**
     * Specifies the type of azurerm endpoint, either `WorkloadIdentityFederation`, `ManagedServiceIdentity` or `ServicePrincipal`. Defaults to `ServicePrincipal` for backwards compatibility. `ManagedServiceIdentity` has not yet been implemented for this resource.
     * 
     */
    @Import(name="serviceEndpointAuthenticationScheme")
    private @Nullable Output<String> serviceEndpointAuthenticationScheme;

    /**
     * @return Specifies the type of azurerm endpoint, either `WorkloadIdentityFederation`, `ManagedServiceIdentity` or `ServicePrincipal`. Defaults to `ServicePrincipal` for backwards compatibility. `ManagedServiceIdentity` has not yet been implemented for this resource.
     * 
     */
    public Optional<Output<String>> serviceEndpointAuthenticationScheme() {
        return Optional.ofNullable(this.serviceEndpointAuthenticationScheme);
    }

    /**
     * The name you will use to refer to this service connection in task inputs.
     * 
     */
    @Import(name="serviceEndpointName", required=true)
    private Output<String> serviceEndpointName;

    /**
     * @return The name you will use to refer to this service connection in task inputs.
     * 
     */
    public Output<String> serviceEndpointName() {
        return this.serviceEndpointName;
    }

    private ServiceEndpointAzureEcrArgs() {}

    private ServiceEndpointAzureEcrArgs(ServiceEndpointAzureEcrArgs $) {
        this.authorization = $.authorization;
        this.azurecrName = $.azurecrName;
        this.azurecrSpnTenantid = $.azurecrSpnTenantid;
        this.azurecrSubscriptionId = $.azurecrSubscriptionId;
        this.azurecrSubscriptionName = $.azurecrSubscriptionName;
        this.credentials = $.credentials;
        this.description = $.description;
        this.projectId = $.projectId;
        this.resourceGroup = $.resourceGroup;
        this.serviceEndpointAuthenticationScheme = $.serviceEndpointAuthenticationScheme;
        this.serviceEndpointName = $.serviceEndpointName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServiceEndpointAzureEcrArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServiceEndpointAzureEcrArgs $;

        public Builder() {
            $ = new ServiceEndpointAzureEcrArgs();
        }

        public Builder(ServiceEndpointAzureEcrArgs defaults) {
            $ = new ServiceEndpointAzureEcrArgs(Objects.requireNonNull(defaults));
        }

        public Builder authorization(@Nullable Output<Map<String,String>> authorization) {
            $.authorization = authorization;
            return this;
        }

        public Builder authorization(Map<String,String> authorization) {
            return authorization(Output.of(authorization));
        }

        /**
         * @param azurecrName The Azure container registry name.
         * 
         * @return builder
         * 
         */
        public Builder azurecrName(Output<String> azurecrName) {
            $.azurecrName = azurecrName;
            return this;
        }

        /**
         * @param azurecrName The Azure container registry name.
         * 
         * @return builder
         * 
         */
        public Builder azurecrName(String azurecrName) {
            return azurecrName(Output.of(azurecrName));
        }

        /**
         * @param azurecrSpnTenantid The tenant id of the service principal.
         * 
         * @return builder
         * 
         */
        public Builder azurecrSpnTenantid(Output<String> azurecrSpnTenantid) {
            $.azurecrSpnTenantid = azurecrSpnTenantid;
            return this;
        }

        /**
         * @param azurecrSpnTenantid The tenant id of the service principal.
         * 
         * @return builder
         * 
         */
        public Builder azurecrSpnTenantid(String azurecrSpnTenantid) {
            return azurecrSpnTenantid(Output.of(azurecrSpnTenantid));
        }

        /**
         * @param azurecrSubscriptionId The subscription id of the Azure targets.
         * 
         * @return builder
         * 
         */
        public Builder azurecrSubscriptionId(Output<String> azurecrSubscriptionId) {
            $.azurecrSubscriptionId = azurecrSubscriptionId;
            return this;
        }

        /**
         * @param azurecrSubscriptionId The subscription id of the Azure targets.
         * 
         * @return builder
         * 
         */
        public Builder azurecrSubscriptionId(String azurecrSubscriptionId) {
            return azurecrSubscriptionId(Output.of(azurecrSubscriptionId));
        }

        /**
         * @param azurecrSubscriptionName The subscription name of the Azure targets.
         * 
         * @return builder
         * 
         */
        public Builder azurecrSubscriptionName(Output<String> azurecrSubscriptionName) {
            $.azurecrSubscriptionName = azurecrSubscriptionName;
            return this;
        }

        /**
         * @param azurecrSubscriptionName The subscription name of the Azure targets.
         * 
         * @return builder
         * 
         */
        public Builder azurecrSubscriptionName(String azurecrSubscriptionName) {
            return azurecrSubscriptionName(Output.of(azurecrSubscriptionName));
        }

        /**
         * @param credentials A `credentials` block.
         * 
         * @return builder
         * 
         */
        public Builder credentials(@Nullable Output<ServiceEndpointAzureEcrCredentialsArgs> credentials) {
            $.credentials = credentials;
            return this;
        }

        /**
         * @param credentials A `credentials` block.
         * 
         * @return builder
         * 
         */
        public Builder credentials(ServiceEndpointAzureEcrCredentialsArgs credentials) {
            return credentials(Output.of(credentials));
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
         * @param resourceGroup The resource group to which the container registry belongs.
         * 
         * @return builder
         * 
         */
        public Builder resourceGroup(@Nullable Output<String> resourceGroup) {
            $.resourceGroup = resourceGroup;
            return this;
        }

        /**
         * @param resourceGroup The resource group to which the container registry belongs.
         * 
         * @return builder
         * 
         */
        public Builder resourceGroup(String resourceGroup) {
            return resourceGroup(Output.of(resourceGroup));
        }

        /**
         * @param serviceEndpointAuthenticationScheme Specifies the type of azurerm endpoint, either `WorkloadIdentityFederation`, `ManagedServiceIdentity` or `ServicePrincipal`. Defaults to `ServicePrincipal` for backwards compatibility. `ManagedServiceIdentity` has not yet been implemented for this resource.
         * 
         * @return builder
         * 
         */
        public Builder serviceEndpointAuthenticationScheme(@Nullable Output<String> serviceEndpointAuthenticationScheme) {
            $.serviceEndpointAuthenticationScheme = serviceEndpointAuthenticationScheme;
            return this;
        }

        /**
         * @param serviceEndpointAuthenticationScheme Specifies the type of azurerm endpoint, either `WorkloadIdentityFederation`, `ManagedServiceIdentity` or `ServicePrincipal`. Defaults to `ServicePrincipal` for backwards compatibility. `ManagedServiceIdentity` has not yet been implemented for this resource.
         * 
         * @return builder
         * 
         */
        public Builder serviceEndpointAuthenticationScheme(String serviceEndpointAuthenticationScheme) {
            return serviceEndpointAuthenticationScheme(Output.of(serviceEndpointAuthenticationScheme));
        }

        /**
         * @param serviceEndpointName The name you will use to refer to this service connection in task inputs.
         * 
         * @return builder
         * 
         */
        public Builder serviceEndpointName(Output<String> serviceEndpointName) {
            $.serviceEndpointName = serviceEndpointName;
            return this;
        }

        /**
         * @param serviceEndpointName The name you will use to refer to this service connection in task inputs.
         * 
         * @return builder
         * 
         */
        public Builder serviceEndpointName(String serviceEndpointName) {
            return serviceEndpointName(Output.of(serviceEndpointName));
        }

        public ServiceEndpointAzureEcrArgs build() {
            if ($.azurecrName == null) {
                throw new MissingRequiredPropertyException("ServiceEndpointAzureEcrArgs", "azurecrName");
            }
            if ($.azurecrSpnTenantid == null) {
                throw new MissingRequiredPropertyException("ServiceEndpointAzureEcrArgs", "azurecrSpnTenantid");
            }
            if ($.azurecrSubscriptionId == null) {
                throw new MissingRequiredPropertyException("ServiceEndpointAzureEcrArgs", "azurecrSubscriptionId");
            }
            if ($.azurecrSubscriptionName == null) {
                throw new MissingRequiredPropertyException("ServiceEndpointAzureEcrArgs", "azurecrSubscriptionName");
            }
            if ($.projectId == null) {
                throw new MissingRequiredPropertyException("ServiceEndpointAzureEcrArgs", "projectId");
            }
            if ($.serviceEndpointName == null) {
                throw new MissingRequiredPropertyException("ServiceEndpointAzureEcrArgs", "serviceEndpointName");
            }
            return $;
        }
    }

}
