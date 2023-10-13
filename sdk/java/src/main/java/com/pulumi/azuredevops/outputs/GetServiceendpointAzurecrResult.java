// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Map;
import java.util.Objects;

@CustomType
public final class GetServiceendpointAzurecrResult {
    /**
     * @return The Object ID of the Service Principal.
     * 
     */
    private String appObjectId;
    /**
     * @return Specifies the Authorization Scheme Map.
     * 
     */
    private Map<String,String> authorization;
    /**
     * @return The ID of Service Principal Role Assignment.
     * 
     */
    private String azSpnRoleAssignmentId;
    /**
     * @return The Service Principal Role Permissions.
     * 
     */
    private String azSpnRolePermissions;
    /**
     * @return The Azure Container Registry name.
     * 
     */
    private String azurecrName;
    /**
     * @return The Tenant ID of the service principal.
     * 
     */
    private String azurecrSpnTenantid;
    /**
     * @return The Subscription ID of the Azure targets.
     * 
     */
    private String azurecrSubscriptionId;
    /**
     * @return The Subscription Name of the Azure targets.
     * 
     */
    private String azurecrSubscriptionName;
    /**
     * @return The Service Endpoint description.
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
     * @return The Resource Group to which the Container Registry belongs.
     * 
     */
    private String resourceGroup;
    private String serviceEndpointId;
    private String serviceEndpointName;
    /**
     * @return The Application(Client) ID of the Service Principal.
     * 
     */
    private String servicePrincipalId;
    /**
     * @return The ID of the Service Principal.
     * 
     */
    private String spnObjectId;

    private GetServiceendpointAzurecrResult() {}
    /**
     * @return The Object ID of the Service Principal.
     * 
     */
    public String appObjectId() {
        return this.appObjectId;
    }
    /**
     * @return Specifies the Authorization Scheme Map.
     * 
     */
    public Map<String,String> authorization() {
        return this.authorization;
    }
    /**
     * @return The ID of Service Principal Role Assignment.
     * 
     */
    public String azSpnRoleAssignmentId() {
        return this.azSpnRoleAssignmentId;
    }
    /**
     * @return The Service Principal Role Permissions.
     * 
     */
    public String azSpnRolePermissions() {
        return this.azSpnRolePermissions;
    }
    /**
     * @return The Azure Container Registry name.
     * 
     */
    public String azurecrName() {
        return this.azurecrName;
    }
    /**
     * @return The Tenant ID of the service principal.
     * 
     */
    public String azurecrSpnTenantid() {
        return this.azurecrSpnTenantid;
    }
    /**
     * @return The Subscription ID of the Azure targets.
     * 
     */
    public String azurecrSubscriptionId() {
        return this.azurecrSubscriptionId;
    }
    /**
     * @return The Subscription Name of the Azure targets.
     * 
     */
    public String azurecrSubscriptionName() {
        return this.azurecrSubscriptionName;
    }
    /**
     * @return The Service Endpoint description.
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
     * @return The Resource Group to which the Container Registry belongs.
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
    /**
     * @return The Application(Client) ID of the Service Principal.
     * 
     */
    public String servicePrincipalId() {
        return this.servicePrincipalId;
    }
    /**
     * @return The ID of the Service Principal.
     * 
     */
    public String spnObjectId() {
        return this.spnObjectId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetServiceendpointAzurecrResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String appObjectId;
        private Map<String,String> authorization;
        private String azSpnRoleAssignmentId;
        private String azSpnRolePermissions;
        private String azurecrName;
        private String azurecrSpnTenantid;
        private String azurecrSubscriptionId;
        private String azurecrSubscriptionName;
        private String description;
        private String id;
        private String projectId;
        private String resourceGroup;
        private String serviceEndpointId;
        private String serviceEndpointName;
        private String servicePrincipalId;
        private String spnObjectId;
        public Builder() {}
        public Builder(GetServiceendpointAzurecrResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.appObjectId = defaults.appObjectId;
    	      this.authorization = defaults.authorization;
    	      this.azSpnRoleAssignmentId = defaults.azSpnRoleAssignmentId;
    	      this.azSpnRolePermissions = defaults.azSpnRolePermissions;
    	      this.azurecrName = defaults.azurecrName;
    	      this.azurecrSpnTenantid = defaults.azurecrSpnTenantid;
    	      this.azurecrSubscriptionId = defaults.azurecrSubscriptionId;
    	      this.azurecrSubscriptionName = defaults.azurecrSubscriptionName;
    	      this.description = defaults.description;
    	      this.id = defaults.id;
    	      this.projectId = defaults.projectId;
    	      this.resourceGroup = defaults.resourceGroup;
    	      this.serviceEndpointId = defaults.serviceEndpointId;
    	      this.serviceEndpointName = defaults.serviceEndpointName;
    	      this.servicePrincipalId = defaults.servicePrincipalId;
    	      this.spnObjectId = defaults.spnObjectId;
        }

        @CustomType.Setter
        public Builder appObjectId(String appObjectId) {
            this.appObjectId = Objects.requireNonNull(appObjectId);
            return this;
        }
        @CustomType.Setter
        public Builder authorization(Map<String,String> authorization) {
            this.authorization = Objects.requireNonNull(authorization);
            return this;
        }
        @CustomType.Setter
        public Builder azSpnRoleAssignmentId(String azSpnRoleAssignmentId) {
            this.azSpnRoleAssignmentId = Objects.requireNonNull(azSpnRoleAssignmentId);
            return this;
        }
        @CustomType.Setter
        public Builder azSpnRolePermissions(String azSpnRolePermissions) {
            this.azSpnRolePermissions = Objects.requireNonNull(azSpnRolePermissions);
            return this;
        }
        @CustomType.Setter
        public Builder azurecrName(String azurecrName) {
            this.azurecrName = Objects.requireNonNull(azurecrName);
            return this;
        }
        @CustomType.Setter
        public Builder azurecrSpnTenantid(String azurecrSpnTenantid) {
            this.azurecrSpnTenantid = Objects.requireNonNull(azurecrSpnTenantid);
            return this;
        }
        @CustomType.Setter
        public Builder azurecrSubscriptionId(String azurecrSubscriptionId) {
            this.azurecrSubscriptionId = Objects.requireNonNull(azurecrSubscriptionId);
            return this;
        }
        @CustomType.Setter
        public Builder azurecrSubscriptionName(String azurecrSubscriptionName) {
            this.azurecrSubscriptionName = Objects.requireNonNull(azurecrSubscriptionName);
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
        @CustomType.Setter
        public Builder servicePrincipalId(String servicePrincipalId) {
            this.servicePrincipalId = Objects.requireNonNull(servicePrincipalId);
            return this;
        }
        @CustomType.Setter
        public Builder spnObjectId(String spnObjectId) {
            this.spnObjectId = Objects.requireNonNull(spnObjectId);
            return this;
        }
        public GetServiceendpointAzurecrResult build() {
            final var o = new GetServiceendpointAzurecrResult();
            o.appObjectId = appObjectId;
            o.authorization = authorization;
            o.azSpnRoleAssignmentId = azSpnRoleAssignmentId;
            o.azSpnRolePermissions = azSpnRolePermissions;
            o.azurecrName = azurecrName;
            o.azurecrSpnTenantid = azurecrSpnTenantid;
            o.azurecrSubscriptionId = azurecrSubscriptionId;
            o.azurecrSubscriptionName = azurecrSubscriptionName;
            o.description = description;
            o.id = id;
            o.projectId = projectId;
            o.resourceGroup = resourceGroup;
            o.serviceEndpointId = serviceEndpointId;
            o.serviceEndpointName = serviceEndpointName;
            o.servicePrincipalId = servicePrincipalId;
            o.spnObjectId = spnObjectId;
            return o;
        }
    }
}