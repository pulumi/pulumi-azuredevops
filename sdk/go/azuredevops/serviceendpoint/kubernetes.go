// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package serviceendpoint

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
// - [Azure DevOps Service REST API 5.1 - Endpoints](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-5.1)
//
// ## Import
//
// Azure DevOps Service Endpoint Kubernetes can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**
//
// ```sh
//  $ pulumi import azuredevops:ServiceEndpoint/kubernetes:Kubernetes serviceendpoint 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
// ```
//
// Deprecated: azuredevops.serviceendpoint.Kubernetes has been deprecated in favor of azuredevops.ServiceEndpointKubernetes
type Kubernetes struct {
	pulumi.CustomResourceState

	// The hostname (in form of URI) of the Kubernetes API.
	ApiserverUrl  pulumi.StringOutput    `pulumi:"apiserverUrl"`
	Authorization pulumi.StringMapOutput `pulumi:"authorization"`
	// The authentication method used to authenticate on the Kubernetes cluster. The value should be one of AzureSubscription, Kubeconfig, ServiceAccount.
	AuthorizationType pulumi.StringOutput `pulumi:"authorizationType"`
	// The configuration for authorization_type="AzureSubscription".
	AzureSubscriptions KubernetesAzureSubscriptionArrayOutput `pulumi:"azureSubscriptions"`
	Description        pulumi.StringPtrOutput                 `pulumi:"description"`
	// The configuration for authorization_type="Kubeconfig".
	Kubeconfigs KubernetesKubeconfigArrayOutput `pulumi:"kubeconfigs"`
	// The project ID or project name.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The configuration for authorization_type="ServiceAccount". This type uses the credentials of a service account currently deployed to the cluster.
	ServiceAccounts KubernetesServiceAccountArrayOutput `pulumi:"serviceAccounts"`
	// The Service Endpoint name.
	ServiceEndpointName pulumi.StringOutput `pulumi:"serviceEndpointName"`
}

// NewKubernetes registers a new resource with the given unique name, arguments, and options.
func NewKubernetes(ctx *pulumi.Context,
	name string, args *KubernetesArgs, opts ...pulumi.ResourceOption) (*Kubernetes, error) {
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
	var resource Kubernetes
	err := ctx.RegisterResource("azuredevops:ServiceEndpoint/kubernetes:Kubernetes", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKubernetes gets an existing Kubernetes resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKubernetes(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KubernetesState, opts ...pulumi.ResourceOption) (*Kubernetes, error) {
	var resource Kubernetes
	err := ctx.ReadResource("azuredevops:ServiceEndpoint/kubernetes:Kubernetes", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Kubernetes resources.
type kubernetesState struct {
	// The hostname (in form of URI) of the Kubernetes API.
	ApiserverUrl  *string           `pulumi:"apiserverUrl"`
	Authorization map[string]string `pulumi:"authorization"`
	// The authentication method used to authenticate on the Kubernetes cluster. The value should be one of AzureSubscription, Kubeconfig, ServiceAccount.
	AuthorizationType *string `pulumi:"authorizationType"`
	// The configuration for authorization_type="AzureSubscription".
	AzureSubscriptions []KubernetesAzureSubscription `pulumi:"azureSubscriptions"`
	Description        *string                       `pulumi:"description"`
	// The configuration for authorization_type="Kubeconfig".
	Kubeconfigs []KubernetesKubeconfig `pulumi:"kubeconfigs"`
	// The project ID or project name.
	ProjectId *string `pulumi:"projectId"`
	// The configuration for authorization_type="ServiceAccount". This type uses the credentials of a service account currently deployed to the cluster.
	ServiceAccounts []KubernetesServiceAccount `pulumi:"serviceAccounts"`
	// The Service Endpoint name.
	ServiceEndpointName *string `pulumi:"serviceEndpointName"`
}

type KubernetesState struct {
	// The hostname (in form of URI) of the Kubernetes API.
	ApiserverUrl  pulumi.StringPtrInput
	Authorization pulumi.StringMapInput
	// The authentication method used to authenticate on the Kubernetes cluster. The value should be one of AzureSubscription, Kubeconfig, ServiceAccount.
	AuthorizationType pulumi.StringPtrInput
	// The configuration for authorization_type="AzureSubscription".
	AzureSubscriptions KubernetesAzureSubscriptionArrayInput
	Description        pulumi.StringPtrInput
	// The configuration for authorization_type="Kubeconfig".
	Kubeconfigs KubernetesKubeconfigArrayInput
	// The project ID or project name.
	ProjectId pulumi.StringPtrInput
	// The configuration for authorization_type="ServiceAccount". This type uses the credentials of a service account currently deployed to the cluster.
	ServiceAccounts KubernetesServiceAccountArrayInput
	// The Service Endpoint name.
	ServiceEndpointName pulumi.StringPtrInput
}

func (KubernetesState) ElementType() reflect.Type {
	return reflect.TypeOf((*kubernetesState)(nil)).Elem()
}

type kubernetesArgs struct {
	// The hostname (in form of URI) of the Kubernetes API.
	ApiserverUrl  string            `pulumi:"apiserverUrl"`
	Authorization map[string]string `pulumi:"authorization"`
	// The authentication method used to authenticate on the Kubernetes cluster. The value should be one of AzureSubscription, Kubeconfig, ServiceAccount.
	AuthorizationType string `pulumi:"authorizationType"`
	// The configuration for authorization_type="AzureSubscription".
	AzureSubscriptions []KubernetesAzureSubscription `pulumi:"azureSubscriptions"`
	Description        *string                       `pulumi:"description"`
	// The configuration for authorization_type="Kubeconfig".
	Kubeconfigs []KubernetesKubeconfig `pulumi:"kubeconfigs"`
	// The project ID or project name.
	ProjectId string `pulumi:"projectId"`
	// The configuration for authorization_type="ServiceAccount". This type uses the credentials of a service account currently deployed to the cluster.
	ServiceAccounts []KubernetesServiceAccount `pulumi:"serviceAccounts"`
	// The Service Endpoint name.
	ServiceEndpointName string `pulumi:"serviceEndpointName"`
}

// The set of arguments for constructing a Kubernetes resource.
type KubernetesArgs struct {
	// The hostname (in form of URI) of the Kubernetes API.
	ApiserverUrl  pulumi.StringInput
	Authorization pulumi.StringMapInput
	// The authentication method used to authenticate on the Kubernetes cluster. The value should be one of AzureSubscription, Kubeconfig, ServiceAccount.
	AuthorizationType pulumi.StringInput
	// The configuration for authorization_type="AzureSubscription".
	AzureSubscriptions KubernetesAzureSubscriptionArrayInput
	Description        pulumi.StringPtrInput
	// The configuration for authorization_type="Kubeconfig".
	Kubeconfigs KubernetesKubeconfigArrayInput
	// The project ID or project name.
	ProjectId pulumi.StringInput
	// The configuration for authorization_type="ServiceAccount". This type uses the credentials of a service account currently deployed to the cluster.
	ServiceAccounts KubernetesServiceAccountArrayInput
	// The Service Endpoint name.
	ServiceEndpointName pulumi.StringInput
}

func (KubernetesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*kubernetesArgs)(nil)).Elem()
}

type KubernetesInput interface {
	pulumi.Input

	ToKubernetesOutput() KubernetesOutput
	ToKubernetesOutputWithContext(ctx context.Context) KubernetesOutput
}

func (*Kubernetes) ElementType() reflect.Type {
	return reflect.TypeOf((**Kubernetes)(nil)).Elem()
}

func (i *Kubernetes) ToKubernetesOutput() KubernetesOutput {
	return i.ToKubernetesOutputWithContext(context.Background())
}

func (i *Kubernetes) ToKubernetesOutputWithContext(ctx context.Context) KubernetesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesOutput)
}

// KubernetesArrayInput is an input type that accepts KubernetesArray and KubernetesArrayOutput values.
// You can construct a concrete instance of `KubernetesArrayInput` via:
//
//          KubernetesArray{ KubernetesArgs{...} }
type KubernetesArrayInput interface {
	pulumi.Input

	ToKubernetesArrayOutput() KubernetesArrayOutput
	ToKubernetesArrayOutputWithContext(context.Context) KubernetesArrayOutput
}

type KubernetesArray []KubernetesInput

func (KubernetesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Kubernetes)(nil)).Elem()
}

func (i KubernetesArray) ToKubernetesArrayOutput() KubernetesArrayOutput {
	return i.ToKubernetesArrayOutputWithContext(context.Background())
}

func (i KubernetesArray) ToKubernetesArrayOutputWithContext(ctx context.Context) KubernetesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesArrayOutput)
}

// KubernetesMapInput is an input type that accepts KubernetesMap and KubernetesMapOutput values.
// You can construct a concrete instance of `KubernetesMapInput` via:
//
//          KubernetesMap{ "key": KubernetesArgs{...} }
type KubernetesMapInput interface {
	pulumi.Input

	ToKubernetesMapOutput() KubernetesMapOutput
	ToKubernetesMapOutputWithContext(context.Context) KubernetesMapOutput
}

type KubernetesMap map[string]KubernetesInput

func (KubernetesMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Kubernetes)(nil)).Elem()
}

func (i KubernetesMap) ToKubernetesMapOutput() KubernetesMapOutput {
	return i.ToKubernetesMapOutputWithContext(context.Background())
}

func (i KubernetesMap) ToKubernetesMapOutputWithContext(ctx context.Context) KubernetesMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesMapOutput)
}

type KubernetesOutput struct{ *pulumi.OutputState }

func (KubernetesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Kubernetes)(nil)).Elem()
}

func (o KubernetesOutput) ToKubernetesOutput() KubernetesOutput {
	return o
}

func (o KubernetesOutput) ToKubernetesOutputWithContext(ctx context.Context) KubernetesOutput {
	return o
}

type KubernetesArrayOutput struct{ *pulumi.OutputState }

func (KubernetesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Kubernetes)(nil)).Elem()
}

func (o KubernetesArrayOutput) ToKubernetesArrayOutput() KubernetesArrayOutput {
	return o
}

func (o KubernetesArrayOutput) ToKubernetesArrayOutputWithContext(ctx context.Context) KubernetesArrayOutput {
	return o
}

func (o KubernetesArrayOutput) Index(i pulumi.IntInput) KubernetesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Kubernetes {
		return vs[0].([]*Kubernetes)[vs[1].(int)]
	}).(KubernetesOutput)
}

type KubernetesMapOutput struct{ *pulumi.OutputState }

func (KubernetesMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Kubernetes)(nil)).Elem()
}

func (o KubernetesMapOutput) ToKubernetesMapOutput() KubernetesMapOutput {
	return o
}

func (o KubernetesMapOutput) ToKubernetesMapOutputWithContext(ctx context.Context) KubernetesMapOutput {
	return o
}

func (o KubernetesMapOutput) MapIndex(k pulumi.StringInput) KubernetesOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Kubernetes {
		return vs[0].(map[string]*Kubernetes)[vs[1].(string)]
	}).(KubernetesOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*KubernetesInput)(nil)).Elem(), &Kubernetes{})
	pulumi.RegisterInputType(reflect.TypeOf((*KubernetesArrayInput)(nil)).Elem(), KubernetesArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*KubernetesMapInput)(nil)).Elem(), KubernetesMap{})
	pulumi.RegisterOutputType(KubernetesOutput{})
	pulumi.RegisterOutputType(KubernetesArrayOutput{})
	pulumi.RegisterOutputType(KubernetesMapOutput{})
}
