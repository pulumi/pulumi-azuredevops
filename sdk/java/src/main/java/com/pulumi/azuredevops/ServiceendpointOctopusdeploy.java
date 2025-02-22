// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops;

import com.pulumi.azuredevops.ServiceendpointOctopusdeployArgs;
import com.pulumi.azuredevops.Utilities;
import com.pulumi.azuredevops.inputs.ServiceendpointOctopusdeployState;
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
 * Manages an Octopus Deploy service endpoint within Azure DevOps. Using this service endpoint requires you to install [Octopus Deploy](https://marketplace.visualstudio.com/items?itemName=octopusdeploy.octopus-deploy-build-release-tasks).
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
 * import com.pulumi.azuredevops.ServiceendpointOctopusdeploy;
 * import com.pulumi.azuredevops.ServiceendpointOctopusdeployArgs;
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
 *         var exampleServiceendpointOctopusdeploy = new ServiceendpointOctopusdeploy("exampleServiceendpointOctopusdeploy", ServiceendpointOctopusdeployArgs.builder()
 *             .projectId(example.id())
 *             .url("https://octopus.com")
 *             .apiKey("000000000000000000000000000000000000")
 *             .serviceEndpointName("Example Octopus Deploy")
 *             .description("Managed by Pulumi")
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
 * - [Azure DevOps Service REST API 7.0 - Agent Pools](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-7.0)
 * 
 * ## Import
 * 
 * Azure DevOps Octopus Deploy Service Endpoint can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**
 * 
 * ```sh
 * $ pulumi import azuredevops:index/serviceendpointOctopusdeploy:ServiceendpointOctopusdeploy example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
 * ```
 * 
 */
@ResourceType(type="azuredevops:index/serviceendpointOctopusdeploy:ServiceendpointOctopusdeploy")
public class ServiceendpointOctopusdeploy extends com.pulumi.resources.CustomResource {
    /**
     * API key to connect to Octopus Deploy.
     * 
     */
    @Export(name="apiKey", refs={String.class}, tree="[0]")
    private Output<String> apiKey;

    /**
     * @return API key to connect to Octopus Deploy.
     * 
     */
    public Output<String> apiKey() {
        return this.apiKey;
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
     * Whether to ignore SSL errors when connecting to the Octopus server from the agent. Default to `false`.
     * 
     */
    @Export(name="ignoreSslError", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> ignoreSslError;

    /**
     * @return Whether to ignore SSL errors when connecting to the Octopus server from the agent. Default to `false`.
     * 
     */
    public Output<Optional<Boolean>> ignoreSslError() {
        return Codegen.optional(this.ignoreSslError);
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
     * Octopus Server url.
     * 
     */
    @Export(name="url", refs={String.class}, tree="[0]")
    private Output<String> url;

    /**
     * @return Octopus Server url.
     * 
     */
    public Output<String> url() {
        return this.url;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ServiceendpointOctopusdeploy(java.lang.String name) {
        this(name, ServiceendpointOctopusdeployArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ServiceendpointOctopusdeploy(java.lang.String name, ServiceendpointOctopusdeployArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ServiceendpointOctopusdeploy(java.lang.String name, ServiceendpointOctopusdeployArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/serviceendpointOctopusdeploy:ServiceendpointOctopusdeploy", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private ServiceendpointOctopusdeploy(java.lang.String name, Output<java.lang.String> id, @Nullable ServiceendpointOctopusdeployState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/serviceendpointOctopusdeploy:ServiceendpointOctopusdeploy", name, state, makeResourceOptions(options, id), false);
    }

    private static ServiceendpointOctopusdeployArgs makeArgs(ServiceendpointOctopusdeployArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? ServiceendpointOctopusdeployArgs.Empty : args;
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
    public static ServiceendpointOctopusdeploy get(java.lang.String name, Output<java.lang.String> id, @Nullable ServiceendpointOctopusdeployState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ServiceendpointOctopusdeploy(name, id, state, options);
    }
}
