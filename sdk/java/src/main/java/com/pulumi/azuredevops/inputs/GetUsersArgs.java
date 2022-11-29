// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetUsersArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetUsersArgs Empty = new GetUsersArgs();

    /**
     * The type of source provider for the `origin_id` parameter (ex:AD, AAD, MSA) The supported origins are listed below.
     * 
     */
    @Import(name="origin")
    private @Nullable Output<String> origin;

    /**
     * @return The type of source provider for the `origin_id` parameter (ex:AD, AAD, MSA) The supported origins are listed below.
     * 
     */
    public Optional<Output<String>> origin() {
        return Optional.ofNullable(this.origin);
    }

    /**
     * The unique identifier from the system of origin.
     * 
     */
    @Import(name="originId")
    private @Nullable Output<String> originId;

    /**
     * @return The unique identifier from the system of origin.
     * 
     */
    public Optional<Output<String>> originId() {
        return Optional.ofNullable(this.originId);
    }

    /**
     * The PrincipalName of this graph member from the source provider.
     * 
     */
    @Import(name="principalName")
    private @Nullable Output<String> principalName;

    /**
     * @return The PrincipalName of this graph member from the source provider.
     * 
     */
    public Optional<Output<String>> principalName() {
        return Optional.ofNullable(this.principalName);
    }

    /**
     * A list of user subject subtypes to reduce the retrieved results, e.g. `msa`, `aad`, `svc` (service identity), `imp` (imported identity), etc. The supported subject types are listed below.
     * 
     */
    @Import(name="subjectTypes")
    private @Nullable Output<List<String>> subjectTypes;

    /**
     * @return A list of user subject subtypes to reduce the retrieved results, e.g. `msa`, `aad`, `svc` (service identity), `imp` (imported identity), etc. The supported subject types are listed below.
     * 
     */
    public Optional<Output<List<String>>> subjectTypes() {
        return Optional.ofNullable(this.subjectTypes);
    }

    private GetUsersArgs() {}

    private GetUsersArgs(GetUsersArgs $) {
        this.origin = $.origin;
        this.originId = $.originId;
        this.principalName = $.principalName;
        this.subjectTypes = $.subjectTypes;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetUsersArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetUsersArgs $;

        public Builder() {
            $ = new GetUsersArgs();
        }

        public Builder(GetUsersArgs defaults) {
            $ = new GetUsersArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param origin The type of source provider for the `origin_id` parameter (ex:AD, AAD, MSA) The supported origins are listed below.
         * 
         * @return builder
         * 
         */
        public Builder origin(@Nullable Output<String> origin) {
            $.origin = origin;
            return this;
        }

        /**
         * @param origin The type of source provider for the `origin_id` parameter (ex:AD, AAD, MSA) The supported origins are listed below.
         * 
         * @return builder
         * 
         */
        public Builder origin(String origin) {
            return origin(Output.of(origin));
        }

        /**
         * @param originId The unique identifier from the system of origin.
         * 
         * @return builder
         * 
         */
        public Builder originId(@Nullable Output<String> originId) {
            $.originId = originId;
            return this;
        }

        /**
         * @param originId The unique identifier from the system of origin.
         * 
         * @return builder
         * 
         */
        public Builder originId(String originId) {
            return originId(Output.of(originId));
        }

        /**
         * @param principalName The PrincipalName of this graph member from the source provider.
         * 
         * @return builder
         * 
         */
        public Builder principalName(@Nullable Output<String> principalName) {
            $.principalName = principalName;
            return this;
        }

        /**
         * @param principalName The PrincipalName of this graph member from the source provider.
         * 
         * @return builder
         * 
         */
        public Builder principalName(String principalName) {
            return principalName(Output.of(principalName));
        }

        /**
         * @param subjectTypes A list of user subject subtypes to reduce the retrieved results, e.g. `msa`, `aad`, `svc` (service identity), `imp` (imported identity), etc. The supported subject types are listed below.
         * 
         * @return builder
         * 
         */
        public Builder subjectTypes(@Nullable Output<List<String>> subjectTypes) {
            $.subjectTypes = subjectTypes;
            return this;
        }

        /**
         * @param subjectTypes A list of user subject subtypes to reduce the retrieved results, e.g. `msa`, `aad`, `svc` (service identity), `imp` (imported identity), etc. The supported subject types are listed below.
         * 
         * @return builder
         * 
         */
        public Builder subjectTypes(List<String> subjectTypes) {
            return subjectTypes(Output.of(subjectTypes));
        }

        /**
         * @param subjectTypes A list of user subject subtypes to reduce the retrieved results, e.g. `msa`, `aad`, `svc` (service identity), `imp` (imported identity), etc. The supported subject types are listed below.
         * 
         * @return builder
         * 
         */
        public Builder subjectTypes(String... subjectTypes) {
            return subjectTypes(List.of(subjectTypes));
        }

        public GetUsersArgs build() {
            return $;
        }
    }

}
