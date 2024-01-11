// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

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
// ### Service Principal Manual AzureRM Service Endpoint (Subscription Scoped)
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
// ### Service Principal Manual AzureRM Service Endpoint (ManagementGroup Scoped)
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
// ### Service Principal Automatic AzureRM Service Endpoint
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
// ### Workload Identity Federation Manual AzureRM Service Endpoint (Subscription Scoped)
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-azuredevops/sdk/v2/go/azuredevops"
//	"github.com/pulumi/pulumi-azurerm/sdk/v1/go/azurerm"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			serviceConnectionName := "example-federated-sc"
//			exampleProject, err := azuredevops.NewProject(ctx, "exampleProject", &azuredevops.ProjectArgs{
//				Visibility:       pulumi.String("private"),
//				VersionControl:   pulumi.String("Git"),
//				WorkItemTemplate: pulumi.String("Agile"),
//				Description:      pulumi.String("Managed by Terraform"),
//			})
//			if err != nil {
//				return err
//			}
//			identity, err := index.NewAzurerm_resource_group(ctx, "identity", &index.Azurerm_resource_groupArgs{
//				Name:     "identity",
//				Location: "UK South",
//			})
//			if err != nil {
//				return err
//			}
//			exampleazurerm_user_assigned_identity, err := index.NewAzurerm_user_assigned_identity(ctx, "exampleazurerm_user_assigned_identity", &index.Azurerm_user_assigned_identityArgs{
//				Location:          identity.Location,
//				Name:              "example-identity",
//				ResourceGroupName: "azurerm_resource_group.identity.name",
//			})
//			if err != nil {
//				return err
//			}
//			exampleServiceEndpointAzureRM, err := azuredevops.NewServiceEndpointAzureRM(ctx, "exampleServiceEndpointAzureRM", &azuredevops.ServiceEndpointAzureRMArgs{
//				ProjectId:                           exampleProject.ID(),
//				ServiceEndpointName:                 pulumi.String(serviceConnectionName),
//				Description:                         pulumi.String("Managed by Terraform"),
//				ServiceEndpointAuthenticationScheme: pulumi.String("WorkloadIdentityFederation"),
//				Credentials: &azuredevops.ServiceEndpointAzureRMCredentialsArgs{
//					Serviceprincipalid: exampleazurerm_user_assigned_identity.ClientId,
//				},
//				AzurermSpnTenantid:      pulumi.String("00000000-0000-0000-0000-000000000000"),
//				AzurermSubscriptionId:   pulumi.String("00000000-0000-0000-0000-000000000000"),
//				AzurermSubscriptionName: pulumi.String("Example Subscription Name"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = index.NewAzurerm_federated_identity_credential(ctx, "exampleazurerm_federated_identity_credential", &index.Azurerm_federated_identity_credentialArgs{
//				Name:              "example-federated-credential",
//				ResourceGroupName: identity.Name,
//				ParentId:          exampleazurerm_user_assigned_identity.Id,
//				Audience: []string{
//					"api://AzureADTokenExchange",
//				},
//				Issuer:  exampleServiceEndpointAzureRM.WorkloadIdentityFederationIssuer,
//				Subject: exampleServiceEndpointAzureRM.WorkloadIdentityFederationSubject,
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Workload Identity Federation Automatic AzureRM Service Endpoint
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
// ### Managed Identity AzureRM Service Endpoint
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
// ## Relevant Links
//
// - [Azure DevOps Service REST API 7.0 - Service End points](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-7.0)
//
// ## Import
//
// Azure DevOps Service Endpoint Azure Resource Manage can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**
//
// ```sh
//
//	$ pulumi import azuredevops:index/serviceEndpointAzureRM:ServiceEndpointAzureRM example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
//
// ```
type ServiceEndpointAzureRM struct {
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
	Credentials ServiceEndpointAzureRMCredentialsPtrOutput `pulumi:"credentials"`
	// Service connection description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The Cloud Environment to use. Defaults to `AzureCloud`. Possible values are `AzureCloud`, `AzureChinaCloud`. Changing this forces a new resource to be created.
	//
	// > **NOTE:** One of either `Subscription` scoped i.e. `azurermSubscriptionId`, `azurermSubscriptionName` or `ManagementGroup` scoped i.e. `azurermManagementGroupId`, `azurermManagementGroupName` values must be specified.
	Environment pulumi.StringPtrOutput `pulumi:"environment"`
	// A `features` block.
	Features ServiceEndpointAzureRMFeaturesPtrOutput `pulumi:"features"`
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

// NewServiceEndpointAzureRM registers a new resource with the given unique name, arguments, and options.
func NewServiceEndpointAzureRM(ctx *pulumi.Context,
	name string, args *ServiceEndpointAzureRMArgs, opts ...pulumi.ResourceOption) (*ServiceEndpointAzureRM, error) {
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
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azuredevops:ServiceEndpoint/azureRM:AzureRM"),
		},
	})
	opts = append(opts, aliases)
	opts = internal.PkgResourceDefaultOpts(opts)
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
	Credentials *ServiceEndpointAzureRMCredentials `pulumi:"credentials"`
	// Service connection description.
	Description *string `pulumi:"description"`
	// The Cloud Environment to use. Defaults to `AzureCloud`. Possible values are `AzureCloud`, `AzureChinaCloud`. Changing this forces a new resource to be created.
	//
	// > **NOTE:** One of either `Subscription` scoped i.e. `azurermSubscriptionId`, `azurermSubscriptionName` or `ManagementGroup` scoped i.e. `azurermManagementGroupId`, `azurermManagementGroupName` values must be specified.
	Environment *string `pulumi:"environment"`
	// A `features` block.
	Features *ServiceEndpointAzureRMFeatures `pulumi:"features"`
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

type ServiceEndpointAzureRMState struct {
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
	Credentials ServiceEndpointAzureRMCredentialsPtrInput
	// Service connection description.
	Description pulumi.StringPtrInput
	// The Cloud Environment to use. Defaults to `AzureCloud`. Possible values are `AzureCloud`, `AzureChinaCloud`. Changing this forces a new resource to be created.
	//
	// > **NOTE:** One of either `Subscription` scoped i.e. `azurermSubscriptionId`, `azurermSubscriptionName` or `ManagementGroup` scoped i.e. `azurermManagementGroupId`, `azurermManagementGroupName` values must be specified.
	Environment pulumi.StringPtrInput
	// A `features` block.
	Features ServiceEndpointAzureRMFeaturesPtrInput
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

func (ServiceEndpointAzureRMState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceEndpointAzureRMState)(nil)).Elem()
}

type serviceEndpointAzureRMArgs struct {
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
	Credentials *ServiceEndpointAzureRMCredentials `pulumi:"credentials"`
	// Service connection description.
	Description *string `pulumi:"description"`
	// The Cloud Environment to use. Defaults to `AzureCloud`. Possible values are `AzureCloud`, `AzureChinaCloud`. Changing this forces a new resource to be created.
	//
	// > **NOTE:** One of either `Subscription` scoped i.e. `azurermSubscriptionId`, `azurermSubscriptionName` or `ManagementGroup` scoped i.e. `azurermManagementGroupId`, `azurermManagementGroupName` values must be specified.
	Environment *string `pulumi:"environment"`
	// A `features` block.
	Features *ServiceEndpointAzureRMFeatures `pulumi:"features"`
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

// The set of arguments for constructing a ServiceEndpointAzureRM resource.
type ServiceEndpointAzureRMArgs struct {
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
	Credentials ServiceEndpointAzureRMCredentialsPtrInput
	// Service connection description.
	Description pulumi.StringPtrInput
	// The Cloud Environment to use. Defaults to `AzureCloud`. Possible values are `AzureCloud`, `AzureChinaCloud`. Changing this forces a new resource to be created.
	//
	// > **NOTE:** One of either `Subscription` scoped i.e. `azurermSubscriptionId`, `azurermSubscriptionName` or `ManagementGroup` scoped i.e. `azurermManagementGroupId`, `azurermManagementGroupName` values must be specified.
	Environment pulumi.StringPtrInput
	// A `features` block.
	Features ServiceEndpointAzureRMFeaturesPtrInput
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

func (ServiceEndpointAzureRMArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceEndpointAzureRMArgs)(nil)).Elem()
}

type ServiceEndpointAzureRMInput interface {
	pulumi.Input

	ToServiceEndpointAzureRMOutput() ServiceEndpointAzureRMOutput
	ToServiceEndpointAzureRMOutputWithContext(ctx context.Context) ServiceEndpointAzureRMOutput
}

func (*ServiceEndpointAzureRM) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceEndpointAzureRM)(nil)).Elem()
}

func (i *ServiceEndpointAzureRM) ToServiceEndpointAzureRMOutput() ServiceEndpointAzureRMOutput {
	return i.ToServiceEndpointAzureRMOutputWithContext(context.Background())
}

func (i *ServiceEndpointAzureRM) ToServiceEndpointAzureRMOutputWithContext(ctx context.Context) ServiceEndpointAzureRMOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointAzureRMOutput)
}

// ServiceEndpointAzureRMArrayInput is an input type that accepts ServiceEndpointAzureRMArray and ServiceEndpointAzureRMArrayOutput values.
// You can construct a concrete instance of `ServiceEndpointAzureRMArrayInput` via:
//
//	ServiceEndpointAzureRMArray{ ServiceEndpointAzureRMArgs{...} }
type ServiceEndpointAzureRMArrayInput interface {
	pulumi.Input

	ToServiceEndpointAzureRMArrayOutput() ServiceEndpointAzureRMArrayOutput
	ToServiceEndpointAzureRMArrayOutputWithContext(context.Context) ServiceEndpointAzureRMArrayOutput
}

type ServiceEndpointAzureRMArray []ServiceEndpointAzureRMInput

func (ServiceEndpointAzureRMArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceEndpointAzureRM)(nil)).Elem()
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
//	ServiceEndpointAzureRMMap{ "key": ServiceEndpointAzureRMArgs{...} }
type ServiceEndpointAzureRMMapInput interface {
	pulumi.Input

	ToServiceEndpointAzureRMMapOutput() ServiceEndpointAzureRMMapOutput
	ToServiceEndpointAzureRMMapOutputWithContext(context.Context) ServiceEndpointAzureRMMapOutput
}

type ServiceEndpointAzureRMMap map[string]ServiceEndpointAzureRMInput

func (ServiceEndpointAzureRMMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceEndpointAzureRM)(nil)).Elem()
}

func (i ServiceEndpointAzureRMMap) ToServiceEndpointAzureRMMapOutput() ServiceEndpointAzureRMMapOutput {
	return i.ToServiceEndpointAzureRMMapOutputWithContext(context.Background())
}

func (i ServiceEndpointAzureRMMap) ToServiceEndpointAzureRMMapOutputWithContext(ctx context.Context) ServiceEndpointAzureRMMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointAzureRMMapOutput)
}

type ServiceEndpointAzureRMOutput struct{ *pulumi.OutputState }

func (ServiceEndpointAzureRMOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceEndpointAzureRM)(nil)).Elem()
}

func (o ServiceEndpointAzureRMOutput) ToServiceEndpointAzureRMOutput() ServiceEndpointAzureRMOutput {
	return o
}

func (o ServiceEndpointAzureRMOutput) ToServiceEndpointAzureRMOutputWithContext(ctx context.Context) ServiceEndpointAzureRMOutput {
	return o
}

func (o ServiceEndpointAzureRMOutput) Authorization() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ServiceEndpointAzureRM) pulumi.StringMapOutput { return v.Authorization }).(pulumi.StringMapOutput)
}

// The Management group ID of the Azure targets.
func (o ServiceEndpointAzureRMOutput) AzurermManagementGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceEndpointAzureRM) pulumi.StringPtrOutput { return v.AzurermManagementGroupId }).(pulumi.StringPtrOutput)
}

// The Management group Name of the targets.
func (o ServiceEndpointAzureRMOutput) AzurermManagementGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceEndpointAzureRM) pulumi.StringPtrOutput { return v.AzurermManagementGroupName }).(pulumi.StringPtrOutput)
}

// The Tenant ID if the service principal.
func (o ServiceEndpointAzureRMOutput) AzurermSpnTenantid() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEndpointAzureRM) pulumi.StringOutput { return v.AzurermSpnTenantid }).(pulumi.StringOutput)
}

// The Subscription ID of the Azure targets.
func (o ServiceEndpointAzureRMOutput) AzurermSubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceEndpointAzureRM) pulumi.StringPtrOutput { return v.AzurermSubscriptionId }).(pulumi.StringPtrOutput)
}

// The Subscription Name of the targets.
func (o ServiceEndpointAzureRMOutput) AzurermSubscriptionName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceEndpointAzureRM) pulumi.StringPtrOutput { return v.AzurermSubscriptionName }).(pulumi.StringPtrOutput)
}

// A `credentials` block.
func (o ServiceEndpointAzureRMOutput) Credentials() ServiceEndpointAzureRMCredentialsPtrOutput {
	return o.ApplyT(func(v *ServiceEndpointAzureRM) ServiceEndpointAzureRMCredentialsPtrOutput { return v.Credentials }).(ServiceEndpointAzureRMCredentialsPtrOutput)
}

// Service connection description.
func (o ServiceEndpointAzureRMOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceEndpointAzureRM) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The Cloud Environment to use. Defaults to `AzureCloud`. Possible values are `AzureCloud`, `AzureChinaCloud`. Changing this forces a new resource to be created.
//
// > **NOTE:** One of either `Subscription` scoped i.e. `azurermSubscriptionId`, `azurermSubscriptionName` or `ManagementGroup` scoped i.e. `azurermManagementGroupId`, `azurermManagementGroupName` values must be specified.
func (o ServiceEndpointAzureRMOutput) Environment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceEndpointAzureRM) pulumi.StringPtrOutput { return v.Environment }).(pulumi.StringPtrOutput)
}

// A `features` block.
func (o ServiceEndpointAzureRMOutput) Features() ServiceEndpointAzureRMFeaturesPtrOutput {
	return o.ApplyT(func(v *ServiceEndpointAzureRM) ServiceEndpointAzureRMFeaturesPtrOutput { return v.Features }).(ServiceEndpointAzureRMFeaturesPtrOutput)
}

// The ID of the project.
func (o ServiceEndpointAzureRMOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEndpointAzureRM) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The resource group used for scope of automatic service endpoint.
func (o ServiceEndpointAzureRMOutput) ResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceEndpointAzureRM) pulumi.StringPtrOutput { return v.ResourceGroup }).(pulumi.StringPtrOutput)
}

// Specifies the type of azurerm endpoint, either `WorkloadIdentityFederation`, `ManagedServiceIdentity` or `ServicePrincipal`. Defaults to `ServicePrincipal` for backwards compatibility.
//
// > **NOTE:** The `WorkloadIdentityFederation` authentication scheme is currently in private preview. Your organisation must be part of the preview and the feature toggle must be turned on to use it. More details can be found [here](https://aka.ms/azdo-rm-workload-identity).
func (o ServiceEndpointAzureRMOutput) ServiceEndpointAuthenticationScheme() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceEndpointAzureRM) pulumi.StringPtrOutput { return v.ServiceEndpointAuthenticationScheme }).(pulumi.StringPtrOutput)
}

// The Service Endpoint Name.
func (o ServiceEndpointAzureRMOutput) ServiceEndpointName() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEndpointAzureRM) pulumi.StringOutput { return v.ServiceEndpointName }).(pulumi.StringOutput)
}

// The Application(Client) ID of the Service Principal.
func (o ServiceEndpointAzureRMOutput) ServicePrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEndpointAzureRM) pulumi.StringOutput { return v.ServicePrincipalId }).(pulumi.StringOutput)
}

// The issuer if `serviceEndpointAuthenticationScheme` is set to `WorkloadIdentityFederation`. This looks like `https://vstoken.dev.azure.com/00000000-0000-0000-0000-000000000000`, where the GUID is the Organization ID of your Azure DevOps Organisation.
func (o ServiceEndpointAzureRMOutput) WorkloadIdentityFederationIssuer() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEndpointAzureRM) pulumi.StringOutput { return v.WorkloadIdentityFederationIssuer }).(pulumi.StringOutput)
}

// The subject if `serviceEndpointAuthenticationScheme` is set to `WorkloadIdentityFederation`. This looks like `sc://<organisation>/<project>/<service-connection-name>`.
func (o ServiceEndpointAzureRMOutput) WorkloadIdentityFederationSubject() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEndpointAzureRM) pulumi.StringOutput { return v.WorkloadIdentityFederationSubject }).(pulumi.StringOutput)
}

type ServiceEndpointAzureRMArrayOutput struct{ *pulumi.OutputState }

func (ServiceEndpointAzureRMArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceEndpointAzureRM)(nil)).Elem()
}

func (o ServiceEndpointAzureRMArrayOutput) ToServiceEndpointAzureRMArrayOutput() ServiceEndpointAzureRMArrayOutput {
	return o
}

func (o ServiceEndpointAzureRMArrayOutput) ToServiceEndpointAzureRMArrayOutputWithContext(ctx context.Context) ServiceEndpointAzureRMArrayOutput {
	return o
}

func (o ServiceEndpointAzureRMArrayOutput) Index(i pulumi.IntInput) ServiceEndpointAzureRMOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServiceEndpointAzureRM {
		return vs[0].([]*ServiceEndpointAzureRM)[vs[1].(int)]
	}).(ServiceEndpointAzureRMOutput)
}

type ServiceEndpointAzureRMMapOutput struct{ *pulumi.OutputState }

func (ServiceEndpointAzureRMMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceEndpointAzureRM)(nil)).Elem()
}

func (o ServiceEndpointAzureRMMapOutput) ToServiceEndpointAzureRMMapOutput() ServiceEndpointAzureRMMapOutput {
	return o
}

func (o ServiceEndpointAzureRMMapOutput) ToServiceEndpointAzureRMMapOutputWithContext(ctx context.Context) ServiceEndpointAzureRMMapOutput {
	return o
}

func (o ServiceEndpointAzureRMMapOutput) MapIndex(k pulumi.StringInput) ServiceEndpointAzureRMOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServiceEndpointAzureRM {
		return vs[0].(map[string]*ServiceEndpointAzureRM)[vs[1].(string)]
	}).(ServiceEndpointAzureRMOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceEndpointAzureRMInput)(nil)).Elem(), &ServiceEndpointAzureRM{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceEndpointAzureRMArrayInput)(nil)).Elem(), ServiceEndpointAzureRMArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceEndpointAzureRMMapInput)(nil)).Elem(), ServiceEndpointAzureRMMap{})
	pulumi.RegisterOutputType(ServiceEndpointAzureRMOutput{})
	pulumi.RegisterOutputType(ServiceEndpointAzureRMArrayOutput{})
	pulumi.RegisterOutputType(ServiceEndpointAzureRMMapOutput{})
}
