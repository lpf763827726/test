FROM ubuntu:22.04

RUN apt-get clean
RUN apt-get update --fix-missing
RUN apt-get upgrade -y
RUN apt-get install ca-certificates -y
RUN apt-get install gnupg2 -y

# RUN apt-get install libssl-dev -y
RUN apt-get install curl -y
# RUN ln -s /usr/lib/libssl.so /usr/lib/libssl.so.1.1
RUN curl -O http://nz2.archive.ubuntu.com/ubuntu/pool/main/o/openssl/libssl1.1_1.1.1f-1ubuntu2.20_amd64.deb
RUN dpkg -i libssl1.1_1.1.1f-1ubuntu2.20_amd64.deb

EXPOSE 8080

WORKDIR /app

ADD target/x86_64-unknown-linux-gnu/release/gitfog-master /app/gitfog-master
ADD migrations /app/migrations
ADD sqlx /app/sqlx

ENTRYPOINT ["./gitfog-master"]