// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package core

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a project within Azure DevOps.
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
// 		_, err := azuredevops.NewProject(ctx, "project", &azuredevops.ProjectArgs{
// 			Description: pulumi.String("Test Project Description"),
// 			Features: pulumi.StringMap{
// 				"artifacts": pulumi.String("disabled"),
// 				"testplans": pulumi.String("disabled"),
// 			},
// 			ProjectName:      pulumi.String("Test Project"),
// 			VersionControl:   pulumi.String("Git"),
// 			Visibility:       pulumi.String("private"),
// 			WorkItemTemplate: pulumi.String("Agile"),
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
// * [Azure DevOps Service REST API 5.1 - Projects](https://docs.microsoft.com/en-us/rest/api/azure/devops/core/projects?view=azure-devops-rest-5.1)
//
// ## PAT Permissions Required
//
// - **Project & Team**: Read, Write, & Manage
//
// Deprecated: azuredevops.core.Project has been deprecated in favor of azuredevops.Project
type Project struct {
	pulumi.CustomResourceState

	// The Description of the Project.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Defines the status (`enabled`, `disabled`) of the project features.\
	// Valid features `boards`, `repositories`, `pipelines`, `testplans`, `artifacts`
	Features pulumi.StringMapOutput `pulumi:"features"`
	// The Process Template ID used by the Project.
	ProcessTemplateId pulumi.StringOutput `pulumi:"processTemplateId"`
	// The Project Name.
	ProjectName pulumi.StringOutput `pulumi:"projectName"`
	// Specifies the version control system. Valid values: `Git` or `Tfvc`. Defaults to `Git`.
	VersionControl pulumi.StringPtrOutput `pulumi:"versionControl"`
	// Specifies the visibility of the Project. Valid values: `private` or `public`. Defaults to `private`.
	Visibility pulumi.StringPtrOutput `pulumi:"visibility"`
	// Specifies the work item template. Defaults to `Agile`.
	WorkItemTemplate pulumi.StringPtrOutput `pulumi:"workItemTemplate"`
}

// NewProject registers a new resource with the given unique name, arguments, and options.
func NewProject(ctx *pulumi.Context,
	name string, args *ProjectArgs, opts ...pulumi.ResourceOption) (*Project, error) {
	if args == nil || args.ProjectName == nil {
		return nil, errors.New("missing required argument 'ProjectName'")
	}
	if args == nil {
		args = &ProjectArgs{}
	}
	var resource Project
	err := ctx.RegisterResource("azuredevops:Core/project:Project", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProject gets an existing Project resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProject(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectState, opts ...pulumi.ResourceOption) (*Project, error) {
	var resource Project
	err := ctx.ReadResource("azuredevops:Core/project:Project", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Project resources.
type projectState struct {
	// The Description of the Project.
	Description *string `pulumi:"description"`
	// Defines the status (`enabled`, `disabled`) of the project features.\
	// Valid features `boards`, `repositories`, `pipelines`, `testplans`, `artifacts`
	Features map[string]string `pulumi:"features"`
	// The Process Template ID used by the Project.
	ProcessTemplateId *string `pulumi:"processTemplateId"`
	// The Project Name.
	ProjectName *string `pulumi:"projectName"`
	// Specifies the version control system. Valid values: `Git` or `Tfvc`. Defaults to `Git`.
	VersionControl *string `pulumi:"versionControl"`
	// Specifies the visibility of the Project. Valid values: `private` or `public`. Defaults to `private`.
	Visibility *string `pulumi:"visibility"`
	// Specifies the work item template. Defaults to `Agile`.
	WorkItemTemplate *string `pulumi:"workItemTemplate"`
}

type ProjectState struct {
	// The Description of the Project.
	Description pulumi.StringPtrInput
	// Defines the status (`enabled`, `disabled`) of the project features.\
	// Valid features `boards`, `repositories`, `pipelines`, `testplans`, `artifacts`
	Features pulumi.StringMapInput
	// The Process Template ID used by the Project.
	ProcessTemplateId pulumi.StringPtrInput
	// The Project Name.
	ProjectName pulumi.StringPtrInput
	// Specifies the version control system. Valid values: `Git` or `Tfvc`. Defaults to `Git`.
	VersionControl pulumi.StringPtrInput
	// Specifies the visibility of the Project. Valid values: `private` or `public`. Defaults to `private`.
	Visibility pulumi.StringPtrInput
	// Specifies the work item template. Defaults to `Agile`.
	WorkItemTemplate pulumi.StringPtrInput
}

func (ProjectState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectState)(nil)).Elem()
}

type projectArgs struct {
	// The Description of the Project.
	Description *string `pulumi:"description"`
	// Defines the status (`enabled`, `disabled`) of the project features.\
	// Valid features `boards`, `repositories`, `pipelines`, `testplans`, `artifacts`
	Features map[string]string `pulumi:"features"`
	// The Project Name.
	ProjectName string `pulumi:"projectName"`
	// Specifies the version control system. Valid values: `Git` or `Tfvc`. Defaults to `Git`.
	VersionControl *string `pulumi:"versionControl"`
	// Specifies the visibility of the Project. Valid values: `private` or `public`. Defaults to `private`.
	Visibility *string `pulumi:"visibility"`
	// Specifies the work item template. Defaults to `Agile`.
	WorkItemTemplate *string `pulumi:"workItemTemplate"`
}

// The set of arguments for constructing a Project resource.
type ProjectArgs struct {
	// The Description of the Project.
	Description pulumi.StringPtrInput
	// Defines the status (`enabled`, `disabled`) of the project features.\
	// Valid features `boards`, `repositories`, `pipelines`, `testplans`, `artifacts`
	Features pulumi.StringMapInput
	// The Project Name.
	ProjectName pulumi.StringInput
	// Specifies the version control system. Valid values: `Git` or `Tfvc`. Defaults to `Git`.
	VersionControl pulumi.StringPtrInput
	// Specifies the visibility of the Project. Valid values: `private` or `public`. Defaults to `private`.
	Visibility pulumi.StringPtrInput
	// Specifies the work item template. Defaults to `Agile`.
	WorkItemTemplate pulumi.StringPtrInput
}

func (ProjectArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectArgs)(nil)).Elem()
}