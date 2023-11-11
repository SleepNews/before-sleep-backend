# 基础镜像选择 golang
FROM golang:1.21-bookworm

# 设置工作目录
WORKDIR /app

# 将项目文件复制到容器中
COPY . .

# 设置go proxy
RUN go env -w GOPROXY=https://goproxy.cn,direct

# 构建 Go 项目
RUN go build -o main .

# 运行 Go 项目
CMD ["./main", "--generate-test-data"]