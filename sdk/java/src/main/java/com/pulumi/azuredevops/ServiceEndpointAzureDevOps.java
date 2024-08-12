// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops;

import com.pulumi.azuredevops.ServiceEndpointAzureDevOpsArgs;
import com.pulumi.azuredevops.Utilities;
import com.pulumi.azuredevops.inputs.ServiceEndpointAzureDevOpsState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages an Azure DevOps service endpoint within Azure DevOps.
 * 
 * &gt; **Note** This resource is duplicate with `azuredevops.ServiceEndpointPipeline`,  will be removed in the future, use `azuredevops.ServiceEndpointPipeline` instead.
 * 
 * &gt; **Note** Prerequisite: Extension [Configurable Pipeline Runner](https://marketplace.visualstudio.com/items?itemName=CSE-DevOps.RunPipelines) has been installed for the organization.
 * 
 * ## Example Usage
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
 * import com.pulumi.azuredevops.ServiceEndpointAzureDevOps;
 * import com.pulumi.azuredevops.ServiceEndpointAzureDevOpsArgs;
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
 *             .description("Managed by Terraform")
 *             .build());
 * 
 *         var exampleServiceEndpointAzureDevOps = new ServiceEndpointAzureDevOps("exampleServiceEndpointAzureDevOps", ServiceEndpointAzureDevOpsArgs.builder()
 *             .projectId(example.id())
 *             .serviceEndpointName("Example Azure DevOps")
 *             .orgUrl("https://dev.azure.com/testorganization")
 *             .releaseApiUrl("https://vsrm.dev.azure.com/testorganization")
 *             .personalAccessToken("0000000000000000000000000000000000000000000000000000")
 *             .description("Managed by Terraform")
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
 * - [Azure DevOps Service REST API 7.0 - Service Endpoints](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-7.0)
 * 
 * ## Import
 * 
 * Azure DevOps Service Endpoint Azure DevOps can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**
 * 
 * ```sh
 * $ pulumi import azuredevops:index/serviceEndpointAzureDevOps:ServiceEndpointAzureDevOps example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
 * ```
 * 
 */
@ResourceType(type="azuredevops:index/serviceEndpointAzureDevOps:ServiceEndpointAzureDevOps")
public class ServiceEndpointAzureDevOps extends com.pulumi.resources.CustomResource {
    @Export(name="authorization", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> authorization;

    public Output<Map<String,String>> authorization() {
        return this.authorization;
    }
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The organization URL.
     * 
     */
    @Export(name="orgUrl", refs={String.class}, tree="[0]")
    private Output<String> orgUrl;

    /**
     * @return The organization URL.
     * 
     */
    public Output<String> orgUrl() {
        return this.orgUrl;
    }
    /**
     * The Azure DevOps personal access token.
     * 
     */
    @Export(name="personalAccessToken", refs={String.class}, tree="[0]")
    private Output<String> personalAccessToken;

    /**
     * @return The Azure DevOps personal access token.
     * 
     */
    public Output<String> personalAccessToken() {
        return this.personalAccessToken;
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
     * The URL of the release API.
     * 
     */
    @Export(name="releaseApiUrl", refs={String.class}, tree="[0]")
    private Output<String> releaseApiUrl;

    /**
     * @return The URL of the release API.
     * 
     */
    public Output<String> releaseApiUrl() {
        return this.releaseApiUrl;
    }
    /**
     * The Service Endpoint name.
     * 
     */
    @Export(name="serviceEndpointName", refs={String.class}, tree="[0]")
    private Output<String> serviceEndpointName;

    /**
     * @return The Service Endpoint name.
     * 
     */
    public Output<String> serviceEndpointName() {
        return this.serviceEndpointName;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ServiceEndpointAzureDevOps(java.lang.String name) {
        this(name, ServiceEndpointAzureDevOpsArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ServiceEndpointAzureDevOps(java.lang.String name, ServiceEndpointAzureDevOpsArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ServiceEndpointAzureDevOps(java.lang.String name, ServiceEndpointAzureDevOpsArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/serviceEndpointAzureDevOps:ServiceEndpointAzureDevOps", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private ServiceEndpointAzureDevOps(java.lang.String name, Output<java.lang.String> id, @Nullable ServiceEndpointAzureDevOpsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/serviceEndpointAzureDevOps:ServiceEndpointAzureDevOps", name, state, makeResourceOptions(options, id), false);
    }

    private static ServiceEndpointAzureDevOpsArgs makeArgs(ServiceEndpointAzureDevOpsArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? ServiceEndpointAzureDevOpsArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "personalAccessToken"
            ))
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
    public static ServiceEndpointAzureDevOps get(java.lang.String name, Output<java.lang.String> id, @Nullable ServiceEndpointAzureDevOpsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ServiceEndpointAzureDevOps(name, id, state, options);
    }
}
