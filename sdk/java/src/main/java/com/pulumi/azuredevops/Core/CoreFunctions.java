// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.Core;

import com.pulumi.azuredevops.Core.inputs.GetProjectArgs;
import com.pulumi.azuredevops.Core.inputs.GetProjectPlainArgs;
import com.pulumi.azuredevops.Core.inputs.GetProjectsArgs;
import com.pulumi.azuredevops.Core.inputs.GetProjectsPlainArgs;
import com.pulumi.azuredevops.Core.outputs.GetClientConfigResult;
import com.pulumi.azuredevops.Core.outputs.GetProjectResult;
import com.pulumi.azuredevops.Core.outputs.GetProjectsResult;
import com.pulumi.azuredevops.Utilities;
import com.pulumi.core.Output;
import com.pulumi.core.TypeShape;
import com.pulumi.deployment.Deployment;
import com.pulumi.deployment.InvokeOptions;
import com.pulumi.resources.InvokeArgs;
import java.util.concurrent.CompletableFuture;

public final class CoreFunctions {
    /**
     * Use this data source to access information about the Azure DevOps organization configured for the provider.
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
     *         final var example = AzuredevopsFunctions.getClientConfig();
     * 
     *         ctx.export(&#34;orgUrl&#34;, example.applyValue(getClientConfigResult -&gt; getClientConfigResult.organizationUrl()));
     *     }
     * }
     * ```
     * 
     * @deprecated
     * azuredevops.core.getClientConfig has been deprecated in favor of azuredevops.getClientConfig
     * 
     */
    @Deprecated /* azuredevops.core.getClientConfig has been deprecated in favor of azuredevops.getClientConfig */
    public static Output<GetClientConfigResult> getClientConfig() {
        return getClientConfig(InvokeArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * Use this data source to access information about the Azure DevOps organization configured for the provider.
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
     *         final var example = AzuredevopsFunctions.getClientConfig();
     * 
     *         ctx.export(&#34;orgUrl&#34;, example.applyValue(getClientConfigResult -&gt; getClientConfigResult.organizationUrl()));
     *     }
     * }
     * ```
     * 
     * @deprecated
     * azuredevops.core.getClientConfig has been deprecated in favor of azuredevops.getClientConfig
     * 
     */
    @Deprecated /* azuredevops.core.getClientConfig has been deprecated in favor of azuredevops.getClientConfig */
    public static CompletableFuture<GetClientConfigResult> getClientConfigPlain() {
        return getClientConfigPlain(InvokeArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * Use this data source to access information about the Azure DevOps organization configured for the provider.
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
     *         final var example = AzuredevopsFunctions.getClientConfig();
     * 
     *         ctx.export(&#34;orgUrl&#34;, example.applyValue(getClientConfigResult -&gt; getClientConfigResult.organizationUrl()));
     *     }
     * }
     * ```
     * 
     * @deprecated
     * azuredevops.core.getClientConfig has been deprecated in favor of azuredevops.getClientConfig
     * 
     */
    @Deprecated /* azuredevops.core.getClientConfig has been deprecated in favor of azuredevops.getClientConfig */
    public static Output<GetClientConfigResult> getClientConfig(InvokeArgs args) {
        return getClientConfig(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source to access information about the Azure DevOps organization configured for the provider.
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
     *         final var example = AzuredevopsFunctions.getClientConfig();
     * 
     *         ctx.export(&#34;orgUrl&#34;, example.applyValue(getClientConfigResult -&gt; getClientConfigResult.organizationUrl()));
     *     }
     * }
     * ```
     * 
     * @deprecated
     * azuredevops.core.getClientConfig has been deprecated in favor of azuredevops.getClientConfig
     * 
     */
    @Deprecated /* azuredevops.core.getClientConfig has been deprecated in favor of azuredevops.getClientConfig */
    public static CompletableFuture<GetClientConfigResult> getClientConfigPlain(InvokeArgs args) {
        return getClientConfigPlain(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source to access information about the Azure DevOps organization configured for the provider.
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
     *         final var example = AzuredevopsFunctions.getClientConfig();
     * 
     *         ctx.export(&#34;orgUrl&#34;, example.applyValue(getClientConfigResult -&gt; getClientConfigResult.organizationUrl()));
     *     }
     * }
     * ```
     * 
     * @deprecated
     * azuredevops.core.getClientConfig has been deprecated in favor of azuredevops.getClientConfig
     * 
     */
    @Deprecated /* azuredevops.core.getClientConfig has been deprecated in favor of azuredevops.getClientConfig */
    public static Output<GetClientConfigResult> getClientConfig(InvokeArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("azuredevops:Core/getClientConfig:getClientConfig", TypeShape.of(GetClientConfigResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source to access information about the Azure DevOps organization configured for the provider.
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
     *         final var example = AzuredevopsFunctions.getClientConfig();
     * 
     *         ctx.export(&#34;orgUrl&#34;, example.applyValue(getClientConfigResult -&gt; getClientConfigResult.organizationUrl()));
     *     }
     * }
     * ```
     * 
     * @deprecated
     * azuredevops.core.getClientConfig has been deprecated in favor of azuredevops.getClientConfig
     * 
     */
    @Deprecated /* azuredevops.core.getClientConfig has been deprecated in favor of azuredevops.getClientConfig */
    public static CompletableFuture<GetClientConfigResult> getClientConfigPlain(InvokeArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("azuredevops:Core/getClientConfig:getClientConfig", TypeShape.of(GetClientConfigResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source to access information about an existing Project within Azure DevOps.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.azuredevops.AzuredevopsFunctions;
     * import com.pulumi.azuredevops.inputs.GetProjectArgs;
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
     *         final var example = AzuredevopsFunctions.getProject(GetProjectArgs.builder()
     *             .name(&#34;Example Project&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;project&#34;, example.applyValue(getProjectResult -&gt; getProjectResult));
     *     }
     * }
     * ```
     * ## Relevant Links
     * 
     * - [Azure DevOps Service REST API 7.0 - Projects - Get](https://docs.microsoft.com/en-us/rest/api/azure/devops/core/projects/get?view=azure-devops-rest-7.0)
     * 
     * @deprecated
     * azuredevops.core.getProject has been deprecated in favor of azuredevops.getProject
     * 
     */
    @Deprecated /* azuredevops.core.getProject has been deprecated in favor of azuredevops.getProject */
    public static Output<GetProjectResult> getProject() {
        return getProject(GetProjectArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * Use this data source to access information about an existing Project within Azure DevOps.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.azuredevops.AzuredevopsFunctions;
     * import com.pulumi.azuredevops.inputs.GetProjectArgs;
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
     *         final var example = AzuredevopsFunctions.getProject(GetProjectArgs.builder()
     *             .name(&#34;Example Project&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;project&#34;, example.applyValue(getProjectResult -&gt; getProjectResult));
     *     }
     * }
     * ```
     * ## Relevant Links
     * 
     * - [Azure DevOps Service REST API 7.0 - Projects - Get](https://docs.microsoft.com/en-us/rest/api/azure/devops/core/projects/get?view=azure-devops-rest-7.0)
     * 
     * @deprecated
     * azuredevops.core.getProject has been deprecated in favor of azuredevops.getProject
     * 
     */
    @Deprecated /* azuredevops.core.getProject has been deprecated in favor of azuredevops.getProject */
    public static CompletableFuture<GetProjectResult> getProjectPlain() {
        return getProjectPlain(GetProjectPlainArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * Use this data source to access information about an existing Project within Azure DevOps.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.azuredevops.AzuredevopsFunctions;
     * import com.pulumi.azuredevops.inputs.GetProjectArgs;
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
     *         final var example = AzuredevopsFunctions.getProject(GetProjectArgs.builder()
     *             .name(&#34;Example Project&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;project&#34;, example.applyValue(getProjectResult -&gt; getProjectResult));
     *     }
     * }
     * ```
     * ## Relevant Links
     * 
     * - [Azure DevOps Service REST API 7.0 - Projects - Get](https://docs.microsoft.com/en-us/rest/api/azure/devops/core/projects/get?view=azure-devops-rest-7.0)
     * 
     * @deprecated
     * azuredevops.core.getProject has been deprecated in favor of azuredevops.getProject
     * 
     */
    @Deprecated /* azuredevops.core.getProject has been deprecated in favor of azuredevops.getProject */
    public static Output<GetProjectResult> getProject(GetProjectArgs args) {
        return getProject(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source to access information about an existing Project within Azure DevOps.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.azuredevops.AzuredevopsFunctions;
     * import com.pulumi.azuredevops.inputs.GetProjectArgs;
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
     *         final var example = AzuredevopsFunctions.getProject(GetProjectArgs.builder()
     *             .name(&#34;Example Project&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;project&#34;, example.applyValue(getProjectResult -&gt; getProjectResult));
     *     }
     * }
     * ```
     * ## Relevant Links
     * 
     * - [Azure DevOps Service REST API 7.0 - Projects - Get](https://docs.microsoft.com/en-us/rest/api/azure/devops/core/projects/get?view=azure-devops-rest-7.0)
     * 
     * @deprecated
     * azuredevops.core.getProject has been deprecated in favor of azuredevops.getProject
     * 
     */
    @Deprecated /* azuredevops.core.getProject has been deprecated in favor of azuredevops.getProject */
    public static CompletableFuture<GetProjectResult> getProjectPlain(GetProjectPlainArgs args) {
        return getProjectPlain(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source to access information about an existing Project within Azure DevOps.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.azuredevops.AzuredevopsFunctions;
     * import com.pulumi.azuredevops.inputs.GetProjectArgs;
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
     *         final var example = AzuredevopsFunctions.getProject(GetProjectArgs.builder()
     *             .name(&#34;Example Project&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;project&#34;, example.applyValue(getProjectResult -&gt; getProjectResult));
     *     }
     * }
     * ```
     * ## Relevant Links
     * 
     * - [Azure DevOps Service REST API 7.0 - Projects - Get](https://docs.microsoft.com/en-us/rest/api/azure/devops/core/projects/get?view=azure-devops-rest-7.0)
     * 
     * @deprecated
     * azuredevops.core.getProject has been deprecated in favor of azuredevops.getProject
     * 
     */
    @Deprecated /* azuredevops.core.getProject has been deprecated in favor of azuredevops.getProject */
    public static Output<GetProjectResult> getProject(GetProjectArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("azuredevops:Core/getProject:getProject", TypeShape.of(GetProjectResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source to access information about an existing Project within Azure DevOps.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.azuredevops.AzuredevopsFunctions;
     * import com.pulumi.azuredevops.inputs.GetProjectArgs;
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
     *         final var example = AzuredevopsFunctions.getProject(GetProjectArgs.builder()
     *             .name(&#34;Example Project&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;project&#34;, example.applyValue(getProjectResult -&gt; getProjectResult));
     *     }
     * }
     * ```
     * ## Relevant Links
     * 
     * - [Azure DevOps Service REST API 7.0 - Projects - Get](https://docs.microsoft.com/en-us/rest/api/azure/devops/core/projects/get?view=azure-devops-rest-7.0)
     * 
     * @deprecated
     * azuredevops.core.getProject has been deprecated in favor of azuredevops.getProject
     * 
     */
    @Deprecated /* azuredevops.core.getProject has been deprecated in favor of azuredevops.getProject */
    public static CompletableFuture<GetProjectResult> getProjectPlain(GetProjectPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("azuredevops:Core/getProject:getProject", TypeShape.of(GetProjectResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source to access information about existing Projects within Azure DevOps.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.azuredevops.AzuredevopsFunctions;
     * import com.pulumi.azuredevops.inputs.GetProjectsArgs;
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
     *         final var example = AzuredevopsFunctions.getProjects(GetProjectsArgs.builder()
     *             .name(&#34;Example Project&#34;)
     *             .state(&#34;wellFormed&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;projectId&#34;, example.applyValue(getProjectsResult -&gt; getProjectsResult.projects()).stream().map(element -&gt; element.projectId()).collect(toList()));
     *         ctx.export(&#34;name&#34;, example.applyValue(getProjectsResult -&gt; getProjectsResult.projects()).stream().map(element -&gt; element.name()).collect(toList()));
     *         ctx.export(&#34;projectUrl&#34;, example.applyValue(getProjectsResult -&gt; getProjectsResult.projects()).stream().map(element -&gt; element.projectUrl()).collect(toList()));
     *         ctx.export(&#34;state&#34;, example.applyValue(getProjectsResult -&gt; getProjectsResult.projects()).stream().map(element -&gt; element.state()).collect(toList()));
     *     }
     * }
     * ```
     * ## Relevant Links
     * 
     * - [Azure DevOps Service REST API 7.0 - Projects - Get](https://docs.microsoft.com/en-us/rest/api/azure/devops/core/projects/get?view=azure-devops-rest-7.0)
     * 
     * @deprecated
     * azuredevops.core.getProjects has been deprecated in favor of azuredevops.getProjects
     * 
     */
    @Deprecated /* azuredevops.core.getProjects has been deprecated in favor of azuredevops.getProjects */
    public static Output<GetProjectsResult> getProjects() {
        return getProjects(GetProjectsArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * Use this data source to access information about existing Projects within Azure DevOps.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.azuredevops.AzuredevopsFunctions;
     * import com.pulumi.azuredevops.inputs.GetProjectsArgs;
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
     *         final var example = AzuredevopsFunctions.getProjects(GetProjectsArgs.builder()
     *             .name(&#34;Example Project&#34;)
     *             .state(&#34;wellFormed&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;projectId&#34;, example.applyValue(getProjectsResult -&gt; getProjectsResult.projects()).stream().map(element -&gt; element.projectId()).collect(toList()));
     *         ctx.export(&#34;name&#34;, example.applyValue(getProjectsResult -&gt; getProjectsResult.projects()).stream().map(element -&gt; element.name()).collect(toList()));
     *         ctx.export(&#34;projectUrl&#34;, example.applyValue(getProjectsResult -&gt; getProjectsResult.projects()).stream().map(element -&gt; element.projectUrl()).collect(toList()));
     *         ctx.export(&#34;state&#34;, example.applyValue(getProjectsResult -&gt; getProjectsResult.projects()).stream().map(element -&gt; element.state()).collect(toList()));
     *     }
     * }
     * ```
     * ## Relevant Links
     * 
     * - [Azure DevOps Service REST API 7.0 - Projects - Get](https://docs.microsoft.com/en-us/rest/api/azure/devops/core/projects/get?view=azure-devops-rest-7.0)
     * 
     * @deprecated
     * azuredevops.core.getProjects has been deprecated in favor of azuredevops.getProjects
     * 
     */
    @Deprecated /* azuredevops.core.getProjects has been deprecated in favor of azuredevops.getProjects */
    public static CompletableFuture<GetProjectsResult> getProjectsPlain() {
        return getProjectsPlain(GetProjectsPlainArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * Use this data source to access information about existing Projects within Azure DevOps.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.azuredevops.AzuredevopsFunctions;
     * import com.pulumi.azuredevops.inputs.GetProjectsArgs;
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
     *         final var example = AzuredevopsFunctions.getProjects(GetProjectsArgs.builder()
     *             .name(&#34;Example Project&#34;)
     *             .state(&#34;wellFormed&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;projectId&#34;, example.applyValue(getProjectsResult -&gt; getProjectsResult.projects()).stream().map(element -&gt; element.projectId()).collect(toList()));
     *         ctx.export(&#34;name&#34;, example.applyValue(getProjectsResult -&gt; getProjectsResult.projects()).stream().map(element -&gt; element.name()).collect(toList()));
     *         ctx.export(&#34;projectUrl&#34;, example.applyValue(getProjectsResult -&gt; getProjectsResult.projects()).stream().map(element -&gt; element.projectUrl()).collect(toList()));
     *         ctx.export(&#34;state&#34;, example.applyValue(getProjectsResult -&gt; getProjectsResult.projects()).stream().map(element -&gt; element.state()).collect(toList()));
     *     }
     * }
     * ```
     * ## Relevant Links
     * 
     * - [Azure DevOps Service REST API 7.0 - Projects - Get](https://docs.microsoft.com/en-us/rest/api/azure/devops/core/projects/get?view=azure-devops-rest-7.0)
     * 
     * @deprecated
     * azuredevops.core.getProjects has been deprecated in favor of azuredevops.getProjects
     * 
     */
    @Deprecated /* azuredevops.core.getProjects has been deprecated in favor of azuredevops.getProjects */
    public static Output<GetProjectsResult> getProjects(GetProjectsArgs args) {
        return getProjects(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source to access information about existing Projects within Azure DevOps.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.azuredevops.AzuredevopsFunctions;
     * import com.pulumi.azuredevops.inputs.GetProjectsArgs;
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
     *         final var example = AzuredevopsFunctions.getProjects(GetProjectsArgs.builder()
     *             .name(&#34;Example Project&#34;)
     *             .state(&#34;wellFormed&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;projectId&#34;, example.applyValue(getProjectsResult -&gt; getProjectsResult.projects()).stream().map(element -&gt; element.projectId()).collect(toList()));
     *         ctx.export(&#34;name&#34;, example.applyValue(getProjectsResult -&gt; getProjectsResult.projects()).stream().map(element -&gt; element.name()).collect(toList()));
     *         ctx.export(&#34;projectUrl&#34;, example.applyValue(getProjectsResult -&gt; getProjectsResult.projects()).stream().map(element -&gt; element.projectUrl()).collect(toList()));
     *         ctx.export(&#34;state&#34;, example.applyValue(getProjectsResult -&gt; getProjectsResult.projects()).stream().map(element -&gt; element.state()).collect(toList()));
     *     }
     * }
     * ```
     * ## Relevant Links
     * 
     * - [Azure DevOps Service REST API 7.0 - Projects - Get](https://docs.microsoft.com/en-us/rest/api/azure/devops/core/projects/get?view=azure-devops-rest-7.0)
     * 
     * @deprecated
     * azuredevops.core.getProjects has been deprecated in favor of azuredevops.getProjects
     * 
     */
    @Deprecated /* azuredevops.core.getProjects has been deprecated in favor of azuredevops.getProjects */
    public static CompletableFuture<GetProjectsResult> getProjectsPlain(GetProjectsPlainArgs args) {
        return getProjectsPlain(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source to access information about existing Projects within Azure DevOps.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.azuredevops.AzuredevopsFunctions;
     * import com.pulumi.azuredevops.inputs.GetProjectsArgs;
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
     *         final var example = AzuredevopsFunctions.getProjects(GetProjectsArgs.builder()
     *             .name(&#34;Example Project&#34;)
     *             .state(&#34;wellFormed&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;projectId&#34;, example.applyValue(getProjectsResult -&gt; getProjectsResult.projects()).stream().map(element -&gt; element.projectId()).collect(toList()));
     *         ctx.export(&#34;name&#34;, example.applyValue(getProjectsResult -&gt; getProjectsResult.projects()).stream().map(element -&gt; element.name()).collect(toList()));
     *         ctx.export(&#34;projectUrl&#34;, example.applyValue(getProjectsResult -&gt; getProjectsResult.projects()).stream().map(element -&gt; element.projectUrl()).collect(toList()));
     *         ctx.export(&#34;state&#34;, example.applyValue(getProjectsResult -&gt; getProjectsResult.projects()).stream().map(element -&gt; element.state()).collect(toList()));
     *     }
     * }
     * ```
     * ## Relevant Links
     * 
     * - [Azure DevOps Service REST API 7.0 - Projects - Get](https://docs.microsoft.com/en-us/rest/api/azure/devops/core/projects/get?view=azure-devops-rest-7.0)
     * 
     * @deprecated
     * azuredevops.core.getProjects has been deprecated in favor of azuredevops.getProjects
     * 
     */
    @Deprecated /* azuredevops.core.getProjects has been deprecated in favor of azuredevops.getProjects */
    public static Output<GetProjectsResult> getProjects(GetProjectsArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("azuredevops:Core/getProjects:getProjects", TypeShape.of(GetProjectsResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source to access information about existing Projects within Azure DevOps.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.azuredevops.AzuredevopsFunctions;
     * import com.pulumi.azuredevops.inputs.GetProjectsArgs;
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
     *         final var example = AzuredevopsFunctions.getProjects(GetProjectsArgs.builder()
     *             .name(&#34;Example Project&#34;)
     *             .state(&#34;wellFormed&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;projectId&#34;, example.applyValue(getProjectsResult -&gt; getProjectsResult.projects()).stream().map(element -&gt; element.projectId()).collect(toList()));
     *         ctx.export(&#34;name&#34;, example.applyValue(getProjectsResult -&gt; getProjectsResult.projects()).stream().map(element -&gt; element.name()).collect(toList()));
     *         ctx.export(&#34;projectUrl&#34;, example.applyValue(getProjectsResult -&gt; getProjectsResult.projects()).stream().map(element -&gt; element.projectUrl()).collect(toList()));
     *         ctx.export(&#34;state&#34;, example.applyValue(getProjectsResult -&gt; getProjectsResult.projects()).stream().map(element -&gt; element.state()).collect(toList()));
     *     }
     * }
     * ```
     * ## Relevant Links
     * 
     * - [Azure DevOps Service REST API 7.0 - Projects - Get](https://docs.microsoft.com/en-us/rest/api/azure/devops/core/projects/get?view=azure-devops-rest-7.0)
     * 
     * @deprecated
     * azuredevops.core.getProjects has been deprecated in favor of azuredevops.getProjects
     * 
     */
    @Deprecated /* azuredevops.core.getProjects has been deprecated in favor of azuredevops.getProjects */
    public static CompletableFuture<GetProjectsResult> getProjectsPlain(GetProjectsPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("azuredevops:Core/getProjects:getProjects", TypeShape.of(GetProjectsResult.class), args, Utilities.withVersion(options));
    }
}
