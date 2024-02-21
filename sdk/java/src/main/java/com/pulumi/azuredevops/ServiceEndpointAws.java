// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops;

import com.pulumi.azuredevops.ServiceEndpointAwsArgs;
import com.pulumi.azuredevops.Utilities;
import com.pulumi.azuredevops.inputs.ServiceEndpointAwsState;
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
 * Manages a AWS service endpoint within Azure DevOps. Using this service endpoint requires you to first install [AWS Toolkit for Azure DevOps](https://marketplace.visualstudio.com/items?itemName=AmazonWebServices.aws-vsts-tools).
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
 * import com.pulumi.azuredevops.ServiceEndpointAws;
 * import com.pulumi.azuredevops.ServiceEndpointAwsArgs;
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
 *         var exampleServiceEndpointAws = new ServiceEndpointAws(&#34;exampleServiceEndpointAws&#34;, ServiceEndpointAwsArgs.builder()        
 *             .projectId(exampleProject.id())
 *             .serviceEndpointName(&#34;Example AWS&#34;)
 *             .accessKeyId(&#34;00000000-0000-0000-0000-000000000000&#34;)
 *             .secretAccessKey(&#34;accesskey&#34;)
 *             .description(&#34;Managed by AzureDevOps&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ## Relevant Links
 * 
 * * [aws-toolkit-azure-devops](https://github.com/aws/aws-toolkit-azure-devops)
 * * [Azure DevOps Service REST API 7.0 - Agent Pools](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-7.0)
 * 
 * ## Import
 * 
 * Azure DevOps Service Endpoint AWS can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**
 * 
 * ```sh
 * $ pulumi import azuredevops:index/serviceEndpointAws:ServiceEndpointAws azuredevops_serviceendpoint_aws.example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
 * ```
 * 
 */
@ResourceType(type="azuredevops:index/serviceEndpointAws:ServiceEndpointAws")
public class ServiceEndpointAws extends com.pulumi.resources.CustomResource {
    /**
     * The AWS access key ID for signing programmatic requests.
     * 
     */
    @Export(name="accessKeyId", refs={String.class}, tree="[0]")
    private Output<String> accessKeyId;

    /**
     * @return The AWS access key ID for signing programmatic requests.
     * 
     */
    public Output<String> accessKeyId() {
        return this.accessKeyId;
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
     * A unique identifier that is used by third parties when assuming roles in their customers&#39; accounts, aka cross-account role access.
     * 
     */
    @Export(name="externalId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> externalId;

    /**
     * @return A unique identifier that is used by third parties when assuming roles in their customers&#39; accounts, aka cross-account role access.
     * 
     */
    public Output<Optional<String>> externalId() {
        return Codegen.optional(this.externalId);
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
     * Optional identifier for the assumed role session.
     * 
     */
    @Export(name="roleSessionName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> roleSessionName;

    /**
     * @return Optional identifier for the assumed role session.
     * 
     */
    public Output<Optional<String>> roleSessionName() {
        return Codegen.optional(this.roleSessionName);
    }
    /**
     * The Amazon Resource Name (ARN) of the role to assume.
     * 
     */
    @Export(name="roleToAssume", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> roleToAssume;

    /**
     * @return The Amazon Resource Name (ARN) of the role to assume.
     * 
     */
    public Output<Optional<String>> roleToAssume() {
        return Codegen.optional(this.roleToAssume);
    }
    /**
     * The AWS secret access key for signing programmatic requests.
     * 
     */
    @Export(name="secretAccessKey", refs={String.class}, tree="[0]")
    private Output<String> secretAccessKey;

    /**
     * @return The AWS secret access key for signing programmatic requests.
     * 
     */
    public Output<String> secretAccessKey() {
        return this.secretAccessKey;
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
     * The AWS session token for signing programmatic requests.
     * 
     */
    @Export(name="sessionToken", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> sessionToken;

    /**
     * @return The AWS session token for signing programmatic requests.
     * 
     */
    public Output<Optional<String>> sessionToken() {
        return Codegen.optional(this.sessionToken);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ServiceEndpointAws(String name) {
        this(name, ServiceEndpointAwsArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ServiceEndpointAws(String name, ServiceEndpointAwsArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ServiceEndpointAws(String name, ServiceEndpointAwsArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/serviceEndpointAws:ServiceEndpointAws", name, args == null ? ServiceEndpointAwsArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ServiceEndpointAws(String name, Output<String> id, @Nullable ServiceEndpointAwsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/serviceEndpointAws:ServiceEndpointAws", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "secretAccessKey",
                "sessionToken"
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
    public static ServiceEndpointAws get(String name, Output<String> id, @Nullable ServiceEndpointAwsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ServiceEndpointAws(name, id, state, options);
    }
}
