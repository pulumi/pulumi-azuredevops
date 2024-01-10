// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.Agent.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class QueueState extends com.pulumi.resources.ResourceArgs {

    public static final QueueState Empty = new QueueState();

    /**
     * The ID of the organization agent pool. Conflicts with `name`.
     * 
     * &gt; **NOTE:**
     * One of `name` or `agent_pool_id` must be specified, but not both.
     * When `agent_pool_id` is specified, the agent queue name will be derived from the agent pool name.
     * 
     */
    @Import(name="agentPoolId")
    private @Nullable Output<Integer> agentPoolId;

    /**
     * @return The ID of the organization agent pool. Conflicts with `name`.
     * 
     * &gt; **NOTE:**
     * One of `name` or `agent_pool_id` must be specified, but not both.
     * When `agent_pool_id` is specified, the agent queue name will be derived from the agent pool name.
     * 
     */
    public Optional<Output<Integer>> agentPoolId() {
        return Optional.ofNullable(this.agentPoolId);
    }

    /**
     * The name of the agent queue. Defaults to the ID of the agent pool. Conflicts with `agent_pool_id`.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the agent queue. Defaults to the ID of the agent pool. Conflicts with `agent_pool_id`.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The ID of the project in which to create the resource.
     * 
     */
    @Import(name="projectId")
    private @Nullable Output<String> projectId;

    /**
     * @return The ID of the project in which to create the resource.
     * 
     */
    public Optional<Output<String>> projectId() {
        return Optional.ofNullable(this.projectId);
    }

    private QueueState() {}

    private QueueState(QueueState $) {
        this.agentPoolId = $.agentPoolId;
        this.name = $.name;
        this.projectId = $.projectId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(QueueState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private QueueState $;

        public Builder() {
            $ = new QueueState();
        }

        public Builder(QueueState defaults) {
            $ = new QueueState(Objects.requireNonNull(defaults));
        }

        /**
         * @param agentPoolId The ID of the organization agent pool. Conflicts with `name`.
         * 
         * &gt; **NOTE:**
         * One of `name` or `agent_pool_id` must be specified, but not both.
         * When `agent_pool_id` is specified, the agent queue name will be derived from the agent pool name.
         * 
         * @return builder
         * 
         */
        public Builder agentPoolId(@Nullable Output<Integer> agentPoolId) {
            $.agentPoolId = agentPoolId;
            return this;
        }

        /**
         * @param agentPoolId The ID of the organization agent pool. Conflicts with `name`.
         * 
         * &gt; **NOTE:**
         * One of `name` or `agent_pool_id` must be specified, but not both.
         * When `agent_pool_id` is specified, the agent queue name will be derived from the agent pool name.
         * 
         * @return builder
         * 
         */
        public Builder agentPoolId(Integer agentPoolId) {
            return agentPoolId(Output.of(agentPoolId));
        }

        /**
         * @param name The name of the agent queue. Defaults to the ID of the agent pool. Conflicts with `agent_pool_id`.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the agent queue. Defaults to the ID of the agent pool. Conflicts with `agent_pool_id`.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param projectId The ID of the project in which to create the resource.
         * 
         * @return builder
         * 
         */
        public Builder projectId(@Nullable Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param projectId The ID of the project in which to create the resource.
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
        }

        public QueueState build() {
            return $;
        }
    }

}
