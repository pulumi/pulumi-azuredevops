// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GitPermissionsArgs extends com.pulumi.resources.ResourceArgs {

    public static final GitPermissionsArgs Empty = new GitPermissionsArgs();

    /**
     * The name of the branch to assign the permissions.
     * 
     */
    @Import(name="branchName")
    private @Nullable Output<String> branchName;

    /**
     * @return The name of the branch to assign the permissions.
     * 
     */
    public Optional<Output<String>> branchName() {
        return Optional.ofNullable(this.branchName);
    }

    /**
     * the permissions to assign. The follwing permissions are available
     * 
     */
    @Import(name="permissions", required=true)
    private Output<Map<String,String>> permissions;

    /**
     * @return the permissions to assign. The follwing permissions are available
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

    /**
     * The ID of the GIT repository to assign the permissions
     * 
     */
    @Import(name="repositoryId")
    private @Nullable Output<String> repositoryId;

    /**
     * @return The ID of the GIT repository to assign the permissions
     * 
     */
    public Optional<Output<String>> repositoryId() {
        return Optional.ofNullable(this.repositoryId);
    }

    private GitPermissionsArgs() {}

    private GitPermissionsArgs(GitPermissionsArgs $) {
        this.branchName = $.branchName;
        this.permissions = $.permissions;
        this.principal = $.principal;
        this.projectId = $.projectId;
        this.replace = $.replace;
        this.repositoryId = $.repositoryId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GitPermissionsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GitPermissionsArgs $;

        public Builder() {
            $ = new GitPermissionsArgs();
        }

        public Builder(GitPermissionsArgs defaults) {
            $ = new GitPermissionsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param branchName The name of the branch to assign the permissions.
         * 
         * @return builder
         * 
         */
        public Builder branchName(@Nullable Output<String> branchName) {
            $.branchName = branchName;
            return this;
        }

        /**
         * @param branchName The name of the branch to assign the permissions.
         * 
         * @return builder
         * 
         */
        public Builder branchName(String branchName) {
            return branchName(Output.of(branchName));
        }

        /**
         * @param permissions the permissions to assign. The follwing permissions are available
         * 
         * @return builder
         * 
         */
        public Builder permissions(Output<Map<String,String>> permissions) {
            $.permissions = permissions;
            return this;
        }

        /**
         * @param permissions the permissions to assign. The follwing permissions are available
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

        /**
         * @param repositoryId The ID of the GIT repository to assign the permissions
         * 
         * @return builder
         * 
         */
        public Builder repositoryId(@Nullable Output<String> repositoryId) {
            $.repositoryId = repositoryId;
            return this;
        }

        /**
         * @param repositoryId The ID of the GIT repository to assign the permissions
         * 
         * @return builder
         * 
         */
        public Builder repositoryId(String repositoryId) {
            return repositoryId(Output.of(repositoryId));
        }

        public GitPermissionsArgs build() {
            $.permissions = Objects.requireNonNull($.permissions, "expected parameter 'permissions' to be non-null");
            $.principal = Objects.requireNonNull($.principal, "expected parameter 'principal' to be non-null");
            $.projectId = Objects.requireNonNull($.projectId, "expected parameter 'projectId' to be non-null");
            return $;
        }
    }

}
