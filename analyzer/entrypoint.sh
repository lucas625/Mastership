#!/bin/sh
python manage.py migrate
gunicorn --bind 0.0.0.0:8000 --log-leve info --timeout 300 --workers 4 analyzer.wsgi
