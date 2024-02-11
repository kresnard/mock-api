package apimodel

import (
	"mock_api/entities"
	"mock_api/pkg"
)

func GetAll() ([]entities.Api, error) {
	rows, err := pkg.DB.Query(`SELECT * FROM apis ORDER BY updated_at desc`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var apis []entities.Api

	for rows.Next() {
		var api entities.Api
		err = rows.Scan(&api.Id, &api.Name, &api.Url, &api.Method, &api.Response, &api.CreatedAt, &api.UpdatedAt)
		if err != nil {
			return apis, err
		}

		apis = append(apis, api)
	}
	return apis, nil
}

func Create(api entities.Api) bool {
	var lastId int64
	err := pkg.DB.QueryRow(`
		INSERT INTO apis (name, url, method, response, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`,
		api.Name, api.Url, api.Method, api.Response, api.CreatedAt, api.UpdatedAt).Scan(&lastId)

	if err != nil {
		return false
	}

	return lastId > 0
}

func Detail(id int) entities.Api {
	row := pkg.DB.QueryRow(`SELECT * FROM apis where id = $1`, id)

	var api entities.Api
	if err := row.Scan(&api.Id, &api.Name, &api.Url, &api.Method, &api.Response, &api.CreatedAt, &api.UpdatedAt); err != nil {
		panic(err.Error())
	}
	return api
}

func Update(api entities.Api) bool {
	var updatedId int64
	err := pkg.DB.QueryRow(`
		UPDATE apis
		SET name = $1, url = $2, method = $3, response = $4, updated_at = $5
		WHERE id = $6
		RETURNING id`,
		api.Name, api.Url, api.Method, api.Response, api.UpdatedAt, api.Id).Scan(&updatedId)

	if err != nil {
		return false
	}

	return updatedId > 0
}

func Delete(api entities.Api) error {
	_, err := pkg.DB.Exec(`DELETE FROM apis where id = $1`, api.Id)

	return err
}
