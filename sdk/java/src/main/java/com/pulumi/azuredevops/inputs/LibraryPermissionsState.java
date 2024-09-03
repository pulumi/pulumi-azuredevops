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


public final class LibraryPermissionsState extends com.pulumi.resources.ResourceArgs {

    public static final LibraryPermissionsState Empty = new LibraryPermissionsState();

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
     * &lt;table&gt;
     * &lt;thead&gt;
     * &lt;tr&gt;
     * &lt;th&gt;Permission&lt;/th&gt;
     * &lt;th&gt;Description&lt;/th&gt;
     * &lt;/tr&gt;
     * &lt;/thead&gt;
     * &lt;tbody&gt;
     * &lt;tr&gt;
     * &lt;td&gt;View&lt;/td&gt;
     * &lt;td&gt;View library item&lt;/td&gt;
     * &lt;/tr&gt;
     * &lt;tr&gt;
     * &lt;td&gt;Administer&lt;/td&gt;
     * &lt;td&gt;Administer library item&lt;/td&gt;
     * &lt;/tr&gt;
     * &lt;tr&gt;
     * &lt;td&gt;Create&lt;/td&gt;
     * &lt;td&gt;Create library item&lt;/td&gt;
     * &lt;/tr&gt;
     * &lt;tr&gt;
     * &lt;td&gt;ViewSecrets&lt;/td&gt;
     * &lt;td&gt;View library item secrets&lt;/td&gt;
     * &lt;/tr&gt;
     * &lt;tr&gt;
     * &lt;td&gt;Use&lt;/td&gt;
     * &lt;td&gt;Use library item&lt;/td&gt;
     * &lt;/tr&gt;
     * &lt;tr&gt;
     * &lt;td&gt;Owner&lt;/td&gt;
     * &lt;td&gt;Owner library item&lt;/td&gt;
     * &lt;/tr&gt;
     * &lt;/tbody&gt;
     * &lt;/table&gt;
     * 
     */
    @Import(name="replace")
    private @Nullable Output<Boolean> replace;

    /**
     * @return Replace (`true`) or merge (`false`) the permissions. Default: `true`
     * &lt;table&gt;
     * &lt;thead&gt;
     * &lt;tr&gt;
     * &lt;th&gt;Permission&lt;/th&gt;
     * &lt;th&gt;Description&lt;/th&gt;
     * &lt;/tr&gt;
     * &lt;/thead&gt;
     * &lt;tbody&gt;
     * &lt;tr&gt;
     * &lt;td&gt;View&lt;/td&gt;
     * &lt;td&gt;View library item&lt;/td&gt;
     * &lt;/tr&gt;
     * &lt;tr&gt;
     * &lt;td&gt;Administer&lt;/td&gt;
     * &lt;td&gt;Administer library item&lt;/td&gt;
     * &lt;/tr&gt;
     * &lt;tr&gt;
     * &lt;td&gt;Create&lt;/td&gt;
     * &lt;td&gt;Create library item&lt;/td&gt;
     * &lt;/tr&gt;
     * &lt;tr&gt;
     * &lt;td&gt;ViewSecrets&lt;/td&gt;
     * &lt;td&gt;View library item secrets&lt;/td&gt;
     * &lt;/tr&gt;
     * &lt;tr&gt;
     * &lt;td&gt;Use&lt;/td&gt;
     * &lt;td&gt;Use library item&lt;/td&gt;
     * &lt;/tr&gt;
     * &lt;tr&gt;
     * &lt;td&gt;Owner&lt;/td&gt;
     * &lt;td&gt;Owner library item&lt;/td&gt;
     * &lt;/tr&gt;
     * &lt;/tbody&gt;
     * &lt;/table&gt;
     * 
     */
    public Optional<Output<Boolean>> replace() {
        return Optional.ofNullable(this.replace);
    }

    private LibraryPermissionsState() {}

    private LibraryPermissionsState(LibraryPermissionsState $) {
        this.permissions = $.permissions;
        this.principal = $.principal;
        this.projectId = $.projectId;
        this.replace = $.replace;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(LibraryPermissionsState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private LibraryPermissionsState $;

        public Builder() {
            $ = new LibraryPermissionsState();
        }

        public Builder(LibraryPermissionsState defaults) {
            $ = new LibraryPermissionsState(Objects.requireNonNull(defaults));
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
         * &lt;table&gt;
         * &lt;thead&gt;
         * &lt;tr&gt;
         * &lt;th&gt;Permission&lt;/th&gt;
         * &lt;th&gt;Description&lt;/th&gt;
         * &lt;/tr&gt;
         * &lt;/thead&gt;
         * &lt;tbody&gt;
         * &lt;tr&gt;
         * &lt;td&gt;View&lt;/td&gt;
         * &lt;td&gt;View library item&lt;/td&gt;
         * &lt;/tr&gt;
         * &lt;tr&gt;
         * &lt;td&gt;Administer&lt;/td&gt;
         * &lt;td&gt;Administer library item&lt;/td&gt;
         * &lt;/tr&gt;
         * &lt;tr&gt;
         * &lt;td&gt;Create&lt;/td&gt;
         * &lt;td&gt;Create library item&lt;/td&gt;
         * &lt;/tr&gt;
         * &lt;tr&gt;
         * &lt;td&gt;ViewSecrets&lt;/td&gt;
         * &lt;td&gt;View library item secrets&lt;/td&gt;
         * &lt;/tr&gt;
         * &lt;tr&gt;
         * &lt;td&gt;Use&lt;/td&gt;
         * &lt;td&gt;Use library item&lt;/td&gt;
         * &lt;/tr&gt;
         * &lt;tr&gt;
         * &lt;td&gt;Owner&lt;/td&gt;
         * &lt;td&gt;Owner library item&lt;/td&gt;
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
         * @param replace Replace (`true`) or merge (`false`) the permissions. Default: `true`
         * &lt;table&gt;
         * &lt;thead&gt;
         * &lt;tr&gt;
         * &lt;th&gt;Permission&lt;/th&gt;
         * &lt;th&gt;Description&lt;/th&gt;
         * &lt;/tr&gt;
         * &lt;/thead&gt;
         * &lt;tbody&gt;
         * &lt;tr&gt;
         * &lt;td&gt;View&lt;/td&gt;
         * &lt;td&gt;View library item&lt;/td&gt;
         * &lt;/tr&gt;
         * &lt;tr&gt;
         * &lt;td&gt;Administer&lt;/td&gt;
         * &lt;td&gt;Administer library item&lt;/td&gt;
         * &lt;/tr&gt;
         * &lt;tr&gt;
         * &lt;td&gt;Create&lt;/td&gt;
         * &lt;td&gt;Create library item&lt;/td&gt;
         * &lt;/tr&gt;
         * &lt;tr&gt;
         * &lt;td&gt;ViewSecrets&lt;/td&gt;
         * &lt;td&gt;View library item secrets&lt;/td&gt;
         * &lt;/tr&gt;
         * &lt;tr&gt;
         * &lt;td&gt;Use&lt;/td&gt;
         * &lt;td&gt;Use library item&lt;/td&gt;
         * &lt;/tr&gt;
         * &lt;tr&gt;
         * &lt;td&gt;Owner&lt;/td&gt;
         * &lt;td&gt;Owner library item&lt;/td&gt;
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

        public LibraryPermissionsState build() {
            return $;
        }
    }

}
