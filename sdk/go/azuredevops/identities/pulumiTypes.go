// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package identities

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azuredevops/sdk/v2/go/azuredevops/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = internal.GetEnvOrDefault

type GetUsersUser struct {
	// The descriptor is the primary way to reference the graph subject while the system is running. This field will uniquely identify the same graph subject across both Accounts and Organizations.
	Descriptor string `pulumi:"descriptor"`
	// This is the non-unique display name of the graph subject. To change this field, you must alter its value in the source provider.
	DisplayName string `pulumi:"displayName"`
	// The user ID.
	Id string `pulumi:"id"`
	// The email address of record for a given graph member. This may be different than the principal name.
	MailAddress string `pulumi:"mailAddress"`
	// The type of source provider for the `originId` parameter (ex:AD, AAD, MSA) The supported origins are listed below.
	Origin string `pulumi:"origin"`
	// The unique identifier from the system of origin.
	//
	// DataSource without specifying any arguments will return all users inside an organization.
	//
	// List of possible subject types
	//
	// ```go
	// package main
	//
	// import (
	// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	// )
	//
	// func main() {
	// 	pulumi.Run(func(ctx *pulumi.Context) error {
	// 		return nil
	// 	})
	// }
	// ```
	//
	// List of possible origins
	//
	// ```go
	// package main
	//
	// import (
	// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	// )
	//
	// func main() {
	// 	pulumi.Run(func(ctx *pulumi.Context) error {
	// 		return nil
	// 	})
	// }
	// ```
	OriginId *string `pulumi:"originId"`
	// The PrincipalName of this graph member from the source provider.
	PrincipalName string `pulumi:"principalName"`
}

// GetUsersUserInput is an input type that accepts GetUsersUserArgs and GetUsersUserOutput values.
// You can construct a concrete instance of `GetUsersUserInput` via:
//
//	GetUsersUserArgs{...}
type GetUsersUserInput interface {
	pulumi.Input

	ToGetUsersUserOutput() GetUsersUserOutput
	ToGetUsersUserOutputWithContext(context.Context) GetUsersUserOutput
}

type GetUsersUserArgs struct {
	// The descriptor is the primary way to reference the graph subject while the system is running. This field will uniquely identify the same graph subject across both Accounts and Organizations.
	Descriptor pulumi.StringInput `pulumi:"descriptor"`
	// This is the non-unique display name of the graph subject. To change this field, you must alter its value in the source provider.
	DisplayName pulumi.StringInput `pulumi:"displayName"`
	// The user ID.
	Id pulumi.StringInput `pulumi:"id"`
	// The email address of record for a given graph member. This may be different than the principal name.
	MailAddress pulumi.StringInput `pulumi:"mailAddress"`
	// The type of source provider for the `originId` parameter (ex:AD, AAD, MSA) The supported origins are listed below.
	Origin pulumi.StringInput `pulumi:"origin"`
	// The unique identifier from the system of origin.
	//
	// DataSource without specifying any arguments will return all users inside an organization.
	//
	// List of possible subject types
	//
	// ```go
	// package main
	//
	// import (
	// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	// )
	//
	// func main() {
	// 	pulumi.Run(func(ctx *pulumi.Context) error {
	// 		return nil
	// 	})
	// }
	// ```
	//
	// List of possible origins
	//
	// ```go
	// package main
	//
	// import (
	// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	// )
	//
	// func main() {
	// 	pulumi.Run(func(ctx *pulumi.Context) error {
	// 		return nil
	// 	})
	// }
	// ```
	OriginId pulumi.StringPtrInput `pulumi:"originId"`
	// The PrincipalName of this graph member from the source provider.
	PrincipalName pulumi.StringInput `pulumi:"principalName"`
}

func (GetUsersUserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetUsersUser)(nil)).Elem()
}

func (i GetUsersUserArgs) ToGetUsersUserOutput() GetUsersUserOutput {
	return i.ToGetUsersUserOutputWithContext(context.Background())
}

func (i GetUsersUserArgs) ToGetUsersUserOutputWithContext(ctx context.Context) GetUsersUserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetUsersUserOutput)
}

// GetUsersUserArrayInput is an input type that accepts GetUsersUserArray and GetUsersUserArrayOutput values.
// You can construct a concrete instance of `GetUsersUserArrayInput` via:
//
//	GetUsersUserArray{ GetUsersUserArgs{...} }
type GetUsersUserArrayInput interface {
	pulumi.Input

	ToGetUsersUserArrayOutput() GetUsersUserArrayOutput
	ToGetUsersUserArrayOutputWithContext(context.Context) GetUsersUserArrayOutput
}

type GetUsersUserArray []GetUsersUserInput

func (GetUsersUserArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetUsersUser)(nil)).Elem()
}

func (i GetUsersUserArray) ToGetUsersUserArrayOutput() GetUsersUserArrayOutput {
	return i.ToGetUsersUserArrayOutputWithContext(context.Background())
}

func (i GetUsersUserArray) ToGetUsersUserArrayOutputWithContext(ctx context.Context) GetUsersUserArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetUsersUserArrayOutput)
}

type GetUsersUserOutput struct{ *pulumi.OutputState }

func (GetUsersUserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetUsersUser)(nil)).Elem()
}

func (o GetUsersUserOutput) ToGetUsersUserOutput() GetUsersUserOutput {
	return o
}

func (o GetUsersUserOutput) ToGetUsersUserOutputWithContext(ctx context.Context) GetUsersUserOutput {
	return o
}

// The descriptor is the primary way to reference the graph subject while the system is running. This field will uniquely identify the same graph subject across both Accounts and Organizations.
func (o GetUsersUserOutput) Descriptor() pulumi.StringOutput {
	return o.ApplyT(func(v GetUsersUser) string { return v.Descriptor }).(pulumi.StringOutput)
}

// This is the non-unique display name of the graph subject. To change this field, you must alter its value in the source provider.
func (o GetUsersUserOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v GetUsersUser) string { return v.DisplayName }).(pulumi.StringOutput)
}

// The user ID.
func (o GetUsersUserOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetUsersUser) string { return v.Id }).(pulumi.StringOutput)
}

// The email address of record for a given graph member. This may be different than the principal name.
func (o GetUsersUserOutput) MailAddress() pulumi.StringOutput {
	return o.ApplyT(func(v GetUsersUser) string { return v.MailAddress }).(pulumi.StringOutput)
}

// The type of source provider for the `originId` parameter (ex:AD, AAD, MSA) The supported origins are listed below.
func (o GetUsersUserOutput) Origin() pulumi.StringOutput {
	return o.ApplyT(func(v GetUsersUser) string { return v.Origin }).(pulumi.StringOutput)
}

// The unique identifier from the system of origin.
//
// DataSource without specifying any arguments will return all users inside an organization.
//
// # List of possible subject types
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			return nil
//		})
//	}
//
// ```
//
// # List of possible origins
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			return nil
//		})
//	}
//
// ```
func (o GetUsersUserOutput) OriginId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetUsersUser) *string { return v.OriginId }).(pulumi.StringPtrOutput)
}

// The PrincipalName of this graph member from the source provider.
func (o GetUsersUserOutput) PrincipalName() pulumi.StringOutput {
	return o.ApplyT(func(v GetUsersUser) string { return v.PrincipalName }).(pulumi.StringOutput)
}

type GetUsersUserArrayOutput struct{ *pulumi.OutputState }

func (GetUsersUserArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetUsersUser)(nil)).Elem()
}

func (o GetUsersUserArrayOutput) ToGetUsersUserArrayOutput() GetUsersUserArrayOutput {
	return o
}

func (o GetUsersUserArrayOutput) ToGetUsersUserArrayOutputWithContext(ctx context.Context) GetUsersUserArrayOutput {
	return o
}

func (o GetUsersUserArrayOutput) Index(i pulumi.IntInput) GetUsersUserOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetUsersUser {
		return vs[0].([]GetUsersUser)[vs[1].(int)]
	}).(GetUsersUserOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GetUsersUserInput)(nil)).Elem(), GetUsersUserArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetUsersUserArrayInput)(nil)).Elem(), GetUsersUserArray{})
	pulumi.RegisterOutputType(GetUsersUserOutput{})
	pulumi.RegisterOutputType(GetUsersUserArrayOutput{})
}
