// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Immutable;

namespace Pulumi.AzureDevOps
{
    public static class Config
    {
        [global::System.Diagnostics.CodeAnalysis.SuppressMessage("Microsoft.Design", "IDE1006", Justification = 
        "Double underscore prefix used to avoid conflicts with variable names.")]
        private sealed class __Value<T>
        {
            private readonly Func<T> _getter;
            private T _value = default!;
            private bool _set;

            public __Value(Func<T> getter)
            {
                _getter = getter;
            }

            public T Get() => _set ? _value : _getter();

            public void Set(T value)
            {
                _value = value;
                _set = true;
            }
        }

        private static readonly global::Pulumi.Config __config = new global::Pulumi.Config("azuredevops");

        private static readonly __Value<string?> _orgServiceUrl = new __Value<string?>(() => __config.Get("orgServiceUrl") ?? Utilities.GetEnv("AZDO_ORG_SERVICE_URL"));
        /// <summary>
        /// The url of the Azure DevOps instance which should be used.
        /// </summary>
        public static string? OrgServiceUrl
        {
            get => _orgServiceUrl.Get();
            set => _orgServiceUrl.Set(value);
        }

        private static readonly __Value<string?> _personalAccessToken = new __Value<string?>(() => __config.Get("personalAccessToken"));
        /// <summary>
        /// The personal access token which should be used.
        /// </summary>
        public static string? PersonalAccessToken
        {
            get => _personalAccessToken.Get();
            set => _personalAccessToken.Set(value);
        }

    }
}
