from fastapi import FastAPI, UploadFile, File
from lib import OnnxModel

app = FastAPI()
onnx = OnnxModel("model.onnx")

CLASSES = [
    'BLACK_KNIGHT',
    'BLACK_BISHOP',
    'BLACK_ADVISOR',
    'BLACK_KING',
    'BLACK_ROOK',
    'BLACK_CANNON',
    'BLACK_PAWN',
    'WHITE_ROOK',
    'WHITE_KNIGHT',
    'WHITE_ADVISOR',
    'WHITE_KING',
    'WHITE_BISHOP',
    'WHITE_CANNON',
    'WHITE_PAWN',
    'BROAD'
]


@app.post("/dector")
async def dector(file: UploadFile = File(...)):
    body = await file.read()
    await file.close()
    output = onnx.inference(body)
    data = onnx.filter_box(output, 0.7)
    result = []
    for (left, top, right, bottom, confidence, label) in data:
        result.append({
            "left": float(left),
            "top": float(top),
            "right": float(right),
            "bottom": float(bottom),
            "confidence": float(confidence),
            "label": int(label),
        })

    return {"data": result}
