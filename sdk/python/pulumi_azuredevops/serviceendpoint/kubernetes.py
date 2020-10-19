# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['Kubernetes']

warnings.warn("azuredevops.serviceendpoint.Kubernetes has been deprecated in favor of azuredevops.ServiceEndpointKubernetes", DeprecationWarning)


class Kubernetes(pulumi.CustomResource):
    warnings.warn("azuredevops.serviceendpoint.Kubernetes has been deprecated in favor of azuredevops.ServiceEndpointKubernetes", DeprecationWarning)

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
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a Kubernetes service endpoint within Azure DevOps.

        ## Relevant Links

        - [Azure DevOps Service REST API 5.1 - Endpoints](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-5.1)

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
        pulumi.log.warn("Kubernetes is deprecated: azuredevops.serviceendpoint.Kubernetes has been deprecated in favor of azuredevops.ServiceEndpointKubernetes")
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if apiserver_url is None:
                raise TypeError("Missing required property 'apiserver_url'")
            __props__['apiserver_url'] = apiserver_url
            __props__['authorization'] = authorization
            if authorization_type is None:
                raise TypeError("Missing required property 'authorization_type'")
            __props__['authorization_type'] = authorization_type
            __props__['azure_subscriptions'] = azure_subscriptions
            __props__['description'] = description
            __props__['kubeconfigs'] = kubeconfigs
            if project_id is None:
                raise TypeError("Missing required property 'project_id'")
            __props__['project_id'] = project_id
            __props__['service_accounts'] = service_accounts
            if service_endpoint_name is None:
                raise TypeError("Missing required property 'service_endpoint_name'")
            __props__['service_endpoint_name'] = service_endpoint_name
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

        __props__ = dict()

        __props__["apiserver_url"] = apiserver_url
        __props__["authorization"] = authorization
        __props__["authorization_type"] = authorization_type
        __props__["azure_subscriptions"] = azure_subscriptions
        __props__["description"] = description
        __props__["kubeconfigs"] = kubeconfigs
        __props__["project_id"] = project_id
        __props__["service_accounts"] = service_accounts
        __props__["service_endpoint_name"] = service_endpoint_name
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

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

