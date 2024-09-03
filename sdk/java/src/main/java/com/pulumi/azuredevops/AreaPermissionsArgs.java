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


public final class AreaPermissionsArgs extends com.pulumi.resources.ResourceArgs {

    public static final AreaPermissionsArgs Empty = new AreaPermissionsArgs();

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
     */
    @Import(name="permissions", required=true)
    private Output<Map<String,String>> permissions;

    /**
     * @return the permissions to assign. The following permissions are available.
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
     * Replace (`true`) or merge (`false`) the permissions. Default: `true`.
     * &lt;table&gt;
     * &lt;thead&gt;
     * &lt;tr&gt;
     * &lt;th&gt;Permission&lt;/th&gt;
     * &lt;th&gt;Description&lt;/th&gt;
     * &lt;/tr&gt;
     * &lt;/thead&gt;
     * &lt;tbody&gt;
     * &lt;tr&gt;
     * &lt;td&gt;GENERIC_READ&lt;/td&gt;
     * &lt;td&gt;View permissions for this node&lt;/td&gt;
     * &lt;/tr&gt;
     * &lt;tr&gt;
     * &lt;td&gt;GENERIC_WRITE&lt;/td&gt;
     * &lt;td&gt;Edit this node&lt;/td&gt;
     * &lt;/tr&gt;
     * &lt;tr&gt;
     * &lt;td&gt;CREATE_CHILDREN&lt;/td&gt;
     * &lt;td&gt;Create child nodes&lt;/td&gt;
     * &lt;/tr&gt;
     * &lt;tr&gt;
     * &lt;td&gt;DELETE&lt;/td&gt;
     * &lt;td&gt;Delete this node&lt;/td&gt;
     * &lt;/tr&gt;
     * &lt;tr&gt;
     * &lt;td&gt;WORK_ITEM_READ&lt;/td&gt;
     * &lt;td&gt;View work items in this node&lt;/td&gt;
     * &lt;/tr&gt;
     * &lt;tr&gt;
     * &lt;td&gt;WORK_ITEM_WRITE&lt;/td&gt;
     * &lt;td&gt;Edit work items in this node&lt;/td&gt;
     * &lt;/tr&gt;
     * &lt;tr&gt;
     * &lt;td&gt;MANAGE_TEST_PLANS&lt;/td&gt;
     * &lt;td&gt;Manage test plans&lt;/td&gt;
     * &lt;/tr&gt;
     * &lt;tr&gt;
     * &lt;td&gt;MANAGE_TEST_SUITES&lt;/td&gt;
     * &lt;td&gt;Manage test suites&lt;/td&gt;
     * &lt;/tr&gt;
     * &lt;tr&gt;
     * &lt;td&gt;WORK_ITEM_SAVE_COMMENT&lt;/td&gt;
     * &lt;td&gt;Edit work item comments in this node&lt;/td&gt;
     * &lt;/tr&gt;
     * &lt;/tbody&gt;
     * &lt;/table&gt;
     * 
     */
    @Import(name="replace")
    private @Nullable Output<Boolean> replace;

    /**
     * @return Replace (`true`) or merge (`false`) the permissions. Default: `true`.
     * &lt;table&gt;
     * &lt;thead&gt;
     * &lt;tr&gt;
     * &lt;th&gt;Permission&lt;/th&gt;
     * &lt;th&gt;Description&lt;/th&gt;
     * &lt;/tr&gt;
     * &lt;/thead&gt;
     * &lt;tbody&gt;
     * &lt;tr&gt;
     * &lt;td&gt;GENERIC_READ&lt;/td&gt;
     * &lt;td&gt;View permissions for this node&lt;/td&gt;
     * &lt;/tr&gt;
     * &lt;tr&gt;
     * &lt;td&gt;GENERIC_WRITE&lt;/td&gt;
     * &lt;td&gt;Edit this node&lt;/td&gt;
     * &lt;/tr&gt;
     * &lt;tr&gt;
     * &lt;td&gt;CREATE_CHILDREN&lt;/td&gt;
     * &lt;td&gt;Create child nodes&lt;/td&gt;
     * &lt;/tr&gt;
     * &lt;tr&gt;
     * &lt;td&gt;DELETE&lt;/td&gt;
     * &lt;td&gt;Delete this node&lt;/td&gt;
     * &lt;/tr&gt;
     * &lt;tr&gt;
     * &lt;td&gt;WORK_ITEM_READ&lt;/td&gt;
     * &lt;td&gt;View work items in this node&lt;/td&gt;
     * &lt;/tr&gt;
     * &lt;tr&gt;
     * &lt;td&gt;WORK_ITEM_WRITE&lt;/td&gt;
     * &lt;td&gt;Edit work items in this node&lt;/td&gt;
     * &lt;/tr&gt;
     * &lt;tr&gt;
     * &lt;td&gt;MANAGE_TEST_PLANS&lt;/td&gt;
     * &lt;td&gt;Manage test plans&lt;/td&gt;
     * &lt;/tr&gt;
     * &lt;tr&gt;
     * &lt;td&gt;MANAGE_TEST_SUITES&lt;/td&gt;
     * &lt;td&gt;Manage test suites&lt;/td&gt;
     * &lt;/tr&gt;
     * &lt;tr&gt;
     * &lt;td&gt;WORK_ITEM_SAVE_COMMENT&lt;/td&gt;
     * &lt;td&gt;Edit work item comments in this node&lt;/td&gt;
     * &lt;/tr&gt;
     * &lt;/tbody&gt;
     * &lt;/table&gt;
     * 
     */
    public Optional<Output<Boolean>> replace() {
        return Optional.ofNullable(this.replace);
    }

    private AreaPermissionsArgs() {}

    private AreaPermissionsArgs(AreaPermissionsArgs $) {
        this.path = $.path;
        this.permissions = $.permissions;
        this.principal = $.principal;
        this.projectId = $.projectId;
        this.replace = $.replace;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AreaPermissionsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AreaPermissionsArgs $;

        public Builder() {
            $ = new AreaPermissionsArgs();
        }

        public Builder(AreaPermissionsArgs defaults) {
            $ = new AreaPermissionsArgs(Objects.requireNonNull(defaults));
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
         * @param replace Replace (`true`) or merge (`false`) the permissions. Default: `true`.
         * &lt;table&gt;
         * &lt;thead&gt;
         * &lt;tr&gt;
         * &lt;th&gt;Permission&lt;/th&gt;
         * &lt;th&gt;Description&lt;/th&gt;
         * &lt;/tr&gt;
         * &lt;/thead&gt;
         * &lt;tbody&gt;
         * &lt;tr&gt;
         * &lt;td&gt;GENERIC_READ&lt;/td&gt;
         * &lt;td&gt;View permissions for this node&lt;/td&gt;
         * &lt;/tr&gt;
         * &lt;tr&gt;
         * &lt;td&gt;GENERIC_WRITE&lt;/td&gt;
         * &lt;td&gt;Edit this node&lt;/td&gt;
         * &lt;/tr&gt;
         * &lt;tr&gt;
         * &lt;td&gt;CREATE_CHILDREN&lt;/td&gt;
         * &lt;td&gt;Create child nodes&lt;/td&gt;
         * &lt;/tr&gt;
         * &lt;tr&gt;
         * &lt;td&gt;DELETE&lt;/td&gt;
         * &lt;td&gt;Delete this node&lt;/td&gt;
         * &lt;/tr&gt;
         * &lt;tr&gt;
         * &lt;td&gt;WORK_ITEM_READ&lt;/td&gt;
         * &lt;td&gt;View work items in this node&lt;/td&gt;
         * &lt;/tr&gt;
         * &lt;tr&gt;
         * &lt;td&gt;WORK_ITEM_WRITE&lt;/td&gt;
         * &lt;td&gt;Edit work items in this node&lt;/td&gt;
         * &lt;/tr&gt;
         * &lt;tr&gt;
         * &lt;td&gt;MANAGE_TEST_PLANS&lt;/td&gt;
         * &lt;td&gt;Manage test plans&lt;/td&gt;
         * &lt;/tr&gt;
         * &lt;tr&gt;
         * &lt;td&gt;MANAGE_TEST_SUITES&lt;/td&gt;
         * &lt;td&gt;Manage test suites&lt;/td&gt;
         * &lt;/tr&gt;
         * &lt;tr&gt;
         * &lt;td&gt;WORK_ITEM_SAVE_COMMENT&lt;/td&gt;
         * &lt;td&gt;Edit work item comments in this node&lt;/td&gt;
         * &lt;/tr&gt;
         * &lt;/tbody&gt;
         * &lt;/table&gt;
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
         * &lt;table&gt;
         * &lt;thead&gt;
         * &lt;tr&gt;
         * &lt;th&gt;Permission&lt;/th&gt;
         * &lt;th&gt;Description&lt;/th&gt;
         * &lt;/tr&gt;
         * &lt;/thead&gt;
         * &lt;tbody&gt;
         * &lt;tr&gt;
         * &lt;td&gt;GENERIC_READ&lt;/td&gt;
         * &lt;td&gt;View permissions for this node&lt;/td&gt;
         * &lt;/tr&gt;
         * &lt;tr&gt;
         * &lt;td&gt;GENERIC_WRITE&lt;/td&gt;
         * &lt;td&gt;Edit this node&lt;/td&gt;
         * &lt;/tr&gt;
         * &lt;tr&gt;
         * &lt;td&gt;CREATE_CHILDREN&lt;/td&gt;
         * &lt;td&gt;Create child nodes&lt;/td&gt;
         * &lt;/tr&gt;
         * &lt;tr&gt;
         * &lt;td&gt;DELETE&lt;/td&gt;
         * &lt;td&gt;Delete this node&lt;/td&gt;
         * &lt;/tr&gt;
         * &lt;tr&gt;
         * &lt;td&gt;WORK_ITEM_READ&lt;/td&gt;
         * &lt;td&gt;View work items in this node&lt;/td&gt;
         * &lt;/tr&gt;
         * &lt;tr&gt;
         * &lt;td&gt;WORK_ITEM_WRITE&lt;/td&gt;
         * &lt;td&gt;Edit work items in this node&lt;/td&gt;
         * &lt;/tr&gt;
         * &lt;tr&gt;
         * &lt;td&gt;MANAGE_TEST_PLANS&lt;/td&gt;
         * &lt;td&gt;Manage test plans&lt;/td&gt;
         * &lt;/tr&gt;
         * &lt;tr&gt;
         * &lt;td&gt;MANAGE_TEST_SUITES&lt;/td&gt;
         * &lt;td&gt;Manage test suites&lt;/td&gt;
         * &lt;/tr&gt;
         * &lt;tr&gt;
         * &lt;td&gt;WORK_ITEM_SAVE_COMMENT&lt;/td&gt;
         * &lt;td&gt;Edit work item comments in this node&lt;/td&gt;
         * &lt;/tr&gt;
         * &lt;/tbody&gt;
         * &lt;/table&gt;
         * 
         * @return builder
         * 
         */
        public Builder replace(Boolean replace) {
            return replace(Output.of(replace));
        }

        public AreaPermissionsArgs build() {
            if ($.permissions == null) {
                throw new MissingRequiredPropertyException("AreaPermissionsArgs", "permissions");
            }
            if ($.principal == null) {
                throw new MissingRequiredPropertyException("AreaPermissionsArgs", "principal");
            }
            if ($.projectId == null) {
                throw new MissingRequiredPropertyException("AreaPermissionsArgs", "projectId");
            }
            return $;
        }
    }

}
