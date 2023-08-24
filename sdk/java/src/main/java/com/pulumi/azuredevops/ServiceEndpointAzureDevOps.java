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
 * ```java
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
 *         var exampleProject = new Project(&#34;exampleProject&#34;, ProjectArgs.builder()        
 *             .visibility(&#34;private&#34;)
 *             .versionControl(&#34;Git&#34;)
 *             .workItemTemplate(&#34;Agile&#34;)
 *             .description(&#34;Managed by Terraform&#34;)
 *             .build());
 * 
 *         var exampleServiceEndpointAzureDevOps = new ServiceEndpointAzureDevOps(&#34;exampleServiceEndpointAzureDevOps&#34;, ServiceEndpointAzureDevOpsArgs.builder()        
 *             .projectId(exampleProject.id())
 *             .serviceEndpointName(&#34;Example Azure DevOps&#34;)
 *             .orgUrl(&#34;https://dev.azure.com/testorganization&#34;)
 *             .releaseApiUrl(&#34;https://vsrm.dev.azure.com/testorganization&#34;)
 *             .personalAccessToken(&#34;0000000000000000000000000000000000000000000000000000&#34;)
 *             .description(&#34;Managed by Terraform&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ## Relevant Links
 * 
 * - [Azure DevOps Service REST API 6.0 - Service Endpoints](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-6.0)
 * 
 * ## Import
 * 
 * Azure DevOps Service Endpoint Azure DevOps can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**
 * 
 * ```sh
 *  $ pulumi import azuredevops:index/serviceEndpointAzureDevOps:ServiceEndpointAzureDevOps example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
 * ```
 * 
 */
@ResourceType(type="azuredevops:index/serviceEndpointAzureDevOps:ServiceEndpointAzureDevOps")
public class ServiceEndpointAzureDevOps extends com.pulumi.resources.CustomResource {
    @Export(name="authorization", type=Map.class, parameters={String.class, String.class})
    private Output<Map<String,String>> authorization;

    public Output<Map<String,String>> authorization() {
        return this.authorization;
    }
    @Export(name="description", type=String.class, parameters={})
    private Output</* @Nullable */ String> description;

    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The organization URL.
     * 
     */
    @Export(name="orgUrl", type=String.class, parameters={})
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
    @Export(name="personalAccessToken", type=String.class, parameters={})
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
    @Export(name="projectId", type=String.class, parameters={})
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
    @Export(name="releaseApiUrl", type=String.class, parameters={})
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
    @Export(name="serviceEndpointName", type=String.class, parameters={})
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
    public ServiceEndpointAzureDevOps(String name) {
        this(name, ServiceEndpointAzureDevOpsArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ServiceEndpointAzureDevOps(String name, ServiceEndpointAzureDevOpsArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ServiceEndpointAzureDevOps(String name, ServiceEndpointAzureDevOpsArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/serviceEndpointAzureDevOps:ServiceEndpointAzureDevOps", name, args == null ? ServiceEndpointAzureDevOpsArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ServiceEndpointAzureDevOps(String name, Output<String> id, @Nullable ServiceEndpointAzureDevOpsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/serviceEndpointAzureDevOps:ServiceEndpointAzureDevOps", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
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
    public static ServiceEndpointAzureDevOps get(String name, Output<String> id, @Nullable ServiceEndpointAzureDevOpsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ServiceEndpointAzureDevOps(name, id, state, options);
    }
}
