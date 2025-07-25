// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureDevOps.Outputs
{

    [OutputType]
    public sealed class ServicehookStorageQueuePipelinesRunStateChangedEvent
    {
        /// <summary>
        /// The pipeline ID that will generate an event. If not specified, all pipelines in the project will trigger the event.
        /// </summary>
        public readonly string? PipelineId;
        /// <summary>
        /// Which run result should generate an event. Only valid if published_event is `RunStateChanged`. If not specified, all results will trigger the event.
        /// </summary>
        public readonly string? RunResultFilter;
        /// <summary>
        /// Which run state should generate an event. Only valid if published_event is `RunStateChanged`. If not specified, all states will trigger the event.
        /// </summary>
        public readonly string? RunStateFilter;

        [OutputConstructor]
        private ServicehookStorageQueuePipelinesRunStateChangedEvent(
            string? pipelineId,

            string? runResultFilter,

            string? runStateFilter)
        {
            PipelineId = pipelineId;
            RunResultFilter = runResultFilter;
            RunStateFilter = runStateFilter;
        }
    }
}
