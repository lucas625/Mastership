from __future__ import annotations

from typing import List

from core.data.workload import Workload


class Analyzer:

    def __init__(
            self,
            label: str,
            workload: Workload,
            mean: float,
            standard_deviation: float,
            failures: int,
            rtts_in_milliseconds: List[float]
    ):
        self._label = label
        self._workload = workload
        self._mean = mean
        self._standard_deviation = standard_deviation
        self._failures = failures
        self._rtts_in_milliseconds = rtts_in_milliseconds

    def __str__(self) -> str:
        return f'{self.label} - {str(self.workload)}'

    def __repr__(self) -> str:
        return f'{self.label} - {repr(self.workload)}'

    def __eq__(self, other: Analyzer) -> bool:
        are_equal = isinstance(other, Analyzer)
        if are_equal:
            are_equal = (
                    self.label == other.label and
                    self.workload == other.workload and
                    self.mean == other.mean and
                    self.standard_deviation == other.standard_deviation and
                    self.failures == other.failures and
                    self.rtts_in_milliseconds == other.rtts_in_milliseconds
            )
        return are_equal

    @property
    def label(self) -> str:
        return self._label

    @property
    def workload(self) -> Workload:
        return self._workload

    @property
    def mean(self) -> float:
        return self._mean

    @property
    def standard_deviation(self) -> float:
        return self._standard_deviation

    @property
    def failures(self) -> int:
        return self._failures

    @property
    def rtts_in_milliseconds(self) -> List[float]:
        return self._rtts_in_milliseconds
