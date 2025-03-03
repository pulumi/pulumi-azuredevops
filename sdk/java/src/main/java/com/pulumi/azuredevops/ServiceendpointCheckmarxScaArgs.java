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


public final class ServiceendpointCheckmarxScaArgs extends com.pulumi.resources.ResourceArgs {

    public static final ServiceendpointCheckmarxScaArgs Empty = new ServiceendpointCheckmarxScaArgs();

    /**
     * The Access Control URL of the Checkmarx SCA.
     * 
     */
    @Import(name="accessControlUrl", required=true)
    private Output<String> accessControlUrl;

    /**
     * @return The Access Control URL of the Checkmarx SCA.
     * 
     */
    public Output<String> accessControlUrl() {
        return this.accessControlUrl;
    }

    /**
     * The account of the Checkmarx SCA.
     * 
     */
    @Import(name="account", required=true)
    private Output<String> account;

    /**
     * @return The account of the Checkmarx SCA.
     * 
     */
    public Output<String> account() {
        return this.account;
    }

    @Import(name="description")
    private @Nullable Output<String> description;

    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The password of the Checkmarx SCA.
     * 
     */
    @Import(name="password", required=true)
    private Output<String> password;

    /**
     * @return The password of the Checkmarx SCA.
     * 
     */
    public Output<String> password() {
        return this.password;
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
     * The Server URL of the Checkmarx SCA.
     * 
     */
    @Import(name="serverUrl", required=true)
    private Output<String> serverUrl;

    /**
     * @return The Server URL of the Checkmarx SCA.
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
     * The username of the Checkmarx SCA.
     * 
     */
    @Import(name="username", required=true)
    private Output<String> username;

    /**
     * @return The username of the Checkmarx SCA.
     * 
     */
    public Output<String> username() {
        return this.username;
    }

    /**
     * The Web App URL of the Checkmarx SCA.
     * 
     */
    @Import(name="webAppUrl", required=true)
    private Output<String> webAppUrl;

    /**
     * @return The Web App URL of the Checkmarx SCA.
     * 
     */
    public Output<String> webAppUrl() {
        return this.webAppUrl;
    }

    private ServiceendpointCheckmarxScaArgs() {}

    private ServiceendpointCheckmarxScaArgs(ServiceendpointCheckmarxScaArgs $) {
        this.accessControlUrl = $.accessControlUrl;
        this.account = $.account;
        this.description = $.description;
        this.password = $.password;
        this.projectId = $.projectId;
        this.serverUrl = $.serverUrl;
        this.serviceEndpointName = $.serviceEndpointName;
        this.team = $.team;
        this.username = $.username;
        this.webAppUrl = $.webAppUrl;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServiceendpointCheckmarxScaArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServiceendpointCheckmarxScaArgs $;

        public Builder() {
            $ = new ServiceendpointCheckmarxScaArgs();
        }

        public Builder(ServiceendpointCheckmarxScaArgs defaults) {
            $ = new ServiceendpointCheckmarxScaArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param accessControlUrl The Access Control URL of the Checkmarx SCA.
         * 
         * @return builder
         * 
         */
        public Builder accessControlUrl(Output<String> accessControlUrl) {
            $.accessControlUrl = accessControlUrl;
            return this;
        }

        /**
         * @param accessControlUrl The Access Control URL of the Checkmarx SCA.
         * 
         * @return builder
         * 
         */
        public Builder accessControlUrl(String accessControlUrl) {
            return accessControlUrl(Output.of(accessControlUrl));
        }

        /**
         * @param account The account of the Checkmarx SCA.
         * 
         * @return builder
         * 
         */
        public Builder account(Output<String> account) {
            $.account = account;
            return this;
        }

        /**
         * @param account The account of the Checkmarx SCA.
         * 
         * @return builder
         * 
         */
        public Builder account(String account) {
            return account(Output.of(account));
        }

        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param password The password of the Checkmarx SCA.
         * 
         * @return builder
         * 
         */
        public Builder password(Output<String> password) {
            $.password = password;
            return this;
        }

        /**
         * @param password The password of the Checkmarx SCA.
         * 
         * @return builder
         * 
         */
        public Builder password(String password) {
            return password(Output.of(password));
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
         * @param serverUrl The Server URL of the Checkmarx SCA.
         * 
         * @return builder
         * 
         */
        public Builder serverUrl(Output<String> serverUrl) {
            $.serverUrl = serverUrl;
            return this;
        }

        /**
         * @param serverUrl The Server URL of the Checkmarx SCA.
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
         * @param username The username of the Checkmarx SCA.
         * 
         * @return builder
         * 
         */
        public Builder username(Output<String> username) {
            $.username = username;
            return this;
        }

        /**
         * @param username The username of the Checkmarx SCA.
         * 
         * @return builder
         * 
         */
        public Builder username(String username) {
            return username(Output.of(username));
        }

        /**
         * @param webAppUrl The Web App URL of the Checkmarx SCA.
         * 
         * @return builder
         * 
         */
        public Builder webAppUrl(Output<String> webAppUrl) {
            $.webAppUrl = webAppUrl;
            return this;
        }

        /**
         * @param webAppUrl The Web App URL of the Checkmarx SCA.
         * 
         * @return builder
         * 
         */
        public Builder webAppUrl(String webAppUrl) {
            return webAppUrl(Output.of(webAppUrl));
        }

        public ServiceendpointCheckmarxScaArgs build() {
            if ($.accessControlUrl == null) {
                throw new MissingRequiredPropertyException("ServiceendpointCheckmarxScaArgs", "accessControlUrl");
            }
            if ($.account == null) {
                throw new MissingRequiredPropertyException("ServiceendpointCheckmarxScaArgs", "account");
            }
            if ($.password == null) {
                throw new MissingRequiredPropertyException("ServiceendpointCheckmarxScaArgs", "password");
            }
            if ($.projectId == null) {
                throw new MissingRequiredPropertyException("ServiceendpointCheckmarxScaArgs", "projectId");
            }
            if ($.serverUrl == null) {
                throw new MissingRequiredPropertyException("ServiceendpointCheckmarxScaArgs", "serverUrl");
            }
            if ($.serviceEndpointName == null) {
                throw new MissingRequiredPropertyException("ServiceendpointCheckmarxScaArgs", "serviceEndpointName");
            }
            if ($.username == null) {
                throw new MissingRequiredPropertyException("ServiceendpointCheckmarxScaArgs", "username");
            }
            if ($.webAppUrl == null) {
                throw new MissingRequiredPropertyException("ServiceendpointCheckmarxScaArgs", "webAppUrl");
            }
            return $;
        }
    }

}
