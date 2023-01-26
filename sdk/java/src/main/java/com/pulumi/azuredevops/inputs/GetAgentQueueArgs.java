// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class GetAgentQueueArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetAgentQueueArgs Empty = new GetAgentQueueArgs();

    /**
     * Name of the Agent Queue.
     * 
     */
    @Import(name="name", required=true)
    private Output<String> name;

    /**
     * @return Name of the Agent Queue.
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    /**
     * The Project Id.
     * 
     */
    @Import(name="projectId", required=true)
    private Output<String> projectId;

    /**
     * @return The Project Id.
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }

    private GetAgentQueueArgs() {}

    private GetAgentQueueArgs(GetAgentQueueArgs $) {
        this.name = $.name;
        this.projectId = $.projectId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetAgentQueueArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetAgentQueueArgs $;

        public Builder() {
            $ = new GetAgentQueueArgs();
        }

        public Builder(GetAgentQueueArgs defaults) {
            $ = new GetAgentQueueArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param name Name of the Agent Queue.
         * 
         * @return builder
         * 
         */
        public Builder name(Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the Agent Queue.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param projectId The Project Id.
         * 
         * @return builder
         * 
         */
        public Builder projectId(Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param projectId The Project Id.
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
        }

        public GetAgentQueueArgs build() {
            $.name = Objects.requireNonNull($.name, "expected parameter 'name' to be non-null");
            $.projectId = Objects.requireNonNull($.projectId, "expected parameter 'projectId' to be non-null");
            return $;
        }
    }

}
