package context

import (
	"net/http"
)

// NewContext return the Context with Input and Output
func NewContext() *Context {
	return &Context{
		Input:  NewInput(),
		Output: NewOutput(),
	}
}

// Context Http request context struct including BeegoInput, BeegoOutput, http.Request and http.ResponseWriter.
// BeegoInput and BeegoOutput provides some api to operate request and response more easily.
type Context struct {
	Input          *BeegoInput
	Output         *BeegoOutput
	Request        *http.Request
	ResponseWriter *Response
	_xsrfToken     string
}

//Response is a wrapper for the http.ResponseWriter
//started set to true if response was written to then don't execute other handler
type Response struct {
	http.ResponseWriter
	Started bool
	Status  int
}
