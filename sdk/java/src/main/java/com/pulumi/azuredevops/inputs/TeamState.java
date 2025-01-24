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


public final class TeamState extends com.pulumi.resources.ResourceArgs {

    public static final TeamState Empty = new TeamState();

    /**
     * List of subject descriptors to define administrators of the team.
     * 
     * &gt; **NOTE:** It&#39;s possible to define team administrators both within the
     * `azuredevops.Team` resource via the `administrators` block and by using the
     * `azuredevops.TeamAdministrators` resource. However it&#39;s not possible to use
     * both methods to manage team administrators, since there&#39;ll be conflicts.
     * 
     */
    @Import(name="administrators")
    private @Nullable Output<List<String>> administrators;

    /**
     * @return List of subject descriptors to define administrators of the team.
     * 
     * &gt; **NOTE:** It&#39;s possible to define team administrators both within the
     * `azuredevops.Team` resource via the `administrators` block and by using the
     * `azuredevops.TeamAdministrators` resource. However it&#39;s not possible to use
     * both methods to manage team administrators, since there&#39;ll be conflicts.
     * 
     */
    public Optional<Output<List<String>>> administrators() {
        return Optional.ofNullable(this.administrators);
    }

    /**
     * The description of the Team.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return The description of the Team.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The descriptor of the Team.
     * 
     */
    @Import(name="descriptor")
    private @Nullable Output<String> descriptor;

    /**
     * @return The descriptor of the Team.
     * 
     */
    public Optional<Output<String>> descriptor() {
        return Optional.ofNullable(this.descriptor);
    }

    /**
     * List of subject descriptors to define members of the team.
     * 
     * &gt; **NOTE:** It&#39;s possible to define team members both within the
     * `azuredevops.Team` resource via the `members` block and by using the
     * `azuredevops.TeamMembers` resource. However it&#39;s not possible to use
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
     * `azuredevops.TeamMembers` resource. However it&#39;s not possible to use
     * both methods to manage team members, since there&#39;ll be conflicts.
     * 
     */
    public Optional<Output<List<String>>> members() {
        return Optional.ofNullable(this.members);
    }

    /**
     * The name of the Team.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the Team.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
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

    private TeamState() {}

    private TeamState(TeamState $) {
        this.administrators = $.administrators;
        this.description = $.description;
        this.descriptor = $.descriptor;
        this.members = $.members;
        this.name = $.name;
        this.projectId = $.projectId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(TeamState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private TeamState $;

        public Builder() {
            $ = new TeamState();
        }

        public Builder(TeamState defaults) {
            $ = new TeamState(Objects.requireNonNull(defaults));
        }

        /**
         * @param administrators List of subject descriptors to define administrators of the team.
         * 
         * &gt; **NOTE:** It&#39;s possible to define team administrators both within the
         * `azuredevops.Team` resource via the `administrators` block and by using the
         * `azuredevops.TeamAdministrators` resource. However it&#39;s not possible to use
         * both methods to manage team administrators, since there&#39;ll be conflicts.
         * 
         * @return builder
         * 
         */
        public Builder administrators(@Nullable Output<List<String>> administrators) {
            $.administrators = administrators;
            return this;
        }

        /**
         * @param administrators List of subject descriptors to define administrators of the team.
         * 
         * &gt; **NOTE:** It&#39;s possible to define team administrators both within the
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
         * &gt; **NOTE:** It&#39;s possible to define team administrators both within the
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
         * @param description The description of the Team.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The description of the Team.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param descriptor The descriptor of the Team.
         * 
         * @return builder
         * 
         */
        public Builder descriptor(@Nullable Output<String> descriptor) {
            $.descriptor = descriptor;
            return this;
        }

        /**
         * @param descriptor The descriptor of the Team.
         * 
         * @return builder
         * 
         */
        public Builder descriptor(String descriptor) {
            return descriptor(Output.of(descriptor));
        }

        /**
         * @param members List of subject descriptors to define members of the team.
         * 
         * &gt; **NOTE:** It&#39;s possible to define team members both within the
         * `azuredevops.Team` resource via the `members` block and by using the
         * `azuredevops.TeamMembers` resource. However it&#39;s not possible to use
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
         * `azuredevops.TeamMembers` resource. However it&#39;s not possible to use
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
         * `azuredevops.TeamMembers` resource. However it&#39;s not possible to use
         * both methods to manage team members, since there&#39;ll be conflicts.
         * 
         * @return builder
         * 
         */
        public Builder members(String... members) {
            return members(List.of(members));
        }

        /**
         * @param name The name of the Team.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the Team.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
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

        public TeamState build() {
            return $;
        }
    }

}
