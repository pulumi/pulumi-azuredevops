// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops;

import com.pulumi.azuredevops.ServiceendpointGcpTerraformArgs;
import com.pulumi.azuredevops.Utilities;
import com.pulumi.azuredevops.inputs.ServiceendpointGcpTerraformState;
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
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.azuredevops.Project;
 * import com.pulumi.azuredevops.ProjectArgs;
 * import com.pulumi.azuredevops.ServiceendpointGcpTerraform;
 * import com.pulumi.azuredevops.ServiceendpointGcpTerraformArgs;
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
 *         var example = new Project(&#34;example&#34;, ProjectArgs.builder()        
 *             .name(&#34;Example Project&#34;)
 *             .visibility(&#34;private&#34;)
 *             .versionControl(&#34;Git&#34;)
 *             .workItemTemplate(&#34;Agile&#34;)
 *             .description(&#34;Managed by Terraform&#34;)
 *             .build());
 * 
 *         var exampleServiceendpointGcpTerraform = new ServiceendpointGcpTerraform(&#34;exampleServiceendpointGcpTerraform&#34;, ServiceendpointGcpTerraformArgs.builder()        
 *             .projectId(example.id())
 *             .tokenUri(&#34;https://oauth2.example.com/token&#34;)
 *             .clientEmail(&#34;gcp-sa-example@example.iam.gserviceaccount.com&#34;)
 *             .privateKey(&#34;0000000000000000000000000000000000000&#34;)
 *             .serviceEndpointName(&#34;Example GCP Terraform extension&#34;)
 *             .gcpProjectId(&#34;Example GCP Project&#34;)
 *             .description(&#34;Managed by Terraform&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Relevant Links
 * 
 * - [Azure DevOps Service REST API 7.1 - Service Endpoints](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-7.1)
 * 
 * ## Import
 * 
 * Azure DevOps Service Endpoint GCP can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**
 * 
 * ```sh
 * $ pulumi import azuredevops:index/serviceendpointGcpTerraform:ServiceendpointGcpTerraform azuredevops_serviceendpoint_gcp_terraform.example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
 * ```
 * 
 */
@ResourceType(type="azuredevops:index/serviceendpointGcpTerraform:ServiceendpointGcpTerraform")
public class ServiceendpointGcpTerraform extends com.pulumi.resources.CustomResource {
    @Export(name="authorization", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> authorization;

    public Output<Map<String,String>> authorization() {
        return this.authorization;
    }
    /**
     * The client email field in the JSON key file for creating the JSON Web Token.
     * 
     */
    @Export(name="clientEmail", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> clientEmail;

    /**
     * @return The client email field in the JSON key file for creating the JSON Web Token.
     * 
     */
    public Output<Optional<String>> clientEmail() {
        return Codegen.optional(this.clientEmail);
    }
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * GCP project associated with the Service Connection.
     * 
     */
    @Export(name="gcpProjectId", refs={String.class}, tree="[0]")
    private Output<String> gcpProjectId;

    /**
     * @return GCP project associated with the Service Connection.
     * 
     */
    public Output<String> gcpProjectId() {
        return this.gcpProjectId;
    }
    /**
     * The client email field in the JSON key file for creating the JSON Web Token.
     * 
     */
    @Export(name="privateKey", refs={String.class}, tree="[0]")
    private Output<String> privateKey;

    /**
     * @return The client email field in the JSON key file for creating the JSON Web Token.
     * 
     */
    public Output<String> privateKey() {
        return this.privateKey;
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
     * Scope to be provided.
     * 
     */
    @Export(name="scope", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> scope;

    /**
     * @return Scope to be provided.
     * 
     */
    public Output<Optional<String>> scope() {
        return Codegen.optional(this.scope);
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
     * The token uri field in the JSON key file for creating the JSON Web Token.
     * 
     */
    @Export(name="tokenUri", refs={String.class}, tree="[0]")
    private Output<String> tokenUri;

    /**
     * @return The token uri field in the JSON key file for creating the JSON Web Token.
     * 
     */
    public Output<String> tokenUri() {
        return this.tokenUri;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ServiceendpointGcpTerraform(String name) {
        this(name, ServiceendpointGcpTerraformArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ServiceendpointGcpTerraform(String name, ServiceendpointGcpTerraformArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ServiceendpointGcpTerraform(String name, ServiceendpointGcpTerraformArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/serviceendpointGcpTerraform:ServiceendpointGcpTerraform", name, args == null ? ServiceendpointGcpTerraformArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ServiceendpointGcpTerraform(String name, Output<String> id, @Nullable ServiceendpointGcpTerraformState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/serviceendpointGcpTerraform:ServiceendpointGcpTerraform", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "privateKey"
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
    public static ServiceendpointGcpTerraform get(String name, Output<String> id, @Nullable ServiceendpointGcpTerraformState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ServiceendpointGcpTerraform(name, id, state, options);
    }
}
