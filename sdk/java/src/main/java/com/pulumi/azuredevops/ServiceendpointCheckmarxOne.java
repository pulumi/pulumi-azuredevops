// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops;

import com.pulumi.azuredevops.ServiceendpointCheckmarxOneArgs;
import com.pulumi.azuredevops.Utilities;
import com.pulumi.azuredevops.inputs.ServiceendpointCheckmarxOneState;
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
 * Manages a Checkmarx One service endpoint within Azure DevOps. Using this service endpoint requires you to install: [Checkmarx AST](https://marketplace.visualstudio.com/items?itemName=checkmarx.checkmarx-ast-azure-plugin)
 * 
 * ## Example Usage
 * 
 * ### Authorize with API Key
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
 * import com.pulumi.azuredevops.ServiceendpointCheckmarxOne;
 * import com.pulumi.azuredevops.ServiceendpointCheckmarxOneArgs;
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
 *         var exampleServiceendpointCheckmarxOne = new ServiceendpointCheckmarxOne("exampleServiceendpointCheckmarxOne", ServiceendpointCheckmarxOneArgs.builder()
 *             .projectId(example.id())
 *             .serviceEndpointName("Example Checkmarx One")
 *             .serverUrl("https://server.com")
 *             .apiKey("apikey")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ### Authorize with Client ID and Secret
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
 * import com.pulumi.azuredevops.ServiceendpointCheckmarxOne;
 * import com.pulumi.azuredevops.ServiceendpointCheckmarxOneArgs;
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
 *         var exampleServiceendpointCheckmarxOne = new ServiceendpointCheckmarxOne("exampleServiceendpointCheckmarxOne", ServiceendpointCheckmarxOneArgs.builder()
 *             .projectId(example.id())
 *             .serviceEndpointName("Example Checkmarx One")
 *             .serverUrl("https://server.com")
 *             .clientId("clientid")
 *             .clientSecret("secret")
 *             .authorizationUrl("https://authurl.com")
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
 * Azure DevOps Service Endpoint Check Marx One can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**
 * 
 * ```sh
 * $ pulumi import azuredevops:index/serviceendpointCheckmarxOne:ServiceendpointCheckmarxOne example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
 * ```
 * 
 */
@ResourceType(type="azuredevops:index/serviceendpointCheckmarxOne:ServiceendpointCheckmarxOne")
public class ServiceendpointCheckmarxOne extends com.pulumi.resources.CustomResource {
    /**
     * The account of the Checkmarx One. Conflict with `client_id` and `client_secret`.
     * 
     */
    @Export(name="apiKey", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> apiKey;

    /**
     * @return The account of the Checkmarx One. Conflict with `client_id` and `client_secret`.
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
    /**
     * The URL of Checkmarx Authorization. Used when using `client_id` and `client_secret` authorization.
     * 
     */
    @Export(name="authorizationUrl", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> authorizationUrl;

    /**
     * @return The URL of Checkmarx Authorization. Used when using `client_id` and `client_secret` authorization.
     * 
     */
    public Output<Optional<String>> authorizationUrl() {
        return Codegen.optional(this.authorizationUrl);
    }
    /**
     * The Client ID of the Checkmarx One. Conflict with `api_key`
     * 
     */
    @Export(name="clientId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> clientId;

    /**
     * @return The Client ID of the Checkmarx One. Conflict with `api_key`
     * 
     */
    public Output<Optional<String>> clientId() {
        return Codegen.optional(this.clientId);
    }
    /**
     * The Client Secret of the Checkmarx One. Conflict with `api_key`
     * 
     * &gt; **Note** At least one of `api_key` and `client_id`, `client_secret` must be set
     * 
     */
    @Export(name="clientSecret", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> clientSecret;

    /**
     * @return The Client Secret of the Checkmarx One. Conflict with `api_key`
     * 
     * &gt; **Note** At least one of `api_key` and `client_id`, `client_secret` must be set
     * 
     */
    public Output<Optional<String>> clientSecret() {
        return Codegen.optional(this.clientSecret);
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
     * The Server URL of the Checkmarx One Service.
     * 
     */
    @Export(name="serverUrl", refs={String.class}, tree="[0]")
    private Output<String> serverUrl;

    /**
     * @return The Server URL of the Checkmarx One Service.
     * 
     */
    public Output<String> serverUrl() {
        return this.serverUrl;
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
    public ServiceendpointCheckmarxOne(java.lang.String name) {
        this(name, ServiceendpointCheckmarxOneArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ServiceendpointCheckmarxOne(java.lang.String name, ServiceendpointCheckmarxOneArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ServiceendpointCheckmarxOne(java.lang.String name, ServiceendpointCheckmarxOneArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/serviceendpointCheckmarxOne:ServiceendpointCheckmarxOne", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private ServiceendpointCheckmarxOne(java.lang.String name, Output<java.lang.String> id, @Nullable ServiceendpointCheckmarxOneState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/serviceendpointCheckmarxOne:ServiceendpointCheckmarxOne", name, state, makeResourceOptions(options, id), false);
    }

    private static ServiceendpointCheckmarxOneArgs makeArgs(ServiceendpointCheckmarxOneArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? ServiceendpointCheckmarxOneArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "apiKey",
                "clientSecret"
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
    public static ServiceendpointCheckmarxOne get(java.lang.String name, Output<java.lang.String> id, @Nullable ServiceendpointCheckmarxOneState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ServiceendpointCheckmarxOne(name, id, state, options);
    }
}
