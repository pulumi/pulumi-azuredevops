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
    'VariableGroupKeyVault',
    'VariableGroupVariable',
]

@pulumi.output_type
class VariableGroupKeyVault(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "serviceEndpointId":
            suggest = "service_endpoint_id"
        elif key == "searchDepth":
            suggest = "search_depth"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in VariableGroupKeyVault. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        VariableGroupKeyVault.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        VariableGroupKeyVault.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 name: str,
                 service_endpoint_id: str,
                 search_depth: Optional[int] = None):
        """
        :param str name: The name of the Azure key vault to link secrets from as variables.
        :param str service_endpoint_id: The id of the Azure subscription endpoint to access the key vault.
        :param int search_depth: Set the Azure Key Vault Secret search depth. Defaults to `20`.
        """
        VariableGroupKeyVault._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            name=name,
            service_endpoint_id=service_endpoint_id,
            search_depth=search_depth,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             name: Optional[str] = None,
             service_endpoint_id: Optional[str] = None,
             search_depth: Optional[int] = None,
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
    def name(self) -> str:
        """
        The name of the Azure key vault to link secrets from as variables.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="serviceEndpointId")
    def service_endpoint_id(self) -> str:
        """
        The id of the Azure subscription endpoint to access the key vault.
        """
        return pulumi.get(self, "service_endpoint_id")

    @property
    @pulumi.getter(name="searchDepth")
    def search_depth(self) -> Optional[int]:
        """
        Set the Azure Key Vault Secret search depth. Defaults to `20`.
        """
        return pulumi.get(self, "search_depth")


@pulumi.output_type
class VariableGroupVariable(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "contentType":
            suggest = "content_type"
        elif key == "isSecret":
            suggest = "is_secret"
        elif key == "secretValue":
            suggest = "secret_value"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in VariableGroupVariable. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        VariableGroupVariable.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        VariableGroupVariable.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 name: str,
                 content_type: Optional[str] = None,
                 enabled: Optional[bool] = None,
                 expires: Optional[str] = None,
                 is_secret: Optional[bool] = None,
                 secret_value: Optional[str] = None,
                 value: Optional[str] = None):
        """
        :param str name: The key value used for the variable. Must be unique within the Variable Group.
        :param bool is_secret: A boolean flag describing if the variable value is sensitive. Defaults to `false`.
        :param str secret_value: The secret value of the variable. If omitted, it will default to empty string. Used when `is_secret` set to `true`.
        :param str value: The value of the variable. If omitted, it will default to empty string.
        """
        VariableGroupVariable._configure(
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
             name: Optional[str] = None,
             content_type: Optional[str] = None,
             enabled: Optional[bool] = None,
             expires: Optional[str] = None,
             is_secret: Optional[bool] = None,
             secret_value: Optional[str] = None,
             value: Optional[str] = None,
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
    def name(self) -> str:
        """
        The key value used for the variable. Must be unique within the Variable Group.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="contentType")
    def content_type(self) -> Optional[str]:
        return pulumi.get(self, "content_type")

    @property
    @pulumi.getter
    def enabled(self) -> Optional[bool]:
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter
    def expires(self) -> Optional[str]:
        return pulumi.get(self, "expires")

    @property
    @pulumi.getter(name="isSecret")
    def is_secret(self) -> Optional[bool]:
        """
        A boolean flag describing if the variable value is sensitive. Defaults to `false`.
        """
        return pulumi.get(self, "is_secret")

    @property
    @pulumi.getter(name="secretValue")
    def secret_value(self) -> Optional[str]:
        """
        The secret value of the variable. If omitted, it will default to empty string. Used when `is_secret` set to `true`.
        """
        return pulumi.get(self, "secret_value")

    @property
    @pulumi.getter
    def value(self) -> Optional[str]:
        """
        The value of the variable. If omitted, it will default to empty string.
        """
        return pulumi.get(self, "value")


