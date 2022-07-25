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
	"github.com/pulumi/pulumi-azuredevops/provider/v2/pkg/version"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
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
	p := shimv2.NewProvider(azuredevops.Provider())

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
			"azuredevops_build_definition_permissions":  {Tok: makeResource(mainMod, "BuildDefinitionPermissions")},
			"azuredevops_serviceendpoint_runpipeline":   {Tok: makeResource(mainMod, "ServiceEndpointPipeline")},
			"azuredevops_serviceendpoint_artifactory":   {Tok: makeResource(mainMod, "ServiceEndpointArtifactory")},
			"azuredevops_serviceendpoint_sonarqube":     {Tok: makeResource(mainMod, "ServiceEndpointSonarQube")},
			"azuredevops_serviceendpoint_ssh":           {Tok: makeResource(mainMod, "ServiceEndpointSsh")},
			"azuredevops_serviceendpoint_npm":           {Tok: makeResource(mainMod, "ServiceEndpointNpm")},
			"azuredevops_serviceendpoint_azuredevops":   {Tok: makeResource(mainMod, "ServiceEndpointAzureDevOps")},
			"azuredevops_serviceendpoint_generic":       {Tok: makeResource(mainMod, "ServiceEndpointGeneric")},
			"azuredevops_serviceendpoint_generic_git":   {Tok: makeResource(mainMod, "ServiceEndpointGenericGit")},
			"azuredevops_serviceendpoint_servicefabric": {Tok: makeResource(mainMod, "ServiceEndpointServiceFabric")},
			"azuredevops_serviceendpoint_github_enterprise": {
				Tok: makeResource(mainMod, "ServiceEndpointGitHubEnterprise"),
			},
			"azuredevops_branch_policy_merge_types":  {Tok: makeResource(mainMod, "BranchPolicyMergeTypes")},
			"azuredevops_branch_policy_status_check": {Tok: makeResource(mainMod, "BranchPolicyStatusCheck")},
			"azuredevops_repository_policy_author_email_pattern": {
				Tok: makeResource(mainMod, "RepositoryPolicyAuthorEmailPattern"),
			},
			"azuredevops_repository_policy_file_path_pattern": {
				Tok: makeResource(mainMod, "RepositoryPolicyFilePathPattern"),
			},
			"azuredevops_repository_policy_case_enforcement": {
				Tok: makeResource(mainMod, "RepositoryPolicyCaseEnforcement"),
			},
			"azuredevops_repository_policy_check_credentials": {
				Tok: makeResource(mainMod, "RepositoryPolicyCheckCredentials"),
			},
			"azuredevops_repository_policy_max_file_size": {
				Tok: makeResource(mainMod, "RepositoryPolicyMaxFileSize"),
			},
			"azuredevops_repository_policy_max_path_length": {
				Tok: makeResource(mainMod, "RepositoryPolicyMaxPathLength"),
			},
			"azuredevops_repository_policy_reserved_names": {
				Tok: makeResource(mainMod, "RepositoryPolicyReservedNames"),
			},
			"azuredevops_team":                            {Tok: makeResource(mainMod, "Team")},
			"azuredevops_team_administrators":             {Tok: makeResource(mainMod, "TeamAdministrators")},
			"azuredevops_team_members":                    {Tok: makeResource(mainMod, "TeamMembers")},
			"azuredevops_environment":                     {Tok: makeResource(mainMod, "Environment")},
			"azuredevops_git_repository_file":             {Tok: makeResource(mainMod, "GitRepositoryFile")},
			"azuredevops_serviceendpoint_argocd":          {Tok: makeResource(mainMod, "ServiceendpointArgocd")},
			"azuredevops_serviceendpoint_permissions":     {Tok: makeResource(mainMod, "ServiceendpointPermissions")},
			"azuredevops_servicehook_permissions":         {Tok: makeResource(mainMod, "ServicehookPermissions")},
			"azuredevops_tagging_permissions":             {Tok: makeResource(mainMod, "TaggingPermissions")},
			"azuredevops_project_pipeline_settings":       {Tok: makeResource(mainMod, "ProjectPipelineSettings")},
			"azuredevops_serviceendpoint_incomingwebhook": {Tok: makeResource(mainMod, "ServiceendpointIncomingwebhook")},
			"azuredevops_serviceendpoint_octopusdeploy":   {Tok: makeResource(mainMod, "ServiceendpointOctopusdeploy")},
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
			"azuredevops_team": {
				Tok: makeDataSource(mainMod, "getTeam"),
				Docs: &tfbridge.DocInfo{
					Source: "data_team.html.markdown",
				},
			},
			"azuredevops_teams": {
				Tok: makeDataSource(mainMod, "getTeams"),
				Docs: &tfbridge.DocInfo{
					Source: "data_teams.html.markdown",
				},
			},
			"azuredevops_groups":           {Tok: makeDataSource(mainMod, "getGroups")},
			"azuredevops_variable_group":   {Tok: makeDataSource(mainMod, "getVariableGroup")},
			"azuredevops_build_definition": {Tok: makeDataSource(mainMod, "getBuildDefinition")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
		},
		Python: &tfbridge.PythonInfo{
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
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
				"Pulumi": "3.*",
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
