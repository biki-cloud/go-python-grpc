from __future__ import print_function

import logging

import grpc

import hello_pb2
import hello_pb2_grpc


def do_request():
    msg = "bikibiki"
    with grpc.insecure_channel("localhost:50051") as channel:
        stub = hello_pb2_grpc.HelloServiceStub(channel)
        stub.Hello(hello_pb2.MsgRequest(message=msg))


if __name__ == '__main__':
    logging.basicConfig()
    do_request()
