// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azuredevops/sdk/v2/go/azuredevops/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Manages a Build Folder.
//
// ## Import
//
// Build Folders can be imported using the `project name/path` or `project id/path`, e.g.
//
// ```sh
//
//	$ pulumi import azuredevops:index/buildFolder:BuildFolder example "Example Project/\\ExampleFolder"
//
// ```
//
//	or
//
// ```sh
//
//	$ pulumi import azuredevops:index/buildFolder:BuildFolder example 00000000-0000-0000-0000-000000000000/\\ExampleFolder
//
// ```
type BuildFolder struct {
	pulumi.CustomResourceState

	// Folder Description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The folder path.
	Path pulumi.StringOutput `pulumi:"path"`
	// The ID of the project in which the folder will be created.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
}

// NewBuildFolder registers a new resource with the given unique name, arguments, and options.
func NewBuildFolder(ctx *pulumi.Context,
	name string, args *BuildFolderArgs, opts ...pulumi.ResourceOption) (*BuildFolder, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Path == nil {
		return nil, errors.New("invalid value for required argument 'Path'")
	}
	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource BuildFolder
	err := ctx.RegisterResource("azuredevops:index/buildFolder:BuildFolder", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBuildFolder gets an existing BuildFolder resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBuildFolder(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BuildFolderState, opts ...pulumi.ResourceOption) (*BuildFolder, error) {
	var resource BuildFolder
	err := ctx.ReadResource("azuredevops:index/buildFolder:BuildFolder", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BuildFolder resources.
type buildFolderState struct {
	// Folder Description.
	Description *string `pulumi:"description"`
	// The folder path.
	Path *string `pulumi:"path"`
	// The ID of the project in which the folder will be created.
	ProjectId *string `pulumi:"projectId"`
}

type BuildFolderState struct {
	// Folder Description.
	Description pulumi.StringPtrInput
	// The folder path.
	Path pulumi.StringPtrInput
	// The ID of the project in which the folder will be created.
	ProjectId pulumi.StringPtrInput
}

func (BuildFolderState) ElementType() reflect.Type {
	return reflect.TypeOf((*buildFolderState)(nil)).Elem()
}

type buildFolderArgs struct {
	// Folder Description.
	Description *string `pulumi:"description"`
	// The folder path.
	Path string `pulumi:"path"`
	// The ID of the project in which the folder will be created.
	ProjectId string `pulumi:"projectId"`
}

// The set of arguments for constructing a BuildFolder resource.
type BuildFolderArgs struct {
	// Folder Description.
	Description pulumi.StringPtrInput
	// The folder path.
	Path pulumi.StringInput
	// The ID of the project in which the folder will be created.
	ProjectId pulumi.StringInput
}

func (BuildFolderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*buildFolderArgs)(nil)).Elem()
}

type BuildFolderInput interface {
	pulumi.Input

	ToBuildFolderOutput() BuildFolderOutput
	ToBuildFolderOutputWithContext(ctx context.Context) BuildFolderOutput
}

func (*BuildFolder) ElementType() reflect.Type {
	return reflect.TypeOf((**BuildFolder)(nil)).Elem()
}

func (i *BuildFolder) ToBuildFolderOutput() BuildFolderOutput {
	return i.ToBuildFolderOutputWithContext(context.Background())
}

func (i *BuildFolder) ToBuildFolderOutputWithContext(ctx context.Context) BuildFolderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BuildFolderOutput)
}

func (i *BuildFolder) ToOutput(ctx context.Context) pulumix.Output[*BuildFolder] {
	return pulumix.Output[*BuildFolder]{
		OutputState: i.ToBuildFolderOutputWithContext(ctx).OutputState,
	}
}

// BuildFolderArrayInput is an input type that accepts BuildFolderArray and BuildFolderArrayOutput values.
// You can construct a concrete instance of `BuildFolderArrayInput` via:
//
//	BuildFolderArray{ BuildFolderArgs{...} }
type BuildFolderArrayInput interface {
	pulumi.Input

	ToBuildFolderArrayOutput() BuildFolderArrayOutput
	ToBuildFolderArrayOutputWithContext(context.Context) BuildFolderArrayOutput
}

type BuildFolderArray []BuildFolderInput

func (BuildFolderArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BuildFolder)(nil)).Elem()
}

func (i BuildFolderArray) ToBuildFolderArrayOutput() BuildFolderArrayOutput {
	return i.ToBuildFolderArrayOutputWithContext(context.Background())
}

func (i BuildFolderArray) ToBuildFolderArrayOutputWithContext(ctx context.Context) BuildFolderArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BuildFolderArrayOutput)
}

func (i BuildFolderArray) ToOutput(ctx context.Context) pulumix.Output[[]*BuildFolder] {
	return pulumix.Output[[]*BuildFolder]{
		OutputState: i.ToBuildFolderArrayOutputWithContext(ctx).OutputState,
	}
}

// BuildFolderMapInput is an input type that accepts BuildFolderMap and BuildFolderMapOutput values.
// You can construct a concrete instance of `BuildFolderMapInput` via:
//
//	BuildFolderMap{ "key": BuildFolderArgs{...} }
type BuildFolderMapInput interface {
	pulumi.Input

	ToBuildFolderMapOutput() BuildFolderMapOutput
	ToBuildFolderMapOutputWithContext(context.Context) BuildFolderMapOutput
}

type BuildFolderMap map[string]BuildFolderInput

func (BuildFolderMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BuildFolder)(nil)).Elem()
}

func (i BuildFolderMap) ToBuildFolderMapOutput() BuildFolderMapOutput {
	return i.ToBuildFolderMapOutputWithContext(context.Background())
}

func (i BuildFolderMap) ToBuildFolderMapOutputWithContext(ctx context.Context) BuildFolderMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BuildFolderMapOutput)
}

func (i BuildFolderMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*BuildFolder] {
	return pulumix.Output[map[string]*BuildFolder]{
		OutputState: i.ToBuildFolderMapOutputWithContext(ctx).OutputState,
	}
}

type BuildFolderOutput struct{ *pulumi.OutputState }

func (BuildFolderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BuildFolder)(nil)).Elem()
}

func (o BuildFolderOutput) ToBuildFolderOutput() BuildFolderOutput {
	return o
}

func (o BuildFolderOutput) ToBuildFolderOutputWithContext(ctx context.Context) BuildFolderOutput {
	return o
}

func (o BuildFolderOutput) ToOutput(ctx context.Context) pulumix.Output[*BuildFolder] {
	return pulumix.Output[*BuildFolder]{
		OutputState: o.OutputState,
	}
}

// Folder Description.
func (o BuildFolderOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BuildFolder) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The folder path.
func (o BuildFolderOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v *BuildFolder) pulumi.StringOutput { return v.Path }).(pulumi.StringOutput)
}

// The ID of the project in which the folder will be created.
func (o BuildFolderOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *BuildFolder) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

type BuildFolderArrayOutput struct{ *pulumi.OutputState }

func (BuildFolderArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BuildFolder)(nil)).Elem()
}

func (o BuildFolderArrayOutput) ToBuildFolderArrayOutput() BuildFolderArrayOutput {
	return o
}

func (o BuildFolderArrayOutput) ToBuildFolderArrayOutputWithContext(ctx context.Context) BuildFolderArrayOutput {
	return o
}

func (o BuildFolderArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*BuildFolder] {
	return pulumix.Output[[]*BuildFolder]{
		OutputState: o.OutputState,
	}
}

func (o BuildFolderArrayOutput) Index(i pulumi.IntInput) BuildFolderOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BuildFolder {
		return vs[0].([]*BuildFolder)[vs[1].(int)]
	}).(BuildFolderOutput)
}

type BuildFolderMapOutput struct{ *pulumi.OutputState }

func (BuildFolderMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BuildFolder)(nil)).Elem()
}

func (o BuildFolderMapOutput) ToBuildFolderMapOutput() BuildFolderMapOutput {
	return o
}

func (o BuildFolderMapOutput) ToBuildFolderMapOutputWithContext(ctx context.Context) BuildFolderMapOutput {
	return o
}

func (o BuildFolderMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*BuildFolder] {
	return pulumix.Output[map[string]*BuildFolder]{
		OutputState: o.OutputState,
	}
}

func (o BuildFolderMapOutput) MapIndex(k pulumi.StringInput) BuildFolderOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BuildFolder {
		return vs[0].(map[string]*BuildFolder)[vs[1].(string)]
	}).(BuildFolderOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BuildFolderInput)(nil)).Elem(), &BuildFolder{})
	pulumi.RegisterInputType(reflect.TypeOf((*BuildFolderArrayInput)(nil)).Elem(), BuildFolderArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BuildFolderMapInput)(nil)).Elem(), BuildFolderMap{})
	pulumi.RegisterOutputType(BuildFolderOutput{})
	pulumi.RegisterOutputType(BuildFolderArrayOutput{})
	pulumi.RegisterOutputType(BuildFolderMapOutput{})
}
