import cv2
import numpy as np
import onnxruntime

class OnnxModel():
    def __init__(self, onnxpath):
        self.onnx_session = onnxruntime.InferenceSession(onnxpath)
        self.input_name = [node.name for node in self.onnx_session.get_inputs()]
        self.output_name = [node.name for node in self.onnx_session.get_outputs()]

    def get_input_feed(self, img_tensor):
        input_feed = {}
        for name in self.input_name:
            input_feed[name] = img_tensor
        return input_feed

    def inference(self, body):
        data = np.frombuffer(body, np.uint8)
        img = cv2.imdecode(data, cv2.IMREAD_COLOR)
        # img = cv2.imread(img_path)
        or_img = cv2.resize(img, (640, 640))
        img = or_img[:, :, ::-1].transpose(2, 0, 1)
        img = img.astype(dtype=np.float32) / 255.0
        img = np.expand_dims(img, axis=0)
        input_feed = self.get_input_feed(img)
        return self.onnx_session.run(None, input_feed)[0]

    @staticmethod
    def nms(dets, thresh):
        x1 = dets[:, 0]
        y1 = dets[:, 1]
        x2 = dets[:, 2]
        y2 = dets[:, 3]
        areas = (y2 - y1 + 1) * (x2 - x1 + 1)
        scores = dets[:, 4]
        keep = []
        index = scores.argsort()[::-1]
        while index.size > 0:
            i = index[0]
            keep.append(i)
            x11 = np.maximum(x1[i], x1[index[1:]])
            y11 = np.maximum(y1[i], y1[index[1:]])
            x22 = np.minimum(x2[i], x2[index[1:]])
            y22 = np.minimum(y2[i], y2[index[1:]])
            w = np.maximum(0, x22 - x11 + 1)
            h = np.maximum(0, y22 - y11 + 1)
            overlaps = w * h
            ious = overlaps / (areas[i] + areas[index[1:]] - overlaps)
            idx = np.where(ious <= thresh)[0]
            index = index[idx + 1]
        return keep

    @staticmethod
    def xywh2xyxy(x):
        # [x, y, w, h] to [x1, y1, x2, y2]
        y = np.copy(x)
        y[:, 0] = x[:, 0] - x[:, 2] / 2
        y[:, 1] = x[:, 1] - x[:, 3] / 2
        y[:, 2] = x[:, 0] + x[:, 2] / 2
        y[:, 3] = x[:, 1] + x[:, 3] / 2
        return y


    def filter_box(self, org_box, conf_thres):  # 过滤掉无用的框
        org_box = np.squeeze(org_box)
        conf = org_box[..., 4] > conf_thres
        box = org_box[conf == True]
        cls_cinf = box[..., 5:]
        cls = []
        for i in range(len(cls_cinf)):
            cls.append(int(np.argmax(cls_cinf[i])))
        all_cls = list(set(cls))
        output = []
        for i in range(len(all_cls)):
            curr_cls = all_cls[i]
            curr_cls_box = []
            curr_out_box = []
            for j in range(len(cls)):
                if cls[j] == curr_cls:
                    box[j][5] = curr_cls
                    curr_cls_box.append(box[j][:6])
            curr_cls_box = np.array(curr_cls_box)
            curr_cls_box = self.xywh2xyxy(curr_cls_box)
            curr_out_box = self.nms(curr_cls_box, conf_thres)
            # yield curr_cls_box
            for k in curr_out_box:
                yield curr_cls_box[k]
                # output.append(curr_cls_box[k])
        # return output

    def dector(self, filepath: str, threshold=0.7) -> list[list]:
        output = self.inference(filepath)
        return self.filter_box(output, threshold)

if __name__ == "__main__":
    model = OnnxModel('model.onnx')
    with open("test.png", "rb") as fp:
        output = model.inference(fp.read())
    result = model.filter_box(output, 0.7)
    for i in result:
        print(i)
