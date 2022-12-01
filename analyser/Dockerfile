FROM python:3.9-slim

# Install dependencies only
RUN pip install --upgrade pip
COPY requirements.txt /analyser/
RUN pip install -r /analyser/requirements.txt

# Setup the rest of the project to be executed
COPY . /analyser
WORKDIR /analyser

# Env vars
ENV SECRET_KEY ""
ENV DEBUG false

ENTRYPOINT ["./entrypoint.sh"]
