package mockmodel

import "mock_api/pkg"

func Get(url string) (string, error) {
	row := pkg.DB.QueryRow(`SELECT response FROM apis WHERE url = $1`, url)

	response := ""
	if err := row.Scan(&response); err != nil {
		return response, err
	}
	return response, nil
}
