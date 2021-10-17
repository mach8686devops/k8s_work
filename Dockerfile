
# 打包依赖阶段使用golang作为基础镜像
FROM golang:1.16 as builder

# 启用go module
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /app


COPY . ./

# 指定运行时环境变量
ENV PORT=8080

EXPOSE 8080

RUN GOOS=linux GOARCH=amd64 go build work01.go

ENTRYPOINT ["./work01"]

