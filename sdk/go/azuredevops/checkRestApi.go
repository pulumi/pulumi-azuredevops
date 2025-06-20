// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azuredevops/sdk/v3/go/azuredevops/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Rest API check on a resource within Azure DevOps.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-azuredevops/sdk/v3/go/azuredevops"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			example, err := azuredevops.NewProject(ctx, "example", &azuredevops.ProjectArgs{
//				Name: pulumi.String("Example Project"),
//			})
//			if err != nil {
//				return err
//			}
//			exampleServiceEndpointGeneric, err := azuredevops.NewServiceEndpointGeneric(ctx, "example", &azuredevops.ServiceEndpointGenericArgs{
//				ProjectId:           example.ID(),
//				ServerUrl:           pulumi.String("https://some-server.example.com"),
//				ServiceEndpointName: pulumi.String("Example Generic"),
//				Username:            pulumi.String("username"),
//				Password:            pulumi.String("password"),
//				Description:         pulumi.String("Managed by Pulumi"),
//			})
//			if err != nil {
//				return err
//			}
//			exampleAzure, err := azuredevops.NewServiceEndpointGeneric(ctx, "example_azure", &azuredevops.ServiceEndpointGenericArgs{
//				ProjectId:           example.ID(),
//				ServerUrl:           pulumi.String("https://dev.azure.com/"),
//				ServiceEndpointName: pulumi.String("Example Generic Azure"),
//				Username:            pulumi.String("username"),
//				Password:            pulumi.String("dummy"),
//			})
//			if err != nil {
//				return err
//			}
//			exampleVariableGroup, err := azuredevops.NewVariableGroup(ctx, "example", &azuredevops.VariableGroupArgs{
//				ProjectId:   example.ID(),
//				Name:        pulumi.String("Example Variable Group"),
//				AllowAccess: pulumi.Bool(true),
//				Variables: azuredevops.VariableGroupVariableArray{
//					&azuredevops.VariableGroupVariableArgs{
//						Name:  pulumi.String("FOO"),
//						Value: pulumi.String("BAR"),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = azuredevops.NewCheckRestApi(ctx, "example", &azuredevops.CheckRestApiArgs{
//				ProjectId:                    example.ID(),
//				TargetResourceId:             exampleServiceEndpointGeneric.ID(),
//				TargetResourceType:           pulumi.String("endpoint"),
//				DisplayName:                  pulumi.String("Example REST API Check"),
//				ConnectedServiceNameSelector: pulumi.String("connectedServiceName"),
//				ConnectedServiceName:         exampleAzure.ServiceEndpointName,
//				Method:                       pulumi.String("POST"),
//				Headers:                      pulumi.String("{\"contentType\":\"application/json\"}"),
//				Body:                         pulumi.String("{\"params\":\"value\"}"),
//				CompletionEvent:              pulumi.String("ApiResponse"),
//				SuccessCriteria:              pulumi.String("eq(root['status'], '200')"),
//				UrlSuffix:                    pulumi.String("user/1"),
//				RetryInterval:                pulumi.Int(4000),
//				VariableGroupName:            exampleVariableGroup.Name,
//				Timeout:                      pulumi.Int(40000),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Relevant Links
//
// - [Define approvals and checks](https://learn.microsoft.com/en-us/azure/devops/pipelines/process/approvals?view=azure-devops&tabs=check-pass)
//
// ## Import
//
// Importing this resource is not supported.
type CheckRestApi struct {
	pulumi.CustomResourceState

	// The Rest API request body.
	Body pulumi.StringPtrOutput `pulumi:"body"`
	// The completion event of the Rest API call. Possible values: `Callback`, `ApiResponse`. Defaults to `Callback`.
	CompletionEvent pulumi.StringPtrOutput `pulumi:"completionEvent"`
	// The name of the Service Connection.
	ConnectedServiceName pulumi.StringOutput `pulumi:"connectedServiceName"`
	// The type of the Service Connection used to invoke the REST API. Possible values: `connectedServiceName`(**Generic** type service connection) and `connectedServiceNameARM`(**Azure Resource Manager** type service connection).
	ConnectedServiceNameSelector pulumi.StringOutput `pulumi:"connectedServiceNameSelector"`
	// The Name of the Rest API check.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The headers of the request in JSON format.
	Headers pulumi.StringPtrOutput `pulumi:"headers"`
	// The HTTP method of the request. Possible values: `OPTIONS`, `GET`, `HEAD`, `POST`, `PUT`, `DELETE`, `TRACE`, `PATCH`
	Method pulumi.StringOutput `pulumi:"method"`
	// The ID of the project. Changing this forces a new resource to be created.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The time between evaluations (minutes).
	//
	// ~>**NOTE** 1. The retry times should less them 10 based on the timeout. For example: `timeout` is `4000` then `retryInterval` should be `0` or no less then `400`.
	// <br>2. `retryInterval` is not required when `completion_event=Callback`.
	RetryInterval pulumi.IntPtrOutput `pulumi:"retryInterval"`
	// The Criteria which defines when to pass the task. No criteria means response content does not influence the result.
	//
	// ~>**NOTE** `successCriteria` is used when `completion_event=ApiResponse`
	SuccessCriteria pulumi.StringPtrOutput `pulumi:"successCriteria"`
	// The ID of the resource being protected by the check. Changing this forces a new resource to be created
	TargetResourceId pulumi.StringOutput `pulumi:"targetResourceId"`
	// The type of resource being protected by the check. Possible values: `endpoint`, `environment`, `queue`, `repository`, `securefile`, `variablegroup`. Changing this forces a new resource to be created.
	TargetResourceType pulumi.StringOutput `pulumi:"targetResourceType"`
	// The timeout in minutes for the Rest API check. Defaults to `1440`.
	Timeout pulumi.IntPtrOutput `pulumi:"timeout"`
	// The URL suffix and parameters.
	UrlSuffix pulumi.StringPtrOutput `pulumi:"urlSuffix"`
	// The name of the Variable Group.
	VariableGroupName pulumi.StringPtrOutput `pulumi:"variableGroupName"`
	// The version of the Rest API check.
	Version pulumi.IntOutput `pulumi:"version"`
}

// NewCheckRestApi registers a new resource with the given unique name, arguments, and options.
func NewCheckRestApi(ctx *pulumi.Context,
	name string, args *CheckRestApiArgs, opts ...pulumi.ResourceOption) (*CheckRestApi, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConnectedServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ConnectedServiceName'")
	}
	if args.ConnectedServiceNameSelector == nil {
		return nil, errors.New("invalid value for required argument 'ConnectedServiceNameSelector'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.Method == nil {
		return nil, errors.New("invalid value for required argument 'Method'")
	}
	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.TargetResourceId == nil {
		return nil, errors.New("invalid value for required argument 'TargetResourceId'")
	}
	if args.TargetResourceType == nil {
		return nil, errors.New("invalid value for required argument 'TargetResourceType'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource CheckRestApi
	err := ctx.RegisterResource("azuredevops:index/checkRestApi:CheckRestApi", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCheckRestApi gets an existing CheckRestApi resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCheckRestApi(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CheckRestApiState, opts ...pulumi.ResourceOption) (*CheckRestApi, error) {
	var resource CheckRestApi
	err := ctx.ReadResource("azuredevops:index/checkRestApi:CheckRestApi", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CheckRestApi resources.
type checkRestApiState struct {
	// The Rest API request body.
	Body *string `pulumi:"body"`
	// The completion event of the Rest API call. Possible values: `Callback`, `ApiResponse`. Defaults to `Callback`.
	CompletionEvent *string `pulumi:"completionEvent"`
	// The name of the Service Connection.
	ConnectedServiceName *string `pulumi:"connectedServiceName"`
	// The type of the Service Connection used to invoke the REST API. Possible values: `connectedServiceName`(**Generic** type service connection) and `connectedServiceNameARM`(**Azure Resource Manager** type service connection).
	ConnectedServiceNameSelector *string `pulumi:"connectedServiceNameSelector"`
	// The Name of the Rest API check.
	DisplayName *string `pulumi:"displayName"`
	// The headers of the request in JSON format.
	Headers *string `pulumi:"headers"`
	// The HTTP method of the request. Possible values: `OPTIONS`, `GET`, `HEAD`, `POST`, `PUT`, `DELETE`, `TRACE`, `PATCH`
	Method *string `pulumi:"method"`
	// The ID of the project. Changing this forces a new resource to be created.
	ProjectId *string `pulumi:"projectId"`
	// The time between evaluations (minutes).
	//
	// ~>**NOTE** 1. The retry times should less them 10 based on the timeout. For example: `timeout` is `4000` then `retryInterval` should be `0` or no less then `400`.
	// <br>2. `retryInterval` is not required when `completion_event=Callback`.
	RetryInterval *int `pulumi:"retryInterval"`
	// The Criteria which defines when to pass the task. No criteria means response content does not influence the result.
	//
	// ~>**NOTE** `successCriteria` is used when `completion_event=ApiResponse`
	SuccessCriteria *string `pulumi:"successCriteria"`
	// The ID of the resource being protected by the check. Changing this forces a new resource to be created
	TargetResourceId *string `pulumi:"targetResourceId"`
	// The type of resource being protected by the check. Possible values: `endpoint`, `environment`, `queue`, `repository`, `securefile`, `variablegroup`. Changing this forces a new resource to be created.
	TargetResourceType *string `pulumi:"targetResourceType"`
	// The timeout in minutes for the Rest API check. Defaults to `1440`.
	Timeout *int `pulumi:"timeout"`
	// The URL suffix and parameters.
	UrlSuffix *string `pulumi:"urlSuffix"`
	// The name of the Variable Group.
	VariableGroupName *string `pulumi:"variableGroupName"`
	// The version of the Rest API check.
	Version *int `pulumi:"version"`
}

type CheckRestApiState struct {
	// The Rest API request body.
	Body pulumi.StringPtrInput
	// The completion event of the Rest API call. Possible values: `Callback`, `ApiResponse`. Defaults to `Callback`.
	CompletionEvent pulumi.StringPtrInput
	// The name of the Service Connection.
	ConnectedServiceName pulumi.StringPtrInput
	// The type of the Service Connection used to invoke the REST API. Possible values: `connectedServiceName`(**Generic** type service connection) and `connectedServiceNameARM`(**Azure Resource Manager** type service connection).
	ConnectedServiceNameSelector pulumi.StringPtrInput
	// The Name of the Rest API check.
	DisplayName pulumi.StringPtrInput
	// The headers of the request in JSON format.
	Headers pulumi.StringPtrInput
	// The HTTP method of the request. Possible values: `OPTIONS`, `GET`, `HEAD`, `POST`, `PUT`, `DELETE`, `TRACE`, `PATCH`
	Method pulumi.StringPtrInput
	// The ID of the project. Changing this forces a new resource to be created.
	ProjectId pulumi.StringPtrInput
	// The time between evaluations (minutes).
	//
	// ~>**NOTE** 1. The retry times should less them 10 based on the timeout. For example: `timeout` is `4000` then `retryInterval` should be `0` or no less then `400`.
	// <br>2. `retryInterval` is not required when `completion_event=Callback`.
	RetryInterval pulumi.IntPtrInput
	// The Criteria which defines when to pass the task. No criteria means response content does not influence the result.
	//
	// ~>**NOTE** `successCriteria` is used when `completion_event=ApiResponse`
	SuccessCriteria pulumi.StringPtrInput
	// The ID of the resource being protected by the check. Changing this forces a new resource to be created
	TargetResourceId pulumi.StringPtrInput
	// The type of resource being protected by the check. Possible values: `endpoint`, `environment`, `queue`, `repository`, `securefile`, `variablegroup`. Changing this forces a new resource to be created.
	TargetResourceType pulumi.StringPtrInput
	// The timeout in minutes for the Rest API check. Defaults to `1440`.
	Timeout pulumi.IntPtrInput
	// The URL suffix and parameters.
	UrlSuffix pulumi.StringPtrInput
	// The name of the Variable Group.
	VariableGroupName pulumi.StringPtrInput
	// The version of the Rest API check.
	Version pulumi.IntPtrInput
}

func (CheckRestApiState) ElementType() reflect.Type {
	return reflect.TypeOf((*checkRestApiState)(nil)).Elem()
}

type checkRestApiArgs struct {
	// The Rest API request body.
	Body *string `pulumi:"body"`
	// The completion event of the Rest API call. Possible values: `Callback`, `ApiResponse`. Defaults to `Callback`.
	CompletionEvent *string `pulumi:"completionEvent"`
	// The name of the Service Connection.
	ConnectedServiceName string `pulumi:"connectedServiceName"`
	// The type of the Service Connection used to invoke the REST API. Possible values: `connectedServiceName`(**Generic** type service connection) and `connectedServiceNameARM`(**Azure Resource Manager** type service connection).
	ConnectedServiceNameSelector string `pulumi:"connectedServiceNameSelector"`
	// The Name of the Rest API check.
	DisplayName string `pulumi:"displayName"`
	// The headers of the request in JSON format.
	Headers *string `pulumi:"headers"`
	// The HTTP method of the request. Possible values: `OPTIONS`, `GET`, `HEAD`, `POST`, `PUT`, `DELETE`, `TRACE`, `PATCH`
	Method string `pulumi:"method"`
	// The ID of the project. Changing this forces a new resource to be created.
	ProjectId string `pulumi:"projectId"`
	// The time between evaluations (minutes).
	//
	// ~>**NOTE** 1. The retry times should less them 10 based on the timeout. For example: `timeout` is `4000` then `retryInterval` should be `0` or no less then `400`.
	// <br>2. `retryInterval` is not required when `completion_event=Callback`.
	RetryInterval *int `pulumi:"retryInterval"`
	// The Criteria which defines when to pass the task. No criteria means response content does not influence the result.
	//
	// ~>**NOTE** `successCriteria` is used when `completion_event=ApiResponse`
	SuccessCriteria *string `pulumi:"successCriteria"`
	// The ID of the resource being protected by the check. Changing this forces a new resource to be created
	TargetResourceId string `pulumi:"targetResourceId"`
	// The type of resource being protected by the check. Possible values: `endpoint`, `environment`, `queue`, `repository`, `securefile`, `variablegroup`. Changing this forces a new resource to be created.
	TargetResourceType string `pulumi:"targetResourceType"`
	// The timeout in minutes for the Rest API check. Defaults to `1440`.
	Timeout *int `pulumi:"timeout"`
	// The URL suffix and parameters.
	UrlSuffix *string `pulumi:"urlSuffix"`
	// The name of the Variable Group.
	VariableGroupName *string `pulumi:"variableGroupName"`
}

// The set of arguments for constructing a CheckRestApi resource.
type CheckRestApiArgs struct {
	// The Rest API request body.
	Body pulumi.StringPtrInput
	// The completion event of the Rest API call. Possible values: `Callback`, `ApiResponse`. Defaults to `Callback`.
	CompletionEvent pulumi.StringPtrInput
	// The name of the Service Connection.
	ConnectedServiceName pulumi.StringInput
	// The type of the Service Connection used to invoke the REST API. Possible values: `connectedServiceName`(**Generic** type service connection) and `connectedServiceNameARM`(**Azure Resource Manager** type service connection).
	ConnectedServiceNameSelector pulumi.StringInput
	// The Name of the Rest API check.
	DisplayName pulumi.StringInput
	// The headers of the request in JSON format.
	Headers pulumi.StringPtrInput
	// The HTTP method of the request. Possible values: `OPTIONS`, `GET`, `HEAD`, `POST`, `PUT`, `DELETE`, `TRACE`, `PATCH`
	Method pulumi.StringInput
	// The ID of the project. Changing this forces a new resource to be created.
	ProjectId pulumi.StringInput
	// The time between evaluations (minutes).
	//
	// ~>**NOTE** 1. The retry times should less them 10 based on the timeout. For example: `timeout` is `4000` then `retryInterval` should be `0` or no less then `400`.
	// <br>2. `retryInterval` is not required when `completion_event=Callback`.
	RetryInterval pulumi.IntPtrInput
	// The Criteria which defines when to pass the task. No criteria means response content does not influence the result.
	//
	// ~>**NOTE** `successCriteria` is used when `completion_event=ApiResponse`
	SuccessCriteria pulumi.StringPtrInput
	// The ID of the resource being protected by the check. Changing this forces a new resource to be created
	TargetResourceId pulumi.StringInput
	// The type of resource being protected by the check. Possible values: `endpoint`, `environment`, `queue`, `repository`, `securefile`, `variablegroup`. Changing this forces a new resource to be created.
	TargetResourceType pulumi.StringInput
	// The timeout in minutes for the Rest API check. Defaults to `1440`.
	Timeout pulumi.IntPtrInput
	// The URL suffix and parameters.
	UrlSuffix pulumi.StringPtrInput
	// The name of the Variable Group.
	VariableGroupName pulumi.StringPtrInput
}

func (CheckRestApiArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*checkRestApiArgs)(nil)).Elem()
}

type CheckRestApiInput interface {
	pulumi.Input

	ToCheckRestApiOutput() CheckRestApiOutput
	ToCheckRestApiOutputWithContext(ctx context.Context) CheckRestApiOutput
}

func (*CheckRestApi) ElementType() reflect.Type {
	return reflect.TypeOf((**CheckRestApi)(nil)).Elem()
}

func (i *CheckRestApi) ToCheckRestApiOutput() CheckRestApiOutput {
	return i.ToCheckRestApiOutputWithContext(context.Background())
}

func (i *CheckRestApi) ToCheckRestApiOutputWithContext(ctx context.Context) CheckRestApiOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CheckRestApiOutput)
}

// CheckRestApiArrayInput is an input type that accepts CheckRestApiArray and CheckRestApiArrayOutput values.
// You can construct a concrete instance of `CheckRestApiArrayInput` via:
//
//	CheckRestApiArray{ CheckRestApiArgs{...} }
type CheckRestApiArrayInput interface {
	pulumi.Input

	ToCheckRestApiArrayOutput() CheckRestApiArrayOutput
	ToCheckRestApiArrayOutputWithContext(context.Context) CheckRestApiArrayOutput
}

type CheckRestApiArray []CheckRestApiInput

func (CheckRestApiArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CheckRestApi)(nil)).Elem()
}

func (i CheckRestApiArray) ToCheckRestApiArrayOutput() CheckRestApiArrayOutput {
	return i.ToCheckRestApiArrayOutputWithContext(context.Background())
}

func (i CheckRestApiArray) ToCheckRestApiArrayOutputWithContext(ctx context.Context) CheckRestApiArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CheckRestApiArrayOutput)
}

// CheckRestApiMapInput is an input type that accepts CheckRestApiMap and CheckRestApiMapOutput values.
// You can construct a concrete instance of `CheckRestApiMapInput` via:
//
//	CheckRestApiMap{ "key": CheckRestApiArgs{...} }
type CheckRestApiMapInput interface {
	pulumi.Input

	ToCheckRestApiMapOutput() CheckRestApiMapOutput
	ToCheckRestApiMapOutputWithContext(context.Context) CheckRestApiMapOutput
}

type CheckRestApiMap map[string]CheckRestApiInput

func (CheckRestApiMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CheckRestApi)(nil)).Elem()
}

func (i CheckRestApiMap) ToCheckRestApiMapOutput() CheckRestApiMapOutput {
	return i.ToCheckRestApiMapOutputWithContext(context.Background())
}

func (i CheckRestApiMap) ToCheckRestApiMapOutputWithContext(ctx context.Context) CheckRestApiMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CheckRestApiMapOutput)
}

type CheckRestApiOutput struct{ *pulumi.OutputState }

func (CheckRestApiOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CheckRestApi)(nil)).Elem()
}

func (o CheckRestApiOutput) ToCheckRestApiOutput() CheckRestApiOutput {
	return o
}

func (o CheckRestApiOutput) ToCheckRestApiOutputWithContext(ctx context.Context) CheckRestApiOutput {
	return o
}

// The Rest API request body.
func (o CheckRestApiOutput) Body() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CheckRestApi) pulumi.StringPtrOutput { return v.Body }).(pulumi.StringPtrOutput)
}

// The completion event of the Rest API call. Possible values: `Callback`, `ApiResponse`. Defaults to `Callback`.
func (o CheckRestApiOutput) CompletionEvent() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CheckRestApi) pulumi.StringPtrOutput { return v.CompletionEvent }).(pulumi.StringPtrOutput)
}

// The name of the Service Connection.
func (o CheckRestApiOutput) ConnectedServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *CheckRestApi) pulumi.StringOutput { return v.ConnectedServiceName }).(pulumi.StringOutput)
}

// The type of the Service Connection used to invoke the REST API. Possible values: `connectedServiceName`(**Generic** type service connection) and `connectedServiceNameARM`(**Azure Resource Manager** type service connection).
func (o CheckRestApiOutput) ConnectedServiceNameSelector() pulumi.StringOutput {
	return o.ApplyT(func(v *CheckRestApi) pulumi.StringOutput { return v.ConnectedServiceNameSelector }).(pulumi.StringOutput)
}

// The Name of the Rest API check.
func (o CheckRestApiOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *CheckRestApi) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// The headers of the request in JSON format.
func (o CheckRestApiOutput) Headers() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CheckRestApi) pulumi.StringPtrOutput { return v.Headers }).(pulumi.StringPtrOutput)
}

// The HTTP method of the request. Possible values: `OPTIONS`, `GET`, `HEAD`, `POST`, `PUT`, `DELETE`, `TRACE`, `PATCH`
func (o CheckRestApiOutput) Method() pulumi.StringOutput {
	return o.ApplyT(func(v *CheckRestApi) pulumi.StringOutput { return v.Method }).(pulumi.StringOutput)
}

// The ID of the project. Changing this forces a new resource to be created.
func (o CheckRestApiOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *CheckRestApi) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The time between evaluations (minutes).
//
// ~>**NOTE** 1. The retry times should less them 10 based on the timeout. For example: `timeout` is `4000` then `retryInterval` should be `0` or no less then `400`.
// <br>2. `retryInterval` is not required when `completion_event=Callback`.
func (o CheckRestApiOutput) RetryInterval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CheckRestApi) pulumi.IntPtrOutput { return v.RetryInterval }).(pulumi.IntPtrOutput)
}

// The Criteria which defines when to pass the task. No criteria means response content does not influence the result.
//
// ~>**NOTE** `successCriteria` is used when `completion_event=ApiResponse`
func (o CheckRestApiOutput) SuccessCriteria() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CheckRestApi) pulumi.StringPtrOutput { return v.SuccessCriteria }).(pulumi.StringPtrOutput)
}

// The ID of the resource being protected by the check. Changing this forces a new resource to be created
func (o CheckRestApiOutput) TargetResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *CheckRestApi) pulumi.StringOutput { return v.TargetResourceId }).(pulumi.StringOutput)
}

// The type of resource being protected by the check. Possible values: `endpoint`, `environment`, `queue`, `repository`, `securefile`, `variablegroup`. Changing this forces a new resource to be created.
func (o CheckRestApiOutput) TargetResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *CheckRestApi) pulumi.StringOutput { return v.TargetResourceType }).(pulumi.StringOutput)
}

// The timeout in minutes for the Rest API check. Defaults to `1440`.
func (o CheckRestApiOutput) Timeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CheckRestApi) pulumi.IntPtrOutput { return v.Timeout }).(pulumi.IntPtrOutput)
}

// The URL suffix and parameters.
func (o CheckRestApiOutput) UrlSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CheckRestApi) pulumi.StringPtrOutput { return v.UrlSuffix }).(pulumi.StringPtrOutput)
}

// The name of the Variable Group.
func (o CheckRestApiOutput) VariableGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CheckRestApi) pulumi.StringPtrOutput { return v.VariableGroupName }).(pulumi.StringPtrOutput)
}

// The version of the Rest API check.
func (o CheckRestApiOutput) Version() pulumi.IntOutput {
	return o.ApplyT(func(v *CheckRestApi) pulumi.IntOutput { return v.Version }).(pulumi.IntOutput)
}

type CheckRestApiArrayOutput struct{ *pulumi.OutputState }

func (CheckRestApiArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CheckRestApi)(nil)).Elem()
}

func (o CheckRestApiArrayOutput) ToCheckRestApiArrayOutput() CheckRestApiArrayOutput {
	return o
}

func (o CheckRestApiArrayOutput) ToCheckRestApiArrayOutputWithContext(ctx context.Context) CheckRestApiArrayOutput {
	return o
}

func (o CheckRestApiArrayOutput) Index(i pulumi.IntInput) CheckRestApiOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CheckRestApi {
		return vs[0].([]*CheckRestApi)[vs[1].(int)]
	}).(CheckRestApiOutput)
}

type CheckRestApiMapOutput struct{ *pulumi.OutputState }

func (CheckRestApiMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CheckRestApi)(nil)).Elem()
}

func (o CheckRestApiMapOutput) ToCheckRestApiMapOutput() CheckRestApiMapOutput {
	return o
}

func (o CheckRestApiMapOutput) ToCheckRestApiMapOutputWithContext(ctx context.Context) CheckRestApiMapOutput {
	return o
}

func (o CheckRestApiMapOutput) MapIndex(k pulumi.StringInput) CheckRestApiOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CheckRestApi {
		return vs[0].(map[string]*CheckRestApi)[vs[1].(string)]
	}).(CheckRestApiOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CheckRestApiInput)(nil)).Elem(), &CheckRestApi{})
	pulumi.RegisterInputType(reflect.TypeOf((*CheckRestApiArrayInput)(nil)).Elem(), CheckRestApiArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CheckRestApiMapInput)(nil)).Elem(), CheckRestApiMap{})
	pulumi.RegisterOutputType(CheckRestApiOutput{})
	pulumi.RegisterOutputType(CheckRestApiArrayOutput{})
	pulumi.RegisterOutputType(CheckRestApiMapOutput{})
}
