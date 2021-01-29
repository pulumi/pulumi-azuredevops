// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package azuredevops

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Bitbucket service endpoint within Azure DevOps.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azuredevops/sdk/go/azuredevops/"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		project, err := azuredevops.NewProject(ctx, "project", &azuredevops.ProjectArgs{
// 			Visibility:       pulumi.String("private"),
// 			VersionControl:   pulumi.String("Git"),
// 			WorkItemTemplate: pulumi.String("Agile"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = azuredevops.NewServiceEndpointBitBucket(ctx, "serviceendpoint", &azuredevops.ServiceEndpointBitBucketArgs{
// 			ProjectId:           project.ID(),
// 			Username:            pulumi.String("username"),
// 			Password:            pulumi.String("password"),
// 			ServiceEndpointName: pulumi.String("Sample Bitbucket"),
// 			Description:         pulumi.String("Managed by Terraform"),
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
// - [Azure DevOps Service REST API 5.1 - Agent Pools](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-5.1)
//
// ## Import
//
// Azure DevOps Service Endpoint Bitbucket can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**
//
// ```sh
//  $ pulumi import azuredevops:index/serviceEndpointBitBucket:ServiceEndpointBitBucket serviceendpoint 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
// ```
type ServiceEndpointBitBucket struct {
	pulumi.CustomResourceState

	Authorization pulumi.StringMapOutput `pulumi:"authorization"`
	Description   pulumi.StringPtrOutput `pulumi:"description"`
	// Bitbucket account password.
	Password pulumi.StringOutput `pulumi:"password"`
	// A bcrypted hash of the attribute 'password'
	PasswordHash pulumi.StringOutput `pulumi:"passwordHash"`
	// The project ID or project name.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The Service Endpoint name.
	ServiceEndpointName pulumi.StringOutput `pulumi:"serviceEndpointName"`
	// Bitbucket account username.
	Username pulumi.StringOutput `pulumi:"username"`
}

// NewServiceEndpointBitBucket registers a new resource with the given unique name, arguments, and options.
func NewServiceEndpointBitBucket(ctx *pulumi.Context,
	name string, args *ServiceEndpointBitBucketArgs, opts ...pulumi.ResourceOption) (*ServiceEndpointBitBucket, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Password == nil {
		return nil, errors.New("invalid value for required argument 'Password'")
	}
	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.ServiceEndpointName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceEndpointName'")
	}
	if args.Username == nil {
		return nil, errors.New("invalid value for required argument 'Username'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azuredevops:ServiceEndpoint/bitBucket:BitBucket"),
		},
	})
	opts = append(opts, aliases)
	var resource ServiceEndpointBitBucket
	err := ctx.RegisterResource("azuredevops:index/serviceEndpointBitBucket:ServiceEndpointBitBucket", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceEndpointBitBucket gets an existing ServiceEndpointBitBucket resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceEndpointBitBucket(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceEndpointBitBucketState, opts ...pulumi.ResourceOption) (*ServiceEndpointBitBucket, error) {
	var resource ServiceEndpointBitBucket
	err := ctx.ReadResource("azuredevops:index/serviceEndpointBitBucket:ServiceEndpointBitBucket", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceEndpointBitBucket resources.
type serviceEndpointBitBucketState struct {
	Authorization map[string]string `pulumi:"authorization"`
	Description   *string           `pulumi:"description"`
	// Bitbucket account password.
	Password *string `pulumi:"password"`
	// A bcrypted hash of the attribute 'password'
	PasswordHash *string `pulumi:"passwordHash"`
	// The project ID or project name.
	ProjectId *string `pulumi:"projectId"`
	// The Service Endpoint name.
	ServiceEndpointName *string `pulumi:"serviceEndpointName"`
	// Bitbucket account username.
	Username *string `pulumi:"username"`
}

type ServiceEndpointBitBucketState struct {
	Authorization pulumi.StringMapInput
	Description   pulumi.StringPtrInput
	// Bitbucket account password.
	Password pulumi.StringPtrInput
	// A bcrypted hash of the attribute 'password'
	PasswordHash pulumi.StringPtrInput
	// The project ID or project name.
	ProjectId pulumi.StringPtrInput
	// The Service Endpoint name.
	ServiceEndpointName pulumi.StringPtrInput
	// Bitbucket account username.
	Username pulumi.StringPtrInput
}

func (ServiceEndpointBitBucketState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceEndpointBitBucketState)(nil)).Elem()
}

type serviceEndpointBitBucketArgs struct {
	Authorization map[string]string `pulumi:"authorization"`
	Description   *string           `pulumi:"description"`
	// Bitbucket account password.
	Password string `pulumi:"password"`
	// The project ID or project name.
	ProjectId string `pulumi:"projectId"`
	// The Service Endpoint name.
	ServiceEndpointName string `pulumi:"serviceEndpointName"`
	// Bitbucket account username.
	Username string `pulumi:"username"`
}

// The set of arguments for constructing a ServiceEndpointBitBucket resource.
type ServiceEndpointBitBucketArgs struct {
	Authorization pulumi.StringMapInput
	Description   pulumi.StringPtrInput
	// Bitbucket account password.
	Password pulumi.StringInput
	// The project ID or project name.
	ProjectId pulumi.StringInput
	// The Service Endpoint name.
	ServiceEndpointName pulumi.StringInput
	// Bitbucket account username.
	Username pulumi.StringInput
}

func (ServiceEndpointBitBucketArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceEndpointBitBucketArgs)(nil)).Elem()
}

type ServiceEndpointBitBucketInput interface {
	pulumi.Input

	ToServiceEndpointBitBucketOutput() ServiceEndpointBitBucketOutput
	ToServiceEndpointBitBucketOutputWithContext(ctx context.Context) ServiceEndpointBitBucketOutput
}

func (*ServiceEndpointBitBucket) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceEndpointBitBucket)(nil))
}

func (i *ServiceEndpointBitBucket) ToServiceEndpointBitBucketOutput() ServiceEndpointBitBucketOutput {
	return i.ToServiceEndpointBitBucketOutputWithContext(context.Background())
}

func (i *ServiceEndpointBitBucket) ToServiceEndpointBitBucketOutputWithContext(ctx context.Context) ServiceEndpointBitBucketOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointBitBucketOutput)
}

type ServiceEndpointBitBucketOutput struct {
	*pulumi.OutputState
}

func (ServiceEndpointBitBucketOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceEndpointBitBucket)(nil))
}

func (o ServiceEndpointBitBucketOutput) ToServiceEndpointBitBucketOutput() ServiceEndpointBitBucketOutput {
	return o
}

func (o ServiceEndpointBitBucketOutput) ToServiceEndpointBitBucketOutputWithContext(ctx context.Context) ServiceEndpointBitBucketOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ServiceEndpointBitBucketOutput{})
}
