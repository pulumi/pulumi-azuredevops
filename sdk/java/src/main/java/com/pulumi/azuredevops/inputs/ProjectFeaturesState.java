// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ProjectFeaturesState extends com.pulumi.resources.ResourceArgs {

    public static final ProjectFeaturesState Empty = new ProjectFeaturesState();

    /**
     * Defines the status (`enabled`, `disabled`) of the project features.\
     * Valid features `boards`, `repositories`, `pipelines`, `testplans`, `artifacts`**NOTE:**\
     * It&#39;s possible to define project features both within the `azuredevops.ProjectFeatures` resource and
     * via the `features` block by using the `azuredevops.Project` resource.
     * However it&#39;s not possible to use both methods to manage features, since there&#39;ll be conflicts.
     * 
     */
    @Import(name="features")
    private @Nullable Output<Map<String,String>> features;

    /**
     * @return Defines the status (`enabled`, `disabled`) of the project features.\
     * Valid features `boards`, `repositories`, `pipelines`, `testplans`, `artifacts`**NOTE:**\
     * It&#39;s possible to define project features both within the `azuredevops.ProjectFeatures` resource and
     * via the `features` block by using the `azuredevops.Project` resource.
     * However it&#39;s not possible to use both methods to manage features, since there&#39;ll be conflicts.
     * 
     */
    public Optional<Output<Map<String,String>>> features() {
        return Optional.ofNullable(this.features);
    }

    @Import(name="projectId")
    private @Nullable Output<String> projectId;

    public Optional<Output<String>> projectId() {
        return Optional.ofNullable(this.projectId);
    }

    private ProjectFeaturesState() {}

    private ProjectFeaturesState(ProjectFeaturesState $) {
        this.features = $.features;
        this.projectId = $.projectId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ProjectFeaturesState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ProjectFeaturesState $;

        public Builder() {
            $ = new ProjectFeaturesState();
        }

        public Builder(ProjectFeaturesState defaults) {
            $ = new ProjectFeaturesState(Objects.requireNonNull(defaults));
        }

        /**
         * @param features Defines the status (`enabled`, `disabled`) of the project features.\
         * Valid features `boards`, `repositories`, `pipelines`, `testplans`, `artifacts`**NOTE:**\
         * It&#39;s possible to define project features both within the `azuredevops.ProjectFeatures` resource and
         * via the `features` block by using the `azuredevops.Project` resource.
         * However it&#39;s not possible to use both methods to manage features, since there&#39;ll be conflicts.
         * 
         * @return builder
         * 
         */
        public Builder features(@Nullable Output<Map<String,String>> features) {
            $.features = features;
            return this;
        }

        /**
         * @param features Defines the status (`enabled`, `disabled`) of the project features.\
         * Valid features `boards`, `repositories`, `pipelines`, `testplans`, `artifacts`**NOTE:**\
         * It&#39;s possible to define project features both within the `azuredevops.ProjectFeatures` resource and
         * via the `features` block by using the `azuredevops.Project` resource.
         * However it&#39;s not possible to use both methods to manage features, since there&#39;ll be conflicts.
         * 
         * @return builder
         * 
         */
        public Builder features(Map<String,String> features) {
            return features(Output.of(features));
        }

        public Builder projectId(@Nullable Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
        }

        public ProjectFeaturesState build() {
            return $;
        }
    }

}
