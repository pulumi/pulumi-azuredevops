// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops;

import com.pulumi.azuredevops.UserArgs;
import com.pulumi.azuredevops.Utilities;
import com.pulumi.azuredevops.inputs.UserState;
import com.pulumi.core.Alias;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a user entitlement within Azure DevOps.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.azuredevops.User;
 * import com.pulumi.azuredevops.UserArgs;
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
 *         var example = new User(&#34;example&#34;, UserArgs.builder()        
 *             .principalName(&#34;foo@contoso.com&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ## Relevant Links
 * 
 * - [Azure DevOps Service REST API 7.0 - User Entitlements - Add](https://docs.microsoft.com/en-us/rest/api/azure/devops/memberentitlementmanagement/user-entitlements/add?view=azure-devops-rest-7.0)
 * - [Programmatic mapping of access levels](https://docs.microsoft.com/en-us/azure/devops/organizations/security/access-levels?view=azure-devops#programmatic-mapping-of-access-levels)
 * 
 * ## PAT Permissions Required
 * 
 * - **Member Entitlement Management**: Read &amp; Write
 * 
 * ## Import
 * 
 * The resources allows the import via the UUID of a user entitlement or by using the principal name of a user owning an entitlement.
 * 
 */
@ResourceType(type="azuredevops:index/user:User")
public class User extends com.pulumi.resources.CustomResource {
    /**
     * Type of Account License. Valid values: `advanced`, `earlyAdopter`, `express`, `none`, `professional`, or `stakeholder`. Defaults to `express`. In addition the value `basic` is allowed which is an alias for `express` and reflects the name of the `express` license used in the Azure DevOps web interface.
     * 
     */
    @Export(name="accountLicenseType", type=String.class, parameters={})
    private Output</* @Nullable */ String> accountLicenseType;

    /**
     * @return Type of Account License. Valid values: `advanced`, `earlyAdopter`, `express`, `none`, `professional`, or `stakeholder`. Defaults to `express`. In addition the value `basic` is allowed which is an alias for `express` and reflects the name of the `express` license used in the Azure DevOps web interface.
     * 
     */
    public Output<Optional<String>> accountLicenseType() {
        return Codegen.optional(this.accountLicenseType);
    }
    /**
     * The descriptor is the primary way to reference the graph subject while the system is running. This field will uniquely identify the user graph subject.
     * 
     */
    @Export(name="descriptor", type=String.class, parameters={})
    private Output<String> descriptor;

    /**
     * @return The descriptor is the primary way to reference the graph subject while the system is running. This field will uniquely identify the user graph subject.
     * 
     */
    public Output<String> descriptor() {
        return this.descriptor;
    }
    /**
     * The source of the licensing (e.g. Account. MSDN etc.) Valid values: `account` (Default), `auto`, `msdn`, `none`, `profile`, `trial`
     * 
     * &gt; **NOTE:** A user can only be referenced by it&#39;s `principal_name` or by the combination of `origin_id` and `origin`.
     * 
     */
    @Export(name="licensingSource", type=String.class, parameters={})
    private Output</* @Nullable */ String> licensingSource;

    /**
     * @return The source of the licensing (e.g. Account. MSDN etc.) Valid values: `account` (Default), `auto`, `msdn`, `none`, `profile`, `trial`
     * 
     * &gt; **NOTE:** A user can only be referenced by it&#39;s `principal_name` or by the combination of `origin_id` and `origin`.
     * 
     */
    public Output<Optional<String>> licensingSource() {
        return Codegen.optional(this.licensingSource);
    }
    /**
     * The type of source provider for the origin identifier.
     * 
     */
    @Export(name="origin", type=String.class, parameters={})
    private Output<String> origin;

    /**
     * @return The type of source provider for the origin identifier.
     * 
     */
    public Output<String> origin() {
        return this.origin;
    }
    /**
     * The unique identifier from the system of origin. Typically a sid, object id or Guid. e.g. Used for member of other tenant on Azure Active Directory.
     * 
     */
    @Export(name="originId", type=String.class, parameters={})
    private Output<String> originId;

    /**
     * @return The unique identifier from the system of origin. Typically a sid, object id or Guid. e.g. Used for member of other tenant on Azure Active Directory.
     * 
     */
    public Output<String> originId() {
        return this.originId;
    }
    /**
     * The principal name is the PrincipalName of a graph member from the source provider. Usually, e-mail address.
     * 
     */
    @Export(name="principalName", type=String.class, parameters={})
    private Output<String> principalName;

    /**
     * @return The principal name is the PrincipalName of a graph member from the source provider. Usually, e-mail address.
     * 
     */
    public Output<String> principalName() {
        return this.principalName;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public User(String name) {
        this(name, UserArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public User(String name, @Nullable UserArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public User(String name, @Nullable UserArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/user:User", name, args == null ? UserArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private User(String name, Output<String> id, @Nullable UserState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/user:User", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .aliases(List.of(
                Output.of(Alias.builder().type("azuredevops:Entitlement/user:User").build())
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
    public static User get(String name, Output<String> id, @Nullable UserState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new User(name, id, state, options);
    }
}
