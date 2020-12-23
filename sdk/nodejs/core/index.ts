// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./getClientConfig";
export * from "./getProject";
export * from "./getProjects";
export * from "./project";
export * from "./projectFeatures";

// Import resources to register:
import { Project } from "./project";
import { ProjectFeatures } from "./projectFeatures";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "azuredevops:Core/project:Project":
                return new Project(name, <any>undefined, { urn })
            case "azuredevops:Core/projectFeatures:ProjectFeatures":
                return new ProjectFeatures(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("azuredevops", "Core/project", _module)
pulumi.runtime.registerResourceModule("azuredevops", "Core/projectFeatures", _module)
