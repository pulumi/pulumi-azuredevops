// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class TeamMembersState extends com.pulumi.resources.ResourceArgs {

    public static final TeamMembersState Empty = new TeamMembersState();

    /**
     * List of subject descriptors to define members of the team.
     * 
     * &gt; **NOTE:** It&#39;s possible to define team members both within the
     * `azuredevops.Team` resource via the `members` block and by using the
     * `azuredevops.TeamMembers` resource. However, it&#39;s not possible to use
     * both methods to manage team members, since there&#39;ll be conflicts.
     * 
     */
    @Import(name="members")
    private @Nullable Output<List<String>> members;

    /**
     * @return List of subject descriptors to define members of the team.
     * 
     * &gt; **NOTE:** It&#39;s possible to define team members both within the
     * `azuredevops.Team` resource via the `members` block and by using the
     * `azuredevops.TeamMembers` resource. However, it&#39;s not possible to use
     * both methods to manage team members, since there&#39;ll be conflicts.
     * 
     */
    public Optional<Output<List<String>>> members() {
        return Optional.ofNullable(this.members);
    }

    /**
     * The mode how the resource manages team members. Possible values: `add`, `overwrite`. Defaults to `add`.
     * 
     * &gt; **NOTE:** 1. `mode = add`: the resource will ensure that all specified members will be part of the referenced team
     * &lt;br&gt;2. `mode = overwrite`: the resource will replace all existing members with the members specified within the `members` block
     * 
     */
    @Import(name="mode")
    private @Nullable Output<String> mode;

    /**
     * @return The mode how the resource manages team members. Possible values: `add`, `overwrite`. Defaults to `add`.
     * 
     * &gt; **NOTE:** 1. `mode = add`: the resource will ensure that all specified members will be part of the referenced team
     * &lt;br&gt;2. `mode = overwrite`: the resource will replace all existing members with the members specified within the `members` block
     * 
     */
    public Optional<Output<String>> mode() {
        return Optional.ofNullable(this.mode);
    }

    /**
     * The Project ID.
     * 
     */
    @Import(name="projectId")
    private @Nullable Output<String> projectId;

    /**
     * @return The Project ID.
     * 
     */
    public Optional<Output<String>> projectId() {
        return Optional.ofNullable(this.projectId);
    }

    /**
     * The ID of the Team.
     * 
     */
    @Import(name="teamId")
    private @Nullable Output<String> teamId;

    /**
     * @return The ID of the Team.
     * 
     */
    public Optional<Output<String>> teamId() {
        return Optional.ofNullable(this.teamId);
    }

    private TeamMembersState() {}

    private TeamMembersState(TeamMembersState $) {
        this.members = $.members;
        this.mode = $.mode;
        this.projectId = $.projectId;
        this.teamId = $.teamId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(TeamMembersState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private TeamMembersState $;

        public Builder() {
            $ = new TeamMembersState();
        }

        public Builder(TeamMembersState defaults) {
            $ = new TeamMembersState(Objects.requireNonNull(defaults));
        }

        /**
         * @param members List of subject descriptors to define members of the team.
         * 
         * &gt; **NOTE:** It&#39;s possible to define team members both within the
         * `azuredevops.Team` resource via the `members` block and by using the
         * `azuredevops.TeamMembers` resource. However, it&#39;s not possible to use
         * both methods to manage team members, since there&#39;ll be conflicts.
         * 
         * @return builder
         * 
         */
        public Builder members(@Nullable Output<List<String>> members) {
            $.members = members;
            return this;
        }

        /**
         * @param members List of subject descriptors to define members of the team.
         * 
         * &gt; **NOTE:** It&#39;s possible to define team members both within the
         * `azuredevops.Team` resource via the `members` block and by using the
         * `azuredevops.TeamMembers` resource. However, it&#39;s not possible to use
         * both methods to manage team members, since there&#39;ll be conflicts.
         * 
         * @return builder
         * 
         */
        public Builder members(List<String> members) {
            return members(Output.of(members));
        }

        /**
         * @param members List of subject descriptors to define members of the team.
         * 
         * &gt; **NOTE:** It&#39;s possible to define team members both within the
         * `azuredevops.Team` resource via the `members` block and by using the
         * `azuredevops.TeamMembers` resource. However, it&#39;s not possible to use
         * both methods to manage team members, since there&#39;ll be conflicts.
         * 
         * @return builder
         * 
         */
        public Builder members(String... members) {
            return members(List.of(members));
        }

        /**
         * @param mode The mode how the resource manages team members. Possible values: `add`, `overwrite`. Defaults to `add`.
         * 
         * &gt; **NOTE:** 1. `mode = add`: the resource will ensure that all specified members will be part of the referenced team
         * &lt;br&gt;2. `mode = overwrite`: the resource will replace all existing members with the members specified within the `members` block
         * 
         * @return builder
         * 
         */
        public Builder mode(@Nullable Output<String> mode) {
            $.mode = mode;
            return this;
        }

        /**
         * @param mode The mode how the resource manages team members. Possible values: `add`, `overwrite`. Defaults to `add`.
         * 
         * &gt; **NOTE:** 1. `mode = add`: the resource will ensure that all specified members will be part of the referenced team
         * &lt;br&gt;2. `mode = overwrite`: the resource will replace all existing members with the members specified within the `members` block
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
        public Builder projectId(@Nullable Output<String> projectId) {
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
        public Builder teamId(@Nullable Output<String> teamId) {
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

        public TeamMembersState build() {
            return $;
        }
    }

}
