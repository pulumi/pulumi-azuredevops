// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class ServiceendpointExternaltfsAuthPersonalArgs extends com.pulumi.resources.ResourceArgs {

    public static final ServiceendpointExternaltfsAuthPersonalArgs Empty = new ServiceendpointExternaltfsAuthPersonalArgs();

    /**
     * The Personal Access Token for Azure DevOps Organization.
     * 
     */
    @Import(name="personalAccessToken", required=true)
    private Output<String> personalAccessToken;

    /**
     * @return The Personal Access Token for Azure DevOps Organization.
     * 
     */
    public Output<String> personalAccessToken() {
        return this.personalAccessToken;
    }

    private ServiceendpointExternaltfsAuthPersonalArgs() {}

    private ServiceendpointExternaltfsAuthPersonalArgs(ServiceendpointExternaltfsAuthPersonalArgs $) {
        this.personalAccessToken = $.personalAccessToken;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServiceendpointExternaltfsAuthPersonalArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServiceendpointExternaltfsAuthPersonalArgs $;

        public Builder() {
            $ = new ServiceendpointExternaltfsAuthPersonalArgs();
        }

        public Builder(ServiceendpointExternaltfsAuthPersonalArgs defaults) {
            $ = new ServiceendpointExternaltfsAuthPersonalArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param personalAccessToken The Personal Access Token for Azure DevOps Organization.
         * 
         * @return builder
         * 
         */
        public Builder personalAccessToken(Output<String> personalAccessToken) {
            $.personalAccessToken = personalAccessToken;
            return this;
        }

        /**
         * @param personalAccessToken The Personal Access Token for Azure DevOps Organization.
         * 
         * @return builder
         * 
         */
        public Builder personalAccessToken(String personalAccessToken) {
            return personalAccessToken(Output.of(personalAccessToken));
        }

        public ServiceendpointExternaltfsAuthPersonalArgs build() {
            if ($.personalAccessToken == null) {
                throw new MissingRequiredPropertyException("ServiceendpointExternaltfsAuthPersonalArgs", "personalAccessToken");
            }
            return $;
        }
    }

}
