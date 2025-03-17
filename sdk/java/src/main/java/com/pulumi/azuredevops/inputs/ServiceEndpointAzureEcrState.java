// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.inputs;

import com.pulumi.azuredevops.inputs.ServiceEndpointAzureEcrCredentialsArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ServiceEndpointAzureEcrState extends com.pulumi.resources.ResourceArgs {

    public static final ServiceEndpointAzureEcrState Empty = new ServiceEndpointAzureEcrState();

    @Import(name="appObjectId")
    private @Nullable Output<String> appObjectId;

    public Optional<Output<String>> appObjectId() {
        return Optional.ofNullable(this.appObjectId);
    }

    @Import(name="authorization")
    private @Nullable Output<Map<String,String>> authorization;

    public Optional<Output<Map<String,String>>> authorization() {
        return Optional.ofNullable(this.authorization);
    }

    @Import(name="azSpnRoleAssignmentId")
    private @Nullable Output<String> azSpnRoleAssignmentId;

    public Optional<Output<String>> azSpnRoleAssignmentId() {
        return Optional.ofNullable(this.azSpnRoleAssignmentId);
    }

    @Import(name="azSpnRolePermissions")
    private @Nullable Output<String> azSpnRolePermissions;

    public Optional<Output<String>> azSpnRolePermissions() {
        return Optional.ofNullable(this.azSpnRolePermissions);
    }

    /**
     * The Azure container registry name.
     * 
     */
    @Import(name="azurecrName")
    private @Nullable Output<String> azurecrName;

    /**
     * @return The Azure container registry name.
     * 
     */
    public Optional<Output<String>> azurecrName() {
        return Optional.ofNullable(this.azurecrName);
    }

    /**
     * The tenant id of the service principal.
     * 
     */
    @Import(name="azurecrSpnTenantid")
    private @Nullable Output<String> azurecrSpnTenantid;

    /**
     * @return The tenant id of the service principal.
     * 
     */
    public Optional<Output<String>> azurecrSpnTenantid() {
        return Optional.ofNullable(this.azurecrSpnTenantid);
    }

    /**
     * The subscription id of the Azure targets.
     * 
     */
    @Import(name="azurecrSubscriptionId")
    private @Nullable Output<String> azurecrSubscriptionId;

    /**
     * @return The subscription id of the Azure targets.
     * 
     */
    public Optional<Output<String>> azurecrSubscriptionId() {
        return Optional.ofNullable(this.azurecrSubscriptionId);
    }

    /**
     * The subscription name of the Azure targets.
     * 
     */
    @Import(name="azurecrSubscriptionName")
    private @Nullable Output<String> azurecrSubscriptionName;

    /**
     * @return The subscription name of the Azure targets.
     * 
     */
    public Optional<Output<String>> azurecrSubscriptionName() {
        return Optional.ofNullable(this.azurecrSubscriptionName);
    }

    /**
     * A `credentials` block as defined below.
     * 
     */
    @Import(name="credentials")
    private @Nullable Output<ServiceEndpointAzureEcrCredentialsArgs> credentials;

    /**
     * @return A `credentials` block as defined below.
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
    @Import(name="serviceEndpointName")
    private @Nullable Output<String> serviceEndpointName;

    /**
     * @return The name you will use to refer to this service connection in task inputs.
     * 
     */
    public Optional<Output<String>> serviceEndpointName() {
        return Optional.ofNullable(this.serviceEndpointName);
    }

    /**
     * The Application(Client) ID of the Service Principal.
     * 
     */
    @Import(name="servicePrincipalId")
    private @Nullable Output<String> servicePrincipalId;

    /**
     * @return The Application(Client) ID of the Service Principal.
     * 
     */
    public Optional<Output<String>> servicePrincipalId() {
        return Optional.ofNullable(this.servicePrincipalId);
    }

    @Import(name="spnObjectId")
    private @Nullable Output<String> spnObjectId;

    public Optional<Output<String>> spnObjectId() {
        return Optional.ofNullable(this.spnObjectId);
    }

    /**
     * The issuer if `service_endpoint_authentication_scheme` is set to `WorkloadIdentityFederation`. This looks like `https://vstoken.dev.azure.com/00000000-0000-0000-0000-000000000000`, where the GUID is the Organization ID of your Azure DevOps Organisation.
     * 
     */
    @Import(name="workloadIdentityFederationIssuer")
    private @Nullable Output<String> workloadIdentityFederationIssuer;

    /**
     * @return The issuer if `service_endpoint_authentication_scheme` is set to `WorkloadIdentityFederation`. This looks like `https://vstoken.dev.azure.com/00000000-0000-0000-0000-000000000000`, where the GUID is the Organization ID of your Azure DevOps Organisation.
     * 
     */
    public Optional<Output<String>> workloadIdentityFederationIssuer() {
        return Optional.ofNullable(this.workloadIdentityFederationIssuer);
    }

    /**
     * The subject if `service_endpoint_authentication_scheme` is set to `WorkloadIdentityFederation`. This looks like `sc://&lt;organisation&gt;/&lt;project&gt;/&lt;service-connection-name&gt;`.
     * 
     */
    @Import(name="workloadIdentityFederationSubject")
    private @Nullable Output<String> workloadIdentityFederationSubject;

    /**
     * @return The subject if `service_endpoint_authentication_scheme` is set to `WorkloadIdentityFederation`. This looks like `sc://&lt;organisation&gt;/&lt;project&gt;/&lt;service-connection-name&gt;`.
     * 
     */
    public Optional<Output<String>> workloadIdentityFederationSubject() {
        return Optional.ofNullable(this.workloadIdentityFederationSubject);
    }

    private ServiceEndpointAzureEcrState() {}

    private ServiceEndpointAzureEcrState(ServiceEndpointAzureEcrState $) {
        this.appObjectId = $.appObjectId;
        this.authorization = $.authorization;
        this.azSpnRoleAssignmentId = $.azSpnRoleAssignmentId;
        this.azSpnRolePermissions = $.azSpnRolePermissions;
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
        this.servicePrincipalId = $.servicePrincipalId;
        this.spnObjectId = $.spnObjectId;
        this.workloadIdentityFederationIssuer = $.workloadIdentityFederationIssuer;
        this.workloadIdentityFederationSubject = $.workloadIdentityFederationSubject;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServiceEndpointAzureEcrState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServiceEndpointAzureEcrState $;

        public Builder() {
            $ = new ServiceEndpointAzureEcrState();
        }

        public Builder(ServiceEndpointAzureEcrState defaults) {
            $ = new ServiceEndpointAzureEcrState(Objects.requireNonNull(defaults));
        }

        public Builder appObjectId(@Nullable Output<String> appObjectId) {
            $.appObjectId = appObjectId;
            return this;
        }

        public Builder appObjectId(String appObjectId) {
            return appObjectId(Output.of(appObjectId));
        }

        public Builder authorization(@Nullable Output<Map<String,String>> authorization) {
            $.authorization = authorization;
            return this;
        }

        public Builder authorization(Map<String,String> authorization) {
            return authorization(Output.of(authorization));
        }

        public Builder azSpnRoleAssignmentId(@Nullable Output<String> azSpnRoleAssignmentId) {
            $.azSpnRoleAssignmentId = azSpnRoleAssignmentId;
            return this;
        }

        public Builder azSpnRoleAssignmentId(String azSpnRoleAssignmentId) {
            return azSpnRoleAssignmentId(Output.of(azSpnRoleAssignmentId));
        }

        public Builder azSpnRolePermissions(@Nullable Output<String> azSpnRolePermissions) {
            $.azSpnRolePermissions = azSpnRolePermissions;
            return this;
        }

        public Builder azSpnRolePermissions(String azSpnRolePermissions) {
            return azSpnRolePermissions(Output.of(azSpnRolePermissions));
        }

        /**
         * @param azurecrName The Azure container registry name.
         * 
         * @return builder
         * 
         */
        public Builder azurecrName(@Nullable Output<String> azurecrName) {
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
        public Builder azurecrSpnTenantid(@Nullable Output<String> azurecrSpnTenantid) {
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
        public Builder azurecrSubscriptionId(@Nullable Output<String> azurecrSubscriptionId) {
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
        public Builder azurecrSubscriptionName(@Nullable Output<String> azurecrSubscriptionName) {
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
         * @param credentials A `credentials` block as defined below.
         * 
         * @return builder
         * 
         */
        public Builder credentials(@Nullable Output<ServiceEndpointAzureEcrCredentialsArgs> credentials) {
            $.credentials = credentials;
            return this;
        }

        /**
         * @param credentials A `credentials` block as defined below.
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
        public Builder serviceEndpointName(@Nullable Output<String> serviceEndpointName) {
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

        /**
         * @param servicePrincipalId The Application(Client) ID of the Service Principal.
         * 
         * @return builder
         * 
         */
        public Builder servicePrincipalId(@Nullable Output<String> servicePrincipalId) {
            $.servicePrincipalId = servicePrincipalId;
            return this;
        }

        /**
         * @param servicePrincipalId The Application(Client) ID of the Service Principal.
         * 
         * @return builder
         * 
         */
        public Builder servicePrincipalId(String servicePrincipalId) {
            return servicePrincipalId(Output.of(servicePrincipalId));
        }

        public Builder spnObjectId(@Nullable Output<String> spnObjectId) {
            $.spnObjectId = spnObjectId;
            return this;
        }

        public Builder spnObjectId(String spnObjectId) {
            return spnObjectId(Output.of(spnObjectId));
        }

        /**
         * @param workloadIdentityFederationIssuer The issuer if `service_endpoint_authentication_scheme` is set to `WorkloadIdentityFederation`. This looks like `https://vstoken.dev.azure.com/00000000-0000-0000-0000-000000000000`, where the GUID is the Organization ID of your Azure DevOps Organisation.
         * 
         * @return builder
         * 
         */
        public Builder workloadIdentityFederationIssuer(@Nullable Output<String> workloadIdentityFederationIssuer) {
            $.workloadIdentityFederationIssuer = workloadIdentityFederationIssuer;
            return this;
        }

        /**
         * @param workloadIdentityFederationIssuer The issuer if `service_endpoint_authentication_scheme` is set to `WorkloadIdentityFederation`. This looks like `https://vstoken.dev.azure.com/00000000-0000-0000-0000-000000000000`, where the GUID is the Organization ID of your Azure DevOps Organisation.
         * 
         * @return builder
         * 
         */
        public Builder workloadIdentityFederationIssuer(String workloadIdentityFederationIssuer) {
            return workloadIdentityFederationIssuer(Output.of(workloadIdentityFederationIssuer));
        }

        /**
         * @param workloadIdentityFederationSubject The subject if `service_endpoint_authentication_scheme` is set to `WorkloadIdentityFederation`. This looks like `sc://&lt;organisation&gt;/&lt;project&gt;/&lt;service-connection-name&gt;`.
         * 
         * @return builder
         * 
         */
        public Builder workloadIdentityFederationSubject(@Nullable Output<String> workloadIdentityFederationSubject) {
            $.workloadIdentityFederationSubject = workloadIdentityFederationSubject;
            return this;
        }

        /**
         * @param workloadIdentityFederationSubject The subject if `service_endpoint_authentication_scheme` is set to `WorkloadIdentityFederation`. This looks like `sc://&lt;organisation&gt;/&lt;project&gt;/&lt;service-connection-name&gt;`.
         * 
         * @return builder
         * 
         */
        public Builder workloadIdentityFederationSubject(String workloadIdentityFederationSubject) {
            return workloadIdentityFederationSubject(Output.of(workloadIdentityFederationSubject));
        }

        public ServiceEndpointAzureEcrState build() {
            return $;
        }
    }

}
