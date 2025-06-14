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

// Manages a Work Item in Azure Devops.
//
// ## Example Usage
//
// ### Basic usage
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
//			_, err := azuredevops.NewProject(ctx, "example", &azuredevops.ProjectArgs{
//				Name:             pulumi.String("Example Project"),
//				WorkItemTemplate: pulumi.String("Agile"),
//				VersionControl:   pulumi.String("Git"),
//				Visibility:       pulumi.String("private"),
//				Description:      pulumi.String("Managed by Pulumi"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = azuredevops.NewWorkitem(ctx, "example", &azuredevops.WorkitemArgs{
//				ProjectId: pulumi.Any(exampleAzuredevopsProject.Id),
//				Title:     pulumi.String("Example Work Item"),
//				Type:      pulumi.String("Issue"),
//				State:     pulumi.String("Active"),
//				Tags: pulumi.StringArray{
//					pulumi.String("Tag"),
//				},
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
// ### With custom fields
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
//			_, err := azuredevops.NewProject(ctx, "example", &azuredevops.ProjectArgs{
//				Name:             pulumi.String("Example Project"),
//				WorkItemTemplate: pulumi.String("Agile"),
//				VersionControl:   pulumi.String("Git"),
//				Visibility:       pulumi.String("private"),
//				Description:      pulumi.String("Managed by Pulumi"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = azuredevops.NewWorkitem(ctx, "example", &azuredevops.WorkitemArgs{
//				ProjectId: pulumi.Any(exampleAzuredevopsProject.Id),
//				Title:     pulumi.String("Example Work Item"),
//				Type:      pulumi.String("Issue"),
//				State:     pulumi.String("Active"),
//				Tags: pulumi.StringArray{
//					pulumi.String("Tag"),
//				},
//				CustomFields: pulumi.StringMap{
//					"example": pulumi.String("example"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### With Parent Work Item
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
//				Name:             pulumi.String("Example Project"),
//				WorkItemTemplate: pulumi.String("Agile"),
//				VersionControl:   pulumi.String("Git"),
//				Visibility:       pulumi.String("private"),
//			})
//			if err != nil {
//				return err
//			}
//			epic, err := azuredevops.NewWorkitem(ctx, "epic", &azuredevops.WorkitemArgs{
//				ProjectId: example.ID(),
//				Title:     pulumi.String("Example EPIC Title"),
//				Type:      pulumi.String("Epic"),
//				State:     pulumi.String("New"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = azuredevops.NewWorkitem(ctx, "example", &azuredevops.WorkitemArgs{
//				ProjectId: example.ID(),
//				Title:     pulumi.String("Example Work Item"),
//				Type:      pulumi.String("Issue"),
//				State:     pulumi.String("Active"),
//				Tags: pulumi.StringArray{
//					pulumi.String("Tag"),
//				},
//				ParentId: epic.ID(),
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
// ## Import
//
// Azure DevOps Work Item can be imported using the Project ID and Work Item ID, e.g.
//
// ```sh
// $ pulumi import azuredevops:index/workitem:Workitem example 00000000-0000-0000-0000-000000000000/0
// ```
type Workitem struct {
	pulumi.CustomResourceState

	// Specifies the area where the Work Item is used.
	AreaPath pulumi.StringOutput `pulumi:"areaPath"`
	// Specifies a list with Custom Fields for the Work Item.
	CustomFields pulumi.StringMapOutput `pulumi:"customFields"`
	// Specifies the iteration in which the Work Item is used.
	IterationPath pulumi.StringOutput `pulumi:"iterationPath"`
	// The parent work item.
	ParentId pulumi.IntPtrOutput `pulumi:"parentId"`
	// The ID of the Project.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// A `relations` blocks as documented below.
	Relations WorkitemRelationArrayOutput `pulumi:"relations"`
	// The state of the Work Item. The four main states that are defined for the User Story (`Agile`) are `New`, `Active`, `Resolved`, and `Closed`. See [Workflow states](https://learn.microsoft.com/en-us/azure/devops/boards/work-items/workflow-and-state-categories?view=azure-devops&tabs=agile-process#workflow-states) for more details.
	State pulumi.StringOutput `pulumi:"state"`
	// Specifies a list of Tags.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// The Title of the Work Item.
	Title pulumi.StringOutput `pulumi:"title"`
	// The Type of the Work Item. The work item type varies depending on the process used when creating the project(`Agile`, `Basic`, `Scrum`, `Scrum`). See [Work Item Types](https://learn.microsoft.com/en-us/azure/devops/boards/work-items/about-work-items?view=azure-devops) for more details.
	Type pulumi.StringOutput `pulumi:"type"`
	// The URL of the Work Item.
	Url pulumi.StringOutput `pulumi:"url"`
}

// NewWorkitem registers a new resource with the given unique name, arguments, and options.
func NewWorkitem(ctx *pulumi.Context,
	name string, args *WorkitemArgs, opts ...pulumi.ResourceOption) (*Workitem, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.Title == nil {
		return nil, errors.New("invalid value for required argument 'Title'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Workitem
	err := ctx.RegisterResource("azuredevops:index/workitem:Workitem", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkitem gets an existing Workitem resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkitem(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkitemState, opts ...pulumi.ResourceOption) (*Workitem, error) {
	var resource Workitem
	err := ctx.ReadResource("azuredevops:index/workitem:Workitem", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Workitem resources.
type workitemState struct {
	// Specifies the area where the Work Item is used.
	AreaPath *string `pulumi:"areaPath"`
	// Specifies a list with Custom Fields for the Work Item.
	CustomFields map[string]string `pulumi:"customFields"`
	// Specifies the iteration in which the Work Item is used.
	IterationPath *string `pulumi:"iterationPath"`
	// The parent work item.
	ParentId *int `pulumi:"parentId"`
	// The ID of the Project.
	ProjectId *string `pulumi:"projectId"`
	// A `relations` blocks as documented below.
	Relations []WorkitemRelation `pulumi:"relations"`
	// The state of the Work Item. The four main states that are defined for the User Story (`Agile`) are `New`, `Active`, `Resolved`, and `Closed`. See [Workflow states](https://learn.microsoft.com/en-us/azure/devops/boards/work-items/workflow-and-state-categories?view=azure-devops&tabs=agile-process#workflow-states) for more details.
	State *string `pulumi:"state"`
	// Specifies a list of Tags.
	Tags []string `pulumi:"tags"`
	// The Title of the Work Item.
	Title *string `pulumi:"title"`
	// The Type of the Work Item. The work item type varies depending on the process used when creating the project(`Agile`, `Basic`, `Scrum`, `Scrum`). See [Work Item Types](https://learn.microsoft.com/en-us/azure/devops/boards/work-items/about-work-items?view=azure-devops) for more details.
	Type *string `pulumi:"type"`
	// The URL of the Work Item.
	Url *string `pulumi:"url"`
}

type WorkitemState struct {
	// Specifies the area where the Work Item is used.
	AreaPath pulumi.StringPtrInput
	// Specifies a list with Custom Fields for the Work Item.
	CustomFields pulumi.StringMapInput
	// Specifies the iteration in which the Work Item is used.
	IterationPath pulumi.StringPtrInput
	// The parent work item.
	ParentId pulumi.IntPtrInput
	// The ID of the Project.
	ProjectId pulumi.StringPtrInput
	// A `relations` blocks as documented below.
	Relations WorkitemRelationArrayInput
	// The state of the Work Item. The four main states that are defined for the User Story (`Agile`) are `New`, `Active`, `Resolved`, and `Closed`. See [Workflow states](https://learn.microsoft.com/en-us/azure/devops/boards/work-items/workflow-and-state-categories?view=azure-devops&tabs=agile-process#workflow-states) for more details.
	State pulumi.StringPtrInput
	// Specifies a list of Tags.
	Tags pulumi.StringArrayInput
	// The Title of the Work Item.
	Title pulumi.StringPtrInput
	// The Type of the Work Item. The work item type varies depending on the process used when creating the project(`Agile`, `Basic`, `Scrum`, `Scrum`). See [Work Item Types](https://learn.microsoft.com/en-us/azure/devops/boards/work-items/about-work-items?view=azure-devops) for more details.
	Type pulumi.StringPtrInput
	// The URL of the Work Item.
	Url pulumi.StringPtrInput
}

func (WorkitemState) ElementType() reflect.Type {
	return reflect.TypeOf((*workitemState)(nil)).Elem()
}

type workitemArgs struct {
	// Specifies the area where the Work Item is used.
	AreaPath *string `pulumi:"areaPath"`
	// Specifies a list with Custom Fields for the Work Item.
	CustomFields map[string]string `pulumi:"customFields"`
	// Specifies the iteration in which the Work Item is used.
	IterationPath *string `pulumi:"iterationPath"`
	// The parent work item.
	ParentId *int `pulumi:"parentId"`
	// The ID of the Project.
	ProjectId string `pulumi:"projectId"`
	// The state of the Work Item. The four main states that are defined for the User Story (`Agile`) are `New`, `Active`, `Resolved`, and `Closed`. See [Workflow states](https://learn.microsoft.com/en-us/azure/devops/boards/work-items/workflow-and-state-categories?view=azure-devops&tabs=agile-process#workflow-states) for more details.
	State *string `pulumi:"state"`
	// Specifies a list of Tags.
	Tags []string `pulumi:"tags"`
	// The Title of the Work Item.
	Title string `pulumi:"title"`
	// The Type of the Work Item. The work item type varies depending on the process used when creating the project(`Agile`, `Basic`, `Scrum`, `Scrum`). See [Work Item Types](https://learn.microsoft.com/en-us/azure/devops/boards/work-items/about-work-items?view=azure-devops) for more details.
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a Workitem resource.
type WorkitemArgs struct {
	// Specifies the area where the Work Item is used.
	AreaPath pulumi.StringPtrInput
	// Specifies a list with Custom Fields for the Work Item.
	CustomFields pulumi.StringMapInput
	// Specifies the iteration in which the Work Item is used.
	IterationPath pulumi.StringPtrInput
	// The parent work item.
	ParentId pulumi.IntPtrInput
	// The ID of the Project.
	ProjectId pulumi.StringInput
	// The state of the Work Item. The four main states that are defined for the User Story (`Agile`) are `New`, `Active`, `Resolved`, and `Closed`. See [Workflow states](https://learn.microsoft.com/en-us/azure/devops/boards/work-items/workflow-and-state-categories?view=azure-devops&tabs=agile-process#workflow-states) for more details.
	State pulumi.StringPtrInput
	// Specifies a list of Tags.
	Tags pulumi.StringArrayInput
	// The Title of the Work Item.
	Title pulumi.StringInput
	// The Type of the Work Item. The work item type varies depending on the process used when creating the project(`Agile`, `Basic`, `Scrum`, `Scrum`). See [Work Item Types](https://learn.microsoft.com/en-us/azure/devops/boards/work-items/about-work-items?view=azure-devops) for more details.
	Type pulumi.StringInput
}

func (WorkitemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workitemArgs)(nil)).Elem()
}

type WorkitemInput interface {
	pulumi.Input

	ToWorkitemOutput() WorkitemOutput
	ToWorkitemOutputWithContext(ctx context.Context) WorkitemOutput
}

func (*Workitem) ElementType() reflect.Type {
	return reflect.TypeOf((**Workitem)(nil)).Elem()
}

func (i *Workitem) ToWorkitemOutput() WorkitemOutput {
	return i.ToWorkitemOutputWithContext(context.Background())
}

func (i *Workitem) ToWorkitemOutputWithContext(ctx context.Context) WorkitemOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkitemOutput)
}

// WorkitemArrayInput is an input type that accepts WorkitemArray and WorkitemArrayOutput values.
// You can construct a concrete instance of `WorkitemArrayInput` via:
//
//	WorkitemArray{ WorkitemArgs{...} }
type WorkitemArrayInput interface {
	pulumi.Input

	ToWorkitemArrayOutput() WorkitemArrayOutput
	ToWorkitemArrayOutputWithContext(context.Context) WorkitemArrayOutput
}

type WorkitemArray []WorkitemInput

func (WorkitemArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Workitem)(nil)).Elem()
}

func (i WorkitemArray) ToWorkitemArrayOutput() WorkitemArrayOutput {
	return i.ToWorkitemArrayOutputWithContext(context.Background())
}

func (i WorkitemArray) ToWorkitemArrayOutputWithContext(ctx context.Context) WorkitemArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkitemArrayOutput)
}

// WorkitemMapInput is an input type that accepts WorkitemMap and WorkitemMapOutput values.
// You can construct a concrete instance of `WorkitemMapInput` via:
//
//	WorkitemMap{ "key": WorkitemArgs{...} }
type WorkitemMapInput interface {
	pulumi.Input

	ToWorkitemMapOutput() WorkitemMapOutput
	ToWorkitemMapOutputWithContext(context.Context) WorkitemMapOutput
}

type WorkitemMap map[string]WorkitemInput

func (WorkitemMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Workitem)(nil)).Elem()
}

func (i WorkitemMap) ToWorkitemMapOutput() WorkitemMapOutput {
	return i.ToWorkitemMapOutputWithContext(context.Background())
}

func (i WorkitemMap) ToWorkitemMapOutputWithContext(ctx context.Context) WorkitemMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkitemMapOutput)
}

type WorkitemOutput struct{ *pulumi.OutputState }

func (WorkitemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Workitem)(nil)).Elem()
}

func (o WorkitemOutput) ToWorkitemOutput() WorkitemOutput {
	return o
}

func (o WorkitemOutput) ToWorkitemOutputWithContext(ctx context.Context) WorkitemOutput {
	return o
}

// Specifies the area where the Work Item is used.
func (o WorkitemOutput) AreaPath() pulumi.StringOutput {
	return o.ApplyT(func(v *Workitem) pulumi.StringOutput { return v.AreaPath }).(pulumi.StringOutput)
}

// Specifies a list with Custom Fields for the Work Item.
func (o WorkitemOutput) CustomFields() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Workitem) pulumi.StringMapOutput { return v.CustomFields }).(pulumi.StringMapOutput)
}

// Specifies the iteration in which the Work Item is used.
func (o WorkitemOutput) IterationPath() pulumi.StringOutput {
	return o.ApplyT(func(v *Workitem) pulumi.StringOutput { return v.IterationPath }).(pulumi.StringOutput)
}

// The parent work item.
func (o WorkitemOutput) ParentId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Workitem) pulumi.IntPtrOutput { return v.ParentId }).(pulumi.IntPtrOutput)
}

// The ID of the Project.
func (o WorkitemOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *Workitem) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// A `relations` blocks as documented below.
func (o WorkitemOutput) Relations() WorkitemRelationArrayOutput {
	return o.ApplyT(func(v *Workitem) WorkitemRelationArrayOutput { return v.Relations }).(WorkitemRelationArrayOutput)
}

// The state of the Work Item. The four main states that are defined for the User Story (`Agile`) are `New`, `Active`, `Resolved`, and `Closed`. See [Workflow states](https://learn.microsoft.com/en-us/azure/devops/boards/work-items/workflow-and-state-categories?view=azure-devops&tabs=agile-process#workflow-states) for more details.
func (o WorkitemOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *Workitem) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// Specifies a list of Tags.
func (o WorkitemOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Workitem) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// The Title of the Work Item.
func (o WorkitemOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func(v *Workitem) pulumi.StringOutput { return v.Title }).(pulumi.StringOutput)
}

// The Type of the Work Item. The work item type varies depending on the process used when creating the project(`Agile`, `Basic`, `Scrum`, `Scrum`). See [Work Item Types](https://learn.microsoft.com/en-us/azure/devops/boards/work-items/about-work-items?view=azure-devops) for more details.
func (o WorkitemOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Workitem) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The URL of the Work Item.
func (o WorkitemOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *Workitem) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

type WorkitemArrayOutput struct{ *pulumi.OutputState }

func (WorkitemArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Workitem)(nil)).Elem()
}

func (o WorkitemArrayOutput) ToWorkitemArrayOutput() WorkitemArrayOutput {
	return o
}

func (o WorkitemArrayOutput) ToWorkitemArrayOutputWithContext(ctx context.Context) WorkitemArrayOutput {
	return o
}

func (o WorkitemArrayOutput) Index(i pulumi.IntInput) WorkitemOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Workitem {
		return vs[0].([]*Workitem)[vs[1].(int)]
	}).(WorkitemOutput)
}

type WorkitemMapOutput struct{ *pulumi.OutputState }

func (WorkitemMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Workitem)(nil)).Elem()
}

func (o WorkitemMapOutput) ToWorkitemMapOutput() WorkitemMapOutput {
	return o
}

func (o WorkitemMapOutput) ToWorkitemMapOutputWithContext(ctx context.Context) WorkitemMapOutput {
	return o
}

func (o WorkitemMapOutput) MapIndex(k pulumi.StringInput) WorkitemOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Workitem {
		return vs[0].(map[string]*Workitem)[vs[1].(string)]
	}).(WorkitemOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WorkitemInput)(nil)).Elem(), &Workitem{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkitemArrayInput)(nil)).Elem(), WorkitemArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkitemMapInput)(nil)).Elem(), WorkitemMap{})
	pulumi.RegisterOutputType(WorkitemOutput{})
	pulumi.RegisterOutputType(WorkitemArrayOutput{})
	pulumi.RegisterOutputType(WorkitemMapOutput{})
}
