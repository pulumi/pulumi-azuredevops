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
	"path"

	// Allow embedding provider metadata
	_ "embed"

	"github.com/microsoft/terraform-provider-azuredevops/azuredevops"
	"github.com/pulumi/pulumi-azuredevops/provider/v2/pkg/version"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	tks "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge/tokens"
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

//go:embed cmd/pulumi-resource-azuredevops/bridge-metadata.json
var metadata []byte

// Provider returns additional overlaid schema and metadata associated with the provider.
func Provider() tfbridge.ProviderInfo {
	p := shimv2.NewProvider(azuredevops.Provider())

	makeResource := func(name string) tokens.Type {
		return tfbridge.MakeResource(mainPkg, mainMod, name)
	}

	makeDataSource := func(name string) tokens.ModuleMember {
		return tfbridge.MakeDataSource(mainPkg, mainMod, name)
	}

	prov := tfbridge.ProviderInfo{
		P:                p,
		Name:             "azuredevops",
		Description:      "A Pulumi package for creating and managing Azure DevOps.",
		Keywords:         []string{"pulumi", "azuredevops"},
		License:          "Apache-2.0",
		Homepage:         "https://pulumi.io",
		Repository:       "https://github.com/pulumi/pulumi-azuredevops",
		GitHubOrg:        "microsoft",
		UpstreamRepoPath: "./upstream",
		Version:          version.Version,
		MetadataInfo:     tfbridge.NewProviderMetadata(metadata),
		Config: map[string]*tfbridge.SchemaInfo{
			"org_service_url": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"AZDO_ORG_SERVICE_URL"},
				},
			},
		},
		Resources: map[string]*tfbridge.ResourceInfo{
			"azuredevops_serviceendpoint_aws": {
				Tok: makeResource("ServiceEndpointAws"),
			},
			"azuredevops_serviceendpoint_azurecr": {
				Tok: makeResource("ServiceEndpointAzureEcr"),
			},
			"azuredevops_workitemquery_permissions": {
				Tok: makeResource("WorkItemQueryPermissions"),
			},
			"azuredevops_serviceendpoint_runpipeline": {Tok: makeResource("ServiceEndpointPipeline")},
			"azuredevops_serviceendpoint_artifactory": {Tok: makeResource("ServiceEndpointArtifactory")},
			"azuredevops_serviceendpoint_sonarqube":   {Tok: makeResource("ServiceEndpointSonarQube")},
			"azuredevops_serviceendpoint_ssh":         {Tok: makeResource("ServiceEndpointSsh")},
			"azuredevops_serviceendpoint_npm":         {Tok: makeResource("ServiceEndpointNpm")},
			"azuredevops_serviceendpoint_azuredevops": {Tok: makeResource("ServiceEndpointAzureDevOps")},
			"azuredevops_serviceendpoint_generic":     {Tok: makeResource("ServiceEndpointGeneric")},
			"azuredevops_serviceendpoint_generic_git": {Tok: makeResource("ServiceEndpointGenericGit")},
			"azuredevops_serviceendpoint_github_enterprise": {
				Tok: makeResource("ServiceEndpointGitHubEnterprise"),
			},
			"azuredevops_serviceendpoint_servicefabric": {
				Tok: makeResource("ServiceEndpointServiceFabric"),
			},
			"azuredevops_serviceendpoint_sonarcloud": {
				Tok: makeResource("ServiceEndpointSonarCloud"),
			},
			"azuredevops_iteration_permissions": { // iteration -> Iterative
				Tok: makeResource("IterativePermissions"),
			},

			"azuredevops_serviceendpoint_argocd":          {Tok: makeResource("ServiceendpointArgocd")},
			"azuredevops_serviceendpoint_permissions":     {Tok: makeResource("ServiceendpointPermissions")},
			"azuredevops_servicehook_permissions":         {Tok: makeResource("ServicehookPermissions")},
			"azuredevops_serviceendpoint_incomingwebhook": {Tok: makeResource("ServiceendpointIncomingwebhook")},
			"azuredevops_serviceendpoint_octopusdeploy":   {Tok: makeResource("ServiceendpointOctopusdeploy")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"azuredevops_serviceendpoint_azurerm": {Tok: makeDataSource("getServiceEndpointAzureRM")},
			"azuredevops_serviceendpoint_github":  {Tok: makeDataSource("getServiceEndpointGithub")},
			"azuredevops_team": {
				Docs: &tfbridge.DocInfo{
					Source: "data_team.html.markdown",
				},
			},
			"azuredevops_teams": {
				Docs: &tfbridge.DocInfo{
					Source: "data_teams.html.markdown",
				},
			},
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
			PyProject: struct{ Enabled bool }{true},
		},

		Golang: &tfbridge.GolangInfo{
			ImportBasePath: path.Join(
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

	makeResourceL := func(mod, name string) tokens.Type {
		return tfbridge.MakeResource(mainPkg, mod, name)
	}

	prov.RenameResourceWithAlias("azuredevops_resource_authorization",
		makeResourceL("Security", "ResourceAuthorization"),
		makeResource("ResourceAuthorization"), "Security", mainMod, nil)
	prov.RenameResourceWithAlias("azuredevops_branch_policy_build_validation",
		makeResourceL("Policy", "BranchPolicyBuildValidation"),
		makeResource("BranchPolicyBuildValidation"), "Policy", mainMod, nil)
	prov.RenameResourceWithAlias("azuredevops_branch_policy_min_reviewers",
		makeResourceL("Policy", "BranchPolicyMinReviewers"),
		makeResource("BranchPolicyMinReviewers"), "Policy", mainMod, nil)
	prov.RenameResourceWithAlias("azuredevops_build_definition",
		makeResourceL("Build", "BuildDefinition"),
		makeResource("BuildDefinition"), "Build", mainMod, nil)
	prov.RenameResourceWithAlias("azuredevops_project",
		makeResourceL("Core", "Project"),
		makeResource("Project"), "Core", mainMod, nil)
	prov.RenameResourceWithAlias("azuredevops_project_features",
		makeResourceL("Core", "ProjectFeatures"),
		makeResource("ProjectFeatures"), "Core", mainMod, nil)
	prov.RenameResourceWithAlias("azuredevops_variable_group",
		makeResourceL("Pipeline", "VariableGroup"),
		makeResource("VariableGroup"), "Pipeline", mainMod, nil)
	prov.RenameResourceWithAlias("azuredevops_serviceendpoint_azurerm",
		makeResourceL("ServiceEndpoint", "AzureRM"),
		makeResource("ServiceEndpointAzureRM"), "ServiceEndpoint", mainMod, nil)
	prov.RenameResourceWithAlias("azuredevops_serviceendpoint_bitbucket",
		makeResourceL("ServiceEndpoint", "BitBucket"),
		makeResource("ServiceEndpointBitBucket"), "ServiceEndpoint", mainMod, nil)
	prov.RenameResourceWithAlias("azuredevops_serviceendpoint_dockerregistry",
		makeResourceL("ServiceEndpoint", "DockerRegistry"),
		makeResource("ServiceEndpointDockerRegistry"), "ServiceEndpoint", mainMod, &tfbridge.ResourceInfo{
			Fields: map[string]*tfbridge.SchemaInfo{
				"docker_registry": {
					CSharpName: "DockerRegistryUrl",
				},
			},
		})
	prov.RenameResourceWithAlias("azuredevops_serviceendpoint_github",
		makeResourceL("ServiceEndpoint", "GitHub"),
		makeResource("ServiceEndpointGitHub"), "ServiceEndpoint", mainMod, nil)
	prov.RenameResourceWithAlias("azuredevops_serviceendpoint_kubernetes",
		makeResourceL("ServiceEndpoint", "Kubernetes"),
		makeResource("ServiceEndpointKubernetes"), "ServiceEndpoint", mainMod, nil)
	prov.RenameResourceWithAlias("azuredevops_git_repository",
		makeResourceL("Repository", "Git"),
		makeResource("Git"), "Repository", mainMod, nil)
	prov.RenameResourceWithAlias("azuredevops_user_entitlement",
		makeResourceL("Entitlement", "User"),
		makeResource("User"), "Entitlement", mainMod, nil)
	prov.RenameResourceWithAlias("azuredevops_group_membership",
		makeResourceL("Identities", "GroupMembership"),
		makeResource("GroupMembership"), "Identities", mainMod, nil)
	prov.RenameResourceWithAlias("azuredevops_agent_pool",
		makeResourceL("Agent", "Pool"),
		makeResource("Pool"), "Agent", mainMod, nil)
	prov.RenameResourceWithAlias("azuredevops_agent_queue",
		makeResourceL("Agent", "Queue"),
		makeResource("Queue"), "Agent", mainMod, nil)
	prov.RenameResourceWithAlias("azuredevops_group",
		makeResourceL("Identities", "Group"),
		makeResource("Group"), "Identities", mainMod, nil)

	makeDataSourceL := func(mod, name string) tokens.ModuleMember {
		return tfbridge.MakeDataSource(mainPkg, mod, name)
	}

	prov.RenameDataSource("azuredevops_agent_pool", makeDataSourceL("Agent", "getPool"),
		makeDataSource("getPool"), "Agent", mainMod, nil)
	prov.RenameDataSource("azuredevops_agent_pools", makeDataSourceL("Agent", "getPools"),
		makeDataSource("getPools"), "Agent", mainMod, nil)
	prov.RenameDataSource("azuredevops_client_config", makeDataSourceL("Core", "getClientConfig"),
		makeDataSource("getClientConfig"), "Core", mainMod, nil)
	prov.RenameDataSource("azuredevops_group", makeDataSourceL("Identities", "getGroup"),
		makeDataSource("getGroup"), "Identities", mainMod, nil)
	prov.RenameDataSource("azuredevops_project", makeDataSourceL("Core", "getProject"),
		makeDataSource("getProject"), "Core", mainMod, nil)
	prov.RenameDataSource("azuredevops_projects", makeDataSourceL("Core", "getProjects"),
		makeDataSource("getProjects"), "Core", mainMod, nil)
	prov.RenameDataSource("azuredevops_git_repositories",
		makeDataSourceL("Repository", "getRepositories"),
		makeDataSource("getRepositories"), "Repository", mainMod, nil)
	prov.RenameDataSource("azuredevops_users",
		makeDataSourceL("Identities", "getUsers"),
		makeDataSource("getUsers"), "Identities", mainMod, nil)

	prov.MustComputeTokens(tks.SingleModule("azuredevops_",
		mainMod, tks.MakeStandard(mainPkg)))

	prov.SetAutonaming(255, "-")

	prov.MustApplyAutoAliases()

	return prov
}
