// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops;

import com.pulumi.azuredevops.ServiceendpointJenkinsArgs;
import com.pulumi.azuredevops.Utilities;
import com.pulumi.azuredevops.inputs.ServiceendpointJenkinsState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a Jenkins service endpoint within Azure DevOps, which can be used as a resource in YAML pipelines to connect to a Jenkins instance.
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
 * import com.pulumi.azuredevops.ServiceendpointJenkins;
 * import com.pulumi.azuredevops.ServiceendpointJenkinsArgs;
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
 *         var exampleServiceendpointJenkins = new ServiceendpointJenkins("exampleServiceendpointJenkins", ServiceendpointJenkinsArgs.builder()
 *             .projectId(example.id())
 *             .serviceEndpointName("jenkins-example")
 *             .description("Service Endpoint for 'Jenkins' (Managed by Terraform)")
 *             .url("https://example.com")
 *             .acceptUntrustedCerts(false)
 *             .username("username")
 *             .password("password")
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
 * Azure DevOps Jenkins Service Endpoint can be imported using the `projectId/id` or `projectName/id`, e.g.
 * 
 * ```sh
 * $ pulumi import azuredevops:index/serviceendpointJenkins:ServiceendpointJenkins example projectName/00000000-0000-0000-0000-000000000000
 * ```
 * 
 */
@ResourceType(type="azuredevops:index/serviceendpointJenkins:ServiceendpointJenkins")
public class ServiceendpointJenkins extends com.pulumi.resources.CustomResource {
    /**
     * Allows the Jenkins clients to accept self-signed SSL server certificates. Defaults to `false.`
     * 
     */
    @Export(name="acceptUntrustedCerts", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> acceptUntrustedCerts;

    /**
     * @return Allows the Jenkins clients to accept self-signed SSL server certificates. Defaults to `false.`
     * 
     */
    public Output<Optional<Boolean>> acceptUntrustedCerts() {
        return Codegen.optional(this.acceptUntrustedCerts);
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
     * The Service Endpoint password to authenticate at the Jenkins Instance.
     * 
     */
    @Export(name="password", refs={String.class}, tree="[0]")
    private Output<String> password;

    /**
     * @return The Service Endpoint password to authenticate at the Jenkins Instance.
     * 
     */
    public Output<String> password() {
        return this.password;
    }
    /**
     * The ID of the project. Changing this forces a new Service Connection Jenkins to be created.
     * 
     */
    @Export(name="projectId", refs={String.class}, tree="[0]")
    private Output<String> projectId;

    /**
     * @return The ID of the project. Changing this forces a new Service Connection Jenkins to be created.
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }
    /**
     * The name of the service endpoint. Changing this forces a new Service Connection Jenkins to be created.
     * 
     */
    @Export(name="serviceEndpointName", refs={String.class}, tree="[0]")
    private Output<String> serviceEndpointName;

    /**
     * @return The name of the service endpoint. Changing this forces a new Service Connection Jenkins to be created.
     * 
     */
    public Output<String> serviceEndpointName() {
        return this.serviceEndpointName;
    }
    /**
     * The Service Endpoint url.
     * 
     */
    @Export(name="url", refs={String.class}, tree="[0]")
    private Output<String> url;

    /**
     * @return The Service Endpoint url.
     * 
     */
    public Output<String> url() {
        return this.url;
    }
    /**
     * The Service Endpoint username to authenticate at the Jenkins Instance.
     * 
     */
    @Export(name="username", refs={String.class}, tree="[0]")
    private Output<String> username;

    /**
     * @return The Service Endpoint username to authenticate at the Jenkins Instance.
     * 
     */
    public Output<String> username() {
        return this.username;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ServiceendpointJenkins(java.lang.String name) {
        this(name, ServiceendpointJenkinsArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ServiceendpointJenkins(java.lang.String name, ServiceendpointJenkinsArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ServiceendpointJenkins(java.lang.String name, ServiceendpointJenkinsArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/serviceendpointJenkins:ServiceendpointJenkins", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private ServiceendpointJenkins(java.lang.String name, Output<java.lang.String> id, @Nullable ServiceendpointJenkinsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/serviceendpointJenkins:ServiceendpointJenkins", name, state, makeResourceOptions(options, id), false);
    }

    private static ServiceendpointJenkinsArgs makeArgs(ServiceendpointJenkinsArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? ServiceendpointJenkinsArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "password"
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
    public static ServiceendpointJenkins get(java.lang.String name, Output<java.lang.String> id, @Nullable ServiceendpointJenkinsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ServiceendpointJenkins(name, id, state, options);
    }
}
