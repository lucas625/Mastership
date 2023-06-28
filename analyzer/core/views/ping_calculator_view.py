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
        headers = {
            'Content-type': 'application/json',
            'Accept': 'text/plain'
        }
        try:
            calculator_response = requests.post(f'{settings.CALCULATOR_ADDRESS}/add', json=data, headers=headers)
        except requests.RequestException as exc:
            data = {'details': f'Request Error: {repr(exc)};;; class: {exc.__class__()}'}
            return Response(status=HTTP_500_INTERNAL_SERVER_ERROR, data=data)
        except Exception as exc:
            data = {'details': f'Unknown Error: {repr(exc)}'}
            return Response(status=HTTP_500_INTERNAL_SERVER_ERROR, data=data)

        try:
            response_data = calculator_response.json()
        except Exception:
            response_data = f'response as string: {str(calculator_response)}'

        return Response(status=calculator_response.status_code, data=response_data)
