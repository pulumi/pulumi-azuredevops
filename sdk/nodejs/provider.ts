// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The provider type for the azuredevops package. By default, resources use package-wide configuration
 * settings, however an explicit `Provider` instance may be created and passed during resource
 * construction to achieve fine-grained programmatic control over provider settings. See the
 * [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
 */
export class Provider extends pulumi.ProviderResource {
    /** @internal */
    public static readonly __pulumiType = 'azuredevops';

    /**
     * Returns true if the given object is an instance of Provider.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Provider {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === "pulumi:providers:" + Provider.__pulumiType;
    }

    /**
     * Base64 encoded certificate to use to authenticate to the service principal.
     */
    public readonly clientCertificate!: pulumi.Output<string | undefined>;
    /**
     * Password for a client certificate password.
     */
    public readonly clientCertificatePassword!: pulumi.Output<string | undefined>;
    /**
     * Path to a certificate to use to authenticate to the service principal.
     */
    public readonly clientCertificatePath!: pulumi.Output<string | undefined>;
    /**
     * The service principal client or managed service principal id which should be used.
     */
    public readonly clientId!: pulumi.Output<string | undefined>;
    /**
     * The service principal client id which should be used during an apply operation in Terraform Cloud.
     */
    public readonly clientIdApply!: pulumi.Output<string | undefined>;
    /**
     * The service principal client id which should be used during a plan operation in Terraform Cloud.
     */
    public readonly clientIdPlan!: pulumi.Output<string | undefined>;
    /**
     * Client secret for authenticating to a service principal.
     */
    public readonly clientSecret!: pulumi.Output<string | undefined>;
    /**
     * Path to a file containing a client secret for authenticating to a service principal.
     */
    public readonly clientSecretPath!: pulumi.Output<string | undefined>;
    /**
     * Set the audience when requesting OIDC tokens.
     */
    public readonly oidcAudience!: pulumi.Output<string | undefined>;
    /**
     * The bearer token for the request to the OIDC provider. For use when authenticating as a Service Principal using OpenID
     * Connect.
     */
    public readonly oidcRequestToken!: pulumi.Output<string | undefined>;
    /**
     * The URL for the OIDC provider from which to request an ID token. For use when authenticating as a Service Principal
     * using OpenID Connect.
     */
    public readonly oidcRequestUrl!: pulumi.Output<string | undefined>;
    /**
     * Terraform Cloud dynamic credential provider tag.
     */
    public readonly oidcTfcTag!: pulumi.Output<string | undefined>;
    /**
     * OIDC token to authenticate as a service principal.
     */
    public readonly oidcToken!: pulumi.Output<string | undefined>;
    /**
     * OIDC token from file to authenticate as a service principal.
     */
    public readonly oidcTokenFilePath!: pulumi.Output<string | undefined>;
    /**
     * The url of the Azure DevOps instance which should be used.
     */
    public readonly orgServiceUrl!: pulumi.Output<string | undefined>;
    /**
     * The personal access token which should be used.
     */
    public readonly personalAccessToken!: pulumi.Output<string | undefined>;
    /**
     * The service principal tenant id which should be used.
     */
    public readonly tenantId!: pulumi.Output<string | undefined>;
    /**
     * The service principal tenant id which should be used during an apply operation in Terraform Cloud..
     */
    public readonly tenantIdApply!: pulumi.Output<string | undefined>;
    /**
     * The service principal tenant id which should be used during a plan operation in Terraform Cloud.
     */
    public readonly tenantIdPlan!: pulumi.Output<string | undefined>;

    /**
     * Create a Provider resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ProviderArgs, opts?: pulumi.ResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        {
            resourceInputs["clientCertificate"] = args?.clientCertificate ? pulumi.secret(args.clientCertificate) : undefined;
            resourceInputs["clientCertificatePassword"] = args?.clientCertificatePassword ? pulumi.secret(args.clientCertificatePassword) : undefined;
            resourceInputs["clientCertificatePath"] = args ? args.clientCertificatePath : undefined;
            resourceInputs["clientId"] = args ? args.clientId : undefined;
            resourceInputs["clientIdApply"] = args ? args.clientIdApply : undefined;
            resourceInputs["clientIdPlan"] = args ? args.clientIdPlan : undefined;
            resourceInputs["clientSecret"] = args?.clientSecret ? pulumi.secret(args.clientSecret) : undefined;
            resourceInputs["clientSecretPath"] = args ? args.clientSecretPath : undefined;
            resourceInputs["oidcAudience"] = args ? args.oidcAudience : undefined;
            resourceInputs["oidcRequestToken"] = args ? args.oidcRequestToken : undefined;
            resourceInputs["oidcRequestUrl"] = args ? args.oidcRequestUrl : undefined;
            resourceInputs["oidcTfcTag"] = args ? args.oidcTfcTag : undefined;
            resourceInputs["oidcToken"] = args?.oidcToken ? pulumi.secret(args.oidcToken) : undefined;
            resourceInputs["oidcTokenFilePath"] = args ? args.oidcTokenFilePath : undefined;
            resourceInputs["orgServiceUrl"] = (args ? args.orgServiceUrl : undefined) ?? utilities.getEnv("AZDO_ORG_SERVICE_URL");
            resourceInputs["personalAccessToken"] = args?.personalAccessToken ? pulumi.secret(args.personalAccessToken) : undefined;
            resourceInputs["tenantId"] = args ? args.tenantId : undefined;
            resourceInputs["tenantIdApply"] = args ? args.tenantIdApply : undefined;
            resourceInputs["tenantIdPlan"] = args ? args.tenantIdPlan : undefined;
            resourceInputs["useMsi"] = pulumi.output(args ? args.useMsi : undefined).apply(JSON.stringify);
            resourceInputs["useOidc"] = pulumi.output(args ? args.useOidc : undefined).apply(JSON.stringify);
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["clientCertificate", "clientCertificatePassword", "clientSecret", "oidcToken", "personalAccessToken"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Provider.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Provider resource.
 */
export interface ProviderArgs {
    /**
     * Base64 encoded certificate to use to authenticate to the service principal.
     */
    clientCertificate?: pulumi.Input<string>;
    /**
     * Password for a client certificate password.
     */
    clientCertificatePassword?: pulumi.Input<string>;
    /**
     * Path to a certificate to use to authenticate to the service principal.
     */
    clientCertificatePath?: pulumi.Input<string>;
    /**
     * The service principal client or managed service principal id which should be used.
     */
    clientId?: pulumi.Input<string>;
    /**
     * The service principal client id which should be used during an apply operation in Terraform Cloud.
     */
    clientIdApply?: pulumi.Input<string>;
    /**
     * The service principal client id which should be used during a plan operation in Terraform Cloud.
     */
    clientIdPlan?: pulumi.Input<string>;
    /**
     * Client secret for authenticating to a service principal.
     */
    clientSecret?: pulumi.Input<string>;
    /**
     * Path to a file containing a client secret for authenticating to a service principal.
     */
    clientSecretPath?: pulumi.Input<string>;
    /**
     * Set the audience when requesting OIDC tokens.
     */
    oidcAudience?: pulumi.Input<string>;
    /**
     * The bearer token for the request to the OIDC provider. For use when authenticating as a Service Principal using OpenID
     * Connect.
     */
    oidcRequestToken?: pulumi.Input<string>;
    /**
     * The URL for the OIDC provider from which to request an ID token. For use when authenticating as a Service Principal
     * using OpenID Connect.
     */
    oidcRequestUrl?: pulumi.Input<string>;
    /**
     * Terraform Cloud dynamic credential provider tag.
     */
    oidcTfcTag?: pulumi.Input<string>;
    /**
     * OIDC token to authenticate as a service principal.
     */
    oidcToken?: pulumi.Input<string>;
    /**
     * OIDC token from file to authenticate as a service principal.
     */
    oidcTokenFilePath?: pulumi.Input<string>;
    /**
     * The url of the Azure DevOps instance which should be used.
     */
    orgServiceUrl?: pulumi.Input<string>;
    /**
     * The personal access token which should be used.
     */
    personalAccessToken?: pulumi.Input<string>;
    /**
     * The service principal tenant id which should be used.
     */
    tenantId?: pulumi.Input<string>;
    /**
     * The service principal tenant id which should be used during an apply operation in Terraform Cloud..
     */
    tenantIdApply?: pulumi.Input<string>;
    /**
     * The service principal tenant id which should be used during a plan operation in Terraform Cloud.
     */
    tenantIdPlan?: pulumi.Input<string>;
    /**
     * Use an Azure Managed Service Identity.
     */
    useMsi?: pulumi.Input<boolean>;
    /**
     * Use an OIDC token to authenticate to a service principal.
     */
    useOidc?: pulumi.Input<boolean>;
}
