// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class SecurityroleAssignmentArgs extends com.pulumi.resources.ResourceArgs {

    public static final SecurityroleAssignmentArgs Empty = new SecurityroleAssignmentArgs();

    /**
     * The ID of the identity to authorize.
     * 
     */
    @Import(name="identityId", required=true)
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
    @Import(name="resourceId", required=true)
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
    @Import(name="roleName", required=true)
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
    @Import(name="scope", required=true)
    private Output<String> scope;

    /**
     * @return The scope in which this assignment should exist.
     * 
     */
    public Output<String> scope() {
        return this.scope;
    }

    private SecurityroleAssignmentArgs() {}

    private SecurityroleAssignmentArgs(SecurityroleAssignmentArgs $) {
        this.identityId = $.identityId;
        this.resourceId = $.resourceId;
        this.roleName = $.roleName;
        this.scope = $.scope;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SecurityroleAssignmentArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SecurityroleAssignmentArgs $;

        public Builder() {
            $ = new SecurityroleAssignmentArgs();
        }

        public Builder(SecurityroleAssignmentArgs defaults) {
            $ = new SecurityroleAssignmentArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param identityId The ID of the identity to authorize.
         * 
         * @return builder
         * 
         */
        public Builder identityId(Output<String> identityId) {
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
         * @param resourceId The ID of the resource on which the role is to be assigned.
         * 
         * @return builder
         * 
         */
        public Builder resourceId(Output<String> resourceId) {
            $.resourceId = resourceId;
            return this;
        }

        /**
         * @param resourceId The ID of the resource on which the role is to be assigned.
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
        public Builder roleName(Output<String> roleName) {
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
        public Builder scope(Output<String> scope) {
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

        public SecurityroleAssignmentArgs build() {
            if ($.identityId == null) {
                throw new MissingRequiredPropertyException("SecurityroleAssignmentArgs", "identityId");
            }
            if ($.resourceId == null) {
                throw new MissingRequiredPropertyException("SecurityroleAssignmentArgs", "resourceId");
            }
            if ($.roleName == null) {
                throw new MissingRequiredPropertyException("SecurityroleAssignmentArgs", "roleName");
            }
            if ($.scope == null) {
                throw new MissingRequiredPropertyException("SecurityroleAssignmentArgs", "scope");
            }
            return $;
        }
    }

}