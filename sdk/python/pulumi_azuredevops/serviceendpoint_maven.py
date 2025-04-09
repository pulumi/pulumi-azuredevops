# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
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

__all__ = ['ServiceendpointMavenArgs', 'ServiceendpointMaven']

@pulumi.input_type
class ServiceendpointMavenArgs:
    def __init__(__self__, *,
                 project_id: pulumi.Input[builtins.str],
                 repository_id: pulumi.Input[builtins.str],
                 service_endpoint_name: pulumi.Input[builtins.str],
                 url: pulumi.Input[builtins.str],
                 authentication_basic: Optional[pulumi.Input['ServiceendpointMavenAuthenticationBasicArgs']] = None,
                 authentication_token: Optional[pulumi.Input['ServiceendpointMavenAuthenticationTokenArgs']] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a ServiceendpointMaven resource.
        :param pulumi.Input[builtins.str] project_id: The ID of the project. Changing this forces a new Service Connection Maven to be created.
        :param pulumi.Input[builtins.str] repository_id: The ID of the server that matches the id element of the `repository/mirror` that Maven tries to connect to.
        :param pulumi.Input[builtins.str] service_endpoint_name: The name of the service endpoint. Changing this forces a new Service Connection Maven to be created.
        :param pulumi.Input[builtins.str] url: The URL of the Maven Repository.
        :param pulumi.Input['ServiceendpointMavenAuthenticationBasicArgs'] authentication_basic: A `authentication_basic` block as documented below.
        :param pulumi.Input['ServiceendpointMavenAuthenticationTokenArgs'] authentication_token: A `authentication_token` block as documented below.
        """
        pulumi.set(__self__, "project_id", project_id)
        pulumi.set(__self__, "repository_id", repository_id)
        pulumi.set(__self__, "service_endpoint_name", service_endpoint_name)
        pulumi.set(__self__, "url", url)
        if authentication_basic is not None:
            pulumi.set(__self__, "authentication_basic", authentication_basic)
        if authentication_token is not None:
            pulumi.set(__self__, "authentication_token", authentication_token)
        if description is not None:
            pulumi.set(__self__, "description", description)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Input[builtins.str]:
        """
        The ID of the project. Changing this forces a new Service Connection Maven to be created.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter(name="repositoryId")
    def repository_id(self) -> pulumi.Input[builtins.str]:
        """
        The ID of the server that matches the id element of the `repository/mirror` that Maven tries to connect to.
        """
        return pulumi.get(self, "repository_id")

    @repository_id.setter
    def repository_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "repository_id", value)

    @property
    @pulumi.getter(name="serviceEndpointName")
    def service_endpoint_name(self) -> pulumi.Input[builtins.str]:
        """
        The name of the service endpoint. Changing this forces a new Service Connection Maven to be created.
        """
        return pulumi.get(self, "service_endpoint_name")

    @service_endpoint_name.setter
    def service_endpoint_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "service_endpoint_name", value)

    @property
    @pulumi.getter
    def url(self) -> pulumi.Input[builtins.str]:
        """
        The URL of the Maven Repository.
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "url", value)

    @property
    @pulumi.getter(name="authenticationBasic")
    def authentication_basic(self) -> Optional[pulumi.Input['ServiceendpointMavenAuthenticationBasicArgs']]:
        """
        A `authentication_basic` block as documented below.
        """
        return pulumi.get(self, "authentication_basic")

    @authentication_basic.setter
    def authentication_basic(self, value: Optional[pulumi.Input['ServiceendpointMavenAuthenticationBasicArgs']]):
        pulumi.set(self, "authentication_basic", value)

    @property
    @pulumi.getter(name="authenticationToken")
    def authentication_token(self) -> Optional[pulumi.Input['ServiceendpointMavenAuthenticationTokenArgs']]:
        """
        A `authentication_token` block as documented below.
        """
        return pulumi.get(self, "authentication_token")

    @authentication_token.setter
    def authentication_token(self, value: Optional[pulumi.Input['ServiceendpointMavenAuthenticationTokenArgs']]):
        pulumi.set(self, "authentication_token", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "description", value)


@pulumi.input_type
class _ServiceendpointMavenState:
    def __init__(__self__, *,
                 authentication_basic: Optional[pulumi.Input['ServiceendpointMavenAuthenticationBasicArgs']] = None,
                 authentication_token: Optional[pulumi.Input['ServiceendpointMavenAuthenticationTokenArgs']] = None,
                 authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 project_id: Optional[pulumi.Input[builtins.str]] = None,
                 repository_id: Optional[pulumi.Input[builtins.str]] = None,
                 service_endpoint_name: Optional[pulumi.Input[builtins.str]] = None,
                 url: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering ServiceendpointMaven resources.
        :param pulumi.Input['ServiceendpointMavenAuthenticationBasicArgs'] authentication_basic: A `authentication_basic` block as documented below.
        :param pulumi.Input['ServiceendpointMavenAuthenticationTokenArgs'] authentication_token: A `authentication_token` block as documented below.
        :param pulumi.Input[builtins.str] project_id: The ID of the project. Changing this forces a new Service Connection Maven to be created.
        :param pulumi.Input[builtins.str] repository_id: The ID of the server that matches the id element of the `repository/mirror` that Maven tries to connect to.
        :param pulumi.Input[builtins.str] service_endpoint_name: The name of the service endpoint. Changing this forces a new Service Connection Maven to be created.
        :param pulumi.Input[builtins.str] url: The URL of the Maven Repository.
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
        if repository_id is not None:
            pulumi.set(__self__, "repository_id", repository_id)
        if service_endpoint_name is not None:
            pulumi.set(__self__, "service_endpoint_name", service_endpoint_name)
        if url is not None:
            pulumi.set(__self__, "url", url)

    @property
    @pulumi.getter(name="authenticationBasic")
    def authentication_basic(self) -> Optional[pulumi.Input['ServiceendpointMavenAuthenticationBasicArgs']]:
        """
        A `authentication_basic` block as documented below.
        """
        return pulumi.get(self, "authentication_basic")

    @authentication_basic.setter
    def authentication_basic(self, value: Optional[pulumi.Input['ServiceendpointMavenAuthenticationBasicArgs']]):
        pulumi.set(self, "authentication_basic", value)

    @property
    @pulumi.getter(name="authenticationToken")
    def authentication_token(self) -> Optional[pulumi.Input['ServiceendpointMavenAuthenticationTokenArgs']]:
        """
        A `authentication_token` block as documented below.
        """
        return pulumi.get(self, "authentication_token")

    @authentication_token.setter
    def authentication_token(self, value: Optional[pulumi.Input['ServiceendpointMavenAuthenticationTokenArgs']]):
        pulumi.set(self, "authentication_token", value)

    @property
    @pulumi.getter
    def authorization(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]:
        return pulumi.get(self, "authorization")

    @authorization.setter
    def authorization(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "authorization", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The ID of the project. Changing this forces a new Service Connection Maven to be created.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter(name="repositoryId")
    def repository_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The ID of the server that matches the id element of the `repository/mirror` that Maven tries to connect to.
        """
        return pulumi.get(self, "repository_id")

    @repository_id.setter
    def repository_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "repository_id", value)

    @property
    @pulumi.getter(name="serviceEndpointName")
    def service_endpoint_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of the service endpoint. Changing this forces a new Service Connection Maven to be created.
        """
        return pulumi.get(self, "service_endpoint_name")

    @service_endpoint_name.setter
    def service_endpoint_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "service_endpoint_name", value)

    @property
    @pulumi.getter
    def url(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The URL of the Maven Repository.
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "url", value)


class ServiceendpointMaven(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authentication_basic: Optional[pulumi.Input[Union['ServiceendpointMavenAuthenticationBasicArgs', 'ServiceendpointMavenAuthenticationBasicArgsDict']]] = None,
                 authentication_token: Optional[pulumi.Input[Union['ServiceendpointMavenAuthenticationTokenArgs', 'ServiceendpointMavenAuthenticationTokenArgsDict']]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 project_id: Optional[pulumi.Input[builtins.str]] = None,
                 repository_id: Optional[pulumi.Input[builtins.str]] = None,
                 service_endpoint_name: Optional[pulumi.Input[builtins.str]] = None,
                 url: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Manages a Maven service endpoint within Azure DevOps, which can be used as a resource in YAML pipelines to connect to a Maven instance.

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
        example_serviceendpoint_maven = azuredevops.ServiceendpointMaven("example",
            project_id=example.id,
            service_endpoint_name="maven-example",
            description="Service Endpoint for 'Maven' (Managed by Terraform)",
            url="https://example.com",
            repository_id="example",
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
        example_serviceendpoint_maven = azuredevops.ServiceendpointMaven("example",
            project_id=example.id,
            service_endpoint_name="maven-example",
            description="Service Endpoint for 'Maven' (Managed by Terraform)",
            url="https://example.com",
            repository_id="example",
            authentication_basic={
                "username": "username",
                "password": "password",
            })
        ```

        ## Import

        Azure DevOps Maven Service Connection can be imported using the `projectId/id` or `projectName/id`, e.g.

        ```sh
        $ pulumi import azuredevops:index/serviceendpointMaven:ServiceendpointMaven example projectName/00000000-0000-0000-0000-000000000000
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Union['ServiceendpointMavenAuthenticationBasicArgs', 'ServiceendpointMavenAuthenticationBasicArgsDict']] authentication_basic: A `authentication_basic` block as documented below.
        :param pulumi.Input[Union['ServiceendpointMavenAuthenticationTokenArgs', 'ServiceendpointMavenAuthenticationTokenArgsDict']] authentication_token: A `authentication_token` block as documented below.
        :param pulumi.Input[builtins.str] project_id: The ID of the project. Changing this forces a new Service Connection Maven to be created.
        :param pulumi.Input[builtins.str] repository_id: The ID of the server that matches the id element of the `repository/mirror` that Maven tries to connect to.
        :param pulumi.Input[builtins.str] service_endpoint_name: The name of the service endpoint. Changing this forces a new Service Connection Maven to be created.
        :param pulumi.Input[builtins.str] url: The URL of the Maven Repository.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ServiceendpointMavenArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a Maven service endpoint within Azure DevOps, which can be used as a resource in YAML pipelines to connect to a Maven instance.

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
        example_serviceendpoint_maven = azuredevops.ServiceendpointMaven("example",
            project_id=example.id,
            service_endpoint_name="maven-example",
            description="Service Endpoint for 'Maven' (Managed by Terraform)",
            url="https://example.com",
            repository_id="example",
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
        example_serviceendpoint_maven = azuredevops.ServiceendpointMaven("example",
            project_id=example.id,
            service_endpoint_name="maven-example",
            description="Service Endpoint for 'Maven' (Managed by Terraform)",
            url="https://example.com",
            repository_id="example",
            authentication_basic={
                "username": "username",
                "password": "password",
            })
        ```

        ## Import

        Azure DevOps Maven Service Connection can be imported using the `projectId/id` or `projectName/id`, e.g.

        ```sh
        $ pulumi import azuredevops:index/serviceendpointMaven:ServiceendpointMaven example projectName/00000000-0000-0000-0000-000000000000
        ```

        :param str resource_name: The name of the resource.
        :param ServiceendpointMavenArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ServiceendpointMavenArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authentication_basic: Optional[pulumi.Input[Union['ServiceendpointMavenAuthenticationBasicArgs', 'ServiceendpointMavenAuthenticationBasicArgsDict']]] = None,
                 authentication_token: Optional[pulumi.Input[Union['ServiceendpointMavenAuthenticationTokenArgs', 'ServiceendpointMavenAuthenticationTokenArgsDict']]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 project_id: Optional[pulumi.Input[builtins.str]] = None,
                 repository_id: Optional[pulumi.Input[builtins.str]] = None,
                 service_endpoint_name: Optional[pulumi.Input[builtins.str]] = None,
                 url: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ServiceendpointMavenArgs.__new__(ServiceendpointMavenArgs)

            __props__.__dict__["authentication_basic"] = authentication_basic
            __props__.__dict__["authentication_token"] = authentication_token
            __props__.__dict__["description"] = description
            if project_id is None and not opts.urn:
                raise TypeError("Missing required property 'project_id'")
            __props__.__dict__["project_id"] = project_id
            if repository_id is None and not opts.urn:
                raise TypeError("Missing required property 'repository_id'")
            __props__.__dict__["repository_id"] = repository_id
            if service_endpoint_name is None and not opts.urn:
                raise TypeError("Missing required property 'service_endpoint_name'")
            __props__.__dict__["service_endpoint_name"] = service_endpoint_name
            if url is None and not opts.urn:
                raise TypeError("Missing required property 'url'")
            __props__.__dict__["url"] = url
            __props__.__dict__["authorization"] = None
        super(ServiceendpointMaven, __self__).__init__(
            'azuredevops:index/serviceendpointMaven:ServiceendpointMaven',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            authentication_basic: Optional[pulumi.Input[Union['ServiceendpointMavenAuthenticationBasicArgs', 'ServiceendpointMavenAuthenticationBasicArgsDict']]] = None,
            authentication_token: Optional[pulumi.Input[Union['ServiceendpointMavenAuthenticationTokenArgs', 'ServiceendpointMavenAuthenticationTokenArgsDict']]] = None,
            authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
            description: Optional[pulumi.Input[builtins.str]] = None,
            project_id: Optional[pulumi.Input[builtins.str]] = None,
            repository_id: Optional[pulumi.Input[builtins.str]] = None,
            service_endpoint_name: Optional[pulumi.Input[builtins.str]] = None,
            url: Optional[pulumi.Input[builtins.str]] = None) -> 'ServiceendpointMaven':
        """
        Get an existing ServiceendpointMaven resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Union['ServiceendpointMavenAuthenticationBasicArgs', 'ServiceendpointMavenAuthenticationBasicArgsDict']] authentication_basic: A `authentication_basic` block as documented below.
        :param pulumi.Input[Union['ServiceendpointMavenAuthenticationTokenArgs', 'ServiceendpointMavenAuthenticationTokenArgsDict']] authentication_token: A `authentication_token` block as documented below.
        :param pulumi.Input[builtins.str] project_id: The ID of the project. Changing this forces a new Service Connection Maven to be created.
        :param pulumi.Input[builtins.str] repository_id: The ID of the server that matches the id element of the `repository/mirror` that Maven tries to connect to.
        :param pulumi.Input[builtins.str] service_endpoint_name: The name of the service endpoint. Changing this forces a new Service Connection Maven to be created.
        :param pulumi.Input[builtins.str] url: The URL of the Maven Repository.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ServiceendpointMavenState.__new__(_ServiceendpointMavenState)

        __props__.__dict__["authentication_basic"] = authentication_basic
        __props__.__dict__["authentication_token"] = authentication_token
        __props__.__dict__["authorization"] = authorization
        __props__.__dict__["description"] = description
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["repository_id"] = repository_id
        __props__.__dict__["service_endpoint_name"] = service_endpoint_name
        __props__.__dict__["url"] = url
        return ServiceendpointMaven(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="authenticationBasic")
    def authentication_basic(self) -> pulumi.Output[Optional['outputs.ServiceendpointMavenAuthenticationBasic']]:
        """
        A `authentication_basic` block as documented below.
        """
        return pulumi.get(self, "authentication_basic")

    @property
    @pulumi.getter(name="authenticationToken")
    def authentication_token(self) -> pulumi.Output[Optional['outputs.ServiceendpointMavenAuthenticationToken']]:
        """
        A `authentication_token` block as documented below.
        """
        return pulumi.get(self, "authentication_token")

    @property
    @pulumi.getter
    def authorization(self) -> pulumi.Output[Mapping[str, builtins.str]]:
        return pulumi.get(self, "authorization")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[builtins.str]]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[builtins.str]:
        """
        The ID of the project. Changing this forces a new Service Connection Maven to be created.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="repositoryId")
    def repository_id(self) -> pulumi.Output[builtins.str]:
        """
        The ID of the server that matches the id element of the `repository/mirror` that Maven tries to connect to.
        """
        return pulumi.get(self, "repository_id")

    @property
    @pulumi.getter(name="serviceEndpointName")
    def service_endpoint_name(self) -> pulumi.Output[builtins.str]:
        """
        The name of the service endpoint. Changing this forces a new Service Connection Maven to be created.
        """
        return pulumi.get(self, "service_endpoint_name")

    @property
    @pulumi.getter
    def url(self) -> pulumi.Output[builtins.str]:
        """
        The URL of the Maven Repository.
        """
        return pulumi.get(self, "url")

