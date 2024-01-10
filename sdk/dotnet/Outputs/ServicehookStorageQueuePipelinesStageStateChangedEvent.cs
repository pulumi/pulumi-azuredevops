// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureDevOps.Outputs
{

    [OutputType]
    public sealed class ServicehookStorageQueuePipelinesStageStateChangedEvent
    {
        /// <summary>
        /// The pipeline ID that will generate an event.
        /// </summary>
        public readonly string? PipelineId;
        /// <summary>
        /// Which stage should generate an event. Only valid if published_event is `StageStateChanged`. If not specified, all stages will trigger the event.
        /// </summary>
        public readonly string? StageName;
        /// <summary>
        /// Which stage result should generate an event. Only valid if published_event is `StageStateChanged`. If not specified, all results will trigger the event.
        /// </summary>
        public readonly string? StageResultFilter;
        /// <summary>
        /// Which stage state should generate an event. Only valid if published_event is `StageStateChanged`. If not specified, all states will trigger the event.
        /// </summary>
        public readonly string? StageStateFilter;

        [OutputConstructor]
        private ServicehookStorageQueuePipelinesStageStateChangedEvent(
            string? pipelineId,

            string? stageName,

            string? stageResultFilter,

            string? stageStateFilter)
        {
            PipelineId = pipelineId;
            StageName = stageName;
            StageResultFilter = stageResultFilter;
            StageStateFilter = stageStateFilter;
        }
    }
}
