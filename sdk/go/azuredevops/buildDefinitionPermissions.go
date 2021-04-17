// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages permissions for a Build Definition
//
// > **Note** Permissions can be assigned to group principals and not to single user principals.
//
// ## Relevant Links
//
// * [Azure DevOps Service REST API 5.1 - Security](https://docs.microsoft.com/en-us/rest/api/azure/devops/security/?view=azure-devops-rest-5.1)
//
// ## PAT Permissions Required
//
// - **Project & Team**: vso.security_manage - Grants the ability to read, write, and manage security permissions.
//
// ## Import
//
// The resource does not support import.
type BuildDefinitionPermissions struct {
	pulumi.CustomResourceState

	// The id of the build definition to assign the permissions.
	BuildDefinitionId pulumi.StringOutput `pulumi:"buildDefinitionId"`
	// the permissions to assign. The following permissions are available.
	Permissions pulumi.StringMapOutput `pulumi:"permissions"`
	// The **group** principal to assign the permissions.
	Principal pulumi.StringOutput `pulumi:"principal"`
	// The ID of the project to assign the permissions.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// Replace (`true`) or merge (`false`) the permissions. Default: `true`.
	Replace pulumi.BoolPtrOutput `pulumi:"replace"`
}

// NewBuildDefinitionPermissions registers a new resource with the given unique name, arguments, and options.
func NewBuildDefinitionPermissions(ctx *pulumi.Context,
	name string, args *BuildDefinitionPermissionsArgs, opts ...pulumi.ResourceOption) (*BuildDefinitionPermissions, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BuildDefinitionId == nil {
		return nil, errors.New("invalid value for required argument 'BuildDefinitionId'")
	}
	if args.Permissions == nil {
		return nil, errors.New("invalid value for required argument 'Permissions'")
	}
	if args.Principal == nil {
		return nil, errors.New("invalid value for required argument 'Principal'")
	}
	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	var resource BuildDefinitionPermissions
	err := ctx.RegisterResource("azuredevops:index/buildDefinitionPermissions:BuildDefinitionPermissions", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBuildDefinitionPermissions gets an existing BuildDefinitionPermissions resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBuildDefinitionPermissions(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BuildDefinitionPermissionsState, opts ...pulumi.ResourceOption) (*BuildDefinitionPermissions, error) {
	var resource BuildDefinitionPermissions
	err := ctx.ReadResource("azuredevops:index/buildDefinitionPermissions:BuildDefinitionPermissions", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BuildDefinitionPermissions resources.
type buildDefinitionPermissionsState struct {
	// The id of the build definition to assign the permissions.
	BuildDefinitionId *string `pulumi:"buildDefinitionId"`
	// the permissions to assign. The following permissions are available.
	Permissions map[string]string `pulumi:"permissions"`
	// The **group** principal to assign the permissions.
	Principal *string `pulumi:"principal"`
	// The ID of the project to assign the permissions.
	ProjectId *string `pulumi:"projectId"`
	// Replace (`true`) or merge (`false`) the permissions. Default: `true`.
	Replace *bool `pulumi:"replace"`
}

type BuildDefinitionPermissionsState struct {
	// The id of the build definition to assign the permissions.
	BuildDefinitionId pulumi.StringPtrInput
	// the permissions to assign. The following permissions are available.
	Permissions pulumi.StringMapInput
	// The **group** principal to assign the permissions.
	Principal pulumi.StringPtrInput
	// The ID of the project to assign the permissions.
	ProjectId pulumi.StringPtrInput
	// Replace (`true`) or merge (`false`) the permissions. Default: `true`.
	Replace pulumi.BoolPtrInput
}

func (BuildDefinitionPermissionsState) ElementType() reflect.Type {
	return reflect.TypeOf((*buildDefinitionPermissionsState)(nil)).Elem()
}

type buildDefinitionPermissionsArgs struct {
	// The id of the build definition to assign the permissions.
	BuildDefinitionId string `pulumi:"buildDefinitionId"`
	// the permissions to assign. The following permissions are available.
	Permissions map[string]string `pulumi:"permissions"`
	// The **group** principal to assign the permissions.
	Principal string `pulumi:"principal"`
	// The ID of the project to assign the permissions.
	ProjectId string `pulumi:"projectId"`
	// Replace (`true`) or merge (`false`) the permissions. Default: `true`.
	Replace *bool `pulumi:"replace"`
}

// The set of arguments for constructing a BuildDefinitionPermissions resource.
type BuildDefinitionPermissionsArgs struct {
	// The id of the build definition to assign the permissions.
	BuildDefinitionId pulumi.StringInput
	// the permissions to assign. The following permissions are available.
	Permissions pulumi.StringMapInput
	// The **group** principal to assign the permissions.
	Principal pulumi.StringInput
	// The ID of the project to assign the permissions.
	ProjectId pulumi.StringInput
	// Replace (`true`) or merge (`false`) the permissions. Default: `true`.
	Replace pulumi.BoolPtrInput
}

func (BuildDefinitionPermissionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*buildDefinitionPermissionsArgs)(nil)).Elem()
}

type BuildDefinitionPermissionsInput interface {
	pulumi.Input

	ToBuildDefinitionPermissionsOutput() BuildDefinitionPermissionsOutput
	ToBuildDefinitionPermissionsOutputWithContext(ctx context.Context) BuildDefinitionPermissionsOutput
}

func (*BuildDefinitionPermissions) ElementType() reflect.Type {
	return reflect.TypeOf((*BuildDefinitionPermissions)(nil))
}

func (i *BuildDefinitionPermissions) ToBuildDefinitionPermissionsOutput() BuildDefinitionPermissionsOutput {
	return i.ToBuildDefinitionPermissionsOutputWithContext(context.Background())
}

func (i *BuildDefinitionPermissions) ToBuildDefinitionPermissionsOutputWithContext(ctx context.Context) BuildDefinitionPermissionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BuildDefinitionPermissionsOutput)
}

func (i *BuildDefinitionPermissions) ToBuildDefinitionPermissionsPtrOutput() BuildDefinitionPermissionsPtrOutput {
	return i.ToBuildDefinitionPermissionsPtrOutputWithContext(context.Background())
}

func (i *BuildDefinitionPermissions) ToBuildDefinitionPermissionsPtrOutputWithContext(ctx context.Context) BuildDefinitionPermissionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BuildDefinitionPermissionsPtrOutput)
}

type BuildDefinitionPermissionsPtrInput interface {
	pulumi.Input

	ToBuildDefinitionPermissionsPtrOutput() BuildDefinitionPermissionsPtrOutput
	ToBuildDefinitionPermissionsPtrOutputWithContext(ctx context.Context) BuildDefinitionPermissionsPtrOutput
}

type buildDefinitionPermissionsPtrType BuildDefinitionPermissionsArgs

func (*buildDefinitionPermissionsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BuildDefinitionPermissions)(nil))
}

func (i *buildDefinitionPermissionsPtrType) ToBuildDefinitionPermissionsPtrOutput() BuildDefinitionPermissionsPtrOutput {
	return i.ToBuildDefinitionPermissionsPtrOutputWithContext(context.Background())
}

func (i *buildDefinitionPermissionsPtrType) ToBuildDefinitionPermissionsPtrOutputWithContext(ctx context.Context) BuildDefinitionPermissionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BuildDefinitionPermissionsPtrOutput)
}

// BuildDefinitionPermissionsArrayInput is an input type that accepts BuildDefinitionPermissionsArray and BuildDefinitionPermissionsArrayOutput values.
// You can construct a concrete instance of `BuildDefinitionPermissionsArrayInput` via:
//
//          BuildDefinitionPermissionsArray{ BuildDefinitionPermissionsArgs{...} }
type BuildDefinitionPermissionsArrayInput interface {
	pulumi.Input

	ToBuildDefinitionPermissionsArrayOutput() BuildDefinitionPermissionsArrayOutput
	ToBuildDefinitionPermissionsArrayOutputWithContext(context.Context) BuildDefinitionPermissionsArrayOutput
}

type BuildDefinitionPermissionsArray []BuildDefinitionPermissionsInput

func (BuildDefinitionPermissionsArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*BuildDefinitionPermissions)(nil))
}

func (i BuildDefinitionPermissionsArray) ToBuildDefinitionPermissionsArrayOutput() BuildDefinitionPermissionsArrayOutput {
	return i.ToBuildDefinitionPermissionsArrayOutputWithContext(context.Background())
}

func (i BuildDefinitionPermissionsArray) ToBuildDefinitionPermissionsArrayOutputWithContext(ctx context.Context) BuildDefinitionPermissionsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BuildDefinitionPermissionsArrayOutput)
}

// BuildDefinitionPermissionsMapInput is an input type that accepts BuildDefinitionPermissionsMap and BuildDefinitionPermissionsMapOutput values.
// You can construct a concrete instance of `BuildDefinitionPermissionsMapInput` via:
//
//          BuildDefinitionPermissionsMap{ "key": BuildDefinitionPermissionsArgs{...} }
type BuildDefinitionPermissionsMapInput interface {
	pulumi.Input

	ToBuildDefinitionPermissionsMapOutput() BuildDefinitionPermissionsMapOutput
	ToBuildDefinitionPermissionsMapOutputWithContext(context.Context) BuildDefinitionPermissionsMapOutput
}

type BuildDefinitionPermissionsMap map[string]BuildDefinitionPermissionsInput

func (BuildDefinitionPermissionsMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*BuildDefinitionPermissions)(nil))
}

func (i BuildDefinitionPermissionsMap) ToBuildDefinitionPermissionsMapOutput() BuildDefinitionPermissionsMapOutput {
	return i.ToBuildDefinitionPermissionsMapOutputWithContext(context.Background())
}

func (i BuildDefinitionPermissionsMap) ToBuildDefinitionPermissionsMapOutputWithContext(ctx context.Context) BuildDefinitionPermissionsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BuildDefinitionPermissionsMapOutput)
}

type BuildDefinitionPermissionsOutput struct {
	*pulumi.OutputState
}

func (BuildDefinitionPermissionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BuildDefinitionPermissions)(nil))
}

func (o BuildDefinitionPermissionsOutput) ToBuildDefinitionPermissionsOutput() BuildDefinitionPermissionsOutput {
	return o
}

func (o BuildDefinitionPermissionsOutput) ToBuildDefinitionPermissionsOutputWithContext(ctx context.Context) BuildDefinitionPermissionsOutput {
	return o
}

func (o BuildDefinitionPermissionsOutput) ToBuildDefinitionPermissionsPtrOutput() BuildDefinitionPermissionsPtrOutput {
	return o.ToBuildDefinitionPermissionsPtrOutputWithContext(context.Background())
}

func (o BuildDefinitionPermissionsOutput) ToBuildDefinitionPermissionsPtrOutputWithContext(ctx context.Context) BuildDefinitionPermissionsPtrOutput {
	return o.ApplyT(func(v BuildDefinitionPermissions) *BuildDefinitionPermissions {
		return &v
	}).(BuildDefinitionPermissionsPtrOutput)
}

type BuildDefinitionPermissionsPtrOutput struct {
	*pulumi.OutputState
}

func (BuildDefinitionPermissionsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BuildDefinitionPermissions)(nil))
}

func (o BuildDefinitionPermissionsPtrOutput) ToBuildDefinitionPermissionsPtrOutput() BuildDefinitionPermissionsPtrOutput {
	return o
}

func (o BuildDefinitionPermissionsPtrOutput) ToBuildDefinitionPermissionsPtrOutputWithContext(ctx context.Context) BuildDefinitionPermissionsPtrOutput {
	return o
}

type BuildDefinitionPermissionsArrayOutput struct{ *pulumi.OutputState }

func (BuildDefinitionPermissionsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BuildDefinitionPermissions)(nil))
}

func (o BuildDefinitionPermissionsArrayOutput) ToBuildDefinitionPermissionsArrayOutput() BuildDefinitionPermissionsArrayOutput {
	return o
}

func (o BuildDefinitionPermissionsArrayOutput) ToBuildDefinitionPermissionsArrayOutputWithContext(ctx context.Context) BuildDefinitionPermissionsArrayOutput {
	return o
}

func (o BuildDefinitionPermissionsArrayOutput) Index(i pulumi.IntInput) BuildDefinitionPermissionsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) BuildDefinitionPermissions {
		return vs[0].([]BuildDefinitionPermissions)[vs[1].(int)]
	}).(BuildDefinitionPermissionsOutput)
}

type BuildDefinitionPermissionsMapOutput struct{ *pulumi.OutputState }

func (BuildDefinitionPermissionsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]BuildDefinitionPermissions)(nil))
}

func (o BuildDefinitionPermissionsMapOutput) ToBuildDefinitionPermissionsMapOutput() BuildDefinitionPermissionsMapOutput {
	return o
}

func (o BuildDefinitionPermissionsMapOutput) ToBuildDefinitionPermissionsMapOutputWithContext(ctx context.Context) BuildDefinitionPermissionsMapOutput {
	return o
}

func (o BuildDefinitionPermissionsMapOutput) MapIndex(k pulumi.StringInput) BuildDefinitionPermissionsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) BuildDefinitionPermissions {
		return vs[0].(map[string]BuildDefinitionPermissions)[vs[1].(string)]
	}).(BuildDefinitionPermissionsOutput)
}

func init() {
	pulumi.RegisterOutputType(BuildDefinitionPermissionsOutput{})
	pulumi.RegisterOutputType(BuildDefinitionPermissionsPtrOutput{})
	pulumi.RegisterOutputType(BuildDefinitionPermissionsArrayOutput{})
	pulumi.RegisterOutputType(BuildDefinitionPermissionsMapOutput{})
}
