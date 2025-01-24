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

__all__ = ['ServiceendpointJfrogArtifactoryV2Args', 'ServiceendpointJfrogArtifactoryV2']

@pulumi.input_type
class ServiceendpointJfrogArtifactoryV2Args:
    def __init__(__self__, *,
                 project_id: pulumi.Input[str],
                 service_endpoint_name: pulumi.Input[str],
                 url: pulumi.Input[str],
                 authentication_basic: Optional[pulumi.Input['ServiceendpointJfrogArtifactoryV2AuthenticationBasicArgs']] = None,
                 authentication_token: Optional[pulumi.Input['ServiceendpointJfrogArtifactoryV2AuthenticationTokenArgs']] = None,
                 authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ServiceendpointJfrogArtifactoryV2 resource.
        :param pulumi.Input[str] project_id: The ID of the project.
        :param pulumi.Input[str] service_endpoint_name: The Service Endpoint name.
        :param pulumi.Input[str] url: URL of the Artifactory server to connect with.
               
               > **NOTE:** URL should not end in a slash character.
        :param pulumi.Input['ServiceendpointJfrogArtifactoryV2AuthenticationBasicArgs'] authentication_basic: An `authentication_basic` block as documented below.
        :param pulumi.Input['ServiceendpointJfrogArtifactoryV2AuthenticationTokenArgs'] authentication_token: An `authentication_token` block as documented below.
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
        URL of the Artifactory server to connect with.

        > **NOTE:** URL should not end in a slash character.
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: pulumi.Input[str]):
        pulumi.set(self, "url", value)

    @property
    @pulumi.getter(name="authenticationBasic")
    def authentication_basic(self) -> Optional[pulumi.Input['ServiceendpointJfrogArtifactoryV2AuthenticationBasicArgs']]:
        """
        An `authentication_basic` block as documented below.
        """
        return pulumi.get(self, "authentication_basic")

    @authentication_basic.setter
    def authentication_basic(self, value: Optional[pulumi.Input['ServiceendpointJfrogArtifactoryV2AuthenticationBasicArgs']]):
        pulumi.set(self, "authentication_basic", value)

    @property
    @pulumi.getter(name="authenticationToken")
    def authentication_token(self) -> Optional[pulumi.Input['ServiceendpointJfrogArtifactoryV2AuthenticationTokenArgs']]:
        """
        An `authentication_token` block as documented below.
        """
        return pulumi.get(self, "authentication_token")

    @authentication_token.setter
    def authentication_token(self, value: Optional[pulumi.Input['ServiceendpointJfrogArtifactoryV2AuthenticationTokenArgs']]):
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
class _ServiceendpointJfrogArtifactoryV2State:
    def __init__(__self__, *,
                 authentication_basic: Optional[pulumi.Input['ServiceendpointJfrogArtifactoryV2AuthenticationBasicArgs']] = None,
                 authentication_token: Optional[pulumi.Input['ServiceendpointJfrogArtifactoryV2AuthenticationTokenArgs']] = None,
                 authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 service_endpoint_name: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ServiceendpointJfrogArtifactoryV2 resources.
        :param pulumi.Input['ServiceendpointJfrogArtifactoryV2AuthenticationBasicArgs'] authentication_basic: An `authentication_basic` block as documented below.
        :param pulumi.Input['ServiceendpointJfrogArtifactoryV2AuthenticationTokenArgs'] authentication_token: An `authentication_token` block as documented below.
        :param pulumi.Input[str] description: The Service Endpoint description.
        :param pulumi.Input[str] project_id: The ID of the project.
        :param pulumi.Input[str] service_endpoint_name: The Service Endpoint name.
        :param pulumi.Input[str] url: URL of the Artifactory server to connect with.
               
               > **NOTE:** URL should not end in a slash character.
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
    def authentication_basic(self) -> Optional[pulumi.Input['ServiceendpointJfrogArtifactoryV2AuthenticationBasicArgs']]:
        """
        An `authentication_basic` block as documented below.
        """
        return pulumi.get(self, "authentication_basic")

    @authentication_basic.setter
    def authentication_basic(self, value: Optional[pulumi.Input['ServiceendpointJfrogArtifactoryV2AuthenticationBasicArgs']]):
        pulumi.set(self, "authentication_basic", value)

    @property
    @pulumi.getter(name="authenticationToken")
    def authentication_token(self) -> Optional[pulumi.Input['ServiceendpointJfrogArtifactoryV2AuthenticationTokenArgs']]:
        """
        An `authentication_token` block as documented below.
        """
        return pulumi.get(self, "authentication_token")

    @authentication_token.setter
    def authentication_token(self, value: Optional[pulumi.Input['ServiceendpointJfrogArtifactoryV2AuthenticationTokenArgs']]):
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


class ServiceendpointJfrogArtifactoryV2(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authentication_basic: Optional[pulumi.Input[Union['ServiceendpointJfrogArtifactoryV2AuthenticationBasicArgs', 'ServiceendpointJfrogArtifactoryV2AuthenticationBasicArgsDict']]] = None,
                 authentication_token: Optional[pulumi.Input[Union['ServiceendpointJfrogArtifactoryV2AuthenticationTokenArgs', 'ServiceendpointJfrogArtifactoryV2AuthenticationTokenArgsDict']]] = None,
                 authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 service_endpoint_name: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages a JFrog Artifactory V2 service endpoint within an Azure DevOps organization.

        > **Note:** Using this service endpoint requires you to first install [JFrog Extension](https://marketplace.visualstudio.com/items?itemName=JFrog.jfrog-azure-devops-extension).

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
        example_serviceendpoint_jfrog_artifactory_v2 = azuredevops.ServiceendpointJfrogArtifactoryV2("example",
            project_id=example.id,
            service_endpoint_name="Example JFrog Artifactory V2",
            description="Managed by Pulumi",
            url="https://artifactory.my.com",
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
        example_serviceendpoint_jfrog_artifactory_v2 = azuredevops.ServiceendpointJfrogArtifactoryV2("example",
            project_id=example.id,
            service_endpoint_name="Example JFrog Artifactory V2",
            description="Managed by Pulumi",
            url="https://artifactory.my.com",
            authentication_basic={
                "username": "username",
                "password": "password",
            })
        ```

        ## Relevant Links

        * [Azure DevOps Service Connections](https://docs.microsoft.com/en-us/azure/devops/pipelines/library/service-endpoints?view=azure-devops&tabs=yaml)
        * [Artifactory User Token](https://docs.artifactory.org/latest/user-guide/user-token/)

        ## Import

        Azure DevOps JFrog Artifactory V2 Service Endpoint can be imported using the **projectID/serviceEndpointID**, e.g.

        ```sh
        $ pulumi import azuredevops:index/serviceendpointJfrogArtifactoryV2:ServiceendpointJfrogArtifactoryV2 example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Union['ServiceendpointJfrogArtifactoryV2AuthenticationBasicArgs', 'ServiceendpointJfrogArtifactoryV2AuthenticationBasicArgsDict']] authentication_basic: An `authentication_basic` block as documented below.
        :param pulumi.Input[Union['ServiceendpointJfrogArtifactoryV2AuthenticationTokenArgs', 'ServiceendpointJfrogArtifactoryV2AuthenticationTokenArgsDict']] authentication_token: An `authentication_token` block as documented below.
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
                 args: ServiceendpointJfrogArtifactoryV2Args,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a JFrog Artifactory V2 service endpoint within an Azure DevOps organization.

        > **Note:** Using this service endpoint requires you to first install [JFrog Extension](https://marketplace.visualstudio.com/items?itemName=JFrog.jfrog-azure-devops-extension).

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
        example_serviceendpoint_jfrog_artifactory_v2 = azuredevops.ServiceendpointJfrogArtifactoryV2("example",
            project_id=example.id,
            service_endpoint_name="Example JFrog Artifactory V2",
            description="Managed by Pulumi",
            url="https://artifactory.my.com",
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
        example_serviceendpoint_jfrog_artifactory_v2 = azuredevops.ServiceendpointJfrogArtifactoryV2("example",
            project_id=example.id,
            service_endpoint_name="Example JFrog Artifactory V2",
            description="Managed by Pulumi",
            url="https://artifactory.my.com",
            authentication_basic={
                "username": "username",
                "password": "password",
            })
        ```

        ## Relevant Links

        * [Azure DevOps Service Connections](https://docs.microsoft.com/en-us/azure/devops/pipelines/library/service-endpoints?view=azure-devops&tabs=yaml)
        * [Artifactory User Token](https://docs.artifactory.org/latest/user-guide/user-token/)

        ## Import

        Azure DevOps JFrog Artifactory V2 Service Endpoint can be imported using the **projectID/serviceEndpointID**, e.g.

        ```sh
        $ pulumi import azuredevops:index/serviceendpointJfrogArtifactoryV2:ServiceendpointJfrogArtifactoryV2 example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
        ```

        :param str resource_name: The name of the resource.
        :param ServiceendpointJfrogArtifactoryV2Args args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ServiceendpointJfrogArtifactoryV2Args, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authentication_basic: Optional[pulumi.Input[Union['ServiceendpointJfrogArtifactoryV2AuthenticationBasicArgs', 'ServiceendpointJfrogArtifactoryV2AuthenticationBasicArgsDict']]] = None,
                 authentication_token: Optional[pulumi.Input[Union['ServiceendpointJfrogArtifactoryV2AuthenticationTokenArgs', 'ServiceendpointJfrogArtifactoryV2AuthenticationTokenArgsDict']]] = None,
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
            __props__ = ServiceendpointJfrogArtifactoryV2Args.__new__(ServiceendpointJfrogArtifactoryV2Args)

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
        super(ServiceendpointJfrogArtifactoryV2, __self__).__init__(
            'azuredevops:index/serviceendpointJfrogArtifactoryV2:ServiceendpointJfrogArtifactoryV2',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            authentication_basic: Optional[pulumi.Input[Union['ServiceendpointJfrogArtifactoryV2AuthenticationBasicArgs', 'ServiceendpointJfrogArtifactoryV2AuthenticationBasicArgsDict']]] = None,
            authentication_token: Optional[pulumi.Input[Union['ServiceendpointJfrogArtifactoryV2AuthenticationTokenArgs', 'ServiceendpointJfrogArtifactoryV2AuthenticationTokenArgsDict']]] = None,
            authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            description: Optional[pulumi.Input[str]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            service_endpoint_name: Optional[pulumi.Input[str]] = None,
            url: Optional[pulumi.Input[str]] = None) -> 'ServiceendpointJfrogArtifactoryV2':
        """
        Get an existing ServiceendpointJfrogArtifactoryV2 resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Union['ServiceendpointJfrogArtifactoryV2AuthenticationBasicArgs', 'ServiceendpointJfrogArtifactoryV2AuthenticationBasicArgsDict']] authentication_basic: An `authentication_basic` block as documented below.
        :param pulumi.Input[Union['ServiceendpointJfrogArtifactoryV2AuthenticationTokenArgs', 'ServiceendpointJfrogArtifactoryV2AuthenticationTokenArgsDict']] authentication_token: An `authentication_token` block as documented below.
        :param pulumi.Input[str] description: The Service Endpoint description.
        :param pulumi.Input[str] project_id: The ID of the project.
        :param pulumi.Input[str] service_endpoint_name: The Service Endpoint name.
        :param pulumi.Input[str] url: URL of the Artifactory server to connect with.
               
               > **NOTE:** URL should not end in a slash character.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ServiceendpointJfrogArtifactoryV2State.__new__(_ServiceendpointJfrogArtifactoryV2State)

        __props__.__dict__["authentication_basic"] = authentication_basic
        __props__.__dict__["authentication_token"] = authentication_token
        __props__.__dict__["authorization"] = authorization
        __props__.__dict__["description"] = description
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["service_endpoint_name"] = service_endpoint_name
        __props__.__dict__["url"] = url
        return ServiceendpointJfrogArtifactoryV2(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="authenticationBasic")
    def authentication_basic(self) -> pulumi.Output[Optional['outputs.ServiceendpointJfrogArtifactoryV2AuthenticationBasic']]:
        """
        An `authentication_basic` block as documented below.
        """
        return pulumi.get(self, "authentication_basic")

    @property
    @pulumi.getter(name="authenticationToken")
    def authentication_token(self) -> pulumi.Output[Optional['outputs.ServiceendpointJfrogArtifactoryV2AuthenticationToken']]:
        """
        An `authentication_token` block as documented below.
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

