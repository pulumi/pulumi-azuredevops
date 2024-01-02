// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.ServiceEndpoint.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GitHubAuthPersonal {
    /**
     * @return The Personal Access Token for GitHub.
     * 
     */
    private String personalAccessToken;

    private GitHubAuthPersonal() {}
    /**
     * @return The Personal Access Token for GitHub.
     * 
     */
    public String personalAccessToken() {
        return this.personalAccessToken;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GitHubAuthPersonal defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String personalAccessToken;
        public Builder() {}
        public Builder(GitHubAuthPersonal defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.personalAccessToken = defaults.personalAccessToken;
        }

        @CustomType.Setter
        public Builder personalAccessToken(String personalAccessToken) {
            if (personalAccessToken == null) {
              throw new MissingRequiredPropertyException("GitHubAuthPersonal", "personalAccessToken");
            }
            this.personalAccessToken = personalAccessToken;
            return this;
        }
        public GitHubAuthPersonal build() {
            final var _resultValue = new GitHubAuthPersonal();
            _resultValue.personalAccessToken = personalAccessToken;
            return _resultValue;
        }
    }
}
