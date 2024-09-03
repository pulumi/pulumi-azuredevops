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


public final class ProjectPermissionsState extends com.pulumi.resources.ResourceArgs {

    public static final ProjectPermissionsState Empty = new ProjectPermissionsState();

    /**
     * the permissions to assign. The following permissions are available
     * 
     * | Permission  |             Description              |
     * |-------------|--------------------------------------|
     * | GENERIC_    | View project-level                   |
     * | GENERIC_    | Edit project-level                   |
     * | DELETE      | Delete team                          |
     * | PUBLISH_    | Create test                          |
     * | ADMINISTER_ | Administer a                         |
     * | START_      | Start a                              |
     * | EDIT_       | Edit build                           |
     * | UPDATE_     | Write to build operational           |
     * | DELETE_     | Delete test                          |
     * | VIEW_       | View test                            |
     * | MANAGE_     | Manage test                          |
     * | MANAGE_     | Manage test                          |
     * | WORK_       | Delete and restore work              |
     * | WORK_       | Move work items out of this          |
     * | WORK_       | Permanently delete work              |
     * | RENAME      | Rename team                          |
     * | MANAGE_     | Manage project                       |
     * | MANAGE_     | Manage system project                |
     * | BYPASS_     | Bypass project property              |
     * | BYPASS_     | Bypass rules on work item            |
     * | SUPPRESS_   | Suppress notifications for work item |
     * | UPDATE_     | Update project                       |
     * | CHANGE_     | Change process of team               |
     * | AGILETOOLS_ | Agile backlog                        |
     * | AGILETOOLS_ | Agile                                |
     * 
     */
    @Import(name="permissions")
    private @Nullable Output<Map<String,String>> permissions;

    /**
     * @return the permissions to assign. The following permissions are available
     * 
     * | Permission  |             Description              |
     * |-------------|--------------------------------------|
     * | GENERIC_    | View project-level                   |
     * | GENERIC_    | Edit project-level                   |
     * | DELETE      | Delete team                          |
     * | PUBLISH_    | Create test                          |
     * | ADMINISTER_ | Administer a                         |
     * | START_      | Start a                              |
     * | EDIT_       | Edit build                           |
     * | UPDATE_     | Write to build operational           |
     * | DELETE_     | Delete test                          |
     * | VIEW_       | View test                            |
     * | MANAGE_     | Manage test                          |
     * | MANAGE_     | Manage test                          |
     * | WORK_       | Delete and restore work              |
     * | WORK_       | Move work items out of this          |
     * | WORK_       | Permanently delete work              |
     * | RENAME      | Rename team                          |
     * | MANAGE_     | Manage project                       |
     * | MANAGE_     | Manage system project                |
     * | BYPASS_     | Bypass project property              |
     * | BYPASS_     | Bypass rules on work item            |
     * | SUPPRESS_   | Suppress notifications for work item |
     * | UPDATE_     | Update project                       |
     * | CHANGE_     | Change process of team               |
     * | AGILETOOLS_ | Agile backlog                        |
     * | AGILETOOLS_ | Agile                                |
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
     * Replace (`true`) or merge (`false`) the permissions. Default: `true`
     * 
     */
    @Import(name="replace")
    private @Nullable Output<Boolean> replace;

    /**
     * @return Replace (`true`) or merge (`false`) the permissions. Default: `true`
     * 
     */
    public Optional<Output<Boolean>> replace() {
        return Optional.ofNullable(this.replace);
    }

    private ProjectPermissionsState() {}

    private ProjectPermissionsState(ProjectPermissionsState $) {
        this.permissions = $.permissions;
        this.principal = $.principal;
        this.projectId = $.projectId;
        this.replace = $.replace;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ProjectPermissionsState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ProjectPermissionsState $;

        public Builder() {
            $ = new ProjectPermissionsState();
        }

        public Builder(ProjectPermissionsState defaults) {
            $ = new ProjectPermissionsState(Objects.requireNonNull(defaults));
        }

        /**
         * @param permissions the permissions to assign. The following permissions are available
         * 
         * | Permission  |             Description              |
         * |-------------|--------------------------------------|
         * | GENERIC_    | View project-level                   |
         * | GENERIC_    | Edit project-level                   |
         * | DELETE      | Delete team                          |
         * | PUBLISH_    | Create test                          |
         * | ADMINISTER_ | Administer a                         |
         * | START_      | Start a                              |
         * | EDIT_       | Edit build                           |
         * | UPDATE_     | Write to build operational           |
         * | DELETE_     | Delete test                          |
         * | VIEW_       | View test                            |
         * | MANAGE_     | Manage test                          |
         * | MANAGE_     | Manage test                          |
         * | WORK_       | Delete and restore work              |
         * | WORK_       | Move work items out of this          |
         * | WORK_       | Permanently delete work              |
         * | RENAME      | Rename team                          |
         * | MANAGE_     | Manage project                       |
         * | MANAGE_     | Manage system project                |
         * | BYPASS_     | Bypass project property              |
         * | BYPASS_     | Bypass rules on work item            |
         * | SUPPRESS_   | Suppress notifications for work item |
         * | UPDATE_     | Update project                       |
         * | CHANGE_     | Change process of team               |
         * | AGILETOOLS_ | Agile backlog                        |
         * | AGILETOOLS_ | Agile                                |
         * 
         * @return builder
         * 
         */
        public Builder permissions(@Nullable Output<Map<String,String>> permissions) {
            $.permissions = permissions;
            return this;
        }

        /**
         * @param permissions the permissions to assign. The following permissions are available
         * 
         * | Permission  |             Description              |
         * |-------------|--------------------------------------|
         * | GENERIC_    | View project-level                   |
         * | GENERIC_    | Edit project-level                   |
         * | DELETE      | Delete team                          |
         * | PUBLISH_    | Create test                          |
         * | ADMINISTER_ | Administer a                         |
         * | START_      | Start a                              |
         * | EDIT_       | Edit build                           |
         * | UPDATE_     | Write to build operational           |
         * | DELETE_     | Delete test                          |
         * | VIEW_       | View test                            |
         * | MANAGE_     | Manage test                          |
         * | MANAGE_     | Manage test                          |
         * | WORK_       | Delete and restore work              |
         * | WORK_       | Move work items out of this          |
         * | WORK_       | Permanently delete work              |
         * | RENAME      | Rename team                          |
         * | MANAGE_     | Manage project                       |
         * | MANAGE_     | Manage system project                |
         * | BYPASS_     | Bypass project property              |
         * | BYPASS_     | Bypass rules on work item            |
         * | SUPPRESS_   | Suppress notifications for work item |
         * | UPDATE_     | Update project                       |
         * | CHANGE_     | Change process of team               |
         * | AGILETOOLS_ | Agile backlog                        |
         * | AGILETOOLS_ | Agile                                |
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
         * @param replace Replace (`true`) or merge (`false`) the permissions. Default: `true`
         * 
         * @return builder
         * 
         */
        public Builder replace(@Nullable Output<Boolean> replace) {
            $.replace = replace;
            return this;
        }

        /**
         * @param replace Replace (`true`) or merge (`false`) the permissions. Default: `true`
         * 
         * @return builder
         * 
         */
        public Builder replace(Boolean replace) {
            return replace(Output.of(replace));
        }

        public ProjectPermissionsState build() {
            return $;
        }
    }

}
