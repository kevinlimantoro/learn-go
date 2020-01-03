package util

import (
	"compress/gzip"
	"encoding/json"
	"net/http"
)

// Message is a base model in which the server will return
func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

// RespondUncompress is the function that wrap how server will write the response to client
func RespondUncompress(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// Respond is the function that wrap how server will write the response to client
func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Accept-Charset", "utf-8")
	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Content-Encoding", "gzip")
	gz := gzip.NewWriter(w)
	defer gz.Close()
	json.NewEncoder(gz).Encode(data)
}

// type gzipResponseWriter struct {
// 	io.Writer
// 	http.ResponseWriter
// }

// func (w gzipResponseWriter) Write(b []byte) (int, error) {
// 	return w.Writer.Write(b)
// }

// func MakeGzipHandler(fn http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
// 			fn(w, r)
// 			return
// 		}
// 		w.Header().Set("Content-Encoding", "gzip")
// 		gz := gzip.NewWriter(w)
// 		defer gz.Close()
// 		gzr := gzipResponseWriter{Writer: gz, ResponseWriter: w}
// 		fn(gzr, r)
// 	}
// }
