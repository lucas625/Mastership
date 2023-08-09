import requests
from django.conf import settings
from rest_framework.response import Response
from rest_framework.status import HTTP_200_OK
from rest_framework.status import HTTP_500_INTERNAL_SERVER_ERROR
from rest_framework.views import APIView


SHOULD_PING_KEY = 'shouldPing'

class PingView(APIView):

    http_method_names = ['post']

    def post(self, request, *args, **kwargs):
        if SHOULD_PING_KEY in request.data and not request.data[SHOULD_PING_KEY]:
            settings.LOGGER.info("No extra ping was necessary, returning")
            return Response(status=HTTP_200_OK, data={'message': "ping went ok"})

        settings.LOGGER.info("Will continue ping")

        data = {
            SHOULD_PING_KEY: False,
        }

        headers = {
            'Content-type': 'application/json',
        }

        try:
            security_checker_response = requests.post(
                f'{settings.SECURITY_CHECKER_ADDRESS}/api/security-checker/ping',
                json=data,
                headers=headers
            )
        except requests.RequestException as exc:
            data = {'details': f'Request Error: {repr(exc)};;; class: {exc.__class__()}'}
            return Response(status=HTTP_500_INTERNAL_SERVER_ERROR, data=data)
        except Exception as exc:
            data = {'details': f'Unknown Error: {repr(exc)}'}
            return Response(status=HTTP_500_INTERNAL_SERVER_ERROR, data=data)

        try:
            response_data = security_checker_response.json()
        except Exception:
            response_data = f'response as string: {str(security_checker_response)}'

        return Response(status=security_checker_response.status_code, data=response_data)
