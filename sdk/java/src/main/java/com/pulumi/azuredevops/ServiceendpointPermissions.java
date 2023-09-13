// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops;

import com.pulumi.azuredevops.ServiceendpointPermissionsArgs;
import com.pulumi.azuredevops.Utilities;
import com.pulumi.azuredevops.inputs.ServiceendpointPermissionsState;
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
 * Manages permissions for a Service Endpoint
 * 
 * &gt; **Note** Permissions can be assigned to group principals and not to single user principals.
 * 
 * ## Permission levels
 * 
 * Permission for Service Endpoints within Azure DevOps can be applied on two different levels.
 * Those levels are reflected by specifying (or omitting) values for the arguments `project_id` and `serviceendpoint_id`.
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
 * import com.pulumi.azuredevops.AzuredevopsFunctions;
 * import com.pulumi.azuredevops.inputs.GetGroupArgs;
 * import com.pulumi.azuredevops.ServiceendpointPermissions;
 * import com.pulumi.azuredevops.ServiceendpointPermissionsArgs;
 * import com.pulumi.azuredevops.ServiceEndpointDockerRegistry;
 * import com.pulumi.azuredevops.ServiceEndpointDockerRegistryArgs;
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
 *             .workItemTemplate(&#34;Agile&#34;)
 *             .versionControl(&#34;Git&#34;)
 *             .visibility(&#34;private&#34;)
 *             .description(&#34;Managed by Terraform&#34;)
 *             .build());
 * 
 *         final var example-readers = AzuredevopsFunctions.getGroup(GetGroupArgs.builder()
 *             .projectId(exampleProject.id())
 *             .name(&#34;Readers&#34;)
 *             .build());
 * 
 *         var example_root_permissions = new ServiceendpointPermissions(&#34;example-root-permissions&#34;, ServiceendpointPermissionsArgs.builder()        
 *             .projectId(exampleProject.id())
 *             .principal(example_readers.applyValue(example_readers -&gt; example_readers.id()))
 *             .permissions(Map.ofEntries(
 *                 Map.entry(&#34;Use&#34;, &#34;allow&#34;),
 *                 Map.entry(&#34;Administer&#34;, &#34;allow&#34;),
 *                 Map.entry(&#34;Create&#34;, &#34;allow&#34;),
 *                 Map.entry(&#34;ViewAuthorization&#34;, &#34;allow&#34;),
 *                 Map.entry(&#34;ViewEndpoint&#34;, &#34;allow&#34;)
 *             ))
 *             .build());
 * 
 *         var exampleServiceEndpointDockerRegistry = new ServiceEndpointDockerRegistry(&#34;exampleServiceEndpointDockerRegistry&#34;, ServiceEndpointDockerRegistryArgs.builder()        
 *             .projectId(exampleProject.id())
 *             .serviceEndpointName(&#34;Example Docker Hub&#34;)
 *             .dockerUsername(&#34;username&#34;)
 *             .dockerEmail(&#34;email@example.com&#34;)
 *             .dockerPassword(&#34;password&#34;)
 *             .registryType(&#34;DockerHub&#34;)
 *             .build());
 * 
 *         var example_permissions = new ServiceendpointPermissions(&#34;example-permissions&#34;, ServiceendpointPermissionsArgs.builder()        
 *             .projectId(exampleProject.id())
 *             .principal(example_readers.applyValue(example_readers -&gt; example_readers.id()))
 *             .serviceendpointId(exampleServiceEndpointDockerRegistry.id())
 *             .permissions(Map.ofEntries(
 *                 Map.entry(&#34;Use&#34;, &#34;allow&#34;),
 *                 Map.entry(&#34;Administer&#34;, &#34;deny&#34;),
 *                 Map.entry(&#34;Create&#34;, &#34;deny&#34;),
 *                 Map.entry(&#34;ViewAuthorization&#34;, &#34;allow&#34;),
 *                 Map.entry(&#34;ViewEndpoint&#34;, &#34;allow&#34;)
 *             ))
 *             .build());
 * 
 *     }
 * }
 * ```
 * ## Relevant Links
 * 
 * * [Azure DevOps Service REST API 7.0 - Security](https://docs.microsoft.com/en-us/rest/api/azure/devops/security/?view=azure-devops-rest-7.0)
 * 
 * ## PAT Permissions Required
 * 
 * - **Project &amp; Team**: vso.security_manage - Grants the ability to read, write, and manage security permissions.
 * 
 * ## Import
 * 
 * The resource does not support import.
 * 
 */
@ResourceType(type="azuredevops:index/serviceendpointPermissions:ServiceendpointPermissions")
public class ServiceendpointPermissions extends com.pulumi.resources.CustomResource {
    /**
     * the permissions to assign. The following permissions are available.
     * 
     */
    @Export(name="permissions", type=Map.class, parameters={String.class, String.class})
    private Output<Map<String,String>> permissions;

    /**
     * @return the permissions to assign. The following permissions are available.
     * 
     */
    public Output<Map<String,String>> permissions() {
        return this.permissions;
    }
    /**
     * The **group** principal to assign the permissions.
     * 
     */
    @Export(name="principal", type=String.class, parameters={})
    private Output<String> principal;

    /**
     * @return The **group** principal to assign the permissions.
     * 
     */
    public Output<String> principal() {
        return this.principal;
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
     * Replace (`true`) or merge (`false`) the permissions. Default: `true`
     * 
     * | Permission        | Description                         |
     * | ----------------- | ----------------------------------- |
     * | Use               | Use service endpoint                |
     * | Administer        | Full control over service endpoints |
     * | Create            | Create service endpoints            |
     * | ViewAuthorization | View authorizations                 |
     * | ViewEndpoint      | View service endpoint properties    |
     * 
     */
    @Export(name="replace", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> replace;

    /**
     * @return Replace (`true`) or merge (`false`) the permissions. Default: `true`
     * 
     * | Permission        | Description                         |
     * | ----------------- | ----------------------------------- |
     * | Use               | Use service endpoint                |
     * | Administer        | Full control over service endpoints |
     * | Create            | Create service endpoints            |
     * | ViewAuthorization | View authorizations                 |
     * | ViewEndpoint      | View service endpoint properties    |
     * 
     */
    public Output<Optional<Boolean>> replace() {
        return Codegen.optional(this.replace);
    }
    /**
     * The id of the service endpoint to assign the permissions.
     * 
     */
    @Export(name="serviceendpointId", type=String.class, parameters={})
    private Output</* @Nullable */ String> serviceendpointId;

    /**
     * @return The id of the service endpoint to assign the permissions.
     * 
     */
    public Output<Optional<String>> serviceendpointId() {
        return Codegen.optional(this.serviceendpointId);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ServiceendpointPermissions(String name) {
        this(name, ServiceendpointPermissionsArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ServiceendpointPermissions(String name, ServiceendpointPermissionsArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ServiceendpointPermissions(String name, ServiceendpointPermissionsArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/serviceendpointPermissions:ServiceendpointPermissions", name, args == null ? ServiceendpointPermissionsArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ServiceendpointPermissions(String name, Output<String> id, @Nullable ServiceendpointPermissionsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/serviceendpointPermissions:ServiceendpointPermissions", name, state, makeResourceOptions(options, id));
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
    public static ServiceendpointPermissions get(String name, Output<String> id, @Nullable ServiceendpointPermissionsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ServiceendpointPermissions(name, id, state, options);
    }
}
