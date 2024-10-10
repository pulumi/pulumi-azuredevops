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

__all__ = [
    'GetIdentityGroupsResult',
    'AwaitableGetIdentityGroupsResult',
    'get_identity_groups',
    'get_identity_groups_output',
]

@pulumi.output_type
class GetIdentityGroupsResult:
    """
    A collection of values returned by getIdentityGroups.
    """
    def __init__(__self__, groups=None, id=None, project_id=None):
        if groups and not isinstance(groups, list):
            raise TypeError("Expected argument 'groups' to be a list")
        pulumi.set(__self__, "groups", groups)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)

    @property
    @pulumi.getter
    def groups(self) -> Sequence['outputs.GetIdentityGroupsGroupResult']:
        """
        A set of existing groups in your Azure DevOps Organization or project with details about every single group which includes:
        """
        return pulumi.get(self, "groups")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[str]:
        return pulumi.get(self, "project_id")


class AwaitableGetIdentityGroupsResult(GetIdentityGroupsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetIdentityGroupsResult(
            groups=self.groups,
            id=self.id,
            project_id=self.project_id)


def get_identity_groups(project_id: Optional[str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetIdentityGroupsResult:
    """
    Use this data source to access information about existing Groups within Azure DevOps On-Premise(Azure DevOps Server).

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azuredevops as azuredevops

    example = azuredevops.get_project(name="Example Project")
    # load all existing groups inside an organization
    example_all_groups = azuredevops.get_identity_groups()
    # load all existing groups inside a specific project
    example_project_groups = azuredevops.get_identity_groups(project_id=example.id)
    ```

    ## Relevant Links

    - [Azure DevOps Service REST API 7.0 - Identities](https://docs.microsoft.com/en-us/rest/api/azure/devops/ims/?view=azure-devops-rest-7.2)


    :param str project_id: The Project ID. If no project ID is specified all groups of an organization will be returned
    """
    __args__ = dict()
    __args__['projectId'] = project_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azuredevops:index/getIdentityGroups:getIdentityGroups', __args__, opts=opts, typ=GetIdentityGroupsResult).value

    return AwaitableGetIdentityGroupsResult(
        groups=pulumi.get(__ret__, 'groups'),
        id=pulumi.get(__ret__, 'id'),
        project_id=pulumi.get(__ret__, 'project_id'))
def get_identity_groups_output(project_id: Optional[pulumi.Input[Optional[str]]] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetIdentityGroupsResult]:
    """
    Use this data source to access information about existing Groups within Azure DevOps On-Premise(Azure DevOps Server).

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azuredevops as azuredevops

    example = azuredevops.get_project(name="Example Project")
    # load all existing groups inside an organization
    example_all_groups = azuredevops.get_identity_groups()
    # load all existing groups inside a specific project
    example_project_groups = azuredevops.get_identity_groups(project_id=example.id)
    ```

    ## Relevant Links

    - [Azure DevOps Service REST API 7.0 - Identities](https://docs.microsoft.com/en-us/rest/api/azure/devops/ims/?view=azure-devops-rest-7.2)


    :param str project_id: The Project ID. If no project ID is specified all groups of an organization will be returned
    """
    __args__ = dict()
    __args__['projectId'] = project_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azuredevops:index/getIdentityGroups:getIdentityGroups', __args__, opts=opts, typ=GetIdentityGroupsResult)
    return __ret__.apply(lambda __response__: GetIdentityGroupsResult(
        groups=pulumi.get(__response__, 'groups'),
        id=pulumi.get(__response__, 'id'),
        project_id=pulumi.get(__response__, 'project_id')))
