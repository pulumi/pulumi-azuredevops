// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./branchPolicyBuildValidation";
export * from "./branchPolicyMinReviewers";

// Import resources to register:
import { BranchPolicyBuildValidation } from "./branchPolicyBuildValidation";
import { BranchPolicyMinReviewers } from "./branchPolicyMinReviewers";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "azuredevops:Policy/branchPolicyBuildValidation:BranchPolicyBuildValidation":
                return new BranchPolicyBuildValidation(name, <any>undefined, { urn })
            case "azuredevops:Policy/branchPolicyMinReviewers:BranchPolicyMinReviewers":
                return new BranchPolicyMinReviewers(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("azuredevops", "Policy/branchPolicyBuildValidation", _module)
pulumi.runtime.registerResourceModule("azuredevops", "Policy/branchPolicyMinReviewers", _module)
