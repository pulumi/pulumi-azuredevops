// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Map;
import java.util.Objects;

@CustomType
public final class GetServiceEndpointAzureRMResult {
    /**
     * @return Specifies the Authorization Scheme Map.
     * 
     */
    private Map<String,String> authorization;
    /**
     * @return Specified the Management Group ID of the Service Endpoint is target, if available.
     * 
     */
    private String azurermManagementGroupId;
    /**
     * @return Specified the Management Group Name of the Service Endpoint target, if available.
     * 
     */
    private String azurermManagementGroupName;
    /**
     * @return Specifies the Tenant ID of the Azure targets.
     * 
     */
    private String azurermSpnTenantid;
    /**
     * @return Specifies the Subscription ID of the Service Endpoint target, if available.
     * 
     */
    private String azurermSubscriptionId;
    /**
     * @return Specifies the Subscription Name of the Service Endpoint target, if available.
     * 
     */
    private String azurermSubscriptionName;
    /**
     * @return Specifies the description of the Service Endpoint.
     * 
     */
    private String description;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private String projectId;
    /**
     * @return Specifies the Resource Group of the Service Endpoint target, if available.
     * 
     */
    private String resourceGroup;
    private String serviceEndpointId;
    private String serviceEndpointName;

    private GetServiceEndpointAzureRMResult() {}
    /**
     * @return Specifies the Authorization Scheme Map.
     * 
     */
    public Map<String,String> authorization() {
        return this.authorization;
    }
    /**
     * @return Specified the Management Group ID of the Service Endpoint is target, if available.
     * 
     */
    public String azurermManagementGroupId() {
        return this.azurermManagementGroupId;
    }
    /**
     * @return Specified the Management Group Name of the Service Endpoint target, if available.
     * 
     */
    public String azurermManagementGroupName() {
        return this.azurermManagementGroupName;
    }
    /**
     * @return Specifies the Tenant ID of the Azure targets.
     * 
     */
    public String azurermSpnTenantid() {
        return this.azurermSpnTenantid;
    }
    /**
     * @return Specifies the Subscription ID of the Service Endpoint target, if available.
     * 
     */
    public String azurermSubscriptionId() {
        return this.azurermSubscriptionId;
    }
    /**
     * @return Specifies the Subscription Name of the Service Endpoint target, if available.
     * 
     */
    public String azurermSubscriptionName() {
        return this.azurermSubscriptionName;
    }
    /**
     * @return Specifies the description of the Service Endpoint.
     * 
     */
    public String description() {
        return this.description;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public String projectId() {
        return this.projectId;
    }
    /**
     * @return Specifies the Resource Group of the Service Endpoint target, if available.
     * 
     */
    public String resourceGroup() {
        return this.resourceGroup;
    }
    public String serviceEndpointId() {
        return this.serviceEndpointId;
    }
    public String serviceEndpointName() {
        return this.serviceEndpointName;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetServiceEndpointAzureRMResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Map<String,String> authorization;
        private String azurermManagementGroupId;
        private String azurermManagementGroupName;
        private String azurermSpnTenantid;
        private String azurermSubscriptionId;
        private String azurermSubscriptionName;
        private String description;
        private String id;
        private String projectId;
        private String resourceGroup;
        private String serviceEndpointId;
        private String serviceEndpointName;
        public Builder() {}
        public Builder(GetServiceEndpointAzureRMResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.authorization = defaults.authorization;
    	      this.azurermManagementGroupId = defaults.azurermManagementGroupId;
    	      this.azurermManagementGroupName = defaults.azurermManagementGroupName;
    	      this.azurermSpnTenantid = defaults.azurermSpnTenantid;
    	      this.azurermSubscriptionId = defaults.azurermSubscriptionId;
    	      this.azurermSubscriptionName = defaults.azurermSubscriptionName;
    	      this.description = defaults.description;
    	      this.id = defaults.id;
    	      this.projectId = defaults.projectId;
    	      this.resourceGroup = defaults.resourceGroup;
    	      this.serviceEndpointId = defaults.serviceEndpointId;
    	      this.serviceEndpointName = defaults.serviceEndpointName;
        }

        @CustomType.Setter
        public Builder authorization(Map<String,String> authorization) {
            this.authorization = Objects.requireNonNull(authorization);
            return this;
        }
        @CustomType.Setter
        public Builder azurermManagementGroupId(String azurermManagementGroupId) {
            this.azurermManagementGroupId = Objects.requireNonNull(azurermManagementGroupId);
            return this;
        }
        @CustomType.Setter
        public Builder azurermManagementGroupName(String azurermManagementGroupName) {
            this.azurermManagementGroupName = Objects.requireNonNull(azurermManagementGroupName);
            return this;
        }
        @CustomType.Setter
        public Builder azurermSpnTenantid(String azurermSpnTenantid) {
            this.azurermSpnTenantid = Objects.requireNonNull(azurermSpnTenantid);
            return this;
        }
        @CustomType.Setter
        public Builder azurermSubscriptionId(String azurermSubscriptionId) {
            this.azurermSubscriptionId = Objects.requireNonNull(azurermSubscriptionId);
            return this;
        }
        @CustomType.Setter
        public Builder azurermSubscriptionName(String azurermSubscriptionName) {
            this.azurermSubscriptionName = Objects.requireNonNull(azurermSubscriptionName);
            return this;
        }
        @CustomType.Setter
        public Builder description(String description) {
            this.description = Objects.requireNonNull(description);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder projectId(String projectId) {
            this.projectId = Objects.requireNonNull(projectId);
            return this;
        }
        @CustomType.Setter
        public Builder resourceGroup(String resourceGroup) {
            this.resourceGroup = Objects.requireNonNull(resourceGroup);
            return this;
        }
        @CustomType.Setter
        public Builder serviceEndpointId(String serviceEndpointId) {
            this.serviceEndpointId = Objects.requireNonNull(serviceEndpointId);
            return this;
        }
        @CustomType.Setter
        public Builder serviceEndpointName(String serviceEndpointName) {
            this.serviceEndpointName = Objects.requireNonNull(serviceEndpointName);
            return this;
        }
        public GetServiceEndpointAzureRMResult build() {
            final var o = new GetServiceEndpointAzureRMResult();
            o.authorization = authorization;
            o.azurermManagementGroupId = azurermManagementGroupId;
            o.azurermManagementGroupName = azurermManagementGroupName;
            o.azurermSpnTenantid = azurermSpnTenantid;
            o.azurermSubscriptionId = azurermSubscriptionId;
            o.azurermSubscriptionName = azurermSubscriptionName;
            o.description = description;
            o.id = id;
            o.projectId = projectId;
            o.resourceGroup = resourceGroup;
            o.serviceEndpointId = serviceEndpointId;
            o.serviceEndpointName = serviceEndpointName;
            return o;
        }
    }
}
