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

	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/pulumi/pulumi-terraform-bridge/v2/pkg/tfbridge"
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v2/go/common/tokens"
	"github.com/terraform-providers/terraform-provider-azuredevops/azuredevops"
)

// all of the token components used below.
const (
	// packages:
	mainPkg = "azuredevops"

	// modules:
	//mainMod = "index"
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
func preConfigureCallback(vars resource.PropertyMap, c *terraform.ResourceConfig) error {
	return nil
}

// Provider returns additional overlaid schema and metadata associated with the provider.
func Provider() tfbridge.ProviderInfo {
	p := azuredevops.Provider()

	prov := tfbridge.ProviderInfo{
		P:           p,
		Name:        "azuredevops",
		Description: "A Pulumi package for creating and managing Azure DevOps.",
		Keywords:    []string{"pulumi", "azuredevops"},
		License:     "Apache-2.0",
		Homepage:    "https://pulumi.io",
		Repository:  "https://github.com/pulumi/pulumi-azuredevops",
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
			"azuredevops_resource_authorization": {
				Tok: makeResource("Security", "ResourceAuthorization"),
				Docs: &tfbridge.DocInfo{
					Source: "resource_authorization.html.markdown",
				},
			},
			"azuredevops_branch_policy_build_validation": {
				Tok: makeResource("Policy", "BranchPolicyBuildValidation"),
				Docs: &tfbridge.DocInfo{
					Source: "branch_policy_build_validation.html.markdown",
				},
			},
			"azuredevops_branch_policy_min_reviewers": {
				Tok: makeResource("Policy", "BranchPolicyMinReviewers"),
				Docs: &tfbridge.DocInfo{
					Source: "branch_policy_min_reviewers.html.markdown",
				},
			},
			"azuredevops_build_definition": {
				Tok: makeResource("Build", "BuildDefinition"),
				Docs: &tfbridge.DocInfo{
					Source: "build_definition.html.markdown",
				},
			},
			"azuredevops_project": {
				Tok: makeResource("Core", "Project"),
				Docs: &tfbridge.DocInfo{
					Source: "project.html.markdown",
				},
			},
			"azuredevops_project_features": {
				Tok: makeResource("Core", "ProjectFeatures"),
				Docs: &tfbridge.DocInfo{
					Source: "project_features.html.markdown",
				},
			},
			"azuredevops_variable_group": {
				Tok: makeResource("Pipeline", "VariableGroup"),
				Docs: &tfbridge.DocInfo{
					Source: "variable_group.html.markdown",
				},
			},
			"azuredevops_serviceendpoint_azurerm": {
				Tok: makeResource("ServiceEndpoint", "AzureRM"),
				Docs: &tfbridge.DocInfo{
					Source: "serviceendpoint_azurerm.html.markdown",
				},
			},
			"azuredevops_serviceendpoint_bitbucket": {
				Tok: makeResource("ServiceEndpoint", "BitBucket"),
				Docs: &tfbridge.DocInfo{
					Source: "serviceendpoint_bitbucket.html.markdown",
				},
			},
			"azuredevops_serviceendpoint_dockerregistry": {
				Tok: makeResource("ServiceEndpoint", "DockerRegistry"),
				Docs: &tfbridge.DocInfo{
					Source: "serviceendpoint_dockerregistry.html.markdown",
				},
				Fields: map[string]*tfbridge.SchemaInfo{
					"docker_registry": {
						CSharpName: "DockerRegistryUrl",
					},
				},
			},
			"azuredevops_serviceendpoint_github": {
				Tok: makeResource("ServiceEndpoint", "GitHub"),
				Docs: &tfbridge.DocInfo{
					Source: "serviceendpoint_github.html.markdown",
				},
			},
			"azuredevops_serviceendpoint_kubernetes": {
				Tok: makeResource("ServiceEndpoint", "Kubernetes"),
				Docs: &tfbridge.DocInfo{
					Source: "serviceendpoint_kubernetes.html.markdown",
				},
			},
			"azuredevops_git_repository": {
				Tok: makeResource("Repository", "Git"),
				Docs: &tfbridge.DocInfo{
					Source: "azure_git_repository.html.markdown",
				},
			},
			"azuredevops_user_entitlement": {
				Tok: makeResource("Entitlement", "User"),
				Docs: &tfbridge.DocInfo{
					Source: "user_entitlement.html.markdown",
				},
			},
			"azuredevops_group_membership": {
				Tok: makeResource("Identities", "GroupMembership"),
				Docs: &tfbridge.DocInfo{
					Source: "group_membership.html.markdown",
				},
			},
			"azuredevops_agent_pool": {
				Tok: makeResource("Agent", "Pool"),
				Docs: &tfbridge.DocInfo{
					Source: "agent_pool.html.markdown",
				},
			},
			"azuredevops_agent_queue": {
				Tok: makeResource("Agent", "Queue"),
				Docs: &tfbridge.DocInfo{
					Source: "agent_queue.html.markdown",
				},
			},
			"azuredevops_group": {
				Tok: makeResource("Identities", "Group"),
				Docs: &tfbridge.DocInfo{
					Source: "group.html.markdown",
				},
			},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"azuredevops_agent_pool": {
				Tok: makeDataSource("Agent", "getPool"),
				Docs: &tfbridge.DocInfo{
					Source: "data_agent_pool.html.markdown",
				},
			},
			"azuredevops_agent_pools": {
				Tok: makeDataSource("Agent", "getPools"),
				Docs: &tfbridge.DocInfo{
					Source: "data_agent_pools.html.markdown",
				},
			},
			"azuredevops_client_config": {
				Tok: makeDataSource("Core", "getClientConfig"),
				Docs: &tfbridge.DocInfo{
					Source: "data_client_config.html.markdown",
				},
			},
			"azuredevops_group": {
				Tok: makeDataSource("Identities", "getGroup"),
				Docs: &tfbridge.DocInfo{
					Source: "data_group.html.markdown",
				},
			},
			"azuredevops_project": {
				Tok: makeDataSource("Core", "getProject"),
				Docs: &tfbridge.DocInfo{
					Source: "data_project.html.markdown",
				},
			},
			"azuredevops_projects": {
				Tok: makeDataSource("Core", "getProjects"),
				Docs: &tfbridge.DocInfo{
					Source: "data_projects.html.markdown",
				},
			},
			"azuredevops_git_repositories": {
				Tok: makeDataSource("Repository", "getRepositories"),
				Docs: &tfbridge.DocInfo{
					Source: "data_git_repositories.html.markdown",
				},
			},
			"azuredevops_users": {
				Tok: makeDataSource("Identities", "getUsers"),
				Docs: &tfbridge.DocInfo{
					Source: "data_users.html.markdown",
				},
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
				"pulumi": ">=2.0.0,<3.0.0",
			},
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

	// For all resources with name properties, we will add an auto-name property.  Make sure to skip those that
	// already have a name mapping entry, since those may have custom overrides set above (e.g., for length).
	const nameProperty = "name"
	for resname, res := range prov.Resources {
		if resourceSchema := p.ResourcesMap[resname]; resourceSchema != nil {
			// Only apply auto-name to input properties (Optional || Required) named `name`
			if tfs, has := resourceSchema.Schema[nameProperty]; has && (tfs.Optional || tfs.Required) {
				if _, hasfield := res.Fields[nameProperty]; !hasfield {
					if res.Fields == nil {
						res.Fields = make(map[string]*tfbridge.SchemaInfo)
					}
					res.Fields[nameProperty] = tfbridge.AutoName(nameProperty, 255)
				}
			}
		}
	}

	return prov
}
