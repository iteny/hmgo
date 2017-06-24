package hmgo

import "net/http"

// Handle is a function that can be registered to a route to Handle Http
// Handle是一个函数，那就可以注册一个路由来处理Http
// Request the same http.HandlerFunc,but has a third parameter for the values of Wildcard (variables).
// 请求和http.HandlerFunc相同，但第三个参数为通配符(变量)的值。
type Handle func(http.ResponseWriter, *http.Request, Params)

// Param is a single url parameter,consisting of a key and a value.
// Param是单一url参数，由一个键和值组成。
type Param struct {
	Key   string
	Value string
}

// Params is a Param slice,as returned by the router.
// Params类型是一个Param结构体切片,由路由返回。
// The slice is ordered,the first Url Parameter is also the first slice value.
// 切片被排序，第一个Url参数也是第一个切片的值。
// It's therefore safe to read values by the index.
// 所以从索引读取值是安全的。
type Params []Param

// ByName returns the value of the first Param which key matched the given name.
// If no matching Param is found,an empty string is returned.
func (ps Params) ByName(name string) string {
	for i := range ps {
		if ps[i].Key == name {
			return ps[i].Value
		}
	}
	return ""
}
