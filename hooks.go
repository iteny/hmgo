package hmgo

import (
	"mime"
	"net/http"
)

//加载mime类型
func loadMine() error {
	for k, v := range mimemaps {
		mime.AddExtensionType(k, v)
	}
	return nil
}

//加载自定义http错误处理
func loadHttpErrorHandler() error {
	m := map[string]func(http.ResponseWriter, *http.Request){
		"401": unauthorized,
		"402": unauthorized,
		"403": unauthorized,
		"404": unauthorized,
		"405": unauthorized,
		"500": unauthorized,
		"501": unauthorized,
		"502": unauthorized,
		"503": unauthorized,
		"504": unauthorized,
	}
	for k, v := range m {
		if _, ok := ErrorMaps[k]; !ok {
			ErrorHandler(k, v)
		}
	}
	return nil
}
