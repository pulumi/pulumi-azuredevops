# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs

__all__ = [
    'GetVariableGroupResult',
    'AwaitableGetVariableGroupResult',
    'get_variable_group',
    'get_variable_group_output',
]

@pulumi.output_type
class GetVariableGroupResult:
    """
    A collection of values returned by getVariableGroup.
    """
    def __init__(__self__, allow_access=None, description=None, id=None, key_vaults=None, name=None, project_id=None, variables=None):
        if allow_access and not isinstance(allow_access, bool):
            raise TypeError("Expected argument 'allow_access' to be a bool")
        pulumi.set(__self__, "allow_access", allow_access)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if key_vaults and not isinstance(key_vaults, list):
            raise TypeError("Expected argument 'key_vaults' to be a list")
        pulumi.set(__self__, "key_vaults", key_vaults)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if variables and not isinstance(variables, list):
            raise TypeError("Expected argument 'variables' to be a list")
        pulumi.set(__self__, "variables", variables)

    @property
    @pulumi.getter(name="allowAccess")
    def allow_access(self) -> bool:
        """
        Boolean that indicate if this Variable Group is shared by all pipelines of this project.
        """
        return pulumi.get(self, "allow_access")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        The description of the Variable Group.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="keyVaults")
    def key_vaults(self) -> Sequence['outputs.GetVariableGroupKeyVaultResult']:
        """
        A list of `key_vault` blocks as documented below.
        """
        return pulumi.get(self, "key_vaults")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the Azure key vault to link secrets from as variables.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> str:
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def variables(self) -> Sequence['outputs.GetVariableGroupVariableResult']:
        """
        One or more `variable` blocks as documented below.
        """
        return pulumi.get(self, "variables")


class AwaitableGetVariableGroupResult(GetVariableGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVariableGroupResult(
            allow_access=self.allow_access,
            description=self.description,
            id=self.id,
            key_vaults=self.key_vaults,
            name=self.name,
            project_id=self.project_id,
            variables=self.variables)


def get_variable_group(name: Optional[str] = None,
                       project_id: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVariableGroupResult:
    """
    Use this data source to access information about existing Variable Groups within Azure DevOps.

    > **Note:** Secret values are masked by service and cannot be obtained through API. [Set secret variables](https://docs.microsoft.com/en-us/azure/devops/pipelines/process/variables?view=azure-devops&tabs=yaml%2Cbatch#secret-variables)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azuredevops as azuredevops

    example_project = azuredevops.get_project(name="Example Project")
    example_variable_group = azuredevops.get_variable_group(project_id=example_project.id,
        name="Example Variable Group")
    pulumi.export("id", example_variable_group.id)
    ```
    ## Relevant Links

    - [Azure DevOps Service REST API 6.0 - Variable Groups](https://docs.microsoft.com/en-us/rest/api/azure/devops/distributedtask/variablegroups?view=azure-devops-rest-6.0)


    :param str name: The name of the Variable Group to retrieve.
    :param str project_id: The project ID.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['projectId'] = project_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azuredevops:index/getVariableGroup:getVariableGroup', __args__, opts=opts, typ=GetVariableGroupResult).value

    return AwaitableGetVariableGroupResult(
        allow_access=pulumi.get(__ret__, 'allow_access'),
        description=pulumi.get(__ret__, 'description'),
        id=pulumi.get(__ret__, 'id'),
        key_vaults=pulumi.get(__ret__, 'key_vaults'),
        name=pulumi.get(__ret__, 'name'),
        project_id=pulumi.get(__ret__, 'project_id'),
        variables=pulumi.get(__ret__, 'variables'))


@_utilities.lift_output_func(get_variable_group)
def get_variable_group_output(name: Optional[pulumi.Input[str]] = None,
                              project_id: Optional[pulumi.Input[str]] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetVariableGroupResult]:
    """
    Use this data source to access information about existing Variable Groups within Azure DevOps.

    > **Note:** Secret values are masked by service and cannot be obtained through API. [Set secret variables](https://docs.microsoft.com/en-us/azure/devops/pipelines/process/variables?view=azure-devops&tabs=yaml%2Cbatch#secret-variables)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azuredevops as azuredevops

    example_project = azuredevops.get_project(name="Example Project")
    example_variable_group = azuredevops.get_variable_group(project_id=example_project.id,
        name="Example Variable Group")
    pulumi.export("id", example_variable_group.id)
    ```
    ## Relevant Links

    - [Azure DevOps Service REST API 6.0 - Variable Groups](https://docs.microsoft.com/en-us/rest/api/azure/devops/distributedtask/variablegroups?view=azure-devops-rest-6.0)


    :param str name: The name of the Variable Group to retrieve.
    :param str project_id: The project ID.
    """
    ...
