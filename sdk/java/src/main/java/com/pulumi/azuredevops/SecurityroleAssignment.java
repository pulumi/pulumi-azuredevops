// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops;

import com.pulumi.azuredevops.SecurityroleAssignmentArgs;
import com.pulumi.azuredevops.Utilities;
import com.pulumi.azuredevops.inputs.SecurityroleAssignmentState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Manages assignment of security roles to various resources within Azure DevOps organization.
 * 
 */
@ResourceType(type="azuredevops:index/securityroleAssignment:SecurityroleAssignment")
public class SecurityroleAssignment extends com.pulumi.resources.CustomResource {
    /**
     * The ID of the identity to authorize.
     * 
     */
    @Export(name="identityId", refs={String.class}, tree="[0]")
    private Output<String> identityId;

    /**
     * @return The ID of the identity to authorize.
     * 
     */
    public Output<String> identityId() {
        return this.identityId;
    }
    /**
     * The ID of the resource on which the role is to be assigned.
     * 
     */
    @Export(name="resourceId", refs={String.class}, tree="[0]")
    private Output<String> resourceId;

    /**
     * @return The ID of the resource on which the role is to be assigned.
     * 
     */
    public Output<String> resourceId() {
        return this.resourceId;
    }
    /**
     * Name of the role to assign.
     * 
     */
    @Export(name="roleName", refs={String.class}, tree="[0]")
    private Output<String> roleName;

    /**
     * @return Name of the role to assign.
     * 
     */
    public Output<String> roleName() {
        return this.roleName;
    }
    /**
     * The scope in which this assignment should exist.
     * 
     */
    @Export(name="scope", refs={String.class}, tree="[0]")
    private Output<String> scope;

    /**
     * @return The scope in which this assignment should exist.
     * 
     */
    public Output<String> scope() {
        return this.scope;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public SecurityroleAssignment(String name) {
        this(name, SecurityroleAssignmentArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public SecurityroleAssignment(String name, SecurityroleAssignmentArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public SecurityroleAssignment(String name, SecurityroleAssignmentArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/securityroleAssignment:SecurityroleAssignment", name, args == null ? SecurityroleAssignmentArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private SecurityroleAssignment(String name, Output<String> id, @Nullable SecurityroleAssignmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/securityroleAssignment:SecurityroleAssignment", name, state, makeResourceOptions(options, id));
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
    public static SecurityroleAssignment get(String name, Output<String> id, @Nullable SecurityroleAssignmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new SecurityroleAssignment(name, id, state, options);
    }
}