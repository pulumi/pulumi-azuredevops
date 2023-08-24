// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetTeamResult {
    /**
     * @return List of subject descriptors for `administrators` of the team.
     * 
     */
    private List<String> administrators;
    /**
     * @return Team description.
     * 
     */
    private String description;
    /**
     * @return The descriptor of the Team.
     * 
     */
    private String descriptor;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return List of subject descriptors for `members` of the team.
     * 
     */
    private List<String> members;
    private String name;
    private String projectId;
    private @Nullable Integer top;

    private GetTeamResult() {}
    /**
     * @return List of subject descriptors for `administrators` of the team.
     * 
     */
    public List<String> administrators() {
        return this.administrators;
    }
    /**
     * @return Team description.
     * 
     */
    public String description() {
        return this.description;
    }
    /**
     * @return The descriptor of the Team.
     * 
     */
    public String descriptor() {
        return this.descriptor;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return List of subject descriptors for `members` of the team.
     * 
     */
    public List<String> members() {
        return this.members;
    }
    public String name() {
        return this.name;
    }
    public String projectId() {
        return this.projectId;
    }
    public Optional<Integer> top() {
        return Optional.ofNullable(this.top);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetTeamResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<String> administrators;
        private String description;
        private String descriptor;
        private String id;
        private List<String> members;
        private String name;
        private String projectId;
        private @Nullable Integer top;
        public Builder() {}
        public Builder(GetTeamResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.administrators = defaults.administrators;
    	      this.description = defaults.description;
    	      this.descriptor = defaults.descriptor;
    	      this.id = defaults.id;
    	      this.members = defaults.members;
    	      this.name = defaults.name;
    	      this.projectId = defaults.projectId;
    	      this.top = defaults.top;
        }

        @CustomType.Setter
        public Builder administrators(List<String> administrators) {
            this.administrators = Objects.requireNonNull(administrators);
            return this;
        }
        public Builder administrators(String... administrators) {
            return administrators(List.of(administrators));
        }
        @CustomType.Setter
        public Builder description(String description) {
            this.description = Objects.requireNonNull(description);
            return this;
        }
        @CustomType.Setter
        public Builder descriptor(String descriptor) {
            this.descriptor = Objects.requireNonNull(descriptor);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder members(List<String> members) {
            this.members = Objects.requireNonNull(members);
            return this;
        }
        public Builder members(String... members) {
            return members(List.of(members));
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder projectId(String projectId) {
            this.projectId = Objects.requireNonNull(projectId);
            return this;
        }
        @CustomType.Setter
        public Builder top(@Nullable Integer top) {
            this.top = top;
            return this;
        }
        public GetTeamResult build() {
            final var o = new GetTeamResult();
            o.administrators = administrators;
            o.description = description;
            o.descriptor = descriptor;
            o.id = id;
            o.members = members;
            o.name = name;
            o.projectId = projectId;
            o.top = top;
            return o;
        }
    }
}
