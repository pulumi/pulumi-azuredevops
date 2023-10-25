# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['ServiceendpointJfrogXrayV2Args', 'ServiceendpointJfrogXrayV2']

@pulumi.input_type
class ServiceendpointJfrogXrayV2Args:
    def __init__(__self__, *,
                 project_id: pulumi.Input[str],
                 service_endpoint_name: pulumi.Input[str],
                 url: pulumi.Input[str],
                 authentication_basic: Optional[pulumi.Input['ServiceendpointJfrogXrayV2AuthenticationBasicArgs']] = None,
                 authentication_token: Optional[pulumi.Input['ServiceendpointJfrogXrayV2AuthenticationTokenArgs']] = None,
                 authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ServiceendpointJfrogXrayV2 resource.
        :param pulumi.Input[str] project_id: The ID of the project.
        :param pulumi.Input[str] service_endpoint_name: The Service Endpoint name.
        :param pulumi.Input[str] url: URL of the Artifactory server to connect with.
               
               > **NOTE:** URL should not end in a slash character.
        :param pulumi.Input['ServiceendpointJfrogXrayV2AuthenticationBasicArgs'] authentication_basic: A `authentication_basic` block as documented below.
        :param pulumi.Input['ServiceendpointJfrogXrayV2AuthenticationTokenArgs'] authentication_token: A `authentication_token` block as documented below.
        :param pulumi.Input[str] description: The Service Endpoint description.
        """
        ServiceendpointJfrogXrayV2Args._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            project_id=project_id,
            service_endpoint_name=service_endpoint_name,
            url=url,
            authentication_basic=authentication_basic,
            authentication_token=authentication_token,
            authorization=authorization,
            description=description,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             project_id: Optional[pulumi.Input[str]] = None,
             service_endpoint_name: Optional[pulumi.Input[str]] = None,
             url: Optional[pulumi.Input[str]] = None,
             authentication_basic: Optional[pulumi.Input['ServiceendpointJfrogXrayV2AuthenticationBasicArgs']] = None,
             authentication_token: Optional[pulumi.Input['ServiceendpointJfrogXrayV2AuthenticationTokenArgs']] = None,
             authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
             description: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if project_id is None and 'projectId' in kwargs:
            project_id = kwargs['projectId']
        if project_id is None:
            raise TypeError("Missing 'project_id' argument")
        if service_endpoint_name is None and 'serviceEndpointName' in kwargs:
            service_endpoint_name = kwargs['serviceEndpointName']
        if service_endpoint_name is None:
            raise TypeError("Missing 'service_endpoint_name' argument")
        if url is None:
            raise TypeError("Missing 'url' argument")
        if authentication_basic is None and 'authenticationBasic' in kwargs:
            authentication_basic = kwargs['authenticationBasic']
        if authentication_token is None and 'authenticationToken' in kwargs:
            authentication_token = kwargs['authenticationToken']

        _setter("project_id", project_id)
        _setter("service_endpoint_name", service_endpoint_name)
        _setter("url", url)
        if authentication_basic is not None:
            _setter("authentication_basic", authentication_basic)
        if authentication_token is not None:
            _setter("authentication_token", authentication_token)
        if authorization is not None:
            _setter("authorization", authorization)
        if description is not None:
            _setter("description", description)

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
        URL of the Artifactory server to connect with.

        > **NOTE:** URL should not end in a slash character.
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: pulumi.Input[str]):
        pulumi.set(self, "url", value)

    @property
    @pulumi.getter(name="authenticationBasic")
    def authentication_basic(self) -> Optional[pulumi.Input['ServiceendpointJfrogXrayV2AuthenticationBasicArgs']]:
        """
        A `authentication_basic` block as documented below.
        """
        return pulumi.get(self, "authentication_basic")

    @authentication_basic.setter
    def authentication_basic(self, value: Optional[pulumi.Input['ServiceendpointJfrogXrayV2AuthenticationBasicArgs']]):
        pulumi.set(self, "authentication_basic", value)

    @property
    @pulumi.getter(name="authenticationToken")
    def authentication_token(self) -> Optional[pulumi.Input['ServiceendpointJfrogXrayV2AuthenticationTokenArgs']]:
        """
        A `authentication_token` block as documented below.
        """
        return pulumi.get(self, "authentication_token")

    @authentication_token.setter
    def authentication_token(self, value: Optional[pulumi.Input['ServiceendpointJfrogXrayV2AuthenticationTokenArgs']]):
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
class _ServiceendpointJfrogXrayV2State:
    def __init__(__self__, *,
                 authentication_basic: Optional[pulumi.Input['ServiceendpointJfrogXrayV2AuthenticationBasicArgs']] = None,
                 authentication_token: Optional[pulumi.Input['ServiceendpointJfrogXrayV2AuthenticationTokenArgs']] = None,
                 authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 service_endpoint_name: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ServiceendpointJfrogXrayV2 resources.
        :param pulumi.Input['ServiceendpointJfrogXrayV2AuthenticationBasicArgs'] authentication_basic: A `authentication_basic` block as documented below.
        :param pulumi.Input['ServiceendpointJfrogXrayV2AuthenticationTokenArgs'] authentication_token: A `authentication_token` block as documented below.
        :param pulumi.Input[str] description: The Service Endpoint description.
        :param pulumi.Input[str] project_id: The ID of the project.
        :param pulumi.Input[str] service_endpoint_name: The Service Endpoint name.
        :param pulumi.Input[str] url: URL of the Artifactory server to connect with.
               
               > **NOTE:** URL should not end in a slash character.
        """
        _ServiceendpointJfrogXrayV2State._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            authentication_basic=authentication_basic,
            authentication_token=authentication_token,
            authorization=authorization,
            description=description,
            project_id=project_id,
            service_endpoint_name=service_endpoint_name,
            url=url,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             authentication_basic: Optional[pulumi.Input['ServiceendpointJfrogXrayV2AuthenticationBasicArgs']] = None,
             authentication_token: Optional[pulumi.Input['ServiceendpointJfrogXrayV2AuthenticationTokenArgs']] = None,
             authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
             description: Optional[pulumi.Input[str]] = None,
             project_id: Optional[pulumi.Input[str]] = None,
             service_endpoint_name: Optional[pulumi.Input[str]] = None,
             url: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if authentication_basic is None and 'authenticationBasic' in kwargs:
            authentication_basic = kwargs['authenticationBasic']
        if authentication_token is None and 'authenticationToken' in kwargs:
            authentication_token = kwargs['authenticationToken']
        if project_id is None and 'projectId' in kwargs:
            project_id = kwargs['projectId']
        if service_endpoint_name is None and 'serviceEndpointName' in kwargs:
            service_endpoint_name = kwargs['serviceEndpointName']

        if authentication_basic is not None:
            _setter("authentication_basic", authentication_basic)
        if authentication_token is not None:
            _setter("authentication_token", authentication_token)
        if authorization is not None:
            _setter("authorization", authorization)
        if description is not None:
            _setter("description", description)
        if project_id is not None:
            _setter("project_id", project_id)
        if service_endpoint_name is not None:
            _setter("service_endpoint_name", service_endpoint_name)
        if url is not None:
            _setter("url", url)

    @property
    @pulumi.getter(name="authenticationBasic")
    def authentication_basic(self) -> Optional[pulumi.Input['ServiceendpointJfrogXrayV2AuthenticationBasicArgs']]:
        """
        A `authentication_basic` block as documented below.
        """
        return pulumi.get(self, "authentication_basic")

    @authentication_basic.setter
    def authentication_basic(self, value: Optional[pulumi.Input['ServiceendpointJfrogXrayV2AuthenticationBasicArgs']]):
        pulumi.set(self, "authentication_basic", value)

    @property
    @pulumi.getter(name="authenticationToken")
    def authentication_token(self) -> Optional[pulumi.Input['ServiceendpointJfrogXrayV2AuthenticationTokenArgs']]:
        """
        A `authentication_token` block as documented below.
        """
        return pulumi.get(self, "authentication_token")

    @authentication_token.setter
    def authentication_token(self, value: Optional[pulumi.Input['ServiceendpointJfrogXrayV2AuthenticationTokenArgs']]):
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
        URL of the Artifactory server to connect with.

        > **NOTE:** URL should not end in a slash character.
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url", value)


class ServiceendpointJfrogXrayV2(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authentication_basic: Optional[pulumi.Input[pulumi.InputType['ServiceendpointJfrogXrayV2AuthenticationBasicArgs']]] = None,
                 authentication_token: Optional[pulumi.Input[pulumi.InputType['ServiceendpointJfrogXrayV2AuthenticationTokenArgs']]] = None,
                 authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 service_endpoint_name: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages an JFrog XRay V2 server endpoint within an Azure DevOps organization.

        > **Note:** Using this service endpoint requires you to first install [JFrog Extension](https://marketplace.visualstudio.com/items?itemName=JFrog.jfrog-azure-devops-extension).

        ## Relevant Links

        * [Azure DevOps Service Connections](https://docs.microsoft.com/en-us/azure/devops/pipelines/library/service-endpoints?view=azure-devops&tabs=yaml)
        * [Artifactory User Token](https://docs.artifactory.org/latest/user-guide/user-token/)

        ## Import

        Azure DevOps Service Endpoint JFrog XRay V2 can be imported using the **projectID/serviceEndpointID**, e.g.

        ```sh
         $ pulumi import azuredevops:index/serviceendpointJfrogXrayV2:ServiceendpointJfrogXrayV2 example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['ServiceendpointJfrogXrayV2AuthenticationBasicArgs']] authentication_basic: A `authentication_basic` block as documented below.
        :param pulumi.Input[pulumi.InputType['ServiceendpointJfrogXrayV2AuthenticationTokenArgs']] authentication_token: A `authentication_token` block as documented below.
        :param pulumi.Input[str] description: The Service Endpoint description.
        :param pulumi.Input[str] project_id: The ID of the project.
        :param pulumi.Input[str] service_endpoint_name: The Service Endpoint name.
        :param pulumi.Input[str] url: URL of the Artifactory server to connect with.
               
               > **NOTE:** URL should not end in a slash character.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ServiceendpointJfrogXrayV2Args,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages an JFrog XRay V2 server endpoint within an Azure DevOps organization.

        > **Note:** Using this service endpoint requires you to first install [JFrog Extension](https://marketplace.visualstudio.com/items?itemName=JFrog.jfrog-azure-devops-extension).

        ## Relevant Links

        * [Azure DevOps Service Connections](https://docs.microsoft.com/en-us/azure/devops/pipelines/library/service-endpoints?view=azure-devops&tabs=yaml)
        * [Artifactory User Token](https://docs.artifactory.org/latest/user-guide/user-token/)

        ## Import

        Azure DevOps Service Endpoint JFrog XRay V2 can be imported using the **projectID/serviceEndpointID**, e.g.

        ```sh
         $ pulumi import azuredevops:index/serviceendpointJfrogXrayV2:ServiceendpointJfrogXrayV2 example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
        ```

        :param str resource_name: The name of the resource.
        :param ServiceendpointJfrogXrayV2Args args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ServiceendpointJfrogXrayV2Args, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            ServiceendpointJfrogXrayV2Args._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authentication_basic: Optional[pulumi.Input[pulumi.InputType['ServiceendpointJfrogXrayV2AuthenticationBasicArgs']]] = None,
                 authentication_token: Optional[pulumi.Input[pulumi.InputType['ServiceendpointJfrogXrayV2AuthenticationTokenArgs']]] = None,
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
            __props__ = ServiceendpointJfrogXrayV2Args.__new__(ServiceendpointJfrogXrayV2Args)

            authentication_basic = _utilities.configure(authentication_basic, ServiceendpointJfrogXrayV2AuthenticationBasicArgs, True)
            __props__.__dict__["authentication_basic"] = authentication_basic
            authentication_token = _utilities.configure(authentication_token, ServiceendpointJfrogXrayV2AuthenticationTokenArgs, True)
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
        super(ServiceendpointJfrogXrayV2, __self__).__init__(
            'azuredevops:index/serviceendpointJfrogXrayV2:ServiceendpointJfrogXrayV2',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            authentication_basic: Optional[pulumi.Input[pulumi.InputType['ServiceendpointJfrogXrayV2AuthenticationBasicArgs']]] = None,
            authentication_token: Optional[pulumi.Input[pulumi.InputType['ServiceendpointJfrogXrayV2AuthenticationTokenArgs']]] = None,
            authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            description: Optional[pulumi.Input[str]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            service_endpoint_name: Optional[pulumi.Input[str]] = None,
            url: Optional[pulumi.Input[str]] = None) -> 'ServiceendpointJfrogXrayV2':
        """
        Get an existing ServiceendpointJfrogXrayV2 resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['ServiceendpointJfrogXrayV2AuthenticationBasicArgs']] authentication_basic: A `authentication_basic` block as documented below.
        :param pulumi.Input[pulumi.InputType['ServiceendpointJfrogXrayV2AuthenticationTokenArgs']] authentication_token: A `authentication_token` block as documented below.
        :param pulumi.Input[str] description: The Service Endpoint description.
        :param pulumi.Input[str] project_id: The ID of the project.
        :param pulumi.Input[str] service_endpoint_name: The Service Endpoint name.
        :param pulumi.Input[str] url: URL of the Artifactory server to connect with.
               
               > **NOTE:** URL should not end in a slash character.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ServiceendpointJfrogXrayV2State.__new__(_ServiceendpointJfrogXrayV2State)

        __props__.__dict__["authentication_basic"] = authentication_basic
        __props__.__dict__["authentication_token"] = authentication_token
        __props__.__dict__["authorization"] = authorization
        __props__.__dict__["description"] = description
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["service_endpoint_name"] = service_endpoint_name
        __props__.__dict__["url"] = url
        return ServiceendpointJfrogXrayV2(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="authenticationBasic")
    def authentication_basic(self) -> pulumi.Output[Optional['outputs.ServiceendpointJfrogXrayV2AuthenticationBasic']]:
        """
        A `authentication_basic` block as documented below.
        """
        return pulumi.get(self, "authentication_basic")

    @property
    @pulumi.getter(name="authenticationToken")
    def authentication_token(self) -> pulumi.Output[Optional['outputs.ServiceendpointJfrogXrayV2AuthenticationToken']]:
        """
        A `authentication_token` block as documented below.
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
        URL of the Artifactory server to connect with.

        > **NOTE:** URL should not end in a slash character.
        """
        return pulumi.get(self, "url")

