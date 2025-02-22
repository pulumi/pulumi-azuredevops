// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AreaPermissionsState extends com.pulumi.resources.ResourceArgs {

    public static final AreaPermissionsState Empty = new AreaPermissionsState();

    /**
     * The name of the branch to assign the permissions.
     * 
     */
    @Import(name="path")
    private @Nullable Output<String> path;

    /**
     * @return The name of the branch to assign the permissions.
     * 
     */
    public Optional<Output<String>> path() {
        return Optional.ofNullable(this.path);
    }

    /**
     * the permissions to assign. The following permissions are available.
     * 
     * | Permission             | Description                          |
     * |------------------------|--------------------------------------|
     * | GENERIC_READ           | View permissions for this node       |
     * | GENERIC_WRITE          | Edit this node                       |
     * | CREATE_CHILDREN        | Create child nodes                   |
     * | DELETE                 | Delete this node                     |
     * | WORK_ITEM_READ         | View work items in this node         |
     * | WORK_ITEM_WRITE        | Edit work items in this node         |
     * | MANAGE_TEST_PLANS      | Manage test plans                    |
     * | MANAGE_TEST_SUITES     | Manage test suites                   |
     * | WORK_ITEM_SAVE_COMMENT | Edit work item comments in this node |
     * 
     */
    @Import(name="permissions")
    private @Nullable Output<Map<String,String>> permissions;

    /**
     * @return the permissions to assign. The following permissions are available.
     * 
     * | Permission             | Description                          |
     * |------------------------|--------------------------------------|
     * | GENERIC_READ           | View permissions for this node       |
     * | GENERIC_WRITE          | Edit this node                       |
     * | CREATE_CHILDREN        | Create child nodes                   |
     * | DELETE                 | Delete this node                     |
     * | WORK_ITEM_READ         | View work items in this node         |
     * | WORK_ITEM_WRITE        | Edit work items in this node         |
     * | MANAGE_TEST_PLANS      | Manage test plans                    |
     * | MANAGE_TEST_SUITES     | Manage test suites                   |
     * | WORK_ITEM_SAVE_COMMENT | Edit work item comments in this node |
     * 
     */
    public Optional<Output<Map<String,String>>> permissions() {
        return Optional.ofNullable(this.permissions);
    }

    /**
     * The **group** principal to assign the permissions.
     * 
     */
    @Import(name="principal")
    private @Nullable Output<String> principal;

    /**
     * @return The **group** principal to assign the permissions.
     * 
     */
    public Optional<Output<String>> principal() {
        return Optional.ofNullable(this.principal);
    }

    /**
     * The ID of the project to assign the permissions.
     * 
     */
    @Import(name="projectId")
    private @Nullable Output<String> projectId;

    /**
     * @return The ID of the project to assign the permissions.
     * 
     */
    public Optional<Output<String>> projectId() {
        return Optional.ofNullable(this.projectId);
    }

    /**
     * Replace (`true`) or merge (`false`) the permissions. Default: `true`.
     * 
     */
    @Import(name="replace")
    private @Nullable Output<Boolean> replace;

    /**
     * @return Replace (`true`) or merge (`false`) the permissions. Default: `true`.
     * 
     */
    public Optional<Output<Boolean>> replace() {
        return Optional.ofNullable(this.replace);
    }

    private AreaPermissionsState() {}

    private AreaPermissionsState(AreaPermissionsState $) {
        this.path = $.path;
        this.permissions = $.permissions;
        this.principal = $.principal;
        this.projectId = $.projectId;
        this.replace = $.replace;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AreaPermissionsState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AreaPermissionsState $;

        public Builder() {
            $ = new AreaPermissionsState();
        }

        public Builder(AreaPermissionsState defaults) {
            $ = new AreaPermissionsState(Objects.requireNonNull(defaults));
        }

        /**
         * @param path The name of the branch to assign the permissions.
         * 
         * @return builder
         * 
         */
        public Builder path(@Nullable Output<String> path) {
            $.path = path;
            return this;
        }

        /**
         * @param path The name of the branch to assign the permissions.
         * 
         * @return builder
         * 
         */
        public Builder path(String path) {
            return path(Output.of(path));
        }

        /**
         * @param permissions the permissions to assign. The following permissions are available.
         * 
         * | Permission             | Description                          |
         * |------------------------|--------------------------------------|
         * | GENERIC_READ           | View permissions for this node       |
         * | GENERIC_WRITE          | Edit this node                       |
         * | CREATE_CHILDREN        | Create child nodes                   |
         * | DELETE                 | Delete this node                     |
         * | WORK_ITEM_READ         | View work items in this node         |
         * | WORK_ITEM_WRITE        | Edit work items in this node         |
         * | MANAGE_TEST_PLANS      | Manage test plans                    |
         * | MANAGE_TEST_SUITES     | Manage test suites                   |
         * | WORK_ITEM_SAVE_COMMENT | Edit work item comments in this node |
         * 
         * @return builder
         * 
         */
        public Builder permissions(@Nullable Output<Map<String,String>> permissions) {
            $.permissions = permissions;
            return this;
        }

        /**
         * @param permissions the permissions to assign. The following permissions are available.
         * 
         * | Permission             | Description                          |
         * |------------------------|--------------------------------------|
         * | GENERIC_READ           | View permissions for this node       |
         * | GENERIC_WRITE          | Edit this node                       |
         * | CREATE_CHILDREN        | Create child nodes                   |
         * | DELETE                 | Delete this node                     |
         * | WORK_ITEM_READ         | View work items in this node         |
         * | WORK_ITEM_WRITE        | Edit work items in this node         |
         * | MANAGE_TEST_PLANS      | Manage test plans                    |
         * | MANAGE_TEST_SUITES     | Manage test suites                   |
         * | WORK_ITEM_SAVE_COMMENT | Edit work item comments in this node |
         * 
         * @return builder
         * 
         */
        public Builder permissions(Map<String,String> permissions) {
            return permissions(Output.of(permissions));
        }

        /**
         * @param principal The **group** principal to assign the permissions.
         * 
         * @return builder
         * 
         */
        public Builder principal(@Nullable Output<String> principal) {
            $.principal = principal;
            return this;
        }

        /**
         * @param principal The **group** principal to assign the permissions.
         * 
         * @return builder
         * 
         */
        public Builder principal(String principal) {
            return principal(Output.of(principal));
        }

        /**
         * @param projectId The ID of the project to assign the permissions.
         * 
         * @return builder
         * 
         */
        public Builder projectId(@Nullable Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param projectId The ID of the project to assign the permissions.
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
        }

        /**
         * @param replace Replace (`true`) or merge (`false`) the permissions. Default: `true`.
         * 
         * @return builder
         * 
         */
        public Builder replace(@Nullable Output<Boolean> replace) {
            $.replace = replace;
            return this;
        }

        /**
         * @param replace Replace (`true`) or merge (`false`) the permissions. Default: `true`.
         * 
         * @return builder
         * 
         */
        public Builder replace(Boolean replace) {
            return replace(Output.of(replace));
        }

        public AreaPermissionsState build() {
            return $;
        }
    }

}
