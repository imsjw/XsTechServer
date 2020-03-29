package http

import (
	"bufio"
	"fmt"
	"frame/interfaces"
	"frame/log"
	"net"
	"time"
)

func Analysis(conn net.Conn) (req *Request, resp *Response) {
	req = NewRequest(conn)
	req.analysis()
	resp = NewResponse(req)
	return
}

func OutResponse(req *Request, resp *Response, i *interfaces.Interface) {
	w := bufio.NewWriter(resp.GetConn())
	//输出响应行
	w.WriteString(fmt.Sprint("HTTP/1.1 ", resp.GetStatusCode(), " ", resp.GetStatusMsg(), "\r\n"))
	//设置响应数据发送时间
	resp.header["Date"] = time.Now().UTC().Format("Sun, 02 Jan 2006 15:04:05 GMT")
	//设置body长度
	resp.header["Content-Length"] = fmt.Sprint(len(resp.body))
	//输出header
	for k, v := range resp.header {
		w.WriteString(k)
		w.WriteString(": ")
		w.WriteString(v)
		w.WriteString("\r\n")
	}
	w.WriteString("\r\n")
	w.Write(resp.GetBody())
	w.Flush()

	log.Info("frame.protocol.http.OutResponse\nmethod:", req.method, " url:", req.url, "\nreqBody: ", string(req.body), "\nrespBody:", string(resp.body))

}
