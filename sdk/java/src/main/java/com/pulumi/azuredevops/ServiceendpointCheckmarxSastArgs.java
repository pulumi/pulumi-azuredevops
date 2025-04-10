// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ServiceendpointCheckmarxSastArgs extends com.pulumi.resources.ResourceArgs {

    public static final ServiceendpointCheckmarxSastArgs Empty = new ServiceendpointCheckmarxSastArgs();

    @Import(name="description")
    private @Nullable Output<String> description;

    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The password of the Checkmarx SAST.
     * 
     */
    @Import(name="password", required=true)
    private Output<String> password;

    /**
     * @return The password of the Checkmarx SAST.
     * 
     */
    public Output<String> password() {
        return this.password;
    }

    /**
     * Predefined sets of queries that you can select when Creating, Configuring and Branching Projects. Predefined presets are provided by Checkmarx and you can configure your own. You can also import and export presets (on the server).In Service Connection if preset(optional) value is added, then it will igonres Preset available in pipeline and uses preset available in service connection only.If Preset is blank in service connection then it will use pipelines preset.
     * 
     */
    @Import(name="preset")
    private @Nullable Output<String> preset;

    /**
     * @return Predefined sets of queries that you can select when Creating, Configuring and Branching Projects. Predefined presets are provided by Checkmarx and you can configure your own. You can also import and export presets (on the server).In Service Connection if preset(optional) value is added, then it will igonres Preset available in pipeline and uses preset available in service connection only.If Preset is blank in service connection then it will use pipelines preset.
     * 
     */
    public Optional<Output<String>> preset() {
        return Optional.ofNullable(this.preset);
    }

    /**
     * The ID of the project.
     * 
     */
    @Import(name="projectId", required=true)
    private Output<String> projectId;

    /**
     * @return The ID of the project.
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }

    /**
     * The Server URL of the Checkmarx SAST.
     * 
     */
    @Import(name="serverUrl", required=true)
    private Output<String> serverUrl;

    /**
     * @return The Server URL of the Checkmarx SAST.
     * 
     */
    public Output<String> serverUrl() {
        return this.serverUrl;
    }

    /**
     * The Service Endpoint name.
     * 
     */
    @Import(name="serviceEndpointName", required=true)
    private Output<String> serviceEndpointName;

    /**
     * @return The Service Endpoint name.
     * 
     */
    public Output<String> serviceEndpointName() {
        return this.serviceEndpointName;
    }

    /**
     * The full team name of the Checkmarx.
     * 
     */
    @Import(name="team")
    private @Nullable Output<String> team;

    /**
     * @return The full team name of the Checkmarx.
     * 
     */
    public Optional<Output<String>> team() {
        return Optional.ofNullable(this.team);
    }

    /**
     * The username of the Checkmarx SAST.
     * 
     */
    @Import(name="username", required=true)
    private Output<String> username;

    /**
     * @return The username of the Checkmarx SAST.
     * 
     */
    public Output<String> username() {
        return this.username;
    }

    private ServiceendpointCheckmarxSastArgs() {}

    private ServiceendpointCheckmarxSastArgs(ServiceendpointCheckmarxSastArgs $) {
        this.description = $.description;
        this.password = $.password;
        this.preset = $.preset;
        this.projectId = $.projectId;
        this.serverUrl = $.serverUrl;
        this.serviceEndpointName = $.serviceEndpointName;
        this.team = $.team;
        this.username = $.username;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServiceendpointCheckmarxSastArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServiceendpointCheckmarxSastArgs $;

        public Builder() {
            $ = new ServiceendpointCheckmarxSastArgs();
        }

        public Builder(ServiceendpointCheckmarxSastArgs defaults) {
            $ = new ServiceendpointCheckmarxSastArgs(Objects.requireNonNull(defaults));
        }

        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param password The password of the Checkmarx SAST.
         * 
         * @return builder
         * 
         */
        public Builder password(Output<String> password) {
            $.password = password;
            return this;
        }

        /**
         * @param password The password of the Checkmarx SAST.
         * 
         * @return builder
         * 
         */
        public Builder password(String password) {
            return password(Output.of(password));
        }

        /**
         * @param preset Predefined sets of queries that you can select when Creating, Configuring and Branching Projects. Predefined presets are provided by Checkmarx and you can configure your own. You can also import and export presets (on the server).In Service Connection if preset(optional) value is added, then it will igonres Preset available in pipeline and uses preset available in service connection only.If Preset is blank in service connection then it will use pipelines preset.
         * 
         * @return builder
         * 
         */
        public Builder preset(@Nullable Output<String> preset) {
            $.preset = preset;
            return this;
        }

        /**
         * @param preset Predefined sets of queries that you can select when Creating, Configuring and Branching Projects. Predefined presets are provided by Checkmarx and you can configure your own. You can also import and export presets (on the server).In Service Connection if preset(optional) value is added, then it will igonres Preset available in pipeline and uses preset available in service connection only.If Preset is blank in service connection then it will use pipelines preset.
         * 
         * @return builder
         * 
         */
        public Builder preset(String preset) {
            return preset(Output.of(preset));
        }

        /**
         * @param projectId The ID of the project.
         * 
         * @return builder
         * 
         */
        public Builder projectId(Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param projectId The ID of the project.
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
        }

        /**
         * @param serverUrl The Server URL of the Checkmarx SAST.
         * 
         * @return builder
         * 
         */
        public Builder serverUrl(Output<String> serverUrl) {
            $.serverUrl = serverUrl;
            return this;
        }

        /**
         * @param serverUrl The Server URL of the Checkmarx SAST.
         * 
         * @return builder
         * 
         */
        public Builder serverUrl(String serverUrl) {
            return serverUrl(Output.of(serverUrl));
        }

        /**
         * @param serviceEndpointName The Service Endpoint name.
         * 
         * @return builder
         * 
         */
        public Builder serviceEndpointName(Output<String> serviceEndpointName) {
            $.serviceEndpointName = serviceEndpointName;
            return this;
        }

        /**
         * @param serviceEndpointName The Service Endpoint name.
         * 
         * @return builder
         * 
         */
        public Builder serviceEndpointName(String serviceEndpointName) {
            return serviceEndpointName(Output.of(serviceEndpointName));
        }

        /**
         * @param team The full team name of the Checkmarx.
         * 
         * @return builder
         * 
         */
        public Builder team(@Nullable Output<String> team) {
            $.team = team;
            return this;
        }

        /**
         * @param team The full team name of the Checkmarx.
         * 
         * @return builder
         * 
         */
        public Builder team(String team) {
            return team(Output.of(team));
        }

        /**
         * @param username The username of the Checkmarx SAST.
         * 
         * @return builder
         * 
         */
        public Builder username(Output<String> username) {
            $.username = username;
            return this;
        }

        /**
         * @param username The username of the Checkmarx SAST.
         * 
         * @return builder
         * 
         */
        public Builder username(String username) {
            return username(Output.of(username));
        }

        public ServiceendpointCheckmarxSastArgs build() {
            if ($.password == null) {
                throw new MissingRequiredPropertyException("ServiceendpointCheckmarxSastArgs", "password");
            }
            if ($.projectId == null) {
                throw new MissingRequiredPropertyException("ServiceendpointCheckmarxSastArgs", "projectId");
            }
            if ($.serverUrl == null) {
                throw new MissingRequiredPropertyException("ServiceendpointCheckmarxSastArgs", "serverUrl");
            }
            if ($.serviceEndpointName == null) {
                throw new MissingRequiredPropertyException("ServiceendpointCheckmarxSastArgs", "serviceEndpointName");
            }
            if ($.username == null) {
                throw new MissingRequiredPropertyException("ServiceendpointCheckmarxSastArgs", "username");
            }
            return $;
        }
    }

}
