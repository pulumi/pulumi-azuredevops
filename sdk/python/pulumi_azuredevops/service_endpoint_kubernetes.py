# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins as _builtins
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['ServiceEndpointKubernetesArgs', 'ServiceEndpointKubernetes']

@pulumi.input_type
class ServiceEndpointKubernetesArgs:
    def __init__(__self__, *,
                 apiserver_url: pulumi.Input[_builtins.str],
                 authorization_type: pulumi.Input[_builtins.str],
                 project_id: pulumi.Input[_builtins.str],
                 service_endpoint_name: pulumi.Input[_builtins.str],
                 azure_subscriptions: Optional[pulumi.Input[Sequence[pulumi.Input['ServiceEndpointKubernetesAzureSubscriptionArgs']]]] = None,
                 description: Optional[pulumi.Input[_builtins.str]] = None,
                 kubeconfig: Optional[pulumi.Input['ServiceEndpointKubernetesKubeconfigArgs']] = None,
                 service_account: Optional[pulumi.Input['ServiceEndpointKubernetesServiceAccountArgs']] = None):
        """
        The set of arguments for constructing a ServiceEndpointKubernetes resource.
        :param pulumi.Input[_builtins.str] apiserver_url: The hostname (in form of URI) of the Kubernetes API.
        :param pulumi.Input[_builtins.str] authorization_type: The authentication method used to authenticate on the Kubernetes cluster. The value should be one of AzureSubscription, Kubeconfig, ServiceAccount.
        :param pulumi.Input[_builtins.str] project_id: The ID of the project.
        :param pulumi.Input[_builtins.str] service_endpoint_name: The Service Endpoint name.
        :param pulumi.Input[Sequence[pulumi.Input['ServiceEndpointKubernetesAzureSubscriptionArgs']]] azure_subscriptions: An `azure_subscription` block as defined below.
        :param pulumi.Input['ServiceEndpointKubernetesKubeconfigArgs'] kubeconfig: A `kubeconfig` block as defined below.
        :param pulumi.Input['ServiceEndpointKubernetesServiceAccountArgs'] service_account: A `service_account` block as defined below.
        """
        pulumi.set(__self__, "apiserver_url", apiserver_url)
        pulumi.set(__self__, "authorization_type", authorization_type)
        pulumi.set(__self__, "project_id", project_id)
        pulumi.set(__self__, "service_endpoint_name", service_endpoint_name)
        if azure_subscriptions is not None:
            pulumi.set(__self__, "azure_subscriptions", azure_subscriptions)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if kubeconfig is not None:
            pulumi.set(__self__, "kubeconfig", kubeconfig)
        if service_account is not None:
            pulumi.set(__self__, "service_account", service_account)

    @_builtins.property
    @pulumi.getter(name="apiserverUrl")
    def apiserver_url(self) -> pulumi.Input[_builtins.str]:
        """
        The hostname (in form of URI) of the Kubernetes API.
        """
        return pulumi.get(self, "apiserver_url")

    @apiserver_url.setter
    def apiserver_url(self, value: pulumi.Input[_builtins.str]):
        pulumi.set(self, "apiserver_url", value)

    @_builtins.property
    @pulumi.getter(name="authorizationType")
    def authorization_type(self) -> pulumi.Input[_builtins.str]:
        """
        The authentication method used to authenticate on the Kubernetes cluster. The value should be one of AzureSubscription, Kubeconfig, ServiceAccount.
        """
        return pulumi.get(self, "authorization_type")

    @authorization_type.setter
    def authorization_type(self, value: pulumi.Input[_builtins.str]):
        pulumi.set(self, "authorization_type", value)

    @_builtins.property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Input[_builtins.str]:
        """
        The ID of the project.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: pulumi.Input[_builtins.str]):
        pulumi.set(self, "project_id", value)

    @_builtins.property
    @pulumi.getter(name="serviceEndpointName")
    def service_endpoint_name(self) -> pulumi.Input[_builtins.str]:
        """
        The Service Endpoint name.
        """
        return pulumi.get(self, "service_endpoint_name")

    @service_endpoint_name.setter
    def service_endpoint_name(self, value: pulumi.Input[_builtins.str]):
        pulumi.set(self, "service_endpoint_name", value)

    @_builtins.property
    @pulumi.getter(name="azureSubscriptions")
    def azure_subscriptions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ServiceEndpointKubernetesAzureSubscriptionArgs']]]]:
        """
        An `azure_subscription` block as defined below.
        """
        return pulumi.get(self, "azure_subscriptions")

    @azure_subscriptions.setter
    def azure_subscriptions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ServiceEndpointKubernetesAzureSubscriptionArgs']]]]):
        pulumi.set(self, "azure_subscriptions", value)

    @_builtins.property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[_builtins.str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "description", value)

    @_builtins.property
    @pulumi.getter
    def kubeconfig(self) -> Optional[pulumi.Input['ServiceEndpointKubernetesKubeconfigArgs']]:
        """
        A `kubeconfig` block as defined below.
        """
        return pulumi.get(self, "kubeconfig")

    @kubeconfig.setter
    def kubeconfig(self, value: Optional[pulumi.Input['ServiceEndpointKubernetesKubeconfigArgs']]):
        pulumi.set(self, "kubeconfig", value)

    @_builtins.property
    @pulumi.getter(name="serviceAccount")
    def service_account(self) -> Optional[pulumi.Input['ServiceEndpointKubernetesServiceAccountArgs']]:
        """
        A `service_account` block as defined below.
        """
        return pulumi.get(self, "service_account")

    @service_account.setter
    def service_account(self, value: Optional[pulumi.Input['ServiceEndpointKubernetesServiceAccountArgs']]):
        pulumi.set(self, "service_account", value)


@pulumi.input_type
class _ServiceEndpointKubernetesState:
    def __init__(__self__, *,
                 apiserver_url: Optional[pulumi.Input[_builtins.str]] = None,
                 authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]]] = None,
                 authorization_type: Optional[pulumi.Input[_builtins.str]] = None,
                 azure_subscriptions: Optional[pulumi.Input[Sequence[pulumi.Input['ServiceEndpointKubernetesAzureSubscriptionArgs']]]] = None,
                 description: Optional[pulumi.Input[_builtins.str]] = None,
                 kubeconfig: Optional[pulumi.Input['ServiceEndpointKubernetesKubeconfigArgs']] = None,
                 project_id: Optional[pulumi.Input[_builtins.str]] = None,
                 service_account: Optional[pulumi.Input['ServiceEndpointKubernetesServiceAccountArgs']] = None,
                 service_endpoint_name: Optional[pulumi.Input[_builtins.str]] = None):
        """
        Input properties used for looking up and filtering ServiceEndpointKubernetes resources.
        :param pulumi.Input[_builtins.str] apiserver_url: The hostname (in form of URI) of the Kubernetes API.
        :param pulumi.Input[_builtins.str] authorization_type: The authentication method used to authenticate on the Kubernetes cluster. The value should be one of AzureSubscription, Kubeconfig, ServiceAccount.
        :param pulumi.Input[Sequence[pulumi.Input['ServiceEndpointKubernetesAzureSubscriptionArgs']]] azure_subscriptions: An `azure_subscription` block as defined below.
        :param pulumi.Input['ServiceEndpointKubernetesKubeconfigArgs'] kubeconfig: A `kubeconfig` block as defined below.
        :param pulumi.Input[_builtins.str] project_id: The ID of the project.
        :param pulumi.Input['ServiceEndpointKubernetesServiceAccountArgs'] service_account: A `service_account` block as defined below.
        :param pulumi.Input[_builtins.str] service_endpoint_name: The Service Endpoint name.
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
        if kubeconfig is not None:
            pulumi.set(__self__, "kubeconfig", kubeconfig)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if service_account is not None:
            pulumi.set(__self__, "service_account", service_account)
        if service_endpoint_name is not None:
            pulumi.set(__self__, "service_endpoint_name", service_endpoint_name)

    @_builtins.property
    @pulumi.getter(name="apiserverUrl")
    def apiserver_url(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The hostname (in form of URI) of the Kubernetes API.
        """
        return pulumi.get(self, "apiserver_url")

    @apiserver_url.setter
    def apiserver_url(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "apiserver_url", value)

    @_builtins.property
    @pulumi.getter
    def authorization(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]]]:
        return pulumi.get(self, "authorization")

    @authorization.setter
    def authorization(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]]]):
        pulumi.set(self, "authorization", value)

    @_builtins.property
    @pulumi.getter(name="authorizationType")
    def authorization_type(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The authentication method used to authenticate on the Kubernetes cluster. The value should be one of AzureSubscription, Kubeconfig, ServiceAccount.
        """
        return pulumi.get(self, "authorization_type")

    @authorization_type.setter
    def authorization_type(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "authorization_type", value)

    @_builtins.property
    @pulumi.getter(name="azureSubscriptions")
    def azure_subscriptions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ServiceEndpointKubernetesAzureSubscriptionArgs']]]]:
        """
        An `azure_subscription` block as defined below.
        """
        return pulumi.get(self, "azure_subscriptions")

    @azure_subscriptions.setter
    def azure_subscriptions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ServiceEndpointKubernetesAzureSubscriptionArgs']]]]):
        pulumi.set(self, "azure_subscriptions", value)

    @_builtins.property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[_builtins.str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "description", value)

    @_builtins.property
    @pulumi.getter
    def kubeconfig(self) -> Optional[pulumi.Input['ServiceEndpointKubernetesKubeconfigArgs']]:
        """
        A `kubeconfig` block as defined below.
        """
        return pulumi.get(self, "kubeconfig")

    @kubeconfig.setter
    def kubeconfig(self, value: Optional[pulumi.Input['ServiceEndpointKubernetesKubeconfigArgs']]):
        pulumi.set(self, "kubeconfig", value)

    @_builtins.property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The ID of the project.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "project_id", value)

    @_builtins.property
    @pulumi.getter(name="serviceAccount")
    def service_account(self) -> Optional[pulumi.Input['ServiceEndpointKubernetesServiceAccountArgs']]:
        """
        A `service_account` block as defined below.
        """
        return pulumi.get(self, "service_account")

    @service_account.setter
    def service_account(self, value: Optional[pulumi.Input['ServiceEndpointKubernetesServiceAccountArgs']]):
        pulumi.set(self, "service_account", value)

    @_builtins.property
    @pulumi.getter(name="serviceEndpointName")
    def service_endpoint_name(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The Service Endpoint name.
        """
        return pulumi.get(self, "service_endpoint_name")

    @service_endpoint_name.setter
    def service_endpoint_name(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "service_endpoint_name", value)


@pulumi.type_token("azuredevops:index/serviceEndpointKubernetes:ServiceEndpointKubernetes")
class ServiceEndpointKubernetes(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 apiserver_url: Optional[pulumi.Input[_builtins.str]] = None,
                 authorization_type: Optional[pulumi.Input[_builtins.str]] = None,
                 azure_subscriptions: Optional[pulumi.Input[Sequence[pulumi.Input[Union['ServiceEndpointKubernetesAzureSubscriptionArgs', 'ServiceEndpointKubernetesAzureSubscriptionArgsDict']]]]] = None,
                 description: Optional[pulumi.Input[_builtins.str]] = None,
                 kubeconfig: Optional[pulumi.Input[Union['ServiceEndpointKubernetesKubeconfigArgs', 'ServiceEndpointKubernetesKubeconfigArgsDict']]] = None,
                 project_id: Optional[pulumi.Input[_builtins.str]] = None,
                 service_account: Optional[pulumi.Input[Union['ServiceEndpointKubernetesServiceAccountArgs', 'ServiceEndpointKubernetesServiceAccountArgsDict']]] = None,
                 service_endpoint_name: Optional[pulumi.Input[_builtins.str]] = None,
                 __props__=None):
        """
        Manages a Kubernetes service endpoint within Azure DevOps.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        example = azuredevops.Project("example",
            name="Example Project",
            visibility="private",
            version_control="Git",
            work_item_template="Agile",
            description="Managed by Pulumi")
        example_azure = azuredevops.ServiceEndpointKubernetes("example-azure",
            project_id=example.id,
            service_endpoint_name="Example Kubernetes",
            apiserver_url="https://sample-kubernetes-cluster.hcp.westeurope.azmk8s.io",
            authorization_type="AzureSubscription",
            azure_subscriptions=[{
                "subscription_id": "00000000-0000-0000-0000-000000000000",
                "subscription_name": "Example",
                "tenant_id": "00000000-0000-0000-0000-000000000000",
                "resourcegroup_id": "example-rg",
                "namespace": "default",
                "cluster_name": "example-aks",
            }])
        example_kubeconfig = azuredevops.ServiceEndpointKubernetes("example-kubeconfig",
            project_id=example.id,
            service_endpoint_name="Example Kubernetes",
            apiserver_url="https://sample-kubernetes-cluster.hcp.westeurope.azmk8s.io",
            authorization_type="Kubeconfig",
            kubeconfig={
                "kube_config": \"\"\"                              apiVersion: v1
                                      clusters:
                                      - cluster:
                                          certificate-authority: fake-ca-file
                                          server: https://1.2.3.4
                                        name: development
                                      contexts:
                                      - context:
                                          cluster: development
                                          namespace: frontend
                                          user: developer
                                        name: dev-frontend
                                      current-context: dev-frontend
                                      kind: Config
                                      preferences: {}
                                      users:
                                      - name: developer
                                        user:
                                          client-certificate: fake-cert-file
                                          client-key: fake-key-file
        \"\"\",
                "accept_untrusted_certs": True,
                "cluster_context": "dev-frontend",
            })
        example_service_account = azuredevops.ServiceEndpointKubernetes("example-service-account",
            project_id=example.id,
            service_endpoint_name="Example Kubernetes",
            apiserver_url="https://sample-kubernetes-cluster.hcp.westeurope.azmk8s.io",
            authorization_type="ServiceAccount",
            service_account={
                "token": "000000000000000000000000",
                "ca_cert": "0000000000000000000000000000000",
            })
        ```

        ## Relevant Links

        - [Azure DevOps Service REST API 7.0 - Endpoints](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-7.0)

        ## Import

        Azure DevOps Kubernetes Service Endpoint can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**

        ```sh
        $ pulumi import azuredevops:index/serviceEndpointKubernetes:ServiceEndpointKubernetes example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[_builtins.str] apiserver_url: The hostname (in form of URI) of the Kubernetes API.
        :param pulumi.Input[_builtins.str] authorization_type: The authentication method used to authenticate on the Kubernetes cluster. The value should be one of AzureSubscription, Kubeconfig, ServiceAccount.
        :param pulumi.Input[Sequence[pulumi.Input[Union['ServiceEndpointKubernetesAzureSubscriptionArgs', 'ServiceEndpointKubernetesAzureSubscriptionArgsDict']]]] azure_subscriptions: An `azure_subscription` block as defined below.
        :param pulumi.Input[Union['ServiceEndpointKubernetesKubeconfigArgs', 'ServiceEndpointKubernetesKubeconfigArgsDict']] kubeconfig: A `kubeconfig` block as defined below.
        :param pulumi.Input[_builtins.str] project_id: The ID of the project.
        :param pulumi.Input[Union['ServiceEndpointKubernetesServiceAccountArgs', 'ServiceEndpointKubernetesServiceAccountArgsDict']] service_account: A `service_account` block as defined below.
        :param pulumi.Input[_builtins.str] service_endpoint_name: The Service Endpoint name.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ServiceEndpointKubernetesArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a Kubernetes service endpoint within Azure DevOps.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        example = azuredevops.Project("example",
            name="Example Project",
            visibility="private",
            version_control="Git",
            work_item_template="Agile",
            description="Managed by Pulumi")
        example_azure = azuredevops.ServiceEndpointKubernetes("example-azure",
            project_id=example.id,
            service_endpoint_name="Example Kubernetes",
            apiserver_url="https://sample-kubernetes-cluster.hcp.westeurope.azmk8s.io",
            authorization_type="AzureSubscription",
            azure_subscriptions=[{
                "subscription_id": "00000000-0000-0000-0000-000000000000",
                "subscription_name": "Example",
                "tenant_id": "00000000-0000-0000-0000-000000000000",
                "resourcegroup_id": "example-rg",
                "namespace": "default",
                "cluster_name": "example-aks",
            }])
        example_kubeconfig = azuredevops.ServiceEndpointKubernetes("example-kubeconfig",
            project_id=example.id,
            service_endpoint_name="Example Kubernetes",
            apiserver_url="https://sample-kubernetes-cluster.hcp.westeurope.azmk8s.io",
            authorization_type="Kubeconfig",
            kubeconfig={
                "kube_config": \"\"\"                              apiVersion: v1
                                      clusters:
                                      - cluster:
                                          certificate-authority: fake-ca-file
                                          server: https://1.2.3.4
                                        name: development
                                      contexts:
                                      - context:
                                          cluster: development
                                          namespace: frontend
                                          user: developer
                                        name: dev-frontend
                                      current-context: dev-frontend
                                      kind: Config
                                      preferences: {}
                                      users:
                                      - name: developer
                                        user:
                                          client-certificate: fake-cert-file
                                          client-key: fake-key-file
        \"\"\",
                "accept_untrusted_certs": True,
                "cluster_context": "dev-frontend",
            })
        example_service_account = azuredevops.ServiceEndpointKubernetes("example-service-account",
            project_id=example.id,
            service_endpoint_name="Example Kubernetes",
            apiserver_url="https://sample-kubernetes-cluster.hcp.westeurope.azmk8s.io",
            authorization_type="ServiceAccount",
            service_account={
                "token": "000000000000000000000000",
                "ca_cert": "0000000000000000000000000000000",
            })
        ```

        ## Relevant Links

        - [Azure DevOps Service REST API 7.0 - Endpoints](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-7.0)

        ## Import

        Azure DevOps Kubernetes Service Endpoint can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**

        ```sh
        $ pulumi import azuredevops:index/serviceEndpointKubernetes:ServiceEndpointKubernetes example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
        ```

        :param str resource_name: The name of the resource.
        :param ServiceEndpointKubernetesArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ServiceEndpointKubernetesArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 apiserver_url: Optional[pulumi.Input[_builtins.str]] = None,
                 authorization_type: Optional[pulumi.Input[_builtins.str]] = None,
                 azure_subscriptions: Optional[pulumi.Input[Sequence[pulumi.Input[Union['ServiceEndpointKubernetesAzureSubscriptionArgs', 'ServiceEndpointKubernetesAzureSubscriptionArgsDict']]]]] = None,
                 description: Optional[pulumi.Input[_builtins.str]] = None,
                 kubeconfig: Optional[pulumi.Input[Union['ServiceEndpointKubernetesKubeconfigArgs', 'ServiceEndpointKubernetesKubeconfigArgsDict']]] = None,
                 project_id: Optional[pulumi.Input[_builtins.str]] = None,
                 service_account: Optional[pulumi.Input[Union['ServiceEndpointKubernetesServiceAccountArgs', 'ServiceEndpointKubernetesServiceAccountArgsDict']]] = None,
                 service_endpoint_name: Optional[pulumi.Input[_builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ServiceEndpointKubernetesArgs.__new__(ServiceEndpointKubernetesArgs)

            if apiserver_url is None and not opts.urn:
                raise TypeError("Missing required property 'apiserver_url'")
            __props__.__dict__["apiserver_url"] = apiserver_url
            if authorization_type is None and not opts.urn:
                raise TypeError("Missing required property 'authorization_type'")
            __props__.__dict__["authorization_type"] = authorization_type
            __props__.__dict__["azure_subscriptions"] = azure_subscriptions
            __props__.__dict__["description"] = description
            __props__.__dict__["kubeconfig"] = kubeconfig
            if project_id is None and not opts.urn:
                raise TypeError("Missing required property 'project_id'")
            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["service_account"] = service_account
            if service_endpoint_name is None and not opts.urn:
                raise TypeError("Missing required property 'service_endpoint_name'")
            __props__.__dict__["service_endpoint_name"] = service_endpoint_name
            __props__.__dict__["authorization"] = None
        super(ServiceEndpointKubernetes, __self__).__init__(
            'azuredevops:index/serviceEndpointKubernetes:ServiceEndpointKubernetes',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            apiserver_url: Optional[pulumi.Input[_builtins.str]] = None,
            authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]]] = None,
            authorization_type: Optional[pulumi.Input[_builtins.str]] = None,
            azure_subscriptions: Optional[pulumi.Input[Sequence[pulumi.Input[Union['ServiceEndpointKubernetesAzureSubscriptionArgs', 'ServiceEndpointKubernetesAzureSubscriptionArgsDict']]]]] = None,
            description: Optional[pulumi.Input[_builtins.str]] = None,
            kubeconfig: Optional[pulumi.Input[Union['ServiceEndpointKubernetesKubeconfigArgs', 'ServiceEndpointKubernetesKubeconfigArgsDict']]] = None,
            project_id: Optional[pulumi.Input[_builtins.str]] = None,
            service_account: Optional[pulumi.Input[Union['ServiceEndpointKubernetesServiceAccountArgs', 'ServiceEndpointKubernetesServiceAccountArgsDict']]] = None,
            service_endpoint_name: Optional[pulumi.Input[_builtins.str]] = None) -> 'ServiceEndpointKubernetes':
        """
        Get an existing ServiceEndpointKubernetes resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[_builtins.str] apiserver_url: The hostname (in form of URI) of the Kubernetes API.
        :param pulumi.Input[_builtins.str] authorization_type: The authentication method used to authenticate on the Kubernetes cluster. The value should be one of AzureSubscription, Kubeconfig, ServiceAccount.
        :param pulumi.Input[Sequence[pulumi.Input[Union['ServiceEndpointKubernetesAzureSubscriptionArgs', 'ServiceEndpointKubernetesAzureSubscriptionArgsDict']]]] azure_subscriptions: An `azure_subscription` block as defined below.
        :param pulumi.Input[Union['ServiceEndpointKubernetesKubeconfigArgs', 'ServiceEndpointKubernetesKubeconfigArgsDict']] kubeconfig: A `kubeconfig` block as defined below.
        :param pulumi.Input[_builtins.str] project_id: The ID of the project.
        :param pulumi.Input[Union['ServiceEndpointKubernetesServiceAccountArgs', 'ServiceEndpointKubernetesServiceAccountArgsDict']] service_account: A `service_account` block as defined below.
        :param pulumi.Input[_builtins.str] service_endpoint_name: The Service Endpoint name.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ServiceEndpointKubernetesState.__new__(_ServiceEndpointKubernetesState)

        __props__.__dict__["apiserver_url"] = apiserver_url
        __props__.__dict__["authorization"] = authorization
        __props__.__dict__["authorization_type"] = authorization_type
        __props__.__dict__["azure_subscriptions"] = azure_subscriptions
        __props__.__dict__["description"] = description
        __props__.__dict__["kubeconfig"] = kubeconfig
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["service_account"] = service_account
        __props__.__dict__["service_endpoint_name"] = service_endpoint_name
        return ServiceEndpointKubernetes(resource_name, opts=opts, __props__=__props__)

    @_builtins.property
    @pulumi.getter(name="apiserverUrl")
    def apiserver_url(self) -> pulumi.Output[_builtins.str]:
        """
        The hostname (in form of URI) of the Kubernetes API.
        """
        return pulumi.get(self, "apiserver_url")

    @_builtins.property
    @pulumi.getter
    def authorization(self) -> pulumi.Output[Mapping[str, _builtins.str]]:
        return pulumi.get(self, "authorization")

    @_builtins.property
    @pulumi.getter(name="authorizationType")
    def authorization_type(self) -> pulumi.Output[_builtins.str]:
        """
        The authentication method used to authenticate on the Kubernetes cluster. The value should be one of AzureSubscription, Kubeconfig, ServiceAccount.
        """
        return pulumi.get(self, "authorization_type")

    @_builtins.property
    @pulumi.getter(name="azureSubscriptions")
    def azure_subscriptions(self) -> pulumi.Output[Optional[Sequence['outputs.ServiceEndpointKubernetesAzureSubscription']]]:
        """
        An `azure_subscription` block as defined below.
        """
        return pulumi.get(self, "azure_subscriptions")

    @_builtins.property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[_builtins.str]]:
        return pulumi.get(self, "description")

    @_builtins.property
    @pulumi.getter
    def kubeconfig(self) -> pulumi.Output[Optional['outputs.ServiceEndpointKubernetesKubeconfig']]:
        """
        A `kubeconfig` block as defined below.
        """
        return pulumi.get(self, "kubeconfig")

    @_builtins.property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[_builtins.str]:
        """
        The ID of the project.
        """
        return pulumi.get(self, "project_id")

    @_builtins.property
    @pulumi.getter(name="serviceAccount")
    def service_account(self) -> pulumi.Output[Optional['outputs.ServiceEndpointKubernetesServiceAccount']]:
        """
        A `service_account` block as defined below.
        """
        return pulumi.get(self, "service_account")

    @_builtins.property
    @pulumi.getter(name="serviceEndpointName")
    def service_endpoint_name(self) -> pulumi.Output[_builtins.str]:
        """
        The Service Endpoint name.
        """
        return pulumi.get(self, "service_endpoint_name")

