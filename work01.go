package main

import (
	"log"
	"fmt"
	"net/http"
	"encoding/json"
	"os"
)

func main() {
	http.HandleFunc("/healthz", index) // index 为向 url发送请求时，调用的函数
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type Header struct {
	Connection string `json:"connection"`
	Accept     string `json:"accept"`
	Cookie     string `json:"cookie"`
	JAVA_HOME  string `json:"java_home"`
}

func index(writer http.ResponseWriter, request *http.Request) {
	//接收客户端 request，并将 request 中带的 header 写入 response header
	//读取当前系统的环境变量中的 VERSION 配置，并写入 response header
	//Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
	//当访问 localhost/healthz 时，应返回 200
	var mapOne map[string]string
	mapOne = make(map[string]string)

	mapOne["Connection"] = request.Header.Get("Connection")
	mapOne["Accept"] = request.Header.Get("Accept")
	mapOne["Cookie"] = request.Header.Get("Cookie")
	mapOne["JAVA_HOME"] = os.Getenv("JAVA_HOME")
	strJson, err := json.Marshal(mapOne)
	if err != nil {
		fmt.Printf("json.Marshal failed, err:%v\n", err)
		return
	}

	writer.Write(strJson)
	writer.WriteHeader(200)

}
