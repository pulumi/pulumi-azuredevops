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

	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	tks "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge/tokens"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfgen"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"

	"github.com/pulumi/pulumi-azuredevops/provider/v3/pkg/version"
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
		P:            p,
		Name:         "azuredevops",
		Description:  "A Pulumi package for creating and managing Azure DevOps.",
		Keywords:     []string{"pulumi", "azuredevops"},
		License:      "Apache-2.0",
		Homepage:     "https://pulumi.io",
		Repository:   "https://github.com/pulumi/pulumi-azuredevops",
		GitHubOrg:    "microsoft",
		Version:      version.Version,
		MetadataInfo: tfbridge.NewProviderMetadata(metadata),
		DocRules:     &tfbridge.DocRuleInfo{EditRules: docEditRules},
		Config: map[string]*tfbridge.SchemaInfo{
			"org_service_url": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"AZDO_ORG_SERVICE_URL"},
				},
			},
		},
		Resources: map[string]*tfbridge.ResourceInfo{
			"azuredevops_agent_pool":                  {Tok: makeResource("Pool")},
			"azuredevops_agent_queue":                 {Tok: makeResource("Queue")},
			"azuredevops_branch_policy_min_reviewers": {Tok: makeResource("BranchPolicyMinReviewers")},
			"azuredevops_build_definition":            {Tok: makeResource("BuildDefinition")},
			"azuredevops_environment_resource_kubernetes": {
				Docs: &tfbridge.DocInfo{
					Source: "kubernetes_resource.html.markdown",
				},
			},
			"azuredevops_git_repository":   {Tok: makeResource("Git")},
			"azuredevops_group":            {Tok: makeResource("Group")},
			"azuredevops_group_membership": {Tok: makeResource("GroupMembership")},
			"azuredevops_iteration_permissions": {
				// iteration -> Iterative
				Tok: makeResource("IterativePermissions"),
			},
			"azuredevops_project":                           {Tok: makeResource("Project")},
			"azuredevops_project_features":                  {Tok: makeResource("ProjectFeatures")},
			"azuredevops_serviceendpoint_argocd":            {Tok: makeResource("ServiceendpointArgocd")},
			"azuredevops_serviceendpoint_artifactory":       {Tok: makeResource("ServiceEndpointArtifactory")},
			"azuredevops_serviceendpoint_aws":               {Tok: makeResource("ServiceEndpointAws")},
			"azuredevops_serviceendpoint_azurecr":           {Tok: makeResource("ServiceEndpointAzureEcr")},
			"azuredevops_serviceendpoint_azuredevops":       {Tok: makeResource("ServiceEndpointAzureDevOps")},
			"azuredevops_serviceendpoint_azurerm":           {Tok: makeResource("ServiceEndpointAzureRM")},
			"azuredevops_serviceendpoint_bitbucket":         {Tok: makeResource("ServiceEndpointBitBucket")},
			"azuredevops_serviceendpoint_generic":           {Tok: makeResource("ServiceEndpointGeneric")},
			"azuredevops_serviceendpoint_generic_git":       {Tok: makeResource("ServiceEndpointGenericGit")},
			"azuredevops_serviceendpoint_github":            {Tok: makeResource("ServiceEndpointGitHub")},
			"azuredevops_serviceendpoint_github_enterprise": {Tok: makeResource("ServiceEndpointGitHubEnterprise")},
			"azuredevops_serviceendpoint_incomingwebhook":   {Tok: makeResource("ServiceendpointIncomingwebhook")},
			"azuredevops_serviceendpoint_kubernetes":        {Tok: makeResource("ServiceEndpointKubernetes")},
			"azuredevops_serviceendpoint_npm":               {Tok: makeResource("ServiceEndpointNpm")},
			"azuredevops_serviceendpoint_octopusdeploy":     {Tok: makeResource("ServiceendpointOctopusdeploy")},
			"azuredevops_serviceendpoint_permissions":       {Tok: makeResource("ServiceendpointPermissions")},
			"azuredevops_serviceendpoint_runpipeline":       {Tok: makeResource("ServiceEndpointPipeline")},
			"azuredevops_serviceendpoint_servicefabric":     {Tok: makeResource("ServiceEndpointServiceFabric")},
			"azuredevops_serviceendpoint_sonarcloud":        {Tok: makeResource("ServiceEndpointSonarCloud")},
			"azuredevops_serviceendpoint_sonarqube":         {Tok: makeResource("ServiceEndpointSonarQube")},
			"azuredevops_serviceendpoint_ssh":               {Tok: makeResource("ServiceEndpointSsh")},
			"azuredevops_servicehook_permissions":           {Tok: makeResource("ServicehookPermissions")},
			"azuredevops_user_entitlement":                  {Tok: makeResource("User")},
			"azuredevops_variable_group":                    {Tok: makeResource("VariableGroup")},
			"azuredevops_workitemquery_permissions":         {Tok: makeResource("WorkItemQueryPermissions")},

			"azuredevops_serviceendpoint_dockerregistry": {
				Tok: makeResource("ServiceEndpointDockerRegistry"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"docker_registry": {
						CSharpName: "DockerRegistryUrl",
					},
				},
			},
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

			"azuredevops_agent_pool":       {Tok: makeDataSource("getPool")},
			"azuredevops_agent_pools":      {Tok: makeDataSource("getPools")},
			"azuredevops_client_config":    {Tok: makeDataSource("getClientConfig")},
			"azuredevops_group":            {Tok: makeDataSource("getGroup")},
			"azuredevops_project":          {Tok: makeDataSource("getProject")},
			"azuredevops_projects":         {Tok: makeDataSource("getProjects")},
			"azuredevops_git_repositories": {Tok: makeDataSource("getRepositories")},
			"azuredevops_users":            {Tok: makeDataSource("getUsers")},
			"azuredevops_identity_user": {
				// Pluralizing because the data source returns multiple users and the docs file is pluralized.
				Tok: makeDataSource("getIdentityUsers"),
				Docs: &tfbridge.DocInfo{
					Source: "identity_users.html.markdown",
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
			RespectSchemaVersion: true,
		},
		Python: &tfbridge.PythonInfo{
			RespectSchemaVersion: true,
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
			PyProject:  struct{ Enabled bool }{true},
			InputTypes: tfbridge.PythonInputTypeClassesAndDicts,
		},

		Golang: &tfbridge.GolangInfo{
			ImportBasePath: path.Join(
				fmt.Sprintf("github.com/pulumi/pulumi-%[1]s/sdk/", mainPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
			RespectSchemaVersion:           true,
		},
		CSharp: &tfbridge.CSharpInfo{
			RespectSchemaVersion: true,
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
			Namespaces: map[string]string{
				mainPkg: "AzureDevOps",
			},
		},
	}

	prov.MustComputeTokens(tks.SingleModule("azuredevops_",
		mainMod, tks.MakeStandard(mainPkg)))

	prov.SetAutonaming(255, "-")

	prov.MustApplyAutoAliases()

	return prov
}

func docEditRules(defaults []tfbridge.DocsEdit) []tfbridge.DocsEdit {
	return append(
		defaults,
		skipAuthenticationSection,
	)
}

// The Authentication section contains links to internal guides.
var skipAuthenticationSection = tfbridge.DocsEdit{
	Path: "index.md",
	Edit: func(_ string, content []byte) ([]byte, error) {
		return tfgen.SkipSectionByHeaderContent(content, func(headerText string) bool {
			return headerText == "Authentication"
		})
	},
}
