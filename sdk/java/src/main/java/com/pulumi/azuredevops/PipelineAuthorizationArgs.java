// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class PipelineAuthorizationArgs extends com.pulumi.resources.ResourceArgs {

    public static final PipelineAuthorizationArgs Empty = new PipelineAuthorizationArgs();

    /**
     * The ID of the pipeline. If not configured, all pipelines will be authorized. Changing this forces a new resource to be created.
     * 
     */
    @Import(name="pipelineId")
    private @Nullable Output<Integer> pipelineId;

    /**
     * @return The ID of the pipeline. If not configured, all pipelines will be authorized. Changing this forces a new resource to be created.
     * 
     */
    public Optional<Output<Integer>> pipelineId() {
        return Optional.ofNullable(this.pipelineId);
    }

    /**
     * The  ID of the project. Changing this forces a new resource to be created
     * 
     */
    @Import(name="projectId", required=true)
    private Output<String> projectId;

    /**
     * @return The  ID of the project. Changing this forces a new resource to be created
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }

    /**
     * The ID of the resource to authorize. Changing this forces a new resource to be created
     * 
     */
    @Import(name="resourceId", required=true)
    private Output<String> resourceId;

    /**
     * @return The ID of the resource to authorize. Changing this forces a new resource to be created
     * 
     */
    public Output<String> resourceId() {
        return this.resourceId;
    }

    /**
     * The type of the resource to authorize. Valid values: `endpoint`, `queue`, `variablegroup`, `environment`. Changing this forces a new resource to be created
     * 
     */
    @Import(name="type", required=true)
    private Output<String> type;

    /**
     * @return The type of the resource to authorize. Valid values: `endpoint`, `queue`, `variablegroup`, `environment`. Changing this forces a new resource to be created
     * 
     */
    public Output<String> type() {
        return this.type;
    }

    private PipelineAuthorizationArgs() {}

    private PipelineAuthorizationArgs(PipelineAuthorizationArgs $) {
        this.pipelineId = $.pipelineId;
        this.projectId = $.projectId;
        this.resourceId = $.resourceId;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(PipelineAuthorizationArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private PipelineAuthorizationArgs $;

        public Builder() {
            $ = new PipelineAuthorizationArgs();
        }

        public Builder(PipelineAuthorizationArgs defaults) {
            $ = new PipelineAuthorizationArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param pipelineId The ID of the pipeline. If not configured, all pipelines will be authorized. Changing this forces a new resource to be created.
         * 
         * @return builder
         * 
         */
        public Builder pipelineId(@Nullable Output<Integer> pipelineId) {
            $.pipelineId = pipelineId;
            return this;
        }

        /**
         * @param pipelineId The ID of the pipeline. If not configured, all pipelines will be authorized. Changing this forces a new resource to be created.
         * 
         * @return builder
         * 
         */
        public Builder pipelineId(Integer pipelineId) {
            return pipelineId(Output.of(pipelineId));
        }

        /**
         * @param projectId The  ID of the project. Changing this forces a new resource to be created
         * 
         * @return builder
         * 
         */
        public Builder projectId(Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param projectId The  ID of the project. Changing this forces a new resource to be created
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
        }

        /**
         * @param resourceId The ID of the resource to authorize. Changing this forces a new resource to be created
         * 
         * @return builder
         * 
         */
        public Builder resourceId(Output<String> resourceId) {
            $.resourceId = resourceId;
            return this;
        }

        /**
         * @param resourceId The ID of the resource to authorize. Changing this forces a new resource to be created
         * 
         * @return builder
         * 
         */
        public Builder resourceId(String resourceId) {
            return resourceId(Output.of(resourceId));
        }

        /**
         * @param type The type of the resource to authorize. Valid values: `endpoint`, `queue`, `variablegroup`, `environment`. Changing this forces a new resource to be created
         * 
         * @return builder
         * 
         */
        public Builder type(Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type The type of the resource to authorize. Valid values: `endpoint`, `queue`, `variablegroup`, `environment`. Changing this forces a new resource to be created
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        public PipelineAuthorizationArgs build() {
            $.projectId = Objects.requireNonNull($.projectId, "expected parameter 'projectId' to be non-null");
            $.resourceId = Objects.requireNonNull($.resourceId, "expected parameter 'resourceId' to be non-null");
            $.type = Objects.requireNonNull($.type, "expected parameter 'type' to be non-null");
            return $;
        }
    }

}
