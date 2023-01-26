// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Bitbucket service endpoint within Azure DevOps.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-azuredevops/sdk/v2/go/azuredevops"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			exampleProject, err := azuredevops.NewProject(ctx, "exampleProject", &azuredevops.ProjectArgs{
//				Visibility:       pulumi.String("private"),
//				VersionControl:   pulumi.String("Git"),
//				WorkItemTemplate: pulumi.String("Agile"),
//				Description:      pulumi.String("Managed by Terraform"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = azuredevops.NewServiceEndpointBitBucket(ctx, "exampleServiceEndpointBitBucket", &azuredevops.ServiceEndpointBitBucketArgs{
//				ProjectId:           exampleProject.ID(),
//				Username:            pulumi.String("username"),
//				Password:            pulumi.String("password"),
//				ServiceEndpointName: pulumi.String("Example Bitbucket"),
//				Description:         pulumi.String("Managed by Terraform"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ## Relevant Links
//
// - [Azure DevOps Service REST API 6.0 - Agent Pools](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-6.0)
//
// ## Import
//
// Azure DevOps Service Endpoint Bitbucket can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**
//
// ```sh
//
//	$ pulumi import azuredevops:index/serviceEndpointBitBucket:ServiceEndpointBitBucket example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
//
// ```
type ServiceEndpointBitBucket struct {
	pulumi.CustomResourceState

	Authorization pulumi.StringMapOutput `pulumi:"authorization"`
	Description   pulumi.StringPtrOutput `pulumi:"description"`
	// Bitbucket account password.
	Password pulumi.StringOutput `pulumi:"password"`
	// A bcrypted hash of the attribute 'password'
	PasswordHash pulumi.StringOutput `pulumi:"passwordHash"`
	// The ID of the project.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The Service Endpoint name.
	ServiceEndpointName pulumi.StringOutput `pulumi:"serviceEndpointName"`
	// Bitbucket account username.
	Username pulumi.StringOutput `pulumi:"username"`
}

// NewServiceEndpointBitBucket registers a new resource with the given unique name, arguments, and options.
func NewServiceEndpointBitBucket(ctx *pulumi.Context,
	name string, args *ServiceEndpointBitBucketArgs, opts ...pulumi.ResourceOption) (*ServiceEndpointBitBucket, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Password == nil {
		return nil, errors.New("invalid value for required argument 'Password'")
	}
	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.ServiceEndpointName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceEndpointName'")
	}
	if args.Username == nil {
		return nil, errors.New("invalid value for required argument 'Username'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azuredevops:ServiceEndpoint/bitBucket:BitBucket"),
		},
	})
	opts = append(opts, aliases)
	if args.Password != nil {
		args.Password = pulumi.ToSecret(args.Password).(pulumi.StringOutput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"password",
		"passwordHash",
	})
	opts = append(opts, secrets)
	var resource ServiceEndpointBitBucket
	err := ctx.RegisterResource("azuredevops:index/serviceEndpointBitBucket:ServiceEndpointBitBucket", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceEndpointBitBucket gets an existing ServiceEndpointBitBucket resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceEndpointBitBucket(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceEndpointBitBucketState, opts ...pulumi.ResourceOption) (*ServiceEndpointBitBucket, error) {
	var resource ServiceEndpointBitBucket
	err := ctx.ReadResource("azuredevops:index/serviceEndpointBitBucket:ServiceEndpointBitBucket", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceEndpointBitBucket resources.
type serviceEndpointBitBucketState struct {
	Authorization map[string]string `pulumi:"authorization"`
	Description   *string           `pulumi:"description"`
	// Bitbucket account password.
	Password *string `pulumi:"password"`
	// A bcrypted hash of the attribute 'password'
	PasswordHash *string `pulumi:"passwordHash"`
	// The ID of the project.
	ProjectId *string `pulumi:"projectId"`
	// The Service Endpoint name.
	ServiceEndpointName *string `pulumi:"serviceEndpointName"`
	// Bitbucket account username.
	Username *string `pulumi:"username"`
}

type ServiceEndpointBitBucketState struct {
	Authorization pulumi.StringMapInput
	Description   pulumi.StringPtrInput
	// Bitbucket account password.
	Password pulumi.StringPtrInput
	// A bcrypted hash of the attribute 'password'
	PasswordHash pulumi.StringPtrInput
	// The ID of the project.
	ProjectId pulumi.StringPtrInput
	// The Service Endpoint name.
	ServiceEndpointName pulumi.StringPtrInput
	// Bitbucket account username.
	Username pulumi.StringPtrInput
}

func (ServiceEndpointBitBucketState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceEndpointBitBucketState)(nil)).Elem()
}

type serviceEndpointBitBucketArgs struct {
	Authorization map[string]string `pulumi:"authorization"`
	Description   *string           `pulumi:"description"`
	// Bitbucket account password.
	Password string `pulumi:"password"`
	// The ID of the project.
	ProjectId string `pulumi:"projectId"`
	// The Service Endpoint name.
	ServiceEndpointName string `pulumi:"serviceEndpointName"`
	// Bitbucket account username.
	Username string `pulumi:"username"`
}

// The set of arguments for constructing a ServiceEndpointBitBucket resource.
type ServiceEndpointBitBucketArgs struct {
	Authorization pulumi.StringMapInput
	Description   pulumi.StringPtrInput
	// Bitbucket account password.
	Password pulumi.StringInput
	// The ID of the project.
	ProjectId pulumi.StringInput
	// The Service Endpoint name.
	ServiceEndpointName pulumi.StringInput
	// Bitbucket account username.
	Username pulumi.StringInput
}

func (ServiceEndpointBitBucketArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceEndpointBitBucketArgs)(nil)).Elem()
}

type ServiceEndpointBitBucketInput interface {
	pulumi.Input

	ToServiceEndpointBitBucketOutput() ServiceEndpointBitBucketOutput
	ToServiceEndpointBitBucketOutputWithContext(ctx context.Context) ServiceEndpointBitBucketOutput
}

func (*ServiceEndpointBitBucket) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceEndpointBitBucket)(nil)).Elem()
}

func (i *ServiceEndpointBitBucket) ToServiceEndpointBitBucketOutput() ServiceEndpointBitBucketOutput {
	return i.ToServiceEndpointBitBucketOutputWithContext(context.Background())
}

func (i *ServiceEndpointBitBucket) ToServiceEndpointBitBucketOutputWithContext(ctx context.Context) ServiceEndpointBitBucketOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointBitBucketOutput)
}

// ServiceEndpointBitBucketArrayInput is an input type that accepts ServiceEndpointBitBucketArray and ServiceEndpointBitBucketArrayOutput values.
// You can construct a concrete instance of `ServiceEndpointBitBucketArrayInput` via:
//
//	ServiceEndpointBitBucketArray{ ServiceEndpointBitBucketArgs{...} }
type ServiceEndpointBitBucketArrayInput interface {
	pulumi.Input

	ToServiceEndpointBitBucketArrayOutput() ServiceEndpointBitBucketArrayOutput
	ToServiceEndpointBitBucketArrayOutputWithContext(context.Context) ServiceEndpointBitBucketArrayOutput
}

type ServiceEndpointBitBucketArray []ServiceEndpointBitBucketInput

func (ServiceEndpointBitBucketArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceEndpointBitBucket)(nil)).Elem()
}

func (i ServiceEndpointBitBucketArray) ToServiceEndpointBitBucketArrayOutput() ServiceEndpointBitBucketArrayOutput {
	return i.ToServiceEndpointBitBucketArrayOutputWithContext(context.Background())
}

func (i ServiceEndpointBitBucketArray) ToServiceEndpointBitBucketArrayOutputWithContext(ctx context.Context) ServiceEndpointBitBucketArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointBitBucketArrayOutput)
}

// ServiceEndpointBitBucketMapInput is an input type that accepts ServiceEndpointBitBucketMap and ServiceEndpointBitBucketMapOutput values.
// You can construct a concrete instance of `ServiceEndpointBitBucketMapInput` via:
//
//	ServiceEndpointBitBucketMap{ "key": ServiceEndpointBitBucketArgs{...} }
type ServiceEndpointBitBucketMapInput interface {
	pulumi.Input

	ToServiceEndpointBitBucketMapOutput() ServiceEndpointBitBucketMapOutput
	ToServiceEndpointBitBucketMapOutputWithContext(context.Context) ServiceEndpointBitBucketMapOutput
}

type ServiceEndpointBitBucketMap map[string]ServiceEndpointBitBucketInput

func (ServiceEndpointBitBucketMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceEndpointBitBucket)(nil)).Elem()
}

func (i ServiceEndpointBitBucketMap) ToServiceEndpointBitBucketMapOutput() ServiceEndpointBitBucketMapOutput {
	return i.ToServiceEndpointBitBucketMapOutputWithContext(context.Background())
}

func (i ServiceEndpointBitBucketMap) ToServiceEndpointBitBucketMapOutputWithContext(ctx context.Context) ServiceEndpointBitBucketMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointBitBucketMapOutput)
}

type ServiceEndpointBitBucketOutput struct{ *pulumi.OutputState }

func (ServiceEndpointBitBucketOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceEndpointBitBucket)(nil)).Elem()
}

func (o ServiceEndpointBitBucketOutput) ToServiceEndpointBitBucketOutput() ServiceEndpointBitBucketOutput {
	return o
}

func (o ServiceEndpointBitBucketOutput) ToServiceEndpointBitBucketOutputWithContext(ctx context.Context) ServiceEndpointBitBucketOutput {
	return o
}

func (o ServiceEndpointBitBucketOutput) Authorization() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ServiceEndpointBitBucket) pulumi.StringMapOutput { return v.Authorization }).(pulumi.StringMapOutput)
}

func (o ServiceEndpointBitBucketOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceEndpointBitBucket) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Bitbucket account password.
func (o ServiceEndpointBitBucketOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEndpointBitBucket) pulumi.StringOutput { return v.Password }).(pulumi.StringOutput)
}

// A bcrypted hash of the attribute 'password'
func (o ServiceEndpointBitBucketOutput) PasswordHash() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEndpointBitBucket) pulumi.StringOutput { return v.PasswordHash }).(pulumi.StringOutput)
}

// The ID of the project.
func (o ServiceEndpointBitBucketOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEndpointBitBucket) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The Service Endpoint name.
func (o ServiceEndpointBitBucketOutput) ServiceEndpointName() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEndpointBitBucket) pulumi.StringOutput { return v.ServiceEndpointName }).(pulumi.StringOutput)
}

// Bitbucket account username.
func (o ServiceEndpointBitBucketOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEndpointBitBucket) pulumi.StringOutput { return v.Username }).(pulumi.StringOutput)
}

type ServiceEndpointBitBucketArrayOutput struct{ *pulumi.OutputState }

func (ServiceEndpointBitBucketArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceEndpointBitBucket)(nil)).Elem()
}

func (o ServiceEndpointBitBucketArrayOutput) ToServiceEndpointBitBucketArrayOutput() ServiceEndpointBitBucketArrayOutput {
	return o
}

func (o ServiceEndpointBitBucketArrayOutput) ToServiceEndpointBitBucketArrayOutputWithContext(ctx context.Context) ServiceEndpointBitBucketArrayOutput {
	return o
}

func (o ServiceEndpointBitBucketArrayOutput) Index(i pulumi.IntInput) ServiceEndpointBitBucketOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServiceEndpointBitBucket {
		return vs[0].([]*ServiceEndpointBitBucket)[vs[1].(int)]
	}).(ServiceEndpointBitBucketOutput)
}

type ServiceEndpointBitBucketMapOutput struct{ *pulumi.OutputState }

func (ServiceEndpointBitBucketMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceEndpointBitBucket)(nil)).Elem()
}

func (o ServiceEndpointBitBucketMapOutput) ToServiceEndpointBitBucketMapOutput() ServiceEndpointBitBucketMapOutput {
	return o
}

func (o ServiceEndpointBitBucketMapOutput) ToServiceEndpointBitBucketMapOutputWithContext(ctx context.Context) ServiceEndpointBitBucketMapOutput {
	return o
}

func (o ServiceEndpointBitBucketMapOutput) MapIndex(k pulumi.StringInput) ServiceEndpointBitBucketOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServiceEndpointBitBucket {
		return vs[0].(map[string]*ServiceEndpointBitBucket)[vs[1].(string)]
	}).(ServiceEndpointBitBucketOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceEndpointBitBucketInput)(nil)).Elem(), &ServiceEndpointBitBucket{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceEndpointBitBucketArrayInput)(nil)).Elem(), ServiceEndpointBitBucketArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceEndpointBitBucketMapInput)(nil)).Elem(), ServiceEndpointBitBucketMap{})
	pulumi.RegisterOutputType(ServiceEndpointBitBucketOutput{})
	pulumi.RegisterOutputType(ServiceEndpointBitBucketArrayOutput{})
	pulumi.RegisterOutputType(ServiceEndpointBitBucketMapOutput{})
}
