package apicontroller

import (
	"html/template"
	"mock_api/entities"
	"mock_api/models/apimodel"
	"net/http"
	"strconv"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	apis, err := apimodel.GetAll()
	if err != nil {
		panic(err)
	}

	data := map[string]any{
		"apis": apis,
	}

	temp, err := template.ParseFiles("views/api/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)

}

func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/api/create.html")
		if err != nil {
			panic(err)
		}

		temp.Execute(w, nil)
	}

	if r.Method == "POST" {
		var api entities.Api
		api.Name = r.FormValue("name")
		api.Url = r.FormValue("url")
		api.Method = "GET"
		api.Response = r.FormValue("response")
		api.CreatedAt = time.Now()
		api.UpdatedAt = time.Now()

		ok := apimodel.Create(api)
		if !ok {
			temp, _ := template.ParseFiles("views/api/create.html")
			temp.Execute(w, nil)
		}
		http.Redirect(w, r, "/apis", http.StatusSeeOther)
	}

}

func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/api/edit.html")
		if err != nil {
			panic(err)
		}

		idString := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		api := apimodel.Detail(id)
		data := map[string]any{
			"api": api,
		}

		temp.Execute(w, data)
	}

	if r.Method == "POST" {
		var api entities.Api

		idString := r.FormValue("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		api.Id = uint(id)
		api.Name = r.FormValue("name")
		api.Url = r.FormValue("url")
		api.Method = "GET"
		api.Response = r.FormValue("response")
		api.UpdatedAt = time.Now()

		ok := apimodel.Update(api)
		if !ok {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
		}
		http.Redirect(w, r, r.Header.Get("/apis"), http.StatusSeeOther)
	}

}

func Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}
	var api entities.Api
	api.Id = uint(id)

	if err = apimodel.Delete(api); err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/apis", http.StatusSeeOther)
}
