# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetClientConfigResult',
    'AwaitableGetClientConfigResult',
    'get_client_config',
    'get_client_config_output',
]

warnings.warn("""azuredevops.core.getClientConfig has been deprecated in favor of azuredevops.getClientConfig""", DeprecationWarning)

@pulumi.output_type
class GetClientConfigResult:
    """
    A collection of values returned by getClientConfig.
    """
    def __init__(__self__, id=None, organization_url=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if organization_url and not isinstance(organization_url, str):
            raise TypeError("Expected argument 'organization_url' to be a str")
        pulumi.set(__self__, "organization_url", organization_url)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="organizationUrl")
    def organization_url(self) -> str:
        return pulumi.get(self, "organization_url")


class AwaitableGetClientConfigResult(GetClientConfigResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetClientConfigResult(
            id=self.id,
            organization_url=self.organization_url)


def get_client_config(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetClientConfigResult:
    """
    Use this data source to access information about the Azure DevOps organization configured for the provider.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azuredevops as azuredevops

    example = azuredevops.get_client_config()
    pulumi.export("orgUrl", example.organization_url)
    ```
    """
    pulumi.log.warn("""get_client_config is deprecated: azuredevops.core.getClientConfig has been deprecated in favor of azuredevops.getClientConfig""")
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azuredevops:Core/getClientConfig:getClientConfig', __args__, opts=opts, typ=GetClientConfigResult).value

    return AwaitableGetClientConfigResult(
        id=pulumi.get(__ret__, 'id'),
        organization_url=pulumi.get(__ret__, 'organization_url'))


@_utilities.lift_output_func(get_client_config)
def get_client_config_output(opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetClientConfigResult]:
    """
    Use this data source to access information about the Azure DevOps organization configured for the provider.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azuredevops as azuredevops

    example = azuredevops.get_client_config()
    pulumi.export("orgUrl", example.organization_url)
    ```
    """
    pulumi.log.warn("""get_client_config is deprecated: azuredevops.core.getClientConfig has been deprecated in favor of azuredevops.getClientConfig""")
    ...
