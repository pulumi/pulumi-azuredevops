// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azuredevops/sdk/v3/go/azuredevops/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages Feed within Azure DevOps organization.
//
// ## Example Usage
//
// ### Create Feed in the scope of whole Organization
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
//			_, err := azuredevops.NewFeed(ctx, "example", &azuredevops.FeedArgs{
//				Name: pulumi.String("examplefeed"),
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
// ### Create Feed in the scope of a Project
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
//				Name:             pulumi.String("Example Project"),
//				Visibility:       pulumi.String("private"),
//				VersionControl:   pulumi.String("Git"),
//				WorkItemTemplate: pulumi.String("Agile"),
//				Description:      pulumi.String("Managed by Pulumi"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = azuredevops.NewFeed(ctx, "example", &azuredevops.FeedArgs{
//				Name:      pulumi.String("examplefeed"),
//				ProjectId: example.ID(),
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
// ### Create Feed with Soft Delete
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
//			_, err := azuredevops.NewFeed(ctx, "example", &azuredevops.FeedArgs{
//				Name: pulumi.String("examplefeed"),
//				Features: azuredevops.FeedFeatureArray{
//					&azuredevops.FeedFeatureArgs{
//						PermanentDelete: pulumi.Bool(false),
//					},
//				},
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
//
// ## Import
//
// Azure DevOps Feed can be imported using the Project ID and Feed ID or Feed ID e.g.:
//
// ```sh
// $ pulumi import azuredevops:index/feed:Feed example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
// ```
//
// or
//
// ```sh
// $ pulumi import azuredevops:index/feed:Feed example 00000000-0000-0000-0000-000000000000
// ```
type Feed struct {
	pulumi.CustomResourceState

	// A `features` blocks as documented below.
	//
	// > **Note** *Because of ADO limitations feed name can be **reserved** for up to 15 minutes after permanent delete of the feed*
	Features FeedFeatureArrayOutput `pulumi:"features"`
	// The name of the Feed.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the Project Feed is created in. If not specified, feed will be created at the organization level.
	ProjectId pulumi.StringPtrOutput `pulumi:"projectId"`
}

// NewFeed registers a new resource with the given unique name, arguments, and options.
func NewFeed(ctx *pulumi.Context,
	name string, args *FeedArgs, opts ...pulumi.ResourceOption) (*Feed, error) {
	if args == nil {
		args = &FeedArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Feed
	err := ctx.RegisterResource("azuredevops:index/feed:Feed", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFeed gets an existing Feed resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFeed(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FeedState, opts ...pulumi.ResourceOption) (*Feed, error) {
	var resource Feed
	err := ctx.ReadResource("azuredevops:index/feed:Feed", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Feed resources.
type feedState struct {
	// A `features` blocks as documented below.
	//
	// > **Note** *Because of ADO limitations feed name can be **reserved** for up to 15 minutes after permanent delete of the feed*
	Features []FeedFeature `pulumi:"features"`
	// The name of the Feed.
	Name *string `pulumi:"name"`
	// The ID of the Project Feed is created in. If not specified, feed will be created at the organization level.
	ProjectId *string `pulumi:"projectId"`
}

type FeedState struct {
	// A `features` blocks as documented below.
	//
	// > **Note** *Because of ADO limitations feed name can be **reserved** for up to 15 minutes after permanent delete of the feed*
	Features FeedFeatureArrayInput
	// The name of the Feed.
	Name pulumi.StringPtrInput
	// The ID of the Project Feed is created in. If not specified, feed will be created at the organization level.
	ProjectId pulumi.StringPtrInput
}

func (FeedState) ElementType() reflect.Type {
	return reflect.TypeOf((*feedState)(nil)).Elem()
}

type feedArgs struct {
	// A `features` blocks as documented below.
	//
	// > **Note** *Because of ADO limitations feed name can be **reserved** for up to 15 minutes after permanent delete of the feed*
	Features []FeedFeature `pulumi:"features"`
	// The name of the Feed.
	Name *string `pulumi:"name"`
	// The ID of the Project Feed is created in. If not specified, feed will be created at the organization level.
	ProjectId *string `pulumi:"projectId"`
}

// The set of arguments for constructing a Feed resource.
type FeedArgs struct {
	// A `features` blocks as documented below.
	//
	// > **Note** *Because of ADO limitations feed name can be **reserved** for up to 15 minutes after permanent delete of the feed*
	Features FeedFeatureArrayInput
	// The name of the Feed.
	Name pulumi.StringPtrInput
	// The ID of the Project Feed is created in. If not specified, feed will be created at the organization level.
	ProjectId pulumi.StringPtrInput
}

func (FeedArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*feedArgs)(nil)).Elem()
}

type FeedInput interface {
	pulumi.Input

	ToFeedOutput() FeedOutput
	ToFeedOutputWithContext(ctx context.Context) FeedOutput
}

func (*Feed) ElementType() reflect.Type {
	return reflect.TypeOf((**Feed)(nil)).Elem()
}

func (i *Feed) ToFeedOutput() FeedOutput {
	return i.ToFeedOutputWithContext(context.Background())
}

func (i *Feed) ToFeedOutputWithContext(ctx context.Context) FeedOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FeedOutput)
}

// FeedArrayInput is an input type that accepts FeedArray and FeedArrayOutput values.
// You can construct a concrete instance of `FeedArrayInput` via:
//
//	FeedArray{ FeedArgs{...} }
type FeedArrayInput interface {
	pulumi.Input

	ToFeedArrayOutput() FeedArrayOutput
	ToFeedArrayOutputWithContext(context.Context) FeedArrayOutput
}

type FeedArray []FeedInput

func (FeedArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Feed)(nil)).Elem()
}

func (i FeedArray) ToFeedArrayOutput() FeedArrayOutput {
	return i.ToFeedArrayOutputWithContext(context.Background())
}

func (i FeedArray) ToFeedArrayOutputWithContext(ctx context.Context) FeedArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FeedArrayOutput)
}

// FeedMapInput is an input type that accepts FeedMap and FeedMapOutput values.
// You can construct a concrete instance of `FeedMapInput` via:
//
//	FeedMap{ "key": FeedArgs{...} }
type FeedMapInput interface {
	pulumi.Input

	ToFeedMapOutput() FeedMapOutput
	ToFeedMapOutputWithContext(context.Context) FeedMapOutput
}

type FeedMap map[string]FeedInput

func (FeedMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Feed)(nil)).Elem()
}

func (i FeedMap) ToFeedMapOutput() FeedMapOutput {
	return i.ToFeedMapOutputWithContext(context.Background())
}

func (i FeedMap) ToFeedMapOutputWithContext(ctx context.Context) FeedMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FeedMapOutput)
}

type FeedOutput struct{ *pulumi.OutputState }

func (FeedOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Feed)(nil)).Elem()
}

func (o FeedOutput) ToFeedOutput() FeedOutput {
	return o
}

func (o FeedOutput) ToFeedOutputWithContext(ctx context.Context) FeedOutput {
	return o
}

// A `features` blocks as documented below.
//
// > **Note** *Because of ADO limitations feed name can be **reserved** for up to 15 minutes after permanent delete of the feed*
func (o FeedOutput) Features() FeedFeatureArrayOutput {
	return o.ApplyT(func(v *Feed) FeedFeatureArrayOutput { return v.Features }).(FeedFeatureArrayOutput)
}

// The name of the Feed.
func (o FeedOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Feed) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ID of the Project Feed is created in. If not specified, feed will be created at the organization level.
func (o FeedOutput) ProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Feed) pulumi.StringPtrOutput { return v.ProjectId }).(pulumi.StringPtrOutput)
}

type FeedArrayOutput struct{ *pulumi.OutputState }

func (FeedArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Feed)(nil)).Elem()
}

func (o FeedArrayOutput) ToFeedArrayOutput() FeedArrayOutput {
	return o
}

func (o FeedArrayOutput) ToFeedArrayOutputWithContext(ctx context.Context) FeedArrayOutput {
	return o
}

func (o FeedArrayOutput) Index(i pulumi.IntInput) FeedOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Feed {
		return vs[0].([]*Feed)[vs[1].(int)]
	}).(FeedOutput)
}

type FeedMapOutput struct{ *pulumi.OutputState }

func (FeedMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Feed)(nil)).Elem()
}

func (o FeedMapOutput) ToFeedMapOutput() FeedMapOutput {
	return o
}

func (o FeedMapOutput) ToFeedMapOutputWithContext(ctx context.Context) FeedMapOutput {
	return o
}

func (o FeedMapOutput) MapIndex(k pulumi.StringInput) FeedOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Feed {
		return vs[0].(map[string]*Feed)[vs[1].(string)]
	}).(FeedOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FeedInput)(nil)).Elem(), &Feed{})
	pulumi.RegisterInputType(reflect.TypeOf((*FeedArrayInput)(nil)).Elem(), FeedArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FeedMapInput)(nil)).Elem(), FeedMap{})
	pulumi.RegisterOutputType(FeedOutput{})
	pulumi.RegisterOutputType(FeedArrayOutput{})
	pulumi.RegisterOutputType(FeedMapOutput{})
}
