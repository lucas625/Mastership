from rest_framework.views import APIView


class AnalyserView(APIView):

    http_method_names = ['post']

    def post(self, request, *args, **kwargs):
        # TODO: Validate data
        # TODO: call business
        pass
