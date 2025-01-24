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


public final class LibraryPermissionsArgs extends com.pulumi.resources.ResourceArgs {

    public static final LibraryPermissionsArgs Empty = new LibraryPermissionsArgs();

    /**
     * the permissions to assign. The following permissions are available.
     * 
     * | Permission  | Description               |
     * |-------------|---------------------------|
     * | View        | View library item         |
     * | Administer  | Administer library item   |
     * | Create      | Create library item       |
     * | ViewSecrets | View library item secrets |
     * | Use         | Use library item          |
     * | Owner       | Owner library item        |
     * 
     */
    @Import(name="permissions", required=true)
    private Output<Map<String,String>> permissions;

    /**
     * @return the permissions to assign. The following permissions are available.
     * 
     * | Permission  | Description               |
     * |-------------|---------------------------|
     * | View        | View library item         |
     * | Administer  | Administer library item   |
     * | Create      | Create library item       |
     * | ViewSecrets | View library item secrets |
     * | Use         | Use library item          |
     * | Owner       | Owner library item        |
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
     * The ID of the project.
     * 
     */
    @Import(name="projectId", required=true)
    private Output<String> projectId;

    /**
     * @return The ID of the project.
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
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

    private LibraryPermissionsArgs() {}

    private LibraryPermissionsArgs(LibraryPermissionsArgs $) {
        this.permissions = $.permissions;
        this.principal = $.principal;
        this.projectId = $.projectId;
        this.replace = $.replace;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(LibraryPermissionsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private LibraryPermissionsArgs $;

        public Builder() {
            $ = new LibraryPermissionsArgs();
        }

        public Builder(LibraryPermissionsArgs defaults) {
            $ = new LibraryPermissionsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param permissions the permissions to assign. The following permissions are available.
         * 
         * | Permission  | Description               |
         * |-------------|---------------------------|
         * | View        | View library item         |
         * | Administer  | Administer library item   |
         * | Create      | Create library item       |
         * | ViewSecrets | View library item secrets |
         * | Use         | Use library item          |
         * | Owner       | Owner library item        |
         * 
         * @return builder
         * 
         */
        public Builder permissions(Output<Map<String,String>> permissions) {
            $.permissions = permissions;
            return this;
        }

        /**
         * @param permissions the permissions to assign. The following permissions are available.
         * 
         * | Permission  | Description               |
         * |-------------|---------------------------|
         * | View        | View library item         |
         * | Administer  | Administer library item   |
         * | Create      | Create library item       |
         * | ViewSecrets | View library item secrets |
         * | Use         | Use library item          |
         * | Owner       | Owner library item        |
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
         * @param projectId The ID of the project.
         * 
         * @return builder
         * 
         */
        public Builder projectId(Output<String> projectId) {
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

        public LibraryPermissionsArgs build() {
            if ($.permissions == null) {
                throw new MissingRequiredPropertyException("LibraryPermissionsArgs", "permissions");
            }
            if ($.principal == null) {
                throw new MissingRequiredPropertyException("LibraryPermissionsArgs", "principal");
            }
            if ($.projectId == null) {
                throw new MissingRequiredPropertyException("LibraryPermissionsArgs", "projectId");
            }
            return $;
        }
    }

}
