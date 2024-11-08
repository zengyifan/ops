# Builder Stage
FROM golang:1.23-alpine as builder

# 设置工作目录
WORKDIR /app

# 安装构建依赖
RUN apk add --no-cache git make

# 复制 go mod 和 sum 文件
RUN cd ..
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

# copy source code
COPY . .

RUN make build

# Distribution Stage
FROM golang:1.23-alpine

# 设置工作目录
WORKDIR /app

# 从构建者阶段复制构建好的二进制文件
COPY --from=builder /app/_output/platforms/linux/amd64/app /app/app
COPY --from=builder /app/configs/config.yaml /etc/app/config.yaml

# 暴露应用端口
EXPOSE 8888

# 为最终镜像设置运行时命令
ENTRYPOINT ["/app/app"]
CMD ["-c", "/etc/app/config.yaml"]




