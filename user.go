package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
)

/* TODO, Add Token Server, make data server get the correct user
 *       Add Parameters for Public Data Server and Private Data Server
 *       How to Close Web Query from Other Crawler?
 */
func UserMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s := r.Header.Get("Authorization")
		s1 := strings.Replace(s, "Basic ", "", 1)
		fmt.Println(s, s1)
		ue, err := base64.StdEncoding.DecodeString(s1)
		if err == nil {
			//TODO
			fmt.Println("User Middleware :", string(ue))
		} else {
			//TODO
			fmt.Println(err)
		}
		next.ServeHTTP(w, r)
	})
}

type DataIndex struct {
	Genome string      `json:"genome"`
	Dbname string      `json:"dbname"`
	Data   interface{} `json:"data"` // map[string]string or map[string][]string? could be uri or more sofisticated data structure such as binindex image
	Format string      `json:"format"`
}

func cred(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", "Authorization")
		next.ServeHTTP(w, r)
	})
}
