// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class WorkItemQueryPermissionsArgs extends com.pulumi.resources.ResourceArgs {

    public static final WorkItemQueryPermissionsArgs Empty = new WorkItemQueryPermissionsArgs();

    /**
     * Path to a query or folder beneath `Shared Queries`
     * 
     */
    @Import(name="path")
    private @Nullable Output<String> path;

    /**
     * @return Path to a query or folder beneath `Shared Queries`
     * 
     */
    public Optional<Output<String>> path() {
        return Optional.ofNullable(this.path);
    }

    /**
     * the permissions to assign. The following permissions are available
     * 
     * | Permissions              | Description                        |
     * |--------------------------|------------------------------------|
     * | Read                     | Read                               |
     * | Contribute               | Contribute                         |
     * | Delete                   | Delete                             |
     * | ManagePermissions        | Manage Permissions                 |
     * 
     */
    @Import(name="permissions", required=true)
    private Output<Map<String,String>> permissions;

    /**
     * @return the permissions to assign. The following permissions are available
     * 
     * | Permissions              | Description                        |
     * |--------------------------|------------------------------------|
     * | Read                     | Read                               |
     * | Contribute               | Contribute                         |
     * | Delete                   | Delete                             |
     * | ManagePermissions        | Manage Permissions                 |
     * 
     */
    public Output<Map<String,String>> permissions() {
        return this.permissions;
    }

    /**
     * The **group** principal to assign the permissions.
     * 
     */
    @Import(name="principal", required=true)
    private Output<String> principal;

    /**
     * @return The **group** principal to assign the permissions.
     * 
     */
    public Output<String> principal() {
        return this.principal;
    }

    /**
     * The ID of the project to assign the permissions.
     * 
     */
    @Import(name="projectId", required=true)
    private Output<String> projectId;

    /**
     * @return The ID of the project to assign the permissions.
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }

    /**
     * Replace (`true`) or merge (`false`) the permissions. Defaults to `true`
     * 
     */
    @Import(name="replace")
    private @Nullable Output<Boolean> replace;

    /**
     * @return Replace (`true`) or merge (`false`) the permissions. Defaults to `true`
     * 
     */
    public Optional<Output<Boolean>> replace() {
        return Optional.ofNullable(this.replace);
    }

    private WorkItemQueryPermissionsArgs() {}

    private WorkItemQueryPermissionsArgs(WorkItemQueryPermissionsArgs $) {
        this.path = $.path;
        this.permissions = $.permissions;
        this.principal = $.principal;
        this.projectId = $.projectId;
        this.replace = $.replace;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(WorkItemQueryPermissionsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private WorkItemQueryPermissionsArgs $;

        public Builder() {
            $ = new WorkItemQueryPermissionsArgs();
        }

        public Builder(WorkItemQueryPermissionsArgs defaults) {
            $ = new WorkItemQueryPermissionsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param path Path to a query or folder beneath `Shared Queries`
         * 
         * @return builder
         * 
         */
        public Builder path(@Nullable Output<String> path) {
            $.path = path;
            return this;
        }

        /**
         * @param path Path to a query or folder beneath `Shared Queries`
         * 
         * @return builder
         * 
         */
        public Builder path(String path) {
            return path(Output.of(path));
        }

        /**
         * @param permissions the permissions to assign. The following permissions are available
         * 
         * | Permissions              | Description                        |
         * |--------------------------|------------------------------------|
         * | Read                     | Read                               |
         * | Contribute               | Contribute                         |
         * | Delete                   | Delete                             |
         * | ManagePermissions        | Manage Permissions                 |
         * 
         * @return builder
         * 
         */
        public Builder permissions(Output<Map<String,String>> permissions) {
            $.permissions = permissions;
            return this;
        }

        /**
         * @param permissions the permissions to assign. The following permissions are available
         * 
         * | Permissions              | Description                        |
         * |--------------------------|------------------------------------|
         * | Read                     | Read                               |
         * | Contribute               | Contribute                         |
         * | Delete                   | Delete                             |
         * | ManagePermissions        | Manage Permissions                 |
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
        public Builder principal(Output<String> principal) {
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
        public Builder projectId(Output<String> projectId) {
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
         * @param replace Replace (`true`) or merge (`false`) the permissions. Defaults to `true`
         * 
         * @return builder
         * 
         */
        public Builder replace(@Nullable Output<Boolean> replace) {
            $.replace = replace;
            return this;
        }

        /**
         * @param replace Replace (`true`) or merge (`false`) the permissions. Defaults to `true`
         * 
         * @return builder
         * 
         */
        public Builder replace(Boolean replace) {
            return replace(Output.of(replace));
        }

        public WorkItemQueryPermissionsArgs build() {
            if ($.permissions == null) {
                throw new MissingRequiredPropertyException("WorkItemQueryPermissionsArgs", "permissions");
            }
            if ($.principal == null) {
                throw new MissingRequiredPropertyException("WorkItemQueryPermissionsArgs", "principal");
            }
            if ($.projectId == null) {
                throw new MissingRequiredPropertyException("WorkItemQueryPermissionsArgs", "projectId");
            }
            return $;
        }
    }

}
