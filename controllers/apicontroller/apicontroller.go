package apicontroller

import (
	"encoding/json"
	"html/template"
	"log"
	"mock_api/entities"
	"mock_api/models/apimodel"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

type apiController struct {
	repo apimodel.RepositoryApi
}

func NewApiController(repo apimodel.RepositoryApi) apiController {
	return apiController{repo: repo}
}

func (c *apiController) Index(w http.ResponseWriter, r *http.Request) {
	apis, err := c.repo.GetAll()
	if err != nil {
		panic(err)
	}

	data := map[string]any{
		"apis": apis,
	}
	log.Println(data)

	temp, err := template.ParseFiles("views/api/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)

}

func (c *apiController) Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/api/create.html")
		if err != nil {
			panic(err)
		}

		temp.Execute(w, nil)
	}

	if r.Method == "POST" {
		var api entities.Api
		if r.FormValue("name") == "" || r.FormValue("url") == "" || r.FormValue("method") == "" || r.FormValue("response") == "" {
			log.Println("form value can't be empty")
			temp, _ := template.ParseFiles("views/api/create.html")
			temp.Execute(w, nil)
			return
		}

		pattern := `^/`
		matched, _ := regexp.MatchString(pattern, r.FormValue("url"))
		if matched {
			log.Println("URL should not start with a slash")
			temp, _ := template.ParseFiles("views/api/create.html")
			temp.Execute(w, nil)
			return
		}

		pattern = `\s`
		matched, _ = regexp.MatchString(pattern, r.FormValue("url"))
		if matched {
			log.Println("URL can't contain whitespace")
			temp, _ := template.ParseFiles("views/api/create.html")
			temp.Execute(w, nil)
			return
		}

		upperMethod := strings.ToUpper(r.FormValue("method"))
		method := upperMethod
		switch method {
		case "GET", "POST", "PUT", "PATCH", "DELETE":
			log.Println("Method is valid.")
		default:
			log.Println("Invalid method.")
			temp, _ := template.ParseFiles("views/api/create.html")
			temp.Execute(w, nil)
		}

		var parsedJson interface{}
		err := json.Unmarshal([]byte(r.FormValue("response")), &parsedJson)
		if err != nil {
			log.Println(err)
			temp, _ := template.ParseFiles("views/api/create.html")
			temp.Execute(w, nil)
			return
		}
		api.ApiMapper(r)

		ok := c.repo.Create(api)
		if !ok {
			temp, _ := template.ParseFiles("views/api/create.html")
			temp.Execute(w, nil)
		}
		http.Redirect(w, r, "/apis", http.StatusSeeOther)
	}

}

func (c *apiController) Edit(w http.ResponseWriter, r *http.Request) {
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

		api := c.repo.Detail(id)
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

		if r.FormValue("name") == "" || r.FormValue("url") == "" || r.FormValue("method") == "" || r.FormValue("response") == "" {
			log.Println("form value can't be empty")
			temp, _ := template.ParseFiles("views/api/create.html")
			temp.Execute(w, nil)
			return
		}

		pattern := `^/`
		matched, _ := regexp.MatchString(pattern, r.FormValue("url"))
		if matched {
			log.Println("URL should not start with a slash")
			temp, _ := template.ParseFiles("views/api/create.html")
			temp.Execute(w, nil)
			return
		}

		pattern = `\s`
		matched, _ = regexp.MatchString(pattern, r.FormValue("url"))
		if matched {
			log.Println("URL can't contain whitespace")
			temp, _ := template.ParseFiles("views/api/create.html")
			temp.Execute(w, nil)
			return
		}

		upperMethod := strings.ToUpper(r.FormValue("method"))
		method := upperMethod
		switch method {
		case "GET", "POST", "PUT", "PATCH", "DELETE":
			log.Println("Method is valid.")
		default:
			log.Println("Invalid method.")
			temp, _ := template.ParseFiles("views/api/create.html")
			temp.Execute(w, nil)
		}

		var parsedJson interface{}
		err = json.Unmarshal([]byte(r.FormValue("response")), &parsedJson)
		if err != nil {
			log.Println(err)
			temp, _ := template.ParseFiles("views/api/create.html")
			temp.Execute(w, nil)
			return
		}

		api.ApiMapper(r)
		api.Id = uint(id)

		ok := c.repo.Update(api)
		if !ok {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
		}
		http.Redirect(w, r, r.Header.Get("/apis"), http.StatusSeeOther)
	}

}

func (c *apiController) Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}
	var api entities.Api
	api.Id = uint(id)

	if err = c.repo.Delete(api); err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/apis", http.StatusSeeOther)
}
