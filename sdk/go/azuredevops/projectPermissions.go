// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages permissions for a AzureDevOps project
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
type ProjectPermissions struct {
	pulumi.CustomResourceState

	// the permissions to assign. The following permissions are available
	Permissions pulumi.StringMapOutput `pulumi:"permissions"`
	// The **group** principal to assign the permissions.
	Principal pulumi.StringOutput `pulumi:"principal"`
	// The ID of the project to assign the permissions.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// Replace (`true`) or merge (`false`) the permissions. Default: `true`
	Replace pulumi.BoolPtrOutput `pulumi:"replace"`
}

// NewProjectPermissions registers a new resource with the given unique name, arguments, and options.
func NewProjectPermissions(ctx *pulumi.Context,
	name string, args *ProjectPermissionsArgs, opts ...pulumi.ResourceOption) (*ProjectPermissions, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
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
	var resource ProjectPermissions
	err := ctx.RegisterResource("azuredevops:index/projectPermissions:ProjectPermissions", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProjectPermissions gets an existing ProjectPermissions resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProjectPermissions(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectPermissionsState, opts ...pulumi.ResourceOption) (*ProjectPermissions, error) {
	var resource ProjectPermissions
	err := ctx.ReadResource("azuredevops:index/projectPermissions:ProjectPermissions", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProjectPermissions resources.
type projectPermissionsState struct {
	// the permissions to assign. The following permissions are available
	Permissions map[string]string `pulumi:"permissions"`
	// The **group** principal to assign the permissions.
	Principal *string `pulumi:"principal"`
	// The ID of the project to assign the permissions.
	ProjectId *string `pulumi:"projectId"`
	// Replace (`true`) or merge (`false`) the permissions. Default: `true`
	Replace *bool `pulumi:"replace"`
}

type ProjectPermissionsState struct {
	// the permissions to assign. The following permissions are available
	Permissions pulumi.StringMapInput
	// The **group** principal to assign the permissions.
	Principal pulumi.StringPtrInput
	// The ID of the project to assign the permissions.
	ProjectId pulumi.StringPtrInput
	// Replace (`true`) or merge (`false`) the permissions. Default: `true`
	Replace pulumi.BoolPtrInput
}

func (ProjectPermissionsState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectPermissionsState)(nil)).Elem()
}

type projectPermissionsArgs struct {
	// the permissions to assign. The following permissions are available
	Permissions map[string]string `pulumi:"permissions"`
	// The **group** principal to assign the permissions.
	Principal string `pulumi:"principal"`
	// The ID of the project to assign the permissions.
	ProjectId string `pulumi:"projectId"`
	// Replace (`true`) or merge (`false`) the permissions. Default: `true`
	Replace *bool `pulumi:"replace"`
}

// The set of arguments for constructing a ProjectPermissions resource.
type ProjectPermissionsArgs struct {
	// the permissions to assign. The following permissions are available
	Permissions pulumi.StringMapInput
	// The **group** principal to assign the permissions.
	Principal pulumi.StringInput
	// The ID of the project to assign the permissions.
	ProjectId pulumi.StringInput
	// Replace (`true`) or merge (`false`) the permissions. Default: `true`
	Replace pulumi.BoolPtrInput
}

func (ProjectPermissionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectPermissionsArgs)(nil)).Elem()
}

type ProjectPermissionsInput interface {
	pulumi.Input

	ToProjectPermissionsOutput() ProjectPermissionsOutput
	ToProjectPermissionsOutputWithContext(ctx context.Context) ProjectPermissionsOutput
}

func (ProjectPermissions) ElementType() reflect.Type {
	return reflect.TypeOf((*ProjectPermissions)(nil)).Elem()
}

func (i ProjectPermissions) ToProjectPermissionsOutput() ProjectPermissionsOutput {
	return i.ToProjectPermissionsOutputWithContext(context.Background())
}

func (i ProjectPermissions) ToProjectPermissionsOutputWithContext(ctx context.Context) ProjectPermissionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectPermissionsOutput)
}

type ProjectPermissionsOutput struct {
	*pulumi.OutputState
}

func (ProjectPermissionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProjectPermissionsOutput)(nil)).Elem()
}

func (o ProjectPermissionsOutput) ToProjectPermissionsOutput() ProjectPermissionsOutput {
	return o
}

func (o ProjectPermissionsOutput) ToProjectPermissionsOutputWithContext(ctx context.Context) ProjectPermissionsOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ProjectPermissionsOutput{})
}
