
# 打包依赖阶段使用golang作为基础镜像
FROM golang:1.16 as builder

# 启用go module
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /app

COPY . .

# 指定OS等，并go build
RUN GOOS=linux GOARCH=amd64 go build .


# 运行阶段指定scratch作为基础镜像
FROM alpine

WORKDIR /app


COPY --from=builder . ./

# 指定运行时环境变量
ENV PORT=80

EXPOSE 80

RUN go build work01.go

RUN ./work01

