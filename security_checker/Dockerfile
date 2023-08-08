FROM python:3.9-slim

# Install dependencies only
RUN pip install --upgrade pip
COPY requirements.txt /security_checker/
RUN pip install -r /security_checker/requirements.txt

# Setup the rest of the project to be executed
COPY . /security_checker
WORKDIR /security_checker

# Env vars
ENV SECRET_KEY ""
ENV DEBUG false
ENV CALCULATOR_ADDRESS "http://msc-calculator:8000"
ENV SECURITY_CHECKER_ADDRESS "http://msc-security-checker:8000"

ENTRYPOINT ["./entrypoint.sh"]
