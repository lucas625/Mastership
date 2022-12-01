#!/bin/sh
python manage.py migrate
gunicorn --bind 0.0.0.0:8082 --log-leve info --timeout 300 --workers 4 analyser.wsgi
