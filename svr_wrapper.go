package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
)

func newSvrWrapper(svr http.Handler) http.Handler {
	return &svrWrapper{
		svr: svr,
	}
}

type svrWrapper struct {
	svr http.Handler
}

type respWriterWrapper struct {
	buf bytes.Buffer
	w   http.ResponseWriter
}

func (r *respWriterWrapper) Header() http.Header {
	return r.w.Header()
}

func (r *respWriterWrapper) Write(b []byte) (_ int, _ error) {
	r.buf.Write(b)
	return r.w.Write(b)
}

func (r *respWriterWrapper) WriteHeader(statusCode int) {
	r.w.WriteHeader(statusCode)
}

func (s *svrWrapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rawxml, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	defer r.Body.Close()
	r.Body = io.NopCloser(bytes.NewBuffer(rawxml))

	fmt.Println(string(rawxml))

	ww := &respWriterWrapper{
		w: w,
	}
	s.svr.ServeHTTP(ww, r)

	log.Println(ww.buf.String())
}
