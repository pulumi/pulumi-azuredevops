// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops;

import com.pulumi.azuredevops.ServicePrincipalEntitlementArgs;
import com.pulumi.azuredevops.Utilities;
import com.pulumi.azuredevops.inputs.ServicePrincipalEntitlementState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a Service Principal Entitlement.
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.azuredevops.ServicePrincipalEntitlement;
 * import com.pulumi.azuredevops.ServicePrincipalEntitlementArgs;
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
 *         var example = new ServicePrincipalEntitlement("example", ServicePrincipalEntitlementArgs.builder()
 *             .originId("00000000-0000-0000-0000-000000000000")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Service Principal Entitlements can be imported using the `resource id`.
 * 
 * The `resource id` can be found using DEV Tools in the `Users` section of the ADO organization.
 * 
 * ```sh
 * $ pulumi import azuredevops:index/servicePrincipalEntitlement:ServicePrincipalEntitlement example 8480c6eb-ce60-47e9-88df-eca3c801638b
 * ```
 * 
 */
@ResourceType(type="azuredevops:index/servicePrincipalEntitlement:ServicePrincipalEntitlement")
public class ServicePrincipalEntitlement extends com.pulumi.resources.CustomResource {
    /**
     * Type of Account License. Possible values are: `advanced`, `earlyAdopter`, `express`, `none`, `professional`, or `stakeholder`. Defaults to `express`. In addition the value `basic` is allowed which is an alias for `express` and reflects the name of the `express` license used in the Azure DevOps web interface.
     * 
     */
    @Export(name="accountLicenseType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> accountLicenseType;

    /**
     * @return Type of Account License. Possible values are: `advanced`, `earlyAdopter`, `express`, `none`, `professional`, or `stakeholder`. Defaults to `express`. In addition the value `basic` is allowed which is an alias for `express` and reflects the name of the `express` license used in the Azure DevOps web interface.
     * 
     */
    public Output<Optional<String>> accountLicenseType() {
        return Codegen.optional(this.accountLicenseType);
    }
    /**
     * The descriptor is the primary way to reference the graph subject while the system is running. This field will uniquely identify the user graph subject.
     * 
     */
    @Export(name="descriptor", refs={String.class}, tree="[0]")
    private Output<String> descriptor;

    /**
     * @return The descriptor is the primary way to reference the graph subject while the system is running. This field will uniquely identify the user graph subject.
     * 
     */
    public Output<String> descriptor() {
        return this.descriptor;
    }
    /**
     * The display name of service principal.
     * 
     */
    @Export(name="displayName", refs={String.class}, tree="[0]")
    private Output<String> displayName;

    /**
     * @return The display name of service principal.
     * 
     */
    public Output<String> displayName() {
        return this.displayName;
    }
    /**
     * The source of the licensing (e.g. Account. MSDN etc.) Possible values are: `account` (Default), `auto`, `msdn`, `none`, `profile`, `trial`
     * 
     */
    @Export(name="licensingSource", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> licensingSource;

    /**
     * @return The source of the licensing (e.g. Account. MSDN etc.) Possible values are: `account` (Default), `auto`, `msdn`, `none`, `profile`, `trial`
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
     * The Object ID of the service principal in Entra ID. Changing this forces a new Service Principal Entitlement to be created.
     * 
     */
    @Export(name="originId", refs={String.class}, tree="[0]")
    private Output<String> originId;

    /**
     * @return The Object ID of the service principal in Entra ID. Changing this forces a new Service Principal Entitlement to be created.
     * 
     */
    public Output<String> originId() {
        return this.originId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ServicePrincipalEntitlement(java.lang.String name) {
        this(name, ServicePrincipalEntitlementArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ServicePrincipalEntitlement(java.lang.String name, ServicePrincipalEntitlementArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ServicePrincipalEntitlement(java.lang.String name, ServicePrincipalEntitlementArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/servicePrincipalEntitlement:ServicePrincipalEntitlement", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private ServicePrincipalEntitlement(java.lang.String name, Output<java.lang.String> id, @Nullable ServicePrincipalEntitlementState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/servicePrincipalEntitlement:ServicePrincipalEntitlement", name, state, makeResourceOptions(options, id), false);
    }

    private static ServicePrincipalEntitlementArgs makeArgs(ServicePrincipalEntitlementArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? ServicePrincipalEntitlementArgs.Empty : args;
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
    public static ServicePrincipalEntitlement get(java.lang.String name, Output<java.lang.String> id, @Nullable ServicePrincipalEntitlementState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ServicePrincipalEntitlement(name, id, state, options);
    }
}
