FROM python:3.8

RUN mkdir -p /app
WORKDIR /app

RUN pip install pipenv
COPY Pipfile Pipfile.lock .
RUN pipenv install

COPY gateway.py tweet-gateway.proto .

RUN pipenv run python -m grpc_tools.protoc --proto_path=. --python_out=. --grpc_python_out=. tweet-gateway.proto

CMD [ "pipenv", "run", "python", "-u", "gateway.py" ]
