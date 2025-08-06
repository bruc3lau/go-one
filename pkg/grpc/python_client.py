import grpc
from helloworld import helloworld_pb2
from helloworld import helloworld_pb2_grpc


def run():
    # NOTE: The server address and port must match the server's configuration.
    with grpc.insecure_channel('localhost:50051') as channel:
        stub = helloworld_pb2_grpc.GreeterStub(channel)
        response = stub.SayHello(helloworld_pb2.HelloRequest(name='you'))
    print("Greeter client received: " + response.message)


if __name__ == '__main__':
    run()
