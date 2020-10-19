// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getIteration(args: GetIterationArgs, opts?: pulumi.InvokeOptions): Promise<GetIterationResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azuredevops:index/getIteration:getIteration", {
        "fetchChildren": args.fetchChildren,
        "path": args.path,
        "projectId": args.projectId,
    }, opts);
}

/**
 * A collection of arguments for invoking getIteration.
 */
export interface GetIterationArgs {
    readonly fetchChildren?: boolean;
    readonly path?: string;
    readonly projectId: string;
}

/**
 * A collection of values returned by getIteration.
 */
export interface GetIterationResult {
    readonly childrens: outputs.GetIterationChildren[];
    readonly fetchChildren?: boolean;
    readonly hasChildren: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    readonly path: string;
    readonly projectId: string;
}