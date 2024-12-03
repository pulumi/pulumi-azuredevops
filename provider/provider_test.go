package azuredevops

import (
	"context"
	"path"
	"testing"

	_ "embed"

	"github.com/pulumi/providertest/providers"
	"github.com/pulumi/providertest/pulumitest"
	"github.com/pulumi/providertest/pulumitest/opttest"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	pulumirpc "github.com/pulumi/pulumi/sdk/v3/proto/go"

	"github.com/pulumi/pulumi-azuredevops/provider/v3/pkg/version"
)

func TestVariableGroupValidation(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}
	pt := pulumitest.NewPulumiTest(t, path.Join("test-programs", "variable-group"),
		opttest.AttachProviderServer("azuredevops", newTestProvider))

	// This fails if running the upstream validation via the bridge
	pt.Up(t)
}

// Use the non-embedded schema to avoid having to run generation before running the tests.
//
//go:embed cmd/pulumi-resource-azuredevops/schema.json
var schemaBytes []byte

func newTestProvider(_ providers.PulumiTest) (pulumirpc.ResourceProviderServer, error) {
	version.Version = "0.0.1"
	providerInfo := Provider()

	return tfbridge.NewProvider(context.Background(), nil, "azuredevops",
		version.Version, providerInfo.P, providerInfo, schemaBytes), nil
}
