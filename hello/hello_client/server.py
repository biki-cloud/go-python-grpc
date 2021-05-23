from concurrent import futures

import grpc

import hello_pb2
import hello_pb2_grpc


class HelloServiceServer(hello_pb2_grpc.HelloServiceServicer):
    def __init__(self):
        pass

    def Hello(self, request: hello_pb2.MsgRequest, context):
        msg = request.message
        print("Hello function is invoked from client")
        print(f"recieved {msg}")
        return hello_pb2.MsgResponse(
            result=f"recieved {msg} thanks to client"
        )


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    hello_pb2_grpc.add_HelloServiceServicer_to_server(HelloServiceServer(), server)
    server.add_insecure_port('[::]:50051')
    server.start()
    print('Starting gRPC sample server...')
    server.wait_for_termination()


if __name__ == '__main__':
    serve()
