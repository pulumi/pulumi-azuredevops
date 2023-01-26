// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a AWS service endpoint within Azure DevOps. Using this service endpoint requires you to first install [AWS Toolkit for Azure DevOps](https://marketplace.visualstudio.com/items?itemName=AmazonWebServices.aws-vsts-tools).
//
// ## Example Usage
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
//			_, err = azuredevops.NewServiceEndpointAws(ctx, "exampleServiceEndpointAws", &azuredevops.ServiceEndpointAwsArgs{
//				ProjectId:           exampleProject.ID(),
//				ServiceEndpointName: pulumi.String("Example AWS"),
//				AccessKeyId:         pulumi.String("00000000-0000-0000-0000-000000000000"),
//				SecretAccessKey:     pulumi.String("accesskey"),
//				Description:         pulumi.String("Managed by AzureDevOps"),
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
// * [aws-toolkit-azure-devops](https://github.com/aws/aws-toolkit-azure-devops)
// * [Azure DevOps Service REST API 6.0 - Agent Pools](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-6.0)
//
// ## Import
//
// Azure DevOps Service Endpoint AWS can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**
//
// ```sh
//
//	$ pulumi import azuredevops:index/serviceEndpointAws:ServiceEndpointAws azuredevops_serviceendpoint_aws.example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
//
// ```
type ServiceEndpointAws struct {
	pulumi.CustomResourceState

	// The AWS access key ID for signing programmatic requests.
	AccessKeyId   pulumi.StringOutput    `pulumi:"accessKeyId"`
	Authorization pulumi.StringMapOutput `pulumi:"authorization"`
	Description   pulumi.StringPtrOutput `pulumi:"description"`
	// A unique identifier that is used by third parties when assuming roles in their customers' accounts, aka cross-account role access.
	ExternalId pulumi.StringPtrOutput `pulumi:"externalId"`
	// The ID of the project.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// Optional identifier for the assumed role session.
	RoleSessionName pulumi.StringPtrOutput `pulumi:"roleSessionName"`
	// The Amazon Resource Name (ARN) of the role to assume.
	RoleToAssume pulumi.StringPtrOutput `pulumi:"roleToAssume"`
	// The AWS secret access key for signing programmatic requests.
	SecretAccessKey pulumi.StringOutput `pulumi:"secretAccessKey"`
	// A bcrypted hash of the attribute 'secret_access_key'
	SecretAccessKeyHash pulumi.StringOutput `pulumi:"secretAccessKeyHash"`
	// The Service Endpoint name.
	ServiceEndpointName pulumi.StringOutput `pulumi:"serviceEndpointName"`
	// The AWS session token for signing programmatic requests.
	SessionToken pulumi.StringPtrOutput `pulumi:"sessionToken"`
	// A bcrypted hash of the attribute 'session_token'
	SessionTokenHash pulumi.StringOutput `pulumi:"sessionTokenHash"`
}

// NewServiceEndpointAws registers a new resource with the given unique name, arguments, and options.
func NewServiceEndpointAws(ctx *pulumi.Context,
	name string, args *ServiceEndpointAwsArgs, opts ...pulumi.ResourceOption) (*ServiceEndpointAws, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccessKeyId == nil {
		return nil, errors.New("invalid value for required argument 'AccessKeyId'")
	}
	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.SecretAccessKey == nil {
		return nil, errors.New("invalid value for required argument 'SecretAccessKey'")
	}
	if args.ServiceEndpointName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceEndpointName'")
	}
	var resource ServiceEndpointAws
	err := ctx.RegisterResource("azuredevops:index/serviceEndpointAws:ServiceEndpointAws", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceEndpointAws gets an existing ServiceEndpointAws resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceEndpointAws(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceEndpointAwsState, opts ...pulumi.ResourceOption) (*ServiceEndpointAws, error) {
	var resource ServiceEndpointAws
	err := ctx.ReadResource("azuredevops:index/serviceEndpointAws:ServiceEndpointAws", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceEndpointAws resources.
type serviceEndpointAwsState struct {
	// The AWS access key ID for signing programmatic requests.
	AccessKeyId   *string           `pulumi:"accessKeyId"`
	Authorization map[string]string `pulumi:"authorization"`
	Description   *string           `pulumi:"description"`
	// A unique identifier that is used by third parties when assuming roles in their customers' accounts, aka cross-account role access.
	ExternalId *string `pulumi:"externalId"`
	// The ID of the project.
	ProjectId *string `pulumi:"projectId"`
	// Optional identifier for the assumed role session.
	RoleSessionName *string `pulumi:"roleSessionName"`
	// The Amazon Resource Name (ARN) of the role to assume.
	RoleToAssume *string `pulumi:"roleToAssume"`
	// The AWS secret access key for signing programmatic requests.
	SecretAccessKey *string `pulumi:"secretAccessKey"`
	// A bcrypted hash of the attribute 'secret_access_key'
	SecretAccessKeyHash *string `pulumi:"secretAccessKeyHash"`
	// The Service Endpoint name.
	ServiceEndpointName *string `pulumi:"serviceEndpointName"`
	// The AWS session token for signing programmatic requests.
	SessionToken *string `pulumi:"sessionToken"`
	// A bcrypted hash of the attribute 'session_token'
	SessionTokenHash *string `pulumi:"sessionTokenHash"`
}

type ServiceEndpointAwsState struct {
	// The AWS access key ID for signing programmatic requests.
	AccessKeyId   pulumi.StringPtrInput
	Authorization pulumi.StringMapInput
	Description   pulumi.StringPtrInput
	// A unique identifier that is used by third parties when assuming roles in their customers' accounts, aka cross-account role access.
	ExternalId pulumi.StringPtrInput
	// The ID of the project.
	ProjectId pulumi.StringPtrInput
	// Optional identifier for the assumed role session.
	RoleSessionName pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) of the role to assume.
	RoleToAssume pulumi.StringPtrInput
	// The AWS secret access key for signing programmatic requests.
	SecretAccessKey pulumi.StringPtrInput
	// A bcrypted hash of the attribute 'secret_access_key'
	SecretAccessKeyHash pulumi.StringPtrInput
	// The Service Endpoint name.
	ServiceEndpointName pulumi.StringPtrInput
	// The AWS session token for signing programmatic requests.
	SessionToken pulumi.StringPtrInput
	// A bcrypted hash of the attribute 'session_token'
	SessionTokenHash pulumi.StringPtrInput
}

func (ServiceEndpointAwsState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceEndpointAwsState)(nil)).Elem()
}

type serviceEndpointAwsArgs struct {
	// The AWS access key ID for signing programmatic requests.
	AccessKeyId   string            `pulumi:"accessKeyId"`
	Authorization map[string]string `pulumi:"authorization"`
	Description   *string           `pulumi:"description"`
	// A unique identifier that is used by third parties when assuming roles in their customers' accounts, aka cross-account role access.
	ExternalId *string `pulumi:"externalId"`
	// The ID of the project.
	ProjectId string `pulumi:"projectId"`
	// Optional identifier for the assumed role session.
	RoleSessionName *string `pulumi:"roleSessionName"`
	// The Amazon Resource Name (ARN) of the role to assume.
	RoleToAssume *string `pulumi:"roleToAssume"`
	// The AWS secret access key for signing programmatic requests.
	SecretAccessKey string `pulumi:"secretAccessKey"`
	// The Service Endpoint name.
	ServiceEndpointName string `pulumi:"serviceEndpointName"`
	// The AWS session token for signing programmatic requests.
	SessionToken *string `pulumi:"sessionToken"`
}

// The set of arguments for constructing a ServiceEndpointAws resource.
type ServiceEndpointAwsArgs struct {
	// The AWS access key ID for signing programmatic requests.
	AccessKeyId   pulumi.StringInput
	Authorization pulumi.StringMapInput
	Description   pulumi.StringPtrInput
	// A unique identifier that is used by third parties when assuming roles in their customers' accounts, aka cross-account role access.
	ExternalId pulumi.StringPtrInput
	// The ID of the project.
	ProjectId pulumi.StringInput
	// Optional identifier for the assumed role session.
	RoleSessionName pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) of the role to assume.
	RoleToAssume pulumi.StringPtrInput
	// The AWS secret access key for signing programmatic requests.
	SecretAccessKey pulumi.StringInput
	// The Service Endpoint name.
	ServiceEndpointName pulumi.StringInput
	// The AWS session token for signing programmatic requests.
	SessionToken pulumi.StringPtrInput
}

func (ServiceEndpointAwsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceEndpointAwsArgs)(nil)).Elem()
}

type ServiceEndpointAwsInput interface {
	pulumi.Input

	ToServiceEndpointAwsOutput() ServiceEndpointAwsOutput
	ToServiceEndpointAwsOutputWithContext(ctx context.Context) ServiceEndpointAwsOutput
}

func (*ServiceEndpointAws) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceEndpointAws)(nil)).Elem()
}

func (i *ServiceEndpointAws) ToServiceEndpointAwsOutput() ServiceEndpointAwsOutput {
	return i.ToServiceEndpointAwsOutputWithContext(context.Background())
}

func (i *ServiceEndpointAws) ToServiceEndpointAwsOutputWithContext(ctx context.Context) ServiceEndpointAwsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointAwsOutput)
}

// ServiceEndpointAwsArrayInput is an input type that accepts ServiceEndpointAwsArray and ServiceEndpointAwsArrayOutput values.
// You can construct a concrete instance of `ServiceEndpointAwsArrayInput` via:
//
//	ServiceEndpointAwsArray{ ServiceEndpointAwsArgs{...} }
type ServiceEndpointAwsArrayInput interface {
	pulumi.Input

	ToServiceEndpointAwsArrayOutput() ServiceEndpointAwsArrayOutput
	ToServiceEndpointAwsArrayOutputWithContext(context.Context) ServiceEndpointAwsArrayOutput
}

type ServiceEndpointAwsArray []ServiceEndpointAwsInput

func (ServiceEndpointAwsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceEndpointAws)(nil)).Elem()
}

func (i ServiceEndpointAwsArray) ToServiceEndpointAwsArrayOutput() ServiceEndpointAwsArrayOutput {
	return i.ToServiceEndpointAwsArrayOutputWithContext(context.Background())
}

func (i ServiceEndpointAwsArray) ToServiceEndpointAwsArrayOutputWithContext(ctx context.Context) ServiceEndpointAwsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointAwsArrayOutput)
}

// ServiceEndpointAwsMapInput is an input type that accepts ServiceEndpointAwsMap and ServiceEndpointAwsMapOutput values.
// You can construct a concrete instance of `ServiceEndpointAwsMapInput` via:
//
//	ServiceEndpointAwsMap{ "key": ServiceEndpointAwsArgs{...} }
type ServiceEndpointAwsMapInput interface {
	pulumi.Input

	ToServiceEndpointAwsMapOutput() ServiceEndpointAwsMapOutput
	ToServiceEndpointAwsMapOutputWithContext(context.Context) ServiceEndpointAwsMapOutput
}

type ServiceEndpointAwsMap map[string]ServiceEndpointAwsInput

func (ServiceEndpointAwsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceEndpointAws)(nil)).Elem()
}

func (i ServiceEndpointAwsMap) ToServiceEndpointAwsMapOutput() ServiceEndpointAwsMapOutput {
	return i.ToServiceEndpointAwsMapOutputWithContext(context.Background())
}

func (i ServiceEndpointAwsMap) ToServiceEndpointAwsMapOutputWithContext(ctx context.Context) ServiceEndpointAwsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointAwsMapOutput)
}

type ServiceEndpointAwsOutput struct{ *pulumi.OutputState }

func (ServiceEndpointAwsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceEndpointAws)(nil)).Elem()
}

func (o ServiceEndpointAwsOutput) ToServiceEndpointAwsOutput() ServiceEndpointAwsOutput {
	return o
}

func (o ServiceEndpointAwsOutput) ToServiceEndpointAwsOutputWithContext(ctx context.Context) ServiceEndpointAwsOutput {
	return o
}

// The AWS access key ID for signing programmatic requests.
func (o ServiceEndpointAwsOutput) AccessKeyId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEndpointAws) pulumi.StringOutput { return v.AccessKeyId }).(pulumi.StringOutput)
}

func (o ServiceEndpointAwsOutput) Authorization() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ServiceEndpointAws) pulumi.StringMapOutput { return v.Authorization }).(pulumi.StringMapOutput)
}

func (o ServiceEndpointAwsOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceEndpointAws) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// A unique identifier that is used by third parties when assuming roles in their customers' accounts, aka cross-account role access.
func (o ServiceEndpointAwsOutput) ExternalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceEndpointAws) pulumi.StringPtrOutput { return v.ExternalId }).(pulumi.StringPtrOutput)
}

// The ID of the project.
func (o ServiceEndpointAwsOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEndpointAws) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// Optional identifier for the assumed role session.
func (o ServiceEndpointAwsOutput) RoleSessionName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceEndpointAws) pulumi.StringPtrOutput { return v.RoleSessionName }).(pulumi.StringPtrOutput)
}

// The Amazon Resource Name (ARN) of the role to assume.
func (o ServiceEndpointAwsOutput) RoleToAssume() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceEndpointAws) pulumi.StringPtrOutput { return v.RoleToAssume }).(pulumi.StringPtrOutput)
}

// The AWS secret access key for signing programmatic requests.
func (o ServiceEndpointAwsOutput) SecretAccessKey() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEndpointAws) pulumi.StringOutput { return v.SecretAccessKey }).(pulumi.StringOutput)
}

// A bcrypted hash of the attribute 'secret_access_key'
func (o ServiceEndpointAwsOutput) SecretAccessKeyHash() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEndpointAws) pulumi.StringOutput { return v.SecretAccessKeyHash }).(pulumi.StringOutput)
}

// The Service Endpoint name.
func (o ServiceEndpointAwsOutput) ServiceEndpointName() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEndpointAws) pulumi.StringOutput { return v.ServiceEndpointName }).(pulumi.StringOutput)
}

// The AWS session token for signing programmatic requests.
func (o ServiceEndpointAwsOutput) SessionToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceEndpointAws) pulumi.StringPtrOutput { return v.SessionToken }).(pulumi.StringPtrOutput)
}

// A bcrypted hash of the attribute 'session_token'
func (o ServiceEndpointAwsOutput) SessionTokenHash() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEndpointAws) pulumi.StringOutput { return v.SessionTokenHash }).(pulumi.StringOutput)
}

type ServiceEndpointAwsArrayOutput struct{ *pulumi.OutputState }

func (ServiceEndpointAwsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceEndpointAws)(nil)).Elem()
}

func (o ServiceEndpointAwsArrayOutput) ToServiceEndpointAwsArrayOutput() ServiceEndpointAwsArrayOutput {
	return o
}

func (o ServiceEndpointAwsArrayOutput) ToServiceEndpointAwsArrayOutputWithContext(ctx context.Context) ServiceEndpointAwsArrayOutput {
	return o
}

func (o ServiceEndpointAwsArrayOutput) Index(i pulumi.IntInput) ServiceEndpointAwsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServiceEndpointAws {
		return vs[0].([]*ServiceEndpointAws)[vs[1].(int)]
	}).(ServiceEndpointAwsOutput)
}

type ServiceEndpointAwsMapOutput struct{ *pulumi.OutputState }

func (ServiceEndpointAwsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceEndpointAws)(nil)).Elem()
}

func (o ServiceEndpointAwsMapOutput) ToServiceEndpointAwsMapOutput() ServiceEndpointAwsMapOutput {
	return o
}

func (o ServiceEndpointAwsMapOutput) ToServiceEndpointAwsMapOutputWithContext(ctx context.Context) ServiceEndpointAwsMapOutput {
	return o
}

func (o ServiceEndpointAwsMapOutput) MapIndex(k pulumi.StringInput) ServiceEndpointAwsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServiceEndpointAws {
		return vs[0].(map[string]*ServiceEndpointAws)[vs[1].(string)]
	}).(ServiceEndpointAwsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceEndpointAwsInput)(nil)).Elem(), &ServiceEndpointAws{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceEndpointAwsArrayInput)(nil)).Elem(), ServiceEndpointAwsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceEndpointAwsMapInput)(nil)).Elem(), ServiceEndpointAwsMap{})
	pulumi.RegisterOutputType(ServiceEndpointAwsOutput{})
	pulumi.RegisterOutputType(ServiceEndpointAwsArrayOutput{})
	pulumi.RegisterOutputType(ServiceEndpointAwsMapOutput{})
}
