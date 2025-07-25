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

__all__ = ['FeedRetentionPolicyArgs', 'FeedRetentionPolicy']

@pulumi.input_type
class FeedRetentionPolicyArgs:
    def __init__(__self__, *,
                 count_limit: pulumi.Input[_builtins.int],
                 days_to_keep_recently_downloaded_packages: pulumi.Input[_builtins.int],
                 feed_id: pulumi.Input[_builtins.str],
                 project_id: Optional[pulumi.Input[_builtins.str]] = None):
        """
        The set of arguments for constructing a FeedRetentionPolicy resource.
        :param pulumi.Input[_builtins.int] count_limit: The maximum number of versions per package.
        :param pulumi.Input[_builtins.int] days_to_keep_recently_downloaded_packages: The days to keep recently downloaded packages.
        :param pulumi.Input[_builtins.str] feed_id: The ID of the Feed. Changing this forces a new resource to be created.
        :param pulumi.Input[_builtins.str] project_id: The ID of the Project. If not specified, Feed will be created at the organization level. Changing this forces a new resource to be created.
        """
        pulumi.set(__self__, "count_limit", count_limit)
        pulumi.set(__self__, "days_to_keep_recently_downloaded_packages", days_to_keep_recently_downloaded_packages)
        pulumi.set(__self__, "feed_id", feed_id)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)

    @_builtins.property
    @pulumi.getter(name="countLimit")
    def count_limit(self) -> pulumi.Input[_builtins.int]:
        """
        The maximum number of versions per package.
        """
        return pulumi.get(self, "count_limit")

    @count_limit.setter
    def count_limit(self, value: pulumi.Input[_builtins.int]):
        pulumi.set(self, "count_limit", value)

    @_builtins.property
    @pulumi.getter(name="daysToKeepRecentlyDownloadedPackages")
    def days_to_keep_recently_downloaded_packages(self) -> pulumi.Input[_builtins.int]:
        """
        The days to keep recently downloaded packages.
        """
        return pulumi.get(self, "days_to_keep_recently_downloaded_packages")

    @days_to_keep_recently_downloaded_packages.setter
    def days_to_keep_recently_downloaded_packages(self, value: pulumi.Input[_builtins.int]):
        pulumi.set(self, "days_to_keep_recently_downloaded_packages", value)

    @_builtins.property
    @pulumi.getter(name="feedId")
    def feed_id(self) -> pulumi.Input[_builtins.str]:
        """
        The ID of the Feed. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "feed_id")

    @feed_id.setter
    def feed_id(self, value: pulumi.Input[_builtins.str]):
        pulumi.set(self, "feed_id", value)

    @_builtins.property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The ID of the Project. If not specified, Feed will be created at the organization level. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "project_id", value)


@pulumi.input_type
class _FeedRetentionPolicyState:
    def __init__(__self__, *,
                 count_limit: Optional[pulumi.Input[_builtins.int]] = None,
                 days_to_keep_recently_downloaded_packages: Optional[pulumi.Input[_builtins.int]] = None,
                 feed_id: Optional[pulumi.Input[_builtins.str]] = None,
                 project_id: Optional[pulumi.Input[_builtins.str]] = None):
        """
        Input properties used for looking up and filtering FeedRetentionPolicy resources.
        :param pulumi.Input[_builtins.int] count_limit: The maximum number of versions per package.
        :param pulumi.Input[_builtins.int] days_to_keep_recently_downloaded_packages: The days to keep recently downloaded packages.
        :param pulumi.Input[_builtins.str] feed_id: The ID of the Feed. Changing this forces a new resource to be created.
        :param pulumi.Input[_builtins.str] project_id: The ID of the Project. If not specified, Feed will be created at the organization level. Changing this forces a new resource to be created.
        """
        if count_limit is not None:
            pulumi.set(__self__, "count_limit", count_limit)
        if days_to_keep_recently_downloaded_packages is not None:
            pulumi.set(__self__, "days_to_keep_recently_downloaded_packages", days_to_keep_recently_downloaded_packages)
        if feed_id is not None:
            pulumi.set(__self__, "feed_id", feed_id)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)

    @_builtins.property
    @pulumi.getter(name="countLimit")
    def count_limit(self) -> Optional[pulumi.Input[_builtins.int]]:
        """
        The maximum number of versions per package.
        """
        return pulumi.get(self, "count_limit")

    @count_limit.setter
    def count_limit(self, value: Optional[pulumi.Input[_builtins.int]]):
        pulumi.set(self, "count_limit", value)

    @_builtins.property
    @pulumi.getter(name="daysToKeepRecentlyDownloadedPackages")
    def days_to_keep_recently_downloaded_packages(self) -> Optional[pulumi.Input[_builtins.int]]:
        """
        The days to keep recently downloaded packages.
        """
        return pulumi.get(self, "days_to_keep_recently_downloaded_packages")

    @days_to_keep_recently_downloaded_packages.setter
    def days_to_keep_recently_downloaded_packages(self, value: Optional[pulumi.Input[_builtins.int]]):
        pulumi.set(self, "days_to_keep_recently_downloaded_packages", value)

    @_builtins.property
    @pulumi.getter(name="feedId")
    def feed_id(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The ID of the Feed. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "feed_id")

    @feed_id.setter
    def feed_id(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "feed_id", value)

    @_builtins.property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The ID of the Project. If not specified, Feed will be created at the organization level. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "project_id", value)


@pulumi.type_token("azuredevops:index/feedRetentionPolicy:FeedRetentionPolicy")
class FeedRetentionPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 count_limit: Optional[pulumi.Input[_builtins.int]] = None,
                 days_to_keep_recently_downloaded_packages: Optional[pulumi.Input[_builtins.int]] = None,
                 feed_id: Optional[pulumi.Input[_builtins.str]] = None,
                 project_id: Optional[pulumi.Input[_builtins.str]] = None,
                 __props__=None):
        """
        Manages the Feed Retention Policy within Azure DevOps.

        ## Example Usage

        ### Project Feed
        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        example = azuredevops.Project("example",
            name="Example Project",
            visibility="private",
            version_control="Git",
            work_item_template="Agile",
            description="Managed by Pulumi")
        example_feed = azuredevops.Feed("example",
            name="ExampleFeed",
            project_id=example.id)
        example_feed_retention_policy = azuredevops.FeedRetentionPolicy("example",
            project_id=example.id,
            feed_id=example_feed.id,
            count_limit=20,
            days_to_keep_recently_downloaded_packages=30)
        ```

        ### Organization Feed
        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        example = azuredevops.Feed("example", name="examplefeed")
        example_feed_retention_policy = azuredevops.FeedRetentionPolicy("example",
            feed_id=example.id,
            count_limit=20,
            days_to_keep_recently_downloaded_packages=30)
        ```

        ## Relevant Links

        - [Azure DevOps Service REST API 7.0 - Feed Management](https://learn.microsoft.com/en-us/rest/api/azure/devops/artifacts/feed-management?view=azure-devops-rest-7.0)

        ## Import

        Azure DevOps Feed Retention Policy can be imported using the Project ID and Feed ID or Feed ID e.g.:

        ```sh
        $ pulumi import azuredevops:index/feedRetentionPolicy:FeedRetentionPolicy example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
        ```

        or

        ```sh
        $ pulumi import azuredevops:index/feedRetentionPolicy:FeedRetentionPolicy example 00000000-0000-0000-0000-000000000000
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[_builtins.int] count_limit: The maximum number of versions per package.
        :param pulumi.Input[_builtins.int] days_to_keep_recently_downloaded_packages: The days to keep recently downloaded packages.
        :param pulumi.Input[_builtins.str] feed_id: The ID of the Feed. Changing this forces a new resource to be created.
        :param pulumi.Input[_builtins.str] project_id: The ID of the Project. If not specified, Feed will be created at the organization level. Changing this forces a new resource to be created.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FeedRetentionPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages the Feed Retention Policy within Azure DevOps.

        ## Example Usage

        ### Project Feed
        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        example = azuredevops.Project("example",
            name="Example Project",
            visibility="private",
            version_control="Git",
            work_item_template="Agile",
            description="Managed by Pulumi")
        example_feed = azuredevops.Feed("example",
            name="ExampleFeed",
            project_id=example.id)
        example_feed_retention_policy = azuredevops.FeedRetentionPolicy("example",
            project_id=example.id,
            feed_id=example_feed.id,
            count_limit=20,
            days_to_keep_recently_downloaded_packages=30)
        ```

        ### Organization Feed
        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        example = azuredevops.Feed("example", name="examplefeed")
        example_feed_retention_policy = azuredevops.FeedRetentionPolicy("example",
            feed_id=example.id,
            count_limit=20,
            days_to_keep_recently_downloaded_packages=30)
        ```

        ## Relevant Links

        - [Azure DevOps Service REST API 7.0 - Feed Management](https://learn.microsoft.com/en-us/rest/api/azure/devops/artifacts/feed-management?view=azure-devops-rest-7.0)

        ## Import

        Azure DevOps Feed Retention Policy can be imported using the Project ID and Feed ID or Feed ID e.g.:

        ```sh
        $ pulumi import azuredevops:index/feedRetentionPolicy:FeedRetentionPolicy example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
        ```

        or

        ```sh
        $ pulumi import azuredevops:index/feedRetentionPolicy:FeedRetentionPolicy example 00000000-0000-0000-0000-000000000000
        ```

        :param str resource_name: The name of the resource.
        :param FeedRetentionPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FeedRetentionPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 count_limit: Optional[pulumi.Input[_builtins.int]] = None,
                 days_to_keep_recently_downloaded_packages: Optional[pulumi.Input[_builtins.int]] = None,
                 feed_id: Optional[pulumi.Input[_builtins.str]] = None,
                 project_id: Optional[pulumi.Input[_builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FeedRetentionPolicyArgs.__new__(FeedRetentionPolicyArgs)

            if count_limit is None and not opts.urn:
                raise TypeError("Missing required property 'count_limit'")
            __props__.__dict__["count_limit"] = count_limit
            if days_to_keep_recently_downloaded_packages is None and not opts.urn:
                raise TypeError("Missing required property 'days_to_keep_recently_downloaded_packages'")
            __props__.__dict__["days_to_keep_recently_downloaded_packages"] = days_to_keep_recently_downloaded_packages
            if feed_id is None and not opts.urn:
                raise TypeError("Missing required property 'feed_id'")
            __props__.__dict__["feed_id"] = feed_id
            __props__.__dict__["project_id"] = project_id
        super(FeedRetentionPolicy, __self__).__init__(
            'azuredevops:index/feedRetentionPolicy:FeedRetentionPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            count_limit: Optional[pulumi.Input[_builtins.int]] = None,
            days_to_keep_recently_downloaded_packages: Optional[pulumi.Input[_builtins.int]] = None,
            feed_id: Optional[pulumi.Input[_builtins.str]] = None,
            project_id: Optional[pulumi.Input[_builtins.str]] = None) -> 'FeedRetentionPolicy':
        """
        Get an existing FeedRetentionPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[_builtins.int] count_limit: The maximum number of versions per package.
        :param pulumi.Input[_builtins.int] days_to_keep_recently_downloaded_packages: The days to keep recently downloaded packages.
        :param pulumi.Input[_builtins.str] feed_id: The ID of the Feed. Changing this forces a new resource to be created.
        :param pulumi.Input[_builtins.str] project_id: The ID of the Project. If not specified, Feed will be created at the organization level. Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FeedRetentionPolicyState.__new__(_FeedRetentionPolicyState)

        __props__.__dict__["count_limit"] = count_limit
        __props__.__dict__["days_to_keep_recently_downloaded_packages"] = days_to_keep_recently_downloaded_packages
        __props__.__dict__["feed_id"] = feed_id
        __props__.__dict__["project_id"] = project_id
        return FeedRetentionPolicy(resource_name, opts=opts, __props__=__props__)

    @_builtins.property
    @pulumi.getter(name="countLimit")
    def count_limit(self) -> pulumi.Output[_builtins.int]:
        """
        The maximum number of versions per package.
        """
        return pulumi.get(self, "count_limit")

    @_builtins.property
    @pulumi.getter(name="daysToKeepRecentlyDownloadedPackages")
    def days_to_keep_recently_downloaded_packages(self) -> pulumi.Output[_builtins.int]:
        """
        The days to keep recently downloaded packages.
        """
        return pulumi.get(self, "days_to_keep_recently_downloaded_packages")

    @_builtins.property
    @pulumi.getter(name="feedId")
    def feed_id(self) -> pulumi.Output[_builtins.str]:
        """
        The ID of the Feed. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "feed_id")

    @_builtins.property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[Optional[_builtins.str]]:
        """
        The ID of the Project. If not specified, Feed will be created at the organization level. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "project_id")

