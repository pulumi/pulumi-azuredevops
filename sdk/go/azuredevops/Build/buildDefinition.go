// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package Build

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Build Definition within Azure DevOps.
//
//
// ## Relevant Links
//
// * [Azure DevOps Service REST API 5.1 - Build Definitions](https://docs.microsoft.com/en-us/rest/api/azure/devops/build/definitions?view=azure-devops-rest-5.1)
type BuildDefinition struct {
	pulumi.CustomResourceState

	// The agent pool that should execute the build. Defaults to `Hosted Ubuntu 1604`.
	AgentPoolName pulumi.StringPtrOutput `pulumi:"agentPoolName"`
	// Continuous Integration Integration trigger.
	CiTrigger BuildDefinitionCiTriggerPtrOutput `pulumi:"ciTrigger"`
	// The name of the build definition.
	Name pulumi.StringOutput    `pulumi:"name"`
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
	if args == nil || args.ProjectId == nil {
		return nil, errors.New("missing required argument 'ProjectId'")
	}
	if args == nil || args.Repository == nil {
		return nil, errors.New("missing required argument 'Repository'")
	}
	if args == nil {
		args = &BuildDefinitionArgs{}
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
	// The agent pool that should execute the build. Defaults to `Hosted Ubuntu 1604`.
	AgentPoolName *string `pulumi:"agentPoolName"`
	// Continuous Integration Integration trigger.
	CiTrigger *BuildDefinitionCiTrigger `pulumi:"ciTrigger"`
	// The name of the build definition.
	Name *string `pulumi:"name"`
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
	// The agent pool that should execute the build. Defaults to `Hosted Ubuntu 1604`.
	AgentPoolName pulumi.StringPtrInput
	// Continuous Integration Integration trigger.
	CiTrigger BuildDefinitionCiTriggerPtrInput
	// The name of the build definition.
	Name pulumi.StringPtrInput
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
	// The agent pool that should execute the build. Defaults to `Hosted Ubuntu 1604`.
	AgentPoolName *string `pulumi:"agentPoolName"`
	// Continuous Integration Integration trigger.
	CiTrigger *BuildDefinitionCiTrigger `pulumi:"ciTrigger"`
	// The name of the build definition.
	Name *string `pulumi:"name"`
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
	// The agent pool that should execute the build. Defaults to `Hosted Ubuntu 1604`.
	AgentPoolName pulumi.StringPtrInput
	// Continuous Integration Integration trigger.
	CiTrigger BuildDefinitionCiTriggerPtrInput
	// The name of the build definition.
	Name pulumi.StringPtrInput
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
