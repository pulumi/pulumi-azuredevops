// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.outputs;

import com.pulumi.azuredevops.outputs.BuildDefinitionScheduleBranchFilter;
import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class BuildDefinitionSchedule {
    /**
     * @return A `branch_filter` block as defined below.
     * 
     */
    private List<BuildDefinitionScheduleBranchFilter> branchFilters;
    /**
     * @return When to build. Possible values are: `Mon`, `Tue`, `Wed`, `Thu`, `Fri`, `Sat`, `Sun`.
     * 
     */
    private List<String> daysToBuilds;
    /**
     * @return The ID of the schedule job
     * 
     */
    private @Nullable String scheduleJobId;
    /**
     * @return Schedule builds if the source or pipeline has changed. Defaults to `true`.
     * 
     */
    private @Nullable Boolean scheduleOnlyWithChanges;
    /**
     * @return Build start hour. Possible values are: `0 ~ 23`. Defaults to `0`.
     * 
     */
    private @Nullable Integer startHours;
    /**
     * @return Build start minute. Possible values are: `0 ~ 59`. Defaults to `0`.
     * 
     */
    private @Nullable Integer startMinutes;
    /**
     * @return Build time zone. Defaults to `(UTC) Coordinated Universal Time`. Possible values are:
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
    private @Nullable String timeZone;

    private BuildDefinitionSchedule() {}
    /**
     * @return A `branch_filter` block as defined below.
     * 
     */
    public List<BuildDefinitionScheduleBranchFilter> branchFilters() {
        return this.branchFilters;
    }
    /**
     * @return When to build. Possible values are: `Mon`, `Tue`, `Wed`, `Thu`, `Fri`, `Sat`, `Sun`.
     * 
     */
    public List<String> daysToBuilds() {
        return this.daysToBuilds;
    }
    /**
     * @return The ID of the schedule job
     * 
     */
    public Optional<String> scheduleJobId() {
        return Optional.ofNullable(this.scheduleJobId);
    }
    /**
     * @return Schedule builds if the source or pipeline has changed. Defaults to `true`.
     * 
     */
    public Optional<Boolean> scheduleOnlyWithChanges() {
        return Optional.ofNullable(this.scheduleOnlyWithChanges);
    }
    /**
     * @return Build start hour. Possible values are: `0 ~ 23`. Defaults to `0`.
     * 
     */
    public Optional<Integer> startHours() {
        return Optional.ofNullable(this.startHours);
    }
    /**
     * @return Build start minute. Possible values are: `0 ~ 59`. Defaults to `0`.
     * 
     */
    public Optional<Integer> startMinutes() {
        return Optional.ofNullable(this.startMinutes);
    }
    /**
     * @return Build time zone. Defaults to `(UTC) Coordinated Universal Time`. Possible values are:
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
    public Optional<String> timeZone() {
        return Optional.ofNullable(this.timeZone);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(BuildDefinitionSchedule defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<BuildDefinitionScheduleBranchFilter> branchFilters;
        private List<String> daysToBuilds;
        private @Nullable String scheduleJobId;
        private @Nullable Boolean scheduleOnlyWithChanges;
        private @Nullable Integer startHours;
        private @Nullable Integer startMinutes;
        private @Nullable String timeZone;
        public Builder() {}
        public Builder(BuildDefinitionSchedule defaults) {
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
        public Builder branchFilters(List<BuildDefinitionScheduleBranchFilter> branchFilters) {
            if (branchFilters == null) {
              throw new MissingRequiredPropertyException("BuildDefinitionSchedule", "branchFilters");
            }
            this.branchFilters = branchFilters;
            return this;
        }
        public Builder branchFilters(BuildDefinitionScheduleBranchFilter... branchFilters) {
            return branchFilters(List.of(branchFilters));
        }
        @CustomType.Setter
        public Builder daysToBuilds(List<String> daysToBuilds) {
            if (daysToBuilds == null) {
              throw new MissingRequiredPropertyException("BuildDefinitionSchedule", "daysToBuilds");
            }
            this.daysToBuilds = daysToBuilds;
            return this;
        }
        public Builder daysToBuilds(String... daysToBuilds) {
            return daysToBuilds(List.of(daysToBuilds));
        }
        @CustomType.Setter
        public Builder scheduleJobId(@Nullable String scheduleJobId) {

            this.scheduleJobId = scheduleJobId;
            return this;
        }
        @CustomType.Setter
        public Builder scheduleOnlyWithChanges(@Nullable Boolean scheduleOnlyWithChanges) {

            this.scheduleOnlyWithChanges = scheduleOnlyWithChanges;
            return this;
        }
        @CustomType.Setter
        public Builder startHours(@Nullable Integer startHours) {

            this.startHours = startHours;
            return this;
        }
        @CustomType.Setter
        public Builder startMinutes(@Nullable Integer startMinutes) {

            this.startMinutes = startMinutes;
            return this;
        }
        @CustomType.Setter
        public Builder timeZone(@Nullable String timeZone) {

            this.timeZone = timeZone;
            return this;
        }
        public BuildDefinitionSchedule build() {
            final var _resultValue = new BuildDefinitionSchedule();
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
