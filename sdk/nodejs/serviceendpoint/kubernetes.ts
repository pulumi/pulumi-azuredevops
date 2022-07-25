// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages a Kubernetes service endpoint within Azure DevOps.
 *
 * ## Relevant Links
 *
 * - [Azure DevOps Service REST API 6.0 - Endpoints](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-6.0)
 *
 * ## Import
 *
 * Azure DevOps Service Endpoint Kubernetes can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**
 *
 * ```sh
 *  $ pulumi import azuredevops:ServiceEndpoint/kubernetes:Kubernetes example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
 * ```
 *
 * @deprecated azuredevops.serviceendpoint.Kubernetes has been deprecated in favor of azuredevops.ServiceEndpointKubernetes
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
        pulumi.log.warn("Kubernetes is deprecated: azuredevops.serviceendpoint.Kubernetes has been deprecated in favor of azuredevops.ServiceEndpointKubernetes")
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
     * The hostname (in form of URI) of the Kubernetes API.
     */
    public readonly apiserverUrl!: pulumi.Output<string>;
    public readonly authorization!: pulumi.Output<{[key: string]: string}>;
    /**
     * The authentication method used to authenticate on the Kubernetes cluster. The value should be one of AzureSubscription, Kubeconfig, ServiceAccount.
     */
    public readonly authorizationType!: pulumi.Output<string>;
    /**
     * A `azureSubscription` block defined blow.
     */
    public readonly azureSubscriptions!: pulumi.Output<outputs.ServiceEndpoint.KubernetesAzureSubscription[] | undefined>;
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * A `kubeconfig` block defined blow.
     */
    public readonly kubeconfigs!: pulumi.Output<outputs.ServiceEndpoint.KubernetesKubeconfig[] | undefined>;
    /**
     * The ID of the project.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * A `serviceAccount` block defined blow.
     */
    public readonly serviceAccount!: pulumi.Output<outputs.ServiceEndpoint.KubernetesServiceAccount | undefined>;
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
    /** @deprecated azuredevops.serviceendpoint.Kubernetes has been deprecated in favor of azuredevops.ServiceEndpointKubernetes */
    constructor(name: string, args: KubernetesArgs, opts?: pulumi.CustomResourceOptions)
    /** @deprecated azuredevops.serviceendpoint.Kubernetes has been deprecated in favor of azuredevops.ServiceEndpointKubernetes */
    constructor(name: string, argsOrState?: KubernetesArgs | KubernetesState, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("Kubernetes is deprecated: azuredevops.serviceendpoint.Kubernetes has been deprecated in favor of azuredevops.ServiceEndpointKubernetes")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as KubernetesState | undefined;
            resourceInputs["apiserverUrl"] = state ? state.apiserverUrl : undefined;
            resourceInputs["authorization"] = state ? state.authorization : undefined;
            resourceInputs["authorizationType"] = state ? state.authorizationType : undefined;
            resourceInputs["azureSubscriptions"] = state ? state.azureSubscriptions : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["kubeconfigs"] = state ? state.kubeconfigs : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["serviceAccount"] = state ? state.serviceAccount : undefined;
            resourceInputs["serviceEndpointName"] = state ? state.serviceEndpointName : undefined;
        } else {
            const args = argsOrState as KubernetesArgs | undefined;
            if ((!args || args.apiserverUrl === undefined) && !opts.urn) {
                throw new Error("Missing required property 'apiserverUrl'");
            }
            if ((!args || args.authorizationType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'authorizationType'");
            }
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            if ((!args || args.serviceEndpointName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceEndpointName'");
            }
            resourceInputs["apiserverUrl"] = args ? args.apiserverUrl : undefined;
            resourceInputs["authorization"] = args ? args.authorization : undefined;
            resourceInputs["authorizationType"] = args ? args.authorizationType : undefined;
            resourceInputs["azureSubscriptions"] = args ? args.azureSubscriptions : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["kubeconfigs"] = args ? args.kubeconfigs : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["serviceAccount"] = args ? args.serviceAccount : undefined;
            resourceInputs["serviceEndpointName"] = args ? args.serviceEndpointName : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Kubernetes.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Kubernetes resources.
 */
export interface KubernetesState {
    /**
     * The hostname (in form of URI) of the Kubernetes API.
     */
    apiserverUrl?: pulumi.Input<string>;
    authorization?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The authentication method used to authenticate on the Kubernetes cluster. The value should be one of AzureSubscription, Kubeconfig, ServiceAccount.
     */
    authorizationType?: pulumi.Input<string>;
    /**
     * A `azureSubscription` block defined blow.
     */
    azureSubscriptions?: pulumi.Input<pulumi.Input<inputs.ServiceEndpoint.KubernetesAzureSubscription>[]>;
    description?: pulumi.Input<string>;
    /**
     * A `kubeconfig` block defined blow.
     */
    kubeconfigs?: pulumi.Input<pulumi.Input<inputs.ServiceEndpoint.KubernetesKubeconfig>[]>;
    /**
     * The ID of the project.
     */
    projectId?: pulumi.Input<string>;
    /**
     * A `serviceAccount` block defined blow.
     */
    serviceAccount?: pulumi.Input<inputs.ServiceEndpoint.KubernetesServiceAccount>;
    /**
     * The Service Endpoint name.
     */
    serviceEndpointName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Kubernetes resource.
 */
export interface KubernetesArgs {
    /**
     * The hostname (in form of URI) of the Kubernetes API.
     */
    apiserverUrl: pulumi.Input<string>;
    authorization?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The authentication method used to authenticate on the Kubernetes cluster. The value should be one of AzureSubscription, Kubeconfig, ServiceAccount.
     */
    authorizationType: pulumi.Input<string>;
    /**
     * A `azureSubscription` block defined blow.
     */
    azureSubscriptions?: pulumi.Input<pulumi.Input<inputs.ServiceEndpoint.KubernetesAzureSubscription>[]>;
    description?: pulumi.Input<string>;
    /**
     * A `kubeconfig` block defined blow.
     */
    kubeconfigs?: pulumi.Input<pulumi.Input<inputs.ServiceEndpoint.KubernetesKubeconfig>[]>;
    /**
     * The ID of the project.
     */
    projectId: pulumi.Input<string>;
    /**
     * A `serviceAccount` block defined blow.
     */
    serviceAccount?: pulumi.Input<inputs.ServiceEndpoint.KubernetesServiceAccount>;
    /**
     * The Service Endpoint name.
     */
    serviceEndpointName: pulumi.Input<string>;
}
