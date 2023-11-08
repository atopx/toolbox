import sys
import grpc
import asyncio
from concurrent import futures

from lib import OnnxModel, ccdector_pb2 as pb, ccdector_pb2_grpc as rpc


class Dector(rpc.DectorServicer):

    def __init__(self, onnx_model):
        self.onnx = onnx_model

    def Predict(self, params, context):
        output = self.onnx.inference(params.file)
        result = self.onnx.filter_box(output, 0.7)
        reply = pb.Reply(data=[])
        for (left, top, right, bottom, confidence, label) in result:
            record = pb.Reply.Record(
                label=int(label),
                confidence=float(confidence),
                left=float(left),
                top=float(top),
                right=float(right),
                bottom=float(bottom),
            )
            reply.data.append(record)
        return reply


async def start_server():
    # start rpc service
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=40), options=[
        ('grpc.max_send_message_length', 64 * 1024 * 1024),
        ('grpc.max_receive_message_length', 64 * 1024 * 1024),
        ('grpc.enable_retries', 1),
    ])
    rpc.add_DectorServicer_to_server(
        Dector(OnnxModel('model.onnx')), server)  # 加入服务

    server.add_insecure_port('[::]:20088')
    server.start()
    print("start service listen [::]:20088")
    server.wait_for_termination()


if __name__ == '__main__':
    if sys.version_info < (3, 10):
        loop = asyncio.get_event_loop()
    else:
        loop = asyncio.new_event_loop()
    loop.run_until_complete(asyncio.wait([start_server()]))
    loop.close()
