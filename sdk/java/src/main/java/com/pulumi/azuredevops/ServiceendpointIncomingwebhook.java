// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops;

import com.pulumi.azuredevops.ServiceendpointIncomingwebhookArgs;
import com.pulumi.azuredevops.Utilities;
import com.pulumi.azuredevops.inputs.ServiceendpointIncomingwebhookState;
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
 * Manages an Incoming WebHook service endpoint within Azure DevOps, which can be used as a resource in YAML pipelines to subscribe to a webhook event.
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
 * import com.pulumi.azuredevops.ServiceendpointIncomingwebhook;
 * import com.pulumi.azuredevops.ServiceendpointIncomingwebhookArgs;
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
 *         var exampleServiceendpointIncomingwebhook = new ServiceendpointIncomingwebhook("exampleServiceendpointIncomingwebhook", ServiceendpointIncomingwebhookArgs.builder()
 *             .projectId(example.id())
 *             .webhookName("example_webhook")
 *             .secret("secret")
 *             .httpHeader("X-Hub-Signature")
 *             .serviceEndpointName("Example IncomingWebhook")
 *             .description("Managed by Pulumi")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Azure DevOps Incoming WebHook Service Endpoint can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**
 * 
 * ```sh
 * $ pulumi import azuredevops:index/serviceendpointIncomingwebhook:ServiceendpointIncomingwebhook example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
 * ```
 * 
 */
@ResourceType(type="azuredevops:index/serviceendpointIncomingwebhook:ServiceendpointIncomingwebhook")
public class ServiceendpointIncomingwebhook extends com.pulumi.resources.CustomResource {
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
     * Http header name on which checksum will be sent.
     * 
     */
    @Export(name="httpHeader", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> httpHeader;

    /**
     * @return Http header name on which checksum will be sent.
     * 
     */
    public Output<Optional<String>> httpHeader() {
        return Codegen.optional(this.httpHeader);
    }
    /**
     * The ID of the project. Changing this forces a new Service Connection Incoming WebHook to be created.
     * 
     */
    @Export(name="projectId", refs={String.class}, tree="[0]")
    private Output<String> projectId;

    /**
     * @return The ID of the project. Changing this forces a new Service Connection Incoming WebHook to be created.
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }
    /**
     * Secret for the WebHook. WebHook service will use this secret to calculate the payload checksum.
     * 
     */
    @Export(name="secret", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> secret;

    /**
     * @return Secret for the WebHook. WebHook service will use this secret to calculate the payload checksum.
     * 
     */
    public Output<Optional<String>> secret() {
        return Codegen.optional(this.secret);
    }
    /**
     * The name of the service endpoint. Changing this forces a new Service Connection Incoming WebHook to be created.
     * 
     */
    @Export(name="serviceEndpointName", refs={String.class}, tree="[0]")
    private Output<String> serviceEndpointName;

    /**
     * @return The name of the service endpoint. Changing this forces a new Service Connection Incoming WebHook to be created.
     * 
     */
    public Output<String> serviceEndpointName() {
        return this.serviceEndpointName;
    }
    /**
     * The name of the WebHook.
     * 
     */
    @Export(name="webhookName", refs={String.class}, tree="[0]")
    private Output<String> webhookName;

    /**
     * @return The name of the WebHook.
     * 
     */
    public Output<String> webhookName() {
        return this.webhookName;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ServiceendpointIncomingwebhook(java.lang.String name) {
        this(name, ServiceendpointIncomingwebhookArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ServiceendpointIncomingwebhook(java.lang.String name, ServiceendpointIncomingwebhookArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ServiceendpointIncomingwebhook(java.lang.String name, ServiceendpointIncomingwebhookArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/serviceendpointIncomingwebhook:ServiceendpointIncomingwebhook", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private ServiceendpointIncomingwebhook(java.lang.String name, Output<java.lang.String> id, @Nullable ServiceendpointIncomingwebhookState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/serviceendpointIncomingwebhook:ServiceendpointIncomingwebhook", name, state, makeResourceOptions(options, id), false);
    }

    private static ServiceendpointIncomingwebhookArgs makeArgs(ServiceendpointIncomingwebhookArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? ServiceendpointIncomingwebhookArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "secret"
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
    public static ServiceendpointIncomingwebhook get(java.lang.String name, Output<java.lang.String> id, @Nullable ServiceendpointIncomingwebhookState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ServiceendpointIncomingwebhook(name, id, state, options);
    }
}
