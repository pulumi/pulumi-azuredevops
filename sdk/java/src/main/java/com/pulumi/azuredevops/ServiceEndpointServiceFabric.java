// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops;

import com.pulumi.azuredevops.ServiceEndpointServiceFabricArgs;
import com.pulumi.azuredevops.Utilities;
import com.pulumi.azuredevops.inputs.ServiceEndpointServiceFabricState;
import com.pulumi.azuredevops.outputs.ServiceEndpointServiceFabricAzureActiveDirectory;
import com.pulumi.azuredevops.outputs.ServiceEndpointServiceFabricCertificate;
import com.pulumi.azuredevops.outputs.ServiceEndpointServiceFabricNone;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a Service Fabric service endpoint within Azure DevOps.
 * 
 * ## Example Usage
 * ### Client Certificate Authentication
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.azuredevops.Project;
 * import com.pulumi.azuredevops.ProjectArgs;
 * import com.pulumi.azuredevops.ServiceEndpointServiceFabric;
 * import com.pulumi.azuredevops.ServiceEndpointServiceFabricArgs;
 * import com.pulumi.azuredevops.inputs.ServiceEndpointServiceFabricCertificateArgs;
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
 *         var exampleServiceEndpointServiceFabric = new ServiceEndpointServiceFabric(&#34;exampleServiceEndpointServiceFabric&#34;, ServiceEndpointServiceFabricArgs.builder()        
 *             .projectId(exampleProject.id())
 *             .serviceEndpointName(&#34;Example Service Fabric&#34;)
 *             .description(&#34;Managed by Terraform&#34;)
 *             .clusterEndpoint(&#34;tcp://test&#34;)
 *             .certificate(ServiceEndpointServiceFabricCertificateArgs.builder()
 *                 .serverCertificateLookup(&#34;Thumbprint&#34;)
 *                 .serverCertificateThumbprint(&#34;0000000000000000000000000000000000000000&#34;)
 *                 .clientCertificate(Base64.getEncoder().encodeToString(Files.readAllBytes(Paths.get(&#34;certificate.pfx&#34;))))
 *                 .clientCertificatePassword(&#34;password&#34;)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Azure Active Directory Authentication
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.azuredevops.Project;
 * import com.pulumi.azuredevops.ProjectArgs;
 * import com.pulumi.azuredevops.ServiceEndpointServiceFabric;
 * import com.pulumi.azuredevops.ServiceEndpointServiceFabricArgs;
 * import com.pulumi.azuredevops.inputs.ServiceEndpointServiceFabricAzureActiveDirectoryArgs;
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
 *         var project = new Project(&#34;project&#34;, ProjectArgs.builder()        
 *             .visibility(&#34;private&#34;)
 *             .versionControl(&#34;Git&#34;)
 *             .workItemTemplate(&#34;Agile&#34;)
 *             .build());
 * 
 *         var test = new ServiceEndpointServiceFabric(&#34;test&#34;, ServiceEndpointServiceFabricArgs.builder()        
 *             .projectId(project.id())
 *             .serviceEndpointName(&#34;Sample Service Fabric&#34;)
 *             .description(&#34;Managed by Terraform&#34;)
 *             .clusterEndpoint(&#34;tcp://test&#34;)
 *             .azureActiveDirectory(ServiceEndpointServiceFabricAzureActiveDirectoryArgs.builder()
 *                 .serverCertificateLookup(&#34;Thumbprint&#34;)
 *                 .serverCertificateThumbprint(&#34;0000000000000000000000000000000000000000&#34;)
 *                 .username(&#34;username&#34;)
 *                 .password(&#34;password&#34;)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Windows Authentication
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.azuredevops.Project;
 * import com.pulumi.azuredevops.ProjectArgs;
 * import com.pulumi.azuredevops.ServiceEndpointServiceFabric;
 * import com.pulumi.azuredevops.ServiceEndpointServiceFabricArgs;
 * import com.pulumi.azuredevops.inputs.ServiceEndpointServiceFabricNoneArgs;
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
 *         var project = new Project(&#34;project&#34;, ProjectArgs.builder()        
 *             .visibility(&#34;private&#34;)
 *             .versionControl(&#34;Git&#34;)
 *             .workItemTemplate(&#34;Agile&#34;)
 *             .build());
 * 
 *         var test = new ServiceEndpointServiceFabric(&#34;test&#34;, ServiceEndpointServiceFabricArgs.builder()        
 *             .projectId(project.id())
 *             .serviceEndpointName(&#34;Sample Service Fabric&#34;)
 *             .description(&#34;Managed by Terraform&#34;)
 *             .clusterEndpoint(&#34;tcp://test&#34;)
 *             .none(ServiceEndpointServiceFabricNoneArgs.builder()
 *                 .unsecured(false)
 *                 .clusterSpn(&#34;HTTP/www.contoso.com&#34;)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ## Relevant Links
 * 
 * - [Azure DevOps Service REST API 6.0 - Endpoints](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-6.0)
 * 
 * ## Import
 * 
 * Azure DevOps Service Endpoint Service Fabric can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**
 * 
 * ```sh
 *  $ pulumi import azuredevops:index/serviceEndpointServiceFabric:ServiceEndpointServiceFabric example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
 * ```
 * 
 */
@ResourceType(type="azuredevops:index/serviceEndpointServiceFabric:ServiceEndpointServiceFabric")
public class ServiceEndpointServiceFabric extends com.pulumi.resources.CustomResource {
    @Export(name="authorization", type=Map.class, parameters={String.class, String.class})
    private Output<Map<String,String>> authorization;

    public Output<Map<String,String>> authorization() {
        return this.authorization;
    }
    @Export(name="azureActiveDirectory", type=ServiceEndpointServiceFabricAzureActiveDirectory.class, parameters={})
    private Output</* @Nullable */ ServiceEndpointServiceFabricAzureActiveDirectory> azureActiveDirectory;

    public Output<Optional<ServiceEndpointServiceFabricAzureActiveDirectory>> azureActiveDirectory() {
        return Codegen.optional(this.azureActiveDirectory);
    }
    @Export(name="certificate", type=ServiceEndpointServiceFabricCertificate.class, parameters={})
    private Output</* @Nullable */ ServiceEndpointServiceFabricCertificate> certificate;

    public Output<Optional<ServiceEndpointServiceFabricCertificate>> certificate() {
        return Codegen.optional(this.certificate);
    }
    /**
     * Client connection endpoint for the cluster. Prefix the value with &#39;tcp://&#39;;. This value overrides the publish profile.
     * 
     */
    @Export(name="clusterEndpoint", type=String.class, parameters={})
    private Output<String> clusterEndpoint;

    /**
     * @return Client connection endpoint for the cluster. Prefix the value with &#39;tcp://&#39;;. This value overrides the publish profile.
     * 
     */
    public Output<String> clusterEndpoint() {
        return this.clusterEndpoint;
    }
    @Export(name="description", type=String.class, parameters={})
    private Output</* @Nullable */ String> description;

    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    @Export(name="none", type=ServiceEndpointServiceFabricNone.class, parameters={})
    private Output</* @Nullable */ ServiceEndpointServiceFabricNone> none;

    public Output<Optional<ServiceEndpointServiceFabricNone>> none() {
        return Codegen.optional(this.none);
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
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ServiceEndpointServiceFabric(String name) {
        this(name, ServiceEndpointServiceFabricArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ServiceEndpointServiceFabric(String name, ServiceEndpointServiceFabricArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ServiceEndpointServiceFabric(String name, ServiceEndpointServiceFabricArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/serviceEndpointServiceFabric:ServiceEndpointServiceFabric", name, args == null ? ServiceEndpointServiceFabricArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ServiceEndpointServiceFabric(String name, Output<String> id, @Nullable ServiceEndpointServiceFabricState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/serviceEndpointServiceFabric:ServiceEndpointServiceFabric", name, state, makeResourceOptions(options, id));
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
    public static ServiceEndpointServiceFabric get(String name, Output<String> id, @Nullable ServiceEndpointServiceFabricState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ServiceEndpointServiceFabric(name, id, state, options);
    }
}