// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops;

import com.pulumi.azuredevops.GroupEntitlementArgs;
import com.pulumi.azuredevops.Utilities;
import com.pulumi.azuredevops.inputs.GroupEntitlementState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a group entitlement within Azure DevOps.
 * 
 * ## Example Usage
 * 
 * ### With an Azure DevOps local group managed by this resource
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.azuredevops.GroupEntitlement;
 * import com.pulumi.azuredevops.GroupEntitlementArgs;
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
 *         var example = new GroupEntitlement("example", GroupEntitlementArgs.builder()
 *             .displayName("Group Name")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ### With group origin ID
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.azuredevops.GroupEntitlement;
 * import com.pulumi.azuredevops.GroupEntitlementArgs;
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
 *         var example = new GroupEntitlement("example", GroupEntitlementArgs.builder()
 *             .origin("aad")
 *             .originId("00000000-0000-0000-0000-000000000000")
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
 * - [Azure DevOps Service REST API 7.0 - Group Entitlements](https://learn.microsoft.com/en-us/rest/api/azure/devops/memberentitlementmanagement/group-entitlements?view=azure-devops-rest-7.1)
 * - [Programmatic mapping of access levels](https://docs.microsoft.com/en-us/azure/devops/organizations/security/access-levels?view=azure-devops#programmatic-mapping-of-access-levels)
 * 
 * ## PAT Permissions Required
 * 
 * - **Member Entitlement Management**: Read &amp; Write
 * 
 * ## Import
 * 
 * The resource allows the import via the ID of a group entitlement, which is a UUID.
 * 
 * ```sh
 * $ pulumi import azuredevops:index/groupEntitlement:GroupEntitlement example 00000000-0000-0000-0000-000000000000
 * ```
 * 
 */
@ResourceType(type="azuredevops:index/groupEntitlement:GroupEntitlement")
public class GroupEntitlement extends com.pulumi.resources.CustomResource {
    /**
     * Type of Account License. Possible values are: `advanced`, `earlyAdopter`, `express`, `none`, `professional`, or `stakeholder`. Defaults to `express`. In addition, the value `basic` is allowed which is an alias for `express` and reflects the name of the `express` license used in the Azure DevOps web interface.
     * 
     */
    @Export(name="accountLicenseType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> accountLicenseType;

    /**
     * @return Type of Account License. Possible values are: `advanced`, `earlyAdopter`, `express`, `none`, `professional`, or `stakeholder`. Defaults to `express`. In addition, the value `basic` is allowed which is an alias for `express` and reflects the name of the `express` license used in the Azure DevOps web interface.
     * 
     */
    public Output<Optional<String>> accountLicenseType() {
        return Codegen.optional(this.accountLicenseType);
    }
    /**
     * The descriptor is the primary way to reference the graph subject while the system is running. This field will uniquely identify the group graph subject.
     * 
     */
    @Export(name="descriptor", refs={String.class}, tree="[0]")
    private Output<String> descriptor;

    /**
     * @return The descriptor is the primary way to reference the graph subject while the system is running. This field will uniquely identify the group graph subject.
     * 
     */
    public Output<String> descriptor() {
        return this.descriptor;
    }
    /**
     * The display name is the name used in Azure DevOps UI. Cannot be set together with `origin_id` and `origin`.
     * 
     */
    @Export(name="displayName", refs={String.class}, tree="[0]")
    private Output<String> displayName;

    /**
     * @return The display name is the name used in Azure DevOps UI. Cannot be set together with `origin_id` and `origin`.
     * 
     */
    public Output<String> displayName() {
        return this.displayName;
    }
    /**
     * The source of the licensing (e.g. Account. MSDN etc.). Possible values are: `account`, `auto`, `msdn`, `none`, `profile`, `trial`. Defaults to `account`.
     * 
     * &gt; **NOTE:** A existing group in Azure AD can only be referenced by the combination of `origin_id` and `origin`.
     * 
     */
    @Export(name="licensingSource", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> licensingSource;

    /**
     * @return The source of the licensing (e.g. Account. MSDN etc.). Possible values are: `account`, `auto`, `msdn`, `none`, `profile`, `trial`. Defaults to `account`.
     * 
     * &gt; **NOTE:** A existing group in Azure AD can only be referenced by the combination of `origin_id` and `origin`.
     * 
     */
    public Output<Optional<String>> licensingSource() {
        return Codegen.optional(this.licensingSource);
    }
    /**
     * The type of source provider for the origin identifier.
     * 
     */
    @Export(name="origin", refs={String.class}, tree="[0]")
    private Output<String> origin;

    /**
     * @return The type of source provider for the origin identifier.
     * 
     */
    public Output<String> origin() {
        return this.origin;
    }
    /**
     * The unique identifier from the system of origin. Typically, a sid, object id or Guid. e.g. Used for member of other tenant on Azure Active Directory.
     * 
     */
    @Export(name="originId", refs={String.class}, tree="[0]")
    private Output<String> originId;

    /**
     * @return The unique identifier from the system of origin. Typically, a sid, object id or Guid. e.g. Used for member of other tenant on Azure Active Directory.
     * 
     */
    public Output<String> originId() {
        return this.originId;
    }
    /**
     * The principal name of a graph member on Azure DevOps
     * 
     */
    @Export(name="principalName", refs={String.class}, tree="[0]")
    private Output<String> principalName;

    /**
     * @return The principal name of a graph member on Azure DevOps
     * 
     */
    public Output<String> principalName() {
        return this.principalName;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public GroupEntitlement(java.lang.String name) {
        this(name, GroupEntitlementArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public GroupEntitlement(java.lang.String name, @Nullable GroupEntitlementArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public GroupEntitlement(java.lang.String name, @Nullable GroupEntitlementArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/groupEntitlement:GroupEntitlement", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private GroupEntitlement(java.lang.String name, Output<java.lang.String> id, @Nullable GroupEntitlementState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/groupEntitlement:GroupEntitlement", name, state, makeResourceOptions(options, id), false);
    }

    private static GroupEntitlementArgs makeArgs(@Nullable GroupEntitlementArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? GroupEntitlementArgs.Empty : args;
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
    public static GroupEntitlement get(java.lang.String name, Output<java.lang.String> id, @Nullable GroupEntitlementState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new GroupEntitlement(name, id, state, options);
    }
}
