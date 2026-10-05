package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, req *http.Request) {
		bytes, err := os.ReadFile(path)
		if err != nil {
			panic(err)
		}
		writer.Header().Set("Content-Type", "text/html; charset=utf-8")
		// TODO figure out how to inject a stylesheet
		fmt.Printf("Received a request from %s\n", req.RemoteAddr)
		_, err = writer.Write(bytes)
		if err != nil {
			panic(err)
		}
	})
	fmt.Println("Listening on 127.0.0.1:8080")
	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		panic(err)
	}

}
