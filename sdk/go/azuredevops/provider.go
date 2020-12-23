// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The provider type for the azuredevops package. By default, resources use package-wide configuration
// settings, however an explicit `Provider` instance may be created and passed during resource
// construction to achieve fine-grained programmatic control over provider settings. See the
// [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
type Provider struct {
	pulumi.ProviderResourceState
}

// NewProvider registers a new resource with the given unique name, arguments, and options.
func NewProvider(ctx *pulumi.Context,
	name string, args *ProviderArgs, opts ...pulumi.ResourceOption) (*Provider, error) {
	if args == nil {
		args = &ProviderArgs{}
	}

	if args.OrgServiceUrl == nil {
		args.OrgServiceUrl = pulumi.StringPtr(getEnvOrDefault("", nil, "AZDO_ORG_SERVICE_URL").(string))
	}
	if args.PersonalAccessToken == nil {
		args.PersonalAccessToken = pulumi.StringPtr(getEnvOrDefault("", nil, "AZDO_PERSONAL_ACCESS_TOKEN").(string))
	}
	var resource Provider
	err := ctx.RegisterResource("pulumi:providers:azuredevops", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type providerArgs struct {
	// The url of the Azure DevOps instance which should be used.
	OrgServiceUrl *string `pulumi:"orgServiceUrl"`
	// The personal access token which should be used.
	PersonalAccessToken *string `pulumi:"personalAccessToken"`
}

// The set of arguments for constructing a Provider resource.
type ProviderArgs struct {
	// The url of the Azure DevOps instance which should be used.
	OrgServiceUrl pulumi.StringPtrInput
	// The personal access token which should be used.
	PersonalAccessToken pulumi.StringPtrInput
}

func (ProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*providerArgs)(nil)).Elem()
}

type ProviderInput interface {
	pulumi.Input

	ToProviderOutput() ProviderOutput
	ToProviderOutputWithContext(ctx context.Context) ProviderOutput
}

func (Provider) ElementType() reflect.Type {
	return reflect.TypeOf((*Provider)(nil)).Elem()
}

func (i Provider) ToProviderOutput() ProviderOutput {
	return i.ToProviderOutputWithContext(context.Background())
}

func (i Provider) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderOutput)
}

type ProviderOutput struct {
	*pulumi.OutputState
}

func (ProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProviderOutput)(nil)).Elem()
}

func (o ProviderOutput) ToProviderOutput() ProviderOutput {
	return o
}

func (o ProviderOutput) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ProviderOutput{})
}
