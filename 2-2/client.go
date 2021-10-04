package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	version := os.Getenv("VERSION")
	log.Printf("version: %s", version)

	engine := &Engine{
		version: version,
	}
	log.Fatal(http.ListenAndServe(":80", engine))
}

type Engine struct {
	version string
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("VERSION", e.version)
	for k, v := range req.Header {
		w.Header().Add(k, fmt.Sprintf("%q", v))
	}

	// 目前没有要处理的逻辑，所以都返回 200
	w.WriteHeader(200)

	fmt.Printf("client ip: %s, http code: %d\n", req.Host, 200)

	switch req.URL.Path {
	case "/":
		fmt.Printf("URL.Path = %q\n", req.URL.Path)
	case "/healthz":
		for k, v := range w.Header() {
			fmt.Fprintf(w, "Header[%s]: %q\n", k, v)
		}
		w.Write([]byte("ok"))
	default:
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}
