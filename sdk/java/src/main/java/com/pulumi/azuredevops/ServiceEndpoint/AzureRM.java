// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.ServiceEndpoint;

import com.pulumi.azuredevops.ServiceEndpoint.AzureRMArgs;
import com.pulumi.azuredevops.ServiceEndpoint.inputs.AzureRMState;
import com.pulumi.azuredevops.ServiceEndpoint.outputs.AzureRMCredentials;
import com.pulumi.azuredevops.ServiceEndpoint.outputs.AzureRMFeatures;
import com.pulumi.azuredevops.Utilities;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages Manual or Automatic AzureRM service endpoint within Azure DevOps.
 * 
 * ## Requirements (Manual AzureRM Service Endpoint)
 * 
 * Before to create a service end point in Azure DevOps, you need to create a Service Principal in your Azure subscription.
 * 
 * For detailed steps to create a service principal with Azure cli see the [documentation](https://docs.microsoft.com/en-us/cli/azure/create-an-azure-service-principal-azure-cli?view=azure-cli-latest)
 * 
 * ## Example Usage
 * ### Service Principal Manual AzureRM Service Endpoint (Subscription Scoped)
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.azuredevops.Project;
 * import com.pulumi.azuredevops.ProjectArgs;
 * import com.pulumi.azuredevops.ServiceEndpointAzureRM;
 * import com.pulumi.azuredevops.ServiceEndpointAzureRMArgs;
 * import com.pulumi.azuredevops.inputs.ServiceEndpointAzureRMCredentialsArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var exampleProject = new Project(&#34;exampleProject&#34;, ProjectArgs.builder()        
 *             .visibility(&#34;private&#34;)
 *             .versionControl(&#34;Git&#34;)
 *             .workItemTemplate(&#34;Agile&#34;)
 *             .description(&#34;Managed by Terraform&#34;)
 *             .build());
 * 
 *         var exampleServiceEndpointAzureRM = new ServiceEndpointAzureRM(&#34;exampleServiceEndpointAzureRM&#34;, ServiceEndpointAzureRMArgs.builder()        
 *             .projectId(exampleProject.id())
 *             .serviceEndpointName(&#34;Example AzureRM&#34;)
 *             .description(&#34;Managed by Terraform&#34;)
 *             .serviceEndpointAuthenticationScheme(&#34;ServicePrincipal&#34;)
 *             .credentials(ServiceEndpointAzureRMCredentialsArgs.builder()
 *                 .serviceprincipalid(&#34;00000000-0000-0000-0000-000000000000&#34;)
 *                 .serviceprincipalkey(&#34;xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx&#34;)
 *                 .build())
 *             .azurermSpnTenantid(&#34;00000000-0000-0000-0000-000000000000&#34;)
 *             .azurermSubscriptionId(&#34;00000000-0000-0000-0000-000000000000&#34;)
 *             .azurermSubscriptionName(&#34;Example Subscription Name&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Service Principal Manual AzureRM Service Endpoint (ManagementGroup Scoped)
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.azuredevops.Project;
 * import com.pulumi.azuredevops.ProjectArgs;
 * import com.pulumi.azuredevops.ServiceEndpointAzureRM;
 * import com.pulumi.azuredevops.ServiceEndpointAzureRMArgs;
 * import com.pulumi.azuredevops.inputs.ServiceEndpointAzureRMCredentialsArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var exampleProject = new Project(&#34;exampleProject&#34;, ProjectArgs.builder()        
 *             .visibility(&#34;private&#34;)
 *             .versionControl(&#34;Git&#34;)
 *             .workItemTemplate(&#34;Agile&#34;)
 *             .description(&#34;Managed by Terraform&#34;)
 *             .build());
 * 
 *         var exampleServiceEndpointAzureRM = new ServiceEndpointAzureRM(&#34;exampleServiceEndpointAzureRM&#34;, ServiceEndpointAzureRMArgs.builder()        
 *             .projectId(exampleProject.id())
 *             .serviceEndpointName(&#34;Example AzureRM&#34;)
 *             .description(&#34;Managed by Terraform&#34;)
 *             .serviceEndpointAuthenticationScheme(&#34;ServicePrincipal&#34;)
 *             .credentials(ServiceEndpointAzureRMCredentialsArgs.builder()
 *                 .serviceprincipalid(&#34;00000000-0000-0000-0000-000000000000&#34;)
 *                 .serviceprincipalkey(&#34;xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx&#34;)
 *                 .build())
 *             .azurermSpnTenantid(&#34;00000000-0000-0000-0000-000000000000&#34;)
 *             .azurermManagementGroupId(&#34;managementGroup&#34;)
 *             .azurermManagementGroupName(&#34;managementGroup&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Service Principal Automatic AzureRM Service Endpoint
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.azuredevops.Project;
 * import com.pulumi.azuredevops.ProjectArgs;
 * import com.pulumi.azuredevops.ServiceEndpointAzureRM;
 * import com.pulumi.azuredevops.ServiceEndpointAzureRMArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var exampleProject = new Project(&#34;exampleProject&#34;, ProjectArgs.builder()        
 *             .visibility(&#34;private&#34;)
 *             .versionControl(&#34;Git&#34;)
 *             .workItemTemplate(&#34;Agile&#34;)
 *             .build());
 * 
 *         var exampleServiceEndpointAzureRM = new ServiceEndpointAzureRM(&#34;exampleServiceEndpointAzureRM&#34;, ServiceEndpointAzureRMArgs.builder()        
 *             .projectId(exampleProject.id())
 *             .serviceEndpointName(&#34;Example AzureRM&#34;)
 *             .serviceEndpointAuthenticationScheme(&#34;ServicePrincipal&#34;)
 *             .azurermSpnTenantid(&#34;00000000-0000-0000-0000-000000000000&#34;)
 *             .azurermSubscriptionId(&#34;00000000-0000-0000-0000-000000000000&#34;)
 *             .azurermSubscriptionName(&#34;Example Subscription Name&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Workload Identity Federation Manual AzureRM Service Endpoint (Subscription Scoped)
 * 
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.azuredevops.Project;
 * import com.pulumi.azuredevops.ProjectArgs;
 * import com.pulumi.azure.core.ResourceGroup;
 * import com.pulumi.azure.core.ResourceGroupArgs;
 * import com.pulumi.azure.authorization.UserAssignedIdentity;
 * import com.pulumi.azure.authorization.UserAssignedIdentityArgs;
 * import com.pulumi.azuredevops.ServiceEndpointAzureRM;
 * import com.pulumi.azuredevops.ServiceEndpointAzureRMArgs;
 * import com.pulumi.azuredevops.inputs.ServiceEndpointAzureRMCredentialsArgs;
 * import com.pulumi.azure.armmsi.FederatedIdentityCredential;
 * import com.pulumi.azure.armmsi.FederatedIdentityCredentialArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         final var serviceConnectionName = &#34;example-federated-sc&#34;;
 * 
 *         var exampleProject = new Project(&#34;exampleProject&#34;, ProjectArgs.builder()        
 *             .visibility(&#34;private&#34;)
 *             .versionControl(&#34;Git&#34;)
 *             .workItemTemplate(&#34;Agile&#34;)
 *             .description(&#34;Managed by Terraform&#34;)
 *             .build());
 * 
 *         var identity = new ResourceGroup(&#34;identity&#34;, ResourceGroupArgs.builder()        
 *             .location(&#34;UK South&#34;)
 *             .build());
 * 
 *         var exampleUserAssignedIdentity = new UserAssignedIdentity(&#34;exampleUserAssignedIdentity&#34;, UserAssignedIdentityArgs.builder()        
 *             .location(identity.location())
 *             .resourceGroupName(&#34;azurerm_resource_group.identity.name&#34;)
 *             .build());
 * 
 *         var exampleServiceEndpointAzureRM = new ServiceEndpointAzureRM(&#34;exampleServiceEndpointAzureRM&#34;, ServiceEndpointAzureRMArgs.builder()        
 *             .projectId(exampleProject.id())
 *             .serviceEndpointName(serviceConnectionName)
 *             .description(&#34;Managed by Terraform&#34;)
 *             .serviceEndpointAuthenticationScheme(&#34;WorkloadIdentityFederation&#34;)
 *             .credentials(ServiceEndpointAzureRMCredentialsArgs.builder()
 *                 .serviceprincipalid(exampleUserAssignedIdentity.clientId())
 *                 .build())
 *             .azurermSpnTenantid(&#34;00000000-0000-0000-0000-000000000000&#34;)
 *             .azurermSubscriptionId(&#34;00000000-0000-0000-0000-000000000000&#34;)
 *             .azurermSubscriptionName(&#34;Example Subscription Name&#34;)
 *             .build());
 * 
 *         var exampleFederatedIdentityCredential = new FederatedIdentityCredential(&#34;exampleFederatedIdentityCredential&#34;, FederatedIdentityCredentialArgs.builder()        
 *             .resourceGroupName(identity.name())
 *             .parentId(exampleUserAssignedIdentity.id())
 *             .audience(&#34;api://AzureADTokenExchange&#34;)
 *             .issuer(exampleServiceEndpointAzureRM.workloadIdentityFederationIssuer())
 *             .subject(exampleServiceEndpointAzureRM.workloadIdentityFederationSubject())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Workload Identity Federation Automatic AzureRM Service Endpoint
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.azuredevops.Project;
 * import com.pulumi.azuredevops.ProjectArgs;
 * import com.pulumi.azuredevops.ServiceEndpointAzureRM;
 * import com.pulumi.azuredevops.ServiceEndpointAzureRMArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var exampleProject = new Project(&#34;exampleProject&#34;, ProjectArgs.builder()        
 *             .visibility(&#34;private&#34;)
 *             .versionControl(&#34;Git&#34;)
 *             .workItemTemplate(&#34;Agile&#34;)
 *             .build());
 * 
 *         var exampleServiceEndpointAzureRM = new ServiceEndpointAzureRM(&#34;exampleServiceEndpointAzureRM&#34;, ServiceEndpointAzureRMArgs.builder()        
 *             .projectId(exampleProject.id())
 *             .serviceEndpointName(&#34;Example AzureRM&#34;)
 *             .serviceEndpointAuthenticationScheme(&#34;WorkloadIdentityFederation&#34;)
 *             .azurermSpnTenantid(&#34;00000000-0000-0000-0000-000000000000&#34;)
 *             .azurermSubscriptionId(&#34;00000000-0000-0000-0000-000000000000&#34;)
 *             .azurermSubscriptionName(&#34;Example Subscription Name&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Managed Identity AzureRM Service Endpoint
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.azuredevops.Project;
 * import com.pulumi.azuredevops.ProjectArgs;
 * import com.pulumi.azuredevops.ServiceEndpointAzureRM;
 * import com.pulumi.azuredevops.ServiceEndpointAzureRMArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var exampleProject = new Project(&#34;exampleProject&#34;, ProjectArgs.builder()        
 *             .visibility(&#34;private&#34;)
 *             .versionControl(&#34;Git&#34;)
 *             .workItemTemplate(&#34;Agile&#34;)
 *             .build());
 * 
 *         var exampleServiceEndpointAzureRM = new ServiceEndpointAzureRM(&#34;exampleServiceEndpointAzureRM&#34;, ServiceEndpointAzureRMArgs.builder()        
 *             .projectId(exampleProject.id())
 *             .serviceEndpointName(&#34;Example AzureRM&#34;)
 *             .serviceEndpointAuthenticationScheme(&#34;ManagedServiceIdentity&#34;)
 *             .azurermSpnTenantid(&#34;00000000-0000-0000-0000-000000000000&#34;)
 *             .azurermSubscriptionId(&#34;00000000-0000-0000-0000-000000000000&#34;)
 *             .azurermSubscriptionName(&#34;Example Subscription Name&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ## Relevant Links
 * 
 * - [Azure DevOps Service REST API 7.0 - Service End points](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-7.0)
 * 
 * ## Import
 * 
 * Azure DevOps Service Endpoint Azure Resource Manage can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**
 * 
 * ```sh
 *  $ pulumi import azuredevops:ServiceEndpoint/azureRM:AzureRM example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
 * ```
 * 
 * @deprecated
 * azuredevops.serviceendpoint.AzureRM has been deprecated in favor of azuredevops.ServiceEndpointAzureRM
 * 
 */
@Deprecated /* azuredevops.serviceendpoint.AzureRM has been deprecated in favor of azuredevops.ServiceEndpointAzureRM */
@ResourceType(type="azuredevops:ServiceEndpoint/azureRM:AzureRM")
public class AzureRM extends com.pulumi.resources.CustomResource {
    @Export(name="authorization", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> authorization;

    public Output<Map<String,String>> authorization() {
        return this.authorization;
    }
    /**
     * The Management group ID of the Azure targets.
     * 
     */
    @Export(name="azurermManagementGroupId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> azurermManagementGroupId;

    /**
     * @return The Management group ID of the Azure targets.
     * 
     */
    public Output<Optional<String>> azurermManagementGroupId() {
        return Codegen.optional(this.azurermManagementGroupId);
    }
    /**
     * The Management group Name of the targets.
     * 
     */
    @Export(name="azurermManagementGroupName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> azurermManagementGroupName;

    /**
     * @return The Management group Name of the targets.
     * 
     */
    public Output<Optional<String>> azurermManagementGroupName() {
        return Codegen.optional(this.azurermManagementGroupName);
    }
    /**
     * The Tenant ID if the service principal.
     * 
     */
    @Export(name="azurermSpnTenantid", refs={String.class}, tree="[0]")
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
    @Export(name="azurermSubscriptionId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> azurermSubscriptionId;

    /**
     * @return The Subscription ID of the Azure targets.
     * 
     */
    public Output<Optional<String>> azurermSubscriptionId() {
        return Codegen.optional(this.azurermSubscriptionId);
    }
    /**
     * The Subscription Name of the targets.
     * 
     */
    @Export(name="azurermSubscriptionName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> azurermSubscriptionName;

    /**
     * @return The Subscription Name of the targets.
     * 
     */
    public Output<Optional<String>> azurermSubscriptionName() {
        return Codegen.optional(this.azurermSubscriptionName);
    }
    /**
     * A `credentials` block.
     * 
     */
    @Export(name="credentials", refs={AzureRMCredentials.class}, tree="[0]")
    private Output</* @Nullable */ AzureRMCredentials> credentials;

    /**
     * @return A `credentials` block.
     * 
     */
    public Output<Optional<AzureRMCredentials>> credentials() {
        return Codegen.optional(this.credentials);
    }
    /**
     * Service connection description.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return Service connection description.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The Cloud Environment to use. Defaults to `AzureCloud`. Possible values are `AzureCloud`, `AzureChinaCloud`. Changing this forces a new resource to be created.
     * 
     * &gt; **NOTE:** One of either `Subscription` scoped i.e. `azurerm_subscription_id`, `azurerm_subscription_name` or `ManagementGroup` scoped i.e. `azurerm_management_group_id`, `azurerm_management_group_name` values must be specified.
     * 
     */
    @Export(name="environment", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> environment;

    /**
     * @return The Cloud Environment to use. Defaults to `AzureCloud`. Possible values are `AzureCloud`, `AzureChinaCloud`. Changing this forces a new resource to be created.
     * 
     * &gt; **NOTE:** One of either `Subscription` scoped i.e. `azurerm_subscription_id`, `azurerm_subscription_name` or `ManagementGroup` scoped i.e. `azurerm_management_group_id`, `azurerm_management_group_name` values must be specified.
     * 
     */
    public Output<Optional<String>> environment() {
        return Codegen.optional(this.environment);
    }
    /**
     * A `features` block.
     * 
     */
    @Export(name="features", refs={AzureRMFeatures.class}, tree="[0]")
    private Output</* @Nullable */ AzureRMFeatures> features;

    /**
     * @return A `features` block.
     * 
     */
    public Output<Optional<AzureRMFeatures>> features() {
        return Codegen.optional(this.features);
    }
    /**
     * The ID of the project.
     * 
     */
    @Export(name="projectId", refs={String.class}, tree="[0]")
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
    @Export(name="resourceGroup", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> resourceGroup;

    /**
     * @return The resource group used for scope of automatic service endpoint.
     * 
     */
    public Output<Optional<String>> resourceGroup() {
        return Codegen.optional(this.resourceGroup);
    }
    /**
     * Specifies the type of azurerm endpoint, either `WorkloadIdentityFederation`, `ManagedServiceIdentity` or `ServicePrincipal`. Defaults to `ServicePrincipal` for backwards compatibility.
     * 
     * &gt; **NOTE:** The `WorkloadIdentityFederation` authentication scheme is currently in private preview. Your organisation must be part of the preview and the feature toggle must be turned on to use it. More details can be found [here](https://aka.ms/azdo-rm-workload-identity).
     * 
     */
    @Export(name="serviceEndpointAuthenticationScheme", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> serviceEndpointAuthenticationScheme;

    /**
     * @return Specifies the type of azurerm endpoint, either `WorkloadIdentityFederation`, `ManagedServiceIdentity` or `ServicePrincipal`. Defaults to `ServicePrincipal` for backwards compatibility.
     * 
     * &gt; **NOTE:** The `WorkloadIdentityFederation` authentication scheme is currently in private preview. Your organisation must be part of the preview and the feature toggle must be turned on to use it. More details can be found [here](https://aka.ms/azdo-rm-workload-identity).
     * 
     */
    public Output<Optional<String>> serviceEndpointAuthenticationScheme() {
        return Codegen.optional(this.serviceEndpointAuthenticationScheme);
    }
    /**
     * The Service Endpoint Name.
     * 
     */
    @Export(name="serviceEndpointName", refs={String.class}, tree="[0]")
    private Output<String> serviceEndpointName;

    /**
     * @return The Service Endpoint Name.
     * 
     */
    public Output<String> serviceEndpointName() {
        return this.serviceEndpointName;
    }
    /**
     * The Application(Client) ID of the Service Principal.
     * 
     */
    @Export(name="servicePrincipalId", refs={String.class}, tree="[0]")
    private Output<String> servicePrincipalId;

    /**
     * @return The Application(Client) ID of the Service Principal.
     * 
     */
    public Output<String> servicePrincipalId() {
        return this.servicePrincipalId;
    }
    /**
     * The issuer if `service_endpoint_authentication_scheme` is set to `WorkloadIdentityFederation`. This looks like `https://vstoken.dev.azure.com/00000000-0000-0000-0000-000000000000`, where the GUID is the Organization ID of your Azure DevOps Organisation.
     * 
     */
    @Export(name="workloadIdentityFederationIssuer", refs={String.class}, tree="[0]")
    private Output<String> workloadIdentityFederationIssuer;

    /**
     * @return The issuer if `service_endpoint_authentication_scheme` is set to `WorkloadIdentityFederation`. This looks like `https://vstoken.dev.azure.com/00000000-0000-0000-0000-000000000000`, where the GUID is the Organization ID of your Azure DevOps Organisation.
     * 
     */
    public Output<String> workloadIdentityFederationIssuer() {
        return this.workloadIdentityFederationIssuer;
    }
    /**
     * The subject if `service_endpoint_authentication_scheme` is set to `WorkloadIdentityFederation`. This looks like `sc://&lt;organisation&gt;/&lt;project&gt;/&lt;service-connection-name&gt;`.
     * 
     */
    @Export(name="workloadIdentityFederationSubject", refs={String.class}, tree="[0]")
    private Output<String> workloadIdentityFederationSubject;

    /**
     * @return The subject if `service_endpoint_authentication_scheme` is set to `WorkloadIdentityFederation`. This looks like `sc://&lt;organisation&gt;/&lt;project&gt;/&lt;service-connection-name&gt;`.
     * 
     */
    public Output<String> workloadIdentityFederationSubject() {
        return this.workloadIdentityFederationSubject;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AzureRM(String name) {
        this(name, AzureRMArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AzureRM(String name, AzureRMArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AzureRM(String name, AzureRMArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:ServiceEndpoint/azureRM:AzureRM", name, args == null ? AzureRMArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private AzureRM(String name, Output<String> id, @Nullable AzureRMState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:ServiceEndpoint/azureRM:AzureRM", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static AzureRM get(String name, Output<String> id, @Nullable AzureRMState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AzureRM(name, id, state, options);
    }
}
