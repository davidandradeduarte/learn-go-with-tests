package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	// GreetToStdout()
	log.Fatal(http.ListenAndServe(":5050", http.HandlerFunc(MyGreeterHandler)))
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "Http")
}

func GreetToStdout() {
	Greet(os.Stdout, "Stdout")
}

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}
