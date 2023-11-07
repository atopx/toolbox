# FROM tiangolo/uvicorn-gunicorn:python3.9
FROM tiangolo/uvicorn-gunicorn:python3.9-slim

LABEL maintainer="atopx <3940422@qq.com>"

WORKDIR /app

# RUN sed -i 's/http:\/\/deb.debian.org/http:\/\/mirrors.aliyun.com/g' /etc/apt/sources.list.d/debian.sources
# RUN apt update && apt install -y libgl1-mesa-glx libglib2.0-dev
# RUN apt clean

COPY requirements.txt /tmp/requirements.txt
RUN pip install --no-cache-dir -r /tmp/requirements.txt \
    -i http://mirrors.aliyun.com/pypi/simple \
    --trusted-host mirrors.aliyun.com
COPY . /app

CMD ["uvicorn", "main:app", "--host", "0.0.0.0"]