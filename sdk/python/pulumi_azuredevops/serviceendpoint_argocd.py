# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
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

__all__ = ['ServiceendpointArgocdArgs', 'ServiceendpointArgocd']

@pulumi.input_type
class ServiceendpointArgocdArgs:
    def __init__(__self__, *,
                 project_id: pulumi.Input[str],
                 service_endpoint_name: pulumi.Input[str],
                 url: pulumi.Input[str],
                 authentication_basic: Optional[pulumi.Input['ServiceendpointArgocdAuthenticationBasicArgs']] = None,
                 authentication_token: Optional[pulumi.Input['ServiceendpointArgocdAuthenticationTokenArgs']] = None,
                 authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ServiceendpointArgocd resource.
        :param pulumi.Input[str] project_id: The ID of the project.
        :param pulumi.Input[str] service_endpoint_name: The Service Endpoint name.
        :param pulumi.Input[str] url: URL of the ArgoCD server to connect with.
        :param pulumi.Input['ServiceendpointArgocdAuthenticationBasicArgs'] authentication_basic: An `authentication_basic` block for the ArgoCD as documented below.
               
               > **NOTE:** `authentication_basic` and `authentication_token` conflict with each other, only one is required.
        :param pulumi.Input['ServiceendpointArgocdAuthenticationTokenArgs'] authentication_token: An `authentication_token` block for the ArgoCD as documented below.
        :param pulumi.Input[str] description: The Service Endpoint description.
        """
        pulumi.set(__self__, "project_id", project_id)
        pulumi.set(__self__, "service_endpoint_name", service_endpoint_name)
        pulumi.set(__self__, "url", url)
        if authentication_basic is not None:
            pulumi.set(__self__, "authentication_basic", authentication_basic)
        if authentication_token is not None:
            pulumi.set(__self__, "authentication_token", authentication_token)
        if authorization is not None:
            pulumi.set(__self__, "authorization", authorization)
        if description is not None:
            pulumi.set(__self__, "description", description)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Input[str]:
        """
        The ID of the project.
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
    def url(self) -> pulumi.Input[str]:
        """
        URL of the ArgoCD server to connect with.
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: pulumi.Input[str]):
        pulumi.set(self, "url", value)

    @property
    @pulumi.getter(name="authenticationBasic")
    def authentication_basic(self) -> Optional[pulumi.Input['ServiceendpointArgocdAuthenticationBasicArgs']]:
        """
        An `authentication_basic` block for the ArgoCD as documented below.

        > **NOTE:** `authentication_basic` and `authentication_token` conflict with each other, only one is required.
        """
        return pulumi.get(self, "authentication_basic")

    @authentication_basic.setter
    def authentication_basic(self, value: Optional[pulumi.Input['ServiceendpointArgocdAuthenticationBasicArgs']]):
        pulumi.set(self, "authentication_basic", value)

    @property
    @pulumi.getter(name="authenticationToken")
    def authentication_token(self) -> Optional[pulumi.Input['ServiceendpointArgocdAuthenticationTokenArgs']]:
        """
        An `authentication_token` block for the ArgoCD as documented below.
        """
        return pulumi.get(self, "authentication_token")

    @authentication_token.setter
    def authentication_token(self, value: Optional[pulumi.Input['ServiceendpointArgocdAuthenticationTokenArgs']]):
        pulumi.set(self, "authentication_token", value)

    @property
    @pulumi.getter
    def authorization(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        return pulumi.get(self, "authorization")

    @authorization.setter
    def authorization(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "authorization", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The Service Endpoint description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)


@pulumi.input_type
class _ServiceendpointArgocdState:
    def __init__(__self__, *,
                 authentication_basic: Optional[pulumi.Input['ServiceendpointArgocdAuthenticationBasicArgs']] = None,
                 authentication_token: Optional[pulumi.Input['ServiceendpointArgocdAuthenticationTokenArgs']] = None,
                 authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 service_endpoint_name: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ServiceendpointArgocd resources.
        :param pulumi.Input['ServiceendpointArgocdAuthenticationBasicArgs'] authentication_basic: An `authentication_basic` block for the ArgoCD as documented below.
               
               > **NOTE:** `authentication_basic` and `authentication_token` conflict with each other, only one is required.
        :param pulumi.Input['ServiceendpointArgocdAuthenticationTokenArgs'] authentication_token: An `authentication_token` block for the ArgoCD as documented below.
        :param pulumi.Input[str] description: The Service Endpoint description.
        :param pulumi.Input[str] project_id: The ID of the project.
        :param pulumi.Input[str] service_endpoint_name: The Service Endpoint name.
        :param pulumi.Input[str] url: URL of the ArgoCD server to connect with.
        """
        if authentication_basic is not None:
            pulumi.set(__self__, "authentication_basic", authentication_basic)
        if authentication_token is not None:
            pulumi.set(__self__, "authentication_token", authentication_token)
        if authorization is not None:
            pulumi.set(__self__, "authorization", authorization)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if service_endpoint_name is not None:
            pulumi.set(__self__, "service_endpoint_name", service_endpoint_name)
        if url is not None:
            pulumi.set(__self__, "url", url)

    @property
    @pulumi.getter(name="authenticationBasic")
    def authentication_basic(self) -> Optional[pulumi.Input['ServiceendpointArgocdAuthenticationBasicArgs']]:
        """
        An `authentication_basic` block for the ArgoCD as documented below.

        > **NOTE:** `authentication_basic` and `authentication_token` conflict with each other, only one is required.
        """
        return pulumi.get(self, "authentication_basic")

    @authentication_basic.setter
    def authentication_basic(self, value: Optional[pulumi.Input['ServiceendpointArgocdAuthenticationBasicArgs']]):
        pulumi.set(self, "authentication_basic", value)

    @property
    @pulumi.getter(name="authenticationToken")
    def authentication_token(self) -> Optional[pulumi.Input['ServiceendpointArgocdAuthenticationTokenArgs']]:
        """
        An `authentication_token` block for the ArgoCD as documented below.
        """
        return pulumi.get(self, "authentication_token")

    @authentication_token.setter
    def authentication_token(self, value: Optional[pulumi.Input['ServiceendpointArgocdAuthenticationTokenArgs']]):
        pulumi.set(self, "authentication_token", value)

    @property
    @pulumi.getter
    def authorization(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        return pulumi.get(self, "authorization")

    @authorization.setter
    def authorization(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "authorization", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The Service Endpoint description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the project.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

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

    @property
    @pulumi.getter
    def url(self) -> Optional[pulumi.Input[str]]:
        """
        URL of the ArgoCD server to connect with.
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url", value)


class ServiceendpointArgocd(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authentication_basic: Optional[pulumi.Input[Union['ServiceendpointArgocdAuthenticationBasicArgs', 'ServiceendpointArgocdAuthenticationBasicArgsDict']]] = None,
                 authentication_token: Optional[pulumi.Input[Union['ServiceendpointArgocdAuthenticationTokenArgs', 'ServiceendpointArgocdAuthenticationTokenArgsDict']]] = None,
                 authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 service_endpoint_name: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages a ArgoCD service endpoint within Azure DevOps. Using this service endpoint requires you to first install [Argo CD Extension](https://marketplace.visualstudio.com/items?itemName=scb-tomasmortensen.vsix-argocd).

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        example = azuredevops.Project("example",
            name="Example Project",
            visibility="private",
            version_control="Git",
            work_item_template="Agile")
        example_serviceendpoint_argocd = azuredevops.ServiceendpointArgocd("example",
            project_id=example.id,
            service_endpoint_name="Example ArgoCD",
            description="Managed by Pulumi",
            url="https://argocd.my.com",
            authentication_token={
                "token": "0000000000000000000000000000000000000000",
            })
        ```
        Alternatively a username and password may be used.

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        example = azuredevops.Project("example",
            name="Example Project",
            visibility="private",
            version_control="Git",
            work_item_template="Agile",
            description="Managed by Pulumi")
        example_serviceendpoint_argocd = azuredevops.ServiceendpointArgocd("example",
            project_id=example.id,
            service_endpoint_name="Example ArgoCD",
            description="Managed by Pulumi",
            url="https://argocd.my.com",
            authentication_basic={
                "username": "username",
                "password": "password",
            })
        ```
        ## Relevant Links

        - [Azure DevOps Service Connections](https://docs.microsoft.com/en-us/azure/devops/pipelines/library/service-endpoints?view=azure-devops&tabs=yaml)
        - [ArgoCD Project/User Token](https://argo-cd.readthedocs.io/en/stable/user-guide/commands/argocd_account_generate-token/)
        - [Argo CD Extension](https://marketplace.visualstudio.com/items?itemName=scb-tomasmortensen.vsix-argocd)

        ## Import

        Azure DevOps Service Endpoint ArgoCD can be imported using the **projectID/serviceEndpointID**, e.g.

        ```sh
        $ pulumi import azuredevops:index/serviceendpointArgocd:ServiceendpointArgocd example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Union['ServiceendpointArgocdAuthenticationBasicArgs', 'ServiceendpointArgocdAuthenticationBasicArgsDict']] authentication_basic: An `authentication_basic` block for the ArgoCD as documented below.
               
               > **NOTE:** `authentication_basic` and `authentication_token` conflict with each other, only one is required.
        :param pulumi.Input[Union['ServiceendpointArgocdAuthenticationTokenArgs', 'ServiceendpointArgocdAuthenticationTokenArgsDict']] authentication_token: An `authentication_token` block for the ArgoCD as documented below.
        :param pulumi.Input[str] description: The Service Endpoint description.
        :param pulumi.Input[str] project_id: The ID of the project.
        :param pulumi.Input[str] service_endpoint_name: The Service Endpoint name.
        :param pulumi.Input[str] url: URL of the ArgoCD server to connect with.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ServiceendpointArgocdArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a ArgoCD service endpoint within Azure DevOps. Using this service endpoint requires you to first install [Argo CD Extension](https://marketplace.visualstudio.com/items?itemName=scb-tomasmortensen.vsix-argocd).

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        example = azuredevops.Project("example",
            name="Example Project",
            visibility="private",
            version_control="Git",
            work_item_template="Agile")
        example_serviceendpoint_argocd = azuredevops.ServiceendpointArgocd("example",
            project_id=example.id,
            service_endpoint_name="Example ArgoCD",
            description="Managed by Pulumi",
            url="https://argocd.my.com",
            authentication_token={
                "token": "0000000000000000000000000000000000000000",
            })
        ```
        Alternatively a username and password may be used.

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        example = azuredevops.Project("example",
            name="Example Project",
            visibility="private",
            version_control="Git",
            work_item_template="Agile",
            description="Managed by Pulumi")
        example_serviceendpoint_argocd = azuredevops.ServiceendpointArgocd("example",
            project_id=example.id,
            service_endpoint_name="Example ArgoCD",
            description="Managed by Pulumi",
            url="https://argocd.my.com",
            authentication_basic={
                "username": "username",
                "password": "password",
            })
        ```
        ## Relevant Links

        - [Azure DevOps Service Connections](https://docs.microsoft.com/en-us/azure/devops/pipelines/library/service-endpoints?view=azure-devops&tabs=yaml)
        - [ArgoCD Project/User Token](https://argo-cd.readthedocs.io/en/stable/user-guide/commands/argocd_account_generate-token/)
        - [Argo CD Extension](https://marketplace.visualstudio.com/items?itemName=scb-tomasmortensen.vsix-argocd)

        ## Import

        Azure DevOps Service Endpoint ArgoCD can be imported using the **projectID/serviceEndpointID**, e.g.

        ```sh
        $ pulumi import azuredevops:index/serviceendpointArgocd:ServiceendpointArgocd example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
        ```

        :param str resource_name: The name of the resource.
        :param ServiceendpointArgocdArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ServiceendpointArgocdArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authentication_basic: Optional[pulumi.Input[Union['ServiceendpointArgocdAuthenticationBasicArgs', 'ServiceendpointArgocdAuthenticationBasicArgsDict']]] = None,
                 authentication_token: Optional[pulumi.Input[Union['ServiceendpointArgocdAuthenticationTokenArgs', 'ServiceendpointArgocdAuthenticationTokenArgsDict']]] = None,
                 authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 service_endpoint_name: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ServiceendpointArgocdArgs.__new__(ServiceendpointArgocdArgs)

            __props__.__dict__["authentication_basic"] = authentication_basic
            __props__.__dict__["authentication_token"] = authentication_token
            __props__.__dict__["authorization"] = authorization
            __props__.__dict__["description"] = description
            if project_id is None and not opts.urn:
                raise TypeError("Missing required property 'project_id'")
            __props__.__dict__["project_id"] = project_id
            if service_endpoint_name is None and not opts.urn:
                raise TypeError("Missing required property 'service_endpoint_name'")
            __props__.__dict__["service_endpoint_name"] = service_endpoint_name
            if url is None and not opts.urn:
                raise TypeError("Missing required property 'url'")
            __props__.__dict__["url"] = url
        super(ServiceendpointArgocd, __self__).__init__(
            'azuredevops:index/serviceendpointArgocd:ServiceendpointArgocd',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            authentication_basic: Optional[pulumi.Input[Union['ServiceendpointArgocdAuthenticationBasicArgs', 'ServiceendpointArgocdAuthenticationBasicArgsDict']]] = None,
            authentication_token: Optional[pulumi.Input[Union['ServiceendpointArgocdAuthenticationTokenArgs', 'ServiceendpointArgocdAuthenticationTokenArgsDict']]] = None,
            authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            description: Optional[pulumi.Input[str]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            service_endpoint_name: Optional[pulumi.Input[str]] = None,
            url: Optional[pulumi.Input[str]] = None) -> 'ServiceendpointArgocd':
        """
        Get an existing ServiceendpointArgocd resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Union['ServiceendpointArgocdAuthenticationBasicArgs', 'ServiceendpointArgocdAuthenticationBasicArgsDict']] authentication_basic: An `authentication_basic` block for the ArgoCD as documented below.
               
               > **NOTE:** `authentication_basic` and `authentication_token` conflict with each other, only one is required.
        :param pulumi.Input[Union['ServiceendpointArgocdAuthenticationTokenArgs', 'ServiceendpointArgocdAuthenticationTokenArgsDict']] authentication_token: An `authentication_token` block for the ArgoCD as documented below.
        :param pulumi.Input[str] description: The Service Endpoint description.
        :param pulumi.Input[str] project_id: The ID of the project.
        :param pulumi.Input[str] service_endpoint_name: The Service Endpoint name.
        :param pulumi.Input[str] url: URL of the ArgoCD server to connect with.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ServiceendpointArgocdState.__new__(_ServiceendpointArgocdState)

        __props__.__dict__["authentication_basic"] = authentication_basic
        __props__.__dict__["authentication_token"] = authentication_token
        __props__.__dict__["authorization"] = authorization
        __props__.__dict__["description"] = description
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["service_endpoint_name"] = service_endpoint_name
        __props__.__dict__["url"] = url
        return ServiceendpointArgocd(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="authenticationBasic")
    def authentication_basic(self) -> pulumi.Output[Optional['outputs.ServiceendpointArgocdAuthenticationBasic']]:
        """
        An `authentication_basic` block for the ArgoCD as documented below.

        > **NOTE:** `authentication_basic` and `authentication_token` conflict with each other, only one is required.
        """
        return pulumi.get(self, "authentication_basic")

    @property
    @pulumi.getter(name="authenticationToken")
    def authentication_token(self) -> pulumi.Output[Optional['outputs.ServiceendpointArgocdAuthenticationToken']]:
        """
        An `authentication_token` block for the ArgoCD as documented below.
        """
        return pulumi.get(self, "authentication_token")

    @property
    @pulumi.getter
    def authorization(self) -> pulumi.Output[Mapping[str, str]]:
        return pulumi.get(self, "authorization")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The Service Endpoint description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        """
        The ID of the project.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="serviceEndpointName")
    def service_endpoint_name(self) -> pulumi.Output[str]:
        """
        The Service Endpoint name.
        """
        return pulumi.get(self, "service_endpoint_name")

    @property
    @pulumi.getter
    def url(self) -> pulumi.Output[str]:
        """
        URL of the ArgoCD server to connect with.
        """
        return pulumi.get(self, "url")

