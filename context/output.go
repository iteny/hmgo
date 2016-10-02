package context

// BeegoOutput does work for sending response header.
type BeegoOutput struct {
	Context    *Context
	Status     int
	EnableGzip bool
}

// NewOutput returns new BeegoOutput.
// it contains nothing now.
func NewOutput() *BeegoOutput {
	return &BeegoOutput{}
}
