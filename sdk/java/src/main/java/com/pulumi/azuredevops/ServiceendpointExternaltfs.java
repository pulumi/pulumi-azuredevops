// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops;

import com.pulumi.azuredevops.ServiceendpointExternaltfsArgs;
import com.pulumi.azuredevops.Utilities;
import com.pulumi.azuredevops.inputs.ServiceendpointExternaltfsState;
import com.pulumi.azuredevops.outputs.ServiceendpointExternaltfsAuthPersonal;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages an Azure Repository/Team Foundation Server service endpoint within Azure DevOps.
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
 * import com.pulumi.azuredevops.ServiceendpointExternaltfs;
 * import com.pulumi.azuredevops.ServiceendpointExternaltfsArgs;
 * import com.pulumi.azuredevops.inputs.ServiceendpointExternaltfsAuthPersonalArgs;
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
 *         var exampleServiceendpointExternaltfs = new ServiceendpointExternaltfs("exampleServiceendpointExternaltfs", ServiceendpointExternaltfsArgs.builder()
 *             .projectId(example.id())
 *             .serviceEndpointName("Example External TFS Name")
 *             .connectionUrl("https://dev.azure.com/myorganization")
 *             .description("Managed by Pulumi")
 *             .authPersonal(ServiceendpointExternaltfsAuthPersonalArgs.builder()
 *                 .personalAccessToken("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
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
 * - [Azure DevOps Service REST API 7.0 - Service Endpoints](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-7.0)
 * 
 * ## Import
 * 
 * Azure DevOps Azure Repository/Team Foundation Server Service Endpoint can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**
 * 
 * ```sh
 * $ pulumi import azuredevops:index/serviceendpointExternaltfs:ServiceendpointExternaltfs example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
 * ```
 * 
 */
@ResourceType(type="azuredevops:index/serviceendpointExternaltfs:ServiceendpointExternaltfs")
public class ServiceendpointExternaltfs extends com.pulumi.resources.CustomResource {
    /**
     * An `auth_personal` block as documented below. Allows connecting using a personal access token.
     * 
     */
    @Export(name="authPersonal", refs={ServiceendpointExternaltfsAuthPersonal.class}, tree="[0]")
    private Output<ServiceendpointExternaltfsAuthPersonal> authPersonal;

    /**
     * @return An `auth_personal` block as documented below. Allows connecting using a personal access token.
     * 
     */
    public Output<ServiceendpointExternaltfsAuthPersonal> authPersonal() {
        return this.authPersonal;
    }
    @Export(name="authorization", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> authorization;

    public Output<Map<String,String>> authorization() {
        return this.authorization;
    }
    /**
     * Azure DevOps Organization or TFS Project Collection Url.
     * 
     */
    @Export(name="connectionUrl", refs={String.class}, tree="[0]")
    private Output<String> connectionUrl;

    /**
     * @return Azure DevOps Organization or TFS Project Collection Url.
     * 
     */
    public Output<String> connectionUrl() {
        return this.connectionUrl;
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
    public ServiceendpointExternaltfs(java.lang.String name) {
        this(name, ServiceendpointExternaltfsArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ServiceendpointExternaltfs(java.lang.String name, ServiceendpointExternaltfsArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ServiceendpointExternaltfs(java.lang.String name, ServiceendpointExternaltfsArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/serviceendpointExternaltfs:ServiceendpointExternaltfs", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private ServiceendpointExternaltfs(java.lang.String name, Output<java.lang.String> id, @Nullable ServiceendpointExternaltfsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/serviceendpointExternaltfs:ServiceendpointExternaltfs", name, state, makeResourceOptions(options, id), false);
    }

    private static ServiceendpointExternaltfsArgs makeArgs(ServiceendpointExternaltfsArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? ServiceendpointExternaltfsArgs.Empty : args;
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
    public static ServiceendpointExternaltfs get(java.lang.String name, Output<java.lang.String> id, @Nullable ServiceendpointExternaltfsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ServiceendpointExternaltfs(name, id, state, options);
    }
}
