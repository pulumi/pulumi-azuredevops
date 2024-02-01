// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ServicehookStorageQueuePipelinesStageStateChangedEventArgs extends com.pulumi.resources.ResourceArgs {

    public static final ServicehookStorageQueuePipelinesStageStateChangedEventArgs Empty = new ServicehookStorageQueuePipelinesStageStateChangedEventArgs();

    /**
     * The pipeline ID that will generate an event.
     * 
     */
    @Import(name="pipelineId")
    private @Nullable Output<String> pipelineId;

    /**
     * @return The pipeline ID that will generate an event.
     * 
     */
    public Optional<Output<String>> pipelineId() {
        return Optional.ofNullable(this.pipelineId);
    }

    /**
     * Which stage should generate an event. Only valid if published_event is `StageStateChanged`. If not specified, all stages will trigger the event.
     * 
     */
    @Import(name="stageName")
    private @Nullable Output<String> stageName;

    /**
     * @return Which stage should generate an event. Only valid if published_event is `StageStateChanged`. If not specified, all stages will trigger the event.
     * 
     */
    public Optional<Output<String>> stageName() {
        return Optional.ofNullable(this.stageName);
    }

    /**
     * Which stage result should generate an event. Only valid if published_event is `StageStateChanged`. If not specified, all results will trigger the event.
     * 
     */
    @Import(name="stageResultFilter")
    private @Nullable Output<String> stageResultFilter;

    /**
     * @return Which stage result should generate an event. Only valid if published_event is `StageStateChanged`. If not specified, all results will trigger the event.
     * 
     */
    public Optional<Output<String>> stageResultFilter() {
        return Optional.ofNullable(this.stageResultFilter);
    }

    /**
     * Which stage state should generate an event. Only valid if published_event is `StageStateChanged`. If not specified, all states will trigger the event.
     * 
     */
    @Import(name="stageStateFilter")
    private @Nullable Output<String> stageStateFilter;

    /**
     * @return Which stage state should generate an event. Only valid if published_event is `StageStateChanged`. If not specified, all states will trigger the event.
     * 
     */
    public Optional<Output<String>> stageStateFilter() {
        return Optional.ofNullable(this.stageStateFilter);
    }

    private ServicehookStorageQueuePipelinesStageStateChangedEventArgs() {}

    private ServicehookStorageQueuePipelinesStageStateChangedEventArgs(ServicehookStorageQueuePipelinesStageStateChangedEventArgs $) {
        this.pipelineId = $.pipelineId;
        this.stageName = $.stageName;
        this.stageResultFilter = $.stageResultFilter;
        this.stageStateFilter = $.stageStateFilter;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServicehookStorageQueuePipelinesStageStateChangedEventArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServicehookStorageQueuePipelinesStageStateChangedEventArgs $;

        public Builder() {
            $ = new ServicehookStorageQueuePipelinesStageStateChangedEventArgs();
        }

        public Builder(ServicehookStorageQueuePipelinesStageStateChangedEventArgs defaults) {
            $ = new ServicehookStorageQueuePipelinesStageStateChangedEventArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param pipelineId The pipeline ID that will generate an event.
         * 
         * @return builder
         * 
         */
        public Builder pipelineId(@Nullable Output<String> pipelineId) {
            $.pipelineId = pipelineId;
            return this;
        }

        /**
         * @param pipelineId The pipeline ID that will generate an event.
         * 
         * @return builder
         * 
         */
        public Builder pipelineId(String pipelineId) {
            return pipelineId(Output.of(pipelineId));
        }

        /**
         * @param stageName Which stage should generate an event. Only valid if published_event is `StageStateChanged`. If not specified, all stages will trigger the event.
         * 
         * @return builder
         * 
         */
        public Builder stageName(@Nullable Output<String> stageName) {
            $.stageName = stageName;
            return this;
        }

        /**
         * @param stageName Which stage should generate an event. Only valid if published_event is `StageStateChanged`. If not specified, all stages will trigger the event.
         * 
         * @return builder
         * 
         */
        public Builder stageName(String stageName) {
            return stageName(Output.of(stageName));
        }

        /**
         * @param stageResultFilter Which stage result should generate an event. Only valid if published_event is `StageStateChanged`. If not specified, all results will trigger the event.
         * 
         * @return builder
         * 
         */
        public Builder stageResultFilter(@Nullable Output<String> stageResultFilter) {
            $.stageResultFilter = stageResultFilter;
            return this;
        }

        /**
         * @param stageResultFilter Which stage result should generate an event. Only valid if published_event is `StageStateChanged`. If not specified, all results will trigger the event.
         * 
         * @return builder
         * 
         */
        public Builder stageResultFilter(String stageResultFilter) {
            return stageResultFilter(Output.of(stageResultFilter));
        }

        /**
         * @param stageStateFilter Which stage state should generate an event. Only valid if published_event is `StageStateChanged`. If not specified, all states will trigger the event.
         * 
         * @return builder
         * 
         */
        public Builder stageStateFilter(@Nullable Output<String> stageStateFilter) {
            $.stageStateFilter = stageStateFilter;
            return this;
        }

        /**
         * @param stageStateFilter Which stage state should generate an event. Only valid if published_event is `StageStateChanged`. If not specified, all states will trigger the event.
         * 
         * @return builder
         * 
         */
        public Builder stageStateFilter(String stageStateFilter) {
            return stageStateFilter(Output.of(stageStateFilter));
        }

        public ServicehookStorageQueuePipelinesStageStateChangedEventArgs build() {
            return $;
        }
    }

}