from concurrent import futures
import dataclasses
import json
import logging
import os
from typing import Dict

import grpc
from grpc_reflection.v1alpha import reflection
import twitter

import tweet_gateway_pb2 
import tweet_gateway_pb2_grpc


@dataclasses.dataclass
class Credentials:
    consumer_key: str
    consumer_secret: str
    access_token_key: str
    access_token_secret: str


class TweetGateway(tweet_gateway_pb2_grpc.TweetGateway):
    def __init__(self, config: Dict[str, Credentials]):
        self.config: Dict[str, twitter.Api] = {}
        for user in config:
            cred = config[user]
            self.config[user] = twitter.Api(
                    consumer_key=cred.consumer_key,
                    consumer_secret=cred.consumer_secret,
                    access_token_key=cred.access_token_key,
                    access_token_secret=cred.access_token_secret,
                    sleep_on_rate_limit=True)

    def tweet(self, request: tweet_gateway_pb2.TweetRequest, context):
        username = request.username
        content = request.content

        if username == "":
            context.abort(grpc.StatusCode.INVALID_ARGUMENT, "username must not be empty")

        if content == "":
            context.abort(grpc.StatusCode.INVALID_ARGUMENT, "content must not be empty")

        if not username in self.config:
            context.abort(grpc.StatusCode.INVALID_ARGUMENT, "username is not in config")

        api = self.config[username]
        api.PostUpdate(content)

        return tweet_gateway_pb2.TweetResponse()


def serve():
    CONFIG = os.getenv("TG_CONFIG")
    with open(CONFIG, "r") as f:
        config = json.loads(f.read())

    config = {
        user: Credentials(
            cred["consumer_key"],
            cred["consumer_secret"],
            cred["access_token_key"],
            cred["access_token_secret"]) for user, cred in config.items()
    }

    server = grpc.server(futures.ThreadPoolExecutor(max_workers=16))
    servicer = TweetGateway(config)
    tweet_gateway_pb2_grpc.add_TweetGatewayServicer_to_server(servicer, server)

    SERVICE_NAMES = (
        tweet_gateway_pb2.DESCRIPTOR.services_by_name["TweetGateway"].full_name,
        reflection.SERVICE_NAME,
    )
    reflection.enable_server_reflection(SERVICE_NAMES, server)

    server.add_insecure_port("0.0.0.0:31337")
    server.start()
    server.wait_for_termination()

if __name__ == "__main__":
    logging.basicConfig()
    serve()

