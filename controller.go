package hmgo

import (
	"net/http"
)

type Controller struct {
	Wr http.ResponseWriter
	R  *http.Request
}
