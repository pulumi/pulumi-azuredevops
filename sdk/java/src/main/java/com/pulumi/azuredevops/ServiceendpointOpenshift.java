// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops;

import com.pulumi.azuredevops.ServiceendpointOpenshiftArgs;
import com.pulumi.azuredevops.Utilities;
import com.pulumi.azuredevops.inputs.ServiceendpointOpenshiftState;
import com.pulumi.azuredevops.outputs.ServiceendpointOpenshiftAuthBasic;
import com.pulumi.azuredevops.outputs.ServiceendpointOpenshiftAuthNone;
import com.pulumi.azuredevops.outputs.ServiceendpointOpenshiftAuthToken;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages an Openshift service endpoint within Azure DevOps organization. Using this service endpoint requires you to first install the [OpenShift Extension](https://marketplace.visualstudio.com/items?itemName=redhat.openshift-vsts).
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
 * import com.pulumi.azuredevops.ServiceendpointOpenshift;
 * import com.pulumi.azuredevops.ServiceendpointOpenshiftArgs;
 * import com.pulumi.azuredevops.inputs.ServiceendpointOpenshiftAuthBasicArgs;
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
 *         var exampleServiceendpointOpenshift = new ServiceendpointOpenshift("exampleServiceendpointOpenshift", ServiceendpointOpenshiftArgs.builder()
 *             .projectId(example.id())
 *             .serviceEndpointName("Example Openshift")
 *             .serverUrl("https://example.server")
 *             .certificateAuthorityFile("/opt/file")
 *             .acceptUntrustedCerts(true)
 *             .authBasic(ServiceendpointOpenshiftAuthBasicArgs.builder()
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
 * import com.pulumi.azuredevops.ServiceendpointOpenshift;
 * import com.pulumi.azuredevops.ServiceendpointOpenshiftArgs;
 * import com.pulumi.azuredevops.inputs.ServiceendpointOpenshiftAuthTokenArgs;
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
 *         var exampleServiceendpointOpenshift = new ServiceendpointOpenshift("exampleServiceendpointOpenshift", ServiceendpointOpenshiftArgs.builder()
 *             .projectId(example.id())
 *             .serviceEndpointName("Example Openshift")
 *             .serverUrl("https://example.server")
 *             .certificateAuthorityFile("/opt/file")
 *             .acceptUntrustedCerts(true)
 *             .authToken(ServiceendpointOpenshiftAuthTokenArgs.builder()
 *                 .token("username")
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
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
 * import com.pulumi.azuredevops.ServiceendpointOpenshift;
 * import com.pulumi.azuredevops.ServiceendpointOpenshiftArgs;
 * import com.pulumi.azuredevops.inputs.ServiceendpointOpenshiftAuthNoneArgs;
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
 *         var exampleServiceendpointOpenshift = new ServiceendpointOpenshift("exampleServiceendpointOpenshift", ServiceendpointOpenshiftArgs.builder()
 *             .projectId(example.id())
 *             .serviceEndpointName("Example Openshift")
 *             .serverUrl("https://example.server")
 *             .authNone(ServiceendpointOpenshiftAuthNoneArgs.builder()
 *                 .kubeConfig("config")
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
 * - [Azure DevOps Service REST API 7.1 - Service Endpoints](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-7.1)
 * 
 * ## Import
 * 
 * Azure DevOps Openshift Service Endpoint can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**
 * 
 * ```sh
 * $ pulumi import azuredevops:index/serviceendpointOpenshift:ServiceendpointOpenshift example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
 * ```
 * 
 */
@ResourceType(type="azuredevops:index/serviceendpointOpenshift:ServiceendpointOpenshift")
public class ServiceendpointOpenshift extends com.pulumi.resources.CustomResource {
    /**
     * Set this option to allow clients to accept a self-signed certificate. Available when using `auth_basic` or `auth_token` authorization.
     * 
     */
    @Export(name="acceptUntrustedCerts", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> acceptUntrustedCerts;

    /**
     * @return Set this option to allow clients to accept a self-signed certificate. Available when using `auth_basic` or `auth_token` authorization.
     * 
     */
    public Output<Optional<Boolean>> acceptUntrustedCerts() {
        return Codegen.optional(this.acceptUntrustedCerts);
    }
    /**
     * An `auth_basic` block as documented below.
     * 
     */
    @Export(name="authBasic", refs={ServiceendpointOpenshiftAuthBasic.class}, tree="[0]")
    private Output</* @Nullable */ ServiceendpointOpenshiftAuthBasic> authBasic;

    /**
     * @return An `auth_basic` block as documented below.
     * 
     */
    public Output<Optional<ServiceendpointOpenshiftAuthBasic>> authBasic() {
        return Codegen.optional(this.authBasic);
    }
    /**
     * An `auth_none` block as documented below.
     * 
     */
    @Export(name="authNone", refs={ServiceendpointOpenshiftAuthNone.class}, tree="[0]")
    private Output</* @Nullable */ ServiceendpointOpenshiftAuthNone> authNone;

    /**
     * @return An `auth_none` block as documented below.
     * 
     */
    public Output<Optional<ServiceendpointOpenshiftAuthNone>> authNone() {
        return Codegen.optional(this.authNone);
    }
    /**
     * An `auth_token` block as documented below.
     * 
     */
    @Export(name="authToken", refs={ServiceendpointOpenshiftAuthToken.class}, tree="[0]")
    private Output</* @Nullable */ ServiceendpointOpenshiftAuthToken> authToken;

    /**
     * @return An `auth_token` block as documented below.
     * 
     */
    public Output<Optional<ServiceendpointOpenshiftAuthToken>> authToken() {
        return Codegen.optional(this.authToken);
    }
    @Export(name="authorization", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> authorization;

    public Output<Map<String,String>> authorization() {
        return this.authorization;
    }
    /**
     * The path to a certificate authority file to correctly and securely authenticates with an OpenShift server that uses HTTPS. Available when using `auth_basic` or `auth_token` authorization.
     * 
     */
    @Export(name="certificateAuthorityFile", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> certificateAuthorityFile;

    /**
     * @return The path to a certificate authority file to correctly and securely authenticates with an OpenShift server that uses HTTPS. Available when using `auth_basic` or `auth_token` authorization.
     * 
     */
    public Output<Optional<String>> certificateAuthorityFile() {
        return Codegen.optional(this.certificateAuthorityFile);
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
     * The URL for the OpenShift cluster to connect to.
     * 
     */
    @Export(name="serverUrl", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> serverUrl;

    /**
     * @return The URL for the OpenShift cluster to connect to.
     * 
     */
    public Output<Optional<String>> serverUrl() {
        return Codegen.optional(this.serverUrl);
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
    public ServiceendpointOpenshift(java.lang.String name) {
        this(name, ServiceendpointOpenshiftArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ServiceendpointOpenshift(java.lang.String name, ServiceendpointOpenshiftArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ServiceendpointOpenshift(java.lang.String name, ServiceendpointOpenshiftArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/serviceendpointOpenshift:ServiceendpointOpenshift", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private ServiceendpointOpenshift(java.lang.String name, Output<java.lang.String> id, @Nullable ServiceendpointOpenshiftState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/serviceendpointOpenshift:ServiceendpointOpenshift", name, state, makeResourceOptions(options, id), false);
    }

    private static ServiceendpointOpenshiftArgs makeArgs(ServiceendpointOpenshiftArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? ServiceendpointOpenshiftArgs.Empty : args;
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
    public static ServiceendpointOpenshift get(java.lang.String name, Output<java.lang.String> id, @Nullable ServiceendpointOpenshiftState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ServiceendpointOpenshift(name, id, state, options);
    }
}
