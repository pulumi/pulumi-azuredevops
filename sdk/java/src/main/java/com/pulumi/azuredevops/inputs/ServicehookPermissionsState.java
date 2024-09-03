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


public final class ServicehookPermissionsState extends com.pulumi.resources.ResourceArgs {

    public static final ServicehookPermissionsState Empty = new ServicehookPermissionsState();

    /**
     * the permissions to assign. The following permissions are available.
     * 
     */
    @Import(name="permissions")
    private @Nullable Output<Map<String,String>> permissions;

    /**
     * @return the permissions to assign. The following permissions are available.
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
     * The ID of the project.
     * 
     */
    @Import(name="projectId")
    private @Nullable Output<String> projectId;

    /**
     * @return The ID of the project.
     * 
     */
    public Optional<Output<String>> projectId() {
        return Optional.ofNullable(this.projectId);
    }

    /**
     * Replace (`true`) or merge (`false`) the permissions. Default: `true`
     * 
     * |        Name         | Permission |
     * |---------------------|------------|
     * | ViewSubscriptions   | View       |
     * | EditSubscriptions   | Edit       |
     * | DeleteSubscriptions | Delete     |
     * | PublishEvents       | Publish    |
     * 
     */
    @Import(name="replace")
    private @Nullable Output<Boolean> replace;

    /**
     * @return Replace (`true`) or merge (`false`) the permissions. Default: `true`
     * 
     * |        Name         | Permission |
     * |---------------------|------------|
     * | ViewSubscriptions   | View       |
     * | EditSubscriptions   | Edit       |
     * | DeleteSubscriptions | Delete     |
     * | PublishEvents       | Publish    |
     * 
     */
    public Optional<Output<Boolean>> replace() {
        return Optional.ofNullable(this.replace);
    }

    private ServicehookPermissionsState() {}

    private ServicehookPermissionsState(ServicehookPermissionsState $) {
        this.permissions = $.permissions;
        this.principal = $.principal;
        this.projectId = $.projectId;
        this.replace = $.replace;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServicehookPermissionsState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServicehookPermissionsState $;

        public Builder() {
            $ = new ServicehookPermissionsState();
        }

        public Builder(ServicehookPermissionsState defaults) {
            $ = new ServicehookPermissionsState(Objects.requireNonNull(defaults));
        }

        /**
         * @param permissions the permissions to assign. The following permissions are available.
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
         * @param projectId The ID of the project.
         * 
         * @return builder
         * 
         */
        public Builder projectId(@Nullable Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param projectId The ID of the project.
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
         * |        Name         | Permission |
         * |---------------------|------------|
         * | ViewSubscriptions   | View       |
         * | EditSubscriptions   | Edit       |
         * | DeleteSubscriptions | Delete     |
         * | PublishEvents       | Publish    |
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
         * |        Name         | Permission |
         * |---------------------|------------|
         * | ViewSubscriptions   | View       |
         * | EditSubscriptions   | Edit       |
         * | DeleteSubscriptions | Delete     |
         * | PublishEvents       | Publish    |
         * 
         * @return builder
         * 
         */
        public Builder replace(Boolean replace) {
            return replace(Output.of(replace));
        }

        public ServicehookPermissionsState build() {
            return $;
        }
    }

}
