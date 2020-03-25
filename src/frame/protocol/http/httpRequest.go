package http

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"strings"
)

type Request struct {
	Http
}

func (This *Request) analysis() error {
	read := bufio.NewReader(This.conn)
	requestLine, err := read.ReadString('\n')
	if err != nil {
		if err == io.EOF {
			return err
		}
		return errors.New(fmt.Sprint("[http.analysis() error] [read requestLine error] [bufio.ReadString() error] [", err, "]"))
	}
	requestLine = strings.TrimSuffix(requestLine, "\r\n")
	err = This.analysisRequestLine(requestLine)
	if err != nil {
		return errors.New(fmt.Sprint("[http.analysis() error] ", err))
	}
	//读取并解析header
	for {
		headerItem, err := read.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				return err
			}
			return errors.New(fmt.Sprint("[http.analysis() error] [read header item error] [bufio.ReadString() error] [", err, "]"))
		}
		headerItem = strings.TrimSuffix(headerItem, "\r\n")
		if headerItem == "" {
			break
		}
		err = This.analysisHeaderItem(headerItem)
		if err != nil {
			return err
		}
	}
	//读取body
	contentLength, exist := This.header["Content-Length"]
	if exist {
		len, err := strconv.Atoi(contentLength)
		if err != nil {
			return errors.New(fmt.Sprint("[http.analysis() error] [analysis Content-Length error] [", err, "]"))
		}
		if len > 0 {
			This.body = make([]byte, len, len)
			_, err := read.Read(This.body)
			if err == io.EOF {
				return err
			}
		}
	}
	log.Println("req url:", This.url, " body:", string(This.body))
	return nil
}

func (This *Request) analysisHeaderItem(headerItem string) error {
	strs := strings.SplitN(headerItem, ": ", 2)
	if len(strs) < 2 {
		return errors.New(fmt.Sprint("[http.analysisHeaderItem() error] [header item content: ", headerItem, "]"))
	}
	This.header[strs[0]] = strs[1]
	return nil
}

func (This *Request) analysisRequestLine(requestLine string) error {
	strs := strings.Split(requestLine, " ")
	if len(strs) != 3 {
		return errors.New(fmt.Sprint("[http.analysisRequestLine() error] [requestLine: ", requestLine, "]"))
	}
	This.method = strs[0]
	This.url = strs[1]
	This.protocolVersion = strs[2]
	return nil
}

func (This *Request) GetObjParam(v interface{}) error {
	if This.body == nil {
		return errors.New("[http.GetObjParam() error] [body is nil]")
	}
	return json.Unmarshal(This.body, v)
}

func NewRequest(conn net.Conn) (req *Request) {
	req = new(Request)
	req.conn = conn
	req.header = make(map[string]string)
	return req
}
