package responder

import (
	"fmt"
	"web-api-scaffold/internal/pkg/errno"
)

type Body struct {
	Code errno.Code  `json:"code"`
	Data interface{} `json:"data,omitempty"`
	Msg  string      `json:"msg,omitempty"`
	Err  string      `json:"error,omitempty"`
}

func Succeed() *Body {
	return &Body{
		Code: errno.CodeOk,
		Msg:  errno.CodeOk.Message(),
	}
}

func BuildBody(code errno.Code) *Body {
	return &Body{
		Code: code,
		Msg:  code.Message(),
	}
}

func (r *Body) With(val interface{}) *Body {
	switch v := val.(type) {
	case error:
		return r.setError(v)
	case string:
		return r.setMessage(v)
	default:
		return r.setData(v)
	}
}

func (r *Body) setMessage(msg string) *Body {
	r.Msg = msg
	return r
}

func (r *Body) setData(data interface{}) *Body {
	r.Data = data
	return r
}

func (r *Body) setError(err error) *Body {
	r.Err = err.Error()
	return r
}

func (r *Body) WarpMessage(msg string) *Body {
	if r.Msg != "" {
		r.Msg = fmt.Sprintf("%s, %s", r.Msg, msg)
	}
	return r
}

func (r *Body) WarpError(err string) *Body {
	if r.Err != "" {
		r.Err = fmt.Sprintf("%s, %s", r.Err, err)
	}
	return r
}

func (r *Body) Error() string {
	return r.Msg
}
