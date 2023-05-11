// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.Build.inputs;

import com.pulumi.azuredevops.Build.inputs.BuildDefinitionCiTriggerArgs;
import com.pulumi.azuredevops.Build.inputs.BuildDefinitionPullRequestTriggerArgs;
import com.pulumi.azuredevops.Build.inputs.BuildDefinitionRepositoryArgs;
import com.pulumi.azuredevops.Build.inputs.BuildDefinitionScheduleArgs;
import com.pulumi.azuredevops.Build.inputs.BuildDefinitionVariableArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class BuildDefinitionState extends com.pulumi.resources.ResourceArgs {

    public static final BuildDefinitionState Empty = new BuildDefinitionState();

    /**
     * The agent pool that should execute the build. Defaults to `Azure Pipelines`.
     * 
     */
    @Import(name="agentPoolName")
    private @Nullable Output<String> agentPoolName;

    /**
     * @return The agent pool that should execute the build. Defaults to `Azure Pipelines`.
     * 
     */
    public Optional<Output<String>> agentPoolName() {
        return Optional.ofNullable(this.agentPoolName);
    }

    /**
     * Continuous Integration trigger.
     * 
     */
    @Import(name="ciTrigger")
    private @Nullable Output<BuildDefinitionCiTriggerArgs> ciTrigger;

    /**
     * @return Continuous Integration trigger.
     * 
     */
    public Optional<Output<BuildDefinitionCiTriggerArgs>> ciTrigger() {
        return Optional.ofNullable(this.ciTrigger);
    }

    /**
     * The name of the build definition.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the build definition.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The folder path of the build definition.
     * 
     */
    @Import(name="path")
    private @Nullable Output<String> path;

    /**
     * @return The folder path of the build definition.
     * 
     */
    public Optional<Output<String>> path() {
        return Optional.ofNullable(this.path);
    }

    /**
     * The project ID or project name.
     * 
     */
    @Import(name="projectId")
    private @Nullable Output<String> projectId;

    /**
     * @return The project ID or project name.
     * 
     */
    public Optional<Output<String>> projectId() {
        return Optional.ofNullable(this.projectId);
    }

    /**
     * Pull Request Integration Integration trigger.
     * 
     */
    @Import(name="pullRequestTrigger")
    private @Nullable Output<BuildDefinitionPullRequestTriggerArgs> pullRequestTrigger;

    /**
     * @return Pull Request Integration Integration trigger.
     * 
     */
    public Optional<Output<BuildDefinitionPullRequestTriggerArgs>> pullRequestTrigger() {
        return Optional.ofNullable(this.pullRequestTrigger);
    }

    /**
     * A `repository` block as documented below.
     * 
     */
    @Import(name="repository")
    private @Nullable Output<BuildDefinitionRepositoryArgs> repository;

    /**
     * @return A `repository` block as documented below.
     * 
     */
    public Optional<Output<BuildDefinitionRepositoryArgs>> repository() {
        return Optional.ofNullable(this.repository);
    }

    /**
     * The revision of the build definition
     * 
     */
    @Import(name="revision")
    private @Nullable Output<Integer> revision;

    /**
     * @return The revision of the build definition
     * 
     */
    public Optional<Output<Integer>> revision() {
        return Optional.ofNullable(this.revision);
    }

    @Import(name="schedules")
    private @Nullable Output<List<BuildDefinitionScheduleArgs>> schedules;

    public Optional<Output<List<BuildDefinitionScheduleArgs>>> schedules() {
        return Optional.ofNullable(this.schedules);
    }

    /**
     * A list of variable group IDs (integers) to link to the build definition.
     * 
     */
    @Import(name="variableGroups")
    private @Nullable Output<List<Integer>> variableGroups;

    /**
     * @return A list of variable group IDs (integers) to link to the build definition.
     * 
     */
    public Optional<Output<List<Integer>>> variableGroups() {
        return Optional.ofNullable(this.variableGroups);
    }

    /**
     * A list of `variable` blocks, as documented below.
     * 
     */
    @Import(name="variables")
    private @Nullable Output<List<BuildDefinitionVariableArgs>> variables;

    /**
     * @return A list of `variable` blocks, as documented below.
     * 
     */
    public Optional<Output<List<BuildDefinitionVariableArgs>>> variables() {
        return Optional.ofNullable(this.variables);
    }

    private BuildDefinitionState() {}

    private BuildDefinitionState(BuildDefinitionState $) {
        this.agentPoolName = $.agentPoolName;
        this.ciTrigger = $.ciTrigger;
        this.name = $.name;
        this.path = $.path;
        this.projectId = $.projectId;
        this.pullRequestTrigger = $.pullRequestTrigger;
        this.repository = $.repository;
        this.revision = $.revision;
        this.schedules = $.schedules;
        this.variableGroups = $.variableGroups;
        this.variables = $.variables;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(BuildDefinitionState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private BuildDefinitionState $;

        public Builder() {
            $ = new BuildDefinitionState();
        }

        public Builder(BuildDefinitionState defaults) {
            $ = new BuildDefinitionState(Objects.requireNonNull(defaults));
        }

        /**
         * @param agentPoolName The agent pool that should execute the build. Defaults to `Azure Pipelines`.
         * 
         * @return builder
         * 
         */
        public Builder agentPoolName(@Nullable Output<String> agentPoolName) {
            $.agentPoolName = agentPoolName;
            return this;
        }

        /**
         * @param agentPoolName The agent pool that should execute the build. Defaults to `Azure Pipelines`.
         * 
         * @return builder
         * 
         */
        public Builder agentPoolName(String agentPoolName) {
            return agentPoolName(Output.of(agentPoolName));
        }

        /**
         * @param ciTrigger Continuous Integration trigger.
         * 
         * @return builder
         * 
         */
        public Builder ciTrigger(@Nullable Output<BuildDefinitionCiTriggerArgs> ciTrigger) {
            $.ciTrigger = ciTrigger;
            return this;
        }

        /**
         * @param ciTrigger Continuous Integration trigger.
         * 
         * @return builder
         * 
         */
        public Builder ciTrigger(BuildDefinitionCiTriggerArgs ciTrigger) {
            return ciTrigger(Output.of(ciTrigger));
        }

        /**
         * @param name The name of the build definition.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the build definition.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param path The folder path of the build definition.
         * 
         * @return builder
         * 
         */
        public Builder path(@Nullable Output<String> path) {
            $.path = path;
            return this;
        }

        /**
         * @param path The folder path of the build definition.
         * 
         * @return builder
         * 
         */
        public Builder path(String path) {
            return path(Output.of(path));
        }

        /**
         * @param projectId The project ID or project name.
         * 
         * @return builder
         * 
         */
        public Builder projectId(@Nullable Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param projectId The project ID or project name.
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
        }

        /**
         * @param pullRequestTrigger Pull Request Integration Integration trigger.
         * 
         * @return builder
         * 
         */
        public Builder pullRequestTrigger(@Nullable Output<BuildDefinitionPullRequestTriggerArgs> pullRequestTrigger) {
            $.pullRequestTrigger = pullRequestTrigger;
            return this;
        }

        /**
         * @param pullRequestTrigger Pull Request Integration Integration trigger.
         * 
         * @return builder
         * 
         */
        public Builder pullRequestTrigger(BuildDefinitionPullRequestTriggerArgs pullRequestTrigger) {
            return pullRequestTrigger(Output.of(pullRequestTrigger));
        }

        /**
         * @param repository A `repository` block as documented below.
         * 
         * @return builder
         * 
         */
        public Builder repository(@Nullable Output<BuildDefinitionRepositoryArgs> repository) {
            $.repository = repository;
            return this;
        }

        /**
         * @param repository A `repository` block as documented below.
         * 
         * @return builder
         * 
         */
        public Builder repository(BuildDefinitionRepositoryArgs repository) {
            return repository(Output.of(repository));
        }

        /**
         * @param revision The revision of the build definition
         * 
         * @return builder
         * 
         */
        public Builder revision(@Nullable Output<Integer> revision) {
            $.revision = revision;
            return this;
        }

        /**
         * @param revision The revision of the build definition
         * 
         * @return builder
         * 
         */
        public Builder revision(Integer revision) {
            return revision(Output.of(revision));
        }

        public Builder schedules(@Nullable Output<List<BuildDefinitionScheduleArgs>> schedules) {
            $.schedules = schedules;
            return this;
        }

        public Builder schedules(List<BuildDefinitionScheduleArgs> schedules) {
            return schedules(Output.of(schedules));
        }

        public Builder schedules(BuildDefinitionScheduleArgs... schedules) {
            return schedules(List.of(schedules));
        }

        /**
         * @param variableGroups A list of variable group IDs (integers) to link to the build definition.
         * 
         * @return builder
         * 
         */
        public Builder variableGroups(@Nullable Output<List<Integer>> variableGroups) {
            $.variableGroups = variableGroups;
            return this;
        }

        /**
         * @param variableGroups A list of variable group IDs (integers) to link to the build definition.
         * 
         * @return builder
         * 
         */
        public Builder variableGroups(List<Integer> variableGroups) {
            return variableGroups(Output.of(variableGroups));
        }

        /**
         * @param variableGroups A list of variable group IDs (integers) to link to the build definition.
         * 
         * @return builder
         * 
         */
        public Builder variableGroups(Integer... variableGroups) {
            return variableGroups(List.of(variableGroups));
        }

        /**
         * @param variables A list of `variable` blocks, as documented below.
         * 
         * @return builder
         * 
         */
        public Builder variables(@Nullable Output<List<BuildDefinitionVariableArgs>> variables) {
            $.variables = variables;
            return this;
        }

        /**
         * @param variables A list of `variable` blocks, as documented below.
         * 
         * @return builder
         * 
         */
        public Builder variables(List<BuildDefinitionVariableArgs> variables) {
            return variables(Output.of(variables));
        }

        /**
         * @param variables A list of `variable` blocks, as documented below.
         * 
         * @return builder
         * 
         */
        public Builder variables(BuildDefinitionVariableArgs... variables) {
            return variables(List.of(variables));
        }

        public BuildDefinitionState build() {
            return $;
        }
    }

}