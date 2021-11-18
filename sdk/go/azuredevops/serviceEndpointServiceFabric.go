// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Service Fabric service endpoint within Azure DevOps.
//
// ## Example Usage
// ### Client Certificate Authentication
//
// ```go
// package main
//
// import (
// 	"encoding/base64"
// 	"io/ioutil"
//
// 	"github.com/pulumi/pulumi-azuredevops/sdk/v2/go/azuredevops"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func filebase64OrPanic(path string) pulumi.StringPtrInput {
// 	if fileData, err := ioutil.ReadFile(path); err == nil {
// 		return pulumi.String(base64.StdEncoding.EncodeToString(fileData[:]))
// 	} else {
// 		panic(err.Error())
// 	}
// }
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
// 		_, err = azuredevops.NewServiceEndpointServiceFabric(ctx, "test", &azuredevops.ServiceEndpointServiceFabricArgs{
// 			ProjectId:           project.ID(),
// 			ServiceEndpointName: pulumi.String("Sample Service Fabric"),
// 			Description:         pulumi.String("Managed by Terraform"),
// 			ClusterEndpoint:     pulumi.String("tcp://test"),
// 			Certificate: &ServiceEndpointServiceFabricCertificateArgs{
// 				ServerCertificateLookup:     pulumi.String("Thumbprint"),
// 				ServerCertificateThumbprint: pulumi.String("0000000000000000000000000000000000000000"),
// 				ClientCertificate:           filebase64OrPanic("certificate.pfx"),
// 				ClientCertificatePassword:   pulumi.String("password"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Azure Active Directory Authentication
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
// 		_, err = azuredevops.NewServiceEndpointServiceFabric(ctx, "test", &azuredevops.ServiceEndpointServiceFabricArgs{
// 			ProjectId:           project.ID(),
// 			ServiceEndpointName: pulumi.String("Sample Service Fabric"),
// 			Description:         pulumi.String("Managed by Terraform"),
// 			ClusterEndpoint:     pulumi.String("tcp://test"),
// 			AzureActiveDirectory: &ServiceEndpointServiceFabricAzureActiveDirectoryArgs{
// 				ServerCertificateLookup:     pulumi.String("Thumbprint"),
// 				ServerCertificateThumbprint: pulumi.String("0000000000000000000000000000000000000000"),
// 				Username:                    pulumi.String("username"),
// 				Password:                    pulumi.String("password"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Windows Authentication
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
// 		_, err = azuredevops.NewServiceEndpointServiceFabric(ctx, "test", &azuredevops.ServiceEndpointServiceFabricArgs{
// 			ProjectId:           project.ID(),
// 			ServiceEndpointName: pulumi.String("Sample Service Fabric"),
// 			Description:         pulumi.String("Managed by Terraform"),
// 			ClusterEndpoint:     pulumi.String("tcp://test"),
// 			None: &ServiceEndpointServiceFabricNoneArgs{
// 				Unsecured:  pulumi.Bool(false),
// 				ClusterSpn: pulumi.String("HTTP/www.contoso.com"),
// 			},
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
// - [Azure DevOps Service REST API 5.1 - Endpoints](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-5.1)
//
// ## Import
//
// Azure DevOps Service Endpoint Service Fabric can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**
//
// ```sh
//  $ pulumi import azuredevops:index/serviceEndpointServiceFabric:ServiceEndpointServiceFabric serviceendpoint 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
// ```
type ServiceEndpointServiceFabric struct {
	pulumi.CustomResourceState

	Authorization        pulumi.StringMapOutput                                    `pulumi:"authorization"`
	AzureActiveDirectory ServiceEndpointServiceFabricAzureActiveDirectoryPtrOutput `pulumi:"azureActiveDirectory"`
	Certificate          ServiceEndpointServiceFabricCertificatePtrOutput          `pulumi:"certificate"`
	// Client connection endpoint for the cluster. Prefix the value with 'tcp://';. This value overrides the publish profile.
	ClusterEndpoint pulumi.StringOutput                       `pulumi:"clusterEndpoint"`
	Description     pulumi.StringPtrOutput                    `pulumi:"description"`
	None            ServiceEndpointServiceFabricNonePtrOutput `pulumi:"none"`
	// The project ID or project name.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The Service Endpoint name.
	ServiceEndpointName pulumi.StringOutput `pulumi:"serviceEndpointName"`
}

// NewServiceEndpointServiceFabric registers a new resource with the given unique name, arguments, and options.
func NewServiceEndpointServiceFabric(ctx *pulumi.Context,
	name string, args *ServiceEndpointServiceFabricArgs, opts ...pulumi.ResourceOption) (*ServiceEndpointServiceFabric, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterEndpoint == nil {
		return nil, errors.New("invalid value for required argument 'ClusterEndpoint'")
	}
	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.ServiceEndpointName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceEndpointName'")
	}
	var resource ServiceEndpointServiceFabric
	err := ctx.RegisterResource("azuredevops:index/serviceEndpointServiceFabric:ServiceEndpointServiceFabric", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceEndpointServiceFabric gets an existing ServiceEndpointServiceFabric resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceEndpointServiceFabric(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceEndpointServiceFabricState, opts ...pulumi.ResourceOption) (*ServiceEndpointServiceFabric, error) {
	var resource ServiceEndpointServiceFabric
	err := ctx.ReadResource("azuredevops:index/serviceEndpointServiceFabric:ServiceEndpointServiceFabric", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceEndpointServiceFabric resources.
type serviceEndpointServiceFabricState struct {
	Authorization        map[string]string                                 `pulumi:"authorization"`
	AzureActiveDirectory *ServiceEndpointServiceFabricAzureActiveDirectory `pulumi:"azureActiveDirectory"`
	Certificate          *ServiceEndpointServiceFabricCertificate          `pulumi:"certificate"`
	// Client connection endpoint for the cluster. Prefix the value with 'tcp://';. This value overrides the publish profile.
	ClusterEndpoint *string                           `pulumi:"clusterEndpoint"`
	Description     *string                           `pulumi:"description"`
	None            *ServiceEndpointServiceFabricNone `pulumi:"none"`
	// The project ID or project name.
	ProjectId *string `pulumi:"projectId"`
	// The Service Endpoint name.
	ServiceEndpointName *string `pulumi:"serviceEndpointName"`
}

type ServiceEndpointServiceFabricState struct {
	Authorization        pulumi.StringMapInput
	AzureActiveDirectory ServiceEndpointServiceFabricAzureActiveDirectoryPtrInput
	Certificate          ServiceEndpointServiceFabricCertificatePtrInput
	// Client connection endpoint for the cluster. Prefix the value with 'tcp://';. This value overrides the publish profile.
	ClusterEndpoint pulumi.StringPtrInput
	Description     pulumi.StringPtrInput
	None            ServiceEndpointServiceFabricNonePtrInput
	// The project ID or project name.
	ProjectId pulumi.StringPtrInput
	// The Service Endpoint name.
	ServiceEndpointName pulumi.StringPtrInput
}

func (ServiceEndpointServiceFabricState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceEndpointServiceFabricState)(nil)).Elem()
}

type serviceEndpointServiceFabricArgs struct {
	Authorization        map[string]string                                 `pulumi:"authorization"`
	AzureActiveDirectory *ServiceEndpointServiceFabricAzureActiveDirectory `pulumi:"azureActiveDirectory"`
	Certificate          *ServiceEndpointServiceFabricCertificate          `pulumi:"certificate"`
	// Client connection endpoint for the cluster. Prefix the value with 'tcp://';. This value overrides the publish profile.
	ClusterEndpoint string                            `pulumi:"clusterEndpoint"`
	Description     *string                           `pulumi:"description"`
	None            *ServiceEndpointServiceFabricNone `pulumi:"none"`
	// The project ID or project name.
	ProjectId string `pulumi:"projectId"`
	// The Service Endpoint name.
	ServiceEndpointName string `pulumi:"serviceEndpointName"`
}

// The set of arguments for constructing a ServiceEndpointServiceFabric resource.
type ServiceEndpointServiceFabricArgs struct {
	Authorization        pulumi.StringMapInput
	AzureActiveDirectory ServiceEndpointServiceFabricAzureActiveDirectoryPtrInput
	Certificate          ServiceEndpointServiceFabricCertificatePtrInput
	// Client connection endpoint for the cluster. Prefix the value with 'tcp://';. This value overrides the publish profile.
	ClusterEndpoint pulumi.StringInput
	Description     pulumi.StringPtrInput
	None            ServiceEndpointServiceFabricNonePtrInput
	// The project ID or project name.
	ProjectId pulumi.StringInput
	// The Service Endpoint name.
	ServiceEndpointName pulumi.StringInput
}

func (ServiceEndpointServiceFabricArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceEndpointServiceFabricArgs)(nil)).Elem()
}

type ServiceEndpointServiceFabricInput interface {
	pulumi.Input

	ToServiceEndpointServiceFabricOutput() ServiceEndpointServiceFabricOutput
	ToServiceEndpointServiceFabricOutputWithContext(ctx context.Context) ServiceEndpointServiceFabricOutput
}

func (*ServiceEndpointServiceFabric) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceEndpointServiceFabric)(nil))
}

func (i *ServiceEndpointServiceFabric) ToServiceEndpointServiceFabricOutput() ServiceEndpointServiceFabricOutput {
	return i.ToServiceEndpointServiceFabricOutputWithContext(context.Background())
}

func (i *ServiceEndpointServiceFabric) ToServiceEndpointServiceFabricOutputWithContext(ctx context.Context) ServiceEndpointServiceFabricOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointServiceFabricOutput)
}

func (i *ServiceEndpointServiceFabric) ToServiceEndpointServiceFabricPtrOutput() ServiceEndpointServiceFabricPtrOutput {
	return i.ToServiceEndpointServiceFabricPtrOutputWithContext(context.Background())
}

func (i *ServiceEndpointServiceFabric) ToServiceEndpointServiceFabricPtrOutputWithContext(ctx context.Context) ServiceEndpointServiceFabricPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointServiceFabricPtrOutput)
}

type ServiceEndpointServiceFabricPtrInput interface {
	pulumi.Input

	ToServiceEndpointServiceFabricPtrOutput() ServiceEndpointServiceFabricPtrOutput
	ToServiceEndpointServiceFabricPtrOutputWithContext(ctx context.Context) ServiceEndpointServiceFabricPtrOutput
}

type serviceEndpointServiceFabricPtrType ServiceEndpointServiceFabricArgs

func (*serviceEndpointServiceFabricPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceEndpointServiceFabric)(nil))
}

func (i *serviceEndpointServiceFabricPtrType) ToServiceEndpointServiceFabricPtrOutput() ServiceEndpointServiceFabricPtrOutput {
	return i.ToServiceEndpointServiceFabricPtrOutputWithContext(context.Background())
}

func (i *serviceEndpointServiceFabricPtrType) ToServiceEndpointServiceFabricPtrOutputWithContext(ctx context.Context) ServiceEndpointServiceFabricPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointServiceFabricPtrOutput)
}

// ServiceEndpointServiceFabricArrayInput is an input type that accepts ServiceEndpointServiceFabricArray and ServiceEndpointServiceFabricArrayOutput values.
// You can construct a concrete instance of `ServiceEndpointServiceFabricArrayInput` via:
//
//          ServiceEndpointServiceFabricArray{ ServiceEndpointServiceFabricArgs{...} }
type ServiceEndpointServiceFabricArrayInput interface {
	pulumi.Input

	ToServiceEndpointServiceFabricArrayOutput() ServiceEndpointServiceFabricArrayOutput
	ToServiceEndpointServiceFabricArrayOutputWithContext(context.Context) ServiceEndpointServiceFabricArrayOutput
}

type ServiceEndpointServiceFabricArray []ServiceEndpointServiceFabricInput

func (ServiceEndpointServiceFabricArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceEndpointServiceFabric)(nil)).Elem()
}

func (i ServiceEndpointServiceFabricArray) ToServiceEndpointServiceFabricArrayOutput() ServiceEndpointServiceFabricArrayOutput {
	return i.ToServiceEndpointServiceFabricArrayOutputWithContext(context.Background())
}

func (i ServiceEndpointServiceFabricArray) ToServiceEndpointServiceFabricArrayOutputWithContext(ctx context.Context) ServiceEndpointServiceFabricArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointServiceFabricArrayOutput)
}

// ServiceEndpointServiceFabricMapInput is an input type that accepts ServiceEndpointServiceFabricMap and ServiceEndpointServiceFabricMapOutput values.
// You can construct a concrete instance of `ServiceEndpointServiceFabricMapInput` via:
//
//          ServiceEndpointServiceFabricMap{ "key": ServiceEndpointServiceFabricArgs{...} }
type ServiceEndpointServiceFabricMapInput interface {
	pulumi.Input

	ToServiceEndpointServiceFabricMapOutput() ServiceEndpointServiceFabricMapOutput
	ToServiceEndpointServiceFabricMapOutputWithContext(context.Context) ServiceEndpointServiceFabricMapOutput
}

type ServiceEndpointServiceFabricMap map[string]ServiceEndpointServiceFabricInput

func (ServiceEndpointServiceFabricMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceEndpointServiceFabric)(nil)).Elem()
}

func (i ServiceEndpointServiceFabricMap) ToServiceEndpointServiceFabricMapOutput() ServiceEndpointServiceFabricMapOutput {
	return i.ToServiceEndpointServiceFabricMapOutputWithContext(context.Background())
}

func (i ServiceEndpointServiceFabricMap) ToServiceEndpointServiceFabricMapOutputWithContext(ctx context.Context) ServiceEndpointServiceFabricMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointServiceFabricMapOutput)
}

type ServiceEndpointServiceFabricOutput struct{ *pulumi.OutputState }

func (ServiceEndpointServiceFabricOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceEndpointServiceFabric)(nil))
}

func (o ServiceEndpointServiceFabricOutput) ToServiceEndpointServiceFabricOutput() ServiceEndpointServiceFabricOutput {
	return o
}

func (o ServiceEndpointServiceFabricOutput) ToServiceEndpointServiceFabricOutputWithContext(ctx context.Context) ServiceEndpointServiceFabricOutput {
	return o
}

func (o ServiceEndpointServiceFabricOutput) ToServiceEndpointServiceFabricPtrOutput() ServiceEndpointServiceFabricPtrOutput {
	return o.ToServiceEndpointServiceFabricPtrOutputWithContext(context.Background())
}

func (o ServiceEndpointServiceFabricOutput) ToServiceEndpointServiceFabricPtrOutputWithContext(ctx context.Context) ServiceEndpointServiceFabricPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServiceEndpointServiceFabric) *ServiceEndpointServiceFabric {
		return &v
	}).(ServiceEndpointServiceFabricPtrOutput)
}

type ServiceEndpointServiceFabricPtrOutput struct{ *pulumi.OutputState }

func (ServiceEndpointServiceFabricPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceEndpointServiceFabric)(nil))
}

func (o ServiceEndpointServiceFabricPtrOutput) ToServiceEndpointServiceFabricPtrOutput() ServiceEndpointServiceFabricPtrOutput {
	return o
}

func (o ServiceEndpointServiceFabricPtrOutput) ToServiceEndpointServiceFabricPtrOutputWithContext(ctx context.Context) ServiceEndpointServiceFabricPtrOutput {
	return o
}

func (o ServiceEndpointServiceFabricPtrOutput) Elem() ServiceEndpointServiceFabricOutput {
	return o.ApplyT(func(v *ServiceEndpointServiceFabric) ServiceEndpointServiceFabric {
		if v != nil {
			return *v
		}
		var ret ServiceEndpointServiceFabric
		return ret
	}).(ServiceEndpointServiceFabricOutput)
}

type ServiceEndpointServiceFabricArrayOutput struct{ *pulumi.OutputState }

func (ServiceEndpointServiceFabricArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceEndpointServiceFabric)(nil))
}

func (o ServiceEndpointServiceFabricArrayOutput) ToServiceEndpointServiceFabricArrayOutput() ServiceEndpointServiceFabricArrayOutput {
	return o
}

func (o ServiceEndpointServiceFabricArrayOutput) ToServiceEndpointServiceFabricArrayOutputWithContext(ctx context.Context) ServiceEndpointServiceFabricArrayOutput {
	return o
}

func (o ServiceEndpointServiceFabricArrayOutput) Index(i pulumi.IntInput) ServiceEndpointServiceFabricOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServiceEndpointServiceFabric {
		return vs[0].([]ServiceEndpointServiceFabric)[vs[1].(int)]
	}).(ServiceEndpointServiceFabricOutput)
}

type ServiceEndpointServiceFabricMapOutput struct{ *pulumi.OutputState }

func (ServiceEndpointServiceFabricMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ServiceEndpointServiceFabric)(nil))
}

func (o ServiceEndpointServiceFabricMapOutput) ToServiceEndpointServiceFabricMapOutput() ServiceEndpointServiceFabricMapOutput {
	return o
}

func (o ServiceEndpointServiceFabricMapOutput) ToServiceEndpointServiceFabricMapOutputWithContext(ctx context.Context) ServiceEndpointServiceFabricMapOutput {
	return o
}

func (o ServiceEndpointServiceFabricMapOutput) MapIndex(k pulumi.StringInput) ServiceEndpointServiceFabricOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ServiceEndpointServiceFabric {
		return vs[0].(map[string]ServiceEndpointServiceFabric)[vs[1].(string)]
	}).(ServiceEndpointServiceFabricOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceEndpointServiceFabricInput)(nil)).Elem(), &ServiceEndpointServiceFabric{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceEndpointServiceFabricPtrInput)(nil)).Elem(), &ServiceEndpointServiceFabric{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceEndpointServiceFabricArrayInput)(nil)).Elem(), ServiceEndpointServiceFabricArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceEndpointServiceFabricMapInput)(nil)).Elem(), ServiceEndpointServiceFabricMap{})
	pulumi.RegisterOutputType(ServiceEndpointServiceFabricOutput{})
	pulumi.RegisterOutputType(ServiceEndpointServiceFabricPtrOutput{})
	pulumi.RegisterOutputType(ServiceEndpointServiceFabricArrayOutput{})
	pulumi.RegisterOutputType(ServiceEndpointServiceFabricMapOutput{})
}
