FROM alpine:3.14.0

ENV TZ=Asia/Tehran \
    PATH="/app:${PATH}" \
    PYTHONUNBUFFERED=1

RUN apk add --update --no-cache \
      ca-certificates \
      bash \
      openssl \
      net-tools \
      curl \
      busybox-extras \
      mtr \
      iputils \
      bind-tools \
      py3-pip \
      jq

RUN pip3 install --no-cache-dir httpstat

RUN wget https://github.com/codesenberg/bombardier/releases/download/v1.2.5/bombardier-linux-amd64 && \
      mv bombardier-linux-amd64 bombardier && install bombardier /usr/bin && rm bombardier

ENTRYPOINT ["/bin/bash"]