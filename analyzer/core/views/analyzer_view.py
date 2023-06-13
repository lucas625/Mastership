import http

import requests
from rest_framework.response import Response
from rest_framework.status import HTTP_200_OK
from rest_framework.status import HTTP_500_INTERNAL_SERVER_ERROR
from rest_framework.views import APIView

from core.business import AnalyzerBusiness
from core.serializers import AnalyzerSerializer


class AnalyzerView(APIView):

    http_method_names = ['get', 'post']

    def get(self, request, *args, **kwargs):
        """
        A dummy method for pinging the calculator service when using kubernetes, to see if access to it is enabled or not.
        """
        data = {
            'firstNumber': 10,
            'secondNumber': 20
        }
        try:
            response = requests.post('http://msc-calculator:8000/add', json=data)
        except Exception as exc:
            response = Response(status=HTTP_500_INTERNAL_SERVER_ERROR, data=exc)
        return response

    def post(self, request, *args, **kwargs):
        serializer = AnalyzerSerializer(data=request.data)
        serializer.is_valid(raise_exception=True)
        analyzer = serializer.save()
        business = AnalyzerBusiness(analyzer)
        csv_str = business.do_analysis()
        return Response({'csv_data': csv_str}, HTTP_200_OK)
