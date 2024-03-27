package examples

import (
	"context"
	"os"
	"strings"
	"testing"

	azdo "github.com/pulumi/pulumi-azuredevops/provider/v3"
	"github.com/pulumi/pulumi-azuredevops/provider/v3/pkg/version"
	testutils "github.com/pulumi/pulumi-terraform-bridge/testing/x"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

func init() {
	// This is necessary for gRPC testing. It doesn't effect integration tests, since
	// they use their own binary.
	version.Version = "6.0.0"
}

func getCwd(t *testing.T) string {
	cwd, err := os.Getwd()
	if err != nil {
		t.FailNow()
	}
	return cwd
}

func getBaseOptions(t *testing.T) integration.ProgramTestOptions {
	return integration.ProgramTestOptions{
		ExpectRefreshChanges: true,
	}
}

func replay(t *testing.T, sequence string) {
	info := azdo.Provider()
	ctx := context.Background()
	p := tfbridge.NewProvider(ctx, nil, info.Name, info.Version, info.P, info, []byte("{}"))
	testutils.ReplaySequence(t, p, sequence)
}

// This test checks that the problem fixed in 0001-remove_git_repo_default.patch
// does not fail the check for repo.
// https://github.com/pulumi/pulumi-azuredevops/issues/226
func TestServiceConnectionIdDefaultDoesNotConflict(t *testing.T) {
	repro := strings.ReplaceAll(`
	[
		{
			"method": "/pulumirpc.ResourceProvider/Check",
			"request": {
				"urn": "urn:pulumi:dev::azure_devops_prog::azuredevops:index/git:Git::repo",
				"olds": {},
				"news": {
					"initialization": {
						"initType": "Clean"
					},
					"projectId": "04da6b54-80e4-46f7-96ec-b56ff0331ba9"
				},
				"randomSeed": "Am+8wrqVx8+eCu9XSun5rpfEjbC1iNHy27lOskORu3s="
			},
			"response": {
				"inputs": {
					"__defaults": [
						"name"
					],
					"initialization": {
						"__defaults": [
							"serviceConnectionId"
						],
						"serviceConnectionId": "",
						"initType": "Clean"
					},
					"name": "repo-779583a",
					"projectId": "04da6b54-80e4-46f7-96ec-b56ff0331ba9"
				}
			},
			"metadata": {
				"kind": "resource",
				"mode": "client",
				"name": "azuredevops"
			}
		}
  ]
	`, "$", "`")
	replay(t, repro)
}

// This test checks that the problem fixed in 0002-remove_min_reviwers_defaults.patch
// does not fail the check for repo.
// https://github.com/pulumi/pulumi-azuredevops/issues/227
func TestMinReviewersDefaultDoesNotConflict(t *testing.T) {
	repro := strings.ReplaceAll(`
	[
		{
			"method": "/pulumirpc.ResourceProvider/Check",
			"request": {
				"urn": "urn:pulumi:dev::azure_devops_prog::azuredevops:index/branchPolicyMinReviewers:BranchPolicyMinReviewers::exampleBranchPolicyMinReviewers",
				"olds": {},
				"news": {
					"projectId": "75e12870-434c-455b-b8ac-c0e17cc51c46",
					"settings": {
						"onPushResetAllVotes": true,
						"reviewerCount": 7,
						"scopes": [
							{
								"matchType": "Exact",
								"repositoryId": "54745441-9448-43f9-a5cb-2c3644bbac54",
								"repositoryRef": "refs/heads/master"
							}
						]
					}
				},
				"randomSeed": "TxXqw2+y2hjh5dvv//yChyw2PYzaCeOdsx8RoYlqDBs="
			},
			"response": {
				"inputs": {
					"__defaults": [
						"blocking",
						"enabled"
					],
					"blocking": true,
					"enabled": true,
					"projectId": "75e12870-434c-455b-b8ac-c0e17cc51c46",
					"settings": {
						"__defaults": [
							"allowCompletionWithRejectsOrWaits",
							"lastPusherCannotApprove",
							"onLastIterationRequireVote",
							"onPushResetApprovedVotes",
							"submitterCanVote"
						],
						"allowCompletionWithRejectsOrWaits": false,
						"lastPusherCannotApprove": false,
						"onLastIterationRequireVote": false,
						"onPushResetAllVotes": true,
						"onPushResetApprovedVotes": false,
						"reviewerCount": 7,
						"scopes": [
							{
								"__defaults": [],
								"matchType": "Exact",
								"repositoryId": "54745441-9448-43f9-a5cb-2c3644bbac54",
								"repositoryRef": "refs/heads/master"
							}
						],
						"submitterCanVote": false
					}
				}
			},
			"metadata": {
				"kind": "resource",
				"mode": "client",
				"name": "azuredevops"
			}
		}
  ]
	`, "$", "`")
	replay(t, repro)
}
