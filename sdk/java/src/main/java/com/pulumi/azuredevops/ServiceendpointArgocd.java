// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops;

import com.pulumi.azuredevops.ServiceendpointArgocdArgs;
import com.pulumi.azuredevops.Utilities;
import com.pulumi.azuredevops.inputs.ServiceendpointArgocdState;
import com.pulumi.azuredevops.outputs.ServiceendpointArgocdAuthenticationBasic;
import com.pulumi.azuredevops.outputs.ServiceendpointArgocdAuthenticationToken;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a ArgoCD service endpoint within Azure DevOps. Using this service endpoint requires you to first install [Argo CD Extension](https://marketplace.visualstudio.com/items?itemName=scb-tomasmortensen.vsix-argocd).
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
 * import com.pulumi.azuredevops.ServiceendpointArgocd;
 * import com.pulumi.azuredevops.ServiceendpointArgocdArgs;
 * import com.pulumi.azuredevops.inputs.ServiceendpointArgocdAuthenticationTokenArgs;
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
 *         var exampleServiceendpointArgocd = new ServiceendpointArgocd("exampleServiceendpointArgocd", ServiceendpointArgocdArgs.builder()
 *             .projectId(example.id())
 *             .serviceEndpointName("Example ArgoCD")
 *             .description("Managed by Pulumi")
 *             .url("https://argocd.my.com")
 *             .authenticationToken(ServiceendpointArgocdAuthenticationTokenArgs.builder()
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
 * import com.pulumi.azuredevops.ServiceendpointArgocd;
 * import com.pulumi.azuredevops.ServiceendpointArgocdArgs;
 * import com.pulumi.azuredevops.inputs.ServiceendpointArgocdAuthenticationBasicArgs;
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
 *         var exampleServiceendpointArgocd = new ServiceendpointArgocd("exampleServiceendpointArgocd", ServiceendpointArgocdArgs.builder()
 *             .projectId(example.id())
 *             .serviceEndpointName("Example ArgoCD")
 *             .description("Managed by Pulumi")
 *             .url("https://argocd.my.com")
 *             .authenticationBasic(ServiceendpointArgocdAuthenticationBasicArgs.builder()
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
 * ## Relevant Links
 * 
 * - [Azure DevOps Service Connections](https://docs.microsoft.com/en-us/azure/devops/pipelines/library/service-endpoints?view=azure-devops&amp;tabs=yaml)
 * - [ArgoCD Project/User Token](https://argo-cd.readthedocs.io/en/stable/user-guide/commands/argocd_account_generate-token/)
 * - [Argo CD Extension](https://marketplace.visualstudio.com/items?itemName=scb-tomasmortensen.vsix-argocd)
 * 
 * ## Import
 * 
 * Azure DevOps Service Endpoint ArgoCD can be imported using the **projectID/serviceEndpointID**, e.g.
 * 
 * ```sh
 * $ pulumi import azuredevops:index/serviceendpointArgocd:ServiceendpointArgocd example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
 * ```
 * 
 */
@ResourceType(type="azuredevops:index/serviceendpointArgocd:ServiceendpointArgocd")
public class ServiceendpointArgocd extends com.pulumi.resources.CustomResource {
    /**
     * An `authentication_basic` block for the ArgoCD as documented below.
     * 
     * &gt; **NOTE:** `authentication_basic` and `authentication_token` conflict with each other, only one is required.
     * 
     */
    @Export(name="authenticationBasic", refs={ServiceendpointArgocdAuthenticationBasic.class}, tree="[0]")
    private Output</* @Nullable */ ServiceendpointArgocdAuthenticationBasic> authenticationBasic;

    /**
     * @return An `authentication_basic` block for the ArgoCD as documented below.
     * 
     * &gt; **NOTE:** `authentication_basic` and `authentication_token` conflict with each other, only one is required.
     * 
     */
    public Output<Optional<ServiceendpointArgocdAuthenticationBasic>> authenticationBasic() {
        return Codegen.optional(this.authenticationBasic);
    }
    /**
     * An `authentication_token` block for the ArgoCD as documented below.
     * 
     */
    @Export(name="authenticationToken", refs={ServiceendpointArgocdAuthenticationToken.class}, tree="[0]")
    private Output</* @Nullable */ ServiceendpointArgocdAuthenticationToken> authenticationToken;

    /**
     * @return An `authentication_token` block for the ArgoCD as documented below.
     * 
     */
    public Output<Optional<ServiceendpointArgocdAuthenticationToken>> authenticationToken() {
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
     * URL of the ArgoCD server to connect with.
     * 
     */
    @Export(name="url", refs={String.class}, tree="[0]")
    private Output<String> url;

    /**
     * @return URL of the ArgoCD server to connect with.
     * 
     */
    public Output<String> url() {
        return this.url;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ServiceendpointArgocd(java.lang.String name) {
        this(name, ServiceendpointArgocdArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ServiceendpointArgocd(java.lang.String name, ServiceendpointArgocdArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ServiceendpointArgocd(java.lang.String name, ServiceendpointArgocdArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/serviceendpointArgocd:ServiceendpointArgocd", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private ServiceendpointArgocd(java.lang.String name, Output<java.lang.String> id, @Nullable ServiceendpointArgocdState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/serviceendpointArgocd:ServiceendpointArgocd", name, state, makeResourceOptions(options, id), false);
    }

    private static ServiceendpointArgocdArgs makeArgs(ServiceendpointArgocdArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? ServiceendpointArgocdArgs.Empty : args;
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
    public static ServiceendpointArgocd get(java.lang.String name, Output<java.lang.String> id, @Nullable ServiceendpointArgocdState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ServiceendpointArgocd(name, id, state, options);
    }
}
