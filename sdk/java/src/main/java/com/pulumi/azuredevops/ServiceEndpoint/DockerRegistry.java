// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.ServiceEndpoint;

import com.pulumi.azuredevops.ServiceEndpoint.DockerRegistryArgs;
import com.pulumi.azuredevops.ServiceEndpoint.inputs.DockerRegistryState;
import com.pulumi.azuredevops.Utilities;
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
 * Manages a Docker Registry service endpoint within Azure DevOps.
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
 * import com.pulumi.azuredevops.ServiceEndpointDockerRegistry;
 * import com.pulumi.azuredevops.ServiceEndpointDockerRegistryArgs;
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
 *         var exampleServiceEndpointDockerRegistry = new ServiceEndpointDockerRegistry(&#34;exampleServiceEndpointDockerRegistry&#34;, ServiceEndpointDockerRegistryArgs.builder()        
 *             .projectId(exampleProject.id())
 *             .serviceEndpointName(&#34;Example Docker Hub&#34;)
 *             .dockerUsername(&#34;example&#34;)
 *             .dockerEmail(&#34;email@example.com&#34;)
 *             .dockerPassword(&#34;12345&#34;)
 *             .registryType(&#34;DockerHub&#34;)
 *             .build());
 * 
 *         var example_other = new ServiceEndpointDockerRegistry(&#34;example-other&#34;, ServiceEndpointDockerRegistryArgs.builder()        
 *             .projectId(exampleProject.id())
 *             .serviceEndpointName(&#34;Example Docker Registry&#34;)
 *             .dockerRegistry(&#34;https://sample.azurecr.io/v1&#34;)
 *             .dockerUsername(&#34;sample&#34;)
 *             .dockerPassword(&#34;12345&#34;)
 *             .registryType(&#34;Others&#34;)
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
 * - [Docker Registry Service Connection](https://docs.microsoft.com/en-us/azure/devops/pipelines/library/service-endpoints?view=azure-devops&amp;tabs=yaml#sep-docreg)
 * 
 * ## Import
 * 
 * Azure DevOps Service Endpoint Docker Registry can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**
 * 
 * ```sh
 * $ pulumi import azuredevops:ServiceEndpoint/dockerRegistry:DockerRegistry example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
 * ```
 * 
 * @deprecated
 * azuredevops.serviceendpoint.DockerRegistry has been deprecated in favor of azuredevops.ServiceEndpointDockerRegistry
 * 
 */
@Deprecated /* azuredevops.serviceendpoint.DockerRegistry has been deprecated in favor of azuredevops.ServiceEndpointDockerRegistry */
@ResourceType(type="azuredevops:ServiceEndpoint/dockerRegistry:DockerRegistry")
public class DockerRegistry extends com.pulumi.resources.CustomResource {
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
     * The email for Docker account user.
     * 
     */
    @Export(name="dockerEmail", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> dockerEmail;

    /**
     * @return The email for Docker account user.
     * 
     */
    public Output<Optional<String>> dockerEmail() {
        return Codegen.optional(this.dockerEmail);
    }
    /**
     * The password for the account user identified above.
     * 
     */
    @Export(name="dockerPassword", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> dockerPassword;

    /**
     * @return The password for the account user identified above.
     * 
     */
    public Output<Optional<String>> dockerPassword() {
        return Codegen.optional(this.dockerPassword);
    }
    /**
     * The URL of the Docker registry. (Default: &#34;https://index.docker.io/v1/&#34;)
     * 
     */
    @Export(name="dockerRegistry", refs={String.class}, tree="[0]")
    private Output<String> dockerRegistry;

    /**
     * @return The URL of the Docker registry. (Default: &#34;https://index.docker.io/v1/&#34;)
     * 
     */
    public Output<String> dockerRegistry() {
        return this.dockerRegistry;
    }
    /**
     * The identifier of the Docker account user.
     * 
     */
    @Export(name="dockerUsername", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> dockerUsername;

    /**
     * @return The identifier of the Docker account user.
     * 
     */
    public Output<Optional<String>> dockerUsername() {
        return Codegen.optional(this.dockerUsername);
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
     * Can be &#34;DockerHub&#34; or &#34;Others&#34; (Default &#34;DockerHub&#34;)
     * 
     */
    @Export(name="registryType", refs={String.class}, tree="[0]")
    private Output<String> registryType;

    /**
     * @return Can be &#34;DockerHub&#34; or &#34;Others&#34; (Default &#34;DockerHub&#34;)
     * 
     */
    public Output<String> registryType() {
        return this.registryType;
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
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public DockerRegistry(String name) {
        this(name, DockerRegistryArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public DockerRegistry(String name, DockerRegistryArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public DockerRegistry(String name, DockerRegistryArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:ServiceEndpoint/dockerRegistry:DockerRegistry", name, args == null ? DockerRegistryArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private DockerRegistry(String name, Output<String> id, @Nullable DockerRegistryState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:ServiceEndpoint/dockerRegistry:DockerRegistry", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "dockerPassword"
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
    public static DockerRegistry get(String name, Output<String> id, @Nullable DockerRegistryState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new DockerRegistry(name, id, state, options);
    }
}
