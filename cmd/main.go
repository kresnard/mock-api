package main

import (
	"log"
	apicontroller "mock_api/controllers/apicontroller"
	homecontroller "mock_api/controllers/homecontroller"
	"mock_api/controllers/mockcontroller"
	"mock_api/pkg"
	"net/http"
	"os"
)

var (
	port = ":" + os.Getenv("PORT")
)

func main() {
	pkg.DBConnection()

	http.HandleFunc("/", homecontroller.Welcome)

	http.HandleFunc("/api/", mockcontroller.GetMock)

	http.HandleFunc("/apis/", apicontroller.Index)
	http.HandleFunc("/apis/add", apicontroller.Add)
	http.HandleFunc("/apis/edit", apicontroller.Edit)
	http.HandleFunc("/apis/delete", apicontroller.Delete)

	log.Printf("service serve at port %s", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		panic(err)
	}
}
