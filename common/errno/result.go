package errno

import "encoding/json"

type Result interface {
	i()
	WithData(data interface{}) Result
	WithTrace(id string) Result
	ToString() string
	error
}

type res struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data"`
	Trace string      `json:"trace,omitempty"`
}

func NewResult(code int, msg string) Result {
	return &res{
		Code: code,
		Msg:  msg,
	}
}

func (r *res) i() {
	panic("implement me")
}

func (r *res) WithData(data interface{}) Result {
	r.Data = data
	return r
}

func (r *res) WithTrace(id string) Result {
	r.Trace = id
	return r
}

func (r *res) ToString() string {
	_r := &struct {
		Code  int         `json:"code"`
		Msg   string      `json:"msg"`
		Data  interface{} `json:"data"`
		Trace string      `json:"trace,omitempty"`
	}{
		r.Code,
		r.Msg,
		r.Data,
		r.Trace,
	}

	raw, _ := json.Marshal(_r)
	return string(raw)
}

func (r *res) Error() string {
	return r.Msg
}
