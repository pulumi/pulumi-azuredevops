// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.inputs;

import com.pulumi.azuredevops.inputs.BuildDefinitionScheduleBranchFilterArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class BuildDefinitionScheduleArgs extends com.pulumi.resources.ResourceArgs {

    public static final BuildDefinitionScheduleArgs Empty = new BuildDefinitionScheduleArgs();

    /**
     * block supports the following:
     * 
     */
    @Import(name="branchFilters")
    private @Nullable Output<List<BuildDefinitionScheduleBranchFilterArgs>> branchFilters;

    /**
     * @return block supports the following:
     * 
     */
    public Optional<Output<List<BuildDefinitionScheduleBranchFilterArgs>>> branchFilters() {
        return Optional.ofNullable(this.branchFilters);
    }

    /**
     * When to build. Valid values: `Mon`, `Tue`, `Wed`, `Thu`, `Fri`, `Sat`, `Sun`.
     * 
     */
    @Import(name="daysToBuilds", required=true)
    private Output<List<String>> daysToBuilds;

    /**
     * @return When to build. Valid values: `Mon`, `Tue`, `Wed`, `Thu`, `Fri`, `Sat`, `Sun`.
     * 
     */
    public Output<List<String>> daysToBuilds() {
        return this.daysToBuilds;
    }

    /**
     * The ID of the schedule job
     * 
     */
    @Import(name="scheduleJobId")
    private @Nullable Output<String> scheduleJobId;

    /**
     * @return The ID of the schedule job
     * 
     */
    public Optional<Output<String>> scheduleJobId() {
        return Optional.ofNullable(this.scheduleJobId);
    }

    /**
     * Schedule builds if the source or pipeline has changed. Defaults to `true`.
     * 
     */
    @Import(name="scheduleOnlyWithChanges")
    private @Nullable Output<Boolean> scheduleOnlyWithChanges;

    /**
     * @return Schedule builds if the source or pipeline has changed. Defaults to `true`.
     * 
     */
    public Optional<Output<Boolean>> scheduleOnlyWithChanges() {
        return Optional.ofNullable(this.scheduleOnlyWithChanges);
    }

    /**
     * Build start hour. Defaults to `0`. Valid values: `0 ~ 23`.
     * 
     */
    @Import(name="startHours")
    private @Nullable Output<Integer> startHours;

    /**
     * @return Build start hour. Defaults to `0`. Valid values: `0 ~ 23`.
     * 
     */
    public Optional<Output<Integer>> startHours() {
        return Optional.ofNullable(this.startHours);
    }

    /**
     * Build start minute. Defaults to `0`. Valid values: `0 ~ 59`.
     * 
     */
    @Import(name="startMinutes")
    private @Nullable Output<Integer> startMinutes;

    /**
     * @return Build start minute. Defaults to `0`. Valid values: `0 ~ 59`.
     * 
     */
    public Optional<Output<Integer>> startMinutes() {
        return Optional.ofNullable(this.startMinutes);
    }

    /**
     * Build time zone. Defaults to `(UTC) Coordinated Universal Time`. Valid values:
     * `(UTC-12:00) International Date Line West`,
     * `(UTC-11:00) Coordinated Universal Time-11`,
     * `(UTC-10:00) Aleutian Islands`,
     * `(UTC-10:00) Hawaii`,
     * `(UTC-09:30) Marquesas Islands`,
     * `(UTC-09:00) Alaska`,
     * `(UTC-09:00) Coordinated Universal Time-09`,
     * `(UTC-08:00) Baja California`,
     * `(UTC-08:00) Coordinated Universal Time-08`,
     * `(UTC-08:00) Pacific Time (US &amp;Canada)`,
     * `(UTC-07:00) Arizona`,
     * `(UTC-07:00) Chihuahua, La Paz, Mazatlan`,
     * `(UTC-07:00) Mountain Time (US &amp;Canada)`,
     * `(UTC-07:00) Yukon`,
     * `(UTC-06:00) Central America`,
     * `(UTC-06:00) Central Time (US &amp;Canada)`,
     * `(UTC-06:00) Easter Island`,
     * `(UTC-06:00) Guadalajara, Mexico City, Monterrey`,
     * `(UTC-06:00) Saskatchewan`,
     * `(UTC-05:00) Bogota, Lima, Quito, Rio Branco`,
     * `(UTC-05:00) Chetumal`,
     * `(UTC-05:00) Eastern Time (US &amp;Canada)`,
     * `(UTC-05:00) Haiti`,
     * `(UTC-05:00) Havana`,
     * `(UTC-05:00) Indiana (East)`,
     * `(UTC-05:00) Turks and Caicos`,
     * `(UTC-04:00) Asuncion`,
     * `(UTC-04:00) Atlantic Time (Canada)`,
     * `(UTC-04:00) Caracas`,
     * `(UTC-04:00) Cuiaba`,
     * `(UTC-04:00) Georgetown, La Paz, Manaus, San Juan`,
     * `(UTC-04:00) Santiago`,
     * `(UTC-03:30) Newfoundland`,
     * `(UTC-03:00) Araguaina`,
     * `(UTC-03:00) Brasilia`,
     * `(UTC-03:00) Cayenne, Fortaleza`,
     * `(UTC-03:00) City of Buenos Aires`,
     * `(UTC-03:00) Greenland`,
     * `(UTC-03:00) Montevideo`,
     * `(UTC-03:00) Punta Arenas`,
     * `(UTC-03:00) Saint Pierre and Miquelon`,
     * `(UTC-03:00) Salvador`,
     * `(UTC-02:00) Coordinated Universal Time-02`,
     * `(UTC-02:00) Mid-Atlantic - Old`,
     * `(UTC-01:00) Azores`,
     * `(UTC-01:00) Cabo Verde Is.`,
     * `(UTC) Coordinated Universal Time`,
     * `(UTC+00:00) Dublin, Edinburgh, Lisbon, London`,
     * `(UTC+00:00) Monrovia, Reykjavik`,
     * `(UTC+00:00) Sao Tome`,
     * `(UTC+01:00) Casablanca`,
     * `(UTC+01:00) Amsterdam, Berlin, Bern, Rome, Stockholm, Vienna`,
     * `(UTC+01:00) Belgrade, Bratislava, Budapest, Ljubljana, Prague`,
     * `(UTC+01:00) Brussels, Copenhagen, Madrid, Paris`,
     * `(UTC+01:00) Sarajevo, Skopje, Warsaw, Zagreb`,
     * `(UTC+01:00) West Central Africa`,
     * `(UTC+02:00) Amman`,
     * `(UTC+02:00) Athens, Bucharest`,
     * `(UTC+02:00) Beirut`,
     * `(UTC+02:00) Cairo`,
     * `(UTC+02:00) Chisinau`,
     * `(UTC+02:00) Damascus`,
     * `(UTC+02:00) Gaza, Hebron`,
     * `(UTC+02:00) Harare, Pretoria`,
     * `(UTC+02:00) Helsinki, Kyiv, Riga, Sofia, Tallinn, Vilnius`,
     * `(UTC+02:00) Jerusalem`,
     * `(UTC+02:00) Juba`,
     * `(UTC+02:00) Kaliningrad`,
     * `(UTC+02:00) Khartoum`,
     * `(UTC+02:00) Tripoli`,
     * `(UTC+02:00) Windhoek`,
     * `(UTC+03:00) Baghdad`,
     * `(UTC+03:00) Istanbul`,
     * `(UTC+03:00) Kuwait, Riyadh`,
     * `(UTC+03:00) Minsk`,
     * `(UTC+03:00) Moscow, St. Petersburg`,
     * `(UTC+03:00) Nairobi`,
     * `(UTC+03:00) Volgograd`,
     * `(UTC+03:30) Tehran`,
     * `(UTC+04:00) Abu Dhabi, Muscat`,
     * `(UTC+04:00) Astrakhan, Ulyanovsk`,
     * `(UTC+04:00) Baku`,
     * `(UTC+04:00) Izhevsk, Samara`,
     * `(UTC+04:00) Port Louis`,
     * `(UTC+04:00) Saratov`,
     * `(UTC+04:00) Tbilisi`,
     * `(UTC+04:00) Yerevan`,
     * `(UTC+04:30) Kabul`,
     * `(UTC+05:00) Ashgabat, Tashkent`,
     * `(UTC+05:00) Ekaterinburg`,
     * `(UTC+05:00) Islamabad, Karachi`,
     * `(UTC+05:00) Qyzylorda`,
     * `(UTC+05:30) Chennai, Kolkata, Mumbai, New Delhi`,
     * `(UTC+05:30) Sri Jayawardenepura`,
     * `(UTC+05:45) Kathmandu`,
     * `(UTC+06:00) Astana`,
     * `(UTC+06:00) Dhaka`,
     * `(UTC+06:00) Omsk`,
     * `(UTC+06:30) Yangon (Rangoon)`,
     * `(UTC+07:00) Bangkok, Hanoi, Jakarta`,
     * `(UTC+07:00) Barnaul, Gorno-Altaysk`,
     * `(UTC+07:00) Hovd`,
     * `(UTC+07:00) Krasnoyarsk`,
     * `(UTC+07:00) Novosibirsk`,
     * `(UTC+07:00) Tomsk`,
     * `(UTC+08:00) Beijing, Chongqing, Hong Kong, Urumqi`,
     * `(UTC+08:00) Irkutsk`,
     * `(UTC+08:00) Kuala Lumpur, Singapore`,
     * `(UTC+08:00) Perth`,
     * `(UTC+08:00) Taipei`,
     * `(UTC+08:00) Ulaanbaatar`,
     * `(UTC+08:45) Eucla`,
     * `(UTC+09:00) Chita`,
     * `(UTC+09:00) Osaka, Sapporo, Tokyo`,
     * `(UTC+09:00) Pyongyang`,
     * `(UTC+09:00) Seoul`,
     * `(UTC+09:00) Yakutsk`,
     * `(UTC+09:30) Adelaide`,
     * `(UTC+09:30) Darwin`,
     * `(UTC+10:00) Brisbane`,
     * `(UTC+10:00) Canberra, Melbourne, Sydney`,
     * `(UTC+10:00) Guam, Port Moresby`,
     * `(UTC+10:00) Hobart`,
     * `(UTC+10:00) Vladivostok`,
     * `(UTC+10:30) Lord Howe Island`,
     * `(UTC+11:00) Bougainville Island`,
     * `(UTC+11:00) Chokurdakh`,
     * `(UTC+11:00) Magadan`,
     * `(UTC+11:00) Norfolk Island`,
     * `(UTC+11:00) Sakhalin`,
     * `(UTC+11:00) Solomon Is., New Caledonia`,
     * `(UTC+12:00) Anadyr, Petropavlovsk-Kamchatsky`,
     * `(UTC+12:00) Auckland, Wellington`,
     * `(UTC+12:00) Coordinated Universal Time+12`,
     * `(UTC+12:00) Fiji`,
     * `(UTC+12:00) Petropavlovsk-Kamchatsky - Old`,
     * `(UTC+12:45) Chatham Islands`,
     * `(UTC+13:00) Coordinated Universal Time+13`,
     * `(UTC+13:00) Nuku&#39;alofa`,
     * `(UTC+13:00) Samoa`,
     * `(UTC+14:00) Kiritimati Island`.
     * 
     */
    @Import(name="timeZone")
    private @Nullable Output<String> timeZone;

    /**
     * @return Build time zone. Defaults to `(UTC) Coordinated Universal Time`. Valid values:
     * `(UTC-12:00) International Date Line West`,
     * `(UTC-11:00) Coordinated Universal Time-11`,
     * `(UTC-10:00) Aleutian Islands`,
     * `(UTC-10:00) Hawaii`,
     * `(UTC-09:30) Marquesas Islands`,
     * `(UTC-09:00) Alaska`,
     * `(UTC-09:00) Coordinated Universal Time-09`,
     * `(UTC-08:00) Baja California`,
     * `(UTC-08:00) Coordinated Universal Time-08`,
     * `(UTC-08:00) Pacific Time (US &amp;Canada)`,
     * `(UTC-07:00) Arizona`,
     * `(UTC-07:00) Chihuahua, La Paz, Mazatlan`,
     * `(UTC-07:00) Mountain Time (US &amp;Canada)`,
     * `(UTC-07:00) Yukon`,
     * `(UTC-06:00) Central America`,
     * `(UTC-06:00) Central Time (US &amp;Canada)`,
     * `(UTC-06:00) Easter Island`,
     * `(UTC-06:00) Guadalajara, Mexico City, Monterrey`,
     * `(UTC-06:00) Saskatchewan`,
     * `(UTC-05:00) Bogota, Lima, Quito, Rio Branco`,
     * `(UTC-05:00) Chetumal`,
     * `(UTC-05:00) Eastern Time (US &amp;Canada)`,
     * `(UTC-05:00) Haiti`,
     * `(UTC-05:00) Havana`,
     * `(UTC-05:00) Indiana (East)`,
     * `(UTC-05:00) Turks and Caicos`,
     * `(UTC-04:00) Asuncion`,
     * `(UTC-04:00) Atlantic Time (Canada)`,
     * `(UTC-04:00) Caracas`,
     * `(UTC-04:00) Cuiaba`,
     * `(UTC-04:00) Georgetown, La Paz, Manaus, San Juan`,
     * `(UTC-04:00) Santiago`,
     * `(UTC-03:30) Newfoundland`,
     * `(UTC-03:00) Araguaina`,
     * `(UTC-03:00) Brasilia`,
     * `(UTC-03:00) Cayenne, Fortaleza`,
     * `(UTC-03:00) City of Buenos Aires`,
     * `(UTC-03:00) Greenland`,
     * `(UTC-03:00) Montevideo`,
     * `(UTC-03:00) Punta Arenas`,
     * `(UTC-03:00) Saint Pierre and Miquelon`,
     * `(UTC-03:00) Salvador`,
     * `(UTC-02:00) Coordinated Universal Time-02`,
     * `(UTC-02:00) Mid-Atlantic - Old`,
     * `(UTC-01:00) Azores`,
     * `(UTC-01:00) Cabo Verde Is.`,
     * `(UTC) Coordinated Universal Time`,
     * `(UTC+00:00) Dublin, Edinburgh, Lisbon, London`,
     * `(UTC+00:00) Monrovia, Reykjavik`,
     * `(UTC+00:00) Sao Tome`,
     * `(UTC+01:00) Casablanca`,
     * `(UTC+01:00) Amsterdam, Berlin, Bern, Rome, Stockholm, Vienna`,
     * `(UTC+01:00) Belgrade, Bratislava, Budapest, Ljubljana, Prague`,
     * `(UTC+01:00) Brussels, Copenhagen, Madrid, Paris`,
     * `(UTC+01:00) Sarajevo, Skopje, Warsaw, Zagreb`,
     * `(UTC+01:00) West Central Africa`,
     * `(UTC+02:00) Amman`,
     * `(UTC+02:00) Athens, Bucharest`,
     * `(UTC+02:00) Beirut`,
     * `(UTC+02:00) Cairo`,
     * `(UTC+02:00) Chisinau`,
     * `(UTC+02:00) Damascus`,
     * `(UTC+02:00) Gaza, Hebron`,
     * `(UTC+02:00) Harare, Pretoria`,
     * `(UTC+02:00) Helsinki, Kyiv, Riga, Sofia, Tallinn, Vilnius`,
     * `(UTC+02:00) Jerusalem`,
     * `(UTC+02:00) Juba`,
     * `(UTC+02:00) Kaliningrad`,
     * `(UTC+02:00) Khartoum`,
     * `(UTC+02:00) Tripoli`,
     * `(UTC+02:00) Windhoek`,
     * `(UTC+03:00) Baghdad`,
     * `(UTC+03:00) Istanbul`,
     * `(UTC+03:00) Kuwait, Riyadh`,
     * `(UTC+03:00) Minsk`,
     * `(UTC+03:00) Moscow, St. Petersburg`,
     * `(UTC+03:00) Nairobi`,
     * `(UTC+03:00) Volgograd`,
     * `(UTC+03:30) Tehran`,
     * `(UTC+04:00) Abu Dhabi, Muscat`,
     * `(UTC+04:00) Astrakhan, Ulyanovsk`,
     * `(UTC+04:00) Baku`,
     * `(UTC+04:00) Izhevsk, Samara`,
     * `(UTC+04:00) Port Louis`,
     * `(UTC+04:00) Saratov`,
     * `(UTC+04:00) Tbilisi`,
     * `(UTC+04:00) Yerevan`,
     * `(UTC+04:30) Kabul`,
     * `(UTC+05:00) Ashgabat, Tashkent`,
     * `(UTC+05:00) Ekaterinburg`,
     * `(UTC+05:00) Islamabad, Karachi`,
     * `(UTC+05:00) Qyzylorda`,
     * `(UTC+05:30) Chennai, Kolkata, Mumbai, New Delhi`,
     * `(UTC+05:30) Sri Jayawardenepura`,
     * `(UTC+05:45) Kathmandu`,
     * `(UTC+06:00) Astana`,
     * `(UTC+06:00) Dhaka`,
     * `(UTC+06:00) Omsk`,
     * `(UTC+06:30) Yangon (Rangoon)`,
     * `(UTC+07:00) Bangkok, Hanoi, Jakarta`,
     * `(UTC+07:00) Barnaul, Gorno-Altaysk`,
     * `(UTC+07:00) Hovd`,
     * `(UTC+07:00) Krasnoyarsk`,
     * `(UTC+07:00) Novosibirsk`,
     * `(UTC+07:00) Tomsk`,
     * `(UTC+08:00) Beijing, Chongqing, Hong Kong, Urumqi`,
     * `(UTC+08:00) Irkutsk`,
     * `(UTC+08:00) Kuala Lumpur, Singapore`,
     * `(UTC+08:00) Perth`,
     * `(UTC+08:00) Taipei`,
     * `(UTC+08:00) Ulaanbaatar`,
     * `(UTC+08:45) Eucla`,
     * `(UTC+09:00) Chita`,
     * `(UTC+09:00) Osaka, Sapporo, Tokyo`,
     * `(UTC+09:00) Pyongyang`,
     * `(UTC+09:00) Seoul`,
     * `(UTC+09:00) Yakutsk`,
     * `(UTC+09:30) Adelaide`,
     * `(UTC+09:30) Darwin`,
     * `(UTC+10:00) Brisbane`,
     * `(UTC+10:00) Canberra, Melbourne, Sydney`,
     * `(UTC+10:00) Guam, Port Moresby`,
     * `(UTC+10:00) Hobart`,
     * `(UTC+10:00) Vladivostok`,
     * `(UTC+10:30) Lord Howe Island`,
     * `(UTC+11:00) Bougainville Island`,
     * `(UTC+11:00) Chokurdakh`,
     * `(UTC+11:00) Magadan`,
     * `(UTC+11:00) Norfolk Island`,
     * `(UTC+11:00) Sakhalin`,
     * `(UTC+11:00) Solomon Is., New Caledonia`,
     * `(UTC+12:00) Anadyr, Petropavlovsk-Kamchatsky`,
     * `(UTC+12:00) Auckland, Wellington`,
     * `(UTC+12:00) Coordinated Universal Time+12`,
     * `(UTC+12:00) Fiji`,
     * `(UTC+12:00) Petropavlovsk-Kamchatsky - Old`,
     * `(UTC+12:45) Chatham Islands`,
     * `(UTC+13:00) Coordinated Universal Time+13`,
     * `(UTC+13:00) Nuku&#39;alofa`,
     * `(UTC+13:00) Samoa`,
     * `(UTC+14:00) Kiritimati Island`.
     * 
     */
    public Optional<Output<String>> timeZone() {
        return Optional.ofNullable(this.timeZone);
    }

    private BuildDefinitionScheduleArgs() {}

    private BuildDefinitionScheduleArgs(BuildDefinitionScheduleArgs $) {
        this.branchFilters = $.branchFilters;
        this.daysToBuilds = $.daysToBuilds;
        this.scheduleJobId = $.scheduleJobId;
        this.scheduleOnlyWithChanges = $.scheduleOnlyWithChanges;
        this.startHours = $.startHours;
        this.startMinutes = $.startMinutes;
        this.timeZone = $.timeZone;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(BuildDefinitionScheduleArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private BuildDefinitionScheduleArgs $;

        public Builder() {
            $ = new BuildDefinitionScheduleArgs();
        }

        public Builder(BuildDefinitionScheduleArgs defaults) {
            $ = new BuildDefinitionScheduleArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param branchFilters block supports the following:
         * 
         * @return builder
         * 
         */
        public Builder branchFilters(@Nullable Output<List<BuildDefinitionScheduleBranchFilterArgs>> branchFilters) {
            $.branchFilters = branchFilters;
            return this;
        }

        /**
         * @param branchFilters block supports the following:
         * 
         * @return builder
         * 
         */
        public Builder branchFilters(List<BuildDefinitionScheduleBranchFilterArgs> branchFilters) {
            return branchFilters(Output.of(branchFilters));
        }

        /**
         * @param branchFilters block supports the following:
         * 
         * @return builder
         * 
         */
        public Builder branchFilters(BuildDefinitionScheduleBranchFilterArgs... branchFilters) {
            return branchFilters(List.of(branchFilters));
        }

        /**
         * @param daysToBuilds When to build. Valid values: `Mon`, `Tue`, `Wed`, `Thu`, `Fri`, `Sat`, `Sun`.
         * 
         * @return builder
         * 
         */
        public Builder daysToBuilds(Output<List<String>> daysToBuilds) {
            $.daysToBuilds = daysToBuilds;
            return this;
        }

        /**
         * @param daysToBuilds When to build. Valid values: `Mon`, `Tue`, `Wed`, `Thu`, `Fri`, `Sat`, `Sun`.
         * 
         * @return builder
         * 
         */
        public Builder daysToBuilds(List<String> daysToBuilds) {
            return daysToBuilds(Output.of(daysToBuilds));
        }

        /**
         * @param daysToBuilds When to build. Valid values: `Mon`, `Tue`, `Wed`, `Thu`, `Fri`, `Sat`, `Sun`.
         * 
         * @return builder
         * 
         */
        public Builder daysToBuilds(String... daysToBuilds) {
            return daysToBuilds(List.of(daysToBuilds));
        }

        /**
         * @param scheduleJobId The ID of the schedule job
         * 
         * @return builder
         * 
         */
        public Builder scheduleJobId(@Nullable Output<String> scheduleJobId) {
            $.scheduleJobId = scheduleJobId;
            return this;
        }

        /**
         * @param scheduleJobId The ID of the schedule job
         * 
         * @return builder
         * 
         */
        public Builder scheduleJobId(String scheduleJobId) {
            return scheduleJobId(Output.of(scheduleJobId));
        }

        /**
         * @param scheduleOnlyWithChanges Schedule builds if the source or pipeline has changed. Defaults to `true`.
         * 
         * @return builder
         * 
         */
        public Builder scheduleOnlyWithChanges(@Nullable Output<Boolean> scheduleOnlyWithChanges) {
            $.scheduleOnlyWithChanges = scheduleOnlyWithChanges;
            return this;
        }

        /**
         * @param scheduleOnlyWithChanges Schedule builds if the source or pipeline has changed. Defaults to `true`.
         * 
         * @return builder
         * 
         */
        public Builder scheduleOnlyWithChanges(Boolean scheduleOnlyWithChanges) {
            return scheduleOnlyWithChanges(Output.of(scheduleOnlyWithChanges));
        }

        /**
         * @param startHours Build start hour. Defaults to `0`. Valid values: `0 ~ 23`.
         * 
         * @return builder
         * 
         */
        public Builder startHours(@Nullable Output<Integer> startHours) {
            $.startHours = startHours;
            return this;
        }

        /**
         * @param startHours Build start hour. Defaults to `0`. Valid values: `0 ~ 23`.
         * 
         * @return builder
         * 
         */
        public Builder startHours(Integer startHours) {
            return startHours(Output.of(startHours));
        }

        /**
         * @param startMinutes Build start minute. Defaults to `0`. Valid values: `0 ~ 59`.
         * 
         * @return builder
         * 
         */
        public Builder startMinutes(@Nullable Output<Integer> startMinutes) {
            $.startMinutes = startMinutes;
            return this;
        }

        /**
         * @param startMinutes Build start minute. Defaults to `0`. Valid values: `0 ~ 59`.
         * 
         * @return builder
         * 
         */
        public Builder startMinutes(Integer startMinutes) {
            return startMinutes(Output.of(startMinutes));
        }

        /**
         * @param timeZone Build time zone. Defaults to `(UTC) Coordinated Universal Time`. Valid values:
         * `(UTC-12:00) International Date Line West`,
         * `(UTC-11:00) Coordinated Universal Time-11`,
         * `(UTC-10:00) Aleutian Islands`,
         * `(UTC-10:00) Hawaii`,
         * `(UTC-09:30) Marquesas Islands`,
         * `(UTC-09:00) Alaska`,
         * `(UTC-09:00) Coordinated Universal Time-09`,
         * `(UTC-08:00) Baja California`,
         * `(UTC-08:00) Coordinated Universal Time-08`,
         * `(UTC-08:00) Pacific Time (US &amp;Canada)`,
         * `(UTC-07:00) Arizona`,
         * `(UTC-07:00) Chihuahua, La Paz, Mazatlan`,
         * `(UTC-07:00) Mountain Time (US &amp;Canada)`,
         * `(UTC-07:00) Yukon`,
         * `(UTC-06:00) Central America`,
         * `(UTC-06:00) Central Time (US &amp;Canada)`,
         * `(UTC-06:00) Easter Island`,
         * `(UTC-06:00) Guadalajara, Mexico City, Monterrey`,
         * `(UTC-06:00) Saskatchewan`,
         * `(UTC-05:00) Bogota, Lima, Quito, Rio Branco`,
         * `(UTC-05:00) Chetumal`,
         * `(UTC-05:00) Eastern Time (US &amp;Canada)`,
         * `(UTC-05:00) Haiti`,
         * `(UTC-05:00) Havana`,
         * `(UTC-05:00) Indiana (East)`,
         * `(UTC-05:00) Turks and Caicos`,
         * `(UTC-04:00) Asuncion`,
         * `(UTC-04:00) Atlantic Time (Canada)`,
         * `(UTC-04:00) Caracas`,
         * `(UTC-04:00) Cuiaba`,
         * `(UTC-04:00) Georgetown, La Paz, Manaus, San Juan`,
         * `(UTC-04:00) Santiago`,
         * `(UTC-03:30) Newfoundland`,
         * `(UTC-03:00) Araguaina`,
         * `(UTC-03:00) Brasilia`,
         * `(UTC-03:00) Cayenne, Fortaleza`,
         * `(UTC-03:00) City of Buenos Aires`,
         * `(UTC-03:00) Greenland`,
         * `(UTC-03:00) Montevideo`,
         * `(UTC-03:00) Punta Arenas`,
         * `(UTC-03:00) Saint Pierre and Miquelon`,
         * `(UTC-03:00) Salvador`,
         * `(UTC-02:00) Coordinated Universal Time-02`,
         * `(UTC-02:00) Mid-Atlantic - Old`,
         * `(UTC-01:00) Azores`,
         * `(UTC-01:00) Cabo Verde Is.`,
         * `(UTC) Coordinated Universal Time`,
         * `(UTC+00:00) Dublin, Edinburgh, Lisbon, London`,
         * `(UTC+00:00) Monrovia, Reykjavik`,
         * `(UTC+00:00) Sao Tome`,
         * `(UTC+01:00) Casablanca`,
         * `(UTC+01:00) Amsterdam, Berlin, Bern, Rome, Stockholm, Vienna`,
         * `(UTC+01:00) Belgrade, Bratislava, Budapest, Ljubljana, Prague`,
         * `(UTC+01:00) Brussels, Copenhagen, Madrid, Paris`,
         * `(UTC+01:00) Sarajevo, Skopje, Warsaw, Zagreb`,
         * `(UTC+01:00) West Central Africa`,
         * `(UTC+02:00) Amman`,
         * `(UTC+02:00) Athens, Bucharest`,
         * `(UTC+02:00) Beirut`,
         * `(UTC+02:00) Cairo`,
         * `(UTC+02:00) Chisinau`,
         * `(UTC+02:00) Damascus`,
         * `(UTC+02:00) Gaza, Hebron`,
         * `(UTC+02:00) Harare, Pretoria`,
         * `(UTC+02:00) Helsinki, Kyiv, Riga, Sofia, Tallinn, Vilnius`,
         * `(UTC+02:00) Jerusalem`,
         * `(UTC+02:00) Juba`,
         * `(UTC+02:00) Kaliningrad`,
         * `(UTC+02:00) Khartoum`,
         * `(UTC+02:00) Tripoli`,
         * `(UTC+02:00) Windhoek`,
         * `(UTC+03:00) Baghdad`,
         * `(UTC+03:00) Istanbul`,
         * `(UTC+03:00) Kuwait, Riyadh`,
         * `(UTC+03:00) Minsk`,
         * `(UTC+03:00) Moscow, St. Petersburg`,
         * `(UTC+03:00) Nairobi`,
         * `(UTC+03:00) Volgograd`,
         * `(UTC+03:30) Tehran`,
         * `(UTC+04:00) Abu Dhabi, Muscat`,
         * `(UTC+04:00) Astrakhan, Ulyanovsk`,
         * `(UTC+04:00) Baku`,
         * `(UTC+04:00) Izhevsk, Samara`,
         * `(UTC+04:00) Port Louis`,
         * `(UTC+04:00) Saratov`,
         * `(UTC+04:00) Tbilisi`,
         * `(UTC+04:00) Yerevan`,
         * `(UTC+04:30) Kabul`,
         * `(UTC+05:00) Ashgabat, Tashkent`,
         * `(UTC+05:00) Ekaterinburg`,
         * `(UTC+05:00) Islamabad, Karachi`,
         * `(UTC+05:00) Qyzylorda`,
         * `(UTC+05:30) Chennai, Kolkata, Mumbai, New Delhi`,
         * `(UTC+05:30) Sri Jayawardenepura`,
         * `(UTC+05:45) Kathmandu`,
         * `(UTC+06:00) Astana`,
         * `(UTC+06:00) Dhaka`,
         * `(UTC+06:00) Omsk`,
         * `(UTC+06:30) Yangon (Rangoon)`,
         * `(UTC+07:00) Bangkok, Hanoi, Jakarta`,
         * `(UTC+07:00) Barnaul, Gorno-Altaysk`,
         * `(UTC+07:00) Hovd`,
         * `(UTC+07:00) Krasnoyarsk`,
         * `(UTC+07:00) Novosibirsk`,
         * `(UTC+07:00) Tomsk`,
         * `(UTC+08:00) Beijing, Chongqing, Hong Kong, Urumqi`,
         * `(UTC+08:00) Irkutsk`,
         * `(UTC+08:00) Kuala Lumpur, Singapore`,
         * `(UTC+08:00) Perth`,
         * `(UTC+08:00) Taipei`,
         * `(UTC+08:00) Ulaanbaatar`,
         * `(UTC+08:45) Eucla`,
         * `(UTC+09:00) Chita`,
         * `(UTC+09:00) Osaka, Sapporo, Tokyo`,
         * `(UTC+09:00) Pyongyang`,
         * `(UTC+09:00) Seoul`,
         * `(UTC+09:00) Yakutsk`,
         * `(UTC+09:30) Adelaide`,
         * `(UTC+09:30) Darwin`,
         * `(UTC+10:00) Brisbane`,
         * `(UTC+10:00) Canberra, Melbourne, Sydney`,
         * `(UTC+10:00) Guam, Port Moresby`,
         * `(UTC+10:00) Hobart`,
         * `(UTC+10:00) Vladivostok`,
         * `(UTC+10:30) Lord Howe Island`,
         * `(UTC+11:00) Bougainville Island`,
         * `(UTC+11:00) Chokurdakh`,
         * `(UTC+11:00) Magadan`,
         * `(UTC+11:00) Norfolk Island`,
         * `(UTC+11:00) Sakhalin`,
         * `(UTC+11:00) Solomon Is., New Caledonia`,
         * `(UTC+12:00) Anadyr, Petropavlovsk-Kamchatsky`,
         * `(UTC+12:00) Auckland, Wellington`,
         * `(UTC+12:00) Coordinated Universal Time+12`,
         * `(UTC+12:00) Fiji`,
         * `(UTC+12:00) Petropavlovsk-Kamchatsky - Old`,
         * `(UTC+12:45) Chatham Islands`,
         * `(UTC+13:00) Coordinated Universal Time+13`,
         * `(UTC+13:00) Nuku&#39;alofa`,
         * `(UTC+13:00) Samoa`,
         * `(UTC+14:00) Kiritimati Island`.
         * 
         * @return builder
         * 
         */
        public Builder timeZone(@Nullable Output<String> timeZone) {
            $.timeZone = timeZone;
            return this;
        }

        /**
         * @param timeZone Build time zone. Defaults to `(UTC) Coordinated Universal Time`. Valid values:
         * `(UTC-12:00) International Date Line West`,
         * `(UTC-11:00) Coordinated Universal Time-11`,
         * `(UTC-10:00) Aleutian Islands`,
         * `(UTC-10:00) Hawaii`,
         * `(UTC-09:30) Marquesas Islands`,
         * `(UTC-09:00) Alaska`,
         * `(UTC-09:00) Coordinated Universal Time-09`,
         * `(UTC-08:00) Baja California`,
         * `(UTC-08:00) Coordinated Universal Time-08`,
         * `(UTC-08:00) Pacific Time (US &amp;Canada)`,
         * `(UTC-07:00) Arizona`,
         * `(UTC-07:00) Chihuahua, La Paz, Mazatlan`,
         * `(UTC-07:00) Mountain Time (US &amp;Canada)`,
         * `(UTC-07:00) Yukon`,
         * `(UTC-06:00) Central America`,
         * `(UTC-06:00) Central Time (US &amp;Canada)`,
         * `(UTC-06:00) Easter Island`,
         * `(UTC-06:00) Guadalajara, Mexico City, Monterrey`,
         * `(UTC-06:00) Saskatchewan`,
         * `(UTC-05:00) Bogota, Lima, Quito, Rio Branco`,
         * `(UTC-05:00) Chetumal`,
         * `(UTC-05:00) Eastern Time (US &amp;Canada)`,
         * `(UTC-05:00) Haiti`,
         * `(UTC-05:00) Havana`,
         * `(UTC-05:00) Indiana (East)`,
         * `(UTC-05:00) Turks and Caicos`,
         * `(UTC-04:00) Asuncion`,
         * `(UTC-04:00) Atlantic Time (Canada)`,
         * `(UTC-04:00) Caracas`,
         * `(UTC-04:00) Cuiaba`,
         * `(UTC-04:00) Georgetown, La Paz, Manaus, San Juan`,
         * `(UTC-04:00) Santiago`,
         * `(UTC-03:30) Newfoundland`,
         * `(UTC-03:00) Araguaina`,
         * `(UTC-03:00) Brasilia`,
         * `(UTC-03:00) Cayenne, Fortaleza`,
         * `(UTC-03:00) City of Buenos Aires`,
         * `(UTC-03:00) Greenland`,
         * `(UTC-03:00) Montevideo`,
         * `(UTC-03:00) Punta Arenas`,
         * `(UTC-03:00) Saint Pierre and Miquelon`,
         * `(UTC-03:00) Salvador`,
         * `(UTC-02:00) Coordinated Universal Time-02`,
         * `(UTC-02:00) Mid-Atlantic - Old`,
         * `(UTC-01:00) Azores`,
         * `(UTC-01:00) Cabo Verde Is.`,
         * `(UTC) Coordinated Universal Time`,
         * `(UTC+00:00) Dublin, Edinburgh, Lisbon, London`,
         * `(UTC+00:00) Monrovia, Reykjavik`,
         * `(UTC+00:00) Sao Tome`,
         * `(UTC+01:00) Casablanca`,
         * `(UTC+01:00) Amsterdam, Berlin, Bern, Rome, Stockholm, Vienna`,
         * `(UTC+01:00) Belgrade, Bratislava, Budapest, Ljubljana, Prague`,
         * `(UTC+01:00) Brussels, Copenhagen, Madrid, Paris`,
         * `(UTC+01:00) Sarajevo, Skopje, Warsaw, Zagreb`,
         * `(UTC+01:00) West Central Africa`,
         * `(UTC+02:00) Amman`,
         * `(UTC+02:00) Athens, Bucharest`,
         * `(UTC+02:00) Beirut`,
         * `(UTC+02:00) Cairo`,
         * `(UTC+02:00) Chisinau`,
         * `(UTC+02:00) Damascus`,
         * `(UTC+02:00) Gaza, Hebron`,
         * `(UTC+02:00) Harare, Pretoria`,
         * `(UTC+02:00) Helsinki, Kyiv, Riga, Sofia, Tallinn, Vilnius`,
         * `(UTC+02:00) Jerusalem`,
         * `(UTC+02:00) Juba`,
         * `(UTC+02:00) Kaliningrad`,
         * `(UTC+02:00) Khartoum`,
         * `(UTC+02:00) Tripoli`,
         * `(UTC+02:00) Windhoek`,
         * `(UTC+03:00) Baghdad`,
         * `(UTC+03:00) Istanbul`,
         * `(UTC+03:00) Kuwait, Riyadh`,
         * `(UTC+03:00) Minsk`,
         * `(UTC+03:00) Moscow, St. Petersburg`,
         * `(UTC+03:00) Nairobi`,
         * `(UTC+03:00) Volgograd`,
         * `(UTC+03:30) Tehran`,
         * `(UTC+04:00) Abu Dhabi, Muscat`,
         * `(UTC+04:00) Astrakhan, Ulyanovsk`,
         * `(UTC+04:00) Baku`,
         * `(UTC+04:00) Izhevsk, Samara`,
         * `(UTC+04:00) Port Louis`,
         * `(UTC+04:00) Saratov`,
         * `(UTC+04:00) Tbilisi`,
         * `(UTC+04:00) Yerevan`,
         * `(UTC+04:30) Kabul`,
         * `(UTC+05:00) Ashgabat, Tashkent`,
         * `(UTC+05:00) Ekaterinburg`,
         * `(UTC+05:00) Islamabad, Karachi`,
         * `(UTC+05:00) Qyzylorda`,
         * `(UTC+05:30) Chennai, Kolkata, Mumbai, New Delhi`,
         * `(UTC+05:30) Sri Jayawardenepura`,
         * `(UTC+05:45) Kathmandu`,
         * `(UTC+06:00) Astana`,
         * `(UTC+06:00) Dhaka`,
         * `(UTC+06:00) Omsk`,
         * `(UTC+06:30) Yangon (Rangoon)`,
         * `(UTC+07:00) Bangkok, Hanoi, Jakarta`,
         * `(UTC+07:00) Barnaul, Gorno-Altaysk`,
         * `(UTC+07:00) Hovd`,
         * `(UTC+07:00) Krasnoyarsk`,
         * `(UTC+07:00) Novosibirsk`,
         * `(UTC+07:00) Tomsk`,
         * `(UTC+08:00) Beijing, Chongqing, Hong Kong, Urumqi`,
         * `(UTC+08:00) Irkutsk`,
         * `(UTC+08:00) Kuala Lumpur, Singapore`,
         * `(UTC+08:00) Perth`,
         * `(UTC+08:00) Taipei`,
         * `(UTC+08:00) Ulaanbaatar`,
         * `(UTC+08:45) Eucla`,
         * `(UTC+09:00) Chita`,
         * `(UTC+09:00) Osaka, Sapporo, Tokyo`,
         * `(UTC+09:00) Pyongyang`,
         * `(UTC+09:00) Seoul`,
         * `(UTC+09:00) Yakutsk`,
         * `(UTC+09:30) Adelaide`,
         * `(UTC+09:30) Darwin`,
         * `(UTC+10:00) Brisbane`,
         * `(UTC+10:00) Canberra, Melbourne, Sydney`,
         * `(UTC+10:00) Guam, Port Moresby`,
         * `(UTC+10:00) Hobart`,
         * `(UTC+10:00) Vladivostok`,
         * `(UTC+10:30) Lord Howe Island`,
         * `(UTC+11:00) Bougainville Island`,
         * `(UTC+11:00) Chokurdakh`,
         * `(UTC+11:00) Magadan`,
         * `(UTC+11:00) Norfolk Island`,
         * `(UTC+11:00) Sakhalin`,
         * `(UTC+11:00) Solomon Is., New Caledonia`,
         * `(UTC+12:00) Anadyr, Petropavlovsk-Kamchatsky`,
         * `(UTC+12:00) Auckland, Wellington`,
         * `(UTC+12:00) Coordinated Universal Time+12`,
         * `(UTC+12:00) Fiji`,
         * `(UTC+12:00) Petropavlovsk-Kamchatsky - Old`,
         * `(UTC+12:45) Chatham Islands`,
         * `(UTC+13:00) Coordinated Universal Time+13`,
         * `(UTC+13:00) Nuku&#39;alofa`,
         * `(UTC+13:00) Samoa`,
         * `(UTC+14:00) Kiritimati Island`.
         * 
         * @return builder
         * 
         */
        public Builder timeZone(String timeZone) {
            return timeZone(Output.of(timeZone));
        }

        public BuildDefinitionScheduleArgs build() {
            if ($.daysToBuilds == null) {
                throw new MissingRequiredPropertyException("BuildDefinitionScheduleArgs", "daysToBuilds");
            }
            return $;
        }
    }

}
