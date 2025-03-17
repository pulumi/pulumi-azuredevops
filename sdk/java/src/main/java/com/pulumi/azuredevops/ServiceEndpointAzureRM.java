// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops;

import com.pulumi.azuredevops.ServiceEndpointAzureRMArgs;
import com.pulumi.azuredevops.Utilities;
import com.pulumi.azuredevops.inputs.ServiceEndpointAzureRMState;
import com.pulumi.azuredevops.outputs.ServiceEndpointAzureRMCredentials;
import com.pulumi.azuredevops.outputs.ServiceEndpointAzureRMFeatures;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages Manual or Automatic Azure Resource Manager service endpoint within Azure DevOps.
 * 
 * ~&gt;**NOTE:**
 * If you receive an error message like:```Failed to obtain the Json Web Token(JWT) using service principal client ID. Exception message: A configuration issue is preventing authentication - check the error message from the server for details.```
 * You should check the secret of this Application or if you recently rotate the secret, wait a few minutes for Azure to propagate the secret.
 * 
 * ## Requirements (Manual AzureRM Service Endpoint)
 * 
 * Before to create a service end point in Azure DevOps, you need to create a Service Principal in your Azure subscription.
 * 
 * For detailed steps to create a service principal with Azure cli see the [documentation](https://docs.microsoft.com/en-us/cli/azure/create-an-azure-service-principal-azure-cli?view=azure-cli-latest)
 * 
 * ## Example Usage
 * 
 * ### Service Principal Manual AzureRM Service Endpoint (Subscription Scoped)
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
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
 *         var example = new Project("example", ProjectArgs.builder()
 *             .name("Example Project")
 *             .visibility("private")
 *             .versionControl("Git")
 *             .workItemTemplate("Agile")
 *             .description("Managed by Pulumi")
 *             .build());
 * 
 *         var exampleServiceEndpointAzureRM = new ServiceEndpointAzureRM("exampleServiceEndpointAzureRM", ServiceEndpointAzureRMArgs.builder()
 *             .projectId(example.id())
 *             .serviceEndpointName("Example AzureRM")
 *             .description("Managed by Pulumi")
 *             .serviceEndpointAuthenticationScheme("ServicePrincipal")
 *             .credentials(ServiceEndpointAzureRMCredentialsArgs.builder()
 *                 .serviceprincipalid("00000000-0000-0000-0000-000000000000")
 *                 .serviceprincipalkey("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
 *                 .build())
 *             .azurermSpnTenantid("00000000-0000-0000-0000-000000000000")
 *             .azurermSubscriptionId("00000000-0000-0000-0000-000000000000")
 *             .azurermSubscriptionName("Example Subscription Name")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ### Service Principal Manual AzureRM Service Endpoint (ManagementGroup Scoped)
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
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
 *         var example = new Project("example", ProjectArgs.builder()
 *             .name("Example Project")
 *             .visibility("private")
 *             .versionControl("Git")
 *             .workItemTemplate("Agile")
 *             .description("Managed by Pulumi")
 *             .build());
 * 
 *         var exampleServiceEndpointAzureRM = new ServiceEndpointAzureRM("exampleServiceEndpointAzureRM", ServiceEndpointAzureRMArgs.builder()
 *             .projectId(example.id())
 *             .serviceEndpointName("Example AzureRM")
 *             .description("Managed by Pulumi")
 *             .serviceEndpointAuthenticationScheme("ServicePrincipal")
 *             .credentials(ServiceEndpointAzureRMCredentialsArgs.builder()
 *                 .serviceprincipalid("00000000-0000-0000-0000-000000000000")
 *                 .serviceprincipalkey("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
 *                 .build())
 *             .azurermSpnTenantid("00000000-0000-0000-0000-000000000000")
 *             .azurermManagementGroupId("managementGroup")
 *             .azurermManagementGroupName("managementGroup")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ### Service Principal Automatic AzureRM Service Endpoint
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
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
 *         var example = new Project("example", ProjectArgs.builder()
 *             .name("Example Project")
 *             .visibility("private")
 *             .versionControl("Git")
 *             .workItemTemplate("Agile")
 *             .build());
 * 
 *         var exampleServiceEndpointAzureRM = new ServiceEndpointAzureRM("exampleServiceEndpointAzureRM", ServiceEndpointAzureRMArgs.builder()
 *             .projectId(example.id())
 *             .serviceEndpointName("Example AzureRM")
 *             .serviceEndpointAuthenticationScheme("ServicePrincipal")
 *             .azurermSpnTenantid("00000000-0000-0000-0000-000000000000")
 *             .azurermSubscriptionId("00000000-0000-0000-0000-000000000000")
 *             .azurermSubscriptionName("Example Subscription Name")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ### Workload Identity Federation Manual AzureRM Service Endpoint (Subscription Scoped)
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
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
 *         final var serviceConnectionName = "example-federated-sc";
 * 
 *         var example = new Project("example", ProjectArgs.builder()
 *             .name("Example Project")
 *             .visibility("private")
 *             .versionControl("Git")
 *             .workItemTemplate("Agile")
 *             .description("Managed by Pulumi")
 *             .build());
 * 
 *         var identity = new ResourceGroup("identity", ResourceGroupArgs.builder()
 *             .name("identity")
 *             .location("UK South")
 *             .build());
 * 
 *         var exampleUserAssignedIdentity = new UserAssignedIdentity("exampleUserAssignedIdentity", UserAssignedIdentityArgs.builder()
 *             .location(identity.location())
 *             .name("example-identity")
 *             .resourceGroupName("azurerm_resource_group.identity.name")
 *             .build());
 * 
 *         var exampleServiceEndpointAzureRM = new ServiceEndpointAzureRM("exampleServiceEndpointAzureRM", ServiceEndpointAzureRMArgs.builder()
 *             .projectId(example.id())
 *             .serviceEndpointName(serviceConnectionName)
 *             .description("Managed by Pulumi")
 *             .serviceEndpointAuthenticationScheme("WorkloadIdentityFederation")
 *             .credentials(ServiceEndpointAzureRMCredentialsArgs.builder()
 *                 .serviceprincipalid(exampleUserAssignedIdentity.clientId())
 *                 .build())
 *             .azurermSpnTenantid("00000000-0000-0000-0000-000000000000")
 *             .azurermSubscriptionId("00000000-0000-0000-0000-000000000000")
 *             .azurermSubscriptionName("Example Subscription Name")
 *             .build());
 * 
 *         var exampleFederatedIdentityCredential = new FederatedIdentityCredential("exampleFederatedIdentityCredential", FederatedIdentityCredentialArgs.builder()
 *             .name("example-federated-credential")
 *             .resourceGroupName(identity.name())
 *             .parentId(exampleUserAssignedIdentity.id())
 *             .audience("api://AzureADTokenExchange")
 *             .issuer(exampleServiceEndpointAzureRM.workloadIdentityFederationIssuer())
 *             .subject(exampleServiceEndpointAzureRM.workloadIdentityFederationSubject())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ### Workload Identity Federation Automatic AzureRM Service Endpoint
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
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
 *         var example = new Project("example", ProjectArgs.builder()
 *             .name("Example Project")
 *             .visibility("private")
 *             .versionControl("Git")
 *             .workItemTemplate("Agile")
 *             .build());
 * 
 *         var exampleServiceEndpointAzureRM = new ServiceEndpointAzureRM("exampleServiceEndpointAzureRM", ServiceEndpointAzureRMArgs.builder()
 *             .projectId(example.id())
 *             .serviceEndpointName("Example AzureRM")
 *             .serviceEndpointAuthenticationScheme("WorkloadIdentityFederation")
 *             .azurermSpnTenantid("00000000-0000-0000-0000-000000000000")
 *             .azurermSubscriptionId("00000000-0000-0000-0000-000000000000")
 *             .azurermSubscriptionName("Example Subscription Name")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ### Managed Identity AzureRM Service Endpoint
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
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
 *         var example = new Project("example", ProjectArgs.builder()
 *             .name("Example Project")
 *             .visibility("private")
 *             .versionControl("Git")
 *             .workItemTemplate("Agile")
 *             .build());
 * 
 *         var exampleServiceEndpointAzureRM = new ServiceEndpointAzureRM("exampleServiceEndpointAzureRM", ServiceEndpointAzureRMArgs.builder()
 *             .projectId(example.id())
 *             .serviceEndpointName("Example AzureRM")
 *             .serviceEndpointAuthenticationScheme("ManagedServiceIdentity")
 *             .azurermSpnTenantid("00000000-0000-0000-0000-000000000000")
 *             .azurermSubscriptionId("00000000-0000-0000-0000-000000000000")
 *             .azurermSubscriptionName("Example Subscription Name")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Relevant Links
 * 
 * - [Azure DevOps Service REST API 7.0 - Service End points](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-7.0)
 * 
 * ## Import
 * 
 * Azure DevOps Azure Resource Manager Service Endpoint can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**
 * 
 * ```sh
 * $ pulumi import azuredevops:index/serviceEndpointAzureRM:ServiceEndpointAzureRM example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
 * ```
 * 
 */
@ResourceType(type="azuredevops:index/serviceEndpointAzureRM:ServiceEndpointAzureRM")
public class ServiceEndpointAzureRM extends com.pulumi.resources.CustomResource {
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
     * The Tenant ID of the service principal.
     * 
     */
    @Export(name="azurermSpnTenantid", refs={String.class}, tree="[0]")
    private Output<String> azurermSpnTenantid;

    /**
     * @return The Tenant ID of the service principal.
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
     * A `credentials` block as defined below.
     * 
     */
    @Export(name="credentials", refs={ServiceEndpointAzureRMCredentials.class}, tree="[0]")
    private Output</* @Nullable */ ServiceEndpointAzureRMCredentials> credentials;

    /**
     * @return A `credentials` block as defined below.
     * 
     */
    public Output<Optional<ServiceEndpointAzureRMCredentials>> credentials() {
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
     * The Cloud Environment to use. Defaults to `AzureCloud`. Possible values are `AzureCloud`, `AzureChinaCloud`, `AzureUSGovernment`, `AzureGermanCloud` and `AzureStack`. Changing this forces a new resource to be created.
     * 
     */
    @Export(name="environment", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> environment;

    /**
     * @return The Cloud Environment to use. Defaults to `AzureCloud`. Possible values are `AzureCloud`, `AzureChinaCloud`, `AzureUSGovernment`, `AzureGermanCloud` and `AzureStack`. Changing this forces a new resource to be created.
     * 
     */
    public Output<Optional<String>> environment() {
        return Codegen.optional(this.environment);
    }
    /**
     * A `features` block as defined below.
     * 
     */
    @Export(name="features", refs={ServiceEndpointAzureRMFeatures.class}, tree="[0]")
    private Output</* @Nullable */ ServiceEndpointAzureRMFeatures> features;

    /**
     * @return A `features` block as defined below.
     * 
     */
    public Output<Optional<ServiceEndpointAzureRMFeatures>> features() {
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
     * The server URL of the service endpoint. Changing this forces a new resource to be created.
     * 
     * &gt; **NOTE:** One of either `Subscription` scoped i.e. `azurerm_subscription_id`, `azurerm_subscription_name` or `ManagementGroup` scoped i.e. `azurerm_management_group_id`, `azurerm_management_group_name` values must be specified.
     * 
     */
    @Export(name="serverUrl", refs={String.class}, tree="[0]")
    private Output<String> serverUrl;

    /**
     * @return The server URL of the service endpoint. Changing this forces a new resource to be created.
     * 
     * &gt; **NOTE:** One of either `Subscription` scoped i.e. `azurerm_subscription_id`, `azurerm_subscription_name` or `ManagementGroup` scoped i.e. `azurerm_management_group_id`, `azurerm_management_group_name` values must be specified.
     * 
     */
    public Output<String> serverUrl() {
        return this.serverUrl;
    }
    /**
     * Specifies the type of Azure Resource Manager Service Endpoint. Possible values are `WorkloadIdentityFederation`, `ManagedServiceIdentity` or `ServicePrincipal`. Defaults to `ServicePrincipal` for backwards compatibility.
     * 
     * &gt; **NOTE:** The `WorkloadIdentityFederation` authentication scheme is currently in private preview. Your organisation must be part of the preview and the feature toggle must be turned on to use it. More details can be found [here](https://aka.ms/azdo-rm-workload-identity).
     * 
     */
    @Export(name="serviceEndpointAuthenticationScheme", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> serviceEndpointAuthenticationScheme;

    /**
     * @return Specifies the type of Azure Resource Manager Service Endpoint. Possible values are `WorkloadIdentityFederation`, `ManagedServiceIdentity` or `ServicePrincipal`. Defaults to `ServicePrincipal` for backwards compatibility.
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
    public ServiceEndpointAzureRM(java.lang.String name) {
        this(name, ServiceEndpointAzureRMArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ServiceEndpointAzureRM(java.lang.String name, ServiceEndpointAzureRMArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ServiceEndpointAzureRM(java.lang.String name, ServiceEndpointAzureRMArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/serviceEndpointAzureRM:ServiceEndpointAzureRM", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private ServiceEndpointAzureRM(java.lang.String name, Output<java.lang.String> id, @Nullable ServiceEndpointAzureRMState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/serviceEndpointAzureRM:ServiceEndpointAzureRM", name, state, makeResourceOptions(options, id), false);
    }

    private static ServiceEndpointAzureRMArgs makeArgs(ServiceEndpointAzureRMArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? ServiceEndpointAzureRMArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
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
    public static ServiceEndpointAzureRM get(java.lang.String name, Output<java.lang.String> id, @Nullable ServiceEndpointAzureRMState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ServiceEndpointAzureRM(name, id, state, options);
    }
}
