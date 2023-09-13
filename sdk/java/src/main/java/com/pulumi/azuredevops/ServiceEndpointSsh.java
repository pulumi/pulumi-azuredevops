// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops;

import com.pulumi.azuredevops.ServiceEndpointSshArgs;
import com.pulumi.azuredevops.Utilities;
import com.pulumi.azuredevops.inputs.ServiceEndpointSshState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a SSH service endpoint within Azure DevOps.
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
 * import com.pulumi.azuredevops.ServiceEndpointSsh;
 * import com.pulumi.azuredevops.ServiceEndpointSshArgs;
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
 *         var exampleServiceEndpointSsh = new ServiceEndpointSsh(&#34;exampleServiceEndpointSsh&#34;, ServiceEndpointSshArgs.builder()        
 *             .projectId(exampleProject.id())
 *             .serviceEndpointName(&#34;Example SSH&#34;)
 *             .host(&#34;1.2.3.4&#34;)
 *             .username(&#34;username&#34;)
 *             .description(&#34;Managed by Terraform&#34;)
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
 * Azure DevOps Service Endpoint SSH can be imported using **projectID/serviceEndpointID** or ** projectName/serviceEndpointID**
 * 
 * ```sh
 *  $ pulumi import azuredevops:index/serviceEndpointSsh:ServiceEndpointSsh example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
 * ```
 * 
 */
@ResourceType(type="azuredevops:index/serviceEndpointSsh:ServiceEndpointSsh")
public class ServiceEndpointSsh extends com.pulumi.resources.CustomResource {
    @Export(name="authorization", type=Map.class, parameters={String.class, String.class})
    private Output<Map<String,String>> authorization;

    public Output<Map<String,String>> authorization() {
        return this.authorization;
    }
    @Export(name="description", type=String.class, parameters={})
    private Output</* @Nullable */ String> description;

    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The Host name or IP address of the remote machine.
     * 
     */
    @Export(name="host", type=String.class, parameters={})
    private Output<String> host;

    /**
     * @return The Host name or IP address of the remote machine.
     * 
     */
    public Output<String> host() {
        return this.host;
    }
    /**
     * Password for connecting to the endpoint.
     * 
     */
    @Export(name="password", type=String.class, parameters={})
    private Output</* @Nullable */ String> password;

    /**
     * @return Password for connecting to the endpoint.
     * 
     */
    public Output<Optional<String>> password() {
        return Codegen.optional(this.password);
    }
    /**
     * Port number on the remote machine to use for connecting. Defaults to `22`.
     * 
     */
    @Export(name="port", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> port;

    /**
     * @return Port number on the remote machine to use for connecting. Defaults to `22`.
     * 
     */
    public Output<Optional<Integer>> port() {
        return Codegen.optional(this.port);
    }
    /**
     * Private Key for connecting to the endpoint.
     * 
     */
    @Export(name="privateKey", type=String.class, parameters={})
    private Output</* @Nullable */ String> privateKey;

    /**
     * @return Private Key for connecting to the endpoint.
     * 
     */
    public Output<Optional<String>> privateKey() {
        return Codegen.optional(this.privateKey);
    }
    /**
     * The ID of the project.
     * 
     */
    @Export(name="projectId", type=String.class, parameters={})
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
    @Export(name="serviceEndpointName", type=String.class, parameters={})
    private Output<String> serviceEndpointName;

    /**
     * @return The Service Endpoint name.
     * 
     */
    public Output<String> serviceEndpointName() {
        return this.serviceEndpointName;
    }
    /**
     * Username for connecting to the endpoint.
     * 
     */
    @Export(name="username", type=String.class, parameters={})
    private Output<String> username;

    /**
     * @return Username for connecting to the endpoint.
     * 
     */
    public Output<String> username() {
        return this.username;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ServiceEndpointSsh(String name) {
        this(name, ServiceEndpointSshArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ServiceEndpointSsh(String name, ServiceEndpointSshArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ServiceEndpointSsh(String name, ServiceEndpointSshArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/serviceEndpointSsh:ServiceEndpointSsh", name, args == null ? ServiceEndpointSshArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ServiceEndpointSsh(String name, Output<String> id, @Nullable ServiceEndpointSshState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/serviceEndpointSsh:ServiceEndpointSsh", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "password",
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
    public static ServiceEndpointSsh get(String name, Output<String> id, @Nullable ServiceEndpointSshState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ServiceEndpointSsh(name, id, state, options);
    }
}
