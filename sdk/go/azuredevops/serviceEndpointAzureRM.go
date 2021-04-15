// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages Manual or Automatic AzureRM service endpoint within Azure DevOps.
//
// ## Requirements (Manual AzureRM Service Endpoint)
//
// Before to create a service end point in Azure DevOps, you need to create a Service Principal in your Azure subscription.
//
// For detailed steps to create a service principal with Azure cli see the [documentation](https://docs.microsoft.com/en-us/cli/azure/create-an-azure-service-principal-azure-cli?view=azure-cli-latest)
//
// ## Example Usage
// ### Manual AzureRM Service Endpoint
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azuredevops/sdk/v2/go/azuredevops"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		project, err := azuredevops.NewProject(ctx, "project", &azuredevops.ProjectArgs{
// 			Visibility:       pulumi.String("private"),
// 			VersionControl:   pulumi.String("Git"),
// 			WorkItemTemplate: pulumi.String("Agile"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = azuredevops.NewServiceEndpointAzureRM(ctx, "endpointazure", &azuredevops.ServiceEndpointAzureRMArgs{
// 			ProjectId:           project.ID(),
// 			ServiceEndpointName: pulumi.String("Sample AzureRM"),
// 			Description:         pulumi.String("Managed by Terraform"),
// 			Credentials: &azuredevops.ServiceEndpointAzureRMCredentialsArgs{
// 				Serviceprincipalid:  pulumi.String("00000000-0000-0000-0000-000000000000"),
// 				Serviceprincipalkey: pulumi.String("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
// 			},
// 			AzurermSpnTenantid:      pulumi.String("00000000-0000-0000-0000-000000000000"),
// 			AzurermSubscriptionId:   pulumi.String("00000000-0000-0000-0000-000000000000"),
// 			AzurermSubscriptionName: pulumi.String("Sample Subscription"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Automatic AzureRM Service Endpoint
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azuredevops/sdk/v2/go/azuredevops"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		project, err := azuredevops.NewProject(ctx, "project", &azuredevops.ProjectArgs{
// 			Visibility:       pulumi.String("private"),
// 			VersionControl:   pulumi.String("Git"),
// 			WorkItemTemplate: pulumi.String("Agile"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = azuredevops.NewServiceEndpointAzureRM(ctx, "endpointazure", &azuredevops.ServiceEndpointAzureRMArgs{
// 			ProjectId:               project.ID(),
// 			ServiceEndpointName:     pulumi.String("Sample AzureRM"),
// 			Description:             pulumi.String("Managed by Terraform"),
// 			AzurermSpnTenantid:      pulumi.String("00000000-0000-0000-0000-000000000000"),
// 			AzurermSubscriptionId:   pulumi.String("00000000-0000-0000-0000-000000000000"),
// 			AzurermSubscriptionName: pulumi.String("Microsoft Azure DEMO"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ## Relevant Links
//
// - [Azure DevOps Service REST API 5.1 - Service End points](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-5.1)
//
// ## Import
//
// Azure DevOps Service Endpoint Azure Resource Manage can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**
//
// ```sh
//  $ pulumi import azuredevops:index/serviceEndpointAzureRM:ServiceEndpointAzureRM serviceendpoint 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
// ```
type ServiceEndpointAzureRM struct {
	pulumi.CustomResourceState

	Authorization pulumi.StringMapOutput `pulumi:"authorization"`
	// The tenant id if the service principal.
	AzurermSpnTenantid pulumi.StringOutput `pulumi:"azurermSpnTenantid"`
	// The subscription Id of the Azure targets.
	AzurermSubscriptionId pulumi.StringOutput `pulumi:"azurermSubscriptionId"`
	// The subscription Name of the targets.
	AzurermSubscriptionName pulumi.StringOutput `pulumi:"azurermSubscriptionName"`
	// A `credentials` block.
	Credentials ServiceEndpointAzureRMCredentialsPtrOutput `pulumi:"credentials"`
	// Service connection description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The project ID or project name.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The resource group used for scope of automatic service endpoint.
	ResourceGroup pulumi.StringPtrOutput `pulumi:"resourceGroup"`
	// The Service Endpoint name.
	ServiceEndpointName pulumi.StringOutput `pulumi:"serviceEndpointName"`
}

// NewServiceEndpointAzureRM registers a new resource with the given unique name, arguments, and options.
func NewServiceEndpointAzureRM(ctx *pulumi.Context,
	name string, args *ServiceEndpointAzureRMArgs, opts ...pulumi.ResourceOption) (*ServiceEndpointAzureRM, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AzurermSpnTenantid == nil {
		return nil, errors.New("invalid value for required argument 'AzurermSpnTenantid'")
	}
	if args.AzurermSubscriptionId == nil {
		return nil, errors.New("invalid value for required argument 'AzurermSubscriptionId'")
	}
	if args.AzurermSubscriptionName == nil {
		return nil, errors.New("invalid value for required argument 'AzurermSubscriptionName'")
	}
	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.ServiceEndpointName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceEndpointName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azuredevops:ServiceEndpoint/azureRM:AzureRM"),
		},
	})
	opts = append(opts, aliases)
	var resource ServiceEndpointAzureRM
	err := ctx.RegisterResource("azuredevops:index/serviceEndpointAzureRM:ServiceEndpointAzureRM", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceEndpointAzureRM gets an existing ServiceEndpointAzureRM resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceEndpointAzureRM(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceEndpointAzureRMState, opts ...pulumi.ResourceOption) (*ServiceEndpointAzureRM, error) {
	var resource ServiceEndpointAzureRM
	err := ctx.ReadResource("azuredevops:index/serviceEndpointAzureRM:ServiceEndpointAzureRM", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceEndpointAzureRM resources.
type serviceEndpointAzureRMState struct {
	Authorization map[string]string `pulumi:"authorization"`
	// The tenant id if the service principal.
	AzurermSpnTenantid *string `pulumi:"azurermSpnTenantid"`
	// The subscription Id of the Azure targets.
	AzurermSubscriptionId *string `pulumi:"azurermSubscriptionId"`
	// The subscription Name of the targets.
	AzurermSubscriptionName *string `pulumi:"azurermSubscriptionName"`
	// A `credentials` block.
	Credentials *ServiceEndpointAzureRMCredentials `pulumi:"credentials"`
	// Service connection description.
	Description *string `pulumi:"description"`
	// The project ID or project name.
	ProjectId *string `pulumi:"projectId"`
	// The resource group used for scope of automatic service endpoint.
	ResourceGroup *string `pulumi:"resourceGroup"`
	// The Service Endpoint name.
	ServiceEndpointName *string `pulumi:"serviceEndpointName"`
}

type ServiceEndpointAzureRMState struct {
	Authorization pulumi.StringMapInput
	// The tenant id if the service principal.
	AzurermSpnTenantid pulumi.StringPtrInput
	// The subscription Id of the Azure targets.
	AzurermSubscriptionId pulumi.StringPtrInput
	// The subscription Name of the targets.
	AzurermSubscriptionName pulumi.StringPtrInput
	// A `credentials` block.
	Credentials ServiceEndpointAzureRMCredentialsPtrInput
	// Service connection description.
	Description pulumi.StringPtrInput
	// The project ID or project name.
	ProjectId pulumi.StringPtrInput
	// The resource group used for scope of automatic service endpoint.
	ResourceGroup pulumi.StringPtrInput
	// The Service Endpoint name.
	ServiceEndpointName pulumi.StringPtrInput
}

func (ServiceEndpointAzureRMState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceEndpointAzureRMState)(nil)).Elem()
}

type serviceEndpointAzureRMArgs struct {
	Authorization map[string]string `pulumi:"authorization"`
	// The tenant id if the service principal.
	AzurermSpnTenantid string `pulumi:"azurermSpnTenantid"`
	// The subscription Id of the Azure targets.
	AzurermSubscriptionId string `pulumi:"azurermSubscriptionId"`
	// The subscription Name of the targets.
	AzurermSubscriptionName string `pulumi:"azurermSubscriptionName"`
	// A `credentials` block.
	Credentials *ServiceEndpointAzureRMCredentials `pulumi:"credentials"`
	// Service connection description.
	Description *string `pulumi:"description"`
	// The project ID or project name.
	ProjectId string `pulumi:"projectId"`
	// The resource group used for scope of automatic service endpoint.
	ResourceGroup *string `pulumi:"resourceGroup"`
	// The Service Endpoint name.
	ServiceEndpointName string `pulumi:"serviceEndpointName"`
}

// The set of arguments for constructing a ServiceEndpointAzureRM resource.
type ServiceEndpointAzureRMArgs struct {
	Authorization pulumi.StringMapInput
	// The tenant id if the service principal.
	AzurermSpnTenantid pulumi.StringInput
	// The subscription Id of the Azure targets.
	AzurermSubscriptionId pulumi.StringInput
	// The subscription Name of the targets.
	AzurermSubscriptionName pulumi.StringInput
	// A `credentials` block.
	Credentials ServiceEndpointAzureRMCredentialsPtrInput
	// Service connection description.
	Description pulumi.StringPtrInput
	// The project ID or project name.
	ProjectId pulumi.StringInput
	// The resource group used for scope of automatic service endpoint.
	ResourceGroup pulumi.StringPtrInput
	// The Service Endpoint name.
	ServiceEndpointName pulumi.StringInput
}

func (ServiceEndpointAzureRMArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceEndpointAzureRMArgs)(nil)).Elem()
}

type ServiceEndpointAzureRMInput interface {
	pulumi.Input

	ToServiceEndpointAzureRMOutput() ServiceEndpointAzureRMOutput
	ToServiceEndpointAzureRMOutputWithContext(ctx context.Context) ServiceEndpointAzureRMOutput
}

func (*ServiceEndpointAzureRM) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceEndpointAzureRM)(nil))
}

func (i *ServiceEndpointAzureRM) ToServiceEndpointAzureRMOutput() ServiceEndpointAzureRMOutput {
	return i.ToServiceEndpointAzureRMOutputWithContext(context.Background())
}

func (i *ServiceEndpointAzureRM) ToServiceEndpointAzureRMOutputWithContext(ctx context.Context) ServiceEndpointAzureRMOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointAzureRMOutput)
}

func (i *ServiceEndpointAzureRM) ToServiceEndpointAzureRMPtrOutput() ServiceEndpointAzureRMPtrOutput {
	return i.ToServiceEndpointAzureRMPtrOutputWithContext(context.Background())
}

func (i *ServiceEndpointAzureRM) ToServiceEndpointAzureRMPtrOutputWithContext(ctx context.Context) ServiceEndpointAzureRMPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointAzureRMPtrOutput)
}

type ServiceEndpointAzureRMPtrInput interface {
	pulumi.Input

	ToServiceEndpointAzureRMPtrOutput() ServiceEndpointAzureRMPtrOutput
	ToServiceEndpointAzureRMPtrOutputWithContext(ctx context.Context) ServiceEndpointAzureRMPtrOutput
}

type serviceEndpointAzureRMPtrType ServiceEndpointAzureRMArgs

func (*serviceEndpointAzureRMPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceEndpointAzureRM)(nil))
}

func (i *serviceEndpointAzureRMPtrType) ToServiceEndpointAzureRMPtrOutput() ServiceEndpointAzureRMPtrOutput {
	return i.ToServiceEndpointAzureRMPtrOutputWithContext(context.Background())
}

func (i *serviceEndpointAzureRMPtrType) ToServiceEndpointAzureRMPtrOutputWithContext(ctx context.Context) ServiceEndpointAzureRMPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointAzureRMPtrOutput)
}

// ServiceEndpointAzureRMArrayInput is an input type that accepts ServiceEndpointAzureRMArray and ServiceEndpointAzureRMArrayOutput values.
// You can construct a concrete instance of `ServiceEndpointAzureRMArrayInput` via:
//
//          ServiceEndpointAzureRMArray{ ServiceEndpointAzureRMArgs{...} }
type ServiceEndpointAzureRMArrayInput interface {
	pulumi.Input

	ToServiceEndpointAzureRMArrayOutput() ServiceEndpointAzureRMArrayOutput
	ToServiceEndpointAzureRMArrayOutputWithContext(context.Context) ServiceEndpointAzureRMArrayOutput
}

type ServiceEndpointAzureRMArray []ServiceEndpointAzureRMInput

func (ServiceEndpointAzureRMArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*ServiceEndpointAzureRM)(nil))
}

func (i ServiceEndpointAzureRMArray) ToServiceEndpointAzureRMArrayOutput() ServiceEndpointAzureRMArrayOutput {
	return i.ToServiceEndpointAzureRMArrayOutputWithContext(context.Background())
}

func (i ServiceEndpointAzureRMArray) ToServiceEndpointAzureRMArrayOutputWithContext(ctx context.Context) ServiceEndpointAzureRMArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointAzureRMArrayOutput)
}

// ServiceEndpointAzureRMMapInput is an input type that accepts ServiceEndpointAzureRMMap and ServiceEndpointAzureRMMapOutput values.
// You can construct a concrete instance of `ServiceEndpointAzureRMMapInput` via:
//
//          ServiceEndpointAzureRMMap{ "key": ServiceEndpointAzureRMArgs{...} }
type ServiceEndpointAzureRMMapInput interface {
	pulumi.Input

	ToServiceEndpointAzureRMMapOutput() ServiceEndpointAzureRMMapOutput
	ToServiceEndpointAzureRMMapOutputWithContext(context.Context) ServiceEndpointAzureRMMapOutput
}

type ServiceEndpointAzureRMMap map[string]ServiceEndpointAzureRMInput

func (ServiceEndpointAzureRMMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*ServiceEndpointAzureRM)(nil))
}

func (i ServiceEndpointAzureRMMap) ToServiceEndpointAzureRMMapOutput() ServiceEndpointAzureRMMapOutput {
	return i.ToServiceEndpointAzureRMMapOutputWithContext(context.Background())
}

func (i ServiceEndpointAzureRMMap) ToServiceEndpointAzureRMMapOutputWithContext(ctx context.Context) ServiceEndpointAzureRMMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointAzureRMMapOutput)
}

type ServiceEndpointAzureRMOutput struct {
	*pulumi.OutputState
}

func (ServiceEndpointAzureRMOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceEndpointAzureRM)(nil))
}

func (o ServiceEndpointAzureRMOutput) ToServiceEndpointAzureRMOutput() ServiceEndpointAzureRMOutput {
	return o
}

func (o ServiceEndpointAzureRMOutput) ToServiceEndpointAzureRMOutputWithContext(ctx context.Context) ServiceEndpointAzureRMOutput {
	return o
}

func (o ServiceEndpointAzureRMOutput) ToServiceEndpointAzureRMPtrOutput() ServiceEndpointAzureRMPtrOutput {
	return o.ToServiceEndpointAzureRMPtrOutputWithContext(context.Background())
}

func (o ServiceEndpointAzureRMOutput) ToServiceEndpointAzureRMPtrOutputWithContext(ctx context.Context) ServiceEndpointAzureRMPtrOutput {
	return o.ApplyT(func(v ServiceEndpointAzureRM) *ServiceEndpointAzureRM {
		return &v
	}).(ServiceEndpointAzureRMPtrOutput)
}

type ServiceEndpointAzureRMPtrOutput struct {
	*pulumi.OutputState
}

func (ServiceEndpointAzureRMPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceEndpointAzureRM)(nil))
}

func (o ServiceEndpointAzureRMPtrOutput) ToServiceEndpointAzureRMPtrOutput() ServiceEndpointAzureRMPtrOutput {
	return o
}

func (o ServiceEndpointAzureRMPtrOutput) ToServiceEndpointAzureRMPtrOutputWithContext(ctx context.Context) ServiceEndpointAzureRMPtrOutput {
	return o
}

type ServiceEndpointAzureRMArrayOutput struct{ *pulumi.OutputState }

func (ServiceEndpointAzureRMArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceEndpointAzureRM)(nil))
}

func (o ServiceEndpointAzureRMArrayOutput) ToServiceEndpointAzureRMArrayOutput() ServiceEndpointAzureRMArrayOutput {
	return o
}

func (o ServiceEndpointAzureRMArrayOutput) ToServiceEndpointAzureRMArrayOutputWithContext(ctx context.Context) ServiceEndpointAzureRMArrayOutput {
	return o
}

func (o ServiceEndpointAzureRMArrayOutput) Index(i pulumi.IntInput) ServiceEndpointAzureRMOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServiceEndpointAzureRM {
		return vs[0].([]ServiceEndpointAzureRM)[vs[1].(int)]
	}).(ServiceEndpointAzureRMOutput)
}

type ServiceEndpointAzureRMMapOutput struct{ *pulumi.OutputState }

func (ServiceEndpointAzureRMMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ServiceEndpointAzureRM)(nil))
}

func (o ServiceEndpointAzureRMMapOutput) ToServiceEndpointAzureRMMapOutput() ServiceEndpointAzureRMMapOutput {
	return o
}

func (o ServiceEndpointAzureRMMapOutput) ToServiceEndpointAzureRMMapOutputWithContext(ctx context.Context) ServiceEndpointAzureRMMapOutput {
	return o
}

func (o ServiceEndpointAzureRMMapOutput) MapIndex(k pulumi.StringInput) ServiceEndpointAzureRMOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ServiceEndpointAzureRM {
		return vs[0].(map[string]ServiceEndpointAzureRM)[vs[1].(string)]
	}).(ServiceEndpointAzureRMOutput)
}

func init() {
	pulumi.RegisterOutputType(ServiceEndpointAzureRMOutput{})
	pulumi.RegisterOutputType(ServiceEndpointAzureRMPtrOutput{})
	pulumi.RegisterOutputType(ServiceEndpointAzureRMArrayOutput{})
	pulumi.RegisterOutputType(ServiceEndpointAzureRMMapOutput{})
}
