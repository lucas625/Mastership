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
        try:
            calculator_response = requests.post(f'{settings.CALCULATOR_ADDRESS}/add', json=data)
            response = Response(status=calculator_response.status_code, data=calculator_response.json())
        except requests.ConnectionError as exc:
            data = {'details': f'Connection Error: {str(exc)}'}
            response = Response(status=HTTP_500_INTERNAL_SERVER_ERROR, data=data)
        except requests.HTTPError as exc:
            data = {'details': f'HTTP Error: {str(exc)}'}
            response = Response(status=HTTP_500_INTERNAL_SERVER_ERROR, data=data)
        except requests.Timeout as exc:
            data = {'details': f'Timeout Error: {str(exc)}'}
            response = Response(status=HTTP_500_INTERNAL_SERVER_ERROR, data=data)
        except requests.TooManyRedirects as exc:
            data = {'details': f'Too many redirects Error: {str(exc)}'}
            response = Response(status=HTTP_500_INTERNAL_SERVER_ERROR, data=data)
        except requests.RequestException as exc:
            data = {'details': f'Request Error: {str(exc)}'}
            response = Response(status=HTTP_500_INTERNAL_SERVER_ERROR, data=data)
        except Exception as exc:
            data = {'details': f'Unknown Error: {str(exc)}'}
            response = Response(status=HTTP_500_INTERNAL_SERVER_ERROR, data=data)

        return response
