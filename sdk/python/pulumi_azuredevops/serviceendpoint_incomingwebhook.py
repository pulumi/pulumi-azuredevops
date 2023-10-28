# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ServiceendpointIncomingwebhookArgs', 'ServiceendpointIncomingwebhook']

@pulumi.input_type
class ServiceendpointIncomingwebhookArgs:
    def __init__(__self__, *,
                 project_id: pulumi.Input[str],
                 service_endpoint_name: pulumi.Input[str],
                 webhook_name: pulumi.Input[str],
                 authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 http_header: Optional[pulumi.Input[str]] = None,
                 secret: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ServiceendpointIncomingwebhook resource.
        :param pulumi.Input[str] project_id: The ID of the project. Changing this forces a new Service Connection Incoming WebHook to be created.
        :param pulumi.Input[str] service_endpoint_name: The name of the service endpoint. Changing this forces a new Service Connection Incoming WebHook to be created.
        :param pulumi.Input[str] webhook_name: The name of the WebHook.
        :param pulumi.Input[str] http_header: Http header name on which checksum will be sent.
        :param pulumi.Input[str] secret: Secret for the WebHook. WebHook service will use this secret to calculate the payload checksum.
        """
        pulumi.set(__self__, "project_id", project_id)
        pulumi.set(__self__, "service_endpoint_name", service_endpoint_name)
        pulumi.set(__self__, "webhook_name", webhook_name)
        if authorization is not None:
            pulumi.set(__self__, "authorization", authorization)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if http_header is not None:
            pulumi.set(__self__, "http_header", http_header)
        if secret is not None:
            pulumi.set(__self__, "secret", secret)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Input[str]:
        """
        The ID of the project. Changing this forces a new Service Connection Incoming WebHook to be created.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter(name="serviceEndpointName")
    def service_endpoint_name(self) -> pulumi.Input[str]:
        """
        The name of the service endpoint. Changing this forces a new Service Connection Incoming WebHook to be created.
        """
        return pulumi.get(self, "service_endpoint_name")

    @service_endpoint_name.setter
    def service_endpoint_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "service_endpoint_name", value)

    @property
    @pulumi.getter(name="webhookName")
    def webhook_name(self) -> pulumi.Input[str]:
        """
        The name of the WebHook.
        """
        return pulumi.get(self, "webhook_name")

    @webhook_name.setter
    def webhook_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "webhook_name", value)

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
    @pulumi.getter(name="httpHeader")
    def http_header(self) -> Optional[pulumi.Input[str]]:
        """
        Http header name on which checksum will be sent.
        """
        return pulumi.get(self, "http_header")

    @http_header.setter
    def http_header(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "http_header", value)

    @property
    @pulumi.getter
    def secret(self) -> Optional[pulumi.Input[str]]:
        """
        Secret for the WebHook. WebHook service will use this secret to calculate the payload checksum.
        """
        return pulumi.get(self, "secret")

    @secret.setter
    def secret(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secret", value)


@pulumi.input_type
class _ServiceendpointIncomingwebhookState:
    def __init__(__self__, *,
                 authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 http_header: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 secret: Optional[pulumi.Input[str]] = None,
                 service_endpoint_name: Optional[pulumi.Input[str]] = None,
                 webhook_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ServiceendpointIncomingwebhook resources.
        :param pulumi.Input[str] http_header: Http header name on which checksum will be sent.
        :param pulumi.Input[str] project_id: The ID of the project. Changing this forces a new Service Connection Incoming WebHook to be created.
        :param pulumi.Input[str] secret: Secret for the WebHook. WebHook service will use this secret to calculate the payload checksum.
        :param pulumi.Input[str] service_endpoint_name: The name of the service endpoint. Changing this forces a new Service Connection Incoming WebHook to be created.
        :param pulumi.Input[str] webhook_name: The name of the WebHook.
        """
        if authorization is not None:
            pulumi.set(__self__, "authorization", authorization)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if http_header is not None:
            pulumi.set(__self__, "http_header", http_header)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if secret is not None:
            pulumi.set(__self__, "secret", secret)
        if service_endpoint_name is not None:
            pulumi.set(__self__, "service_endpoint_name", service_endpoint_name)
        if webhook_name is not None:
            pulumi.set(__self__, "webhook_name", webhook_name)

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
    @pulumi.getter(name="httpHeader")
    def http_header(self) -> Optional[pulumi.Input[str]]:
        """
        Http header name on which checksum will be sent.
        """
        return pulumi.get(self, "http_header")

    @http_header.setter
    def http_header(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "http_header", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the project. Changing this forces a new Service Connection Incoming WebHook to be created.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def secret(self) -> Optional[pulumi.Input[str]]:
        """
        Secret for the WebHook. WebHook service will use this secret to calculate the payload checksum.
        """
        return pulumi.get(self, "secret")

    @secret.setter
    def secret(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secret", value)

    @property
    @pulumi.getter(name="serviceEndpointName")
    def service_endpoint_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the service endpoint. Changing this forces a new Service Connection Incoming WebHook to be created.
        """
        return pulumi.get(self, "service_endpoint_name")

    @service_endpoint_name.setter
    def service_endpoint_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_endpoint_name", value)

    @property
    @pulumi.getter(name="webhookName")
    def webhook_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the WebHook.
        """
        return pulumi.get(self, "webhook_name")

    @webhook_name.setter
    def webhook_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "webhook_name", value)


class ServiceendpointIncomingwebhook(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 http_header: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 secret: Optional[pulumi.Input[str]] = None,
                 service_endpoint_name: Optional[pulumi.Input[str]] = None,
                 webhook_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages an Incoming WebHook service endpoint within Azure DevOps, which can be used as a resource in YAML pipelines to subscribe to a webhook event.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        example_project = azuredevops.Project("exampleProject",
            visibility="private",
            version_control="Git",
            work_item_template="Agile",
            description="Managed by Terraform")
        example_serviceendpoint_incomingwebhook = azuredevops.ServiceendpointIncomingwebhook("exampleServiceendpointIncomingwebhook",
            project_id=example_project.id,
            webhook_name="example_webhook",
            secret="secret",
            http_header="X-Hub-Signature",
            service_endpoint_name="Example IncomingWebhook",
            description="Managed by Terraform")
        ```

        ## Import

        Azure DevOps Service Endpoint Incoming WebHook can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**

        ```sh
         $ pulumi import azuredevops:index/serviceendpointIncomingwebhook:ServiceendpointIncomingwebhook example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] http_header: Http header name on which checksum will be sent.
        :param pulumi.Input[str] project_id: The ID of the project. Changing this forces a new Service Connection Incoming WebHook to be created.
        :param pulumi.Input[str] secret: Secret for the WebHook. WebHook service will use this secret to calculate the payload checksum.
        :param pulumi.Input[str] service_endpoint_name: The name of the service endpoint. Changing this forces a new Service Connection Incoming WebHook to be created.
        :param pulumi.Input[str] webhook_name: The name of the WebHook.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ServiceendpointIncomingwebhookArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages an Incoming WebHook service endpoint within Azure DevOps, which can be used as a resource in YAML pipelines to subscribe to a webhook event.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        example_project = azuredevops.Project("exampleProject",
            visibility="private",
            version_control="Git",
            work_item_template="Agile",
            description="Managed by Terraform")
        example_serviceendpoint_incomingwebhook = azuredevops.ServiceendpointIncomingwebhook("exampleServiceendpointIncomingwebhook",
            project_id=example_project.id,
            webhook_name="example_webhook",
            secret="secret",
            http_header="X-Hub-Signature",
            service_endpoint_name="Example IncomingWebhook",
            description="Managed by Terraform")
        ```

        ## Import

        Azure DevOps Service Endpoint Incoming WebHook can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**

        ```sh
         $ pulumi import azuredevops:index/serviceendpointIncomingwebhook:ServiceendpointIncomingwebhook example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
        ```

        :param str resource_name: The name of the resource.
        :param ServiceendpointIncomingwebhookArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ServiceendpointIncomingwebhookArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 http_header: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 secret: Optional[pulumi.Input[str]] = None,
                 service_endpoint_name: Optional[pulumi.Input[str]] = None,
                 webhook_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ServiceendpointIncomingwebhookArgs.__new__(ServiceendpointIncomingwebhookArgs)

            __props__.__dict__["authorization"] = authorization
            __props__.__dict__["description"] = description
            __props__.__dict__["http_header"] = http_header
            if project_id is None and not opts.urn:
                raise TypeError("Missing required property 'project_id'")
            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["secret"] = None if secret is None else pulumi.Output.secret(secret)
            if service_endpoint_name is None and not opts.urn:
                raise TypeError("Missing required property 'service_endpoint_name'")
            __props__.__dict__["service_endpoint_name"] = service_endpoint_name
            if webhook_name is None and not opts.urn:
                raise TypeError("Missing required property 'webhook_name'")
            __props__.__dict__["webhook_name"] = webhook_name
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["secret"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(ServiceendpointIncomingwebhook, __self__).__init__(
            'azuredevops:index/serviceendpointIncomingwebhook:ServiceendpointIncomingwebhook',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            description: Optional[pulumi.Input[str]] = None,
            http_header: Optional[pulumi.Input[str]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            secret: Optional[pulumi.Input[str]] = None,
            service_endpoint_name: Optional[pulumi.Input[str]] = None,
            webhook_name: Optional[pulumi.Input[str]] = None) -> 'ServiceendpointIncomingwebhook':
        """
        Get an existing ServiceendpointIncomingwebhook resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] http_header: Http header name on which checksum will be sent.
        :param pulumi.Input[str] project_id: The ID of the project. Changing this forces a new Service Connection Incoming WebHook to be created.
        :param pulumi.Input[str] secret: Secret for the WebHook. WebHook service will use this secret to calculate the payload checksum.
        :param pulumi.Input[str] service_endpoint_name: The name of the service endpoint. Changing this forces a new Service Connection Incoming WebHook to be created.
        :param pulumi.Input[str] webhook_name: The name of the WebHook.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ServiceendpointIncomingwebhookState.__new__(_ServiceendpointIncomingwebhookState)

        __props__.__dict__["authorization"] = authorization
        __props__.__dict__["description"] = description
        __props__.__dict__["http_header"] = http_header
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["secret"] = secret
        __props__.__dict__["service_endpoint_name"] = service_endpoint_name
        __props__.__dict__["webhook_name"] = webhook_name
        return ServiceendpointIncomingwebhook(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def authorization(self) -> pulumi.Output[Mapping[str, str]]:
        return pulumi.get(self, "authorization")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="httpHeader")
    def http_header(self) -> pulumi.Output[Optional[str]]:
        """
        Http header name on which checksum will be sent.
        """
        return pulumi.get(self, "http_header")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        """
        The ID of the project. Changing this forces a new Service Connection Incoming WebHook to be created.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def secret(self) -> pulumi.Output[Optional[str]]:
        """
        Secret for the WebHook. WebHook service will use this secret to calculate the payload checksum.
        """
        return pulumi.get(self, "secret")

    @property
    @pulumi.getter(name="serviceEndpointName")
    def service_endpoint_name(self) -> pulumi.Output[str]:
        """
        The name of the service endpoint. Changing this forces a new Service Connection Incoming WebHook to be created.
        """
        return pulumi.get(self, "service_endpoint_name")

    @property
    @pulumi.getter(name="webhookName")
    def webhook_name(self) -> pulumi.Output[str]:
        """
        The name of the WebHook.
        """
        return pulumi.get(self, "webhook_name")

