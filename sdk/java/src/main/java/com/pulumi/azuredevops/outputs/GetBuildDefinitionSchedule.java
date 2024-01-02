// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.outputs;

import com.pulumi.azuredevops.outputs.GetBuildDefinitionScheduleBranchFilter;
import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetBuildDefinitionSchedule {
    /**
     * @return A `branch_filter` block as defined above.
     * 
     */
    private List<GetBuildDefinitionScheduleBranchFilter> branchFilters;
    /**
     * @return A list of days to build on.
     * 
     */
    private List<String> daysToBuilds;
    /**
     * @return The ID of the schedule job.
     * 
     */
    private String scheduleJobId;
    /**
     * @return Schedule builds if the source or pipeline has changed.
     * 
     */
    private Boolean scheduleOnlyWithChanges;
    /**
     * @return Build start hour.
     * 
     */
    private Integer startHours;
    /**
     * @return Build start minute.
     * 
     */
    private Integer startMinutes;
    /**
     * @return Build time zone.
     * 
     */
    private String timeZone;

    private GetBuildDefinitionSchedule() {}
    /**
     * @return A `branch_filter` block as defined above.
     * 
     */
    public List<GetBuildDefinitionScheduleBranchFilter> branchFilters() {
        return this.branchFilters;
    }
    /**
     * @return A list of days to build on.
     * 
     */
    public List<String> daysToBuilds() {
        return this.daysToBuilds;
    }
    /**
     * @return The ID of the schedule job.
     * 
     */
    public String scheduleJobId() {
        return this.scheduleJobId;
    }
    /**
     * @return Schedule builds if the source or pipeline has changed.
     * 
     */
    public Boolean scheduleOnlyWithChanges() {
        return this.scheduleOnlyWithChanges;
    }
    /**
     * @return Build start hour.
     * 
     */
    public Integer startHours() {
        return this.startHours;
    }
    /**
     * @return Build start minute.
     * 
     */
    public Integer startMinutes() {
        return this.startMinutes;
    }
    /**
     * @return Build time zone.
     * 
     */
    public String timeZone() {
        return this.timeZone;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetBuildDefinitionSchedule defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<GetBuildDefinitionScheduleBranchFilter> branchFilters;
        private List<String> daysToBuilds;
        private String scheduleJobId;
        private Boolean scheduleOnlyWithChanges;
        private Integer startHours;
        private Integer startMinutes;
        private String timeZone;
        public Builder() {}
        public Builder(GetBuildDefinitionSchedule defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.branchFilters = defaults.branchFilters;
    	      this.daysToBuilds = defaults.daysToBuilds;
    	      this.scheduleJobId = defaults.scheduleJobId;
    	      this.scheduleOnlyWithChanges = defaults.scheduleOnlyWithChanges;
    	      this.startHours = defaults.startHours;
    	      this.startMinutes = defaults.startMinutes;
    	      this.timeZone = defaults.timeZone;
        }

        @CustomType.Setter
        public Builder branchFilters(List<GetBuildDefinitionScheduleBranchFilter> branchFilters) {
            if (branchFilters == null) {
              throw new MissingRequiredPropertyException("GetBuildDefinitionSchedule", "branchFilters");
            }
            this.branchFilters = branchFilters;
            return this;
        }
        public Builder branchFilters(GetBuildDefinitionScheduleBranchFilter... branchFilters) {
            return branchFilters(List.of(branchFilters));
        }
        @CustomType.Setter
        public Builder daysToBuilds(List<String> daysToBuilds) {
            if (daysToBuilds == null) {
              throw new MissingRequiredPropertyException("GetBuildDefinitionSchedule", "daysToBuilds");
            }
            this.daysToBuilds = daysToBuilds;
            return this;
        }
        public Builder daysToBuilds(String... daysToBuilds) {
            return daysToBuilds(List.of(daysToBuilds));
        }
        @CustomType.Setter
        public Builder scheduleJobId(String scheduleJobId) {
            if (scheduleJobId == null) {
              throw new MissingRequiredPropertyException("GetBuildDefinitionSchedule", "scheduleJobId");
            }
            this.scheduleJobId = scheduleJobId;
            return this;
        }
        @CustomType.Setter
        public Builder scheduleOnlyWithChanges(Boolean scheduleOnlyWithChanges) {
            if (scheduleOnlyWithChanges == null) {
              throw new MissingRequiredPropertyException("GetBuildDefinitionSchedule", "scheduleOnlyWithChanges");
            }
            this.scheduleOnlyWithChanges = scheduleOnlyWithChanges;
            return this;
        }
        @CustomType.Setter
        public Builder startHours(Integer startHours) {
            if (startHours == null) {
              throw new MissingRequiredPropertyException("GetBuildDefinitionSchedule", "startHours");
            }
            this.startHours = startHours;
            return this;
        }
        @CustomType.Setter
        public Builder startMinutes(Integer startMinutes) {
            if (startMinutes == null) {
              throw new MissingRequiredPropertyException("GetBuildDefinitionSchedule", "startMinutes");
            }
            this.startMinutes = startMinutes;
            return this;
        }
        @CustomType.Setter
        public Builder timeZone(String timeZone) {
            if (timeZone == null) {
              throw new MissingRequiredPropertyException("GetBuildDefinitionSchedule", "timeZone");
            }
            this.timeZone = timeZone;
            return this;
        }
        public GetBuildDefinitionSchedule build() {
            final var _resultValue = new GetBuildDefinitionSchedule();
            _resultValue.branchFilters = branchFilters;
            _resultValue.daysToBuilds = daysToBuilds;
            _resultValue.scheduleJobId = scheduleJobId;
            _resultValue.scheduleOnlyWithChanges = scheduleOnlyWithChanges;
            _resultValue.startHours = startHours;
            _resultValue.startMinutes = startMinutes;
            _resultValue.timeZone = timeZone;
            return _resultValue;
        }
    }
}
