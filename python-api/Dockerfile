FROM python:3.8-alpine

WORKDIR /app

COPY requirements.txt .
# コンテナ内で必要なパッケージをインストール
RUN apk add --no-cache build-base \
 && apk --no-cache add postgresql-dev \
 && python3 -m pip install psycopg2 \
 && pip3 install --no-cache-dir --trusted-host pypi.python.org -r requirements.txt \
 && apk del build-base

EXPOSE 8000
# FastAPIを8000ポートで待機
CMD ["uvicorn", "main:app", "--reload", "--host", "0.0.0.0", "--port", "8000"]