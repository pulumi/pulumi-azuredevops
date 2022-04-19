// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a project within Azure DevOps.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azuredevops/sdk/v2/go/azuredevops"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := azuredevops.NewProject(ctx, "example", &azuredevops.ProjectArgs{
// 			Description: pulumi.String("Managed by Terraform"),
// 			Features: pulumi.StringMap{
// 				"artifacts": pulumi.String("disabled"),
// 				"testplans": pulumi.String("disabled"),
// 			},
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
// - [Azure DevOps Service REST API 6.0 - Projects](https://docs.microsoft.com/en-us/rest/api/azure/devops/core/projects?view=azure-devops-rest-6.0)
//
// ## PAT Permissions Required
//
// - **Project & Team**: Read, Write, & Manage
//
// ## Import
//
// Azure DevOps Projects can be imported using the project name or by the project Guid, e.g.
//
// ```sh
//  $ pulumi import azuredevops:index/project:Project example "Example Project"
// ```
//
//  or
//
// ```sh
//  $ pulumi import azuredevops:index/project:Project example 00000000-0000-0000-0000-000000000000
// ```
type Project struct {
	pulumi.CustomResourceState

	// The Description of the Project.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Defines the status (`enabled`, `disabled`) of the project features.\
	// Valid features are `boards`, `repositories`, `pipelines`, `testplans`, `artifacts`
	Features pulumi.StringMapOutput `pulumi:"features"`
	// The Project Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The Process Template ID used by the Project.
	ProcessTemplateId pulumi.StringOutput `pulumi:"processTemplateId"`
	// Specifies the version control system. Valid values: `Git` or `Tfvc`. Defaults to `Git`.
	VersionControl pulumi.StringPtrOutput `pulumi:"versionControl"`
	// Specifies the visibility of the Project. Valid values: `private` or `public`. Defaults to `private`.
	Visibility pulumi.StringPtrOutput `pulumi:"visibility"`
	// Specifies the work item template. Valid values: `Agile`, `Basic`, `CMMI` or `Scrum`. Defaults to `Agile`.
	WorkItemTemplate pulumi.StringPtrOutput `pulumi:"workItemTemplate"`
}

// NewProject registers a new resource with the given unique name, arguments, and options.
func NewProject(ctx *pulumi.Context,
	name string, args *ProjectArgs, opts ...pulumi.ResourceOption) (*Project, error) {
	if args == nil {
		args = &ProjectArgs{}
	}

	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azuredevops:Core/project:Project"),
		},
	})
	opts = append(opts, aliases)
	var resource Project
	err := ctx.RegisterResource("azuredevops:index/project:Project", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azuredevops:index/project:Project", name, id, state, &resource, opts...)
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
	// Valid features are `boards`, `repositories`, `pipelines`, `testplans`, `artifacts`
	Features map[string]string `pulumi:"features"`
	// The Project Name.
	Name *string `pulumi:"name"`
	// The Process Template ID used by the Project.
	ProcessTemplateId *string `pulumi:"processTemplateId"`
	// Specifies the version control system. Valid values: `Git` or `Tfvc`. Defaults to `Git`.
	VersionControl *string `pulumi:"versionControl"`
	// Specifies the visibility of the Project. Valid values: `private` or `public`. Defaults to `private`.
	Visibility *string `pulumi:"visibility"`
	// Specifies the work item template. Valid values: `Agile`, `Basic`, `CMMI` or `Scrum`. Defaults to `Agile`.
	WorkItemTemplate *string `pulumi:"workItemTemplate"`
}

type ProjectState struct {
	// The Description of the Project.
	Description pulumi.StringPtrInput
	// Defines the status (`enabled`, `disabled`) of the project features.\
	// Valid features are `boards`, `repositories`, `pipelines`, `testplans`, `artifacts`
	Features pulumi.StringMapInput
	// The Project Name.
	Name pulumi.StringPtrInput
	// The Process Template ID used by the Project.
	ProcessTemplateId pulumi.StringPtrInput
	// Specifies the version control system. Valid values: `Git` or `Tfvc`. Defaults to `Git`.
	VersionControl pulumi.StringPtrInput
	// Specifies the visibility of the Project. Valid values: `private` or `public`. Defaults to `private`.
	Visibility pulumi.StringPtrInput
	// Specifies the work item template. Valid values: `Agile`, `Basic`, `CMMI` or `Scrum`. Defaults to `Agile`.
	WorkItemTemplate pulumi.StringPtrInput
}

func (ProjectState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectState)(nil)).Elem()
}

type projectArgs struct {
	// The Description of the Project.
	Description *string `pulumi:"description"`
	// Defines the status (`enabled`, `disabled`) of the project features.\
	// Valid features are `boards`, `repositories`, `pipelines`, `testplans`, `artifacts`
	Features map[string]string `pulumi:"features"`
	// The Project Name.
	Name *string `pulumi:"name"`
	// Specifies the version control system. Valid values: `Git` or `Tfvc`. Defaults to `Git`.
	VersionControl *string `pulumi:"versionControl"`
	// Specifies the visibility of the Project. Valid values: `private` or `public`. Defaults to `private`.
	Visibility *string `pulumi:"visibility"`
	// Specifies the work item template. Valid values: `Agile`, `Basic`, `CMMI` or `Scrum`. Defaults to `Agile`.
	WorkItemTemplate *string `pulumi:"workItemTemplate"`
}

// The set of arguments for constructing a Project resource.
type ProjectArgs struct {
	// The Description of the Project.
	Description pulumi.StringPtrInput
	// Defines the status (`enabled`, `disabled`) of the project features.\
	// Valid features are `boards`, `repositories`, `pipelines`, `testplans`, `artifacts`
	Features pulumi.StringMapInput
	// The Project Name.
	Name pulumi.StringPtrInput
	// Specifies the version control system. Valid values: `Git` or `Tfvc`. Defaults to `Git`.
	VersionControl pulumi.StringPtrInput
	// Specifies the visibility of the Project. Valid values: `private` or `public`. Defaults to `private`.
	Visibility pulumi.StringPtrInput
	// Specifies the work item template. Valid values: `Agile`, `Basic`, `CMMI` or `Scrum`. Defaults to `Agile`.
	WorkItemTemplate pulumi.StringPtrInput
}

func (ProjectArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectArgs)(nil)).Elem()
}

type ProjectInput interface {
	pulumi.Input

	ToProjectOutput() ProjectOutput
	ToProjectOutputWithContext(ctx context.Context) ProjectOutput
}

func (*Project) ElementType() reflect.Type {
	return reflect.TypeOf((**Project)(nil)).Elem()
}

func (i *Project) ToProjectOutput() ProjectOutput {
	return i.ToProjectOutputWithContext(context.Background())
}

func (i *Project) ToProjectOutputWithContext(ctx context.Context) ProjectOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectOutput)
}

// ProjectArrayInput is an input type that accepts ProjectArray and ProjectArrayOutput values.
// You can construct a concrete instance of `ProjectArrayInput` via:
//
//          ProjectArray{ ProjectArgs{...} }
type ProjectArrayInput interface {
	pulumi.Input

	ToProjectArrayOutput() ProjectArrayOutput
	ToProjectArrayOutputWithContext(context.Context) ProjectArrayOutput
}

type ProjectArray []ProjectInput

func (ProjectArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Project)(nil)).Elem()
}

func (i ProjectArray) ToProjectArrayOutput() ProjectArrayOutput {
	return i.ToProjectArrayOutputWithContext(context.Background())
}

func (i ProjectArray) ToProjectArrayOutputWithContext(ctx context.Context) ProjectArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectArrayOutput)
}

// ProjectMapInput is an input type that accepts ProjectMap and ProjectMapOutput values.
// You can construct a concrete instance of `ProjectMapInput` via:
//
//          ProjectMap{ "key": ProjectArgs{...} }
type ProjectMapInput interface {
	pulumi.Input

	ToProjectMapOutput() ProjectMapOutput
	ToProjectMapOutputWithContext(context.Context) ProjectMapOutput
}

type ProjectMap map[string]ProjectInput

func (ProjectMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Project)(nil)).Elem()
}

func (i ProjectMap) ToProjectMapOutput() ProjectMapOutput {
	return i.ToProjectMapOutputWithContext(context.Background())
}

func (i ProjectMap) ToProjectMapOutputWithContext(ctx context.Context) ProjectMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectMapOutput)
}

type ProjectOutput struct{ *pulumi.OutputState }

func (ProjectOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Project)(nil)).Elem()
}

func (o ProjectOutput) ToProjectOutput() ProjectOutput {
	return o
}

func (o ProjectOutput) ToProjectOutputWithContext(ctx context.Context) ProjectOutput {
	return o
}

type ProjectArrayOutput struct{ *pulumi.OutputState }

func (ProjectArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Project)(nil)).Elem()
}

func (o ProjectArrayOutput) ToProjectArrayOutput() ProjectArrayOutput {
	return o
}

func (o ProjectArrayOutput) ToProjectArrayOutputWithContext(ctx context.Context) ProjectArrayOutput {
	return o
}

func (o ProjectArrayOutput) Index(i pulumi.IntInput) ProjectOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Project {
		return vs[0].([]*Project)[vs[1].(int)]
	}).(ProjectOutput)
}

type ProjectMapOutput struct{ *pulumi.OutputState }

func (ProjectMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Project)(nil)).Elem()
}

func (o ProjectMapOutput) ToProjectMapOutput() ProjectMapOutput {
	return o
}

func (o ProjectMapOutput) ToProjectMapOutputWithContext(ctx context.Context) ProjectMapOutput {
	return o
}

func (o ProjectMapOutput) MapIndex(k pulumi.StringInput) ProjectOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Project {
		return vs[0].(map[string]*Project)[vs[1].(string)]
	}).(ProjectOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectInput)(nil)).Elem(), &Project{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectArrayInput)(nil)).Elem(), ProjectArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectMapInput)(nil)).Elem(), ProjectMap{})
	pulumi.RegisterOutputType(ProjectOutput{})
	pulumi.RegisterOutputType(ProjectArrayOutput{})
	pulumi.RegisterOutputType(ProjectMapOutput{})
}
