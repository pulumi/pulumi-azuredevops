// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops;

import com.pulumi.azuredevops.CheckBusinessHoursArgs;
import com.pulumi.azuredevops.Utilities;
import com.pulumi.azuredevops.inputs.CheckBusinessHoursState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a business hours check on a resource within Azure DevOps.
 * 
 * ## Example Usage
 * 
 * ### Protect a service connection
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.azuredevops.Project;
 * import com.pulumi.azuredevops.ProjectArgs;
 * import com.pulumi.azuredevops.ServiceEndpointGeneric;
 * import com.pulumi.azuredevops.ServiceEndpointGenericArgs;
 * import com.pulumi.azuredevops.CheckBusinessHours;
 * import com.pulumi.azuredevops.CheckBusinessHoursArgs;
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
 *         var example = new Project("example", ProjectArgs.builder()
 *             .name("Example Project")
 *             .build());
 * 
 *         var exampleServiceEndpointGeneric = new ServiceEndpointGeneric("exampleServiceEndpointGeneric", ServiceEndpointGenericArgs.builder()
 *             .projectId(example.id())
 *             .serverUrl("https://some-server.example.com")
 *             .username("username")
 *             .password("password")
 *             .serviceEndpointName("Example Generic")
 *             .description("Managed by Pulumi")
 *             .build());
 * 
 *         var exampleCheckBusinessHours = new CheckBusinessHours("exampleCheckBusinessHours", CheckBusinessHoursArgs.builder()
 *             .projectId(example.id())
 *             .displayName("Managed by Pulumi")
 *             .targetResourceId(exampleServiceEndpointGeneric.id())
 *             .targetResourceType("endpoint")
 *             .startTime("07:00")
 *             .endTime("15:30")
 *             .timeZone("UTC")
 *             .monday(true)
 *             .tuesday(true)
 *             .timeout(1440)
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ### Protect an environment
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.azuredevops.Project;
 * import com.pulumi.azuredevops.ProjectArgs;
 * import com.pulumi.azuredevops.Environment;
 * import com.pulumi.azuredevops.EnvironmentArgs;
 * import com.pulumi.azuredevops.CheckBusinessHours;
 * import com.pulumi.azuredevops.CheckBusinessHoursArgs;
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
 *         var example = new Project("example", ProjectArgs.builder()
 *             .name("Example Project")
 *             .build());
 * 
 *         var exampleEnvironment = new Environment("exampleEnvironment", EnvironmentArgs.builder()
 *             .projectId(example.id())
 *             .name("Example Environment")
 *             .build());
 * 
 *         var exampleCheckBusinessHours = new CheckBusinessHours("exampleCheckBusinessHours", CheckBusinessHoursArgs.builder()
 *             .projectId(example.id())
 *             .displayName("Managed by Pulumi")
 *             .targetResourceId(exampleEnvironment.id())
 *             .targetResourceType("environment")
 *             .startTime("07:00")
 *             .endTime("15:30")
 *             .timeZone("UTC")
 *             .monday(true)
 *             .tuesday(true)
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ### Protect an agent queue
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.azuredevops.Project;
 * import com.pulumi.azuredevops.ProjectArgs;
 * import com.pulumi.azuredevops.Pool;
 * import com.pulumi.azuredevops.PoolArgs;
 * import com.pulumi.azuredevops.Queue;
 * import com.pulumi.azuredevops.QueueArgs;
 * import com.pulumi.azuredevops.CheckBusinessHours;
 * import com.pulumi.azuredevops.CheckBusinessHoursArgs;
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
 *         var example = new Project("example", ProjectArgs.builder()
 *             .name("Example Project")
 *             .build());
 * 
 *         var examplePool = new Pool("examplePool", PoolArgs.builder()
 *             .name("example-pool")
 *             .build());
 * 
 *         var exampleQueue = new Queue("exampleQueue", QueueArgs.builder()
 *             .projectId(example.id())
 *             .agentPoolId(examplePool.id())
 *             .build());
 * 
 *         var exampleCheckBusinessHours = new CheckBusinessHours("exampleCheckBusinessHours", CheckBusinessHoursArgs.builder()
 *             .projectId(example.id())
 *             .displayName("Managed by Pulumi")
 *             .targetResourceId(exampleQueue.id())
 *             .targetResourceType("queue")
 *             .startTime("07:00")
 *             .endTime("15:30")
 *             .timeZone("UTC")
 *             .monday(true)
 *             .tuesday(true)
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ### Protect a repository
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.azuredevops.Project;
 * import com.pulumi.azuredevops.ProjectArgs;
 * import com.pulumi.azuredevops.Git;
 * import com.pulumi.azuredevops.GitArgs;
 * import com.pulumi.azuredevops.inputs.GitInitializationArgs;
 * import com.pulumi.azuredevops.CheckBusinessHours;
 * import com.pulumi.azuredevops.CheckBusinessHoursArgs;
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
 *         var example = new Project("example", ProjectArgs.builder()
 *             .name("Example Project")
 *             .build());
 * 
 *         var exampleGit = new Git("exampleGit", GitArgs.builder()
 *             .projectId(example.id())
 *             .name("Example Empty Git Repository")
 *             .initialization(GitInitializationArgs.builder()
 *                 .initType("Clean")
 *                 .build())
 *             .build());
 * 
 *         var exampleCheckBusinessHours = new CheckBusinessHours("exampleCheckBusinessHours", CheckBusinessHoursArgs.builder()
 *             .projectId(example.id())
 *             .displayName("Managed by Pulumi")
 *             .targetResourceId(Output.tuple(example.id(), exampleGit.id()).applyValue(values -> {
 *                 var exampleId = values.t1;
 *                 var exampleGitId = values.t2;
 *                 return String.format("%s.%s", exampleId,exampleGitId);
 *             }))
 *             .targetResourceType("repository")
 *             .startTime("07:00")
 *             .endTime("15:30")
 *             .timeZone("UTC")
 *             .monday(true)
 *             .tuesday(true)
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ### Protect a variable group
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.azuredevops.Project;
 * import com.pulumi.azuredevops.ProjectArgs;
 * import com.pulumi.azuredevops.VariableGroup;
 * import com.pulumi.azuredevops.VariableGroupArgs;
 * import com.pulumi.azuredevops.inputs.VariableGroupVariableArgs;
 * import com.pulumi.azuredevops.CheckBusinessHours;
 * import com.pulumi.azuredevops.CheckBusinessHoursArgs;
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
 *         var example = new Project("example", ProjectArgs.builder()
 *             .name("Example Project")
 *             .build());
 * 
 *         var exampleVariableGroup = new VariableGroup("exampleVariableGroup", VariableGroupArgs.builder()
 *             .projectId(example.id())
 *             .name("Example Variable Group")
 *             .description("Example Variable Group Description")
 *             .allowAccess(true)
 *             .variables(            
 *                 VariableGroupVariableArgs.builder()
 *                     .name("key1")
 *                     .value("val1")
 *                     .build(),
 *                 VariableGroupVariableArgs.builder()
 *                     .name("key2")
 *                     .secretValue("val2")
 *                     .isSecret(true)
 *                     .build())
 *             .build());
 * 
 *         var exampleCheckBusinessHours = new CheckBusinessHours("exampleCheckBusinessHours", CheckBusinessHoursArgs.builder()
 *             .projectId(example.id())
 *             .displayName("Managed by Pulumi")
 *             .targetResourceId(exampleVariableGroup.id())
 *             .targetResourceType("variablegroup")
 *             .startTime("07:00")
 *             .endTime("15:30")
 *             .timeZone("UTC")
 *             .monday(true)
 *             .tuesday(true)
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Relevant Links
 * 
 * - [Define approvals and checks](https://learn.microsoft.com/en-us/azure/devops/pipelines/process/approvals?view=azure-devops&amp;tabs=check-pass)
 * 
 * ## Supported Time Zones
 * 
 * - AUS Central Standard Time
 * - AUS Eastern Standard Time
 * - Afghanistan Standard Time
 * - Alaskan Standard Time
 * - Aleutian Standard Time
 * - Altai Standard Time
 * - Arab Standard Time
 * - Arabian Standard Time
 * - Arabic Standard Time
 * - Argentina Standard Time
 * - Astrakhan Standard Time
 * - Atlantic Standard Time
 * - Aus Central W. Standard Time
 * - Azerbaijan Standard Time
 * - Azores Standard Time
 * - Bahia Standard Time
 * - Bangladesh Standard Time
 * - Belarus Standard Time
 * - Bougainville Standard Time
 * - Canada Central Standard Time
 * - Cape Verde Standard Time
 * - Caucasus Standard Time
 * - Cen. Australia Standard Time
 * - Central America Standard Time
 * - Central Asia Standard Time
 * - Central Brazilian Standard Time
 * - Central Europe Standard Time
 * - Central European Standard Time
 * - Central Pacific Standard Time
 * - Central Standard Time (Mexico)
 * - Central Standard Time
 * - Chatham Islands Standard Time
 * - China Standard Time
 * - Cuba Standard Time
 * - Dateline Standard Time
 * - E. Africa Standard Time
 * - E. Australia Standard Time
 * - E. Europe Standard Time
 * - E. South America Standard Time
 * - Easter Island Standard Time
 * - Eastern Standard Time (Mexico)
 * - Eastern Standard Time
 * - Egypt Standard Time
 * - Ekaterinburg Standard Time
 * - FLE Standard Time
 * - Fiji Standard Time
 * - GMT Standard Time
 * - GTB Standard Time
 * - Georgian Standard Time
 * - Greenland Standard Time
 * - Greenwich Standard Time
 * - Haiti Standard Time
 * - Hawaiian Standard Time
 * - India Standard Time
 * - Iran Standard Time
 * - Israel Standard Time
 * - Jordan Standard Time
 * - Kaliningrad Standard Time
 * - Kamchatka Standard Time
 * - Korea Standard Time
 * - Libya Standard Time
 * - Line Islands Standard Time
 * - Lord Howe Standard Time
 * - Magadan Standard Time
 * - Magallanes Standard Time
 * - Marquesas Standard Time
 * - Mauritius Standard Time
 * - Mid-Atlantic Standard Time
 * - Middle East Standard Time
 * - Montevideo Standard Time
 * - Morocco Standard Time
 * - Mountain Standard Time (Mexico)
 * - Mountain Standard Time
 * - Myanmar Standard Time
 * - N. Central Asia Standard Time
 * - Namibia Standard Time
 * - Nepal Standard Time
 * - New Zealand Standard Time
 * - Newfoundland Standard Time
 * - Norfolk Standard Time
 * - North Asia East Standard Time
 * - North Asia Standard Time
 * - North Korea Standard Time
 * - Omsk Standard Time
 * - Pacific SA Standard Time
 * - Pacific Standard Time (Mexico)
 * - Pacific Standard Time
 * - Pakistan Standard Time
 * - Paraguay Standard Time
 * - Qyzylorda Standard Time
 * - Romance Standard Time
 * - Russia Time Zone 10
 * - Russia Time Zone 11
 * - Russia Time Zone 3
 * - Russian Standard Time
 * - SA Eastern Standard Time
 * - SA Pacific Standard Time
 * - SA Western Standard Time
 * - SE Asia Standard Time
 * - Saint Pierre Standard Time
 * - Sakhalin Standard Time
 * - Samoa Standard Time
 * - Sao Tome Standard Time
 * - Saratov Standard Time
 * - Singapore Standard Time
 * - South Africa Standard Time
 * - South Sudan Standard Time
 * - Sri Lanka Standard Time
 * - Sudan Standard Time
 * - Syria Standard Time
 * - Taipei Standard Time
 * - Tasmania Standard Time
 * - Tocantins Standard Time
 * - Tokyo Standard Time
 * - Tomsk Standard Time
 * - Tonga Standard Time
 * - Transbaikal Standard Time
 * - Turkey Standard Time
 * - Turks And Caicos Standard Time
 * - US Eastern Standard Time
 * - US Mountain Standard Time
 * - UTC
 * - UTC+12
 * - UTC+13
 * - UTC-02
 * - UTC-08
 * - UTC-09
 * - UTC-11
 * - Ulaanbaatar Standard Time
 * - Venezuela Standard Time
 * - Vladivostok Standard Time
 * - Volgograd Standard Time
 * - W. Australia Standard Time
 * - W. Central Africa Standard Time
 * - W. Europe Standard Time
 * - W. Mongolia Standard Time
 * - West Asia Standard Time
 * - West Bank Standard Time
 * - West Pacific Standard Time
 * - Yakutsk Standard Time
 * - Yukon Standard Time
 * 
 * ## Import
 * 
 * Importing this resource is not supported.
 * 
 */
@ResourceType(type="azuredevops:index/checkBusinessHours:CheckBusinessHours")
public class CheckBusinessHours extends com.pulumi.resources.CustomResource {
    /**
     * The name of the business hours check displayed in the web UI.
     * 
     */
    @Export(name="displayName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> displayName;

    /**
     * @return The name of the business hours check displayed in the web UI.
     * 
     */
    public Output<Optional<String>> displayName() {
        return Codegen.optional(this.displayName);
    }
    /**
     * The end of the time period that this check will be allowed to pass, specified as 24-hour time with leading zeros.
     * 
     */
    @Export(name="endTime", refs={String.class}, tree="[0]")
    private Output<String> endTime;

    /**
     * @return The end of the time period that this check will be allowed to pass, specified as 24-hour time with leading zeros.
     * 
     */
    public Output<String> endTime() {
        return this.endTime;
    }
    /**
     * This check will pass on Fridays. Defaults to `false`.
     * 
     */
    @Export(name="friday", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> friday;

    /**
     * @return This check will pass on Fridays. Defaults to `false`.
     * 
     */
    public Output<Optional<Boolean>> friday() {
        return Codegen.optional(this.friday);
    }
    /**
     * This check will pass on Mondays. Defaults to `false`.
     * 
     */
    @Export(name="monday", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> monday;

    /**
     * @return This check will pass on Mondays. Defaults to `false`.
     * 
     */
    public Output<Optional<Boolean>> monday() {
        return Codegen.optional(this.monday);
    }
    /**
     * The project ID.
     * 
     */
    @Export(name="projectId", refs={String.class}, tree="[0]")
    private Output<String> projectId;

    /**
     * @return The project ID.
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }
    /**
     * This check will pass on Saturdays. Defaults to `false`.
     * 
     */
    @Export(name="saturday", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> saturday;

    /**
     * @return This check will pass on Saturdays. Defaults to `false`.
     * 
     */
    public Output<Optional<Boolean>> saturday() {
        return Codegen.optional(this.saturday);
    }
    /**
     * The beginning of the time period that this check will be allowed to pass, specified as 24-hour time with leading zeros.
     * 
     */
    @Export(name="startTime", refs={String.class}, tree="[0]")
    private Output<String> startTime;

    /**
     * @return The beginning of the time period that this check will be allowed to pass, specified as 24-hour time with leading zeros.
     * 
     */
    public Output<String> startTime() {
        return this.startTime;
    }
    /**
     * This check will pass on Sundays. Defaults to `false`.
     * 
     */
    @Export(name="sunday", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> sunday;

    /**
     * @return This check will pass on Sundays. Defaults to `false`.
     * 
     */
    public Output<Optional<Boolean>> sunday() {
        return Codegen.optional(this.sunday);
    }
    /**
     * The ID of the resource being protected by the check.
     * 
     */
    @Export(name="targetResourceId", refs={String.class}, tree="[0]")
    private Output<String> targetResourceId;

    /**
     * @return The ID of the resource being protected by the check.
     * 
     */
    public Output<String> targetResourceId() {
        return this.targetResourceId;
    }
    /**
     * The type of resource being protected by the check. Possible values are: `endpoint`, `environment`, `queue`, `repository`, `securefile`, `variablegroup`.
     * 
     */
    @Export(name="targetResourceType", refs={String.class}, tree="[0]")
    private Output<String> targetResourceType;

    /**
     * @return The type of resource being protected by the check. Possible values are: `endpoint`, `environment`, `queue`, `repository`, `securefile`, `variablegroup`.
     * 
     */
    public Output<String> targetResourceType() {
        return this.targetResourceType;
    }
    /**
     * This check will pass on Thursdays. Defaults to `false`.
     * 
     */
    @Export(name="thursday", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> thursday;

    /**
     * @return This check will pass on Thursdays. Defaults to `false`.
     * 
     */
    public Output<Optional<Boolean>> thursday() {
        return Codegen.optional(this.thursday);
    }
    /**
     * The time zone this check will be evaluated in. See below for supported values.
     * 
     */
    @Export(name="timeZone", refs={String.class}, tree="[0]")
    private Output<String> timeZone;

    /**
     * @return The time zone this check will be evaluated in. See below for supported values.
     * 
     */
    public Output<String> timeZone() {
        return this.timeZone;
    }
    /**
     * The timeout in minutes for the business hours check. Defaults to `1440`.
     * 
     */
    @Export(name="timeout", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> timeout;

    /**
     * @return The timeout in minutes for the business hours check. Defaults to `1440`.
     * 
     */
    public Output<Optional<Integer>> timeout() {
        return Codegen.optional(this.timeout);
    }
    /**
     * This check will pass on Tuesday. Defaults to `false`.
     * 
     */
    @Export(name="tuesday", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> tuesday;

    /**
     * @return This check will pass on Tuesday. Defaults to `false`.
     * 
     */
    public Output<Optional<Boolean>> tuesday() {
        return Codegen.optional(this.tuesday);
    }
    /**
     * The version of the check.
     * 
     */
    @Export(name="version", refs={Integer.class}, tree="[0]")
    private Output<Integer> version;

    /**
     * @return The version of the check.
     * 
     */
    public Output<Integer> version() {
        return this.version;
    }
    /**
     * This check will pass on Wednesdays. Defaults to `false`.
     * 
     */
    @Export(name="wednesday", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> wednesday;

    /**
     * @return This check will pass on Wednesdays. Defaults to `false`.
     * 
     */
    public Output<Optional<Boolean>> wednesday() {
        return Codegen.optional(this.wednesday);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public CheckBusinessHours(java.lang.String name) {
        this(name, CheckBusinessHoursArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public CheckBusinessHours(java.lang.String name, CheckBusinessHoursArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public CheckBusinessHours(java.lang.String name, CheckBusinessHoursArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/checkBusinessHours:CheckBusinessHours", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private CheckBusinessHours(java.lang.String name, Output<java.lang.String> id, @Nullable CheckBusinessHoursState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/checkBusinessHours:CheckBusinessHours", name, state, makeResourceOptions(options, id), false);
    }

    private static CheckBusinessHoursArgs makeArgs(CheckBusinessHoursArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? CheckBusinessHoursArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
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
    public static CheckBusinessHours get(java.lang.String name, Output<java.lang.String> id, @Nullable CheckBusinessHoursState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new CheckBusinessHours(name, id, state, options);
    }
}
