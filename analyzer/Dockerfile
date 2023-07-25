FROM python:3.9-slim

# Install dependencies only
RUN pip install --upgrade pip
COPY requirements.txt /analyzer/
RUN pip install -r /analyzer/requirements.txt

# Setup the rest of the project to be executed
COPY . /analyzer
WORKDIR /analyzer

# Env vars
ENV SECRET_KEY ""
ENV DEBUG false
ENV CALCULATOR_ADDRESS "http://msc-calculator:8000"

ENTRYPOINT ["./entrypoint.sh"]
