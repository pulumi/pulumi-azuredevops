// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
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
//  $ pulumi import azuredevops:index/serviceEndpointKubernetes:ServiceEndpointKubernetes serviceendpoint 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
// ```
type ServiceEndpointKubernetes struct {
	pulumi.CustomResourceState

	// The Service Endpoint description.
	ApiserverUrl  pulumi.StringOutput    `pulumi:"apiserverUrl"`
	Authorization pulumi.StringMapOutput `pulumi:"authorization"`
	// The authentication method used to authenticate on the Kubernetes cluster. The value should be one of AzureSubscription, Kubeconfig, ServiceAccount.
	AuthorizationType pulumi.StringOutput `pulumi:"authorizationType"`
	// The configuration for authorization_type="AzureSubscription".
	AzureSubscriptions ServiceEndpointKubernetesAzureSubscriptionArrayOutput `pulumi:"azureSubscriptions"`
	Description        pulumi.StringPtrOutput                                `pulumi:"description"`
	// The configuration for authorization_type="Kubeconfig".
	Kubeconfigs ServiceEndpointKubernetesKubeconfigArrayOutput `pulumi:"kubeconfigs"`
	// The project ID or project name.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The configuration for authorization_type="ServiceAccount". This type uses the credentials of a service account currently deployed to the cluster.
	ServiceAccounts ServiceEndpointKubernetesServiceAccountArrayOutput `pulumi:"serviceAccounts"`
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
	// The Service Endpoint description.
	ApiserverUrl  *string           `pulumi:"apiserverUrl"`
	Authorization map[string]string `pulumi:"authorization"`
	// The authentication method used to authenticate on the Kubernetes cluster. The value should be one of AzureSubscription, Kubeconfig, ServiceAccount.
	AuthorizationType *string `pulumi:"authorizationType"`
	// The configuration for authorization_type="AzureSubscription".
	AzureSubscriptions []ServiceEndpointKubernetesAzureSubscription `pulumi:"azureSubscriptions"`
	Description        *string                                      `pulumi:"description"`
	// The configuration for authorization_type="Kubeconfig".
	Kubeconfigs []ServiceEndpointKubernetesKubeconfig `pulumi:"kubeconfigs"`
	// The project ID or project name.
	ProjectId *string `pulumi:"projectId"`
	// The configuration for authorization_type="ServiceAccount". This type uses the credentials of a service account currently deployed to the cluster.
	ServiceAccounts []ServiceEndpointKubernetesServiceAccount `pulumi:"serviceAccounts"`
	// The Service Endpoint name.
	ServiceEndpointName *string `pulumi:"serviceEndpointName"`
}

type ServiceEndpointKubernetesState struct {
	// The Service Endpoint description.
	ApiserverUrl  pulumi.StringPtrInput
	Authorization pulumi.StringMapInput
	// The authentication method used to authenticate on the Kubernetes cluster. The value should be one of AzureSubscription, Kubeconfig, ServiceAccount.
	AuthorizationType pulumi.StringPtrInput
	// The configuration for authorization_type="AzureSubscription".
	AzureSubscriptions ServiceEndpointKubernetesAzureSubscriptionArrayInput
	Description        pulumi.StringPtrInput
	// The configuration for authorization_type="Kubeconfig".
	Kubeconfigs ServiceEndpointKubernetesKubeconfigArrayInput
	// The project ID or project name.
	ProjectId pulumi.StringPtrInput
	// The configuration for authorization_type="ServiceAccount". This type uses the credentials of a service account currently deployed to the cluster.
	ServiceAccounts ServiceEndpointKubernetesServiceAccountArrayInput
	// The Service Endpoint name.
	ServiceEndpointName pulumi.StringPtrInput
}

func (ServiceEndpointKubernetesState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceEndpointKubernetesState)(nil)).Elem()
}

type serviceEndpointKubernetesArgs struct {
	// The Service Endpoint description.
	ApiserverUrl  string            `pulumi:"apiserverUrl"`
	Authorization map[string]string `pulumi:"authorization"`
	// The authentication method used to authenticate on the Kubernetes cluster. The value should be one of AzureSubscription, Kubeconfig, ServiceAccount.
	AuthorizationType string `pulumi:"authorizationType"`
	// The configuration for authorization_type="AzureSubscription".
	AzureSubscriptions []ServiceEndpointKubernetesAzureSubscription `pulumi:"azureSubscriptions"`
	Description        *string                                      `pulumi:"description"`
	// The configuration for authorization_type="Kubeconfig".
	Kubeconfigs []ServiceEndpointKubernetesKubeconfig `pulumi:"kubeconfigs"`
	// The project ID or project name.
	ProjectId string `pulumi:"projectId"`
	// The configuration for authorization_type="ServiceAccount". This type uses the credentials of a service account currently deployed to the cluster.
	ServiceAccounts []ServiceEndpointKubernetesServiceAccount `pulumi:"serviceAccounts"`
	// The Service Endpoint name.
	ServiceEndpointName string `pulumi:"serviceEndpointName"`
}

// The set of arguments for constructing a ServiceEndpointKubernetes resource.
type ServiceEndpointKubernetesArgs struct {
	// The Service Endpoint description.
	ApiserverUrl  pulumi.StringInput
	Authorization pulumi.StringMapInput
	// The authentication method used to authenticate on the Kubernetes cluster. The value should be one of AzureSubscription, Kubeconfig, ServiceAccount.
	AuthorizationType pulumi.StringInput
	// The configuration for authorization_type="AzureSubscription".
	AzureSubscriptions ServiceEndpointKubernetesAzureSubscriptionArrayInput
	Description        pulumi.StringPtrInput
	// The configuration for authorization_type="Kubeconfig".
	Kubeconfigs ServiceEndpointKubernetesKubeconfigArrayInput
	// The project ID or project name.
	ProjectId pulumi.StringInput
	// The configuration for authorization_type="ServiceAccount". This type uses the credentials of a service account currently deployed to the cluster.
	ServiceAccounts ServiceEndpointKubernetesServiceAccountArrayInput
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

func (ServiceEndpointKubernetes) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceEndpointKubernetes)(nil)).Elem()
}

func (i ServiceEndpointKubernetes) ToServiceEndpointKubernetesOutput() ServiceEndpointKubernetesOutput {
	return i.ToServiceEndpointKubernetesOutputWithContext(context.Background())
}

func (i ServiceEndpointKubernetes) ToServiceEndpointKubernetesOutputWithContext(ctx context.Context) ServiceEndpointKubernetesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointKubernetesOutput)
}

type ServiceEndpointKubernetesOutput struct {
	*pulumi.OutputState
}

func (ServiceEndpointKubernetesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceEndpointKubernetesOutput)(nil)).Elem()
}

func (o ServiceEndpointKubernetesOutput) ToServiceEndpointKubernetesOutput() ServiceEndpointKubernetesOutput {
	return o
}

func (o ServiceEndpointKubernetesOutput) ToServiceEndpointKubernetesOutputWithContext(ctx context.Context) ServiceEndpointKubernetesOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ServiceEndpointKubernetesOutput{})
}
