// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.util.Objects;


public final class BuildDefinitionPullRequestTriggerForksArgs extends com.pulumi.resources.ResourceArgs {

    public static final BuildDefinitionPullRequestTriggerForksArgs Empty = new BuildDefinitionPullRequestTriggerForksArgs();

    /**
     * Build pull requests from forks of this repository.
     * 
     */
    @Import(name="enabled", required=true)
    private Output<Boolean> enabled;

    /**
     * @return Build pull requests from forks of this repository.
     * 
     */
    public Output<Boolean> enabled() {
        return this.enabled;
    }

    /**
     * Make secrets available to builds of forks.
     * 
     */
    @Import(name="shareSecrets", required=true)
    private Output<Boolean> shareSecrets;

    /**
     * @return Make secrets available to builds of forks.
     * 
     */
    public Output<Boolean> shareSecrets() {
        return this.shareSecrets;
    }

    private BuildDefinitionPullRequestTriggerForksArgs() {}

    private BuildDefinitionPullRequestTriggerForksArgs(BuildDefinitionPullRequestTriggerForksArgs $) {
        this.enabled = $.enabled;
        this.shareSecrets = $.shareSecrets;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(BuildDefinitionPullRequestTriggerForksArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private BuildDefinitionPullRequestTriggerForksArgs $;

        public Builder() {
            $ = new BuildDefinitionPullRequestTriggerForksArgs();
        }

        public Builder(BuildDefinitionPullRequestTriggerForksArgs defaults) {
            $ = new BuildDefinitionPullRequestTriggerForksArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param enabled Build pull requests from forks of this repository.
         * 
         * @return builder
         * 
         */
        public Builder enabled(Output<Boolean> enabled) {
            $.enabled = enabled;
            return this;
        }

        /**
         * @param enabled Build pull requests from forks of this repository.
         * 
         * @return builder
         * 
         */
        public Builder enabled(Boolean enabled) {
            return enabled(Output.of(enabled));
        }

        /**
         * @param shareSecrets Make secrets available to builds of forks.
         * 
         * @return builder
         * 
         */
        public Builder shareSecrets(Output<Boolean> shareSecrets) {
            $.shareSecrets = shareSecrets;
            return this;
        }

        /**
         * @param shareSecrets Make secrets available to builds of forks.
         * 
         * @return builder
         * 
         */
        public Builder shareSecrets(Boolean shareSecrets) {
            return shareSecrets(Output.of(shareSecrets));
        }

        public BuildDefinitionPullRequestTriggerForksArgs build() {
            if ($.enabled == null) {
                throw new MissingRequiredPropertyException("BuildDefinitionPullRequestTriggerForksArgs", "enabled");
            }
            if ($.shareSecrets == null) {
                throw new MissingRequiredPropertyException("BuildDefinitionPullRequestTriggerForksArgs", "shareSecrets");
            }
            return $;
        }
    }

}
