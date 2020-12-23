// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

// Export members:
export * from "./areaPermissions";
export * from "./branchPolicyAutoReviewers";
export * from "./branchPolicyBuildValidation";
export * from "./branchPolicyCommentResolution";
export * from "./branchPolicyMinReviewers";
export * from "./branchPolicyWorkItemLinking";
export * from "./buildDefinition";
export * from "./getAgentQueue";
export * from "./getArea";
export * from "./getClientConfig";
export * from "./getGitRepository";
export * from "./getGroup";
export * from "./getIteration";
export * from "./getPool";
export * from "./getPools";
export * from "./getProject";
export * from "./getProjects";
export * from "./getRepositories";
export * from "./getUsers";
export * from "./git";
export * from "./gitPermissions";
export * from "./group";
export * from "./groupMembership";
export * from "./iterativePermissions";
export * from "./pool";
export * from "./project";
export * from "./projectFeatures";
export * from "./projectPermissions";
export * from "./provider";
export * from "./queue";
export * from "./resourceAuthorization";
export * from "./serviceEndpointAws";
export * from "./serviceEndpointAzureEcr";
export * from "./serviceEndpointAzureRM";
export * from "./serviceEndpointBitBucket";
export * from "./serviceEndpointDockerRegistry";
export * from "./serviceEndpointGitHub";
export * from "./serviceEndpointKubernetes";
export * from "./user";
export * from "./variableGroup";
export * from "./workItemQueryPermissions";

// Export sub-modules:
import * as agent from "./agent";
import * as build from "./build";
import * as config from "./config";
import * as core from "./core";
import * as entitlement from "./entitlement";
import * as identities from "./identities";
import * as pipeline from "./pipeline";
import * as policy from "./policy";
import * as repository from "./repository";
import * as security from "./security";
import * as serviceendpoint from "./serviceendpoint";
import * as types from "./types";

export {
    agent,
    build,
    config,
    core,
    entitlement,
    identities,
    pipeline,
    policy,
    repository,
    security,
    serviceendpoint,
    types,
};

// Import resources to register:
import { AreaPermissions } from "./areaPermissions";
import { BranchPolicyAutoReviewers } from "./branchPolicyAutoReviewers";
import { BranchPolicyBuildValidation } from "./branchPolicyBuildValidation";
import { BranchPolicyCommentResolution } from "./branchPolicyCommentResolution";
import { BranchPolicyMinReviewers } from "./branchPolicyMinReviewers";
import { BranchPolicyWorkItemLinking } from "./branchPolicyWorkItemLinking";
import { BuildDefinition } from "./buildDefinition";
import { Git } from "./git";
import { GitPermissions } from "./gitPermissions";
import { Group } from "./group";
import { GroupMembership } from "./groupMembership";
import { IterativePermissions } from "./iterativePermissions";
import { Pool } from "./pool";
import { Project } from "./project";
import { ProjectFeatures } from "./projectFeatures";
import { ProjectPermissions } from "./projectPermissions";
import { Queue } from "./queue";
import { ResourceAuthorization } from "./resourceAuthorization";
import { ServiceEndpointAws } from "./serviceEndpointAws";
import { ServiceEndpointAzureEcr } from "./serviceEndpointAzureEcr";
import { ServiceEndpointAzureRM } from "./serviceEndpointAzureRM";
import { ServiceEndpointBitBucket } from "./serviceEndpointBitBucket";
import { ServiceEndpointDockerRegistry } from "./serviceEndpointDockerRegistry";
import { ServiceEndpointGitHub } from "./serviceEndpointGitHub";
import { ServiceEndpointKubernetes } from "./serviceEndpointKubernetes";
import { User } from "./user";
import { VariableGroup } from "./variableGroup";
import { WorkItemQueryPermissions } from "./workItemQueryPermissions";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "azuredevops:index/areaPermissions:AreaPermissions":
                return new AreaPermissions(name, <any>undefined, { urn })
            case "azuredevops:index/branchPolicyAutoReviewers:BranchPolicyAutoReviewers":
                return new BranchPolicyAutoReviewers(name, <any>undefined, { urn })
            case "azuredevops:index/branchPolicyBuildValidation:BranchPolicyBuildValidation":
                return new BranchPolicyBuildValidation(name, <any>undefined, { urn })
            case "azuredevops:index/branchPolicyCommentResolution:BranchPolicyCommentResolution":
                return new BranchPolicyCommentResolution(name, <any>undefined, { urn })
            case "azuredevops:index/branchPolicyMinReviewers:BranchPolicyMinReviewers":
                return new BranchPolicyMinReviewers(name, <any>undefined, { urn })
            case "azuredevops:index/branchPolicyWorkItemLinking:BranchPolicyWorkItemLinking":
                return new BranchPolicyWorkItemLinking(name, <any>undefined, { urn })
            case "azuredevops:index/buildDefinition:BuildDefinition":
                return new BuildDefinition(name, <any>undefined, { urn })
            case "azuredevops:index/git:Git":
                return new Git(name, <any>undefined, { urn })
            case "azuredevops:index/gitPermissions:GitPermissions":
                return new GitPermissions(name, <any>undefined, { urn })
            case "azuredevops:index/group:Group":
                return new Group(name, <any>undefined, { urn })
            case "azuredevops:index/groupMembership:GroupMembership":
                return new GroupMembership(name, <any>undefined, { urn })
            case "azuredevops:index/iterativePermissions:IterativePermissions":
                return new IterativePermissions(name, <any>undefined, { urn })
            case "azuredevops:index/pool:Pool":
                return new Pool(name, <any>undefined, { urn })
            case "azuredevops:index/project:Project":
                return new Project(name, <any>undefined, { urn })
            case "azuredevops:index/projectFeatures:ProjectFeatures":
                return new ProjectFeatures(name, <any>undefined, { urn })
            case "azuredevops:index/projectPermissions:ProjectPermissions":
                return new ProjectPermissions(name, <any>undefined, { urn })
            case "azuredevops:index/queue:Queue":
                return new Queue(name, <any>undefined, { urn })
            case "azuredevops:index/resourceAuthorization:ResourceAuthorization":
                return new ResourceAuthorization(name, <any>undefined, { urn })
            case "azuredevops:index/serviceEndpointAws:ServiceEndpointAws":
                return new ServiceEndpointAws(name, <any>undefined, { urn })
            case "azuredevops:index/serviceEndpointAzureEcr:ServiceEndpointAzureEcr":
                return new ServiceEndpointAzureEcr(name, <any>undefined, { urn })
            case "azuredevops:index/serviceEndpointAzureRM:ServiceEndpointAzureRM":
                return new ServiceEndpointAzureRM(name, <any>undefined, { urn })
            case "azuredevops:index/serviceEndpointBitBucket:ServiceEndpointBitBucket":
                return new ServiceEndpointBitBucket(name, <any>undefined, { urn })
            case "azuredevops:index/serviceEndpointDockerRegistry:ServiceEndpointDockerRegistry":
                return new ServiceEndpointDockerRegistry(name, <any>undefined, { urn })
            case "azuredevops:index/serviceEndpointGitHub:ServiceEndpointGitHub":
                return new ServiceEndpointGitHub(name, <any>undefined, { urn })
            case "azuredevops:index/serviceEndpointKubernetes:ServiceEndpointKubernetes":
                return new ServiceEndpointKubernetes(name, <any>undefined, { urn })
            case "azuredevops:index/user:User":
                return new User(name, <any>undefined, { urn })
            case "azuredevops:index/variableGroup:VariableGroup":
                return new VariableGroup(name, <any>undefined, { urn })
            case "azuredevops:index/workItemQueryPermissions:WorkItemQueryPermissions":
                return new WorkItemQueryPermissions(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("azuredevops", "index/areaPermissions", _module)
pulumi.runtime.registerResourceModule("azuredevops", "index/branchPolicyAutoReviewers", _module)
pulumi.runtime.registerResourceModule("azuredevops", "index/branchPolicyBuildValidation", _module)
pulumi.runtime.registerResourceModule("azuredevops", "index/branchPolicyCommentResolution", _module)
pulumi.runtime.registerResourceModule("azuredevops", "index/branchPolicyMinReviewers", _module)
pulumi.runtime.registerResourceModule("azuredevops", "index/branchPolicyWorkItemLinking", _module)
pulumi.runtime.registerResourceModule("azuredevops", "index/buildDefinition", _module)
pulumi.runtime.registerResourceModule("azuredevops", "index/git", _module)
pulumi.runtime.registerResourceModule("azuredevops", "index/gitPermissions", _module)
pulumi.runtime.registerResourceModule("azuredevops", "index/group", _module)
pulumi.runtime.registerResourceModule("azuredevops", "index/groupMembership", _module)
pulumi.runtime.registerResourceModule("azuredevops", "index/iterativePermissions", _module)
pulumi.runtime.registerResourceModule("azuredevops", "index/pool", _module)
pulumi.runtime.registerResourceModule("azuredevops", "index/project", _module)
pulumi.runtime.registerResourceModule("azuredevops", "index/projectFeatures", _module)
pulumi.runtime.registerResourceModule("azuredevops", "index/projectPermissions", _module)
pulumi.runtime.registerResourceModule("azuredevops", "index/queue", _module)
pulumi.runtime.registerResourceModule("azuredevops", "index/resourceAuthorization", _module)
pulumi.runtime.registerResourceModule("azuredevops", "index/serviceEndpointAws", _module)
pulumi.runtime.registerResourceModule("azuredevops", "index/serviceEndpointAzureEcr", _module)
pulumi.runtime.registerResourceModule("azuredevops", "index/serviceEndpointAzureRM", _module)
pulumi.runtime.registerResourceModule("azuredevops", "index/serviceEndpointBitBucket", _module)
pulumi.runtime.registerResourceModule("azuredevops", "index/serviceEndpointDockerRegistry", _module)
pulumi.runtime.registerResourceModule("azuredevops", "index/serviceEndpointGitHub", _module)
pulumi.runtime.registerResourceModule("azuredevops", "index/serviceEndpointKubernetes", _module)
pulumi.runtime.registerResourceModule("azuredevops", "index/user", _module)
pulumi.runtime.registerResourceModule("azuredevops", "index/variableGroup", _module)
pulumi.runtime.registerResourceModule("azuredevops", "index/workItemQueryPermissions", _module)

import { Provider } from "./provider";

pulumi.runtime.registerResourcePackage("azuredevops", {
    version: utilities.getVersion(),
    constructProvider: (name: string, type: string, urn: string): pulumi.ProviderResource => {
        if (type !== "pulumi:providers:azuredevops") {
            throw new Error(`unknown provider type ${type}`);
        }
        return new Provider(name, <any>undefined, { urn });
    },
});
