// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azuredevops/sdk/v3/go/azuredevops/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about an Environment.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
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
//			exampleProject, err := azuredevops.NewProject(ctx, "example", &azuredevops.ProjectArgs{
//				Name:             pulumi.String("Example Project"),
//				WorkItemTemplate: pulumi.String("Agile"),
//				VersionControl:   pulumi.String("Git"),
//				Visibility:       pulumi.String("private"),
//				Description:      pulumi.String("Managed by Terraform"),
//			})
//			if err != nil {
//				return err
//			}
//			exampleEnvironment, err := azuredevops.NewEnvironment(ctx, "example", &azuredevops.EnvironmentArgs{
//				ProjectId:   exampleProject.ID(),
//				Name:        pulumi.String("Example Environment"),
//				Description: pulumi.String("Managed by Terraform"),
//			})
//			if err != nil {
//				return err
//			}
//			_ = pulumi.All(exampleProject.ID(), exampleEnvironment.ID()).ApplyT(func(_args []interface{}) (azuredevops.GetEnvironmentResult, error) {
//				exampleProjectId := _args[0].(string)
//				exampleEnvironmentId := _args[1].(string)
//				return azuredevops.LookupEnvironmentOutput(ctx, azuredevops.GetEnvironmentOutputArgs{
//					ProjectId:     exampleProjectId,
//					EnvironmentId: exampleEnvironmentId,
//				}, nil), nil
//			}).(azuredevops.GetEnvironmentResultOutput)
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ## Relevant Links
//
// * [Azure DevOps Service REST API 7.0 - Environments](https://docs.microsoft.com/en-us/rest/api/azure/devops/distributedtask/environments?view=azure-devops-rest-7.0)
func LookupEnvironment(ctx *pulumi.Context, args *LookupEnvironmentArgs, opts ...pulumi.InvokeOption) (*LookupEnvironmentResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupEnvironmentResult
	err := ctx.Invoke("azuredevops:index/getEnvironment:getEnvironment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getEnvironment.
type LookupEnvironmentArgs struct {
	// The ID of the Environment.
	EnvironmentId *int `pulumi:"environmentId"`
	// Name of the Environment.
	//
	// > **NOTE:** One of either `environmentId` or `name` must be specified.
	Name *string `pulumi:"name"`
	// The ID of the project.
	ProjectId string `pulumi:"projectId"`
}

// A collection of values returned by getEnvironment.
type LookupEnvironmentResult struct {
	// A description for the Environment.
	Description   string `pulumi:"description"`
	EnvironmentId *int   `pulumi:"environmentId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The name of the Environment.
	Name      string `pulumi:"name"`
	ProjectId string `pulumi:"projectId"`
}

func LookupEnvironmentOutput(ctx *pulumi.Context, args LookupEnvironmentOutputArgs, opts ...pulumi.InvokeOption) LookupEnvironmentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEnvironmentResult, error) {
			args := v.(LookupEnvironmentArgs)
			r, err := LookupEnvironment(ctx, &args, opts...)
			var s LookupEnvironmentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEnvironmentResultOutput)
}

// A collection of arguments for invoking getEnvironment.
type LookupEnvironmentOutputArgs struct {
	// The ID of the Environment.
	EnvironmentId pulumi.IntPtrInput `pulumi:"environmentId"`
	// Name of the Environment.
	//
	// > **NOTE:** One of either `environmentId` or `name` must be specified.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The ID of the project.
	ProjectId pulumi.StringInput `pulumi:"projectId"`
}

func (LookupEnvironmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEnvironmentArgs)(nil)).Elem()
}

// A collection of values returned by getEnvironment.
type LookupEnvironmentResultOutput struct{ *pulumi.OutputState }

func (LookupEnvironmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEnvironmentResult)(nil)).Elem()
}

func (o LookupEnvironmentResultOutput) ToLookupEnvironmentResultOutput() LookupEnvironmentResultOutput {
	return o
}

func (o LookupEnvironmentResultOutput) ToLookupEnvironmentResultOutputWithContext(ctx context.Context) LookupEnvironmentResultOutput {
	return o
}

// A description for the Environment.
func (o LookupEnvironmentResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupEnvironmentResultOutput) EnvironmentId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) *int { return v.EnvironmentId }).(pulumi.IntPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupEnvironmentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the Environment.
func (o LookupEnvironmentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupEnvironmentResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEnvironmentResultOutput{})
}
