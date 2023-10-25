# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ProjectFeaturesArgs', 'ProjectFeatures']

@pulumi.input_type
class ProjectFeaturesArgs:
    def __init__(__self__, *,
                 features: pulumi.Input[Mapping[str, pulumi.Input[str]]],
                 project_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a ProjectFeatures resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] features: Defines the status (`enabled`, `disabled`) of the project features.  
               Valid features `boards`, `repositories`, `pipelines`, `testplans`, `artifacts`
               
               > **NOTE:**
               > It's possible to define project features both within the `ProjectFeatures` resource and
               > via the `features` block by using the `Project` resource.
               > However it's not possible to use both methods to manage features, since there'll be conflicts.
        """
        ProjectFeaturesArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            features=features,
            project_id=project_id,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             features: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
             project_id: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if features is None:
            raise TypeError("Missing 'features' argument")
        if project_id is None and 'projectId' in kwargs:
            project_id = kwargs['projectId']
        if project_id is None:
            raise TypeError("Missing 'project_id' argument")

        _setter("features", features)
        _setter("project_id", project_id)

    @property
    @pulumi.getter
    def features(self) -> pulumi.Input[Mapping[str, pulumi.Input[str]]]:
        """
        Defines the status (`enabled`, `disabled`) of the project features.  
        Valid features `boards`, `repositories`, `pipelines`, `testplans`, `artifacts`

        > **NOTE:**
        > It's possible to define project features both within the `ProjectFeatures` resource and
        > via the `features` block by using the `Project` resource.
        > However it's not possible to use both methods to manage features, since there'll be conflicts.
        """
        return pulumi.get(self, "features")

    @features.setter
    def features(self, value: pulumi.Input[Mapping[str, pulumi.Input[str]]]):
        pulumi.set(self, "features", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "project_id", value)


@pulumi.input_type
class _ProjectFeaturesState:
    def __init__(__self__, *,
                 features: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 project_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ProjectFeatures resources.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] features: Defines the status (`enabled`, `disabled`) of the project features.  
               Valid features `boards`, `repositories`, `pipelines`, `testplans`, `artifacts`
               
               > **NOTE:**
               > It's possible to define project features both within the `ProjectFeatures` resource and
               > via the `features` block by using the `Project` resource.
               > However it's not possible to use both methods to manage features, since there'll be conflicts.
        """
        _ProjectFeaturesState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            features=features,
            project_id=project_id,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             features: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
             project_id: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if project_id is None and 'projectId' in kwargs:
            project_id = kwargs['projectId']

        if features is not None:
            _setter("features", features)
        if project_id is not None:
            _setter("project_id", project_id)

    @property
    @pulumi.getter
    def features(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Defines the status (`enabled`, `disabled`) of the project features.  
        Valid features `boards`, `repositories`, `pipelines`, `testplans`, `artifacts`

        > **NOTE:**
        > It's possible to define project features both within the `ProjectFeatures` resource and
        > via the `features` block by using the `Project` resource.
        > However it's not possible to use both methods to manage features, since there'll be conflicts.
        """
        return pulumi.get(self, "features")

    @features.setter
    def features(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "features", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)


warnings.warn("""azuredevops.core.ProjectFeatures has been deprecated in favor of azuredevops.ProjectFeatures""", DeprecationWarning)


class ProjectFeatures(pulumi.CustomResource):
    warnings.warn("""azuredevops.core.ProjectFeatures has been deprecated in favor of azuredevops.ProjectFeatures""", DeprecationWarning)

    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 features: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages features for Azure DevOps projects

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        example = azuredevops.Project("example",
            visibility="private",
            version_control="Git",
            work_item_template="Agile",
            description="Managed by Terraform")
        example_features = azuredevops.ProjectFeatures("example-features",
            project_id=example.id,
            features={
                "testplans": "disabled",
                "artifacts": "enabled",
            })
        ```
        ## Relevant Links

        No official documentation available

        ## PAT Permissions Required

        - **Project & Team**: Read, Write, & Manage

        ## Import

        Azure DevOps feature settings can be imported using the project id, e.g.

        ```sh
         $ pulumi import azuredevops:Core/projectFeatures:ProjectFeatures example 00000000-0000-0000-0000-000000000000
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] features: Defines the status (`enabled`, `disabled`) of the project features.  
               Valid features `boards`, `repositories`, `pipelines`, `testplans`, `artifacts`
               
               > **NOTE:**
               > It's possible to define project features both within the `ProjectFeatures` resource and
               > via the `features` block by using the `Project` resource.
               > However it's not possible to use both methods to manage features, since there'll be conflicts.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ProjectFeaturesArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages features for Azure DevOps projects

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        example = azuredevops.Project("example",
            visibility="private",
            version_control="Git",
            work_item_template="Agile",
            description="Managed by Terraform")
        example_features = azuredevops.ProjectFeatures("example-features",
            project_id=example.id,
            features={
                "testplans": "disabled",
                "artifacts": "enabled",
            })
        ```
        ## Relevant Links

        No official documentation available

        ## PAT Permissions Required

        - **Project & Team**: Read, Write, & Manage

        ## Import

        Azure DevOps feature settings can be imported using the project id, e.g.

        ```sh
         $ pulumi import azuredevops:Core/projectFeatures:ProjectFeatures example 00000000-0000-0000-0000-000000000000
        ```

        :param str resource_name: The name of the resource.
        :param ProjectFeaturesArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProjectFeaturesArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            ProjectFeaturesArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 features: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        pulumi.log.warn("""ProjectFeatures is deprecated: azuredevops.core.ProjectFeatures has been deprecated in favor of azuredevops.ProjectFeatures""")
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ProjectFeaturesArgs.__new__(ProjectFeaturesArgs)

            if features is None and not opts.urn:
                raise TypeError("Missing required property 'features'")
            __props__.__dict__["features"] = features
            if project_id is None and not opts.urn:
                raise TypeError("Missing required property 'project_id'")
            __props__.__dict__["project_id"] = project_id
        super(ProjectFeatures, __self__).__init__(
            'azuredevops:Core/projectFeatures:ProjectFeatures',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            features: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            project_id: Optional[pulumi.Input[str]] = None) -> 'ProjectFeatures':
        """
        Get an existing ProjectFeatures resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] features: Defines the status (`enabled`, `disabled`) of the project features.  
               Valid features `boards`, `repositories`, `pipelines`, `testplans`, `artifacts`
               
               > **NOTE:**
               > It's possible to define project features both within the `ProjectFeatures` resource and
               > via the `features` block by using the `Project` resource.
               > However it's not possible to use both methods to manage features, since there'll be conflicts.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ProjectFeaturesState.__new__(_ProjectFeaturesState)

        __props__.__dict__["features"] = features
        __props__.__dict__["project_id"] = project_id
        return ProjectFeatures(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def features(self) -> pulumi.Output[Mapping[str, str]]:
        """
        Defines the status (`enabled`, `disabled`) of the project features.  
        Valid features `boards`, `repositories`, `pipelines`, `testplans`, `artifacts`

        > **NOTE:**
        > It's possible to define project features both within the `ProjectFeatures` resource and
        > via the `features` block by using the `Project` resource.
        > However it's not possible to use both methods to manage features, since there'll be conflicts.
        """
        return pulumi.get(self, "features")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "project_id")

