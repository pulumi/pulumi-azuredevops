# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['KubernetesArgs', 'Kubernetes']

@pulumi.input_type
class KubernetesArgs:
    def __init__(__self__, *,
                 apiserver_url: pulumi.Input[str],
                 authorization_type: pulumi.Input[str],
                 project_id: pulumi.Input[str],
                 service_endpoint_name: pulumi.Input[str],
                 authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 azure_subscriptions: Optional[pulumi.Input[Sequence[pulumi.Input['KubernetesAzureSubscriptionArgs']]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 kubeconfigs: Optional[pulumi.Input[Sequence[pulumi.Input['KubernetesKubeconfigArgs']]]] = None,
                 service_accounts: Optional[pulumi.Input[Sequence[pulumi.Input['KubernetesServiceAccountArgs']]]] = None):
        """
        The set of arguments for constructing a Kubernetes resource.
        :param pulumi.Input[str] apiserver_url: The Service Endpoint description.
        :param pulumi.Input[str] authorization_type: The authentication method used to authenticate on the Kubernetes cluster. The value should be one of AzureSubscription, Kubeconfig, ServiceAccount.
        :param pulumi.Input[str] project_id: The project ID or project name.
        :param pulumi.Input[str] service_endpoint_name: The Service Endpoint name.
        :param pulumi.Input[Sequence[pulumi.Input['KubernetesAzureSubscriptionArgs']]] azure_subscriptions: The configuration for authorization_type="AzureSubscription".
        :param pulumi.Input[Sequence[pulumi.Input['KubernetesKubeconfigArgs']]] kubeconfigs: The configuration for authorization_type="Kubeconfig".
        :param pulumi.Input[Sequence[pulumi.Input['KubernetesServiceAccountArgs']]] service_accounts: The configuration for authorization_type="ServiceAccount". This type uses the credentials of a service account currently deployed to the cluster.
        """
        pulumi.set(__self__, "apiserver_url", apiserver_url)
        pulumi.set(__self__, "authorization_type", authorization_type)
        pulumi.set(__self__, "project_id", project_id)
        pulumi.set(__self__, "service_endpoint_name", service_endpoint_name)
        if authorization is not None:
            pulumi.set(__self__, "authorization", authorization)
        if azure_subscriptions is not None:
            pulumi.set(__self__, "azure_subscriptions", azure_subscriptions)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if kubeconfigs is not None:
            pulumi.set(__self__, "kubeconfigs", kubeconfigs)
        if service_accounts is not None:
            pulumi.set(__self__, "service_accounts", service_accounts)

    @property
    @pulumi.getter(name="apiserverUrl")
    def apiserver_url(self) -> pulumi.Input[str]:
        """
        The Service Endpoint description.
        """
        return pulumi.get(self, "apiserver_url")

    @apiserver_url.setter
    def apiserver_url(self, value: pulumi.Input[str]):
        pulumi.set(self, "apiserver_url", value)

    @property
    @pulumi.getter(name="authorizationType")
    def authorization_type(self) -> pulumi.Input[str]:
        """
        The authentication method used to authenticate on the Kubernetes cluster. The value should be one of AzureSubscription, Kubeconfig, ServiceAccount.
        """
        return pulumi.get(self, "authorization_type")

    @authorization_type.setter
    def authorization_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "authorization_type", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Input[str]:
        """
        The project ID or project name.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter(name="serviceEndpointName")
    def service_endpoint_name(self) -> pulumi.Input[str]:
        """
        The Service Endpoint name.
        """
        return pulumi.get(self, "service_endpoint_name")

    @service_endpoint_name.setter
    def service_endpoint_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "service_endpoint_name", value)

    @property
    @pulumi.getter
    def authorization(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        return pulumi.get(self, "authorization")

    @authorization.setter
    def authorization(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "authorization", value)

    @property
    @pulumi.getter(name="azureSubscriptions")
    def azure_subscriptions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['KubernetesAzureSubscriptionArgs']]]]:
        """
        The configuration for authorization_type="AzureSubscription".
        """
        return pulumi.get(self, "azure_subscriptions")

    @azure_subscriptions.setter
    def azure_subscriptions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['KubernetesAzureSubscriptionArgs']]]]):
        pulumi.set(self, "azure_subscriptions", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def kubeconfigs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['KubernetesKubeconfigArgs']]]]:
        """
        The configuration for authorization_type="Kubeconfig".
        """
        return pulumi.get(self, "kubeconfigs")

    @kubeconfigs.setter
    def kubeconfigs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['KubernetesKubeconfigArgs']]]]):
        pulumi.set(self, "kubeconfigs", value)

    @property
    @pulumi.getter(name="serviceAccounts")
    def service_accounts(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['KubernetesServiceAccountArgs']]]]:
        """
        The configuration for authorization_type="ServiceAccount". This type uses the credentials of a service account currently deployed to the cluster.
        """
        return pulumi.get(self, "service_accounts")

    @service_accounts.setter
    def service_accounts(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['KubernetesServiceAccountArgs']]]]):
        pulumi.set(self, "service_accounts", value)


@pulumi.input_type
class _KubernetesState:
    def __init__(__self__, *,
                 apiserver_url: Optional[pulumi.Input[str]] = None,
                 authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 authorization_type: Optional[pulumi.Input[str]] = None,
                 azure_subscriptions: Optional[pulumi.Input[Sequence[pulumi.Input['KubernetesAzureSubscriptionArgs']]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 kubeconfigs: Optional[pulumi.Input[Sequence[pulumi.Input['KubernetesKubeconfigArgs']]]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 service_accounts: Optional[pulumi.Input[Sequence[pulumi.Input['KubernetesServiceAccountArgs']]]] = None,
                 service_endpoint_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Kubernetes resources.
        :param pulumi.Input[str] apiserver_url: The Service Endpoint description.
        :param pulumi.Input[str] authorization_type: The authentication method used to authenticate on the Kubernetes cluster. The value should be one of AzureSubscription, Kubeconfig, ServiceAccount.
        :param pulumi.Input[Sequence[pulumi.Input['KubernetesAzureSubscriptionArgs']]] azure_subscriptions: The configuration for authorization_type="AzureSubscription".
        :param pulumi.Input[Sequence[pulumi.Input['KubernetesKubeconfigArgs']]] kubeconfigs: The configuration for authorization_type="Kubeconfig".
        :param pulumi.Input[str] project_id: The project ID or project name.
        :param pulumi.Input[Sequence[pulumi.Input['KubernetesServiceAccountArgs']]] service_accounts: The configuration for authorization_type="ServiceAccount". This type uses the credentials of a service account currently deployed to the cluster.
        :param pulumi.Input[str] service_endpoint_name: The Service Endpoint name.
        """
        if apiserver_url is not None:
            pulumi.set(__self__, "apiserver_url", apiserver_url)
        if authorization is not None:
            pulumi.set(__self__, "authorization", authorization)
        if authorization_type is not None:
            pulumi.set(__self__, "authorization_type", authorization_type)
        if azure_subscriptions is not None:
            pulumi.set(__self__, "azure_subscriptions", azure_subscriptions)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if kubeconfigs is not None:
            pulumi.set(__self__, "kubeconfigs", kubeconfigs)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if service_accounts is not None:
            pulumi.set(__self__, "service_accounts", service_accounts)
        if service_endpoint_name is not None:
            pulumi.set(__self__, "service_endpoint_name", service_endpoint_name)

    @property
    @pulumi.getter(name="apiserverUrl")
    def apiserver_url(self) -> Optional[pulumi.Input[str]]:
        """
        The Service Endpoint description.
        """
        return pulumi.get(self, "apiserver_url")

    @apiserver_url.setter
    def apiserver_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "apiserver_url", value)

    @property
    @pulumi.getter
    def authorization(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        return pulumi.get(self, "authorization")

    @authorization.setter
    def authorization(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "authorization", value)

    @property
    @pulumi.getter(name="authorizationType")
    def authorization_type(self) -> Optional[pulumi.Input[str]]:
        """
        The authentication method used to authenticate on the Kubernetes cluster. The value should be one of AzureSubscription, Kubeconfig, ServiceAccount.
        """
        return pulumi.get(self, "authorization_type")

    @authorization_type.setter
    def authorization_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "authorization_type", value)

    @property
    @pulumi.getter(name="azureSubscriptions")
    def azure_subscriptions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['KubernetesAzureSubscriptionArgs']]]]:
        """
        The configuration for authorization_type="AzureSubscription".
        """
        return pulumi.get(self, "azure_subscriptions")

    @azure_subscriptions.setter
    def azure_subscriptions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['KubernetesAzureSubscriptionArgs']]]]):
        pulumi.set(self, "azure_subscriptions", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def kubeconfigs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['KubernetesKubeconfigArgs']]]]:
        """
        The configuration for authorization_type="Kubeconfig".
        """
        return pulumi.get(self, "kubeconfigs")

    @kubeconfigs.setter
    def kubeconfigs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['KubernetesKubeconfigArgs']]]]):
        pulumi.set(self, "kubeconfigs", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        The project ID or project name.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter(name="serviceAccounts")
    def service_accounts(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['KubernetesServiceAccountArgs']]]]:
        """
        The configuration for authorization_type="ServiceAccount". This type uses the credentials of a service account currently deployed to the cluster.
        """
        return pulumi.get(self, "service_accounts")

    @service_accounts.setter
    def service_accounts(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['KubernetesServiceAccountArgs']]]]):
        pulumi.set(self, "service_accounts", value)

    @property
    @pulumi.getter(name="serviceEndpointName")
    def service_endpoint_name(self) -> Optional[pulumi.Input[str]]:
        """
        The Service Endpoint name.
        """
        return pulumi.get(self, "service_endpoint_name")

    @service_endpoint_name.setter
    def service_endpoint_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_endpoint_name", value)


warnings.warn("""azuredevops.serviceendpoint.Kubernetes has been deprecated in favor of azuredevops.ServiceEndpointKubernetes""", DeprecationWarning)


class Kubernetes(pulumi.CustomResource):
    warnings.warn("""azuredevops.serviceendpoint.Kubernetes has been deprecated in favor of azuredevops.ServiceEndpointKubernetes""", DeprecationWarning)

    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 apiserver_url: Optional[pulumi.Input[str]] = None,
                 authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 authorization_type: Optional[pulumi.Input[str]] = None,
                 azure_subscriptions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['KubernetesAzureSubscriptionArgs']]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 kubeconfigs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['KubernetesKubeconfigArgs']]]]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 service_accounts: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['KubernetesServiceAccountArgs']]]]] = None,
                 service_endpoint_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages a Kubernetes service endpoint within Azure DevOps.

        ## Relevant Links

        - [Azure DevOps Service REST API 5.1 - Endpoints](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-5.1)

        ## Import

        Azure DevOps Service Endpoint Kubernetes can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**

        ```sh
         $ pulumi import azuredevops:ServiceEndpoint/kubernetes:Kubernetes serviceendpoint 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] apiserver_url: The Service Endpoint description.
        :param pulumi.Input[str] authorization_type: The authentication method used to authenticate on the Kubernetes cluster. The value should be one of AzureSubscription, Kubeconfig, ServiceAccount.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['KubernetesAzureSubscriptionArgs']]]] azure_subscriptions: The configuration for authorization_type="AzureSubscription".
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['KubernetesKubeconfigArgs']]]] kubeconfigs: The configuration for authorization_type="Kubeconfig".
        :param pulumi.Input[str] project_id: The project ID or project name.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['KubernetesServiceAccountArgs']]]] service_accounts: The configuration for authorization_type="ServiceAccount". This type uses the credentials of a service account currently deployed to the cluster.
        :param pulumi.Input[str] service_endpoint_name: The Service Endpoint name.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: KubernetesArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a Kubernetes service endpoint within Azure DevOps.

        ## Relevant Links

        - [Azure DevOps Service REST API 5.1 - Endpoints](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-5.1)

        ## Import

        Azure DevOps Service Endpoint Kubernetes can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**

        ```sh
         $ pulumi import azuredevops:ServiceEndpoint/kubernetes:Kubernetes serviceendpoint 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
        ```

        :param str resource_name: The name of the resource.
        :param KubernetesArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(KubernetesArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 apiserver_url: Optional[pulumi.Input[str]] = None,
                 authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 authorization_type: Optional[pulumi.Input[str]] = None,
                 azure_subscriptions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['KubernetesAzureSubscriptionArgs']]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 kubeconfigs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['KubernetesKubeconfigArgs']]]]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 service_accounts: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['KubernetesServiceAccountArgs']]]]] = None,
                 service_endpoint_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        pulumi.log.warn("""Kubernetes is deprecated: azuredevops.serviceendpoint.Kubernetes has been deprecated in favor of azuredevops.ServiceEndpointKubernetes""")
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = KubernetesArgs.__new__(KubernetesArgs)

            if apiserver_url is None and not opts.urn:
                raise TypeError("Missing required property 'apiserver_url'")
            __props__.__dict__["apiserver_url"] = apiserver_url
            __props__.__dict__["authorization"] = authorization
            if authorization_type is None and not opts.urn:
                raise TypeError("Missing required property 'authorization_type'")
            __props__.__dict__["authorization_type"] = authorization_type
            __props__.__dict__["azure_subscriptions"] = azure_subscriptions
            __props__.__dict__["description"] = description
            __props__.__dict__["kubeconfigs"] = kubeconfigs
            if project_id is None and not opts.urn:
                raise TypeError("Missing required property 'project_id'")
            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["service_accounts"] = service_accounts
            if service_endpoint_name is None and not opts.urn:
                raise TypeError("Missing required property 'service_endpoint_name'")
            __props__.__dict__["service_endpoint_name"] = service_endpoint_name
        super(Kubernetes, __self__).__init__(
            'azuredevops:ServiceEndpoint/kubernetes:Kubernetes',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            apiserver_url: Optional[pulumi.Input[str]] = None,
            authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            authorization_type: Optional[pulumi.Input[str]] = None,
            azure_subscriptions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['KubernetesAzureSubscriptionArgs']]]]] = None,
            description: Optional[pulumi.Input[str]] = None,
            kubeconfigs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['KubernetesKubeconfigArgs']]]]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            service_accounts: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['KubernetesServiceAccountArgs']]]]] = None,
            service_endpoint_name: Optional[pulumi.Input[str]] = None) -> 'Kubernetes':
        """
        Get an existing Kubernetes resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] apiserver_url: The Service Endpoint description.
        :param pulumi.Input[str] authorization_type: The authentication method used to authenticate on the Kubernetes cluster. The value should be one of AzureSubscription, Kubeconfig, ServiceAccount.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['KubernetesAzureSubscriptionArgs']]]] azure_subscriptions: The configuration for authorization_type="AzureSubscription".
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['KubernetesKubeconfigArgs']]]] kubeconfigs: The configuration for authorization_type="Kubeconfig".
        :param pulumi.Input[str] project_id: The project ID or project name.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['KubernetesServiceAccountArgs']]]] service_accounts: The configuration for authorization_type="ServiceAccount". This type uses the credentials of a service account currently deployed to the cluster.
        :param pulumi.Input[str] service_endpoint_name: The Service Endpoint name.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _KubernetesState.__new__(_KubernetesState)

        __props__.__dict__["apiserver_url"] = apiserver_url
        __props__.__dict__["authorization"] = authorization
        __props__.__dict__["authorization_type"] = authorization_type
        __props__.__dict__["azure_subscriptions"] = azure_subscriptions
        __props__.__dict__["description"] = description
        __props__.__dict__["kubeconfigs"] = kubeconfigs
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["service_accounts"] = service_accounts
        __props__.__dict__["service_endpoint_name"] = service_endpoint_name
        return Kubernetes(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="apiserverUrl")
    def apiserver_url(self) -> pulumi.Output[str]:
        """
        The Service Endpoint description.
        """
        return pulumi.get(self, "apiserver_url")

    @property
    @pulumi.getter
    def authorization(self) -> pulumi.Output[Mapping[str, str]]:
        return pulumi.get(self, "authorization")

    @property
    @pulumi.getter(name="authorizationType")
    def authorization_type(self) -> pulumi.Output[str]:
        """
        The authentication method used to authenticate on the Kubernetes cluster. The value should be one of AzureSubscription, Kubeconfig, ServiceAccount.
        """
        return pulumi.get(self, "authorization_type")

    @property
    @pulumi.getter(name="azureSubscriptions")
    def azure_subscriptions(self) -> pulumi.Output[Optional[Sequence['outputs.KubernetesAzureSubscription']]]:
        """
        The configuration for authorization_type="AzureSubscription".
        """
        return pulumi.get(self, "azure_subscriptions")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def kubeconfigs(self) -> pulumi.Output[Optional[Sequence['outputs.KubernetesKubeconfig']]]:
        """
        The configuration for authorization_type="Kubeconfig".
        """
        return pulumi.get(self, "kubeconfigs")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        """
        The project ID or project name.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="serviceAccounts")
    def service_accounts(self) -> pulumi.Output[Optional[Sequence['outputs.KubernetesServiceAccount']]]:
        """
        The configuration for authorization_type="ServiceAccount". This type uses the credentials of a service account currently deployed to the cluster.
        """
        return pulumi.get(self, "service_accounts")

    @property
    @pulumi.getter(name="serviceEndpointName")
    def service_endpoint_name(self) -> pulumi.Output[str]:
        """
        The Service Endpoint name.
        """
        return pulumi.get(self, "service_endpoint_name")

