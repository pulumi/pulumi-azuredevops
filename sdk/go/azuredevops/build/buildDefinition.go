// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package build

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Build Definition within Azure DevOps.
//
// ## Example Usage
// ### Tfs
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
// 		exampleProject, err := azuredevops.NewProject(ctx, "exampleProject", &azuredevops.ProjectArgs{
// 			Visibility:       pulumi.String("private"),
// 			VersionControl:   pulumi.String("Git"),
// 			WorkItemTemplate: pulumi.String("Agile"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleGit, err := azuredevops.NewGit(ctx, "exampleGit", &azuredevops.GitArgs{
// 			ProjectId: exampleProject.ID(),
// 			Initialization: &GitInitializationArgs{
// 				InitType: pulumi.String("Clean"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleVariableGroup, err := azuredevops.NewVariableGroup(ctx, "exampleVariableGroup", &azuredevops.VariableGroupArgs{
// 			ProjectId:   exampleProject.ID(),
// 			Description: pulumi.String("Managed by Terraform"),
// 			AllowAccess: pulumi.Bool(true),
// 			Variables: VariableGroupVariableArray{
// 				&VariableGroupVariableArgs{
// 					Name:  pulumi.String("FOO"),
// 					Value: pulumi.String("BAR"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = azuredevops.NewBuildDefinition(ctx, "exampleBuildDefinition", &azuredevops.BuildDefinitionArgs{
// 			ProjectId: exampleProject.ID(),
// 			Path:      pulumi.String("\\ExampleFolder"),
// 			CiTrigger: &BuildDefinitionCiTriggerArgs{
// 				UseYaml: pulumi.Bool(true),
// 			},
// 			Schedules: BuildDefinitionScheduleArray{
// 				&BuildDefinitionScheduleArgs{
// 					BranchFilters: BuildDefinitionScheduleBranchFilterArray{
// 						&BuildDefinitionScheduleBranchFilterArgs{
// 							Includes: pulumi.StringArray{
// 								pulumi.String("master"),
// 							},
// 							Excludes: pulumi.StringArray{
// 								pulumi.String("test"),
// 								pulumi.String("regression"),
// 							},
// 						},
// 					},
// 					DaysToBuilds: pulumi.StringArray{
// 						pulumi.String("Wed"),
// 						pulumi.String("Sun"),
// 					},
// 					ScheduleOnlyWithChanges: pulumi.Bool(true),
// 					StartHours:              pulumi.Int(10),
// 					StartMinutes:            pulumi.Int(59),
// 					TimeZone:                pulumi.String("(UTC) Coordinated Universal Time"),
// 				},
// 			},
// 			Repository: &BuildDefinitionRepositoryArgs{
// 				RepoType:   pulumi.String("TfsGit"),
// 				RepoId:     exampleGit.ID(),
// 				BranchName: exampleGit.DefaultBranch,
// 				YmlPath:    pulumi.String("azure-pipelines.yml"),
// 			},
// 			VariableGroups: pulumi.IntArray{
// 				exampleVariableGroup.ID(),
// 			},
// 			Variables: BuildDefinitionVariableArray{
// 				&BuildDefinitionVariableArgs{
// 					Name:  pulumi.String("PipelineVariable"),
// 					Value: pulumi.String("Go Microsoft!"),
// 				},
// 				&BuildDefinitionVariableArgs{
// 					Name:        pulumi.String("PipelineSecret"),
// 					SecretValue: pulumi.String("ZGV2cw"),
// 					IsSecret:    pulumi.Bool(true),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### GitHub Enterprise
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
// 		exampleProject, err := azuredevops.NewProject(ctx, "exampleProject", &azuredevops.ProjectArgs{
// 			Visibility:       pulumi.String("private"),
// 			VersionControl:   pulumi.String("Git"),
// 			WorkItemTemplate: pulumi.String("Agile"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleServiceEndpointGitHubEnterprise, err := azuredevops.NewServiceEndpointGitHubEnterprise(ctx, "exampleServiceEndpointGitHubEnterprise", &azuredevops.ServiceEndpointGitHubEnterpriseArgs{
// 			ProjectId:           exampleProject.ID(),
// 			ServiceEndpointName: pulumi.String("Example GitHub Enterprise"),
// 			Url:                 pulumi.String("https://github.contoso.com"),
// 			Description:         pulumi.String("Managed by Terraform"),
// 			AuthPersonal: &ServiceEndpointGitHubEnterpriseAuthPersonalArgs{
// 				PersonalAccessToken: pulumi.String("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = azuredevops.NewBuildDefinition(ctx, "exampleBuildDefinition", &azuredevops.BuildDefinitionArgs{
// 			ProjectId: exampleProject.ID(),
// 			Path:      pulumi.String("\\ExampleFolder"),
// 			CiTrigger: &BuildDefinitionCiTriggerArgs{
// 				UseYaml: pulumi.Bool(true),
// 			},
// 			Repository: &BuildDefinitionRepositoryArgs{
// 				RepoType:            pulumi.String("GitHubEnterprise"),
// 				RepoId:              pulumi.String("<GitHub Org>/<Repo Name>"),
// 				GithubEnterpriseUrl: pulumi.String("https://github.company.com"),
// 				BranchName:          pulumi.String("master"),
// 				YmlPath:             pulumi.String("azure-pipelines.yml"),
// 				ServiceConnectionId: exampleServiceEndpointGitHubEnterprise.ID(),
// 			},
// 			Schedules: BuildDefinitionScheduleArray{
// 				&BuildDefinitionScheduleArgs{
// 					BranchFilters: BuildDefinitionScheduleBranchFilterArray{
// 						&BuildDefinitionScheduleBranchFilterArgs{
// 							Includes: pulumi.StringArray{
// 								pulumi.String("main"),
// 							},
// 							Excludes: pulumi.StringArray{
// 								pulumi.String("test"),
// 								pulumi.String("regression"),
// 							},
// 						},
// 					},
// 					DaysToBuilds: pulumi.StringArray{
// 						pulumi.String("Wed"),
// 						pulumi.String("Sun"),
// 					},
// 					ScheduleOnlyWithChanges: pulumi.Bool(true),
// 					StartHours:              pulumi.Int(10),
// 					StartMinutes:            pulumi.Int(59),
// 					TimeZone:                pulumi.String("(UTC) Coordinated Universal Time"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ## Remarks
//
// The path attribute can not end in `\` unless the path is the root value of `\`.
//
// Valid path values (yaml encoded) include:
// - `\\`
// - `\\ExampleFolder`
// - `\\Nested\\Example Folder`
//
// The value of `\\ExampleFolder\\` would be invalid.
//
// ## Relevant Links
//
// - [Azure DevOps Service REST API 6.0 - Build Definitions](https://docs.microsoft.com/en-us/rest/api/azure/devops/build/definitions?view=azure-devops-rest-6.0)
//
// ## Import
//
// Azure DevOps Build Definitions can be imported using the project name/definitions Id or by the project Guid/definitions Id, e.g.
//
// ```sh
//  $ pulumi import azuredevops:Build/buildDefinition:BuildDefinition example "Example Project"/10
// ```
//
//  or
//
// ```sh
//  $ pulumi import azuredevops:Build/buildDefinition:BuildDefinition example 00000000-0000-0000-0000-000000000000/0
// ```
//
// Deprecated: azuredevops.build.BuildDefinition has been deprecated in favor of azuredevops.BuildDefinition
type BuildDefinition struct {
	pulumi.CustomResourceState

	// The agent pool that should execute the build. Defaults to `Azure Pipelines`.
	AgentPoolName pulumi.StringPtrOutput `pulumi:"agentPoolName"`
	// Continuous Integration trigger.
	CiTrigger BuildDefinitionCiTriggerPtrOutput `pulumi:"ciTrigger"`
	// The name of the build definition.
	Name pulumi.StringOutput `pulumi:"name"`
	// The folder path of the build definition.
	Path pulumi.StringPtrOutput `pulumi:"path"`
	// The project ID or project name.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// Pull Request Integration Integration trigger.
	PullRequestTrigger BuildDefinitionPullRequestTriggerPtrOutput `pulumi:"pullRequestTrigger"`
	// A `repository` block as documented below.
	Repository BuildDefinitionRepositoryOutput `pulumi:"repository"`
	// The revision of the build definition
	Revision  pulumi.IntOutput                   `pulumi:"revision"`
	Schedules BuildDefinitionScheduleArrayOutput `pulumi:"schedules"`
	// A list of variable group IDs (integers) to link to the build definition.
	VariableGroups pulumi.IntArrayOutput `pulumi:"variableGroups"`
	// A list of `variable` blocks, as documented below.
	Variables BuildDefinitionVariableArrayOutput `pulumi:"variables"`
}

// NewBuildDefinition registers a new resource with the given unique name, arguments, and options.
func NewBuildDefinition(ctx *pulumi.Context,
	name string, args *BuildDefinitionArgs, opts ...pulumi.ResourceOption) (*BuildDefinition, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.Repository == nil {
		return nil, errors.New("invalid value for required argument 'Repository'")
	}
	var resource BuildDefinition
	err := ctx.RegisterResource("azuredevops:Build/buildDefinition:BuildDefinition", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBuildDefinition gets an existing BuildDefinition resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBuildDefinition(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BuildDefinitionState, opts ...pulumi.ResourceOption) (*BuildDefinition, error) {
	var resource BuildDefinition
	err := ctx.ReadResource("azuredevops:Build/buildDefinition:BuildDefinition", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BuildDefinition resources.
type buildDefinitionState struct {
	// The agent pool that should execute the build. Defaults to `Azure Pipelines`.
	AgentPoolName *string `pulumi:"agentPoolName"`
	// Continuous Integration trigger.
	CiTrigger *BuildDefinitionCiTrigger `pulumi:"ciTrigger"`
	// The name of the build definition.
	Name *string `pulumi:"name"`
	// The folder path of the build definition.
	Path *string `pulumi:"path"`
	// The project ID or project name.
	ProjectId *string `pulumi:"projectId"`
	// Pull Request Integration Integration trigger.
	PullRequestTrigger *BuildDefinitionPullRequestTrigger `pulumi:"pullRequestTrigger"`
	// A `repository` block as documented below.
	Repository *BuildDefinitionRepository `pulumi:"repository"`
	// The revision of the build definition
	Revision  *int                      `pulumi:"revision"`
	Schedules []BuildDefinitionSchedule `pulumi:"schedules"`
	// A list of variable group IDs (integers) to link to the build definition.
	VariableGroups []int `pulumi:"variableGroups"`
	// A list of `variable` blocks, as documented below.
	Variables []BuildDefinitionVariable `pulumi:"variables"`
}

type BuildDefinitionState struct {
	// The agent pool that should execute the build. Defaults to `Azure Pipelines`.
	AgentPoolName pulumi.StringPtrInput
	// Continuous Integration trigger.
	CiTrigger BuildDefinitionCiTriggerPtrInput
	// The name of the build definition.
	Name pulumi.StringPtrInput
	// The folder path of the build definition.
	Path pulumi.StringPtrInput
	// The project ID or project name.
	ProjectId pulumi.StringPtrInput
	// Pull Request Integration Integration trigger.
	PullRequestTrigger BuildDefinitionPullRequestTriggerPtrInput
	// A `repository` block as documented below.
	Repository BuildDefinitionRepositoryPtrInput
	// The revision of the build definition
	Revision  pulumi.IntPtrInput
	Schedules BuildDefinitionScheduleArrayInput
	// A list of variable group IDs (integers) to link to the build definition.
	VariableGroups pulumi.IntArrayInput
	// A list of `variable` blocks, as documented below.
	Variables BuildDefinitionVariableArrayInput
}

func (BuildDefinitionState) ElementType() reflect.Type {
	return reflect.TypeOf((*buildDefinitionState)(nil)).Elem()
}

type buildDefinitionArgs struct {
	// The agent pool that should execute the build. Defaults to `Azure Pipelines`.
	AgentPoolName *string `pulumi:"agentPoolName"`
	// Continuous Integration trigger.
	CiTrigger *BuildDefinitionCiTrigger `pulumi:"ciTrigger"`
	// The name of the build definition.
	Name *string `pulumi:"name"`
	// The folder path of the build definition.
	Path *string `pulumi:"path"`
	// The project ID or project name.
	ProjectId string `pulumi:"projectId"`
	// Pull Request Integration Integration trigger.
	PullRequestTrigger *BuildDefinitionPullRequestTrigger `pulumi:"pullRequestTrigger"`
	// A `repository` block as documented below.
	Repository BuildDefinitionRepository `pulumi:"repository"`
	Schedules  []BuildDefinitionSchedule `pulumi:"schedules"`
	// A list of variable group IDs (integers) to link to the build definition.
	VariableGroups []int `pulumi:"variableGroups"`
	// A list of `variable` blocks, as documented below.
	Variables []BuildDefinitionVariable `pulumi:"variables"`
}

// The set of arguments for constructing a BuildDefinition resource.
type BuildDefinitionArgs struct {
	// The agent pool that should execute the build. Defaults to `Azure Pipelines`.
	AgentPoolName pulumi.StringPtrInput
	// Continuous Integration trigger.
	CiTrigger BuildDefinitionCiTriggerPtrInput
	// The name of the build definition.
	Name pulumi.StringPtrInput
	// The folder path of the build definition.
	Path pulumi.StringPtrInput
	// The project ID or project name.
	ProjectId pulumi.StringInput
	// Pull Request Integration Integration trigger.
	PullRequestTrigger BuildDefinitionPullRequestTriggerPtrInput
	// A `repository` block as documented below.
	Repository BuildDefinitionRepositoryInput
	Schedules  BuildDefinitionScheduleArrayInput
	// A list of variable group IDs (integers) to link to the build definition.
	VariableGroups pulumi.IntArrayInput
	// A list of `variable` blocks, as documented below.
	Variables BuildDefinitionVariableArrayInput
}

func (BuildDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*buildDefinitionArgs)(nil)).Elem()
}

type BuildDefinitionInput interface {
	pulumi.Input

	ToBuildDefinitionOutput() BuildDefinitionOutput
	ToBuildDefinitionOutputWithContext(ctx context.Context) BuildDefinitionOutput
}

func (*BuildDefinition) ElementType() reflect.Type {
	return reflect.TypeOf((**BuildDefinition)(nil)).Elem()
}

func (i *BuildDefinition) ToBuildDefinitionOutput() BuildDefinitionOutput {
	return i.ToBuildDefinitionOutputWithContext(context.Background())
}

func (i *BuildDefinition) ToBuildDefinitionOutputWithContext(ctx context.Context) BuildDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BuildDefinitionOutput)
}

// BuildDefinitionArrayInput is an input type that accepts BuildDefinitionArray and BuildDefinitionArrayOutput values.
// You can construct a concrete instance of `BuildDefinitionArrayInput` via:
//
//          BuildDefinitionArray{ BuildDefinitionArgs{...} }
type BuildDefinitionArrayInput interface {
	pulumi.Input

	ToBuildDefinitionArrayOutput() BuildDefinitionArrayOutput
	ToBuildDefinitionArrayOutputWithContext(context.Context) BuildDefinitionArrayOutput
}

type BuildDefinitionArray []BuildDefinitionInput

func (BuildDefinitionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BuildDefinition)(nil)).Elem()
}

func (i BuildDefinitionArray) ToBuildDefinitionArrayOutput() BuildDefinitionArrayOutput {
	return i.ToBuildDefinitionArrayOutputWithContext(context.Background())
}

func (i BuildDefinitionArray) ToBuildDefinitionArrayOutputWithContext(ctx context.Context) BuildDefinitionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BuildDefinitionArrayOutput)
}

// BuildDefinitionMapInput is an input type that accepts BuildDefinitionMap and BuildDefinitionMapOutput values.
// You can construct a concrete instance of `BuildDefinitionMapInput` via:
//
//          BuildDefinitionMap{ "key": BuildDefinitionArgs{...} }
type BuildDefinitionMapInput interface {
	pulumi.Input

	ToBuildDefinitionMapOutput() BuildDefinitionMapOutput
	ToBuildDefinitionMapOutputWithContext(context.Context) BuildDefinitionMapOutput
}

type BuildDefinitionMap map[string]BuildDefinitionInput

func (BuildDefinitionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BuildDefinition)(nil)).Elem()
}

func (i BuildDefinitionMap) ToBuildDefinitionMapOutput() BuildDefinitionMapOutput {
	return i.ToBuildDefinitionMapOutputWithContext(context.Background())
}

func (i BuildDefinitionMap) ToBuildDefinitionMapOutputWithContext(ctx context.Context) BuildDefinitionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BuildDefinitionMapOutput)
}

type BuildDefinitionOutput struct{ *pulumi.OutputState }

func (BuildDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BuildDefinition)(nil)).Elem()
}

func (o BuildDefinitionOutput) ToBuildDefinitionOutput() BuildDefinitionOutput {
	return o
}

func (o BuildDefinitionOutput) ToBuildDefinitionOutputWithContext(ctx context.Context) BuildDefinitionOutput {
	return o
}

type BuildDefinitionArrayOutput struct{ *pulumi.OutputState }

func (BuildDefinitionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BuildDefinition)(nil)).Elem()
}

func (o BuildDefinitionArrayOutput) ToBuildDefinitionArrayOutput() BuildDefinitionArrayOutput {
	return o
}

func (o BuildDefinitionArrayOutput) ToBuildDefinitionArrayOutputWithContext(ctx context.Context) BuildDefinitionArrayOutput {
	return o
}

func (o BuildDefinitionArrayOutput) Index(i pulumi.IntInput) BuildDefinitionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BuildDefinition {
		return vs[0].([]*BuildDefinition)[vs[1].(int)]
	}).(BuildDefinitionOutput)
}

type BuildDefinitionMapOutput struct{ *pulumi.OutputState }

func (BuildDefinitionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BuildDefinition)(nil)).Elem()
}

func (o BuildDefinitionMapOutput) ToBuildDefinitionMapOutput() BuildDefinitionMapOutput {
	return o
}

func (o BuildDefinitionMapOutput) ToBuildDefinitionMapOutputWithContext(ctx context.Context) BuildDefinitionMapOutput {
	return o
}

func (o BuildDefinitionMapOutput) MapIndex(k pulumi.StringInput) BuildDefinitionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BuildDefinition {
		return vs[0].(map[string]*BuildDefinition)[vs[1].(string)]
	}).(BuildDefinitionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BuildDefinitionInput)(nil)).Elem(), &BuildDefinition{})
	pulumi.RegisterInputType(reflect.TypeOf((*BuildDefinitionArrayInput)(nil)).Elem(), BuildDefinitionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BuildDefinitionMapInput)(nil)).Elem(), BuildDefinitionMap{})
	pulumi.RegisterOutputType(BuildDefinitionOutput{})
	pulumi.RegisterOutputType(BuildDefinitionArrayOutput{})
	pulumi.RegisterOutputType(BuildDefinitionMapOutput{})
}
