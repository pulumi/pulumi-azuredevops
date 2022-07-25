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
    public sealed class BuildDefinitionSchedule
    {
        /// <summary>
        /// block supports the following:
        /// </summary>
        public readonly ImmutableArray<Outputs.BuildDefinitionScheduleBranchFilter> BranchFilters;
        public readonly ImmutableArray<string> DaysToBuilds;
        /// <summary>
        /// The ID of the schedule job
        /// </summary>
        public readonly string? ScheduleJobId;
        public readonly bool? ScheduleOnlyWithChanges;
        public readonly int? StartHours;
        public readonly int? StartMinutes;
        public readonly string? TimeZone;

        [OutputConstructor]
        private BuildDefinitionSchedule(
            ImmutableArray<Outputs.BuildDefinitionScheduleBranchFilter> branchFilters,

            ImmutableArray<string> daysToBuilds,

            string? scheduleJobId,

            bool? scheduleOnlyWithChanges,

            int? startHours,

            int? startMinutes,

            string? timeZone)
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
