// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azuredevops/sdk/v3/go/azuredevops/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a business hours check on a resource within Azure DevOps.
//
// ## Example Usage
//
// ### Protect a service connection
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
//				Username:            pulumi.String("username"),
//				Password:            pulumi.String("password"),
//				ServiceEndpointName: pulumi.String("Example Generic"),
//				Description:         pulumi.String("Managed by Pulumi"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = azuredevops.NewCheckBusinessHours(ctx, "example", &azuredevops.CheckBusinessHoursArgs{
//				ProjectId:          example.ID(),
//				DisplayName:        pulumi.String("Managed by Pulumi"),
//				TargetResourceId:   exampleServiceEndpointGeneric.ID(),
//				TargetResourceType: pulumi.String("endpoint"),
//				StartTime:          pulumi.String("07:00"),
//				EndTime:            pulumi.String("15:30"),
//				TimeZone:           pulumi.String("UTC"),
//				Monday:             pulumi.Bool(true),
//				Tuesday:            pulumi.Bool(true),
//				Timeout:            pulumi.Int(1440),
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
// ### Protect an environment
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
//			exampleEnvironment, err := azuredevops.NewEnvironment(ctx, "example", &azuredevops.EnvironmentArgs{
//				ProjectId: example.ID(),
//				Name:      pulumi.String("Example Environment"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = azuredevops.NewCheckBusinessHours(ctx, "example", &azuredevops.CheckBusinessHoursArgs{
//				ProjectId:          example.ID(),
//				DisplayName:        pulumi.String("Managed by Pulumi"),
//				TargetResourceId:   exampleEnvironment.ID(),
//				TargetResourceType: pulumi.String("environment"),
//				StartTime:          pulumi.String("07:00"),
//				EndTime:            pulumi.String("15:30"),
//				TimeZone:           pulumi.String("UTC"),
//				Monday:             pulumi.Bool(true),
//				Tuesday:            pulumi.Bool(true),
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
// ### Protect an agent queue
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
//			examplePool, err := azuredevops.NewPool(ctx, "example", &azuredevops.PoolArgs{
//				Name: pulumi.String("example-pool"),
//			})
//			if err != nil {
//				return err
//			}
//			exampleQueue, err := azuredevops.NewQueue(ctx, "example", &azuredevops.QueueArgs{
//				ProjectId:   example.ID(),
//				AgentPoolId: examplePool.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = azuredevops.NewCheckBusinessHours(ctx, "example", &azuredevops.CheckBusinessHoursArgs{
//				ProjectId:          example.ID(),
//				DisplayName:        pulumi.String("Managed by Pulumi"),
//				TargetResourceId:   exampleQueue.ID(),
//				TargetResourceType: pulumi.String("queue"),
//				StartTime:          pulumi.String("07:00"),
//				EndTime:            pulumi.String("15:30"),
//				TimeZone:           pulumi.String("UTC"),
//				Monday:             pulumi.Bool(true),
//				Tuesday:            pulumi.Bool(true),
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
// ### Protect a repository
//
// ```go
// package main
//
// import (
//
//	"fmt"
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
//			exampleGit, err := azuredevops.NewGit(ctx, "example", &azuredevops.GitArgs{
//				ProjectId: example.ID(),
//				Name:      pulumi.String("Example Empty Git Repository"),
//				Initialization: &azuredevops.GitInitializationArgs{
//					InitType: pulumi.String("Clean"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = azuredevops.NewCheckBusinessHours(ctx, "example", &azuredevops.CheckBusinessHoursArgs{
//				ProjectId:   example.ID(),
//				DisplayName: pulumi.String("Managed by Pulumi"),
//				TargetResourceId: pulumi.All(example.ID(), exampleGit.ID()).ApplyT(func(_args []interface{}) (string, error) {
//					exampleId := _args[0].(string)
//					exampleGitId := _args[1].(string)
//					return fmt.Sprintf("%v.%v", exampleId, exampleGitId), nil
//				}).(pulumi.StringOutput),
//				TargetResourceType: pulumi.String("repository"),
//				StartTime:          pulumi.String("07:00"),
//				EndTime:            pulumi.String("15:30"),
//				TimeZone:           pulumi.String("UTC"),
//				Monday:             pulumi.Bool(true),
//				Tuesday:            pulumi.Bool(true),
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
// ### Protect a variable group
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
//			exampleVariableGroup, err := azuredevops.NewVariableGroup(ctx, "example", &azuredevops.VariableGroupArgs{
//				ProjectId:   example.ID(),
//				Name:        pulumi.String("Example Variable Group"),
//				Description: pulumi.String("Example Variable Group Description"),
//				AllowAccess: pulumi.Bool(true),
//				Variables: azuredevops.VariableGroupVariableArray{
//					&azuredevops.VariableGroupVariableArgs{
//						Name:  pulumi.String("key1"),
//						Value: pulumi.String("val1"),
//					},
//					&azuredevops.VariableGroupVariableArgs{
//						Name:        pulumi.String("key2"),
//						SecretValue: pulumi.String("val2"),
//						IsSecret:    pulumi.Bool(true),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = azuredevops.NewCheckBusinessHours(ctx, "example", &azuredevops.CheckBusinessHoursArgs{
//				ProjectId:          example.ID(),
//				DisplayName:        pulumi.String("Managed by Pulumi"),
//				TargetResourceId:   exampleVariableGroup.ID(),
//				TargetResourceType: pulumi.String("variablegroup"),
//				StartTime:          pulumi.String("07:00"),
//				EndTime:            pulumi.String("15:30"),
//				TimeZone:           pulumi.String("UTC"),
//				Monday:             pulumi.Bool(true),
//				Tuesday:            pulumi.Bool(true),
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
// ## Supported Time Zones
//
// - AUS Central Standard Time
// - AUS Eastern Standard Time
// - Afghanistan Standard Time
// - Alaskan Standard Time
// - Aleutian Standard Time
// - Altai Standard Time
// - Arab Standard Time
// - Arabian Standard Time
// - Arabic Standard Time
// - Argentina Standard Time
// - Astrakhan Standard Time
// - Atlantic Standard Time
// - Aus Central W. Standard Time
// - Azerbaijan Standard Time
// - Azores Standard Time
// - Bahia Standard Time
// - Bangladesh Standard Time
// - Belarus Standard Time
// - Bougainville Standard Time
// - Canada Central Standard Time
// - Cape Verde Standard Time
// - Caucasus Standard Time
// - Cen. Australia Standard Time
// - Central America Standard Time
// - Central Asia Standard Time
// - Central Brazilian Standard Time
// - Central Europe Standard Time
// - Central European Standard Time
// - Central Pacific Standard Time
// - Central Standard Time (Mexico)
// - Central Standard Time
// - Chatham Islands Standard Time
// - China Standard Time
// - Cuba Standard Time
// - Dateline Standard Time
// - E. Africa Standard Time
// - E. Australia Standard Time
// - E. Europe Standard Time
// - E. South America Standard Time
// - Easter Island Standard Time
// - Eastern Standard Time (Mexico)
// - Eastern Standard Time
// - Egypt Standard Time
// - Ekaterinburg Standard Time
// - FLE Standard Time
// - Fiji Standard Time
// - GMT Standard Time
// - GTB Standard Time
// - Georgian Standard Time
// - Greenland Standard Time
// - Greenwich Standard Time
// - Haiti Standard Time
// - Hawaiian Standard Time
// - India Standard Time
// - Iran Standard Time
// - Israel Standard Time
// - Jordan Standard Time
// - Kaliningrad Standard Time
// - Kamchatka Standard Time
// - Korea Standard Time
// - Libya Standard Time
// - Line Islands Standard Time
// - Lord Howe Standard Time
// - Magadan Standard Time
// - Magallanes Standard Time
// - Marquesas Standard Time
// - Mauritius Standard Time
// - Mid-Atlantic Standard Time
// - Middle East Standard Time
// - Montevideo Standard Time
// - Morocco Standard Time
// - Mountain Standard Time (Mexico)
// - Mountain Standard Time
// - Myanmar Standard Time
// - N. Central Asia Standard Time
// - Namibia Standard Time
// - Nepal Standard Time
// - New Zealand Standard Time
// - Newfoundland Standard Time
// - Norfolk Standard Time
// - North Asia East Standard Time
// - North Asia Standard Time
// - North Korea Standard Time
// - Omsk Standard Time
// - Pacific SA Standard Time
// - Pacific Standard Time (Mexico)
// - Pacific Standard Time
// - Pakistan Standard Time
// - Paraguay Standard Time
// - Qyzylorda Standard Time
// - Romance Standard Time
// - Russia Time Zone 10
// - Russia Time Zone 11
// - Russia Time Zone 3
// - Russian Standard Time
// - SA Eastern Standard Time
// - SA Pacific Standard Time
// - SA Western Standard Time
// - SE Asia Standard Time
// - Saint Pierre Standard Time
// - Sakhalin Standard Time
// - Samoa Standard Time
// - Sao Tome Standard Time
// - Saratov Standard Time
// - Singapore Standard Time
// - South Africa Standard Time
// - South Sudan Standard Time
// - Sri Lanka Standard Time
// - Sudan Standard Time
// - Syria Standard Time
// - Taipei Standard Time
// - Tasmania Standard Time
// - Tocantins Standard Time
// - Tokyo Standard Time
// - Tomsk Standard Time
// - Tonga Standard Time
// - Transbaikal Standard Time
// - Turkey Standard Time
// - Turks And Caicos Standard Time
// - US Eastern Standard Time
// - US Mountain Standard Time
// - UTC
// - UTC+12
// - UTC+13
// - UTC-02
// - UTC-08
// - UTC-09
// - UTC-11
// - Ulaanbaatar Standard Time
// - Venezuela Standard Time
// - Vladivostok Standard Time
// - Volgograd Standard Time
// - W. Australia Standard Time
// - W. Central Africa Standard Time
// - W. Europe Standard Time
// - W. Mongolia Standard Time
// - West Asia Standard Time
// - West Bank Standard Time
// - West Pacific Standard Time
// - Yakutsk Standard Time
// - Yukon Standard Time
//
// ## Import
//
// Importing this resource is not supported.
type CheckBusinessHours struct {
	pulumi.CustomResourceState

	// The name of the business hours check displayed in the web UI.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// The end of the time period that this check will be allowed to pass, specified as 24-hour time with leading zeros.
	EndTime pulumi.StringOutput `pulumi:"endTime"`
	// This check will pass on Fridays. Defaults to `false`.
	Friday pulumi.BoolPtrOutput `pulumi:"friday"`
	// This check will pass on Mondays. Defaults to `false`.
	Monday pulumi.BoolPtrOutput `pulumi:"monday"`
	// The project ID.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// This check will pass on Saturdays. Defaults to `false`.
	Saturday pulumi.BoolPtrOutput `pulumi:"saturday"`
	// The beginning of the time period that this check will be allowed to pass, specified as 24-hour time with leading zeros.
	StartTime pulumi.StringOutput `pulumi:"startTime"`
	// This check will pass on Sundays. Defaults to `false`.
	Sunday pulumi.BoolPtrOutput `pulumi:"sunday"`
	// The ID of the resource being protected by the check.
	TargetResourceId pulumi.StringOutput `pulumi:"targetResourceId"`
	// The type of resource being protected by the check. Possible values are: `endpoint`, `environment`, `queue`, `repository`, `securefile`, `variablegroup`.
	TargetResourceType pulumi.StringOutput `pulumi:"targetResourceType"`
	// This check will pass on Thursdays. Defaults to `false`.
	Thursday pulumi.BoolPtrOutput `pulumi:"thursday"`
	// The time zone this check will be evaluated in. See below for supported values.
	TimeZone pulumi.StringOutput `pulumi:"timeZone"`
	// The timeout in minutes for the business hours check. Defaults to `1440`.
	Timeout pulumi.IntPtrOutput `pulumi:"timeout"`
	// This check will pass on Tuesday. Defaults to `false`.
	Tuesday pulumi.BoolPtrOutput `pulumi:"tuesday"`
	// The version of the check.
	Version pulumi.IntOutput `pulumi:"version"`
	// This check will pass on Wednesdays. Defaults to `false`.
	Wednesday pulumi.BoolPtrOutput `pulumi:"wednesday"`
}

// NewCheckBusinessHours registers a new resource with the given unique name, arguments, and options.
func NewCheckBusinessHours(ctx *pulumi.Context,
	name string, args *CheckBusinessHoursArgs, opts ...pulumi.ResourceOption) (*CheckBusinessHours, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EndTime == nil {
		return nil, errors.New("invalid value for required argument 'EndTime'")
	}
	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.StartTime == nil {
		return nil, errors.New("invalid value for required argument 'StartTime'")
	}
	if args.TargetResourceId == nil {
		return nil, errors.New("invalid value for required argument 'TargetResourceId'")
	}
	if args.TargetResourceType == nil {
		return nil, errors.New("invalid value for required argument 'TargetResourceType'")
	}
	if args.TimeZone == nil {
		return nil, errors.New("invalid value for required argument 'TimeZone'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource CheckBusinessHours
	err := ctx.RegisterResource("azuredevops:index/checkBusinessHours:CheckBusinessHours", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCheckBusinessHours gets an existing CheckBusinessHours resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCheckBusinessHours(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CheckBusinessHoursState, opts ...pulumi.ResourceOption) (*CheckBusinessHours, error) {
	var resource CheckBusinessHours
	err := ctx.ReadResource("azuredevops:index/checkBusinessHours:CheckBusinessHours", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CheckBusinessHours resources.
type checkBusinessHoursState struct {
	// The name of the business hours check displayed in the web UI.
	DisplayName *string `pulumi:"displayName"`
	// The end of the time period that this check will be allowed to pass, specified as 24-hour time with leading zeros.
	EndTime *string `pulumi:"endTime"`
	// This check will pass on Fridays. Defaults to `false`.
	Friday *bool `pulumi:"friday"`
	// This check will pass on Mondays. Defaults to `false`.
	Monday *bool `pulumi:"monday"`
	// The project ID.
	ProjectId *string `pulumi:"projectId"`
	// This check will pass on Saturdays. Defaults to `false`.
	Saturday *bool `pulumi:"saturday"`
	// The beginning of the time period that this check will be allowed to pass, specified as 24-hour time with leading zeros.
	StartTime *string `pulumi:"startTime"`
	// This check will pass on Sundays. Defaults to `false`.
	Sunday *bool `pulumi:"sunday"`
	// The ID of the resource being protected by the check.
	TargetResourceId *string `pulumi:"targetResourceId"`
	// The type of resource being protected by the check. Possible values are: `endpoint`, `environment`, `queue`, `repository`, `securefile`, `variablegroup`.
	TargetResourceType *string `pulumi:"targetResourceType"`
	// This check will pass on Thursdays. Defaults to `false`.
	Thursday *bool `pulumi:"thursday"`
	// The time zone this check will be evaluated in. See below for supported values.
	TimeZone *string `pulumi:"timeZone"`
	// The timeout in minutes for the business hours check. Defaults to `1440`.
	Timeout *int `pulumi:"timeout"`
	// This check will pass on Tuesday. Defaults to `false`.
	Tuesday *bool `pulumi:"tuesday"`
	// The version of the check.
	Version *int `pulumi:"version"`
	// This check will pass on Wednesdays. Defaults to `false`.
	Wednesday *bool `pulumi:"wednesday"`
}

type CheckBusinessHoursState struct {
	// The name of the business hours check displayed in the web UI.
	DisplayName pulumi.StringPtrInput
	// The end of the time period that this check will be allowed to pass, specified as 24-hour time with leading zeros.
	EndTime pulumi.StringPtrInput
	// This check will pass on Fridays. Defaults to `false`.
	Friday pulumi.BoolPtrInput
	// This check will pass on Mondays. Defaults to `false`.
	Monday pulumi.BoolPtrInput
	// The project ID.
	ProjectId pulumi.StringPtrInput
	// This check will pass on Saturdays. Defaults to `false`.
	Saturday pulumi.BoolPtrInput
	// The beginning of the time period that this check will be allowed to pass, specified as 24-hour time with leading zeros.
	StartTime pulumi.StringPtrInput
	// This check will pass on Sundays. Defaults to `false`.
	Sunday pulumi.BoolPtrInput
	// The ID of the resource being protected by the check.
	TargetResourceId pulumi.StringPtrInput
	// The type of resource being protected by the check. Possible values are: `endpoint`, `environment`, `queue`, `repository`, `securefile`, `variablegroup`.
	TargetResourceType pulumi.StringPtrInput
	// This check will pass on Thursdays. Defaults to `false`.
	Thursday pulumi.BoolPtrInput
	// The time zone this check will be evaluated in. See below for supported values.
	TimeZone pulumi.StringPtrInput
	// The timeout in minutes for the business hours check. Defaults to `1440`.
	Timeout pulumi.IntPtrInput
	// This check will pass on Tuesday. Defaults to `false`.
	Tuesday pulumi.BoolPtrInput
	// The version of the check.
	Version pulumi.IntPtrInput
	// This check will pass on Wednesdays. Defaults to `false`.
	Wednesday pulumi.BoolPtrInput
}

func (CheckBusinessHoursState) ElementType() reflect.Type {
	return reflect.TypeOf((*checkBusinessHoursState)(nil)).Elem()
}

type checkBusinessHoursArgs struct {
	// The name of the business hours check displayed in the web UI.
	DisplayName *string `pulumi:"displayName"`
	// The end of the time period that this check will be allowed to pass, specified as 24-hour time with leading zeros.
	EndTime string `pulumi:"endTime"`
	// This check will pass on Fridays. Defaults to `false`.
	Friday *bool `pulumi:"friday"`
	// This check will pass on Mondays. Defaults to `false`.
	Monday *bool `pulumi:"monday"`
	// The project ID.
	ProjectId string `pulumi:"projectId"`
	// This check will pass on Saturdays. Defaults to `false`.
	Saturday *bool `pulumi:"saturday"`
	// The beginning of the time period that this check will be allowed to pass, specified as 24-hour time with leading zeros.
	StartTime string `pulumi:"startTime"`
	// This check will pass on Sundays. Defaults to `false`.
	Sunday *bool `pulumi:"sunday"`
	// The ID of the resource being protected by the check.
	TargetResourceId string `pulumi:"targetResourceId"`
	// The type of resource being protected by the check. Possible values are: `endpoint`, `environment`, `queue`, `repository`, `securefile`, `variablegroup`.
	TargetResourceType string `pulumi:"targetResourceType"`
	// This check will pass on Thursdays. Defaults to `false`.
	Thursday *bool `pulumi:"thursday"`
	// The time zone this check will be evaluated in. See below for supported values.
	TimeZone string `pulumi:"timeZone"`
	// The timeout in minutes for the business hours check. Defaults to `1440`.
	Timeout *int `pulumi:"timeout"`
	// This check will pass on Tuesday. Defaults to `false`.
	Tuesday *bool `pulumi:"tuesday"`
	// This check will pass on Wednesdays. Defaults to `false`.
	Wednesday *bool `pulumi:"wednesday"`
}

// The set of arguments for constructing a CheckBusinessHours resource.
type CheckBusinessHoursArgs struct {
	// The name of the business hours check displayed in the web UI.
	DisplayName pulumi.StringPtrInput
	// The end of the time period that this check will be allowed to pass, specified as 24-hour time with leading zeros.
	EndTime pulumi.StringInput
	// This check will pass on Fridays. Defaults to `false`.
	Friday pulumi.BoolPtrInput
	// This check will pass on Mondays. Defaults to `false`.
	Monday pulumi.BoolPtrInput
	// The project ID.
	ProjectId pulumi.StringInput
	// This check will pass on Saturdays. Defaults to `false`.
	Saturday pulumi.BoolPtrInput
	// The beginning of the time period that this check will be allowed to pass, specified as 24-hour time with leading zeros.
	StartTime pulumi.StringInput
	// This check will pass on Sundays. Defaults to `false`.
	Sunday pulumi.BoolPtrInput
	// The ID of the resource being protected by the check.
	TargetResourceId pulumi.StringInput
	// The type of resource being protected by the check. Possible values are: `endpoint`, `environment`, `queue`, `repository`, `securefile`, `variablegroup`.
	TargetResourceType pulumi.StringInput
	// This check will pass on Thursdays. Defaults to `false`.
	Thursday pulumi.BoolPtrInput
	// The time zone this check will be evaluated in. See below for supported values.
	TimeZone pulumi.StringInput
	// The timeout in minutes for the business hours check. Defaults to `1440`.
	Timeout pulumi.IntPtrInput
	// This check will pass on Tuesday. Defaults to `false`.
	Tuesday pulumi.BoolPtrInput
	// This check will pass on Wednesdays. Defaults to `false`.
	Wednesday pulumi.BoolPtrInput
}

func (CheckBusinessHoursArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*checkBusinessHoursArgs)(nil)).Elem()
}

type CheckBusinessHoursInput interface {
	pulumi.Input

	ToCheckBusinessHoursOutput() CheckBusinessHoursOutput
	ToCheckBusinessHoursOutputWithContext(ctx context.Context) CheckBusinessHoursOutput
}

func (*CheckBusinessHours) ElementType() reflect.Type {
	return reflect.TypeOf((**CheckBusinessHours)(nil)).Elem()
}

func (i *CheckBusinessHours) ToCheckBusinessHoursOutput() CheckBusinessHoursOutput {
	return i.ToCheckBusinessHoursOutputWithContext(context.Background())
}

func (i *CheckBusinessHours) ToCheckBusinessHoursOutputWithContext(ctx context.Context) CheckBusinessHoursOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CheckBusinessHoursOutput)
}

// CheckBusinessHoursArrayInput is an input type that accepts CheckBusinessHoursArray and CheckBusinessHoursArrayOutput values.
// You can construct a concrete instance of `CheckBusinessHoursArrayInput` via:
//
//	CheckBusinessHoursArray{ CheckBusinessHoursArgs{...} }
type CheckBusinessHoursArrayInput interface {
	pulumi.Input

	ToCheckBusinessHoursArrayOutput() CheckBusinessHoursArrayOutput
	ToCheckBusinessHoursArrayOutputWithContext(context.Context) CheckBusinessHoursArrayOutput
}

type CheckBusinessHoursArray []CheckBusinessHoursInput

func (CheckBusinessHoursArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CheckBusinessHours)(nil)).Elem()
}

func (i CheckBusinessHoursArray) ToCheckBusinessHoursArrayOutput() CheckBusinessHoursArrayOutput {
	return i.ToCheckBusinessHoursArrayOutputWithContext(context.Background())
}

func (i CheckBusinessHoursArray) ToCheckBusinessHoursArrayOutputWithContext(ctx context.Context) CheckBusinessHoursArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CheckBusinessHoursArrayOutput)
}

// CheckBusinessHoursMapInput is an input type that accepts CheckBusinessHoursMap and CheckBusinessHoursMapOutput values.
// You can construct a concrete instance of `CheckBusinessHoursMapInput` via:
//
//	CheckBusinessHoursMap{ "key": CheckBusinessHoursArgs{...} }
type CheckBusinessHoursMapInput interface {
	pulumi.Input

	ToCheckBusinessHoursMapOutput() CheckBusinessHoursMapOutput
	ToCheckBusinessHoursMapOutputWithContext(context.Context) CheckBusinessHoursMapOutput
}

type CheckBusinessHoursMap map[string]CheckBusinessHoursInput

func (CheckBusinessHoursMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CheckBusinessHours)(nil)).Elem()
}

func (i CheckBusinessHoursMap) ToCheckBusinessHoursMapOutput() CheckBusinessHoursMapOutput {
	return i.ToCheckBusinessHoursMapOutputWithContext(context.Background())
}

func (i CheckBusinessHoursMap) ToCheckBusinessHoursMapOutputWithContext(ctx context.Context) CheckBusinessHoursMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CheckBusinessHoursMapOutput)
}

type CheckBusinessHoursOutput struct{ *pulumi.OutputState }

func (CheckBusinessHoursOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CheckBusinessHours)(nil)).Elem()
}

func (o CheckBusinessHoursOutput) ToCheckBusinessHoursOutput() CheckBusinessHoursOutput {
	return o
}

func (o CheckBusinessHoursOutput) ToCheckBusinessHoursOutputWithContext(ctx context.Context) CheckBusinessHoursOutput {
	return o
}

// The name of the business hours check displayed in the web UI.
func (o CheckBusinessHoursOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CheckBusinessHours) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// The end of the time period that this check will be allowed to pass, specified as 24-hour time with leading zeros.
func (o CheckBusinessHoursOutput) EndTime() pulumi.StringOutput {
	return o.ApplyT(func(v *CheckBusinessHours) pulumi.StringOutput { return v.EndTime }).(pulumi.StringOutput)
}

// This check will pass on Fridays. Defaults to `false`.
func (o CheckBusinessHoursOutput) Friday() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CheckBusinessHours) pulumi.BoolPtrOutput { return v.Friday }).(pulumi.BoolPtrOutput)
}

// This check will pass on Mondays. Defaults to `false`.
func (o CheckBusinessHoursOutput) Monday() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CheckBusinessHours) pulumi.BoolPtrOutput { return v.Monday }).(pulumi.BoolPtrOutput)
}

// The project ID.
func (o CheckBusinessHoursOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *CheckBusinessHours) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// This check will pass on Saturdays. Defaults to `false`.
func (o CheckBusinessHoursOutput) Saturday() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CheckBusinessHours) pulumi.BoolPtrOutput { return v.Saturday }).(pulumi.BoolPtrOutput)
}

// The beginning of the time period that this check will be allowed to pass, specified as 24-hour time with leading zeros.
func (o CheckBusinessHoursOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v *CheckBusinessHours) pulumi.StringOutput { return v.StartTime }).(pulumi.StringOutput)
}

// This check will pass on Sundays. Defaults to `false`.
func (o CheckBusinessHoursOutput) Sunday() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CheckBusinessHours) pulumi.BoolPtrOutput { return v.Sunday }).(pulumi.BoolPtrOutput)
}

// The ID of the resource being protected by the check.
func (o CheckBusinessHoursOutput) TargetResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *CheckBusinessHours) pulumi.StringOutput { return v.TargetResourceId }).(pulumi.StringOutput)
}

// The type of resource being protected by the check. Possible values are: `endpoint`, `environment`, `queue`, `repository`, `securefile`, `variablegroup`.
func (o CheckBusinessHoursOutput) TargetResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *CheckBusinessHours) pulumi.StringOutput { return v.TargetResourceType }).(pulumi.StringOutput)
}

// This check will pass on Thursdays. Defaults to `false`.
func (o CheckBusinessHoursOutput) Thursday() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CheckBusinessHours) pulumi.BoolPtrOutput { return v.Thursday }).(pulumi.BoolPtrOutput)
}

// The time zone this check will be evaluated in. See below for supported values.
func (o CheckBusinessHoursOutput) TimeZone() pulumi.StringOutput {
	return o.ApplyT(func(v *CheckBusinessHours) pulumi.StringOutput { return v.TimeZone }).(pulumi.StringOutput)
}

// The timeout in minutes for the business hours check. Defaults to `1440`.
func (o CheckBusinessHoursOutput) Timeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CheckBusinessHours) pulumi.IntPtrOutput { return v.Timeout }).(pulumi.IntPtrOutput)
}

// This check will pass on Tuesday. Defaults to `false`.
func (o CheckBusinessHoursOutput) Tuesday() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CheckBusinessHours) pulumi.BoolPtrOutput { return v.Tuesday }).(pulumi.BoolPtrOutput)
}

// The version of the check.
func (o CheckBusinessHoursOutput) Version() pulumi.IntOutput {
	return o.ApplyT(func(v *CheckBusinessHours) pulumi.IntOutput { return v.Version }).(pulumi.IntOutput)
}

// This check will pass on Wednesdays. Defaults to `false`.
func (o CheckBusinessHoursOutput) Wednesday() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CheckBusinessHours) pulumi.BoolPtrOutput { return v.Wednesday }).(pulumi.BoolPtrOutput)
}

type CheckBusinessHoursArrayOutput struct{ *pulumi.OutputState }

func (CheckBusinessHoursArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CheckBusinessHours)(nil)).Elem()
}

func (o CheckBusinessHoursArrayOutput) ToCheckBusinessHoursArrayOutput() CheckBusinessHoursArrayOutput {
	return o
}

func (o CheckBusinessHoursArrayOutput) ToCheckBusinessHoursArrayOutputWithContext(ctx context.Context) CheckBusinessHoursArrayOutput {
	return o
}

func (o CheckBusinessHoursArrayOutput) Index(i pulumi.IntInput) CheckBusinessHoursOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CheckBusinessHours {
		return vs[0].([]*CheckBusinessHours)[vs[1].(int)]
	}).(CheckBusinessHoursOutput)
}

type CheckBusinessHoursMapOutput struct{ *pulumi.OutputState }

func (CheckBusinessHoursMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CheckBusinessHours)(nil)).Elem()
}

func (o CheckBusinessHoursMapOutput) ToCheckBusinessHoursMapOutput() CheckBusinessHoursMapOutput {
	return o
}

func (o CheckBusinessHoursMapOutput) ToCheckBusinessHoursMapOutputWithContext(ctx context.Context) CheckBusinessHoursMapOutput {
	return o
}

func (o CheckBusinessHoursMapOutput) MapIndex(k pulumi.StringInput) CheckBusinessHoursOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CheckBusinessHours {
		return vs[0].(map[string]*CheckBusinessHours)[vs[1].(string)]
	}).(CheckBusinessHoursOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CheckBusinessHoursInput)(nil)).Elem(), &CheckBusinessHours{})
	pulumi.RegisterInputType(reflect.TypeOf((*CheckBusinessHoursArrayInput)(nil)).Elem(), CheckBusinessHoursArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CheckBusinessHoursMapInput)(nil)).Elem(), CheckBusinessHoursMap{})
	pulumi.RegisterOutputType(CheckBusinessHoursOutput{})
	pulumi.RegisterOutputType(CheckBusinessHoursArrayOutput{})
	pulumi.RegisterOutputType(CheckBusinessHoursMapOutput{})
}
