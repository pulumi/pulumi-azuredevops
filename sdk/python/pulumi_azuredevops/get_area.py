# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs

__all__ = [
    'GetAreaResult',
    'AwaitableGetAreaResult',
    'get_area',
    'get_area_output',
]

@pulumi.output_type
class GetAreaResult:
    """
    A collection of values returned by getArea.
    """
    def __init__(__self__, childrens=None, fetch_children=None, has_children=None, id=None, name=None, path=None, project_id=None):
        if childrens and not isinstance(childrens, list):
            raise TypeError("Expected argument 'childrens' to be a list")
        pulumi.set(__self__, "childrens", childrens)
        if fetch_children and not isinstance(fetch_children, bool):
            raise TypeError("Expected argument 'fetch_children' to be a bool")
        pulumi.set(__self__, "fetch_children", fetch_children)
        if has_children and not isinstance(has_children, bool):
            raise TypeError("Expected argument 'has_children' to be a bool")
        pulumi.set(__self__, "has_children", has_children)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if path and not isinstance(path, str):
            raise TypeError("Expected argument 'path' to be a str")
        pulumi.set(__self__, "path", path)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)

    @property
    @pulumi.getter
    def childrens(self) -> Sequence['outputs.GetAreaChildrenResult']:
        """
        A list of `children` blocks as defined below, empty if `has_children == false`
        """
        return pulumi.get(self, "childrens")

    @property
    @pulumi.getter(name="fetchChildren")
    def fetch_children(self) -> Optional[bool]:
        return pulumi.get(self, "fetch_children")

    @property
    @pulumi.getter(name="hasChildren")
    def has_children(self) -> bool:
        """
        Indicator if the child Area node has child nodes
        """
        return pulumi.get(self, "has_children")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the child Area node
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def path(self) -> str:
        """
        The complete path (in relative URL format) of the child Area
        """
        return pulumi.get(self, "path")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> str:
        """
        The project ID of the child Area node
        """
        return pulumi.get(self, "project_id")


class AwaitableGetAreaResult(GetAreaResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAreaResult(
            childrens=self.childrens,
            fetch_children=self.fetch_children,
            has_children=self.has_children,
            id=self.id,
            name=self.name,
            path=self.path,
            project_id=self.project_id)


def get_area(fetch_children: Optional[bool] = None,
             path: Optional[str] = None,
             project_id: Optional[str] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAreaResult:
    """
    Use this data source to access information about an existing Area (Component) within Azure DevOps.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azuredevops as azuredevops

    #---------------------------------------------------------------------------
    # Azure DevOps project
    project = azuredevops.Project("project",
        work_item_template="Agile",
        version_control="Git",
        visibility="private",
        description="Managed by Terraform")
    area = project.id.apply(lambda id: azuredevops.get_area_output(project_id=id,
        path="/",
        fetch_children=False))
    ```
    ## Relevant Links

    - [Azure DevOps Service REST API 5.1 - Classification Nodes - Get Classification Nodes](https://docs.microsoft.com/en-us/rest/api/azure/devops/wit/classification%20nodes/get%20classification%20nodes?view=azure-devops-rest-5.1)

    ## PAT Permissions Required

    - **Project & Team**: vso.work - Grants the ability to read work items, queries, boards, area and iterations paths, and other work item tracking related metadata. Also grants the ability to execute queries, search work items and to receive notifications about work item events via service hooks.


    :param bool fetch_children: Read children nodes, _Depth_: 1, _Default_: `true`
    :param str path: The path to the Area; _Format_: URL relative; if omitted, or value `"/"` is used, the root Area will be returned
    :param str project_id: The project ID.
    """
    __args__ = dict()
    __args__['fetchChildren'] = fetch_children
    __args__['path'] = path
    __args__['projectId'] = project_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azuredevops:index/getArea:getArea', __args__, opts=opts, typ=GetAreaResult).value

    return AwaitableGetAreaResult(
        childrens=__ret__.childrens,
        fetch_children=__ret__.fetch_children,
        has_children=__ret__.has_children,
        id=__ret__.id,
        name=__ret__.name,
        path=__ret__.path,
        project_id=__ret__.project_id)


@_utilities.lift_output_func(get_area)
def get_area_output(fetch_children: Optional[pulumi.Input[Optional[bool]]] = None,
                    path: Optional[pulumi.Input[Optional[str]]] = None,
                    project_id: Optional[pulumi.Input[str]] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAreaResult]:
    """
    Use this data source to access information about an existing Area (Component) within Azure DevOps.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azuredevops as azuredevops

    #---------------------------------------------------------------------------
    # Azure DevOps project
    project = azuredevops.Project("project",
        work_item_template="Agile",
        version_control="Git",
        visibility="private",
        description="Managed by Terraform")
    area = project.id.apply(lambda id: azuredevops.get_area_output(project_id=id,
        path="/",
        fetch_children=False))
    ```
    ## Relevant Links

    - [Azure DevOps Service REST API 5.1 - Classification Nodes - Get Classification Nodes](https://docs.microsoft.com/en-us/rest/api/azure/devops/wit/classification%20nodes/get%20classification%20nodes?view=azure-devops-rest-5.1)

    ## PAT Permissions Required

    - **Project & Team**: vso.work - Grants the ability to read work items, queries, boards, area and iterations paths, and other work item tracking related metadata. Also grants the ability to execute queries, search work items and to receive notifications about work item events via service hooks.


    :param bool fetch_children: Read children nodes, _Depth_: 1, _Default_: `true`
    :param str path: The path to the Area; _Format_: URL relative; if omitted, or value `"/"` is used, the root Area will be returned
    :param str project_id: The project ID.
    """
    ...
