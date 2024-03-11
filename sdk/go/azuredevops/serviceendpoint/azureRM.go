// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package serviceendpoint

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azuredevops/sdk/v2/go/azuredevops/internal"
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
//
// ### Service Principal Manual AzureRM Service Endpoint (Subscription Scoped)
//
// <!--Start PulumiCodeChooser -->
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
//			_, err = azuredevops.NewServiceEndpointAzureRM(ctx, "exampleServiceEndpointAzureRM", &azuredevops.ServiceEndpointAzureRMArgs{
//				ProjectId:                           exampleProject.ID(),
//				ServiceEndpointName:                 pulumi.String("Example AzureRM"),
//				Description:                         pulumi.String("Managed by Terraform"),
//				ServiceEndpointAuthenticationScheme: pulumi.String("ServicePrincipal"),
//				Credentials: &azuredevops.ServiceEndpointAzureRMCredentialsArgs{
//					Serviceprincipalid:  pulumi.String("00000000-0000-0000-0000-000000000000"),
//					Serviceprincipalkey: pulumi.String("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
//				},
//				AzurermSpnTenantid:      pulumi.String("00000000-0000-0000-0000-000000000000"),
//				AzurermSubscriptionId:   pulumi.String("00000000-0000-0000-0000-000000000000"),
//				AzurermSubscriptionName: pulumi.String("Example Subscription Name"),
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
// ### Service Principal Manual AzureRM Service Endpoint (ManagementGroup Scoped)
//
// <!--Start PulumiCodeChooser -->
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
//			_, err = azuredevops.NewServiceEndpointAzureRM(ctx, "exampleServiceEndpointAzureRM", &azuredevops.ServiceEndpointAzureRMArgs{
//				ProjectId:                           exampleProject.ID(),
//				ServiceEndpointName:                 pulumi.String("Example AzureRM"),
//				Description:                         pulumi.String("Managed by Terraform"),
//				ServiceEndpointAuthenticationScheme: pulumi.String("ServicePrincipal"),
//				Credentials: &azuredevops.ServiceEndpointAzureRMCredentialsArgs{
//					Serviceprincipalid:  pulumi.String("00000000-0000-0000-0000-000000000000"),
//					Serviceprincipalkey: pulumi.String("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
//				},
//				AzurermSpnTenantid:         pulumi.String("00000000-0000-0000-0000-000000000000"),
//				AzurermManagementGroupId:   pulumi.String("managementGroup"),
//				AzurermManagementGroupName: pulumi.String("managementGroup"),
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
// ### Service Principal Automatic AzureRM Service Endpoint
//
// <!--Start PulumiCodeChooser -->
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
//			})
//			if err != nil {
//				return err
//			}
//			_, err = azuredevops.NewServiceEndpointAzureRM(ctx, "exampleServiceEndpointAzureRM", &azuredevops.ServiceEndpointAzureRMArgs{
//				ProjectId:                           exampleProject.ID(),
//				ServiceEndpointName:                 pulumi.String("Example AzureRM"),
//				ServiceEndpointAuthenticationScheme: pulumi.String("ServicePrincipal"),
//				AzurermSpnTenantid:                  pulumi.String("00000000-0000-0000-0000-000000000000"),
//				AzurermSubscriptionId:               pulumi.String("00000000-0000-0000-0000-000000000000"),
//				AzurermSubscriptionName:             pulumi.String("Example Subscription Name"),
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
// ### Workload Identity Federation Automatic AzureRM Service Endpoint
//
// <!--Start PulumiCodeChooser -->
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
//			})
//			if err != nil {
//				return err
//			}
//			_, err = azuredevops.NewServiceEndpointAzureRM(ctx, "exampleServiceEndpointAzureRM", &azuredevops.ServiceEndpointAzureRMArgs{
//				ProjectId:                           exampleProject.ID(),
//				ServiceEndpointName:                 pulumi.String("Example AzureRM"),
//				ServiceEndpointAuthenticationScheme: pulumi.String("WorkloadIdentityFederation"),
//				AzurermSpnTenantid:                  pulumi.String("00000000-0000-0000-0000-000000000000"),
//				AzurermSubscriptionId:               pulumi.String("00000000-0000-0000-0000-000000000000"),
//				AzurermSubscriptionName:             pulumi.String("Example Subscription Name"),
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
// ### Managed Identity AzureRM Service Endpoint
//
// <!--Start PulumiCodeChooser -->
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
//			})
//			if err != nil {
//				return err
//			}
//			_, err = azuredevops.NewServiceEndpointAzureRM(ctx, "exampleServiceEndpointAzureRM", &azuredevops.ServiceEndpointAzureRMArgs{
//				ProjectId:                           exampleProject.ID(),
//				ServiceEndpointName:                 pulumi.String("Example AzureRM"),
//				ServiceEndpointAuthenticationScheme: pulumi.String("ManagedServiceIdentity"),
//				AzurermSpnTenantid:                  pulumi.String("00000000-0000-0000-0000-000000000000"),
//				AzurermSubscriptionId:               pulumi.String("00000000-0000-0000-0000-000000000000"),
//				AzurermSubscriptionName:             pulumi.String("Example Subscription Name"),
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
// - [Azure DevOps Service REST API 7.0 - Service End points](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-7.0)
//
// ## Import
//
// Azure DevOps Service Endpoint Azure Resource Manage can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**
//
// ```sh
// $ pulumi import azuredevops:ServiceEndpoint/azureRM:AzureRM example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
// ```
//
// Deprecated: azuredevops.serviceendpoint.AzureRM has been deprecated in favor of azuredevops.ServiceEndpointAzureRM
type AzureRM struct {
	pulumi.CustomResourceState

	Authorization pulumi.StringMapOutput `pulumi:"authorization"`
	// The Management group ID of the Azure targets.
	AzurermManagementGroupId pulumi.StringPtrOutput `pulumi:"azurermManagementGroupId"`
	// The Management group Name of the targets.
	AzurermManagementGroupName pulumi.StringPtrOutput `pulumi:"azurermManagementGroupName"`
	// The Tenant ID if the service principal.
	AzurermSpnTenantid pulumi.StringOutput `pulumi:"azurermSpnTenantid"`
	// The Subscription ID of the Azure targets.
	AzurermSubscriptionId pulumi.StringPtrOutput `pulumi:"azurermSubscriptionId"`
	// The Subscription Name of the targets.
	AzurermSubscriptionName pulumi.StringPtrOutput `pulumi:"azurermSubscriptionName"`
	// A `credentials` block.
	Credentials AzureRMCredentialsPtrOutput `pulumi:"credentials"`
	// Service connection description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The Cloud Environment to use. Defaults to `AzureCloud`. Possible values are `AzureCloud`, `AzureChinaCloud`. Changing this forces a new resource to be created.
	//
	// > **NOTE:** One of either `Subscription` scoped i.e. `azurermSubscriptionId`, `azurermSubscriptionName` or `ManagementGroup` scoped i.e. `azurermManagementGroupId`, `azurermManagementGroupName` values must be specified.
	Environment pulumi.StringPtrOutput `pulumi:"environment"`
	// A `features` block.
	Features AzureRMFeaturesPtrOutput `pulumi:"features"`
	// The ID of the project.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The resource group used for scope of automatic service endpoint.
	ResourceGroup pulumi.StringPtrOutput `pulumi:"resourceGroup"`
	// Specifies the type of azurerm endpoint, either `WorkloadIdentityFederation`, `ManagedServiceIdentity` or `ServicePrincipal`. Defaults to `ServicePrincipal` for backwards compatibility.
	//
	// > **NOTE:** The `WorkloadIdentityFederation` authentication scheme is currently in private preview. Your organisation must be part of the preview and the feature toggle must be turned on to use it. More details can be found [here](https://aka.ms/azdo-rm-workload-identity).
	ServiceEndpointAuthenticationScheme pulumi.StringPtrOutput `pulumi:"serviceEndpointAuthenticationScheme"`
	// The Service Endpoint Name.
	ServiceEndpointName pulumi.StringOutput `pulumi:"serviceEndpointName"`
	// The Application(Client) ID of the Service Principal.
	ServicePrincipalId pulumi.StringOutput `pulumi:"servicePrincipalId"`
	// The issuer if `serviceEndpointAuthenticationScheme` is set to `WorkloadIdentityFederation`. This looks like `https://vstoken.dev.azure.com/00000000-0000-0000-0000-000000000000`, where the GUID is the Organization ID of your Azure DevOps Organisation.
	WorkloadIdentityFederationIssuer pulumi.StringOutput `pulumi:"workloadIdentityFederationIssuer"`
	// The subject if `serviceEndpointAuthenticationScheme` is set to `WorkloadIdentityFederation`. This looks like `sc://<organisation>/<project>/<service-connection-name>`.
	WorkloadIdentityFederationSubject pulumi.StringOutput `pulumi:"workloadIdentityFederationSubject"`
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
	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.ServiceEndpointName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceEndpointName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
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
	// The Management group ID of the Azure targets.
	AzurermManagementGroupId *string `pulumi:"azurermManagementGroupId"`
	// The Management group Name of the targets.
	AzurermManagementGroupName *string `pulumi:"azurermManagementGroupName"`
	// The Tenant ID if the service principal.
	AzurermSpnTenantid *string `pulumi:"azurermSpnTenantid"`
	// The Subscription ID of the Azure targets.
	AzurermSubscriptionId *string `pulumi:"azurermSubscriptionId"`
	// The Subscription Name of the targets.
	AzurermSubscriptionName *string `pulumi:"azurermSubscriptionName"`
	// A `credentials` block.
	Credentials *AzureRMCredentials `pulumi:"credentials"`
	// Service connection description.
	Description *string `pulumi:"description"`
	// The Cloud Environment to use. Defaults to `AzureCloud`. Possible values are `AzureCloud`, `AzureChinaCloud`. Changing this forces a new resource to be created.
	//
	// > **NOTE:** One of either `Subscription` scoped i.e. `azurermSubscriptionId`, `azurermSubscriptionName` or `ManagementGroup` scoped i.e. `azurermManagementGroupId`, `azurermManagementGroupName` values must be specified.
	Environment *string `pulumi:"environment"`
	// A `features` block.
	Features *AzureRMFeatures `pulumi:"features"`
	// The ID of the project.
	ProjectId *string `pulumi:"projectId"`
	// The resource group used for scope of automatic service endpoint.
	ResourceGroup *string `pulumi:"resourceGroup"`
	// Specifies the type of azurerm endpoint, either `WorkloadIdentityFederation`, `ManagedServiceIdentity` or `ServicePrincipal`. Defaults to `ServicePrincipal` for backwards compatibility.
	//
	// > **NOTE:** The `WorkloadIdentityFederation` authentication scheme is currently in private preview. Your organisation must be part of the preview and the feature toggle must be turned on to use it. More details can be found [here](https://aka.ms/azdo-rm-workload-identity).
	ServiceEndpointAuthenticationScheme *string `pulumi:"serviceEndpointAuthenticationScheme"`
	// The Service Endpoint Name.
	ServiceEndpointName *string `pulumi:"serviceEndpointName"`
	// The Application(Client) ID of the Service Principal.
	ServicePrincipalId *string `pulumi:"servicePrincipalId"`
	// The issuer if `serviceEndpointAuthenticationScheme` is set to `WorkloadIdentityFederation`. This looks like `https://vstoken.dev.azure.com/00000000-0000-0000-0000-000000000000`, where the GUID is the Organization ID of your Azure DevOps Organisation.
	WorkloadIdentityFederationIssuer *string `pulumi:"workloadIdentityFederationIssuer"`
	// The subject if `serviceEndpointAuthenticationScheme` is set to `WorkloadIdentityFederation`. This looks like `sc://<organisation>/<project>/<service-connection-name>`.
	WorkloadIdentityFederationSubject *string `pulumi:"workloadIdentityFederationSubject"`
}

type AzureRMState struct {
	Authorization pulumi.StringMapInput
	// The Management group ID of the Azure targets.
	AzurermManagementGroupId pulumi.StringPtrInput
	// The Management group Name of the targets.
	AzurermManagementGroupName pulumi.StringPtrInput
	// The Tenant ID if the service principal.
	AzurermSpnTenantid pulumi.StringPtrInput
	// The Subscription ID of the Azure targets.
	AzurermSubscriptionId pulumi.StringPtrInput
	// The Subscription Name of the targets.
	AzurermSubscriptionName pulumi.StringPtrInput
	// A `credentials` block.
	Credentials AzureRMCredentialsPtrInput
	// Service connection description.
	Description pulumi.StringPtrInput
	// The Cloud Environment to use. Defaults to `AzureCloud`. Possible values are `AzureCloud`, `AzureChinaCloud`. Changing this forces a new resource to be created.
	//
	// > **NOTE:** One of either `Subscription` scoped i.e. `azurermSubscriptionId`, `azurermSubscriptionName` or `ManagementGroup` scoped i.e. `azurermManagementGroupId`, `azurermManagementGroupName` values must be specified.
	Environment pulumi.StringPtrInput
	// A `features` block.
	Features AzureRMFeaturesPtrInput
	// The ID of the project.
	ProjectId pulumi.StringPtrInput
	// The resource group used for scope of automatic service endpoint.
	ResourceGroup pulumi.StringPtrInput
	// Specifies the type of azurerm endpoint, either `WorkloadIdentityFederation`, `ManagedServiceIdentity` or `ServicePrincipal`. Defaults to `ServicePrincipal` for backwards compatibility.
	//
	// > **NOTE:** The `WorkloadIdentityFederation` authentication scheme is currently in private preview. Your organisation must be part of the preview and the feature toggle must be turned on to use it. More details can be found [here](https://aka.ms/azdo-rm-workload-identity).
	ServiceEndpointAuthenticationScheme pulumi.StringPtrInput
	// The Service Endpoint Name.
	ServiceEndpointName pulumi.StringPtrInput
	// The Application(Client) ID of the Service Principal.
	ServicePrincipalId pulumi.StringPtrInput
	// The issuer if `serviceEndpointAuthenticationScheme` is set to `WorkloadIdentityFederation`. This looks like `https://vstoken.dev.azure.com/00000000-0000-0000-0000-000000000000`, where the GUID is the Organization ID of your Azure DevOps Organisation.
	WorkloadIdentityFederationIssuer pulumi.StringPtrInput
	// The subject if `serviceEndpointAuthenticationScheme` is set to `WorkloadIdentityFederation`. This looks like `sc://<organisation>/<project>/<service-connection-name>`.
	WorkloadIdentityFederationSubject pulumi.StringPtrInput
}

func (AzureRMState) ElementType() reflect.Type {
	return reflect.TypeOf((*azureRMState)(nil)).Elem()
}

type azureRMArgs struct {
	Authorization map[string]string `pulumi:"authorization"`
	// The Management group ID of the Azure targets.
	AzurermManagementGroupId *string `pulumi:"azurermManagementGroupId"`
	// The Management group Name of the targets.
	AzurermManagementGroupName *string `pulumi:"azurermManagementGroupName"`
	// The Tenant ID if the service principal.
	AzurermSpnTenantid string `pulumi:"azurermSpnTenantid"`
	// The Subscription ID of the Azure targets.
	AzurermSubscriptionId *string `pulumi:"azurermSubscriptionId"`
	// The Subscription Name of the targets.
	AzurermSubscriptionName *string `pulumi:"azurermSubscriptionName"`
	// A `credentials` block.
	Credentials *AzureRMCredentials `pulumi:"credentials"`
	// Service connection description.
	Description *string `pulumi:"description"`
	// The Cloud Environment to use. Defaults to `AzureCloud`. Possible values are `AzureCloud`, `AzureChinaCloud`. Changing this forces a new resource to be created.
	//
	// > **NOTE:** One of either `Subscription` scoped i.e. `azurermSubscriptionId`, `azurermSubscriptionName` or `ManagementGroup` scoped i.e. `azurermManagementGroupId`, `azurermManagementGroupName` values must be specified.
	Environment *string `pulumi:"environment"`
	// A `features` block.
	Features *AzureRMFeatures `pulumi:"features"`
	// The ID of the project.
	ProjectId string `pulumi:"projectId"`
	// The resource group used for scope of automatic service endpoint.
	ResourceGroup *string `pulumi:"resourceGroup"`
	// Specifies the type of azurerm endpoint, either `WorkloadIdentityFederation`, `ManagedServiceIdentity` or `ServicePrincipal`. Defaults to `ServicePrincipal` for backwards compatibility.
	//
	// > **NOTE:** The `WorkloadIdentityFederation` authentication scheme is currently in private preview. Your organisation must be part of the preview and the feature toggle must be turned on to use it. More details can be found [here](https://aka.ms/azdo-rm-workload-identity).
	ServiceEndpointAuthenticationScheme *string `pulumi:"serviceEndpointAuthenticationScheme"`
	// The Service Endpoint Name.
	ServiceEndpointName string `pulumi:"serviceEndpointName"`
}

// The set of arguments for constructing a AzureRM resource.
type AzureRMArgs struct {
	Authorization pulumi.StringMapInput
	// The Management group ID of the Azure targets.
	AzurermManagementGroupId pulumi.StringPtrInput
	// The Management group Name of the targets.
	AzurermManagementGroupName pulumi.StringPtrInput
	// The Tenant ID if the service principal.
	AzurermSpnTenantid pulumi.StringInput
	// The Subscription ID of the Azure targets.
	AzurermSubscriptionId pulumi.StringPtrInput
	// The Subscription Name of the targets.
	AzurermSubscriptionName pulumi.StringPtrInput
	// A `credentials` block.
	Credentials AzureRMCredentialsPtrInput
	// Service connection description.
	Description pulumi.StringPtrInput
	// The Cloud Environment to use. Defaults to `AzureCloud`. Possible values are `AzureCloud`, `AzureChinaCloud`. Changing this forces a new resource to be created.
	//
	// > **NOTE:** One of either `Subscription` scoped i.e. `azurermSubscriptionId`, `azurermSubscriptionName` or `ManagementGroup` scoped i.e. `azurermManagementGroupId`, `azurermManagementGroupName` values must be specified.
	Environment pulumi.StringPtrInput
	// A `features` block.
	Features AzureRMFeaturesPtrInput
	// The ID of the project.
	ProjectId pulumi.StringInput
	// The resource group used for scope of automatic service endpoint.
	ResourceGroup pulumi.StringPtrInput
	// Specifies the type of azurerm endpoint, either `WorkloadIdentityFederation`, `ManagedServiceIdentity` or `ServicePrincipal`. Defaults to `ServicePrincipal` for backwards compatibility.
	//
	// > **NOTE:** The `WorkloadIdentityFederation` authentication scheme is currently in private preview. Your organisation must be part of the preview and the feature toggle must be turned on to use it. More details can be found [here](https://aka.ms/azdo-rm-workload-identity).
	ServiceEndpointAuthenticationScheme pulumi.StringPtrInput
	// The Service Endpoint Name.
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
	return reflect.TypeOf((**AzureRM)(nil)).Elem()
}

func (i *AzureRM) ToAzureRMOutput() AzureRMOutput {
	return i.ToAzureRMOutputWithContext(context.Background())
}

func (i *AzureRM) ToAzureRMOutputWithContext(ctx context.Context) AzureRMOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureRMOutput)
}

// AzureRMArrayInput is an input type that accepts AzureRMArray and AzureRMArrayOutput values.
// You can construct a concrete instance of `AzureRMArrayInput` via:
//
//	AzureRMArray{ AzureRMArgs{...} }
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
//	AzureRMMap{ "key": AzureRMArgs{...} }
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
	return reflect.TypeOf((**AzureRM)(nil)).Elem()
}

func (o AzureRMOutput) ToAzureRMOutput() AzureRMOutput {
	return o
}

func (o AzureRMOutput) ToAzureRMOutputWithContext(ctx context.Context) AzureRMOutput {
	return o
}

func (o AzureRMOutput) Authorization() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AzureRM) pulumi.StringMapOutput { return v.Authorization }).(pulumi.StringMapOutput)
}

// The Management group ID of the Azure targets.
func (o AzureRMOutput) AzurermManagementGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureRM) pulumi.StringPtrOutput { return v.AzurermManagementGroupId }).(pulumi.StringPtrOutput)
}

// The Management group Name of the targets.
func (o AzureRMOutput) AzurermManagementGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureRM) pulumi.StringPtrOutput { return v.AzurermManagementGroupName }).(pulumi.StringPtrOutput)
}

// The Tenant ID if the service principal.
func (o AzureRMOutput) AzurermSpnTenantid() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureRM) pulumi.StringOutput { return v.AzurermSpnTenantid }).(pulumi.StringOutput)
}

// The Subscription ID of the Azure targets.
func (o AzureRMOutput) AzurermSubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureRM) pulumi.StringPtrOutput { return v.AzurermSubscriptionId }).(pulumi.StringPtrOutput)
}

// The Subscription Name of the targets.
func (o AzureRMOutput) AzurermSubscriptionName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureRM) pulumi.StringPtrOutput { return v.AzurermSubscriptionName }).(pulumi.StringPtrOutput)
}

// A `credentials` block.
func (o AzureRMOutput) Credentials() AzureRMCredentialsPtrOutput {
	return o.ApplyT(func(v *AzureRM) AzureRMCredentialsPtrOutput { return v.Credentials }).(AzureRMCredentialsPtrOutput)
}

// Service connection description.
func (o AzureRMOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureRM) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The Cloud Environment to use. Defaults to `AzureCloud`. Possible values are `AzureCloud`, `AzureChinaCloud`. Changing this forces a new resource to be created.
//
// > **NOTE:** One of either `Subscription` scoped i.e. `azurermSubscriptionId`, `azurermSubscriptionName` or `ManagementGroup` scoped i.e. `azurermManagementGroupId`, `azurermManagementGroupName` values must be specified.
func (o AzureRMOutput) Environment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureRM) pulumi.StringPtrOutput { return v.Environment }).(pulumi.StringPtrOutput)
}

// A `features` block.
func (o AzureRMOutput) Features() AzureRMFeaturesPtrOutput {
	return o.ApplyT(func(v *AzureRM) AzureRMFeaturesPtrOutput { return v.Features }).(AzureRMFeaturesPtrOutput)
}

// The ID of the project.
func (o AzureRMOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureRM) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The resource group used for scope of automatic service endpoint.
func (o AzureRMOutput) ResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureRM) pulumi.StringPtrOutput { return v.ResourceGroup }).(pulumi.StringPtrOutput)
}

// Specifies the type of azurerm endpoint, either `WorkloadIdentityFederation`, `ManagedServiceIdentity` or `ServicePrincipal`. Defaults to `ServicePrincipal` for backwards compatibility.
//
// > **NOTE:** The `WorkloadIdentityFederation` authentication scheme is currently in private preview. Your organisation must be part of the preview and the feature toggle must be turned on to use it. More details can be found [here](https://aka.ms/azdo-rm-workload-identity).
func (o AzureRMOutput) ServiceEndpointAuthenticationScheme() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureRM) pulumi.StringPtrOutput { return v.ServiceEndpointAuthenticationScheme }).(pulumi.StringPtrOutput)
}

// The Service Endpoint Name.
func (o AzureRMOutput) ServiceEndpointName() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureRM) pulumi.StringOutput { return v.ServiceEndpointName }).(pulumi.StringOutput)
}

// The Application(Client) ID of the Service Principal.
func (o AzureRMOutput) ServicePrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureRM) pulumi.StringOutput { return v.ServicePrincipalId }).(pulumi.StringOutput)
}

// The issuer if `serviceEndpointAuthenticationScheme` is set to `WorkloadIdentityFederation`. This looks like `https://vstoken.dev.azure.com/00000000-0000-0000-0000-000000000000`, where the GUID is the Organization ID of your Azure DevOps Organisation.
func (o AzureRMOutput) WorkloadIdentityFederationIssuer() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureRM) pulumi.StringOutput { return v.WorkloadIdentityFederationIssuer }).(pulumi.StringOutput)
}

// The subject if `serviceEndpointAuthenticationScheme` is set to `WorkloadIdentityFederation`. This looks like `sc://<organisation>/<project>/<service-connection-name>`.
func (o AzureRMOutput) WorkloadIdentityFederationSubject() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureRM) pulumi.StringOutput { return v.WorkloadIdentityFederationSubject }).(pulumi.StringOutput)
}

type AzureRMArrayOutput struct{ *pulumi.OutputState }

func (AzureRMArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AzureRM)(nil)).Elem()
}

func (o AzureRMArrayOutput) ToAzureRMArrayOutput() AzureRMArrayOutput {
	return o
}

func (o AzureRMArrayOutput) ToAzureRMArrayOutputWithContext(ctx context.Context) AzureRMArrayOutput {
	return o
}

func (o AzureRMArrayOutput) Index(i pulumi.IntInput) AzureRMOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AzureRM {
		return vs[0].([]*AzureRM)[vs[1].(int)]
	}).(AzureRMOutput)
}

type AzureRMMapOutput struct{ *pulumi.OutputState }

func (AzureRMMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AzureRM)(nil)).Elem()
}

func (o AzureRMMapOutput) ToAzureRMMapOutput() AzureRMMapOutput {
	return o
}

func (o AzureRMMapOutput) ToAzureRMMapOutputWithContext(ctx context.Context) AzureRMMapOutput {
	return o
}

func (o AzureRMMapOutput) MapIndex(k pulumi.StringInput) AzureRMOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AzureRM {
		return vs[0].(map[string]*AzureRM)[vs[1].(string)]
	}).(AzureRMOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AzureRMInput)(nil)).Elem(), &AzureRM{})
	pulumi.RegisterInputType(reflect.TypeOf((*AzureRMArrayInput)(nil)).Elem(), AzureRMArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AzureRMMapInput)(nil)).Elem(), AzureRMMap{})
	pulumi.RegisterOutputType(AzureRMOutput{})
	pulumi.RegisterOutputType(AzureRMArrayOutput{})
	pulumi.RegisterOutputType(AzureRMMapOutput{})
}
