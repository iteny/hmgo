package hmgo

import (
	hmcontext "github.com/iteny/hmgo/context"
	"sync"
)

// default filter execution points
const (
	BeforeStatic = iota
	BeforeRouter
	BeforeExec
	AfterExec
	FinishRouter
)

const (
	routerTypeBeego = iota
	routerTypeRESTFul
	routerTypeHandler
)

// ControllerRegister containers registered router rules, controller handlers and filters.
type ControllerRegister struct {
	routers      map[string]*Tree
	enableFilter bool
	filters      [FinishRouter + 1][]*FilterRouter
	pool         sync.Pool
}

func NewControllerRegister() *ControllerRegister {
	cr := &ControllerRegister{
		routers: make(map[string]*Tree),
	}
	cr.pool.New = func() interface{} {
		return hmcontext.NewContext()
	}
	return cr
}
