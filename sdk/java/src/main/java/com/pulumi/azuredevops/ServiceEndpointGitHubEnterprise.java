// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops;

import com.pulumi.azuredevops.ServiceEndpointGitHubEnterpriseArgs;
import com.pulumi.azuredevops.Utilities;
import com.pulumi.azuredevops.inputs.ServiceEndpointGitHubEnterpriseState;
import com.pulumi.azuredevops.outputs.ServiceEndpointGitHubEnterpriseAuthPersonal;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a GitHub Enterprise Server service endpoint within Azure DevOps.
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
 * import com.pulumi.azuredevops.ServiceEndpointGitHubEnterprise;
 * import com.pulumi.azuredevops.ServiceEndpointGitHubEnterpriseArgs;
 * import com.pulumi.azuredevops.inputs.ServiceEndpointGitHubEnterpriseAuthPersonalArgs;
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
 *         var exampleServiceEndpointGitHubEnterprise = new ServiceEndpointGitHubEnterprise("exampleServiceEndpointGitHubEnterprise", ServiceEndpointGitHubEnterpriseArgs.builder()
 *             .projectId(example.id())
 *             .serviceEndpointName("Example GitHub Enterprise")
 *             .url("https://github.contoso.com")
 *             .description("Managed by Pulumi")
 *             .authPersonal(ServiceEndpointGitHubEnterpriseAuthPersonalArgs.builder()
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
 * Azure DevOps GitHub Enterprise Server Service Endpoint can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**
 * 
 * ```sh
 * $ pulumi import azuredevops:index/serviceEndpointGitHubEnterprise:ServiceEndpointGitHubEnterprise example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
 * ```
 * 
 */
@ResourceType(type="azuredevops:index/serviceEndpointGitHubEnterprise:ServiceEndpointGitHubEnterprise")
public class ServiceEndpointGitHubEnterprise extends com.pulumi.resources.CustomResource {
    /**
     * An `auth_personal` block as documented below. Allows connecting using a personal access token.
     * 
     */
    @Export(name="authPersonal", refs={ServiceEndpointGitHubEnterpriseAuthPersonal.class}, tree="[0]")
    private Output<ServiceEndpointGitHubEnterpriseAuthPersonal> authPersonal;

    /**
     * @return An `auth_personal` block as documented below. Allows connecting using a personal access token.
     * 
     */
    public Output<ServiceEndpointGitHubEnterpriseAuthPersonal> authPersonal() {
        return this.authPersonal;
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
     * GitHub Enterprise Server Url.
     * 
     */
    @Export(name="url", refs={String.class}, tree="[0]")
    private Output<String> url;

    /**
     * @return GitHub Enterprise Server Url.
     * 
     */
    public Output<String> url() {
        return this.url;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ServiceEndpointGitHubEnterprise(java.lang.String name) {
        this(name, ServiceEndpointGitHubEnterpriseArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ServiceEndpointGitHubEnterprise(java.lang.String name, ServiceEndpointGitHubEnterpriseArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ServiceEndpointGitHubEnterprise(java.lang.String name, ServiceEndpointGitHubEnterpriseArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/serviceEndpointGitHubEnterprise:ServiceEndpointGitHubEnterprise", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private ServiceEndpointGitHubEnterprise(java.lang.String name, Output<java.lang.String> id, @Nullable ServiceEndpointGitHubEnterpriseState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/serviceEndpointGitHubEnterprise:ServiceEndpointGitHubEnterprise", name, state, makeResourceOptions(options, id), false);
    }

    private static ServiceEndpointGitHubEnterpriseArgs makeArgs(ServiceEndpointGitHubEnterpriseArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? ServiceEndpointGitHubEnterpriseArgs.Empty : args;
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
    public static ServiceEndpointGitHubEnterprise get(java.lang.String name, Output<java.lang.String> id, @Nullable ServiceEndpointGitHubEnterpriseState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ServiceEndpointGitHubEnterprise(name, id, state, options);
    }
}
