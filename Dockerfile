# This dockerfile uses extends image https://hub.docker.com/_/golang
# VERSION v1.0.0
# Author: sinlov
# dockerfile offical document https://docs.docker.com/engine/reference/builder/
# https://hub.docker.com/_/golang?tab=description
FROM golang:1.13.8-alpine

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk --no-cache add make git gcc libtool musl-dev

COPY $PWD /usr/src/myapp
WORKDIR /usr/src/myapp
RUN make initDockerImagesMod


CMD ["tail",  "-f", "/etc/alpine-release"]
#ENTRYPOINT [ "go", "env" ]

