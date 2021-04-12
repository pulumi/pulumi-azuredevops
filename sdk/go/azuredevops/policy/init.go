// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package policy

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-azuredevops/sdk/go/azuredevops"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "azuredevops:Policy/branchPolicyBuildValidation:BranchPolicyBuildValidation":
		r = &BranchPolicyBuildValidation{}
	case "azuredevops:Policy/branchPolicyMinReviewers:BranchPolicyMinReviewers":
		r = &BranchPolicyMinReviewers{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := azuredevops.PkgVersion()
	if err != nil {
		fmt.Println("failed to determine package version. defaulting to v1: %v", err)
	}
	pulumi.RegisterResourceModule(
		"azuredevops",
		"Policy/branchPolicyBuildValidation",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azuredevops",
		"Policy/branchPolicyMinReviewers",
		&module{version},
	)
}
