// *** WARNING: this file was generated by pulumi-language-dotnet. ***
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

        private static readonly __Value<string?> _clientCertificate = new __Value<string?>(() => __config.Get("clientCertificate"));
        /// <summary>
        /// Base64 encoded certificate to use to authenticate to the service principal.
        /// </summary>
        public static string? ClientCertificate
        {
            get => _clientCertificate.Get();
            set => _clientCertificate.Set(value);
        }

        private static readonly __Value<string?> _clientCertificatePassword = new __Value<string?>(() => __config.Get("clientCertificatePassword"));
        /// <summary>
        /// Password for a client certificate password.
        /// </summary>
        public static string? ClientCertificatePassword
        {
            get => _clientCertificatePassword.Get();
            set => _clientCertificatePassword.Set(value);
        }

        private static readonly __Value<string?> _clientCertificatePath = new __Value<string?>(() => __config.Get("clientCertificatePath"));
        /// <summary>
        /// Path to a certificate to use to authenticate to the service principal.
        /// </summary>
        public static string? ClientCertificatePath
        {
            get => _clientCertificatePath.Get();
            set => _clientCertificatePath.Set(value);
        }

        private static readonly __Value<string?> _clientId = new __Value<string?>(() => __config.Get("clientId"));
        /// <summary>
        /// The service principal client or managed service principal id which should be used.
        /// </summary>
        public static string? ClientId
        {
            get => _clientId.Get();
            set => _clientId.Set(value);
        }

        private static readonly __Value<string?> _clientIdApply = new __Value<string?>(() => __config.Get("clientIdApply"));
        public static string? ClientIdApply
        {
            get => _clientIdApply.Get();
            set => _clientIdApply.Set(value);
        }

        private static readonly __Value<string?> _clientIdPlan = new __Value<string?>(() => __config.Get("clientIdPlan"));
        public static string? ClientIdPlan
        {
            get => _clientIdPlan.Get();
            set => _clientIdPlan.Set(value);
        }

        private static readonly __Value<string?> _clientSecret = new __Value<string?>(() => __config.Get("clientSecret"));
        /// <summary>
        /// Client secret for authenticating to a service principal.
        /// </summary>
        public static string? ClientSecret
        {
            get => _clientSecret.Get();
            set => _clientSecret.Set(value);
        }

        private static readonly __Value<string?> _clientSecretPath = new __Value<string?>(() => __config.Get("clientSecretPath"));
        /// <summary>
        /// Path to a file containing a client secret for authenticating to a service principal.
        /// </summary>
        public static string? ClientSecretPath
        {
            get => _clientSecretPath.Get();
            set => _clientSecretPath.Set(value);
        }

        private static readonly __Value<string?> _oidcAudience = new __Value<string?>(() => __config.Get("oidcAudience"));
        /// <summary>
        /// Set the audience when requesting OIDC tokens.
        /// </summary>
        public static string? OidcAudience
        {
            get => _oidcAudience.Get();
            set => _oidcAudience.Set(value);
        }

        private static readonly __Value<string?> _oidcRequestToken = new __Value<string?>(() => __config.Get("oidcRequestToken"));
        /// <summary>
        /// The bearer token for the request to the OIDC provider. For use when authenticating as a Service Principal using OpenID
        /// Connect.
        /// </summary>
        public static string? OidcRequestToken
        {
            get => _oidcRequestToken.Get();
            set => _oidcRequestToken.Set(value);
        }

        private static readonly __Value<string?> _oidcRequestUrl = new __Value<string?>(() => __config.Get("oidcRequestUrl"));
        /// <summary>
        /// The URL for the OIDC provider from which to request an ID token. For use when authenticating as a Service Principal
        /// using OpenID Connect.
        /// </summary>
        public static string? OidcRequestUrl
        {
            get => _oidcRequestUrl.Get();
            set => _oidcRequestUrl.Set(value);
        }

        private static readonly __Value<string?> _oidcTfcTag = new __Value<string?>(() => __config.Get("oidcTfcTag"));
        public static string? OidcTfcTag
        {
            get => _oidcTfcTag.Get();
            set => _oidcTfcTag.Set(value);
        }

        private static readonly __Value<string?> _oidcToken = new __Value<string?>(() => __config.Get("oidcToken"));
        /// <summary>
        /// OIDC token to authenticate as a service principal.
        /// </summary>
        public static string? OidcToken
        {
            get => _oidcToken.Get();
            set => _oidcToken.Set(value);
        }

        private static readonly __Value<string?> _oidcTokenFilePath = new __Value<string?>(() => __config.Get("oidcTokenFilePath"));
        /// <summary>
        /// OIDC token from file to authenticate as a service principal.
        /// </summary>
        public static string? OidcTokenFilePath
        {
            get => _oidcTokenFilePath.Get();
            set => _oidcTokenFilePath.Set(value);
        }

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

        private static readonly __Value<string?> _tenantId = new __Value<string?>(() => __config.Get("tenantId"));
        /// <summary>
        /// The service principal tenant id which should be used.
        /// </summary>
        public static string? TenantId
        {
            get => _tenantId.Get();
            set => _tenantId.Set(value);
        }

        private static readonly __Value<string?> _tenantIdApply = new __Value<string?>(() => __config.Get("tenantIdApply"));
        public static string? TenantIdApply
        {
            get => _tenantIdApply.Get();
            set => _tenantIdApply.Set(value);
        }

        private static readonly __Value<string?> _tenantIdPlan = new __Value<string?>(() => __config.Get("tenantIdPlan"));
        public static string? TenantIdPlan
        {
            get => _tenantIdPlan.Get();
            set => _tenantIdPlan.Set(value);
        }

        private static readonly __Value<bool?> _useMsi = new __Value<bool?>(() => __config.GetBoolean("useMsi"));
        /// <summary>
        /// Use an Azure Managed Service Identity.
        /// </summary>
        public static bool? UseMsi
        {
            get => _useMsi.Get();
            set => _useMsi.Set(value);
        }

        private static readonly __Value<bool?> _useOidc = new __Value<bool?>(() => __config.GetBoolean("useOidc"));
        /// <summary>
        /// Use an OIDC token to authenticate to a service principal.
        /// </summary>
        public static bool? UseOidc
        {
            get => _useOidc.Get();
            set => _useOidc.Set(value);
        }

    }
}
