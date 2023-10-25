# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['GitHubArgs', 'GitHub']

@pulumi.input_type
class GitHubArgs:
    def __init__(__self__, *,
                 project_id: pulumi.Input[str],
                 service_endpoint_name: pulumi.Input[str],
                 auth_oauth: Optional[pulumi.Input['GitHubAuthOauthArgs']] = None,
                 auth_personal: Optional[pulumi.Input['GitHubAuthPersonalArgs']] = None,
                 authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a GitHub resource.
        :param pulumi.Input[str] project_id: The ID of the project.
        :param pulumi.Input[str] service_endpoint_name: The Service Endpoint name.
        :param pulumi.Input['GitHubAuthPersonalArgs'] auth_personal: An `auth_personal` block as documented below. Allows connecting using a personal access token.
        """
        GitHubArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            project_id=project_id,
            service_endpoint_name=service_endpoint_name,
            auth_oauth=auth_oauth,
            auth_personal=auth_personal,
            authorization=authorization,
            description=description,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             project_id: Optional[pulumi.Input[str]] = None,
             service_endpoint_name: Optional[pulumi.Input[str]] = None,
             auth_oauth: Optional[pulumi.Input['GitHubAuthOauthArgs']] = None,
             auth_personal: Optional[pulumi.Input['GitHubAuthPersonalArgs']] = None,
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
        if auth_oauth is None and 'authOauth' in kwargs:
            auth_oauth = kwargs['authOauth']
        if auth_personal is None and 'authPersonal' in kwargs:
            auth_personal = kwargs['authPersonal']

        _setter("project_id", project_id)
        _setter("service_endpoint_name", service_endpoint_name)
        if auth_oauth is not None:
            _setter("auth_oauth", auth_oauth)
        if auth_personal is not None:
            _setter("auth_personal", auth_personal)
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
    @pulumi.getter(name="authOauth")
    def auth_oauth(self) -> Optional[pulumi.Input['GitHubAuthOauthArgs']]:
        return pulumi.get(self, "auth_oauth")

    @auth_oauth.setter
    def auth_oauth(self, value: Optional[pulumi.Input['GitHubAuthOauthArgs']]):
        pulumi.set(self, "auth_oauth", value)

    @property
    @pulumi.getter(name="authPersonal")
    def auth_personal(self) -> Optional[pulumi.Input['GitHubAuthPersonalArgs']]:
        """
        An `auth_personal` block as documented below. Allows connecting using a personal access token.
        """
        return pulumi.get(self, "auth_personal")

    @auth_personal.setter
    def auth_personal(self, value: Optional[pulumi.Input['GitHubAuthPersonalArgs']]):
        pulumi.set(self, "auth_personal", value)

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
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)


@pulumi.input_type
class _GitHubState:
    def __init__(__self__, *,
                 auth_oauth: Optional[pulumi.Input['GitHubAuthOauthArgs']] = None,
                 auth_personal: Optional[pulumi.Input['GitHubAuthPersonalArgs']] = None,
                 authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 service_endpoint_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering GitHub resources.
        :param pulumi.Input['GitHubAuthPersonalArgs'] auth_personal: An `auth_personal` block as documented below. Allows connecting using a personal access token.
        :param pulumi.Input[str] project_id: The ID of the project.
        :param pulumi.Input[str] service_endpoint_name: The Service Endpoint name.
        """
        _GitHubState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            auth_oauth=auth_oauth,
            auth_personal=auth_personal,
            authorization=authorization,
            description=description,
            project_id=project_id,
            service_endpoint_name=service_endpoint_name,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             auth_oauth: Optional[pulumi.Input['GitHubAuthOauthArgs']] = None,
             auth_personal: Optional[pulumi.Input['GitHubAuthPersonalArgs']] = None,
             authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
             description: Optional[pulumi.Input[str]] = None,
             project_id: Optional[pulumi.Input[str]] = None,
             service_endpoint_name: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if auth_oauth is None and 'authOauth' in kwargs:
            auth_oauth = kwargs['authOauth']
        if auth_personal is None and 'authPersonal' in kwargs:
            auth_personal = kwargs['authPersonal']
        if project_id is None and 'projectId' in kwargs:
            project_id = kwargs['projectId']
        if service_endpoint_name is None and 'serviceEndpointName' in kwargs:
            service_endpoint_name = kwargs['serviceEndpointName']

        if auth_oauth is not None:
            _setter("auth_oauth", auth_oauth)
        if auth_personal is not None:
            _setter("auth_personal", auth_personal)
        if authorization is not None:
            _setter("authorization", authorization)
        if description is not None:
            _setter("description", description)
        if project_id is not None:
            _setter("project_id", project_id)
        if service_endpoint_name is not None:
            _setter("service_endpoint_name", service_endpoint_name)

    @property
    @pulumi.getter(name="authOauth")
    def auth_oauth(self) -> Optional[pulumi.Input['GitHubAuthOauthArgs']]:
        return pulumi.get(self, "auth_oauth")

    @auth_oauth.setter
    def auth_oauth(self, value: Optional[pulumi.Input['GitHubAuthOauthArgs']]):
        pulumi.set(self, "auth_oauth", value)

    @property
    @pulumi.getter(name="authPersonal")
    def auth_personal(self) -> Optional[pulumi.Input['GitHubAuthPersonalArgs']]:
        """
        An `auth_personal` block as documented below. Allows connecting using a personal access token.
        """
        return pulumi.get(self, "auth_personal")

    @auth_personal.setter
    def auth_personal(self, value: Optional[pulumi.Input['GitHubAuthPersonalArgs']]):
        pulumi.set(self, "auth_personal", value)

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


warnings.warn("""azuredevops.serviceendpoint.GitHub has been deprecated in favor of azuredevops.ServiceEndpointGitHub""", DeprecationWarning)


class GitHub(pulumi.CustomResource):
    warnings.warn("""azuredevops.serviceendpoint.GitHub has been deprecated in favor of azuredevops.ServiceEndpointGitHub""", DeprecationWarning)

    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auth_oauth: Optional[pulumi.Input[pulumi.InputType['GitHubAuthOauthArgs']]] = None,
                 auth_personal: Optional[pulumi.Input[pulumi.InputType['GitHubAuthPersonalArgs']]] = None,
                 authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 service_endpoint_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages a GitHub service endpoint within Azure DevOps.

        ## Relevant Links

        - [Azure DevOps Service REST API 7.0 - Service Endpoints](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-7.0)

        ## Import

        Azure DevOps Service Endpoint GitHub can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**

        ```sh
         $ pulumi import azuredevops:ServiceEndpoint/gitHub:GitHub example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['GitHubAuthPersonalArgs']] auth_personal: An `auth_personal` block as documented below. Allows connecting using a personal access token.
        :param pulumi.Input[str] project_id: The ID of the project.
        :param pulumi.Input[str] service_endpoint_name: The Service Endpoint name.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: GitHubArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a GitHub service endpoint within Azure DevOps.

        ## Relevant Links

        - [Azure DevOps Service REST API 7.0 - Service Endpoints](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-7.0)

        ## Import

        Azure DevOps Service Endpoint GitHub can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**

        ```sh
         $ pulumi import azuredevops:ServiceEndpoint/gitHub:GitHub example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
        ```

        :param str resource_name: The name of the resource.
        :param GitHubArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GitHubArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            GitHubArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auth_oauth: Optional[pulumi.Input[pulumi.InputType['GitHubAuthOauthArgs']]] = None,
                 auth_personal: Optional[pulumi.Input[pulumi.InputType['GitHubAuthPersonalArgs']]] = None,
                 authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 service_endpoint_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        pulumi.log.warn("""GitHub is deprecated: azuredevops.serviceendpoint.GitHub has been deprecated in favor of azuredevops.ServiceEndpointGitHub""")
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = GitHubArgs.__new__(GitHubArgs)

            auth_oauth = _utilities.configure(auth_oauth, GitHubAuthOauthArgs, True)
            __props__.__dict__["auth_oauth"] = auth_oauth
            auth_personal = _utilities.configure(auth_personal, GitHubAuthPersonalArgs, True)
            __props__.__dict__["auth_personal"] = auth_personal
            __props__.__dict__["authorization"] = authorization
            __props__.__dict__["description"] = description
            if project_id is None and not opts.urn:
                raise TypeError("Missing required property 'project_id'")
            __props__.__dict__["project_id"] = project_id
            if service_endpoint_name is None and not opts.urn:
                raise TypeError("Missing required property 'service_endpoint_name'")
            __props__.__dict__["service_endpoint_name"] = service_endpoint_name
        super(GitHub, __self__).__init__(
            'azuredevops:ServiceEndpoint/gitHub:GitHub',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            auth_oauth: Optional[pulumi.Input[pulumi.InputType['GitHubAuthOauthArgs']]] = None,
            auth_personal: Optional[pulumi.Input[pulumi.InputType['GitHubAuthPersonalArgs']]] = None,
            authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            description: Optional[pulumi.Input[str]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            service_endpoint_name: Optional[pulumi.Input[str]] = None) -> 'GitHub':
        """
        Get an existing GitHub resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['GitHubAuthPersonalArgs']] auth_personal: An `auth_personal` block as documented below. Allows connecting using a personal access token.
        :param pulumi.Input[str] project_id: The ID of the project.
        :param pulumi.Input[str] service_endpoint_name: The Service Endpoint name.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _GitHubState.__new__(_GitHubState)

        __props__.__dict__["auth_oauth"] = auth_oauth
        __props__.__dict__["auth_personal"] = auth_personal
        __props__.__dict__["authorization"] = authorization
        __props__.__dict__["description"] = description
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["service_endpoint_name"] = service_endpoint_name
        return GitHub(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="authOauth")
    def auth_oauth(self) -> pulumi.Output[Optional['outputs.GitHubAuthOauth']]:
        return pulumi.get(self, "auth_oauth")

    @property
    @pulumi.getter(name="authPersonal")
    def auth_personal(self) -> pulumi.Output[Optional['outputs.GitHubAuthPersonal']]:
        """
        An `auth_personal` block as documented below. Allows connecting using a personal access token.
        """
        return pulumi.get(self, "auth_personal")

    @property
    @pulumi.getter
    def authorization(self) -> pulumi.Output[Mapping[str, str]]:
        return pulumi.get(self, "authorization")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
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

