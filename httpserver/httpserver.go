package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	version := os.Getenv("VERSION")
	if version == "" {
		version = "0.0.1"
	}
	log.Printf("version: %s", version)

	engine := &Engine{
		version: version,
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		_ = http.ListenAndServe("0.0.0.0:8080", engine)
	}()

	<-sc
	fmt.Println("server is closing")
	err := engine.Shutdown(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println("server is closed")
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
			_, _ = fmt.Fprintf(w, "Header[%s]: %q\n", k, v)
		}
		_, _ = w.Write([]byte("ok"))
	default:
		_, _ = fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}

func (e *Engine) Shutdown(ctx context.Context) error {
	fmt.Println("server shutdown")
	return nil
}
