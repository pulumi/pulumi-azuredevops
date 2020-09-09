// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Use this data source to access information about an existing users within Azure DevOps.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * // Load single user by using it's principal name
 * const user = pulumi.output(azuredevops.getUsers({
 *     principalName: "contoso-user@contoso.onmicrosoft.com",
 * }, { async: true }));
 * // Load all users know inside an organization
 * const all_users = pulumi.output(azuredevops.getUsers({ async: true }));
 * // Load all users know inside an organization originating from a specific source (origin)
 * const all_from_origin = pulumi.output(azuredevops.getUsers({
 *     origin: "aad",
 * }, { async: true }));
 * // Load all users know inside an organization filtered by their subject types
 * const all_from_subject_types = pulumi.output(azuredevops.getUsers({
 *     subjectTypes: [
 *         "aad",
 *         "msa",
 *     ],
 * }, { async: true }));
 * // Load a single user by origin and origin ID
 * const all_from_origin_id = pulumi.output(azuredevops.getUsers({
 *     origin: "aad",
 *     originId: "a7ead982-8438-4cd2-b9e3-c3aa51a7b675",
 * }, { async: true }));
 * ```
 * ## Relevant Links
 *
 * - [Azure DevOps Service REST API 5.1 - Graph Users API](https://docs.microsoft.com/en-us/rest/api/azure/devops/graph/users?view=azure-devops-rest-5.1)
 */
export function getUsers(args?: GetUsersArgs, opts?: pulumi.InvokeOptions): Promise<GetUsersResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azuredevops:index/getUsers:getUsers", {
        "origin": args.origin,
        "originId": args.originId,
        "principalName": args.principalName,
        "subjectTypes": args.subjectTypes,
    }, opts);
}

/**
 * A collection of arguments for invoking getUsers.
 */
export interface GetUsersArgs {
    /**
     * The type of source provider for the `originId` parameter (ex:AD, AAD, MSA) The supported origins are listed below.
     */
    readonly origin?: string;
    /**
     * The unique identifier from the system of origin.
     */
    readonly originId?: string;
    /**
     * The PrincipalName of this graph member from the source provider.
     */
    readonly principalName?: string;
    /**
     * A list of user subject subtypes to reduce the retrieved results, e.g. `msa`, `aad`, `svc` (service identity), `imp` (imported identity), etc. The supported subject types are listed below.
     */
    readonly subjectTypes?: string[];
}

/**
 * A collection of values returned by getUsers.
 */
export interface GetUsersResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The type of source provider for the origin identifier (ex:AD, AAD, MSA)
     */
    readonly origin?: string;
    /**
     * The unique identifier from the system of origin. Typically a sid, object id or Guid. Linking and unlinking operations can cause this value to change for a user because the user is not backed by a different provider and has a different unique id in the new provider.
     */
    readonly originId?: string;
    /**
     * This is the PrincipalName of this graph member from the source provider. The source provider may change this field over time and it is not guaranteed to be immutable for the life of the graph member by VSTS.
     */
    readonly principalName?: string;
    readonly subjectTypes?: string[];
    /**
     * A list of existing users in your Azure DevOps Organization with details about every single user which includes:
     */
    readonly users: outputs.GetUsersUser[];
}