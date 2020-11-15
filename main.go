package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

var flagPort = flag.Int("p", 8000, "the port number")
var flagDir = flag.String("d", "", "the directory to serve (default to the current directory)")

func main() {
	flag.Parse()
	if *flagDir == "" {
		wd, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		*flagDir = wd
	}

	fmt.Printf("Serving %v on port %v\n", *flagDir, *flagPort)
	fmt.Printf("Available on http://localhost:%v\n", *flagPort)

	panic(http.ListenAndServe(fmt.Sprintf(":%d", *flagPort), http.FileServer(http.Dir(*flagDir))))
}
