# ベースイメージ：Go が入ってる公式イメージをそのまま使う（単一ステージ）
FROM golang:1.25.4-bookworm

WORKDIR /app

# モジュールファイル（あれば）を先にコピーしておくと依存ダウンロードが速くなるけど、
# 最低限ならまとめてコピーしても問題ない
COPY . .

# ビルド
RUN go build -o /go-backend main.go

# ポート公開（ドキュメント的に）
EXPOSE 8081

# 実行
CMD ["/go-backend"]
