// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package core

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages features for Azure DevOps projects
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azuredevops/sdk/go/azuredevops"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		tf_project_test_001, err := azuredevops.LookupProject(ctx, &azuredevops.LookupProjectArgs{
// 			ProjectName: "Test Project",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = azuredevops.NewProjectFeatures(ctx, "my_project_features", &azuredevops.ProjectFeaturesArgs{
// 			ProjectId: pulumi.String(tf_project_test_001.Id),
// 			Features: pulumi.StringMap{
// 				"testplans": pulumi.String("disabled"),
// 				"artifacts": pulumi.String("enabled"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ## Relevant Links
//
// No official documentation available
//
// ## PAT Permissions Required
//
// - **Project & Team**: Read, Write, & Manage
//
// Deprecated: azuredevops.core.ProjectFeatures has been deprecated in favor of azuredevops.ProjectFeatures
type ProjectFeatures struct {
	pulumi.CustomResourceState

	// Defines the status (`enabled`, `disabled`) of the project features.\
	// Valid features `boards`, `repositories`, `pipelines`, `testplans`, `artifacts`
	Features  pulumi.StringMapOutput `pulumi:"features"`
	ProjectId pulumi.StringOutput    `pulumi:"projectId"`
}

// NewProjectFeatures registers a new resource with the given unique name, arguments, and options.
func NewProjectFeatures(ctx *pulumi.Context,
	name string, args *ProjectFeaturesArgs, opts ...pulumi.ResourceOption) (*ProjectFeatures, error) {
	if args == nil || args.Features == nil {
		return nil, errors.New("missing required argument 'Features'")
	}
	if args == nil || args.ProjectId == nil {
		return nil, errors.New("missing required argument 'ProjectId'")
	}
	if args == nil {
		args = &ProjectFeaturesArgs{}
	}
	var resource ProjectFeatures
	err := ctx.RegisterResource("azuredevops:Core/projectFeatures:ProjectFeatures", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProjectFeatures gets an existing ProjectFeatures resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProjectFeatures(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectFeaturesState, opts ...pulumi.ResourceOption) (*ProjectFeatures, error) {
	var resource ProjectFeatures
	err := ctx.ReadResource("azuredevops:Core/projectFeatures:ProjectFeatures", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProjectFeatures resources.
type projectFeaturesState struct {
	// Defines the status (`enabled`, `disabled`) of the project features.\
	// Valid features `boards`, `repositories`, `pipelines`, `testplans`, `artifacts`
	Features  map[string]string `pulumi:"features"`
	ProjectId *string           `pulumi:"projectId"`
}

type ProjectFeaturesState struct {
	// Defines the status (`enabled`, `disabled`) of the project features.\
	// Valid features `boards`, `repositories`, `pipelines`, `testplans`, `artifacts`
	Features  pulumi.StringMapInput
	ProjectId pulumi.StringPtrInput
}

func (ProjectFeaturesState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectFeaturesState)(nil)).Elem()
}

type projectFeaturesArgs struct {
	// Defines the status (`enabled`, `disabled`) of the project features.\
	// Valid features `boards`, `repositories`, `pipelines`, `testplans`, `artifacts`
	Features  map[string]string `pulumi:"features"`
	ProjectId string            `pulumi:"projectId"`
}

// The set of arguments for constructing a ProjectFeatures resource.
type ProjectFeaturesArgs struct {
	// Defines the status (`enabled`, `disabled`) of the project features.\
	// Valid features `boards`, `repositories`, `pipelines`, `testplans`, `artifacts`
	Features  pulumi.StringMapInput
	ProjectId pulumi.StringInput
}

func (ProjectFeaturesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectFeaturesArgs)(nil)).Elem()
}