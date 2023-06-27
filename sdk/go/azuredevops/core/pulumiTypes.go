// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package core

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type GetProjectsProject struct {
	// Name of the Project, if not specified all projects will be returned.
	Name string `pulumi:"name"`
	// The ID of the Project.
	ProjectId string `pulumi:"projectId"`
	// Url to the full version of the object.
	ProjectUrl string `pulumi:"projectUrl"`
	// State of the Project, if not specified all projects will be returned. Valid values are `all`, `deleting`, `new`, `wellFormed`, `createPending`, `unchanged`,`deleted`.
	//
	// DataSource without specifying any arguments will return all projects.
	State string `pulumi:"state"`
}

// GetProjectsProjectInput is an input type that accepts GetProjectsProjectArgs and GetProjectsProjectOutput values.
// You can construct a concrete instance of `GetProjectsProjectInput` via:
//
//	GetProjectsProjectArgs{...}
type GetProjectsProjectInput interface {
	pulumi.Input

	ToGetProjectsProjectOutput() GetProjectsProjectOutput
	ToGetProjectsProjectOutputWithContext(context.Context) GetProjectsProjectOutput
}

type GetProjectsProjectArgs struct {
	// Name of the Project, if not specified all projects will be returned.
	Name pulumi.StringInput `pulumi:"name"`
	// The ID of the Project.
	ProjectId pulumi.StringInput `pulumi:"projectId"`
	// Url to the full version of the object.
	ProjectUrl pulumi.StringInput `pulumi:"projectUrl"`
	// State of the Project, if not specified all projects will be returned. Valid values are `all`, `deleting`, `new`, `wellFormed`, `createPending`, `unchanged`,`deleted`.
	//
	// DataSource without specifying any arguments will return all projects.
	State pulumi.StringInput `pulumi:"state"`
}

func (GetProjectsProjectArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProjectsProject)(nil)).Elem()
}

func (i GetProjectsProjectArgs) ToGetProjectsProjectOutput() GetProjectsProjectOutput {
	return i.ToGetProjectsProjectOutputWithContext(context.Background())
}

func (i GetProjectsProjectArgs) ToGetProjectsProjectOutputWithContext(ctx context.Context) GetProjectsProjectOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetProjectsProjectOutput)
}

// GetProjectsProjectArrayInput is an input type that accepts GetProjectsProjectArray and GetProjectsProjectArrayOutput values.
// You can construct a concrete instance of `GetProjectsProjectArrayInput` via:
//
//	GetProjectsProjectArray{ GetProjectsProjectArgs{...} }
type GetProjectsProjectArrayInput interface {
	pulumi.Input

	ToGetProjectsProjectArrayOutput() GetProjectsProjectArrayOutput
	ToGetProjectsProjectArrayOutputWithContext(context.Context) GetProjectsProjectArrayOutput
}

type GetProjectsProjectArray []GetProjectsProjectInput

func (GetProjectsProjectArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetProjectsProject)(nil)).Elem()
}

func (i GetProjectsProjectArray) ToGetProjectsProjectArrayOutput() GetProjectsProjectArrayOutput {
	return i.ToGetProjectsProjectArrayOutputWithContext(context.Background())
}

func (i GetProjectsProjectArray) ToGetProjectsProjectArrayOutputWithContext(ctx context.Context) GetProjectsProjectArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetProjectsProjectArrayOutput)
}

type GetProjectsProjectOutput struct{ *pulumi.OutputState }

func (GetProjectsProjectOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProjectsProject)(nil)).Elem()
}

func (o GetProjectsProjectOutput) ToGetProjectsProjectOutput() GetProjectsProjectOutput {
	return o
}

func (o GetProjectsProjectOutput) ToGetProjectsProjectOutputWithContext(ctx context.Context) GetProjectsProjectOutput {
	return o
}

// Name of the Project, if not specified all projects will be returned.
func (o GetProjectsProjectOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetProjectsProject) string { return v.Name }).(pulumi.StringOutput)
}

// The ID of the Project.
func (o GetProjectsProjectOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v GetProjectsProject) string { return v.ProjectId }).(pulumi.StringOutput)
}

// Url to the full version of the object.
func (o GetProjectsProjectOutput) ProjectUrl() pulumi.StringOutput {
	return o.ApplyT(func(v GetProjectsProject) string { return v.ProjectUrl }).(pulumi.StringOutput)
}

// State of the Project, if not specified all projects will be returned. Valid values are `all`, `deleting`, `new`, `wellFormed`, `createPending`, `unchanged`,`deleted`.
//
// DataSource without specifying any arguments will return all projects.
func (o GetProjectsProjectOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v GetProjectsProject) string { return v.State }).(pulumi.StringOutput)
}

type GetProjectsProjectArrayOutput struct{ *pulumi.OutputState }

func (GetProjectsProjectArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetProjectsProject)(nil)).Elem()
}

func (o GetProjectsProjectArrayOutput) ToGetProjectsProjectArrayOutput() GetProjectsProjectArrayOutput {
	return o
}

func (o GetProjectsProjectArrayOutput) ToGetProjectsProjectArrayOutputWithContext(ctx context.Context) GetProjectsProjectArrayOutput {
	return o
}

func (o GetProjectsProjectArrayOutput) Index(i pulumi.IntInput) GetProjectsProjectOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetProjectsProject {
		return vs[0].([]GetProjectsProject)[vs[1].(int)]
	}).(GetProjectsProjectOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GetProjectsProjectInput)(nil)).Elem(), GetProjectsProjectArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetProjectsProjectArrayInput)(nil)).Elem(), GetProjectsProjectArray{})
	pulumi.RegisterOutputType(GetProjectsProjectOutput{})
	pulumi.RegisterOutputType(GetProjectsProjectArrayOutput{})
}
