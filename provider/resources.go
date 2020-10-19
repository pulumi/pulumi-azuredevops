// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package azuredevops

import (
	"unicode"

	"github.com/microsoft/terraform-provider-azuredevops/azuredevops"
	"github.com/pulumi/pulumi-terraform-bridge/v2/pkg/tfbridge"
	shim "github.com/pulumi/pulumi-terraform-bridge/v2/pkg/tfshim"
	shimv1 "github.com/pulumi/pulumi-terraform-bridge/v2/pkg/tfshim/sdk-v1"
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v2/go/common/tokens"
)

// all of the token components used below.
const (
	// packages:
	mainPkg = "azuredevops"

	// modules:
	mainMod = "index"
)

// makeMember manufactures a type token for the package and the given module and type.
func makeMember(mod string, mem string) tokens.ModuleMember {
	return tokens.ModuleMember(mainPkg + ":" + mod + ":" + mem)
}

// makeType manufactures a type token for the package and the given module and type.
func makeType(mod string, typ string) tokens.Type {
	return tokens.Type(makeMember(mod, typ))
}

// makeDataSource manufactures a standard resource token given a module and resource name.  It
// automatically uses the main package and names the file by simply lower casing the data source's
// first character.
func makeDataSource(mod string, res string) tokens.ModuleMember {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return makeMember(mod+"/"+fn, res)
}

// makeResource manufactures a standard resource token given a module and resource name.  It
// automatically uses the main package and names the file by simply lower casing the resource's
// first character.
func makeResource(mod string, res string) tokens.Type {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return makeType(mod+"/"+fn, res)
}

// preConfigureCallback is called before the providerConfigure function of the underlying provider.
// It should validate that the provider can be configured, and provide actionable errors in the case
// it cannot be. Configuration variables can be read from `vars` using the `stringValue` function -
// for example `stringValue(vars, "accessKey")`.
func preConfigureCallback(vars resource.PropertyMap, c shim.ResourceConfig) error {
	return nil
}

// Provider returns additional overlaid schema and metadata associated with the provider.
func Provider() tfbridge.ProviderInfo {
	p := shimv1.NewProvider(azuredevops.Provider())

	prov := tfbridge.ProviderInfo{
		P:           p,
		Name:        "azuredevops",
		Description: "A Pulumi package for creating and managing Azure DevOps.",
		Keywords:    []string{"pulumi", "azuredevops"},
		License:     "Apache-2.0",
		Homepage:    "https://pulumi.io",
		Repository:  "https://github.com/pulumi/pulumi-azuredevops",
		GitHubOrg:   "microsoft",
		Config: map[string]*tfbridge.SchemaInfo{
			"org_service_url": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"AZDO_ORG_SERVICE_URL"},
				},
			},
			"personal_access_token": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"AZDO_PERSONAL_ACCESS_TOKEN"},
				},
			},
		},
		PreConfigureCallback: preConfigureCallback,
		Resources: map[string]*tfbridge.ResourceInfo{
			"azuredevops_area_permissions": {
				Tok: makeResource(mainMod, "AreaPermissions"),
			},
			"azuredevops_branch_policy_auto_reviewers": {
				Tok: makeResource(mainMod, "BranchPolicyAutoReviewers"),
			},
			"azuredevops_branch_policy_comment_resolution": {
				Tok: makeResource(mainMod, "BranchPolicyCommentResolution"),
			},
			"azuredevops_branch_policy_work_item_linking": {
				Tok: makeResource(mainMod, "BranchPolicyWorkItemLinking"),
			},
			"azuredevops_git_permissions": {
				Tok: makeResource(mainMod, "GitPermissions"),
			},
			"azuredevops_iteration_permissions": {
				Tok: makeResource(mainMod, "IterativePermissions"),
			},
			"azuredevops_project_permissions": {
				Tok: makeResource(mainMod, "ProjectPermissions"),
			},
			"azuredevops_serviceendpoint_aws": {
				Tok: makeResource(mainMod, "ServiceEndpointAws"),
			},
			"azuredevops_serviceendpoint_azurecr": {
				Tok: makeResource(mainMod, "ServiceEndpointAzureEcr"),
			},
			"azuredevops_workitemquery_permissions": {
				Tok: makeResource(mainMod, "WorkItemQueryPermissions"),
			},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"azuredevops_area": {
				Tok: makeDataSource(mainMod, "getArea"),
			},
			"azuredevops_git_repository": {
				Tok: makeDataSource(mainMod, "getGitRepository"),
			},
			"azuredevops_iteration": {
				Tok: makeDataSource(mainMod, "getIteration"),
			},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^2.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^8.0.25", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
		},
		Python: &tfbridge.PythonInfo{
			Requires: map[string]string{
				"pulumi": ">=2.9.0,<3.0.0",
			},
			UsesIOClasses: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi":                       "2.*",
				"System.Collections.Immutable": "1.6.0",
			},
			Namespaces: map[string]string{
				mainPkg: "AzureDevOps",
			},
		},
	}

	prov.RenameResourceWithAlias("azuredevops_resource_authorization",
		makeResource("Security", "ResourceAuthorization"),
		makeResource(mainMod, "ResourceAuthorization"), "Security", mainMod, &tfbridge.ResourceInfo{
			Docs: &tfbridge.DocInfo{Source: "resource_authorization.html.markdown"},
		})
	prov.RenameResourceWithAlias("azuredevops_branch_policy_build_validation",
		makeResource("Policy", "BranchPolicyBuildValidation"),
		makeResource(mainMod, "BranchPolicyBuildValidation"), "Policy", mainMod, &tfbridge.ResourceInfo{
			Docs: &tfbridge.DocInfo{Source: "branch_policy_build_validation.html.markdown"},
		})
	prov.RenameResourceWithAlias("azuredevops_branch_policy_min_reviewers",
		makeResource("Policy", "BranchPolicyMinReviewers"),
		makeResource(mainMod, "BranchPolicyMinReviewers"), "Policy", mainMod, &tfbridge.ResourceInfo{
			Docs: &tfbridge.DocInfo{Source: "branch_policy_min_reviewers.html.markdown"},
		})
	prov.RenameResourceWithAlias("azuredevops_build_definition",
		makeResource("Build", "BuildDefinition"),
		makeResource(mainMod, "BuildDefinition"), "Build", mainMod, &tfbridge.ResourceInfo{
			Docs: &tfbridge.DocInfo{Source: "build_definition.html.markdown"},
		})
	prov.RenameResourceWithAlias("azuredevops_project",
		makeResource("Core", "Project"),
		makeResource(mainMod, "Project"), "Core", mainMod, &tfbridge.ResourceInfo{
			Docs: &tfbridge.DocInfo{Source: "project.html.markdown"},
		})
	prov.RenameResourceWithAlias("azuredevops_project_features",
		makeResource("Core", "ProjectFeatures"),
		makeResource(mainMod, "ProjectFeatures"), "Core", mainMod, &tfbridge.ResourceInfo{
			Docs: &tfbridge.DocInfo{Source: "project_features.html.markdown"},
		})
	prov.RenameResourceWithAlias("azuredevops_variable_group",
		makeResource("Pipeline", "VariableGroup"),
		makeResource(mainMod, "VariableGroup"), "Pipeline", mainMod, &tfbridge.ResourceInfo{
			Docs: &tfbridge.DocInfo{Source: "variable_group.html.markdown"},
		})
	prov.RenameResourceWithAlias("azuredevops_serviceendpoint_azurerm",
		makeResource("ServiceEndpoint", "AzureRM"),
		makeResource(mainMod, "ServiceEndpointAzureRM"), "ServiceEndpoint", mainMod, &tfbridge.ResourceInfo{
			Docs: &tfbridge.DocInfo{Source: "serviceendpoint_azurerm.html.markdown"},
		})
	prov.RenameResourceWithAlias("azuredevops_serviceendpoint_bitbucket",
		makeResource("ServiceEndpoint", "BitBucket"),
		makeResource(mainMod, "ServiceEndpointBitBucket"), "ServiceEndpoint", mainMod, &tfbridge.ResourceInfo{
			Docs: &tfbridge.DocInfo{Source: "serviceendpoint_bitbucket.html.markdown"},
		})
	prov.RenameResourceWithAlias("azuredevops_serviceendpoint_dockerregistry",
		makeResource("ServiceEndpoint", "DockerRegistry"),
		makeResource(mainMod, "ServiceEndpointDockerRegistry"), "ServiceEndpoint", mainMod, &tfbridge.ResourceInfo{
			Docs: &tfbridge.DocInfo{
				Source: "serviceendpoint_dockerregistry.html.markdown",
			},
			Fields: map[string]*tfbridge.SchemaInfo{
				"docker_registry": {
					CSharpName: "DockerRegistryUrl",
				},
			},
		})
	prov.RenameResourceWithAlias("azuredevops_serviceendpoint_github",
		makeResource("ServiceEndpoint", "GitHub"),
		makeResource(mainMod, "ServiceEndpointGitHub"), "ServiceEndpoint", mainMod, &tfbridge.ResourceInfo{
			Docs: &tfbridge.DocInfo{Source: "serviceendpoint_github.html.markdown"},
		})
	prov.RenameResourceWithAlias("azuredevops_serviceendpoint_kubernetes",
		makeResource("ServiceEndpoint", "Kubernetes"),
		makeResource(mainMod, "ServiceEndpointKubernetes"), "ServiceEndpoint", mainMod, &tfbridge.ResourceInfo{
			Docs: &tfbridge.DocInfo{Source: "serviceendpoint_kubernetes.html.markdown"},
		})
	prov.RenameResourceWithAlias("azuredevops_git_repository",
		makeResource("Repository", "Git"),
		makeResource(mainMod, "Git"), "Repository", mainMod, &tfbridge.ResourceInfo{
			Docs: &tfbridge.DocInfo{Source: "azure_git_repository.html.markdown"},
		})
	prov.RenameResourceWithAlias("azuredevops_user_entitlement",
		makeResource("Entitlement", "User"),
		makeResource(mainMod, "User"), "Entitlement", mainMod, &tfbridge.ResourceInfo{
			Docs: &tfbridge.DocInfo{Source: "user_entitlement.html.markdown"},
		})
	prov.RenameResourceWithAlias("azuredevops_group_membership",
		makeResource("Identities", "GroupMembership"),
		makeResource(mainMod, "GroupMembership"), "Identities", mainMod, &tfbridge.ResourceInfo{
			Docs: &tfbridge.DocInfo{Source: "group_membership.html.markdown"},
		})
	prov.RenameResourceWithAlias("azuredevops_agent_pool",
		makeResource("Agent", "Pool"),
		makeResource(mainMod, "Pool"), "Agent", mainMod, &tfbridge.ResourceInfo{
			Docs: &tfbridge.DocInfo{Source: "agent_pool.html.markdown"},
		})
	prov.RenameResourceWithAlias("azuredevops_agent_queue",
		makeResource("Agent", "Queue"),
		makeResource(mainMod, "Queue"), "Agent", mainMod, &tfbridge.ResourceInfo{
			Docs: &tfbridge.DocInfo{Source: "agent_queue.html.markdown"},
		})
	prov.RenameResourceWithAlias("azuredevops_group",
		makeResource("Identities", "Group"),
		makeResource(mainMod, "Group"), "Identities", mainMod, &tfbridge.ResourceInfo{
			Docs: &tfbridge.DocInfo{Source: "group.html.markdown"},
		})

	prov.RenameDataSource("azuredevops_agent_pool", makeDataSource("Agent", "getPool"),
		makeDataSource(mainMod, "getPool"), "Agent", mainMod, &tfbridge.DataSourceInfo{
			Docs: &tfbridge.DocInfo{Source: "data_agent_pool.html.markdown"},
		})
	prov.RenameDataSource("azuredevops_agent_pools", makeDataSource("Agent", "getPools"),
		makeDataSource(mainMod, "getPools"), "Agent", mainMod, &tfbridge.DataSourceInfo{
			Docs: &tfbridge.DocInfo{Source: "data_agent_pools.html.markdown"},
		})
	prov.RenameDataSource("azuredevops_client_config", makeDataSource("Core", "getClientConfig"),
		makeDataSource(mainMod, "getClientConfig"), "Core", mainMod, &tfbridge.DataSourceInfo{
			Docs: &tfbridge.DocInfo{Source: "data_client_config.html.markdown"},
		})
	prov.RenameDataSource("azuredevops_group", makeDataSource("Identities", "getGroup"),
		makeDataSource(mainMod, "getGroup"), "Identities", mainMod, &tfbridge.DataSourceInfo{
			Docs: &tfbridge.DocInfo{Source: "data_group.html.markdown"},
		})
	prov.RenameDataSource("azuredevops_project", makeDataSource("Core", "getProject"),
		makeDataSource(mainMod, "getProject"), "Core", mainMod, &tfbridge.DataSourceInfo{
			Docs: &tfbridge.DocInfo{Source: "data_project.html.markdown"},
		})
	prov.RenameDataSource("azuredevops_projects", makeDataSource("Core", "getProjects"),
		makeDataSource(mainMod, "getProjects"), "Core", mainMod, &tfbridge.DataSourceInfo{
			Docs: &tfbridge.DocInfo{Source: "data_projects.html.markdown"},
		})
	prov.RenameDataSource("azuredevops_git_repositories",
		makeDataSource("Repository", "getRepositories"),
		makeDataSource(mainMod, "getRepositories"), "Repository", mainMod, &tfbridge.DataSourceInfo{
			Docs: &tfbridge.DocInfo{Source: "data_git_repositories.html.markdown"},
		})
	prov.RenameDataSource("azuredevops_users",
		makeDataSource("Identities", "getUsers"),
		makeDataSource(mainMod, "getUsers"), "Identities", mainMod, &tfbridge.DataSourceInfo{
			Docs: &tfbridge.DocInfo{Source: "data_users.html.markdown"},
		})

	prov.SetAutonaming(255, "-")

	return prov
}
