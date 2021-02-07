package main

import (
	"crypto/subtle"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"time"
)

var portFlag = flag.Int("p", 8000, "port number")
var dirFlag = flag.String("d", "", "alternate directory to serve")
var usernameFlag = flag.String("username", "", "username for basic authentication")
var passwordFlag = flag.String("password", "", "password for basic authentication")
var versionFlag = flag.Bool("version", false, "print goserve version")

type responseRecord struct {
	http.ResponseWriter
	status int
}

func (r *responseRecord) WriteHeader(status int) {
	r.status = status
	r.ResponseWriter.WriteHeader(status)
}

func defaultHypen(s string) string {
	if s == "" {
		return "-"
	}
	return s
}

func handleLog(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		record := &responseRecord{
			ResponseWriter: w,
			status:         200,
		}
		h.ServeHTTP(record, r)

		host, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil || host == "" {
			host = "-"
		}

		fmt.Printf("%s - - %s \"%s %s %s\" %d - \"%s\" \"%s\"\n",
			host,
			time.Now().Format("[02/Jan/2006:15:04:05 -0700]"),
			r.Method, r.URL, r.Proto,
			record.status,
			defaultHypen(r.Referer()),
			defaultHypen(r.UserAgent()),
		)
	}
}

func auth(h http.Handler, username, password string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()

		if !ok || subtle.ConstantTimeCompare([]byte(user), []byte(username)) != 1 ||
			subtle.ConstantTimeCompare([]byte(pass), []byte(password)) != 1 {

			w.Header().Set("WWW-Authenticate", `Basic realm=""`)
			http.Error(w, "Authorization Required", http.StatusUnauthorized)
			return
		}

		h.ServeHTTP(w, r)
	}
}

func main() {
	flag.Parse()

	if *versionFlag {
		fmt.Println("goserve version 0.1.0")
		return
	}

	if *dirFlag == "" {
		wd, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		*dirFlag = wd
	}

	h := http.FileServer(http.Dir(*dirFlag))
	if *usernameFlag != "" && *passwordFlag != "" {
		h = auth(h, *usernameFlag, *passwordFlag)
	}
	h = handleLog(h)

	fmt.Printf("Serving %v on port %v\n", *dirFlag, *portFlag)
	fmt.Printf("Available on http://localhost:%v\n", *portFlag)
	panic(http.ListenAndServe(fmt.Sprintf(":%d", *portFlag), h))
}
