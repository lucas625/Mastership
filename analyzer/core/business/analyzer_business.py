import io
import csv
from typing import List

import numpy

from core.data import Analyzer

_RTT_90_KEY = 'rtt90'
_RTT_95_KEY = 'rtt95'
_RTT_99_KEY = 'rtt99'


class AnalyzerBusiness:

    def __init__(self, data: Analyzer):
        self._data = data

    def do_analysis(self) -> str:
        numpy_array = numpy.array(self._data.rtts_in_milliseconds)
        row = {
            _RTT_90_KEY: numpy.percentile(numpy_array, 90),
            _RTT_95_KEY: numpy.percentile(numpy_array, 95),
            _RTT_99_KEY: numpy.percentile(numpy_array, 99)
        }
        return self._write_csv(row)

    def _write_csv(self, row: dict) -> str:
        output = io.StringIO()
        writer = csv.DictWriter(output, fieldnames=self._get_headers())
        writer.writeheader()
        writer.writerow(row)
        return output.getvalue()

    @staticmethod
    def _get_headers() -> List[str]:
        return [_RTT_90_KEY, _RTT_95_KEY, _RTT_99_KEY]

