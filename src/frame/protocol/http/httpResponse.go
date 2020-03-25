package http

import "encoding/json"

type Response struct {
	Http
	statusCode int
	statusMsg  string
}

func (This *Response) GetStatusCode() int {
	return This.statusCode
}

func (This *Response) SetStatusCode(code int) {
	This.statusCode = code
}

func (This *Response) GetStatusMsg() string {
	return This.statusMsg
}

func (This *Response) SetStatusMsg(msg string) {
	This.statusMsg = msg
}

func (This *Response) SetStrResult(str string) {
	This.body = []byte(str)
}

func (This *Response) SetObjResult(v interface{}) error {
	bytes, err := json.Marshal(v)
	if err != nil {
		return err
	}
	This.body = bytes
	return nil
}

func NewResponse(r *Request) (resp *Response) {
	resp = new(Response)
	resp.conn = r.GetConn()
	resp.header = make(map[string]string)

	resp.statusCode = 200
	resp.statusMsg = "OK"

	resp.header["Server"] = "XSTech"
	resp.header["Content-Type"] = "application/json"

	connection, exist := r.header["Connection"]
	if exist && connection == "keep-alive" {
		resp.header["Connection"] = "keep-alive"
	}
	return resp
}
