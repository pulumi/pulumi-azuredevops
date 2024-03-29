// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops;

import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;

public final class Config {

    private static final com.pulumi.Config config = com.pulumi.Config.of("azuredevops");
/**
 * Base64 encoded certificate to use to authenticate to the service principal.
 * 
 */
    public Optional<String> clientCertificate() {
        return Codegen.stringProp("clientCertificate").config(config).get();
    }
/**
 * Password for a client certificate password.
 * 
 */
    public Optional<String> clientCertificatePassword() {
        return Codegen.stringProp("clientCertificatePassword").config(config).get();
    }
/**
 * Path to a certificate to use to authenticate to the service principal.
 * 
 */
    public Optional<String> clientCertificatePath() {
        return Codegen.stringProp("clientCertificatePath").config(config).get();
    }
/**
 * The service principal client or managed service principal id which should be used.
 * 
 */
    public Optional<String> clientId() {
        return Codegen.stringProp("clientId").config(config).get();
    }
    public Optional<String> clientIdApply() {
        return Codegen.stringProp("clientIdApply").config(config).get();
    }
    public Optional<String> clientIdPlan() {
        return Codegen.stringProp("clientIdPlan").config(config).get();
    }
/**
 * Client secret for authenticating to a service principal.
 * 
 */
    public Optional<String> clientSecret() {
        return Codegen.stringProp("clientSecret").config(config).get();
    }
/**
 * Path to a file containing a client secret for authenticating to a service principal.
 * 
 */
    public Optional<String> clientSecretPath() {
        return Codegen.stringProp("clientSecretPath").config(config).get();
    }
/**
 * Set the audience when requesting OIDC tokens.
 * 
 */
    public Optional<String> oidcAudience() {
        return Codegen.stringProp("oidcAudience").config(config).get();
    }
/**
 * The bearer token for the request to the OIDC provider. For use when authenticating as a Service Principal using OpenID
 * Connect.
 * 
 */
    public Optional<String> oidcRequestToken() {
        return Codegen.stringProp("oidcRequestToken").config(config).get();
    }
/**
 * The URL for the OIDC provider from which to request an ID token. For use when authenticating as a Service Principal
 * using OpenID Connect.
 * 
 */
    public Optional<String> oidcRequestUrl() {
        return Codegen.stringProp("oidcRequestUrl").config(config).get();
    }
    public Optional<String> oidcTfcTag() {
        return Codegen.stringProp("oidcTfcTag").config(config).get();
    }
/**
 * OIDC token to authenticate as a service principal.
 * 
 */
    public Optional<String> oidcToken() {
        return Codegen.stringProp("oidcToken").config(config).get();
    }
/**
 * OIDC token from file to authenticate as a service principal.
 * 
 */
    public Optional<String> oidcTokenFilePath() {
        return Codegen.stringProp("oidcTokenFilePath").config(config).get();
    }
/**
 * The url of the Azure DevOps instance which should be used.
 * 
 */
    public Optional<String> orgServiceUrl() {
        return Codegen.stringProp("orgServiceUrl").config(config).env("AZDO_ORG_SERVICE_URL").get();
    }
/**
 * The personal access token which should be used.
 * 
 */
    public Optional<String> personalAccessToken() {
        return Codegen.stringProp("personalAccessToken").config(config).get();
    }
/**
 * The service principal tenant id which should be used.
 * 
 */
    public Optional<String> tenantId() {
        return Codegen.stringProp("tenantId").config(config).get();
    }
    public Optional<String> tenantIdApply() {
        return Codegen.stringProp("tenantIdApply").config(config).get();
    }
    public Optional<String> tenantIdPlan() {
        return Codegen.stringProp("tenantIdPlan").config(config).get();
    }
/**
 * Use an Azure Managed Service Identity.
 * 
 */
    public Optional<Boolean> useMsi() {
        return Codegen.booleanProp("useMsi").config(config).get();
    }
/**
 * Use an OIDC token to authenticate to a service principal.
 * 
 */
    public Optional<Boolean> useOidc() {
        return Codegen.booleanProp("useOidc").config(config).get();
    }
}
