// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.ServiceEndpoint.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GitHubAuthPersonal {
    /**
     * @return The Personal Access Token for GitHub.
     * 
     */
    private String personalAccessToken;
    private @Nullable String personalAccessTokenHash;

    private GitHubAuthPersonal() {}
    /**
     * @return The Personal Access Token for GitHub.
     * 
     */
    public String personalAccessToken() {
        return this.personalAccessToken;
    }
    public Optional<String> personalAccessTokenHash() {
        return Optional.ofNullable(this.personalAccessTokenHash);
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
        private @Nullable String personalAccessTokenHash;
        public Builder() {}
        public Builder(GitHubAuthPersonal defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.personalAccessToken = defaults.personalAccessToken;
    	      this.personalAccessTokenHash = defaults.personalAccessTokenHash;
        }

        @CustomType.Setter
        public Builder personalAccessToken(String personalAccessToken) {
            this.personalAccessToken = Objects.requireNonNull(personalAccessToken);
            return this;
        }
        @CustomType.Setter
        public Builder personalAccessTokenHash(@Nullable String personalAccessTokenHash) {
            this.personalAccessTokenHash = personalAccessTokenHash;
            return this;
        }
        public GitHubAuthPersonal build() {
            final var o = new GitHubAuthPersonal();
            o.personalAccessToken = personalAccessToken;
            o.personalAccessTokenHash = personalAccessTokenHash;
            return o;
        }
    }
}