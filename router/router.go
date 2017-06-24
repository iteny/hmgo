package router

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
// ByName函数返回匹配给定name的第一个参数的值。
// If no matching Param is found,returns an empty string.
// 如果找不到匹配的参数，返回一个空的字符串。
func (ps Params) ByName(name string) string {
	for i := range ps {
		if ps[i].Key == name {
			return ps[i].Value
		}
	}
	return ""
}

// Router is a http.Handler which can be used to dispatch requests to different Handler functions via configurable.
// Router结构体是一个http.Handler,可以通过配置向不同的Handler函数发送请求。
type Router struct {
	trees map[string]*node
	// Enables automatic redirection if the current route can't be matched but a
	// handler for the path with (without) the trailing slash exists.
	// For example if /foo/ is requested but a route only exists for /foo, the
	// client is redirected to /foo with http status code 301 for GET requests
	// and 307 for all other request methods.
	// 启用自动重定向，如果当前路由不能匹配，但是一个handler路径有(没有)尾部斜杠的存在。
	// 例如请求/foo/，但是路径只有/foo,此客户端就重定向到/foo,Get请求的http状态码301和所有其他的请求方法。
	RedirectTrailingSlash bool
	// If enabled, the router tries to fix the current request path, if no
	// handle is registered for it.
	// First superfluous path elements like ../ or // are removed.
	// Afterwards the router does a case-insensitive lookup of the cleaned path.
	// If a handle can be found for this route, the router makes a redirection
	// to the corrected path with status code 301 for GET requests and 307 for
	// all other request methods.
	// For example /FOO and /..//Foo could be redirected to /foo.
	// RedirectTrailingSlash is independent of this option.
	// 如果启用，此路由尝试解决当前请求的路径，如果Handle没有注册。
	// 第一个多余的路径元素比如 ../ 或者 // 删除。
	// 之后路由对清除的路径进行不区分大小写的查询。
	// 如果可以为此路由找到handle,路由器将重定向到具有状态码301的更正路径，用于GET请求，307位所有其他请求的方法。
	// 例如 /FOO 和 /..//Foo 能够重定向到 /foo。
	// RedirectTrailingSlash于此选项无关。
	RedirectFixedPath bool
	// If enabled, the router checks if another method is allowed for the
	// current route, if the current request can not be routed.
	// If this is the case, the request is answered with 'Method Not Allowed'
	// and HTTP status code 405.
	// If no other Method is allowed, the request is delegated to the NotFound
	// handler.
	// 如果启用，路由器将检查其他方法是否合法当前路径，如果当前请求不能路由。
	// 如果是这种情况，请求以'Method Not Allowed'和'HTTP status code 405'应答。
	// 如果没有其他方法是合法的，此请求被委派给NotFound Handler。
	HandleMethodNotAllowed bool
	// If enabled, the router automatically replies to OPTIONS requests.
	// Custom OPTIONS handlers take priority over automatic replies.
	// 如果启用，此路由将自动回复选项请求。
	// 自定义选项hanlders程序优先于自动回复。
	HandleOPTIONS bool
	// Configurable http.Handler which is called when no matching route is
	// found. If it is not set, http.NotFound is used.
	// 无匹配路由时调用可配置http.Handler查询,如果未设置，则使用http.NotFound
	NotFound http.Handler
	// Configurable http.Handler which is called when a request
	// cannot be routed and HandleMethodNotAllowed is true.
	// If it is not set, http.Error with http.StatusMethodNotAllowed is used.
	// The "Allow" header with allowed request methods is set before the handler
	// is called.
	// 可配置http.Handler，在请求时调用不能路由，HandleMethodNotAllowed为true。如果未设置，则使用http.StatusMethodNotAllowed的http.Error。允许的请求方法的"允许"标头被调用之前设置。
	MethodNotAllowed http.Handler
	// Function to handle panics recovered from http handlers.
	// It should be used to generate a error page and return the http error code
	// 500 (Internal Server Error).
	// The handler can be used to keep your server from crashing because of
	// unrecovered panics.
	// 从http Handlers处理程序恢复异常功能。
	// 它应该用于生成错误页面并返回http错误码500(内部服务器错误)。
	// 改Handler程序可用于保持你的服务器不会因为未收回的异常而崩溃。
	PanicHandler func(http.ResponseWriter, *http.Request, interface{})
}

// New returns a new initialized Router.
// New函数返回一个新的初始化Router结构体
// Path auto-correction,including trailing slashes,is enabled by default.
// 默认情况下启用路径自动更正，包括尾部斜杠。
func New() *Router {
	return &Router{
		RedirectTrailingSlash:  true,
		RedirectFixedPath:      true,
		HandleMethodNotAllowed: true,
		HandleOPTIONS:          true,
	}
}

// GET is a shortcut for router.Handle("GET",path,handle)
func (r *Router) GET(path string, handle Handle) {
	r.Handle("GET", path, handle)
}

// HEAD is a shortcut for router.Handle("HEAD", path, handle)
func (r *Router) HEAD(path string, handle Handle) {
	r.Handle("HEAD", path, handle)
}

// OPTIONS is a shortcut for router.Handle("OPTIONS", path, handle)
func (r *Router) OPTIONS(path string, handle Handle) {
	r.Handle("OPTIONS", path, handle)
}

// POST is a shortcut for router.Handle("POST", path, handle)
func (r *Router) POST(path string, handle Handle) {
	r.Handle("POST", path, handle)
}

// PUT is a shortcut for router.Handle("PUT", path, handle)
func (r *Router) PUT(path string, handle Handle) {
	r.Handle("PUT", path, handle)
}

// PATCH is a shortcut for router.Handle("PATCH", path, handle)
func (r *Router) PATCH(path string, handle Handle) {
	r.Handle("PATCH", path, handle)
}

// DELETE is a shortcut for router.Handle("DELETE", path, handle)
func (r *Router) DELETE(path string, handle Handle) {
	r.Handle("DELETE", path, handle)
}

// Handle registers a new request handle with the given path and method.
//
// For GET, POST, PUT, PATCH and DELETE requests the respective shortcut
// functions can be used.
//
// This function is intended for bulk loading and to allow the usage of less
// frequently used, non-standardized or custom methods (e.g. for internal
// communication with a proxy).
func (r *Router) Handle(method, path string, handle Handle) {
	if path[0] != '/' {
		panic("path must begin with '/' in path '" + path + "'")
	}

	if r.trees == nil {
		r.trees = make(map[string]*node)
	}

	root := r.trees[method]
	if root == nil {
		root = new(node)
		r.trees[method] = root
	}

	root.addRoute(path, handle)
}

// Handler is an adapter which allows the usage of an http.Handler as a
// request handle.
func (r *Router) Handler(method, path string, handler http.Handler) {
	r.Handle(method, path,
		func(w http.ResponseWriter, req *http.Request, _ Params) {
			handler.ServeHTTP(w, req)
		},
	)
}

// HandlerFunc is an adapter which allows the usage of an http.HandlerFunc as a
// request handle.
func (r *Router) HandlerFunc(method, path string, handler http.HandlerFunc) {
	r.Handler(method, path, handler)
}

// ServeFiles serves files from the given file system root.
// The path must end with "/*filepath", files are then served from the local
// path /defined/root/dir/*filepath.
// For example if root is "/etc" and *filepath is "passwd", the local file
// "/etc/passwd" would be served.
// Internally a http.FileServer is used, therefore http.NotFound is used instead
// of the Router's NotFound handler.
// To use the operating system's file system implementation,
// use http.Dir:
//     router.ServeFiles("/src/*filepath", http.Dir("/var/www"))
func (r *Router) ServeFiles(path string, root http.FileSystem) {
	if len(path) < 10 || path[len(path)-10:] != "/*filepath" {
		panic("path must end with /*filepath in path '" + path + "'")
	}

	fileServer := http.FileServer(root)

	r.GET(path, func(w http.ResponseWriter, req *http.Request, ps Params) {
		req.URL.Path = ps.ByName("filepath")
		fileServer.ServeHTTP(w, req)
	})
}

func (r *Router) recv(w http.ResponseWriter, req *http.Request) {
	if rcv := recover(); rcv != nil {
		r.PanicHandler(w, req, rcv)
	}
}

// Lookup allows the manual lookup of a method + path combo.
// This is e.g. useful to build a framework around this router.
// If the path was found, it returns the handle function and the path parameter
// values. Otherwise the third return value indicates whether a redirection to
// the same path with an extra / without the trailing slash should be performed.
func (r *Router) Lookup(method, path string) (Handle, Params, bool) {
	if root := r.trees[method]; root != nil {
		return root.getValue(path)
	}
	return nil, nil, false
}

func (r *Router) allowed(path, reqMethod string) (allow string) {
	if path == "*" { // server-wide
		for method := range r.trees {
			if method == "OPTIONS" {
				continue
			}

			// add request method to list of allowed methods
			if len(allow) == 0 {
				allow = method
			} else {
				allow += ", " + method
			}
		}
	} else { // specific path
		for method := range r.trees {
			// Skip the requested method - we already tried this one
			if method == reqMethod || method == "OPTIONS" {
				continue
			}

			handle, _, _ := r.trees[method].getValue(path)
			if handle != nil {
				// add request method to list of allowed methods
				if len(allow) == 0 {
					allow = method
				} else {
					allow += ", " + method
				}
			}
		}
	}
	if len(allow) > 0 {
		allow += ", OPTIONS"
	}
	return
}

// ServeHTTP makes the router implement the http.Handler interface.
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if r.PanicHandler != nil {
		defer r.recv(w, req)
	}

	path := req.URL.Path

	if root := r.trees[req.Method]; root != nil {
		if handle, ps, tsr := root.getValue(path); handle != nil {
			handle(w, req, ps)
			return
		} else if req.Method != "CONNECT" && path != "/" {
			code := 301 // Permanent redirect, request with GET method
			if req.Method != "GET" {
				// Temporary redirect, request with same method
				// As of Go 1.3, Go does not support status code 308.
				code = 307
			}

			if tsr && r.RedirectTrailingSlash {
				if len(path) > 1 && path[len(path)-1] == '/' {
					req.URL.Path = path[:len(path)-1]
				} else {
					req.URL.Path = path + "/"
				}
				http.Redirect(w, req, req.URL.String(), code)
				return
			}

			// Try to fix the request path
			if r.RedirectFixedPath {
				fixedPath, found := root.findCaseInsensitivePath(
					CleanPath(path),
					r.RedirectTrailingSlash,
				)
				if found {
					req.URL.Path = string(fixedPath)
					http.Redirect(w, req, req.URL.String(), code)
					return
				}
			}
		}
	}

	if req.Method == "OPTIONS" {
		// Handle OPTIONS requests
		if r.HandleOPTIONS {
			if allow := r.allowed(path, req.Method); len(allow) > 0 {
				w.Header().Set("Allow", allow)
				return
			}
		}
	} else {
		// Handle 405
		if r.HandleMethodNotAllowed {
			if allow := r.allowed(path, req.Method); len(allow) > 0 {
				w.Header().Set("Allow", allow)
				if r.MethodNotAllowed != nil {
					r.MethodNotAllowed.ServeHTTP(w, req)
				} else {
					http.Error(w,
						http.StatusText(http.StatusMethodNotAllowed),
						http.StatusMethodNotAllowed,
					)
				}
				return
			}
		}
	}

	// Handle 404
	if r.NotFound != nil {
		r.NotFound.ServeHTTP(w, req)
	} else {
		http.NotFound(w, req)
	}
}
