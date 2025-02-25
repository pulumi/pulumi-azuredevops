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


public final class BuildDefinitionPermissionsState extends com.pulumi.resources.ResourceArgs {

    public static final BuildDefinitionPermissionsState Empty = new BuildDefinitionPermissionsState();

    /**
     * The id of the build definition to assign the permissions.
     * 
     */
    @Import(name="buildDefinitionId")
    private @Nullable Output<String> buildDefinitionId;

    /**
     * @return The id of the build definition to assign the permissions.
     * 
     */
    public Optional<Output<String>> buildDefinitionId() {
        return Optional.ofNullable(this.buildDefinitionId);
    }

    /**
     * the permissions to assign. The following permissions are available.
     * 
     * | Permission                               | Description                           |
     * |------------------------------------------|---------------------------------------|
     * | ViewBuilds                               | View builds                           |
     * | EditBuildQuality                         | Edit build quality                    |
     * | RetainIndefinitely                       | Retain indefinitely                   |
     * | DeleteBuilds                             | Delete builds                         |
     * | ManageBuildQualities                     | Manage build qualities                |
     * | DestroyBuilds                            | Destroy builds                        |
     * | UpdateBuildInformation                   | Update build information              |
     * | QueueBuilds                              | Queue builds                          |
     * | ManageBuildQueue                         | Manage build queue                    |
     * | StopBuilds                               | Stop builds                           |
     * | ViewBuildDefinition                      | View build pipeline                   |
     * | EditBuildDefinition                      | Edit build pipeline                   |
     * | DeleteBuildDefinition                    | Delete build pipeline                 |
     * | OverrideBuildCheckInValidation           | Override check-in validation by build |
     * | AdministerBuildPermissions               | Administer build permissions          |
     * | CreateBuildDefinition                    | Create build pipeline                 |
     * | EditPipelineQueueConfigurationPermission | Edit queue build configuration        |
     * 
     */
    @Import(name="permissions")
    private @Nullable Output<Map<String,String>> permissions;

    /**
     * @return the permissions to assign. The following permissions are available.
     * 
     * | Permission                               | Description                           |
     * |------------------------------------------|---------------------------------------|
     * | ViewBuilds                               | View builds                           |
     * | EditBuildQuality                         | Edit build quality                    |
     * | RetainIndefinitely                       | Retain indefinitely                   |
     * | DeleteBuilds                             | Delete builds                         |
     * | ManageBuildQualities                     | Manage build qualities                |
     * | DestroyBuilds                            | Destroy builds                        |
     * | UpdateBuildInformation                   | Update build information              |
     * | QueueBuilds                              | Queue builds                          |
     * | ManageBuildQueue                         | Manage build queue                    |
     * | StopBuilds                               | Stop builds                           |
     * | ViewBuildDefinition                      | View build pipeline                   |
     * | EditBuildDefinition                      | Edit build pipeline                   |
     * | DeleteBuildDefinition                    | Delete build pipeline                 |
     * | OverrideBuildCheckInValidation           | Override check-in validation by build |
     * | AdministerBuildPermissions               | Administer build permissions          |
     * | CreateBuildDefinition                    | Create build pipeline                 |
     * | EditPipelineQueueConfigurationPermission | Edit queue build configuration        |
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

    private BuildDefinitionPermissionsState() {}

    private BuildDefinitionPermissionsState(BuildDefinitionPermissionsState $) {
        this.buildDefinitionId = $.buildDefinitionId;
        this.permissions = $.permissions;
        this.principal = $.principal;
        this.projectId = $.projectId;
        this.replace = $.replace;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(BuildDefinitionPermissionsState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private BuildDefinitionPermissionsState $;

        public Builder() {
            $ = new BuildDefinitionPermissionsState();
        }

        public Builder(BuildDefinitionPermissionsState defaults) {
            $ = new BuildDefinitionPermissionsState(Objects.requireNonNull(defaults));
        }

        /**
         * @param buildDefinitionId The id of the build definition to assign the permissions.
         * 
         * @return builder
         * 
         */
        public Builder buildDefinitionId(@Nullable Output<String> buildDefinitionId) {
            $.buildDefinitionId = buildDefinitionId;
            return this;
        }

        /**
         * @param buildDefinitionId The id of the build definition to assign the permissions.
         * 
         * @return builder
         * 
         */
        public Builder buildDefinitionId(String buildDefinitionId) {
            return buildDefinitionId(Output.of(buildDefinitionId));
        }

        /**
         * @param permissions the permissions to assign. The following permissions are available.
         * 
         * | Permission                               | Description                           |
         * |------------------------------------------|---------------------------------------|
         * | ViewBuilds                               | View builds                           |
         * | EditBuildQuality                         | Edit build quality                    |
         * | RetainIndefinitely                       | Retain indefinitely                   |
         * | DeleteBuilds                             | Delete builds                         |
         * | ManageBuildQualities                     | Manage build qualities                |
         * | DestroyBuilds                            | Destroy builds                        |
         * | UpdateBuildInformation                   | Update build information              |
         * | QueueBuilds                              | Queue builds                          |
         * | ManageBuildQueue                         | Manage build queue                    |
         * | StopBuilds                               | Stop builds                           |
         * | ViewBuildDefinition                      | View build pipeline                   |
         * | EditBuildDefinition                      | Edit build pipeline                   |
         * | DeleteBuildDefinition                    | Delete build pipeline                 |
         * | OverrideBuildCheckInValidation           | Override check-in validation by build |
         * | AdministerBuildPermissions               | Administer build permissions          |
         * | CreateBuildDefinition                    | Create build pipeline                 |
         * | EditPipelineQueueConfigurationPermission | Edit queue build configuration        |
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
         * | Permission                               | Description                           |
         * |------------------------------------------|---------------------------------------|
         * | ViewBuilds                               | View builds                           |
         * | EditBuildQuality                         | Edit build quality                    |
         * | RetainIndefinitely                       | Retain indefinitely                   |
         * | DeleteBuilds                             | Delete builds                         |
         * | ManageBuildQualities                     | Manage build qualities                |
         * | DestroyBuilds                            | Destroy builds                        |
         * | UpdateBuildInformation                   | Update build information              |
         * | QueueBuilds                              | Queue builds                          |
         * | ManageBuildQueue                         | Manage build queue                    |
         * | StopBuilds                               | Stop builds                           |
         * | ViewBuildDefinition                      | View build pipeline                   |
         * | EditBuildDefinition                      | Edit build pipeline                   |
         * | DeleteBuildDefinition                    | Delete build pipeline                 |
         * | OverrideBuildCheckInValidation           | Override check-in validation by build |
         * | AdministerBuildPermissions               | Administer build permissions          |
         * | CreateBuildDefinition                    | Create build pipeline                 |
         * | EditPipelineQueueConfigurationPermission | Edit queue build configuration        |
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

        public BuildDefinitionPermissionsState build() {
            return $;
        }
    }

}
