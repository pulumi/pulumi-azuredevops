// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class GetGroupMembershipPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetGroupMembershipPlainArgs Empty = new GetGroupMembershipPlainArgs();

    /**
     * The descriptor of the group.
     * 
     */
    @Import(name="groupDescriptor", required=true)
    private String groupDescriptor;

    /**
     * @return The descriptor of the group.
     * 
     */
    public String groupDescriptor() {
        return this.groupDescriptor;
    }

    private GetGroupMembershipPlainArgs() {}

    private GetGroupMembershipPlainArgs(GetGroupMembershipPlainArgs $) {
        this.groupDescriptor = $.groupDescriptor;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetGroupMembershipPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetGroupMembershipPlainArgs $;

        public Builder() {
            $ = new GetGroupMembershipPlainArgs();
        }

        public Builder(GetGroupMembershipPlainArgs defaults) {
            $ = new GetGroupMembershipPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param groupDescriptor The descriptor of the group.
         * 
         * @return builder
         * 
         */
        public Builder groupDescriptor(String groupDescriptor) {
            $.groupDescriptor = groupDescriptor;
            return this;
        }

        public GetGroupMembershipPlainArgs build() {
            if ($.groupDescriptor == null) {
                throw new MissingRequiredPropertyException("GetGroupMembershipPlainArgs", "groupDescriptor");
            }
            return $;
        }
    }

}
