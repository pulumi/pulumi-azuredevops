// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azuredevops/sdk/v3/go/azuredevops/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages permissions for service hooks
//
// ## Permission levels
//
// Permissions for service hooks within Azure DevOps can be applied on the Organizational level or, if the optional attribute `projectId` is specified, on Project level.
// Those levels are reflected by specifying (or omitting) values for the argument `projectId`.
//
// ## Example Usage
//
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
//				WorkItemTemplate: pulumi.String("Agile"),
//				VersionControl:   pulumi.String("Git"),
//				Visibility:       pulumi.String("private"),
//				Description:      pulumi.String("Managed by Terraform"),
//			})
//			if err != nil {
//				return err
//			}
//			example_readers := azuredevops.LookupGroupOutput(ctx, azuredevops.GetGroupOutputArgs{
//				ProjectId: example.ID(),
//				Name:      pulumi.String("Readers"),
//			}, nil)
//			_, err = azuredevops.NewServicehookPermissions(ctx, "example-permissions", &azuredevops.ServicehookPermissionsArgs{
//				ProjectId: example.ID(),
//				Principal: pulumi.String(example_readers.ApplyT(func(example_readers azuredevops.GetGroupResult) (*string, error) {
//					return &example_readers.Id, nil
//				}).(pulumi.StringPtrOutput)),
//				Permissions: pulumi.StringMap{
//					"ViewSubscriptions":   pulumi.String("allow"),
//					"EditSubscriptions":   pulumi.String("allow"),
//					"DeleteSubscriptions": pulumi.String("allow"),
//					"PublishEvents":       pulumi.String("allow"),
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
// * [Azure DevOps Service REST API 7.0 - Security](https://docs.microsoft.com/en-us/rest/api/azure/devops/security/?view=azure-devops-rest-7.0)
//
// ## PAT Permissions Required
//
// - **Project & Team**: vso.security_manage - Grants the ability to read, write, and manage security permissions.
//
// ## Import
//
// The resource does not support import.
type ServicehookPermissions struct {
	pulumi.CustomResourceState

	// the permissions to assign. The following permissions are available.
	Permissions pulumi.StringMapOutput `pulumi:"permissions"`
	// The **group** principal to assign the permissions.
	Principal pulumi.StringOutput `pulumi:"principal"`
	// The ID of the project.
	ProjectId pulumi.StringPtrOutput `pulumi:"projectId"`
	// Replace (`true`) or merge (`false`) the permissions. Default: `true`
	//
	// | Name               | Permission Description   |
	// | ------------------ | ------------------------ |
	// | ViewSubscriptions  | View Subscriptions       |
	// | EditSubscriptions  | Edit Subscription        |
	// | DeleteSubscriptions| Delete Subscriptions     |
	// | PublishEvents      | Publish Events           |
	Replace pulumi.BoolPtrOutput `pulumi:"replace"`
}

// NewServicehookPermissions registers a new resource with the given unique name, arguments, and options.
func NewServicehookPermissions(ctx *pulumi.Context,
	name string, args *ServicehookPermissionsArgs, opts ...pulumi.ResourceOption) (*ServicehookPermissions, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Permissions == nil {
		return nil, errors.New("invalid value for required argument 'Permissions'")
	}
	if args.Principal == nil {
		return nil, errors.New("invalid value for required argument 'Principal'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ServicehookPermissions
	err := ctx.RegisterResource("azuredevops:index/servicehookPermissions:ServicehookPermissions", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServicehookPermissions gets an existing ServicehookPermissions resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServicehookPermissions(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServicehookPermissionsState, opts ...pulumi.ResourceOption) (*ServicehookPermissions, error) {
	var resource ServicehookPermissions
	err := ctx.ReadResource("azuredevops:index/servicehookPermissions:ServicehookPermissions", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServicehookPermissions resources.
type servicehookPermissionsState struct {
	// the permissions to assign. The following permissions are available.
	Permissions map[string]string `pulumi:"permissions"`
	// The **group** principal to assign the permissions.
	Principal *string `pulumi:"principal"`
	// The ID of the project.
	ProjectId *string `pulumi:"projectId"`
	// Replace (`true`) or merge (`false`) the permissions. Default: `true`
	//
	// | Name               | Permission Description   |
	// | ------------------ | ------------------------ |
	// | ViewSubscriptions  | View Subscriptions       |
	// | EditSubscriptions  | Edit Subscription        |
	// | DeleteSubscriptions| Delete Subscriptions     |
	// | PublishEvents      | Publish Events           |
	Replace *bool `pulumi:"replace"`
}

type ServicehookPermissionsState struct {
	// the permissions to assign. The following permissions are available.
	Permissions pulumi.StringMapInput
	// The **group** principal to assign the permissions.
	Principal pulumi.StringPtrInput
	// The ID of the project.
	ProjectId pulumi.StringPtrInput
	// Replace (`true`) or merge (`false`) the permissions. Default: `true`
	//
	// | Name               | Permission Description   |
	// | ------------------ | ------------------------ |
	// | ViewSubscriptions  | View Subscriptions       |
	// | EditSubscriptions  | Edit Subscription        |
	// | DeleteSubscriptions| Delete Subscriptions     |
	// | PublishEvents      | Publish Events           |
	Replace pulumi.BoolPtrInput
}

func (ServicehookPermissionsState) ElementType() reflect.Type {
	return reflect.TypeOf((*servicehookPermissionsState)(nil)).Elem()
}

type servicehookPermissionsArgs struct {
	// the permissions to assign. The following permissions are available.
	Permissions map[string]string `pulumi:"permissions"`
	// The **group** principal to assign the permissions.
	Principal string `pulumi:"principal"`
	// The ID of the project.
	ProjectId *string `pulumi:"projectId"`
	// Replace (`true`) or merge (`false`) the permissions. Default: `true`
	//
	// | Name               | Permission Description   |
	// | ------------------ | ------------------------ |
	// | ViewSubscriptions  | View Subscriptions       |
	// | EditSubscriptions  | Edit Subscription        |
	// | DeleteSubscriptions| Delete Subscriptions     |
	// | PublishEvents      | Publish Events           |
	Replace *bool `pulumi:"replace"`
}

// The set of arguments for constructing a ServicehookPermissions resource.
type ServicehookPermissionsArgs struct {
	// the permissions to assign. The following permissions are available.
	Permissions pulumi.StringMapInput
	// The **group** principal to assign the permissions.
	Principal pulumi.StringInput
	// The ID of the project.
	ProjectId pulumi.StringPtrInput
	// Replace (`true`) or merge (`false`) the permissions. Default: `true`
	//
	// | Name               | Permission Description   |
	// | ------------------ | ------------------------ |
	// | ViewSubscriptions  | View Subscriptions       |
	// | EditSubscriptions  | Edit Subscription        |
	// | DeleteSubscriptions| Delete Subscriptions     |
	// | PublishEvents      | Publish Events           |
	Replace pulumi.BoolPtrInput
}

func (ServicehookPermissionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*servicehookPermissionsArgs)(nil)).Elem()
}

type ServicehookPermissionsInput interface {
	pulumi.Input

	ToServicehookPermissionsOutput() ServicehookPermissionsOutput
	ToServicehookPermissionsOutputWithContext(ctx context.Context) ServicehookPermissionsOutput
}

func (*ServicehookPermissions) ElementType() reflect.Type {
	return reflect.TypeOf((**ServicehookPermissions)(nil)).Elem()
}

func (i *ServicehookPermissions) ToServicehookPermissionsOutput() ServicehookPermissionsOutput {
	return i.ToServicehookPermissionsOutputWithContext(context.Background())
}

func (i *ServicehookPermissions) ToServicehookPermissionsOutputWithContext(ctx context.Context) ServicehookPermissionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicehookPermissionsOutput)
}

// ServicehookPermissionsArrayInput is an input type that accepts ServicehookPermissionsArray and ServicehookPermissionsArrayOutput values.
// You can construct a concrete instance of `ServicehookPermissionsArrayInput` via:
//
//	ServicehookPermissionsArray{ ServicehookPermissionsArgs{...} }
type ServicehookPermissionsArrayInput interface {
	pulumi.Input

	ToServicehookPermissionsArrayOutput() ServicehookPermissionsArrayOutput
	ToServicehookPermissionsArrayOutputWithContext(context.Context) ServicehookPermissionsArrayOutput
}

type ServicehookPermissionsArray []ServicehookPermissionsInput

func (ServicehookPermissionsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServicehookPermissions)(nil)).Elem()
}

func (i ServicehookPermissionsArray) ToServicehookPermissionsArrayOutput() ServicehookPermissionsArrayOutput {
	return i.ToServicehookPermissionsArrayOutputWithContext(context.Background())
}

func (i ServicehookPermissionsArray) ToServicehookPermissionsArrayOutputWithContext(ctx context.Context) ServicehookPermissionsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicehookPermissionsArrayOutput)
}

// ServicehookPermissionsMapInput is an input type that accepts ServicehookPermissionsMap and ServicehookPermissionsMapOutput values.
// You can construct a concrete instance of `ServicehookPermissionsMapInput` via:
//
//	ServicehookPermissionsMap{ "key": ServicehookPermissionsArgs{...} }
type ServicehookPermissionsMapInput interface {
	pulumi.Input

	ToServicehookPermissionsMapOutput() ServicehookPermissionsMapOutput
	ToServicehookPermissionsMapOutputWithContext(context.Context) ServicehookPermissionsMapOutput
}

type ServicehookPermissionsMap map[string]ServicehookPermissionsInput

func (ServicehookPermissionsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServicehookPermissions)(nil)).Elem()
}

func (i ServicehookPermissionsMap) ToServicehookPermissionsMapOutput() ServicehookPermissionsMapOutput {
	return i.ToServicehookPermissionsMapOutputWithContext(context.Background())
}

func (i ServicehookPermissionsMap) ToServicehookPermissionsMapOutputWithContext(ctx context.Context) ServicehookPermissionsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicehookPermissionsMapOutput)
}

type ServicehookPermissionsOutput struct{ *pulumi.OutputState }

func (ServicehookPermissionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServicehookPermissions)(nil)).Elem()
}

func (o ServicehookPermissionsOutput) ToServicehookPermissionsOutput() ServicehookPermissionsOutput {
	return o
}

func (o ServicehookPermissionsOutput) ToServicehookPermissionsOutputWithContext(ctx context.Context) ServicehookPermissionsOutput {
	return o
}

// the permissions to assign. The following permissions are available.
func (o ServicehookPermissionsOutput) Permissions() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ServicehookPermissions) pulumi.StringMapOutput { return v.Permissions }).(pulumi.StringMapOutput)
}

// The **group** principal to assign the permissions.
func (o ServicehookPermissionsOutput) Principal() pulumi.StringOutput {
	return o.ApplyT(func(v *ServicehookPermissions) pulumi.StringOutput { return v.Principal }).(pulumi.StringOutput)
}

// The ID of the project.
func (o ServicehookPermissionsOutput) ProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServicehookPermissions) pulumi.StringPtrOutput { return v.ProjectId }).(pulumi.StringPtrOutput)
}

// Replace (`true`) or merge (`false`) the permissions. Default: `true`
//
// | Name               | Permission Description   |
// | ------------------ | ------------------------ |
// | ViewSubscriptions  | View Subscriptions       |
// | EditSubscriptions  | Edit Subscription        |
// | DeleteSubscriptions| Delete Subscriptions     |
// | PublishEvents      | Publish Events           |
func (o ServicehookPermissionsOutput) Replace() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ServicehookPermissions) pulumi.BoolPtrOutput { return v.Replace }).(pulumi.BoolPtrOutput)
}

type ServicehookPermissionsArrayOutput struct{ *pulumi.OutputState }

func (ServicehookPermissionsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServicehookPermissions)(nil)).Elem()
}

func (o ServicehookPermissionsArrayOutput) ToServicehookPermissionsArrayOutput() ServicehookPermissionsArrayOutput {
	return o
}

func (o ServicehookPermissionsArrayOutput) ToServicehookPermissionsArrayOutputWithContext(ctx context.Context) ServicehookPermissionsArrayOutput {
	return o
}

func (o ServicehookPermissionsArrayOutput) Index(i pulumi.IntInput) ServicehookPermissionsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServicehookPermissions {
		return vs[0].([]*ServicehookPermissions)[vs[1].(int)]
	}).(ServicehookPermissionsOutput)
}

type ServicehookPermissionsMapOutput struct{ *pulumi.OutputState }

func (ServicehookPermissionsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServicehookPermissions)(nil)).Elem()
}

func (o ServicehookPermissionsMapOutput) ToServicehookPermissionsMapOutput() ServicehookPermissionsMapOutput {
	return o
}

func (o ServicehookPermissionsMapOutput) ToServicehookPermissionsMapOutputWithContext(ctx context.Context) ServicehookPermissionsMapOutput {
	return o
}

func (o ServicehookPermissionsMapOutput) MapIndex(k pulumi.StringInput) ServicehookPermissionsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServicehookPermissions {
		return vs[0].(map[string]*ServicehookPermissions)[vs[1].(string)]
	}).(ServicehookPermissionsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServicehookPermissionsInput)(nil)).Elem(), &ServicehookPermissions{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServicehookPermissionsArrayInput)(nil)).Elem(), ServicehookPermissionsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServicehookPermissionsMapInput)(nil)).Elem(), ServicehookPermissionsMap{})
	pulumi.RegisterOutputType(ServicehookPermissionsOutput{})
	pulumi.RegisterOutputType(ServicehookPermissionsArrayOutput{})
	pulumi.RegisterOutputType(ServicehookPermissionsMapOutput{})
}
