// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ## # Kubernetes
//
// Manages a Kubernetes service endpoint within Azure DevOps.
//
// ## Relevant Links
//
// * [Azure DevOps Service REST API 5.1 - Endpoints](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-5.1)
type Kubernetes struct {
	pulumi.CustomResourceState

	// The Service Endpoint description.
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
	if args == nil || args.ApiserverUrl == nil {
		return nil, errors.New("missing required argument 'ApiserverUrl'")
	}
	if args == nil || args.AuthorizationType == nil {
		return nil, errors.New("missing required argument 'AuthorizationType'")
	}
	if args == nil || args.ProjectId == nil {
		return nil, errors.New("missing required argument 'ProjectId'")
	}
	if args == nil || args.ServiceEndpointName == nil {
		return nil, errors.New("missing required argument 'ServiceEndpointName'")
	}
	if args == nil {
		args = &KubernetesArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azuredevops:ServiceEndpoint/kubernetes:Kubernetes"),
		},
	})
	opts = append(opts, aliases)
	var resource Kubernetes
	err := ctx.RegisterResource("azuredevops:index/kubernetes:Kubernetes", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azuredevops:index/kubernetes:Kubernetes", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Kubernetes resources.
type kubernetesState struct {
	// The Service Endpoint description.
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
	// The Service Endpoint description.
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
	// The Service Endpoint description.
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
	// The Service Endpoint description.
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
