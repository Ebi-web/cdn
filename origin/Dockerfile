# ステージ1: ビルド環境
FROM golang:1.20.5 AS builder

WORKDIR /app
COPY . .

# 依存関係のダウンロードとアプリケーションのビルド
RUN go mod download
RUN go build -o origin-server

# ステージ2: 実行環境
FROM golang:1.20.5

WORKDIR /app
COPY --from=builder /app/origin-server .
COPY --from=builder /app/computer_man1_smile.png .

# サーバーの起動
CMD ["./origin-server"]
