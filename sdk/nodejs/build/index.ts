// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { BuildDefinitionArgs, BuildDefinitionState } from "./buildDefinition";
export type BuildDefinition = import("./buildDefinition").BuildDefinition;
export const BuildDefinition: typeof import("./buildDefinition").BuildDefinition = null as any;
utilities.lazyLoad(exports, ["BuildDefinition"], () => require("./buildDefinition"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "azuredevops:Build/buildDefinition:BuildDefinition":
                return new BuildDefinition(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("azuredevops", "Build/buildDefinition", _module)
