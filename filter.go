package hmgo

import (
	hmcontext "github.com/iteny/hmgo/context"
)

// FilterFunc defines a filter function which is invoked before the controller handler is executed.
type FilterFunc func(*hmcontext.Context)

// FilterRouter defines a filter operation which is invoked before the controller handler is executed.
// It can match the URL against a pattern, and execute a filter function
// when a request with a matching URL arrives.
type FilterRouter struct {
	filterFunc     FilterFunc
	tree           *Tree
	pattern        string
	returnOnOutput bool
	resetParams    bool
}
