// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package serviceendpoint

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type AzureRMCredentials struct {
	// The service principal application Id
	Serviceprincipalid string `pulumi:"serviceprincipalid"`
	// The service principal secret.
	Serviceprincipalkey     string  `pulumi:"serviceprincipalkey"`
	ServiceprincipalkeyHash *string `pulumi:"serviceprincipalkeyHash"`
}

// AzureRMCredentialsInput is an input type that accepts AzureRMCredentialsArgs and AzureRMCredentialsOutput values.
// You can construct a concrete instance of `AzureRMCredentialsInput` via:
//
//          AzureRMCredentialsArgs{...}
type AzureRMCredentialsInput interface {
	pulumi.Input

	ToAzureRMCredentialsOutput() AzureRMCredentialsOutput
	ToAzureRMCredentialsOutputWithContext(context.Context) AzureRMCredentialsOutput
}

type AzureRMCredentialsArgs struct {
	// The service principal application Id
	Serviceprincipalid pulumi.StringInput `pulumi:"serviceprincipalid"`
	// The service principal secret.
	Serviceprincipalkey     pulumi.StringInput    `pulumi:"serviceprincipalkey"`
	ServiceprincipalkeyHash pulumi.StringPtrInput `pulumi:"serviceprincipalkeyHash"`
}

func (AzureRMCredentialsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureRMCredentials)(nil)).Elem()
}

func (i AzureRMCredentialsArgs) ToAzureRMCredentialsOutput() AzureRMCredentialsOutput {
	return i.ToAzureRMCredentialsOutputWithContext(context.Background())
}

func (i AzureRMCredentialsArgs) ToAzureRMCredentialsOutputWithContext(ctx context.Context) AzureRMCredentialsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureRMCredentialsOutput)
}

func (i AzureRMCredentialsArgs) ToAzureRMCredentialsPtrOutput() AzureRMCredentialsPtrOutput {
	return i.ToAzureRMCredentialsPtrOutputWithContext(context.Background())
}

func (i AzureRMCredentialsArgs) ToAzureRMCredentialsPtrOutputWithContext(ctx context.Context) AzureRMCredentialsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureRMCredentialsOutput).ToAzureRMCredentialsPtrOutputWithContext(ctx)
}

// AzureRMCredentialsPtrInput is an input type that accepts AzureRMCredentialsArgs, AzureRMCredentialsPtr and AzureRMCredentialsPtrOutput values.
// You can construct a concrete instance of `AzureRMCredentialsPtrInput` via:
//
//          AzureRMCredentialsArgs{...}
//
//  or:
//
//          nil
type AzureRMCredentialsPtrInput interface {
	pulumi.Input

	ToAzureRMCredentialsPtrOutput() AzureRMCredentialsPtrOutput
	ToAzureRMCredentialsPtrOutputWithContext(context.Context) AzureRMCredentialsPtrOutput
}

type azureRMCredentialsPtrType AzureRMCredentialsArgs

func AzureRMCredentialsPtr(v *AzureRMCredentialsArgs) AzureRMCredentialsPtrInput {
	return (*azureRMCredentialsPtrType)(v)
}

func (*azureRMCredentialsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureRMCredentials)(nil)).Elem()
}

func (i *azureRMCredentialsPtrType) ToAzureRMCredentialsPtrOutput() AzureRMCredentialsPtrOutput {
	return i.ToAzureRMCredentialsPtrOutputWithContext(context.Background())
}

func (i *azureRMCredentialsPtrType) ToAzureRMCredentialsPtrOutputWithContext(ctx context.Context) AzureRMCredentialsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureRMCredentialsPtrOutput)
}

type AzureRMCredentialsOutput struct{ *pulumi.OutputState }

func (AzureRMCredentialsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureRMCredentials)(nil)).Elem()
}

func (o AzureRMCredentialsOutput) ToAzureRMCredentialsOutput() AzureRMCredentialsOutput {
	return o
}

func (o AzureRMCredentialsOutput) ToAzureRMCredentialsOutputWithContext(ctx context.Context) AzureRMCredentialsOutput {
	return o
}

func (o AzureRMCredentialsOutput) ToAzureRMCredentialsPtrOutput() AzureRMCredentialsPtrOutput {
	return o.ToAzureRMCredentialsPtrOutputWithContext(context.Background())
}

func (o AzureRMCredentialsOutput) ToAzureRMCredentialsPtrOutputWithContext(ctx context.Context) AzureRMCredentialsPtrOutput {
	return o.ApplyT(func(v AzureRMCredentials) *AzureRMCredentials {
		return &v
	}).(AzureRMCredentialsPtrOutput)
}

// The service principal application Id
func (o AzureRMCredentialsOutput) Serviceprincipalid() pulumi.StringOutput {
	return o.ApplyT(func(v AzureRMCredentials) string { return v.Serviceprincipalid }).(pulumi.StringOutput)
}

// The service principal secret.
func (o AzureRMCredentialsOutput) Serviceprincipalkey() pulumi.StringOutput {
	return o.ApplyT(func(v AzureRMCredentials) string { return v.Serviceprincipalkey }).(pulumi.StringOutput)
}

func (o AzureRMCredentialsOutput) ServiceprincipalkeyHash() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureRMCredentials) *string { return v.ServiceprincipalkeyHash }).(pulumi.StringPtrOutput)
}

type AzureRMCredentialsPtrOutput struct{ *pulumi.OutputState }

func (AzureRMCredentialsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureRMCredentials)(nil)).Elem()
}

func (o AzureRMCredentialsPtrOutput) ToAzureRMCredentialsPtrOutput() AzureRMCredentialsPtrOutput {
	return o
}

func (o AzureRMCredentialsPtrOutput) ToAzureRMCredentialsPtrOutputWithContext(ctx context.Context) AzureRMCredentialsPtrOutput {
	return o
}

func (o AzureRMCredentialsPtrOutput) Elem() AzureRMCredentialsOutput {
	return o.ApplyT(func(v *AzureRMCredentials) AzureRMCredentials { return *v }).(AzureRMCredentialsOutput)
}

// The service principal application Id
func (o AzureRMCredentialsPtrOutput) Serviceprincipalid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureRMCredentials) *string {
		if v == nil {
			return nil
		}
		return &v.Serviceprincipalid
	}).(pulumi.StringPtrOutput)
}

// The service principal secret.
func (o AzureRMCredentialsPtrOutput) Serviceprincipalkey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureRMCredentials) *string {
		if v == nil {
			return nil
		}
		return &v.Serviceprincipalkey
	}).(pulumi.StringPtrOutput)
}

func (o AzureRMCredentialsPtrOutput) ServiceprincipalkeyHash() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureRMCredentials) *string {
		if v == nil {
			return nil
		}
		return v.ServiceprincipalkeyHash
	}).(pulumi.StringPtrOutput)
}

type GitHubAuthOauth struct {
	OauthConfigurationId string `pulumi:"oauthConfigurationId"`
}

// GitHubAuthOauthInput is an input type that accepts GitHubAuthOauthArgs and GitHubAuthOauthOutput values.
// You can construct a concrete instance of `GitHubAuthOauthInput` via:
//
//          GitHubAuthOauthArgs{...}
type GitHubAuthOauthInput interface {
	pulumi.Input

	ToGitHubAuthOauthOutput() GitHubAuthOauthOutput
	ToGitHubAuthOauthOutputWithContext(context.Context) GitHubAuthOauthOutput
}

type GitHubAuthOauthArgs struct {
	OauthConfigurationId pulumi.StringInput `pulumi:"oauthConfigurationId"`
}

func (GitHubAuthOauthArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GitHubAuthOauth)(nil)).Elem()
}

func (i GitHubAuthOauthArgs) ToGitHubAuthOauthOutput() GitHubAuthOauthOutput {
	return i.ToGitHubAuthOauthOutputWithContext(context.Background())
}

func (i GitHubAuthOauthArgs) ToGitHubAuthOauthOutputWithContext(ctx context.Context) GitHubAuthOauthOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GitHubAuthOauthOutput)
}

func (i GitHubAuthOauthArgs) ToGitHubAuthOauthPtrOutput() GitHubAuthOauthPtrOutput {
	return i.ToGitHubAuthOauthPtrOutputWithContext(context.Background())
}

func (i GitHubAuthOauthArgs) ToGitHubAuthOauthPtrOutputWithContext(ctx context.Context) GitHubAuthOauthPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GitHubAuthOauthOutput).ToGitHubAuthOauthPtrOutputWithContext(ctx)
}

// GitHubAuthOauthPtrInput is an input type that accepts GitHubAuthOauthArgs, GitHubAuthOauthPtr and GitHubAuthOauthPtrOutput values.
// You can construct a concrete instance of `GitHubAuthOauthPtrInput` via:
//
//          GitHubAuthOauthArgs{...}
//
//  or:
//
//          nil
type GitHubAuthOauthPtrInput interface {
	pulumi.Input

	ToGitHubAuthOauthPtrOutput() GitHubAuthOauthPtrOutput
	ToGitHubAuthOauthPtrOutputWithContext(context.Context) GitHubAuthOauthPtrOutput
}

type gitHubAuthOauthPtrType GitHubAuthOauthArgs

func GitHubAuthOauthPtr(v *GitHubAuthOauthArgs) GitHubAuthOauthPtrInput {
	return (*gitHubAuthOauthPtrType)(v)
}

func (*gitHubAuthOauthPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GitHubAuthOauth)(nil)).Elem()
}

func (i *gitHubAuthOauthPtrType) ToGitHubAuthOauthPtrOutput() GitHubAuthOauthPtrOutput {
	return i.ToGitHubAuthOauthPtrOutputWithContext(context.Background())
}

func (i *gitHubAuthOauthPtrType) ToGitHubAuthOauthPtrOutputWithContext(ctx context.Context) GitHubAuthOauthPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GitHubAuthOauthPtrOutput)
}

type GitHubAuthOauthOutput struct{ *pulumi.OutputState }

func (GitHubAuthOauthOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GitHubAuthOauth)(nil)).Elem()
}

func (o GitHubAuthOauthOutput) ToGitHubAuthOauthOutput() GitHubAuthOauthOutput {
	return o
}

func (o GitHubAuthOauthOutput) ToGitHubAuthOauthOutputWithContext(ctx context.Context) GitHubAuthOauthOutput {
	return o
}

func (o GitHubAuthOauthOutput) ToGitHubAuthOauthPtrOutput() GitHubAuthOauthPtrOutput {
	return o.ToGitHubAuthOauthPtrOutputWithContext(context.Background())
}

func (o GitHubAuthOauthOutput) ToGitHubAuthOauthPtrOutputWithContext(ctx context.Context) GitHubAuthOauthPtrOutput {
	return o.ApplyT(func(v GitHubAuthOauth) *GitHubAuthOauth {
		return &v
	}).(GitHubAuthOauthPtrOutput)
}
func (o GitHubAuthOauthOutput) OauthConfigurationId() pulumi.StringOutput {
	return o.ApplyT(func(v GitHubAuthOauth) string { return v.OauthConfigurationId }).(pulumi.StringOutput)
}

type GitHubAuthOauthPtrOutput struct{ *pulumi.OutputState }

func (GitHubAuthOauthPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GitHubAuthOauth)(nil)).Elem()
}

func (o GitHubAuthOauthPtrOutput) ToGitHubAuthOauthPtrOutput() GitHubAuthOauthPtrOutput {
	return o
}

func (o GitHubAuthOauthPtrOutput) ToGitHubAuthOauthPtrOutputWithContext(ctx context.Context) GitHubAuthOauthPtrOutput {
	return o
}

func (o GitHubAuthOauthPtrOutput) Elem() GitHubAuthOauthOutput {
	return o.ApplyT(func(v *GitHubAuthOauth) GitHubAuthOauth { return *v }).(GitHubAuthOauthOutput)
}

func (o GitHubAuthOauthPtrOutput) OauthConfigurationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GitHubAuthOauth) *string {
		if v == nil {
			return nil
		}
		return &v.OauthConfigurationId
	}).(pulumi.StringPtrOutput)
}

type GitHubAuthPersonal struct {
	// The Personal Access Token for Github.
	PersonalAccessToken     string  `pulumi:"personalAccessToken"`
	PersonalAccessTokenHash *string `pulumi:"personalAccessTokenHash"`
}

// GitHubAuthPersonalInput is an input type that accepts GitHubAuthPersonalArgs and GitHubAuthPersonalOutput values.
// You can construct a concrete instance of `GitHubAuthPersonalInput` via:
//
//          GitHubAuthPersonalArgs{...}
type GitHubAuthPersonalInput interface {
	pulumi.Input

	ToGitHubAuthPersonalOutput() GitHubAuthPersonalOutput
	ToGitHubAuthPersonalOutputWithContext(context.Context) GitHubAuthPersonalOutput
}

type GitHubAuthPersonalArgs struct {
	// The Personal Access Token for Github.
	PersonalAccessToken     pulumi.StringInput    `pulumi:"personalAccessToken"`
	PersonalAccessTokenHash pulumi.StringPtrInput `pulumi:"personalAccessTokenHash"`
}

func (GitHubAuthPersonalArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GitHubAuthPersonal)(nil)).Elem()
}

func (i GitHubAuthPersonalArgs) ToGitHubAuthPersonalOutput() GitHubAuthPersonalOutput {
	return i.ToGitHubAuthPersonalOutputWithContext(context.Background())
}

func (i GitHubAuthPersonalArgs) ToGitHubAuthPersonalOutputWithContext(ctx context.Context) GitHubAuthPersonalOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GitHubAuthPersonalOutput)
}

func (i GitHubAuthPersonalArgs) ToGitHubAuthPersonalPtrOutput() GitHubAuthPersonalPtrOutput {
	return i.ToGitHubAuthPersonalPtrOutputWithContext(context.Background())
}

func (i GitHubAuthPersonalArgs) ToGitHubAuthPersonalPtrOutputWithContext(ctx context.Context) GitHubAuthPersonalPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GitHubAuthPersonalOutput).ToGitHubAuthPersonalPtrOutputWithContext(ctx)
}

// GitHubAuthPersonalPtrInput is an input type that accepts GitHubAuthPersonalArgs, GitHubAuthPersonalPtr and GitHubAuthPersonalPtrOutput values.
// You can construct a concrete instance of `GitHubAuthPersonalPtrInput` via:
//
//          GitHubAuthPersonalArgs{...}
//
//  or:
//
//          nil
type GitHubAuthPersonalPtrInput interface {
	pulumi.Input

	ToGitHubAuthPersonalPtrOutput() GitHubAuthPersonalPtrOutput
	ToGitHubAuthPersonalPtrOutputWithContext(context.Context) GitHubAuthPersonalPtrOutput
}

type gitHubAuthPersonalPtrType GitHubAuthPersonalArgs

func GitHubAuthPersonalPtr(v *GitHubAuthPersonalArgs) GitHubAuthPersonalPtrInput {
	return (*gitHubAuthPersonalPtrType)(v)
}

func (*gitHubAuthPersonalPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GitHubAuthPersonal)(nil)).Elem()
}

func (i *gitHubAuthPersonalPtrType) ToGitHubAuthPersonalPtrOutput() GitHubAuthPersonalPtrOutput {
	return i.ToGitHubAuthPersonalPtrOutputWithContext(context.Background())
}

func (i *gitHubAuthPersonalPtrType) ToGitHubAuthPersonalPtrOutputWithContext(ctx context.Context) GitHubAuthPersonalPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GitHubAuthPersonalPtrOutput)
}

type GitHubAuthPersonalOutput struct{ *pulumi.OutputState }

func (GitHubAuthPersonalOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GitHubAuthPersonal)(nil)).Elem()
}

func (o GitHubAuthPersonalOutput) ToGitHubAuthPersonalOutput() GitHubAuthPersonalOutput {
	return o
}

func (o GitHubAuthPersonalOutput) ToGitHubAuthPersonalOutputWithContext(ctx context.Context) GitHubAuthPersonalOutput {
	return o
}

func (o GitHubAuthPersonalOutput) ToGitHubAuthPersonalPtrOutput() GitHubAuthPersonalPtrOutput {
	return o.ToGitHubAuthPersonalPtrOutputWithContext(context.Background())
}

func (o GitHubAuthPersonalOutput) ToGitHubAuthPersonalPtrOutputWithContext(ctx context.Context) GitHubAuthPersonalPtrOutput {
	return o.ApplyT(func(v GitHubAuthPersonal) *GitHubAuthPersonal {
		return &v
	}).(GitHubAuthPersonalPtrOutput)
}

// The Personal Access Token for Github.
func (o GitHubAuthPersonalOutput) PersonalAccessToken() pulumi.StringOutput {
	return o.ApplyT(func(v GitHubAuthPersonal) string { return v.PersonalAccessToken }).(pulumi.StringOutput)
}

func (o GitHubAuthPersonalOutput) PersonalAccessTokenHash() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GitHubAuthPersonal) *string { return v.PersonalAccessTokenHash }).(pulumi.StringPtrOutput)
}

type GitHubAuthPersonalPtrOutput struct{ *pulumi.OutputState }

func (GitHubAuthPersonalPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GitHubAuthPersonal)(nil)).Elem()
}

func (o GitHubAuthPersonalPtrOutput) ToGitHubAuthPersonalPtrOutput() GitHubAuthPersonalPtrOutput {
	return o
}

func (o GitHubAuthPersonalPtrOutput) ToGitHubAuthPersonalPtrOutputWithContext(ctx context.Context) GitHubAuthPersonalPtrOutput {
	return o
}

func (o GitHubAuthPersonalPtrOutput) Elem() GitHubAuthPersonalOutput {
	return o.ApplyT(func(v *GitHubAuthPersonal) GitHubAuthPersonal { return *v }).(GitHubAuthPersonalOutput)
}

// The Personal Access Token for Github.
func (o GitHubAuthPersonalPtrOutput) PersonalAccessToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GitHubAuthPersonal) *string {
		if v == nil {
			return nil
		}
		return &v.PersonalAccessToken
	}).(pulumi.StringPtrOutput)
}

func (o GitHubAuthPersonalPtrOutput) PersonalAccessTokenHash() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GitHubAuthPersonal) *string {
		if v == nil {
			return nil
		}
		return v.PersonalAccessTokenHash
	}).(pulumi.StringPtrOutput)
}

type KubernetesAzureSubscription struct {
	// Azure environment refers to whether the public cloud offering or domestic (government) clouds are being used. Currently, only the public cloud is supported. The value must be AzureCloud. This is also the default-value.
	AzureEnvironment *string `pulumi:"azureEnvironment"`
	// Set this option to allow use cluster admin credentials.
	ClusterAdmin *bool `pulumi:"clusterAdmin"`
	// The name of the Kubernetes cluster.
	ClusterName string `pulumi:"clusterName"`
	// The Kubernetes namespace. Default value is "default".
	Namespace *string `pulumi:"namespace"`
	// The resource group name, to which the Kubernetes cluster is deployed.
	ResourcegroupId string `pulumi:"resourcegroupId"`
	// The id of the Azure subscription.
	SubscriptionId string `pulumi:"subscriptionId"`
	// The name of the Azure subscription.
	SubscriptionName string `pulumi:"subscriptionName"`
	// The id of the tenant used by the subscription.
	TenantId string `pulumi:"tenantId"`
}

// KubernetesAzureSubscriptionInput is an input type that accepts KubernetesAzureSubscriptionArgs and KubernetesAzureSubscriptionOutput values.
// You can construct a concrete instance of `KubernetesAzureSubscriptionInput` via:
//
//          KubernetesAzureSubscriptionArgs{...}
type KubernetesAzureSubscriptionInput interface {
	pulumi.Input

	ToKubernetesAzureSubscriptionOutput() KubernetesAzureSubscriptionOutput
	ToKubernetesAzureSubscriptionOutputWithContext(context.Context) KubernetesAzureSubscriptionOutput
}

type KubernetesAzureSubscriptionArgs struct {
	// Azure environment refers to whether the public cloud offering or domestic (government) clouds are being used. Currently, only the public cloud is supported. The value must be AzureCloud. This is also the default-value.
	AzureEnvironment pulumi.StringPtrInput `pulumi:"azureEnvironment"`
	// Set this option to allow use cluster admin credentials.
	ClusterAdmin pulumi.BoolPtrInput `pulumi:"clusterAdmin"`
	// The name of the Kubernetes cluster.
	ClusterName pulumi.StringInput `pulumi:"clusterName"`
	// The Kubernetes namespace. Default value is "default".
	Namespace pulumi.StringPtrInput `pulumi:"namespace"`
	// The resource group name, to which the Kubernetes cluster is deployed.
	ResourcegroupId pulumi.StringInput `pulumi:"resourcegroupId"`
	// The id of the Azure subscription.
	SubscriptionId pulumi.StringInput `pulumi:"subscriptionId"`
	// The name of the Azure subscription.
	SubscriptionName pulumi.StringInput `pulumi:"subscriptionName"`
	// The id of the tenant used by the subscription.
	TenantId pulumi.StringInput `pulumi:"tenantId"`
}

func (KubernetesAzureSubscriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KubernetesAzureSubscription)(nil)).Elem()
}

func (i KubernetesAzureSubscriptionArgs) ToKubernetesAzureSubscriptionOutput() KubernetesAzureSubscriptionOutput {
	return i.ToKubernetesAzureSubscriptionOutputWithContext(context.Background())
}

func (i KubernetesAzureSubscriptionArgs) ToKubernetesAzureSubscriptionOutputWithContext(ctx context.Context) KubernetesAzureSubscriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesAzureSubscriptionOutput)
}

// KubernetesAzureSubscriptionArrayInput is an input type that accepts KubernetesAzureSubscriptionArray and KubernetesAzureSubscriptionArrayOutput values.
// You can construct a concrete instance of `KubernetesAzureSubscriptionArrayInput` via:
//
//          KubernetesAzureSubscriptionArray{ KubernetesAzureSubscriptionArgs{...} }
type KubernetesAzureSubscriptionArrayInput interface {
	pulumi.Input

	ToKubernetesAzureSubscriptionArrayOutput() KubernetesAzureSubscriptionArrayOutput
	ToKubernetesAzureSubscriptionArrayOutputWithContext(context.Context) KubernetesAzureSubscriptionArrayOutput
}

type KubernetesAzureSubscriptionArray []KubernetesAzureSubscriptionInput

func (KubernetesAzureSubscriptionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]KubernetesAzureSubscription)(nil)).Elem()
}

func (i KubernetesAzureSubscriptionArray) ToKubernetesAzureSubscriptionArrayOutput() KubernetesAzureSubscriptionArrayOutput {
	return i.ToKubernetesAzureSubscriptionArrayOutputWithContext(context.Background())
}

func (i KubernetesAzureSubscriptionArray) ToKubernetesAzureSubscriptionArrayOutputWithContext(ctx context.Context) KubernetesAzureSubscriptionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesAzureSubscriptionArrayOutput)
}

type KubernetesAzureSubscriptionOutput struct{ *pulumi.OutputState }

func (KubernetesAzureSubscriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KubernetesAzureSubscription)(nil)).Elem()
}

func (o KubernetesAzureSubscriptionOutput) ToKubernetesAzureSubscriptionOutput() KubernetesAzureSubscriptionOutput {
	return o
}

func (o KubernetesAzureSubscriptionOutput) ToKubernetesAzureSubscriptionOutputWithContext(ctx context.Context) KubernetesAzureSubscriptionOutput {
	return o
}

// Azure environment refers to whether the public cloud offering or domestic (government) clouds are being used. Currently, only the public cloud is supported. The value must be AzureCloud. This is also the default-value.
func (o KubernetesAzureSubscriptionOutput) AzureEnvironment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KubernetesAzureSubscription) *string { return v.AzureEnvironment }).(pulumi.StringPtrOutput)
}

// Set this option to allow use cluster admin credentials.
func (o KubernetesAzureSubscriptionOutput) ClusterAdmin() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v KubernetesAzureSubscription) *bool { return v.ClusterAdmin }).(pulumi.BoolPtrOutput)
}

// The name of the Kubernetes cluster.
func (o KubernetesAzureSubscriptionOutput) ClusterName() pulumi.StringOutput {
	return o.ApplyT(func(v KubernetesAzureSubscription) string { return v.ClusterName }).(pulumi.StringOutput)
}

// The Kubernetes namespace. Default value is "default".
func (o KubernetesAzureSubscriptionOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KubernetesAzureSubscription) *string { return v.Namespace }).(pulumi.StringPtrOutput)
}

// The resource group name, to which the Kubernetes cluster is deployed.
func (o KubernetesAzureSubscriptionOutput) ResourcegroupId() pulumi.StringOutput {
	return o.ApplyT(func(v KubernetesAzureSubscription) string { return v.ResourcegroupId }).(pulumi.StringOutput)
}

// The id of the Azure subscription.
func (o KubernetesAzureSubscriptionOutput) SubscriptionId() pulumi.StringOutput {
	return o.ApplyT(func(v KubernetesAzureSubscription) string { return v.SubscriptionId }).(pulumi.StringOutput)
}

// The name of the Azure subscription.
func (o KubernetesAzureSubscriptionOutput) SubscriptionName() pulumi.StringOutput {
	return o.ApplyT(func(v KubernetesAzureSubscription) string { return v.SubscriptionName }).(pulumi.StringOutput)
}

// The id of the tenant used by the subscription.
func (o KubernetesAzureSubscriptionOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v KubernetesAzureSubscription) string { return v.TenantId }).(pulumi.StringOutput)
}

type KubernetesAzureSubscriptionArrayOutput struct{ *pulumi.OutputState }

func (KubernetesAzureSubscriptionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]KubernetesAzureSubscription)(nil)).Elem()
}

func (o KubernetesAzureSubscriptionArrayOutput) ToKubernetesAzureSubscriptionArrayOutput() KubernetesAzureSubscriptionArrayOutput {
	return o
}

func (o KubernetesAzureSubscriptionArrayOutput) ToKubernetesAzureSubscriptionArrayOutputWithContext(ctx context.Context) KubernetesAzureSubscriptionArrayOutput {
	return o
}

func (o KubernetesAzureSubscriptionArrayOutput) Index(i pulumi.IntInput) KubernetesAzureSubscriptionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) KubernetesAzureSubscription {
		return vs[0].([]KubernetesAzureSubscription)[vs[1].(int)]
	}).(KubernetesAzureSubscriptionOutput)
}

type KubernetesKubeconfig struct {
	// Set this option to allow clients to accept a self-signed certificate.
	AcceptUntrustedCerts *bool `pulumi:"acceptUntrustedCerts"`
	// Context within the kubeconfig file that is to be used for identifying the cluster. Default value is the current-context set in kubeconfig.
	ClusterContext *string `pulumi:"clusterContext"`
	// The content of the kubeconfig in yaml notation to be used to communicate with the API-Server of Kubernetes.
	KubeConfig     string  `pulumi:"kubeConfig"`
	KubeConfigHash *string `pulumi:"kubeConfigHash"`
}

// KubernetesKubeconfigInput is an input type that accepts KubernetesKubeconfigArgs and KubernetesKubeconfigOutput values.
// You can construct a concrete instance of `KubernetesKubeconfigInput` via:
//
//          KubernetesKubeconfigArgs{...}
type KubernetesKubeconfigInput interface {
	pulumi.Input

	ToKubernetesKubeconfigOutput() KubernetesKubeconfigOutput
	ToKubernetesKubeconfigOutputWithContext(context.Context) KubernetesKubeconfigOutput
}

type KubernetesKubeconfigArgs struct {
	// Set this option to allow clients to accept a self-signed certificate.
	AcceptUntrustedCerts pulumi.BoolPtrInput `pulumi:"acceptUntrustedCerts"`
	// Context within the kubeconfig file that is to be used for identifying the cluster. Default value is the current-context set in kubeconfig.
	ClusterContext pulumi.StringPtrInput `pulumi:"clusterContext"`
	// The content of the kubeconfig in yaml notation to be used to communicate with the API-Server of Kubernetes.
	KubeConfig     pulumi.StringInput    `pulumi:"kubeConfig"`
	KubeConfigHash pulumi.StringPtrInput `pulumi:"kubeConfigHash"`
}

func (KubernetesKubeconfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KubernetesKubeconfig)(nil)).Elem()
}

func (i KubernetesKubeconfigArgs) ToKubernetesKubeconfigOutput() KubernetesKubeconfigOutput {
	return i.ToKubernetesKubeconfigOutputWithContext(context.Background())
}

func (i KubernetesKubeconfigArgs) ToKubernetesKubeconfigOutputWithContext(ctx context.Context) KubernetesKubeconfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesKubeconfigOutput)
}

// KubernetesKubeconfigArrayInput is an input type that accepts KubernetesKubeconfigArray and KubernetesKubeconfigArrayOutput values.
// You can construct a concrete instance of `KubernetesKubeconfigArrayInput` via:
//
//          KubernetesKubeconfigArray{ KubernetesKubeconfigArgs{...} }
type KubernetesKubeconfigArrayInput interface {
	pulumi.Input

	ToKubernetesKubeconfigArrayOutput() KubernetesKubeconfigArrayOutput
	ToKubernetesKubeconfigArrayOutputWithContext(context.Context) KubernetesKubeconfigArrayOutput
}

type KubernetesKubeconfigArray []KubernetesKubeconfigInput

func (KubernetesKubeconfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]KubernetesKubeconfig)(nil)).Elem()
}

func (i KubernetesKubeconfigArray) ToKubernetesKubeconfigArrayOutput() KubernetesKubeconfigArrayOutput {
	return i.ToKubernetesKubeconfigArrayOutputWithContext(context.Background())
}

func (i KubernetesKubeconfigArray) ToKubernetesKubeconfigArrayOutputWithContext(ctx context.Context) KubernetesKubeconfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesKubeconfigArrayOutput)
}

type KubernetesKubeconfigOutput struct{ *pulumi.OutputState }

func (KubernetesKubeconfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KubernetesKubeconfig)(nil)).Elem()
}

func (o KubernetesKubeconfigOutput) ToKubernetesKubeconfigOutput() KubernetesKubeconfigOutput {
	return o
}

func (o KubernetesKubeconfigOutput) ToKubernetesKubeconfigOutputWithContext(ctx context.Context) KubernetesKubeconfigOutput {
	return o
}

// Set this option to allow clients to accept a self-signed certificate.
func (o KubernetesKubeconfigOutput) AcceptUntrustedCerts() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v KubernetesKubeconfig) *bool { return v.AcceptUntrustedCerts }).(pulumi.BoolPtrOutput)
}

// Context within the kubeconfig file that is to be used for identifying the cluster. Default value is the current-context set in kubeconfig.
func (o KubernetesKubeconfigOutput) ClusterContext() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KubernetesKubeconfig) *string { return v.ClusterContext }).(pulumi.StringPtrOutput)
}

// The content of the kubeconfig in yaml notation to be used to communicate with the API-Server of Kubernetes.
func (o KubernetesKubeconfigOutput) KubeConfig() pulumi.StringOutput {
	return o.ApplyT(func(v KubernetesKubeconfig) string { return v.KubeConfig }).(pulumi.StringOutput)
}

func (o KubernetesKubeconfigOutput) KubeConfigHash() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KubernetesKubeconfig) *string { return v.KubeConfigHash }).(pulumi.StringPtrOutput)
}

type KubernetesKubeconfigArrayOutput struct{ *pulumi.OutputState }

func (KubernetesKubeconfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]KubernetesKubeconfig)(nil)).Elem()
}

func (o KubernetesKubeconfigArrayOutput) ToKubernetesKubeconfigArrayOutput() KubernetesKubeconfigArrayOutput {
	return o
}

func (o KubernetesKubeconfigArrayOutput) ToKubernetesKubeconfigArrayOutputWithContext(ctx context.Context) KubernetesKubeconfigArrayOutput {
	return o
}

func (o KubernetesKubeconfigArrayOutput) Index(i pulumi.IntInput) KubernetesKubeconfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) KubernetesKubeconfig {
		return vs[0].([]KubernetesKubeconfig)[vs[1].(int)]
	}).(KubernetesKubeconfigOutput)
}

type KubernetesServiceAccount struct {
	// The certificate from a Kubernetes secret object.
	CaCert     string  `pulumi:"caCert"`
	CaCertHash *string `pulumi:"caCertHash"`
	// The token from a Kubernetes secret object.
	Token     string  `pulumi:"token"`
	TokenHash *string `pulumi:"tokenHash"`
}

// KubernetesServiceAccountInput is an input type that accepts KubernetesServiceAccountArgs and KubernetesServiceAccountOutput values.
// You can construct a concrete instance of `KubernetesServiceAccountInput` via:
//
//          KubernetesServiceAccountArgs{...}
type KubernetesServiceAccountInput interface {
	pulumi.Input

	ToKubernetesServiceAccountOutput() KubernetesServiceAccountOutput
	ToKubernetesServiceAccountOutputWithContext(context.Context) KubernetesServiceAccountOutput
}

type KubernetesServiceAccountArgs struct {
	// The certificate from a Kubernetes secret object.
	CaCert     pulumi.StringInput    `pulumi:"caCert"`
	CaCertHash pulumi.StringPtrInput `pulumi:"caCertHash"`
	// The token from a Kubernetes secret object.
	Token     pulumi.StringInput    `pulumi:"token"`
	TokenHash pulumi.StringPtrInput `pulumi:"tokenHash"`
}

func (KubernetesServiceAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KubernetesServiceAccount)(nil)).Elem()
}

func (i KubernetesServiceAccountArgs) ToKubernetesServiceAccountOutput() KubernetesServiceAccountOutput {
	return i.ToKubernetesServiceAccountOutputWithContext(context.Background())
}

func (i KubernetesServiceAccountArgs) ToKubernetesServiceAccountOutputWithContext(ctx context.Context) KubernetesServiceAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesServiceAccountOutput)
}

// KubernetesServiceAccountArrayInput is an input type that accepts KubernetesServiceAccountArray and KubernetesServiceAccountArrayOutput values.
// You can construct a concrete instance of `KubernetesServiceAccountArrayInput` via:
//
//          KubernetesServiceAccountArray{ KubernetesServiceAccountArgs{...} }
type KubernetesServiceAccountArrayInput interface {
	pulumi.Input

	ToKubernetesServiceAccountArrayOutput() KubernetesServiceAccountArrayOutput
	ToKubernetesServiceAccountArrayOutputWithContext(context.Context) KubernetesServiceAccountArrayOutput
}

type KubernetesServiceAccountArray []KubernetesServiceAccountInput

func (KubernetesServiceAccountArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]KubernetesServiceAccount)(nil)).Elem()
}

func (i KubernetesServiceAccountArray) ToKubernetesServiceAccountArrayOutput() KubernetesServiceAccountArrayOutput {
	return i.ToKubernetesServiceAccountArrayOutputWithContext(context.Background())
}

func (i KubernetesServiceAccountArray) ToKubernetesServiceAccountArrayOutputWithContext(ctx context.Context) KubernetesServiceAccountArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesServiceAccountArrayOutput)
}

type KubernetesServiceAccountOutput struct{ *pulumi.OutputState }

func (KubernetesServiceAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KubernetesServiceAccount)(nil)).Elem()
}

func (o KubernetesServiceAccountOutput) ToKubernetesServiceAccountOutput() KubernetesServiceAccountOutput {
	return o
}

func (o KubernetesServiceAccountOutput) ToKubernetesServiceAccountOutputWithContext(ctx context.Context) KubernetesServiceAccountOutput {
	return o
}

// The certificate from a Kubernetes secret object.
func (o KubernetesServiceAccountOutput) CaCert() pulumi.StringOutput {
	return o.ApplyT(func(v KubernetesServiceAccount) string { return v.CaCert }).(pulumi.StringOutput)
}

func (o KubernetesServiceAccountOutput) CaCertHash() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KubernetesServiceAccount) *string { return v.CaCertHash }).(pulumi.StringPtrOutput)
}

// The token from a Kubernetes secret object.
func (o KubernetesServiceAccountOutput) Token() pulumi.StringOutput {
	return o.ApplyT(func(v KubernetesServiceAccount) string { return v.Token }).(pulumi.StringOutput)
}

func (o KubernetesServiceAccountOutput) TokenHash() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KubernetesServiceAccount) *string { return v.TokenHash }).(pulumi.StringPtrOutput)
}

type KubernetesServiceAccountArrayOutput struct{ *pulumi.OutputState }

func (KubernetesServiceAccountArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]KubernetesServiceAccount)(nil)).Elem()
}

func (o KubernetesServiceAccountArrayOutput) ToKubernetesServiceAccountArrayOutput() KubernetesServiceAccountArrayOutput {
	return o
}

func (o KubernetesServiceAccountArrayOutput) ToKubernetesServiceAccountArrayOutputWithContext(ctx context.Context) KubernetesServiceAccountArrayOutput {
	return o
}

func (o KubernetesServiceAccountArrayOutput) Index(i pulumi.IntInput) KubernetesServiceAccountOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) KubernetesServiceAccount {
		return vs[0].([]KubernetesServiceAccount)[vs[1].(int)]
	}).(KubernetesServiceAccountOutput)
}

func init() {
	pulumi.RegisterOutputType(AzureRMCredentialsOutput{})
	pulumi.RegisterOutputType(AzureRMCredentialsPtrOutput{})
	pulumi.RegisterOutputType(GitHubAuthOauthOutput{})
	pulumi.RegisterOutputType(GitHubAuthOauthPtrOutput{})
	pulumi.RegisterOutputType(GitHubAuthPersonalOutput{})
	pulumi.RegisterOutputType(GitHubAuthPersonalPtrOutput{})
	pulumi.RegisterOutputType(KubernetesAzureSubscriptionOutput{})
	pulumi.RegisterOutputType(KubernetesAzureSubscriptionArrayOutput{})
	pulumi.RegisterOutputType(KubernetesKubeconfigOutput{})
	pulumi.RegisterOutputType(KubernetesKubeconfigArrayOutput{})
	pulumi.RegisterOutputType(KubernetesServiceAccountOutput{})
	pulumi.RegisterOutputType(KubernetesServiceAccountArrayOutput{})
}
