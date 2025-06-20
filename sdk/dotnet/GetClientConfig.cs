// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureDevOps
{
    public static class GetClientConfig
    {
        /// <summary>
        /// Use this data source to access information about the Azure DevOps organization configured for the provider.
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using AzureDevOps = Pulumi.AzureDevOps;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = AzureDevOps.GetClientConfig.Invoke();
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["orgUrl"] = example.Apply(getClientConfigResult =&gt; getClientConfigResult.OrganizationUrl),
        ///     };
        /// });
        /// ```
        /// </summary>
        public static Task<GetClientConfigResult> InvokeAsync(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetClientConfigResult>("azuredevops:index/getClientConfig:getClientConfig", InvokeArgs.Empty, options.WithDefaults());

        /// <summary>
        /// Use this data source to access information about the Azure DevOps organization configured for the provider.
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using AzureDevOps = Pulumi.AzureDevOps;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = AzureDevOps.GetClientConfig.Invoke();
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["orgUrl"] = example.Apply(getClientConfigResult =&gt; getClientConfigResult.OrganizationUrl),
        ///     };
        /// });
        /// ```
        /// </summary>
        public static Output<GetClientConfigResult> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetClientConfigResult>("azuredevops:index/getClientConfig:getClientConfig", InvokeArgs.Empty, options.WithDefaults());

        /// <summary>
        /// Use this data source to access information about the Azure DevOps organization configured for the provider.
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using AzureDevOps = Pulumi.AzureDevOps;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = AzureDevOps.GetClientConfig.Invoke();
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["orgUrl"] = example.Apply(getClientConfigResult =&gt; getClientConfigResult.OrganizationUrl),
        ///     };
        /// });
        /// ```
        /// </summary>
        public static Output<GetClientConfigResult> Invoke(InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetClientConfigResult>("azuredevops:index/getClientConfig:getClientConfig", InvokeArgs.Empty, options.WithDefaults());
    }


    [OutputType]
    public sealed class GetClientConfigResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the organization.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The URL of the organization.
        /// </summary>
        public readonly string OrganizationUrl;
        /// <summary>
        /// The owner ID of the organization.
        /// </summary>
        public readonly string OwnerId;
        /// <summary>
        /// The status of the organization.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// The Tenant ID of the connected Azure Directory.
        /// </summary>
        public readonly string TenantId;

        [OutputConstructor]
        private GetClientConfigResult(
            string id,

            string name,

            string organizationUrl,

            string ownerId,

            string status,

            string tenantId)
        {
            Id = id;
            Name = name;
            OrganizationUrl = organizationUrl;
            OwnerId = ownerId;
            Status = status;
            TenantId = tenantId;
        }
    }
}
