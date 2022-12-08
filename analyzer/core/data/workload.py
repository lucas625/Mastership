from __future__ import annotations


class Workload:

    def __init__(self, interactions: int, interval_between_batches_in_milliseconds: int, batch_size: int):
        self._interactions = interactions
        self._interval_between_batches_in_milliseconds = interval_between_batches_in_milliseconds
        self._batch_size = batch_size

    def __str__(self) -> str:
        return (
            f'interactions: {self.interactions}, '
            f'interval: {self.interval_between_batches_in_milliseconds}, '
            f'batch size: {self.batch_size}'
        )

    def __repr__(self) -> str:
        return (
            f'interactions: {self.interactions}, '
            f'interval: {self.interval_between_batches_in_milliseconds}, '
            f'batch size: {self.batch_size}'
        )

    def __eq__(self, other: Workload) -> bool:
        are_equal = isinstance(other, Workload)
        if are_equal:
            are_equal = (
                self.interactions == other.interactions and
                self.interval_between_batches_in_milliseconds == other.interval_between_batches_in_milliseconds and
                self.batch_size == other.batch_size
            )
        return are_equal

    @property
    def interactions(self) -> int:
        return self._interactions

    @property
    def interval_between_batches_in_milliseconds(self) -> int:
        return self._interval_between_batches_in_milliseconds

    @property
    def batch_size(self) -> int:
        return self._batch_size
