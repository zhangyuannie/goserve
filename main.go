package main

import (
	"crypto/subtle"
	"flag"
	"fmt"
	"net/http"
	"os"
)

var portFlag = flag.Int("p", 8000, "port number")
var dirFlag = flag.String("d", "", "directory to serve (default current directory)")
var usernameFlag = flag.String("username", "", "username for basic authentication (default none)")
var passwordFlag = flag.String("password", "", "password for basic authentication (default none)")

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

	fmt.Printf("Serving %v on port %v\n", *dirFlag, *portFlag)
	fmt.Printf("Available on http://localhost:%v\n", *portFlag)
	panic(http.ListenAndServe(fmt.Sprintf(":%d", *portFlag), h))
}
