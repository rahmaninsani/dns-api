package app

import (
	"github.com/rahmaninsani/dns-api/helper"
	"net/http"
	"strings"
)

type HttpParams map[string]string
type HandlerFunc func(writer http.ResponseWriter, response *http.Request, params HttpParams)
type PanicHandlerFunc func(http.ResponseWriter, *http.Request, any)

type Router struct {
	routes       map[string]map[string]HandlerFunc
	PanicHandler PanicHandlerFunc
}

func NewRouter() *Router {
	return &Router{
		routes: make(map[string]map[string]HandlerFunc),
	}
}

func (router *Router) GET(path string, handlerFunc HandlerFunc) {
	router.addRoute(http.MethodGet, path, handlerFunc)
}

func (router *Router) POST(path string, handlerFunc HandlerFunc) {
	router.addRoute(http.MethodPost, path, handlerFunc)
}

func (router *Router) PUT(path string, handlerFunc HandlerFunc) {
	router.addRoute(http.MethodPut, path, handlerFunc)
}

func (router *Router) PATCH(path string, handlerFunc HandlerFunc) {
	router.addRoute(http.MethodPatch, path, handlerFunc)
}

func (router *Router) DELETE(path string, handlerFunc HandlerFunc) {
	router.addRoute(http.MethodDelete, path, handlerFunc)
}

func (router *Router) addRoute(method, path string, handlerFunc HandlerFunc) {
	if router.routes[method] == nil {
		router.routes[method] = make(map[string]HandlerFunc)
	}
	router.routes[method][path] = handlerFunc
}

func (router *Router) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	var handler HandlerFunc
	var params map[string]string

	for path, handlerFunc := range router.routes[request.Method] {
		if match, p := matchPath(path, request.URL.Path); match {
			handler = handlerFunc
			params = p
			break
		}
	}

	if handler == nil {
		response := helper.ToResponse(http.StatusNotFound, "Endpoint Not Found", nil)
		helper.WriteToResponseBody(writer, response)
		return
	}

	// Handle panic
	defer func() {
		if err := recover(); err != nil {
			router.PanicHandler(writer, request, err)
		}
	}()

	handler(writer, request, params)
}

func matchPath(pattern, path string) (bool, map[string]string) {
	patternParts := strings.Split(pattern, "/")
	pathParts := strings.Split(path, "/")

	if len(patternParts) != len(pathParts) {
		return false, nil
	}

	var params HttpParams = make(map[string]string)

	for i, part := range patternParts {
		if strings.HasPrefix(part, ":") {
			params[strings.TrimPrefix(part, ":")] = pathParts[i]
		} else if part != pathParts[i] {
			return false, nil
		}
	}

	return true, params
}
