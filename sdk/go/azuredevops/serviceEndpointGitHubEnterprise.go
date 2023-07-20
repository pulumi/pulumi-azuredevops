// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azuredevops/sdk/v2/go/azuredevops/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a GitHub Enterprise Server service endpoint within Azure DevOps.
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
//			_, err = azuredevops.NewServiceEndpointGitHubEnterprise(ctx, "exampleServiceEndpointGitHubEnterprise", &azuredevops.ServiceEndpointGitHubEnterpriseArgs{
//				ProjectId:           exampleProject.ID(),
//				ServiceEndpointName: pulumi.String("Example GitHub Enterprise"),
//				Url:                 pulumi.String("https://github.contoso.com"),
//				Description:         pulumi.String("Managed by Terraform"),
//				AuthPersonal: &azuredevops.ServiceEndpointGitHubEnterpriseAuthPersonalArgs{
//					PersonalAccessToken: pulumi.String("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
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
// ## Relevant Links
//
// - [Azure DevOps Service REST API 6.0 - Service Endpoints](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-6.0)
//
// ## Import
//
// Azure DevOps Service Endpoint GitHub Enterprise Server can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**
//
// ```sh
//
//	$ pulumi import azuredevops:index/serviceEndpointGitHubEnterprise:ServiceEndpointGitHubEnterprise example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
//
// ```
type ServiceEndpointGitHubEnterprise struct {
	pulumi.CustomResourceState

	AuthPersonal  ServiceEndpointGitHubEnterpriseAuthPersonalOutput `pulumi:"authPersonal"`
	Authorization pulumi.StringMapOutput                            `pulumi:"authorization"`
	Description   pulumi.StringPtrOutput                            `pulumi:"description"`
	// The ID of the project.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The Service Endpoint name.
	ServiceEndpointName pulumi.StringOutput `pulumi:"serviceEndpointName"`
	// GitHub Enterprise Server Url.
	Url pulumi.StringOutput `pulumi:"url"`
}

// NewServiceEndpointGitHubEnterprise registers a new resource with the given unique name, arguments, and options.
func NewServiceEndpointGitHubEnterprise(ctx *pulumi.Context,
	name string, args *ServiceEndpointGitHubEnterpriseArgs, opts ...pulumi.ResourceOption) (*ServiceEndpointGitHubEnterprise, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AuthPersonal == nil {
		return nil, errors.New("invalid value for required argument 'AuthPersonal'")
	}
	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.ServiceEndpointName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceEndpointName'")
	}
	if args.Url == nil {
		return nil, errors.New("invalid value for required argument 'Url'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ServiceEndpointGitHubEnterprise
	err := ctx.RegisterResource("azuredevops:index/serviceEndpointGitHubEnterprise:ServiceEndpointGitHubEnterprise", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceEndpointGitHubEnterprise gets an existing ServiceEndpointGitHubEnterprise resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceEndpointGitHubEnterprise(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceEndpointGitHubEnterpriseState, opts ...pulumi.ResourceOption) (*ServiceEndpointGitHubEnterprise, error) {
	var resource ServiceEndpointGitHubEnterprise
	err := ctx.ReadResource("azuredevops:index/serviceEndpointGitHubEnterprise:ServiceEndpointGitHubEnterprise", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceEndpointGitHubEnterprise resources.
type serviceEndpointGitHubEnterpriseState struct {
	AuthPersonal  *ServiceEndpointGitHubEnterpriseAuthPersonal `pulumi:"authPersonal"`
	Authorization map[string]string                            `pulumi:"authorization"`
	Description   *string                                      `pulumi:"description"`
	// The ID of the project.
	ProjectId *string `pulumi:"projectId"`
	// The Service Endpoint name.
	ServiceEndpointName *string `pulumi:"serviceEndpointName"`
	// GitHub Enterprise Server Url.
	Url *string `pulumi:"url"`
}

type ServiceEndpointGitHubEnterpriseState struct {
	AuthPersonal  ServiceEndpointGitHubEnterpriseAuthPersonalPtrInput
	Authorization pulumi.StringMapInput
	Description   pulumi.StringPtrInput
	// The ID of the project.
	ProjectId pulumi.StringPtrInput
	// The Service Endpoint name.
	ServiceEndpointName pulumi.StringPtrInput
	// GitHub Enterprise Server Url.
	Url pulumi.StringPtrInput
}

func (ServiceEndpointGitHubEnterpriseState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceEndpointGitHubEnterpriseState)(nil)).Elem()
}

type serviceEndpointGitHubEnterpriseArgs struct {
	AuthPersonal  ServiceEndpointGitHubEnterpriseAuthPersonal `pulumi:"authPersonal"`
	Authorization map[string]string                           `pulumi:"authorization"`
	Description   *string                                     `pulumi:"description"`
	// The ID of the project.
	ProjectId string `pulumi:"projectId"`
	// The Service Endpoint name.
	ServiceEndpointName string `pulumi:"serviceEndpointName"`
	// GitHub Enterprise Server Url.
	Url string `pulumi:"url"`
}

// The set of arguments for constructing a ServiceEndpointGitHubEnterprise resource.
type ServiceEndpointGitHubEnterpriseArgs struct {
	AuthPersonal  ServiceEndpointGitHubEnterpriseAuthPersonalInput
	Authorization pulumi.StringMapInput
	Description   pulumi.StringPtrInput
	// The ID of the project.
	ProjectId pulumi.StringInput
	// The Service Endpoint name.
	ServiceEndpointName pulumi.StringInput
	// GitHub Enterprise Server Url.
	Url pulumi.StringInput
}

func (ServiceEndpointGitHubEnterpriseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceEndpointGitHubEnterpriseArgs)(nil)).Elem()
}

type ServiceEndpointGitHubEnterpriseInput interface {
	pulumi.Input

	ToServiceEndpointGitHubEnterpriseOutput() ServiceEndpointGitHubEnterpriseOutput
	ToServiceEndpointGitHubEnterpriseOutputWithContext(ctx context.Context) ServiceEndpointGitHubEnterpriseOutput
}

func (*ServiceEndpointGitHubEnterprise) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceEndpointGitHubEnterprise)(nil)).Elem()
}

func (i *ServiceEndpointGitHubEnterprise) ToServiceEndpointGitHubEnterpriseOutput() ServiceEndpointGitHubEnterpriseOutput {
	return i.ToServiceEndpointGitHubEnterpriseOutputWithContext(context.Background())
}

func (i *ServiceEndpointGitHubEnterprise) ToServiceEndpointGitHubEnterpriseOutputWithContext(ctx context.Context) ServiceEndpointGitHubEnterpriseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointGitHubEnterpriseOutput)
}

// ServiceEndpointGitHubEnterpriseArrayInput is an input type that accepts ServiceEndpointGitHubEnterpriseArray and ServiceEndpointGitHubEnterpriseArrayOutput values.
// You can construct a concrete instance of `ServiceEndpointGitHubEnterpriseArrayInput` via:
//
//	ServiceEndpointGitHubEnterpriseArray{ ServiceEndpointGitHubEnterpriseArgs{...} }
type ServiceEndpointGitHubEnterpriseArrayInput interface {
	pulumi.Input

	ToServiceEndpointGitHubEnterpriseArrayOutput() ServiceEndpointGitHubEnterpriseArrayOutput
	ToServiceEndpointGitHubEnterpriseArrayOutputWithContext(context.Context) ServiceEndpointGitHubEnterpriseArrayOutput
}

type ServiceEndpointGitHubEnterpriseArray []ServiceEndpointGitHubEnterpriseInput

func (ServiceEndpointGitHubEnterpriseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceEndpointGitHubEnterprise)(nil)).Elem()
}

func (i ServiceEndpointGitHubEnterpriseArray) ToServiceEndpointGitHubEnterpriseArrayOutput() ServiceEndpointGitHubEnterpriseArrayOutput {
	return i.ToServiceEndpointGitHubEnterpriseArrayOutputWithContext(context.Background())
}

func (i ServiceEndpointGitHubEnterpriseArray) ToServiceEndpointGitHubEnterpriseArrayOutputWithContext(ctx context.Context) ServiceEndpointGitHubEnterpriseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointGitHubEnterpriseArrayOutput)
}

// ServiceEndpointGitHubEnterpriseMapInput is an input type that accepts ServiceEndpointGitHubEnterpriseMap and ServiceEndpointGitHubEnterpriseMapOutput values.
// You can construct a concrete instance of `ServiceEndpointGitHubEnterpriseMapInput` via:
//
//	ServiceEndpointGitHubEnterpriseMap{ "key": ServiceEndpointGitHubEnterpriseArgs{...} }
type ServiceEndpointGitHubEnterpriseMapInput interface {
	pulumi.Input

	ToServiceEndpointGitHubEnterpriseMapOutput() ServiceEndpointGitHubEnterpriseMapOutput
	ToServiceEndpointGitHubEnterpriseMapOutputWithContext(context.Context) ServiceEndpointGitHubEnterpriseMapOutput
}

type ServiceEndpointGitHubEnterpriseMap map[string]ServiceEndpointGitHubEnterpriseInput

func (ServiceEndpointGitHubEnterpriseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceEndpointGitHubEnterprise)(nil)).Elem()
}

func (i ServiceEndpointGitHubEnterpriseMap) ToServiceEndpointGitHubEnterpriseMapOutput() ServiceEndpointGitHubEnterpriseMapOutput {
	return i.ToServiceEndpointGitHubEnterpriseMapOutputWithContext(context.Background())
}

func (i ServiceEndpointGitHubEnterpriseMap) ToServiceEndpointGitHubEnterpriseMapOutputWithContext(ctx context.Context) ServiceEndpointGitHubEnterpriseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointGitHubEnterpriseMapOutput)
}

type ServiceEndpointGitHubEnterpriseOutput struct{ *pulumi.OutputState }

func (ServiceEndpointGitHubEnterpriseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceEndpointGitHubEnterprise)(nil)).Elem()
}

func (o ServiceEndpointGitHubEnterpriseOutput) ToServiceEndpointGitHubEnterpriseOutput() ServiceEndpointGitHubEnterpriseOutput {
	return o
}

func (o ServiceEndpointGitHubEnterpriseOutput) ToServiceEndpointGitHubEnterpriseOutputWithContext(ctx context.Context) ServiceEndpointGitHubEnterpriseOutput {
	return o
}

func (o ServiceEndpointGitHubEnterpriseOutput) AuthPersonal() ServiceEndpointGitHubEnterpriseAuthPersonalOutput {
	return o.ApplyT(func(v *ServiceEndpointGitHubEnterprise) ServiceEndpointGitHubEnterpriseAuthPersonalOutput {
		return v.AuthPersonal
	}).(ServiceEndpointGitHubEnterpriseAuthPersonalOutput)
}

func (o ServiceEndpointGitHubEnterpriseOutput) Authorization() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ServiceEndpointGitHubEnterprise) pulumi.StringMapOutput { return v.Authorization }).(pulumi.StringMapOutput)
}

func (o ServiceEndpointGitHubEnterpriseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceEndpointGitHubEnterprise) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The ID of the project.
func (o ServiceEndpointGitHubEnterpriseOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEndpointGitHubEnterprise) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The Service Endpoint name.
func (o ServiceEndpointGitHubEnterpriseOutput) ServiceEndpointName() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEndpointGitHubEnterprise) pulumi.StringOutput { return v.ServiceEndpointName }).(pulumi.StringOutput)
}

// GitHub Enterprise Server Url.
func (o ServiceEndpointGitHubEnterpriseOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEndpointGitHubEnterprise) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

type ServiceEndpointGitHubEnterpriseArrayOutput struct{ *pulumi.OutputState }

func (ServiceEndpointGitHubEnterpriseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceEndpointGitHubEnterprise)(nil)).Elem()
}

func (o ServiceEndpointGitHubEnterpriseArrayOutput) ToServiceEndpointGitHubEnterpriseArrayOutput() ServiceEndpointGitHubEnterpriseArrayOutput {
	return o
}

func (o ServiceEndpointGitHubEnterpriseArrayOutput) ToServiceEndpointGitHubEnterpriseArrayOutputWithContext(ctx context.Context) ServiceEndpointGitHubEnterpriseArrayOutput {
	return o
}

func (o ServiceEndpointGitHubEnterpriseArrayOutput) Index(i pulumi.IntInput) ServiceEndpointGitHubEnterpriseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServiceEndpointGitHubEnterprise {
		return vs[0].([]*ServiceEndpointGitHubEnterprise)[vs[1].(int)]
	}).(ServiceEndpointGitHubEnterpriseOutput)
}

type ServiceEndpointGitHubEnterpriseMapOutput struct{ *pulumi.OutputState }

func (ServiceEndpointGitHubEnterpriseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceEndpointGitHubEnterprise)(nil)).Elem()
}

func (o ServiceEndpointGitHubEnterpriseMapOutput) ToServiceEndpointGitHubEnterpriseMapOutput() ServiceEndpointGitHubEnterpriseMapOutput {
	return o
}

func (o ServiceEndpointGitHubEnterpriseMapOutput) ToServiceEndpointGitHubEnterpriseMapOutputWithContext(ctx context.Context) ServiceEndpointGitHubEnterpriseMapOutput {
	return o
}

func (o ServiceEndpointGitHubEnterpriseMapOutput) MapIndex(k pulumi.StringInput) ServiceEndpointGitHubEnterpriseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServiceEndpointGitHubEnterprise {
		return vs[0].(map[string]*ServiceEndpointGitHubEnterprise)[vs[1].(string)]
	}).(ServiceEndpointGitHubEnterpriseOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceEndpointGitHubEnterpriseInput)(nil)).Elem(), &ServiceEndpointGitHubEnterprise{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceEndpointGitHubEnterpriseArrayInput)(nil)).Elem(), ServiceEndpointGitHubEnterpriseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceEndpointGitHubEnterpriseMapInput)(nil)).Elem(), ServiceEndpointGitHubEnterpriseMap{})
	pulumi.RegisterOutputType(ServiceEndpointGitHubEnterpriseOutput{})
	pulumi.RegisterOutputType(ServiceEndpointGitHubEnterpriseArrayOutput{})
	pulumi.RegisterOutputType(ServiceEndpointGitHubEnterpriseMapOutput{})
}
