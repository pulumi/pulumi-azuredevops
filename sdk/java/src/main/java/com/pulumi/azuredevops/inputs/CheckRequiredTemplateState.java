// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.inputs;

import com.pulumi.azuredevops.inputs.CheckRequiredTemplateRequiredTemplateArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class CheckRequiredTemplateState extends com.pulumi.resources.ResourceArgs {

    public static final CheckRequiredTemplateState Empty = new CheckRequiredTemplateState();

    /**
     * The project ID. Changing this forces a new Required Template Check to be created.
     * 
     */
    @Import(name="projectId")
    private @Nullable Output<String> projectId;

    /**
     * @return The project ID. Changing this forces a new Required Template Check to be created.
     * 
     */
    public Optional<Output<String>> projectId() {
        return Optional.ofNullable(this.projectId);
    }

    /**
     * One or more `required_template` blocks documented below.
     * 
     */
    @Import(name="requiredTemplates")
    private @Nullable Output<List<CheckRequiredTemplateRequiredTemplateArgs>> requiredTemplates;

    /**
     * @return One or more `required_template` blocks documented below.
     * 
     */
    public Optional<Output<List<CheckRequiredTemplateRequiredTemplateArgs>>> requiredTemplates() {
        return Optional.ofNullable(this.requiredTemplates);
    }

    /**
     * The ID of the resource being protected by the check. Changing this forces a new Required Template Check to be created.
     * 
     */
    @Import(name="targetResourceId")
    private @Nullable Output<String> targetResourceId;

    /**
     * @return The ID of the resource being protected by the check. Changing this forces a new Required Template Check to be created.
     * 
     */
    public Optional<Output<String>> targetResourceId() {
        return Optional.ofNullable(this.targetResourceId);
    }

    /**
     * The type of resource being protected by the check. Valid values: `endpoint`, `environment`, `queue`, `repository`, `securefile`, `variablegroup`. Changing this forces a new Required Template Check to be created.
     * 
     */
    @Import(name="targetResourceType")
    private @Nullable Output<String> targetResourceType;

    /**
     * @return The type of resource being protected by the check. Valid values: `endpoint`, `environment`, `queue`, `repository`, `securefile`, `variablegroup`. Changing this forces a new Required Template Check to be created.
     * 
     */
    public Optional<Output<String>> targetResourceType() {
        return Optional.ofNullable(this.targetResourceType);
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

    private CheckRequiredTemplateState() {}

    private CheckRequiredTemplateState(CheckRequiredTemplateState $) {
        this.projectId = $.projectId;
        this.requiredTemplates = $.requiredTemplates;
        this.targetResourceId = $.targetResourceId;
        this.targetResourceType = $.targetResourceType;
        this.version = $.version;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(CheckRequiredTemplateState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private CheckRequiredTemplateState $;

        public Builder() {
            $ = new CheckRequiredTemplateState();
        }

        public Builder(CheckRequiredTemplateState defaults) {
            $ = new CheckRequiredTemplateState(Objects.requireNonNull(defaults));
        }

        /**
         * @param projectId The project ID. Changing this forces a new Required Template Check to be created.
         * 
         * @return builder
         * 
         */
        public Builder projectId(@Nullable Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param projectId The project ID. Changing this forces a new Required Template Check to be created.
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
        }

        /**
         * @param requiredTemplates One or more `required_template` blocks documented below.
         * 
         * @return builder
         * 
         */
        public Builder requiredTemplates(@Nullable Output<List<CheckRequiredTemplateRequiredTemplateArgs>> requiredTemplates) {
            $.requiredTemplates = requiredTemplates;
            return this;
        }

        /**
         * @param requiredTemplates One or more `required_template` blocks documented below.
         * 
         * @return builder
         * 
         */
        public Builder requiredTemplates(List<CheckRequiredTemplateRequiredTemplateArgs> requiredTemplates) {
            return requiredTemplates(Output.of(requiredTemplates));
        }

        /**
         * @param requiredTemplates One or more `required_template` blocks documented below.
         * 
         * @return builder
         * 
         */
        public Builder requiredTemplates(CheckRequiredTemplateRequiredTemplateArgs... requiredTemplates) {
            return requiredTemplates(List.of(requiredTemplates));
        }

        /**
         * @param targetResourceId The ID of the resource being protected by the check. Changing this forces a new Required Template Check to be created.
         * 
         * @return builder
         * 
         */
        public Builder targetResourceId(@Nullable Output<String> targetResourceId) {
            $.targetResourceId = targetResourceId;
            return this;
        }

        /**
         * @param targetResourceId The ID of the resource being protected by the check. Changing this forces a new Required Template Check to be created.
         * 
         * @return builder
         * 
         */
        public Builder targetResourceId(String targetResourceId) {
            return targetResourceId(Output.of(targetResourceId));
        }

        /**
         * @param targetResourceType The type of resource being protected by the check. Valid values: `endpoint`, `environment`, `queue`, `repository`, `securefile`, `variablegroup`. Changing this forces a new Required Template Check to be created.
         * 
         * @return builder
         * 
         */
        public Builder targetResourceType(@Nullable Output<String> targetResourceType) {
            $.targetResourceType = targetResourceType;
            return this;
        }

        /**
         * @param targetResourceType The type of resource being protected by the check. Valid values: `endpoint`, `environment`, `queue`, `repository`, `securefile`, `variablegroup`. Changing this forces a new Required Template Check to be created.
         * 
         * @return builder
         * 
         */
        public Builder targetResourceType(String targetResourceType) {
            return targetResourceType(Output.of(targetResourceType));
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

        public CheckRequiredTemplateState build() {
            return $;
        }
    }

}
