from rest_framework.response import Response
from rest_framework.status import HTTP_200_OK
from rest_framework.views import APIView

from core.business import AnalyzerBusiness
from core.serializers import AnalyzerSerializer


class AnalyzerView(APIView):

    http_method_names = ['post']

    def post(self, request, *args, **kwargs):
        serializer = AnalyzerSerializer(data=request.data)
        serializer.is_valid(raise_exception=True)
        analyzer = serializer.save()
        business = AnalyzerBusiness(analyzer)
        csv_str = business.do_analysis()
        return Response({'csv_data': csv_str}, HTTP_200_OK)
