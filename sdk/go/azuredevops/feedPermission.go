// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azuredevops/sdk/v3/go/azuredevops/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages creation of the Feed Permission within Azure DevOps organization.
//
// ## Example Usage
//
// ### Create Feed Permission
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-azuredevops/sdk/v3/go/azuredevops"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			example, err := azuredevops.NewProject(ctx, "example", &azuredevops.ProjectArgs{
//				Name: pulumi.String("Example Project"),
//			})
//			if err != nil {
//				return err
//			}
//			exampleGroup, err := azuredevops.NewGroup(ctx, "example", &azuredevops.GroupArgs{
//				Scope:       example.ID(),
//				DisplayName: pulumi.String("Example group"),
//				Description: pulumi.String("Example description"),
//			})
//			if err != nil {
//				return err
//			}
//			exampleFeed, err := azuredevops.NewFeed(ctx, "example", &azuredevops.FeedArgs{
//				Name: pulumi.String("examplefeed"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = azuredevops.NewFeedPermission(ctx, "permission", &azuredevops.FeedPermissionArgs{
//				FeedId:             exampleFeed.ID(),
//				Role:               pulumi.String("reader"),
//				IdentityDescriptor: exampleGroup.Descriptor,
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Relevant Links
//
// - [Azure DevOps Service REST API 7.0 - Feed Management](https://learn.microsoft.com/en-us/rest/api/azure/devops/artifacts/feed-management?view=azure-devops-rest-7.0)
type FeedPermission struct {
	pulumi.CustomResourceState

	// The display name of the assignment
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// The ID of the Feed.
	FeedId pulumi.StringOutput `pulumi:"feedId"`
	// The Descriptor of identity you want to assign a role.
	IdentityDescriptor pulumi.StringOutput `pulumi:"identityDescriptor"`
	// The ID of the identity.
	IdentityId pulumi.StringOutput `pulumi:"identityId"`
	// The ID of the Project Feed is created in. If not specified, feed will be created at the organization level.
	ProjectId pulumi.StringPtrOutput `pulumi:"projectId"`
	// The role to be assigned. Possible values are: `reader`, `contributor`, `collaborator`, `administrator`
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewFeedPermission registers a new resource with the given unique name, arguments, and options.
func NewFeedPermission(ctx *pulumi.Context,
	name string, args *FeedPermissionArgs, opts ...pulumi.ResourceOption) (*FeedPermission, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FeedId == nil {
		return nil, errors.New("invalid value for required argument 'FeedId'")
	}
	if args.IdentityDescriptor == nil {
		return nil, errors.New("invalid value for required argument 'IdentityDescriptor'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource FeedPermission
	err := ctx.RegisterResource("azuredevops:index/feedPermission:FeedPermission", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFeedPermission gets an existing FeedPermission resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFeedPermission(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FeedPermissionState, opts ...pulumi.ResourceOption) (*FeedPermission, error) {
	var resource FeedPermission
	err := ctx.ReadResource("azuredevops:index/feedPermission:FeedPermission", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FeedPermission resources.
type feedPermissionState struct {
	// The display name of the assignment
	DisplayName *string `pulumi:"displayName"`
	// The ID of the Feed.
	FeedId *string `pulumi:"feedId"`
	// The Descriptor of identity you want to assign a role.
	IdentityDescriptor *string `pulumi:"identityDescriptor"`
	// The ID of the identity.
	IdentityId *string `pulumi:"identityId"`
	// The ID of the Project Feed is created in. If not specified, feed will be created at the organization level.
	ProjectId *string `pulumi:"projectId"`
	// The role to be assigned. Possible values are: `reader`, `contributor`, `collaborator`, `administrator`
	Role *string `pulumi:"role"`
}

type FeedPermissionState struct {
	// The display name of the assignment
	DisplayName pulumi.StringPtrInput
	// The ID of the Feed.
	FeedId pulumi.StringPtrInput
	// The Descriptor of identity you want to assign a role.
	IdentityDescriptor pulumi.StringPtrInput
	// The ID of the identity.
	IdentityId pulumi.StringPtrInput
	// The ID of the Project Feed is created in. If not specified, feed will be created at the organization level.
	ProjectId pulumi.StringPtrInput
	// The role to be assigned. Possible values are: `reader`, `contributor`, `collaborator`, `administrator`
	Role pulumi.StringPtrInput
}

func (FeedPermissionState) ElementType() reflect.Type {
	return reflect.TypeOf((*feedPermissionState)(nil)).Elem()
}

type feedPermissionArgs struct {
	// The display name of the assignment
	DisplayName *string `pulumi:"displayName"`
	// The ID of the Feed.
	FeedId string `pulumi:"feedId"`
	// The Descriptor of identity you want to assign a role.
	IdentityDescriptor string `pulumi:"identityDescriptor"`
	// The ID of the Project Feed is created in. If not specified, feed will be created at the organization level.
	ProjectId *string `pulumi:"projectId"`
	// The role to be assigned. Possible values are: `reader`, `contributor`, `collaborator`, `administrator`
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a FeedPermission resource.
type FeedPermissionArgs struct {
	// The display name of the assignment
	DisplayName pulumi.StringPtrInput
	// The ID of the Feed.
	FeedId pulumi.StringInput
	// The Descriptor of identity you want to assign a role.
	IdentityDescriptor pulumi.StringInput
	// The ID of the Project Feed is created in. If not specified, feed will be created at the organization level.
	ProjectId pulumi.StringPtrInput
	// The role to be assigned. Possible values are: `reader`, `contributor`, `collaborator`, `administrator`
	Role pulumi.StringInput
}

func (FeedPermissionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*feedPermissionArgs)(nil)).Elem()
}

type FeedPermissionInput interface {
	pulumi.Input

	ToFeedPermissionOutput() FeedPermissionOutput
	ToFeedPermissionOutputWithContext(ctx context.Context) FeedPermissionOutput
}

func (*FeedPermission) ElementType() reflect.Type {
	return reflect.TypeOf((**FeedPermission)(nil)).Elem()
}

func (i *FeedPermission) ToFeedPermissionOutput() FeedPermissionOutput {
	return i.ToFeedPermissionOutputWithContext(context.Background())
}

func (i *FeedPermission) ToFeedPermissionOutputWithContext(ctx context.Context) FeedPermissionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FeedPermissionOutput)
}

// FeedPermissionArrayInput is an input type that accepts FeedPermissionArray and FeedPermissionArrayOutput values.
// You can construct a concrete instance of `FeedPermissionArrayInput` via:
//
//	FeedPermissionArray{ FeedPermissionArgs{...} }
type FeedPermissionArrayInput interface {
	pulumi.Input

	ToFeedPermissionArrayOutput() FeedPermissionArrayOutput
	ToFeedPermissionArrayOutputWithContext(context.Context) FeedPermissionArrayOutput
}

type FeedPermissionArray []FeedPermissionInput

func (FeedPermissionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FeedPermission)(nil)).Elem()
}

func (i FeedPermissionArray) ToFeedPermissionArrayOutput() FeedPermissionArrayOutput {
	return i.ToFeedPermissionArrayOutputWithContext(context.Background())
}

func (i FeedPermissionArray) ToFeedPermissionArrayOutputWithContext(ctx context.Context) FeedPermissionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FeedPermissionArrayOutput)
}

// FeedPermissionMapInput is an input type that accepts FeedPermissionMap and FeedPermissionMapOutput values.
// You can construct a concrete instance of `FeedPermissionMapInput` via:
//
//	FeedPermissionMap{ "key": FeedPermissionArgs{...} }
type FeedPermissionMapInput interface {
	pulumi.Input

	ToFeedPermissionMapOutput() FeedPermissionMapOutput
	ToFeedPermissionMapOutputWithContext(context.Context) FeedPermissionMapOutput
}

type FeedPermissionMap map[string]FeedPermissionInput

func (FeedPermissionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FeedPermission)(nil)).Elem()
}

func (i FeedPermissionMap) ToFeedPermissionMapOutput() FeedPermissionMapOutput {
	return i.ToFeedPermissionMapOutputWithContext(context.Background())
}

func (i FeedPermissionMap) ToFeedPermissionMapOutputWithContext(ctx context.Context) FeedPermissionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FeedPermissionMapOutput)
}

type FeedPermissionOutput struct{ *pulumi.OutputState }

func (FeedPermissionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FeedPermission)(nil)).Elem()
}

func (o FeedPermissionOutput) ToFeedPermissionOutput() FeedPermissionOutput {
	return o
}

func (o FeedPermissionOutput) ToFeedPermissionOutputWithContext(ctx context.Context) FeedPermissionOutput {
	return o
}

// The display name of the assignment
func (o FeedPermissionOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FeedPermission) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// The ID of the Feed.
func (o FeedPermissionOutput) FeedId() pulumi.StringOutput {
	return o.ApplyT(func(v *FeedPermission) pulumi.StringOutput { return v.FeedId }).(pulumi.StringOutput)
}

// The Descriptor of identity you want to assign a role.
func (o FeedPermissionOutput) IdentityDescriptor() pulumi.StringOutput {
	return o.ApplyT(func(v *FeedPermission) pulumi.StringOutput { return v.IdentityDescriptor }).(pulumi.StringOutput)
}

// The ID of the identity.
func (o FeedPermissionOutput) IdentityId() pulumi.StringOutput {
	return o.ApplyT(func(v *FeedPermission) pulumi.StringOutput { return v.IdentityId }).(pulumi.StringOutput)
}

// The ID of the Project Feed is created in. If not specified, feed will be created at the organization level.
func (o FeedPermissionOutput) ProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FeedPermission) pulumi.StringPtrOutput { return v.ProjectId }).(pulumi.StringPtrOutput)
}

// The role to be assigned. Possible values are: `reader`, `contributor`, `collaborator`, `administrator`
func (o FeedPermissionOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *FeedPermission) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

type FeedPermissionArrayOutput struct{ *pulumi.OutputState }

func (FeedPermissionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FeedPermission)(nil)).Elem()
}

func (o FeedPermissionArrayOutput) ToFeedPermissionArrayOutput() FeedPermissionArrayOutput {
	return o
}

func (o FeedPermissionArrayOutput) ToFeedPermissionArrayOutputWithContext(ctx context.Context) FeedPermissionArrayOutput {
	return o
}

func (o FeedPermissionArrayOutput) Index(i pulumi.IntInput) FeedPermissionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FeedPermission {
		return vs[0].([]*FeedPermission)[vs[1].(int)]
	}).(FeedPermissionOutput)
}

type FeedPermissionMapOutput struct{ *pulumi.OutputState }

func (FeedPermissionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FeedPermission)(nil)).Elem()
}

func (o FeedPermissionMapOutput) ToFeedPermissionMapOutput() FeedPermissionMapOutput {
	return o
}

func (o FeedPermissionMapOutput) ToFeedPermissionMapOutputWithContext(ctx context.Context) FeedPermissionMapOutput {
	return o
}

func (o FeedPermissionMapOutput) MapIndex(k pulumi.StringInput) FeedPermissionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FeedPermission {
		return vs[0].(map[string]*FeedPermission)[vs[1].(string)]
	}).(FeedPermissionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FeedPermissionInput)(nil)).Elem(), &FeedPermission{})
	pulumi.RegisterInputType(reflect.TypeOf((*FeedPermissionArrayInput)(nil)).Elem(), FeedPermissionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FeedPermissionMapInput)(nil)).Elem(), FeedPermissionMap{})
	pulumi.RegisterOutputType(FeedPermissionOutput{})
	pulumi.RegisterOutputType(FeedPermissionArrayOutput{})
	pulumi.RegisterOutputType(FeedPermissionMapOutput{})
}
