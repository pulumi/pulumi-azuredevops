// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * ## # azuredevops.ServiceEndpoint.Kubernetes
 *
 * Manages a Kubernetes service endpoint within Azure DevOps.
 *
 *
 * ## Relevant Links
 *
 * * [Azure DevOps Service REST API 5.1 - Endpoints](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-5.1)
 */
export class Kubernetes extends pulumi.CustomResource {
    /**
     * Get an existing Kubernetes resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: KubernetesState, opts?: pulumi.CustomResourceOptions): Kubernetes {
        return new Kubernetes(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuredevops:ServiceEndpoint/kubernetes:Kubernetes';

    /**
     * Returns true if the given object is an instance of Kubernetes.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Kubernetes {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Kubernetes.__pulumiType;
    }

    /**
     * The Service Endpoint description.
     */
    public readonly apiserverUrl!: pulumi.Output<string>;
    public readonly authorization!: pulumi.Output<{[key: string]: string}>;
    /**
     * The authentication method used to authenticate on the Kubernetes cluster. The value should be one of AzureSubscription, Kubeconfig, ServiceAccount.
     */
    public readonly authorizationType!: pulumi.Output<string>;
    /**
     * The configuration for authorization_type="AzureSubscription".
     */
    public readonly azureSubscriptions!: pulumi.Output<outputs.ServiceEndpoint.KubernetesAzureSubscription[] | undefined>;
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The configuration for authorization_type="Kubeconfig".
     */
    public readonly kubeconfigs!: pulumi.Output<outputs.ServiceEndpoint.KubernetesKubeconfig[] | undefined>;
    /**
     * The project ID or project name.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * The configuration for authorization_type="ServiceAccount". This type uses the credentials of a service account currently deployed to the cluster.
     */
    public readonly serviceAccounts!: pulumi.Output<outputs.ServiceEndpoint.KubernetesServiceAccount[] | undefined>;
    /**
     * The Service Endpoint name.
     */
    public readonly serviceEndpointName!: pulumi.Output<string>;

    /**
     * Create a Kubernetes resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: KubernetesArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: KubernetesArgs | KubernetesState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as KubernetesState | undefined;
            inputs["apiserverUrl"] = state ? state.apiserverUrl : undefined;
            inputs["authorization"] = state ? state.authorization : undefined;
            inputs["authorizationType"] = state ? state.authorizationType : undefined;
            inputs["azureSubscriptions"] = state ? state.azureSubscriptions : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["kubeconfigs"] = state ? state.kubeconfigs : undefined;
            inputs["projectId"] = state ? state.projectId : undefined;
            inputs["serviceAccounts"] = state ? state.serviceAccounts : undefined;
            inputs["serviceEndpointName"] = state ? state.serviceEndpointName : undefined;
        } else {
            const args = argsOrState as KubernetesArgs | undefined;
            if (!args || args.apiserverUrl === undefined) {
                throw new Error("Missing required property 'apiserverUrl'");
            }
            if (!args || args.authorizationType === undefined) {
                throw new Error("Missing required property 'authorizationType'");
            }
            if (!args || args.projectId === undefined) {
                throw new Error("Missing required property 'projectId'");
            }
            if (!args || args.serviceEndpointName === undefined) {
                throw new Error("Missing required property 'serviceEndpointName'");
            }
            inputs["apiserverUrl"] = args ? args.apiserverUrl : undefined;
            inputs["authorization"] = args ? args.authorization : undefined;
            inputs["authorizationType"] = args ? args.authorizationType : undefined;
            inputs["azureSubscriptions"] = args ? args.azureSubscriptions : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["kubeconfigs"] = args ? args.kubeconfigs : undefined;
            inputs["projectId"] = args ? args.projectId : undefined;
            inputs["serviceAccounts"] = args ? args.serviceAccounts : undefined;
            inputs["serviceEndpointName"] = args ? args.serviceEndpointName : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Kubernetes.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Kubernetes resources.
 */
export interface KubernetesState {
    /**
     * The Service Endpoint description.
     */
    readonly apiserverUrl?: pulumi.Input<string>;
    readonly authorization?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The authentication method used to authenticate on the Kubernetes cluster. The value should be one of AzureSubscription, Kubeconfig, ServiceAccount.
     */
    readonly authorizationType?: pulumi.Input<string>;
    /**
     * The configuration for authorization_type="AzureSubscription".
     */
    readonly azureSubscriptions?: pulumi.Input<pulumi.Input<inputs.ServiceEndpoint.KubernetesAzureSubscription>[]>;
    readonly description?: pulumi.Input<string>;
    /**
     * The configuration for authorization_type="Kubeconfig".
     */
    readonly kubeconfigs?: pulumi.Input<pulumi.Input<inputs.ServiceEndpoint.KubernetesKubeconfig>[]>;
    /**
     * The project ID or project name.
     */
    readonly projectId?: pulumi.Input<string>;
    /**
     * The configuration for authorization_type="ServiceAccount". This type uses the credentials of a service account currently deployed to the cluster.
     */
    readonly serviceAccounts?: pulumi.Input<pulumi.Input<inputs.ServiceEndpoint.KubernetesServiceAccount>[]>;
    /**
     * The Service Endpoint name.
     */
    readonly serviceEndpointName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Kubernetes resource.
 */
export interface KubernetesArgs {
    /**
     * The Service Endpoint description.
     */
    readonly apiserverUrl: pulumi.Input<string>;
    readonly authorization?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The authentication method used to authenticate on the Kubernetes cluster. The value should be one of AzureSubscription, Kubeconfig, ServiceAccount.
     */
    readonly authorizationType: pulumi.Input<string>;
    /**
     * The configuration for authorization_type="AzureSubscription".
     */
    readonly azureSubscriptions?: pulumi.Input<pulumi.Input<inputs.ServiceEndpoint.KubernetesAzureSubscription>[]>;
    readonly description?: pulumi.Input<string>;
    /**
     * The configuration for authorization_type="Kubeconfig".
     */
    readonly kubeconfigs?: pulumi.Input<pulumi.Input<inputs.ServiceEndpoint.KubernetesKubeconfig>[]>;
    /**
     * The project ID or project name.
     */
    readonly projectId: pulumi.Input<string>;
    /**
     * The configuration for authorization_type="ServiceAccount". This type uses the credentials of a service account currently deployed to the cluster.
     */
    readonly serviceAccounts?: pulumi.Input<pulumi.Input<inputs.ServiceEndpoint.KubernetesServiceAccount>[]>;
    /**
     * The Service Endpoint name.
     */
    readonly serviceEndpointName: pulumi.Input<string>;
}
