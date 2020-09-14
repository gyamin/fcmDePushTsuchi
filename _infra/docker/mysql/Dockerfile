# ベースイメージ
FROM mysql:5.7

MAINTAINER gaymin

# ENV設定
ENV MYSQL_ROOT_USER root
ENV MYSQL_ROOT_PASSWORD root
ENV MYSQL_DATABASE go_db
ENV MYSQL_USER admin
ENV MYSQL_PASSWORD 1234Qwer

# 設定ファイルコピー
ADD conf/my.cnf /etc/my.cnf

VOLUME /var/lib/mysql