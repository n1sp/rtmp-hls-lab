FROM ubuntu:22.04

# パッケージ更新 + nginx,nginx-rtmpをインストール
RUN apt update && \
    apt install -y nginx libnginx-mod-rtmp && \
    apt clean

# HLS 出力先ディレクトリを作成
RUN mkdir -p /tmp/hls && chmod 777 /tmp/hls

# nginx.conf を配置
COPY conf/nginx.conf /etc/nginx/nginx.conf

# ポート開放
EXPOSE 1935 8080

# nginx をフォアグラウンドで起動
CMD ["nginx", "-g", "daemon off;"]
