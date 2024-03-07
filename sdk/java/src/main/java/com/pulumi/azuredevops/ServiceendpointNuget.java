// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops;

import com.pulumi.azuredevops.ServiceendpointNugetArgs;
import com.pulumi.azuredevops.Utilities;
import com.pulumi.azuredevops.inputs.ServiceendpointNugetState;
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
 * Manages a NuGet service endpoint within Azure DevOps.
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
 * import com.pulumi.azuredevops.ServiceendpointNuget;
 * import com.pulumi.azuredevops.ServiceendpointNugetArgs;
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
 *         var exampleServiceendpointNuget = new ServiceendpointNuget(&#34;exampleServiceendpointNuget&#34;, ServiceendpointNugetArgs.builder()        
 *             .projectId(exampleProject.id())
 *             .apiKey(&#34;apikey&#34;)
 *             .serviceEndpointName(&#34;Example NuGet&#34;)
 *             .description(&#34;Managed by Terraform&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Relevant Links
 * 
 * - [Azure DevOps Service REST API 7.0 - Agent Pools](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-7.0)
 * 
 * ## Import
 * 
 * Azure DevOps Service Endpoint NuGet can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**
 * 
 * ```sh
 * $ pulumi import azuredevops:index/serviceendpointNuget:ServiceendpointNuget example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
 * ```
 * 
 */
@ResourceType(type="azuredevops:index/serviceendpointNuget:ServiceendpointNuget")
public class ServiceendpointNuget extends com.pulumi.resources.CustomResource {
    /**
     * The API Key used to connect to the endpoint.
     * 
     */
    @Export(name="apiKey", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> apiKey;

    /**
     * @return The API Key used to connect to the endpoint.
     * 
     */
    public Output<Optional<String>> apiKey() {
        return Codegen.optional(this.apiKey);
    }
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
     * The URL for the feed. This will generally end with `index.json`.
     * 
     */
    @Export(name="feedUrl", refs={String.class}, tree="[0]")
    private Output<String> feedUrl;

    /**
     * @return The URL for the feed. This will generally end with `index.json`.
     * 
     */
    public Output<String> feedUrl() {
        return this.feedUrl;
    }
    /**
     * The account password used to connect to the endpoint
     * 
     * &gt; **Note** Only one of `api_key` or `personal_access_token` or  `username`, `password` can be set at the same time.
     * 
     */
    @Export(name="password", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> password;

    /**
     * @return The account password used to connect to the endpoint
     * 
     * &gt; **Note** Only one of `api_key` or `personal_access_token` or  `username`, `password` can be set at the same time.
     * 
     */
    public Output<Optional<String>> password() {
        return Codegen.optional(this.password);
    }
    /**
     * The Personal access token used to  connect to the endpoint. Personal access tokens are applicable only for NuGet feeds hosted on other Azure DevOps Services organizations or Azure DevOps Server 2019 (or later).
     * 
     */
    @Export(name="personalAccessToken", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> personalAccessToken;

    /**
     * @return The Personal access token used to  connect to the endpoint. Personal access tokens are applicable only for NuGet feeds hosted on other Azure DevOps Services organizations or Azure DevOps Server 2019 (or later).
     * 
     */
    public Output<Optional<String>> personalAccessToken() {
        return Codegen.optional(this.personalAccessToken);
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
     * The account username used to connect to the endpoint.
     * 
     */
    @Export(name="username", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> username;

    /**
     * @return The account username used to connect to the endpoint.
     * 
     */
    public Output<Optional<String>> username() {
        return Codegen.optional(this.username);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ServiceendpointNuget(String name) {
        this(name, ServiceendpointNugetArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ServiceendpointNuget(String name, ServiceendpointNugetArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ServiceendpointNuget(String name, ServiceendpointNugetArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/serviceendpointNuget:ServiceendpointNuget", name, args == null ? ServiceendpointNugetArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ServiceendpointNuget(String name, Output<String> id, @Nullable ServiceendpointNugetState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/serviceendpointNuget:ServiceendpointNuget", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "apiKey",
                "password",
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
    public static ServiceendpointNuget get(String name, Output<String> id, @Nullable ServiceendpointNugetState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ServiceendpointNuget(name, id, state, options);
    }
}
