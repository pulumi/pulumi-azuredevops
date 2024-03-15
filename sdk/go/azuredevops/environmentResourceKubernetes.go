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

type EnvironmentResourceKubernetes struct {
	pulumi.CustomResourceState

	ClusterName       pulumi.StringPtrOutput   `pulumi:"clusterName"`
	EnvironmentId     pulumi.IntOutput         `pulumi:"environmentId"`
	Name              pulumi.StringOutput      `pulumi:"name"`
	Namespace         pulumi.StringOutput      `pulumi:"namespace"`
	ProjectId         pulumi.StringOutput      `pulumi:"projectId"`
	ServiceEndpointId pulumi.StringOutput      `pulumi:"serviceEndpointId"`
	Tags              pulumi.StringArrayOutput `pulumi:"tags"`
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
	ClusterName       *string  `pulumi:"clusterName"`
	EnvironmentId     *int     `pulumi:"environmentId"`
	Name              *string  `pulumi:"name"`
	Namespace         *string  `pulumi:"namespace"`
	ProjectId         *string  `pulumi:"projectId"`
	ServiceEndpointId *string  `pulumi:"serviceEndpointId"`
	Tags              []string `pulumi:"tags"`
}

type EnvironmentResourceKubernetesState struct {
	ClusterName       pulumi.StringPtrInput
	EnvironmentId     pulumi.IntPtrInput
	Name              pulumi.StringPtrInput
	Namespace         pulumi.StringPtrInput
	ProjectId         pulumi.StringPtrInput
	ServiceEndpointId pulumi.StringPtrInput
	Tags              pulumi.StringArrayInput
}

func (EnvironmentResourceKubernetesState) ElementType() reflect.Type {
	return reflect.TypeOf((*environmentResourceKubernetesState)(nil)).Elem()
}

type environmentResourceKubernetesArgs struct {
	ClusterName       *string  `pulumi:"clusterName"`
	EnvironmentId     int      `pulumi:"environmentId"`
	Name              *string  `pulumi:"name"`
	Namespace         string   `pulumi:"namespace"`
	ProjectId         string   `pulumi:"projectId"`
	ServiceEndpointId string   `pulumi:"serviceEndpointId"`
	Tags              []string `pulumi:"tags"`
}

// The set of arguments for constructing a EnvironmentResourceKubernetes resource.
type EnvironmentResourceKubernetesArgs struct {
	ClusterName       pulumi.StringPtrInput
	EnvironmentId     pulumi.IntInput
	Name              pulumi.StringPtrInput
	Namespace         pulumi.StringInput
	ProjectId         pulumi.StringInput
	ServiceEndpointId pulumi.StringInput
	Tags              pulumi.StringArrayInput
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

func (o EnvironmentResourceKubernetesOutput) ClusterName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnvironmentResourceKubernetes) pulumi.StringPtrOutput { return v.ClusterName }).(pulumi.StringPtrOutput)
}

func (o EnvironmentResourceKubernetesOutput) EnvironmentId() pulumi.IntOutput {
	return o.ApplyT(func(v *EnvironmentResourceKubernetes) pulumi.IntOutput { return v.EnvironmentId }).(pulumi.IntOutput)
}

func (o EnvironmentResourceKubernetesOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *EnvironmentResourceKubernetes) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o EnvironmentResourceKubernetesOutput) Namespace() pulumi.StringOutput {
	return o.ApplyT(func(v *EnvironmentResourceKubernetes) pulumi.StringOutput { return v.Namespace }).(pulumi.StringOutput)
}

func (o EnvironmentResourceKubernetesOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *EnvironmentResourceKubernetes) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

func (o EnvironmentResourceKubernetesOutput) ServiceEndpointId() pulumi.StringOutput {
	return o.ApplyT(func(v *EnvironmentResourceKubernetes) pulumi.StringOutput { return v.ServiceEndpointId }).(pulumi.StringOutput)
}

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
