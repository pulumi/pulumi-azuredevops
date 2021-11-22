// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package serviceendpoint

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
// 			Credentials: &ServiceEndpointAzureRMCredentialsArgs{
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
//  $ pulumi import azuredevops:ServiceEndpoint/azureRM:AzureRM serviceendpoint 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
// ```
//
// Deprecated: azuredevops.serviceendpoint.AzureRM has been deprecated in favor of azuredevops.ServiceEndpointAzureRM
type AzureRM struct {
	pulumi.CustomResourceState

	Authorization pulumi.StringMapOutput `pulumi:"authorization"`
	// The tenant id if the service principal.
	AzurermSpnTenantid pulumi.StringOutput `pulumi:"azurermSpnTenantid"`
	// The subscription Id of the Azure targets.
	AzurermSubscriptionId pulumi.StringOutput `pulumi:"azurermSubscriptionId"`
	// The subscription Name of the targets.
	AzurermSubscriptionName pulumi.StringOutput `pulumi:"azurermSubscriptionName"`
	// A `credentials` block.
	Credentials AzureRMCredentialsPtrOutput `pulumi:"credentials"`
	// Service connection description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The project ID or project name.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The resource group used for scope of automatic service endpoint.
	ResourceGroup pulumi.StringPtrOutput `pulumi:"resourceGroup"`
	// The Service Endpoint name.
	ServiceEndpointName pulumi.StringOutput `pulumi:"serviceEndpointName"`
}

// NewAzureRM registers a new resource with the given unique name, arguments, and options.
func NewAzureRM(ctx *pulumi.Context,
	name string, args *AzureRMArgs, opts ...pulumi.ResourceOption) (*AzureRM, error) {
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
	var resource AzureRM
	err := ctx.RegisterResource("azuredevops:ServiceEndpoint/azureRM:AzureRM", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAzureRM gets an existing AzureRM resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAzureRM(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AzureRMState, opts ...pulumi.ResourceOption) (*AzureRM, error) {
	var resource AzureRM
	err := ctx.ReadResource("azuredevops:ServiceEndpoint/azureRM:AzureRM", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AzureRM resources.
type azureRMState struct {
	Authorization map[string]string `pulumi:"authorization"`
	// The tenant id if the service principal.
	AzurermSpnTenantid *string `pulumi:"azurermSpnTenantid"`
	// The subscription Id of the Azure targets.
	AzurermSubscriptionId *string `pulumi:"azurermSubscriptionId"`
	// The subscription Name of the targets.
	AzurermSubscriptionName *string `pulumi:"azurermSubscriptionName"`
	// A `credentials` block.
	Credentials *AzureRMCredentials `pulumi:"credentials"`
	// Service connection description.
	Description *string `pulumi:"description"`
	// The project ID or project name.
	ProjectId *string `pulumi:"projectId"`
	// The resource group used for scope of automatic service endpoint.
	ResourceGroup *string `pulumi:"resourceGroup"`
	// The Service Endpoint name.
	ServiceEndpointName *string `pulumi:"serviceEndpointName"`
}

type AzureRMState struct {
	Authorization pulumi.StringMapInput
	// The tenant id if the service principal.
	AzurermSpnTenantid pulumi.StringPtrInput
	// The subscription Id of the Azure targets.
	AzurermSubscriptionId pulumi.StringPtrInput
	// The subscription Name of the targets.
	AzurermSubscriptionName pulumi.StringPtrInput
	// A `credentials` block.
	Credentials AzureRMCredentialsPtrInput
	// Service connection description.
	Description pulumi.StringPtrInput
	// The project ID or project name.
	ProjectId pulumi.StringPtrInput
	// The resource group used for scope of automatic service endpoint.
	ResourceGroup pulumi.StringPtrInput
	// The Service Endpoint name.
	ServiceEndpointName pulumi.StringPtrInput
}

func (AzureRMState) ElementType() reflect.Type {
	return reflect.TypeOf((*azureRMState)(nil)).Elem()
}

type azureRMArgs struct {
	Authorization map[string]string `pulumi:"authorization"`
	// The tenant id if the service principal.
	AzurermSpnTenantid string `pulumi:"azurermSpnTenantid"`
	// The subscription Id of the Azure targets.
	AzurermSubscriptionId string `pulumi:"azurermSubscriptionId"`
	// The subscription Name of the targets.
	AzurermSubscriptionName string `pulumi:"azurermSubscriptionName"`
	// A `credentials` block.
	Credentials *AzureRMCredentials `pulumi:"credentials"`
	// Service connection description.
	Description *string `pulumi:"description"`
	// The project ID or project name.
	ProjectId string `pulumi:"projectId"`
	// The resource group used for scope of automatic service endpoint.
	ResourceGroup *string `pulumi:"resourceGroup"`
	// The Service Endpoint name.
	ServiceEndpointName string `pulumi:"serviceEndpointName"`
}

// The set of arguments for constructing a AzureRM resource.
type AzureRMArgs struct {
	Authorization pulumi.StringMapInput
	// The tenant id if the service principal.
	AzurermSpnTenantid pulumi.StringInput
	// The subscription Id of the Azure targets.
	AzurermSubscriptionId pulumi.StringInput
	// The subscription Name of the targets.
	AzurermSubscriptionName pulumi.StringInput
	// A `credentials` block.
	Credentials AzureRMCredentialsPtrInput
	// Service connection description.
	Description pulumi.StringPtrInput
	// The project ID or project name.
	ProjectId pulumi.StringInput
	// The resource group used for scope of automatic service endpoint.
	ResourceGroup pulumi.StringPtrInput
	// The Service Endpoint name.
	ServiceEndpointName pulumi.StringInput
}

func (AzureRMArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*azureRMArgs)(nil)).Elem()
}

type AzureRMInput interface {
	pulumi.Input

	ToAzureRMOutput() AzureRMOutput
	ToAzureRMOutputWithContext(ctx context.Context) AzureRMOutput
}

func (*AzureRM) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureRM)(nil))
}

func (i *AzureRM) ToAzureRMOutput() AzureRMOutput {
	return i.ToAzureRMOutputWithContext(context.Background())
}

func (i *AzureRM) ToAzureRMOutputWithContext(ctx context.Context) AzureRMOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureRMOutput)
}

func (i *AzureRM) ToAzureRMPtrOutput() AzureRMPtrOutput {
	return i.ToAzureRMPtrOutputWithContext(context.Background())
}

func (i *AzureRM) ToAzureRMPtrOutputWithContext(ctx context.Context) AzureRMPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureRMPtrOutput)
}

type AzureRMPtrInput interface {
	pulumi.Input

	ToAzureRMPtrOutput() AzureRMPtrOutput
	ToAzureRMPtrOutputWithContext(ctx context.Context) AzureRMPtrOutput
}

type azureRMPtrType AzureRMArgs

func (*azureRMPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureRM)(nil))
}

func (i *azureRMPtrType) ToAzureRMPtrOutput() AzureRMPtrOutput {
	return i.ToAzureRMPtrOutputWithContext(context.Background())
}

func (i *azureRMPtrType) ToAzureRMPtrOutputWithContext(ctx context.Context) AzureRMPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureRMPtrOutput)
}

// AzureRMArrayInput is an input type that accepts AzureRMArray and AzureRMArrayOutput values.
// You can construct a concrete instance of `AzureRMArrayInput` via:
//
//          AzureRMArray{ AzureRMArgs{...} }
type AzureRMArrayInput interface {
	pulumi.Input

	ToAzureRMArrayOutput() AzureRMArrayOutput
	ToAzureRMArrayOutputWithContext(context.Context) AzureRMArrayOutput
}

type AzureRMArray []AzureRMInput

func (AzureRMArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AzureRM)(nil)).Elem()
}

func (i AzureRMArray) ToAzureRMArrayOutput() AzureRMArrayOutput {
	return i.ToAzureRMArrayOutputWithContext(context.Background())
}

func (i AzureRMArray) ToAzureRMArrayOutputWithContext(ctx context.Context) AzureRMArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureRMArrayOutput)
}

// AzureRMMapInput is an input type that accepts AzureRMMap and AzureRMMapOutput values.
// You can construct a concrete instance of `AzureRMMapInput` via:
//
//          AzureRMMap{ "key": AzureRMArgs{...} }
type AzureRMMapInput interface {
	pulumi.Input

	ToAzureRMMapOutput() AzureRMMapOutput
	ToAzureRMMapOutputWithContext(context.Context) AzureRMMapOutput
}

type AzureRMMap map[string]AzureRMInput

func (AzureRMMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AzureRM)(nil)).Elem()
}

func (i AzureRMMap) ToAzureRMMapOutput() AzureRMMapOutput {
	return i.ToAzureRMMapOutputWithContext(context.Background())
}

func (i AzureRMMap) ToAzureRMMapOutputWithContext(ctx context.Context) AzureRMMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureRMMapOutput)
}

type AzureRMOutput struct{ *pulumi.OutputState }

func (AzureRMOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureRM)(nil))
}

func (o AzureRMOutput) ToAzureRMOutput() AzureRMOutput {
	return o
}

func (o AzureRMOutput) ToAzureRMOutputWithContext(ctx context.Context) AzureRMOutput {
	return o
}

func (o AzureRMOutput) ToAzureRMPtrOutput() AzureRMPtrOutput {
	return o.ToAzureRMPtrOutputWithContext(context.Background())
}

func (o AzureRMOutput) ToAzureRMPtrOutputWithContext(ctx context.Context) AzureRMPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureRM) *AzureRM {
		return &v
	}).(AzureRMPtrOutput)
}

type AzureRMPtrOutput struct{ *pulumi.OutputState }

func (AzureRMPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureRM)(nil))
}

func (o AzureRMPtrOutput) ToAzureRMPtrOutput() AzureRMPtrOutput {
	return o
}

func (o AzureRMPtrOutput) ToAzureRMPtrOutputWithContext(ctx context.Context) AzureRMPtrOutput {
	return o
}

func (o AzureRMPtrOutput) Elem() AzureRMOutput {
	return o.ApplyT(func(v *AzureRM) AzureRM {
		if v != nil {
			return *v
		}
		var ret AzureRM
		return ret
	}).(AzureRMOutput)
}

type AzureRMArrayOutput struct{ *pulumi.OutputState }

func (AzureRMArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AzureRM)(nil))
}

func (o AzureRMArrayOutput) ToAzureRMArrayOutput() AzureRMArrayOutput {
	return o
}

func (o AzureRMArrayOutput) ToAzureRMArrayOutputWithContext(ctx context.Context) AzureRMArrayOutput {
	return o
}

func (o AzureRMArrayOutput) Index(i pulumi.IntInput) AzureRMOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AzureRM {
		return vs[0].([]AzureRM)[vs[1].(int)]
	}).(AzureRMOutput)
}

type AzureRMMapOutput struct{ *pulumi.OutputState }

func (AzureRMMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]AzureRM)(nil))
}

func (o AzureRMMapOutput) ToAzureRMMapOutput() AzureRMMapOutput {
	return o
}

func (o AzureRMMapOutput) ToAzureRMMapOutputWithContext(ctx context.Context) AzureRMMapOutput {
	return o
}

func (o AzureRMMapOutput) MapIndex(k pulumi.StringInput) AzureRMOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) AzureRM {
		return vs[0].(map[string]AzureRM)[vs[1].(string)]
	}).(AzureRMOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AzureRMInput)(nil)).Elem(), &AzureRM{})
	pulumi.RegisterInputType(reflect.TypeOf((*AzureRMPtrInput)(nil)).Elem(), &AzureRM{})
	pulumi.RegisterInputType(reflect.TypeOf((*AzureRMArrayInput)(nil)).Elem(), AzureRMArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AzureRMMapInput)(nil)).Elem(), AzureRMMap{})
	pulumi.RegisterOutputType(AzureRMOutput{})
	pulumi.RegisterOutputType(AzureRMPtrOutput{})
	pulumi.RegisterOutputType(AzureRMArrayOutput{})
	pulumi.RegisterOutputType(AzureRMMapOutput{})
}
