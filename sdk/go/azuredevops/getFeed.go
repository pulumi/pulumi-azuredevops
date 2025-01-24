// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azuredevops/sdk/v3/go/azuredevops/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about existing Feed within a given project in Azure DevOps.
//
// ## Example Usage
//
// ### Basic Example
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
//			_, err := azuredevops.LookupFeed(ctx, &azuredevops.LookupFeedArgs{
//				Name: pulumi.StringRef("releases"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ### Access feed within a project
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
//			example, err := azuredevops.LookupProject(ctx, &azuredevops.LookupProjectArgs{
//				Name: pulumi.StringRef("Example Project"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = azuredevops.LookupFeed(ctx, &azuredevops.LookupFeedArgs{
//				Name:      pulumi.StringRef("releases"),
//				ProjectId: pulumi.StringRef(example.Id),
//			}, nil)
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
// - [Azure DevOps Service REST API 7.0 - Feed - Get](https://learn.microsoft.com/en-us/rest/api/azure/devops/artifacts/feed-management/get-feed?view=azure-devops-rest-7.0)
func LookupFeed(ctx *pulumi.Context, args *LookupFeedArgs, opts ...pulumi.InvokeOption) (*LookupFeedResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupFeedResult
	err := ctx.Invoke("azuredevops:index/getFeed:getFeed", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFeed.
type LookupFeedArgs struct {
	// The ID of the Feed.
	//
	// > **Note** Only one of `name` or `feedId` can be set at the same time.
	FeedId *string `pulumi:"feedId"`
	// The Name of the Feed.
	Name *string `pulumi:"name"`
	// ID of the Project Feed is created in.
	ProjectId *string `pulumi:"projectId"`
}

// A collection of values returned by getFeed.
type LookupFeedResult struct {
	// The ID of the Feed.
	FeedId *string `pulumi:"feedId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The name of the Feed.
	Name *string `pulumi:"name"`
	// The ID of the Project.
	ProjectId *string `pulumi:"projectId"`
}

func LookupFeedOutput(ctx *pulumi.Context, args LookupFeedOutputArgs, opts ...pulumi.InvokeOption) LookupFeedResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupFeedResultOutput, error) {
			args := v.(LookupFeedArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azuredevops:index/getFeed:getFeed", args, LookupFeedResultOutput{}, options).(LookupFeedResultOutput), nil
		}).(LookupFeedResultOutput)
}

// A collection of arguments for invoking getFeed.
type LookupFeedOutputArgs struct {
	// The ID of the Feed.
	//
	// > **Note** Only one of `name` or `feedId` can be set at the same time.
	FeedId pulumi.StringPtrInput `pulumi:"feedId"`
	// The Name of the Feed.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// ID of the Project Feed is created in.
	ProjectId pulumi.StringPtrInput `pulumi:"projectId"`
}

func (LookupFeedOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFeedArgs)(nil)).Elem()
}

// A collection of values returned by getFeed.
type LookupFeedResultOutput struct{ *pulumi.OutputState }

func (LookupFeedResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFeedResult)(nil)).Elem()
}

func (o LookupFeedResultOutput) ToLookupFeedResultOutput() LookupFeedResultOutput {
	return o
}

func (o LookupFeedResultOutput) ToLookupFeedResultOutputWithContext(ctx context.Context) LookupFeedResultOutput {
	return o
}

// The ID of the Feed.
func (o LookupFeedResultOutput) FeedId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFeedResult) *string { return v.FeedId }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupFeedResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFeedResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the Feed.
func (o LookupFeedResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFeedResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The ID of the Project.
func (o LookupFeedResultOutput) ProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFeedResult) *string { return v.ProjectId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFeedResultOutput{})
}
