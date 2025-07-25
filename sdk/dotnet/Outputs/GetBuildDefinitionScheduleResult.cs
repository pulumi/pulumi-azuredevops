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
    public sealed class GetBuildDefinitionScheduleResult
    {
        /// <summary>
        /// A `branch_filter` block as defined above.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetBuildDefinitionScheduleBranchFilterResult> BranchFilters;
        /// <summary>
        /// A list of days to build on.
        /// </summary>
        public readonly ImmutableArray<string> DaysToBuilds;
        /// <summary>
        /// The ID of the schedule job.
        /// </summary>
        public readonly string ScheduleJobId;
        /// <summary>
        /// Schedule builds if the source or pipeline has changed.
        /// </summary>
        public readonly bool ScheduleOnlyWithChanges;
        /// <summary>
        /// Build start hour.
        /// </summary>
        public readonly int StartHours;
        /// <summary>
        /// Build start minute.
        /// </summary>
        public readonly int StartMinutes;
        /// <summary>
        /// Build time zone.
        /// </summary>
        public readonly string TimeZone;

        [OutputConstructor]
        private GetBuildDefinitionScheduleResult(
            ImmutableArray<Outputs.GetBuildDefinitionScheduleBranchFilterResult> branchFilters,

            ImmutableArray<string> daysToBuilds,

            string scheduleJobId,

            bool scheduleOnlyWithChanges,

            int startHours,

            int startMinutes,

            string timeZone)
        {
            BranchFilters = branchFilters;
            DaysToBuilds = daysToBuilds;
            ScheduleJobId = scheduleJobId;
            ScheduleOnlyWithChanges = scheduleOnlyWithChanges;
            StartHours = startHours;
            StartMinutes = startMinutes;
            TimeZone = timeZone;
        }
    }
}
