// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.Core.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ProjectState extends com.pulumi.resources.ResourceArgs {

    public static final ProjectState Empty = new ProjectState();

    /**
     * The Description of the Project.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return The Description of the Project.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * Defines the status (`enabled`, `disabled`) of the project features.
     * Valid features are `boards`, `repositories`, `pipelines`, `testplans`, `artifacts`**NOTE:**
     * It&#39;s possible to define project features both within the `azuredevops.ProjectFeatures` resource and
     * via the `features` block by using the `azuredevops.Project` resource.
     * However it&#39;s not possible to use both methods to manage features, since there&#39;ll be conflicts.
     * 
     */
    @Import(name="features")
    private @Nullable Output<Map<String,String>> features;

    /**
     * @return Defines the status (`enabled`, `disabled`) of the project features.
     * Valid features are `boards`, `repositories`, `pipelines`, `testplans`, `artifacts`**NOTE:**
     * It&#39;s possible to define project features both within the `azuredevops.ProjectFeatures` resource and
     * via the `features` block by using the `azuredevops.Project` resource.
     * However it&#39;s not possible to use both methods to manage features, since there&#39;ll be conflicts.
     * 
     */
    public Optional<Output<Map<String,String>>> features() {
        return Optional.ofNullable(this.features);
    }

    /**
     * The Project Name.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The Project Name.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The Process Template ID used by the Project.
     * 
     */
    @Import(name="processTemplateId")
    private @Nullable Output<String> processTemplateId;

    /**
     * @return The Process Template ID used by the Project.
     * 
     */
    public Optional<Output<String>> processTemplateId() {
        return Optional.ofNullable(this.processTemplateId);
    }

    /**
     * Specifies the version control system. Valid values: `Git` or `Tfvc`. Defaults to `Git`.
     * 
     */
    @Import(name="versionControl")
    private @Nullable Output<String> versionControl;

    /**
     * @return Specifies the version control system. Valid values: `Git` or `Tfvc`. Defaults to `Git`.
     * 
     */
    public Optional<Output<String>> versionControl() {
        return Optional.ofNullable(this.versionControl);
    }

    /**
     * Specifies the visibility of the Project. Valid values: `private` or `public`. Defaults to `private`.
     * 
     */
    @Import(name="visibility")
    private @Nullable Output<String> visibility;

    /**
     * @return Specifies the visibility of the Project. Valid values: `private` or `public`. Defaults to `private`.
     * 
     */
    public Optional<Output<String>> visibility() {
        return Optional.ofNullable(this.visibility);
    }

    /**
     * Specifies the work item template. Valid values: `Agile`, `Basic`, `CMMI`, `Scrum` or a custom, pre-existing one. Defaults to `Agile`. An empty string will use the parent organization default.
     * 
     */
    @Import(name="workItemTemplate")
    private @Nullable Output<String> workItemTemplate;

    /**
     * @return Specifies the work item template. Valid values: `Agile`, `Basic`, `CMMI`, `Scrum` or a custom, pre-existing one. Defaults to `Agile`. An empty string will use the parent organization default.
     * 
     */
    public Optional<Output<String>> workItemTemplate() {
        return Optional.ofNullable(this.workItemTemplate);
    }

    private ProjectState() {}

    private ProjectState(ProjectState $) {
        this.description = $.description;
        this.features = $.features;
        this.name = $.name;
        this.processTemplateId = $.processTemplateId;
        this.versionControl = $.versionControl;
        this.visibility = $.visibility;
        this.workItemTemplate = $.workItemTemplate;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ProjectState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ProjectState $;

        public Builder() {
            $ = new ProjectState();
        }

        public Builder(ProjectState defaults) {
            $ = new ProjectState(Objects.requireNonNull(defaults));
        }

        /**
         * @param description The Description of the Project.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The Description of the Project.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param features Defines the status (`enabled`, `disabled`) of the project features.
         * Valid features are `boards`, `repositories`, `pipelines`, `testplans`, `artifacts`**NOTE:**
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
         * @param features Defines the status (`enabled`, `disabled`) of the project features.
         * Valid features are `boards`, `repositories`, `pipelines`, `testplans`, `artifacts`**NOTE:**
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

        /**
         * @param name The Project Name.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The Project Name.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param processTemplateId The Process Template ID used by the Project.
         * 
         * @return builder
         * 
         */
        public Builder processTemplateId(@Nullable Output<String> processTemplateId) {
            $.processTemplateId = processTemplateId;
            return this;
        }

        /**
         * @param processTemplateId The Process Template ID used by the Project.
         * 
         * @return builder
         * 
         */
        public Builder processTemplateId(String processTemplateId) {
            return processTemplateId(Output.of(processTemplateId));
        }

        /**
         * @param versionControl Specifies the version control system. Valid values: `Git` or `Tfvc`. Defaults to `Git`.
         * 
         * @return builder
         * 
         */
        public Builder versionControl(@Nullable Output<String> versionControl) {
            $.versionControl = versionControl;
            return this;
        }

        /**
         * @param versionControl Specifies the version control system. Valid values: `Git` or `Tfvc`. Defaults to `Git`.
         * 
         * @return builder
         * 
         */
        public Builder versionControl(String versionControl) {
            return versionControl(Output.of(versionControl));
        }

        /**
         * @param visibility Specifies the visibility of the Project. Valid values: `private` or `public`. Defaults to `private`.
         * 
         * @return builder
         * 
         */
        public Builder visibility(@Nullable Output<String> visibility) {
            $.visibility = visibility;
            return this;
        }

        /**
         * @param visibility Specifies the visibility of the Project. Valid values: `private` or `public`. Defaults to `private`.
         * 
         * @return builder
         * 
         */
        public Builder visibility(String visibility) {
            return visibility(Output.of(visibility));
        }

        /**
         * @param workItemTemplate Specifies the work item template. Valid values: `Agile`, `Basic`, `CMMI`, `Scrum` or a custom, pre-existing one. Defaults to `Agile`. An empty string will use the parent organization default.
         * 
         * @return builder
         * 
         */
        public Builder workItemTemplate(@Nullable Output<String> workItemTemplate) {
            $.workItemTemplate = workItemTemplate;
            return this;
        }

        /**
         * @param workItemTemplate Specifies the work item template. Valid values: `Agile`, `Basic`, `CMMI`, `Scrum` or a custom, pre-existing one. Defaults to `Agile`. An empty string will use the parent organization default.
         * 
         * @return builder
         * 
         */
        public Builder workItemTemplate(String workItemTemplate) {
            return workItemTemplate(Output.of(workItemTemplate));
        }

        public ProjectState build() {
            return $;
        }
    }

}
