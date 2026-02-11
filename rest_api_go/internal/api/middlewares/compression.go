package middlewares

import (
	"compress/gzip"
	"fmt"
	"net/http"
	"strings"
)

func Compression(next http.Handler)http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		//Check if the client accepts gzip encoding
		if !strings.Contains(r.Header.Get("Accept-Encoding"),"gzip"){
			next.ServeHTTP(w,r)
			return
		}
		//Set the content encoding header
		w.Header().Set("Content-Encoding","gzip")
		//Create a gzip writer
		gzipWriter := gzip.NewWriter(w)
		defer gzipWriter.Close()
		//Create a custom response writer
		w = &gzipResponseWriter{ResponseWriter:w,Writer:gzipWriter}
		
		//Call the next handler
		next.ServeHTTP(w,r)
		fmt.Println("Sent Response from Compression Middleware")
	})
}

//gzipResponseWriter wraps http.ResponseWriter to write gzipped responses
type gzipResponseWriter struct {
	http.ResponseWriter
	Writer *gzip.Writer
}

func(g *gzipResponseWriter)Write(b []byte)(int,error){
	return g.Writer.Write(b)
}