// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Build Definition within Azure DevOps.
//
// ## Example Usage
// ### GitHub Enterprise
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azuredevops/sdk/go/azuredevops"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := azuredevops.NewBuildDefinition(ctx, "sampleDotnetcoreAppRelease", &azuredevops.BuildDefinitionArgs{
// 			ProjectId: pulumi.Any(azuredevops_project.Project.Id),
// 			Path:      pulumi.String("\\ExampleFolder"),
// 			CiTrigger: &azuredevops.BuildDefinitionCiTriggerArgs{
// 				UseYaml: pulumi.Bool(true),
// 			},
// 			Repository: &azuredevops.BuildDefinitionRepositoryArgs{
// 				RepoType:            pulumi.String("GitHubEnterprise"),
// 				RepoId:              pulumi.String("<GitHub Org>/<Repo Name>"),
// 				GithubEnterpriseUrl: pulumi.String("https://github.company.com"),
// 				BranchName:          pulumi.String("master"),
// 				YmlPath:             pulumi.String("azure-pipelines.yml"),
// 				ServiceConnectionId: pulumi.String("..."),
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
// - [Azure DevOps Service REST API 5.1 - Build Definitions](https://docs.microsoft.com/en-us/rest/api/azure/devops/build/definitions?view=azure-devops-rest-5.1)
//
// ## Import
//
// Azure DevOps Build Definitions can be imported using the project name/definitions Id or by the project Guid/definitions Id, e.g.
//
// ```sh
//  $ pulumi import azuredevops:index/buildDefinition:BuildDefinition build "Test Project"/10
// ```
//
//  or
//
// ```sh
//  $ pulumi import azuredevops:index/buildDefinition:BuildDefinition build 00000000-0000-0000-0000-000000000000/0
// ```
type BuildDefinition struct {
	pulumi.CustomResourceState

	// The agent pool that should execute the build.
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
	Revision pulumi.IntOutput `pulumi:"revision"`
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
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azuredevops:Build/buildDefinition:BuildDefinition"),
		},
	})
	opts = append(opts, aliases)
	var resource BuildDefinition
	err := ctx.RegisterResource("azuredevops:index/buildDefinition:BuildDefinition", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azuredevops:index/buildDefinition:BuildDefinition", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BuildDefinition resources.
type buildDefinitionState struct {
	// The agent pool that should execute the build.
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
	Revision *int `pulumi:"revision"`
	// A list of variable group IDs (integers) to link to the build definition.
	VariableGroups []int `pulumi:"variableGroups"`
	// A list of `variable` blocks, as documented below.
	Variables []BuildDefinitionVariable `pulumi:"variables"`
}

type BuildDefinitionState struct {
	// The agent pool that should execute the build.
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
	Revision pulumi.IntPtrInput
	// A list of variable group IDs (integers) to link to the build definition.
	VariableGroups pulumi.IntArrayInput
	// A list of `variable` blocks, as documented below.
	Variables BuildDefinitionVariableArrayInput
}

func (BuildDefinitionState) ElementType() reflect.Type {
	return reflect.TypeOf((*buildDefinitionState)(nil)).Elem()
}

type buildDefinitionArgs struct {
	// The agent pool that should execute the build.
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
	// A list of variable group IDs (integers) to link to the build definition.
	VariableGroups []int `pulumi:"variableGroups"`
	// A list of `variable` blocks, as documented below.
	Variables []BuildDefinitionVariable `pulumi:"variables"`
}

// The set of arguments for constructing a BuildDefinition resource.
type BuildDefinitionArgs struct {
	// The agent pool that should execute the build.
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
	return reflect.TypeOf((*BuildDefinition)(nil))
}

func (i *BuildDefinition) ToBuildDefinitionOutput() BuildDefinitionOutput {
	return i.ToBuildDefinitionOutputWithContext(context.Background())
}

func (i *BuildDefinition) ToBuildDefinitionOutputWithContext(ctx context.Context) BuildDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BuildDefinitionOutput)
}

func (i *BuildDefinition) ToBuildDefinitionPtrOutput() BuildDefinitionPtrOutput {
	return i.ToBuildDefinitionPtrOutputWithContext(context.Background())
}

func (i *BuildDefinition) ToBuildDefinitionPtrOutputWithContext(ctx context.Context) BuildDefinitionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BuildDefinitionPtrOutput)
}

type BuildDefinitionPtrInput interface {
	pulumi.Input

	ToBuildDefinitionPtrOutput() BuildDefinitionPtrOutput
	ToBuildDefinitionPtrOutputWithContext(ctx context.Context) BuildDefinitionPtrOutput
}

type buildDefinitionPtrType BuildDefinitionArgs

func (*buildDefinitionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BuildDefinition)(nil))
}

func (i *buildDefinitionPtrType) ToBuildDefinitionPtrOutput() BuildDefinitionPtrOutput {
	return i.ToBuildDefinitionPtrOutputWithContext(context.Background())
}

func (i *buildDefinitionPtrType) ToBuildDefinitionPtrOutputWithContext(ctx context.Context) BuildDefinitionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BuildDefinitionPtrOutput)
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
	return reflect.TypeOf(([]*BuildDefinition)(nil))
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
	return reflect.TypeOf((map[string]*BuildDefinition)(nil))
}

func (i BuildDefinitionMap) ToBuildDefinitionMapOutput() BuildDefinitionMapOutput {
	return i.ToBuildDefinitionMapOutputWithContext(context.Background())
}

func (i BuildDefinitionMap) ToBuildDefinitionMapOutputWithContext(ctx context.Context) BuildDefinitionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BuildDefinitionMapOutput)
}

type BuildDefinitionOutput struct {
	*pulumi.OutputState
}

func (BuildDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BuildDefinition)(nil))
}

func (o BuildDefinitionOutput) ToBuildDefinitionOutput() BuildDefinitionOutput {
	return o
}

func (o BuildDefinitionOutput) ToBuildDefinitionOutputWithContext(ctx context.Context) BuildDefinitionOutput {
	return o
}

func (o BuildDefinitionOutput) ToBuildDefinitionPtrOutput() BuildDefinitionPtrOutput {
	return o.ToBuildDefinitionPtrOutputWithContext(context.Background())
}

func (o BuildDefinitionOutput) ToBuildDefinitionPtrOutputWithContext(ctx context.Context) BuildDefinitionPtrOutput {
	return o.ApplyT(func(v BuildDefinition) *BuildDefinition {
		return &v
	}).(BuildDefinitionPtrOutput)
}

type BuildDefinitionPtrOutput struct {
	*pulumi.OutputState
}

func (BuildDefinitionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BuildDefinition)(nil))
}

func (o BuildDefinitionPtrOutput) ToBuildDefinitionPtrOutput() BuildDefinitionPtrOutput {
	return o
}

func (o BuildDefinitionPtrOutput) ToBuildDefinitionPtrOutputWithContext(ctx context.Context) BuildDefinitionPtrOutput {
	return o
}

type BuildDefinitionArrayOutput struct{ *pulumi.OutputState }

func (BuildDefinitionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BuildDefinition)(nil))
}

func (o BuildDefinitionArrayOutput) ToBuildDefinitionArrayOutput() BuildDefinitionArrayOutput {
	return o
}

func (o BuildDefinitionArrayOutput) ToBuildDefinitionArrayOutputWithContext(ctx context.Context) BuildDefinitionArrayOutput {
	return o
}

func (o BuildDefinitionArrayOutput) Index(i pulumi.IntInput) BuildDefinitionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) BuildDefinition {
		return vs[0].([]BuildDefinition)[vs[1].(int)]
	}).(BuildDefinitionOutput)
}

type BuildDefinitionMapOutput struct{ *pulumi.OutputState }

func (BuildDefinitionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]BuildDefinition)(nil))
}

func (o BuildDefinitionMapOutput) ToBuildDefinitionMapOutput() BuildDefinitionMapOutput {
	return o
}

func (o BuildDefinitionMapOutput) ToBuildDefinitionMapOutputWithContext(ctx context.Context) BuildDefinitionMapOutput {
	return o
}

func (o BuildDefinitionMapOutput) MapIndex(k pulumi.StringInput) BuildDefinitionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) BuildDefinition {
		return vs[0].(map[string]BuildDefinition)[vs[1].(string)]
	}).(BuildDefinitionOutput)
}

func init() {
	pulumi.RegisterOutputType(BuildDefinitionOutput{})
	pulumi.RegisterOutputType(BuildDefinitionPtrOutput{})
	pulumi.RegisterOutputType(BuildDefinitionArrayOutput{})
	pulumi.RegisterOutputType(BuildDefinitionMapOutput{})
}
