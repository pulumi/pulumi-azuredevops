package examples

import (
	"context"
	"strings"
	"testing"

	azdo "github.com/pulumi/pulumi-azuredevops/provider/v2"
	"github.com/pulumi/pulumi-azuredevops/provider/v2/pkg/version"
	testutils "github.com/pulumi/pulumi-terraform-bridge/testing/x"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
)

func init() {
	// This is necessary for gRPC testing. It doesn't effect integration tests, since
	// they use their own binary.
	version.Version = "6.0.0"
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
						"__defaults": [],
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
