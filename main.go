package main

import (
	"api_go/src/router"
	_ "github.com/lib/pq"
	"net/http"
	"os"
)

func main() {
	err := http.ListenAndServe(":8080", router.Router())
	if err != nil {
		print("Error: ", err.Error())
		os.Exit(1)
	}
}
