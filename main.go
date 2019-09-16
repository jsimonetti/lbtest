package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
	"time"
)

func handlerNotFound(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		handlerNotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Host: %s\n", r.Host)
	fmt.Fprintf(w, "Date: %s\n", time.Now().UTC().Format(time.StampMilli))
}

func handlerDetails(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Host: %s\n", r.Host)
	hn, _ := os.Hostname()
	fmt.Fprintf(w, "Server: %s\n", hn)
	fmt.Fprintf(w, "RemoteAddr: %s\n", r.RemoteAddr)
	fmt.Fprintf(w, "Date: %s\n\n", time.Now().UTC().Format(time.StampMilli))
	fmt.Fprintf(w, "Proto: %s\n", r.Proto)
	fmt.Fprintf(w, "Method: %s\n", r.Method)
	fmt.Fprintf(w, "URL: %s\n\n", r.URL.Path)

	fmt.Fprintf(w, "Headers: \n")
	var keys []string
	for k, _ := range r.Header {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		v := r.Header[k]
		fmt.Fprintf(w, "%s: %s\n", k, v)
	}
	fmt.Fprintf(w, "\nCookies: \n")
	keys = []string{}
	for _, c := range r.Cookies() {
		keys = append(keys, c.Name)
	}
	sort.Strings(keys)
	for _, k := range keys {
		v, _ := r.Cookie(k)
		fmt.Fprintf(w, "%s: %s\n", k, v.Value)
	}
}

func handleMonitor(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK\n")
}
func main() {
	http.HandleFunc("/monitor/", handleMonitor)
	http.HandleFunc("/details/", handlerDetails)
	http.HandleFunc("/", handlerIndex)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
