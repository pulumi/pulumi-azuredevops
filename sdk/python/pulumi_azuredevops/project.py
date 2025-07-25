# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins as _builtins
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

__all__ = ['ProjectArgs', 'Project']

@pulumi.input_type
class ProjectArgs:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[_builtins.str]] = None,
                 features: Optional[pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]]] = None,
                 name: Optional[pulumi.Input[_builtins.str]] = None,
                 version_control: Optional[pulumi.Input[_builtins.str]] = None,
                 visibility: Optional[pulumi.Input[_builtins.str]] = None,
                 work_item_template: Optional[pulumi.Input[_builtins.str]] = None):
        """
        The set of arguments for constructing a Project resource.
        :param pulumi.Input[_builtins.str] description: The Description of the Project.
               *
        :param pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]] features: Defines the status (`enabled`, `disabled`) of the project features. Valid features are `boards`, `repositories`, `pipelines`, `testplans`, `artifacts`
               
               | Features     | Possible Values   |
               |--------------|-------------------|
               | boards       | enabled, disabled |
               | repositories | enabled, disabled |
               | pipelines    | enabled, disabled |
               | testplans    | enabled, disabled |
               | artifacts    | enabled, disabled |
               
               > **NOTE:** It's possible to define project features both within the `ProjectFeatures` resource and
               via the `features` block by using the `Project` resource.
               However it's not possible to use both methods to manage features, since there'll be conflicts.
        :param pulumi.Input[_builtins.str] name: The Project Name.
        :param pulumi.Input[_builtins.str] version_control: Specifies the version control system. Possbile values are: `Git` or `Tfvc`. Defaults to `Git`.
        :param pulumi.Input[_builtins.str] visibility: Specifies the visibility of the Project. Possible values are: `private` or `public`. Defaults to `private`.
        :param pulumi.Input[_builtins.str] work_item_template: Specifies the work item template. Possible values are: `Agile`, `Basic`, `CMMI`, `Scrum` or a custom, pre-existing one. Defaults to `Agile`. An empty string will use the parent organization default.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if features is not None:
            pulumi.set(__self__, "features", features)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if version_control is not None:
            pulumi.set(__self__, "version_control", version_control)
        if visibility is not None:
            pulumi.set(__self__, "visibility", visibility)
        if work_item_template is not None:
            pulumi.set(__self__, "work_item_template", work_item_template)

    @_builtins.property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The Description of the Project.
        *
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "description", value)

    @_builtins.property
    @pulumi.getter
    def features(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]]]:
        """
        Defines the status (`enabled`, `disabled`) of the project features. Valid features are `boards`, `repositories`, `pipelines`, `testplans`, `artifacts`

        | Features     | Possible Values   |
        |--------------|-------------------|
        | boards       | enabled, disabled |
        | repositories | enabled, disabled |
        | pipelines    | enabled, disabled |
        | testplans    | enabled, disabled |
        | artifacts    | enabled, disabled |

        > **NOTE:** It's possible to define project features both within the `ProjectFeatures` resource and
        via the `features` block by using the `Project` resource.
        However it's not possible to use both methods to manage features, since there'll be conflicts.
        """
        return pulumi.get(self, "features")

    @features.setter
    def features(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]]]):
        pulumi.set(self, "features", value)

    @_builtins.property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The Project Name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "name", value)

    @_builtins.property
    @pulumi.getter(name="versionControl")
    def version_control(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Specifies the version control system. Possbile values are: `Git` or `Tfvc`. Defaults to `Git`.
        """
        return pulumi.get(self, "version_control")

    @version_control.setter
    def version_control(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "version_control", value)

    @_builtins.property
    @pulumi.getter
    def visibility(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Specifies the visibility of the Project. Possible values are: `private` or `public`. Defaults to `private`.
        """
        return pulumi.get(self, "visibility")

    @visibility.setter
    def visibility(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "visibility", value)

    @_builtins.property
    @pulumi.getter(name="workItemTemplate")
    def work_item_template(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Specifies the work item template. Possible values are: `Agile`, `Basic`, `CMMI`, `Scrum` or a custom, pre-existing one. Defaults to `Agile`. An empty string will use the parent organization default.
        """
        return pulumi.get(self, "work_item_template")

    @work_item_template.setter
    def work_item_template(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "work_item_template", value)


@pulumi.input_type
class _ProjectState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[_builtins.str]] = None,
                 features: Optional[pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]]] = None,
                 name: Optional[pulumi.Input[_builtins.str]] = None,
                 process_template_id: Optional[pulumi.Input[_builtins.str]] = None,
                 version_control: Optional[pulumi.Input[_builtins.str]] = None,
                 visibility: Optional[pulumi.Input[_builtins.str]] = None,
                 work_item_template: Optional[pulumi.Input[_builtins.str]] = None):
        """
        Input properties used for looking up and filtering Project resources.
        :param pulumi.Input[_builtins.str] description: The Description of the Project.
               *
        :param pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]] features: Defines the status (`enabled`, `disabled`) of the project features. Valid features are `boards`, `repositories`, `pipelines`, `testplans`, `artifacts`
               
               | Features     | Possible Values   |
               |--------------|-------------------|
               | boards       | enabled, disabled |
               | repositories | enabled, disabled |
               | pipelines    | enabled, disabled |
               | testplans    | enabled, disabled |
               | artifacts    | enabled, disabled |
               
               > **NOTE:** It's possible to define project features both within the `ProjectFeatures` resource and
               via the `features` block by using the `Project` resource.
               However it's not possible to use both methods to manage features, since there'll be conflicts.
        :param pulumi.Input[_builtins.str] name: The Project Name.
        :param pulumi.Input[_builtins.str] process_template_id: The Process Template ID used by the Project.
        :param pulumi.Input[_builtins.str] version_control: Specifies the version control system. Possbile values are: `Git` or `Tfvc`. Defaults to `Git`.
        :param pulumi.Input[_builtins.str] visibility: Specifies the visibility of the Project. Possible values are: `private` or `public`. Defaults to `private`.
        :param pulumi.Input[_builtins.str] work_item_template: Specifies the work item template. Possible values are: `Agile`, `Basic`, `CMMI`, `Scrum` or a custom, pre-existing one. Defaults to `Agile`. An empty string will use the parent organization default.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if features is not None:
            pulumi.set(__self__, "features", features)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if process_template_id is not None:
            pulumi.set(__self__, "process_template_id", process_template_id)
        if version_control is not None:
            pulumi.set(__self__, "version_control", version_control)
        if visibility is not None:
            pulumi.set(__self__, "visibility", visibility)
        if work_item_template is not None:
            pulumi.set(__self__, "work_item_template", work_item_template)

    @_builtins.property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The Description of the Project.
        *
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "description", value)

    @_builtins.property
    @pulumi.getter
    def features(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]]]:
        """
        Defines the status (`enabled`, `disabled`) of the project features. Valid features are `boards`, `repositories`, `pipelines`, `testplans`, `artifacts`

        | Features     | Possible Values   |
        |--------------|-------------------|
        | boards       | enabled, disabled |
        | repositories | enabled, disabled |
        | pipelines    | enabled, disabled |
        | testplans    | enabled, disabled |
        | artifacts    | enabled, disabled |

        > **NOTE:** It's possible to define project features both within the `ProjectFeatures` resource and
        via the `features` block by using the `Project` resource.
        However it's not possible to use both methods to manage features, since there'll be conflicts.
        """
        return pulumi.get(self, "features")

    @features.setter
    def features(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]]]):
        pulumi.set(self, "features", value)

    @_builtins.property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The Project Name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "name", value)

    @_builtins.property
    @pulumi.getter(name="processTemplateId")
    def process_template_id(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The Process Template ID used by the Project.
        """
        return pulumi.get(self, "process_template_id")

    @process_template_id.setter
    def process_template_id(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "process_template_id", value)

    @_builtins.property
    @pulumi.getter(name="versionControl")
    def version_control(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Specifies the version control system. Possbile values are: `Git` or `Tfvc`. Defaults to `Git`.
        """
        return pulumi.get(self, "version_control")

    @version_control.setter
    def version_control(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "version_control", value)

    @_builtins.property
    @pulumi.getter
    def visibility(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Specifies the visibility of the Project. Possible values are: `private` or `public`. Defaults to `private`.
        """
        return pulumi.get(self, "visibility")

    @visibility.setter
    def visibility(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "visibility", value)

    @_builtins.property
    @pulumi.getter(name="workItemTemplate")
    def work_item_template(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Specifies the work item template. Possible values are: `Agile`, `Basic`, `CMMI`, `Scrum` or a custom, pre-existing one. Defaults to `Agile`. An empty string will use the parent organization default.
        """
        return pulumi.get(self, "work_item_template")

    @work_item_template.setter
    def work_item_template(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "work_item_template", value)


@pulumi.type_token("azuredevops:index/project:Project")
class Project(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[_builtins.str]] = None,
                 features: Optional[pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]]] = None,
                 name: Optional[pulumi.Input[_builtins.str]] = None,
                 version_control: Optional[pulumi.Input[_builtins.str]] = None,
                 visibility: Optional[pulumi.Input[_builtins.str]] = None,
                 work_item_template: Optional[pulumi.Input[_builtins.str]] = None,
                 __props__=None):
        """
        Manages a project within Azure DevOps.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        example = azuredevops.Project("example",
            name="Example Project",
            visibility="private",
            version_control="Git",
            work_item_template="Agile",
            description="Managed by Pulumi",
            features={
                "testplans": "disabled",
                "artifacts": "disabled",
            })
        ```

        ## Relevant Links

        - [Azure DevOps Service REST API 7.0 - Projects](https://docs.microsoft.com/en-us/rest/api/azure/devops/core/projects?view=azure-devops-rest-7.0)

        ## PAT Permissions Required

        - **Project & Team**: Read, Write, & Manage
        - **Work Items**: Read

        ## Import

        Azure DevOps Projects can be imported using the project name or by the project Guid, e.g.

        ```sh
        $ pulumi import azuredevops:index/project:Project example "Example Project"
        ```

        or

        ```sh
        $ pulumi import azuredevops:index/project:Project example 00000000-0000-0000-0000-000000000000
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[_builtins.str] description: The Description of the Project.
               *
        :param pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]] features: Defines the status (`enabled`, `disabled`) of the project features. Valid features are `boards`, `repositories`, `pipelines`, `testplans`, `artifacts`
               
               | Features     | Possible Values   |
               |--------------|-------------------|
               | boards       | enabled, disabled |
               | repositories | enabled, disabled |
               | pipelines    | enabled, disabled |
               | testplans    | enabled, disabled |
               | artifacts    | enabled, disabled |
               
               > **NOTE:** It's possible to define project features both within the `ProjectFeatures` resource and
               via the `features` block by using the `Project` resource.
               However it's not possible to use both methods to manage features, since there'll be conflicts.
        :param pulumi.Input[_builtins.str] name: The Project Name.
        :param pulumi.Input[_builtins.str] version_control: Specifies the version control system. Possbile values are: `Git` or `Tfvc`. Defaults to `Git`.
        :param pulumi.Input[_builtins.str] visibility: Specifies the visibility of the Project. Possible values are: `private` or `public`. Defaults to `private`.
        :param pulumi.Input[_builtins.str] work_item_template: Specifies the work item template. Possible values are: `Agile`, `Basic`, `CMMI`, `Scrum` or a custom, pre-existing one. Defaults to `Agile`. An empty string will use the parent organization default.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ProjectArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a project within Azure DevOps.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        example = azuredevops.Project("example",
            name="Example Project",
            visibility="private",
            version_control="Git",
            work_item_template="Agile",
            description="Managed by Pulumi",
            features={
                "testplans": "disabled",
                "artifacts": "disabled",
            })
        ```

        ## Relevant Links

        - [Azure DevOps Service REST API 7.0 - Projects](https://docs.microsoft.com/en-us/rest/api/azure/devops/core/projects?view=azure-devops-rest-7.0)

        ## PAT Permissions Required

        - **Project & Team**: Read, Write, & Manage
        - **Work Items**: Read

        ## Import

        Azure DevOps Projects can be imported using the project name or by the project Guid, e.g.

        ```sh
        $ pulumi import azuredevops:index/project:Project example "Example Project"
        ```

        or

        ```sh
        $ pulumi import azuredevops:index/project:Project example 00000000-0000-0000-0000-000000000000
        ```

        :param str resource_name: The name of the resource.
        :param ProjectArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProjectArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[_builtins.str]] = None,
                 features: Optional[pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]]] = None,
                 name: Optional[pulumi.Input[_builtins.str]] = None,
                 version_control: Optional[pulumi.Input[_builtins.str]] = None,
                 visibility: Optional[pulumi.Input[_builtins.str]] = None,
                 work_item_template: Optional[pulumi.Input[_builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ProjectArgs.__new__(ProjectArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["features"] = features
            __props__.__dict__["name"] = name
            __props__.__dict__["version_control"] = version_control
            __props__.__dict__["visibility"] = visibility
            __props__.__dict__["work_item_template"] = work_item_template
            __props__.__dict__["process_template_id"] = None
        super(Project, __self__).__init__(
            'azuredevops:index/project:Project',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[_builtins.str]] = None,
            features: Optional[pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]]] = None,
            name: Optional[pulumi.Input[_builtins.str]] = None,
            process_template_id: Optional[pulumi.Input[_builtins.str]] = None,
            version_control: Optional[pulumi.Input[_builtins.str]] = None,
            visibility: Optional[pulumi.Input[_builtins.str]] = None,
            work_item_template: Optional[pulumi.Input[_builtins.str]] = None) -> 'Project':
        """
        Get an existing Project resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[_builtins.str] description: The Description of the Project.
               *
        :param pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]] features: Defines the status (`enabled`, `disabled`) of the project features. Valid features are `boards`, `repositories`, `pipelines`, `testplans`, `artifacts`
               
               | Features     | Possible Values   |
               |--------------|-------------------|
               | boards       | enabled, disabled |
               | repositories | enabled, disabled |
               | pipelines    | enabled, disabled |
               | testplans    | enabled, disabled |
               | artifacts    | enabled, disabled |
               
               > **NOTE:** It's possible to define project features both within the `ProjectFeatures` resource and
               via the `features` block by using the `Project` resource.
               However it's not possible to use both methods to manage features, since there'll be conflicts.
        :param pulumi.Input[_builtins.str] name: The Project Name.
        :param pulumi.Input[_builtins.str] process_template_id: The Process Template ID used by the Project.
        :param pulumi.Input[_builtins.str] version_control: Specifies the version control system. Possbile values are: `Git` or `Tfvc`. Defaults to `Git`.
        :param pulumi.Input[_builtins.str] visibility: Specifies the visibility of the Project. Possible values are: `private` or `public`. Defaults to `private`.
        :param pulumi.Input[_builtins.str] work_item_template: Specifies the work item template. Possible values are: `Agile`, `Basic`, `CMMI`, `Scrum` or a custom, pre-existing one. Defaults to `Agile`. An empty string will use the parent organization default.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ProjectState.__new__(_ProjectState)

        __props__.__dict__["description"] = description
        __props__.__dict__["features"] = features
        __props__.__dict__["name"] = name
        __props__.__dict__["process_template_id"] = process_template_id
        __props__.__dict__["version_control"] = version_control
        __props__.__dict__["visibility"] = visibility
        __props__.__dict__["work_item_template"] = work_item_template
        return Project(resource_name, opts=opts, __props__=__props__)

    @_builtins.property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[_builtins.str]]:
        """
        The Description of the Project.
        *
        """
        return pulumi.get(self, "description")

    @_builtins.property
    @pulumi.getter
    def features(self) -> pulumi.Output[Optional[Mapping[str, _builtins.str]]]:
        """
        Defines the status (`enabled`, `disabled`) of the project features. Valid features are `boards`, `repositories`, `pipelines`, `testplans`, `artifacts`

        | Features     | Possible Values   |
        |--------------|-------------------|
        | boards       | enabled, disabled |
        | repositories | enabled, disabled |
        | pipelines    | enabled, disabled |
        | testplans    | enabled, disabled |
        | artifacts    | enabled, disabled |

        > **NOTE:** It's possible to define project features both within the `ProjectFeatures` resource and
        via the `features` block by using the `Project` resource.
        However it's not possible to use both methods to manage features, since there'll be conflicts.
        """
        return pulumi.get(self, "features")

    @_builtins.property
    @pulumi.getter
    def name(self) -> pulumi.Output[_builtins.str]:
        """
        The Project Name.
        """
        return pulumi.get(self, "name")

    @_builtins.property
    @pulumi.getter(name="processTemplateId")
    def process_template_id(self) -> pulumi.Output[_builtins.str]:
        """
        The Process Template ID used by the Project.
        """
        return pulumi.get(self, "process_template_id")

    @_builtins.property
    @pulumi.getter(name="versionControl")
    def version_control(self) -> pulumi.Output[Optional[_builtins.str]]:
        """
        Specifies the version control system. Possbile values are: `Git` or `Tfvc`. Defaults to `Git`.
        """
        return pulumi.get(self, "version_control")

    @_builtins.property
    @pulumi.getter
    def visibility(self) -> pulumi.Output[Optional[_builtins.str]]:
        """
        Specifies the visibility of the Project. Possible values are: `private` or `public`. Defaults to `private`.
        """
        return pulumi.get(self, "visibility")

    @_builtins.property
    @pulumi.getter(name="workItemTemplate")
    def work_item_template(self) -> pulumi.Output[Optional[_builtins.str]]:
        """
        Specifies the work item template. Possible values are: `Agile`, `Basic`, `CMMI`, `Scrum` or a custom, pre-existing one. Defaults to `Agile`. An empty string will use the parent organization default.
        """
        return pulumi.get(self, "work_item_template")

