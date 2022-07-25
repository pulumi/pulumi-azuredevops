// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Kubernetes service endpoint within Azure DevOps.
//
// ## Relevant Links
//
// - [Azure DevOps Service REST API 6.0 - Endpoints](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-6.0)
//
// ## Import
//
// Azure DevOps Service Endpoint Kubernetes can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**
//
// ```sh
//  $ pulumi import azuredevops:index/serviceEndpointKubernetes:ServiceEndpointKubernetes example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
// ```
type ServiceEndpointKubernetes struct {
	pulumi.CustomResourceState

	// The hostname (in form of URI) of the Kubernetes API.
	ApiserverUrl  pulumi.StringOutput    `pulumi:"apiserverUrl"`
	Authorization pulumi.StringMapOutput `pulumi:"authorization"`
	// The authentication method used to authenticate on the Kubernetes cluster. The value should be one of AzureSubscription, Kubeconfig, ServiceAccount.
	AuthorizationType pulumi.StringOutput `pulumi:"authorizationType"`
	// A `azureSubscription` block defined blow.
	AzureSubscriptions ServiceEndpointKubernetesAzureSubscriptionArrayOutput `pulumi:"azureSubscriptions"`
	Description        pulumi.StringPtrOutput                                `pulumi:"description"`
	// A `kubeconfig` block defined blow.
	Kubeconfigs ServiceEndpointKubernetesKubeconfigArrayOutput `pulumi:"kubeconfigs"`
	// The ID of the project.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// A `serviceAccount` block defined blow.
	ServiceAccount ServiceEndpointKubernetesServiceAccountPtrOutput `pulumi:"serviceAccount"`
	// The Service Endpoint name.
	ServiceEndpointName pulumi.StringOutput `pulumi:"serviceEndpointName"`
}

// NewServiceEndpointKubernetes registers a new resource with the given unique name, arguments, and options.
func NewServiceEndpointKubernetes(ctx *pulumi.Context,
	name string, args *ServiceEndpointKubernetesArgs, opts ...pulumi.ResourceOption) (*ServiceEndpointKubernetes, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiserverUrl == nil {
		return nil, errors.New("invalid value for required argument 'ApiserverUrl'")
	}
	if args.AuthorizationType == nil {
		return nil, errors.New("invalid value for required argument 'AuthorizationType'")
	}
	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.ServiceEndpointName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceEndpointName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azuredevops:ServiceEndpoint/kubernetes:Kubernetes"),
		},
	})
	opts = append(opts, aliases)
	var resource ServiceEndpointKubernetes
	err := ctx.RegisterResource("azuredevops:index/serviceEndpointKubernetes:ServiceEndpointKubernetes", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceEndpointKubernetes gets an existing ServiceEndpointKubernetes resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceEndpointKubernetes(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceEndpointKubernetesState, opts ...pulumi.ResourceOption) (*ServiceEndpointKubernetes, error) {
	var resource ServiceEndpointKubernetes
	err := ctx.ReadResource("azuredevops:index/serviceEndpointKubernetes:ServiceEndpointKubernetes", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceEndpointKubernetes resources.
type serviceEndpointKubernetesState struct {
	// The hostname (in form of URI) of the Kubernetes API.
	ApiserverUrl  *string           `pulumi:"apiserverUrl"`
	Authorization map[string]string `pulumi:"authorization"`
	// The authentication method used to authenticate on the Kubernetes cluster. The value should be one of AzureSubscription, Kubeconfig, ServiceAccount.
	AuthorizationType *string `pulumi:"authorizationType"`
	// A `azureSubscription` block defined blow.
	AzureSubscriptions []ServiceEndpointKubernetesAzureSubscription `pulumi:"azureSubscriptions"`
	Description        *string                                      `pulumi:"description"`
	// A `kubeconfig` block defined blow.
	Kubeconfigs []ServiceEndpointKubernetesKubeconfig `pulumi:"kubeconfigs"`
	// The ID of the project.
	ProjectId *string `pulumi:"projectId"`
	// A `serviceAccount` block defined blow.
	ServiceAccount *ServiceEndpointKubernetesServiceAccount `pulumi:"serviceAccount"`
	// The Service Endpoint name.
	ServiceEndpointName *string `pulumi:"serviceEndpointName"`
}

type ServiceEndpointKubernetesState struct {
	// The hostname (in form of URI) of the Kubernetes API.
	ApiserverUrl  pulumi.StringPtrInput
	Authorization pulumi.StringMapInput
	// The authentication method used to authenticate on the Kubernetes cluster. The value should be one of AzureSubscription, Kubeconfig, ServiceAccount.
	AuthorizationType pulumi.StringPtrInput
	// A `azureSubscription` block defined blow.
	AzureSubscriptions ServiceEndpointKubernetesAzureSubscriptionArrayInput
	Description        pulumi.StringPtrInput
	// A `kubeconfig` block defined blow.
	Kubeconfigs ServiceEndpointKubernetesKubeconfigArrayInput
	// The ID of the project.
	ProjectId pulumi.StringPtrInput
	// A `serviceAccount` block defined blow.
	ServiceAccount ServiceEndpointKubernetesServiceAccountPtrInput
	// The Service Endpoint name.
	ServiceEndpointName pulumi.StringPtrInput
}

func (ServiceEndpointKubernetesState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceEndpointKubernetesState)(nil)).Elem()
}

type serviceEndpointKubernetesArgs struct {
	// The hostname (in form of URI) of the Kubernetes API.
	ApiserverUrl  string            `pulumi:"apiserverUrl"`
	Authorization map[string]string `pulumi:"authorization"`
	// The authentication method used to authenticate on the Kubernetes cluster. The value should be one of AzureSubscription, Kubeconfig, ServiceAccount.
	AuthorizationType string `pulumi:"authorizationType"`
	// A `azureSubscription` block defined blow.
	AzureSubscriptions []ServiceEndpointKubernetesAzureSubscription `pulumi:"azureSubscriptions"`
	Description        *string                                      `pulumi:"description"`
	// A `kubeconfig` block defined blow.
	Kubeconfigs []ServiceEndpointKubernetesKubeconfig `pulumi:"kubeconfigs"`
	// The ID of the project.
	ProjectId string `pulumi:"projectId"`
	// A `serviceAccount` block defined blow.
	ServiceAccount *ServiceEndpointKubernetesServiceAccount `pulumi:"serviceAccount"`
	// The Service Endpoint name.
	ServiceEndpointName string `pulumi:"serviceEndpointName"`
}

// The set of arguments for constructing a ServiceEndpointKubernetes resource.
type ServiceEndpointKubernetesArgs struct {
	// The hostname (in form of URI) of the Kubernetes API.
	ApiserverUrl  pulumi.StringInput
	Authorization pulumi.StringMapInput
	// The authentication method used to authenticate on the Kubernetes cluster. The value should be one of AzureSubscription, Kubeconfig, ServiceAccount.
	AuthorizationType pulumi.StringInput
	// A `azureSubscription` block defined blow.
	AzureSubscriptions ServiceEndpointKubernetesAzureSubscriptionArrayInput
	Description        pulumi.StringPtrInput
	// A `kubeconfig` block defined blow.
	Kubeconfigs ServiceEndpointKubernetesKubeconfigArrayInput
	// The ID of the project.
	ProjectId pulumi.StringInput
	// A `serviceAccount` block defined blow.
	ServiceAccount ServiceEndpointKubernetesServiceAccountPtrInput
	// The Service Endpoint name.
	ServiceEndpointName pulumi.StringInput
}

func (ServiceEndpointKubernetesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceEndpointKubernetesArgs)(nil)).Elem()
}

type ServiceEndpointKubernetesInput interface {
	pulumi.Input

	ToServiceEndpointKubernetesOutput() ServiceEndpointKubernetesOutput
	ToServiceEndpointKubernetesOutputWithContext(ctx context.Context) ServiceEndpointKubernetesOutput
}

func (*ServiceEndpointKubernetes) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceEndpointKubernetes)(nil)).Elem()
}

func (i *ServiceEndpointKubernetes) ToServiceEndpointKubernetesOutput() ServiceEndpointKubernetesOutput {
	return i.ToServiceEndpointKubernetesOutputWithContext(context.Background())
}

func (i *ServiceEndpointKubernetes) ToServiceEndpointKubernetesOutputWithContext(ctx context.Context) ServiceEndpointKubernetesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointKubernetesOutput)
}

// ServiceEndpointKubernetesArrayInput is an input type that accepts ServiceEndpointKubernetesArray and ServiceEndpointKubernetesArrayOutput values.
// You can construct a concrete instance of `ServiceEndpointKubernetesArrayInput` via:
//
//          ServiceEndpointKubernetesArray{ ServiceEndpointKubernetesArgs{...} }
type ServiceEndpointKubernetesArrayInput interface {
	pulumi.Input

	ToServiceEndpointKubernetesArrayOutput() ServiceEndpointKubernetesArrayOutput
	ToServiceEndpointKubernetesArrayOutputWithContext(context.Context) ServiceEndpointKubernetesArrayOutput
}

type ServiceEndpointKubernetesArray []ServiceEndpointKubernetesInput

func (ServiceEndpointKubernetesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceEndpointKubernetes)(nil)).Elem()
}

func (i ServiceEndpointKubernetesArray) ToServiceEndpointKubernetesArrayOutput() ServiceEndpointKubernetesArrayOutput {
	return i.ToServiceEndpointKubernetesArrayOutputWithContext(context.Background())
}

func (i ServiceEndpointKubernetesArray) ToServiceEndpointKubernetesArrayOutputWithContext(ctx context.Context) ServiceEndpointKubernetesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointKubernetesArrayOutput)
}

// ServiceEndpointKubernetesMapInput is an input type that accepts ServiceEndpointKubernetesMap and ServiceEndpointKubernetesMapOutput values.
// You can construct a concrete instance of `ServiceEndpointKubernetesMapInput` via:
//
//          ServiceEndpointKubernetesMap{ "key": ServiceEndpointKubernetesArgs{...} }
type ServiceEndpointKubernetesMapInput interface {
	pulumi.Input

	ToServiceEndpointKubernetesMapOutput() ServiceEndpointKubernetesMapOutput
	ToServiceEndpointKubernetesMapOutputWithContext(context.Context) ServiceEndpointKubernetesMapOutput
}

type ServiceEndpointKubernetesMap map[string]ServiceEndpointKubernetesInput

func (ServiceEndpointKubernetesMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceEndpointKubernetes)(nil)).Elem()
}

func (i ServiceEndpointKubernetesMap) ToServiceEndpointKubernetesMapOutput() ServiceEndpointKubernetesMapOutput {
	return i.ToServiceEndpointKubernetesMapOutputWithContext(context.Background())
}

func (i ServiceEndpointKubernetesMap) ToServiceEndpointKubernetesMapOutputWithContext(ctx context.Context) ServiceEndpointKubernetesMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointKubernetesMapOutput)
}

type ServiceEndpointKubernetesOutput struct{ *pulumi.OutputState }

func (ServiceEndpointKubernetesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceEndpointKubernetes)(nil)).Elem()
}

func (o ServiceEndpointKubernetesOutput) ToServiceEndpointKubernetesOutput() ServiceEndpointKubernetesOutput {
	return o
}

func (o ServiceEndpointKubernetesOutput) ToServiceEndpointKubernetesOutputWithContext(ctx context.Context) ServiceEndpointKubernetesOutput {
	return o
}

// The hostname (in form of URI) of the Kubernetes API.
func (o ServiceEndpointKubernetesOutput) ApiserverUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEndpointKubernetes) pulumi.StringOutput { return v.ApiserverUrl }).(pulumi.StringOutput)
}

func (o ServiceEndpointKubernetesOutput) Authorization() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ServiceEndpointKubernetes) pulumi.StringMapOutput { return v.Authorization }).(pulumi.StringMapOutput)
}

// The authentication method used to authenticate on the Kubernetes cluster. The value should be one of AzureSubscription, Kubeconfig, ServiceAccount.
func (o ServiceEndpointKubernetesOutput) AuthorizationType() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEndpointKubernetes) pulumi.StringOutput { return v.AuthorizationType }).(pulumi.StringOutput)
}

// A `azureSubscription` block defined blow.
func (o ServiceEndpointKubernetesOutput) AzureSubscriptions() ServiceEndpointKubernetesAzureSubscriptionArrayOutput {
	return o.ApplyT(func(v *ServiceEndpointKubernetes) ServiceEndpointKubernetesAzureSubscriptionArrayOutput {
		return v.AzureSubscriptions
	}).(ServiceEndpointKubernetesAzureSubscriptionArrayOutput)
}

func (o ServiceEndpointKubernetesOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceEndpointKubernetes) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// A `kubeconfig` block defined blow.
func (o ServiceEndpointKubernetesOutput) Kubeconfigs() ServiceEndpointKubernetesKubeconfigArrayOutput {
	return o.ApplyT(func(v *ServiceEndpointKubernetes) ServiceEndpointKubernetesKubeconfigArrayOutput {
		return v.Kubeconfigs
	}).(ServiceEndpointKubernetesKubeconfigArrayOutput)
}

// The ID of the project.
func (o ServiceEndpointKubernetesOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEndpointKubernetes) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// A `serviceAccount` block defined blow.
func (o ServiceEndpointKubernetesOutput) ServiceAccount() ServiceEndpointKubernetesServiceAccountPtrOutput {
	return o.ApplyT(func(v *ServiceEndpointKubernetes) ServiceEndpointKubernetesServiceAccountPtrOutput {
		return v.ServiceAccount
	}).(ServiceEndpointKubernetesServiceAccountPtrOutput)
}

// The Service Endpoint name.
func (o ServiceEndpointKubernetesOutput) ServiceEndpointName() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEndpointKubernetes) pulumi.StringOutput { return v.ServiceEndpointName }).(pulumi.StringOutput)
}

type ServiceEndpointKubernetesArrayOutput struct{ *pulumi.OutputState }

func (ServiceEndpointKubernetesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceEndpointKubernetes)(nil)).Elem()
}

func (o ServiceEndpointKubernetesArrayOutput) ToServiceEndpointKubernetesArrayOutput() ServiceEndpointKubernetesArrayOutput {
	return o
}

func (o ServiceEndpointKubernetesArrayOutput) ToServiceEndpointKubernetesArrayOutputWithContext(ctx context.Context) ServiceEndpointKubernetesArrayOutput {
	return o
}

func (o ServiceEndpointKubernetesArrayOutput) Index(i pulumi.IntInput) ServiceEndpointKubernetesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServiceEndpointKubernetes {
		return vs[0].([]*ServiceEndpointKubernetes)[vs[1].(int)]
	}).(ServiceEndpointKubernetesOutput)
}

type ServiceEndpointKubernetesMapOutput struct{ *pulumi.OutputState }

func (ServiceEndpointKubernetesMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceEndpointKubernetes)(nil)).Elem()
}

func (o ServiceEndpointKubernetesMapOutput) ToServiceEndpointKubernetesMapOutput() ServiceEndpointKubernetesMapOutput {
	return o
}

func (o ServiceEndpointKubernetesMapOutput) ToServiceEndpointKubernetesMapOutputWithContext(ctx context.Context) ServiceEndpointKubernetesMapOutput {
	return o
}

func (o ServiceEndpointKubernetesMapOutput) MapIndex(k pulumi.StringInput) ServiceEndpointKubernetesOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServiceEndpointKubernetes {
		return vs[0].(map[string]*ServiceEndpointKubernetes)[vs[1].(string)]
	}).(ServiceEndpointKubernetesOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceEndpointKubernetesInput)(nil)).Elem(), &ServiceEndpointKubernetes{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceEndpointKubernetesArrayInput)(nil)).Elem(), ServiceEndpointKubernetesArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceEndpointKubernetesMapInput)(nil)).Elem(), ServiceEndpointKubernetesMap{})
	pulumi.RegisterOutputType(ServiceEndpointKubernetesOutput{})
	pulumi.RegisterOutputType(ServiceEndpointKubernetesArrayOutput{})
	pulumi.RegisterOutputType(ServiceEndpointKubernetesMapOutput{})
}
