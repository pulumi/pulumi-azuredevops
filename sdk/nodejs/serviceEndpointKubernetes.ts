// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Manages a Kubernetes service endpoint within Azure DevOps.
 *
 * ## Relevant Links
 *
 * - [Azure DevOps Service REST API 5.1 - Endpoints](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-5.1)
 *
 * ## Import
 *
 * Azure DevOps Service Endpoint Kubernetes can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**
 *
 * ```sh
 *  $ pulumi import azuredevops:index/serviceEndpointKubernetes:ServiceEndpointKubernetes serviceendpoint 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
 * ```
 */
export class ServiceEndpointKubernetes extends pulumi.CustomResource {
    /**
     * Get an existing ServiceEndpointKubernetes resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServiceEndpointKubernetesState, opts?: pulumi.CustomResourceOptions): ServiceEndpointKubernetes {
        return new ServiceEndpointKubernetes(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuredevops:index/serviceEndpointKubernetes:ServiceEndpointKubernetes';

    /**
     * Returns true if the given object is an instance of ServiceEndpointKubernetes.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServiceEndpointKubernetes {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServiceEndpointKubernetes.__pulumiType;
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
     * The configuration for authorization_type="AzureSubscription".
     */
    public readonly azureSubscriptions!: pulumi.Output<outputs.ServiceEndpointKubernetesAzureSubscription[] | undefined>;
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The configuration for authorization_type="Kubeconfig".
     */
    public readonly kubeconfigs!: pulumi.Output<outputs.ServiceEndpointKubernetesKubeconfig[] | undefined>;
    /**
     * The project ID or project name.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * The configuration for authorization_type="ServiceAccount". This type uses the credentials of a service account currently deployed to the cluster.
     */
    public readonly serviceAccounts!: pulumi.Output<outputs.ServiceEndpointKubernetesServiceAccount[] | undefined>;
    /**
     * The Service Endpoint name.
     */
    public readonly serviceEndpointName!: pulumi.Output<string>;

    /**
     * Create a ServiceEndpointKubernetes resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServiceEndpointKubernetesArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServiceEndpointKubernetesArgs | ServiceEndpointKubernetesState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ServiceEndpointKubernetesState | undefined;
            resourceInputs["apiserverUrl"] = state ? state.apiserverUrl : undefined;
            resourceInputs["authorization"] = state ? state.authorization : undefined;
            resourceInputs["authorizationType"] = state ? state.authorizationType : undefined;
            resourceInputs["azureSubscriptions"] = state ? state.azureSubscriptions : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["kubeconfigs"] = state ? state.kubeconfigs : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["serviceAccounts"] = state ? state.serviceAccounts : undefined;
            resourceInputs["serviceEndpointName"] = state ? state.serviceEndpointName : undefined;
        } else {
            const args = argsOrState as ServiceEndpointKubernetesArgs | undefined;
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
            resourceInputs["serviceAccounts"] = args ? args.serviceAccounts : undefined;
            resourceInputs["serviceEndpointName"] = args ? args.serviceEndpointName : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "azuredevops:ServiceEndpoint/kubernetes:Kubernetes" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(ServiceEndpointKubernetes.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ServiceEndpointKubernetes resources.
 */
export interface ServiceEndpointKubernetesState {
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
     * The configuration for authorization_type="AzureSubscription".
     */
    azureSubscriptions?: pulumi.Input<pulumi.Input<inputs.ServiceEndpointKubernetesAzureSubscription>[]>;
    description?: pulumi.Input<string>;
    /**
     * The configuration for authorization_type="Kubeconfig".
     */
    kubeconfigs?: pulumi.Input<pulumi.Input<inputs.ServiceEndpointKubernetesKubeconfig>[]>;
    /**
     * The project ID or project name.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The configuration for authorization_type="ServiceAccount". This type uses the credentials of a service account currently deployed to the cluster.
     */
    serviceAccounts?: pulumi.Input<pulumi.Input<inputs.ServiceEndpointKubernetesServiceAccount>[]>;
    /**
     * The Service Endpoint name.
     */
    serviceEndpointName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ServiceEndpointKubernetes resource.
 */
export interface ServiceEndpointKubernetesArgs {
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
     * The configuration for authorization_type="AzureSubscription".
     */
    azureSubscriptions?: pulumi.Input<pulumi.Input<inputs.ServiceEndpointKubernetesAzureSubscription>[]>;
    description?: pulumi.Input<string>;
    /**
     * The configuration for authorization_type="Kubeconfig".
     */
    kubeconfigs?: pulumi.Input<pulumi.Input<inputs.ServiceEndpointKubernetesKubeconfig>[]>;
    /**
     * The project ID or project name.
     */
    projectId: pulumi.Input<string>;
    /**
     * The configuration for authorization_type="ServiceAccount". This type uses the credentials of a service account currently deployed to the cluster.
     */
    serviceAccounts?: pulumi.Input<pulumi.Input<inputs.ServiceEndpointKubernetesServiceAccount>[]>;
    /**
     * The Service Endpoint name.
     */
    serviceEndpointName: pulumi.Input<string>;
}
