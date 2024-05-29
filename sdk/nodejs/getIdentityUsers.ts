// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Use this data source to access information about an existing users within Azure DevOps On-Premise(Azure DevOps Server).
 */
export function getIdentityUsers(args: GetIdentityUsersArgs, opts?: pulumi.InvokeOptions): Promise<GetIdentityUsersResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azuredevops:index/getIdentityUsers:getIdentityUsers", {
        "name": args.name,
        "searchFilter": args.searchFilter,
    }, opts);
}

/**
 * A collection of arguments for invoking getIdentityUsers.
 */
export interface GetIdentityUsersArgs {
    /**
     * The PrincipalName of this identity member from the source provider.
     */
    name: string;
    /**
     * The type of search to perform. Default is `General`. Possible values are `AccountName`, `DisplayName`, and `MailAddress`.
     */
    searchFilter?: string;
}

/**
 * A collection of values returned by getIdentityUsers.
 */
export interface GetIdentityUsersResult {
    readonly descriptor: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * This is the PrincipalName of this identity member from the source provider. The source provider may change this field over time and it is not guaranteed to be immutable for the life of the identity member.
     */
    readonly name: string;
    readonly searchFilter?: string;
}
/**
 * Use this data source to access information about an existing users within Azure DevOps On-Premise(Azure DevOps Server).
 */
export function getIdentityUsersOutput(args: GetIdentityUsersOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetIdentityUsersResult> {
    return pulumi.output(args).apply((a: any) => getIdentityUsers(a, opts))
}

/**
 * A collection of arguments for invoking getIdentityUsers.
 */
export interface GetIdentityUsersOutputArgs {
    /**
     * The PrincipalName of this identity member from the source provider.
     */
    name: pulumi.Input<string>;
    /**
     * The type of search to perform. Default is `General`. Possible values are `AccountName`, `DisplayName`, and `MailAddress`.
     */
    searchFilter?: pulumi.Input<string>;
}