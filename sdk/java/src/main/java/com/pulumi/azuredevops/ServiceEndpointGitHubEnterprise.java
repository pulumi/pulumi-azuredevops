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
 * ```java
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
 *         var exampleProject = new Project(&#34;exampleProject&#34;, ProjectArgs.builder()        
 *             .visibility(&#34;private&#34;)
 *             .versionControl(&#34;Git&#34;)
 *             .workItemTemplate(&#34;Agile&#34;)
 *             .description(&#34;Managed by Terraform&#34;)
 *             .build());
 * 
 *         var exampleServiceEndpointGitHubEnterprise = new ServiceEndpointGitHubEnterprise(&#34;exampleServiceEndpointGitHubEnterprise&#34;, ServiceEndpointGitHubEnterpriseArgs.builder()        
 *             .projectId(exampleProject.id())
 *             .serviceEndpointName(&#34;Example GitHub Enterprise&#34;)
 *             .url(&#34;https://github.contoso.com&#34;)
 *             .description(&#34;Managed by Terraform&#34;)
 *             .authPersonal(ServiceEndpointGitHubEnterpriseAuthPersonalArgs.builder()
 *                 .personalAccessToken(&#34;xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx&#34;)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ## Relevant Links
 * 
 * - [Azure DevOps Service REST API 7.0 - Service Endpoints](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-7.0)
 * 
 * ## Import
 * 
 * Azure DevOps Service Endpoint GitHub Enterprise Server can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**
 * 
 * ```sh
 *  $ pulumi import azuredevops:index/serviceEndpointGitHubEnterprise:ServiceEndpointGitHubEnterprise example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
 * ```
 * 
 */
@ResourceType(type="azuredevops:index/serviceEndpointGitHubEnterprise:ServiceEndpointGitHubEnterprise")
public class ServiceEndpointGitHubEnterprise extends com.pulumi.resources.CustomResource {
    @Export(name="authPersonal", refs={ServiceEndpointGitHubEnterpriseAuthPersonal.class}, tree="[0]")
    private Output<ServiceEndpointGitHubEnterpriseAuthPersonal> authPersonal;

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
    public ServiceEndpointGitHubEnterprise(String name) {
        this(name, ServiceEndpointGitHubEnterpriseArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ServiceEndpointGitHubEnterprise(String name, ServiceEndpointGitHubEnterpriseArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ServiceEndpointGitHubEnterprise(String name, ServiceEndpointGitHubEnterpriseArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/serviceEndpointGitHubEnterprise:ServiceEndpointGitHubEnterprise", name, args == null ? ServiceEndpointGitHubEnterpriseArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ServiceEndpointGitHubEnterprise(String name, Output<String> id, @Nullable ServiceEndpointGitHubEnterpriseState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/serviceEndpointGitHubEnterprise:ServiceEndpointGitHubEnterprise", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
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
    public static ServiceEndpointGitHubEnterprise get(String name, Output<String> id, @Nullable ServiceEndpointGitHubEnterpriseState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ServiceEndpointGitHubEnterprise(name, id, state, options);
    }
}
