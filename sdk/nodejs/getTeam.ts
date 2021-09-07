// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export function getTeam(args: GetTeamArgs, opts?: pulumi.InvokeOptions): Promise<GetTeamResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azuredevops:index/getTeam:getTeam", {
        "name": args.name,
        "projectId": args.projectId,
    }, opts);
}

/**
 * A collection of arguments for invoking getTeam.
 */
export interface GetTeamArgs {
    name: string;
    projectId: string;
}

/**
 * A collection of values returned by getTeam.
 */
export interface GetTeamResult {
    readonly administrators: string[];
    readonly description: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly members: string[];
    readonly name: string;
    readonly projectId: string;
}
