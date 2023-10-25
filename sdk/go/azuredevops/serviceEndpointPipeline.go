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

// Manages a Azure DevOps Service Connection service endpoint within Azure DevOps. Allows to run downstream pipelines, monitoring their execution, collecting and consolidating artefacts produced in the delegate pipelines (yaml block `task: RunPipelines@1`). More details on Marketplace page: [RunPipelines](https://marketplace.visualstudio.com/items?itemName=CSE-DevOps.RunPipelines)
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
//			_, err = azuredevops.NewServiceEndpointPipeline(ctx, "exampleServiceEndpointPipeline", &azuredevops.ServiceEndpointPipelineArgs{
//				ProjectId:           exampleProject.ID(),
//				ServiceEndpointName: pulumi.String("Example Pipeline Runner"),
//				OrganizationName:    pulumi.String("Organization Name"),
//				AuthPersonal: &azuredevops.ServiceEndpointPipelineAuthPersonalArgs{
//					PersonalAccessToken: pulumi.String("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
//				},
//				Description: pulumi.String("Managed by Terraform"),
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
// - [Azure DevOps Service REST API 7.0 - Agent Pools](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-7.0)
//
// ## Import
//
// Azure DevOps Service Endpoint can be imported using the `project id`, `service connection id`, e.g.
//
// ```sh
//
//	$ pulumi import azuredevops:index/serviceEndpointPipeline:ServiceEndpointPipeline example projectID/00000000-0000-0000-0000-000000000000
//
// ```
type ServiceEndpointPipeline struct {
	pulumi.CustomResourceState

	// An `authPersonal` block as documented below. Allows connecting using a personal access token.
	AuthPersonal  ServiceEndpointPipelineAuthPersonalOutput `pulumi:"authPersonal"`
	Authorization pulumi.StringMapOutput                    `pulumi:"authorization"`
	Description   pulumi.StringPtrOutput                    `pulumi:"description"`
	// The organization name used for `Organization Url` and `Release API Url` fields.
	OrganizationName pulumi.StringOutput `pulumi:"organizationName"`
	// The ID of the project.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The Service Endpoint name.
	ServiceEndpointName pulumi.StringOutput `pulumi:"serviceEndpointName"`
}

// NewServiceEndpointPipeline registers a new resource with the given unique name, arguments, and options.
func NewServiceEndpointPipeline(ctx *pulumi.Context,
	name string, args *ServiceEndpointPipelineArgs, opts ...pulumi.ResourceOption) (*ServiceEndpointPipeline, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AuthPersonal == nil {
		return nil, errors.New("invalid value for required argument 'AuthPersonal'")
	}
	if args.OrganizationName == nil {
		return nil, errors.New("invalid value for required argument 'OrganizationName'")
	}
	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.ServiceEndpointName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceEndpointName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ServiceEndpointPipeline
	err := ctx.RegisterResource("azuredevops:index/serviceEndpointPipeline:ServiceEndpointPipeline", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceEndpointPipeline gets an existing ServiceEndpointPipeline resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceEndpointPipeline(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceEndpointPipelineState, opts ...pulumi.ResourceOption) (*ServiceEndpointPipeline, error) {
	var resource ServiceEndpointPipeline
	err := ctx.ReadResource("azuredevops:index/serviceEndpointPipeline:ServiceEndpointPipeline", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceEndpointPipeline resources.
type serviceEndpointPipelineState struct {
	// An `authPersonal` block as documented below. Allows connecting using a personal access token.
	AuthPersonal  *ServiceEndpointPipelineAuthPersonal `pulumi:"authPersonal"`
	Authorization map[string]string                    `pulumi:"authorization"`
	Description   *string                              `pulumi:"description"`
	// The organization name used for `Organization Url` and `Release API Url` fields.
	OrganizationName *string `pulumi:"organizationName"`
	// The ID of the project.
	ProjectId *string `pulumi:"projectId"`
	// The Service Endpoint name.
	ServiceEndpointName *string `pulumi:"serviceEndpointName"`
}

type ServiceEndpointPipelineState struct {
	// An `authPersonal` block as documented below. Allows connecting using a personal access token.
	AuthPersonal  ServiceEndpointPipelineAuthPersonalPtrInput
	Authorization pulumi.StringMapInput
	Description   pulumi.StringPtrInput
	// The organization name used for `Organization Url` and `Release API Url` fields.
	OrganizationName pulumi.StringPtrInput
	// The ID of the project.
	ProjectId pulumi.StringPtrInput
	// The Service Endpoint name.
	ServiceEndpointName pulumi.StringPtrInput
}

func (ServiceEndpointPipelineState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceEndpointPipelineState)(nil)).Elem()
}

type serviceEndpointPipelineArgs struct {
	// An `authPersonal` block as documented below. Allows connecting using a personal access token.
	AuthPersonal  ServiceEndpointPipelineAuthPersonal `pulumi:"authPersonal"`
	Authorization map[string]string                   `pulumi:"authorization"`
	Description   *string                             `pulumi:"description"`
	// The organization name used for `Organization Url` and `Release API Url` fields.
	OrganizationName string `pulumi:"organizationName"`
	// The ID of the project.
	ProjectId string `pulumi:"projectId"`
	// The Service Endpoint name.
	ServiceEndpointName string `pulumi:"serviceEndpointName"`
}

// The set of arguments for constructing a ServiceEndpointPipeline resource.
type ServiceEndpointPipelineArgs struct {
	// An `authPersonal` block as documented below. Allows connecting using a personal access token.
	AuthPersonal  ServiceEndpointPipelineAuthPersonalInput
	Authorization pulumi.StringMapInput
	Description   pulumi.StringPtrInput
	// The organization name used for `Organization Url` and `Release API Url` fields.
	OrganizationName pulumi.StringInput
	// The ID of the project.
	ProjectId pulumi.StringInput
	// The Service Endpoint name.
	ServiceEndpointName pulumi.StringInput
}

func (ServiceEndpointPipelineArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceEndpointPipelineArgs)(nil)).Elem()
}

type ServiceEndpointPipelineInput interface {
	pulumi.Input

	ToServiceEndpointPipelineOutput() ServiceEndpointPipelineOutput
	ToServiceEndpointPipelineOutputWithContext(ctx context.Context) ServiceEndpointPipelineOutput
}

func (*ServiceEndpointPipeline) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceEndpointPipeline)(nil)).Elem()
}

func (i *ServiceEndpointPipeline) ToServiceEndpointPipelineOutput() ServiceEndpointPipelineOutput {
	return i.ToServiceEndpointPipelineOutputWithContext(context.Background())
}

func (i *ServiceEndpointPipeline) ToServiceEndpointPipelineOutputWithContext(ctx context.Context) ServiceEndpointPipelineOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointPipelineOutput)
}

func (i *ServiceEndpointPipeline) ToOutput(ctx context.Context) pulumix.Output[*ServiceEndpointPipeline] {
	return pulumix.Output[*ServiceEndpointPipeline]{
		OutputState: i.ToServiceEndpointPipelineOutputWithContext(ctx).OutputState,
	}
}

// ServiceEndpointPipelineArrayInput is an input type that accepts ServiceEndpointPipelineArray and ServiceEndpointPipelineArrayOutput values.
// You can construct a concrete instance of `ServiceEndpointPipelineArrayInput` via:
//
//	ServiceEndpointPipelineArray{ ServiceEndpointPipelineArgs{...} }
type ServiceEndpointPipelineArrayInput interface {
	pulumi.Input

	ToServiceEndpointPipelineArrayOutput() ServiceEndpointPipelineArrayOutput
	ToServiceEndpointPipelineArrayOutputWithContext(context.Context) ServiceEndpointPipelineArrayOutput
}

type ServiceEndpointPipelineArray []ServiceEndpointPipelineInput

func (ServiceEndpointPipelineArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceEndpointPipeline)(nil)).Elem()
}

func (i ServiceEndpointPipelineArray) ToServiceEndpointPipelineArrayOutput() ServiceEndpointPipelineArrayOutput {
	return i.ToServiceEndpointPipelineArrayOutputWithContext(context.Background())
}

func (i ServiceEndpointPipelineArray) ToServiceEndpointPipelineArrayOutputWithContext(ctx context.Context) ServiceEndpointPipelineArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointPipelineArrayOutput)
}

func (i ServiceEndpointPipelineArray) ToOutput(ctx context.Context) pulumix.Output[[]*ServiceEndpointPipeline] {
	return pulumix.Output[[]*ServiceEndpointPipeline]{
		OutputState: i.ToServiceEndpointPipelineArrayOutputWithContext(ctx).OutputState,
	}
}

// ServiceEndpointPipelineMapInput is an input type that accepts ServiceEndpointPipelineMap and ServiceEndpointPipelineMapOutput values.
// You can construct a concrete instance of `ServiceEndpointPipelineMapInput` via:
//
//	ServiceEndpointPipelineMap{ "key": ServiceEndpointPipelineArgs{...} }
type ServiceEndpointPipelineMapInput interface {
	pulumi.Input

	ToServiceEndpointPipelineMapOutput() ServiceEndpointPipelineMapOutput
	ToServiceEndpointPipelineMapOutputWithContext(context.Context) ServiceEndpointPipelineMapOutput
}

type ServiceEndpointPipelineMap map[string]ServiceEndpointPipelineInput

func (ServiceEndpointPipelineMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceEndpointPipeline)(nil)).Elem()
}

func (i ServiceEndpointPipelineMap) ToServiceEndpointPipelineMapOutput() ServiceEndpointPipelineMapOutput {
	return i.ToServiceEndpointPipelineMapOutputWithContext(context.Background())
}

func (i ServiceEndpointPipelineMap) ToServiceEndpointPipelineMapOutputWithContext(ctx context.Context) ServiceEndpointPipelineMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointPipelineMapOutput)
}

func (i ServiceEndpointPipelineMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*ServiceEndpointPipeline] {
	return pulumix.Output[map[string]*ServiceEndpointPipeline]{
		OutputState: i.ToServiceEndpointPipelineMapOutputWithContext(ctx).OutputState,
	}
}

type ServiceEndpointPipelineOutput struct{ *pulumi.OutputState }

func (ServiceEndpointPipelineOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceEndpointPipeline)(nil)).Elem()
}

func (o ServiceEndpointPipelineOutput) ToServiceEndpointPipelineOutput() ServiceEndpointPipelineOutput {
	return o
}

func (o ServiceEndpointPipelineOutput) ToServiceEndpointPipelineOutputWithContext(ctx context.Context) ServiceEndpointPipelineOutput {
	return o
}

func (o ServiceEndpointPipelineOutput) ToOutput(ctx context.Context) pulumix.Output[*ServiceEndpointPipeline] {
	return pulumix.Output[*ServiceEndpointPipeline]{
		OutputState: o.OutputState,
	}
}

// An `authPersonal` block as documented below. Allows connecting using a personal access token.
func (o ServiceEndpointPipelineOutput) AuthPersonal() ServiceEndpointPipelineAuthPersonalOutput {
	return o.ApplyT(func(v *ServiceEndpointPipeline) ServiceEndpointPipelineAuthPersonalOutput { return v.AuthPersonal }).(ServiceEndpointPipelineAuthPersonalOutput)
}

func (o ServiceEndpointPipelineOutput) Authorization() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ServiceEndpointPipeline) pulumi.StringMapOutput { return v.Authorization }).(pulumi.StringMapOutput)
}

func (o ServiceEndpointPipelineOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceEndpointPipeline) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The organization name used for `Organization Url` and `Release API Url` fields.
func (o ServiceEndpointPipelineOutput) OrganizationName() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEndpointPipeline) pulumi.StringOutput { return v.OrganizationName }).(pulumi.StringOutput)
}

// The ID of the project.
func (o ServiceEndpointPipelineOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEndpointPipeline) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The Service Endpoint name.
func (o ServiceEndpointPipelineOutput) ServiceEndpointName() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEndpointPipeline) pulumi.StringOutput { return v.ServiceEndpointName }).(pulumi.StringOutput)
}

type ServiceEndpointPipelineArrayOutput struct{ *pulumi.OutputState }

func (ServiceEndpointPipelineArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceEndpointPipeline)(nil)).Elem()
}

func (o ServiceEndpointPipelineArrayOutput) ToServiceEndpointPipelineArrayOutput() ServiceEndpointPipelineArrayOutput {
	return o
}

func (o ServiceEndpointPipelineArrayOutput) ToServiceEndpointPipelineArrayOutputWithContext(ctx context.Context) ServiceEndpointPipelineArrayOutput {
	return o
}

func (o ServiceEndpointPipelineArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*ServiceEndpointPipeline] {
	return pulumix.Output[[]*ServiceEndpointPipeline]{
		OutputState: o.OutputState,
	}
}

func (o ServiceEndpointPipelineArrayOutput) Index(i pulumi.IntInput) ServiceEndpointPipelineOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServiceEndpointPipeline {
		return vs[0].([]*ServiceEndpointPipeline)[vs[1].(int)]
	}).(ServiceEndpointPipelineOutput)
}

type ServiceEndpointPipelineMapOutput struct{ *pulumi.OutputState }

func (ServiceEndpointPipelineMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceEndpointPipeline)(nil)).Elem()
}

func (o ServiceEndpointPipelineMapOutput) ToServiceEndpointPipelineMapOutput() ServiceEndpointPipelineMapOutput {
	return o
}

func (o ServiceEndpointPipelineMapOutput) ToServiceEndpointPipelineMapOutputWithContext(ctx context.Context) ServiceEndpointPipelineMapOutput {
	return o
}

func (o ServiceEndpointPipelineMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*ServiceEndpointPipeline] {
	return pulumix.Output[map[string]*ServiceEndpointPipeline]{
		OutputState: o.OutputState,
	}
}

func (o ServiceEndpointPipelineMapOutput) MapIndex(k pulumi.StringInput) ServiceEndpointPipelineOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServiceEndpointPipeline {
		return vs[0].(map[string]*ServiceEndpointPipeline)[vs[1].(string)]
	}).(ServiceEndpointPipelineOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceEndpointPipelineInput)(nil)).Elem(), &ServiceEndpointPipeline{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceEndpointPipelineArrayInput)(nil)).Elem(), ServiceEndpointPipelineArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceEndpointPipelineMapInput)(nil)).Elem(), ServiceEndpointPipelineMap{})
	pulumi.RegisterOutputType(ServiceEndpointPipelineOutput{})
	pulumi.RegisterOutputType(ServiceEndpointPipelineArrayOutput{})
	pulumi.RegisterOutputType(ServiceEndpointPipelineMapOutput{})
}
