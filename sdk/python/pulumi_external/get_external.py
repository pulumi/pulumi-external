# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetExternalResult',
    'AwaitableGetExternalResult',
    'get_external',
    'get_external_output',
]

@pulumi.output_type
class GetExternalResult:
    """
    A collection of values returned by getExternal.
    """
    def __init__(__self__, id=None, programs=None, query=None, result=None, working_dir=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if programs and not isinstance(programs, list):
            raise TypeError("Expected argument 'programs' to be a list")
        pulumi.set(__self__, "programs", programs)
        if query and not isinstance(query, dict):
            raise TypeError("Expected argument 'query' to be a dict")
        pulumi.set(__self__, "query", query)
        if result and not isinstance(result, dict):
            raise TypeError("Expected argument 'result' to be a dict")
        pulumi.set(__self__, "result", result)
        if working_dir and not isinstance(working_dir, str):
            raise TypeError("Expected argument 'working_dir' to be a str")
        pulumi.set(__self__, "working_dir", working_dir)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The id of the data source. This will always be set to `-`
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def programs(self) -> Sequence[str]:
        return pulumi.get(self, "programs")

    @property
    @pulumi.getter
    def query(self) -> Optional[Mapping[str, str]]:
        """
        A map of string values to pass to the external program as the query arguments. If not supplied, the program will receive an empty object as its input.
        """
        return pulumi.get(self, "query")

    @property
    @pulumi.getter
    def result(self) -> Mapping[str, str]:
        """
        A map of string values returned from the external program.
        """
        return pulumi.get(self, "result")

    @property
    @pulumi.getter(name="workingDir")
    def working_dir(self) -> Optional[str]:
        """
        Working directory of the program. If not supplied, the program will run in the current directory.
        """
        return pulumi.get(self, "working_dir")


class AwaitableGetExternalResult(GetExternalResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetExternalResult(
            id=self.id,
            programs=self.programs,
            query=self.query,
            result=self.result,
            working_dir=self.working_dir)


def get_external(programs: Optional[Sequence[str]] = None,
                 query: Optional[Mapping[str, str]] = None,
                 working_dir: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetExternalResult:
    """
    Use this data source to access information about an existing resource.

    :param Mapping[str, str] query: A map of string values to pass to the external program as the query arguments. If not supplied, the program will receive an empty object as its input.
    :param str working_dir: Working directory of the program. If not supplied, the program will run in the current directory.
    """
    __args__ = dict()
    __args__['programs'] = programs
    __args__['query'] = query
    __args__['workingDir'] = working_dir
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('external:index/getExternal:getExternal', __args__, opts=opts, typ=GetExternalResult).value

    return AwaitableGetExternalResult(
        id=pulumi.get(__ret__, 'id'),
        programs=pulumi.get(__ret__, 'programs'),
        query=pulumi.get(__ret__, 'query'),
        result=pulumi.get(__ret__, 'result'),
        working_dir=pulumi.get(__ret__, 'working_dir'))


@_utilities.lift_output_func(get_external)
def get_external_output(programs: Optional[pulumi.Input[Sequence[str]]] = None,
                        query: Optional[pulumi.Input[Optional[Mapping[str, str]]]] = None,
                        working_dir: Optional[pulumi.Input[Optional[str]]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetExternalResult]:
    """
    Use this data source to access information about an existing resource.

    :param Mapping[str, str] query: A map of string values to pass to the external program as the query arguments. If not supplied, the program will receive an empty object as its input.
    :param str working_dir: Working directory of the program. If not supplied, the program will run in the current directory.
    """
    ...
