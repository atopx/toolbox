import grpc
from lib import OnnxModel, ccdector_pb2 as pb, ccdector_pb2_grpc as rpc


def run():
    with grpc.insecure_channel('localhost:50051') as channel:
        f = open("test.png", 'rb')  # 读取
        data = f.read()
        f.close()

        # 调用rpc
        response = rpc.DectorStub(channel).Predict(pb.Params(file=data))

    for i in response.data:
        print(i)


if __name__ == '__main__':
    run()