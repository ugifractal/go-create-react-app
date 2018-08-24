package api

import (
	"encoding/json"
	"net/http"
)

func Handler() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")

		body, err := json.Marshal(map[string]interface{}{
			"data": "Hello, World",
		})

		if err != nil {
			res.WriteHeader(500)
			return
		}

		res.WriteHeader(200)
		res.Write(body)
	}
}
