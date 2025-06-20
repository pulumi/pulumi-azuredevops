// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azuredevops/sdk/v3/go/azuredevops/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

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
//				Visibility:       pulumi.String("private"),
//				VersionControl:   pulumi.String("Git"),
//				WorkItemTemplate: pulumi.String("Agile"),
//				Description:      pulumi.String("Managed by Pulumi"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = azuredevops.NewServiceendpointGcpTerraform(ctx, "example", &azuredevops.ServiceendpointGcpTerraformArgs{
//				ProjectId:           example.ID(),
//				TokenUri:            pulumi.String("https://oauth2.example.com/token"),
//				ClientEmail:         pulumi.String("gcp-sa-example@example.iam.gserviceaccount.com"),
//				PrivateKey:          pulumi.String("0000000000000000000000000000000000000"),
//				ServiceEndpointName: pulumi.String("Example GCP Terraform extension"),
//				GcpProjectId:        pulumi.String("Example GCP Project"),
//				Description:         pulumi.String("Managed by Pulumi"),
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
// - [Azure DevOps Service REST API 7.1 - Service Endpoints](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-7.1)
//
// ## Import
//
// Azure DevOps GCP for Terraform Service Endpoint can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**
//
// ```sh
// $ pulumi import azuredevops:index/serviceendpointGcpTerraform:ServiceendpointGcpTerraform  azuredevops_serviceendpoint_gcp_terraform.example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
// ```
type ServiceendpointGcpTerraform struct {
	pulumi.CustomResourceState

	Authorization pulumi.StringMapOutput `pulumi:"authorization"`
	// The client email field in the JSON key file for creating the JSON Web Token.
	ClientEmail pulumi.StringPtrOutput `pulumi:"clientEmail"`
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// GCP project associated with the Service Connection.
	GcpProjectId pulumi.StringOutput `pulumi:"gcpProjectId"`
	// The Private Key for connecting to the endpoint.
	PrivateKey pulumi.StringOutput `pulumi:"privateKey"`
	// The ID of the project.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// Scope to be provided.
	Scope pulumi.StringPtrOutput `pulumi:"scope"`
	// The Service Endpoint name.
	ServiceEndpointName pulumi.StringOutput `pulumi:"serviceEndpointName"`
	// The token uri field in the JSON key file for creating the JSON Web Token.
	TokenUri pulumi.StringOutput `pulumi:"tokenUri"`
}

// NewServiceendpointGcpTerraform registers a new resource with the given unique name, arguments, and options.
func NewServiceendpointGcpTerraform(ctx *pulumi.Context,
	name string, args *ServiceendpointGcpTerraformArgs, opts ...pulumi.ResourceOption) (*ServiceendpointGcpTerraform, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.ServiceEndpointName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceEndpointName'")
	}
	if args.PrivateKey != nil {
		args.PrivateKey = pulumi.ToSecret(args.PrivateKey).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"privateKey",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ServiceendpointGcpTerraform
	err := ctx.RegisterResource("azuredevops:index/serviceendpointGcpTerraform:ServiceendpointGcpTerraform", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceendpointGcpTerraform gets an existing ServiceendpointGcpTerraform resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceendpointGcpTerraform(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceendpointGcpTerraformState, opts ...pulumi.ResourceOption) (*ServiceendpointGcpTerraform, error) {
	var resource ServiceendpointGcpTerraform
	err := ctx.ReadResource("azuredevops:index/serviceendpointGcpTerraform:ServiceendpointGcpTerraform", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceendpointGcpTerraform resources.
type serviceendpointGcpTerraformState struct {
	Authorization map[string]string `pulumi:"authorization"`
	// The client email field in the JSON key file for creating the JSON Web Token.
	ClientEmail *string `pulumi:"clientEmail"`
	Description *string `pulumi:"description"`
	// GCP project associated with the Service Connection.
	GcpProjectId *string `pulumi:"gcpProjectId"`
	// The Private Key for connecting to the endpoint.
	PrivateKey *string `pulumi:"privateKey"`
	// The ID of the project.
	ProjectId *string `pulumi:"projectId"`
	// Scope to be provided.
	Scope *string `pulumi:"scope"`
	// The Service Endpoint name.
	ServiceEndpointName *string `pulumi:"serviceEndpointName"`
	// The token uri field in the JSON key file for creating the JSON Web Token.
	TokenUri *string `pulumi:"tokenUri"`
}

type ServiceendpointGcpTerraformState struct {
	Authorization pulumi.StringMapInput
	// The client email field in the JSON key file for creating the JSON Web Token.
	ClientEmail pulumi.StringPtrInput
	Description pulumi.StringPtrInput
	// GCP project associated with the Service Connection.
	GcpProjectId pulumi.StringPtrInput
	// The Private Key for connecting to the endpoint.
	PrivateKey pulumi.StringPtrInput
	// The ID of the project.
	ProjectId pulumi.StringPtrInput
	// Scope to be provided.
	Scope pulumi.StringPtrInput
	// The Service Endpoint name.
	ServiceEndpointName pulumi.StringPtrInput
	// The token uri field in the JSON key file for creating the JSON Web Token.
	TokenUri pulumi.StringPtrInput
}

func (ServiceendpointGcpTerraformState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceendpointGcpTerraformState)(nil)).Elem()
}

type serviceendpointGcpTerraformArgs struct {
	// The client email field in the JSON key file for creating the JSON Web Token.
	ClientEmail *string `pulumi:"clientEmail"`
	Description *string `pulumi:"description"`
	// GCP project associated with the Service Connection.
	GcpProjectId *string `pulumi:"gcpProjectId"`
	// The Private Key for connecting to the endpoint.
	PrivateKey *string `pulumi:"privateKey"`
	// The ID of the project.
	ProjectId string `pulumi:"projectId"`
	// Scope to be provided.
	Scope *string `pulumi:"scope"`
	// The Service Endpoint name.
	ServiceEndpointName string `pulumi:"serviceEndpointName"`
	// The token uri field in the JSON key file for creating the JSON Web Token.
	TokenUri *string `pulumi:"tokenUri"`
}

// The set of arguments for constructing a ServiceendpointGcpTerraform resource.
type ServiceendpointGcpTerraformArgs struct {
	// The client email field in the JSON key file for creating the JSON Web Token.
	ClientEmail pulumi.StringPtrInput
	Description pulumi.StringPtrInput
	// GCP project associated with the Service Connection.
	GcpProjectId pulumi.StringPtrInput
	// The Private Key for connecting to the endpoint.
	PrivateKey pulumi.StringPtrInput
	// The ID of the project.
	ProjectId pulumi.StringInput
	// Scope to be provided.
	Scope pulumi.StringPtrInput
	// The Service Endpoint name.
	ServiceEndpointName pulumi.StringInput
	// The token uri field in the JSON key file for creating the JSON Web Token.
	TokenUri pulumi.StringPtrInput
}

func (ServiceendpointGcpTerraformArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceendpointGcpTerraformArgs)(nil)).Elem()
}

type ServiceendpointGcpTerraformInput interface {
	pulumi.Input

	ToServiceendpointGcpTerraformOutput() ServiceendpointGcpTerraformOutput
	ToServiceendpointGcpTerraformOutputWithContext(ctx context.Context) ServiceendpointGcpTerraformOutput
}

func (*ServiceendpointGcpTerraform) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceendpointGcpTerraform)(nil)).Elem()
}

func (i *ServiceendpointGcpTerraform) ToServiceendpointGcpTerraformOutput() ServiceendpointGcpTerraformOutput {
	return i.ToServiceendpointGcpTerraformOutputWithContext(context.Background())
}

func (i *ServiceendpointGcpTerraform) ToServiceendpointGcpTerraformOutputWithContext(ctx context.Context) ServiceendpointGcpTerraformOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceendpointGcpTerraformOutput)
}

// ServiceendpointGcpTerraformArrayInput is an input type that accepts ServiceendpointGcpTerraformArray and ServiceendpointGcpTerraformArrayOutput values.
// You can construct a concrete instance of `ServiceendpointGcpTerraformArrayInput` via:
//
//	ServiceendpointGcpTerraformArray{ ServiceendpointGcpTerraformArgs{...} }
type ServiceendpointGcpTerraformArrayInput interface {
	pulumi.Input

	ToServiceendpointGcpTerraformArrayOutput() ServiceendpointGcpTerraformArrayOutput
	ToServiceendpointGcpTerraformArrayOutputWithContext(context.Context) ServiceendpointGcpTerraformArrayOutput
}

type ServiceendpointGcpTerraformArray []ServiceendpointGcpTerraformInput

func (ServiceendpointGcpTerraformArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceendpointGcpTerraform)(nil)).Elem()
}

func (i ServiceendpointGcpTerraformArray) ToServiceendpointGcpTerraformArrayOutput() ServiceendpointGcpTerraformArrayOutput {
	return i.ToServiceendpointGcpTerraformArrayOutputWithContext(context.Background())
}

func (i ServiceendpointGcpTerraformArray) ToServiceendpointGcpTerraformArrayOutputWithContext(ctx context.Context) ServiceendpointGcpTerraformArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceendpointGcpTerraformArrayOutput)
}

// ServiceendpointGcpTerraformMapInput is an input type that accepts ServiceendpointGcpTerraformMap and ServiceendpointGcpTerraformMapOutput values.
// You can construct a concrete instance of `ServiceendpointGcpTerraformMapInput` via:
//
//	ServiceendpointGcpTerraformMap{ "key": ServiceendpointGcpTerraformArgs{...} }
type ServiceendpointGcpTerraformMapInput interface {
	pulumi.Input

	ToServiceendpointGcpTerraformMapOutput() ServiceendpointGcpTerraformMapOutput
	ToServiceendpointGcpTerraformMapOutputWithContext(context.Context) ServiceendpointGcpTerraformMapOutput
}

type ServiceendpointGcpTerraformMap map[string]ServiceendpointGcpTerraformInput

func (ServiceendpointGcpTerraformMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceendpointGcpTerraform)(nil)).Elem()
}

func (i ServiceendpointGcpTerraformMap) ToServiceendpointGcpTerraformMapOutput() ServiceendpointGcpTerraformMapOutput {
	return i.ToServiceendpointGcpTerraformMapOutputWithContext(context.Background())
}

func (i ServiceendpointGcpTerraformMap) ToServiceendpointGcpTerraformMapOutputWithContext(ctx context.Context) ServiceendpointGcpTerraformMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceendpointGcpTerraformMapOutput)
}

type ServiceendpointGcpTerraformOutput struct{ *pulumi.OutputState }

func (ServiceendpointGcpTerraformOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceendpointGcpTerraform)(nil)).Elem()
}

func (o ServiceendpointGcpTerraformOutput) ToServiceendpointGcpTerraformOutput() ServiceendpointGcpTerraformOutput {
	return o
}

func (o ServiceendpointGcpTerraformOutput) ToServiceendpointGcpTerraformOutputWithContext(ctx context.Context) ServiceendpointGcpTerraformOutput {
	return o
}

func (o ServiceendpointGcpTerraformOutput) Authorization() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ServiceendpointGcpTerraform) pulumi.StringMapOutput { return v.Authorization }).(pulumi.StringMapOutput)
}

// The client email field in the JSON key file for creating the JSON Web Token.
func (o ServiceendpointGcpTerraformOutput) ClientEmail() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceendpointGcpTerraform) pulumi.StringPtrOutput { return v.ClientEmail }).(pulumi.StringPtrOutput)
}

func (o ServiceendpointGcpTerraformOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceendpointGcpTerraform) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// GCP project associated with the Service Connection.
func (o ServiceendpointGcpTerraformOutput) GcpProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceendpointGcpTerraform) pulumi.StringOutput { return v.GcpProjectId }).(pulumi.StringOutput)
}

// The Private Key for connecting to the endpoint.
func (o ServiceendpointGcpTerraformOutput) PrivateKey() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceendpointGcpTerraform) pulumi.StringOutput { return v.PrivateKey }).(pulumi.StringOutput)
}

// The ID of the project.
func (o ServiceendpointGcpTerraformOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceendpointGcpTerraform) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// Scope to be provided.
func (o ServiceendpointGcpTerraformOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceendpointGcpTerraform) pulumi.StringPtrOutput { return v.Scope }).(pulumi.StringPtrOutput)
}

// The Service Endpoint name.
func (o ServiceendpointGcpTerraformOutput) ServiceEndpointName() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceendpointGcpTerraform) pulumi.StringOutput { return v.ServiceEndpointName }).(pulumi.StringOutput)
}

// The token uri field in the JSON key file for creating the JSON Web Token.
func (o ServiceendpointGcpTerraformOutput) TokenUri() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceendpointGcpTerraform) pulumi.StringOutput { return v.TokenUri }).(pulumi.StringOutput)
}

type ServiceendpointGcpTerraformArrayOutput struct{ *pulumi.OutputState }

func (ServiceendpointGcpTerraformArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceendpointGcpTerraform)(nil)).Elem()
}

func (o ServiceendpointGcpTerraformArrayOutput) ToServiceendpointGcpTerraformArrayOutput() ServiceendpointGcpTerraformArrayOutput {
	return o
}

func (o ServiceendpointGcpTerraformArrayOutput) ToServiceendpointGcpTerraformArrayOutputWithContext(ctx context.Context) ServiceendpointGcpTerraformArrayOutput {
	return o
}

func (o ServiceendpointGcpTerraformArrayOutput) Index(i pulumi.IntInput) ServiceendpointGcpTerraformOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServiceendpointGcpTerraform {
		return vs[0].([]*ServiceendpointGcpTerraform)[vs[1].(int)]
	}).(ServiceendpointGcpTerraformOutput)
}

type ServiceendpointGcpTerraformMapOutput struct{ *pulumi.OutputState }

func (ServiceendpointGcpTerraformMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceendpointGcpTerraform)(nil)).Elem()
}

func (o ServiceendpointGcpTerraformMapOutput) ToServiceendpointGcpTerraformMapOutput() ServiceendpointGcpTerraformMapOutput {
	return o
}

func (o ServiceendpointGcpTerraformMapOutput) ToServiceendpointGcpTerraformMapOutputWithContext(ctx context.Context) ServiceendpointGcpTerraformMapOutput {
	return o
}

func (o ServiceendpointGcpTerraformMapOutput) MapIndex(k pulumi.StringInput) ServiceendpointGcpTerraformOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServiceendpointGcpTerraform {
		return vs[0].(map[string]*ServiceendpointGcpTerraform)[vs[1].(string)]
	}).(ServiceendpointGcpTerraformOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceendpointGcpTerraformInput)(nil)).Elem(), &ServiceendpointGcpTerraform{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceendpointGcpTerraformArrayInput)(nil)).Elem(), ServiceendpointGcpTerraformArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceendpointGcpTerraformMapInput)(nil)).Elem(), ServiceendpointGcpTerraformMap{})
	pulumi.RegisterOutputType(ServiceendpointGcpTerraformOutput{})
	pulumi.RegisterOutputType(ServiceendpointGcpTerraformArrayOutput{})
	pulumi.RegisterOutputType(ServiceendpointGcpTerraformMapOutput{})
}
