// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Map;
import java.util.Objects;


public final class ProjectFeaturesArgs extends com.pulumi.resources.ResourceArgs {

    public static final ProjectFeaturesArgs Empty = new ProjectFeaturesArgs();

    /**
     * Defines the status (`enabled`, `disabled`) of the project features.\
     * Valid features `boards`, `repositories`, `pipelines`, `testplans`, `artifacts`
     * 
     * &gt; **NOTE:**
     * It&#39;s possible to define project features both within the `azuredevops.ProjectFeatures` resource and
     * via the `features` block by using the `azuredevops.Project` resource.
     * However it&#39;s not possible to use both methods to manage features, since there&#39;ll be conflicts.
     * 
     */
    @Import(name="features", required=true)
    private Output<Map<String,String>> features;

    /**
     * @return Defines the status (`enabled`, `disabled`) of the project features.\
     * Valid features `boards`, `repositories`, `pipelines`, `testplans`, `artifacts`
     * 
     * &gt; **NOTE:**
     * It&#39;s possible to define project features both within the `azuredevops.ProjectFeatures` resource and
     * via the `features` block by using the `azuredevops.Project` resource.
     * However it&#39;s not possible to use both methods to manage features, since there&#39;ll be conflicts.
     * 
     */
    public Output<Map<String,String>> features() {
        return this.features;
    }

    @Import(name="projectId", required=true)
    private Output<String> projectId;

    public Output<String> projectId() {
        return this.projectId;
    }

    private ProjectFeaturesArgs() {}

    private ProjectFeaturesArgs(ProjectFeaturesArgs $) {
        this.features = $.features;
        this.projectId = $.projectId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ProjectFeaturesArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ProjectFeaturesArgs $;

        public Builder() {
            $ = new ProjectFeaturesArgs();
        }

        public Builder(ProjectFeaturesArgs defaults) {
            $ = new ProjectFeaturesArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param features Defines the status (`enabled`, `disabled`) of the project features.\
         * Valid features `boards`, `repositories`, `pipelines`, `testplans`, `artifacts`
         * 
         * &gt; **NOTE:**
         * It&#39;s possible to define project features both within the `azuredevops.ProjectFeatures` resource and
         * via the `features` block by using the `azuredevops.Project` resource.
         * However it&#39;s not possible to use both methods to manage features, since there&#39;ll be conflicts.
         * 
         * @return builder
         * 
         */
        public Builder features(Output<Map<String,String>> features) {
            $.features = features;
            return this;
        }

        /**
         * @param features Defines the status (`enabled`, `disabled`) of the project features.\
         * Valid features `boards`, `repositories`, `pipelines`, `testplans`, `artifacts`
         * 
         * &gt; **NOTE:**
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

        public Builder projectId(Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
        }

        public ProjectFeaturesArgs build() {
            if ($.features == null) {
                throw new MissingRequiredPropertyException("ProjectFeaturesArgs", "features");
            }
            if ($.projectId == null) {
                throw new MissingRequiredPropertyException("ProjectFeaturesArgs", "projectId");
            }
            return $;
        }
    }

}
