// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class TeamAdministratorsArgs extends com.pulumi.resources.ResourceArgs {

    public static final TeamAdministratorsArgs Empty = new TeamAdministratorsArgs();

    /**
     * List of subject descriptors to define administrators of the team.
     * 
     * &gt; **NOTE** It&#39;s possible to define team administrators both within the
     * `azuredevops.Team` resource via the `administrators` block and by using the
     * `azuredevops.TeamAdministrators` resource. However it&#39;s not possible to use
     * both methods to manage team administrators, since there&#39;ll be conflicts.
     * 
     */
    @Import(name="administrators", required=true)
    private Output<List<String>> administrators;

    /**
     * @return List of subject descriptors to define administrators of the team.
     * 
     * &gt; **NOTE** It&#39;s possible to define team administrators both within the
     * `azuredevops.Team` resource via the `administrators` block and by using the
     * `azuredevops.TeamAdministrators` resource. However it&#39;s not possible to use
     * both methods to manage team administrators, since there&#39;ll be conflicts.
     * 
     */
    public Output<List<String>> administrators() {
        return this.administrators;
    }

    /**
     * The mode how the resource manages team administrators. Possible values: `add`, `overwrite`. Defaults to `add`.
     * 
     * &gt; **NOTE:** 1. `mode = add`: the resource will ensure that all specified administrators will be part of the referenced team
     * &lt;br&gt; 2. `mode = overwrite`: the resource will replace all existing administrators with the administrators specified within the `administrators` block
     * 
     */
    @Import(name="mode")
    private @Nullable Output<String> mode;

    /**
     * @return The mode how the resource manages team administrators. Possible values: `add`, `overwrite`. Defaults to `add`.
     * 
     * &gt; **NOTE:** 1. `mode = add`: the resource will ensure that all specified administrators will be part of the referenced team
     * &lt;br&gt; 2. `mode = overwrite`: the resource will replace all existing administrators with the administrators specified within the `administrators` block
     * 
     */
    public Optional<Output<String>> mode() {
        return Optional.ofNullable(this.mode);
    }

    /**
     * The Project ID.
     * 
     */
    @Import(name="projectId", required=true)
    private Output<String> projectId;

    /**
     * @return The Project ID.
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }

    /**
     * The ID of the Team.
     * 
     */
    @Import(name="teamId", required=true)
    private Output<String> teamId;

    /**
     * @return The ID of the Team.
     * 
     */
    public Output<String> teamId() {
        return this.teamId;
    }

    private TeamAdministratorsArgs() {}

    private TeamAdministratorsArgs(TeamAdministratorsArgs $) {
        this.administrators = $.administrators;
        this.mode = $.mode;
        this.projectId = $.projectId;
        this.teamId = $.teamId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(TeamAdministratorsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private TeamAdministratorsArgs $;

        public Builder() {
            $ = new TeamAdministratorsArgs();
        }

        public Builder(TeamAdministratorsArgs defaults) {
            $ = new TeamAdministratorsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param administrators List of subject descriptors to define administrators of the team.
         * 
         * &gt; **NOTE** It&#39;s possible to define team administrators both within the
         * `azuredevops.Team` resource via the `administrators` block and by using the
         * `azuredevops.TeamAdministrators` resource. However it&#39;s not possible to use
         * both methods to manage team administrators, since there&#39;ll be conflicts.
         * 
         * @return builder
         * 
         */
        public Builder administrators(Output<List<String>> administrators) {
            $.administrators = administrators;
            return this;
        }

        /**
         * @param administrators List of subject descriptors to define administrators of the team.
         * 
         * &gt; **NOTE** It&#39;s possible to define team administrators both within the
         * `azuredevops.Team` resource via the `administrators` block and by using the
         * `azuredevops.TeamAdministrators` resource. However it&#39;s not possible to use
         * both methods to manage team administrators, since there&#39;ll be conflicts.
         * 
         * @return builder
         * 
         */
        public Builder administrators(List<String> administrators) {
            return administrators(Output.of(administrators));
        }

        /**
         * @param administrators List of subject descriptors to define administrators of the team.
         * 
         * &gt; **NOTE** It&#39;s possible to define team administrators both within the
         * `azuredevops.Team` resource via the `administrators` block and by using the
         * `azuredevops.TeamAdministrators` resource. However it&#39;s not possible to use
         * both methods to manage team administrators, since there&#39;ll be conflicts.
         * 
         * @return builder
         * 
         */
        public Builder administrators(String... administrators) {
            return administrators(List.of(administrators));
        }

        /**
         * @param mode The mode how the resource manages team administrators. Possible values: `add`, `overwrite`. Defaults to `add`.
         * 
         * &gt; **NOTE:** 1. `mode = add`: the resource will ensure that all specified administrators will be part of the referenced team
         * &lt;br&gt; 2. `mode = overwrite`: the resource will replace all existing administrators with the administrators specified within the `administrators` block
         * 
         * @return builder
         * 
         */
        public Builder mode(@Nullable Output<String> mode) {
            $.mode = mode;
            return this;
        }

        /**
         * @param mode The mode how the resource manages team administrators. Possible values: `add`, `overwrite`. Defaults to `add`.
         * 
         * &gt; **NOTE:** 1. `mode = add`: the resource will ensure that all specified administrators will be part of the referenced team
         * &lt;br&gt; 2. `mode = overwrite`: the resource will replace all existing administrators with the administrators specified within the `administrators` block
         * 
         * @return builder
         * 
         */
        public Builder mode(String mode) {
            return mode(Output.of(mode));
        }

        /**
         * @param projectId The Project ID.
         * 
         * @return builder
         * 
         */
        public Builder projectId(Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param projectId The Project ID.
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
        }

        /**
         * @param teamId The ID of the Team.
         * 
         * @return builder
         * 
         */
        public Builder teamId(Output<String> teamId) {
            $.teamId = teamId;
            return this;
        }

        /**
         * @param teamId The ID of the Team.
         * 
         * @return builder
         * 
         */
        public Builder teamId(String teamId) {
            return teamId(Output.of(teamId));
        }

        public TeamAdministratorsArgs build() {
            if ($.administrators == null) {
                throw new MissingRequiredPropertyException("TeamAdministratorsArgs", "administrators");
            }
            if ($.projectId == null) {
                throw new MissingRequiredPropertyException("TeamAdministratorsArgs", "projectId");
            }
            if ($.teamId == null) {
                throw new MissingRequiredPropertyException("TeamAdministratorsArgs", "teamId");
            }
            return $;
        }
    }

}
