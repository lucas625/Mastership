import requests
from django.conf import settings
from rest_framework.response import Response
from rest_framework.status import HTTP_500_INTERNAL_SERVER_ERROR
from rest_framework.views import APIView


class PingCalculatorView(APIView):

    http_method_names = ['post']

    def post(self, request, *args, **kwargs):
        data = {
            'firstNumber': 10,
            'secondNumber': 20
        }
        calculator_response = None
        try:
            calculator_response = requests.post(f'{settings.CALCULATOR_ADDRESS}/add', json=data)
            response = Response(status=calculator_response.status_code, data=calculator_response.json())
        except Exception as exc:
            print(f'Error: {exc}')
            data = (
                {'details': str(exc), 'status_code': calculator_response.status_code} if
                calculator_response else
                {'details': str(exc), 'status_code': 'unknown'}
            )
            response = Response(status=HTTP_500_INTERNAL_SERVER_ERROR, data=data)

        return response
