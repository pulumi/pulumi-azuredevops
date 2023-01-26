// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.Pipeline.inputs;

import com.pulumi.azuredevops.Pipeline.inputs.VariableGroupKeyVaultArgs;
import com.pulumi.azuredevops.Pipeline.inputs.VariableGroupVariableArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class VariableGroupState extends com.pulumi.resources.ResourceArgs {

    public static final VariableGroupState Empty = new VariableGroupState();

    /**
     * Boolean that indicate if this variable group is shared by all pipelines of this project.
     * 
     */
    @Import(name="allowAccess")
    private @Nullable Output<Boolean> allowAccess;

    /**
     * @return Boolean that indicate if this variable group is shared by all pipelines of this project.
     * 
     */
    public Optional<Output<Boolean>> allowAccess() {
        return Optional.ofNullable(this.allowAccess);
    }

    /**
     * The description of the Variable Group.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return The description of the Variable Group.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    @Import(name="keyVault")
    private @Nullable Output<VariableGroupKeyVaultArgs> keyVault;

    public Optional<Output<VariableGroupKeyVaultArgs>> keyVault() {
        return Optional.ofNullable(this.keyVault);
    }

    /**
     * The name of the Variable Group.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the Variable Group.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The project ID or project name.
     * 
     */
    @Import(name="projectId")
    private @Nullable Output<String> projectId;

    /**
     * @return The project ID or project name.
     * 
     */
    public Optional<Output<String>> projectId() {
        return Optional.ofNullable(this.projectId);
    }

    /**
     * One or more `variable` blocks as documented below.
     * 
     */
    @Import(name="variables")
    private @Nullable Output<List<VariableGroupVariableArgs>> variables;

    /**
     * @return One or more `variable` blocks as documented below.
     * 
     */
    public Optional<Output<List<VariableGroupVariableArgs>>> variables() {
        return Optional.ofNullable(this.variables);
    }

    private VariableGroupState() {}

    private VariableGroupState(VariableGroupState $) {
        this.allowAccess = $.allowAccess;
        this.description = $.description;
        this.keyVault = $.keyVault;
        this.name = $.name;
        this.projectId = $.projectId;
        this.variables = $.variables;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(VariableGroupState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private VariableGroupState $;

        public Builder() {
            $ = new VariableGroupState();
        }

        public Builder(VariableGroupState defaults) {
            $ = new VariableGroupState(Objects.requireNonNull(defaults));
        }

        /**
         * @param allowAccess Boolean that indicate if this variable group is shared by all pipelines of this project.
         * 
         * @return builder
         * 
         */
        public Builder allowAccess(@Nullable Output<Boolean> allowAccess) {
            $.allowAccess = allowAccess;
            return this;
        }

        /**
         * @param allowAccess Boolean that indicate if this variable group is shared by all pipelines of this project.
         * 
         * @return builder
         * 
         */
        public Builder allowAccess(Boolean allowAccess) {
            return allowAccess(Output.of(allowAccess));
        }

        /**
         * @param description The description of the Variable Group.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The description of the Variable Group.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        public Builder keyVault(@Nullable Output<VariableGroupKeyVaultArgs> keyVault) {
            $.keyVault = keyVault;
            return this;
        }

        public Builder keyVault(VariableGroupKeyVaultArgs keyVault) {
            return keyVault(Output.of(keyVault));
        }

        /**
         * @param name The name of the Variable Group.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the Variable Group.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param projectId The project ID or project name.
         * 
         * @return builder
         * 
         */
        public Builder projectId(@Nullable Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param projectId The project ID or project name.
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
        }

        /**
         * @param variables One or more `variable` blocks as documented below.
         * 
         * @return builder
         * 
         */
        public Builder variables(@Nullable Output<List<VariableGroupVariableArgs>> variables) {
            $.variables = variables;
            return this;
        }

        /**
         * @param variables One or more `variable` blocks as documented below.
         * 
         * @return builder
         * 
         */
        public Builder variables(List<VariableGroupVariableArgs> variables) {
            return variables(Output.of(variables));
        }

        /**
         * @param variables One or more `variable` blocks as documented below.
         * 
         * @return builder
         * 
         */
        public Builder variables(VariableGroupVariableArgs... variables) {
            return variables(List.of(variables));
        }

        public VariableGroupState build() {
            return $;
        }
    }

}
