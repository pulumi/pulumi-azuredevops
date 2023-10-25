# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'VariableGroupKeyVaultArgs',
    'VariableGroupVariableArgs',
]

@pulumi.input_type
class VariableGroupKeyVaultArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 service_endpoint_id: pulumi.Input[str],
                 search_depth: Optional[pulumi.Input[int]] = None):
        """
        :param pulumi.Input[str] name: The name of the Azure key vault to link secrets from as variables.
        :param pulumi.Input[str] service_endpoint_id: The id of the Azure subscription endpoint to access the key vault.
        :param pulumi.Input[int] search_depth: Set the Azure Key Vault Secret search depth. Defaults to `20`.
        """
        VariableGroupKeyVaultArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            name=name,
            service_endpoint_id=service_endpoint_id,
            search_depth=search_depth,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             name: Optional[pulumi.Input[str]] = None,
             service_endpoint_id: Optional[pulumi.Input[str]] = None,
             search_depth: Optional[pulumi.Input[int]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if name is None:
            raise TypeError("Missing 'name' argument")
        if service_endpoint_id is None and 'serviceEndpointId' in kwargs:
            service_endpoint_id = kwargs['serviceEndpointId']
        if service_endpoint_id is None:
            raise TypeError("Missing 'service_endpoint_id' argument")
        if search_depth is None and 'searchDepth' in kwargs:
            search_depth = kwargs['searchDepth']

        _setter("name", name)
        _setter("service_endpoint_id", service_endpoint_id)
        if search_depth is not None:
            _setter("search_depth", search_depth)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        The name of the Azure key vault to link secrets from as variables.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="serviceEndpointId")
    def service_endpoint_id(self) -> pulumi.Input[str]:
        """
        The id of the Azure subscription endpoint to access the key vault.
        """
        return pulumi.get(self, "service_endpoint_id")

    @service_endpoint_id.setter
    def service_endpoint_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "service_endpoint_id", value)

    @property
    @pulumi.getter(name="searchDepth")
    def search_depth(self) -> Optional[pulumi.Input[int]]:
        """
        Set the Azure Key Vault Secret search depth. Defaults to `20`.
        """
        return pulumi.get(self, "search_depth")

    @search_depth.setter
    def search_depth(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "search_depth", value)


@pulumi.input_type
class VariableGroupVariableArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 content_type: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 expires: Optional[pulumi.Input[str]] = None,
                 is_secret: Optional[pulumi.Input[bool]] = None,
                 secret_value: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] name: The key value used for the variable. Must be unique within the Variable Group.
        :param pulumi.Input[bool] is_secret: A boolean flag describing if the variable value is sensitive. Defaults to `false`.
        :param pulumi.Input[str] secret_value: The secret value of the variable. If omitted, it will default to empty string. Used when `is_secret` set to `true`.
        :param pulumi.Input[str] value: The value of the variable. If omitted, it will default to empty string.
        """
        VariableGroupVariableArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            name=name,
            content_type=content_type,
            enabled=enabled,
            expires=expires,
            is_secret=is_secret,
            secret_value=secret_value,
            value=value,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             name: Optional[pulumi.Input[str]] = None,
             content_type: Optional[pulumi.Input[str]] = None,
             enabled: Optional[pulumi.Input[bool]] = None,
             expires: Optional[pulumi.Input[str]] = None,
             is_secret: Optional[pulumi.Input[bool]] = None,
             secret_value: Optional[pulumi.Input[str]] = None,
             value: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if name is None:
            raise TypeError("Missing 'name' argument")
        if content_type is None and 'contentType' in kwargs:
            content_type = kwargs['contentType']
        if is_secret is None and 'isSecret' in kwargs:
            is_secret = kwargs['isSecret']
        if secret_value is None and 'secretValue' in kwargs:
            secret_value = kwargs['secretValue']

        _setter("name", name)
        if content_type is not None:
            _setter("content_type", content_type)
        if enabled is not None:
            _setter("enabled", enabled)
        if expires is not None:
            _setter("expires", expires)
        if is_secret is not None:
            _setter("is_secret", is_secret)
        if secret_value is not None:
            _setter("secret_value", secret_value)
        if value is not None:
            _setter("value", value)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        The key value used for the variable. Must be unique within the Variable Group.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="contentType")
    def content_type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "content_type")

    @content_type.setter
    def content_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "content_type", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter
    def expires(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "expires")

    @expires.setter
    def expires(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expires", value)

    @property
    @pulumi.getter(name="isSecret")
    def is_secret(self) -> Optional[pulumi.Input[bool]]:
        """
        A boolean flag describing if the variable value is sensitive. Defaults to `false`.
        """
        return pulumi.get(self, "is_secret")

    @is_secret.setter
    def is_secret(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_secret", value)

    @property
    @pulumi.getter(name="secretValue")
    def secret_value(self) -> Optional[pulumi.Input[str]]:
        """
        The secret value of the variable. If omitted, it will default to empty string. Used when `is_secret` set to `true`.
        """
        return pulumi.get(self, "secret_value")

    @secret_value.setter
    def secret_value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secret_value", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[str]]:
        """
        The value of the variable. If omitted, it will default to empty string.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "value", value)


