// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.outputs;

import com.pulumi.azuredevops.outputs.GetTeamsTeam;
import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetTeamsResult {
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return Project identifier.
     * - `id - Team identifier
     * 
     */
    private @Nullable String projectId;
    /**
     * @return A list of existing projects in your Azure DevOps Organization with details about every project which includes:
     * 
     */
    private List<GetTeamsTeam> teams;
    private @Nullable Integer top;

    private GetTeamsResult() {}
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Project identifier.
     * - `id - Team identifier
     * 
     */
    public Optional<String> projectId() {
        return Optional.ofNullable(this.projectId);
    }
    /**
     * @return A list of existing projects in your Azure DevOps Organization with details about every project which includes:
     * 
     */
    public List<GetTeamsTeam> teams() {
        return this.teams;
    }
    public Optional<Integer> top() {
        return Optional.ofNullable(this.top);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetTeamsResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private @Nullable String projectId;
        private List<GetTeamsTeam> teams;
        private @Nullable Integer top;
        public Builder() {}
        public Builder(GetTeamsResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.projectId = defaults.projectId;
    	      this.teams = defaults.teams;
    	      this.top = defaults.top;
        }

        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetTeamsResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder projectId(@Nullable String projectId) {

            this.projectId = projectId;
            return this;
        }
        @CustomType.Setter
        public Builder teams(List<GetTeamsTeam> teams) {
            if (teams == null) {
              throw new MissingRequiredPropertyException("GetTeamsResult", "teams");
            }
            this.teams = teams;
            return this;
        }
        public Builder teams(GetTeamsTeam... teams) {
            return teams(List.of(teams));
        }
        @CustomType.Setter
        public Builder top(@Nullable Integer top) {

            this.top = top;
            return this;
        }
        public GetTeamsResult build() {
            final var _resultValue = new GetTeamsResult();
            _resultValue.id = id;
            _resultValue.projectId = projectId;
            _resultValue.teams = teams;
            _resultValue.top = top;
            return _resultValue;
        }
    }
}
