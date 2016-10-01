package hmgo

import (
	"fmt"
	"net/http"
	"template"
)

// show 401 unauthorized error.
func unauthorized(w http.ResponseWriter, r *http.Request) {
	responseError(w, r,
		401,
		"<br>The page you have requested can't be authorized."+
			"<br>Perhaps you are here because:"+
			"<br><br><ul>"+
			"<br>The credentials you supplied are incorrect"+
			"<br>There are errors in the website address"+
			"</ul>",
	)
}
func responseError(w http.ResponseWriter, r *http.Request, errCode int, errContent string) {
	t, _ := template.New("hmErrorTemp").Parse(errorTpl)
	data := map[string]interface{}{
		"Title":        http.StatusText(errCode),
		"BeegoVersion": VERSION,
		"Content":      template.HTML(errContent),
	}
	t.Execute(w, data)
}
