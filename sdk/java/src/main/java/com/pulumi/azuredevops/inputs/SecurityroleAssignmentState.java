// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SecurityroleAssignmentState extends com.pulumi.resources.ResourceArgs {

    public static final SecurityroleAssignmentState Empty = new SecurityroleAssignmentState();

    /**
     * The ID of the identity to authorize.
     * 
     */
    @Import(name="identityId")
    private @Nullable Output<String> identityId;

    /**
     * @return The ID of the identity to authorize.
     * 
     */
    public Optional<Output<String>> identityId() {
        return Optional.ofNullable(this.identityId);
    }

    /**
     * The ID of the resource on which the role is to be assigned. Changing this forces a new resource to be created.
     * 
     */
    @Import(name="resourceId")
    private @Nullable Output<String> resourceId;

    /**
     * @return The ID of the resource on which the role is to be assigned. Changing this forces a new resource to be created.
     * 
     */
    public Optional<Output<String>> resourceId() {
        return Optional.ofNullable(this.resourceId);
    }

    /**
     * Name of the role to assign.
     * 
     */
    @Import(name="roleName")
    private @Nullable Output<String> roleName;

    /**
     * @return Name of the role to assign.
     * 
     */
    public Optional<Output<String>> roleName() {
        return Optional.ofNullable(this.roleName);
    }

    /**
     * The scope in which this assignment should exist.
     * 
     */
    @Import(name="scope")
    private @Nullable Output<String> scope;

    /**
     * @return The scope in which this assignment should exist.
     * 
     */
    public Optional<Output<String>> scope() {
        return Optional.ofNullable(this.scope);
    }

    private SecurityroleAssignmentState() {}

    private SecurityroleAssignmentState(SecurityroleAssignmentState $) {
        this.identityId = $.identityId;
        this.resourceId = $.resourceId;
        this.roleName = $.roleName;
        this.scope = $.scope;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SecurityroleAssignmentState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SecurityroleAssignmentState $;

        public Builder() {
            $ = new SecurityroleAssignmentState();
        }

        public Builder(SecurityroleAssignmentState defaults) {
            $ = new SecurityroleAssignmentState(Objects.requireNonNull(defaults));
        }

        /**
         * @param identityId The ID of the identity to authorize.
         * 
         * @return builder
         * 
         */
        public Builder identityId(@Nullable Output<String> identityId) {
            $.identityId = identityId;
            return this;
        }

        /**
         * @param identityId The ID of the identity to authorize.
         * 
         * @return builder
         * 
         */
        public Builder identityId(String identityId) {
            return identityId(Output.of(identityId));
        }

        /**
         * @param resourceId The ID of the resource on which the role is to be assigned. Changing this forces a new resource to be created.
         * 
         * @return builder
         * 
         */
        public Builder resourceId(@Nullable Output<String> resourceId) {
            $.resourceId = resourceId;
            return this;
        }

        /**
         * @param resourceId The ID of the resource on which the role is to be assigned. Changing this forces a new resource to be created.
         * 
         * @return builder
         * 
         */
        public Builder resourceId(String resourceId) {
            return resourceId(Output.of(resourceId));
        }

        /**
         * @param roleName Name of the role to assign.
         * 
         * @return builder
         * 
         */
        public Builder roleName(@Nullable Output<String> roleName) {
            $.roleName = roleName;
            return this;
        }

        /**
         * @param roleName Name of the role to assign.
         * 
         * @return builder
         * 
         */
        public Builder roleName(String roleName) {
            return roleName(Output.of(roleName));
        }

        /**
         * @param scope The scope in which this assignment should exist.
         * 
         * @return builder
         * 
         */
        public Builder scope(@Nullable Output<String> scope) {
            $.scope = scope;
            return this;
        }

        /**
         * @param scope The scope in which this assignment should exist.
         * 
         * @return builder
         * 
         */
        public Builder scope(String scope) {
            return scope(Output.of(scope));
        }

        public SecurityroleAssignmentState build() {
            return $;
        }
    }

}
