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

// Manages a Kubernetes Resource for an Environment.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
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
//			exampleEnvironment, err := azuredevops.NewEnvironment(ctx, "example", &azuredevops.EnvironmentArgs{
//				ProjectId: example.ID(),
//				Name:      pulumi.String("Example Environment"),
//			})
//			if err != nil {
//				return err
//			}
//			exampleServiceEndpointKubernetes, err := azuredevops.NewServiceEndpointKubernetes(ctx, "example", &azuredevops.ServiceEndpointKubernetesArgs{
//				ProjectId:           example.ID(),
//				ServiceEndpointName: pulumi.String("Example Kubernetes"),
//				ApiserverUrl:        pulumi.String("https://sample-kubernetes-cluster.hcp.westeurope.azmk8s.io"),
//				AuthorizationType:   pulumi.String("AzureSubscription"),
//				AzureSubscriptions: azuredevops.ServiceEndpointKubernetesAzureSubscriptionArray{
//					&azuredevops.ServiceEndpointKubernetesAzureSubscriptionArgs{
//						SubscriptionId:   pulumi.String("00000000-0000-0000-0000-000000000000"),
//						SubscriptionName: pulumi.String("Example"),
//						TenantId:         pulumi.String("00000000-0000-0000-0000-000000000000"),
//						ResourcegroupId:  pulumi.String("example-rg"),
//						Namespace:        pulumi.String("default"),
//						ClusterName:      pulumi.String("example-aks"),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = azuredevops.NewEnvironmentResourceKubernetes(ctx, "example", &azuredevops.EnvironmentResourceKubernetesArgs{
//				ProjectId:         example.ID(),
//				EnvironmentId:     exampleEnvironment.ID(),
//				ServiceEndpointId: exampleServiceEndpointKubernetes.ID(),
//				Name:              pulumi.String("Example"),
//				Namespace:         pulumi.String("default"),
//				ClusterName:       pulumi.String("example-aks"),
//				Tags: pulumi.StringArray{
//					pulumi.String("tag1"),
//					pulumi.String("tag2"),
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
// <!--End PulumiCodeChooser -->
//
// ## Relevant Links
//
// * [Azure DevOps Service REST API 6.0 - Kubernetes](https://docs.microsoft.com/en-us/rest/api/azure/devops/distributedtask/kubernetes?view=azure-devops-rest-6.0)
//
// ## Import
//
// The resource does not support import.
type EnvironmentResourceKubernetes struct {
	pulumi.CustomResourceState

	// A cluster name for the Kubernetes Resource.
	ClusterName pulumi.StringPtrOutput `pulumi:"clusterName"`
	// The ID of the environment under which to create the Kubernetes Resource.
	EnvironmentId pulumi.IntOutput `pulumi:"environmentId"`
	// The name for the Kubernetes Resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The namespace for the Kubernetes Resource.
	Namespace pulumi.StringOutput `pulumi:"namespace"`
	// The ID of the project.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The ID of the service endpoint to associate with the Kubernetes Resource.
	ServiceEndpointId pulumi.StringOutput `pulumi:"serviceEndpointId"`
	// A set of tags for the Kubernetes Resource.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
}

// NewEnvironmentResourceKubernetes registers a new resource with the given unique name, arguments, and options.
func NewEnvironmentResourceKubernetes(ctx *pulumi.Context,
	name string, args *EnvironmentResourceKubernetesArgs, opts ...pulumi.ResourceOption) (*EnvironmentResourceKubernetes, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EnvironmentId == nil {
		return nil, errors.New("invalid value for required argument 'EnvironmentId'")
	}
	if args.Namespace == nil {
		return nil, errors.New("invalid value for required argument 'Namespace'")
	}
	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.ServiceEndpointId == nil {
		return nil, errors.New("invalid value for required argument 'ServiceEndpointId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource EnvironmentResourceKubernetes
	err := ctx.RegisterResource("azuredevops:index/environmentResourceKubernetes:EnvironmentResourceKubernetes", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEnvironmentResourceKubernetes gets an existing EnvironmentResourceKubernetes resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEnvironmentResourceKubernetes(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EnvironmentResourceKubernetesState, opts ...pulumi.ResourceOption) (*EnvironmentResourceKubernetes, error) {
	var resource EnvironmentResourceKubernetes
	err := ctx.ReadResource("azuredevops:index/environmentResourceKubernetes:EnvironmentResourceKubernetes", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EnvironmentResourceKubernetes resources.
type environmentResourceKubernetesState struct {
	// A cluster name for the Kubernetes Resource.
	ClusterName *string `pulumi:"clusterName"`
	// The ID of the environment under which to create the Kubernetes Resource.
	EnvironmentId *int `pulumi:"environmentId"`
	// The name for the Kubernetes Resource.
	Name *string `pulumi:"name"`
	// The namespace for the Kubernetes Resource.
	Namespace *string `pulumi:"namespace"`
	// The ID of the project.
	ProjectId *string `pulumi:"projectId"`
	// The ID of the service endpoint to associate with the Kubernetes Resource.
	ServiceEndpointId *string `pulumi:"serviceEndpointId"`
	// A set of tags for the Kubernetes Resource.
	Tags []string `pulumi:"tags"`
}

type EnvironmentResourceKubernetesState struct {
	// A cluster name for the Kubernetes Resource.
	ClusterName pulumi.StringPtrInput
	// The ID of the environment under which to create the Kubernetes Resource.
	EnvironmentId pulumi.IntPtrInput
	// The name for the Kubernetes Resource.
	Name pulumi.StringPtrInput
	// The namespace for the Kubernetes Resource.
	Namespace pulumi.StringPtrInput
	// The ID of the project.
	ProjectId pulumi.StringPtrInput
	// The ID of the service endpoint to associate with the Kubernetes Resource.
	ServiceEndpointId pulumi.StringPtrInput
	// A set of tags for the Kubernetes Resource.
	Tags pulumi.StringArrayInput
}

func (EnvironmentResourceKubernetesState) ElementType() reflect.Type {
	return reflect.TypeOf((*environmentResourceKubernetesState)(nil)).Elem()
}

type environmentResourceKubernetesArgs struct {
	// A cluster name for the Kubernetes Resource.
	ClusterName *string `pulumi:"clusterName"`
	// The ID of the environment under which to create the Kubernetes Resource.
	EnvironmentId int `pulumi:"environmentId"`
	// The name for the Kubernetes Resource.
	Name *string `pulumi:"name"`
	// The namespace for the Kubernetes Resource.
	Namespace string `pulumi:"namespace"`
	// The ID of the project.
	ProjectId string `pulumi:"projectId"`
	// The ID of the service endpoint to associate with the Kubernetes Resource.
	ServiceEndpointId string `pulumi:"serviceEndpointId"`
	// A set of tags for the Kubernetes Resource.
	Tags []string `pulumi:"tags"`
}

// The set of arguments for constructing a EnvironmentResourceKubernetes resource.
type EnvironmentResourceKubernetesArgs struct {
	// A cluster name for the Kubernetes Resource.
	ClusterName pulumi.StringPtrInput
	// The ID of the environment under which to create the Kubernetes Resource.
	EnvironmentId pulumi.IntInput
	// The name for the Kubernetes Resource.
	Name pulumi.StringPtrInput
	// The namespace for the Kubernetes Resource.
	Namespace pulumi.StringInput
	// The ID of the project.
	ProjectId pulumi.StringInput
	// The ID of the service endpoint to associate with the Kubernetes Resource.
	ServiceEndpointId pulumi.StringInput
	// A set of tags for the Kubernetes Resource.
	Tags pulumi.StringArrayInput
}

func (EnvironmentResourceKubernetesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*environmentResourceKubernetesArgs)(nil)).Elem()
}

type EnvironmentResourceKubernetesInput interface {
	pulumi.Input

	ToEnvironmentResourceKubernetesOutput() EnvironmentResourceKubernetesOutput
	ToEnvironmentResourceKubernetesOutputWithContext(ctx context.Context) EnvironmentResourceKubernetesOutput
}

func (*EnvironmentResourceKubernetes) ElementType() reflect.Type {
	return reflect.TypeOf((**EnvironmentResourceKubernetes)(nil)).Elem()
}

func (i *EnvironmentResourceKubernetes) ToEnvironmentResourceKubernetesOutput() EnvironmentResourceKubernetesOutput {
	return i.ToEnvironmentResourceKubernetesOutputWithContext(context.Background())
}

func (i *EnvironmentResourceKubernetes) ToEnvironmentResourceKubernetesOutputWithContext(ctx context.Context) EnvironmentResourceKubernetesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentResourceKubernetesOutput)
}

// EnvironmentResourceKubernetesArrayInput is an input type that accepts EnvironmentResourceKubernetesArray and EnvironmentResourceKubernetesArrayOutput values.
// You can construct a concrete instance of `EnvironmentResourceKubernetesArrayInput` via:
//
//	EnvironmentResourceKubernetesArray{ EnvironmentResourceKubernetesArgs{...} }
type EnvironmentResourceKubernetesArrayInput interface {
	pulumi.Input

	ToEnvironmentResourceKubernetesArrayOutput() EnvironmentResourceKubernetesArrayOutput
	ToEnvironmentResourceKubernetesArrayOutputWithContext(context.Context) EnvironmentResourceKubernetesArrayOutput
}

type EnvironmentResourceKubernetesArray []EnvironmentResourceKubernetesInput

func (EnvironmentResourceKubernetesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EnvironmentResourceKubernetes)(nil)).Elem()
}

func (i EnvironmentResourceKubernetesArray) ToEnvironmentResourceKubernetesArrayOutput() EnvironmentResourceKubernetesArrayOutput {
	return i.ToEnvironmentResourceKubernetesArrayOutputWithContext(context.Background())
}

func (i EnvironmentResourceKubernetesArray) ToEnvironmentResourceKubernetesArrayOutputWithContext(ctx context.Context) EnvironmentResourceKubernetesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentResourceKubernetesArrayOutput)
}

// EnvironmentResourceKubernetesMapInput is an input type that accepts EnvironmentResourceKubernetesMap and EnvironmentResourceKubernetesMapOutput values.
// You can construct a concrete instance of `EnvironmentResourceKubernetesMapInput` via:
//
//	EnvironmentResourceKubernetesMap{ "key": EnvironmentResourceKubernetesArgs{...} }
type EnvironmentResourceKubernetesMapInput interface {
	pulumi.Input

	ToEnvironmentResourceKubernetesMapOutput() EnvironmentResourceKubernetesMapOutput
	ToEnvironmentResourceKubernetesMapOutputWithContext(context.Context) EnvironmentResourceKubernetesMapOutput
}

type EnvironmentResourceKubernetesMap map[string]EnvironmentResourceKubernetesInput

func (EnvironmentResourceKubernetesMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EnvironmentResourceKubernetes)(nil)).Elem()
}

func (i EnvironmentResourceKubernetesMap) ToEnvironmentResourceKubernetesMapOutput() EnvironmentResourceKubernetesMapOutput {
	return i.ToEnvironmentResourceKubernetesMapOutputWithContext(context.Background())
}

func (i EnvironmentResourceKubernetesMap) ToEnvironmentResourceKubernetesMapOutputWithContext(ctx context.Context) EnvironmentResourceKubernetesMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentResourceKubernetesMapOutput)
}

type EnvironmentResourceKubernetesOutput struct{ *pulumi.OutputState }

func (EnvironmentResourceKubernetesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnvironmentResourceKubernetes)(nil)).Elem()
}

func (o EnvironmentResourceKubernetesOutput) ToEnvironmentResourceKubernetesOutput() EnvironmentResourceKubernetesOutput {
	return o
}

func (o EnvironmentResourceKubernetesOutput) ToEnvironmentResourceKubernetesOutputWithContext(ctx context.Context) EnvironmentResourceKubernetesOutput {
	return o
}

// A cluster name for the Kubernetes Resource.
func (o EnvironmentResourceKubernetesOutput) ClusterName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnvironmentResourceKubernetes) pulumi.StringPtrOutput { return v.ClusterName }).(pulumi.StringPtrOutput)
}

// The ID of the environment under which to create the Kubernetes Resource.
func (o EnvironmentResourceKubernetesOutput) EnvironmentId() pulumi.IntOutput {
	return o.ApplyT(func(v *EnvironmentResourceKubernetes) pulumi.IntOutput { return v.EnvironmentId }).(pulumi.IntOutput)
}

// The name for the Kubernetes Resource.
func (o EnvironmentResourceKubernetesOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *EnvironmentResourceKubernetes) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The namespace for the Kubernetes Resource.
func (o EnvironmentResourceKubernetesOutput) Namespace() pulumi.StringOutput {
	return o.ApplyT(func(v *EnvironmentResourceKubernetes) pulumi.StringOutput { return v.Namespace }).(pulumi.StringOutput)
}

// The ID of the project.
func (o EnvironmentResourceKubernetesOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *EnvironmentResourceKubernetes) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The ID of the service endpoint to associate with the Kubernetes Resource.
func (o EnvironmentResourceKubernetesOutput) ServiceEndpointId() pulumi.StringOutput {
	return o.ApplyT(func(v *EnvironmentResourceKubernetes) pulumi.StringOutput { return v.ServiceEndpointId }).(pulumi.StringOutput)
}

// A set of tags for the Kubernetes Resource.
func (o EnvironmentResourceKubernetesOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *EnvironmentResourceKubernetes) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

type EnvironmentResourceKubernetesArrayOutput struct{ *pulumi.OutputState }

func (EnvironmentResourceKubernetesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EnvironmentResourceKubernetes)(nil)).Elem()
}

func (o EnvironmentResourceKubernetesArrayOutput) ToEnvironmentResourceKubernetesArrayOutput() EnvironmentResourceKubernetesArrayOutput {
	return o
}

func (o EnvironmentResourceKubernetesArrayOutput) ToEnvironmentResourceKubernetesArrayOutputWithContext(ctx context.Context) EnvironmentResourceKubernetesArrayOutput {
	return o
}

func (o EnvironmentResourceKubernetesArrayOutput) Index(i pulumi.IntInput) EnvironmentResourceKubernetesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EnvironmentResourceKubernetes {
		return vs[0].([]*EnvironmentResourceKubernetes)[vs[1].(int)]
	}).(EnvironmentResourceKubernetesOutput)
}

type EnvironmentResourceKubernetesMapOutput struct{ *pulumi.OutputState }

func (EnvironmentResourceKubernetesMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EnvironmentResourceKubernetes)(nil)).Elem()
}

func (o EnvironmentResourceKubernetesMapOutput) ToEnvironmentResourceKubernetesMapOutput() EnvironmentResourceKubernetesMapOutput {
	return o
}

func (o EnvironmentResourceKubernetesMapOutput) ToEnvironmentResourceKubernetesMapOutputWithContext(ctx context.Context) EnvironmentResourceKubernetesMapOutput {
	return o
}

func (o EnvironmentResourceKubernetesMapOutput) MapIndex(k pulumi.StringInput) EnvironmentResourceKubernetesOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EnvironmentResourceKubernetes {
		return vs[0].(map[string]*EnvironmentResourceKubernetes)[vs[1].(string)]
	}).(EnvironmentResourceKubernetesOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EnvironmentResourceKubernetesInput)(nil)).Elem(), &EnvironmentResourceKubernetes{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnvironmentResourceKubernetesArrayInput)(nil)).Elem(), EnvironmentResourceKubernetesArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnvironmentResourceKubernetesMapInput)(nil)).Elem(), EnvironmentResourceKubernetesMap{})
	pulumi.RegisterOutputType(EnvironmentResourceKubernetesOutput{})
	pulumi.RegisterOutputType(EnvironmentResourceKubernetesArrayOutput{})
	pulumi.RegisterOutputType(EnvironmentResourceKubernetesMapOutput{})
}
