package sofu

import "strings"

type Router struct {
	routes map[string]map[string]HandlerFunc
}

func NewRouter() *Router {
	return &Router{routes: make(map[string]map[string]HandlerFunc)}
}

func (r *Router) GET(path string, handler HandlerFunc) {
	r.addRoute("GET", path, handler)
}

func (r *Router) POST(path string, handler HandlerFunc) {
	r.addRoute("POST", path, handler)
}

func (r *Router) addRoute(method, path string, handler HandlerFunc) {
	if r.routes[method] == nil {
		r.routes[method] = make(map[string]HandlerFunc)
	}
	r.routes[method][path] = handler
}

func (r *Router) Handle(c *Context) {
	method := c.Request.Method
	path := c.Request.Path
	params := make(map[string]string)

	for routePath, handler := range r.routes[method] {
		if matchRoute(routePath, path, params) {
			c.Request.Params = params
			handler(c)
			return
		}
	}
	c.String(404, "Not Found")
}

func matchRoute(routePath, path string, params map[string]string) bool {
	routeParts := strings.Split(routePath, "/")
	pathParts := strings.Split(path, "/")
	if len(routeParts) != len(pathParts) {
		return false
	}
	for i, part := range routeParts {
		if strings.HasPrefix(part, ":") {
			params[part[1:]] = pathParts[i]
		} else if part != pathParts[i] {
			return false
		}
	}
	return true
}
