package main

import (
	"fmt"
	"mock_api/pkg"
	"net/http"
)

func main() {
	pkg.DBConnection()

	fmt.Println("service serve at port 8181")
	err := http.ListenAndServe(":8181", nil)
	if err != nil {
		panic(err)
	}
}
