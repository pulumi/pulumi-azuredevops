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

// Manages permissions for a Service Endpoint
//
// > **Note** Permissions can be assigned to group principals and not to single user principals.
//
// ## Permission levels
//
// Permission for Service Endpoints within Azure DevOps can be applied on two different levels.
// Those levels are reflected by specifying (or omitting) values for the arguments `projectId` and `serviceendpointId`.
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
//			_, err = azuredevops.NewServiceendpointPermissions(ctx, "example-root-permissions", &azuredevops.ServiceendpointPermissionsArgs{
//				ProjectId: example.ID(),
//				Principal: example_readers.ApplyT(func(example_readers azuredevops.GetGroupResult) (*string, error) {
//					return &example_readers.Id, nil
//				}).(pulumi.StringPtrOutput),
//				Permissions: pulumi.StringMap{
//					"Use":               pulumi.String("allow"),
//					"Administer":        pulumi.String("allow"),
//					"Create":            pulumi.String("allow"),
//					"ViewAuthorization": pulumi.String("allow"),
//					"ViewEndpoint":      pulumi.String("allow"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			exampleServiceEndpointDockerRegistry, err := azuredevops.NewServiceEndpointDockerRegistry(ctx, "example", &azuredevops.ServiceEndpointDockerRegistryArgs{
//				ProjectId:           example.ID(),
//				ServiceEndpointName: pulumi.String("Example Docker Hub"),
//				DockerUsername:      pulumi.String("username"),
//				DockerEmail:         pulumi.String("email@example.com"),
//				DockerPassword:      pulumi.String("password"),
//				RegistryType:        pulumi.String("DockerHub"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = azuredevops.NewServiceendpointPermissions(ctx, "example-permissions", &azuredevops.ServiceendpointPermissionsArgs{
//				ProjectId: example.ID(),
//				Principal: example_readers.ApplyT(func(example_readers azuredevops.GetGroupResult) (*string, error) {
//					return &example_readers.Id, nil
//				}).(pulumi.StringPtrOutput),
//				ServiceendpointId: exampleServiceEndpointDockerRegistry.ID(),
//				Permissions: pulumi.StringMap{
//					"Use":               pulumi.String("allow"),
//					"Administer":        pulumi.String("deny"),
//					"Create":            pulumi.String("deny"),
//					"ViewAuthorization": pulumi.String("allow"),
//					"ViewEndpoint":      pulumi.String("allow"),
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
type ServiceendpointPermissions struct {
	pulumi.CustomResourceState

	// the permissions to assign. The following permissions are available.
	Permissions pulumi.StringMapOutput `pulumi:"permissions"`
	// The **group** principal to assign the permissions.
	Principal pulumi.StringOutput `pulumi:"principal"`
	// The ID of the project.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// Replace (`true`) or merge (`false`) the permissions. Default: `true`
	//
	// | Permission        | Description                         |
	// | ----------------- | ----------------------------------- |
	// | Use               | Use service endpoint                |
	// | Administer        | Full control over service endpoints |
	// | Create            | Create service endpoints            |
	// | ViewAuthorization | View authorizations                 |
	// | ViewEndpoint      | View service endpoint properties    |
	Replace pulumi.BoolPtrOutput `pulumi:"replace"`
	// The id of the service endpoint to assign the permissions.
	ServiceendpointId pulumi.StringPtrOutput `pulumi:"serviceendpointId"`
}

// NewServiceendpointPermissions registers a new resource with the given unique name, arguments, and options.
func NewServiceendpointPermissions(ctx *pulumi.Context,
	name string, args *ServiceendpointPermissionsArgs, opts ...pulumi.ResourceOption) (*ServiceendpointPermissions, error) {
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
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ServiceendpointPermissions
	err := ctx.RegisterResource("azuredevops:index/serviceendpointPermissions:ServiceendpointPermissions", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceendpointPermissions gets an existing ServiceendpointPermissions resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceendpointPermissions(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceendpointPermissionsState, opts ...pulumi.ResourceOption) (*ServiceendpointPermissions, error) {
	var resource ServiceendpointPermissions
	err := ctx.ReadResource("azuredevops:index/serviceendpointPermissions:ServiceendpointPermissions", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceendpointPermissions resources.
type serviceendpointPermissionsState struct {
	// the permissions to assign. The following permissions are available.
	Permissions map[string]string `pulumi:"permissions"`
	// The **group** principal to assign the permissions.
	Principal *string `pulumi:"principal"`
	// The ID of the project.
	ProjectId *string `pulumi:"projectId"`
	// Replace (`true`) or merge (`false`) the permissions. Default: `true`
	//
	// | Permission        | Description                         |
	// | ----------------- | ----------------------------------- |
	// | Use               | Use service endpoint                |
	// | Administer        | Full control over service endpoints |
	// | Create            | Create service endpoints            |
	// | ViewAuthorization | View authorizations                 |
	// | ViewEndpoint      | View service endpoint properties    |
	Replace *bool `pulumi:"replace"`
	// The id of the service endpoint to assign the permissions.
	ServiceendpointId *string `pulumi:"serviceendpointId"`
}

type ServiceendpointPermissionsState struct {
	// the permissions to assign. The following permissions are available.
	Permissions pulumi.StringMapInput
	// The **group** principal to assign the permissions.
	Principal pulumi.StringPtrInput
	// The ID of the project.
	ProjectId pulumi.StringPtrInput
	// Replace (`true`) or merge (`false`) the permissions. Default: `true`
	//
	// | Permission        | Description                         |
	// | ----------------- | ----------------------------------- |
	// | Use               | Use service endpoint                |
	// | Administer        | Full control over service endpoints |
	// | Create            | Create service endpoints            |
	// | ViewAuthorization | View authorizations                 |
	// | ViewEndpoint      | View service endpoint properties    |
	Replace pulumi.BoolPtrInput
	// The id of the service endpoint to assign the permissions.
	ServiceendpointId pulumi.StringPtrInput
}

func (ServiceendpointPermissionsState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceendpointPermissionsState)(nil)).Elem()
}

type serviceendpointPermissionsArgs struct {
	// the permissions to assign. The following permissions are available.
	Permissions map[string]string `pulumi:"permissions"`
	// The **group** principal to assign the permissions.
	Principal string `pulumi:"principal"`
	// The ID of the project.
	ProjectId string `pulumi:"projectId"`
	// Replace (`true`) or merge (`false`) the permissions. Default: `true`
	//
	// | Permission        | Description                         |
	// | ----------------- | ----------------------------------- |
	// | Use               | Use service endpoint                |
	// | Administer        | Full control over service endpoints |
	// | Create            | Create service endpoints            |
	// | ViewAuthorization | View authorizations                 |
	// | ViewEndpoint      | View service endpoint properties    |
	Replace *bool `pulumi:"replace"`
	// The id of the service endpoint to assign the permissions.
	ServiceendpointId *string `pulumi:"serviceendpointId"`
}

// The set of arguments for constructing a ServiceendpointPermissions resource.
type ServiceendpointPermissionsArgs struct {
	// the permissions to assign. The following permissions are available.
	Permissions pulumi.StringMapInput
	// The **group** principal to assign the permissions.
	Principal pulumi.StringInput
	// The ID of the project.
	ProjectId pulumi.StringInput
	// Replace (`true`) or merge (`false`) the permissions. Default: `true`
	//
	// | Permission        | Description                         |
	// | ----------------- | ----------------------------------- |
	// | Use               | Use service endpoint                |
	// | Administer        | Full control over service endpoints |
	// | Create            | Create service endpoints            |
	// | ViewAuthorization | View authorizations                 |
	// | ViewEndpoint      | View service endpoint properties    |
	Replace pulumi.BoolPtrInput
	// The id of the service endpoint to assign the permissions.
	ServiceendpointId pulumi.StringPtrInput
}

func (ServiceendpointPermissionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceendpointPermissionsArgs)(nil)).Elem()
}

type ServiceendpointPermissionsInput interface {
	pulumi.Input

	ToServiceendpointPermissionsOutput() ServiceendpointPermissionsOutput
	ToServiceendpointPermissionsOutputWithContext(ctx context.Context) ServiceendpointPermissionsOutput
}

func (*ServiceendpointPermissions) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceendpointPermissions)(nil)).Elem()
}

func (i *ServiceendpointPermissions) ToServiceendpointPermissionsOutput() ServiceendpointPermissionsOutput {
	return i.ToServiceendpointPermissionsOutputWithContext(context.Background())
}

func (i *ServiceendpointPermissions) ToServiceendpointPermissionsOutputWithContext(ctx context.Context) ServiceendpointPermissionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceendpointPermissionsOutput)
}

// ServiceendpointPermissionsArrayInput is an input type that accepts ServiceendpointPermissionsArray and ServiceendpointPermissionsArrayOutput values.
// You can construct a concrete instance of `ServiceendpointPermissionsArrayInput` via:
//
//	ServiceendpointPermissionsArray{ ServiceendpointPermissionsArgs{...} }
type ServiceendpointPermissionsArrayInput interface {
	pulumi.Input

	ToServiceendpointPermissionsArrayOutput() ServiceendpointPermissionsArrayOutput
	ToServiceendpointPermissionsArrayOutputWithContext(context.Context) ServiceendpointPermissionsArrayOutput
}

type ServiceendpointPermissionsArray []ServiceendpointPermissionsInput

func (ServiceendpointPermissionsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceendpointPermissions)(nil)).Elem()
}

func (i ServiceendpointPermissionsArray) ToServiceendpointPermissionsArrayOutput() ServiceendpointPermissionsArrayOutput {
	return i.ToServiceendpointPermissionsArrayOutputWithContext(context.Background())
}

func (i ServiceendpointPermissionsArray) ToServiceendpointPermissionsArrayOutputWithContext(ctx context.Context) ServiceendpointPermissionsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceendpointPermissionsArrayOutput)
}

// ServiceendpointPermissionsMapInput is an input type that accepts ServiceendpointPermissionsMap and ServiceendpointPermissionsMapOutput values.
// You can construct a concrete instance of `ServiceendpointPermissionsMapInput` via:
//
//	ServiceendpointPermissionsMap{ "key": ServiceendpointPermissionsArgs{...} }
type ServiceendpointPermissionsMapInput interface {
	pulumi.Input

	ToServiceendpointPermissionsMapOutput() ServiceendpointPermissionsMapOutput
	ToServiceendpointPermissionsMapOutputWithContext(context.Context) ServiceendpointPermissionsMapOutput
}

type ServiceendpointPermissionsMap map[string]ServiceendpointPermissionsInput

func (ServiceendpointPermissionsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceendpointPermissions)(nil)).Elem()
}

func (i ServiceendpointPermissionsMap) ToServiceendpointPermissionsMapOutput() ServiceendpointPermissionsMapOutput {
	return i.ToServiceendpointPermissionsMapOutputWithContext(context.Background())
}

func (i ServiceendpointPermissionsMap) ToServiceendpointPermissionsMapOutputWithContext(ctx context.Context) ServiceendpointPermissionsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceendpointPermissionsMapOutput)
}

type ServiceendpointPermissionsOutput struct{ *pulumi.OutputState }

func (ServiceendpointPermissionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceendpointPermissions)(nil)).Elem()
}

func (o ServiceendpointPermissionsOutput) ToServiceendpointPermissionsOutput() ServiceendpointPermissionsOutput {
	return o
}

func (o ServiceendpointPermissionsOutput) ToServiceendpointPermissionsOutputWithContext(ctx context.Context) ServiceendpointPermissionsOutput {
	return o
}

// the permissions to assign. The following permissions are available.
func (o ServiceendpointPermissionsOutput) Permissions() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ServiceendpointPermissions) pulumi.StringMapOutput { return v.Permissions }).(pulumi.StringMapOutput)
}

// The **group** principal to assign the permissions.
func (o ServiceendpointPermissionsOutput) Principal() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceendpointPermissions) pulumi.StringOutput { return v.Principal }).(pulumi.StringOutput)
}

// The ID of the project.
func (o ServiceendpointPermissionsOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceendpointPermissions) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// Replace (`true`) or merge (`false`) the permissions. Default: `true`
//
// | Permission        | Description                         |
// | ----------------- | ----------------------------------- |
// | Use               | Use service endpoint                |
// | Administer        | Full control over service endpoints |
// | Create            | Create service endpoints            |
// | ViewAuthorization | View authorizations                 |
// | ViewEndpoint      | View service endpoint properties    |
func (o ServiceendpointPermissionsOutput) Replace() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ServiceendpointPermissions) pulumi.BoolPtrOutput { return v.Replace }).(pulumi.BoolPtrOutput)
}

// The id of the service endpoint to assign the permissions.
func (o ServiceendpointPermissionsOutput) ServiceendpointId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceendpointPermissions) pulumi.StringPtrOutput { return v.ServiceendpointId }).(pulumi.StringPtrOutput)
}

type ServiceendpointPermissionsArrayOutput struct{ *pulumi.OutputState }

func (ServiceendpointPermissionsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceendpointPermissions)(nil)).Elem()
}

func (o ServiceendpointPermissionsArrayOutput) ToServiceendpointPermissionsArrayOutput() ServiceendpointPermissionsArrayOutput {
	return o
}

func (o ServiceendpointPermissionsArrayOutput) ToServiceendpointPermissionsArrayOutputWithContext(ctx context.Context) ServiceendpointPermissionsArrayOutput {
	return o
}

func (o ServiceendpointPermissionsArrayOutput) Index(i pulumi.IntInput) ServiceendpointPermissionsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServiceendpointPermissions {
		return vs[0].([]*ServiceendpointPermissions)[vs[1].(int)]
	}).(ServiceendpointPermissionsOutput)
}

type ServiceendpointPermissionsMapOutput struct{ *pulumi.OutputState }

func (ServiceendpointPermissionsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceendpointPermissions)(nil)).Elem()
}

func (o ServiceendpointPermissionsMapOutput) ToServiceendpointPermissionsMapOutput() ServiceendpointPermissionsMapOutput {
	return o
}

func (o ServiceendpointPermissionsMapOutput) ToServiceendpointPermissionsMapOutputWithContext(ctx context.Context) ServiceendpointPermissionsMapOutput {
	return o
}

func (o ServiceendpointPermissionsMapOutput) MapIndex(k pulumi.StringInput) ServiceendpointPermissionsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServiceendpointPermissions {
		return vs[0].(map[string]*ServiceendpointPermissions)[vs[1].(string)]
	}).(ServiceendpointPermissionsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceendpointPermissionsInput)(nil)).Elem(), &ServiceendpointPermissions{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceendpointPermissionsArrayInput)(nil)).Elem(), ServiceendpointPermissionsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceendpointPermissionsMapInput)(nil)).Elem(), ServiceendpointPermissionsMap{})
	pulumi.RegisterOutputType(ServiceendpointPermissionsOutput{})
	pulumi.RegisterOutputType(ServiceendpointPermissionsArrayOutput{})
	pulumi.RegisterOutputType(ServiceendpointPermissionsMapOutput{})
}
