// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops;

import com.pulumi.azuredevops.ServiceendpointJfrogPlatformV2Args;
import com.pulumi.azuredevops.Utilities;
import com.pulumi.azuredevops.inputs.ServiceendpointJfrogPlatformV2State;
import com.pulumi.azuredevops.outputs.ServiceendpointJfrogPlatformV2AuthenticationBasic;
import com.pulumi.azuredevops.outputs.ServiceendpointJfrogPlatformV2AuthenticationToken;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a JFrog Platform V2 service endpoint within an Azure DevOps organization.
 * 
 * &gt; **Note:** Using this service endpoint requires you to first install [JFrog Extension](https://marketplace.visualstudio.com/items?itemName=JFrog.jfrog-azure-devops-extension).
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
 * import com.pulumi.azuredevops.ServiceendpointJfrogPlatformV2;
 * import com.pulumi.azuredevops.ServiceendpointJfrogPlatformV2Args;
 * import com.pulumi.azuredevops.inputs.ServiceendpointJfrogPlatformV2AuthenticationTokenArgs;
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
 *         var exampleServiceendpointJfrogPlatformV2 = new ServiceendpointJfrogPlatformV2("exampleServiceendpointJfrogPlatformV2", ServiceendpointJfrogPlatformV2Args.builder()
 *             .projectId(example.id())
 *             .serviceEndpointName("Example Artifactory")
 *             .description("Managed by Pulumi")
 *             .url("https://artifactory.my.com")
 *             .authenticationToken(ServiceendpointJfrogPlatformV2AuthenticationTokenArgs.builder()
 *                 .token("0000000000000000000000000000000000000000")
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * Alternatively a username and password may be used.
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
 * import com.pulumi.azuredevops.ServiceendpointJfrogPlatformV2;
 * import com.pulumi.azuredevops.ServiceendpointJfrogPlatformV2Args;
 * import com.pulumi.azuredevops.inputs.ServiceendpointJfrogPlatformV2AuthenticationBasicArgs;
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
 *         var exampleServiceendpointJfrogPlatformV2 = new ServiceendpointJfrogPlatformV2("exampleServiceendpointJfrogPlatformV2", ServiceendpointJfrogPlatformV2Args.builder()
 *             .projectId(example.id())
 *             .serviceEndpointName("Example Artifactory")
 *             .description("Managed by Pulumi")
 *             .url("https://artifactory.my.com")
 *             .authenticationBasic(ServiceendpointJfrogPlatformV2AuthenticationBasicArgs.builder()
 *                 .username("username")
 *                 .password("password")
 *                 .build())
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
 * * [Azure DevOps Service Connections](https://docs.microsoft.com/en-us/azure/devops/pipelines/library/service-endpoints?view=azure-devops&amp;tabs=yaml)
 * * [Artifactory User Token](https://docs.artifactory.org/latest/user-guide/user-token/)
 * 
 * ## Import
 * 
 * Azure DevOps JFrog Platform V2 Service Endpoint can be imported using the **projectID/serviceEndpointID**, e.g.
 * 
 * ```sh
 * $ pulumi import azuredevops:index/serviceendpointJfrogPlatformV2:ServiceendpointJfrogPlatformV2 example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
 * ```
 * 
 */
@ResourceType(type="azuredevops:index/serviceendpointJfrogPlatformV2:ServiceendpointJfrogPlatformV2")
public class ServiceendpointJfrogPlatformV2 extends com.pulumi.resources.CustomResource {
    /**
     * An `authentication_basic` block as documented below.
     * 
     */
    @Export(name="authenticationBasic", refs={ServiceendpointJfrogPlatformV2AuthenticationBasic.class}, tree="[0]")
    private Output</* @Nullable */ ServiceendpointJfrogPlatformV2AuthenticationBasic> authenticationBasic;

    /**
     * @return An `authentication_basic` block as documented below.
     * 
     */
    public Output<Optional<ServiceendpointJfrogPlatformV2AuthenticationBasic>> authenticationBasic() {
        return Codegen.optional(this.authenticationBasic);
    }
    /**
     * An `authentication_token` block as documented below.
     * 
     */
    @Export(name="authenticationToken", refs={ServiceendpointJfrogPlatformV2AuthenticationToken.class}, tree="[0]")
    private Output</* @Nullable */ ServiceendpointJfrogPlatformV2AuthenticationToken> authenticationToken;

    /**
     * @return An `authentication_token` block as documented below.
     * 
     */
    public Output<Optional<ServiceendpointJfrogPlatformV2AuthenticationToken>> authenticationToken() {
        return Codegen.optional(this.authenticationToken);
    }
    @Export(name="authorization", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> authorization;

    public Output<Map<String,String>> authorization() {
        return this.authorization;
    }
    /**
     * The Service Endpoint description.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return The Service Endpoint description.
     * 
     */
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
     * URL of the Artifactory server to connect with.
     * 
     * &gt; **NOTE:** URL should not end in a slash character.
     * 
     */
    @Export(name="url", refs={String.class}, tree="[0]")
    private Output<String> url;

    /**
     * @return URL of the Artifactory server to connect with.
     * 
     * &gt; **NOTE:** URL should not end in a slash character.
     * 
     */
    public Output<String> url() {
        return this.url;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ServiceendpointJfrogPlatformV2(java.lang.String name) {
        this(name, ServiceendpointJfrogPlatformV2Args.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ServiceendpointJfrogPlatformV2(java.lang.String name, ServiceendpointJfrogPlatformV2Args args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ServiceendpointJfrogPlatformV2(java.lang.String name, ServiceendpointJfrogPlatformV2Args args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/serviceendpointJfrogPlatformV2:ServiceendpointJfrogPlatformV2", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private ServiceendpointJfrogPlatformV2(java.lang.String name, Output<java.lang.String> id, @Nullable ServiceendpointJfrogPlatformV2State state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/serviceendpointJfrogPlatformV2:ServiceendpointJfrogPlatformV2", name, state, makeResourceOptions(options, id), false);
    }

    private static ServiceendpointJfrogPlatformV2Args makeArgs(ServiceendpointJfrogPlatformV2Args args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? ServiceendpointJfrogPlatformV2Args.Empty : args;
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
    public static ServiceendpointJfrogPlatformV2 get(java.lang.String name, Output<java.lang.String> id, @Nullable ServiceendpointJfrogPlatformV2State state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ServiceendpointJfrogPlatformV2(name, id, state, options);
    }
}
