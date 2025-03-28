// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class CheckBusinessHoursState extends com.pulumi.resources.ResourceArgs {

    public static final CheckBusinessHoursState Empty = new CheckBusinessHoursState();

    /**
     * The name of the business hours check displayed in the web UI.
     * 
     */
    @Import(name="displayName")
    private @Nullable Output<String> displayName;

    /**
     * @return The name of the business hours check displayed in the web UI.
     * 
     */
    public Optional<Output<String>> displayName() {
        return Optional.ofNullable(this.displayName);
    }

    /**
     * The end of the time period that this check will be allowed to pass, specified as 24-hour time with leading zeros.
     * 
     */
    @Import(name="endTime")
    private @Nullable Output<String> endTime;

    /**
     * @return The end of the time period that this check will be allowed to pass, specified as 24-hour time with leading zeros.
     * 
     */
    public Optional<Output<String>> endTime() {
        return Optional.ofNullable(this.endTime);
    }

    /**
     * This check will pass on Fridays. Defaults to `false`.
     * 
     */
    @Import(name="friday")
    private @Nullable Output<Boolean> friday;

    /**
     * @return This check will pass on Fridays. Defaults to `false`.
     * 
     */
    public Optional<Output<Boolean>> friday() {
        return Optional.ofNullable(this.friday);
    }

    /**
     * This check will pass on Mondays. Defaults to `false`.
     * 
     */
    @Import(name="monday")
    private @Nullable Output<Boolean> monday;

    /**
     * @return This check will pass on Mondays. Defaults to `false`.
     * 
     */
    public Optional<Output<Boolean>> monday() {
        return Optional.ofNullable(this.monday);
    }

    /**
     * The project ID.
     * 
     */
    @Import(name="projectId")
    private @Nullable Output<String> projectId;

    /**
     * @return The project ID.
     * 
     */
    public Optional<Output<String>> projectId() {
        return Optional.ofNullable(this.projectId);
    }

    /**
     * This check will pass on Saturdays. Defaults to `false`.
     * 
     */
    @Import(name="saturday")
    private @Nullable Output<Boolean> saturday;

    /**
     * @return This check will pass on Saturdays. Defaults to `false`.
     * 
     */
    public Optional<Output<Boolean>> saturday() {
        return Optional.ofNullable(this.saturday);
    }

    /**
     * The beginning of the time period that this check will be allowed to pass, specified as 24-hour time with leading zeros.
     * 
     */
    @Import(name="startTime")
    private @Nullable Output<String> startTime;

    /**
     * @return The beginning of the time period that this check will be allowed to pass, specified as 24-hour time with leading zeros.
     * 
     */
    public Optional<Output<String>> startTime() {
        return Optional.ofNullable(this.startTime);
    }

    /**
     * This check will pass on Sundays. Defaults to `false`.
     * 
     */
    @Import(name="sunday")
    private @Nullable Output<Boolean> sunday;

    /**
     * @return This check will pass on Sundays. Defaults to `false`.
     * 
     */
    public Optional<Output<Boolean>> sunday() {
        return Optional.ofNullable(this.sunday);
    }

    /**
     * The ID of the resource being protected by the check.
     * 
     */
    @Import(name="targetResourceId")
    private @Nullable Output<String> targetResourceId;

    /**
     * @return The ID of the resource being protected by the check.
     * 
     */
    public Optional<Output<String>> targetResourceId() {
        return Optional.ofNullable(this.targetResourceId);
    }

    /**
     * The type of resource being protected by the check. Possible values are: `endpoint`, `environment`, `queue`, `repository`, `securefile`, `variablegroup`.
     * 
     */
    @Import(name="targetResourceType")
    private @Nullable Output<String> targetResourceType;

    /**
     * @return The type of resource being protected by the check. Possible values are: `endpoint`, `environment`, `queue`, `repository`, `securefile`, `variablegroup`.
     * 
     */
    public Optional<Output<String>> targetResourceType() {
        return Optional.ofNullable(this.targetResourceType);
    }

    /**
     * This check will pass on Thursdays. Defaults to `false`.
     * 
     */
    @Import(name="thursday")
    private @Nullable Output<Boolean> thursday;

    /**
     * @return This check will pass on Thursdays. Defaults to `false`.
     * 
     */
    public Optional<Output<Boolean>> thursday() {
        return Optional.ofNullable(this.thursday);
    }

    /**
     * The time zone this check will be evaluated in. See below for supported values.
     * 
     */
    @Import(name="timeZone")
    private @Nullable Output<String> timeZone;

    /**
     * @return The time zone this check will be evaluated in. See below for supported values.
     * 
     */
    public Optional<Output<String>> timeZone() {
        return Optional.ofNullable(this.timeZone);
    }

    /**
     * The timeout in minutes for the business hours check. Defaults to `1440`.
     * 
     */
    @Import(name="timeout")
    private @Nullable Output<Integer> timeout;

    /**
     * @return The timeout in minutes for the business hours check. Defaults to `1440`.
     * 
     */
    public Optional<Output<Integer>> timeout() {
        return Optional.ofNullable(this.timeout);
    }

    /**
     * This check will pass on Tuesday. Defaults to `false`.
     * 
     */
    @Import(name="tuesday")
    private @Nullable Output<Boolean> tuesday;

    /**
     * @return This check will pass on Tuesday. Defaults to `false`.
     * 
     */
    public Optional<Output<Boolean>> tuesday() {
        return Optional.ofNullable(this.tuesday);
    }

    /**
     * The version of the check.
     * 
     */
    @Import(name="version")
    private @Nullable Output<Integer> version;

    /**
     * @return The version of the check.
     * 
     */
    public Optional<Output<Integer>> version() {
        return Optional.ofNullable(this.version);
    }

    /**
     * This check will pass on Wednesdays. Defaults to `false`.
     * 
     */
    @Import(name="wednesday")
    private @Nullable Output<Boolean> wednesday;

    /**
     * @return This check will pass on Wednesdays. Defaults to `false`.
     * 
     */
    public Optional<Output<Boolean>> wednesday() {
        return Optional.ofNullable(this.wednesday);
    }

    private CheckBusinessHoursState() {}

    private CheckBusinessHoursState(CheckBusinessHoursState $) {
        this.displayName = $.displayName;
        this.endTime = $.endTime;
        this.friday = $.friday;
        this.monday = $.monday;
        this.projectId = $.projectId;
        this.saturday = $.saturday;
        this.startTime = $.startTime;
        this.sunday = $.sunday;
        this.targetResourceId = $.targetResourceId;
        this.targetResourceType = $.targetResourceType;
        this.thursday = $.thursday;
        this.timeZone = $.timeZone;
        this.timeout = $.timeout;
        this.tuesday = $.tuesday;
        this.version = $.version;
        this.wednesday = $.wednesday;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(CheckBusinessHoursState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private CheckBusinessHoursState $;

        public Builder() {
            $ = new CheckBusinessHoursState();
        }

        public Builder(CheckBusinessHoursState defaults) {
            $ = new CheckBusinessHoursState(Objects.requireNonNull(defaults));
        }

        /**
         * @param displayName The name of the business hours check displayed in the web UI.
         * 
         * @return builder
         * 
         */
        public Builder displayName(@Nullable Output<String> displayName) {
            $.displayName = displayName;
            return this;
        }

        /**
         * @param displayName The name of the business hours check displayed in the web UI.
         * 
         * @return builder
         * 
         */
        public Builder displayName(String displayName) {
            return displayName(Output.of(displayName));
        }

        /**
         * @param endTime The end of the time period that this check will be allowed to pass, specified as 24-hour time with leading zeros.
         * 
         * @return builder
         * 
         */
        public Builder endTime(@Nullable Output<String> endTime) {
            $.endTime = endTime;
            return this;
        }

        /**
         * @param endTime The end of the time period that this check will be allowed to pass, specified as 24-hour time with leading zeros.
         * 
         * @return builder
         * 
         */
        public Builder endTime(String endTime) {
            return endTime(Output.of(endTime));
        }

        /**
         * @param friday This check will pass on Fridays. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder friday(@Nullable Output<Boolean> friday) {
            $.friday = friday;
            return this;
        }

        /**
         * @param friday This check will pass on Fridays. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder friday(Boolean friday) {
            return friday(Output.of(friday));
        }

        /**
         * @param monday This check will pass on Mondays. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder monday(@Nullable Output<Boolean> monday) {
            $.monday = monday;
            return this;
        }

        /**
         * @param monday This check will pass on Mondays. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder monday(Boolean monday) {
            return monday(Output.of(monday));
        }

        /**
         * @param projectId The project ID.
         * 
         * @return builder
         * 
         */
        public Builder projectId(@Nullable Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param projectId The project ID.
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
        }

        /**
         * @param saturday This check will pass on Saturdays. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder saturday(@Nullable Output<Boolean> saturday) {
            $.saturday = saturday;
            return this;
        }

        /**
         * @param saturday This check will pass on Saturdays. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder saturday(Boolean saturday) {
            return saturday(Output.of(saturday));
        }

        /**
         * @param startTime The beginning of the time period that this check will be allowed to pass, specified as 24-hour time with leading zeros.
         * 
         * @return builder
         * 
         */
        public Builder startTime(@Nullable Output<String> startTime) {
            $.startTime = startTime;
            return this;
        }

        /**
         * @param startTime The beginning of the time period that this check will be allowed to pass, specified as 24-hour time with leading zeros.
         * 
         * @return builder
         * 
         */
        public Builder startTime(String startTime) {
            return startTime(Output.of(startTime));
        }

        /**
         * @param sunday This check will pass on Sundays. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder sunday(@Nullable Output<Boolean> sunday) {
            $.sunday = sunday;
            return this;
        }

        /**
         * @param sunday This check will pass on Sundays. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder sunday(Boolean sunday) {
            return sunday(Output.of(sunday));
        }

        /**
         * @param targetResourceId The ID of the resource being protected by the check.
         * 
         * @return builder
         * 
         */
        public Builder targetResourceId(@Nullable Output<String> targetResourceId) {
            $.targetResourceId = targetResourceId;
            return this;
        }

        /**
         * @param targetResourceId The ID of the resource being protected by the check.
         * 
         * @return builder
         * 
         */
        public Builder targetResourceId(String targetResourceId) {
            return targetResourceId(Output.of(targetResourceId));
        }

        /**
         * @param targetResourceType The type of resource being protected by the check. Possible values are: `endpoint`, `environment`, `queue`, `repository`, `securefile`, `variablegroup`.
         * 
         * @return builder
         * 
         */
        public Builder targetResourceType(@Nullable Output<String> targetResourceType) {
            $.targetResourceType = targetResourceType;
            return this;
        }

        /**
         * @param targetResourceType The type of resource being protected by the check. Possible values are: `endpoint`, `environment`, `queue`, `repository`, `securefile`, `variablegroup`.
         * 
         * @return builder
         * 
         */
        public Builder targetResourceType(String targetResourceType) {
            return targetResourceType(Output.of(targetResourceType));
        }

        /**
         * @param thursday This check will pass on Thursdays. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder thursday(@Nullable Output<Boolean> thursday) {
            $.thursday = thursday;
            return this;
        }

        /**
         * @param thursday This check will pass on Thursdays. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder thursday(Boolean thursday) {
            return thursday(Output.of(thursday));
        }

        /**
         * @param timeZone The time zone this check will be evaluated in. See below for supported values.
         * 
         * @return builder
         * 
         */
        public Builder timeZone(@Nullable Output<String> timeZone) {
            $.timeZone = timeZone;
            return this;
        }

        /**
         * @param timeZone The time zone this check will be evaluated in. See below for supported values.
         * 
         * @return builder
         * 
         */
        public Builder timeZone(String timeZone) {
            return timeZone(Output.of(timeZone));
        }

        /**
         * @param timeout The timeout in minutes for the business hours check. Defaults to `1440`.
         * 
         * @return builder
         * 
         */
        public Builder timeout(@Nullable Output<Integer> timeout) {
            $.timeout = timeout;
            return this;
        }

        /**
         * @param timeout The timeout in minutes for the business hours check. Defaults to `1440`.
         * 
         * @return builder
         * 
         */
        public Builder timeout(Integer timeout) {
            return timeout(Output.of(timeout));
        }

        /**
         * @param tuesday This check will pass on Tuesday. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder tuesday(@Nullable Output<Boolean> tuesday) {
            $.tuesday = tuesday;
            return this;
        }

        /**
         * @param tuesday This check will pass on Tuesday. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder tuesday(Boolean tuesday) {
            return tuesday(Output.of(tuesday));
        }

        /**
         * @param version The version of the check.
         * 
         * @return builder
         * 
         */
        public Builder version(@Nullable Output<Integer> version) {
            $.version = version;
            return this;
        }

        /**
         * @param version The version of the check.
         * 
         * @return builder
         * 
         */
        public Builder version(Integer version) {
            return version(Output.of(version));
        }

        /**
         * @param wednesday This check will pass on Wednesdays. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder wednesday(@Nullable Output<Boolean> wednesday) {
            $.wednesday = wednesday;
            return this;
        }

        /**
         * @param wednesday This check will pass on Wednesdays. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder wednesday(Boolean wednesday) {
            return wednesday(Output.of(wednesday));
        }

        public CheckBusinessHoursState build() {
            return $;
        }
    }

}
