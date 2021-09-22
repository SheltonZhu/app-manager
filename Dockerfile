FROM golang:alpine as builder
ENV GOPROXY https://goproxy.cn
ENV GO111MODULE on
ENV TZ Asia/Shanghai
WORKDIR /go/src/app
COPY . .
RUN go get -v
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

FROM  alpine:latest as prod
ENV GOPROXY https://goproxy.cn
ENV GO111MODULE on
ENV TZ Asia/Shanghai
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories
RUN apk update --no-cache
RUN apk add --no-cache ca-certificates
RUN apk add --no-cache tzdata

WORKDIR /root/
COPY --from=builder /go/src/app/docker-api .
ENV GIN_MODE="release"
EXPOSE 4399
CMD ["./docker-api"]