package main

import (
	"log"
	apicontroller "mock_api/controllers/apicontroller"
	homecontroller "mock_api/controllers/homecontroller"
	"mock_api/controllers/mockcontroller"
	"mock_api/models/apimodel"
	"mock_api/pkg"
	"net/http"
	"os"
)

var (
	port = ":" + os.Getenv("PORT")
)

func main() {
	db := pkg.DBConnection()

	repoApi := apimodel.NewRepositoryAPI(db)
	controllerApi := apicontroller.NewApiController(repoApi)

	http.HandleFunc("/", homecontroller.Welcome)

	http.HandleFunc("/api/", mockcontroller.GetMock)

	http.HandleFunc("/apis/", controllerApi.Index)
	http.HandleFunc("/apis/add", controllerApi.Add)
	http.HandleFunc("/apis/edit", controllerApi.Edit)
	http.HandleFunc("/apis/delete", controllerApi.Delete)

	log.Printf("service serve at port %s", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		panic(err)
	}
}
