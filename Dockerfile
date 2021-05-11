#docker build --build-arg GOPROXY=https://goproxy.cn -t liuguangw/forumx:latest .
FROM golang:latest AS builder
ARG GOPROXY
ARG CGO_ENABLED=0
WORKDIR /home
ADD . forumx
RUN cd forumx && go env && make

FROM alpine:latest
#RUN apk --no-cache add ca-certificates
WORKDIR /home/forumx
COPY --from=builder /home/forumx/forumx /home/forumx/.env /home/forumx/LICENSE  /home/forumx/
EXPOSE 3000
CMD ["./forumx"]
