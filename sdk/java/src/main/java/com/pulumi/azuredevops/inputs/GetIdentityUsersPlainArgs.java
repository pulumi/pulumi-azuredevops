// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetIdentityUsersPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetIdentityUsersPlainArgs Empty = new GetIdentityUsersPlainArgs();

    /**
     * The PrincipalName of this identity member from the source provider.
     * 
     */
    @Import(name="name", required=true)
    private String name;

    /**
     * @return The PrincipalName of this identity member from the source provider.
     * 
     */
    public String name() {
        return this.name;
    }

    /**
     * The type of search to perform. Possible values are: `AccountName`, `DisplayName`, and `MailAddress`. Default is `General`.
     * 
     */
    @Import(name="searchFilter")
    private @Nullable String searchFilter;

    /**
     * @return The type of search to perform. Possible values are: `AccountName`, `DisplayName`, and `MailAddress`. Default is `General`.
     * 
     */
    public Optional<String> searchFilter() {
        return Optional.ofNullable(this.searchFilter);
    }

    private GetIdentityUsersPlainArgs() {}

    private GetIdentityUsersPlainArgs(GetIdentityUsersPlainArgs $) {
        this.name = $.name;
        this.searchFilter = $.searchFilter;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetIdentityUsersPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetIdentityUsersPlainArgs $;

        public Builder() {
            $ = new GetIdentityUsersPlainArgs();
        }

        public Builder(GetIdentityUsersPlainArgs defaults) {
            $ = new GetIdentityUsersPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param name The PrincipalName of this identity member from the source provider.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            $.name = name;
            return this;
        }

        /**
         * @param searchFilter The type of search to perform. Possible values are: `AccountName`, `DisplayName`, and `MailAddress`. Default is `General`.
         * 
         * @return builder
         * 
         */
        public Builder searchFilter(@Nullable String searchFilter) {
            $.searchFilter = searchFilter;
            return this;
        }

        public GetIdentityUsersPlainArgs build() {
            if ($.name == null) {
                throw new MissingRequiredPropertyException("GetIdentityUsersPlainArgs", "name");
            }
            return $;
        }
    }

}
