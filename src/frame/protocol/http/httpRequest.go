package http

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"frame/log"
	"net"
	"strconv"
	"strings"
)

type Request struct {
	Http
}

func (This *Request) analysis() {
	read := bufio.NewReader(This.conn)
	requestLine, err := read.ReadString('\n')
	if err != nil {
		errMsg := fmt.Sprint("frame.protocol.http.Request.analysis 读取请求行失败,错误信息:[", err, "]")
		log.Error(errMsg)
		panic(errMsg)
	}
	requestLine = strings.TrimSuffix(requestLine, "\r\n")
	This.analysisRequestLine(requestLine)
	//读取并解析header
	for {
		headerItem, err := read.ReadString('\n')
		if err != nil {
			errMsg := fmt.Sprint("frame.protocol.http.Request.analysis 读取header失败 [", err, "]")
			log.Error(errMsg)
			panic(errMsg)
		}
		headerItem = strings.TrimSuffix(headerItem, "\r\n")
		if headerItem == "" {
			break
		}
		This.analysisHeaderItem(headerItem)
	}
	//读取body
	contentLength, exist := This.header["Content-Length"]
	if exist {
		len, err := strconv.Atoi(contentLength)
		if err != nil {
			errMsg := fmt.Sprint("frame.protocol.http.Request.analysis 解析Content-Length失败,错误信息:[", err, "]")
			log.Error(errMsg)
			panic(errMsg)
		}
		if len > 0 {
			This.body = make([]byte, len, len)
			_, err := read.Read(This.body)
			if err != nil {
				errMsg := fmt.Sprint("frame.protocol.http.Request.analysis 读取body失败,错误信息:[", err, "]")
				log.Error(errMsg)
				panic(errMsg)
			}
		}
	}
	log.Info("frame.protocol.http.Request.analysis http协议解析成功\nmethod:", This.method, " url:[", This.url, "\nbody:", string(This.body))
}

func (This *Request) analysisHeaderItem(headerItem string) {
	strs := strings.SplitN(headerItem, ": ", 2)
	if len(strs) < 2 {
		errMsg := fmt.Sprint("frame.protocol.http.Request.analysisHeaderItem 解析header失败,headerItem:[", headerItem, "]")
		log.Error(errMsg)
		panic(errMsg)
	}
	This.header[strs[0]] = strs[1]
}

func (This *Request) analysisRequestLine(requestLine string) {
	strs := strings.Split(requestLine, " ")
	if len(strs) != 3 {
		errMsg := fmt.Sprint("frame.protocol.http.Request.analysisRequestLine 请求行解析失败,请求行:[", requestLine, "]")
		log.Error(errMsg)
		panic(errMsg)
	}
	This.method = strs[0]
	This.url = strs[1]
	This.protocolVersion = strs[2]
}

func (This *Request) GetObjParam(v interface{}) error {
	if This.body == nil {
		return errors.New("frame.protocol.http.Request.GetObjParam 获取参数失败,body为nil")
	}
	return json.Unmarshal(This.body, v)
}

func NewRequest(conn net.Conn) (req *Request) {
	req = new(Request)
	req.conn = conn
	req.header = make(map[string]string)
	return req
}
