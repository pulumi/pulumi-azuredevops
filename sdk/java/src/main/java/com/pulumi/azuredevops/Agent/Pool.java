// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.Agent;

import com.pulumi.azuredevops.Agent.PoolArgs;
import com.pulumi.azuredevops.Agent.inputs.PoolState;
import com.pulumi.azuredevops.Utilities;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages an agent pool within Azure DevOps.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.azuredevops.Pool;
 * import com.pulumi.azuredevops.PoolArgs;
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
 *         var example = new Pool(&#34;example&#34;, PoolArgs.builder()        
 *             .autoProvision(false)
 *             .autoUpdate(false)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ## Relevant Links
 * 
 * - [Azure DevOps Service REST API 6.0 - Agent Pools](https://docs.microsoft.com/en-us/rest/api/azure/devops/distributedtask/pools?view=azure-devops-rest-6.0)
 * 
 * ## Import
 * 
 * Azure DevOps Agent Pools can be imported using the agent pool ID, e.g.
 * 
 * ```sh
 *  $ pulumi import azuredevops:Agent/pool:Pool example 0
 * ```
 * 
 * @deprecated
 * azuredevops.agent.Pool has been deprecated in favor of azuredevops.Pool
 * 
 */
@Deprecated /* azuredevops.agent.Pool has been deprecated in favor of azuredevops.Pool */
@ResourceType(type="azuredevops:Agent/pool:Pool")
public class Pool extends com.pulumi.resources.CustomResource {
    /**
     * Specifies whether a queue should be automatically provisioned for each project collection. Defaults to `false`.
     * 
     */
    @Export(name="autoProvision", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> autoProvision;

    /**
     * @return Specifies whether a queue should be automatically provisioned for each project collection. Defaults to `false`.
     * 
     */
    public Output<Optional<Boolean>> autoProvision() {
        return Codegen.optional(this.autoProvision);
    }
    /**
     * Specifies whether or not agents within the pool should be automatically updated. Defaults to `true`.
     * 
     */
    @Export(name="autoUpdate", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> autoUpdate;

    /**
     * @return Specifies whether or not agents within the pool should be automatically updated. Defaults to `true`.
     * 
     */
    public Output<Optional<Boolean>> autoUpdate() {
        return Codegen.optional(this.autoUpdate);
    }
    /**
     * The name of the agent pool.
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return The name of the agent pool.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Specifies whether the agent pool type is Automation or Deployment. Defaults to `automation`.
     * 
     */
    @Export(name="poolType", type=String.class, parameters={})
    private Output</* @Nullable */ String> poolType;

    /**
     * @return Specifies whether the agent pool type is Automation or Deployment. Defaults to `automation`.
     * 
     */
    public Output<Optional<String>> poolType() {
        return Codegen.optional(this.poolType);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Pool(String name) {
        this(name, PoolArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Pool(String name, @Nullable PoolArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Pool(String name, @Nullable PoolArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:Agent/pool:Pool", name, args == null ? PoolArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Pool(String name, Output<String> id, @Nullable PoolState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:Agent/pool:Pool", name, state, makeResourceOptions(options, id));
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
    public static Pool get(String name, Output<String> id, @Nullable PoolState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Pool(name, id, state, options);
    }
}
