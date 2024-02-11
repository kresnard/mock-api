package entities

import (
	"net/http"
	"time"
)

type Api struct {
	Id        uint
	Name      string
	Url       string
	Method    string
	Response  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (a *Api) ApiMapper(r *http.Request) {
	a.Name = r.FormValue("name")
	a.Url = r.FormValue("url")
	a.Method = r.FormValue("method")
	a.CreatedAt = time.Now()
	a.UpdatedAt = time.Now()
	a.Response = r.FormValue("response")
}
