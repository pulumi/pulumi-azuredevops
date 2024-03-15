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

// Manage pipeline access permissions to resources.
//
// > **Note** This resource is a replacement for `ResourceAuthorization`.  Pipeline authorizations managed by `ResourceAuthorization` can also be managed by this resource.
//
// > **Note** If both "All Pipeline Authorization" and "Custom Pipeline Authorization" are configured, "All Pipeline Authorization" has higher priority.
//
// ## Example Usage
//
// ### Authorization for all pipelines
//
// <!--Start PulumiCodeChooser -->
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
//			exampleProject, err := azuredevops.NewProject(ctx, "exampleProject", &azuredevops.ProjectArgs{
//				Visibility:       pulumi.String("private"),
//				VersionControl:   pulumi.String("Git"),
//				WorkItemTemplate: pulumi.String("Agile"),
//				Description:      pulumi.String("Managed by Terraform"),
//			})
//			if err != nil {
//				return err
//			}
//			examplePool, err := azuredevops.NewPool(ctx, "examplePool", &azuredevops.PoolArgs{
//				AutoProvision: pulumi.Bool(false),
//				AutoUpdate:    pulumi.Bool(false),
//			})
//			if err != nil {
//				return err
//			}
//			exampleQueue, err := azuredevops.NewQueue(ctx, "exampleQueue", &azuredevops.QueueArgs{
//				ProjectId:   exampleProject.ID(),
//				AgentPoolId: examplePool.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = azuredevops.NewPipelineAuthorization(ctx, "examplePipelineAuthorization", &azuredevops.PipelineAuthorizationArgs{
//				ProjectId:  exampleProject.ID(),
//				ResourceId: exampleQueue.ID(),
//				Type:       pulumi.String("queue"),
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
// ### Authorization for specific pipeline
//
// <!--Start PulumiCodeChooser -->
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
//			exampleProject, err := azuredevops.NewProject(ctx, "exampleProject", &azuredevops.ProjectArgs{
//				Visibility:       pulumi.String("private"),
//				VersionControl:   pulumi.String("Git"),
//				WorkItemTemplate: pulumi.String("Agile"),
//				Description:      pulumi.String("Managed by Terraform"),
//			})
//			if err != nil {
//				return err
//			}
//			examplePool, err := azuredevops.NewPool(ctx, "examplePool", &azuredevops.PoolArgs{
//				AutoProvision: pulumi.Bool(false),
//				AutoUpdate:    pulumi.Bool(false),
//			})
//			if err != nil {
//				return err
//			}
//			exampleQueue, err := azuredevops.NewQueue(ctx, "exampleQueue", &azuredevops.QueueArgs{
//				ProjectId:   exampleProject.ID(),
//				AgentPoolId: examplePool.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			exampleGitRepository := azuredevops.GetGitRepositoryOutput(ctx, azuredevops.GetGitRepositoryOutputArgs{
//				ProjectId: exampleProject.ID(),
//				Name:      pulumi.String("Example Project"),
//			}, nil)
//			exampleBuildDefinition, err := azuredevops.NewBuildDefinition(ctx, "exampleBuildDefinition", &azuredevops.BuildDefinitionArgs{
//				ProjectId: exampleProject.ID(),
//				Repository: &azuredevops.BuildDefinitionRepositoryArgs{
//					RepoType: pulumi.String("TfsGit"),
//					RepoId: exampleGitRepository.ApplyT(func(exampleGitRepository azuredevops.GetGitRepositoryResult) (*string, error) {
//						return &exampleGitRepository.Id, nil
//					}).(pulumi.StringPtrOutput),
//					YmlPath: pulumi.String("azure-pipelines.yml"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = azuredevops.NewPipelineAuthorization(ctx, "examplePipelineAuthorization", &azuredevops.PipelineAuthorizationArgs{
//				ProjectId:  exampleProject.ID(),
//				ResourceId: exampleQueue.ID(),
//				Type:       pulumi.String("queue"),
//				PipelineId: exampleBuildDefinition.ID(),
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
// - [Azure DevOps Service REST API 7.1 - Pipeline Permissions](https://learn.microsoft.com/en-us/rest/api/azure/devops/approvalsandchecks/pipeline-permissions?view=azure-devops-rest-7.1)
type PipelineAuthorization struct {
	pulumi.CustomResourceState

	// The ID of the pipeline. If not configured, all pipelines will be authorized. Changing this forces a new resource to be created.
	PipelineId pulumi.IntPtrOutput `pulumi:"pipelineId"`
	// The  ID of the project. Changing this forces a new resource to be created
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The ID of the resource to authorize. Changing this forces a new resource to be created
	ResourceId pulumi.StringOutput `pulumi:"resourceId"`
	// The type of the resource to authorize. Valid values: `endpoint`, `queue`, `variablegroup`, `environment`, `repository`. Changing this forces a new resource to be created
	//
	// > **Note** `repository` is for AzureDevOps repository. To authorize repository other than
	// Azure DevOps like GitHub you need to use service connection(`endpoint`)  to connect and authorize.
	// Typical process for connecting to GitHub:
	// **Pipeline  <----> Service Connection(`endpoint`) <----> GitHub Repository**
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewPipelineAuthorization registers a new resource with the given unique name, arguments, and options.
func NewPipelineAuthorization(ctx *pulumi.Context,
	name string, args *PipelineAuthorizationArgs, opts ...pulumi.ResourceOption) (*PipelineAuthorization, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.ResourceId == nil {
		return nil, errors.New("invalid value for required argument 'ResourceId'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PipelineAuthorization
	err := ctx.RegisterResource("azuredevops:index/pipelineAuthorization:PipelineAuthorization", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPipelineAuthorization gets an existing PipelineAuthorization resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPipelineAuthorization(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PipelineAuthorizationState, opts ...pulumi.ResourceOption) (*PipelineAuthorization, error) {
	var resource PipelineAuthorization
	err := ctx.ReadResource("azuredevops:index/pipelineAuthorization:PipelineAuthorization", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PipelineAuthorization resources.
type pipelineAuthorizationState struct {
	// The ID of the pipeline. If not configured, all pipelines will be authorized. Changing this forces a new resource to be created.
	PipelineId *int `pulumi:"pipelineId"`
	// The  ID of the project. Changing this forces a new resource to be created
	ProjectId *string `pulumi:"projectId"`
	// The ID of the resource to authorize. Changing this forces a new resource to be created
	ResourceId *string `pulumi:"resourceId"`
	// The type of the resource to authorize. Valid values: `endpoint`, `queue`, `variablegroup`, `environment`, `repository`. Changing this forces a new resource to be created
	//
	// > **Note** `repository` is for AzureDevOps repository. To authorize repository other than
	// Azure DevOps like GitHub you need to use service connection(`endpoint`)  to connect and authorize.
	// Typical process for connecting to GitHub:
	// **Pipeline  <----> Service Connection(`endpoint`) <----> GitHub Repository**
	Type *string `pulumi:"type"`
}

type PipelineAuthorizationState struct {
	// The ID of the pipeline. If not configured, all pipelines will be authorized. Changing this forces a new resource to be created.
	PipelineId pulumi.IntPtrInput
	// The  ID of the project. Changing this forces a new resource to be created
	ProjectId pulumi.StringPtrInput
	// The ID of the resource to authorize. Changing this forces a new resource to be created
	ResourceId pulumi.StringPtrInput
	// The type of the resource to authorize. Valid values: `endpoint`, `queue`, `variablegroup`, `environment`, `repository`. Changing this forces a new resource to be created
	//
	// > **Note** `repository` is for AzureDevOps repository. To authorize repository other than
	// Azure DevOps like GitHub you need to use service connection(`endpoint`)  to connect and authorize.
	// Typical process for connecting to GitHub:
	// **Pipeline  <----> Service Connection(`endpoint`) <----> GitHub Repository**
	Type pulumi.StringPtrInput
}

func (PipelineAuthorizationState) ElementType() reflect.Type {
	return reflect.TypeOf((*pipelineAuthorizationState)(nil)).Elem()
}

type pipelineAuthorizationArgs struct {
	// The ID of the pipeline. If not configured, all pipelines will be authorized. Changing this forces a new resource to be created.
	PipelineId *int `pulumi:"pipelineId"`
	// The  ID of the project. Changing this forces a new resource to be created
	ProjectId string `pulumi:"projectId"`
	// The ID of the resource to authorize. Changing this forces a new resource to be created
	ResourceId string `pulumi:"resourceId"`
	// The type of the resource to authorize. Valid values: `endpoint`, `queue`, `variablegroup`, `environment`, `repository`. Changing this forces a new resource to be created
	//
	// > **Note** `repository` is for AzureDevOps repository. To authorize repository other than
	// Azure DevOps like GitHub you need to use service connection(`endpoint`)  to connect and authorize.
	// Typical process for connecting to GitHub:
	// **Pipeline  <----> Service Connection(`endpoint`) <----> GitHub Repository**
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a PipelineAuthorization resource.
type PipelineAuthorizationArgs struct {
	// The ID of the pipeline. If not configured, all pipelines will be authorized. Changing this forces a new resource to be created.
	PipelineId pulumi.IntPtrInput
	// The  ID of the project. Changing this forces a new resource to be created
	ProjectId pulumi.StringInput
	// The ID of the resource to authorize. Changing this forces a new resource to be created
	ResourceId pulumi.StringInput
	// The type of the resource to authorize. Valid values: `endpoint`, `queue`, `variablegroup`, `environment`, `repository`. Changing this forces a new resource to be created
	//
	// > **Note** `repository` is for AzureDevOps repository. To authorize repository other than
	// Azure DevOps like GitHub you need to use service connection(`endpoint`)  to connect and authorize.
	// Typical process for connecting to GitHub:
	// **Pipeline  <----> Service Connection(`endpoint`) <----> GitHub Repository**
	Type pulumi.StringInput
}

func (PipelineAuthorizationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*pipelineAuthorizationArgs)(nil)).Elem()
}

type PipelineAuthorizationInput interface {
	pulumi.Input

	ToPipelineAuthorizationOutput() PipelineAuthorizationOutput
	ToPipelineAuthorizationOutputWithContext(ctx context.Context) PipelineAuthorizationOutput
}

func (*PipelineAuthorization) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelineAuthorization)(nil)).Elem()
}

func (i *PipelineAuthorization) ToPipelineAuthorizationOutput() PipelineAuthorizationOutput {
	return i.ToPipelineAuthorizationOutputWithContext(context.Background())
}

func (i *PipelineAuthorization) ToPipelineAuthorizationOutputWithContext(ctx context.Context) PipelineAuthorizationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineAuthorizationOutput)
}

// PipelineAuthorizationArrayInput is an input type that accepts PipelineAuthorizationArray and PipelineAuthorizationArrayOutput values.
// You can construct a concrete instance of `PipelineAuthorizationArrayInput` via:
//
//	PipelineAuthorizationArray{ PipelineAuthorizationArgs{...} }
type PipelineAuthorizationArrayInput interface {
	pulumi.Input

	ToPipelineAuthorizationArrayOutput() PipelineAuthorizationArrayOutput
	ToPipelineAuthorizationArrayOutputWithContext(context.Context) PipelineAuthorizationArrayOutput
}

type PipelineAuthorizationArray []PipelineAuthorizationInput

func (PipelineAuthorizationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PipelineAuthorization)(nil)).Elem()
}

func (i PipelineAuthorizationArray) ToPipelineAuthorizationArrayOutput() PipelineAuthorizationArrayOutput {
	return i.ToPipelineAuthorizationArrayOutputWithContext(context.Background())
}

func (i PipelineAuthorizationArray) ToPipelineAuthorizationArrayOutputWithContext(ctx context.Context) PipelineAuthorizationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineAuthorizationArrayOutput)
}

// PipelineAuthorizationMapInput is an input type that accepts PipelineAuthorizationMap and PipelineAuthorizationMapOutput values.
// You can construct a concrete instance of `PipelineAuthorizationMapInput` via:
//
//	PipelineAuthorizationMap{ "key": PipelineAuthorizationArgs{...} }
type PipelineAuthorizationMapInput interface {
	pulumi.Input

	ToPipelineAuthorizationMapOutput() PipelineAuthorizationMapOutput
	ToPipelineAuthorizationMapOutputWithContext(context.Context) PipelineAuthorizationMapOutput
}

type PipelineAuthorizationMap map[string]PipelineAuthorizationInput

func (PipelineAuthorizationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PipelineAuthorization)(nil)).Elem()
}

func (i PipelineAuthorizationMap) ToPipelineAuthorizationMapOutput() PipelineAuthorizationMapOutput {
	return i.ToPipelineAuthorizationMapOutputWithContext(context.Background())
}

func (i PipelineAuthorizationMap) ToPipelineAuthorizationMapOutputWithContext(ctx context.Context) PipelineAuthorizationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineAuthorizationMapOutput)
}

type PipelineAuthorizationOutput struct{ *pulumi.OutputState }

func (PipelineAuthorizationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelineAuthorization)(nil)).Elem()
}

func (o PipelineAuthorizationOutput) ToPipelineAuthorizationOutput() PipelineAuthorizationOutput {
	return o
}

func (o PipelineAuthorizationOutput) ToPipelineAuthorizationOutputWithContext(ctx context.Context) PipelineAuthorizationOutput {
	return o
}

// The ID of the pipeline. If not configured, all pipelines will be authorized. Changing this forces a new resource to be created.
func (o PipelineAuthorizationOutput) PipelineId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *PipelineAuthorization) pulumi.IntPtrOutput { return v.PipelineId }).(pulumi.IntPtrOutput)
}

// The  ID of the project. Changing this forces a new resource to be created
func (o PipelineAuthorizationOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *PipelineAuthorization) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The ID of the resource to authorize. Changing this forces a new resource to be created
func (o PipelineAuthorizationOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *PipelineAuthorization) pulumi.StringOutput { return v.ResourceId }).(pulumi.StringOutput)
}

// The type of the resource to authorize. Valid values: `endpoint`, `queue`, `variablegroup`, `environment`, `repository`. Changing this forces a new resource to be created
//
// > **Note** `repository` is for AzureDevOps repository. To authorize repository other than
// Azure DevOps like GitHub you need to use service connection(`endpoint`)  to connect and authorize.
// Typical process for connecting to GitHub:
// **Pipeline  <----> Service Connection(`endpoint`) <----> GitHub Repository**
func (o PipelineAuthorizationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PipelineAuthorization) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

type PipelineAuthorizationArrayOutput struct{ *pulumi.OutputState }

func (PipelineAuthorizationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PipelineAuthorization)(nil)).Elem()
}

func (o PipelineAuthorizationArrayOutput) ToPipelineAuthorizationArrayOutput() PipelineAuthorizationArrayOutput {
	return o
}

func (o PipelineAuthorizationArrayOutput) ToPipelineAuthorizationArrayOutputWithContext(ctx context.Context) PipelineAuthorizationArrayOutput {
	return o
}

func (o PipelineAuthorizationArrayOutput) Index(i pulumi.IntInput) PipelineAuthorizationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PipelineAuthorization {
		return vs[0].([]*PipelineAuthorization)[vs[1].(int)]
	}).(PipelineAuthorizationOutput)
}

type PipelineAuthorizationMapOutput struct{ *pulumi.OutputState }

func (PipelineAuthorizationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PipelineAuthorization)(nil)).Elem()
}

func (o PipelineAuthorizationMapOutput) ToPipelineAuthorizationMapOutput() PipelineAuthorizationMapOutput {
	return o
}

func (o PipelineAuthorizationMapOutput) ToPipelineAuthorizationMapOutputWithContext(ctx context.Context) PipelineAuthorizationMapOutput {
	return o
}

func (o PipelineAuthorizationMapOutput) MapIndex(k pulumi.StringInput) PipelineAuthorizationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PipelineAuthorization {
		return vs[0].(map[string]*PipelineAuthorization)[vs[1].(string)]
	}).(PipelineAuthorizationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PipelineAuthorizationInput)(nil)).Elem(), &PipelineAuthorization{})
	pulumi.RegisterInputType(reflect.TypeOf((*PipelineAuthorizationArrayInput)(nil)).Elem(), PipelineAuthorizationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PipelineAuthorizationMapInput)(nil)).Elem(), PipelineAuthorizationMap{})
	pulumi.RegisterOutputType(PipelineAuthorizationOutput{})
	pulumi.RegisterOutputType(PipelineAuthorizationArrayOutput{})
	pulumi.RegisterOutputType(PipelineAuthorizationMapOutput{})
}
