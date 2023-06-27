// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops;

import com.pulumi.azuredevops.ServiceEndpointArtifactoryArgs;
import com.pulumi.azuredevops.Utilities;
import com.pulumi.azuredevops.inputs.ServiceEndpointArtifactoryState;
import com.pulumi.azuredevops.outputs.ServiceEndpointArtifactoryAuthenticationBasic;
import com.pulumi.azuredevops.outputs.ServiceEndpointArtifactoryAuthenticationToken;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages an Artifactory server endpoint within an Azure DevOps organization. Using this service endpoint requires you to first install [JFrog Artifactory Extension](https://marketplace.visualstudio.com/items?itemName=JFrog.jfrog-artifactory-vsts-extension).
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
 * import com.pulumi.azuredevops.ServiceEndpointArtifactory;
 * import com.pulumi.azuredevops.ServiceEndpointArtifactoryArgs;
 * import com.pulumi.azuredevops.inputs.ServiceEndpointArtifactoryAuthenticationTokenArgs;
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
 *         var exampleServiceEndpointArtifactory = new ServiceEndpointArtifactory(&#34;exampleServiceEndpointArtifactory&#34;, ServiceEndpointArtifactoryArgs.builder()        
 *             .projectId(exampleProject.id())
 *             .serviceEndpointName(&#34;Example Artifactory&#34;)
 *             .description(&#34;Managed by Terraform&#34;)
 *             .url(&#34;https://artifactory.my.com&#34;)
 *             .authenticationToken(ServiceEndpointArtifactoryAuthenticationTokenArgs.builder()
 *                 .token(&#34;0000000000000000000000000000000000000000&#34;)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * Alternatively a username and password may be used.
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.azuredevops.Project;
 * import com.pulumi.azuredevops.ProjectArgs;
 * import com.pulumi.azuredevops.ServiceEndpointArtifactory;
 * import com.pulumi.azuredevops.ServiceEndpointArtifactoryArgs;
 * import com.pulumi.azuredevops.inputs.ServiceEndpointArtifactoryAuthenticationBasicArgs;
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
 *         var exampleServiceEndpointArtifactory = new ServiceEndpointArtifactory(&#34;exampleServiceEndpointArtifactory&#34;, ServiceEndpointArtifactoryArgs.builder()        
 *             .projectId(exampleProject.id())
 *             .serviceEndpointName(&#34;Example Artifactory&#34;)
 *             .description(&#34;Managed by Terraform&#34;)
 *             .url(&#34;https://artifactory.my.com&#34;)
 *             .authenticationBasic(ServiceEndpointArtifactoryAuthenticationBasicArgs.builder()
 *                 .username(&#34;username&#34;)
 *                 .password(&#34;password&#34;)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ## Relevant Links
 * 
 * * [Azure DevOps Service Connections](https://docs.microsoft.com/en-us/azure/devops/pipelines/library/service-endpoints?view=azure-devops&amp;tabs=yaml)
 * * [Artifactory User Token](https://docs.artifactory.org/latest/user-guide/user-token/)
 * 
 * ## Import
 * 
 * Azure DevOps Service Endpoint Artifactory can be imported using the **projectID/serviceEndpointID**, e.g.
 * 
 * ```sh
 *  $ pulumi import azuredevops:index/serviceEndpointArtifactory:ServiceEndpointArtifactory example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
 * ```
 * 
 */
@ResourceType(type="azuredevops:index/serviceEndpointArtifactory:ServiceEndpointArtifactory")
public class ServiceEndpointArtifactory extends com.pulumi.resources.CustomResource {
    @Export(name="authenticationBasic", type=ServiceEndpointArtifactoryAuthenticationBasic.class, parameters={})
    private Output</* @Nullable */ ServiceEndpointArtifactoryAuthenticationBasic> authenticationBasic;

    public Output<Optional<ServiceEndpointArtifactoryAuthenticationBasic>> authenticationBasic() {
        return Codegen.optional(this.authenticationBasic);
    }
    @Export(name="authenticationToken", type=ServiceEndpointArtifactoryAuthenticationToken.class, parameters={})
    private Output</* @Nullable */ ServiceEndpointArtifactoryAuthenticationToken> authenticationToken;

    public Output<Optional<ServiceEndpointArtifactoryAuthenticationToken>> authenticationToken() {
        return Codegen.optional(this.authenticationToken);
    }
    @Export(name="authorization", type=Map.class, parameters={String.class, String.class})
    private Output<Map<String,String>> authorization;

    public Output<Map<String,String>> authorization() {
        return this.authorization;
    }
    /**
     * The Service Endpoint description.
     * 
     */
    @Export(name="description", type=String.class, parameters={})
    private Output</* @Nullable */ String> description;

    /**
     * @return The Service Endpoint description.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
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
     * URL of the Artifactory server to connect with.
     * 
     * _Note: URL should not end in a slash character._
     * * either `authentication_token` or `authentication_basic` (one is required)
     * 
     */
    @Export(name="url", type=String.class, parameters={})
    private Output<String> url;

    /**
     * @return URL of the Artifactory server to connect with.
     * 
     * _Note: URL should not end in a slash character._
     * * either `authentication_token` or `authentication_basic` (one is required)
     * 
     */
    public Output<String> url() {
        return this.url;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ServiceEndpointArtifactory(String name) {
        this(name, ServiceEndpointArtifactoryArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ServiceEndpointArtifactory(String name, ServiceEndpointArtifactoryArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ServiceEndpointArtifactory(String name, ServiceEndpointArtifactoryArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/serviceEndpointArtifactory:ServiceEndpointArtifactory", name, args == null ? ServiceEndpointArtifactoryArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ServiceEndpointArtifactory(String name, Output<String> id, @Nullable ServiceEndpointArtifactoryState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/serviceEndpointArtifactory:ServiceEndpointArtifactory", name, state, makeResourceOptions(options, id));
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
    public static ServiceEndpointArtifactory get(String name, Output<String> id, @Nullable ServiceEndpointArtifactoryState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ServiceEndpointArtifactory(name, id, state, options);
    }
}
