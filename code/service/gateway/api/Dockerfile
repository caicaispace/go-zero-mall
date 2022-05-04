### 加载基础镜像
FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOOS linux
### 设置 go module 代理
ENV GOPROXY https://goproxy.cn,direct

WORKDIR /build/zero

### 下载依赖文件
ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
COPY service/gateway/api/etc /app/etc
### 编译源代码
RUN go build -ldflags="-s -w" -o /app/gateway service/gateway/api/gateway.go

### 生成docker镜像
FROM alpine

RUN apk update --no-cache && apk add --no-cache ca-certificates tzdata
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/gateway /app/gateway
COPY --from=builder /app/etc /app/etc

CMD ["./gateway", "-f", "etc/gateway.prod.yaml"]
