// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.Agent;

import com.pulumi.azuredevops.Agent.inputs.GetPoolArgs;
import com.pulumi.azuredevops.Agent.inputs.GetPoolPlainArgs;
import com.pulumi.azuredevops.Agent.outputs.GetPoolResult;
import com.pulumi.azuredevops.Agent.outputs.GetPoolsResult;
import com.pulumi.azuredevops.Utilities;
import com.pulumi.core.Output;
import com.pulumi.core.TypeShape;
import com.pulumi.deployment.Deployment;
import com.pulumi.deployment.InvokeOptions;
import com.pulumi.resources.InvokeArgs;
import java.util.concurrent.CompletableFuture;

public final class AgentFunctions {
    /**
     * Use this data source to access information about an existing Agent Pool within Azure DevOps.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.azuredevops.AzuredevopsFunctions;
     * import com.pulumi.azuredevops.inputs.GetPoolArgs;
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
     *         final var example = AzuredevopsFunctions.getPool(GetPoolArgs.builder()
     *             .name(&#34;Example Agent Pool&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;name&#34;, example.applyValue(getPoolResult -&gt; getPoolResult.name()));
     *         ctx.export(&#34;poolType&#34;, example.applyValue(getPoolResult -&gt; getPoolResult.poolType()));
     *         ctx.export(&#34;autoProvision&#34;, example.applyValue(getPoolResult -&gt; getPoolResult.autoProvision()));
     *         ctx.export(&#34;autoUpdate&#34;, example.applyValue(getPoolResult -&gt; getPoolResult.autoUpdate()));
     *     }
     * }
     * ```
     * ## Relevant Links
     * 
     * - [Azure DevOps Service REST API 7.0 - Agent Pools - Get](https://docs.microsoft.com/en-us/rest/api/azure/devops/distributedtask/pools/get?view=azure-devops-rest-7.0)
     * 
     * @deprecated
     * azuredevops.agent.getPool has been deprecated in favor of azuredevops.getPool
     * 
     */
    @Deprecated /* azuredevops.agent.getPool has been deprecated in favor of azuredevops.getPool */
    public static Output<GetPoolResult> getPool(GetPoolArgs args) {
        return getPool(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source to access information about an existing Agent Pool within Azure DevOps.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.azuredevops.AzuredevopsFunctions;
     * import com.pulumi.azuredevops.inputs.GetPoolArgs;
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
     *         final var example = AzuredevopsFunctions.getPool(GetPoolArgs.builder()
     *             .name(&#34;Example Agent Pool&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;name&#34;, example.applyValue(getPoolResult -&gt; getPoolResult.name()));
     *         ctx.export(&#34;poolType&#34;, example.applyValue(getPoolResult -&gt; getPoolResult.poolType()));
     *         ctx.export(&#34;autoProvision&#34;, example.applyValue(getPoolResult -&gt; getPoolResult.autoProvision()));
     *         ctx.export(&#34;autoUpdate&#34;, example.applyValue(getPoolResult -&gt; getPoolResult.autoUpdate()));
     *     }
     * }
     * ```
     * ## Relevant Links
     * 
     * - [Azure DevOps Service REST API 7.0 - Agent Pools - Get](https://docs.microsoft.com/en-us/rest/api/azure/devops/distributedtask/pools/get?view=azure-devops-rest-7.0)
     * 
     * @deprecated
     * azuredevops.agent.getPool has been deprecated in favor of azuredevops.getPool
     * 
     */
    @Deprecated /* azuredevops.agent.getPool has been deprecated in favor of azuredevops.getPool */
    public static CompletableFuture<GetPoolResult> getPoolPlain(GetPoolPlainArgs args) {
        return getPoolPlain(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source to access information about an existing Agent Pool within Azure DevOps.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.azuredevops.AzuredevopsFunctions;
     * import com.pulumi.azuredevops.inputs.GetPoolArgs;
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
     *         final var example = AzuredevopsFunctions.getPool(GetPoolArgs.builder()
     *             .name(&#34;Example Agent Pool&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;name&#34;, example.applyValue(getPoolResult -&gt; getPoolResult.name()));
     *         ctx.export(&#34;poolType&#34;, example.applyValue(getPoolResult -&gt; getPoolResult.poolType()));
     *         ctx.export(&#34;autoProvision&#34;, example.applyValue(getPoolResult -&gt; getPoolResult.autoProvision()));
     *         ctx.export(&#34;autoUpdate&#34;, example.applyValue(getPoolResult -&gt; getPoolResult.autoUpdate()));
     *     }
     * }
     * ```
     * ## Relevant Links
     * 
     * - [Azure DevOps Service REST API 7.0 - Agent Pools - Get](https://docs.microsoft.com/en-us/rest/api/azure/devops/distributedtask/pools/get?view=azure-devops-rest-7.0)
     * 
     * @deprecated
     * azuredevops.agent.getPool has been deprecated in favor of azuredevops.getPool
     * 
     */
    @Deprecated /* azuredevops.agent.getPool has been deprecated in favor of azuredevops.getPool */
    public static Output<GetPoolResult> getPool(GetPoolArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("azuredevops:Agent/getPool:getPool", TypeShape.of(GetPoolResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source to access information about an existing Agent Pool within Azure DevOps.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.azuredevops.AzuredevopsFunctions;
     * import com.pulumi.azuredevops.inputs.GetPoolArgs;
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
     *         final var example = AzuredevopsFunctions.getPool(GetPoolArgs.builder()
     *             .name(&#34;Example Agent Pool&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;name&#34;, example.applyValue(getPoolResult -&gt; getPoolResult.name()));
     *         ctx.export(&#34;poolType&#34;, example.applyValue(getPoolResult -&gt; getPoolResult.poolType()));
     *         ctx.export(&#34;autoProvision&#34;, example.applyValue(getPoolResult -&gt; getPoolResult.autoProvision()));
     *         ctx.export(&#34;autoUpdate&#34;, example.applyValue(getPoolResult -&gt; getPoolResult.autoUpdate()));
     *     }
     * }
     * ```
     * ## Relevant Links
     * 
     * - [Azure DevOps Service REST API 7.0 - Agent Pools - Get](https://docs.microsoft.com/en-us/rest/api/azure/devops/distributedtask/pools/get?view=azure-devops-rest-7.0)
     * 
     * @deprecated
     * azuredevops.agent.getPool has been deprecated in favor of azuredevops.getPool
     * 
     */
    @Deprecated /* azuredevops.agent.getPool has been deprecated in favor of azuredevops.getPool */
    public static CompletableFuture<GetPoolResult> getPoolPlain(GetPoolPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("azuredevops:Agent/getPool:getPool", TypeShape.of(GetPoolResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source to access information about existing Agent Pools within Azure DevOps.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.azuredevops.AzuredevopsFunctions;
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
     *         final var example = AzuredevopsFunctions.getPools();
     * 
     *         ctx.export(&#34;agentPoolName&#34;, example.applyValue(getPoolsResult -&gt; getPoolsResult.agentPools()).stream().map(element -&gt; element.name()).collect(toList()));
     *         ctx.export(&#34;autoProvision&#34;, example.applyValue(getPoolsResult -&gt; getPoolsResult.agentPools()).stream().map(element -&gt; element.autoProvision()).collect(toList()));
     *         ctx.export(&#34;autoUpdate&#34;, example.applyValue(getPoolsResult -&gt; getPoolsResult.agentPools()).stream().map(element -&gt; element.autoUpdate()).collect(toList()));
     *         ctx.export(&#34;poolType&#34;, example.applyValue(getPoolsResult -&gt; getPoolsResult.agentPools()).stream().map(element -&gt; element.poolType()).collect(toList()));
     *     }
     * }
     * ```
     * ## Relevant Links
     * 
     * - [Azure DevOps Service REST API 7.0 - Agent Pools - Get](https://docs.microsoft.com/en-us/rest/api/azure/devops/distributedtask/pools/get?view=azure-devops-rest-7.0)
     * 
     * @deprecated
     * azuredevops.agent.getPools has been deprecated in favor of azuredevops.getPools
     * 
     */
    @Deprecated /* azuredevops.agent.getPools has been deprecated in favor of azuredevops.getPools */
    public static Output<GetPoolsResult> getPools() {
        return getPools(InvokeArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * Use this data source to access information about existing Agent Pools within Azure DevOps.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.azuredevops.AzuredevopsFunctions;
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
     *         final var example = AzuredevopsFunctions.getPools();
     * 
     *         ctx.export(&#34;agentPoolName&#34;, example.applyValue(getPoolsResult -&gt; getPoolsResult.agentPools()).stream().map(element -&gt; element.name()).collect(toList()));
     *         ctx.export(&#34;autoProvision&#34;, example.applyValue(getPoolsResult -&gt; getPoolsResult.agentPools()).stream().map(element -&gt; element.autoProvision()).collect(toList()));
     *         ctx.export(&#34;autoUpdate&#34;, example.applyValue(getPoolsResult -&gt; getPoolsResult.agentPools()).stream().map(element -&gt; element.autoUpdate()).collect(toList()));
     *         ctx.export(&#34;poolType&#34;, example.applyValue(getPoolsResult -&gt; getPoolsResult.agentPools()).stream().map(element -&gt; element.poolType()).collect(toList()));
     *     }
     * }
     * ```
     * ## Relevant Links
     * 
     * - [Azure DevOps Service REST API 7.0 - Agent Pools - Get](https://docs.microsoft.com/en-us/rest/api/azure/devops/distributedtask/pools/get?view=azure-devops-rest-7.0)
     * 
     * @deprecated
     * azuredevops.agent.getPools has been deprecated in favor of azuredevops.getPools
     * 
     */
    @Deprecated /* azuredevops.agent.getPools has been deprecated in favor of azuredevops.getPools */
    public static CompletableFuture<GetPoolsResult> getPoolsPlain() {
        return getPoolsPlain(InvokeArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * Use this data source to access information about existing Agent Pools within Azure DevOps.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.azuredevops.AzuredevopsFunctions;
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
     *         final var example = AzuredevopsFunctions.getPools();
     * 
     *         ctx.export(&#34;agentPoolName&#34;, example.applyValue(getPoolsResult -&gt; getPoolsResult.agentPools()).stream().map(element -&gt; element.name()).collect(toList()));
     *         ctx.export(&#34;autoProvision&#34;, example.applyValue(getPoolsResult -&gt; getPoolsResult.agentPools()).stream().map(element -&gt; element.autoProvision()).collect(toList()));
     *         ctx.export(&#34;autoUpdate&#34;, example.applyValue(getPoolsResult -&gt; getPoolsResult.agentPools()).stream().map(element -&gt; element.autoUpdate()).collect(toList()));
     *         ctx.export(&#34;poolType&#34;, example.applyValue(getPoolsResult -&gt; getPoolsResult.agentPools()).stream().map(element -&gt; element.poolType()).collect(toList()));
     *     }
     * }
     * ```
     * ## Relevant Links
     * 
     * - [Azure DevOps Service REST API 7.0 - Agent Pools - Get](https://docs.microsoft.com/en-us/rest/api/azure/devops/distributedtask/pools/get?view=azure-devops-rest-7.0)
     * 
     * @deprecated
     * azuredevops.agent.getPools has been deprecated in favor of azuredevops.getPools
     * 
     */
    @Deprecated /* azuredevops.agent.getPools has been deprecated in favor of azuredevops.getPools */
    public static Output<GetPoolsResult> getPools(InvokeArgs args) {
        return getPools(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source to access information about existing Agent Pools within Azure DevOps.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.azuredevops.AzuredevopsFunctions;
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
     *         final var example = AzuredevopsFunctions.getPools();
     * 
     *         ctx.export(&#34;agentPoolName&#34;, example.applyValue(getPoolsResult -&gt; getPoolsResult.agentPools()).stream().map(element -&gt; element.name()).collect(toList()));
     *         ctx.export(&#34;autoProvision&#34;, example.applyValue(getPoolsResult -&gt; getPoolsResult.agentPools()).stream().map(element -&gt; element.autoProvision()).collect(toList()));
     *         ctx.export(&#34;autoUpdate&#34;, example.applyValue(getPoolsResult -&gt; getPoolsResult.agentPools()).stream().map(element -&gt; element.autoUpdate()).collect(toList()));
     *         ctx.export(&#34;poolType&#34;, example.applyValue(getPoolsResult -&gt; getPoolsResult.agentPools()).stream().map(element -&gt; element.poolType()).collect(toList()));
     *     }
     * }
     * ```
     * ## Relevant Links
     * 
     * - [Azure DevOps Service REST API 7.0 - Agent Pools - Get](https://docs.microsoft.com/en-us/rest/api/azure/devops/distributedtask/pools/get?view=azure-devops-rest-7.0)
     * 
     * @deprecated
     * azuredevops.agent.getPools has been deprecated in favor of azuredevops.getPools
     * 
     */
    @Deprecated /* azuredevops.agent.getPools has been deprecated in favor of azuredevops.getPools */
    public static CompletableFuture<GetPoolsResult> getPoolsPlain(InvokeArgs args) {
        return getPoolsPlain(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source to access information about existing Agent Pools within Azure DevOps.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.azuredevops.AzuredevopsFunctions;
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
     *         final var example = AzuredevopsFunctions.getPools();
     * 
     *         ctx.export(&#34;agentPoolName&#34;, example.applyValue(getPoolsResult -&gt; getPoolsResult.agentPools()).stream().map(element -&gt; element.name()).collect(toList()));
     *         ctx.export(&#34;autoProvision&#34;, example.applyValue(getPoolsResult -&gt; getPoolsResult.agentPools()).stream().map(element -&gt; element.autoProvision()).collect(toList()));
     *         ctx.export(&#34;autoUpdate&#34;, example.applyValue(getPoolsResult -&gt; getPoolsResult.agentPools()).stream().map(element -&gt; element.autoUpdate()).collect(toList()));
     *         ctx.export(&#34;poolType&#34;, example.applyValue(getPoolsResult -&gt; getPoolsResult.agentPools()).stream().map(element -&gt; element.poolType()).collect(toList()));
     *     }
     * }
     * ```
     * ## Relevant Links
     * 
     * - [Azure DevOps Service REST API 7.0 - Agent Pools - Get](https://docs.microsoft.com/en-us/rest/api/azure/devops/distributedtask/pools/get?view=azure-devops-rest-7.0)
     * 
     * @deprecated
     * azuredevops.agent.getPools has been deprecated in favor of azuredevops.getPools
     * 
     */
    @Deprecated /* azuredevops.agent.getPools has been deprecated in favor of azuredevops.getPools */
    public static Output<GetPoolsResult> getPools(InvokeArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("azuredevops:Agent/getPools:getPools", TypeShape.of(GetPoolsResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source to access information about existing Agent Pools within Azure DevOps.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.azuredevops.AzuredevopsFunctions;
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
     *         final var example = AzuredevopsFunctions.getPools();
     * 
     *         ctx.export(&#34;agentPoolName&#34;, example.applyValue(getPoolsResult -&gt; getPoolsResult.agentPools()).stream().map(element -&gt; element.name()).collect(toList()));
     *         ctx.export(&#34;autoProvision&#34;, example.applyValue(getPoolsResult -&gt; getPoolsResult.agentPools()).stream().map(element -&gt; element.autoProvision()).collect(toList()));
     *         ctx.export(&#34;autoUpdate&#34;, example.applyValue(getPoolsResult -&gt; getPoolsResult.agentPools()).stream().map(element -&gt; element.autoUpdate()).collect(toList()));
     *         ctx.export(&#34;poolType&#34;, example.applyValue(getPoolsResult -&gt; getPoolsResult.agentPools()).stream().map(element -&gt; element.poolType()).collect(toList()));
     *     }
     * }
     * ```
     * ## Relevant Links
     * 
     * - [Azure DevOps Service REST API 7.0 - Agent Pools - Get](https://docs.microsoft.com/en-us/rest/api/azure/devops/distributedtask/pools/get?view=azure-devops-rest-7.0)
     * 
     * @deprecated
     * azuredevops.agent.getPools has been deprecated in favor of azuredevops.getPools
     * 
     */
    @Deprecated /* azuredevops.agent.getPools has been deprecated in favor of azuredevops.getPools */
    public static CompletableFuture<GetPoolsResult> getPoolsPlain(InvokeArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("azuredevops:Agent/getPools:getPools", TypeShape.of(GetPoolsResult.class), args, Utilities.withVersion(options));
    }
}
