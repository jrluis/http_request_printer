package main 

import (
	"fmt"
	"net/http"
	"io"
	"os"
	"strings"
	"time"
)

var portNumber = ""

func main() {
	CheckEnvironmentVariable("PORT")
	
	portNumber = os.Getenv("PORT")
	
	http.HandleFunc("/", handler)
    http.ListenAndServe(":" + portNumber, nil)
}

func CheckEnvironmentVariable(varName string) {
	varValue := os.Getenv(varName)
	
	if len(strings.TrimSpace(varValue)) == 0 {
		fmt.Printf("%s environment variable is not setted\n", varName)
		os.Exit(-1)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	fmt.Printf("[%v] ", t)
	
	io.Copy(os.Stdout, r.Body)
	fmt.Printf("\n")
}

