// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops;

import com.pulumi.azuredevops.ServiceEndpointAzureEcrArgs;
import com.pulumi.azuredevops.Utilities;
import com.pulumi.azuredevops.inputs.ServiceEndpointAzureEcrState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a Azure Container Registry service endpoint within Azure DevOps.
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.azuredevops.Project;
 * import com.pulumi.azuredevops.ProjectArgs;
 * import com.pulumi.azuredevops.ServiceEndpointAzureEcr;
 * import com.pulumi.azuredevops.ServiceEndpointAzureEcrArgs;
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
 *         var exampleServiceEndpointAzureEcr = new ServiceEndpointAzureEcr(&#34;exampleServiceEndpointAzureEcr&#34;, ServiceEndpointAzureEcrArgs.builder()        
 *             .projectId(exampleProject.id())
 *             .serviceEndpointName(&#34;Example AzureCR&#34;)
 *             .resourceGroup(&#34;example-rg&#34;)
 *             .azurecrSpnTenantid(&#34;00000000-0000-0000-0000-000000000000&#34;)
 *             .azurecrName(&#34;ExampleAcr&#34;)
 *             .azurecrSubscriptionId(&#34;00000000-0000-0000-0000-000000000000&#34;)
 *             .azurecrSubscriptionName(&#34;subscription name&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Relevant Links
 * 
 * - [Azure DevOps Service REST API 7.0 - Service Endpoints](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-7.0)
 * - [Azure Container Registry REST API](https://docs.microsoft.com/en-us/rest/api/containerregistry/)
 * 
 * ## Import
 * 
 * Azure DevOps Service Endpoint Azure Container Registry can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**
 * 
 * ```sh
 * $ pulumi import azuredevops:index/serviceEndpointAzureEcr:ServiceEndpointAzureEcr example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
 * ```
 * 
 */
@ResourceType(type="azuredevops:index/serviceEndpointAzureEcr:ServiceEndpointAzureEcr")
public class ServiceEndpointAzureEcr extends com.pulumi.resources.CustomResource {
    @Export(name="appObjectId", refs={String.class}, tree="[0]")
    private Output<String> appObjectId;

    public Output<String> appObjectId() {
        return this.appObjectId;
    }
    @Export(name="authorization", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> authorization;

    public Output<Map<String,String>> authorization() {
        return this.authorization;
    }
    @Export(name="azSpnRoleAssignmentId", refs={String.class}, tree="[0]")
    private Output<String> azSpnRoleAssignmentId;

    public Output<String> azSpnRoleAssignmentId() {
        return this.azSpnRoleAssignmentId;
    }
    @Export(name="azSpnRolePermissions", refs={String.class}, tree="[0]")
    private Output<String> azSpnRolePermissions;

    public Output<String> azSpnRolePermissions() {
        return this.azSpnRolePermissions;
    }
    /**
     * The Azure container registry name.
     * 
     */
    @Export(name="azurecrName", refs={String.class}, tree="[0]")
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
    @Export(name="azurecrSpnTenantid", refs={String.class}, tree="[0]")
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
    @Export(name="azurecrSubscriptionId", refs={String.class}, tree="[0]")
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
    @Export(name="azurecrSubscriptionName", refs={String.class}, tree="[0]")
    private Output<String> azurecrSubscriptionName;

    /**
     * @return The subscription name of the Azure targets.
     * 
     */
    public Output<String> azurecrSubscriptionName() {
        return this.azurecrSubscriptionName;
    }
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
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
     * The resource group to which the container registry belongs.
     * 
     */
    @Export(name="resourceGroup", refs={String.class}, tree="[0]")
    private Output<String> resourceGroup;

    /**
     * @return The resource group to which the container registry belongs.
     * 
     */
    public Output<String> resourceGroup() {
        return this.resourceGroup;
    }
    /**
     * The name you will use to refer to this service connection in task inputs.
     * 
     */
    @Export(name="serviceEndpointName", refs={String.class}, tree="[0]")
    private Output<String> serviceEndpointName;

    /**
     * @return The name you will use to refer to this service connection in task inputs.
     * 
     */
    public Output<String> serviceEndpointName() {
        return this.serviceEndpointName;
    }
    /**
     * The service principal ID.
     * 
     */
    @Export(name="servicePrincipalId", refs={String.class}, tree="[0]")
    private Output<String> servicePrincipalId;

    /**
     * @return The service principal ID.
     * 
     */
    public Output<String> servicePrincipalId() {
        return this.servicePrincipalId;
    }
    @Export(name="spnObjectId", refs={String.class}, tree="[0]")
    private Output<String> spnObjectId;

    public Output<String> spnObjectId() {
        return this.spnObjectId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ServiceEndpointAzureEcr(String name) {
        this(name, ServiceEndpointAzureEcrArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ServiceEndpointAzureEcr(String name, ServiceEndpointAzureEcrArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ServiceEndpointAzureEcr(String name, ServiceEndpointAzureEcrArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/serviceEndpointAzureEcr:ServiceEndpointAzureEcr", name, args == null ? ServiceEndpointAzureEcrArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ServiceEndpointAzureEcr(String name, Output<String> id, @Nullable ServiceEndpointAzureEcrState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/serviceEndpointAzureEcr:ServiceEndpointAzureEcr", name, state, makeResourceOptions(options, id));
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
    public static ServiceEndpointAzureEcr get(String name, Output<String> id, @Nullable ServiceEndpointAzureEcrState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ServiceEndpointAzureEcr(name, id, state, options);
    }
}
