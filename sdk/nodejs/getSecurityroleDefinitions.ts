// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Use this data source to access information about existing Security Role Definitions within a given scope in Azure DevOps.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const example = azuredevops.getSecurityroleDefinitions({
 *     scope: "distributedtask.environmentreferencerole",
 * });
 * export const securityroleDefinitions = exampleAazuredevopsSecurityroleDefinitions.definitions;
 * ```
 *
 * ## Relevant Links
 *
 * - [Azure DevOps Service REST API 7.0 - Roledefinitions - List](https://learn.microsoft.com/en-us/rest/api/azure/devops/securityroles/roledefinitions/list?view=azure-devops-rest-7.1)
 */
export function getSecurityroleDefinitions(args: GetSecurityroleDefinitionsArgs, opts?: pulumi.InvokeOptions): Promise<GetSecurityroleDefinitionsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azuredevops:index/getSecurityroleDefinitions:getSecurityroleDefinitions", {
        "scope": args.scope,
    }, opts);
}

/**
 * A collection of arguments for invoking getSecurityroleDefinitions.
 */
export interface GetSecurityroleDefinitionsArgs {
    /**
     * Name of the Scope for which Security Role Definitions will be returned.
     *
     * DataSource without specifying any arguments will return all projects.
     */
    scope: string;
}

/**
 * A collection of values returned by getSecurityroleDefinitions.
 */
export interface GetSecurityroleDefinitionsResult {
    /**
     * A list of existing Security Role Definitions in a Scope in your Azure DevOps Organization with details about every definition which includes. A `definitions` block as defined below.
     */
    readonly definitions: outputs.GetSecurityroleDefinitionsDefinition[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The scope of the Security Role Definition.
     */
    readonly scope: string;
}
/**
 * Use this data source to access information about existing Security Role Definitions within a given scope in Azure DevOps.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const example = azuredevops.getSecurityroleDefinitions({
 *     scope: "distributedtask.environmentreferencerole",
 * });
 * export const securityroleDefinitions = exampleAazuredevopsSecurityroleDefinitions.definitions;
 * ```
 *
 * ## Relevant Links
 *
 * - [Azure DevOps Service REST API 7.0 - Roledefinitions - List](https://learn.microsoft.com/en-us/rest/api/azure/devops/securityroles/roledefinitions/list?view=azure-devops-rest-7.1)
 */
export function getSecurityroleDefinitionsOutput(args: GetSecurityroleDefinitionsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSecurityroleDefinitionsResult> {
    return pulumi.output(args).apply((a: any) => getSecurityroleDefinitions(a, opts))
}

/**
 * A collection of arguments for invoking getSecurityroleDefinitions.
 */
export interface GetSecurityroleDefinitionsOutputArgs {
    /**
     * Name of the Scope for which Security Role Definitions will be returned.
     *
     * DataSource without specifying any arguments will return all projects.
     */
    scope: pulumi.Input<string>;
}
