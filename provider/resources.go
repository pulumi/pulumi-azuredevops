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
	"fmt"
	"path/filepath"
	"unicode"

	"github.com/microsoft/terraform-provider-azuredevops/azuredevops"
	"github.com/pulumi/pulumi-azuredevops/provider/pkg/version"
	"github.com/pulumi/pulumi-terraform-bridge/v2/pkg/tfbridge"
	shimv1 "github.com/pulumi/pulumi-terraform-bridge/v2/pkg/tfshim/sdk-v1"
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
			"azuredevops_build_definition_permissions": {Tok: makeResource(mainMod, "BuildDefinitionPermissions")},
			"azuredevops_serviceendpoint_runpipeline":  {Tok: makeResource(mainMod, "ServiceEndpointPipeline")},
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
			"azuredevops_agent_queue": {
				Tok: makeDataSource(mainMod, "getAgentQueue"),
			},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^2.15.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^8.0.25", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
		},
		Python: &tfbridge.PythonInfo{
			Requires: map[string]string{
				"pulumi": ">=2.15.0,<3.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/pulumi/pulumi-%[1]s/sdk/", mainPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
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
		makeResource(mainMod, "ResourceAuthorization"), "Security", mainMod, nil)
	prov.RenameResourceWithAlias("azuredevops_branch_policy_build_validation",
		makeResource("Policy", "BranchPolicyBuildValidation"),
		makeResource(mainMod, "BranchPolicyBuildValidation"), "Policy", mainMod, nil)
	prov.RenameResourceWithAlias("azuredevops_branch_policy_min_reviewers",
		makeResource("Policy", "BranchPolicyMinReviewers"),
		makeResource(mainMod, "BranchPolicyMinReviewers"), "Policy", mainMod, nil)
	prov.RenameResourceWithAlias("azuredevops_build_definition",
		makeResource("Build", "BuildDefinition"),
		makeResource(mainMod, "BuildDefinition"), "Build", mainMod, nil)
	prov.RenameResourceWithAlias("azuredevops_project",
		makeResource("Core", "Project"),
		makeResource(mainMod, "Project"), "Core", mainMod, nil)
	prov.RenameResourceWithAlias("azuredevops_project_features",
		makeResource("Core", "ProjectFeatures"),
		makeResource(mainMod, "ProjectFeatures"), "Core", mainMod, nil)
	prov.RenameResourceWithAlias("azuredevops_variable_group",
		makeResource("Pipeline", "VariableGroup"),
		makeResource(mainMod, "VariableGroup"), "Pipeline", mainMod, nil)
	prov.RenameResourceWithAlias("azuredevops_serviceendpoint_azurerm",
		makeResource("ServiceEndpoint", "AzureRM"),
		makeResource(mainMod, "ServiceEndpointAzureRM"), "ServiceEndpoint", mainMod, nil)
	prov.RenameResourceWithAlias("azuredevops_serviceendpoint_bitbucket",
		makeResource("ServiceEndpoint", "BitBucket"),
		makeResource(mainMod, "ServiceEndpointBitBucket"), "ServiceEndpoint", mainMod, nil)
	prov.RenameResourceWithAlias("azuredevops_serviceendpoint_dockerregistry",
		makeResource("ServiceEndpoint", "DockerRegistry"),
		makeResource(mainMod, "ServiceEndpointDockerRegistry"), "ServiceEndpoint", mainMod, &tfbridge.ResourceInfo{
			Fields: map[string]*tfbridge.SchemaInfo{
				"docker_registry": {
					CSharpName: "DockerRegistryUrl",
				},
			},
		})
	prov.RenameResourceWithAlias("azuredevops_serviceendpoint_github",
		makeResource("ServiceEndpoint", "GitHub"),
		makeResource(mainMod, "ServiceEndpointGitHub"), "ServiceEndpoint", mainMod, nil)
	prov.RenameResourceWithAlias("azuredevops_serviceendpoint_kubernetes",
		makeResource("ServiceEndpoint", "Kubernetes"),
		makeResource(mainMod, "ServiceEndpointKubernetes"), "ServiceEndpoint", mainMod, nil)
	prov.RenameResourceWithAlias("azuredevops_git_repository",
		makeResource("Repository", "Git"),
		makeResource(mainMod, "Git"), "Repository", mainMod, nil)
	prov.RenameResourceWithAlias("azuredevops_user_entitlement",
		makeResource("Entitlement", "User"),
		makeResource(mainMod, "User"), "Entitlement", mainMod, nil)
	prov.RenameResourceWithAlias("azuredevops_group_membership",
		makeResource("Identities", "GroupMembership"),
		makeResource(mainMod, "GroupMembership"), "Identities", mainMod, nil)
	prov.RenameResourceWithAlias("azuredevops_agent_pool",
		makeResource("Agent", "Pool"),
		makeResource(mainMod, "Pool"), "Agent", mainMod, nil)
	prov.RenameResourceWithAlias("azuredevops_agent_queue",
		makeResource("Agent", "Queue"),
		makeResource(mainMod, "Queue"), "Agent", mainMod, nil)
	prov.RenameResourceWithAlias("azuredevops_group",
		makeResource("Identities", "Group"),
		makeResource(mainMod, "Group"), "Identities", mainMod, nil)

	prov.RenameDataSource("azuredevops_agent_pool", makeDataSource("Agent", "getPool"),
		makeDataSource(mainMod, "getPool"), "Agent", mainMod, nil)
	prov.RenameDataSource("azuredevops_agent_pools", makeDataSource("Agent", "getPools"),
		makeDataSource(mainMod, "getPools"), "Agent", mainMod, nil)
	prov.RenameDataSource("azuredevops_client_config", makeDataSource("Core", "getClientConfig"),
		makeDataSource(mainMod, "getClientConfig"), "Core", mainMod, nil)
	prov.RenameDataSource("azuredevops_group", makeDataSource("Identities", "getGroup"),
		makeDataSource(mainMod, "getGroup"), "Identities", mainMod, nil)
	prov.RenameDataSource("azuredevops_project", makeDataSource("Core", "getProject"),
		makeDataSource(mainMod, "getProject"), "Core", mainMod, nil)
	prov.RenameDataSource("azuredevops_projects", makeDataSource("Core", "getProjects"),
		makeDataSource(mainMod, "getProjects"), "Core", mainMod, nil)
	prov.RenameDataSource("azuredevops_git_repositories",
		makeDataSource("Repository", "getRepositories"),
		makeDataSource(mainMod, "getRepositories"), "Repository", mainMod, nil)
	prov.RenameDataSource("azuredevops_users",
		makeDataSource("Identities", "getUsers"),
		makeDataSource(mainMod, "getUsers"), "Identities", mainMod, nil)

	prov.SetAutonaming(255, "-")

	return prov
}
