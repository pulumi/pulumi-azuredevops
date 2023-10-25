// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing users within Azure DevOps.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const example = azuredevops.getUsers({
 *     principalName: "contoso-user@contoso.onmicrosoft.com",
 * });
 * const example-all-users = azuredevops.getUsers({});
 * const example-all-from-origin = azuredevops.getUsers({
 *     origin: "aad",
 * });
 * const example-all-from-subjectTypes = azuredevops.getUsers({
 *     subjectTypes: [
 *         "aad",
 *         "msa",
 *     ],
 * });
 * const example-all-from-origin-id = azuredevops.getUsers({
 *     origin: "aad",
 *     originId: "00000000-0000-0000-0000-000000000000",
 * });
 * ```
 * ## Relevant Links
 *
 * - [Azure DevOps Service REST API 7.0 - Graph Users API](https://docs.microsoft.com/en-us/rest/api/azure/devops/graph/users?view=azure-devops-rest-7.0)
 */
/** @deprecated azuredevops.identities.getUsers has been deprecated in favor of azuredevops.getUsers */
export function getUsers(args?: GetUsersArgs, opts?: pulumi.InvokeOptions): Promise<GetUsersResult> {
    pulumi.log.warn("getUsers is deprecated: azuredevops.identities.getUsers has been deprecated in favor of azuredevops.getUsers")
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azuredevops:Identities/getUsers:getUsers", {
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
    origin?: string;
    /**
     * The unique identifier from the system of origin.
     *
     * DataSource without specifying any arguments will return all users inside an organization.
     *
     * List of possible subject types
     *
     * ```typescript
     * import * as pulumi from "@pulumi/pulumi";
     * ```
     *
     * List of possible origins
     *
     * ```typescript
     * import * as pulumi from "@pulumi/pulumi";
     * ```
     */
    originId?: string;
    /**
     * The PrincipalName of this graph member from the source provider.
     */
    principalName?: string;
    /**
     * A list of user subject subtypes to reduce the retrieved results, e.g. `msa`, `aad`, `svc` (service identity), `imp` (imported identity), etc. The supported subject types are listed below.
     */
    subjectTypes?: string[];
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
     * A set of existing users in your Azure DevOps Organization with details about every single user which includes:
     */
    readonly users: outputs.Identities.GetUsersUser[];
}
/**
 * Use this data source to access information about an existing users within Azure DevOps.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const example = azuredevops.getUsers({
 *     principalName: "contoso-user@contoso.onmicrosoft.com",
 * });
 * const example-all-users = azuredevops.getUsers({});
 * const example-all-from-origin = azuredevops.getUsers({
 *     origin: "aad",
 * });
 * const example-all-from-subjectTypes = azuredevops.getUsers({
 *     subjectTypes: [
 *         "aad",
 *         "msa",
 *     ],
 * });
 * const example-all-from-origin-id = azuredevops.getUsers({
 *     origin: "aad",
 *     originId: "00000000-0000-0000-0000-000000000000",
 * });
 * ```
 * ## Relevant Links
 *
 * - [Azure DevOps Service REST API 7.0 - Graph Users API](https://docs.microsoft.com/en-us/rest/api/azure/devops/graph/users?view=azure-devops-rest-7.0)
 */
/** @deprecated azuredevops.identities.getUsers has been deprecated in favor of azuredevops.getUsers */
export function getUsersOutput(args?: GetUsersOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetUsersResult> {
    return pulumi.output(args).apply((a: any) => getUsers(a, opts))
}

/**
 * A collection of arguments for invoking getUsers.
 */
export interface GetUsersOutputArgs {
    /**
     * The type of source provider for the `originId` parameter (ex:AD, AAD, MSA) The supported origins are listed below.
     */
    origin?: pulumi.Input<string>;
    /**
     * The unique identifier from the system of origin.
     *
     * DataSource without specifying any arguments will return all users inside an organization.
     *
     * List of possible subject types
     *
     * ```typescript
     * import * as pulumi from "@pulumi/pulumi";
     * ```
     *
     * List of possible origins
     *
     * ```typescript
     * import * as pulumi from "@pulumi/pulumi";
     * ```
     */
    originId?: pulumi.Input<string>;
    /**
     * The PrincipalName of this graph member from the source provider.
     */
    principalName?: pulumi.Input<string>;
    /**
     * A list of user subject subtypes to reduce the retrieved results, e.g. `msa`, `aad`, `svc` (service identity), `imp` (imported identity), etc. The supported subject types are listed below.
     */
    subjectTypes?: pulumi.Input<pulumi.Input<string>[]>;
}
