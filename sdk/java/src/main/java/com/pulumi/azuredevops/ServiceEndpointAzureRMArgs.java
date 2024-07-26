// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops;

import com.pulumi.azuredevops.inputs.ServiceEndpointAzureRMCredentialsArgs;
import com.pulumi.azuredevops.inputs.ServiceEndpointAzureRMFeaturesArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ServiceEndpointAzureRMArgs extends com.pulumi.resources.ResourceArgs {

    public static final ServiceEndpointAzureRMArgs Empty = new ServiceEndpointAzureRMArgs();

    @Import(name="authorization")
    private @Nullable Output<Map<String,String>> authorization;

    public Optional<Output<Map<String,String>>> authorization() {
        return Optional.ofNullable(this.authorization);
    }

    /**
     * The Management group ID of the Azure targets.
     * 
     */
    @Import(name="azurermManagementGroupId")
    private @Nullable Output<String> azurermManagementGroupId;

    /**
     * @return The Management group ID of the Azure targets.
     * 
     */
    public Optional<Output<String>> azurermManagementGroupId() {
        return Optional.ofNullable(this.azurermManagementGroupId);
    }

    /**
     * The Management group Name of the targets.
     * 
     */
    @Import(name="azurermManagementGroupName")
    private @Nullable Output<String> azurermManagementGroupName;

    /**
     * @return The Management group Name of the targets.
     * 
     */
    public Optional<Output<String>> azurermManagementGroupName() {
        return Optional.ofNullable(this.azurermManagementGroupName);
    }

    /**
     * The Tenant ID if the service principal.
     * 
     */
    @Import(name="azurermSpnTenantid", required=true)
    private Output<String> azurermSpnTenantid;

    /**
     * @return The Tenant ID if the service principal.
     * 
     */
    public Output<String> azurermSpnTenantid() {
        return this.azurermSpnTenantid;
    }

    /**
     * The Subscription ID of the Azure targets.
     * 
     */
    @Import(name="azurermSubscriptionId")
    private @Nullable Output<String> azurermSubscriptionId;

    /**
     * @return The Subscription ID of the Azure targets.
     * 
     */
    public Optional<Output<String>> azurermSubscriptionId() {
        return Optional.ofNullable(this.azurermSubscriptionId);
    }

    /**
     * The Subscription Name of the targets.
     * 
     */
    @Import(name="azurermSubscriptionName")
    private @Nullable Output<String> azurermSubscriptionName;

    /**
     * @return The Subscription Name of the targets.
     * 
     */
    public Optional<Output<String>> azurermSubscriptionName() {
        return Optional.ofNullable(this.azurermSubscriptionName);
    }

    /**
     * A `credentials` block.
     * 
     */
    @Import(name="credentials")
    private @Nullable Output<ServiceEndpointAzureRMCredentialsArgs> credentials;

    /**
     * @return A `credentials` block.
     * 
     */
    public Optional<Output<ServiceEndpointAzureRMCredentialsArgs>> credentials() {
        return Optional.ofNullable(this.credentials);
    }

    /**
     * Service connection description.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Service connection description.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The Cloud Environment to use. Defaults to `AzureCloud`. Possible values are `AzureCloud`, `AzureChinaCloud`, `AzureUSGovernment`, and `AzureGermanCloud`. Changing this forces a new resource to be created.
     * 
     * &gt; **NOTE:** One of either `Subscription` scoped i.e. `azurerm_subscription_id`, `azurerm_subscription_name` or `ManagementGroup` scoped i.e. `azurerm_management_group_id`, `azurerm_management_group_name` values must be specified.
     * 
     */
    @Import(name="environment")
    private @Nullable Output<String> environment;

    /**
     * @return The Cloud Environment to use. Defaults to `AzureCloud`. Possible values are `AzureCloud`, `AzureChinaCloud`, `AzureUSGovernment`, and `AzureGermanCloud`. Changing this forces a new resource to be created.
     * 
     * &gt; **NOTE:** One of either `Subscription` scoped i.e. `azurerm_subscription_id`, `azurerm_subscription_name` or `ManagementGroup` scoped i.e. `azurerm_management_group_id`, `azurerm_management_group_name` values must be specified.
     * 
     */
    public Optional<Output<String>> environment() {
        return Optional.ofNullable(this.environment);
    }

    /**
     * A `features` block.
     * 
     */
    @Import(name="features")
    private @Nullable Output<ServiceEndpointAzureRMFeaturesArgs> features;

    /**
     * @return A `features` block.
     * 
     */
    public Optional<Output<ServiceEndpointAzureRMFeaturesArgs>> features() {
        return Optional.ofNullable(this.features);
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
     * The resource group used for scope of automatic service endpoint.
     * 
     */
    @Import(name="resourceGroup")
    private @Nullable Output<String> resourceGroup;

    /**
     * @return The resource group used for scope of automatic service endpoint.
     * 
     */
    public Optional<Output<String>> resourceGroup() {
        return Optional.ofNullable(this.resourceGroup);
    }

    /**
     * Specifies the type of azurerm endpoint, either `WorkloadIdentityFederation`, `ManagedServiceIdentity` or `ServicePrincipal`. Defaults to `ServicePrincipal` for backwards compatibility.
     * 
     * &gt; **NOTE:** The `WorkloadIdentityFederation` authentication scheme is currently in private preview. Your organisation must be part of the preview and the feature toggle must be turned on to use it. More details can be found [here](https://aka.ms/azdo-rm-workload-identity).
     * 
     */
    @Import(name="serviceEndpointAuthenticationScheme")
    private @Nullable Output<String> serviceEndpointAuthenticationScheme;

    /**
     * @return Specifies the type of azurerm endpoint, either `WorkloadIdentityFederation`, `ManagedServiceIdentity` or `ServicePrincipal`. Defaults to `ServicePrincipal` for backwards compatibility.
     * 
     * &gt; **NOTE:** The `WorkloadIdentityFederation` authentication scheme is currently in private preview. Your organisation must be part of the preview and the feature toggle must be turned on to use it. More details can be found [here](https://aka.ms/azdo-rm-workload-identity).
     * 
     */
    public Optional<Output<String>> serviceEndpointAuthenticationScheme() {
        return Optional.ofNullable(this.serviceEndpointAuthenticationScheme);
    }

    /**
     * The Service Endpoint Name.
     * 
     */
    @Import(name="serviceEndpointName", required=true)
    private Output<String> serviceEndpointName;

    /**
     * @return The Service Endpoint Name.
     * 
     */
    public Output<String> serviceEndpointName() {
        return this.serviceEndpointName;
    }

    private ServiceEndpointAzureRMArgs() {}

    private ServiceEndpointAzureRMArgs(ServiceEndpointAzureRMArgs $) {
        this.authorization = $.authorization;
        this.azurermManagementGroupId = $.azurermManagementGroupId;
        this.azurermManagementGroupName = $.azurermManagementGroupName;
        this.azurermSpnTenantid = $.azurermSpnTenantid;
        this.azurermSubscriptionId = $.azurermSubscriptionId;
        this.azurermSubscriptionName = $.azurermSubscriptionName;
        this.credentials = $.credentials;
        this.description = $.description;
        this.environment = $.environment;
        this.features = $.features;
        this.projectId = $.projectId;
        this.resourceGroup = $.resourceGroup;
        this.serviceEndpointAuthenticationScheme = $.serviceEndpointAuthenticationScheme;
        this.serviceEndpointName = $.serviceEndpointName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServiceEndpointAzureRMArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServiceEndpointAzureRMArgs $;

        public Builder() {
            $ = new ServiceEndpointAzureRMArgs();
        }

        public Builder(ServiceEndpointAzureRMArgs defaults) {
            $ = new ServiceEndpointAzureRMArgs(Objects.requireNonNull(defaults));
        }

        public Builder authorization(@Nullable Output<Map<String,String>> authorization) {
            $.authorization = authorization;
            return this;
        }

        public Builder authorization(Map<String,String> authorization) {
            return authorization(Output.of(authorization));
        }

        /**
         * @param azurermManagementGroupId The Management group ID of the Azure targets.
         * 
         * @return builder
         * 
         */
        public Builder azurermManagementGroupId(@Nullable Output<String> azurermManagementGroupId) {
            $.azurermManagementGroupId = azurermManagementGroupId;
            return this;
        }

        /**
         * @param azurermManagementGroupId The Management group ID of the Azure targets.
         * 
         * @return builder
         * 
         */
        public Builder azurermManagementGroupId(String azurermManagementGroupId) {
            return azurermManagementGroupId(Output.of(azurermManagementGroupId));
        }

        /**
         * @param azurermManagementGroupName The Management group Name of the targets.
         * 
         * @return builder
         * 
         */
        public Builder azurermManagementGroupName(@Nullable Output<String> azurermManagementGroupName) {
            $.azurermManagementGroupName = azurermManagementGroupName;
            return this;
        }

        /**
         * @param azurermManagementGroupName The Management group Name of the targets.
         * 
         * @return builder
         * 
         */
        public Builder azurermManagementGroupName(String azurermManagementGroupName) {
            return azurermManagementGroupName(Output.of(azurermManagementGroupName));
        }

        /**
         * @param azurermSpnTenantid The Tenant ID if the service principal.
         * 
         * @return builder
         * 
         */
        public Builder azurermSpnTenantid(Output<String> azurermSpnTenantid) {
            $.azurermSpnTenantid = azurermSpnTenantid;
            return this;
        }

        /**
         * @param azurermSpnTenantid The Tenant ID if the service principal.
         * 
         * @return builder
         * 
         */
        public Builder azurermSpnTenantid(String azurermSpnTenantid) {
            return azurermSpnTenantid(Output.of(azurermSpnTenantid));
        }

        /**
         * @param azurermSubscriptionId The Subscription ID of the Azure targets.
         * 
         * @return builder
         * 
         */
        public Builder azurermSubscriptionId(@Nullable Output<String> azurermSubscriptionId) {
            $.azurermSubscriptionId = azurermSubscriptionId;
            return this;
        }

        /**
         * @param azurermSubscriptionId The Subscription ID of the Azure targets.
         * 
         * @return builder
         * 
         */
        public Builder azurermSubscriptionId(String azurermSubscriptionId) {
            return azurermSubscriptionId(Output.of(azurermSubscriptionId));
        }

        /**
         * @param azurermSubscriptionName The Subscription Name of the targets.
         * 
         * @return builder
         * 
         */
        public Builder azurermSubscriptionName(@Nullable Output<String> azurermSubscriptionName) {
            $.azurermSubscriptionName = azurermSubscriptionName;
            return this;
        }

        /**
         * @param azurermSubscriptionName The Subscription Name of the targets.
         * 
         * @return builder
         * 
         */
        public Builder azurermSubscriptionName(String azurermSubscriptionName) {
            return azurermSubscriptionName(Output.of(azurermSubscriptionName));
        }

        /**
         * @param credentials A `credentials` block.
         * 
         * @return builder
         * 
         */
        public Builder credentials(@Nullable Output<ServiceEndpointAzureRMCredentialsArgs> credentials) {
            $.credentials = credentials;
            return this;
        }

        /**
         * @param credentials A `credentials` block.
         * 
         * @return builder
         * 
         */
        public Builder credentials(ServiceEndpointAzureRMCredentialsArgs credentials) {
            return credentials(Output.of(credentials));
        }

        /**
         * @param description Service connection description.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Service connection description.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param environment The Cloud Environment to use. Defaults to `AzureCloud`. Possible values are `AzureCloud`, `AzureChinaCloud`, `AzureUSGovernment`, and `AzureGermanCloud`. Changing this forces a new resource to be created.
         * 
         * &gt; **NOTE:** One of either `Subscription` scoped i.e. `azurerm_subscription_id`, `azurerm_subscription_name` or `ManagementGroup` scoped i.e. `azurerm_management_group_id`, `azurerm_management_group_name` values must be specified.
         * 
         * @return builder
         * 
         */
        public Builder environment(@Nullable Output<String> environment) {
            $.environment = environment;
            return this;
        }

        /**
         * @param environment The Cloud Environment to use. Defaults to `AzureCloud`. Possible values are `AzureCloud`, `AzureChinaCloud`, `AzureUSGovernment`, and `AzureGermanCloud`. Changing this forces a new resource to be created.
         * 
         * &gt; **NOTE:** One of either `Subscription` scoped i.e. `azurerm_subscription_id`, `azurerm_subscription_name` or `ManagementGroup` scoped i.e. `azurerm_management_group_id`, `azurerm_management_group_name` values must be specified.
         * 
         * @return builder
         * 
         */
        public Builder environment(String environment) {
            return environment(Output.of(environment));
        }

        /**
         * @param features A `features` block.
         * 
         * @return builder
         * 
         */
        public Builder features(@Nullable Output<ServiceEndpointAzureRMFeaturesArgs> features) {
            $.features = features;
            return this;
        }

        /**
         * @param features A `features` block.
         * 
         * @return builder
         * 
         */
        public Builder features(ServiceEndpointAzureRMFeaturesArgs features) {
            return features(Output.of(features));
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
         * @param resourceGroup The resource group used for scope of automatic service endpoint.
         * 
         * @return builder
         * 
         */
        public Builder resourceGroup(@Nullable Output<String> resourceGroup) {
            $.resourceGroup = resourceGroup;
            return this;
        }

        /**
         * @param resourceGroup The resource group used for scope of automatic service endpoint.
         * 
         * @return builder
         * 
         */
        public Builder resourceGroup(String resourceGroup) {
            return resourceGroup(Output.of(resourceGroup));
        }

        /**
         * @param serviceEndpointAuthenticationScheme Specifies the type of azurerm endpoint, either `WorkloadIdentityFederation`, `ManagedServiceIdentity` or `ServicePrincipal`. Defaults to `ServicePrincipal` for backwards compatibility.
         * 
         * &gt; **NOTE:** The `WorkloadIdentityFederation` authentication scheme is currently in private preview. Your organisation must be part of the preview and the feature toggle must be turned on to use it. More details can be found [here](https://aka.ms/azdo-rm-workload-identity).
         * 
         * @return builder
         * 
         */
        public Builder serviceEndpointAuthenticationScheme(@Nullable Output<String> serviceEndpointAuthenticationScheme) {
            $.serviceEndpointAuthenticationScheme = serviceEndpointAuthenticationScheme;
            return this;
        }

        /**
         * @param serviceEndpointAuthenticationScheme Specifies the type of azurerm endpoint, either `WorkloadIdentityFederation`, `ManagedServiceIdentity` or `ServicePrincipal`. Defaults to `ServicePrincipal` for backwards compatibility.
         * 
         * &gt; **NOTE:** The `WorkloadIdentityFederation` authentication scheme is currently in private preview. Your organisation must be part of the preview and the feature toggle must be turned on to use it. More details can be found [here](https://aka.ms/azdo-rm-workload-identity).
         * 
         * @return builder
         * 
         */
        public Builder serviceEndpointAuthenticationScheme(String serviceEndpointAuthenticationScheme) {
            return serviceEndpointAuthenticationScheme(Output.of(serviceEndpointAuthenticationScheme));
        }

        /**
         * @param serviceEndpointName The Service Endpoint Name.
         * 
         * @return builder
         * 
         */
        public Builder serviceEndpointName(Output<String> serviceEndpointName) {
            $.serviceEndpointName = serviceEndpointName;
            return this;
        }

        /**
         * @param serviceEndpointName The Service Endpoint Name.
         * 
         * @return builder
         * 
         */
        public Builder serviceEndpointName(String serviceEndpointName) {
            return serviceEndpointName(Output.of(serviceEndpointName));
        }

        public ServiceEndpointAzureRMArgs build() {
            if ($.azurermSpnTenantid == null) {
                throw new MissingRequiredPropertyException("ServiceEndpointAzureRMArgs", "azurermSpnTenantid");
            }
            if ($.projectId == null) {
                throw new MissingRequiredPropertyException("ServiceEndpointAzureRMArgs", "projectId");
            }
            if ($.serviceEndpointName == null) {
                throw new MissingRequiredPropertyException("ServiceEndpointAzureRMArgs", "serviceEndpointName");
            }
            return $;
        }
    }

}
