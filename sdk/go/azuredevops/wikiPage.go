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

// Manages Wiki pages within Azure DevOps project.
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
//				Name:        pulumi.String("Example Project"),
//				Description: pulumi.String("Managed by Pulumi"),
//			})
//			if err != nil {
//				return err
//			}
//			exampleWiki, err := azuredevops.NewWiki(ctx, "example", &azuredevops.WikiArgs{
//				ProjectId: example.ID(),
//				Name:      pulumi.String("Example project wiki "),
//				Type:      pulumi.String("projectWiki"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = azuredevops.NewWikiPage(ctx, "example", &azuredevops.WikiPageArgs{
//				ProjectId: example.ID(),
//				WikiId:    exampleWiki.ID(),
//				Path:      pulumi.String("/page"),
//				Content:   pulumi.String("content"),
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
// - [Azure DevOps Service REST API 7.1 - Wiki Page](https://learn.microsoft.com/en-us/rest/api/azure/devops/wiki/pages?view=azure-devops-rest-7.1)
type WikiPage struct {
	pulumi.CustomResourceState

	// The content of the wiki page.
	Content pulumi.StringOutput `pulumi:"content"`
	Etag    pulumi.StringOutput `pulumi:"etag"`
	// The path of the wiki page.
	Path pulumi.StringOutput `pulumi:"path"`
	// The ID of the Project.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The ID of the Wiki.
	WikiId pulumi.StringOutput `pulumi:"wikiId"`
}

// NewWikiPage registers a new resource with the given unique name, arguments, and options.
func NewWikiPage(ctx *pulumi.Context,
	name string, args *WikiPageArgs, opts ...pulumi.ResourceOption) (*WikiPage, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Content == nil {
		return nil, errors.New("invalid value for required argument 'Content'")
	}
	if args.Path == nil {
		return nil, errors.New("invalid value for required argument 'Path'")
	}
	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.WikiId == nil {
		return nil, errors.New("invalid value for required argument 'WikiId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource WikiPage
	err := ctx.RegisterResource("azuredevops:index/wikiPage:WikiPage", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWikiPage gets an existing WikiPage resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWikiPage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WikiPageState, opts ...pulumi.ResourceOption) (*WikiPage, error) {
	var resource WikiPage
	err := ctx.ReadResource("azuredevops:index/wikiPage:WikiPage", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WikiPage resources.
type wikiPageState struct {
	// The content of the wiki page.
	Content *string `pulumi:"content"`
	Etag    *string `pulumi:"etag"`
	// The path of the wiki page.
	Path *string `pulumi:"path"`
	// The ID of the Project.
	ProjectId *string `pulumi:"projectId"`
	// The ID of the Wiki.
	WikiId *string `pulumi:"wikiId"`
}

type WikiPageState struct {
	// The content of the wiki page.
	Content pulumi.StringPtrInput
	Etag    pulumi.StringPtrInput
	// The path of the wiki page.
	Path pulumi.StringPtrInput
	// The ID of the Project.
	ProjectId pulumi.StringPtrInput
	// The ID of the Wiki.
	WikiId pulumi.StringPtrInput
}

func (WikiPageState) ElementType() reflect.Type {
	return reflect.TypeOf((*wikiPageState)(nil)).Elem()
}

type wikiPageArgs struct {
	// The content of the wiki page.
	Content string  `pulumi:"content"`
	Etag    *string `pulumi:"etag"`
	// The path of the wiki page.
	Path string `pulumi:"path"`
	// The ID of the Project.
	ProjectId string `pulumi:"projectId"`
	// The ID of the Wiki.
	WikiId string `pulumi:"wikiId"`
}

// The set of arguments for constructing a WikiPage resource.
type WikiPageArgs struct {
	// The content of the wiki page.
	Content pulumi.StringInput
	Etag    pulumi.StringPtrInput
	// The path of the wiki page.
	Path pulumi.StringInput
	// The ID of the Project.
	ProjectId pulumi.StringInput
	// The ID of the Wiki.
	WikiId pulumi.StringInput
}

func (WikiPageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*wikiPageArgs)(nil)).Elem()
}

type WikiPageInput interface {
	pulumi.Input

	ToWikiPageOutput() WikiPageOutput
	ToWikiPageOutputWithContext(ctx context.Context) WikiPageOutput
}

func (*WikiPage) ElementType() reflect.Type {
	return reflect.TypeOf((**WikiPage)(nil)).Elem()
}

func (i *WikiPage) ToWikiPageOutput() WikiPageOutput {
	return i.ToWikiPageOutputWithContext(context.Background())
}

func (i *WikiPage) ToWikiPageOutputWithContext(ctx context.Context) WikiPageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WikiPageOutput)
}

// WikiPageArrayInput is an input type that accepts WikiPageArray and WikiPageArrayOutput values.
// You can construct a concrete instance of `WikiPageArrayInput` via:
//
//	WikiPageArray{ WikiPageArgs{...} }
type WikiPageArrayInput interface {
	pulumi.Input

	ToWikiPageArrayOutput() WikiPageArrayOutput
	ToWikiPageArrayOutputWithContext(context.Context) WikiPageArrayOutput
}

type WikiPageArray []WikiPageInput

func (WikiPageArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WikiPage)(nil)).Elem()
}

func (i WikiPageArray) ToWikiPageArrayOutput() WikiPageArrayOutput {
	return i.ToWikiPageArrayOutputWithContext(context.Background())
}

func (i WikiPageArray) ToWikiPageArrayOutputWithContext(ctx context.Context) WikiPageArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WikiPageArrayOutput)
}

// WikiPageMapInput is an input type that accepts WikiPageMap and WikiPageMapOutput values.
// You can construct a concrete instance of `WikiPageMapInput` via:
//
//	WikiPageMap{ "key": WikiPageArgs{...} }
type WikiPageMapInput interface {
	pulumi.Input

	ToWikiPageMapOutput() WikiPageMapOutput
	ToWikiPageMapOutputWithContext(context.Context) WikiPageMapOutput
}

type WikiPageMap map[string]WikiPageInput

func (WikiPageMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WikiPage)(nil)).Elem()
}

func (i WikiPageMap) ToWikiPageMapOutput() WikiPageMapOutput {
	return i.ToWikiPageMapOutputWithContext(context.Background())
}

func (i WikiPageMap) ToWikiPageMapOutputWithContext(ctx context.Context) WikiPageMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WikiPageMapOutput)
}

type WikiPageOutput struct{ *pulumi.OutputState }

func (WikiPageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WikiPage)(nil)).Elem()
}

func (o WikiPageOutput) ToWikiPageOutput() WikiPageOutput {
	return o
}

func (o WikiPageOutput) ToWikiPageOutputWithContext(ctx context.Context) WikiPageOutput {
	return o
}

// The content of the wiki page.
func (o WikiPageOutput) Content() pulumi.StringOutput {
	return o.ApplyT(func(v *WikiPage) pulumi.StringOutput { return v.Content }).(pulumi.StringOutput)
}

func (o WikiPageOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *WikiPage) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The path of the wiki page.
func (o WikiPageOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v *WikiPage) pulumi.StringOutput { return v.Path }).(pulumi.StringOutput)
}

// The ID of the Project.
func (o WikiPageOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *WikiPage) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The ID of the Wiki.
func (o WikiPageOutput) WikiId() pulumi.StringOutput {
	return o.ApplyT(func(v *WikiPage) pulumi.StringOutput { return v.WikiId }).(pulumi.StringOutput)
}

type WikiPageArrayOutput struct{ *pulumi.OutputState }

func (WikiPageArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WikiPage)(nil)).Elem()
}

func (o WikiPageArrayOutput) ToWikiPageArrayOutput() WikiPageArrayOutput {
	return o
}

func (o WikiPageArrayOutput) ToWikiPageArrayOutputWithContext(ctx context.Context) WikiPageArrayOutput {
	return o
}

func (o WikiPageArrayOutput) Index(i pulumi.IntInput) WikiPageOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *WikiPage {
		return vs[0].([]*WikiPage)[vs[1].(int)]
	}).(WikiPageOutput)
}

type WikiPageMapOutput struct{ *pulumi.OutputState }

func (WikiPageMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WikiPage)(nil)).Elem()
}

func (o WikiPageMapOutput) ToWikiPageMapOutput() WikiPageMapOutput {
	return o
}

func (o WikiPageMapOutput) ToWikiPageMapOutputWithContext(ctx context.Context) WikiPageMapOutput {
	return o
}

func (o WikiPageMapOutput) MapIndex(k pulumi.StringInput) WikiPageOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *WikiPage {
		return vs[0].(map[string]*WikiPage)[vs[1].(string)]
	}).(WikiPageOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WikiPageInput)(nil)).Elem(), &WikiPage{})
	pulumi.RegisterInputType(reflect.TypeOf((*WikiPageArrayInput)(nil)).Elem(), WikiPageArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WikiPageMapInput)(nil)).Elem(), WikiPageMap{})
	pulumi.RegisterOutputType(WikiPageOutput{})
	pulumi.RegisterOutputType(WikiPageArrayOutput{})
	pulumi.RegisterOutputType(WikiPageMapOutput{})
}
